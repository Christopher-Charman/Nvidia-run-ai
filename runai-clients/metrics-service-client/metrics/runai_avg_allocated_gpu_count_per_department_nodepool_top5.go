package metrics

type RunaiAvgAllocatedGpuCountPerDepartmentNodePoolTop5Filters struct {
	ClusterId  []string `mapstructure:"clusterId,omitempty"`
	Department []string `mapstructure:"department,omitempty"`
	NodePool   []string `mapstructure:"nodepool,omitempty"`
}

type RunaiAvgAllocatedGpuCountPerDepartmentNodePoolTop5Labels struct {
	ClusterId  bool `mapstructure:"clusterId,omitempty"`
	NodePool   bool `mapstructure:"nodepool,omitempty"`
	Department bool `mapstructure:"department,omitempty"`
}

type RunaiAvgAllocatedGpuCountPerDepartmentNodePoolTop5 struct {
	Filters        RunaiAvgAllocatedGpuCountPerDepartmentNodePoolTop5Filters `mapstructure:"filters,omitempty"`
	ResponseLabels RunaiAvgAllocatedGpuCountPerDepartmentNodePoolTop5Labels  `mapstructure:"labels,omitempty"`
	Quantile       *float32
}

func (m RunaiAvgAllocatedGpuCountPerDepartmentNodePoolTop5) GetName() string {
	return "runai_avg_allocated_gpu_count_per_department_nodepool_top5"
}

func (m RunaiAvgAllocatedGpuCountPerDepartmentNodePoolTop5) GetFilters() interface{} {
	return m.Filters
}

func (m RunaiAvgAllocatedGpuCountPerDepartmentNodePoolTop5) GetResponseLabels() interface{} {
	return m.ResponseLabels
}

func (m RunaiAvgAllocatedGpuCountPerDepartmentNodePoolTop5) GetQuantile() *float32 {
	return m.Quantile
}
func (m RunaiAvgAllocatedGpuCountPerDepartmentNodePoolTop5) IsAggregation() bool {
	return true
}

func (m RunaiAvgAllocatedGpuCountPerDepartmentNodePoolTop5) IsQuantile() bool {
	return false
}
