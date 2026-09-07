package audit_service_client

import (
	"fmt"

	openapi "github.com/dsx-ai-factory/run-ai/runai-clients/audit-service-client/generated"
)

func validate(a *openapi.AuditLog) error {
	if a.GetTenantId() == "" {
		return fmt.Errorf("missing tenantId")
	}

	if a.GetAction() == "" {
		return fmt.Errorf("missing action")
	}

	if len(a.GetMetadata()) > MaxMetadataProperties {
		return fmt.Errorf("metadata cannot have more than %d key value pairs", MaxMetadataProperties)
	}

	return nil
}
