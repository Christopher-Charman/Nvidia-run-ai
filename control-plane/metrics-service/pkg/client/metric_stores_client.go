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
	CreateMetricStore(ctx context.Context, tenantMetricStore *api.MetricStore) (*api.MetricStore, error)
	GetMetricStore(ctx context.Context, overrideTenantId int32) (*api.MetricStore, error)
	DeleteMetricStores(ctx context.Context, overrideTenantId int32) error
}

type metricStoresClient struct {
	client *api.APIClient
	config *Config
}

func (m *metricStoresClient) CreateMetricStore(ctx context.Context, metricStore *api.MetricStore) (*api.MetricStore, error) {
	if err := validateMetricStoreData(metricStore); err != nil {
		return nil, NewAPIError(http.StatusBadRequest, err.Error())
	}

	ms, httpResponse, err := m.client.MetricStoresAPI.CreateMetricStore(ctx).MetricStore(*metricStore).Execute()
	if err != nil {
		return nil, NewAPIErrorFromResponse(httpResponse, err)
	}
	return ms, nil
}

func (m *metricStoresClient) GetMetricStore(ctx context.Context, overrideTenantId int32) (*api.MetricStore, error) {
	mss, httpResponse, err := m.client.MetricStoresAPI.GetMetricStores(ctx).OverrideTenantId(overrideTenantId).Execute()
	if err != nil {
		return nil, NewAPIErrorFromResponse(httpResponse, err)
	}

	if len(mss) == 0 {
		return nil, NewAPIError(http.StatusNotFound, "metric store not found")
	}

	return &mss[0], nil
}

func (m *metricStoresClient) DeleteMetricStores(ctx context.Context, overrideTenantId int32) error {
	httpResponse, err := m.client.MetricStoresAPI.DeleteMetricStores(ctx).OverrideTenantId(overrideTenantId).Execute()
	if err != nil {
		return NewAPIErrorFromResponse(httpResponse, err)
	}
	return nil
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
	if metricStore.DatasourceId == "" {
		return fmt.Errorf("tenant metric store configured not appropriately: no datasource id")
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
