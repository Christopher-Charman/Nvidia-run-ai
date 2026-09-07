# \WorkloadsPrioritiesAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkloadPriorities**](WorkloadsPrioritiesAPI.md#CreateWorkloadPriorities) | **Post** /api/v1/workload-priorities | Create workload priorities.
[**GetSyncWorkloadPriorities**](WorkloadsPrioritiesAPI.md#GetSyncWorkloadPriorities) | **Get** /api/v1/workload-priorities-sync | Get workload priorities for agent.
[**GetWorkloadClusterPriorities**](WorkloadsPrioritiesAPI.md#GetWorkloadClusterPriorities) | **Get** /api/v1/workload-cluster-priorities | Get workload priorities.
[**GetWorkloadPriorities**](WorkloadsPrioritiesAPI.md#GetWorkloadPriorities) | **Get** /api/v1/workload-priorities | Get workload priorities.
[**WorkloadClusterPrioritiesStatus**](WorkloadsPrioritiesAPI.md#WorkloadClusterPrioritiesStatus) | **Post** /api/v1/workload-priorities/cluster-status | Sync workload cluster priorities status.
[**WorkloadPrioritiesStatus**](WorkloadsPrioritiesAPI.md#WorkloadPrioritiesStatus) | **Post** /api/v1/workload-priorities/runai-status | Sync workload priorities status.



## CreateWorkloadPriorities

> Priority CreateWorkloadPriorities(ctx).PriorityForCreate(priorityForCreate).Execute()

Create workload priorities.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	priorityForCreate := *openapiclient.NewPriorityForCreate("train", int32(50)) // PriorityForCreate |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadsPrioritiesAPI.CreateWorkloadPriorities(context.Background()).PriorityForCreate(priorityForCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsPrioritiesAPI.CreateWorkloadPriorities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkloadPriorities`: Priority
	fmt.Fprintf(os.Stdout, "Response from `WorkloadsPrioritiesAPI.CreateWorkloadPriorities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkloadPrioritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **priorityForCreate** | [**PriorityForCreate**](PriorityForCreate.md) |  | 

### Return type

[**Priority**](Priority.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyncWorkloadPriorities

> []SyncPriority GetSyncWorkloadPriorities(ctx).ClusterId(clusterId).Execute()

Get workload priorities for agent.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	clusterId := "d73a738f-fab3-430a-8fa3-5241493d7128" // string | The id of the cluster

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadsPrioritiesAPI.GetSyncWorkloadPriorities(context.Background()).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsPrioritiesAPI.GetSyncWorkloadPriorities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyncWorkloadPriorities`: []SyncPriority
	fmt.Fprintf(os.Stdout, "Response from `WorkloadsPrioritiesAPI.GetSyncWorkloadPriorities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncWorkloadPrioritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | The id of the cluster | 

### Return type

[**[]SyncPriority**](SyncPriority.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkloadClusterPriorities

> GetWorkloadClusterPriorities200Response GetWorkloadClusterPriorities(ctx).ClusterId(clusterId).SortOrder(sortOrder).SortBy(sortBy).FilterBy(filterBy).Execute()

Get workload priorities.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	clusterId := "d73a738f-fab3-430a-8fa3-5241493d7128" // string | The id of the cluster
	sortOrder := "sortOrder_example" // string | Sort results in descending or ascending order. (optional) (default to "asc")
	sortBy := "sortBy_example" // string |  (optional)
	filterBy := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadsPrioritiesAPI.GetWorkloadClusterPriorities(context.Background()).ClusterId(clusterId).SortOrder(sortOrder).SortBy(sortBy).FilterBy(filterBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsPrioritiesAPI.GetWorkloadClusterPriorities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkloadClusterPriorities`: GetWorkloadClusterPriorities200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkloadsPrioritiesAPI.GetWorkloadClusterPriorities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkloadClusterPrioritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | The id of the cluster | 
 **sortOrder** | **string** | Sort results in descending or ascending order. | [default to &quot;asc&quot;]
 **sortBy** | **string** |  | 
 **filterBy** | **[]string** |  | 

### Return type

[**GetWorkloadClusterPriorities200Response**](GetWorkloadClusterPriorities200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkloadPriorities

> GetWorkloadPriorities200Response GetWorkloadPriorities(ctx).SortOrder(sortOrder).SortBy(sortBy).FilterBy(filterBy).Execute()

Get workload priorities.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	sortOrder := "sortOrder_example" // string | Sort results in descending or ascending order. (optional) (default to "asc")
	sortBy := "sortBy_example" // string | Defines the field to sort the results by (optional)
	filterBy := []string{"Inner_example"} // []string | Filter results by a parameter. Use the format field-name operator value. Operators are == Equals, != Not equals, <= Less than or equal, >= Greater than or equal, =@ contains, !@ Does not contains, =^ Starts with and =$ Ends with. Dates are in ISO 8601 timestamp format and available for operators ==, !=, <= and >= (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkloadsPrioritiesAPI.GetWorkloadPriorities(context.Background()).SortOrder(sortOrder).SortBy(sortBy).FilterBy(filterBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsPrioritiesAPI.GetWorkloadPriorities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkloadPriorities`: GetWorkloadPriorities200Response
	fmt.Fprintf(os.Stdout, "Response from `WorkloadsPrioritiesAPI.GetWorkloadPriorities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkloadPrioritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sortOrder** | **string** | Sort results in descending or ascending order. | [default to &quot;asc&quot;]
 **sortBy** | **string** | Defines the field to sort the results by | 
 **filterBy** | **[]string** | Filter results by a parameter. Use the format field-name operator value. Operators are &#x3D;&#x3D; Equals, !&#x3D; Not equals, &lt;&#x3D; Less than or equal, &gt;&#x3D; Greater than or equal, &#x3D;@ contains, !@ Does not contains, &#x3D;^ Starts with and &#x3D;$ Ends with. Dates are in ISO 8601 timestamp format and available for operators &#x3D;&#x3D;, !&#x3D;, &lt;&#x3D; and &gt;&#x3D; | 

### Return type

[**GetWorkloadPriorities200Response**](GetWorkloadPriorities200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkloadClusterPrioritiesStatus

> WorkloadClusterPrioritiesStatus(ctx).ClusterId(clusterId).PrioritiesForStatus(prioritiesForStatus).Execute()

Sync workload cluster priorities status.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	clusterId := "d73a738f-fab3-430a-8fa3-5241493d7128" // string | The id of the cluster
	prioritiesForStatus := *openapiclient.NewPrioritiesForStatus() // PrioritiesForStatus |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkloadsPrioritiesAPI.WorkloadClusterPrioritiesStatus(context.Background()).ClusterId(clusterId).PrioritiesForStatus(prioritiesForStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsPrioritiesAPI.WorkloadClusterPrioritiesStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkloadClusterPrioritiesStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | The id of the cluster | 
 **prioritiesForStatus** | [**PrioritiesForStatus**](PrioritiesForStatus.md) |  | 

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


## WorkloadPrioritiesStatus

> WorkloadPrioritiesStatus(ctx).ClusterId(clusterId).PrioritiesForStatus(prioritiesForStatus).Execute()

Sync workload priorities status.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/workloads-client"
)

func main() {
	clusterId := "d73a738f-fab3-430a-8fa3-5241493d7128" // string | The id of the cluster
	prioritiesForStatus := *openapiclient.NewPrioritiesForStatus() // PrioritiesForStatus |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkloadsPrioritiesAPI.WorkloadPrioritiesStatus(context.Background()).ClusterId(clusterId).PrioritiesForStatus(prioritiesForStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkloadsPrioritiesAPI.WorkloadPrioritiesStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkloadPrioritiesStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | The id of the cluster | 
 **prioritiesForStatus** | [**PrioritiesForStatus**](PrioritiesForStatus.md) |  | 

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

