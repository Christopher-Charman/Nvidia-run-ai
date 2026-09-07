# \UIAnalyticsSettingsAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSettings**](UIAnalyticsSettingsAPI.md#GetSettings) | **Get** /api/v1/ui-analytics/settings | Get UI analytics settings
[**GetUiAnalyticsSettingByKey**](UIAnalyticsSettingsAPI.md#GetUiAnalyticsSettingByKey) | **Get** /api/v1/ui-analytics/settings/{key} | Get UI analytics setting by key
[**UpdateUiAnalyticsSettingByKey**](UIAnalyticsSettingsAPI.md#UpdateUiAnalyticsSettingByKey) | **Put** /api/v1/ui-analytics/settings/{key} | Update UI analytics setting by key



## GetSettings

> AnalyticsSettingsGetAllResponse GetSettings(ctx).TenantId(tenantId).Execute()

Get UI analytics settings

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
	resp, r, err := apiClient.UIAnalyticsSettingsAPI.GetSettings(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UIAnalyticsSettingsAPI.GetSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSettings`: AnalyticsSettingsGetAllResponse
	fmt.Fprintf(os.Stdout, "Response from `UIAnalyticsSettingsAPI.GetSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | If exists will return the relevant resource for the tenant ID | [default to &quot;&quot;]

### Return type

[**AnalyticsSettingsGetAllResponse**](AnalyticsSettingsGetAllResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUiAnalyticsSettingByKey

> map[string]interface{} GetUiAnalyticsSettingByKey(ctx, key).TenantId(tenantId).Execute()

Get UI analytics setting by key

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
	key := openapiclient.uiAnalyticsSettingsKeyEnum("consumptionReport") // UiAnalyticsSettingsKeyEnum | The UI analytics setting key
	tenantId := "1008" // string | If exists will return the relevant resource for the tenant ID (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UIAnalyticsSettingsAPI.GetUiAnalyticsSettingByKey(context.Background(), key).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UIAnalyticsSettingsAPI.GetUiAnalyticsSettingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUiAnalyticsSettingByKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `UIAnalyticsSettingsAPI.GetUiAnalyticsSettingByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | [**UiAnalyticsSettingsKeyEnum**](.md) | The UI analytics setting key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUiAnalyticsSettingByKeyRequest struct via the builder pattern


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


## UpdateUiAnalyticsSettingByKey

> UpdateUiAnalyticsSettingByKey(ctx, key).UpdateConsumptionReport(updateConsumptionReport).TenantId(tenantId).Execute()

Update UI analytics setting by key

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
	key := openapiclient.uiAnalyticsSettingsKeyEnum("consumptionReport") // UiAnalyticsSettingsKeyEnum | The UI analytics setting key
	updateConsumptionReport := *openapiclient.NewUpdateConsumptionReport() // UpdateConsumptionReport | 
	tenantId := "1008" // string | If exists will return the relevant resource for the tenant ID (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UIAnalyticsSettingsAPI.UpdateUiAnalyticsSettingByKey(context.Background(), key).UpdateConsumptionReport(updateConsumptionReport).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UIAnalyticsSettingsAPI.UpdateUiAnalyticsSettingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | [**UiAnalyticsSettingsKeyEnum**](.md) | The UI analytics setting key | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUiAnalyticsSettingByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateConsumptionReport** | [**UpdateConsumptionReport**](UpdateConsumptionReport.md) |  | 
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

