# \DepartmentsAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountDepartments**](DepartmentsAPI.md#CountDepartments) | **Get** /api/v1/org-unit/departments/count | Count departments
[**CreateDepartment**](DepartmentsAPI.md#CreateDepartment) | **Post** /api/v1/org-unit/departments | Create department
[**DeleteDepartment**](DepartmentsAPI.md#DeleteDepartment) | **Delete** /api/v1/org-unit/departments/{departmentId} | Delete department
[**GetDepartment**](DepartmentsAPI.md#GetDepartment) | **Get** /api/v1/org-unit/departments/{departmentId} | Get department
[**GetDepartmentAncestors**](DepartmentsAPI.md#GetDepartmentAncestors) | **Get** /internal/org-unit/departments/{departmentId}/ancestors | Get department ancestors
[**GetDepartmentMetrics**](DepartmentsAPI.md#GetDepartmentMetrics) | **Get** /api/v1/org-unit/departments/{departmentId}/metrics | Get department metrics data.
[**GetDepartments**](DepartmentsAPI.md#GetDepartments) | **Get** /api/v1/org-unit/departments | Get departments
[**GetDepartmentsRanking**](DepartmentsAPI.md#GetDepartmentsRanking) | **Get** /api/v1/org-unit/departments/ranking | Get top 5 departments by highest or lowest average resource utilization or allocation
[**GetDepartmentsTelemetry**](DepartmentsAPI.md#GetDepartmentsTelemetry) | **Get** /api/v1/org-unit/departments/telemetry | Get departments telemetry
[**GetSyncDepartments**](DepartmentsAPI.md#GetSyncDepartments) | **Get** /api/v1/org-unit/departments/sync | get departments for sync
[**PatchDepartmentResources**](DepartmentsAPI.md#PatchDepartmentResources) | **Patch** /api/v1/org-unit/departments/{departmentId}/resources | Patch department resources
[**UpdateDepartment**](DepartmentsAPI.md#UpdateDepartment) | **Put** /api/v1/org-unit/departments/{departmentId} | Update department
[**UpdateDepartmentResources**](DepartmentsAPI.md#UpdateDepartmentResources) | **Put** /api/v1/org-unit/departments/{departmentId}/resources | Update department resources
[**UpdateDepartmentStatus**](DepartmentsAPI.md#UpdateDepartmentStatus) | **Put** /api/v1/org-unit/departments/{departmentId}/status | Update Department status from cluster



## CountDepartments

> CountResponse CountDepartments(ctx).FilterBy(filterBy).Execute()

Count departments



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
	filterBy := []string{"Inner_example"} // []string | Filter results by a parameter. Use the format field-name operator value. Operators are == Equals, != Not equals, <= Less than or equal, >= Greater than or equal, =@ contains, !@ Does not contains, =^ Starts with and =$ Ends with. Dates are in ISO 8601 timestamp format and available for operators ==, !=, <= and >=. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DepartmentsAPI.CountDepartments(context.Background()).FilterBy(filterBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.CountDepartments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountDepartments`: CountResponse
	fmt.Fprintf(os.Stdout, "Response from `DepartmentsAPI.CountDepartments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountDepartmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterBy** | **[]string** | Filter results by a parameter. Use the format field-name operator value. Operators are &#x3D;&#x3D; Equals, !&#x3D; Not equals, &lt;&#x3D; Less than or equal, &gt;&#x3D; Greater than or equal, &#x3D;@ contains, !@ Does not contains, &#x3D;^ Starts with and &#x3D;$ Ends with. Dates are in ISO 8601 timestamp format and available for operators &#x3D;&#x3D;, !&#x3D;, &lt;&#x3D; and &gt;&#x3D;. | 

### Return type

[**CountResponse**](CountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDepartment

> Department CreateDepartment(ctx).DepartmentCreationRequest(departmentCreationRequest).Execute()

Create department



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
	departmentCreationRequest := *openapiclient.NewDepartmentCreationRequest("organization1", "71f69d83-ba66-4822-adf5-55ce55efd210", []openapiclient.ResourcesNullable{*openapiclient.NewResourcesNullable()}) // DepartmentCreationRequest | Department to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DepartmentsAPI.CreateDepartment(context.Background()).DepartmentCreationRequest(departmentCreationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.CreateDepartment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDepartment`: Department
	fmt.Fprintf(os.Stdout, "Response from `DepartmentsAPI.CreateDepartment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDepartmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **departmentCreationRequest** | [**DepartmentCreationRequest**](DepartmentCreationRequest.md) | Department to create. | 

### Return type

[**Department**](Department.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDepartment

> DeleteDepartment(ctx, departmentId).Execute()

Delete department



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
	departmentId := "1" // string | The id of the department.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DepartmentsAPI.DeleteDepartment(context.Background(), departmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.DeleteDepartment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**departmentId** | **string** | The id of the department. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDepartmentRequest struct via the builder pattern


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


## GetDepartment

> Department GetDepartment(ctx, departmentId).Execute()

Get department



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
	departmentId := "1" // string | The id of the department.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DepartmentsAPI.GetDepartment(context.Background(), departmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.GetDepartment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDepartment`: Department
	fmt.Fprintf(os.Stdout, "Response from `DepartmentsAPI.GetDepartment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**departmentId** | **string** | The id of the department. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDepartmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Department**](Department.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDepartmentAncestors

> Ancestors GetDepartmentAncestors(ctx, departmentId).Execute()

Get department ancestors

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
	departmentId := "1" // string | The id of the department.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DepartmentsAPI.GetDepartmentAncestors(context.Background(), departmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.GetDepartmentAncestors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDepartmentAncestors`: Ancestors
	fmt.Fprintf(os.Stdout, "Response from `DepartmentsAPI.GetDepartmentAncestors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**departmentId** | **string** | The id of the department. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDepartmentAncestorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Ancestors**](Ancestors.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDepartmentMetrics

> MetricsResponse GetDepartmentMetrics(ctx, departmentId).MetricType(metricType).Start(start).End(end).NumberOfSamples(numberOfSamples).NodepoolName(nodepoolName).Execute()

Get department metrics data.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/org-unit-service-client"
)

func main() {
	departmentId := "1" // string | The id of the department.
	metricType := []openapiclient.OrgUnitMetricType{openapiclient.OrgUnitMetricType("GPU_QUOTA")} // []OrgUnitMetricType | Specify which data to request.
	start := time.Now() // time.Time | Start date of time range to fetch data in ISO 8601 timestamp format.
	end := time.Now() // time.Time | End date of time range to fetch data in ISO 8601 timestamp format.
	numberOfSamples := int32(20) // int32 | The number of samples to take in the specified time range. (optional) (default to 20)
	nodepoolName := "default" // string | Filter using the nodepool. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DepartmentsAPI.GetDepartmentMetrics(context.Background(), departmentId).MetricType(metricType).Start(start).End(end).NumberOfSamples(numberOfSamples).NodepoolName(nodepoolName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.GetDepartmentMetrics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDepartmentMetrics`: MetricsResponse
	fmt.Fprintf(os.Stdout, "Response from `DepartmentsAPI.GetDepartmentMetrics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**departmentId** | **string** | The id of the department. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDepartmentMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricType** | [**[]OrgUnitMetricType**](OrgUnitMetricType.md) | Specify which data to request. | 
 **start** | **time.Time** | Start date of time range to fetch data in ISO 8601 timestamp format. | 
 **end** | **time.Time** | End date of time range to fetch data in ISO 8601 timestamp format. | 
 **numberOfSamples** | **int32** | The number of samples to take in the specified time range. | [default to 20]
 **nodepoolName** | **string** | Filter using the nodepool. | 

### Return type

[**MetricsResponse**](MetricsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDepartments

> GetDepartments200Response GetDepartments(ctx).FilterBy(filterBy).SortBy(sortBy).Verbosity(verbosity).SortOrder(sortOrder).Offset(offset).Limit(limit).Execute()

Get departments



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
	filterBy := []string{"Inner_example"} // []string | Filter results by a parameter. Use the format field-name operator value. Operators are == Equals, != Not equals, <= Less than or equal, >= Greater than or equal, =@ contains, !@ Does not contains, =^ Starts with and =$ Ends with. Dates are in ISO 8601 timestamp format and available for operators ==, !=, <= and >=. (optional)
	sortBy := openapiclient.DepartmentFilterSortFields("name") // DepartmentFilterSortFields | Sort results by a parameters. (optional)
	verbosity := openapiclient.Verbosity("meta") // Verbosity | Departments verbosity. If it is set to \"verbose\", status will be returned. If it is not defined or set to \"brief\" only unit specific data will be returned. (optional)
	sortOrder := "sortOrder_example" // string | Sort results in descending or ascending order. (optional) (default to "asc")
	offset := int32(100) // int32 | The offset of the first item returned in the collection. (optional)
	limit := int32(56) // int32 | The maximum number of entries to return. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DepartmentsAPI.GetDepartments(context.Background()).FilterBy(filterBy).SortBy(sortBy).Verbosity(verbosity).SortOrder(sortOrder).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.GetDepartments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDepartments`: GetDepartments200Response
	fmt.Fprintf(os.Stdout, "Response from `DepartmentsAPI.GetDepartments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDepartmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterBy** | **[]string** | Filter results by a parameter. Use the format field-name operator value. Operators are &#x3D;&#x3D; Equals, !&#x3D; Not equals, &lt;&#x3D; Less than or equal, &gt;&#x3D; Greater than or equal, &#x3D;@ contains, !@ Does not contains, &#x3D;^ Starts with and &#x3D;$ Ends with. Dates are in ISO 8601 timestamp format and available for operators &#x3D;&#x3D;, !&#x3D;, &lt;&#x3D; and &gt;&#x3D;. | 
 **sortBy** | [**DepartmentFilterSortFields**](DepartmentFilterSortFields.md) | Sort results by a parameters. | 
 **verbosity** | [**Verbosity**](Verbosity.md) | Departments verbosity. If it is set to \&quot;verbose\&quot;, status will be returned. If it is not defined or set to \&quot;brief\&quot; only unit specific data will be returned. | 
 **sortOrder** | **string** | Sort results in descending or ascending order. | [default to &quot;asc&quot;]
 **offset** | **int32** | The offset of the first item returned in the collection. | 
 **limit** | **int32** | The maximum number of entries to return. | [default to 50]

### Return type

[**GetDepartments200Response**](GetDepartments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDepartmentsRanking

> GetDepartmentsRanking200Response GetDepartmentsRanking(ctx).MetricType(metricType).Order(order).ClusterId(clusterId).NodepoolName(nodepoolName).Execute()

Get top 5 departments by highest or lowest average resource utilization or allocation



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
	metricType := openapiclient.RankingMetricType("avgGpuAllocation24h") // RankingMetricType | The type of metric to rank by (top/bottom gpu allocation/utilization)
	order := openapiclient.RankingOrderType("top") // RankingOrderType | The order of the ranking (top/bottom)
	clusterId := "clusterId_example" // string | 
	nodepoolName := "nodepoolName_example" // string | The name of the node pool to filter by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DepartmentsAPI.GetDepartmentsRanking(context.Background()).MetricType(metricType).Order(order).ClusterId(clusterId).NodepoolName(nodepoolName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.GetDepartmentsRanking``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDepartmentsRanking`: GetDepartmentsRanking200Response
	fmt.Fprintf(os.Stdout, "Response from `DepartmentsAPI.GetDepartmentsRanking`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDepartmentsRankingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricType** | [**RankingMetricType**](RankingMetricType.md) | The type of metric to rank by (top/bottom gpu allocation/utilization) | 
 **order** | [**RankingOrderType**](RankingOrderType.md) | The order of the ranking (top/bottom) | 
 **clusterId** | **string** |  | 
 **nodepoolName** | **string** | The name of the node pool to filter by | 

### Return type

[**GetDepartmentsRanking200Response**](GetDepartmentsRanking200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDepartmentsTelemetry

> TelemetryResponse GetDepartmentsTelemetry(ctx).TelemetryType(telemetryType).ClusterId(clusterId).NodepoolName(nodepoolName).DepartmentId(departmentId).GroupBy(groupBy).Execute()

Get departments telemetry

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
	telemetryType := openapiclient.OrgUnitTelemetryType("GPU_QUOTA") // OrgUnitTelemetryType | specifies what data to request
	clusterId := "d73a738f-fab3-430a-8fa3-5241493d7128" // string | Filter using the Universally Unique Identifier (UUID) of the cluster. (optional)
	nodepoolName := "default" // string | Filter using the nodepool. (optional)
	departmentId := "1" // string | Filter using the department id. (optional)
	groupBy := []string{"GroupBy_example"} // []string | department fields to group the data by (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DepartmentsAPI.GetDepartmentsTelemetry(context.Background()).TelemetryType(telemetryType).ClusterId(clusterId).NodepoolName(nodepoolName).DepartmentId(departmentId).GroupBy(groupBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.GetDepartmentsTelemetry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDepartmentsTelemetry`: TelemetryResponse
	fmt.Fprintf(os.Stdout, "Response from `DepartmentsAPI.GetDepartmentsTelemetry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDepartmentsTelemetryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **telemetryType** | [**OrgUnitTelemetryType**](OrgUnitTelemetryType.md) | specifies what data to request | 
 **clusterId** | **string** | Filter using the Universally Unique Identifier (UUID) of the cluster. | 
 **nodepoolName** | **string** | Filter using the nodepool. | 
 **departmentId** | **string** | Filter using the department id. | 
 **groupBy** | **[]string** | department fields to group the data by | 

### Return type

[**TelemetryResponse**](TelemetryResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSyncDepartments

> GetSyncDepartments200Response GetSyncDepartments(ctx).ClusterId(clusterId).UpdatedAfter(updatedAfter).Offset(offset).Limit(limit).Execute()

get departments for sync

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/org-unit-service-client"
)

func main() {
	clusterId := "d73a738f-fab3-430a-8fa3-5241493d7128" // string | The id of the cluster
	updatedAfter := time.Now() // time.Time |  (optional)
	offset := int32(100) // int32 | The offset of the first item returned in the collection. (optional)
	limit := int32(56) // int32 | The maximum number of entries to return. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DepartmentsAPI.GetSyncDepartments(context.Background()).ClusterId(clusterId).UpdatedAfter(updatedAfter).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.GetSyncDepartments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSyncDepartments`: GetSyncDepartments200Response
	fmt.Fprintf(os.Stdout, "Response from `DepartmentsAPI.GetSyncDepartments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSyncDepartmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | The id of the cluster | 
 **updatedAfter** | **time.Time** |  | 
 **offset** | **int32** | The offset of the first item returned in the collection. | 
 **limit** | **int32** | The maximum number of entries to return. | [default to 50]

### Return type

[**GetSyncDepartments200Response**](GetSyncDepartments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchDepartmentResources

> []Resources PatchDepartmentResources(ctx, departmentId).ResourcesNullable(resourcesNullable).Execute()

Patch department resources



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
	departmentId := "1" // string | The id of the department.
	resourcesNullable := []openapiclient.ResourcesNullable{*openapiclient.NewResourcesNullable()} // []ResourcesNullable | Department resources to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DepartmentsAPI.PatchDepartmentResources(context.Background(), departmentId).ResourcesNullable(resourcesNullable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.PatchDepartmentResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchDepartmentResources`: []Resources
	fmt.Fprintf(os.Stdout, "Response from `DepartmentsAPI.PatchDepartmentResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**departmentId** | **string** | The id of the department. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchDepartmentResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourcesNullable** | [**[]ResourcesNullable**](ResourcesNullable.md) | Department resources to update. | 

### Return type

[**[]Resources**](Resources.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDepartment

> Department UpdateDepartment(ctx, departmentId).DepartmentUpdateRequest(departmentUpdateRequest).Execute()

Update department



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
	departmentId := "1" // string | The id of the department.
	departmentUpdateRequest := *openapiclient.NewDepartmentUpdateRequest([]openapiclient.ResourcesNullable{*openapiclient.NewResourcesNullable()}) // DepartmentUpdateRequest | Department to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DepartmentsAPI.UpdateDepartment(context.Background(), departmentId).DepartmentUpdateRequest(departmentUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.UpdateDepartment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDepartment`: Department
	fmt.Fprintf(os.Stdout, "Response from `DepartmentsAPI.UpdateDepartment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**departmentId** | **string** | The id of the department. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDepartmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **departmentUpdateRequest** | [**DepartmentUpdateRequest**](DepartmentUpdateRequest.md) | Department to update. | 

### Return type

[**Department**](Department.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDepartmentResources

> []Resources UpdateDepartmentResources(ctx, departmentId).ResourcesNullable(resourcesNullable).Execute()

Update department resources



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
	departmentId := "1" // string | The id of the department.
	resourcesNullable := []openapiclient.ResourcesNullable{*openapiclient.NewResourcesNullable()} // []ResourcesNullable | Department resources to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DepartmentsAPI.UpdateDepartmentResources(context.Background(), departmentId).ResourcesNullable(resourcesNullable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.UpdateDepartmentResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDepartmentResources`: []Resources
	fmt.Fprintf(os.Stdout, "Response from `DepartmentsAPI.UpdateDepartmentResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**departmentId** | **string** | The id of the department. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDepartmentResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourcesNullable** | [**[]ResourcesNullable**](ResourcesNullable.md) | Department resources to update. | 

### Return type

[**[]Resources**](Resources.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDepartmentStatus

> UpdateDepartmentStatus(ctx, departmentId).DepartmentStatusUpdateRequest(departmentStatusUpdateRequest).Execute()

Update Department status from cluster

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
	departmentId := "1" // string | The id of the department.
	departmentStatusUpdateRequest := *openapiclient.NewDepartmentStatusUpdateRequest(*openapiclient.NewDepartmentStatus()) // DepartmentStatusUpdateRequest | Department to update by name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DepartmentsAPI.UpdateDepartmentStatus(context.Background(), departmentId).DepartmentStatusUpdateRequest(departmentStatusUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DepartmentsAPI.UpdateDepartmentStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**departmentId** | **string** | The id of the department. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDepartmentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **departmentStatusUpdateRequest** | [**DepartmentStatusUpdateRequest**](DepartmentStatusUpdateRequest.md) | Department to update by name | 

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

