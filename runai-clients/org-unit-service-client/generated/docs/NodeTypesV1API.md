# \NodeTypesV1API

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNodeTypeV1**](NodeTypesV1API.md#CreateNodeTypeV1) | **Post** /v1/k8s/clusters/{clusterId}/nodetypes | Create a new NodeType.
[**DeleteNodeTypeV1**](NodeTypesV1API.md#DeleteNodeTypeV1) | **Delete** /v1/k8s/clusters/{clusterId}/nodetypes/{nodeTypeId} | Delete a NodeType.
[**GetNodeTypeV1**](NodeTypesV1API.md#GetNodeTypeV1) | **Get** /v1/k8s/clusters/{clusterId}/nodetypes/{nodeTypeId} | Get a NodeType.
[**GetNodeTypesV1**](NodeTypesV1API.md#GetNodeTypesV1) | **Get** /v1/k8s/clusters/{clusterId}/nodetypes | Get all cluster&#39;s NodeTypes.
[**UpdateNodeTypeV1**](NodeTypesV1API.md#UpdateNodeTypeV1) | **Put** /v1/k8s/clusters/{clusterId}/nodetypes/{nodeTypeId} | Update a NodeType.



## CreateNodeTypeV1

> NodeTypeV1 CreateNodeTypeV1(ctx, clusterId).NodeTypeV1(nodeTypeV1).Execute()

Create a new NodeType.

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
	clusterId := "9f55255e-11ed-47c7-acef-fc4054768dbc" // string | The Universally Unique Identifier (UUID) of the cluster.
	nodeTypeV1 := *openapiclient.NewNodeTypeV1() // NodeTypeV1 | The NodeType to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeTypesV1API.CreateNodeTypeV1(context.Background(), clusterId).NodeTypeV1(nodeTypeV1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeTypesV1API.CreateNodeTypeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNodeTypeV1`: NodeTypeV1
	fmt.Fprintf(os.Stdout, "Response from `NodeTypesV1API.CreateNodeTypeV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Universally Unique Identifier (UUID) of the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodeTypeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodeTypeV1** | [**NodeTypeV1**](NodeTypeV1.md) | The NodeType to create. | 

### Return type

[**NodeTypeV1**](NodeTypeV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNodeTypeV1

> DeleteNodeTypeV1(ctx, clusterId, nodeTypeId).Execute()

Delete a NodeType.

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
	clusterId := "9f55255e-11ed-47c7-acef-fc4054768dbc" // string | The Universally Unique Identifier (UUID) of the cluster.
	nodeTypeId := "1" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NodeTypesV1API.DeleteNodeTypeV1(context.Background(), clusterId, nodeTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeTypesV1API.DeleteNodeTypeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Universally Unique Identifier (UUID) of the cluster. | 
**nodeTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodeTypeV1Request struct via the builder pattern


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


## GetNodeTypeV1

> NodeTypeV1 GetNodeTypeV1(ctx, clusterId, nodeTypeId).Execute()

Get a NodeType.

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
	clusterId := "9f55255e-11ed-47c7-acef-fc4054768dbc" // string | The Universally Unique Identifier (UUID) of the cluster.
	nodeTypeId := "1" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeTypesV1API.GetNodeTypeV1(context.Background(), clusterId, nodeTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeTypesV1API.GetNodeTypeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodeTypeV1`: NodeTypeV1
	fmt.Fprintf(os.Stdout, "Response from `NodeTypesV1API.GetNodeTypeV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Universally Unique Identifier (UUID) of the cluster. | 
**nodeTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeTypeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NodeTypeV1**](NodeTypeV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeTypesV1

> []NodeTypeV1 GetNodeTypesV1(ctx, clusterId).Execute()

Get all cluster's NodeTypes.

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
	clusterId := "9f55255e-11ed-47c7-acef-fc4054768dbc" // string | The Universally Unique Identifier (UUID) of the cluster.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeTypesV1API.GetNodeTypesV1(context.Background(), clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeTypesV1API.GetNodeTypesV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodeTypesV1`: []NodeTypeV1
	fmt.Fprintf(os.Stdout, "Response from `NodeTypesV1API.GetNodeTypesV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Universally Unique Identifier (UUID) of the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeTypesV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]NodeTypeV1**](NodeTypeV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNodeTypeV1

> NodeTypeV1 UpdateNodeTypeV1(ctx, clusterId, nodeTypeId).NodeTypeV1(nodeTypeV1).Execute()

Update a NodeType.

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
	clusterId := "9f55255e-11ed-47c7-acef-fc4054768dbc" // string | The Universally Unique Identifier (UUID) of the cluster.
	nodeTypeId := "1" // string | 
	nodeTypeV1 := *openapiclient.NewNodeTypeV1() // NodeTypeV1 | The NodeType to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeTypesV1API.UpdateNodeTypeV1(context.Background(), clusterId, nodeTypeId).NodeTypeV1(nodeTypeV1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeTypesV1API.UpdateNodeTypeV1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNodeTypeV1`: NodeTypeV1
	fmt.Fprintf(os.Stdout, "Response from `NodeTypesV1API.UpdateNodeTypeV1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **string** | The Universally Unique Identifier (UUID) of the cluster. | 
**nodeTypeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNodeTypeV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nodeTypeV1** | [**NodeTypeV1**](NodeTypeV1.md) | The NodeType to create. | 

### Return type

[**NodeTypeV1**](NodeTypeV1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

