package audit_service_client

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	runaipubsubconfig "github.com/run-ai/runai-common-packages/runaipubsub/config"
	"github.com/spf13/viper"
)

const (
	MaxMetadataProperties = 10
	ClientIdMetadataKey   = "clientId"
)

type Config struct {
	PubSub                       *runaipubsubconfig.Pubsub
	Enabled                      bool   `mapstructure:"audit_enabled"`
	EnablePrintAuditLogsToStdout bool   `mapstructure:"audit_enable_print_to_stdout"`
}

func setConfigDefaults() {
	viper.SetDefault("audit_enabled", true)
	viper.SetDefault("audit_enable_print_to_stdout", false)
}

func readConfigIfNotSet(currentConfig *Config) (*Config, error) {
	if currentConfig == nil {
		return readConfig()
	}
	return currentConfig, nil
}

func validateConfig(config *Config) error {
	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return fmt.Errorf("failed to validate config: %w", err)
	}
	return nil
}

func readConfig() (*Config, error) {
	viper.AutomaticEnv()
	setConfigDefaults()
	return &Config{
		Enabled:                      viper.GetBool("audit_enabled"),
		EnablePrintAuditLogsToStdout: viper.GetBool("audit_enable_print_to_stdout"),
	}, nil
}
