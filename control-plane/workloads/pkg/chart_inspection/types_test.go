// SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package chart_inspection

import (
	"encoding/json"
	"testing"
)

func TestWarningJSONAlwaysIncludesPath(t *testing.T) {
	encoded, err := json.Marshal(Warning{Code: "EXAMPLE", Message: "Example"})
	if err != nil {
		t.Fatal(err)
	}
	var got map[string]any
	if err := json.Unmarshal(encoded, &got); err != nil {
		t.Fatal(err)
	}
	if _, exists := got["path"]; !exists {
		t.Fatalf("warning path omitted: %s", encoded)
	}
}
