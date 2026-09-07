# \NotificationsSettingsAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetNotificationsSettingByKey**](NotificationsSettingsAPI.md#GetNotificationsSettingByKey) | **Get** /api/v1/notifications/settings/{key} | Get notifications setting by key
[**GetNotificationsSettings**](NotificationsSettingsAPI.md#GetNotificationsSettings) | **Get** /api/v1/notifications/settings | Get notifications settings
[**UpdateNotificationsSettingByKey**](NotificationsSettingsAPI.md#UpdateNotificationsSettingByKey) | **Put** /api/v1/notifications/settings/{key} | Update notifications setting by key



## GetNotificationsSettingByKey

> map[string]interface{} GetNotificationsSettingByKey(ctx, key).TenantId(tenantId).Execute()

Get notifications setting by key

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
	key := openapiclient.notificationsSettingsKeyEnum("adminMessage") // NotificationsSettingsKeyEnum | The notifications setting key
	tenantId := "1008" // string | If exists will return the relevant resource for the tenant ID (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationsSettingsAPI.GetNotificationsSettingByKey(context.Background(), key).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsSettingsAPI.GetNotificationsSettingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationsSettingByKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NotificationsSettingsAPI.GetNotificationsSettingByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | [**NotificationsSettingsKeyEnum**](.md) | The notifications setting key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsSettingByKeyRequest struct via the builder pattern


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


## GetNotificationsSettings

> NotificationsSettingsGetAllResponse GetNotificationsSettings(ctx).TenantId(tenantId).Execute()

Get notifications settings

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
	resp, r, err := apiClient.NotificationsSettingsAPI.GetNotificationsSettings(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsSettingsAPI.GetNotificationsSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationsSettings`: NotificationsSettingsGetAllResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationsSettingsAPI.GetNotificationsSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | If exists will return the relevant resource for the tenant ID | [default to &quot;&quot;]

### Return type

[**NotificationsSettingsGetAllResponse**](NotificationsSettingsGetAllResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNotificationsSettingByKey

> UpdateNotificationsSettingByKey(ctx, key).UpdateAdminMessage(updateAdminMessage).TenantId(tenantId).Execute()

Update notifications setting by key

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
	key := openapiclient.notificationsSettingsKeyEnum("adminMessage") // NotificationsSettingsKeyEnum | The notifications setting key
	updateAdminMessage := *openapiclient.NewUpdateAdminMessage() // UpdateAdminMessage | 
	tenantId := "1008" // string | If exists will return the relevant resource for the tenant ID (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationsSettingsAPI.UpdateNotificationsSettingByKey(context.Background(), key).UpdateAdminMessage(updateAdminMessage).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationsSettingsAPI.UpdateNotificationsSettingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | [**NotificationsSettingsKeyEnum**](.md) | The notifications setting key | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationsSettingByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateAdminMessage** | [**UpdateAdminMessage**](UpdateAdminMessage.md) |  | 
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

