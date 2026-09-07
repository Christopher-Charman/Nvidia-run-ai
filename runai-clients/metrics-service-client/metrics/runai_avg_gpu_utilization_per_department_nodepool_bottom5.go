package metrics

type RunaiAvgGpuUtilizationPerDepartmentNodePoolBottom5Filters struct {
	ClusterId  []string `mapstructure:"clusterId,omitempty"`
	Department []string `mapstructure:"department,omitempty"`
	NodePool   []string `mapstructure:"nodepool,omitempty"`
}

type RunaiAvgGpuUtilizationPerDepartmentNodePoolBottom5Labels struct {
	ClusterId  bool `mapstructure:"clusterId,omitempty"`
	NodePool   bool `mapstructure:"nodepool,omitempty"`
	Department bool `mapstructure:"department,omitempty"`
}

type RunaiAvgGpuUtilizationPerDepartmentNodePoolBottom5 struct {
	Filters        RunaiAvgGpuUtilizationPerDepartmentNodePoolBottom5Filters `mapstructure:"filters,omitempty"`
	ResponseLabels RunaiAvgGpuUtilizationPerDepartmentNodePoolBottom5Labels  `mapstructure:"labels,omitempty"`
	Quantile       *float32
}

func (m RunaiAvgGpuUtilizationPerDepartmentNodePoolBottom5) GetName() string {
	return "runai_avg_gpu_utilization_per_department_nodepool_bottom5"
}

func (m RunaiAvgGpuUtilizationPerDepartmentNodePoolBottom5) GetFilters() interface{} {
	return m.Filters
}

func (m RunaiAvgGpuUtilizationPerDepartmentNodePoolBottom5) GetResponseLabels() interface{} {
	return m.ResponseLabels
}

func (m RunaiAvgGpuUtilizationPerDepartmentNodePoolBottom5) GetQuantile() *float32 {
	return m.Quantile
}
func (m RunaiAvgGpuUtilizationPerDepartmentNodePoolBottom5) IsAggregation() bool {
	return true
}

func (m RunaiAvgGpuUtilizationPerDepartmentNodePoolBottom5) IsQuantile() bool {
	return false
}
