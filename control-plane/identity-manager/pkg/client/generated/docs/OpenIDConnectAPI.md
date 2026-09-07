# \OpenIDConnectAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOpenidConfiguration**](OpenIDConnectAPI.md#GetOpenidConfiguration) | **Get** /.well-known/openid-configuration | OpenID Connect Discovery



## GetOpenidConfiguration

> OpenIDConfiguration GetOpenidConfiguration(ctx).Execute()

OpenID Connect Discovery



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenIDConnectAPI.GetOpenidConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenIDConnectAPI.GetOpenidConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOpenidConfiguration`: OpenIDConfiguration
	fmt.Fprintf(os.Stdout, "Response from `OpenIDConnectAPI.GetOpenidConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenidConfigurationRequest struct via the builder pattern


### Return type

[**OpenIDConfiguration**](OpenIDConfiguration.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

