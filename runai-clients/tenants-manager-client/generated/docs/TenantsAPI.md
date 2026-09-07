# \TenantsAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTenantInternal**](TenantsAPI.md#CreateTenantInternal) | **Post** /internal/api/v1/tenants | Create tenant
[**DeleteTenantByIdInternal**](TenantsAPI.md#DeleteTenantByIdInternal) | **Delete** /internal/api/v1/tenants/{tenantId} | soft/hard delete tenant
[**GetLicenseQuery**](TenantsAPI.md#GetLicenseQuery) | **Get** /internal/api/v1/tenants/{tenantId}/license/query | get license query api
[**GetTenant**](TenantsAPI.md#GetTenant) | **Get** /internal/api/v1/tenants | Get tenant
[**GetTenantByIdInternal**](TenantsAPI.md#GetTenantByIdInternal) | **Get** /internal/api/v1/tenants/{tenantId} | Get tenant by id
[**GetTenantLicense**](TenantsAPI.md#GetTenantLicense) | **Get** /internal/api/v1/tenants/{tenantId}/license | Get tenant license
[**UploadLicense**](TenantsAPI.md#UploadLicense) | **Post** /internal/api/v1/tenants/{tenantId}/license | Upload license



## CreateTenantInternal

> TenantCreationResponse CreateTenantInternal(ctx).TenantCreationRequest(tenantCreationRequest).Execute()

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
	resp, r, err := apiClient.TenantsAPI.CreateTenantInternal(context.Background()).TenantCreationRequest(tenantCreationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.CreateTenantInternal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTenantInternal`: TenantCreationResponse
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.CreateTenantInternal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantInternalRequest struct via the builder pattern


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


## DeleteTenantByIdInternal

> DeleteTenantById200Response DeleteTenantByIdInternal(ctx, tenantId).IsHardDelete(isHardDelete).Execute()

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
	resp, r, err := apiClient.TenantsAPI.DeleteTenantByIdInternal(context.Background(), tenantId).IsHardDelete(isHardDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.DeleteTenantByIdInternal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTenantByIdInternal`: DeleteTenantById200Response
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.DeleteTenantByIdInternal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **int32** | The id of the tenant to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantByIdInternalRequest struct via the builder pattern


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


## GetLicenseQuery

> []GetLicenseQuery200ResponseInner GetLicenseQuery(ctx, tenantId).Feature(feature).ExpiredBefore(expiredBefore).Execute()

get license query api

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/tenants-manager-client"
)

func main() {
	tenantId := int32(56) // int32 | The id of the tenant to retrieve
	feature := "feature_example" // string |  (optional)
	expiredBefore := time.Now() // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantsAPI.GetLicenseQuery(context.Background(), tenantId).Feature(feature).ExpiredBefore(expiredBefore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.GetLicenseQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLicenseQuery`: []GetLicenseQuery200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.GetLicenseQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **int32** | The id of the tenant to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **feature** | **string** |  | 
 **expiredBefore** | **string** |  | 

### Return type

[**[]GetLicenseQuery200ResponseInner**](GetLicenseQuery200ResponseInner.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenant

> Tenant GetTenant(ctx).TenantName(tenantName).Execute()

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
	tenantName := "tenantName_example" // string | The name of the tenant to retrieve

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantsAPI.GetTenant(context.Background()).TenantName(tenantName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.GetTenant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenant`: Tenant
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.GetTenant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantName** | **string** | The name of the tenant to retrieve | 

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


## GetTenantByIdInternal

> Tenant GetTenantByIdInternal(ctx, tenantId).Execute()

Get tenant by id

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
	resp, r, err := apiClient.TenantsAPI.GetTenantByIdInternal(context.Background(), tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.GetTenantByIdInternal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantByIdInternal`: Tenant
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.GetTenantByIdInternal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **int32** | The id of the tenant to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantByIdInternalRequest struct via the builder pattern


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


## GetTenantLicense

> string GetTenantLicense(ctx, tenantId).Execute()

Get tenant license

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
	resp, r, err := apiClient.TenantsAPI.GetTenantLicense(context.Background(), tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.GetTenantLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantLicense`: string
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.GetTenantLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **int32** | The id of the tenant to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadLicense

> UploadLicense200Response UploadLicense(ctx, tenantId).File(file).Execute()

Upload license

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
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TenantsAPI.UploadLicense(context.Background(), tenantId).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.UploadLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadLicense`: UploadLicense200Response
	fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.UploadLicense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **int32** | The id of the tenant to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadLicenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** |  | 

### Return type

[**UploadLicense200Response**](UploadLicense200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

