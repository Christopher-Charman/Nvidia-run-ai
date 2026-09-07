package metrics

type RunaiAvgGpuUtilizationPerProjectNodePoolBottom5Filters struct {
	ClusterId  []string `mapstructure:"clusterId,omitempty"`
	Department []string `mapstructure:"department,omitempty"`
	Project    []string `mapstructure:"project,omitempty"`
	NodePool   []string `mapstructure:"nodepool,omitempty"`
}

type RunaiAvgGpuUtilizationPerProjectNodePoolBottom5Labels struct {
	ClusterId  bool `mapstructure:"clusterId,omitempty"`
	NodePool   bool `mapstructure:"nodepool,omitempty"`
	Project    bool `mapstructure:"project,omitempty"`
	Department bool `mapstructure:"department,omitempty"`
}

type RunaiAvgGpuUtilizationPerProjectNodePoolBottom5 struct {
	Filters        RunaiAvgGpuUtilizationPerProjectNodePoolBottom5Filters
	ResponseLabels RunaiAvgGpuUtilizationPerProjectNodePoolBottom5Labels
	Quantile       *float32
}

func (m RunaiAvgGpuUtilizationPerProjectNodePoolBottom5) GetName() string {
	return "runai_avg_gpu_utilization_per_project_nodepool_bottom5"
}

func (m RunaiAvgGpuUtilizationPerProjectNodePoolBottom5) GetFilters() interface{} {
	return m.Filters
}

func (m RunaiAvgGpuUtilizationPerProjectNodePoolBottom5) GetResponseLabels() interface{} {
	return m.ResponseLabels
}

func (m RunaiAvgGpuUtilizationPerProjectNodePoolBottom5) GetQuantile() *float32 {
	return m.Quantile
}
func (m RunaiAvgGpuUtilizationPerProjectNodePoolBottom5) IsAggregation() bool {
	return true
}

func (m RunaiAvgGpuUtilizationPerProjectNodePoolBottom5) IsQuantile() bool {
	return false
}
