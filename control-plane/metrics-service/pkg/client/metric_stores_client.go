package client

import (
	"context"
	"fmt"
	"sync"

	api "github.com/dsx-ai-factory/run-ai/control-plane/metrics-service/pkg/client/generated"
	openapi "github.com/dsx-ai-factory/run-ai/control-plane/metrics-service/pkg/client/generated"
)

var o sync.Once

var instance MetricStoresClienter

type MetricStoresClienter interface {
	CreateMetricStore(ctx context.Context, tenantMetricStore *api.MetricStore) (*api.MetricStore, error)
	GetMetricStore(ctx context.Context, overrideTenantId int32) (*api.MetricStore, error)
}

type metricStoresClient struct {
	client *openapi.APIClient
	config *Config
}

func (m *metricStoresClient) CreateMetricStore(ctx context.Context, metricStore *openapi.MetricStore) (*openapi.MetricStore, error) {
	if err := validateMetricStoreData(metricStore); err != nil {
		return nil, err
	}

	ms, _, err := m.client.MetricStoresAPI.CreateMetricStore(ctx).MetricStore(*metricStore).Execute()
	return ms, err
}

func (m *metricStoresClient) GetMetricStore(ctx context.Context, overrideTenantId int32) (*openapi.MetricStore, error) {
	mss, _, err := m.client.MetricStoresAPI.GetMetricStores(ctx).OverrideTenantId(overrideTenantId).Execute()
	if err != nil || len(mss) == 0 {
		return nil, err
	}

	return &mss[0], err
}

func GetMetricStoresClient() MetricStoresClienter {
	o.Do(func() {
		config := readConfig()
		clientConfig := openapi.NewConfiguration()
		clientConfig.Host = config.MetricsServerHost
		clientConfig.Scheme = config.MetricsServerScheme
		instance = &metricStoresClient{
			client: openapi.NewAPIClient(clientConfig),
			config: config,
		}
	})

	return instance
}

func validateMetricStoreData(metricStore *openapi.MetricStore) error {
	if !metricStore.HasDatasourceId() {
		return fmt.Errorf("tenant metric store configured not appropriately: no datasource id")
	}
	if !metricStore.HasMetricStoreDetails() {
		return fmt.Errorf("tenant metric store configured not appropriately: no metric store")
	}
	if metricStore.MetricStoreDetails.GetType() == "" {
		return fmt.Errorf("tenant metric store configured not appropriately: no metric store type")
	}

	if _, ok := metricStore.MetricStoreDetails.GetReadOk(); !ok {
		return fmt.Errorf("tenant metric store configured not appropriately: no metric store read configured")
	}
	if metricStore.MetricStoreDetails.Read.Url == "" {
		return fmt.Errorf("tenant metric store configured not appropriately: no metric store read url")
	}

	if _, ok := metricStore.MetricStoreDetails.GetWriteOk(); !ok {
		return fmt.Errorf("tenant metric store configured not appropriately: no metric store write configured")
	}
	if metricStore.MetricStoreDetails.Write.Url == "" {
		return fmt.Errorf("tenant metric store configured not appropriately: no metric store write url")
	}

	return nil
}
