package postgres_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/internal/postgres"
	pgerrors "github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/internal/postgres/errors"
	ormMocks "github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/orm/mocks"
	"github.com/stretchr/testify/require"
)

const (
	testUser     = "testuser"
	testPassword = "testpass"
	testSchema   = "testschema"
)

var (
	sqlCreateUser     = fmt.Sprintf(`CREATE USER "testuser" WITH PASSWORD '%s';`, testPassword)
	sqlAlterUser      = fmt.Sprintf(`ALTER USER "testuser" WITH LOGIN PASSWORD '%s';`, testPassword)
	sqlGrantUsage     = `GRANT USAGE ON SCHEMA "testschema" TO "testuser";`
	sqlGrantTables    = `GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA "testschema" TO "testuser";`
	sqlGrantSequences = `GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA "testschema" TO "testuser";`
	sqlGrantFunctions = `GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA "testschema" TO "testuser";`
)

func userAlreadyExistsErr() error {
	return &pgconn.PgError{Code: pgerrors.ErrCodeUserAlreadyExists}
}

func setupGrantMocks(db *ormMocks.DB) {
	db.On("Exec", sqlGrantUsage).Return(int64(0), nil)
	db.On("Exec", sqlGrantTables).Return(int64(0), nil)
	db.On("Exec", sqlGrantSequences).Return(int64(0), nil)
	db.On("Exec", sqlGrantFunctions).Return(int64(0), nil)
}

func TestCreate(t *testing.T) {
	db := ormMocks.NewDB(t)
	db.On("Exec", sqlCreateUser).Return(int64(0), nil)
	setupGrantMocks(db)

	err := postgres.NewRWSessionCreator(db).Create(testUser, testPassword, testSchema)
	require.NoError(t, err)
}

func TestCreate_UserAlreadyExists_PasswordUpdated(t *testing.T) {
	db := ormMocks.NewDB(t)
	db.On("Exec", sqlCreateUser).Return(int64(0), userAlreadyExistsErr())
	db.On("Exec", sqlAlterUser).Return(int64(0), nil)
	setupGrantMocks(db)

	err := postgres.NewRWSessionCreator(db).Create(testUser, testPassword, testSchema)
	require.NoError(t, err)
}

func TestCreate_CreateUserFails(t *testing.T) {
	db := ormMocks.NewDB(t)
	db.On("Exec", sqlCreateUser).Return(int64(0), errors.New("no permission"))

	err := postgres.NewRWSessionCreator(db).Create(testUser, testPassword, testSchema)
	require.EqualError(t, err, "no permission")
}

func TestCreate_AlterUserFails(t *testing.T) {
	db := ormMocks.NewDB(t)
	db.On("Exec", sqlCreateUser).Return(int64(0), userAlreadyExistsErr())
	db.On("Exec", sqlAlterUser).Return(int64(0), errors.New("alter failed"))

	err := postgres.NewRWSessionCreator(db).Create(testUser, testPassword, testSchema)
	require.EqualError(t, err, "alter failed")
}

func TestCreate_GrantUsageFails(t *testing.T) {
	db := ormMocks.NewDB(t)
	db.On("Exec", sqlCreateUser).Return(int64(0), nil)
	db.On("Exec", sqlGrantUsage).Return(int64(0), errors.New("grant usage failed"))

	err := postgres.NewRWSessionCreator(db).Create(testUser, testPassword, testSchema)
	require.EqualError(t, err, "grant usage failed")
}

func TestCreate_GrantTablesFails(t *testing.T) {
	db := ormMocks.NewDB(t)
	db.On("Exec", sqlCreateUser).Return(int64(0), nil)
	db.On("Exec", sqlGrantUsage).Return(int64(0), nil)
	db.On("Exec", sqlGrantTables).Return(int64(0), errors.New("grant tables failed"))

	err := postgres.NewRWSessionCreator(db).Create(testUser, testPassword, testSchema)
	require.EqualError(t, err, "grant tables failed")
}

func TestCreate_GrantSequencesFails(t *testing.T) {
	db := ormMocks.NewDB(t)
	db.On("Exec", sqlCreateUser).Return(int64(0), nil)
	db.On("Exec", sqlGrantUsage).Return(int64(0), nil)
	db.On("Exec", sqlGrantTables).Return(int64(0), nil)
	db.On("Exec", sqlGrantSequences).Return(int64(0), errors.New("grant sequences failed"))

	err := postgres.NewRWSessionCreator(db).Create(testUser, testPassword, testSchema)
	require.EqualError(t, err, "grant sequences failed")
}

func TestCreate_GrantFunctionsFails(t *testing.T) {
	db := ormMocks.NewDB(t)
	db.On("Exec", sqlCreateUser).Return(int64(0), nil)
	db.On("Exec", sqlGrantUsage).Return(int64(0), nil)
	db.On("Exec", sqlGrantTables).Return(int64(0), nil)
	db.On("Exec", sqlGrantSequences).Return(int64(0), nil)
	db.On("Exec", sqlGrantFunctions).Return(int64(0), errors.New("grant functions failed"))

	err := postgres.NewRWSessionCreator(db).Create(testUser, testPassword, testSchema)
	require.EqualError(t, err, "grant functions failed")
}

func TestCreate_PasswordWithSingleQuote(t *testing.T) {
	db := ormMocks.NewDB(t)
	db.On("Exec", `CREATE USER "testuser" WITH PASSWORD 'it''s a secret';`).Return(int64(0), nil)
	setupGrantMocks(db)

	err := postgres.NewRWSessionCreator(db).Create(testUser, "it's a secret", testSchema)
	require.NoError(t, err)
}
