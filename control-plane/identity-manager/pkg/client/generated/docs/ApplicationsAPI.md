# \ApplicationsAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplicationInternally**](ApplicationsAPI.md#CreateApplicationInternally) | **Post** /internal/apps | Create an application internallay
[**GetInstallerApp**](ApplicationsAPI.md#GetInstallerApp) | **Get** /api/v1/apps/installer | Get a list of installer applications.



## CreateApplicationInternally

> ApplicationPostResponse CreateApplicationInternally(ctx).InternalApplicationCreationRequest(internalApplicationCreationRequest).Execute()

Create an application internallay



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
	internalApplicationCreationRequest := *openapiclient.NewInternalApplicationCreationRequest("Name_example", "Realm_example") // InternalApplicationCreationRequest | Application object to create

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.CreateApplicationInternally(context.Background()).InternalApplicationCreationRequest(internalApplicationCreationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.CreateApplicationInternally``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplicationInternally`: ApplicationPostResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.CreateApplicationInternally`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationInternallyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **internalApplicationCreationRequest** | [**InternalApplicationCreationRequest**](InternalApplicationCreationRequest.md) | Application object to create | 

### Return type

[**ApplicationPostResponse**](ApplicationPostResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstallerApp

> ApplicationPostResponse GetInstallerApp(ctx).TenantId(tenantId).Execute()

Get a list of installer applications.



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
	tenantId := "1008" // string | If exists will return the relevant resource for the tenant ID (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationsAPI.GetInstallerApp(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationsAPI.GetInstallerApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstallerApp`: ApplicationPostResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationsAPI.GetInstallerApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInstallerAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | If exists will return the relevant resource for the tenant ID | [default to &quot;&quot;]

### Return type

[**ApplicationPostResponse**](ApplicationPostResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

