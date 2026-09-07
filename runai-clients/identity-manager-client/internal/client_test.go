package internal

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/pkg/errors"
	openapi "github.com/dsx-ai-factory/run-ai/runai-clients/identity-manager-client/generated"
	"github.com/dsx-ai-factory/run-ai/runai-clients/identity-manager-client/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"k8s.io/utils/pointer"
)

const (
	CLIENT_ID = "client_id"

	CLUSTER_DOMAIN = "cluster_domain"
	CLUSTER_ID     = "cluster_1"
)

func GetApiClient(redirectUris []string) []openapi.Client {
	attributeMap := make(map[string]interface{})

	attributeMap["runai.client_type"] = "workloads_auth"
	attributeMap["runai.cluster_id"] = CLUSTER_ID
	clientArray := []openapi.Client{
		{
			Id:           CLIENT_ID,
			Attributes:   attributeMap,
			RedirectUris: redirectUris,
		},
	}
	return clientArray
}

func initIdmService(t *testing.T) (mock *mocks.ClientsAPI, idmClient *IdmClient) {
	mock = mocks.NewClientsAPI(t)
	config := openapi.NewConfiguration()
	apiClient := openapi.NewAPIClient(config)
	apiClient.ClientsAPI = mock
	idmClient = &IdmClient{
		ApiClient: apiClient,
	}
	return
}

func TestGetClientsHappyFlow(t *testing.T) {
	clientServiceMock, testSubject := initIdmService(t)
	clients := GetApiClient(nil)
	clientServiceMock.On("GetAllClients", context.Background()).Return(openapi.ApiGetAllClientsRequest{})
	clientServiceMock.On("GetAllClientsExecute", mock.Anything).Return(clients, nil, nil)
	filter := fmt.Sprintf("%s:%s %s:%s", "runai.client_type", "workloads_auth", "runai.cluster_id", CLUSTER_ID)
	clientsRes, err := testSubject.GetInternalClientByFilter(context.Background(), filter)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(clientsRes), 1)
	assert.Equal(t, clientsRes, clients)
}

func TestGetClientsWithErr(t *testing.T) {
	clientServiceMock, testSubject := initIdmService(t)
	err := errors.Errorf("Error getting clients")
	response := http.Response{
		Status: "500",
	}
	clientServiceMock.On("GetAllClients", context.Background()).Return(openapi.ApiGetAllClientsRequest{})
	clientServiceMock.On("GetAllClientsExecute", mock.Anything).Return(nil, &response, err)
	filter := fmt.Sprintf("%s:%s %s:%s", "runai.client_type", "workloads_auth", "runai.cluster_id", CLUSTER_ID)
	clientsRes, errResponse := testSubject.GetInternalClientByFilter(context.Background(), filter)
	var expectedClientRes []openapi.Client = nil
	assert.True(t, errResponse != nil)
	assert.Equal(t, expectedClientRes, clientsRes)

}

func TestUpdateClientUriUriExistsOnClient(t *testing.T) {
	clientServiceMock, testSubject := initIdmService(t)
	redirectUris := []string{"https://test1", "https://test2"}
	clients := GetApiClient(redirectUris)
	clientServiceMock.On("UpdateRedirectUris", mock.Anything, mock.Anything).Return(openapi.ApiUpdateRedirectUrisRequest{})
	clientServiceMock.On("UpdateRedirectUrisExecute", mock.Anything).Return(nil, nil)
	expectedRedirectUris := []string{"https://test1", "https://test2", CLUSTER_DOMAIN + "/*"}
	expectedReq := getExpectedUpdateClientRedirectUrisReq(clientServiceMock, clients[0], expectedRedirectUris)
	err := testSubject.AppendRedirectUriToInternalClient(context.Background(), clients[0], CLUSTER_DOMAIN)
	assert.True(t, err == nil)
	clientServiceMock.AssertCalled(t, "UpdateRedirectUrisExecute", expectedReq)
}

func TestUpdateClientUriEmptyUrisOnClient(t *testing.T) {
	clientServiceMock, testSubject := initIdmService(t)
	clients := GetApiClient(nil)
	clientServiceMock.On("UpdateRedirectUris", mock.Anything, mock.Anything).Return(openapi.ApiUpdateRedirectUrisRequest{})
	clientServiceMock.On("UpdateRedirectUrisExecute", mock.Anything).Return(nil, nil)
	expectedRedirectUris := []string{CLUSTER_DOMAIN + "/*"}
	expectedReq := getExpectedUpdateClientRedirectUrisReq(clientServiceMock, clients[0], expectedRedirectUris)
	err := testSubject.AppendRedirectUriToInternalClient(context.Background(), clients[0], CLUSTER_DOMAIN)
	assert.True(t, err == nil)
	clientServiceMock.AssertCalled(t, "UpdateRedirectUrisExecute", expectedReq)
}

func TestUpdateClientUriWhenAlreadyExist(t *testing.T) {
	clientServiceMock, testSubject := initIdmService(t)
	redirectUris := []string{"https://test1", "https://test2", CLUSTER_DOMAIN + "/*"}
	clients := GetApiClient(redirectUris)
	err := testSubject.AppendRedirectUriToInternalClient(context.Background(), clients[0], CLUSTER_DOMAIN)
	assert.True(t, err == nil)
	clientServiceMock.AssertNotCalled(t, "UpdateRedirectUrisExecute", mock.Anything)
}

func TestCreateClient_HappyFlow(t *testing.T) {
	clientServiceMock, testSubject := initIdmService(t)
	resp := &openapi.CreateInternalClient201Response{
		Id: pointer.String("ID"),
	}
	httpResp := &http.Response{
		StatusCode: http.StatusOK,
	}
	clientServiceMock.On("CreateInternalClient", mock.Anything).Return(openapi.ApiCreateInternalClientRequest{})
	clientServiceMock.On("CreateInternalClientExecute", mock.Anything).Return(resp, httpResp, nil)
	clientId, err := testSubject.CreateClient(context.Background(), "name", "Application", pointer.String("ID"))
	assert.Equal(t, "ID", *clientId)
	assert.Nil(t, err)
}

func TestCreateClient_Error(t *testing.T) {
	clientServiceMock, testSubject := initIdmService(t)
	resp := &openapi.CreateInternalClient201Response{
		Id: pointer.String("ID"),
	}
	clientServiceMock.On("CreateInternalClient", mock.Anything).Return(openapi.ApiCreateInternalClientRequest{})
	clientServiceMock.On("CreateInternalClientExecute", mock.Anything).Return(resp, nil, fmt.Errorf("error"))
	clientId, err := testSubject.CreateClient(context.Background(), "name", "Application", pointer.String("ID"))
	assert.Nil(t, clientId)
	assert.NotNil(t, err)
}

func TestCreateClient_HttpError(t *testing.T) {
	clientServiceMock, testSubject := initIdmService(t)
	resp := &openapi.CreateInternalClient201Response{
		Id: pointer.String("ID"),
	}
	httpResp := &http.Response{
		StatusCode: http.StatusInternalServerError,
	}
	clientServiceMock.On("CreateInternalClient", mock.Anything).Return(openapi.ApiCreateInternalClientRequest{})
	clientServiceMock.On("CreateInternalClientExecute", mock.Anything).Return(resp, httpResp, fmt.Errorf("error"))
	clientId, err := testSubject.CreateClient(context.Background(), "name", "Application", pointer.String("ID"))
	assert.Nil(t, clientId)
	assert.NotNil(t, err)
}

func getExpectedUpdateClientRedirectUrisReq(mock *mocks.ClientsAPI, client openapi.Client, expectedRedirectUris []string) openapi.ApiUpdateRedirectUrisRequest {
	expectedInnerReq := openapi.NewClientRedirectUrisUpdateRequest(expectedRedirectUris)
	expectedUpdateUriReq := mock.UpdateRedirectUris(context.Background(), client.ClientId)
	expectedUpdateUriReq = expectedUpdateUriReq.ClientRedirectUrisUpdateRequest(*expectedInnerReq)
	return expectedUpdateUriReq
}
