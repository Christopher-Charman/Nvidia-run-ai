package metrics

type RunaiAvgGpuUtilizationPerProjectNodePoolTop5Filters struct {
	ClusterId  []string `mapstructure:"clusterId,omitempty"`
	Department []string `mapstructure:"department,omitempty"`
	Project    []string `mapstructure:"project,omitempty"`
	NodePool   []string `mapstructure:"nodepool,omitempty"`
}

type RunaiAvgGpuUtilizationPerProjectNodePoolTop5Labels struct {
	ClusterId  bool `mapstructure:"clusterId,omitempty"`
	NodePool   bool `mapstructure:"nodepool,omitempty"`
	Project    bool `mapstructure:"project,omitempty"`
	Department bool `mapstructure:"department,omitempty"`
}

type RunaiAvgGpuUtilizationPerProjectNodePoolTop5 struct {
	Filters        RunaiAvgGpuUtilizationPerProjectNodePoolTop5Filters
	ResponseLabels RunaiAvgGpuUtilizationPerProjectNodePoolTop5Labels
	Quantile       *float32
}

func (m RunaiAvgGpuUtilizationPerProjectNodePoolTop5) GetName() string {
	return "runai_avg_gpu_utilization_per_project_nodepool_top5"
}

func (m RunaiAvgGpuUtilizationPerProjectNodePoolTop5) GetFilters() interface{} {
	return m.Filters
}

func (m RunaiAvgGpuUtilizationPerProjectNodePoolTop5) GetResponseLabels() interface{} {
	return m.ResponseLabels
}

func (m RunaiAvgGpuUtilizationPerProjectNodePoolTop5) GetQuantile() *float32 {
	return m.Quantile
}
func (m RunaiAvgGpuUtilizationPerProjectNodePoolTop5) IsAggregation() bool {
	return true
}

func (m RunaiAvgGpuUtilizationPerProjectNodePoolTop5) IsQuantile() bool {
	return false
}
