package metrics

import (
	"github.com/google/uuid"
)

type RunaiGpuUtilizationPerTenantFilters struct {
	ClusterId []uuid.UUID `mapstructure:"clusterId,omitempty"`
}

type RunaiGpuUtilizationPerTenantResponseLabels struct{}

type RunaiGpuUtilizationPerTenantMetric struct {
	Filters        RunaiGpuUtilizationPerTenantFilters
	ResponseLabels RunaiGpuUtilizationPerTenantResponseLabels
	Quantile       *float32
}

func (m RunaiGpuUtilizationPerTenantMetric) GetName() string {
	return "runai_gpu_utilization_per_tenant"
}

func (m RunaiGpuUtilizationPerTenantMetric) GetFilters() interface{} {
	return m.Filters
}

func (m RunaiGpuUtilizationPerTenantMetric) GetResponseLabels() interface{} {
	return m.ResponseLabels
}

func (m RunaiGpuUtilizationPerTenantMetric) GetQuantile() *float32 {
	return m.Quantile
}
func (m RunaiGpuUtilizationPerTenantMetric) IsAggregation() bool {
	return false
}

func (m RunaiGpuUtilizationPerTenantMetric) IsQuantile() bool {
	return false
}
