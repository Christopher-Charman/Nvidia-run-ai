package internal

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/run-ai/runai-common-packages/cache/v2"
	openapi "github.com/dsx-ai-factory/run-ai/runai-clients/cluster-service-client/generated"
)

type ClusterClient struct {
	ApiClient *openapi.APIClient
}

func (client *ClusterClient) GetClusterInternalConfig(ctx context.Context, clusterUUID string) (*openapi.GetInternalClusterConfigResponse, error) {
	request := client.ApiClient.ClusterConfigAPI.GetInternalClusterConfig(ctx, clusterUUID)
	internalConfig, response, err := request.Execute()
	if err != nil {
		return nil, err
	}

	if response != nil && response.Body != nil {
		defer response.Body.Close()
	}

	if !checkResponseStatusCodeOK(response) {
		return nil, fmt.Errorf("get internal cluster config failed, status code error: %s", response.Status)
	}

	return internalConfig, nil
}

func (client *ClusterClient) GetClusterControlPlaneConfig(ctx context.Context, clusterUUID string) (*openapi.ClusterControlPlaneConfig, error) {
	request := client.ApiClient.ControlPlaneConfigAPI.GetClusterControlPlaneConfig(ctx, clusterUUID)
	cpConfig, response, err := request.Execute()
	if err != nil {
		return nil, err
	}

	if response != nil && response.Body != nil {
		defer response.Body.Close()
	}

	if !checkResponseStatusCodeOK(response) {
		return nil, fmt.Errorf("get internal cluster cp config failed, status code error: %s", response.Status)
	}

	return cpConfig, nil
}

func (client *ClusterClient) CreateCluster(ctx context.Context, clusterCreateRequest *openapi.ClusterCreationRequest) (*openapi.DisplayedCluster, error) {
	request := client.ApiClient.ClustersAPI.CreateCluster(ctx)
	request = request.ClusterCreationRequest(*clusterCreateRequest)
	displayedCluster, response, err := request.Execute()
	if err != nil {
		return nil, err
	}

	if response != nil && response.Body != nil {
		defer response.Body.Close()
	}

	if !checkResponseStatusCodeOK(response) {
		return nil, fmt.Errorf("create cluster config failed, status code error: %s", response.Status)
	}

	return displayedCluster, nil
}

func (client *ClusterClient) GetClusterByUuid(ctx context.Context, id string) (*openapi.DisplayedCluster, *openapi.Error) {
	cachedCluster, _ := cache.Get[openapi.DisplayedCluster](ctx, id)
	if cachedCluster != nil && cachedCluster.Version.IsSet() {
		return cachedCluster, nil
	}

	clusterRequest := client.ApiClient.ClustersAPI.GetClusterByUuid(ctx, id)
	displayedCluster, response, err := client.ApiClient.ClustersAPI.GetClusterByUuidExecute(clusterRequest)
	if response != nil {
		if response.StatusCode != http.StatusOK {
			defer response.Body.Close()

			body, err := io.ReadAll(response.Body)
			if err != nil {
				return nil, openapi.NewError(int32(response.StatusCode), fmt.Sprintf("Failed to get cluster %s status is %s", id, response.Status))
			}

			return nil, openapi.NewError(int32(response.StatusCode), string(body))
		}

		if displayedCluster != nil {
			cache.Set[openapi.DisplayedCluster](ctx, id, *displayedCluster)
		}

		return displayedCluster, nil
	}

	if err != nil {
		return nil, openapi.NewError(http.StatusServiceUnavailable, err.Error())
	}

	return nil, openapi.NewError(http.StatusInternalServerError, fmt.Sprintf("failed to get cluster %s details", id))
}

func (client *ClusterClient) GetClusters(ctx context.Context) ([]openapi.DisplayedCluster, error) {
	clustersRequest := client.ApiClient.ClustersAPI.GetClusters(ctx)
	displayedClusters, response, err := client.ApiClient.ClustersAPI.GetClustersExecute(clustersRequest)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("Failed to get clusters. status is %s", response.Status))
		return nil, err
	}

	return displayedClusters, nil
}

func (client *ClusterClient) PutClusterConfigStatus(ctx context.Context, clusterUUID string, putClusterConfigStatusRequest *openapi.PutClusterConfigStatusRequest) error {
	request := client.ApiClient.ClusterConfigAPI.PutClusterConfigStatus(ctx, clusterUUID)
	request = request.PutClusterConfigStatusRequest(*putClusterConfigStatusRequest)
	response, err := request.Execute()
	if err != nil {
		return err
	}
	if response != nil && response.Body != nil {
		defer response.Body.Close()
	}
	if !checkResponseStatusCodeOK(response) {
		return fmt.Errorf("put cluster remote config failed, status code error: %s", response.Status)
	}

	return nil
}

func (client *ClusterClient) GetNodePoolsForTenants(ctx context.Context, tenantIds []string) (*openapi.GetNodepoolsInternal200Response, error) {
	request := client.ApiClient.NodePoolsAPI.GetNodepoolsInternal(ctx).TenantId(tenantIds)
	nodePools, response, err := request.Execute()
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("Failed to get nodepools. status is %s", response.Status))
		return nil, err
	}
	return nodePools, nil
}

func (client *ClusterClient) GetNodePoolsByCluster(ctx context.Context, clusterUUID string) (*openapi.GetNodepools200Response, error) {
	nodePools, response, err := client.ApiClient.NodePoolsAPI.
		GetNodepools(ctx).
		FilterBy([]string{"clusterId==" + clusterUUID}).
		Execute()
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("Failed to get nodepools. status is %s", response.Status))
		return nil, err
	}
	return nodePools, nil
}

func (client *ClusterClient) GetClustersCount(ctx context.Context, groupBy string) ([]openapi.ClustersCount, *openapi.Error) {
	countRequest := client.ApiClient.ClustersAPI.GetClustersCount(ctx).GroupBy(groupBy)
	clustersCount, response, err := client.ApiClient.ClustersAPI.GetClustersCountExecute(countRequest)
	if err != nil {
		return nil, openapi.NewError(http.StatusServiceUnavailable, err.Error())
	}
	if response == nil {
		return nil, openapi.NewError(http.StatusServiceUnavailable, "Failed to get parse response of get cluster count request")
	}
	if response.StatusCode != http.StatusOK {
		return nil, openapi.NewError(int32(response.StatusCode), fmt.Sprintf("Failed to get clusters count. status is %s", response.Status))
	}
	return clustersCount, nil
}

func checkResponseStatusCodeOK(response *http.Response) bool {
	return response.StatusCode >= http.StatusOK || response.StatusCode <= http.StatusIMUsed
}
