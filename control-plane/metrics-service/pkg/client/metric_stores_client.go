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
	CreateMetricStore(ctx context.Context, tenantMetricStore *api.MetricStore) (*api.MetricStore, int, error)
	GetMetricStore(ctx context.Context, overrideTenantId int32) (*api.MetricStore, int, error)
}

type metricStoresClient struct {
	client *api.APIClient
	config *Config
}

func (m *metricStoresClient) CreateMetricStore(ctx context.Context, metricStore *api.MetricStore) (*api.MetricStore, int, error) {
	if err := validateMetricStoreData(metricStore); err != nil {
		return nil, 0, err
	}

	ms, resp, err := m.client.MetricStoresAPI.CreateMetricStore(ctx).MetricStore(*metricStore).Execute()

	statusCode := 0
	if resp != nil {
		statusCode = resp.StatusCode
	}

	if err != nil {
		return nil, statusCode, err
	}
	return ms, statusCode, nil
}

func (m *metricStoresClient) GetMetricStore(ctx context.Context, overrideTenantId int32) (*api.MetricStore, int, error) {
	mss, resp, err := m.client.MetricStoresAPI.GetMetricStores(ctx).OverrideTenantId(overrideTenantId).Execute()

	statusCode := 0
	if resp != nil {
		statusCode = resp.StatusCode
	}

	if err != nil {
		return nil, statusCode, err
	}
	if len(mss) == 0 || statusCode == 404 {
		return nil, http.StatusNotFound, fmt.Errorf("metric store not found for tenant")
	}
	return &mss[0], statusCode, nil
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
