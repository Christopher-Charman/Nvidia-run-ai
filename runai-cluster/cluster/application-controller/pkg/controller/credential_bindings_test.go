/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package controller

import (
	"context"

	"github.com/dsx-ai-factory/run-ai/runai-cluster/cluster/sdk/apis/run/v1alpha1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var _ = Describe("credential Secret event mapping", func() {
	It("enqueues every consumer only in the Secret namespace, including unlabeled Secrets", func() {
		scheme := runtime.NewScheme()
		Expect(v1alpha1.AddToScheme(scheme)).To(Succeed())
		cli := fake.NewClientBuilder().WithScheme(scheme).
			WithIndex(&v1alpha1.Application{}, credentialBindingsIndex, credentialSecretNames).Build()
		for _, entry := range []struct{ name, namespace, secret string }{
			{"first", "project", "shared"}, {"second", "project", "shared"},
			{"other-project", "other", "shared"}, {"unrelated", "project", "other-secret"},
		} {
			app := &v1alpha1.Application{ObjectMeta: metav1.ObjectMeta{
				Name: entry.name, Namespace: entry.namespace,
			},
				Spec: v1alpha1.ApplicationSpec{Helm: &v1alpha1.HelmSpec{
					CredentialBindings: []v1alpha1.HelmCredentialBinding{
						{Variable: "one", SecretName: entry.secret},
						{Variable: "two", SecretName: entry.secret},
					},
				}}}
			Expect(cli.Create(context.Background(), app)).To(Succeed())
		}
		c := NewApplicationController(cli)
		secret := &corev1.Secret{ObjectMeta: metav1.ObjectMeta{Name: "shared", Namespace: "project"}}
		requests := c.mapCredentialSecretToApplications(context.Background(), secret)
		Expect(requests).To(ConsistOf(
			reconcile.Request{NamespacedName: types.NamespacedName{Namespace: "project", Name: "first"}},
			reconcile.Request{NamespacedName: types.NamespacedName{Namespace: "project", Name: "second"}},
		))
	})
})
