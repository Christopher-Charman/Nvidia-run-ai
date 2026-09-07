# \WorkspacesAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountAssets**](WorkspacesAPI.md#CountAssets) | **Get** /api/v1/asset/count | Count assets.
[**CreateWorkspace**](WorkspacesAPI.md#CreateWorkspace) | **Post** /api/v1/workspace | Create workspace.
[**CreateWorkspaceV2**](WorkspacesAPI.md#CreateWorkspaceV2) | **Post** /api/v2/workspace | Create a workspace - assets based [Experimental]
[**GetWorkspaceByIdV2**](WorkspacesAPI.md#GetWorkspaceByIdV2) | **Get** /api/v2/workspace/{workloadId} | Get workspace data - assets based [Experimental]
[**GetWorkspaceMerged**](WorkspacesAPI.md#GetWorkspaceMerged) | **Get** /api/v2/workspace/merged | Get merged workspace
[**ListWorkspaces**](WorkspacesAPI.md#ListWorkspaces) | **Get** /v1/k8s/clusters/{clusterUuid}/workspaces | List workspaces in the cluster



## CountAssets

> AssetCountResponse CountAssets(ctx).Kind(kind).Name(name).Scope(scope).ProjectId(projectId).DepartmentId(departmentId).ClusterId(clusterId).Execute()

Count assets.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/assets-service-client"
)

func main() {
	kind := "kind_example" // string | Request results for specific asset kind. (optional)
	name := "name_example" // string | Filter results by name. (optional)
	scope := "scope_example" // string | Filter results by scope. (optional)
	projectId := int32(56) // int32 | Filter results by project id. if scope filter is project, only assets from the specific project will be included in the response. otherwise, the response will include both project, department, cluster and tenant assets. (optional)
	departmentId := "1" // string | Filter using the department id. (optional)
	clusterId := "d73a738f-fab3-430a-8fa3-5241493d7128" // string | Filter using the Universally Unique Identifier (UUID) of the cluster. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.CountAssets(context.Background()).Kind(kind).Name(name).Scope(scope).ProjectId(projectId).DepartmentId(departmentId).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.CountAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountAssets`: AssetCountResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.CountAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** | Request results for specific asset kind. | 
 **name** | **string** | Filter results by name. | 
 **scope** | **string** | Filter results by scope. | 
 **projectId** | **int32** | Filter results by project id. if scope filter is project, only assets from the specific project will be included in the response. otherwise, the response will include both project, department, cluster and tenant assets. | 
 **departmentId** | **string** | Filter using the department id. | 
 **clusterId** | **string** | Filter using the Universally Unique Identifier (UUID) of the cluster. | 

### Return type

[**AssetCountResponse**](AssetCountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkspace

> Workspace CreateWorkspace(ctx).WorkloadCreationRequest(workloadCreationRequest).Execute()

Create workspace.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/assets-service-client"
)

func main() {
	workloadCreationRequest := *openapiclient.NewWorkloadCreationRequest("Name_example", "71f69d83-ba66-4822-adf5-55ce55efd210", int32(123), "Namespace_example", *openapiclient.NewAssetsIds("Environment_example")) // WorkloadCreationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.CreateWorkspace(context.Background()).WorkloadCreationRequest(workloadCreationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.CreateWorkspace``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkspace`: Workspace
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.CreateWorkspace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workloadCreationRequest** | [**WorkloadCreationRequest**](WorkloadCreationRequest.md) |  | 

### Return type

[**Workspace**](Workspace.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkspaceV2

> WorkspaceV2 CreateWorkspaceV2(ctx).WorkspaceCreationRequestV2(workspaceCreationRequestV2).Execute()

Create a workspace - assets based [Experimental]



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/assets-service-client"
)

func main() {
	workspaceCreationRequestV2 := *openapiclient.NewWorkspaceCreationRequestV2(*openapiclient.NewWorkloadCreationMeta("my-workload-name", "1", "71f69d83-ba66-4822-adf5-55ce55efd210"), *openapiclient.NewAssetsIds("Environment_example")) // WorkspaceCreationRequestV2 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.CreateWorkspaceV2(context.Background()).WorkspaceCreationRequestV2(workspaceCreationRequestV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.CreateWorkspaceV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkspaceV2`: WorkspaceV2
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.CreateWorkspaceV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceCreationRequestV2** | [**WorkspaceCreationRequestV2**](WorkspaceCreationRequestV2.md) |  | 

### Return type

[**WorkspaceV2**](WorkspaceV2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkspaceByIdV2

> WorkspaceV2 GetWorkspaceByIdV2(ctx, workloadId).Execute()

Get workspace data - assets based [Experimental]



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/assets-service-client"
)

func main() {
	workloadId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The  Universally Unique Identifier (UUID) of the workload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.GetWorkspaceByIdV2(context.Background(), workloadId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.GetWorkspaceByIdV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkspaceByIdV2`: WorkspaceV2
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.GetWorkspaceByIdV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workloadId** | **string** | The  Universally Unique Identifier (UUID) of the workload. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceByIdV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WorkspaceV2**](WorkspaceV2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkspaceMerged

> WorkspaceV2 GetWorkspaceMerged(ctx).ProjectId(projectId).WorkloadId(workloadId).TemplateId(templateId).Execute()

Get merged workspace



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/assets-service-client"
)

func main() {
	projectId := "4500001" // string | The id of the project (optional)
	workloadId := "d73a738f-fab3-430a-8fa3-5241493d7128" // string | The id of the workload (optional)
	templateId := "d73a738f-fab3-430a-8fa3-5241493d7128" // string | The id of the template (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.GetWorkspaceMerged(context.Background()).ProjectId(projectId).WorkloadId(workloadId).TemplateId(templateId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.GetWorkspaceMerged``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkspaceMerged`: WorkspaceV2
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.GetWorkspaceMerged`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceMergedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** | The id of the project | 
 **workloadId** | **string** | The id of the workload | 
 **templateId** | **string** | The id of the template | 

### Return type

[**WorkspaceV2**](WorkspaceV2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspaces

> WorkspaceListResponse ListWorkspaces(ctx, clusterUuid).NodeId(nodeId).Filter(filter).SortBy(sortBy).SortDirection(sortDirection).From(from).Limit(limit).Execute()

List workspaces in the cluster

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/assets-service-client"
)

func main() {
	clusterUuid := "9f55255e-11ed-47c7-acef-fc4054768dbc" // string | The Universally Unique Identifier (UUID) of the cluster.
	nodeId := "nodeId_example" // string | Unique identifier of the node. (optional)
	filter := "filter_example" // string |  (optional)
	sortBy := "sortBy_example" // string | Order of the results. (optional)
	sortDirection := "sortDirection_example" // string |  (optional)
	from := int32(56) // int32 | Start the response from a given number of result. Used along with 'limit' to retrieve the results paginated. (optional)
	limit := int32(56) // int32 | The maximum number of entries to return. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkspacesAPI.ListWorkspaces(context.Background(), clusterUuid).NodeId(nodeId).Filter(filter).SortBy(sortBy).SortDirection(sortDirection).From(from).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesAPI.ListWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkspaces`: WorkspaceListResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkspacesAPI.ListWorkspaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** | The Universally Unique Identifier (UUID) of the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodeId** | **string** | Unique identifier of the node. | 
 **filter** | **string** |  | 
 **sortBy** | **string** | Order of the results. | 
 **sortDirection** | **string** |  | 
 **from** | **int32** | Start the response from a given number of result. Used along with &#39;limit&#39; to retrieve the results paginated. | 
 **limit** | **int32** | The maximum number of entries to return. | [default to 50]

### Return type

[**WorkspaceListResponse**](WorkspaceListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

