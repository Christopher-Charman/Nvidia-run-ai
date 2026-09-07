package audit_service_client

import (
	"context"

	client_utils "github.com/run-ai/runai-application-packages/client-utils"
	openapi "github.com/dsx-ai-factory/run-ai/runai-clients/audit-service-client/generated"
	"github.com/dsx-ai-factory/run-ai/runai-clients/audit-service-client/internal/api"
)

//go:generate mockery --name IAuditExternalClient
type IAuditExternalClient interface {
	DownloadAuditLogs(ctx context.Context) openapi.ApiDownloadAuditLogsRequest
	GetAuditLogs(ctx context.Context) openapi.ApiGetAuditLogsRequest
}

//go:generate mockery --name IAuditClient
type IAuditClient interface {
	IAuditExternalClient
}

func NewAuditClient(baseURL string) IAuditClient {
	config := openapi.NewConfiguration()
	config.Servers = openapi.ServerConfigurations{{URL: client_utils.GetInternalUrl(baseURL)}}
	return &api.AuditClient{ApiClient: openapi.NewAPIClient(config)}
}
