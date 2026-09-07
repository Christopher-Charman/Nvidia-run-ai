package internal

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	openapi "github.com/dsx-ai-factory/run-ai/runai-clients/email-service-client/generated"
)

type EmailClient struct {
	ApiClient *openapi.APIClient
	Provider  *EmailTemplateProvider
}

func NewEmailClient(baseURL string) *EmailClient {
	config := openapi.NewConfiguration()
	config.Servers = openapi.ServerConfigurations{{URL: baseURL}}
	return &EmailClient{
		ApiClient: openapi.NewAPIClient(config),
		Provider:  NewEmailTemplateProvider(),
	}
}

func (client *EmailClient) Send(ctx context.Context, to []string, templateData interface{}) error {

	request := openapi.NewExportEmailRequest(to, client.Provider.GetTemplate(templateData))
	sendEmailResponse, err := client.ApiClient.ExportAPI.ExportEmail(ctx).ExportEmailRequest(*request).Execute()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to send email to: %v. error: %v", to, sendEmailResponse))
	}
	return nil
}
