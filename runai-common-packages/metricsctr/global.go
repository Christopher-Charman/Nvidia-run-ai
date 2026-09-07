package metricsctr

import (
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/metrics"
	ctrlmetrics "sigs.k8s.io/controller-runtime/pkg/metrics"
)

func Init() {
	ctrlmetrics.Registry = metrics.GetPrometheusRegistry()
}
