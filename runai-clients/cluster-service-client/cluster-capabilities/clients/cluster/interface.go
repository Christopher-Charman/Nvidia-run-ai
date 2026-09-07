package cluster

import (
	"context"

	clusterService "github.com/dsx-ai-factory/run-ai/runai-clients/cluster-service-client/generated"
)

type FactoryInterface interface {
	Client() ClientInterface
}

type ClientInterface interface {
	GetClusterByUuid(ctx context.Context, clusterId string) (*clusterService.DisplayedCluster, *clusterService.Error)
}
