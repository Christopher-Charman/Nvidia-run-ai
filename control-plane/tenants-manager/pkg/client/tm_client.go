package tenants_manager_client

import (
	"context"
	"net/http"
	"strings"

	openapi "github.com/dsx-ai-factory/run-ai/control-plane/tenants-manager/pkg/client/generated"
)

//go:generate mockery --name ITenantsManager
type ITenantsManager interface {
	ListAllTenantsInternal(ctx context.Context) ([]openapi.Tenant, *http.Response, error)
	GetTenantByIdInternal(ctx context.Context, id int32) (*openapi.Tenant, *http.Response, error)
}

type TenantsManagerClient struct {
	apiClient *openapi.APIClient
}

// NewTenantsManagerClient creates a new API client.  Requires a base URL.
func NewTenantsManagerClient(baseURL string) ITenantsManager {
	config := openapi.NewConfiguration()
	config.Host = strings.Replace(baseURL, "http://", "", 1)
	config.Scheme = "http"
	return &TenantsManagerClient{
		apiClient: openapi.NewAPIClient(config),
	}
}

func (t TenantsManagerClient) ListAllTenantsInternal(ctx context.Context) ([]openapi.Tenant, *http.Response, error) {
	tenants, httpResponse, err := t.apiClient.InternalAPI.ManagementGetTenantsInternal(ctx).Execute()
	if err != nil {
		return nil, httpResponse, err
	}
	return tenants, httpResponse, nil
}

func (t TenantsManagerClient) GetTenantByIdInternal(ctx context.Context, id int32) (*openapi.Tenant, *http.Response, error) {
	tenant, httpResponse, err := t.apiClient.TenantsAPI.GetTenantByIdInternal(ctx, id).Execute()
	if err != nil {
		return nil, httpResponse, err
	}
	return tenant, httpResponse, nil
}
