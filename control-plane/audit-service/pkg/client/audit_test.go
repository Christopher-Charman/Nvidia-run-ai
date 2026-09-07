package client

import (
	"testing"
	"time"

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

func TestAuditFromAudit_CopiesAllFields(t *testing.T) {
	al := NewAuditLog()
	al = al.
		WithTenantId("tenant-123").
		WithSubject("user-abc").
		WithSubjectType("User").
		WithSourceIp("10.0.0.1").
		WithAction("Create").
		WithHttpMethod("POST").
		WithUrl("/api/v1/resource").
		WithEntityType("resource").
		WithEntityName("my-resource").
		WithEntityId("entity-777").
		WithClusterName("cluster-a").
		WithClusterId("cluster-id-1").
		WithRequestId("req-999")

	al = al.WithMetadata("k1", "v1").WithMetadata("k2", "v2")

	al = al.WithStatusCode(200)

	ts := time.Date(2025, 11, 24, 12, 0, 0, 0, time.UTC)
	al = al.WithTimestamp(ts)
	al = al.WithRequestId("req-999")

	got := AuditFromAudit(al)

	assert.NotSame(t, &al.Log, &got.Log)
	assert.Equal(t, &al.Log, &got.Log, "inner Log structs should not share the same address")
}
