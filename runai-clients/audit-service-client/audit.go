package audit_service_client

import (
	"context"
	"sync"

	"github.com/run-ai/runai-common-packages/logging"
	openapi "github.com/dsx-ai-factory/run-ai/runai-clients/audit-service-client/generated"
	"github.com/dsx-ai-factory/run-ai/runai-clients/audit-service-client/internal/write"
)

type contextField string

const (
	cfAuditLog contextField = "audit-log"
)

var (
	once                         sync.Once
	singleInstance               writer
	isInitialized                bool
	enablePrintAuditLogsToStdout bool
)

type writer interface {
	Write(log openapi.AuditLog) error
}

func CtxWithAudit(ctx context.Context, al *AuditLog) context.Context {
	return context.WithValue(ctx, cfAuditLog, al)
}

func AuditFromCtx(ctx context.Context) *AuditLog {
	if ctx.Value(cfAuditLog) == nil {
		return &AuditLog{}
	}
	return ctx.Value(cfAuditLog).(*AuditLog)
}

func Init(config *Config) error {
	var err error
	once.Do(func() {
		if config, err = readConfigIfNotSet(config); err != nil {
			return
		}
		if err = validateConfig(config); err != nil {
			return
		}
		enablePrintAuditLogsToStdout = config.EnablePrintAuditLogsToStdout

		singleInstance, err = createWriter(config)
		if singleInstance == nil {
			singleInstance = &write.EmptyWriter{}
		}
		if err == nil {
			isInitialized = true
		}
	})
	if err != nil {
		logging.Errorf("failed initializing audit writer. err: %+v", err)
	}
	return err
}

func IsInitialized() bool {
	return isInitialized
}

func createWriter(config *Config) (writer, error) {
	if !config.Enabled {
		return &write.EmptyWriter{}, nil
	}
	return write.NewPubSubWriter()
}
