/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 NVIDIA CORPORATION & AFFILIATES. All rights reserved.
 * SPDX-License-Identifier: Apache-2.0
 */

package models

import (
	"github.com/google/uuid"

	"github.com/dsx-ai-factory/run-ai/runai-application-packages/queryapi"
)

// WorkloadSlim is a workload as returned by GET /api/v1/workloads-slim.
// It is read from the workloads table rather than the workloads view, so it has no pod aggregates
// and is never stale.
type WorkloadSlim struct {
	ID                   uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name                 string
	Type                 string
	ProjectName          string
	ClusterID            uuid.UUID
	DistributedFramework string
}

// FilterFieldsToColumn returns the fields GET /api/v1/workloads-slim can be filtered by.
func (WorkloadSlim) FilterFieldsToColumn() queryapi.APIFieldsMapping {
	return queryapi.APIFieldsMapping{
		"type":                 queryapi.NewFieldMapping("type"),
		"name":                 queryapi.NewFieldMapping("name"),
		"clusterId":            queryapi.NewFieldMapping("cluster_id"),
		"projectName":          queryapi.NewFieldMapping("project_name"),
		"distributedFramework": queryapi.NewFieldMapping("distributed_framework"),
	}
}
