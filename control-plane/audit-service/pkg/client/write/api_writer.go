package write

import (
	"context"
	"time"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/strategy"
	"github.com/run-ai/runai-common-packages/logging"
	openapi "github.com/dsx-ai-factory/run-ai/control-plane/audit-service/pkg/client/generated"
)

const (
	maxWriteRetries = 3
	writeRetryDelay = 2 * time.Second
)

type AuditSaver interface {
	SaveAuditLog(ctx context.Context, auditLog openapi.AuditLog) error
}

type APIWriter struct {
	client          AuditSaver
	maxWriteRetries uint
	writeRetryDelay time.Duration
}

func NewAPIWriter(client AuditSaver) *APIWriter {
	return &APIWriter{client: client, maxWriteRetries: maxWriteRetries, writeRetryDelay: writeRetryDelay}
}

func (w *APIWriter) Write(log openapi.AuditLog) error {
	go func() {
		err := retry.Retry(
			func(u uint) error {
				return w.client.SaveAuditLog(context.Background(), log)
			},
			strategy.Wait(w.writeRetryDelay),
			strategy.Limit(w.maxWriteRetries),
		)
		if err != nil {
			logging.Warnf("Failed writing audit log using API writer: %v", err)
		}
	}()
	return nil
}

func (w *APIWriter) IsInitialized() bool {
	return true
}
