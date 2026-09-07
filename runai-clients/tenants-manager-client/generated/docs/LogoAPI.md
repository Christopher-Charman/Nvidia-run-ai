# \LogoAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTenantLogo**](LogoAPI.md#GetTenantLogo) | **Get** /api/v1/logo | Get tenant logo. (base64)
[**UploadTenantLogo**](LogoAPI.md#UploadTenantLogo) | **Post** /api/v1/logo | Upload logo for tenant. (base64)



## GetTenantLogo

> GetTenantLogo200Response GetTenantLogo(ctx).Execute()

Get tenant logo. (base64)



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
	resp, r, err := apiClient.LogoAPI.GetTenantLogo(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogoAPI.GetTenantLogo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTenantLogo`: GetTenantLogo200Response
	fmt.Fprintf(os.Stdout, "Response from `LogoAPI.GetTenantLogo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantLogoRequest struct via the builder pattern


### Return type

[**GetTenantLogo200Response**](GetTenantLogo200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadTenantLogo

> UploadTenantLogo(ctx).LogoCreationRequest(logoCreationRequest).Execute()

Upload logo for tenant. (base64)



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
	logoCreationRequest := *openapiclient.NewLogoCreationRequest("Logo_example") // LogoCreationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LogoAPI.UploadTenantLogo(context.Background()).LogoCreationRequest(logoCreationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LogoAPI.UploadTenantLogo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadTenantLogoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **logoCreationRequest** | [**LogoCreationRequest**](LogoCreationRequest.md) |  | 

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

