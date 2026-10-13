// SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package chart_inspection

import (
	"reflect"
	"testing"
)

func TestValuesPathParsesBrackets(t *testing.T) {
	got, err := parsePath(`workers[0]["display.name"].enabled`)
	if err != nil || !reflect.DeepEqual(got, []any{"workers", 0, "display.name", "enabled"}) {
		t.Fatalf("path = %#v, %v", got, err)
	}
}

func TestValuesPathRejectsUnsafeOrAmbiguousKeys(t *testing.T) {
	for _, value := range []string{"", "a.0", "a[-1]", "a[10001]", "a.constructor", `a["prototype"]`, `a["__proto__"]`, "a..b", "a[b]"} {
		t.Run(value, func(t *testing.T) {
			if _, err := parsePath(value); err == nil {
				t.Fatal("invalid path accepted")
			}
		})
	}
}

func TestConditionRestrictedGrammar(t *testing.T) {
	for _, text := range []string{"true", "false", "enabled=true", "name=longhorn", "name=", "name!=", `(enabled=true || count=2) && name!='hidden'`, `a["display.name"]=hello`} {
		t.Run(text, func(t *testing.T) {
			if _, err := parseCondition(text); err != nil {
				t.Fatal(err)
			}
		})
	}
	for _, text := range []string{"", "x", "!x=true", "x == true", "x=true &&", "(x=true", "constructor=true", "a.0=true", "x=true; run()", "x=true & y=false"} {
		t.Run("reject "+text, func(t *testing.T) {
			if _, err := parseCondition(text); err == nil {
				t.Fatal("invalid condition accepted")
			}
		})
	}
}
