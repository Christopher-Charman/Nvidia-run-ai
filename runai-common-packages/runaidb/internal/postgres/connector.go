package postgres

import (
	"fmt"
	"net/url"

	dbconfig "github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/config"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/internal/gormext"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/internal/logger"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/orm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	dbmetrics "github.com/dsx-ai-factory/run-ai/runai-common-packages/metrics/db"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/tracing"
	oteldb "github.com/dsx-ai-factory/run-ai/runai-common-packages/tracing/db"
	"gorm.io/gorm/schema"
)

type Connector struct {
}

func NewConnector() *Connector {
	return &Connector{}
}

func (c *Connector) Connect(conf *dbconfig.Config) (orm.DB, error) {
	dsn := c.buildDSN(conf)
	gormConfig := &gorm.Config{
		Logger: gormext.NewLoggerAdapter(),
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn.String(),
	}), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err = c.registerPlugins(gormDB); err != nil {
		return nil, err
	}
	if err = c.registerSerializers(conf); err != nil {
		return nil, err
	}
	if err = c.configurePool(gormDB, conf); err != nil {
		return nil, err
	}

	return orm.NewDB(gormDB), nil
}

func (c *Connector) registerPlugins(db *gorm.DB) error {
	dbMetricsPlugin := dbmetrics.DBMiddleware(func(e error) {
		logger.GetLogger().Errorf("database metrics middleware error: %v", e)
	})
	if err := db.Use(dbMetricsPlugin); err != nil {
		return fmt.Errorf("failed to add db metrics plugin to db: %w", err)
	}

	dbOtelPlugin, ok := oteldb.DBMiddleware(tracing.GetTracerProvider())
	if !ok {
		return nil
	}
	if err := db.Use(dbOtelPlugin); err != nil {
		return fmt.Errorf("failed to add otel plugin to db: %w", err)
	}

	return nil
}

func (c *Connector) registerSerializers(conf *dbconfig.Config) error {
	if conf.EncryptionKey == "" {
		return nil
	}

	encSerializer, err := gormext.NewSymmetricEncryptionSerializer()
	if err != nil {
		return fmt.Errorf("failed to create symmetric encryption serializer: %w", err)
	}

	schema.RegisterSerializer("encrypted", encSerializer)
	return nil
}

func (c *Connector) configurePool(gormDB *gorm.DB, conf *dbconfig.Config) error {
	sqlDB, err := gormDB.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying sql.DB: %w", err)
	}

	sqlDB.SetMaxOpenConns(conf.Pool.MaxConnections)
	sqlDB.SetMaxIdleConns(conf.Pool.MaxIdleConnections)
	sqlDB.SetConnMaxIdleTime(conf.Pool.MaxIdleConnectionTimeDuration)

	return nil
}

func (c *Connector) buildDSN(conf *dbconfig.Config) url.URL {
	flags := url.Values{}
	flags.Set("timeZone", "UTC")
	flags.Set("sslmode", string(conf.Connection.SslMode))
	flags.Set("statement_timeout", conf.Connection.StatementTimeout)
	flags.Set("search_path", conf.Connection.Schema)

	return url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(conf.Connection.User, conf.Connection.Password),
		Host:     fmt.Sprintf("%s:%d", conf.Connection.Host, conf.Connection.Port),
		Path:     conf.Connection.Name,
		RawQuery: flags.Encode(),
	}
}
