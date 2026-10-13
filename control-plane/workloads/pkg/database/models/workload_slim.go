/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package models

import (
	"github.com/google/uuid"

	"github.com/dsx-ai-factory/run-ai/runai-application-packages/queryapi"
)

// WorkloadSlim is a workload as returned by GET /api/v1/workloads?verbosity=slim.
// It is read from the workloads table rather than the workloads view, so it has no pod aggregates
// and is never stale.
type WorkloadSlim struct {
	ID                   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name                 string
	Type                 string
	ProjectID            string
	ProjectName          string
	ClusterID            uuid.UUID
	ExternalWorkloadName string
}

// FilterFieldsToColumn returns the filterable fields that are stored on the workloads table.
func (WorkloadSlim) FilterFieldsToColumn() queryapi.APIFieldsMapping {
	viewColumns := WorkloadView{}.FilterFieldsToColumn()
	tableColumns := queryapi.APIFieldsMapping{}
	for field := range (Workload{}).filterableColumns() {
		tableColumns[field] = viewColumns[field]
	}
	return tableColumns
}

func (ws WorkloadSlim) SortFieldsToColumn() queryapi.APIFieldsMapping {
	mapping := ws.FilterFieldsToColumn()
	mapping["phase"] = queryapi.NewFieldMapping("phase %s, phase_updated_at %s")
	return mapping
}
