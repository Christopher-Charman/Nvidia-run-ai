package authorization_client

import "github.com/run-ai/runai-common-packages/config"

type Config struct {
	Url        string `mapstructure:"authorization_url"`
	isExternal bool
}

func (c *Config) SetDefaults() {
	c.Url = "http://runai-backend-authorization:8080"
}

var (
	mockClient *IAuthorizationClient
)

func Init(c *Config) {
	if c == nil {
		config.Register[*Config]()
		return
	}
	config.Set(c)
}

func SetMockAuthorizationClient(mc IAuthorizationClient) {
	mockClient = &mc
}
