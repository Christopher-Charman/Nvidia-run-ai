package metrics

type RunaiAvgAllocatedGpuCountPerDepartmentNodePoolBottom5Filters struct {
	ClusterId  []string `mapstructure:"clusterId,omitempty"`
	Department []string `mapstructure:"department,omitempty"`
	NodePool   []string `mapstructure:"nodepool,omitempty"`
}

type RunaiAvgAllocatedGpuCountPerDepartmentNodePoolBottom5Labels struct {
	ClusterId  bool `mapstructure:"clusterId,omitempty"`
	NodePool   bool `mapstructure:"nodepool,omitempty"`
	Department bool `mapstructure:"department,omitempty"`
}

type RunaiAvgAllocatedGpuCountPerDepartmentNodePoolBottom5 struct {
	Filters        RunaiAvgAllocatedGpuCountPerDepartmentNodePoolBottom5Filters `mapstructure:"filters,omitempty"`
	ResponseLabels RunaiAvgAllocatedGpuCountPerDepartmentNodePoolBottom5Labels  `mapstructure:"labels,omitempty"`
	Quantile       *float32
}

func (m RunaiAvgAllocatedGpuCountPerDepartmentNodePoolBottom5) GetName() string {
	return "runai_avg_allocated_gpu_count_per_department_nodepool_bottom5"
}

func (m RunaiAvgAllocatedGpuCountPerDepartmentNodePoolBottom5) GetFilters() interface{} {
	return m.Filters
}

func (m RunaiAvgAllocatedGpuCountPerDepartmentNodePoolBottom5) GetResponseLabels() interface{} {
	return m.ResponseLabels
}

func (m RunaiAvgAllocatedGpuCountPerDepartmentNodePoolBottom5) GetQuantile() *float32 {
	return m.Quantile
}
func (m RunaiAvgAllocatedGpuCountPerDepartmentNodePoolBottom5) IsAggregation() bool {
	return true
}

func (m RunaiAvgAllocatedGpuCountPerDepartmentNodePoolBottom5) IsQuantile() bool {
	return false
}
