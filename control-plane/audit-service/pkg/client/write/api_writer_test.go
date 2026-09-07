package write

import (
	"errors"
	"testing"
	"time"

	openapi "github.com/dsx-ai-factory/run-ai/control-plane/audit-service/pkg/client/generated"
	"github.com/dsx-ai-factory/run-ai/control-plane/audit-service/pkg/client/mocks"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAPIWriter(t *testing.T) {
	mockClient := mocks.NewIAuditClient(t)
	apiWriter := NewAPIWriter(mockClient)

	require.NotNil(t, apiWriter.client)
	require.True(t, apiWriter.IsInitialized())
}

func TestAPIWriter_Write(t *testing.T) {
	al := openapi.AuditLog{}
	mockClient := mocks.NewIAuditClient(t)
	mockClient.On("SaveAuditLog", mock.Anything, al).Return(nil)
	apiWriter := NewAPIWriter(mockClient)

	err := apiWriter.Write(al)
	require.NoError(t, err)

	time.Sleep(100 * time.Millisecond) // wait for the go routine to finish
}

func TestAPIWriter_ShouldRetryOnSaveAuditLogError(t *testing.T) {
	al := openapi.AuditLog{}
	mockClient := mocks.NewIAuditClient(t)
	mockClient.On("SaveAuditLog", mock.Anything, al).Return(errors.New("failed to save audit log")).Once()
	mockClient.On("SaveAuditLog", mock.Anything, al).Return(nil).Once()
	apiWriter := &APIWriter{
		client:          mockClient,
		maxWriteRetries: 2,
		writeRetryDelay: 10 * time.Millisecond,
	}

	err := apiWriter.Write(al)
	require.NoError(t, err) // function should still return nil, since we're using a go routine

	time.Sleep(100 * time.Millisecond) // wait for the go routine to finish
}
