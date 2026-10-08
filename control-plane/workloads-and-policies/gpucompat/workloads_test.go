/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package gpucompat

import (
	"encoding/json"
	workloadsApi "github.com/dsx-ai-factory/run-ai/control-plane/workloads/pkg/client/generated"
	"testing"
)

func TestNullableResponseMemoryAliases(t *testing.T) {
	for _, key := range []string{"gpuPortionRequest", "gpuMemoryPortionRequest", "gpuPortionLimit", "gpuMemoryPortionLimit"} {
		t.Run(key, func(t *testing.T) {
			var spec workloadsApi.SupersetSpec
			if err := json.Unmarshal([]byte(`{"compute":{"`+key+`":null}}`), &spec); err != nil {
				t.Fatal(err)
			}
			response := LegacyCompatibleWorkloadSpec(&spec).GetCompute()
			old, current := response.GpuPortionRequest, response.GpuMemoryPortionRequest
			if key == "gpuPortionLimit" || key == "gpuMemoryPortionLimit" {
				old, current = response.GpuPortionLimit, response.GpuMemoryPortionLimit
			}
			if !old.IsSet() || !current.IsSet() || old.Get() != nil || current.Get() != nil {
				t.Fatal("null alias presence diverged")
			}
		})
	}
}
