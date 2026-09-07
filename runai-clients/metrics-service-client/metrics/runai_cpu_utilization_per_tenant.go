package metrics

import (
	"github.com/google/uuid"
)

type RunaiCpuUtilizationPerTenantFilters struct {
	ClusterId []uuid.UUID `mapstructure:"clusterId,omitempty"`
}

type RunaiCpuUtilizationPerTenantResponseLabels struct{}

type RunaiCpuUtilizationPerTenantMetric struct {
	Filters        RunaiCpuUtilizationPerTenantFilters
	ResponseLabels RunaiCpuUtilizationPerTenantResponseLabels
	Quantile       *float32
}

func (m RunaiCpuUtilizationPerTenantMetric) GetName() string {
	return "runai_cpu_utilization_per_tenant"
}

func (m RunaiCpuUtilizationPerTenantMetric) GetFilters() interface{} {
	return m.Filters
}

func (m RunaiCpuUtilizationPerTenantMetric) GetResponseLabels() interface{} {
	return m.ResponseLabels
}

func (m RunaiCpuUtilizationPerTenantMetric) GetQuantile() *float32 {
	return m.Quantile
}
func (m RunaiCpuUtilizationPerTenantMetric) IsAggregation() bool {
	return false
}

func (m RunaiCpuUtilizationPerTenantMetric) IsQuantile() bool {
	return false
}
