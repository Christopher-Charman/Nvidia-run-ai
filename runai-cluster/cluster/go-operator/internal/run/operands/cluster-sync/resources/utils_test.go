/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package resources

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	"k8s.io/utils/ptr"

	runaiv1 "github.com/dsx-ai-factory/run-ai/runai-cluster/cluster/sdk/apis/run/v1"
)

var _ = Describe("handleEnvVars", func() {
	envValue := func(envVars []v1.EnvVar, name string) (string, bool) {
		for _, envVar := range envVars {
			if envVar.Name == name {
				return envVar.Value, true
			}
		}
		return "", false
	}

	buildConfig := func() *runaiv1.RunaiConfig {
		runaiConfig := &runaiv1.RunaiConfig{}
		runaiConfig.SetDefaultsWhereNeeded()
		return runaiConfig
	}

	It("defaults GPU fractioning readiness reporting to on", func() {
		value, found := envValue(handleEnvVars(buildConfig()), "ENABLE_NODE_GPU_FRACTIONING_READINESS")
		Expect(found).To(BeTrue())
		Expect(value).To(Equal("true"))
	})

	It("honours an explicit opt-out", func() {
		runaiConfig := buildConfig()
		runaiConfig.Spec.ClusterSync.EnableNodeGpuFractioningReadiness = ptr.To(false)

		value, _ := envValue(handleEnvVars(runaiConfig), "ENABLE_NODE_GPU_FRACTIONING_READINESS")
		Expect(value).To(Equal("false"))
	})
})
