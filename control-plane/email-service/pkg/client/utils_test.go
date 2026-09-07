package client_test

import (
	"testing"

	"github.com/dsx-ai-factory/run-ai/control-plane/email-service/pkg/client"
	"github.com/stretchr/testify/assert"
)

func TestStructToMap(t *testing.T) {
	objMap := client.StructToMap(client.WelcomeToRunai{
		Name:      "name",
		Username:  "username",
		Password:  "password",
		TrialLink: "trial_link",
	})
	assert.Equal(t, "name", objMap["name"])
	assert.Equal(t, "username", objMap["username"])
	assert.Equal(t, "password", objMap["password"])
	assert.Equal(t, "trial_link", objMap["trial_link"])
}

func TestGetStructName(t *testing.T) {
	name := client.GetStructName(client.WelcomeToRunai{})
	assert.Equal(t, "welcome_to_runai", name)
}
