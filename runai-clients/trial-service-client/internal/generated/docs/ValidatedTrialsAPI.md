# \ValidatedTrialsAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTrial**](ValidatedTrialsAPI.md#GetTrial) | **Get** /api/v1/trials/{trialId} | Get trial by id
[**VerifyTrial**](ValidatedTrialsAPI.md#VerifyTrial) | **Post** /api/v1/trials/{trialId}/verify | Verify trial



## GetTrial

> Trial GetTrial(ctx, trialId).Execute()

Get trial by id

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/trial-service-client"
)

func main() {
	trialId := "123e4567-e89b-12d3-a456-426614174000" // string | The id of the trial

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidatedTrialsAPI.GetTrial(context.Background(), trialId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidatedTrialsAPI.GetTrial``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrial`: Trial
	fmt.Fprintf(os.Stdout, "Response from `ValidatedTrialsAPI.GetTrial`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trialId** | **string** | The id of the trial | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Trial**](Trial.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyTrial

> VerifyTrial200Response VerifyTrial(ctx, trialId).VerifyTrialRequest(verifyTrialRequest).Execute()

Verify trial

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/trial-service-client"
)

func main() {
	trialId := "123e4567-e89b-12d3-a456-426614174000" // string | The id of the trial
	verifyTrialRequest := *openapiclient.NewVerifyTrialRequest("Password_example") // VerifyTrialRequest | Tenant data

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ValidatedTrialsAPI.VerifyTrial(context.Background(), trialId).VerifyTrialRequest(verifyTrialRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ValidatedTrialsAPI.VerifyTrial``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyTrial`: VerifyTrial200Response
	fmt.Fprintf(os.Stdout, "Response from `ValidatedTrialsAPI.VerifyTrial`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trialId** | **string** | The id of the trial | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyTrialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verifyTrialRequest** | [**VerifyTrialRequest**](VerifyTrialRequest.md) | Tenant data | 

### Return type

[**VerifyTrial200Response**](VerifyTrial200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

