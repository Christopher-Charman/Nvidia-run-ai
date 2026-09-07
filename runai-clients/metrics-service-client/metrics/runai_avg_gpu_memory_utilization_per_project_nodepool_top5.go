package metrics

type RunaiAvgGpuMemoryUtilizationPerProjectNodePoolTop5Filters struct {
	ClusterId  []string `mapstructure:"clusterId,omitempty"`
	Department []string `mapstructure:"department,omitempty"`
	Project    []string `mapstructure:"project,omitempty"`
	NodePool   []string `mapstructure:"nodepool,omitempty"`
}

type RunaiAvgGpuMemoryUtilizationPerProjectNodePoolTop5Labels struct {
	ClusterId  bool `mapstructure:"clusterId,omitempty"`
	NodePool   bool `mapstructure:"nodepool,omitempty"`
	Project    bool `mapstructure:"project,omitempty"`
	Department bool `mapstructure:"department,omitempty"`
}

type RunaiAvgGpuMemoryUtilizationPerProjectNodePoolTop5 struct {
	Filters        RunaiAvgGpuMemoryUtilizationPerProjectNodePoolTop5Filters
	ResponseLabels RunaiAvgGpuMemoryUtilizationPerProjectNodePoolTop5Labels
	Quantile       *float32
}

func (m RunaiAvgGpuMemoryUtilizationPerProjectNodePoolTop5) GetName() string {
	return "runai_avg_gpu_memory_utilization_per_project_nodepool_top5"
}

func (m RunaiAvgGpuMemoryUtilizationPerProjectNodePoolTop5) GetFilters() interface{} {
	return m.Filters
}

func (m RunaiAvgGpuMemoryUtilizationPerProjectNodePoolTop5) GetResponseLabels() interface{} {
	return m.ResponseLabels
}

func (m RunaiAvgGpuMemoryUtilizationPerProjectNodePoolTop5) GetQuantile() *float32 {
	return m.Quantile
}
func (m RunaiAvgGpuMemoryUtilizationPerProjectNodePoolTop5) IsAggregation() bool {
	return true
}

func (m RunaiAvgGpuMemoryUtilizationPerProjectNodePoolTop5) IsQuantile() bool {
	return false
}
