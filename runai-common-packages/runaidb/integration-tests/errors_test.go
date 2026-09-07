package integrationtests

import (
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/orm"
	"github.com/stretchr/testify/require"
)

const missingTableQuery = "SELECT * FROM test_schema.no_such_table"

func TestUndefinedTable_Select(t *testing.T) {
	err := runaidb.DB().Actual().Exec(missingTableQuery).Error
	require.ErrorIs(t, err, orm.ErrUndefinedTable)
}

func TestUndefinedTable_Insert(t *testing.T) {
	err := runaidb.DB().Actual().Exec(
		"INSERT INTO test_schema.no_such_table (name) VALUES (?)", "value",
	).Error
	require.ErrorIs(t, err, orm.ErrUndefinedTable)
}

func TestUndefinedTable_PreservesPgError(t *testing.T) {
	err := runaidb.DB().Actual().Exec(missingTableQuery).Error

	var pgErr *pgconn.PgError
	require.ErrorAs(t, err, &pgErr)
	require.Equal(t, "42P01", pgErr.Code)
	require.Contains(t, pgErr.Message, "no_such_table")
}

func TestUndefinedTable_IsNotOtherSentinels(t *testing.T) {
	err := runaidb.DB().Actual().Exec(missingTableQuery).Error

	require.NotErrorIs(t, err, orm.ErrReadOnlyConnection)
	require.NotErrorIs(t, err, orm.ErrRecordNotFound)
}

func TestReadOnlyConnection_PreservesPgError(t *testing.T) {
	err := requireReadDB(t).Actual().Exec(
		"INSERT INTO test_schema.test_secrets (name, secret) VALUES (?, ?)",
		"read-only-pgerror", "value",
	).Error

	var pgErr *pgconn.PgError
	require.ErrorAs(t, err, &pgErr)
	require.Equal(t, "25006", pgErr.Code)
}
