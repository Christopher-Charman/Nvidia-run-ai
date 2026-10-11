/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package gpumode

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestResolveSubmittedModes(t *testing.T) {
	for _, tc := range []struct{ name, body, want string }{
		{"portion fallback", `{"gpuMemoryPortionRequest":0.25}`, TimeSlicing},
		{"legacy whole GPU", `{"gpuPortionRequest":2}`, TimeSlicing},
		{"devices", `{"gpuDevicesRequest":2}`, TimeSlicing},
		{"absolute memory", `{"gpuMemoryRequest":"1Gi","gpuComputeMode":"sm-sharing"}`, SMSharing},
		{"CPU only", `{"gpuDevicesRequest":0,"gpuComputeMode":"sm-sharing"}`, ""},
		{"no allocation", `{}`, ""},
		{"spoofed result", `{"gpuMemoryPortionRequest":0.5,"computeMode":"sm-sharing"}`, TimeSlicing},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var compute map[string]interface{}
			require.NoError(t, json.Unmarshal([]byte(tc.body), &compute))
			for _, replica := range []string{"worker", "master", "leader"} {
				values := map[string]interface{}{replica: map[string]interface{}{"compute": compute}}
				ResolveEffectiveModes(values)
				if tc.want == "" {
					require.NotContains(t, compute, "computeMode")
				} else {
					require.Equal(t, tc.want, compute["computeMode"])
				}
			}
		})
	}
}

func TestResponseModeNeverBecomesRequestMode(t *testing.T) {
	compute := map[string]interface{}{"gpuDevicesRequest": json.Number("1"), "computeMode": SMSharing}
	arbitrary := map[string]interface{}{"compute": map[string]interface{}{"computeMode": "application value"}}
	values := map[string]interface{}{"compute": compute, "annotations": arbitrary}
	StripEffectiveModes(values)
	require.NotContains(t, compute, "computeMode")
	require.Equal(t, "application value", arbitrary["compute"].(map[string]interface{})["computeMode"])
	apply, cleanup := PolicyFallback()
	apply(values)
	require.Equal(t, TimeSlicing, compute["gpuComputeMode"])
	cleanup()
	require.NotContains(t, compute, "gpuComputeMode")
}

func TestUpdateModeUsesSubmittedAllocation(t *testing.T) {
	resolved := map[string]interface{}{"compute": map[string]interface{}{"gpuDevicesRequest": 1, "gpuComputeMode": SMSharing}}
	submitted := map[string]interface{}{"compute": map[string]interface{}{"cpuCoreRequest": 1, "computeMode": TimeSlicing}}
	ProjectResolvedModes(resolved, submitted)
	compute := submitted["compute"].(map[string]interface{})
	require.Equal(t, SMSharing, compute["gpuComputeMode"])
	require.NotContains(t, compute, "gpuDevicesRequest")
	require.NotContains(t, compute, "computeMode")
}

func TestUnchangedComputePreservesStoredImplicitMode(t *testing.T) {
	compute := map[string]interface{}{"gpuDevicesRequest": 1, "computeMode": TimeSlicing}
	values := map[string]interface{}{"compute": compute}
	cleanup := PreserveImplicitModeForUpdate(values)
	ApplyModeDefaults(values, map[string]interface{}{"compute": map[string]interface{}{"gpuComputeMode": SMSharing}})
	require.Equal(t, TimeSlicing, compute["gpuComputeMode"])
	cleanup()
	require.NotContains(t, compute, "gpuComputeMode")
	ResolveEffectiveModes(values)
	require.Equal(t, TimeSlicing, compute["computeMode"])
}
