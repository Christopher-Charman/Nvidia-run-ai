// SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package chart_inspection

const (
	QuestionsAbsent   = "absent"
	QuestionsUsable   = "usable"
	QuestionsUnusable = "unusable"
)

type Warning struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Path    string `json:"path"`
}

type Result struct {
	ChartName       string           `json:"chartName"`
	ChartVersion    string           `json:"chartVersion"`
	ArchiveSHA256   string           `json:"archiveSHA256"`
	HasQuestions    bool             `json:"hasQuestions"`
	QuestionsStatus string           `json:"questionsStatus"`
	Groups          []map[string]any `json:"groups"`
	Questions       []map[string]any `json:"questions"`
	Warnings        []Warning        `json:"warnings"`
}
