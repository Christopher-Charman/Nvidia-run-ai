package config

import (
	"fmt"
	"time"

	"github.com/aws/smithy-go/ptr"
	"github.com/dsx-ai-factory/run-ai/runai-common-packages/config"
)

type SslMode string

const (
	SslModeDisable    SslMode = "disable"
	SslModeAllow      SslMode = "allow"
	SslModePrefer     SslMode = "prefer"
	SslModeRequire    SslMode = "require"
	SslModeVerifyCA   SslMode = "verify-ca"
	SslModeVerifyFull SslMode = "verify-full"
)

type LogLevel string

const (
	LogLevelError LogLevel = "Error"
	LogLevelWarn  LogLevel = "Warn"
	LogLevelInfo  LogLevel = "Info"
	LogLevelDebug LogLevel = "Debug"
)

type MigrationDirection string

const (
	MigrationDirectionUp    MigrationDirection = "up"
	MigrationDirectionDown  MigrationDirection = "down"
	MigrationDirectionForce MigrationDirection = "force"
)

type ConnectionConfig struct {
	Host             string  `mapstructure:"db_host" validate:"required"`
	Port             int     `mapstructure:"db_port"`
	Name             string  `mapstructure:"db_name"`
	Schema           string  `mapstructure:"db_schema" validate:"required"`
	User             string  `mapstructure:"db_user" config:"secret"`
	Password         string  `mapstructure:"db_password" config:"secret"`
	SslMode          SslMode `mapstructure:"db_ssl_mode"`
	StatementTimeout string  `mapstructure:"db_statement_timeout"`
}

type PoolConfig struct {
	MaxConnections                int           `mapstructure:"db_max_connections"`
	MaxIdleConnections            int           `mapstructure:"db_max_idle_connections"`
	MaxIdleConnectionTimeDuration time.Duration `mapstructure:"db_max_idle_connections_time_duration"`
}

type MigrationConfig struct {
	Auto          *bool              `mapstructure:"db_migration_auto"`
	Files         string             `mapstructure:"db_migration_files"`
	Direction     MigrationDirection `mapstructure:"db_migration_direction" validate:"oneof=up down force"`
	TargetVersion *string            `mapstructure:"db_migration_target_version"`
	LockRetries   int                `mapstructure:"db_migration_lock_retries" validate:"gte=0"`
}

type Config struct {
	Connection        ConnectionConfig `mapstructure:",squash"`
	Pool              PoolConfig       `mapstructure:",squash"`
	Migration         MigrationConfig  `mapstructure:",squash"`
	UseServiceSession *bool            `mapstructure:"db_use_service_session"`
	EncryptionKey     string           `mapstructure:"db_encryption_key" validate:"len=0|len=16|len=24|len=32"`
	LogLevel          LogLevel         `mapstructure:"log_level"`
	ServiceName       string           `mapstructure:"service_name" validate:"required"`
}

func (dbc *Config) SetDefaults() {
	dbc.UseServiceSession = ptr.Bool(true)
	dbc.LogLevel = LogLevelInfo
	dbc.Connection = ConnectionConfig{
		Port:             5432,
		Name:             "backend",
		StatementTimeout: "2min",
		SslMode:          SslModeDisable,
	}
	dbc.Pool = PoolConfig{
		MaxConnections:                20,
		MaxIdleConnections:            2,
		MaxIdleConnectionTimeDuration: time.Hour,
	}
	dbc.Migration = MigrationConfig{
		Files:       "file:///migrations",
		Direction:   MigrationDirectionUp,
		LockRetries: 3,
		Auto:        ptr.Bool(true),
	}
}

func InitConfig() error {
	if err := config.SetKbSecretsFolder(""); err != nil {
		return fmt.Errorf("failed to set secrets folder: %w", err)
	}

	if err := config.Register[*Config](); err != nil {
		return fmt.Errorf("failed to register runaidb config: %w", err)
	}

	return nil
}
