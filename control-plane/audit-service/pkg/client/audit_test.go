package client

import (
	"testing"

	"github.com/dsx-ai-factory/run-ai/control-plane/audit-service/pkg/client/write"
	"github.com/stretchr/testify/assert"
)

func TestInit_EnabledFalse(t *testing.T) {
	err := Init(&Config{
		Enabled: false,
	})
	assert.Nil(t, err)
	assert.IsType(t, &write.EmptyWriter{}, singleInstance)
}
