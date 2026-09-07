package registry

import (
	"testing"

	"github.com/dsx-ai-factory/run-ai/runai-common-packages/metrics/internal/registry/mocks"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/metrics/internal/runaimetrics"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCounter_RegistryTwiceReturnsSame(t *testing.T) {
	prm := mocks.NewRegisterer(t)
	prm.On("MustRegister", mock.Anything)

	registry := NewCounterRegistry(prm)
	m1 := registry.Register("test_counter", "description", make(map[string]string))
	m2 := registry.Register("test_counter", "description", make(map[string]string))
	assert.Same(t, m1, m2)
}

func TestCounter_GetFromRegistryNotNil(t *testing.T) {
	prm := mocks.NewRegisterer(t)
	prm.On("MustRegister", mock.Anything)

	registry := NewCounterRegistry(prm)
	registry.Register("test_counter", "description", make(map[string]string))
	metric := registry.Get("test_counter")
	assert.NotNil(t, metric)
}

func TestCounter_GetFromRegistryNotNilWhenNotRegistered(t *testing.T) {
	prm := mocks.NewRegisterer(t)

	registry := NewCounterRegistry(prm)
	metric := registry.Get("test_counter")
	assert.Equal(t, &DummyMetricCounter{}, metric)
}

func TestCounter_GetFromRegistrySame(t *testing.T) {
	prm := mocks.NewRegisterer(t)
	prm.On("MustRegister", mock.Anything)

	registry := NewCounterRegistry(prm)
	registry.Register("test_counter2", "description", make(map[string]string))
	metric := registry.Get("test_counter2")
	promMetric := metric.(*runaimetrics.Counter)

	assert.NotNil(t, metric)
	assert.Equal(t, "Desc{fqName: \"runai_test_counter2\", help: \"description\", constLabels: {}, variableLabels: {}}", promMetric.PromCounter.Desc().String())
}

func TestCounter_GetFromRegistrySameLables(t *testing.T) {
	prm := mocks.NewRegisterer(t)
	prm.On("MustRegister", mock.Anything)

	registry := NewCounterRegistry(prm)
	registry.Register("test_counter3", "description", map[string]string{"bla": "bla"})
	metric := registry.Get("test_counter3")
	promMetric := metric.(*runaimetrics.Counter)

	assert.NotNil(t, metric)
	assert.Equal(t, "Desc{fqName: \"runai_test_counter3\", help: \"description\", constLabels: {bla=\"bla\"}, variableLabels: {}}", promMetric.PromCounter.Desc().String())
}
