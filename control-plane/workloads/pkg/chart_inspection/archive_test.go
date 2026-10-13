// SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package chart_inspection

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"strings"
	"testing"
)

func TestInspectorBoundsExpandedYAMLAliases(t *testing.T) {
	for _, filename := range []string{"questions.yaml", "questions.yml"} {
		for _, tc := range []struct{ name, value string }{
			{"scalar bytes", "[&large " + strings.Repeat("x", 32<<10) + strings.Repeat(", *large", 1000) + "]"},
			{"nodes", "[&item x" + strings.Repeat(", *item", 26000) + "]"},
		} {
			t.Run(filename+"/"+tc.name, func(t *testing.T) {
				questions := "questions:\n- variable: data\n  label: Data\n  type: array[string]\n  default: " + tc.value + "\n"
				archive := chartArchive(t,
					archiveEntry{name: "demo/Chart.yaml", body: "name: demo\nversion: 1.0.0\n"},
					archiveEntry{name: "demo/" + filename, body: questions},
				)
				inspector := NewInspector(func(context.Context, Selection, string) ([]byte, error) { return archive, nil }, 0)
				result, err := inspector.Inspect(context.Background(), Selection{ChartID: "authorized-upload"}, "tenant", "")
				if err != nil {
					t.Fatal(err)
				}
				if !result.HasQuestions || result.QuestionsStatus != QuestionsUnusable || len(result.Questions) != 0 || !hasWarning(result, "UNUSABLE_QUESTIONS") {
					t.Fatalf("expanded aliases were accepted: %#v", result)
				}
				encoded, err := json.Marshal(result)
				if err != nil {
					t.Fatal(err)
				}
				if len(encoded) > 4096 {
					t.Fatalf("expanded aliases leaked into response: %d bytes", len(encoded))
				}
			})
		}
	}
}

func FuzzYAMLObject(f *testing.F) {
	for _, seed := range []string{
		"questions: [{variable: replicas, type: int, default: 3}]",
		"base: &base {value: example}\ncopy: {<<: *base}",
		"cycle: &cycle [*cycle]",
		"duplicate: 1\nduplicate: 2",
		"value: .nan",
		"value: !!binary aGVsbG8=",
		"value: 18446744073709551617",
	} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, document string) {
		if len(document) > maxMetadataBytes {
			t.Skip()
		}
		value, err := yamlObject([]byte(document))
		if err != nil {
			return
		}
		encoded, err := json.Marshal(value)
		if err != nil {
			t.Fatalf("accepted YAML could not be encoded: %v", err)
		}
		// JSON escaping can multiply scalar bytes by six; leave room for bounded containers.
		if len(encoded) > 6*maxYAMLScalarBytes+(1<<20) {
			t.Fatalf("accepted YAML exceeded the expanded encoding bound: %d bytes", len(encoded))
		}
	})
}

type archiveEntry struct {
	name, body string
	kind       byte
}

func chartArchive(t *testing.T, entries ...archiveEntry) []byte {
	t.Helper()
	var data bytes.Buffer
	gz := gzip.NewWriter(&data)
	tw := tar.NewWriter(gz)
	for _, entry := range entries {
		kind := entry.kind
		if kind == 0 {
			kind = tar.TypeReg
		}
		header := &tar.Header{Name: entry.name, Mode: 0600, Size: int64(len(entry.body)), Typeflag: kind}
		if kind != tar.TypeReg {
			header.Size = 0
			header.Linkname = "elsewhere"
		}
		if err := tw.WriteHeader(header); err != nil {
			t.Fatal(err)
		}
		if kind == tar.TypeReg {
			if _, err := tw.Write([]byte(entry.body)); err != nil {
				t.Fatal(err)
			}
		}
	}
	if err := tw.Close(); err != nil {
		t.Fatal(err)
	}
	if err := gz.Close(); err != nil {
		t.Fatal(err)
	}
	return data.Bytes()
}

func TestParseArchiveReturnsMetadataWithoutQuestions(t *testing.T) {
	got, err := ParseArchive(chartArchive(t, archiveEntry{name: "demo/Chart.yaml", body: "apiVersion: v2\nname: example\nversion: 1.2.3\n"}, archiveEntry{name: "demo/values.yaml", body: "replicas: 3\n"}))
	if err != nil {
		t.Fatal(err)
	}
	if got == nil || got.ChartName != "example" || got.ChartVersion != "1.2.3" || got.HasQuestions || got.QuestionsStatus != "absent" {
		t.Fatalf("metadata = %#v", got)
	}
	encoded, err := json.Marshal(got)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(encoded, []byte("defaultValues")) {
		t.Fatalf("default values leaked: %s", encoded)
	}
}

func TestParseArchiveNormalizesEitherQuestionsFilename(t *testing.T) {
	for _, filename := range []string{"questions.yaml", "questions.yml"} {
		t.Run(filename, func(t *testing.T) {
			got, err := ParseArchive(chartArchive(t,
				archiveEntry{name: "demo/Chart.yaml", body: "name: example\nversion: 1.2.3\n"},
				archiveEntry{name: "demo/values.yaml", body: "replicas: 3\n"},
				archiveEntry{name: "demo/" + filename, body: "groups: [{name: Compute, label: Compute settings}]\nquestions: [{variable: replicas, type: int, default: 4, group: Compute}]\n"},
				archiveEntry{name: "demo/charts/child/questions.yaml", body: "questions: [\n"},
				archiveEntry{name: "demo/charts/child/questions.yml", body: "questions: [\n"},
			))
			if err != nil {
				t.Fatal(err)
			}
			if got.ChartName != "example" || got.ChartVersion != "1.2.3" || !got.HasQuestions || got.QuestionsStatus != QuestionsUsable || len(got.Questions) != 1 || len(got.Groups) != 1 {
				t.Fatalf("questions metadata = %#v", got)
			}
			question := got.Questions[0]
			if question["variable"] != "replicas" || question["label"] != "replicas" || question["type"] != "int" || question["default"] != 4 || question["group"] != "compute" {
				t.Fatalf("normalized question = %#v", question)
			}
			if got.Groups[0]["name"] != "compute" || got.Groups[0]["label"] != "Compute settings" || !hasWarning(got, "LABEL_DEFAULTED") {
				t.Fatalf("normalized groups or warnings = %#v / %#v", got.Groups, got.Warnings)
			}
		})
	}
}

func TestParseArchiveUnusableQuestionsWarningsNameTheirFile(t *testing.T) {
	for _, filename := range []string{"questions.yaml", "questions.yml"} {
		for _, tc := range []struct{ name, body string }{
			{"malformed", "questions: [\n"},
			{"root array", "- variable: replicas\n"},
			{"empty", ""},
			{"invalid form", "groups: [{name: Compute}]\nquestions: [{variable: replicas, type: int, optional_hint: ignored}, {type: string}]\n"},
		} {
			t.Run(filename+"/"+tc.name, func(t *testing.T) {
				got, err := ParseArchive(chartArchive(t,
					archiveEntry{name: "demo/Chart.yaml", body: "name: example\nversion: 1.2.3\n"},
					archiveEntry{name: "demo/" + filename, body: tc.body},
				))
				if err != nil {
					t.Fatal(err)
				}
				if got.ChartName != "example" || got.ChartVersion != "1.2.3" || !got.HasQuestions || got.QuestionsStatus != QuestionsUnusable || len(got.Questions) != 0 || len(got.Groups) != 0 {
					t.Fatalf("invalid declaration was salvaged or metadata lost: %#v", got)
				}
				found := false
				for _, warning := range got.Warnings {
					if warning.Code == "UNUSABLE_QUESTIONS" {
						found = true
						if warning.Path != "$" || !strings.Contains(warning.Message, filename) {
							t.Fatalf("unusable warning lost its source filename: %#v", warning)
						}
					}
				}
				if !found {
					t.Fatalf("missing unusable warning: %#v", got.Warnings)
				}
			})
		}
	}
}

func TestParseArchiveRejectsAmbiguousQuestionsFilenames(t *testing.T) {
	for _, first := range []string{"questions.yaml", "questions.yml"} {
		t.Run(first+" first", func(t *testing.T) {
			second := "questions.yaml"
			if first == second {
				second = "questions.yml"
			}
			got, err := ParseArchive(chartArchive(t,
				archiveEntry{name: "demo/Chart.yaml", body: "name: demo\nversion: 1.0.0\n"},
				archiveEntry{name: "demo/" + first, body: "questions: []\n"},
				archiveEntry{name: "demo/" + second, body: "questions: [\n"},
			))
			if err == nil || got != nil {
				t.Fatalf("ambiguous declarations accepted: result=%#v error=%v", got, err)
			}
		})
	}
}

func TestParseArchiveRejectsUnsafePackages(t *testing.T) {
	chart := archiveEntry{name: "demo/Chart.yaml", body: "name: demo\nversion: 1.0.0\n"}
	for _, tc := range []struct {
		name    string
		entries []archiveEntry
	}{
		{"traversal", []archiveEntry{chart, {name: "demo/../outside", body: "x"}}},
		{"absolute", []archiveEntry{chart, {name: "/outside", body: "x"}}},
		{"backslash traversal", []archiveEntry{chart, {name: "demo\\..\\outside", body: "x"}}},
		{"symlink", []archiveEntry{chart, {name: "demo/link", kind: tar.TypeSymlink}}},
		{"hardlink", []archiveEntry{chart, {name: "demo/link", kind: tar.TypeLink}}},
		{"duplicate metadata", []archiveEntry{chart, chart}},
		{"duplicate questions.yml", []archiveEntry{chart, {name: "demo/questions.yml", body: "questions: []\n"}, {name: "demo/questions.yml", body: "questions: []\n"}}},
		{"two chart roots", []archiveEntry{chart, {name: "other/Chart.yaml", body: "name: other\nversion: 1.0.0\n"}}},
		{"missing metadata", []archiveEntry{{name: "demo/values.yaml", body: "a: 1\n"}}},
		{"metadata scalar", []archiveEntry{{name: "demo/Chart.yaml", body: "plain"}}},
		{"duplicate YAML keys", []archiveEntry{{name: "demo/Chart.yaml", body: "name: demo\nname: other\nversion: 1.0.0\n"}}},
		{"oversized values", []archiveEntry{chart, {name: "demo/values.yaml", body: "value: " + string(bytes.Repeat([]byte("x"), (1<<20)+1))}}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := ParseArchive(chartArchive(t, tc.entries...)); err == nil {
				t.Fatal("unsafe chart accepted")
			}
		})
	}
}

func TestParseArchiveRejectsCorruptGzip(t *testing.T) {
	data := chartArchive(t, archiveEntry{name: "demo/Chart.yaml", body: "name: demo\nversion: 1.0.0\n"})
	data[len(data)-8] ^= 0xff
	if _, err := ParseArchive(data); err == nil {
		t.Fatal("corrupt gzip trailer accepted")
	}
}

func TestParseArchiveIgnoresSubchartMetadata(t *testing.T) {
	data := chartArchive(t,
		archiveEntry{name: "demo/Chart.yaml", body: "name: root\nversion: 1.0.0\n"},
		archiveEntry{name: "demo/charts/child/Chart.yaml", body: "name: child\nversion: 2.0.0\n"},
		archiveEntry{name: "demo/charts/child/questions.yaml", body: "questions: [{variable: x, type: string}]\n"},
		archiveEntry{name: "demo/charts/child/questions.yml", body: "questions: [{variable: y, type: string}]\n"},
	)
	got, err := ParseArchive(data)
	if err != nil {
		t.Fatal(err)
	}
	if got == nil || got.ChartName != "root" || got.HasQuestions {
		t.Fatalf("subchart metadata used: %#v", got)
	}
}

func TestParseArchiveEnforcesSizeAndEntryBounds(t *testing.T) {
	chart := archiveEntry{name: "demo/Chart.yaml", body: "name: demo\nversion: 1.0.0\n"}
	t.Run("compressed", func(t *testing.T) {
		if _, err := ParseArchive(make([]byte, (32<<20)+1)); err == nil {
			t.Fatal("compressed bound ignored")
		}
	})
	t.Run("expanded", func(t *testing.T) {
		data := chartArchive(t, chart, archiveEntry{name: "demo/data", body: strings.Repeat("x", 65<<20)})
		if _, err := ParseArchive(data); err == nil {
			t.Fatal("expanded bound ignored")
		}
	})
	t.Run("entries", func(t *testing.T) {
		entries := []archiveEntry{chart}
		for n := 0; n < 4096; n++ {
			entries = append(entries, archiveEntry{name: "demo/unused"})
		}
		if _, err := ParseArchive(chartArchive(t, entries...)); err == nil {
			t.Fatal("entry bound ignored")
		}
	})
	for _, filename := range []string{"questions.yaml", "questions.yml"} {
		t.Run(filename, func(t *testing.T) {
			data := chartArchive(t, chart, archiveEntry{name: "demo/" + filename, body: strings.Repeat("x", (1<<20)+1)})
			if _, err := ParseArchive(data); err == nil {
				t.Fatal("questions bound ignored")
			}
		})
	}
}

func TestParseArchiveRejectsHiddenArchiveSuffixes(t *testing.T) {
	first := chartArchive(t, archiveEntry{name: "demo/Chart.yaml", body: "name: demo\nversion: 1.0.0\n"})
	second := chartArchive(t, archiveEntry{name: "other/Chart.yaml", body: "name: other\nversion: 2.0.0\n"})
	t.Run("gzip members", func(t *testing.T) {
		if _, err := ParseArchive(append(first, second...)); err == nil {
			t.Fatal("concatenated gzip charts accepted")
		}
	})
	t.Run("tar suffix", func(t *testing.T) {
		var raw bytes.Buffer
		for _, data := range [][]byte{first, second} {
			gz, err := gzip.NewReader(bytes.NewReader(data))
			if err != nil {
				t.Fatal(err)
			}
			if _, err := raw.ReadFrom(gz); err != nil {
				t.Fatal(err)
			}
			if err := gz.Close(); err != nil {
				t.Fatal(err)
			}
		}
		var compressed bytes.Buffer
		gz := gzip.NewWriter(&compressed)
		if _, err := gz.Write(raw.Bytes()); err != nil {
			t.Fatal(err)
		}
		if err := gz.Close(); err != nil {
			t.Fatal(err)
		}
		if _, err := ParseArchive(compressed.Bytes()); err == nil {
			t.Fatal("hidden second tar chart accepted")
		}
	})
}
