/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package resources

import (
	"context"

	"github.com/dsx-ai-factory/run-ai/runai-cluster/cluster/go-operator/internal/run/utils"
	runv1_internal "github.com/dsx-ai-factory/run-ai/runai-cluster/cluster/sdk/apis/run/v1/runai-internal"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	osv1 "github.com/openshift/api/route/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

var _ = Describe("cluster-api routes", func() {
	var (
		kubeClient client.Client
	)

	BeforeEach(func() {
		kubeClient = fake.NewClientBuilder().Build()
	})

	Context("cluster api route", func() {
		It("should create cluster-api route when running in openshift", func() {
			runaiConfig := utils.GetDefaultRunaiConfig()
			runaiConfig.Spec.Internal.Global.Platform.KubernetesFlavor = runv1_internal.KubernetesFlavorOpenShift

			route := &osv1.Route{}
			err := initClusterApiRoute(context.Background(), kubeClient, route, runaiConfig)
			Expect(err).NotTo(HaveOccurred())
			Expect(route).NotTo(BeNil())
		})

		It("should propagate tls secret to cluster-api route", func() {
			runaiConfig := utils.GetDefaultRunaiConfig()
			runaiConfig.Spec.Internal.Global.Platform.KubernetesFlavor = runv1_internal.KubernetesFlavorOpenShift
			runaiConfig.Spec.ClusterApi.Route.TlsSecret = "tls-secret"

			err := kubeClient.Create(context.Background(), &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "tls-secret",
					Namespace: runaiConfig.Namespace,
				},
				Data: map[string][]byte{
					"tls.crt": []byte("tls-cert"),
					"tls.key": []byte("tls-key"),
				},
			})
			Expect(err).NotTo(HaveOccurred())

			route := &osv1.Route{}
			err = initClusterApiRoute(context.Background(), kubeClient, route, runaiConfig)
			Expect(err).NotTo(HaveOccurred())
			Expect(route.Spec.TLS).NotTo(BeNil())
			Expect(route.Spec.TLS.Certificate).To(Equal("tls-cert"))
			Expect(route.Spec.TLS.Key).To(Equal("tls-key"))
		})
	})

	Context("getTlsSecretCertAndKey", func() {
		It("should return tls secret cert and key", func() {
			secret := &corev1.Secret{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "secret",
					Namespace: "namespace",
				},
				Data: map[string][]byte{
					"tls.crt": []byte("cert"),
					"tls.key": []byte("key"),
				},
			}
			k8sReader := fake.NewClientBuilder().WithObjects(secret).Build()

			cert, key, err := getTlsSecretCertAndKey(context.Background(), k8sReader, "secret", "namespace")
			Expect(err).NotTo(HaveOccurred())
			Expect(cert).To(Equal("cert"))
			Expect(key).To(Equal("key"))
		})

		It("should return error if secret not found", func() {
			cert, key, err := getTlsSecretCertAndKey(context.Background(), kubeClient, "secret", "namespace")
			Expect(err).To(HaveOccurred())
			Expect(cert).To(BeEmpty())
			Expect(key).To(BeEmpty())
		})
	})
})
