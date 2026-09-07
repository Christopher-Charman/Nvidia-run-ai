# \AppearanceSettingsAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAppearanceSettingByKey**](AppearanceSettingsAPI.md#GetAppearanceSettingByKey) | **Get** /api/v1/ui-appearance/settings/{key} | Get appearance setting by key
[**GetAppearanceSettings**](AppearanceSettingsAPI.md#GetAppearanceSettings) | **Get** /api/v1/ui-appearance/settings | Get appearance settings
[**UpdateAppearanceSettingByKey**](AppearanceSettingsAPI.md#UpdateAppearanceSettingByKey) | **Put** /api/v1/ui-appearance/settings/{key} | Update appearance setting by key



## GetAppearanceSettingByKey

> map[string]interface{} GetAppearanceSettingByKey(ctx, key).TenantId(tenantId).Execute()

Get appearance setting by key

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
	key := openapiclient.appearanceSettingsKeyEnum("bannerConfig") // AppearanceSettingsKeyEnum | The appearance setting key
	tenantId := "1008" // string | If exists will return the relevant resource for the tenant ID (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppearanceSettingsAPI.GetAppearanceSettingByKey(context.Background(), key).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppearanceSettingsAPI.GetAppearanceSettingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppearanceSettingByKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AppearanceSettingsAPI.GetAppearanceSettingByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | [**AppearanceSettingsKeyEnum**](.md) | The appearance setting key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppearanceSettingByKeyRequest struct via the builder pattern


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


## GetAppearanceSettings

> AppearanceSettingsGetAllResponse GetAppearanceSettings(ctx).TenantId(tenantId).Execute()

Get appearance settings

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
	resp, r, err := apiClient.AppearanceSettingsAPI.GetAppearanceSettings(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppearanceSettingsAPI.GetAppearanceSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAppearanceSettings`: AppearanceSettingsGetAllResponse
	fmt.Fprintf(os.Stdout, "Response from `AppearanceSettingsAPI.GetAppearanceSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAppearanceSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | If exists will return the relevant resource for the tenant ID | [default to &quot;&quot;]

### Return type

[**AppearanceSettingsGetAllResponse**](AppearanceSettingsGetAllResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAppearanceSettingByKey

> UpdateAppearanceSettingByKey(ctx, key).UpdateBannerConfig(updateBannerConfig).TenantId(tenantId).Execute()

Update appearance setting by key

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
	key := openapiclient.appearanceSettingsKeyEnum("bannerConfig") // AppearanceSettingsKeyEnum | The appearance setting key
	updateBannerConfig := *openapiclient.NewUpdateBannerConfig() // UpdateBannerConfig | 
	tenantId := "1008" // string | If exists will return the relevant resource for the tenant ID (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppearanceSettingsAPI.UpdateAppearanceSettingByKey(context.Background(), key).UpdateBannerConfig(updateBannerConfig).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppearanceSettingsAPI.UpdateAppearanceSettingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | [**AppearanceSettingsKeyEnum**](.md) | The appearance setting key | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAppearanceSettingByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBannerConfig** | [**UpdateBannerConfig**](UpdateBannerConfig.md) |  | 
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

