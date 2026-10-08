/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package services

import (
	"context"
	"encoding/json"
	"github.com/dsx-ai-factory/run-ai/control-plane/workloads-manager/internal/generated/pkg/api"
	"github.com/dsx-ai-factory/run-ai/runai-application-packages/jsonutils/gojsonq"
	"testing"
)

func TestGPUMemoryNimPolicyAndUpdate(t *testing.T) {
	for _, oldDefault := range []bool{false, true} {
		inputKey, defaultKey := "gpuMemoryPortionRequest", "gpuPortionRequest"
		if !oldDefault {
			inputKey, defaultKey = defaultKey, inputKey
		}
		t.Run(defaultKey, func(t *testing.T) {
			for _, locked := range []bool{false, true} {
				var spec api.NimServiceSpec
				if err := json.Unmarshal([]byte(`{"compute":{"`+inputKey+`":0.5}}`), &spec); err != nil {
					t.Fatal(err)
				}
				defaults := map[string]any{"compute": map[string]any{defaultKey: json.Number("0.25")}}
				rules := map[string]any{"compute": map[string]any{defaultKey: map[string]any{"canEdit": !locked}}}
				mutated, err := MutateWorkload(&spec, nil, defaults)
				if err != nil {
					t.Fatal(err)
				}
				_, validationErr := validateMutatedSpecVsPolicy(rules, defaults, mutated)
				if (validationErr != nil) != locked {
					t.Fatalf("lock=%v, error=%v", locked, validationErr)
				}
			}
		})
	}
	var full, update api.NimServiceSpec
	if err := json.Unmarshal([]byte(`{"compute":{"gpuPortionRequest":0.25,"gpuPortionLimit":0.5},"image":"image"}`), &full); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal([]byte(`{"compute":{"gpuMemoryPortionRequest":0.75}}`), &update); err != nil {
		t.Fatal(err)
	}
	merged, err := (&NimServicesImpl{}).mergeUpdateSpecIntoFullSpec(context.Background(), update, full, nil)
	if err != nil {
		t.Fatal(err)
	}
	if merged.Compute.GpuMemoryPortionRequest == nil || *merged.Compute.GpuMemoryPortionRequest != .75 || merged.Compute.GpuPortionLimit != nil || merged.Compute.GpuMemoryPortionLimit != nil {
		t.Fatal("compute section replacement changed")
	}
	stored := storedNimSpec(merged)
	response := legacyCompatibleNimSpec(stored)
	values, jsonErr := gojsonq.LoadJson(response)
	if jsonErr != nil {
		t.Fatal(jsonErr)
	}
	fields := values["compute"].(map[string]any)
	if fields["gpuPortionRequest"] != fields["gpuMemoryPortionRequest"] || stored.Compute.GpuMemoryPortionRequest != nil {
		t.Fatal("round trip lost storage/response contract")
	}
}
