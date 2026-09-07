package integrationtests

import (
	"testing"

	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb"
	"github.com/stretchr/testify/require"
)

func TestMigration_SchemaMigrationsTableExists(t *testing.T) {
	var count int64
	err := runaidb.DB().Raw(
		"SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = ? AND table_name = ?",
		"test_schema", "schema_migrations",
	).Row().Scan(&count)
	require.NoError(t, err)
	require.Equal(t, int64(1), count)
}

func TestMigration_VersionIsCorrect(t *testing.T) {
	var version int
	var dirty bool
	err := runaidb.DB().Raw("SELECT version, dirty FROM test_schema.schema_migrations").Row().Scan(&version, &dirty)
	require.NoError(t, err)
	require.Equal(t, 1, version)
	require.False(t, dirty)
}
