# \TokensAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppToken**](TokensAPI.md#AppToken) | **Post** /v1/k8s/auth/oauth/apptoken | get application token
[**GrantToken**](TokensAPI.md#GrantToken) | **Post** /api/v1/token | Create a token



## AppToken

> AppTokenResponse AppToken(ctx).AppTokenRequest(appTokenRequest).Execute()

get application token



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
	appTokenRequest := *openapiclient.NewAppTokenRequest("Id_example", "Name_example", "Secret_example") // AppTokenRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.AppToken(context.Background()).AppTokenRequest(appTokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.AppToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppToken`: AppTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.AppToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appTokenRequest** | [**AppTokenRequest**](AppTokenRequest.md) |  | 

### Return type

[**AppTokenResponse**](AppTokenResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GrantToken

> TokenResponse GrantToken(ctx).TokenRequest(tokenRequest).Execute()

Create a token



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
	tokenRequest := *openapiclient.NewTokenRequest() // TokenRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TokensAPI.GrantToken(context.Background()).TokenRequest(tokenRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TokensAPI.GrantToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GrantToken`: TokenResponse
	fmt.Fprintf(os.Stdout, "Response from `TokensAPI.GrantToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGrantTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenRequest** | [**TokenRequest**](TokenRequest.md) |  | 

### Return type

[**TokenResponse**](TokenResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

