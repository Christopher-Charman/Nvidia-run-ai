/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package networking

import (
	"context"
	"errors"
	"reflect"
	"strings"
	"testing"

	dromosv1alpha1 "github.com/dsx-ai-factory/run-ai/runai-cluster/cluster/dromos/api/v1alpha1"
	"github.com/dsx-ai-factory/run-ai/runai-cluster/cluster/dromos/pkg/reconcile"
	"github.com/dsx-ai-factory/run-ai/runai-cluster/cluster/dromos/pkg/testutil"
	openshiftv1 "github.com/openshift/api/route/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	networkingv1 "k8s.io/api/networking/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

type patchCountingClient struct {
	client.Client
	patches int
}

func (c *patchCountingClient) Patch(
	ctx context.Context,
	obj client.Object,
	patch client.Patch,
	opts ...client.PatchOption,
) error {
	c.patches++
	return c.Client.Patch(ctx, obj, patch, opts...)
}

type createRaceClient struct {
	client.Client
	target     client.ObjectKey
	concurrent client.Object
	injected   bool
}

func (c *createRaceClient) Create(
	ctx context.Context,
	obj client.Object,
	opts ...client.CreateOption,
) error {
	if !c.injected && client.ObjectKeyFromObject(obj) == c.target {
		c.injected = true
		if err := c.Client.Create(ctx, c.concurrent.DeepCopyObject().(client.Object)); err != nil {
			return err
		}
	}
	return c.Client.Create(ctx, obj, opts...)
}

type replaceBeforePatchClient struct {
	client.Client
	target                 client.ObjectKey
	replacement            client.Object
	injected               bool
	patchedUID             types.UID
	patchedResourceVersion string
}

func (c *replaceBeforePatchClient) Patch(
	ctx context.Context,
	obj client.Object,
	patch client.Patch,
	opts ...client.PatchOption,
) error {
	if !c.injected && client.ObjectKeyFromObject(obj) == c.target {
		c.injected = true
		c.patchedUID = obj.GetUID()
		c.patchedResourceVersion = obj.GetResourceVersion()
		live := obj.DeepCopyObject().(client.Object)
		if err := c.Get(ctx, c.target, live); err != nil {
			return err
		}
		if err := c.Delete(ctx, live); err != nil {
			return err
		}
		replacement := c.replacement.DeepCopyObject().(client.Object)
		replacement.SetResourceVersion("")
		if err := c.Create(ctx, replacement); err != nil {
			return err
		}
	}
	return c.Client.Patch(ctx, obj, patch, opts...)
}

type deletePreconditionClient struct {
	client.Client
	uid types.UID
}

func (c *deletePreconditionClient) Delete(
	_ context.Context,
	obj client.Object,
	opts ...client.DeleteOption,
) error {
	deleteOptions := (&client.DeleteOptions{}).ApplyOptions(opts)
	if deleteOptions.Preconditions != nil && deleteOptions.Preconditions.UID != nil {
		c.uid = *deleteOptions.Preconditions.UID
	}
	return apierrors.NewConflict(
		schema.GroupResource{Resource: "services"}, obj.GetName(), errors.New("UID precondition mismatch"),
	)
}

func TestOwnershipEnforcementAllowsOwnedTrafficGraph(t *testing.T) {
	t.Parallel()

	runner := enforcementRunner()
	pod := enforcementPod(runner, "owned-pod", runner.UID)
	k8sClient := enforcementClient(t, pod)
	plugin := newOwnershipTestPlugin(k8sClient)
	state := reconcile.New()

	err := plugin.HandleAdditionalResources(
		context.Background(), runner, nil, nil, k8sClient, state,
	)
	if err != nil {
		t.Fatalf("HandleAdditionalResources() error = %v", err)
	}
	for _, object := range []client.Object{
		&corev1.Service{},
		&networkingv1.Ingress{},
	} {
		if err := k8sClient.Get(
			context.Background(),
			client.ObjectKey{Namespace: runner.Namespace, Name: enforcementServiceName(object)},
			object,
		); err != nil {
			t.Fatalf("read created %T: %v", object, err)
		}
	}
	if len(state.PluginIssues(plugin)) != 0 {
		t.Fatalf("unexpected plugin issues: %v", state.PluginIssues(plugin))
	}
}

func TestOwnershipEnforcementRejectsCreatedServiceSelectingForeignPod(t *testing.T) {
	t.Parallel()

	runner := enforcementRunner()
	foreignPod := enforcementPod(runner, "foreign-pod", "foreign-runner-uid")
	k8sClient := enforcementClient(t, foreignPod)
	plugin := newOwnershipTestPlugin(k8sClient)
	state := reconcile.New()

	err := plugin.HandleAdditionalResources(
		context.Background(), runner, nil, nil, k8sClient, state,
	)
	if err != nil {
		t.Fatalf("HandleAdditionalResources() error = %v, want settled ownership denial", err)
	}
	assertNotFound(t, k8sClient, runner.Namespace, "workload-service", &corev1.Service{})
	assertNotFound(t, k8sClient, runner.Namespace, "workload-ingress", &networkingv1.Ingress{})
	issues := state.PluginIssues(plugin)
	if len(issues) != 1 || !issues[0].IsCritical {
		t.Fatalf("plugin issues = %v, want one critical issue", issues)
	}
}

func TestOwnershipEnforcementDoesNotMutateReferencedServiceAndRemovesExposure(t *testing.T) {
	t.Parallel()

	runner := enforcementRunner()
	runner.Spec.Instructions.Networking.Services[0] = dromosv1alpha1.NetworkingServiceItem{
		Ref: &dromosv1alpha1.NetworkingServiceRef{Name: "referenced-service", Port: 8080},
	}
	referencedService := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: runner.Namespace,
			Name:      "referenced-service",
			UID:       "referenced-service-uid",
			Labels:    map[string]string{"preserve": "true"},
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{"app": "traffic-target"},
			Ports:    []corev1.ServicePort{{Port: 8080}},
		},
	}
	originalService := referencedService.DeepCopy()
	foreignPod := enforcementPod(runner, "foreign-pod", "foreign-runner-uid")
	existingIngress := newTestIngress("workload-ingress", "workload.example.com", "/")
	existingIngress.Namespace = runner.Namespace
	existingIngress.UID = "ingress-uid"
	existingIngress.OwnerReferences = []metav1.OwnerReference{newOwnershipRunnerRef(runner)}
	k8sClient := enforcementClient(t, referencedService, foreignPod, existingIngress)
	plugin := newOwnershipTestPlugin(k8sClient)
	state := reconcile.New()

	err := plugin.HandleAdditionalResources(
		context.Background(), runner, nil, nil, k8sClient, state,
	)
	if err != nil {
		t.Fatalf("HandleAdditionalResources() error = %v, want settled ownership denial", err)
	}
	liveService := &corev1.Service{}
	if err := k8sClient.Get(
		context.Background(), client.ObjectKeyFromObject(referencedService), liveService,
	); err != nil {
		t.Fatalf("referenced Service was removed: %v", err)
	}
	if !reflect.DeepEqual(liveService.Spec, originalService.Spec) ||
		!reflect.DeepEqual(liveService.Labels, originalService.Labels) {
		t.Fatalf("referenced Service was mutated: got %#v, want %#v", liveService, originalService)
	}
	assertNotFound(t, k8sClient, runner.Namespace, "workload-ingress", &networkingv1.Ingress{})
	issues := state.PluginIssues(plugin)
	if len(issues) != 1 || !issues[0].IsCritical {
		t.Fatalf("plugin issues = %v, want one critical issue", issues)
	}
}

func TestReconcileCreatedServicePreservesExistingServiceOnTransientValidationFailure(t *testing.T) {
	t.Parallel()

	runner := enforcementRunner()
	desired := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Namespace: runner.Namespace, Name: "workload-service"},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{"app": "traffic-target"},
			Ports:    []corev1.ServicePort{{Port: 8080}},
		},
	}
	existing := desired.DeepCopy()
	existing.UID = "workload-service-uid"
	existing.OwnerReferences = []metav1.OwnerReference{newOwnershipRunnerRef(runner)}
	k8sClient := enforcementClient(t, existing)
	readErr := errors.New("API server unavailable")
	plugin := newPlugin(k8sClient, uidDefaultingReader{Reader: selectiveOwnershipReader{
		Reader:     k8sClient,
		podListErr: readErr,
	}})

	err := plugin.reconcileCreatedService(
		context.Background(), runner, runner, desired, k8sClient, "workload service",
	)
	if !errors.Is(err, readErr) || !errors.Is(err, errOwnershipRetryNeeded) {
		t.Fatalf("reconcileCreatedService() error = %v, want retryable read failure", err)
	}
	if !errors.Is(err, errOwnershipOperational) {
		t.Fatalf("reconcileCreatedService() error = %v, want operational failure", err)
	}
	if err := k8sClient.Get(
		context.Background(), client.ObjectKeyFromObject(existing), &corev1.Service{},
	); err != nil {
		t.Fatalf("existing Service was removed after transient validation failure: %v", err)
	}
}

func TestReconcileCreatedServiceWaitsForOrphanEndpointSlicesWithoutApplying(t *testing.T) {
	t.Parallel()

	runner := enforcementRunner()
	orphan := &discoveryv1.EndpointSlice{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: runner.Namespace,
			Name:      "orphan-slice",
			Labels:    map[string]string{discoveryv1.LabelServiceName: "workload-service"},
		},
		AddressType: discoveryv1.AddressTypeIPv4,
	}
	k8sClient := enforcementClient(t, orphan)
	countingClient := &patchCountingClient{Client: k8sClient}
	plugin := newOwnershipTestPlugin(k8sClient)
	desired := plugin.buildWorkloadService(
		runner,
		runner.Spec.Instructions.Networking.Services[0].Create,
	)

	err := plugin.reconcileCreatedService(
		context.Background(), runner, runner, desired, countingClient, "workload service",
	)
	if !errors.Is(err, errOwnershipViolation) ||
		!strings.Contains(err.Error(), "still reference absent Service") {
		t.Fatalf("reconcileCreatedService() error = %v, want settled orphan EndpointSlice denial", err)
	}
	if countingClient.patches != 0 {
		t.Fatalf("Service apply calls = %d, want 0 while an orphan EndpointSlice remains", countingClient.patches)
	}
	assertNotFound(t, k8sClient, runner.Namespace, desired.Name, &corev1.Service{})
}

func TestSetOwnerAndApplyRevalidatesConcurrentCreation(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		ownerRefs     []metav1.OwnerReference
		wantViolation bool
	}{
		{
			name:          "foreign object is denied",
			wantViolation: true,
		},
		{
			name:      "owned object is retried",
			ownerRefs: []metav1.OwnerReference{newOwnershipRunnerRef(enforcementRunner())},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			runner := enforcementRunner()
			desired := &corev1.Service{
				ObjectMeta: metav1.ObjectMeta{Namespace: runner.Namespace, Name: "workload-service"},
				Spec: corev1.ServiceSpec{
					Selector: map[string]string{"app": "desired"},
					Ports:    []corev1.ServicePort{{Port: 8080}},
				},
			}
			concurrent := &corev1.Service{
				ObjectMeta: metav1.ObjectMeta{
					Namespace:       runner.Namespace,
					Name:            desired.Name,
					UID:             "concurrent-service-uid",
					Labels:          map[string]string{"preserve": "true"},
					OwnerReferences: tt.ownerRefs,
				},
				Spec: corev1.ServiceSpec{
					Selector: map[string]string{"app": "concurrent"},
					Ports:    []corev1.ServicePort{{Port: 9090}},
				},
			}
			baseClient := fake.NewClientBuilder().WithScheme(newTestScheme()).Build()
			racingClient := &createRaceClient{
				Client: baseClient, target: client.ObjectKeyFromObject(desired), concurrent: concurrent,
			}
			plugin := newOwnershipTestPlugin(racingClient)

			err := plugin.setOwnerAndApply(
				context.Background(), runner, racingClient, runner, desired, "workload service",
			)
			wantRetry := !tt.wantViolation
			if gotRetry := errors.Is(err, errOwnershipRetryNeeded); gotRetry != wantRetry {
				t.Fatalf(
					"setOwnerAndApply() error = %v, want retryable = %v",
					err, wantRetry,
				)
			}
			if errors.Is(err, errOwnershipViolation) != tt.wantViolation {
				t.Fatalf("setOwnerAndApply() error = %v, wantViolation = %v", err, tt.wantViolation)
			}
			if wantRetry && !errors.Is(err, errOwnershipEventDrivenWait) {
				t.Fatalf("setOwnerAndApply() error = %v, want event-driven retry", err)
			}

			live := &corev1.Service{}
			if err := baseClient.Get(context.Background(), client.ObjectKeyFromObject(desired), live); err != nil {
				t.Fatalf("read concurrent Service: %v", err)
			}
			if !reflect.DeepEqual(live.Spec, concurrent.Spec) ||
				!reflect.DeepEqual(live.Labels, concurrent.Labels) ||
				!reflect.DeepEqual(live.OwnerReferences, concurrent.OwnerReferences) {
				t.Fatalf("concurrent Service was mutated: got %#v, want %#v", live, concurrent)
			}
		})
	}
}

func TestSetOwnerAndApplyRejectsReplacementBeforeApply(t *testing.T) {
	t.Parallel()

	runner := enforcementRunner()
	existing := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:       runner.Namespace,
			Name:            "workload-service",
			UID:             "owned-service-uid",
			OwnerReferences: []metav1.OwnerReference{newOwnershipRunnerRef(runner)},
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{"app": "owned"},
			Ports:    []corev1.ServicePort{{Port: 8080}},
		},
	}
	desired := existing.DeepCopy()
	desired.TypeMeta = metav1.TypeMeta{APIVersion: "v1", Kind: "Service"}
	desired.UID = ""
	desired.ResourceVersion = ""
	desired.OwnerReferences = nil
	desired.Spec.Selector = map[string]string{"app": "desired"}
	replacement := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: runner.Namespace,
			Name:      existing.Name,
			UID:       "replacement-service-uid",
			Labels:    map[string]string{"preserve": "true"},
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{"app": "replacement"},
			Ports:    []corev1.ServicePort{{Port: 9090}},
		},
	}
	baseClient := fake.NewClientBuilder().WithScheme(newTestScheme()).WithObjects(existing).Build()
	racingClient := &replaceBeforePatchClient{
		Client: baseClient, target: client.ObjectKeyFromObject(desired), replacement: replacement,
	}
	plugin := newOwnershipTestPlugin(racingClient)

	err := plugin.setOwnerAndApply(
		context.Background(), runner, racingClient, runner, desired, "workload service",
	)
	if !apierrors.IsConflict(err) || !errors.Is(err, errOwnershipRetryNeeded) {
		t.Fatalf("setOwnerAndApply() error = %v, want retryable resource-version conflict", err)
	}
	if !errors.Is(err, errOwnershipEventDrivenWait) {
		t.Fatalf("setOwnerAndApply() error = %v, want event-driven retry", err)
	}
	if errors.Is(err, errOwnershipViolation) {
		t.Fatalf("setOwnerAndApply() error = %v, want transient conflict, not ownership violation", err)
	}
	if racingClient.patchedUID != existing.UID || racingClient.patchedResourceVersion == "" {
		t.Fatalf(
			"apply identity = UID %q resourceVersion %q, want UID %q and observed resourceVersion",
			racingClient.patchedUID, racingClient.patchedResourceVersion, existing.UID,
		)
	}

	live := &corev1.Service{}
	if err := baseClient.Get(context.Background(), client.ObjectKeyFromObject(desired), live); err != nil {
		t.Fatalf("read replacement Service: %v", err)
	}
	if !reflect.DeepEqual(live.Spec, replacement.Spec) ||
		!reflect.DeepEqual(live.Labels, replacement.Labels) || len(live.OwnerReferences) != 0 {
		t.Fatalf("replacement Service was mutated: got %#v, want %#v", live, replacement)
	}
}

func TestDeleteIfOwnedUsesValidatedUIDPrecondition(t *testing.T) {
	t.Parallel()

	runner := enforcementRunner()
	existing := &corev1.Service{ObjectMeta: metav1.ObjectMeta{
		Namespace:       runner.Namespace,
		Name:            "workload-service",
		UID:             "owned-service-uid",
		OwnerReferences: []metav1.OwnerReference{newOwnershipRunnerRef(runner)},
	}}
	baseClient := enforcementClient(t, existing)
	deletingClient := &deletePreconditionClient{Client: baseClient}
	plugin := newOwnershipTestPlugin(baseClient)

	err := plugin.deleteIfOwned(
		context.Background(), runner, existing.DeepCopy(), deletingClient, "workload service",
	)
	if !apierrors.IsConflict(err) || !errors.Is(err, errOwnershipRetryNeeded) {
		t.Fatalf("deleteIfOwned() error = %v, want retryable UID precondition conflict", err)
	}
	if !errors.Is(err, errOwnershipEventDrivenWait) {
		t.Fatalf("deleteIfOwned() error = %v, want event-driven retry", err)
	}
	if deletingClient.uid != existing.UID {
		t.Fatalf("delete UID precondition = %q, want validated UID %q", deletingClient.uid, existing.UID)
	}
	if err := baseClient.Get(
		context.Background(), client.ObjectKeyFromObject(existing), &corev1.Service{},
	); err != nil {
		t.Fatalf("Service was removed after UID precondition conflict: %v", err)
	}
}

func TestReconcileExternalAccessPreservesExistingExposureOnApplyFailure(t *testing.T) {
	t.Parallel()

	runner := enforcementRunner()
	pod := enforcementPod(runner, "owned-pod", runner.UID)
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: runner.Namespace,
			Name:      "workload-service",
			UID:       "workload-service-uid",
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{"app": "traffic-target"},
			Ports:    []corev1.ServicePort{{Port: 8080}},
		},
	}
	desired := NewPlugin(nil).buildIngress(
		runner, runner.Spec.Instructions.Networking.ExternalAccess.Ingress, service.Name, 8080,
	)
	existing := desired.DeepCopy()
	existing.UID = "workload-ingress-uid"
	existing.OwnerReferences = []metav1.OwnerReference{newOwnershipRunnerRef(runner)}
	k8sClient := enforcementClient(t, pod, service, existing)
	plugin := newOwnershipTestPlugin(k8sClient)
	failingClient := &writeFailingClient{Client: k8sClient, name: desired.Name}

	err := plugin.reconcileExternalAccess(
		context.Background(), runner, runner,
		[]managedResource{{object: desired, description: "ingress"}}, failingClient,
	)
	if err == nil || !strings.Contains(err.Error(), "injected write failure") {
		t.Fatalf("reconcileExternalAccess() error = %v, want injected write failure", err)
	}
	if err := k8sClient.Get(
		context.Background(), client.ObjectKeyFromObject(existing), &networkingv1.Ingress{},
	); err != nil {
		t.Fatalf("existing Ingress was removed after apply failure: %v", err)
	}
}

func TestCleanupAfterOwnershipErrorPreservesAuthProxyOnRetryableFailure(t *testing.T) {
	t.Parallel()

	runner := enforcementRunner()
	ownerRef := []metav1.OwnerReference{newOwnershipRunnerRef(runner)}
	existingIngress := &networkingv1.Ingress{ObjectMeta: metav1.ObjectMeta{
		Namespace: runner.Namespace, Name: "workload-ingress", UID: "workload-ingress-uid",
		OwnerReferences: ownerRef,
	}}
	authService := &corev1.Service{ObjectMeta: metav1.ObjectMeta{
		Namespace: runner.Namespace, Name: "workload-auth", UID: "auth-service-uid",
		OwnerReferences: ownerRef,
	}}
	authDeployment := &appsv1.Deployment{ObjectMeta: metav1.ObjectMeta{
		Namespace: runner.Namespace, Name: "workload-auth", UID: "auth-deployment-uid",
		OwnerReferences: ownerRef,
	}}
	k8sClient := enforcementClient(t, existingIngress, authService, authDeployment)
	plugin := newOwnershipTestPlugin(k8sClient)
	readErr := errors.New("API server unavailable")
	enforcementErr := ownershipEnforcementError("validate traffic graph", readErr)

	err := plugin.cleanupAfterOwnershipError(
		context.Background(), runner, enforcementErr,
		[]managedResource{{object: existingIngress, description: "ingress"}},
		"workload-auth", true, k8sClient,
	)
	if !errors.Is(err, readErr) || !errors.Is(err, errOwnershipRetryNeeded) {
		t.Fatalf("cleanupAfterOwnershipError() error = %v, want retryable read failure", err)
	}
	if !errors.Is(err, errOwnershipOperational) {
		t.Fatalf("cleanupAfterOwnershipError() error = %v, want operational failure", err)
	}
	for _, object := range []client.Object{
		&networkingv1.Ingress{},
		&corev1.Service{},
		&appsv1.Deployment{},
	} {
		name := "workload-auth"
		if _, ok := object.(*networkingv1.Ingress); ok {
			name = "workload-ingress"
		}
		if err := k8sClient.Get(
			context.Background(), client.ObjectKey{Namespace: runner.Namespace, Name: name}, object,
		); err != nil {
			t.Fatalf("existing %T was removed after retryable failure: %v", object, err)
		}
	}
}

func TestOwnershipEnforcementRetryDecision(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name            string
		err             error
		wantErr         bool
		wantCritical    bool
		wantEventDriven bool
		wantOperational bool
	}{
		{
			name:         "proven violation settles",
			err:          ownershipEnforcementError("validate traffic graph", ownershipViolationf("foreign Pod")),
			wantCritical: true,
		},
		{
			name: "missing dependency waits for watched event",
			err: ownershipEnforcementError("validate traffic graph", apierrors.NewNotFound(
				schema.GroupResource{Resource: "services"}, "missing-service",
			)),
			wantEventDriven: true,
		},
		{
			name: "resource-version conflict waits for watched event",
			err: ownershipEnforcementError("apply Service", apierrors.NewConflict(
				schema.GroupResource{Resource: "services"}, "workload-service", errors.New("changed"),
			)),
			wantEventDriven: true,
		},
		{
			name:            "transient read failure uses controller retry",
			err:             ownershipEnforcementError("validate traffic graph", errors.New("API server unavailable")),
			wantErr:         true,
			wantOperational: true,
		},
		{
			name: "mixed watched and operational failures use controller retry",
			err: errors.Join(
				ownershipEnforcementError("validate traffic graph", apierrors.NewNotFound(
					schema.GroupResource{Resource: "services"}, "missing-service",
				)),
				ownershipEnforcementError("validate traffic graph", errors.New("API server unavailable")),
			),
			wantErr:         true,
			wantEventDriven: true,
			wantOperational: true,
		},
		{
			name: "mixed causes in one enforcement error use controller retry",
			err: ownershipEnforcementError("validate traffic graph", errors.Join(
				apierrors.NewNotFound(
					schema.GroupResource{Resource: "services"}, "missing-service",
				),
				errors.New("API server unavailable"),
			)),
			wantErr:         true,
			wantOperational: true,
		},
		{
			name: "violation with cleanup failure uses controller retry",
			err: errors.Join(
				ownershipEnforcementError("validate traffic graph", ownershipViolationf("foreign Pod")),
				ownershipEnforcementError("remove unsafe exposure", errors.New("delete failed")),
			),
			wantErr:         true,
			wantCritical:    true,
			wantOperational: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			runner := enforcementRunner()
			plugin := NewPlugin(nil)
			state := reconcile.New()

			err := plugin.recordOwnershipEnforcementError(
				context.Background(), runner, state, tt.err, "ownership validation",
			)
			if (err != nil) != tt.wantErr {
				t.Fatalf("recordOwnershipEnforcementError() error = %v, wantErr = %v", err, tt.wantErr)
			}
			if got := errors.Is(tt.err, errOwnershipEventDrivenWait); got != tt.wantEventDriven {
				t.Fatalf("event-driven classification = %v, want %v", got, tt.wantEventDriven)
			}
			if got := errors.Is(tt.err, errOwnershipOperational); got != tt.wantOperational {
				t.Fatalf("operational classification = %v, want %v", got, tt.wantOperational)
			}
			issues := state.PluginIssues(plugin)
			if len(issues) != 1 || issues[0].IsCritical != tt.wantCritical {
				t.Fatalf("plugin issues = %v, want one issue with critical = %v", issues, tt.wantCritical)
			}
		})
	}
}

func TestOwnershipEnforcementRejectsSameNameForeignObjects(t *testing.T) {
	t.Parallel()

	runner := enforcementRunner()
	objects := []client.Object{
		&corev1.Service{},
		&appsv1.Deployment{},
		&networkingv1.Ingress{},
		&openshiftv1.Route{},
		&gatewayv1.HTTPRoute{},
	}
	for _, desired := range objects {
		desired := desired
		t.Run(reflect.TypeOf(desired).String(), func(t *testing.T) {
			t.Parallel()
			desired.SetNamespace(runner.Namespace)
			desired.SetName("collision")
			existing := desired.DeepCopyObject().(client.Object)
			existing.SetUID(types.UID("foreign-object-uid"))
			existing.SetOwnerReferences([]metav1.OwnerReference{newOwnershipRunnerRef(runner)})
			existing.GetOwnerReferences()[0].UID = "foreign-runner-uid"
			k8sClient := enforcementClient(t, existing)
			plugin := newOwnershipTestPlugin(k8sClient)

			err := plugin.ensureOwnedOrAbsent(
				context.Background(), runner, desired, "collision test object",
			)
			if !errors.Is(err, errOwnershipViolation) {
				t.Fatalf("ensureOwnedOrAbsent() error = %v, want ownership violation", err)
			}
			if err := plugin.deleteIfOwned(
				context.Background(), runner, desired, k8sClient, "collision test object",
			); err != nil {
				t.Fatalf("deleteIfOwned() error = %v, want foreign object preserved", err)
			}
			live := desired.DeepCopyObject().(client.Object)
			if err := k8sClient.Get(context.Background(), client.ObjectKeyFromObject(desired), live); err != nil {
				t.Fatalf("foreign object was removed: %v", err)
			}
			if live.GetUID() != "foreign-object-uid" {
				t.Fatalf("foreign object changed: UID = %q", live.GetUID())
			}
		})
	}
}

func TestOwnershipEnforcementStatusOutcomes(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		cause      error
		critical   bool
		wantPhase  dromosv1alpha1.PluginPhase
		wantStatus metav1.ConditionStatus
		wantReason string
	}{
		{
			name: "missing dependency waits",
			cause: apierrors.NewNotFound(
				schema.GroupResource{Resource: "services"}, "missing-service",
			),
			wantPhase:  dromosv1alpha1.PluginPhaseInitializing,
			wantStatus: metav1.ConditionFalse,
			wantReason: NetworkingOwnershipReasonDependencyMissing,
		},
		{
			name:       "transient read failure waits",
			cause:      errors.New("API server unavailable"),
			wantPhase:  dromosv1alpha1.PluginPhaseInitializing,
			wantStatus: metav1.ConditionUnknown,
			wantReason: NetworkingOwnershipReasonReadFailed,
		},
		{
			name:       "ownership violation fails networking",
			cause:      ownershipViolationf("foreign traffic destination"),
			critical:   true,
			wantPhase:  dromosv1alpha1.PluginPhaseFailed,
			wantStatus: metav1.ConditionFalse,
			wantReason: NetworkingOwnershipReasonDenied,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			runner := enforcementRunner()
			status := newPluginStatus()
			runner.Status.Plugins = map[string]*dromosv1alpha1.PluginStatus{PluginName: status}
			plugin := NewPlugin(nil)
			state := reconcile.New()
			_ = state.AddPluginIssue(
				context.Background(),
				plugin,
				runner,
				ownershipEnforcementError("validate traffic graph", tt.cause),
				tt.critical,
				"ownership validation",
				"Network",
			)

			err := plugin.Status(
				context.Background(),
				runner,
				fake.NewClientBuilder().WithScheme(newTestScheme()).Build(),
				status,
				state,
			)
			if err != nil {
				t.Fatalf("Status() error = %v", err)
			}
			if status.Phase != tt.wantPhase {
				t.Fatalf("phase = %q, want %q", status.Phase, tt.wantPhase)
			}
			condition := findCondition(status.Conditions, NetworkingOwnershipReadyConditionName)
			if condition == nil || condition.Status != tt.wantStatus || condition.Reason != tt.wantReason {
				t.Fatalf("ownership condition = %#v, want status %q reason %q", condition, tt.wantStatus, tt.wantReason)
			}
		})
	}
}

func TestOwnershipEnforcementStatusRecovers(t *testing.T) {
	t.Parallel()

	runner := enforcementRunner()
	status := newPluginStatus()
	status.Conditions = append(status.Conditions, ownershipCondition(
		ownershipStatusDenied, "previous ownership violation",
	))
	runner.Status.Plugins = map[string]*dromosv1alpha1.PluginStatus{PluginName: status}
	plugin := NewPlugin(nil)

	err := plugin.Status(
		context.Background(),
		runner,
		fake.NewClientBuilder().WithScheme(newTestScheme()).Build(),
		status,
		reconcile.New(),
	)
	if err != nil {
		t.Fatalf("Status() error = %v", err)
	}
	condition := findCondition(status.Conditions, NetworkingOwnershipReadyConditionName)
	if condition == nil || condition.Status != metav1.ConditionTrue ||
		condition.Reason != NetworkingOwnershipReasonAuthorized {
		t.Fatalf("recovered ownership condition = %#v", condition)
	}
}

func TestOwnershipDenialClearsPreviouslyReadyExternalAccessStatus(t *testing.T) {
	t.Parallel()

	runner := enforcementRunner()
	status := newPluginStatus()
	status.Conditions = append(status.Conditions, metav1.Condition{
		Type:    ExternalAccessReadyConditionName,
		Status:  metav1.ConditionTrue,
		Reason:  "ExternalAccessReady",
		Message: "foreign ingress exists, url: https://foreign.example.com/",
	})
	status.AdditionalInfo = &runtime.RawExtension{Raw: []byte(`{"endpoints":[{"urls":["https://foreign.example.com/"]}]}`)}
	runner.Status.Plugins = map[string]*dromosv1alpha1.PluginStatus{PluginName: status}
	plugin := NewPlugin(nil)
	state := reconcile.New()
	_ = state.AddPluginIssue(
		context.Background(),
		plugin,
		runner,
		ownershipEnforcementError("validate traffic graph", ownershipViolationf("foreign exposure")),
		true,
		"ownership validation",
		"Network",
	)

	err := plugin.Status(
		context.Background(),
		runner,
		fake.NewClientBuilder().WithScheme(newTestScheme()).Build(),
		status,
		state,
	)
	if err != nil {
		t.Fatalf("Status() error = %v", err)
	}
	if status.Phase != dromosv1alpha1.PluginPhaseFailed {
		t.Fatalf("phase = %q, want %q", status.Phase, dromosv1alpha1.PluginPhaseFailed)
	}
	condition := findCondition(status.Conditions, ExternalAccessReadyConditionName)
	if condition == nil || condition.Status != metav1.ConditionFalse ||
		condition.Reason != NetworkingOwnershipReasonDenied {
		t.Fatalf("external access condition = %#v, want status False reason %q", condition, NetworkingOwnershipReasonDenied)
	}
	if status.AdditionalInfo != nil {
		t.Fatalf("additional info = %#v, want nil", status.AdditionalInfo)
	}
}

func enforcementRunner() *dromosv1alpha1.WorkloadRunner {
	runner := newTestRunner(
		"workload-runner",
		testNamespace,
		&dromosv1alpha1.NetworkingConfig{
			Services: []dromosv1alpha1.NetworkingServiceItem{{
				Create: &corev1.Service{
					ObjectMeta: metav1.ObjectMeta{Name: "workload-service"},
					Spec: corev1.ServiceSpec{
						Selector: map[string]string{"app": "traffic-target"},
						Ports:    []corev1.ServicePort{{Port: 8080}},
					},
				},
			}},
			ExternalAccess: &dromosv1alpha1.NetworkingExternalAccessConfig{
				Ingress: newTestIngress("workload-ingress", "workload.example.com", "/"),
			},
		},
	)
	runner.UID = "workload-runner-uid"
	return runner
}

func enforcementPod(
	runner *dromosv1alpha1.WorkloadRunner,
	name string,
	ownerUID types.UID,
) *corev1.Pod {
	ownerRef := newOwnershipRunnerRef(runner)
	ownerRef.UID = ownerUID
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:       runner.Namespace,
			Name:            name,
			UID:             types.UID(name + "-uid"),
			Labels:          map[string]string{"app": "traffic-target"},
			OwnerReferences: []metav1.OwnerReference{ownerRef},
		},
	}
}

func enforcementClient(t *testing.T, objects ...client.Object) client.Client {
	t.Helper()
	return testutil.WithSSASupport(
		fake.NewClientBuilder().WithScheme(newTestScheme()).WithObjects(objects...).Build(),
	)
}

func enforcementServiceName(object client.Object) string {
	switch object.(type) {
	case *corev1.Service:
		return "workload-service"
	default:
		return "workload-ingress"
	}
}

func assertNotFound(
	t *testing.T,
	k8sClient client.Client,
	namespace, name string,
	object client.Object,
) {
	t.Helper()
	err := k8sClient.Get(context.Background(), client.ObjectKey{Namespace: namespace, Name: name}, object)
	if !apierrors.IsNotFound(err) {
		t.Fatalf("Get(%T %s/%s) error = %v, want NotFound", object, namespace, name, err)
	}
}
