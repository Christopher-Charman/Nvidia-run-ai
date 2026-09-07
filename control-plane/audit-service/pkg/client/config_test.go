package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig_Defaults(t *testing.T) {
	config, _ := readConfig()
	assert.True(t, config.Enabled)
}

func TestReadConfig_Override(t *testing.T) {
	t.Setenv("AUDIT_ENABLED", "false")
	config, _ := readConfig()
	assert.False(t, config.Enabled)
}

func TestValidateConfig_Valid(t *testing.T) {
	config := &Config{}
	err := validateConfig(config)
	assert.NoError(t, err)
}

func TestValidateConfig_NoQueueName(t *testing.T) {
	config := &Config{}
	err := validateConfig(config)
	assert.Nil(t, err)
}

func TestReadConfigIfNotSet_IsSet(t *testing.T) {
	config := &Config{}
	newConfig, _ := readConfigIfNotSet(config)
	assert.Equal(t, config, newConfig)
}

func TestReadConfigIfNotSet_NotSet(t *testing.T) {
	newConfig, _ := readConfigIfNotSet(nil)
	assert.Equal(t, true, newConfig.Enabled)
}
