package workloads_client

import (
	"context"

	client_utils "github.com/run-ai/runai-application-packages/client-utils"
	openapi "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client/generated"
	"github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client/internal"
)

//go:generate mockery --name IWorkloadsExternalClient
type IWorkloadsExternalClient interface {
}

//go:generate mockery --name IWorkloadsClient
type IWorkloadsClient interface {
	IWorkloadsExternalClient
	GetAncestors(ctx context.Context, resourceId string) (*openapi.Ancestors, error)
	GetInferenceAncestors(ctx context.Context, resourceId string) (*openapi.Ancestors, error)
	GetWorkloads(ctx context.Context, limit, offset int32, filter []string, includeDeleted bool, sortBy string, sortOrder string) (*openapi.GetWorkloads200Response, error)
}

func NewWorkloadsClient(baseURL string) IWorkloadsClient {
	clientConfig := openapi.NewConfiguration()
	clientConfig.Servers = openapi.ServerConfigurations{{URL: client_utils.GetInternalUrl(baseURL)}}
	client := openapi.NewAPIClient(clientConfig)
	return &internal.WorkloadsClient{ApiClient: client}
}

func NewWorkloadsExternalClient(baseURL string) IWorkloadsExternalClient {
	return NewWorkloadsClient(client_utils.GetExternalUrl(baseURL))
}
