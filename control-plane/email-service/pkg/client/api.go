package client

import (
	"context"

	client_utils "github.com/run-ai/runai-application-packages/client-utils"
	openapi "github.com/dsx-ai-factory/run-ai/control-plane/email-service/pkg/client/generated"
)

//go:generate mockery --name TemplateProvider
type TemplateProvider interface {
	GetTemplate(templateData interface{}) openapi.Template
}

//go:generate mockery --name IEmailExternalClient
type IEmailExternalClient interface {
}

//go:generate mockery --name IEmailClient
type IEmailClient interface {
	IEmailExternalClient
	Send(ctx context.Context, to []string, templateData interface{}) error
}

func NewEmailExternalClient(baseURL string) IEmailExternalClient {
	return NewEmailClient(client_utils.GetExternalUrl(baseURL))
}
