package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	openapi "github.com/dsx-ai-factory/run-ai/runai-clients/assets-service-client/generated"
)

type AssetsClient struct {
	ApiClient *openapi.APIClient
}

func (c *AssetsClient) SyncPvcsStatus(ctx context.Context, clusterId string, pvcs []openapi.PVCStatusRequestItem) (*http.Response, error) {
	return c.ApiClient.SynchronizationAPI.SyncPvcsStatus(
		ctx, clusterId).PVCStatusRequest(
		openapi.PVCStatusRequest{
			Pvcs: pvcs,
		}).Execute()
}

func (c *AssetsClient) ListComputeAssets(ctx context.Context, name string) (*openapi.ComputeListResponse, *http.Response, error) {
	return c.ApiClient.ComputeAPI.ListComputeAssets(ctx).Name(name).Execute()
}

func (c *AssetsClient) CreateComputeAsset(ctx context.Context, request *openapi.ComputeCreationRequest) (*openapi.ComputeAsset, *http.Response, error) {
	return c.ApiClient.ComputeAPI.CreateComputeAsset(ctx).ComputeCreationRequest(*request).Execute()
}

func (c *AssetsClient) ListEnvironmentAssets(ctx context.Context, name string) (*openapi.EnvironmentListResponse, *http.Response, error) {
	return c.ApiClient.EnvironmentAPI.ListEnvironmentAssets(ctx).Name(name).Execute()
}

func (c *AssetsClient) CreateEnvironmentAsset(ctx context.Context, request *openapi.EnvironmentCreationRequest) (*openapi.EnvironmentAsset, *http.Response, error) {
	return c.ApiClient.EnvironmentAPI.CreateEnvironmentAsset(ctx).EnvironmentCreationRequest(*request).Execute()
}

func (c *AssetsClient) ListModelAssets(ctx context.Context, name string) (*openapi.ModelListResponse, *http.Response, error) {
	return c.ApiClient.ModelAPI.ListModelAssets(ctx).Name(name).Execute()
}

func (c *AssetsClient) CreateModelAsset(ctx context.Context, request *openapi.ModelCreationRequest) (*openapi.ModelAsset, *http.Response, error) {
	return c.ApiClient.ModelAPI.CreateModelAsset(ctx).ModelCreationRequest(*request).Execute()
}

func (c *AssetsClient) ListDatasourceAssets(ctx context.Context, ids []string) (*openapi.DatasourceListResponse, *http.Response, error) {
	return c.ApiClient.DatasourcesAPI.ListDatasourceAssets(ctx).AssetIds(strings.Join(ids, ",")).
		Execute()
}

func (c *AssetsClient) RequestAssetFlatFormat(ctx context.Context, request openapi.AssetFlatRequest) (*openapi.AssetFlatResponse, *openapi.Error) {
	resp, httpResp, err := c.ApiClient.AssetsAPI.RequestAssetFlatFormat(ctx).AssetFlatRequest(request).Execute()
	if httpResp != nil {
		if httpResp.Body != nil {
			defer httpResp.Body.Close()
		}

		if httpResp.StatusCode != http.StatusOK {
			// Read the body
			body, err := io.ReadAll(httpResp.Body)
			if err != nil {
				return nil, openapi.NewError(http.StatusInternalServerError, fmt.Sprintf("failed to read asset payload: %v", err.Error()))
			}

			var apiError openapi.Error

			jsonErr := json.Unmarshal(body, &apiError)
			if jsonErr != nil {
				return nil, openapi.NewError(int32(httpResp.StatusCode), string(body))
			} else {
				return nil, &apiError
			}
		}

		if resp != nil {
			return resp, nil
		}
	}

	return nil, openapi.NewError(http.StatusServiceUnavailable, fmt.Sprintf("%v", err))
}

func (c *AssetsClient) CredentialsSyncUpdates(ctx context.Context, clusterId string, fromUpdate int64, request openapi.ClusterSyncReport) (*openapi.CredentialsUpdatesResp, *openapi.Error) {
	resp, httpResp, err := c.ApiClient.SynchronizationAPI.CredentialsSyncUpdates(ctx).ClusterSyncReport(request).After(fromUpdate).ClusterId(clusterId).Execute()
	if httpResp != nil {
		if httpResp.Body != nil {
			defer httpResp.Body.Close()
		}

		if httpResp.StatusCode != http.StatusOK {
			// Read the body
			body, err := io.ReadAll(httpResp.Body)
			if err != nil {
				return nil, openapi.NewError(http.StatusInternalServerError, fmt.Sprintf("failed to read response payload: %v", err.Error()))
			}

			var apiError openapi.Error

			jsonErr := json.Unmarshal(body, &apiError)
			if jsonErr != nil {
				return nil, openapi.NewError(int32(httpResp.StatusCode), string(body))
			} else {
				return nil, &apiError
			}
		}

		if resp != nil {
			return resp, nil
		}
	}

	return nil, openapi.NewError(http.StatusServiceUnavailable, fmt.Sprintf("%v", err))
}

func (c *AssetsClient) DatasourceSyncUpdates(ctx context.Context, clusterId string, fromUpdate int64, request openapi.ClusterSyncReport) (*openapi.DatasourceUpdatesResp, *openapi.Error) {
	resp, httpResp, err := c.ApiClient.SynchronizationAPI.DatasourceSyncUpdates(ctx).ClusterSyncReport(request).After(fromUpdate).ClusterId(clusterId).Execute()
	if httpResp != nil {
		if httpResp.Body != nil {
			defer httpResp.Body.Close()
		}

		if httpResp.StatusCode != http.StatusOK {
			// Read the body
			body, err := io.ReadAll(httpResp.Body)
			if err != nil {
				return nil, openapi.NewError(http.StatusInternalServerError, fmt.Sprintf("failed to read response payload: %v", err.Error()))
			}

			var apiError openapi.Error

			jsonErr := json.Unmarshal(body, &apiError)
			if jsonErr != nil {
				return nil, openapi.NewError(int32(httpResp.StatusCode), string(body))
			} else {
				return nil, &apiError
			}
		}

		if resp != nil {
			return resp, nil
		}
	}

	return nil, openapi.NewError(http.StatusServiceUnavailable, fmt.Sprintf("%v", err))
}

func (c *AssetsClient) ListTemplatesOfTenant(ctx context.Context) (*openapi.WorkloadTemplateListResponse, *openapi.Error) {
	resp, httpResp, err := c.ApiClient.TemplateAPI.ListTemplates(ctx).Execute()
	if httpResp != nil {
		if httpResp.Body != nil {
			defer httpResp.Body.Close()
		}

		if httpResp.StatusCode != http.StatusOK {
			// Read the body
			body, err := io.ReadAll(httpResp.Body)
			if err != nil {
				return nil, openapi.NewError(http.StatusInternalServerError, fmt.Sprintf("failed to read response payload of list templates: %v", err.Error()))
			}

			var apiError openapi.Error

			jsonErr := json.Unmarshal(body, &apiError)
			if jsonErr != nil {
				return nil, openapi.NewError(int32(httpResp.StatusCode), string(body))
			} else {
				return nil, &apiError
			}
		}

		if resp != nil {
			return resp, nil
		}
	}

	return nil, openapi.NewError(http.StatusServiceUnavailable, fmt.Sprintf("failed to request list of templates: %v", err))
}

func (c *AssetsClient) ListCredentialsAssets(ctx context.Context, privateOnly bool) (*openapi.CredentialsListResponse, *openapi.Error) {
	resp, httpResp, err := c.ApiClient.CredentialsAPI.ListCredentialsAssets(ctx).PrivateOnly(privateOnly).Execute()
	if httpResp != nil {
		if httpResp.Body != nil {
			defer httpResp.Body.Close()
		}

		if httpResp.StatusCode != http.StatusOK {
			// Read the body
			body, err := io.ReadAll(httpResp.Body)
			if err != nil {
				return nil, openapi.NewError(http.StatusInternalServerError, fmt.Sprintf("failed to read response payload of list credentials: %v", err.Error()))
			}

			var apiError openapi.Error

			jsonErr := json.Unmarshal(body, &apiError)
			if jsonErr != nil {
				return nil, openapi.NewError(int32(httpResp.StatusCode), string(body))
			} else {
				return nil, &apiError
			}
		}

		if resp != nil {
			return resp, nil
		}
	}

	return nil, openapi.NewError(http.StatusServiceUnavailable, fmt.Sprintf("failed to request list of credentials: %v", err))
}
