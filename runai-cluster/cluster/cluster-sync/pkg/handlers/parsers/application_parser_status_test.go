/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package parsers_test

import (
	"encoding/json"

	"github.com/dsx-ai-factory/run-ai/runai-cluster/cluster/cluster-sync/pkg/handlers/parsers"
	globalconstants "github.com/dsx-ai-factory/run-ai/runai-cluster/cluster/common/constants"
	"github.com/dsx-ai-factory/run-ai/runai-cluster/cluster/sdk/apis/run/v1alpha1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ = Describe("Application credential status sync", func() {
	newDeployedApplication := func() *v1alpha1.Application {
		releaseVersion := 1
		return &v1alpha1.Application{
			ObjectMeta: metav1.ObjectMeta{
				Name: "bound-app",
				Labels: map[string]string{
					globalconstants.DepartmentIdLabel: "department-id",
					globalconstants.ProjectIdLabel:    "project-id",
					globalconstants.ProjectNameLabel:  "project-name",
				},
			},
			Status: v1alpha1.ApplicationStatus{
				K8sPhase: "deployed",
				ActionStatus: &v1alpha1.ActionStatus{
					Type: "install", Phase: "completed", Message: "install completed",
				},
				HelmStatus: &v1alpha1.HelmStatus{Phase: "deployed", ReleaseVersion: &releaseVersion},
			},
		}
	}

	DescribeTable("forwards credential loss even when Helm remains deployed",
		func(credentialError string) {
			application := newDeployedApplication()
			application.Status.HelmStatus.Error = credentialError

			sync, err := parsers.ParseBackendApplication(application, false)
			Expect(err).NotTo(HaveOccurred())
			Expect(sync.GetK8sPhase()).To(Equal("deployed"))
			Expect(sync.HasActionInfo()).To(BeTrue())
			info := sync.GetActionInfo()
			Expect(info.GetStatus()).To(Equal("completed"))
			Expect(info.GetHelmPhase()).To(Equal("deployed"))
			Expect(info.GetHelmError()).To(Equal(credentialError))

			// Check the wire representation consumed by workloads, not just the Go struct.
			body, err := json.Marshal(sync)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(body)).To(ContainSubstring(`"helmError":"` + credentialError + `"`))
		},
		Entry("shared Secret removed", "application credential Secrets unavailable"),
		Entry("personal source removed", "personal credentials unavailable"),
	)

	It("omits the reconciliation error once the credential is restored", func() {
		application := newDeployedApplication()
		sync, err := parsers.ParseBackendApplication(application, false)
		Expect(err).NotTo(HaveOccurred())
		info := sync.GetActionInfo()
		Expect(info.HasHelmError()).To(BeFalse())
	})
})
