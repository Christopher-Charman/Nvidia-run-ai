// SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package chart_inspection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

func inspectQuestions(t *testing.T, questions, values string) *Result {
	t.Helper()
	got, err := ParseArchive(chartArchive(t, archiveEntry{name: "demo/Chart.yaml", body: "name: demo\nversion: 1.0.0\n"}, archiveEntry{name: "demo/questions.yaml", body: questions}, archiveEntry{name: "demo/values.yaml", body: values}))
	if err != nil {
		t.Fatal(err)
	}
	return got
}

func hasWarning(got *Result, code string) bool {
	for _, warning := range got.Warnings {
		if warning.Code == code {
			return true
		}
	}
	return false
}

func TestDeclarationNormalizesCompatibleMetadata(t *testing.T) {
	got := inspectQuestions(t, `apiVersion: ignored/v1
groups:
  - name: Advanced
    label: Advanced settings
    description: Declared but unused
questions:
  - variable: enabled
    type: boolean
    default: "false"
    show_subquestions_if: true
    vendor_hint: ignored
    subquestions:
      - variable: workers[0].name
        label: Worker
        type: unfamiliar
  - variable: scale
    label: Scale
    type: number
    default: 1.25
`, "enabled: true\nworkers: [{name: original}]\nscale: 1\n")
	if got.QuestionsStatus != "usable" || len(got.Questions) != 2 {
		t.Fatalf("declaration = %#v", got)
	}
	if got.Questions[0]["label"] != "enabled" || got.Questions[0]["default"] != false || got.Questions[0]["show_subquestion_if"] != true {
		t.Fatalf("normalization = %#v", got.Questions[0])
	}
	child := got.Questions[0]["subquestions"].([]any)[0].(map[string]any)
	if child["type"] != "string" || got.Questions[1]["type"] != "float" {
		t.Fatalf("type normalization = %#v", got.Questions)
	}
	if _, exists := got.Questions[0]["vendor_hint"]; exists {
		t.Fatal("unknown property retained")
	}
	if len(got.Groups) != 2 || got.Groups[0]["name"] != "advanced" {
		t.Fatalf("declared groups not preserved: %#v", got.Groups)
	}
	for _, code := range []string{"LABEL_DEFAULTED", "UNKNOWN_TYPE", "IGNORED_PROPERTY"} {
		if !hasWarning(got, code) {
			t.Fatalf("missing %s warning: %#v", code, got.Warnings)
		}
	}
}

func TestDeclarationAcceptsSupportedFieldDefaults(t *testing.T) {
	for _, tc := range []struct{ kind, value, extra string }{
		{"string", `hello`, ""}, {"multiline", `hello`, ""}, {"password", `'<your-password>'`, ""},
		{"storageclass", `fast`, ""}, {"pvc", `data`, ""}, {"configmap", `config`, ""},
		{"int", `3`, "min: 1\n      max: 5"}, {"float", `1.5`, ""}, {"boolean", `true`, ""},
		{"enum", `1`, "options: [1, two]"}, {"radio", `false`, "options: [false, true]"},
		{"hostname", `example.com`, ""}, {"ipaddr", `::1`, ""}, {"cidr", `10.1.2.3/24`, ""},
		{"cron", `'*/5 * * * *'`, ""}, {"array[string]", `[a, b]`, ""}, {"map[string]", `{a: b}`, ""},
	} {
		t.Run(tc.kind, func(t *testing.T) {
			body := fmt.Sprintf("questions:\n  - variable: setting\n    label: Setting\n    type: '%s'\n    default: %s\n", tc.kind, tc.value)
			if tc.extra != "" {
				body += "    " + strings.ReplaceAll(tc.extra, "\n      ", "\n    ") + "\n"
			}
			got := inspectQuestions(t, body, "setting: null\n")
			if got.QuestionsStatus != "usable" {
				t.Fatalf("supported field rejected: %#v", got.Warnings)
			}
		})
	}
}

func TestDeclarationRejectsWholeInvalidForm(t *testing.T) {
	for _, tc := range []struct{ name, body string }{
		{"malformed YAML", "questions: [\n"}, {"root array", "- variable: x\n"},
		{"duplicate keys", "questions: []\nquestions: []\n"}, {"missing questions", "groups: []\n"},
		{"missing variable", "questions: [{type: string}]"}, {"missing type", "questions: [{variable: x}]"},
		{"secret", "questions: [{variable: x, type: secret}]"}, {"credential", "questions: [{variable: x, type: cloud-credential}]"},
		{"numeric dotted path", "questions: [{variable: workers.0.name, type: string}]"},
		{"prototype path", "questions: [{variable: 'a[\"__proto__\"]', type: string}]"},
		{"duplicate destination", "questions: [{variable: x, type: string}, {variable: x, type: string}]"},
		{"overlapping destination", "questions: [{variable: x, type: string}, {variable: x.y, type: string}]"},
		{"bad default", "questions: [{variable: x, type: int, default: nope}]"},
		{"default below bound", "questions: [{variable: x, type: int, min: 2, default: 1}]"},
		{"bad array default", "questions: [{variable: x, type: 'array[string]', default: [one, 2]}]"},
		{"bad IP", "questions: [{variable: x, type: ipaddr, default: 999.1.1.1}]"},
		{"bad hostname", "questions: [{variable: x, type: hostname, default: a_b}]"},
		{"bad cron", "questions: [{variable: x, type: cron, default: '61 * * * *'}]"},
		{"missing enum options", "questions: [{variable: x, type: enum}]"},
		{"duplicate typed options", "questions: [{variable: x, type: enum, options: [1, 1.0]}]"},
		{"unavailable default", "questions: [{variable: x, type: enum, options: [a, b], default: c}]"},
		{"bad option", "questions: [{variable: x, type: enum, options: [{value: a}]}]"},
		{"bad constraint type", "questions: [{variable: x, type: boolean, min: 1}]"},
		{"conflicting gates", "questions: [{variable: x, type: boolean, show_subquestion_if: true, show_subquestions_if: false}]"},
		{"bad condition", "questions: [{variable: x, type: string, show_if: 'y && true'}]"},
		{"condition cycle", "questions: [{variable: x, type: string, show_if: 'y=ready'}, {variable: y, type: string, show_if: 'x=ready'}]"},
		{"parent visibility cycle", "questions: [{variable: x, type: boolean, show_if: 'y=true', subquestions: [{variable: y, type: boolean}]}]"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := inspectQuestions(t, tc.body, "{}\n")
			if got.QuestionsStatus != "unusable" || !got.HasQuestions || len(got.Questions) != 0 || len(got.Groups) != 0 || !hasWarning(got, "UNUSABLE_QUESTIONS") {
				t.Fatalf("invalid declaration was salvaged: %#v", got)
			}
		})
	}
}

func TestDeclarationPreservesFourChildLevels(t *testing.T) {
	makeTree := func(last int) string {
		child := "{variable: p" + fmt.Sprint(last) + ", type: boolean, subquestions: []}"
		for n := last - 1; n >= 0; n-- {
			child = "{variable: p" + fmt.Sprint(n) + ", type: boolean, group: Basics, subquestions: [" + child + "]}"
		}
		return "questions: [" + child + "]"
	}
	got := inspectQuestions(t, makeTree(4), "{}")
	if got.QuestionsStatus != "usable" {
		t.Fatalf("four levels rejected: %#v", got.Warnings)
	}
	child := got.Questions[0]
	for depth := 0; depth < 4; depth++ {
		if child["group"] != "basics" {
			t.Fatalf("group not inherited at %d", depth)
		}
		child = child["subquestions"].([]any)[0].(map[string]any)
	}
	if child["group"] != "basics" {
		t.Fatal("leaf group not inherited")
	}
	if got := inspectQuestions(t, makeTree(5), "{}"); got.QuestionsStatus != "unusable" {
		t.Fatalf("fifth child level accepted: %#v", got)
	}
}

func TestDeclarationWarnsForDynamicValuesAndConditions(t *testing.T) {
	got := inspectQuestions(t, "questions: [{variable: dynamic, type: string, show_if: 'name= || ready=true'}]", "ready: true\n")
	if got.QuestionsStatus != "usable" || !hasWarning(got, "MISSING_DEFAULT_PATH") || !hasWarning(got, "UNKNOWN_CONDITION_REFERENCE") {
		t.Fatalf("dynamic input rejected or unreported: %#v", got)
	}
}

func TestDeclarationRejectsNullGroupsAndComplexGates(t *testing.T) {
	for _, body := range []string{
		"groups: null\nquestions: []",
		"questions: [{variable: x, type: boolean, show_subquestion_if: [true], show_subquestions_if: [true]}]",
		"questions: [{variable: x, type: boolean, subquestions: null}]",
		"questions: [{variable: x, type: boolean, default: 'True'}]",
	} {
		got := inspectQuestions(t, body, "{}")
		if got.QuestionsStatus != "unusable" {
			t.Fatalf("invalid structure accepted: %s", body)
		}
	}
}

func TestDeclarationIgnoresOptionalOptionMetadata(t *testing.T) {
	got := inspectQuestions(t, "questions: [{variable: mode, type: enum, options: [{label: Basic, value: basic, description: extra}]}]", "mode: basic")
	if got.QuestionsStatus != "usable" || !hasWarning(got, "IGNORED_PROPERTY") {
		t.Fatalf("optional option metadata rejected: %#v", got)
	}
	option := got.Questions[0]["options"].([]any)[0].(map[string]any)
	if _, exists := option["description"]; exists {
		t.Fatal("optional option metadata retained")
	}
}

func TestDeclarationWarnsWhenWhitespaceLabelIsDefaulted(t *testing.T) {
	got := inspectQuestions(t, "questions: [{variable: setting, label: '   ', type: string}]", "setting: original")
	if got.QuestionsStatus != "usable" || len(got.Questions) != 1 || got.Questions[0]["label"] != "setting" {
		t.Fatalf("whitespace label was not normalized: %#v", got)
	}
	for _, warning := range got.Warnings {
		if warning.Code == "LABEL_DEFAULTED" && warning.Path == "questions[0].label" {
			return
		}
	}
	t.Fatalf("normalized label has no warning: %#v", got.Warnings)
}

func TestDeclarationRejectsIncompatibleDestinationContainers(t *testing.T) {
	for _, pair := range [][2]string{
		{"workers[0].name", "workers.enabled"},
		{"workers.enabled", "workers[0].name"},
		{"workers[0].slots[0]", "workers[0].slots.name"},
	} {
		got := inspectQuestions(t, fmt.Sprintf("questions: [{variable: '%s', type: string}, {variable: '%s', type: string}]", pair[0], pair[1]), "{}")
		if got.QuestionsStatus != "unusable" || len(got.Questions) != 0 || !hasWarning(got, "UNUSABLE_QUESTIONS") {
			t.Fatalf("incompatible containers accepted: %v: %#v", pair, got)
		}
	}
	for _, pair := range [][2]string{{"workers[0].name", "workers[1].name"}, {"workers.name", "workers.enabled"}} {
		got := inspectQuestions(t, fmt.Sprintf("questions: [{variable: '%s', type: string}, {variable: '%s', type: string}]", pair[0], pair[1]), "{}")
		if got.QuestionsStatus != "usable" {
			t.Fatalf("compatible containers rejected: %v: %#v", pair, got.Warnings)
		}
	}
}

func TestDeclarationAcceptsNumericPrefixTextConditions(t *testing.T) {
	for _, value := range []string{"1.2.3", "1.2.3-beta", "5m"} {
		t.Run(value, func(t *testing.T) {
			body := fmt.Sprintf("questions: [{variable: setting, type: string, show_if: 'image.tag=%s'}]", value)
			got := inspectQuestions(t, body, "image: {tag: original}")
			if got.QuestionsStatus != "usable" || got.Questions[0]["show_if"] != "image.tag="+value {
				t.Fatalf("unquoted text condition rejected or changed: %#v", got)
			}
		})
	}
}

func TestDeclarationPreservesDistinctLargeIntegerOptions(t *testing.T) {
	got := inspectQuestions(t, "questions: [{variable: mode, type: enum, options: [9007199254740992, 9007199254740993], default: 9007199254740993}]", "mode: 9007199254740993")
	if got.QuestionsStatus != "usable" || len(got.Questions) != 1 {
		t.Fatalf("distinct integer options rejected: %#v", got)
	}
	encoded, err := json.Marshal(got.Questions[0])
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(encoded, []byte(`"default":9007199254740993`)) {
		t.Fatalf("large integer default rounded: %s", encoded)
	}
}

func TestDeclarationPreservesIntegersBeyondUint64(t *testing.T) {
	got := inspectQuestions(t, "questions: [{variable: mode, type: enum, options: [18446744073709551616, 18446744073709551617], default: 18446744073709551617}]", "mode: 18446744073709551617")
	if got.QuestionsStatus != QuestionsUsable || len(got.Questions) != 1 {
		t.Fatalf("distinct integers were rounded: %#v", got.Warnings)
	}
	encoded, err := json.Marshal(got.Questions[0])
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(encoded, []byte(`"default":18446744073709551617`)) {
		t.Fatalf("integer default rounded: %s", encoded)
	}
}

func TestDeclarationRejectsExactLargeIntegerBoundViolations(t *testing.T) {
	for _, extra := range []string{
		"max: 9007199254740992, default: 9007199254740993",
		"min: 9007199254740993, default: 9007199254740992",
		"min: 9007199254740993, max: 9007199254740992",
		"min: -9007199254740992, default: -9007199254740993",
		"max: 18446744073709551614, default: 18446744073709551615",
	} {
		got := inspectQuestions(t, "questions: [{variable: number, type: int, "+extra+"}]", "{}")
		if got.QuestionsStatus != "unusable" || len(got.Questions) != 0 {
			t.Fatalf("exact integer bound ignored: %s: %#v", extra, got)
		}
	}
}
