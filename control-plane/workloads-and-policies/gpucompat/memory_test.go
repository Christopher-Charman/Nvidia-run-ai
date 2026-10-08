/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package gpucompat

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestNormalizeLegacyGPUFields(t *testing.T) {
	for _, pair := range memoryAliases {
		for _, tc := range []struct {
			name     string
			input    map[string]any
			want     any
			conflict bool
		}{
			{"old", map[string]any{pair[0]: .25}, .25, false},
			{"new", map[string]any{pair[1]: .25}, .25, false},
			{"equal", map[string]any{pair[0]: .25, pair[1]: .25}, .25, false},
			{"different numeric encoding", map[string]any{pair[0]: json.Number("0.250"), pair[1]: .25}, .25, false},
			{"conflict", map[string]any{pair[0]: .25, pair[1]: .5}, nil, true},
			{"new null", map[string]any{pair[0]: .25, pair[1]: nil}, .25, false},
			{"old null", map[string]any{pair[0]: nil, pair[1]: .25}, .25, false},
			{"zero preserved", map[string]any{pair[0]: 0.0}, 0.0, false},
			{"null preserved", map[string]any{pair[0]: nil}, nil, false},
		} {
			t.Run(pair[0]+"/"+tc.name, func(t *testing.T) {
				source := map[string]any{"compute": tc.input}
				before := cloneMap(source)
				got, err := NormalizeLegacyGPUFields(source)
				if tc.conflict {
					if err == nil || !strings.Contains(err.Error(), pair[0]) || !strings.Contains(err.Error(), pair[1]) {
						t.Fatalf("expected both aliases in conflict error, got %v", err)
					}
				} else {
					if err != nil {
						t.Fatal(err)
					}
					want := map[string]any{"compute": map[string]any{pair[1]: tc.want}}
					if !reflect.DeepEqual(got, want) {
						t.Fatalf("got %#v; want %#v", got, want)
					}
				}
				if !reflect.DeepEqual(source, before) {
					t.Fatal("mutated input source")
				}
			})
		}
	}
}

func TestMemoryStorageAndResponseRoundTrip(t *testing.T) {
	source := map[string]any{"compute": map[string]any{"gpuPortionRequest": .25, "gpuPortionLimit": .5}}
	canonical, err := NormalizeLegacyGPUFields(source)
	if err != nil {
		t.Fatal(err)
	}
	stored, err := ProjectGPUFieldsForStorage(canonical)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(stored, source) {
		t.Fatalf("storage changed: %#v", stored)
	}
	response := LegacyCompatibleGPUFields(stored)
	fields := response["compute"].(map[string]any)
	for _, pair := range memoryAliases {
		if fields[pair[0]] != fields[pair[1]] {
			t.Fatalf("response aliases diverged: %#v", fields)
		}
	}
	fields["gpuPortionRequest"] = .8
	if !reflect.DeepEqual(stored, source) {
		t.Fatal("response changed stored spec")
	}
	if canonical["compute"].(map[string]any)["gpuMemoryPortionRequest"] != .25 {
		t.Fatal("response changed canonical spec")
	}
}

func TestKnownComputeSectionsAndOpaqueFields(t *testing.T) {
	for _, container := range []string{"spec", "masterSpec", "worker", "master", "leader", "defaults", "rules"} {
		t.Run(container, func(t *testing.T) {
			source := map[string]any{container: map[string]any{"compute": map[string]any{"gpuPortionRequest": .25}}}
			got, err := NormalizeLegacyGPUFields(source)
			if err != nil {
				t.Fatal(err)
			}
			compute := got[container].(map[string]any)["compute"].(map[string]any)
			if compute["gpuMemoryPortionRequest"] != .25 {
				t.Fatal(got)
			}
		})
	}
	source := map[string]any{"annotations": map[string]any{"compute": map[string]any{"gpuPortionRequest": "opaque"}}}
	got, err := NormalizeLegacyGPUFields(source)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, source) {
		t.Fatal("rewrote user-defined fields")
	}
	empty := LegacyCompatibleGPUFields(map[string]any{"compute": map[string]any{}})
	if len(empty["compute"].(map[string]any)) != 0 {
		t.Fatal("fabricated absent values")
	}
}

func TestPolicyAliasesAndDerivedReferences(t *testing.T) {
	rules := map[string]any{"compute": map[string]any{
		"gpuPortionRequest": map[string]any{"canEdit": false},
		"cpuCoreRequest":    map[string]any{"defaultFrom": map[string]any{"field": "compute.gpuPortionRequest"}},
	}}
	canonical, err := NormalizeLegacyGPUPolicy(rules)
	if err != nil {
		t.Fatal(err)
	}
	fields := canonical["compute"].(map[string]any)
	if fields["gpuMemoryPortionRequest"].(map[string]any)["canEdit"] != false {
		t.Fatal("lost lock")
	}
	if fields["cpuCoreRequest"].(map[string]any)["defaultFrom"].(map[string]any)["field"] != "compute.gpuMemoryPortionRequest" {
		t.Fatal("lost dependency")
	}
	stored, err := ProjectGPUPolicyForStorage(canonical)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(stored, rules) {
		t.Fatal("stored references do not round trip")
	}
	if CanonicalGPUFieldPath("annotations.gpuPortionRequest") != "annotations.gpuPortionRequest" {
		t.Fatal("rewrote unrelated path")
	}
}
