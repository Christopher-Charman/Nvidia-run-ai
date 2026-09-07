package metrics

type RunaiAvgGpuUtilizationPerDepartmentNodePoolTop5Filters struct {
	ClusterId  []string `mapstructure:"clusterId,omitempty"`
	Department []string `mapstructure:"department,omitempty"`
	NodePool   []string `mapstructure:"nodepool,omitempty"`
}

type RunaiAvgGpuUtilizationPerDepartmentNodePoolTop5Labels struct {
	ClusterId  bool `mapstructure:"clusterId,omitempty"`
	NodePool   bool `mapstructure:"nodepool,omitempty"`
	Department bool `mapstructure:"department,omitempty"`
}

type RunaiAvgGpuUtilizationPerDepartmentNodePoolTop5 struct {
	Filters        RunaiAvgGpuUtilizationPerDepartmentNodePoolTop5Filters `mapstructure:"filters,omitempty"`
	ResponseLabels RunaiAvgGpuUtilizationPerDepartmentNodePoolTop5Labels  `mapstructure:"labels,omitempty"`
	Quantile       *float32
}

func (m RunaiAvgGpuUtilizationPerDepartmentNodePoolTop5) GetName() string {
	return "runai_avg_gpu_utilization_per_department_nodepool_top5"
}

func (m RunaiAvgGpuUtilizationPerDepartmentNodePoolTop5) GetFilters() interface{} {
	return m.Filters
}

func (m RunaiAvgGpuUtilizationPerDepartmentNodePoolTop5) GetResponseLabels() interface{} {
	return m.ResponseLabels
}

func (m RunaiAvgGpuUtilizationPerDepartmentNodePoolTop5) GetQuantile() *float32 {
	return m.Quantile
}
func (m RunaiAvgGpuUtilizationPerDepartmentNodePoolTop5) IsAggregation() bool {
	return true
}

func (m RunaiAvgGpuUtilizationPerDepartmentNodePoolTop5) IsQuantile() bool {
	return false
}
