package write

import (
	"testing"

	"github.com/stretchr/testify/require"

	openapi "github.com/dsx-ai-factory/run-ai/control-plane/audit-service/pkg/client/generated"
)

func TestFallbackWriter(t *testing.T) {
	mockPrimaryWriter := newMockWriter(t)
	mockFallbackWriter := newMockWriter(t)

	fallbackWriter := NewFallbackWriter(mockPrimaryWriter, mockFallbackWriter)

	require.NotNil(t, fallbackWriter.primary)
	require.NotNil(t, fallbackWriter.fallback)
	require.True(t, fallbackWriter.IsInitialized())
}

func TestFallbackWriter_Write_WhenPrimaryWriterIsInitialized(t *testing.T) {
	al := openapi.AuditLog{}

	mockPrimaryWriter := newMockWriter(t)
	mockPrimaryWriter.On("IsInitialized").Return(true)
	mockPrimaryWriter.On("Write", al).Return(nil)
	mockFallbackWriter := newMockWriter(t)

	fallbackWriter := NewFallbackWriter(mockPrimaryWriter, mockFallbackWriter)

	err := fallbackWriter.Write(al)
	require.NoError(t, err)
}

func TestFallbackWriter_Write_WhenPrimaryWriterIsNotInitialized(t *testing.T) {
	al := openapi.AuditLog{}

	mockPrimaryWriter := newMockWriter(t)
	mockPrimaryWriter.On("IsInitialized").Return(false)
	mockFallbackWriter := newMockWriter(t)
	mockFallbackWriter.On("Write", al).Return(nil)

	fallbackWriter := NewFallbackWriter(mockPrimaryWriter, mockFallbackWriter)

	err := fallbackWriter.Write(al)
	require.NoError(t, err)
}
