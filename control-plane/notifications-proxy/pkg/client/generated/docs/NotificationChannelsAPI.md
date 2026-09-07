# \NotificationChannelsAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationChannelConfig**](NotificationChannelsAPI.md#CreateNotificationChannelConfig) | **Post** /api/v1/notification-channels | Create configuration of notification channel
[**CreateSlackApp**](NotificationChannelsAPI.md#CreateSlackApp) | **Post** /api/v1/notification-channels/slack/create-app | Create Slack app
[**DeleteNotificationChannelConfig**](NotificationChannelsAPI.md#DeleteNotificationChannelConfig) | **Delete** /api/v1/notification-channels/{name} | Delete configuration of Notification Channel
[**GetNotificationChannelConfig**](NotificationChannelsAPI.md#GetNotificationChannelConfig) | **Get** /api/v1/notification-channels/{name} | Get configuration of notification channel
[**GetNotificationChannels**](NotificationChannelsAPI.md#GetNotificationChannels) | **Get** /api/v1/notification-channels | Get supported Notification Channels
[**PatchNotificationChannelConfig**](NotificationChannelsAPI.md#PatchNotificationChannelConfig) | **Patch** /api/v1/notification-channels/{name} | Patch configuration of Notification Channel
[**UpdateNotificationChannelConfig**](NotificationChannelsAPI.md#UpdateNotificationChannelConfig) | **Put** /api/v1/notification-channels/{name} | Update configuration of Notification Channel
[**ValidateNotificationChannel**](NotificationChannelsAPI.md#ValidateNotificationChannel) | **Post** /api/v1/validate-notification-channel | Validate configuration of Notification Channel



## CreateNotificationChannelConfig

> CreateNotificationChannelConfig(ctx).NotificationChannel(notificationChannel).Execute()

Create configuration of notification channel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/control-plane/notifications-proxy/pkg/client"
)

func main() {
	notificationChannel := *openapiclient.NewNotificationChannel(openapiclient.NotificationChannelType("email"), openapiclient.NotificationChannelConfig{EmailNotificationChannelConfig: openapiclient.NewEmailNotificationChannelConfig()}, "Name_example") // NotificationChannel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationChannelsAPI.CreateNotificationChannelConfig(context.Background()).NotificationChannel(notificationChannel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsAPI.CreateNotificationChannelConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationChannelConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationChannel** | [**NotificationChannel**](NotificationChannel.md) |  | 

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


## CreateSlackApp

> CreateSlackAppResponse CreateSlackApp(ctx).CreateSlackAppRequestBody(createSlackAppRequestBody).Execute()

Create Slack app



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/control-plane/notifications-proxy/pkg/client"
)

func main() {
	createSlackAppRequestBody := *openapiclient.NewCreateSlackAppRequestBody("AccessToken_example") // CreateSlackAppRequestBody |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationChannelsAPI.CreateSlackApp(context.Background()).CreateSlackAppRequestBody(createSlackAppRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsAPI.CreateSlackApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSlackApp`: CreateSlackAppResponse
	fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsAPI.CreateSlackApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSlackAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSlackAppRequestBody** | [**CreateSlackAppRequestBody**](CreateSlackAppRequestBody.md) |  | 

### Return type

[**CreateSlackAppResponse**](CreateSlackAppResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationChannelConfig

> DeleteNotificationChannelConfig(ctx, name).Execute()

Delete configuration of Notification Channel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/control-plane/notifications-proxy/pkg/client"
)

func main() {
	name := "name_example" // string | Notification Channel name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationChannelsAPI.DeleteNotificationChannelConfig(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsAPI.DeleteNotificationChannelConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Notification Channel name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationChannelConfigRequest struct via the builder pattern


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


## GetNotificationChannelConfig

> NotificationChannel GetNotificationChannelConfig(ctx, name).Execute()

Get configuration of notification channel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/control-plane/notifications-proxy/pkg/client"
)

func main() {
	name := "name_example" // string | Notification Channel name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationChannelsAPI.GetNotificationChannelConfig(context.Background(), name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsAPI.GetNotificationChannelConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationChannelConfig`: NotificationChannel
	fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsAPI.GetNotificationChannelConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Notification Channel name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationChannelConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NotificationChannel**](NotificationChannel.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationChannels

> []NotificationChannelMeta GetNotificationChannels(ctx).Execute()

Get supported Notification Channels



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/control-plane/notifications-proxy/pkg/client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationChannelsAPI.GetNotificationChannels(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsAPI.GetNotificationChannels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationChannels`: []NotificationChannelMeta
	fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsAPI.GetNotificationChannels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationChannelsRequest struct via the builder pattern


### Return type

[**[]NotificationChannelMeta**](NotificationChannelMeta.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchNotificationChannelConfig

> PatchNotificationChannelConfig(ctx, name).NotificationChannelForPatch(notificationChannelForPatch).Execute()

Patch configuration of Notification Channel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/control-plane/notifications-proxy/pkg/client"
)

func main() {
	name := "name_example" // string | Notification Channel name
	notificationChannelForPatch := *openapiclient.NewNotificationChannelForPatch() // NotificationChannelForPatch |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationChannelsAPI.PatchNotificationChannelConfig(context.Background(), name).NotificationChannelForPatch(notificationChannelForPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsAPI.PatchNotificationChannelConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Notification Channel name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNotificationChannelConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationChannelForPatch** | [**NotificationChannelForPatch**](NotificationChannelForPatch.md) |  | 

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


## UpdateNotificationChannelConfig

> UpdateNotificationChannelConfig(ctx, name).NotificationChannel(notificationChannel).Execute()

Update configuration of Notification Channel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/control-plane/notifications-proxy/pkg/client"
)

func main() {
	name := "name_example" // string | Notification Channel name
	notificationChannel := *openapiclient.NewNotificationChannel(openapiclient.NotificationChannelType("email"), openapiclient.NotificationChannelConfig{EmailNotificationChannelConfig: openapiclient.NewEmailNotificationChannelConfig()}, "Name_example") // NotificationChannel |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotificationChannelsAPI.UpdateNotificationChannelConfig(context.Background(), name).NotificationChannel(notificationChannel).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsAPI.UpdateNotificationChannelConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** | Notification Channel name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationChannelConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notificationChannel** | [**NotificationChannel**](NotificationChannel.md) |  | 

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


## ValidateNotificationChannel

> bool ValidateNotificationChannel(ctx).NotificationChannelForValidate(notificationChannelForValidate).Execute()

Validate configuration of Notification Channel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/control-plane/notifications-proxy/pkg/client"
)

func main() {
	notificationChannelForValidate := *openapiclient.NewNotificationChannelForValidate(openapiclient.NotificationChannelType("email"), openapiclient.NotificationChannelConfig{EmailNotificationChannelConfig: openapiclient.NewEmailNotificationChannelConfig()}) // NotificationChannelForValidate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotificationChannelsAPI.ValidateNotificationChannel(context.Background()).NotificationChannelForValidate(notificationChannelForValidate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotificationChannelsAPI.ValidateNotificationChannel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateNotificationChannel`: bool
	fmt.Fprintf(os.Stdout, "Response from `NotificationChannelsAPI.ValidateNotificationChannel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateNotificationChannelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notificationChannelForValidate** | [**NotificationChannelForValidate**](NotificationChannelForValidate.md) |  | 

### Return type

**bool**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

