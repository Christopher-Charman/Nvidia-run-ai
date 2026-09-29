/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package services

import (
	"encoding/json"

	"github.com/dsx-ai-factory/run-ai/control-plane/workloads/internal/generated/pkg/api"
	"github.com/dsx-ai-factory/run-ai/control-plane/workloads/pkg/database/models"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("AI application reconciliation status", func() {
	It("exposes credential revocation and recovery despite a previously deployed release", func() {
		var report api.AiApplicationForSync
		Expect(json.Unmarshal([]byte(`{
			"k8sPhase":"deployed",
			"actionInfo":{
				"status":"completed", "helmPhase":"deployed",
				"message":"install completed",
				"helmError":"personal credentials unavailable"
			}
		}`), &report)).To(Succeed())
		report.Id = uuid.NewString()

		failed, err := mapSyncApplicationToModel("tenant", "cluster", report)
		Expect(err).NotTo(HaveOccurred())
		response := applicationViewToApi(models.ApplicationView{
			ApplicationStatus: failed.ApplicationStatus,
			ActualPhase:       failed.ActualPhase,
			Message:           failed.Message,
		}, api.AIAPPLICATIONVERBOSITY_FULL)
		Expect(response.Phase).To(Equal("failed"))
		Expect(response.ActionStatus).To(HaveValue(Equal("failed")))
		Expect(response.ActionMessage).To(HaveValue(Equal("personal credentials unavailable")))

		report.ActionInfo.HelmError = ""
		recovered, err := mapSyncApplicationToModel("tenant", "cluster", report)
		Expect(err).NotTo(HaveOccurred())
		Expect(recovered.K8sPhase).To(Equal("deployed"))
		Expect(recovered.ActualPhase).To(Equal("deployed"))
		Expect(recovered.Message).To(HaveValue(Equal("install completed")))
	})

	DescribeTable("preserves action status when no reconciliation error exists",
		func(info *api.AiApplicationSyncActionInfo, phase, message string) {
			actualPhase, actualMessage := determineStatus(api.AiApplicationForSync{
				K8sPhase: "deployed", ActionInfo: info,
			})
			Expect(actualPhase).To(Equal(phase))
			if message == "" {
				Expect(actualMessage).To(BeNil())
			} else {
				Expect(actualMessage).To(HaveValue(Equal(message)))
			}
		},
		Entry("no action", nil, "", ""),
		Entry("completed release", &api.AiApplicationSyncActionInfo{Status: "completed", HelmPhase: "deployed"}, "deployed", ""),
		Entry("completed action without Helm phase", &api.AiApplicationSyncActionInfo{Status: "completed"}, "deployed", ""),
		Entry("running uninstall", &api.AiApplicationSyncActionInfo{Status: "running", Message: "uninstalling"}, "running", "uninstalling"),
		Entry("failed action", &api.AiApplicationSyncActionInfo{Status: "failed", Message: "install failed", Error: "chart unavailable"}, "failed", "install failed (chart unavailable)"),
		Entry("reconciliation failure without an action", &api.AiApplicationSyncActionInfo{HelmError: "controller unavailable"}, "failed", "controller unavailable"),
	)
})
