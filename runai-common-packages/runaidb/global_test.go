package runaidb

import (
	"errors"
	"fmt"
	"sync"
	"testing"

	"github.com/aws/smithy-go/ptr"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/config"
	dbconfig "github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/config"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/internal/logger"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/mocks"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/orm"
	ormmocks "github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/orm/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	defer cleanup()

	dbClient := orm.NewDB(nil)
	dbInitializerMock := mocks.NewDBInitializer(t)
	dbInitializerMock.On("Initialize", mock.Anything).Return(dbClient, nil)

	dbInit = dbInitializerMock

	err := Init(&dbconfig.Config{
		Connection: dbconfig.ConnectionConfig{
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",
			Password: "postgres",
			Schema:   "public",
		},
		ServiceName: "test",
	})
	require.NoError(t, err)
	require.Equal(t, dbClient, DB())
}

func TestInit_FailedToInitConfig(t *testing.T) {
	defer cleanup()

	err := Init(nil)
	require.ErrorContains(t, err, "failed to set secrets folder")
}

func TestInit_FailedToInitialize(t *testing.T) {
	defer cleanup()

	dbInitializerMock := mocks.NewDBInitializer(t)
	dbInitializerMock.On("Initialize", mock.Anything).Return(nil, errors.New("failed to initialize"))
	dbInit = dbInitializerMock

	err := Init(&dbconfig.Config{
		Connection: dbconfig.ConnectionConfig{
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",
			Password: "postgres",
			Schema:   "public",
		},
		ServiceName: "test",
	})
	require.EqualError(t, err, "failed to initialize")
	require.Nil(t, DB())
}

func TestDB(t *testing.T) {
	defer cleanup()

	primaryDB = orm.NewDB(nil)
	require.Equal(t, primaryDB, DB())
}

func TestDB_DBNotInitialized(t *testing.T) {
	defer cleanup()

	primaryDB = nil
	require.Nil(t, DB())
}

func TestNewDB(t *testing.T) {
	defer cleanup()

	conf := &dbconfig.Config{
		Connection: dbconfig.ConnectionConfig{Schema: "schema1"},
		Migration:  dbconfig.MigrationConfig{Auto: ptr.Bool(true)},
	}
	err := config.Set(conf)
	require.NoError(t, err)

	dbClient := orm.NewDB(nil)
	dbInitializerMock := mocks.NewDBInitializer(t)
	dbInitializerMock.On("Initialize", &dbconfig.Config{
		Connection: dbconfig.ConnectionConfig{Schema: "schema2"},
		Migration:  dbconfig.MigrationConfig{Auto: ptr.Bool(false)},
	}).Return(dbClient, nil)
	dbInit = dbInitializerMock

	db, err := NewDB("schema2")
	require.NoError(t, err)
	require.Equal(t, dbClient, db)

	loadedDB, ok := schemaDBs.Load("schema2")
	require.True(t, ok)
	require.Equal(t, dbClient, loadedDB.(orm.DB))
}

func TestNewDB_LoadsFromSchemaDBs(t *testing.T) {
	defer cleanup()

	dbClient := orm.NewDB(nil)
	schemaDBs.Store("schema1", dbClient)

	db, err := NewDB("schema1")
	require.NoError(t, err)
	require.Equal(t, dbClient, db)
}

func TestNewDB_LoadsFromPrimaryDB(t *testing.T) {
	defer cleanup()

	conf := &dbconfig.Config{
		Connection: dbconfig.ConnectionConfig{Schema: "primary"},
	}
	err := config.Set(conf)
	require.NoError(t, err)

	primaryDB = orm.NewDB(nil)
	db, err := NewDB("primary")
	require.NoError(t, err)
	require.Equal(t, primaryDB, db)
}

func TestMigrate(t *testing.T) {
	defer cleanup()

	conf := &dbconfig.Config{
		Migration: dbconfig.MigrationConfig{
			Auto:          ptr.Bool(false),
			Direction:     dbconfig.MigrationDirectionUp,
			TargetVersion: ptr.String("2"),
		},
		UseServiceSession: ptr.Bool(true),
	}
	err := config.Set(conf)
	require.NoError(t, err)

	primaryDB = orm.NewDB(nil)

	dbClient := orm.NewDB(nil)
	dbInitializerMock := mocks.NewDBInitializer(t)
	dbInitializerMock.On("Initialize", &dbconfig.Config{
		Migration: dbconfig.MigrationConfig{
			Auto:          ptr.Bool(true),
			Direction:     dbconfig.MigrationDirectionDown,
			TargetVersion: ptr.String("1"),
		},
		UseServiceSession: ptr.Bool(false),
	}).Return(dbClient, nil)
	dbInit = dbInitializerMock

	err = Migrate(&MigrationParams{
		Direction:     dbconfig.MigrationDirectionDown,
		TargetVersion: ptr.String("1"),
	})
	require.NoError(t, err)
}

func TestMigrate_DBNotInitialized(t *testing.T) {
	defer cleanup()

	err := Migrate(&MigrationParams{
		Direction: dbconfig.MigrationDirectionUp,
	})
	require.EqualError(t, err, "DB is not initialized")
}

func TestMigrate_InvalidParams(t *testing.T) {
	defer cleanup()

	primaryDB = orm.NewDB(nil)

	require.EqualError(t, Migrate(nil), "migration parameters are not valid")
	require.EqualError(t, Migrate(&MigrationParams{}), "migration parameters are not valid")
}

func TestMigrate_Failed(t *testing.T) {
	defer cleanup()

	conf := &dbconfig.Config{
		Connection: dbconfig.ConnectionConfig{Schema: "public"},
		Migration: dbconfig.MigrationConfig{
			Auto:      ptr.Bool(true),
			Direction: dbconfig.MigrationDirectionUp,
		},
	}
	err := config.Set(conf)
	require.NoError(t, err)

	primaryDB = orm.NewDB(nil)

	dbInitializerMock := mocks.NewDBInitializer(t)
	dbInitializerMock.On("Initialize", mock.Anything).Return(nil, errors.New("migration error"))
	dbInit = dbInitializerMock

	err = Migrate(&MigrationParams{
		Direction: dbconfig.MigrationDirectionUp,
	})
	require.EqualError(t, err, "failed to migrate: migration error")
}

func TestClose(t *testing.T) {
	defer cleanup()

	dbMock := ormmocks.NewDB(t)
	dbMock.On("Close").Return(nil)
	primaryDB = dbMock
	require.NoError(t, Close())
	require.Nil(t, primaryDB)
}

func TestClose_Failed(t *testing.T) {
	defer cleanup()

	dbMock := ormmocks.NewDB(t)
	dbMock.On("Close").Return(errors.New("failed to close"))
	primaryDB = dbMock
	require.EqualError(t, Close(), "failed to close primary DB: failed to close")
	require.Nil(t, primaryDB)
}

func TestClose_WithSchemaDBs(t *testing.T) {
	defer cleanup()

	primaryDBMock := ormmocks.NewDB(t)
	primaryDBMock.On("Close").Return(nil)
	primaryDB = primaryDBMock

	schemaDBMock1 := ormmocks.NewDB(t)
	schemaDBMock1.On("Close").Return(nil)
	schemaDBs.Store("schema1", schemaDBMock1)

	schemaDBMock2 := ormmocks.NewDB(t)
	schemaDBMock2.On("Close").Return(nil)
	schemaDBs.Store("schema2", schemaDBMock2)

	require.NoError(t, Close())
}

func TestClose_WithSchemaDBs_Failed(t *testing.T) {
	defer cleanup()

	primaryDBMock := ormmocks.NewDB(t)
	primaryDBMock.On("Close").Return(nil)
	primaryDB = primaryDBMock

	schemaDBMock1 := ormmocks.NewDB(t)
	schemaDBMock1.On("Close").Return(fmt.Errorf("failed to close"))
	schemaDBs.Store("schema1", schemaDBMock1)

	require.EqualError(t, Close(), "failed to close schema DB schema1: failed to close")
}

func TestClose_DBNotInitialized(t *testing.T) {
	defer cleanup()

	require.EqualError(t, Close(), "DB is not initialized")
}

func TestSetLogger(t *testing.T) {
	defer cleanup()

	logger.SetLogger(nil)
	require.NotNil(t, logger.GetLogger())

	myLogger := &myLogger{}
	SetLogger(myLogger)
	require.Equal(t, myLogger, logger.GetLogger())
}

func cleanup() {
	config.FlushAll()
	primaryDB = nil
	schemaDBs = sync.Map{}
	once = sync.Once{}
	dbInit = nil
}

type myLogger struct {
}

func (l *myLogger) Infof(format string, args ...any) {
}

func (l *myLogger) Errorf(format string, args ...any) {
}

func (l *myLogger) Warnf(format string, args ...any) {
}

func (l *myLogger) Debugf(format string, args ...any) {
}
