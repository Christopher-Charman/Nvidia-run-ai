package client

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	client_utils "github.com/run-ai/runai-application-packages/client-utils"
	openapi "github.com/dsx-ai-factory/run-ai/control-plane/email-service/pkg/client/generated"
)

type EmailClient struct {
	ApiClient *openapi.APIClient
	Provider  *EmailTemplateProvider
}

func NewEmailClient(baseURL string) *EmailClient {
	baseURL = client_utils.GetInternalUrl(baseURL)
	config := openapi.NewConfiguration()
	config.Servers = openapi.ServerConfigurations{{URL: baseURL}}
	return &EmailClient{
		ApiClient: openapi.NewAPIClient(config),
		Provider:  NewEmailTemplateProvider(),
	}
}

func (ec *EmailClient) Send(ctx context.Context, to []string, templateData interface{}) error {

	request := openapi.NewExportEmailRequest(to, ec.Provider.GetTemplate(templateData))
	sendEmailResponse, err := ec.ApiClient.ExportAPI.ExportEmail(ctx).ExportEmailRequest(*request).Execute()
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to send email to: %v. error: %v", to, sendEmailResponse))
	}
	return nil
}
