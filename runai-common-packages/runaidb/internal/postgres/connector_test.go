package postgres

import (
	"testing"
	"time"

	"github.com/dsx-ai-factory/run-ai/runai-common-packages/config"
	dbconfig "github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/config"
	"github.com/stretchr/testify/require"
)

func defaultTestConfig() *dbconfig.Config {
	return &dbconfig.Config{
		Connection: dbconfig.ConnectionConfig{
			Host:             "localhost",
			Port:             5432,
			Name:             "backend",
			Schema:           "public",
			User:             "admin",
			Password:         "secret",
			SslMode:          dbconfig.SslModeDisable,
			StatementTimeout: "2min",
		},
		Pool: dbconfig.PoolConfig{
			MaxConnections:                20,
			MaxIdleConnections:            5,
			MaxIdleConnectionTimeDuration: 30 * time.Minute,
		},
	}
}

func TestNewConnector(t *testing.T) {
	c := NewConnector()
	require.NotNil(t, c)
}

func TestBuildDSN(t *testing.T) {
	c := NewConnector()
	conf := defaultTestConfig()

	dsn := c.buildDSN(conf)
	strDsn := dsn.String()
	require.Equal(t, "postgres://admin:secret@localhost:5432/backend?search_path=public&sslmode=disable&statement_timeout=2min&timeZone=UTC", strDsn)
}

func TestRegisterSerializers_NoEncryptionKey(t *testing.T) {
	c := NewConnector()
	conf := defaultTestConfig()

	err := c.registerSerializers(conf)
	require.NoError(t, err)
}

func TestRegisterSerializers_WithEncryptionKey(t *testing.T) {
	c := NewConnector()
	conf := defaultTestConfig()
	conf.EncryptionKey = "12345678901234567890123456789012"

	require.NoError(t, config.Set(conf))
	defer config.FlushAll()

	err := c.registerSerializers(conf)
	require.NoError(t, err)
}

func TestRegisterSerializers_WithInvalidEncryptionKey(t *testing.T) {
	c := NewConnector()
	conf := defaultTestConfig()
	conf.EncryptionKey = "invalid-key"

	require.NoError(t, config.Set(conf))
	defer config.FlushAll()

	err := c.registerSerializers(conf)
	require.EqualError(t, err, "failed to create symmetric encryption serializer: encryption key must be 16, 24, or 32 bytes, got: 11")
}
