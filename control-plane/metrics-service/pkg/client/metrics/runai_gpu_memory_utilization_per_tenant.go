package metrics

import (
	"github.com/google/uuid"
)

type RunaiGpuMemoryUtilizationPerTenantFilters struct {
	ClusterId []uuid.UUID `mapstructure:"clusterId,omitempty"`
}

type RunaiGpuMemoryUtilizationPerTenantResponseLabels struct{}

type RunaiGpuMemoryUtilizationPerTenantMetric struct {
	Filters        RunaiGpuMemoryUtilizationPerTenantFilters
	ResponseLabels RunaiGpuMemoryUtilizationPerTenantResponseLabels
	Quantile       *float32
}

func (m RunaiGpuMemoryUtilizationPerTenantMetric) GetName() string {
	return "runai_gpu_memory_utilization_per_tenant"
}

func (m RunaiGpuMemoryUtilizationPerTenantMetric) GetFilters() interface{} {
	return m.Filters
}

func (m RunaiGpuMemoryUtilizationPerTenantMetric) GetResponseLabels() interface{} {
	return m.ResponseLabels
}

func (m RunaiGpuMemoryUtilizationPerTenantMetric) GetQuantile() *float32 {
	return m.Quantile
}
func (m RunaiGpuMemoryUtilizationPerTenantMetric) IsAggregation() bool {
	return false
}

func (m RunaiGpuMemoryUtilizationPerTenantMetric) IsQuantile() bool {
	return false
}
