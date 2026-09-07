package audit_service_client

import (
	"testing"

	"github.com/dsx-ai-factory/run-ai/runai-clients/audit-service-client/internal/write"
	"github.com/stretchr/testify/assert"
)

func TestInit_EnabledFalse(t *testing.T) {
	err := Init(&Config{
		Enabled: false,
	})
	assert.Nil(t, err)
	assert.IsType(t, &write.EmptyWriter{}, singleInstance)
}
