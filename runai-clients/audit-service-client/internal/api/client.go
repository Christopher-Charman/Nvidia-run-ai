package api

import (
	"context"

	openapi "github.com/dsx-ai-factory/run-ai/runai-clients/audit-service-client/generated"
)

type AuditClient struct {
	ApiClient *openapi.APIClient
}

func (a *AuditClient) GetAuditLogs(ctx context.Context) openapi.ApiGetAuditLogsRequest {
	return a.ApiClient.AuditLogsAPI.GetAuditLogs(ctx)
}

func (a *AuditClient) DownloadAuditLogs(ctx context.Context) openapi.ApiDownloadAuditLogsRequest {
	return a.ApiClient.AuditLogsAPI.DownloadAuditLogs(ctx)
}
