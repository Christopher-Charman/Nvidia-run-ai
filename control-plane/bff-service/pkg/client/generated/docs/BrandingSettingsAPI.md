# \BrandingSettingsAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBrandingSettingByKey**](BrandingSettingsAPI.md#GetBrandingSettingByKey) | **Get** /api/v1/branding/settings/{key} | Get branding setting by key
[**GetBrandingSettings**](BrandingSettingsAPI.md#GetBrandingSettings) | **Get** /api/v1/branding/settings | Get branding settings
[**UpdateBrandingSettingByKey**](BrandingSettingsAPI.md#UpdateBrandingSettingByKey) | **Put** /api/v1/branding/settings/{key} | Put branding setting by key



## GetBrandingSettingByKey

> map[string]interface{} GetBrandingSettingByKey(ctx, key).TenantId(tenantId).Execute()

Get branding setting by key

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/control-plane/bff-service/pkg/client"
)

func main() {
	key := openapiclient.brandingSettingsKeyEnum("logo") // BrandingSettingsKeyEnum | The branding setting key
	tenantId := "1008" // string | If exists will return the relevant resource for the tenant ID (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BrandingSettingsAPI.GetBrandingSettingByKey(context.Background(), key).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandingSettingsAPI.GetBrandingSettingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBrandingSettingByKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BrandingSettingsAPI.GetBrandingSettingByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | [**BrandingSettingsKeyEnum**](.md) | The branding setting key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandingSettingByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tenantId** | **string** | If exists will return the relevant resource for the tenant ID | [default to &quot;&quot;]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrandingSettings

> BrandingSettingsGetAllResponse GetBrandingSettings(ctx).TenantId(tenantId).Execute()

Get branding settings

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/control-plane/bff-service/pkg/client"
)

func main() {
	tenantId := "1008" // string | If exists will return the relevant resource for the tenant ID (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BrandingSettingsAPI.GetBrandingSettings(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandingSettingsAPI.GetBrandingSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBrandingSettings`: BrandingSettingsGetAllResponse
	fmt.Fprintf(os.Stdout, "Response from `BrandingSettingsAPI.GetBrandingSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandingSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | If exists will return the relevant resource for the tenant ID | [default to &quot;&quot;]

### Return type

[**BrandingSettingsGetAllResponse**](BrandingSettingsGetAllResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBrandingSettingByKey

> UpdateBrandingSettingByKey(ctx, key).UpdateLogo(updateLogo).TenantId(tenantId).Execute()

Put branding setting by key



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/control-plane/bff-service/pkg/client"
)

func main() {
	key := openapiclient.brandingSettingsKeyEnum("logo") // BrandingSettingsKeyEnum | The branding setting key
	updateLogo := *openapiclient.NewUpdateLogo("Logo_example") // UpdateLogo | 
	tenantId := "1008" // string | If exists will return the relevant resource for the tenant ID (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BrandingSettingsAPI.UpdateBrandingSettingByKey(context.Background(), key).UpdateLogo(updateLogo).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandingSettingsAPI.UpdateBrandingSettingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | [**BrandingSettingsKeyEnum**](.md) | The branding setting key | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBrandingSettingByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateLogo** | [**UpdateLogo**](UpdateLogo.md) |  | 
 **tenantId** | **string** | If exists will return the relevant resource for the tenant ID | [default to &quot;&quot;]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

