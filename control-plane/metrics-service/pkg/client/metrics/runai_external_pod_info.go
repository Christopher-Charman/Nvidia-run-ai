package metrics

import (
	"github.com/google/uuid"
)

type RunaiExternalRunningPodInfoFilters struct {
	ClusterId      []uuid.UUID `mapstructure:"clusterId,omitempty"`
	PodId          []string    `mapstructure:"pod_uuid,omitempty"`
	WorkloadId     []string    `mapstructure:"workload_id,omitempty"`
	DetailedStatus []string    `mapstructure:"detailed_status,omitempty"`
	Phase          []string    `mapstructure:"phase,omitempty"`
	Node           []string    `mapstructure:"node,omitempty"`
}

type RunaiExternalRunningPodInfoResponseLabels struct {
	ClusterId      bool `mapstructure:"clusterId,omitempty"`
	PodId          bool `mapstructure:"pod_uuid,omitempty"`
	WorkloadId     bool `mapstructure:"workload_id,omitempty"`
	DetailedStatus bool `mapstructure:"detailed_status,omitempty"`
	Phase          bool `mapstructure:"phase,omitempty"`
	Node           bool `mapstructure:"node,omitempty"`
}

type RunaiExternalRunningPodInfoMetric struct {
	Filters        RunaiExternalRunningPodInfoFilters
	ResponseLabels RunaiExternalRunningPodInfoResponseLabels
	Quantile       *float32
}

func (m RunaiExternalRunningPodInfoMetric) GetName() string {
	return "runai_external_running_pod_info"
}

func (m RunaiExternalRunningPodInfoMetric) GetFilters() interface{} {
	return m.Filters
}

func (m RunaiExternalRunningPodInfoMetric) GetResponseLabels() interface{} {
	return m.ResponseLabels
}

func (m RunaiExternalRunningPodInfoMetric) GetQuantile() *float32 {
	return m.Quantile
}

func (m RunaiExternalRunningPodInfoMetric) IsAggregation() bool {
	return false
}

func (m RunaiExternalRunningPodInfoMetric) IsQuantile() bool {
	return false
}
