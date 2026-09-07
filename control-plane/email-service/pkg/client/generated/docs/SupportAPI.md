# \SupportAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContactSupport**](SupportAPI.md#ContactSupport) | **Post** /api/v1/support/contact | contact Run:ai support



## ContactSupport

> ContactSupport(ctx).ContactSupportRequest(contactSupportRequest).Execute()

contact Run:ai support



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/control-plane/email-service/pkg/client"
)

func main() {
	contactSupportRequest := *openapiclient.NewContactSupportRequest("Message_example") // ContactSupportRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SupportAPI.ContactSupport(context.Background()).ContactSupportRequest(contactSupportRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SupportAPI.ContactSupport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiContactSupportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contactSupportRequest** | [**ContactSupportRequest**](ContactSupportRequest.md) |  | 

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

