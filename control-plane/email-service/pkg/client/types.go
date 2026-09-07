package client

import openapi "github.com/dsx-ai-factory/run-ai/control-plane/email-service/pkg/client/generated"

type EmailTemplateProvider struct {
}

func NewEmailTemplateProvider() *EmailTemplateProvider {
	return &EmailTemplateProvider{}
}

func (etp *EmailTemplateProvider) getId(templateData interface{}) string {
	return GetStructName(templateData)
}

func (etp *EmailTemplateProvider) getData(templateData interface{}) map[string]interface{} {
	return StructToMap(templateData)
}

func (etp *EmailTemplateProvider) GetTemplate(templateData interface{}) openapi.Template {
	return openapi.Template{Id: etp.getId(templateData), Data: etp.getData(templateData)}
}
