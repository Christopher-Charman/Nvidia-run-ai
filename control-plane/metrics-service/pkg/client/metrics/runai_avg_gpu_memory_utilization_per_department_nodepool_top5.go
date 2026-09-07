package metrics

type RunaiAvgGpuMemoryUtilizationPerDepartmentNodePoolTop5Filters struct {
	ClusterId  []string `mapstructure:"clusterId,omitempty"`
	Department []string `mapstructure:"department,omitempty"`
	NodePool   []string `mapstructure:"nodepool,omitempty"`
}

type RunaiAvgGpuMemoryUtilizationPerDepartmentNodePoolTop5Labels struct {
	ClusterId  bool `mapstructure:"clusterId,omitempty"`
	NodePool   bool `mapstructure:"nodepool,omitempty"`
	Department bool `mapstructure:"department,omitempty"`
}

type RunaiAvgGpuMemoryUtilizationPerDepartmentNodePoolTop5 struct {
	Filters        RunaiAvgGpuMemoryUtilizationPerDepartmentNodePoolTop5Filters `mapstructure:"filters,omitempty"`
	ResponseLabels RunaiAvgGpuMemoryUtilizationPerDepartmentNodePoolTop5Labels  `mapstructure:"labels,omitempty"`
	Quantile       *float32
}

func (m RunaiAvgGpuMemoryUtilizationPerDepartmentNodePoolTop5) GetName() string {
	return "runai_avg_gpu_memory_utilization_per_department_nodepool_top5"
}

func (m RunaiAvgGpuMemoryUtilizationPerDepartmentNodePoolTop5) GetFilters() interface{} {
	return m.Filters
}

func (m RunaiAvgGpuMemoryUtilizationPerDepartmentNodePoolTop5) GetResponseLabels() interface{} {
	return m.ResponseLabels
}

func (m RunaiAvgGpuMemoryUtilizationPerDepartmentNodePoolTop5) GetQuantile() *float32 {
	return m.Quantile
}
func (m RunaiAvgGpuMemoryUtilizationPerDepartmentNodePoolTop5) IsAggregation() bool {
	return true
}

func (m RunaiAvgGpuMemoryUtilizationPerDepartmentNodePoolTop5) IsQuantile() bool {
	return false
}
