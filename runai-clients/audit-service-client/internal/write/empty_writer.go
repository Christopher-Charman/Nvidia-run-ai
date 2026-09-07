package write

import (
	openapi "github.com/dsx-ai-factory/run-ai/runai-clients/audit-service-client/generated"
)

type EmptyWriter struct{}

func (d *EmptyWriter) Write(log openapi.AuditLog) error {
	return nil
}
