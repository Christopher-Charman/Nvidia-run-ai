package email_service_client

import (
	"context"

	client_utils "github.com/run-ai/runai-application-packages/client-utils"
	openapi "github.com/dsx-ai-factory/run-ai/runai-clients/email-service-client/generated"

	"github.com/dsx-ai-factory/run-ai/runai-clients/email-service-client/internal"
)

//go:generate mockery --name TemplateProvider
type TemplateProvider interface {
	GetTemplate(templateData interface{}) openapi.Template
}

func NewEmailTemplateProvider() TemplateProvider {
	return internal.NewEmailTemplateProvider()
}

//go:generate mockery --name IEmailExternalClient
type IEmailExternalClient interface {
}

//go:generate mockery --name IEmailClient
type IEmailClient interface {
	IEmailExternalClient
	Send(ctx context.Context, to []string, templateData interface{}) error
}

func NewEmailClient(baseURL string) IEmailClient {
	return internal.NewEmailClient(client_utils.GetInternalUrl(baseURL))
}

func NewEmailExternalClient(baseURL string) IEmailExternalClient {
	return NewEmailClient(client_utils.GetExternalUrl(baseURL))
}
