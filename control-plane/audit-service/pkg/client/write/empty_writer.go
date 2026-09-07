package write

import (
	openapi "github.com/dsx-ai-factory/run-ai/control-plane/audit-service/pkg/client/generated"
)

type EmptyWriter struct{}

func (d *EmptyWriter) Write(log openapi.AuditLog) error {
	return nil
}

func (d *EmptyWriter) IsInitialized() bool {
	return true
}
