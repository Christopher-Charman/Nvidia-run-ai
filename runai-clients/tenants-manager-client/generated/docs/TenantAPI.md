# \TenantAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTenant**](TenantAPI.md#CreateTenant) | **Post** /api/v1/tenants | Create tenant
[**DeleteTenantById**](TenantAPI.md#DeleteTenantById) | **Delete** /api/v1/tenants/{tenantId} | soft/hard delete tenant
[**GetTenantById**](TenantAPI.md#GetTenantById) | **Get** /api/v1/tenants/{tenantId} | Get tenant
[**GetTenants**](TenantAPI.md#GetTenants) | **Get** /api/v1/tenants | Get a list of tenants.
[**PostTenantEula**](TenantAPI.md#PostTenantEula) | **Post** /api/v1/sign-eula | Post tenant eula



## CreateTenant

> TenantCreationResponse CreateTenant(ctx).TenantCreationRequest(tenantCreationRequest).Execute()

Create tenant

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/tenants-manager-client"
)

func main() {
	tenantCreationRequest := *openapiclient.NewTenantCreationRequest("name", "first@run.ai") // TenantCreationRequest | The tenant to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.CreateTenant(context.Background()).TenantCreationRequest(tenantCreationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.CreateTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTenant`: TenantCreationResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.CreateTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantCreationRequest** | [**TenantCreationRequest**](TenantCreationRequest.md) | The tenant to create. | 

### Return type

[**TenantCreationResponse**](TenantCreationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTenantById

> DeleteTenantById200Response DeleteTenantById(ctx, tenantId).IsHardDelete(isHardDelete).Execute()

soft/hard delete tenant

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/tenants-manager-client"
)

func main() {
	tenantId := int32(56) // int32 | The id of the tenant to retrieve
	isHardDelete := true // bool | Boolean indicator for soft/hard delete (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.DeleteTenantById(context.Background(), tenantId).IsHardDelete(isHardDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.DeleteTenantById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTenantById`: DeleteTenantById200Response
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.DeleteTenantById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **int32** | The id of the tenant to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isHardDelete** | **bool** | Boolean indicator for soft/hard delete | 

### Return type

[**DeleteTenantById200Response**](DeleteTenantById200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantById

> Tenant GetTenantById(ctx, tenantId).Execute()

Get tenant

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/tenants-manager-client"
)

func main() {
	tenantId := int32(56) // int32 | The id of the tenant to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.GetTenantById(context.Background(), tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetTenantById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantById`: Tenant
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetTenantById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **int32** | The id of the tenant to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tenant**](Tenant.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenants

> []Tenant GetTenants(ctx).Execute()

Get a list of tenants.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/tenants-manager-client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantAPI.GetTenants(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.GetTenants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenants`: []Tenant
	fmt.Fprintf(os.Stdout, "Response from `TenantAPI.GetTenants`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantsRequest struct via the builder pattern


### Return type

[**[]Tenant**](Tenant.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostTenantEula

> PostTenantEula(ctx).PostTenantEulaRequest(postTenantEulaRequest).Execute()

Post tenant eula

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/tenants-manager-client"
)

func main() {
	postTenantEulaRequest := *openapiclient.NewPostTenantEulaRequest("test@run.ai") // PostTenantEulaRequest | The user email that signing the eula.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TenantAPI.PostTenantEula(context.Background()).PostTenantEulaRequest(postTenantEulaRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantAPI.PostTenantEula``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostTenantEulaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postTenantEulaRequest** | [**PostTenantEulaRequest**](PostTenantEulaRequest.md) | The user email that signing the eula. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

