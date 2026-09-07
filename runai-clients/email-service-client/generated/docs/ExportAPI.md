# \ExportAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportEmail**](ExportAPI.md#ExportEmail) | **Post** /api/v1/export/email | export email



## ExportEmail

> ExportEmail(ctx).ExportEmailRequest(exportEmailRequest).Execute()

export email



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/email-service-client"
)

func main() {
	exportEmailRequest := *openapiclient.NewExportEmailRequest([]string{"To_example"}, *openapiclient.NewTemplate("Id_example")) // ExportEmailRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ExportAPI.ExportEmail(context.Background()).ExportEmailRequest(exportEmailRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ExportAPI.ExportEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportEmailRequest** | [**ExportEmailRequest**](ExportEmailRequest.md) |  | 

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

