package postgres

import (
	"fmt"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	migratepostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
	dbconfig "github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/config"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/internal/logger"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/orm"
)

type locker interface {
	Acquire(schema string) (bool, error)
	Release(schema string) (bool, error)
}

type Migrator struct {
	db   orm.DB
	lock locker
}

func NewMigrator(db orm.DB) *Migrator {
	return &Migrator{
		db:   db,
		lock: NewLocker(db),
	}
}

func (m *Migrator) Migrate(conf *dbconfig.Config) error {
	if err := m.ensureSchemaExists(conf.Connection.Schema); err != nil {
		return fmt.Errorf("failed to ensure schema %q exists: %w", conf.Connection.Schema, err)
	}

	acquired, err := m.lock.Acquire(conf.Connection.Schema)
	if err != nil {
		return fmt.Errorf("failed to acquire migration lock for schema %q: %w", conf.Connection.Schema, err)
	}
	if !acquired {
		return fmt.Errorf("migration lock for schema %q is already held", conf.Connection.Schema)
	}
	defer func() {
		if _, err := m.lock.Release(conf.Connection.Schema); err != nil {
			logger.GetLogger().Errorf("failed to release migration lock for schema %q: %v", conf.Connection.Schema, err)
		}
	}()

	migRunner, err := m.getMigrationRunner(conf)
	if err != nil {
		return err
	}
	return m.executeMigration(migRunner, conf)
}

func (m *Migrator) getMigrationRunner(conf *dbconfig.Config) (migrationRunner, error) {
	sqlDB, err := m.db.Actual().DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	driver, err := migratepostgres.WithInstance(sqlDB, &migratepostgres.Config{
		SchemaName: conf.Connection.Schema,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres migration driver: %w", err)
	}

	mig, err := migrate.NewWithDatabaseInstance(getMigrationFilesPath(conf), conf.Connection.Name, driver)
	if err != nil {
		return nil, fmt.Errorf("failed to create migration instance: %w", err)
	}
	mig.Log = &migrationLogger{}

	return mig, nil
}

func (m *Migrator) executeMigration(mig migrationRunner, conf *dbconfig.Config) error {
	version, err := evaluateVersion(conf.Migration.TargetVersion)
	if err != nil {
		return err
	}

	strategy, err := newMigrationStrategy(mig, conf.Migration.Direction)
	if err != nil {
		return err
	}
	if err := strategy.execute(version); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration failed: %w", err)
	}

	return nil
}

func evaluateVersion(targetVersion *string) (*uint, error) {
	if targetVersion == nil {
		return nil, nil
	}
	v, err := strconv.ParseUint(*targetVersion, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid target version %q: %w", *targetVersion, err)
	}
	version := uint(v)
	return &version, nil
}

func getMigrationFilesPath(conf *dbconfig.Config) string {
	if conf.Migration.Direction == dbconfig.MigrationDirectionDown {
		return conf.Migration.Files + "/down"
	}
	return conf.Migration.Files + "/up"
}

func (m *Migrator) ensureSchemaExists(schema string) error {
	sanitizedSchema := pgx.Identifier{schema}.Sanitize()
	query := fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s;", sanitizedSchema)
	_, err := m.db.Exec(query)
	return err
}
