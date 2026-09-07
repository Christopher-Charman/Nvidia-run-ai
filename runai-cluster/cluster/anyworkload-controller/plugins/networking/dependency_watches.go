/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package networking

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/dsx-ai-factory/run-ai/runai-cluster/cluster/common/endpoints"
	dromosv1alpha1 "github.com/dsx-ai-factory/run-ai/runai-cluster/cluster/dromos/api/v1alpha1"
	openshiftv1 "github.com/openshift/api/route/v1"
	corev1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

const (
	workloadRunnerNetworkingDependencyIndex  = "networking.run.ai/dependency"
	endpointSlicePodUIDIndex                 = "networking.run.ai/endpoint-pod-uid"
	controlPlaneNetworkingWorkloadIDLabelKey = "run.ai/networking-workload-id"
	serviceDependencyPrefix                  = "service/"
	exposureDependencyPrefix                 = "exposure/"
	podLabelDependencyPrefix                 = "pod-label/"
)

func registerDependencyIndexes(ctx context.Context, indexer client.FieldIndexer) error {
	if err := indexer.IndexField(
		ctx, &dromosv1alpha1.WorkloadRunner{},
		workloadRunnerNetworkingDependencyIndex, workloadRunnerDependencies,
	); err != nil {
		return fmt.Errorf("register %s cache index: %w", workloadRunnerNetworkingDependencyIndex, err)
	}
	if err := indexer.IndexField(
		ctx, &discoveryv1.EndpointSlice{}, endpointSlicePodUIDIndex, endpointSlicePodUIDs,
	); err != nil {
		return fmt.Errorf("register %s cache index: %w", endpointSlicePodUIDIndex, err)
	}
	return nil
}

// workloadRunnerDependencies indexes the configured traffic graph. Typed keys
// keep Service names, exposure objects, and Pod-label anchors in one field
// without collisions, so the WorkloadRunner configuration is parsed once per
// cache update. The index only routes events; ownership is still authorized
// from live objects during reconciliation.
func workloadRunnerDependencies(obj client.Object) []string {
	runner, ok := obj.(*dromosv1alpha1.WorkloadRunner)
	if !ok || runner == nil {
		return nil
	}

	dependencies := sets.New[string]()
	addService := func(name string) {
		if name != "" {
			dependencies.Insert(serviceDependencyKey(name))
		}
	}
	addSelector := func(selector map[string]string) {
		if key := podLabelDependencyForSelector(selector); key != "" {
			dependencies.Insert(key)
		}
	}
	addAuthProxy := func(name string) {
		if name != "" {
			addService(name)
			addSelector(authProxyLabels(runner, name))
		}
	}
	addExposure := func(kind string, exposure client.Object) {
		name := exposure.GetName()
		if name == "" {
			name = runner.Name
		}
		dependencies.Insert(exposureDependencyKey(kind, name))
		backends, _ := exposureServiceBackends(exposure)
		for _, backend := range backends {
			addService(backend.key.Name)
		}
	}
	addExternalAccess := func(
		ingress *networkingv1.Ingress,
		route *openshiftv1.Route,
		httpRoute *gatewayv1.HTTPRoute,
	) {
		if ingress != nil {
			addExposure("Ingress", ingress)
		}
		if route != nil {
			addExposure("Route", route)
		}
		if httpRoute != nil {
			addExposure("HTTPRoute", httpRoute)
		}
	}

	if legacy := extractNetworkingConfig(runner); legacy != nil {
		for i := range legacy.Services {
			if service := &legacy.Services[i]; service.Create != nil {
				addSelector(service.Create.Spec.Selector)
			}
		}
		if hasService(legacy) {
			addService(getWorkloadServiceInfo(legacy, runner.Name).Name)
			if isAuthEnabled(legacy) {
				addAuthProxy(getAuthProxyName(runner))
			}
		}
		if legacy.ExternalAccess != nil {
			addExternalAccess(
				legacy.ExternalAccess.Ingress,
				legacy.ExternalAccess.Route,
				legacy.ExternalAccess.HTTPRoute,
			)
		}
	}

	annotationConfig, err := endpoints.Get(runner)
	if err == nil && annotationConfig != nil {
		for i := range annotationConfig.Endpoints {
			endpoint := &annotationConfig.Endpoints[i]
			for serviceIndex := range endpoint.Services {
				service := &endpoint.Services[serviceIndex]
				if service.Create != nil {
					addSelector(service.Create.Spec.Selector)
				}
				if service.Auth != nil {
					if name := endpointServiceName(service); name != "" {
						addSelector(authProxyLabels(runner, getEndpointAuthProxyName(name)))
					}
				}
			}
			if service := firstService(endpoint); service != nil {
				name := endpointServiceName(service)
				addService(name)
				if service.Auth != nil && name != "" {
					addService(getEndpointAuthProxyName(name))
				}
			}
			if endpoint.ExternalAccess != nil {
				addExternalAccess(
					endpoint.ExternalAccess.Ingress,
					endpoint.ExternalAccess.Route,
					endpoint.ExternalAccess.HTTPRoute,
				)
			}
		}
	}

	return sets.List(dependencies)
}

func podLabelDependencyForSelector(selector map[string]string) string {
	if len(selector) == 0 {
		return ""
	}
	if value, found := selector[controlPlaneNetworkingWorkloadIDLabelKey]; found {
		return podLabelDependencyKey(controlPlaneNetworkingWorkloadIDLabelKey, value)
	}

	keys := make([]string, 0, len(selector))
	for key := range selector {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return podLabelDependencyKey(keys[0], selector[keys[0]])
}

func podLabelDependencyKey(key, value string) string {
	return podLabelDependencyPrefix + key + "=" + value
}

func serviceDependencyKey(name string) string {
	return serviceDependencyPrefix + name
}

func workloadRunnerHasExposureDependency(obj client.Object) bool {
	for _, dependency := range workloadRunnerDependencies(obj) {
		if strings.HasPrefix(dependency, exposureDependencyPrefix) {
			return true
		}
	}
	return false
}

func endpointSlicePodUIDs(obj client.Object) []string {
	slice, ok := obj.(*discoveryv1.EndpointSlice)
	if !ok || slice == nil {
		return nil
	}

	uids := sets.New[string]()
	for i := range slice.Endpoints {
		target := slice.Endpoints[i].TargetRef
		if target == nil || target.Kind != "Pod" || target.UID == "" {
			continue
		}
		uids.Insert(string(target.UID))
	}
	return sets.List(uids)
}

func (p *Plugin) mapServiceDependency(ctx context.Context, obj client.Object) []reconcile.Request {
	return p.mapNetworkingDependency(ctx, obj, serviceDependencyKey(obj.GetName()))
}

func (p *Plugin) mapExposureDependency(ctx context.Context, obj client.Object) []reconcile.Request {
	key := exposureDependencyForObject(obj)
	if key == "" {
		return p.mapToWorkloadRunner(ctx, obj)
	}
	return p.mapNetworkingDependency(ctx, obj, key)
}

func (p *Plugin) mapNetworkingDependency(
	ctx context.Context,
	obj client.Object,
	value string,
) []reconcile.Request {
	return mergeRequests(
		p.mapToWorkloadRunner(ctx, obj),
		p.workloadRunnersForDependency(ctx, obj.GetNamespace(), value),
	)
}

func (p *Plugin) mapEndpointSliceDependency(ctx context.Context, obj client.Object) []reconcile.Request {
	slice, ok := obj.(*discoveryv1.EndpointSlice)
	if !ok {
		log.FromContext(ctx).Error(
			fmt.Errorf("expected *discoveryv1.EndpointSlice, got %T", obj),
			"map EndpointSlice networking dependency",
		)
		return nil
	}
	if slice == nil {
		return nil
	}
	serviceName, found := slice.Labels[discoveryv1.LabelServiceName]
	if !found || serviceName == "" {
		return nil
	}

	requests := p.workloadRunnersForDependency(
		ctx, slice.Namespace, serviceDependencyKey(serviceName),
	)
	if p.reader == nil {
		return requests
	}
	service := &corev1.Service{}
	if err := p.reader.Get(ctx, types.NamespacedName{Namespace: slice.Namespace, Name: serviceName}, service); err == nil {
		requests = mergeRequests(requests, p.mapToWorkloadRunner(ctx, service))
	}
	return requests
}

func (p *Plugin) mapPodDependency(ctx context.Context, obj client.Object) []reconcile.Request {
	pod, ok := obj.(*corev1.Pod)
	if !ok || pod == nil || p.reader == nil {
		return nil
	}

	var requests []reconcile.Request
	for key, value := range pod.Labels {
		requests = mergeRequests(requests, p.workloadRunnersForDependency(
			ctx,
			pod.Namespace,
			podLabelDependencyKey(key, value),
		))
	}

	if pod.UID != "" {
		slices := &discoveryv1.EndpointSliceList{}
		if err := p.reader.List(
			ctx,
			slices,
			client.InNamespace(pod.Namespace),
			client.MatchingFields{endpointSlicePodUIDIndex: string(pod.UID)},
		); err != nil {
			log.FromContext(ctx).Error(err, "map Pod networking dependency", "pod", client.ObjectKeyFromObject(pod))
			return requests
		}
		for i := range slices.Items {
			requests = mergeRequests(requests, p.mapEndpointSliceDependency(ctx, &slices.Items[i]))
		}
	}
	return requests
}

func (p *Plugin) workloadRunnersForDependency(
	ctx context.Context,
	namespace string,
	value string,
) []reconcile.Request {
	if p.reader == nil || namespace == "" || value == "" {
		return nil
	}

	runners := &dromosv1alpha1.WorkloadRunnerList{}
	if err := p.reader.List(
		ctx,
		runners,
		client.InNamespace(namespace),
		client.MatchingFields{workloadRunnerNetworkingDependencyIndex: value},
	); err != nil {
		log.FromContext(ctx).Error(err, "map networking dependency", "namespace", namespace, "value", value)
		return nil
	}

	requests := make([]reconcile.Request, 0, len(runners.Items))
	for i := range runners.Items {
		requests = append(requests, reconcile.Request{NamespacedName: types.NamespacedName{
			Namespace: runners.Items[i].Namespace,
			Name:      runners.Items[i].Name,
		}})
	}
	return mergeRequests(requests)
}

func exposureDependencyForObject(obj client.Object) string {
	switch object := obj.(type) {
	case *networkingv1.Ingress:
		return exposureDependencyKey("Ingress", object.Name)
	case *openshiftv1.Route:
		return exposureDependencyKey("Route", object.Name)
	case *gatewayv1.HTTPRoute:
		return exposureDependencyKey("HTTPRoute", object.Name)
	default:
		return ""
	}
}

func exposureDependencyKey(kind, name string) string {
	return exposureDependencyPrefix + kind + "/" + name
}

func mergeRequests(groups ...[]reconcile.Request) []reconcile.Request {
	unique := sets.New[types.NamespacedName]()
	for _, group := range groups {
		for _, request := range group {
			unique.Insert(request.NamespacedName)
		}
	}
	keys := unique.UnsortedList()
	sort.Slice(keys, func(i, j int) bool {
		if keys[i].Namespace == keys[j].Namespace {
			return keys[i].Name < keys[j].Name
		}
		return keys[i].Namespace < keys[j].Namespace
	})
	requests := make([]reconcile.Request, 0, len(keys))
	for _, key := range keys {
		requests = append(requests, reconcile.Request{NamespacedName: key})
	}
	return requests
}
