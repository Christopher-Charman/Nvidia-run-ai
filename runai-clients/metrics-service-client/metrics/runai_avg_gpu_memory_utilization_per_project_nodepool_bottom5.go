package metrics

type RunaiAvgGpuMemoryUtilizationPerProjectNodePoolBottom5Filters struct {
	ClusterId  []string `mapstructure:"clusterId,omitempty"`
	Department []string `mapstructure:"department,omitempty"`
	Project    []string `mapstructure:"project,omitempty"`
	NodePool   []string `mapstructure:"nodepool,omitempty"`
}

type RunaiAvgGpuMemoryUtilizationPerProjectNodePoolBottom5Labels struct {
	ClusterId  bool `mapstructure:"clusterId,omitempty"`
	NodePool   bool `mapstructure:"nodepool,omitempty"`
	Project    bool `mapstructure:"project,omitempty"`
	Department bool `mapstructure:"department,omitempty"`
}

type RunaiAvgGpuMemoryUtilizationPerProjectNodePoolBottom5 struct {
	Filters        RunaiAvgGpuMemoryUtilizationPerProjectNodePoolBottom5Filters
	ResponseLabels RunaiAvgGpuMemoryUtilizationPerProjectNodePoolBottom5Labels
	Quantile       *float32
}

func (m RunaiAvgGpuMemoryUtilizationPerProjectNodePoolBottom5) GetName() string {
	return "runai_avg_gpu_memory_utilization_per_project_nodepool_bottom5"
}

func (m RunaiAvgGpuMemoryUtilizationPerProjectNodePoolBottom5) GetFilters() interface{} {
	return m.Filters
}

func (m RunaiAvgGpuMemoryUtilizationPerProjectNodePoolBottom5) GetResponseLabels() interface{} {
	return m.ResponseLabels
}

func (m RunaiAvgGpuMemoryUtilizationPerProjectNodePoolBottom5) GetQuantile() *float32 {
	return m.Quantile
}
func (m RunaiAvgGpuMemoryUtilizationPerProjectNodePoolBottom5) IsAggregation() bool {
	return true
}

func (m RunaiAvgGpuMemoryUtilizationPerProjectNodePoolBottom5) IsQuantile() bool {
	return false
}
