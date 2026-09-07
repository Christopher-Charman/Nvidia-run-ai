# \OrgUnitAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgTree**](OrgUnitAPI.md#GetOrgTree) | **Get** /api/v1/org-unit | Get org tree
[**GetOrgUnitsNames**](OrgUnitAPI.md#GetOrgUnitsNames) | **Get** /internal/org-unit/names | Get the names of requested org units.
[**GetPriorities**](OrgUnitAPI.md#GetPriorities) | **Get** /api/v1/org-unit/priorities | Get priorities
[**UpdateOrgUnitSettings**](OrgUnitAPI.md#UpdateOrgUnitSettings) | **Post** /internal/org-unit/settings | 



## GetOrgTree

> GetOrgTree200Response GetOrgTree(ctx).ShowAll(showAll).Execute()

Get org tree

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/org-unit-service-client"
)

func main() {
	showAll := true // bool | flag to show all org unit without filtering (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgUnitAPI.GetOrgTree(context.Background()).ShowAll(showAll).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgUnitAPI.GetOrgTree``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgTree`: GetOrgTree200Response
	fmt.Fprintf(os.Stdout, "Response from `OrgUnitAPI.GetOrgTree`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **showAll** | **bool** | flag to show all org unit without filtering | 

### Return type

[**GetOrgTree200Response**](GetOrgTree200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgUnitsNames

> []OrgUnit GetOrgUnitsNames(ctx).OrgUnitType(orgUnitType).OrgUnitIds(orgUnitIds).Execute()

Get the names of requested org units.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/org-unit-service-client"
)

func main() {
	orgUnitType := "orgUnitType_example" // string | The org unit type.
	orgUnitIds := "7" // string | The org unit unique identifiers, seperated by ','.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgUnitAPI.GetOrgUnitsNames(context.Background()).OrgUnitType(orgUnitType).OrgUnitIds(orgUnitIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgUnitAPI.GetOrgUnitsNames``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgUnitsNames`: []OrgUnit
	fmt.Fprintf(os.Stdout, "Response from `OrgUnitAPI.GetOrgUnitsNames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgUnitsNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgUnitType** | **string** | The org unit type. | 
 **orgUnitIds** | **string** | The org unit unique identifiers, seperated by &#39;,&#39;. | 

### Return type

[**[]OrgUnit**](OrgUnit.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriorities

> GetPriorities200Response GetPriorities(ctx).Execute()

Get priorities



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/org-unit-service-client"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgUnitAPI.GetPriorities(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgUnitAPI.GetPriorities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPriorities`: GetPriorities200Response
	fmt.Fprintf(os.Stdout, "Response from `OrgUnitAPI.GetPriorities`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrioritiesRequest struct via the builder pattern


### Return type

[**GetPriorities200Response**](GetPriorities200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgUnitSettings

> UpdateOrgUnitSettings(ctx).OrgUnitSingleSetting(orgUnitSingleSetting).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/org-unit-service-client"
)

func main() {
	orgUnitSingleSetting := *openapiclient.NewOrgUnitSingleSetting() // OrgUnitSingleSetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgUnitAPI.UpdateOrgUnitSettings(context.Background()).OrgUnitSingleSetting(orgUnitSingleSetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgUnitAPI.UpdateOrgUnitSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgUnitSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgUnitSingleSetting** | [**OrgUnitSingleSetting**](OrgUnitSingleSetting.md) |  | 

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

