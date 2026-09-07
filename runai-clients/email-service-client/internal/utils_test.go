package internal_test

import (
	"testing"

	templates "github.com/dsx-ai-factory/run-ai/runai-clients/email-service-client"
	utils "github.com/dsx-ai-factory/run-ai/runai-clients/email-service-client/internal"
	"github.com/stretchr/testify/assert"
)

func TestStructToMap(t *testing.T) {
	objMap := utils.StructToMap(templates.WelcomeToRunai{
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
	name := utils.GetStructName(templates.WelcomeToRunai{})
	assert.Equal(t, "welcome_to_runai", name)
}
