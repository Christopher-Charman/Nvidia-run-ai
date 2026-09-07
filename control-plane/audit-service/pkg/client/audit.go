package client

import (
	"context"
	"encoding/json"
	"sync"
	"testing"

	"github.com/run-ai/runai-common-packages/logging"
	openapi "github.com/dsx-ai-factory/run-ai/control-plane/audit-service/pkg/client/generated"
	"github.com/dsx-ai-factory/run-ai/control-plane/audit-service/pkg/client/mocks"
	"github.com/dsx-ai-factory/run-ai/control-plane/audit-service/pkg/client/write"
)

type contextField string

const (
	cfAuditLog contextField = "audit-log"
)

var (
	once                         sync.Once
	singleInstance               Writer
	isInitialized                bool
	enablePrintAuditLogsToStdout bool
)

type Writer interface {
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

func createWriter(config *Config) (Writer, error) {
	if !config.Enabled {
		return &write.EmptyWriter{}, nil
	}
	return write.NewPubSubWriter()
}

func SetMock(t *testing.T) mocks.Writer {
	mock := mocks.NewWriter(t)
	singleInstance = mock
	return *mock
}
