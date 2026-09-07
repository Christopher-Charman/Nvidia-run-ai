# \UserApplicationsAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserApplication**](UserApplicationsAPI.md#CreateUserApplication) | **Post** /api/v1/user-applications | Create a user application.
[**DeleteUserApplicationById**](UserApplicationsAPI.md#DeleteUserApplicationById) | **Delete** /api/v1/user-applications/{appId} | Delete a user application byid.
[**DeleteUserApplicationByIdAdministration**](UserApplicationsAPI.md#DeleteUserApplicationByIdAdministration) | **Delete** /api/v1/administration/user-applications/{appId} | Delete a user application by id for adminstrations.
[**GetUserApplicationById**](UserApplicationsAPI.md#GetUserApplicationById) | **Get** /api/v1/user-applications/{appId} | Get user application by id.
[**GetUserApplications**](UserApplicationsAPI.md#GetUserApplications) | **Get** /api/v1/user-applications | Get a list of users applications.
[**GetUserApplicationsAdministration**](UserApplicationsAPI.md#GetUserApplicationsAdministration) | **Get** /api/v1/administration/user-applications | Get a list of all users applications.
[**RegenerateUserApplicationSecret**](UserApplicationsAPI.md#RegenerateUserApplicationSecret) | **Post** /api/v1/user-applications/{appId}/secret | Regenerate a user application secret.



## CreateUserApplication

> UserAppPostResponse CreateUserApplication(ctx).UserAppCreationRequest(userAppCreationRequest).Execute()

Create a user application.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/control-plane/identity-manager/pkg/client"
)

func main() {
	userAppCreationRequest := *openapiclient.NewUserAppCreationRequest("Name_example") // UserAppCreationRequest | Application object to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserApplicationsAPI.CreateUserApplication(context.Background()).UserAppCreationRequest(userAppCreationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApplicationsAPI.CreateUserApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserApplication`: UserAppPostResponse
	fmt.Fprintf(os.Stdout, "Response from `UserApplicationsAPI.CreateUserApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAppCreationRequest** | [**UserAppCreationRequest**](UserAppCreationRequest.md) | Application object to create | 

### Return type

[**UserAppPostResponse**](UserAppPostResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserApplicationById

> DeleteUserApplicationById(ctx, appId).Execute()

Delete a user application byid.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/control-plane/identity-manager/pkg/client"
)

func main() {
	appId := "appId_example" // string | The application id to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserApplicationsAPI.DeleteUserApplicationById(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApplicationsAPI.DeleteUserApplicationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The application id to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserApplicationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserApplicationByIdAdministration

> DeleteUserApplicationByIdAdministration(ctx, appId).Execute()

Delete a user application by id for adminstrations.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/control-plane/identity-manager/pkg/client"
)

func main() {
	appId := "appId_example" // string | The application id to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserApplicationsAPI.DeleteUserApplicationByIdAdministration(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApplicationsAPI.DeleteUserApplicationByIdAdministration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The application id to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserApplicationByIdAdministrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserApplicationById

> UserApp GetUserApplicationById(ctx, appId).Execute()

Get user application by id.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/control-plane/identity-manager/pkg/client"
)

func main() {
	appId := "appId_example" // string | The application id to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserApplicationsAPI.GetUserApplicationById(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApplicationsAPI.GetUserApplicationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserApplicationById`: UserApp
	fmt.Fprintf(os.Stdout, "Response from `UserApplicationsAPI.GetUserApplicationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The application id to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserApplicationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserApp**](UserApp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserApplications

> []UserApp GetUserApplications(ctx).Execute()

Get a list of users applications.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/control-plane/identity-manager/pkg/client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserApplicationsAPI.GetUserApplications(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApplicationsAPI.GetUserApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserApplications`: []UserApp
	fmt.Fprintf(os.Stdout, "Response from `UserApplicationsAPI.GetUserApplications`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserApplicationsRequest struct via the builder pattern


### Return type

[**[]UserApp**](UserApp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserApplicationsAdministration

> []UserAppAdministration GetUserApplicationsAdministration(ctx).ClientId(clientId).CreatedBy(createdBy).Execute()

Get a list of all users applications.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/control-plane/identity-manager/pkg/client"
)

func main() {
	clientId := "3434c-12323ab-4ce2ea6-b44fc-db344d" // string | Filter results by clientId (optional)
	createdBy := "test@run.ai" // string | Filter results by creatdBy (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserApplicationsAPI.GetUserApplicationsAdministration(context.Background()).ClientId(clientId).CreatedBy(createdBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApplicationsAPI.GetUserApplicationsAdministration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserApplicationsAdministration`: []UserAppAdministration
	fmt.Fprintf(os.Stdout, "Response from `UserApplicationsAPI.GetUserApplicationsAdministration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserApplicationsAdministrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | Filter results by clientId | 
 **createdBy** | **string** | Filter results by creatdBy | 

### Return type

[**[]UserAppAdministration**](UserAppAdministration.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegenerateUserApplicationSecret

> RegenerateAccessKeySecret200Response RegenerateUserApplicationSecret(ctx, appId).Execute()

Regenerate a user application secret.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/control-plane/identity-manager/pkg/client"
)

func main() {
	appId := "appId_example" // string | The application id to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserApplicationsAPI.RegenerateUserApplicationSecret(context.Background(), appId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserApplicationsAPI.RegenerateUserApplicationSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegenerateUserApplicationSecret`: RegenerateAccessKeySecret200Response
	fmt.Fprintf(os.Stdout, "Response from `UserApplicationsAPI.RegenerateUserApplicationSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | The application id to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegenerateUserApplicationSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RegenerateAccessKeySecret200Response**](RegenerateAccessKeySecret200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

