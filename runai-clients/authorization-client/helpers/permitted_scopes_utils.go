package helpers

import (
	"context"
	"fmt"
	"strings"

	authorization_api "github.com/dsx-ai-factory/run-ai/runai-clients/authorization-client/generated"
	"github.com/dsx-ai-factory/run-ai/runai-clients/authorization-client/structs"
	"golang.org/x/exp/slices"
)

var (
	PermittedScopesKey = structs.PermittedScopesKey
)

var insufficientPermissionsError = fmt.Errorf("insufficient permissions")

func GetPermittedScopesFromCtx(ctx context.Context) (*authorization_api.PermittedScopes, bool) {
	permittedScopes, ok := ctx.Value(PermittedScopesKey).(*authorization_api.PermittedScopes)

	return permittedScopes, ok
}

func HasPermittedScopes(ctx context.Context) error {
	permittedScopes, ok := ctx.Value(PermittedScopesKey).(*authorization_api.PermittedScopes)
	if !ok || IsPermittedScopesEmpty(permittedScopes) {
		return insufficientPermissionsError
	}
	return nil
}

func IsPermittedScopesEmpty(permittedScopes *authorization_api.PermittedScopes) bool {
	if permittedScopes == nil {
		return true
	}
	if (permittedScopes.System == nil || !*permittedScopes.System) &&
		len(permittedScopes.Tenants) == 0 &&
		len(permittedScopes.Clusters) == 0 &&
		len(permittedScopes.Departments) == 0 &&
		len(permittedScopes.Projects) == 0 {
		return true
	}

	return false
}

func GetPermittedScopesDbFilter(permittedScopes authorization_api.PermittedScopes) string {
	return GetPermittedScopesDbFilterWithColumnNames(
		permittedScopes,
		"tenant_id",
		"cluster_id",
		"department_id",
		"project_id",
	)
}

func GetPermittedScopesDbFilterWithColumnNames(
	permittedScopes authorization_api.PermittedScopes,
	tenantIdColumnName string,
	clusterIdColumnName string,
	departmentIdColumnName string,
	projectIdColumnName string,
) string {
	if permittedScopes.System != nil && *permittedScopes.System {
		return "true"
	}

	orClause := "(false"
	if len(permittedScopes.Tenants) > 0 && tenantIdColumnName != "" {
		orClause = orClause + " OR " + tenantIdColumnName + " IN " + ListItemsForQuery(permittedScopes.Tenants)
	}
	if len(permittedScopes.Clusters) > 0 && clusterIdColumnName != "" {
		orClause = orClause + " OR " + clusterIdColumnName + " IN " + ListItemsForQuery(permittedScopes.Clusters)
	}
	if len(permittedScopes.Departments) > 0 && departmentIdColumnName != "" {
		orClause = orClause + " OR " + departmentIdColumnName + " IN " + ListItemsForQuery(permittedScopes.Departments)
	}
	if len(permittedScopes.Projects) > 0 && projectIdColumnName != "" {
		orClause = orClause + " OR " + projectIdColumnName + " IN " + ListItemsForQuery(permittedScopes.Projects)
	}

	orClause += ")"

	return orClause
}

// IsActionOnResourcePermittedByScopes assumes "permittedScopes" contains the permitted scopes for the required action.
// returns true if at least one of the ancestors is permitted.
func IsActionOnResourcePermittedByScopes(resourceAncestors *Ancestors, permittedScopes *authorization_api.PermittedScopes) bool {
	if IsPermittedScopesEmpty(permittedScopes) {
		return false
	}

	if permittedScopes.GetSystem() {
		return true
	}

	if slices.Contains(permittedScopes.GetTenants(), resourceAncestors.TenantId) {
		return true
	}

	if resourceAncestors.ClusterUuid != nil && slices.Contains(permittedScopes.GetClusters(), *resourceAncestors.ClusterUuid) {
		return true
	}

	if resourceAncestors.DepartmentId != nil && slices.Contains(permittedScopes.GetDepartments(), *resourceAncestors.DepartmentId) {
		return true
	}

	if resourceAncestors.ProjectId != nil && slices.Contains(permittedScopes.GetProjects(), *resourceAncestors.ProjectId) {
		return true
	}

	return false
}

func ListItemsForQuery(ids []string) string {
	base := "("
	for _, id := range ids {
		base += fmt.Sprintf(" '%s',", id)
	}
	return strings.TrimSuffix(base, ",") + ")"
}
