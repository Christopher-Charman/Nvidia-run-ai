/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package networking

import (
	"context"
	"reflect"
	"testing"

	"github.com/dsx-ai-factory/run-ai/runai-cluster/cluster/common/endpoints"
	dromosv1alpha1 "github.com/dsx-ai-factory/run-ai/runai-cluster/cluster/dromos/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

func TestWorkloadRunnerDependencyIndexesLegacyInput(t *testing.T) {
	t.Parallel()

	runner := newTestRunner("runner", "test-namespace", &dromosv1alpha1.NetworkingConfig{
		Services: []dromosv1alpha1.NetworkingServiceItem{{
			Ref:  &dromosv1alpha1.NetworkingServiceRef{Name: "referenced-service", Port: 8080},
			Auth: &dromosv1alpha1.NetworkingAuthConfig{},
		}},
		ExternalAccess: &dromosv1alpha1.NetworkingExternalAccessConfig{
			Ingress: &networkingv1.Ingress{
				ObjectMeta: metav1.ObjectMeta{Name: "public-ingress"},
				Spec: networkingv1.IngressSpec{
					DefaultBackend: &networkingv1.IngressBackend{Service: &networkingv1.IngressServiceBackend{
						Name: "secondary-service",
					}},
					Rules: []networkingv1.IngressRule{{
						IngressRuleValue: networkingv1.IngressRuleValue{HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{{Backend: networkingv1.IngressBackend{
								Service: &networkingv1.IngressServiceBackend{Name: "referenced-service"},
							}}},
						}},
					}},
				},
			},
		},
	})

	if got, want := workloadRunnerDependencies(runner), []string{
		"exposure/Ingress/public-ingress",
		"pod-label/app=runner-auth-proxy",
		"service/referenced-service",
		"service/runner-auth-proxy",
		"service/secondary-service",
	}; !reflect.DeepEqual(got, want) {
		t.Fatalf("networking dependencies = %v, want %v", got, want)
	}
}

func TestWorkloadRunnerDependencyIndexesEndpointInput(t *testing.T) {
	t.Parallel()

	runner := newTestRunner("runner", "test-namespace", nil)
	port := gatewayv1.PortNumber(8080)
	httpRoute := &gatewayv1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{Name: "endpoint-route"},
		Spec: gatewayv1.HTTPRouteSpec{Rules: []gatewayv1.HTTPRouteRule{{
			BackendRefs: []gatewayv1.HTTPBackendRef{{BackendRef: gatewayv1.BackendRef{
				BackendObjectReference: gatewayv1.BackendObjectReference{
					Name: "endpoint-service",
					Port: &port,
				},
			}}},
		}}},
	}
	config := &endpoints.Config{Endpoints: []endpoints.Endpoint{{
		Services: []endpoints.Service{{
			Create: &corev1.Service{ObjectMeta: metav1.ObjectMeta{Name: "endpoint-service"}},
			Auth:   &endpoints.AuthConfig{},
		}},
		ExternalAccess: &endpoints.ExternalAccess{HTTPRoute: httpRoute},
	}}}
	if err := endpoints.Set(runner, config); err != nil {
		t.Fatalf("set endpoint annotation: %v", err)
	}

	if got, want := workloadRunnerDependencies(runner), []string{
		"exposure/HTTPRoute/endpoint-route",
		"pod-label/app=endpoint-service-auth-proxy",
		"service/endpoint-service",
		"service/endpoint-service-auth-proxy",
	}; !reflect.DeepEqual(got, want) {
		t.Fatalf("networking dependencies = %v, want %v", got, want)
	}
}

func TestServiceAndExposureEventsMapOwnerlessDependencies(t *testing.T) {
	t.Parallel()

	runner := newTestRunner("runner", "test-namespace", &dromosv1alpha1.NetworkingConfig{
		Services: []dromosv1alpha1.NetworkingServiceItem{{
			Ref:  &dromosv1alpha1.NetworkingServiceRef{Name: "referenced-service", Port: 8080},
			Auth: &dromosv1alpha1.NetworkingAuthConfig{},
		}},
		ExternalAccess: &dromosv1alpha1.NetworkingExternalAccessConfig{
			Ingress: &networkingv1.Ingress{ObjectMeta: metav1.ObjectMeta{Name: "collision-ingress"}},
		},
	})
	reader := dependencyWatchClient(runner)
	plugin := NewPlugin(reader)
	want := requestsForRunner(runner)

	service := &corev1.Service{ObjectMeta: metav1.ObjectMeta{
		Namespace: runner.Namespace,
		Name:      "referenced-service",
	}}
	if got := plugin.mapServiceDependency(context.Background(), service); !reflect.DeepEqual(got, want) {
		t.Fatalf("Service requests = %v, want %v", got, want)
	}

	ingress := &networkingv1.Ingress{ObjectMeta: metav1.ObjectMeta{
		Namespace: runner.Namespace,
		Name:      "collision-ingress",
	}}
	if got := plugin.mapExposureDependency(context.Background(), ingress); !reflect.DeepEqual(got, want) {
		t.Fatalf("Ingress requests = %v, want %v", got, want)
	}
}

func TestEndpointSliceAndPodEventsMapServiceDependency(t *testing.T) {
	t.Parallel()

	runner := newTestRunner("runner", "test-namespace", &dromosv1alpha1.NetworkingConfig{
		Services: []dromosv1alpha1.NetworkingServiceItem{{
			Create: &corev1.Service{ObjectMeta: metav1.ObjectMeta{Name: "workload-service"}},
		}},
	})
	podUID := types.UID("pod-uid")
	slice := &discoveryv1.EndpointSlice{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: runner.Namespace,
			Name:      "workload-service-slice",
			Labels:    map[string]string{discoveryv1.LabelServiceName: "workload-service"},
		},
		Endpoints: []discoveryv1.Endpoint{{TargetRef: &corev1.ObjectReference{
			Kind:      "Pod",
			Namespace: runner.Namespace,
			Name:      "workload-pod",
			UID:       podUID,
		}}},
	}
	reader := dependencyWatchClient(runner, slice)
	plugin := NewPlugin(reader)
	want := requestsForRunner(runner)

	if got := plugin.mapEndpointSliceDependency(context.Background(), slice); !reflect.DeepEqual(got, want) {
		t.Fatalf("EndpointSlice requests = %v, want %v", got, want)
	}
	pod := &corev1.Pod{ObjectMeta: metav1.ObjectMeta{
		Namespace: runner.Namespace,
		Name:      "workload-pod",
		UID:       podUID,
	}}
	if got := plugin.mapPodDependency(context.Background(), pod); !reflect.DeepEqual(got, want) {
		t.Fatalf("Pod requests = %v, want %v", got, want)
	}
}

func TestPodEventsMapConfiguredServiceSelectorsWithoutEndpointSlices(t *testing.T) {
	t.Parallel()

	markerless := newTestRunner("markerless", "test-namespace", &dromosv1alpha1.NetworkingConfig{
		Services: []dromosv1alpha1.NetworkingServiceItem{{
			Create: &corev1.Service{Spec: corev1.ServiceSpec{
				Selector: map[string]string{"app": "markerless-pod", "component": "worker"},
			}},
		}},
	})
	direct := newTestRunner("direct", "test-namespace", nil)
	if err := endpoints.Set(direct, &endpoints.Config{Endpoints: []endpoints.Endpoint{{
		Services: []endpoints.Service{{Create: &corev1.Service{Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app":                                    "direct-pod",
				controlPlaneNetworkingWorkloadIDLabelKey: "direct-workload-id",
			},
		}}}},
	}}}); err != nil {
		t.Fatalf("set direct endpoint annotation: %v", err)
	}

	plugin := NewPlugin(dependencyWatchClient(markerless, direct))
	tests := []struct {
		name   string
		labels map[string]string
		want   []reconcile.Request
	}{
		{
			name: "markerless selector fallback",
			labels: map[string]string{
				"app":       "markerless-pod",
				"component": "worker",
			},
			want: requestsForRunner(markerless),
		},
		{
			name: "control-plane ownership selector",
			labels: map[string]string{
				"app":                                    "direct-pod",
				controlPlaneNetworkingWorkloadIDLabelKey: "direct-workload-id",
			},
			want: requestsForRunner(direct),
		},
		{
			name:   "partial direct selector does not match anchor",
			labels: map[string]string{"app": "direct-pod"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			pod := &corev1.Pod{ObjectMeta: metav1.ObjectMeta{
				Namespace: markerless.Namespace,
				Name:      "changed-pod",
				UID:       "changed-pod-uid",
				Labels:    tt.labels,
			}}
			got := plugin.mapPodDependency(context.Background(), pod)
			if len(tt.want) == 0 && len(got) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Pod requests = %v, want %v", got, tt.want)
			}
		})
	}
}

func requestsForRunner(runner *dromosv1alpha1.WorkloadRunner) []reconcile.Request {
	return []reconcile.Request{{NamespacedName: types.NamespacedName{
		Namespace: runner.Namespace,
		Name:      runner.Name,
	}}}
}

func dependencyWatchClient(objects ...client.Object) client.Client {
	return fake.NewClientBuilder().
		WithScheme(newTestScheme()).
		WithObjects(objects...).
		WithIndex(
			&dromosv1alpha1.WorkloadRunner{},
			workloadRunnerNetworkingDependencyIndex,
			workloadRunnerDependencies,
		).
		WithIndex(
			&discoveryv1.EndpointSlice{},
			endpointSlicePodUIDIndex,
			endpointSlicePodUIDs,
		).
		Build()
}
