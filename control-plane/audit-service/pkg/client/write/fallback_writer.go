package write

import openapi "github.com/dsx-ai-factory/run-ai/control-plane/audit-service/pkg/client/generated"

//go:generate mockery --name writer --inpackage
type writer interface {
	Write(log openapi.AuditLog) error
	IsInitialized() bool
}

type FallbackWriter struct {
	primary  writer
	fallback writer
}

func NewFallbackWriter(primary, fallback writer) *FallbackWriter {
	return &FallbackWriter{primary: primary, fallback: fallback}
}

func (fw *FallbackWriter) Write(log openapi.AuditLog) error {
	if fw.primary.IsInitialized() {
		return fw.primary.Write(log)
	}
	return fw.fallback.Write(log)
}

func (fw *FallbackWriter) IsInitialized() bool {
	return fw.primary.IsInitialized() || fw.fallback.IsInitialized()
}
