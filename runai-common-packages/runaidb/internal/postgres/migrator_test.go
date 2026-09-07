package postgres

import (
	"fmt"
	"testing"

	"github.com/aws/smithy-go/ptr"
	"github.com/golang-migrate/migrate/v4"
	dbconfig "github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/config"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/internal/logger"
	logmocks "github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/internal/logger/mocks"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/internal/postgres/mocks"
	ormmocks "github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/orm/mocks"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func defaultMigrationConfig(direction dbconfig.MigrationDirection) *dbconfig.Config {
	return &dbconfig.Config{
		Connection: dbconfig.ConnectionConfig{Schema: "public", Name: "mydb"},
		Migration:  dbconfig.MigrationConfig{Direction: direction, Files: "file:///tmp/migrations"},
	}
}

func TestMigrate_EnsureSchemaError(t *testing.T) {
	db := ormmocks.NewDB(t)
	db.On("Exec", `CREATE SCHEMA IF NOT EXISTS "public";`).Return(int64(0), fmt.Errorf("schema error"))

	m := &Migrator{db: db, lock: mocks.NewLocker(t)}
	err := m.Migrate(defaultMigrationConfig(dbconfig.MigrationDirectionUp))
	require.ErrorContains(t, err, `failed to ensure schema "public" exists`)
}

func TestMigrate_AcquireError(t *testing.T) {
	db := ormmocks.NewDB(t)
	db.On("Exec", `CREATE SCHEMA IF NOT EXISTS "public";`).Return(int64(0), nil)

	lock := mocks.NewLocker(t)
	lock.On("Acquire", "public").Return(false, fmt.Errorf("db error"))

	m := &Migrator{db: db, lock: lock}
	err := m.Migrate(defaultMigrationConfig(dbconfig.MigrationDirectionUp))
	require.EqualError(t, err, "failed to acquire migration lock for schema \"public\": db error")
}

func TestMigrate_LockAlreadyHeld(t *testing.T) {
	db := ormmocks.NewDB(t)
	db.On("Exec", `CREATE SCHEMA IF NOT EXISTS "public";`).Return(int64(0), nil)

	lock := mocks.NewLocker(t)
	lock.On("Acquire", "public").Return(false, nil)

	m := &Migrator{db: db, lock: lock}
	err := m.Migrate(defaultMigrationConfig(dbconfig.MigrationDirectionUp))
	require.EqualError(t, err, "migration lock for schema \"public\" is already held")
}

func TestMigrate_ReleaseError_IsLogged(t *testing.T) {
	log := logmocks.NewLogger(t)
	log.On("Errorf", "failed to release migration lock for schema %q: %v", "public", fmt.Errorf("release failed")).Return()
	logger.SetLogger(log)

	db := ormmocks.NewDB(t)
	db.On("Exec", `CREATE SCHEMA IF NOT EXISTS "public";`).Return(int64(0), nil)
	db.On("Actual").Return(&gorm.DB{Config: &gorm.Config{ConnPool: nil}})

	lock := mocks.NewLocker(t)
	lock.On("Acquire", "public").Return(true, nil)
	lock.On("Release", "public").Return(false, fmt.Errorf("release failed"))

	m := &Migrator{db: db, lock: lock}
	err := m.Migrate(defaultMigrationConfig(dbconfig.MigrationDirectionUp))
	require.ErrorContains(t, err, "failed to get underlying sql.DB")
}

func TestExecuteMigration_InvalidVersion(t *testing.T) {
	m := &Migrator{}
	conf := defaultMigrationConfig(dbconfig.MigrationDirectionUp)
	conf.Migration.TargetVersion = ptr.String("bad")

	err := m.executeMigration(nil, conf)
	require.ErrorContains(t, err, `invalid target version "bad"`)
}

func TestExecuteMigration_UnknownDirection(t *testing.T) {
	m := &Migrator{}
	conf := defaultMigrationConfig("unknown")

	err := m.executeMigration(nil, conf)
	require.ErrorContains(t, err, `unknown migration direction "unknown"`)
}

func TestExecuteMigration_ErrNoChange(t *testing.T) {
	mig := mocks.NewMigrationRunner(t)
	mig.On("Up").Return(migrate.ErrNoChange)

	m := &Migrator{}
	err := m.executeMigration(mig, defaultMigrationConfig(dbconfig.MigrationDirectionUp))
	require.NoError(t, err)
}

func TestExecuteMigration_StrategyError(t *testing.T) {
	mig := mocks.NewMigrationRunner(t)
	mig.On("Up").Return(fmt.Errorf("up failed"))

	m := &Migrator{}
	err := m.executeMigration(mig, defaultMigrationConfig(dbconfig.MigrationDirectionUp))
	require.ErrorContains(t, err, "migration failed: up failed")
}

func TestExecuteMigration_Success(t *testing.T) {
	mig := mocks.NewMigrationRunner(t)
	mig.On("Up").Return(nil)

	m := &Migrator{}
	err := m.executeMigration(mig, defaultMigrationConfig(dbconfig.MigrationDirectionUp))
	require.NoError(t, err)
}

func TestEvaluateVersion_Nil(t *testing.T) {
	v, err := evaluateVersion(nil)
	require.NoError(t, err)
	require.Nil(t, v)
}

func TestEvaluateVersion_Valid(t *testing.T) {
	v, err := evaluateVersion(ptr.String("42"))
	require.NoError(t, err)
	require.NotNil(t, v)
	require.Equal(t, uint(42), *v)
}

func TestEvaluateVersion_Invalid(t *testing.T) {
	_, err := evaluateVersion(ptr.String("not-a-number"))
	require.ErrorContains(t, err, `invalid target version "not-a-number"`)
}

func TestGetMigrationFilesPath_Up(t *testing.T) {
	path := getMigrationFilesPath(defaultMigrationConfig(dbconfig.MigrationDirectionUp))
	require.Equal(t, "file:///tmp/migrations/up", path)
}

func TestGetMigrationFilesPath_Down(t *testing.T) {
	path := getMigrationFilesPath(defaultMigrationConfig(dbconfig.MigrationDirectionDown))
	require.Equal(t, "file:///tmp/migrations/down", path)
}
