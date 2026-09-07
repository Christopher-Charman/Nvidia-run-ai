package metrics

type RunaiAvgAllocatedGpuCountPerProjectNodePoolTop5Filters struct {
	ClusterId  []string `mapstructure:"clusterId,omitempty"`
	Department []string `mapstructure:"department,omitempty"`
	Project    []string `mapstructure:"project,omitempty"`
	NodePool   []string `mapstructure:"nodepool,omitempty"`
}

type RunaiAvgAllocatedGpuCountPerProjectNodePoolTop5ResponseLabels struct {
	ClusterId  bool `mapstructure:"clusterId,omitempty"`
	NodePool   bool `mapstructure:"nodepool,omitempty"`
	Project    bool `mapstructure:"project,omitempty"`
	Department bool `mapstructure:"department,omitempty"`
}

type RunaiAvgAllocatedGpuCountPerProjectNodePoolTop5 struct {
	Filters        RunaiAvgAllocatedGpuCountPerProjectNodePoolTop5Filters
	ResponseLabels RunaiAvgAllocatedGpuCountPerProjectNodePoolTop5ResponseLabels
	Quantile       *float32
}

func (m RunaiAvgAllocatedGpuCountPerProjectNodePoolTop5) GetName() string {
	return "runai_avg_allocated_gpu_count_per_project_nodepool_top5"
}

func (m RunaiAvgAllocatedGpuCountPerProjectNodePoolTop5) GetFilters() interface{} {
	return m.Filters
}

func (m RunaiAvgAllocatedGpuCountPerProjectNodePoolTop5) GetResponseLabels() interface{} {
	return m.ResponseLabels
}

func (m RunaiAvgAllocatedGpuCountPerProjectNodePoolTop5) GetQuantile() *float32 {
	return m.Quantile
}
func (m RunaiAvgAllocatedGpuCountPerProjectNodePoolTop5) IsAggregation() bool {
	return true
}

func (m RunaiAvgAllocatedGpuCountPerProjectNodePoolTop5) IsQuantile() bool {
	return false
}
