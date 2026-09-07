package cluster_service_client

import (
	"context"
	"fmt"
	"time"

	client_utils "github.com/run-ai/runai-application-packages/client-utils"
	"github.com/run-ai/runai-common-packages/cache/v2"
	"github.com/run-ai/runai-common-packages/cache/v2/options"
	"github.com/run-ai/runai-common-packages/logging"
	openapi "github.com/dsx-ai-factory/run-ai/runai-clients/cluster-service-client/generated"
	"github.com/dsx-ai-factory/run-ai/runai-clients/cluster-service-client/internal"
)

const (
	authorizationHeader = "Authorization"
	bearerTokenStr      = "Bearer %s"
)

//go:generate mockery --name IClusterExternalClient
type IClusterExternalClient interface {
	GetClusterInternalConfig(ctx context.Context, clusterUUID string) (*openapi.GetInternalClusterConfigResponse, error)
	GetClusterControlPlaneConfig(ctx context.Context, clusterUUID string) (*openapi.ClusterControlPlaneConfig, error)
	CreateCluster(ctx context.Context, clusterCreateRequest *openapi.ClusterCreationRequest) (*openapi.DisplayedCluster, error)
	GetClusterByUuid(ctx context.Context, id string) (*openapi.DisplayedCluster, *openapi.Error)
	GetClusters(ctx context.Context) ([]openapi.DisplayedCluster, error)
	PutClusterConfigStatus(ctx context.Context, clusterUUID string, putClusterConfigStatusRequest *openapi.PutClusterConfigStatusRequest) error
	GetNodePoolsByCluster(ctx context.Context, clusterUUID string) (*openapi.GetNodepools200Response, error)
}

//go:generate mockery --name IClusterClient
type IClusterClient interface {
	IClusterExternalClient
	GetNodePoolsForTenants(ctx context.Context, tenantIds []string) (*openapi.GetNodepoolsInternal200Response, error)
	GetClustersCount(ctx context.Context, groupBy string) ([]openapi.ClustersCount, *openapi.Error)
}

func NewClusterExternalClient(baseURL string, token string) IClusterExternalClient {
	config := clusterClientConfig(baseURL, true, &token)
	return &internal.ClusterClient{
		ApiClient: openapi.NewAPIClient(config),
	}
}

func NewClusterClient(baseURL string) IClusterClient {
	clientConfig := openapi.NewConfiguration()
	clientConfig.Servers = openapi.ServerConfigurations{{URL: client_utils.GetInternalUrl(baseURL)}}
	client := openapi.NewAPIClient(clientConfig)
	return &internal.ClusterClient{ApiClient: client}
}

func NewCachedClusterClient(baseURL string, ttl time.Duration) IClusterClient {
	err := cache.RegisterStore[openapi.DisplayedCluster](options.Ttl(ttl))
	if err != nil {
		logging.Warnf("Failed to register cache store for type DisplayedCluster: %v", err)
	}
	config := clusterClientConfig(baseURL, false, nil)
	return &internal.ClusterClient{
		ApiClient: openapi.NewAPIClient(config),
	}
}

func clusterClientConfig(baseURL string, external bool, token *string) *openapi.Configuration {
	configuration := openapi.NewConfiguration()
	if external {
		configuration.Servers = openapi.ServerConfigurations{{URL: client_utils.GetExternalUrl(baseURL)}}
	} else {
		configuration.Servers = openapi.ServerConfigurations{{URL: client_utils.GetInternalUrl(baseURL)}}
	}

	if token != nil {
		configuration.AddDefaultHeader(authorizationHeader, fmt.Sprintf(bearerTokenStr, *token))
	}
	return configuration
}

func NewClusterServiceClientWithConfig(config *openapi.Configuration) IClusterClient {
	return &internal.ClusterClient{ApiClient: openapi.NewAPIClient(config)}
}

func NewClusterServiceExternalClientWithConfig(config *openapi.Configuration) IClusterExternalClient {
	return NewClusterServiceClientWithConfig(config)
}
