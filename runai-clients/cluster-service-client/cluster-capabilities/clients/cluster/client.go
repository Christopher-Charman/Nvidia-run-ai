package cluster

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/run-ai/runai-common-packages/logging"
	clusterService "github.com/dsx-ai-factory/run-ai/runai-clients/cluster-service-client/generated"
)

// Client -> for the "real" policies client we need an APIClient to the cluster service
type Client struct {
	apiClient *clusterService.APIClient
}

// Factory -> the "real" client factory. during testing it is replaced with factory which
// generated mock clients
type Factory struct{}

func (self Factory) Client() ClientInterface {
	return &Client{
		apiClient: clusterService.NewAPIClient(serverConfig),
	}
}

func (self Client) GetClusterByUuid(ctx context.Context, clusterId string) (*clusterService.DisplayedCluster, *clusterService.Error) {
	resp, r, err := self.apiClient.ClustersAPI.GetClusterByUuid(ctx, clusterId).Verbosity("full").Execute()
	if r != nil {
		if r.StatusCode == http.StatusOK {
			return resp, nil
		}
		if r.StatusCode == http.StatusNotFound {
			return nil, clusterService.NewError(http.StatusNotFound, fmt.Sprintf("cluster %s does not exist", clusterId))
		}
		data, _ := io.ReadAll(r.Body)

		logging.Errorf("failed to get info of cluster %s: %s", clusterId, data)
		return nil, clusterService.NewError(http.StatusServiceUnavailable, fmt.Sprintf("failed to get info of cluster %s", clusterId))
	}

	logging.Errorf("failed to request info of cluster %s: %v", clusterId, err)
	return nil, clusterService.NewError(http.StatusServiceUnavailable, fmt.Sprintf("failed to request info of cluster %s", clusterId))
}
