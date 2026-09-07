package registry

import (
	"testing"

	"github.com/dsx-ai-factory/run-ai/runai-common-packages/metrics/internal/registry/mocks"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/metrics/internal/runaimetrics"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGauge_RegistryTwiceReturnsSame(t *testing.T) {
	prm := mocks.NewRegisterer(t)
	prm.On("MustRegister", mock.Anything)

	registry := NewGaugeRegistry(prm)
	m1 := registry.Register("test_gauge", "description", make(map[string]string))
	m2 := registry.Register("test_gauge", "description", make(map[string]string))
	assert.Same(t, m1, m2)
}

func TestGauge_GetFromRegistryNotNil(t *testing.T) {
	prm := mocks.NewRegisterer(t)
	prm.On("MustRegister", mock.Anything)

	registry := NewGaugeRegistry(prm)
	registry.Register("test_gauge", "description", make(map[string]string))
	metric := registry.Get("test_gauge")
	assert.NotNil(t, metric)
}

func TestGauge_GetFromRegistryNotNilWhenNotRegistered(t *testing.T) {
	prm := mocks.NewRegisterer(t)

	registry := NewGaugeRegistry(prm)
	metric := registry.Get("test_gauge")
	assert.Equal(t, &DummyMetricGauge{}, metric)
}

func TestGauge_GetFromRegistrySame(t *testing.T) {
	prm := mocks.NewRegisterer(t)
	prm.On("MustRegister", mock.Anything)

	registry := NewGaugeRegistry(prm)
	registry.Register("test_gauge2", "description", make(map[string]string))
	metric := registry.Get("test_gauge2")
	promMetric := metric.(*runaimetrics.Gauge)

	assert.NotNil(t, metric)
	assert.Equal(t, "Desc{fqName: \"runai_test_gauge2\", help: \"description\", constLabels: {}, variableLabels: {}}", promMetric.PromGauge.Desc().String())
}

func TestGauge_GetFromRegistrySameLabels(t *testing.T) {
	prm := mocks.NewRegisterer(t)
	prm.On("MustRegister", mock.Anything)

	registry := NewGaugeRegistry(prm)
	registry.Register("test_gauge3", "description", map[string]string{"bla": "bla"})
	metric := registry.Get("test_gauge3")
	promMetric := metric.(*runaimetrics.Gauge)

	assert.NotNil(t, metric)
	assert.Equal(t, "Desc{fqName: \"runai_test_gauge3\", help: \"description\", constLabels: {bla=\"bla\"}, variableLabels: {}}", promMetric.PromGauge.Desc().String())
}
