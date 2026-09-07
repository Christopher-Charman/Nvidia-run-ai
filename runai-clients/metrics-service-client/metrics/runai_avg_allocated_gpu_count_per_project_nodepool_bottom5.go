package metrics

type RunaiAvgAllocatedGpuCountPerProjectNodePoolBottom5Filters struct {
	ClusterId  []string `mapstructure:"clusterId,omitempty"`
	Department []string `mapstructure:"department,omitempty"`
	Project    []string `mapstructure:"project,omitempty"`
	NodePool   []string `mapstructure:"nodepool,omitempty"`
}

type RunaiAvgAllocatedGpuCountPerProjectNodePoolBottom5Labels struct {
	ClusterId  bool `mapstructure:"clusterId,omitempty"`
	NodePool   bool `mapstructure:"nodepool,omitempty"`
	Project    bool `mapstructure:"project,omitempty"`
	Department bool `mapstructure:"department,omitempty"`
}

type RunaiAvgAllocatedGpuCountPerProjectNodePoolBottom5 struct {
	Filters        RunaiAvgAllocatedGpuCountPerProjectNodePoolBottom5Filters
	ResponseLabels RunaiAvgAllocatedGpuCountPerProjectNodePoolBottom5Labels
	Quantile       *float32
}

func (m RunaiAvgAllocatedGpuCountPerProjectNodePoolBottom5) GetName() string {
	return "runai_avg_allocated_gpu_count_per_project_nodepool_bottom5"
}

func (m RunaiAvgAllocatedGpuCountPerProjectNodePoolBottom5) GetFilters() interface{} {
	return m.Filters
}

func (m RunaiAvgAllocatedGpuCountPerProjectNodePoolBottom5) GetResponseLabels() interface{} {
	return m.ResponseLabels
}

func (m RunaiAvgAllocatedGpuCountPerProjectNodePoolBottom5) GetQuantile() *float32 {
	return m.Quantile
}
func (m RunaiAvgAllocatedGpuCountPerProjectNodePoolBottom5) IsAggregation() bool {
	return true
}

func (m RunaiAvgAllocatedGpuCountPerProjectNodePoolBottom5) IsQuantile() bool {
	return false
}
