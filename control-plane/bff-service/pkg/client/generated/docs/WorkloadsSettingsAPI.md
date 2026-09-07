# \WorkloadsSettingsAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWorkloadsSettingByKey**](WorkloadsSettingsAPI.md#GetWorkloadsSettingByKey) | **Get** /api/v1/ui-workloads/settings/{key} | Get workloads setting by key
[**GetWorkloadsSettings**](WorkloadsSettingsAPI.md#GetWorkloadsSettings) | **Get** /api/v1/ui-workloads/settings | Get workloads settings
[**UpdateWorkloadsSettingByKey**](WorkloadsSettingsAPI.md#UpdateWorkloadsSettingByKey) | **Put** /api/v1/ui-workloads/settings/{key} | Update workloads setting by key



## GetWorkloadsSettingByKey

> map[string]interface{} GetWorkloadsSettingByKey(ctx, key).TenantId(tenantId).Execute()

Get workloads setting by key

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
	key := openapiclient.workloadsSettingsKeyEnum("enableNIM") // WorkloadsSettingsKeyEnum | The workloads setting key
	tenantId := "1008" // string | If exists will return the relevant resource for the tenant ID (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadsSettingsAPI.GetWorkloadsSettingByKey(context.Background(), key).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsSettingsAPI.GetWorkloadsSettingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkloadsSettingByKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkloadsSettingsAPI.GetWorkloadsSettingByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | [**WorkloadsSettingsKeyEnum**](.md) | The workloads setting key | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkloadsSettingByKeyRequest struct via the builder pattern


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


## GetWorkloadsSettings

> WorkloadsSettingsGetAllResponse GetWorkloadsSettings(ctx).TenantId(tenantId).Execute()

Get workloads settings

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
	resp, r, err := apiClient.WorkloadsSettingsAPI.GetWorkloadsSettings(context.Background()).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsSettingsAPI.GetWorkloadsSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkloadsSettings`: WorkloadsSettingsGetAllResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkloadsSettingsAPI.GetWorkloadsSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkloadsSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantId** | **string** | If exists will return the relevant resource for the tenant ID | [default to &quot;&quot;]

### Return type

[**WorkloadsSettingsGetAllResponse**](WorkloadsSettingsGetAllResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkloadsSettingByKey

> UpdateWorkloadsSettingByKey(ctx, key).UpdateWorkloadsSettingByKeyRequest(updateWorkloadsSettingByKeyRequest).TenantId(tenantId).Execute()

Update workloads setting by key

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
	key := openapiclient.workloadsSettingsKeyEnum("enableNIM") // WorkloadsSettingsKeyEnum | The workloads setting key
	updateWorkloadsSettingByKeyRequest := *openapiclient.NewUpdateWorkloadsSettingByKeyRequest() // UpdateWorkloadsSettingByKeyRequest | 
	tenantId := "1008" // string | If exists will return the relevant resource for the tenant ID (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkloadsSettingsAPI.UpdateWorkloadsSettingByKey(context.Background(), key).UpdateWorkloadsSettingByKeyRequest(updateWorkloadsSettingByKeyRequest).TenantId(tenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsSettingsAPI.UpdateWorkloadsSettingByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | [**WorkloadsSettingsKeyEnum**](.md) | The workloads setting key | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkloadsSettingByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWorkloadsSettingByKeyRequest** | [**UpdateWorkloadsSettingByKeyRequest**](UpdateWorkloadsSettingByKeyRequest.md) |  | 
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

