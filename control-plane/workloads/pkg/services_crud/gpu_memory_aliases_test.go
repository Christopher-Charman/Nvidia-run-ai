/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package services_crud

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/dsx-ai-factory/run-ai/control-plane/workloads-and-policies/composer"
	"github.com/dsx-ai-factory/run-ai/control-plane/workloads-and-policies/gpucompat"
	"github.com/dsx-ai-factory/run-ai/control-plane/workloads/internal/generated/pkg/api"
	workloadsApi "github.com/dsx-ai-factory/run-ai/control-plane/workloads/pkg/client/generated"
	"github.com/dsx-ai-factory/run-ai/runai-application-packages/jsonutils/gojsonq"
	corev1 "k8s.io/api/core/v1"
)

func TestWorkspaceMemoryAliasPolicyPipeline(t *testing.T) {
	for _, legacyDefault := range []bool{true, false} {
		for _, locked := range []bool{true, false} {
			for _, value := range []float64{.25, .5} {
				t.Run(fmt.Sprintf("legacy-default=%t/locked=%t/value=%v", legacyDefault, locked, value), func(t *testing.T) {
					defaultKey, inputKey := "gpuPortionRequest", "gpuMemoryPortionRequest"
					if !legacyDefault {
						defaultKey, inputKey = inputKey, defaultKey
					}
					var request api.WorkspaceCreationRequest
					body := fmt.Sprintf(`{"spec":{"image":"example/image","compute":{"gpuDevicesRequest":1,"gpuRequestType":"portion",%q:%v}}}`, inputKey, value)
					if err := json.Unmarshal([]byte(body), &request); err != nil {
						t.Fatal(err)
					}
					input, err := gojsonq.LoadJson(request.Spec)
					if err != nil {
						t.Fatal(err)
					}
					defaults := map[string]any{"compute": map[string]any{defaultKey: json.Number("0.25")}}
					rules := map[string]any{"compute": map[string]any{defaultKey: map[string]any{"canEdit": !locked}}}
					mutated, mutateErr := MutateWorkload(input, defaults, false, nil, nil)
					if mutateErr != nil {
						t.Fatal(mutateErr)
					}
					validationErr := validateSpecAgainstPolicy(Workspace, rules, defaults, mutated)
					if locked && value != .25 {
						if validationErr == nil || validationErr.Code != 400 {
							t.Fatalf("lock bypassed: %v", validationErr)
						}
						return
					}
					if validationErr != nil {
						t.Fatal(validationErr)
					}
					if defaults["compute"].(map[string]any)[defaultKey] != json.Number("0.25") {
						t.Fatal("policy default changed")
					}
					spec, _, rebuildErr := buildResponseSpecFromMap[workloadsApi.SupersetSpec, workloadsApi.MasterSpec](Workspace, mutated)
					if rebuildErr != nil {
						t.Fatal(rebuildErr)
					}
					compute := spec.Compute.Get()
					if compute.GpuMemoryPortionRequest.Get() == nil || *compute.GpuMemoryPortionRequest.Get() != value {
						t.Fatal("canonical value lost")
					}
					if compute.GpuPortionRequest.IsSet() {
						t.Fatal("deprecated value in canonical spec")
					}
					assertWorkspaceMemoryAliasRoundTrip(t, spec, value)
				})
			}
		}
	}
}

func assertWorkspaceMemoryAliasRoundTrip(t *testing.T, spec *workloadsApi.SupersetSpec, value float64) {
	t.Helper()
	compute := spec.Compute.Get()
	var sizing composer.GpuSizingInfo
	var resources corev1.ResourceRequirements
	fc := composer.FieldsComposer{}
	err := fc.ComposeGpuRequestsAndLimits(context.Background(), compute.GpuDevicesRequest.Get(), compute.GpuRequestType.Get(), compute.GpuMemoryPortionRequest.Get(), compute.GpuMemoryPortionLimit.Get(), compute.GpuMemoryRequest.Get(), compute.GpuMemoryLimit.Get(), composer.WholeGpusToK8sResources, &resources, nil, &sizing)()
	if err != nil {
		t.Fatal(err)
	}
	if sizing.PortionRequest == nil || *sizing.PortionRequest != value {
		t.Fatal("native allocation lost")
	}
	params := createWorkloadParams{Response: workloadsApi.SupersetWorkload{Spec: spec}}
	record := NewWorkloadFromWorkloadParams(&params, nil)
	stored, marshalErr := json.Marshal(record.Spec)
	if marshalErr != nil {
		t.Fatal(marshalErr)
	}
	if strings.Contains(string(stored), "gpuMemoryPortion") || !strings.Contains(string(stored), "gpuPortionRequest") {
		t.Fatalf("storage format changed: %s", stored)
	}
	response := WorkloadFromDbRecord(record, workloadsApi.PHASE_CREATING)
	returned := response.Spec.Compute.Get()
	if returned.GetGpuPortionRequest() != value || returned.GetGpuMemoryPortionRequest() != value {
		t.Fatal("response aliases differ")
	}
	*returned.GpuPortionRequest.Get() = .9
	if record.Spec.Compute.Get().GetGpuPortionRequest() != value || spec.Compute.Get().GetGpuMemoryPortionRequest() != value {
		t.Fatal("response mutated stored/canonical values")
	}
}

func TestWorkspaceMemoryAliasSourcePrecedence(t *testing.T) {
	input := map[string]any{"compute": map[string]any{"gpuMemoryPortionRequest": json.Number("0.5")}}
	assets := map[string]any{"compute": map[string]any{"gpuPortionRequest": json.Number("0.4")}}
	template := map[string]any{"compute": map[string]any{"gpuMemoryPortionRequest": json.Number("0.3")}}
	defaults := map[string]any{"compute": map[string]any{"gpuPortionRequest": json.Number("0.25")}}
	for _, tc := range []struct {
		name                    string
		input, assets, template map[string]any
		want                    float64
	}{
		{"request", input, assets, template, .5}, {"asset", nil, assets, template, .4}, {"template", nil, nil, template, .3}, {"policy", nil, nil, nil, .25},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := MutateWorkload(tc.input, defaults, false, tc.template, tc.assets)
			if err != nil {
				t.Fatal(err)
			}
			if fmt.Sprint(got["compute"].(map[string]any)["gpuMemoryPortionRequest"]) != fmt.Sprint(tc.want) {
				t.Fatalf("unexpected source precedence: %#v", got)
			}
		})
	}
}

func TestWorkspaceMemoryAliasConflicts(t *testing.T) {
	for _, source := range []string{"request", "policy", "template", "asset"} {
		t.Run(source, func(t *testing.T) {
			conflict := map[string]any{"compute": map[string]any{"gpuPortionRequest": .25, "gpuMemoryPortionRequest": json.Number("0.5")}}
			var input, defaults, template, assets map[string]any
			switch source {
			case "request":
				input = conflict
			case "policy":
				defaults = conflict
			case "template":
				template = conflict
			case "asset":
				assets = conflict
			}
			_, err := MutateWorkload(input, defaults, false, template, assets)
			if err == nil || !strings.Contains(err.Message, "conflict") {
				t.Fatalf("missing alias conflict: %v", err)
			}
		})
	}
}

func TestWorkspaceMemoryAliasesPreserveOmissionAndLimits(t *testing.T) {
	var spec workloadsApi.SupersetSpec
	if err := json.Unmarshal([]byte(`{"compute":{"gpuPortionRequest":0.25,"gpuPortionLimit":0.5}}`), &spec); err != nil {
		t.Fatal(err)
	}
	response := gpucompat.LegacyCompatibleWorkloadSpec(&spec)
	if response.Compute.Get().GetGpuMemoryPortionLimit() != .5 {
		t.Fatal("limit alias lost")
	}
	stored := gpucompat.StoredWorkloadSpec(response)
	if !reflect.DeepEqual(stored, &spec) {
		t.Fatal("legacy persisted representation changed")
	}
	if gpucompat.LegacyCompatibleWorkloadSpec(&workloadsApi.SupersetSpec{}).Compute.IsSet() {
		t.Fatal("fabricated compute section")
	}
}

func TestWorkspaceMemoryAliasUpdateSemantics(t *testing.T) {
	for _, tc := range []struct {
		name, input    string
		request, limit *float64
		conflict       bool
	}{
		{name: "omitted", input: `{}`, request: gpuMemoryFloat(.25), limit: gpuMemoryFloat(.5)},
		{name: "null", input: `{"compute":null}`, request: gpuMemoryFloat(.25), limit: gpuMemoryFloat(.5)},
		{name: "empty section", input: `{"compute":{}}`},
		{name: "new replaces section", input: `{"compute":{"gpuMemoryPortionRequest":0.75}}`, request: gpuMemoryFloat(.75)},
		{name: "old replaces section", input: `{"compute":{"gpuPortionRequest":0.75}}`, request: gpuMemoryFloat(.75)},
		{name: "explicit zero", input: `{"compute":{"gpuMemoryPortionRequest":0}}`, request: gpuMemoryFloat(0)},
		{name: "conflict", input: `{"compute":{"gpuPortionRequest":0.25,"gpuMemoryPortionRequest":0.75}}`, conflict: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var original workloadsApi.SupersetSpec
			if err := json.Unmarshal([]byte(`{"compute":{"gpuPortionRequest":0.25,"gpuPortionLimit":0.5},"image":"image"}`), &original); err != nil {
				t.Fatal(err)
			}
			var update api.InferenceUpdateSpecSpec
			if err := json.Unmarshal([]byte(tc.input), &update); err != nil {
				t.Fatal(err)
			}
			err := copyUpdateSpecIntoFullSpec(context.Background(), &original, &update)
			if tc.conflict {
				if err == nil {
					t.Fatal("conflict accepted")
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			got := original.GetCompute()
			if !reflect.DeepEqual(got.GpuMemoryPortionRequest.Get(), tc.request) || !reflect.DeepEqual(got.GpuMemoryPortionLimit.Get(), tc.limit) {
				t.Fatalf("unexpected section merge: %#v", got)
			}
			if original.GetImage() != "image" {
				t.Fatal("lost unrelated field")
			}
		})
	}
}

func gpuMemoryFloat(value float64) *float64 { return &value }

func TestGPUMemoryTemplateProjectionAndReset(t *testing.T) {
	var input api.SupersetTemplate
	if err := json.Unmarshal([]byte(`{"spec":{"compute":{"gpuPortionRequest":0.25}},"combinedSpec":{"compute":{"gpuPortionRequest":0.25}}}`), &input); err != nil {
		t.Fatal(err)
	}
	canonical, err := normalizeTemplateGPUMemory(&input)
	if err != nil {
		t.Fatal(err)
	}
	if canonical.Spec.Compute.GpuMemoryPortionRequest == nil || canonical.Spec.Compute.GpuPortionRequest != nil {
		t.Fatal("template not canonical")
	}
	projected, err := legacyCompatibleTemplate(canonical)
	if err != nil {
		t.Fatal(err)
	}
	if *projected.Spec.Compute.GpuPortionRequest != .25 || *projected.CombinedSpec.Compute.GpuMemoryPortionRequest != .25 {
		t.Fatal("template response aliases differ")
	}
	request := &UnifiedRequest{template: canonical}
	if err := request.ClearFields(updateOptions{fieldsToClear: []string{"spec.compute.gpuPortionRequest"}}); err != nil {
		t.Fatal(err)
	}
	if request.template.Spec.Compute != nil && request.template.Spec.Compute.GpuMemoryPortionRequest != nil {
		t.Fatal("deprecated reset path missed canonical field")
	}
	if input.Spec.Compute.GpuPortionRequest == nil || *input.Spec.Compute.GpuPortionRequest != .25 {
		t.Fatal("template input was mutated")
	}
}
