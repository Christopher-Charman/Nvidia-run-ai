package integdb

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/aws/smithy-go/ptr"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/config"
	dbconfig "github.com/dsx-ai-factory/run-ai/runai-common-packages/runaidb/config"
	tcpostgres "github.com/testcontainers/testcontainers-go/modules/postgres"
)

type IntegdbConfig struct {
	User             string `mapstructure:"integdb_user"`
	Password         string `mapstructure:"integdb_password"`
	Name             string `mapstructure:"integdb_name"`
	Schema           string `mapstructure:"integdb_schema" validate:"required"`
	MigrationDirPath string `mapstructure:"integdb_migration_dir_path"`
	ImageVersion     string `mapstructure:"integdb_image_version"`
}

func (c *IntegdbConfig) SetDefaults() {
	c.User = "postgres"
	c.Password = "postgres"
	c.Name = "backend"
	c.Schema = "test"
	c.MigrationDirPath = "file:///migrations"
	c.ImageVersion = "16"
}

func SetupPostgresDB(conf *IntegdbConfig) (*dbconfig.Config, func(), error) {
	conf = resolveConfig(conf)

	ctx := context.Background()
	container, err := tcpostgres.Run(ctx,
		"postgres:"+conf.ImageVersion,
		tcpostgres.WithDatabase(conf.Name),
		tcpostgres.WithUsername(conf.User),
		tcpostgres.WithPassword(conf.Password),
		tcpostgres.BasicWaitStrategies(),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to start postgres container: %w", err)
	}

	cleanup := func() {
		_ = container.Terminate(ctx)
	}

	host, err := container.Host(ctx)
	if err != nil {
		cleanup()
		return nil, nil, fmt.Errorf("failed to get container host: %w", err)
	}

	port, err := container.MappedPort(ctx, "5432")
	if err != nil {
		cleanup()
		return nil, nil, fmt.Errorf("failed to get container port: %w", err)
	}

	dbConf := &dbconfig.Config{}
	dbConf.SetDefaults()

	dbConf.Connection.Host = host
	dbConf.Connection.Port = port.Int()
	dbConf.Connection.Name = conf.Name
	dbConf.Connection.Schema = conf.Schema
	dbConf.Connection.User = conf.User
	dbConf.Connection.Password = conf.Password
	dbConf.Connection.SslMode = dbconfig.SslModeDisable

	dbConf.Migration.Files = conf.MigrationDirPath
	dbConf.Migration.Auto = ptr.Bool(conf.MigrationDirPath != "")

	dbConf.ServiceName = "test-integdb"
	dbConf.EncryptionKey = generateEncryptionKey()

	return dbConf, cleanup, nil
}

func resolveConfig(conf *IntegdbConfig) *IntegdbConfig {
	if conf == nil {
		config.Register[*IntegdbConfig]()
		return config.Get[*IntegdbConfig]()
	}

	config.SetWithDefaults(conf)
	return config.Get[*IntegdbConfig]()
}

func generateEncryptionKey() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
