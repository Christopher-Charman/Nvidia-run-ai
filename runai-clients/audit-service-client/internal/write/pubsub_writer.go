package write

import (
	"context"

	"github.com/run-ai/runai-common-packages/logging"
	"github.com/run-ai/runai-common-packages/runaipubsub"
	openapi "github.com/dsx-ai-factory/run-ai/runai-clients/audit-service-client/generated"
)

type PubsubWriter struct {
}

func NewPubSubWriter() (*PubsubWriter, error) {
	runaipubsub.SetLogger(logging.Actual())
	err := runaipubsub.Init(nil)
	if err != nil {
		return nil, err
	}

	return &PubsubWriter{}, nil
}

func (psw *PubsubWriter) Write(log openapi.AuditLog) error {
	return runaipubsub.Publish(context.Background(), &log)
}
