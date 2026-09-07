# \InstallationAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClusterInstallationUrl**](InstallationAPI.md#ClusterInstallationUrl) | **Get** /api/v1/clusters/{clusterUuid}/installation-url | Get cluster installation url
[**MapClusterUrl**](InstallationAPI.md#MapClusterUrl) | **Post** /api/v1/clusters/{clusterUuid}/domain | Map target url to cluster url



## ClusterInstallationUrl

> ClusterInstallationResponse ClusterInstallationUrl(ctx, clusterUuid).Execute()

Get cluster installation url

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/cluster-service-client"
)

func main() {
	clusterUuid := "9f55255e-11ed-47c7-acef-fc4054768dbc" // string | The Universally Unique Identifier (UUID) of the cluster.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstallationAPI.ClusterInstallationUrl(context.Background(), clusterUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallationAPI.ClusterInstallationUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClusterInstallationUrl`: ClusterInstallationResponse
	fmt.Fprintf(os.Stdout, "Response from `InstallationAPI.ClusterInstallationUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** | The Universally Unique Identifier (UUID) of the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClusterInstallationUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterInstallationResponse**](ClusterInstallationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MapClusterUrl

> ClusterUrlResponse MapClusterUrl(ctx, clusterUuid).ClusterUrlRequest(clusterUrlRequest).Execute()

Map target url to cluster url

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/runai-clients/cluster-service-client"
)

func main() {
	clusterUuid := "9f55255e-11ed-47c7-acef-fc4054768dbc" // string | The Universally Unique Identifier (UUID) of the cluster.
	clusterUrlRequest := *openapiclient.NewClusterUrlRequest("TargetHost_example") // ClusterUrlRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstallationAPI.MapClusterUrl(context.Background(), clusterUuid).ClusterUrlRequest(clusterUrlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstallationAPI.MapClusterUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MapClusterUrl`: ClusterUrlResponse
	fmt.Fprintf(os.Stdout, "Response from `InstallationAPI.MapClusterUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** | The Universally Unique Identifier (UUID) of the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMapClusterUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **clusterUrlRequest** | [**ClusterUrlRequest**](ClusterUrlRequest.md) |  | 

### Return type

[**ClusterUrlResponse**](ClusterUrlResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

