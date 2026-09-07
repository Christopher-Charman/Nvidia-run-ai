package metrics

import (
	"github.com/google/uuid"
)

type RunaiMemoryUtilizatioPerTenantFilters struct {
	ClusterId []uuid.UUID `mapstructure:"clusterId,omitempty"`
}

type RunaiMemoryUtilizatioPerTenantResponseLabels struct{}

type RunaiMemoryUtilizatioPerTenantMetric struct {
	Filters        RunaiMemoryUtilizatioPerTenantFilters
	ResponseLabels RunaiMemoryUtilizatioPerTenantResponseLabels
	Quantile       *float32
}

func (m RunaiMemoryUtilizatioPerTenantMetric) GetName() string {
	return "runai_memory_utilization_per_tenant"
}

func (m RunaiMemoryUtilizatioPerTenantMetric) GetFilters() interface{} {
	return m.Filters
}

func (m RunaiMemoryUtilizatioPerTenantMetric) GetResponseLabels() interface{} {
	return m.ResponseLabels
}

func (m RunaiMemoryUtilizatioPerTenantMetric) GetQuantile() *float32 {
	return m.Quantile
}
func (m RunaiMemoryUtilizatioPerTenantMetric) IsAggregation() bool {
	return false
}

func (m RunaiMemoryUtilizatioPerTenantMetric) IsQuantile() bool {
	return false
}
