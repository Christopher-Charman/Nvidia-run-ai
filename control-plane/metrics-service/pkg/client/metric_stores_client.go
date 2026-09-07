package client

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	api "github.com/dsx-ai-factory/run-ai/control-plane/metrics-service/pkg/client/generated"
)

var o sync.Once

var instance MetricStoresClienter

//go:generate mockery --name MetricStoresClienter
type MetricStoresClienter interface {
	CreateMetricStore(ctx context.Context, tenantMetricStore *api.MetricStore) (*api.MetricStore, *http.Response, error)
	GetMetricStore(ctx context.Context, overrideTenantId int32) (*api.MetricStore, *http.Response, error)
	DeleteMetricStores(ctx context.Context, overrideTenantId int32) (*http.Response, error)
}

type metricStoresClient struct {
	client *api.APIClient
	config *Config
}

func (m *metricStoresClient) CreateMetricStore(ctx context.Context, metricStore *api.MetricStore) (*api.MetricStore, *http.Response, error) {
	if err := validateMetricStoreData(metricStore); err != nil {
		return nil, &http.Response{StatusCode: http.StatusBadRequest}, err
	}

	ms, httpResponse, err := m.client.MetricStoresAPI.CreateMetricStore(ctx).MetricStore(*metricStore).Execute()
	if err != nil {
		return nil, httpResponse, err
	}
	return ms, httpResponse, nil
}

func (m *metricStoresClient) GetMetricStore(ctx context.Context, overrideTenantId int32) (*api.MetricStore, *http.Response, error) {
	mss, httpResponse, err := m.client.MetricStoresAPI.GetMetricStores(ctx).OverrideTenantId(overrideTenantId).Execute()
	if err != nil {
		return nil, httpResponse, err
	}

	if len(mss) == 0 {
		return nil, &http.Response{StatusCode: http.StatusNotFound}, fmt.Errorf("metric store not found")
	}

	return &mss[0], httpResponse, nil
}

func (m *metricStoresClient) DeleteMetricStores(ctx context.Context, overrideTenantId int32) (*http.Response, error) {
	httpResponse, err := m.client.MetricStoresAPI.DeleteMetricStores(ctx).OverrideTenantId(overrideTenantId).Execute()
	return httpResponse, err
}

func GetMetricStoresClient() MetricStoresClienter {
	o.Do(func() {
		config := readConfig()
		clientConfig := api.NewConfiguration()
		clientConfig.Host = config.MetricsServerHost
		clientConfig.Scheme = config.MetricsServerScheme
		instance = &metricStoresClient{
			client: api.NewAPIClient(clientConfig),
			config: config,
		}
	})

	return instance
}

func validateMetricStoreData(metricStore *api.MetricStore) error {
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
