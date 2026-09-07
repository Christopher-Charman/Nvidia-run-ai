package client

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/run-ai/runai-common-packages/logging"
	openapi "github.com/dsx-ai-factory/run-ai/control-plane/audit-service/pkg/client/generated"
	"github.com/dsx-ai-factory/run-ai/control-plane/audit-service/pkg/client/write"
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

func AuditFromAudit(al *AuditLog) *AuditLog {
	var copied openapi.AuditLog
	b, _ := json.Marshal(al.Log)
	_ = json.Unmarshal(b, &copied)
	return &AuditLog{Log: copied}
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

	pubsubWriter, err := write.NewPubSubWriter()
	if err != nil {
		return nil, err
	}

	auditClient := NewAuditClient(config.AuditServiceUrl)
	apiWriter := write.NewAPIWriter(auditClient)

	return write.NewFallbackWriter(pubsubWriter, apiWriter), nil
}
