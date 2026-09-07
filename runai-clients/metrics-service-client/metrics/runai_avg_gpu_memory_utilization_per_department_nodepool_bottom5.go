package metrics

type RunaiAvgGpuMemoryUtilizationPerDepartmentNodePoolBottom5Filters struct {
	ClusterId  []string `mapstructure:"clusterId,omitempty"`
	Department []string `mapstructure:"department,omitempty"`
	NodePool   []string `mapstructure:"nodepool,omitempty"`
}

type RunaiAvgGpuMemoryUtilizationPerDepartmentNodePoolBottom5Labels struct {
	ClusterId  bool `mapstructure:"clusterId,omitempty"`
	NodePool   bool `mapstructure:"nodepool,omitempty"`
	Department bool `mapstructure:"department,omitempty"`
}

type RunaiAvgGpuMemoryUtilizationPerDepartmentNodePoolBottom5 struct {
	Filters        RunaiAvgGpuMemoryUtilizationPerDepartmentNodePoolBottom5Filters `mapstructure:"filters,omitempty"`
	ResponseLabels RunaiAvgGpuMemoryUtilizationPerDepartmentNodePoolBottom5Labels  `mapstructure:"labels,omitempty"`
	Quantile       *float32
}

func (m RunaiAvgGpuMemoryUtilizationPerDepartmentNodePoolBottom5) GetName() string {
	return "runai_avg_gpu_memory_utilization_per_department_nodepool_bottom5"
}

func (m RunaiAvgGpuMemoryUtilizationPerDepartmentNodePoolBottom5) GetFilters() interface{} {
	return m.Filters
}

func (m RunaiAvgGpuMemoryUtilizationPerDepartmentNodePoolBottom5) GetResponseLabels() interface{} {
	return m.ResponseLabels
}

func (m RunaiAvgGpuMemoryUtilizationPerDepartmentNodePoolBottom5) GetQuantile() *float32 {
	return m.Quantile
}
func (m RunaiAvgGpuMemoryUtilizationPerDepartmentNodePoolBottom5) IsAggregation() bool {
	return true
}

func (m RunaiAvgGpuMemoryUtilizationPerDepartmentNodePoolBottom5) IsQuantile() bool {
	return false
}
