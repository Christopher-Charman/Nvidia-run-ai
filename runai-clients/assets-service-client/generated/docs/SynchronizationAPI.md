# \SynchronizationAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccessKeySyncInfo**](SynchronizationAPI.md#AccessKeySyncInfo) | **Get** /api/v1/asset/sync/credentials/access-key | Get access keys sync info.
[**AccessKeySyncInfoV2**](SynchronizationAPI.md#AccessKeySyncInfoV2) | **Get** /api/v2/asset/sync/credentials/access-key | Get access keys sync info.
[**CredentialsSyncUpdates**](SynchronizationAPI.md#CredentialsSyncUpdates) | **Post** /api/v1/asset/credentials/updates | Allow cluster to request a list of updates that needs to be synced to the cluster and to send success/failure about the result of previous sync.
[**DatasourceSyncUpdates**](SynchronizationAPI.md#DatasourceSyncUpdates) | **Post** /api/v1/asset/datasource/updates | Allow cluster to request a list of updates that needs to be synced to the cluster and to send success/failure about the result of previous sync.
[**DockerRegistrySyncInfo**](SynchronizationAPI.md#DockerRegistrySyncInfo) | **Get** /api/v1/asset/sync/credentials/docker-registry | Get docker registry sync info.
[**DockerRegistrySyncInfoV2**](SynchronizationAPI.md#DockerRegistrySyncInfoV2) | **Get** /api/v2/asset/sync/credentials/docker-registry | Get docker registry sync info.
[**GenericSecretSyncInfo**](SynchronizationAPI.md#GenericSecretSyncInfo) | **Get** /api/v2/asset/sync/credentials/generic-secret | Get generic secrets sync info.
[**ListAssetsById**](SynchronizationAPI.md#ListAssetsById) | **Get** /api/v1/asset | List assets by their ids
[**PasswordSyncInfo**](SynchronizationAPI.md#PasswordSyncInfo) | **Get** /api/v1/asset/sync/credentials/password | Get passwords sync info.
[**PasswordSyncInfoV2**](SynchronizationAPI.md#PasswordSyncInfoV2) | **Get** /api/v2/asset/sync/credentials/password | Get passwords sync info.
[**PvcSyncInfo**](SynchronizationAPI.md#PvcSyncInfo) | **Get** /api/v1/asset/sync/storage/pvc | Get pvcs sync info.
[**PvcSyncInfoV2**](SynchronizationAPI.md#PvcSyncInfoV2) | **Get** /api/v2/asset/sync/storage/pvc | Get pvcs sync info.
[**SyncAccessKeys**](SynchronizationAPI.md#SyncAccessKeys) | **Post** /api/v1/asset/sync/credentials/access-key | Sync access key assets.
[**SyncAccessKeysV2**](SynchronizationAPI.md#SyncAccessKeysV2) | **Post** /api/v2/asset/sync/credentials/access-key | Sync access key assets.
[**SyncDockerRegistry**](SynchronizationAPI.md#SyncDockerRegistry) | **Post** /api/v1/asset/sync/credentials/docker-registry | Sync docker registry credentials assets.
[**SyncDockerRegistryV2**](SynchronizationAPI.md#SyncDockerRegistryV2) | **Post** /api/v2/asset/sync/credentials/docker-registry | Sync docker registry credentials assets.
[**SyncGenericSecret**](SynchronizationAPI.md#SyncGenericSecret) | **Post** /api/v2/asset/sync/credentials/generic-secret | Sync generic secret assets.
[**SyncPasswords**](SynchronizationAPI.md#SyncPasswords) | **Post** /api/v1/asset/sync/credentials/password | Sync password assets.
[**SyncPasswordsV2**](SynchronizationAPI.md#SyncPasswordsV2) | **Post** /api/v2/asset/sync/credentials/password | Sync password assets.
[**SyncPvcs**](SynchronizationAPI.md#SyncPvcs) | **Post** /api/v1/asset/sync/storage/pvc | Sync pvc assets.
[**SyncPvcsStatus**](SynchronizationAPI.md#SyncPvcsStatus) | **Put** /api/v2/asset/sync/storage/pvc/{clusterUuid}/status | Sync pvc assets status.
[**SyncPvcsV2**](SynchronizationAPI.md#SyncPvcsV2) | **Post** /api/v2/asset/sync/storage/pvc | Sync pvc assets.
[**SyncValidate**](SynchronizationAPI.md#SyncValidate) | **Post** /api/v1/asset/sync/validate | Validate assets sync.



## AccessKeySyncInfo

> SyncInfoResponse AccessKeySyncInfo(ctx).ClusterId(clusterId).Execute()

Get access keys sync info.



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
	clusterId := "d73a738f-fab3-430a-8fa3-5241493d7128" // string | Filter using the Universally Unique Identifier (UUID) of the cluster. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SynchronizationAPI.AccessKeySyncInfo(context.Background()).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationAPI.AccessKeySyncInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessKeySyncInfo`: SyncInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `SynchronizationAPI.AccessKeySyncInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessKeySyncInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Filter using the Universally Unique Identifier (UUID) of the cluster. | 

### Return type

[**SyncInfoResponse**](SyncInfoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccessKeySyncInfoV2

> SyncInfoResponseV2 AccessKeySyncInfoV2(ctx).ClusterId(clusterId).Scope(scope).Execute()

Get access keys sync info.



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
	clusterId := "d73a738f-fab3-430a-8fa3-5241493d7128" // string | Filter using the Universally Unique Identifier (UUID) of the cluster. (optional)
	scope := "scope_example" // string | the scope level of the asset. If empty - tenant scoped assets. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SynchronizationAPI.AccessKeySyncInfoV2(context.Background()).ClusterId(clusterId).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationAPI.AccessKeySyncInfoV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AccessKeySyncInfoV2`: SyncInfoResponseV2
	fmt.Fprintf(os.Stdout, "Response from `SynchronizationAPI.AccessKeySyncInfoV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccessKeySyncInfoV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Filter using the Universally Unique Identifier (UUID) of the cluster. | 
 **scope** | **string** | the scope level of the asset. If empty - tenant scoped assets. | 

### Return type

[**SyncInfoResponseV2**](SyncInfoResponseV2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CredentialsSyncUpdates

> CredentialsUpdatesResp CredentialsSyncUpdates(ctx).ClusterId(clusterId).After(after).ClusterSyncReport(clusterSyncReport).Execute()

Allow cluster to request a list of updates that needs to be synced to the cluster and to send success/failure about the result of previous sync.

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
	clusterId := "d73a738f-fab3-430a-8fa3-5241493d7128" // string | The id of the cluster
	after := int64(123) // int64 | search for updates with update counter greater than the value passed for this parameter (optional)
	clusterSyncReport := *openapiclient.NewClusterSyncReport() // ClusterSyncReport |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SynchronizationAPI.CredentialsSyncUpdates(context.Background()).ClusterId(clusterId).After(after).ClusterSyncReport(clusterSyncReport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationAPI.CredentialsSyncUpdates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CredentialsSyncUpdates`: CredentialsUpdatesResp
	fmt.Fprintf(os.Stdout, "Response from `SynchronizationAPI.CredentialsSyncUpdates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCredentialsSyncUpdatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | The id of the cluster | 
 **after** | **int64** | search for updates with update counter greater than the value passed for this parameter | 
 **clusterSyncReport** | [**ClusterSyncReport**](ClusterSyncReport.md) |  | 

### Return type

[**CredentialsUpdatesResp**](CredentialsUpdatesResp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DatasourceSyncUpdates

> DatasourceUpdatesResp DatasourceSyncUpdates(ctx).ClusterId(clusterId).After(after).ClusterSyncReport(clusterSyncReport).Execute()

Allow cluster to request a list of updates that needs to be synced to the cluster and to send success/failure about the result of previous sync.

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
	clusterId := "d73a738f-fab3-430a-8fa3-5241493d7128" // string | The id of the cluster
	after := int64(123) // int64 | search for updates with update counter greater than the value passed for this parameter (optional)
	clusterSyncReport := *openapiclient.NewClusterSyncReport() // ClusterSyncReport |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SynchronizationAPI.DatasourceSyncUpdates(context.Background()).ClusterId(clusterId).After(after).ClusterSyncReport(clusterSyncReport).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationAPI.DatasourceSyncUpdates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DatasourceSyncUpdates`: DatasourceUpdatesResp
	fmt.Fprintf(os.Stdout, "Response from `SynchronizationAPI.DatasourceSyncUpdates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDatasourceSyncUpdatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | The id of the cluster | 
 **after** | **int64** | search for updates with update counter greater than the value passed for this parameter | 
 **clusterSyncReport** | [**ClusterSyncReport**](ClusterSyncReport.md) |  | 

### Return type

[**DatasourceUpdatesResp**](DatasourceUpdatesResp.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DockerRegistrySyncInfo

> SyncInfoResponse DockerRegistrySyncInfo(ctx).ClusterId(clusterId).Execute()

Get docker registry sync info.



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
	clusterId := "d73a738f-fab3-430a-8fa3-5241493d7128" // string | Filter using the Universally Unique Identifier (UUID) of the cluster. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SynchronizationAPI.DockerRegistrySyncInfo(context.Background()).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationAPI.DockerRegistrySyncInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DockerRegistrySyncInfo`: SyncInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `SynchronizationAPI.DockerRegistrySyncInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDockerRegistrySyncInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Filter using the Universally Unique Identifier (UUID) of the cluster. | 

### Return type

[**SyncInfoResponse**](SyncInfoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DockerRegistrySyncInfoV2

> SyncInfoResponseV2 DockerRegistrySyncInfoV2(ctx).ClusterId(clusterId).Scope(scope).Execute()

Get docker registry sync info.



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
	clusterId := "d73a738f-fab3-430a-8fa3-5241493d7128" // string | Filter using the Universally Unique Identifier (UUID) of the cluster. (optional)
	scope := "scope_example" // string | the scope level of the asset. If empty - tenant scoped assets. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SynchronizationAPI.DockerRegistrySyncInfoV2(context.Background()).ClusterId(clusterId).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationAPI.DockerRegistrySyncInfoV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DockerRegistrySyncInfoV2`: SyncInfoResponseV2
	fmt.Fprintf(os.Stdout, "Response from `SynchronizationAPI.DockerRegistrySyncInfoV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDockerRegistrySyncInfoV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Filter using the Universally Unique Identifier (UUID) of the cluster. | 
 **scope** | **string** | the scope level of the asset. If empty - tenant scoped assets. | 

### Return type

[**SyncInfoResponseV2**](SyncInfoResponseV2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenericSecretSyncInfo

> SyncInfoResponseV2 GenericSecretSyncInfo(ctx).ClusterId(clusterId).Scope(scope).Execute()

Get generic secrets sync info.



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
	clusterId := "d73a738f-fab3-430a-8fa3-5241493d7128" // string | Filter using the Universally Unique Identifier (UUID) of the cluster. (optional)
	scope := "scope_example" // string | the scope level of the asset. If empty - tenant scoped assets. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SynchronizationAPI.GenericSecretSyncInfo(context.Background()).ClusterId(clusterId).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationAPI.GenericSecretSyncInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenericSecretSyncInfo`: SyncInfoResponseV2
	fmt.Fprintf(os.Stdout, "Response from `SynchronizationAPI.GenericSecretSyncInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGenericSecretSyncInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Filter using the Universally Unique Identifier (UUID) of the cluster. | 
 **scope** | **string** | the scope level of the asset. If empty - tenant scoped assets. | 

### Return type

[**SyncInfoResponseV2**](SyncInfoResponseV2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssetsById

> AssetListByIdsResponse ListAssetsById(ctx).Id(id).Execute()

List assets by their ids



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
	id := "id_example" // string | comma seperated list of asset ids

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SynchronizationAPI.ListAssetsById(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationAPI.ListAssetsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAssetsById`: AssetListByIdsResponse
	fmt.Fprintf(os.Stdout, "Response from `SynchronizationAPI.ListAssetsById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAssetsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | comma seperated list of asset ids | 

### Return type

[**AssetListByIdsResponse**](AssetListByIdsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PasswordSyncInfo

> SyncInfoResponse PasswordSyncInfo(ctx).ClusterId(clusterId).Execute()

Get passwords sync info.



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
	clusterId := "d73a738f-fab3-430a-8fa3-5241493d7128" // string | Filter using the Universally Unique Identifier (UUID) of the cluster. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SynchronizationAPI.PasswordSyncInfo(context.Background()).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationAPI.PasswordSyncInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PasswordSyncInfo`: SyncInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `SynchronizationAPI.PasswordSyncInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPasswordSyncInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Filter using the Universally Unique Identifier (UUID) of the cluster. | 

### Return type

[**SyncInfoResponse**](SyncInfoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PasswordSyncInfoV2

> SyncInfoResponseV2 PasswordSyncInfoV2(ctx).ClusterId(clusterId).Scope(scope).Execute()

Get passwords sync info.



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
	clusterId := "d73a738f-fab3-430a-8fa3-5241493d7128" // string | Filter using the Universally Unique Identifier (UUID) of the cluster. (optional)
	scope := "scope_example" // string | the scope level of the asset. If empty - tenant scoped assets. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SynchronizationAPI.PasswordSyncInfoV2(context.Background()).ClusterId(clusterId).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationAPI.PasswordSyncInfoV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PasswordSyncInfoV2`: SyncInfoResponseV2
	fmt.Fprintf(os.Stdout, "Response from `SynchronizationAPI.PasswordSyncInfoV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPasswordSyncInfoV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Filter using the Universally Unique Identifier (UUID) of the cluster. | 
 **scope** | **string** | the scope level of the asset. If empty - tenant scoped assets. | 

### Return type

[**SyncInfoResponseV2**](SyncInfoResponseV2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PvcSyncInfo

> SyncInfoResponse PvcSyncInfo(ctx).ClusterId(clusterId).Execute()

Get pvcs sync info.



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
	clusterId := "d73a738f-fab3-430a-8fa3-5241493d7128" // string | Filter using the Universally Unique Identifier (UUID) of the cluster. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SynchronizationAPI.PvcSyncInfo(context.Background()).ClusterId(clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationAPI.PvcSyncInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PvcSyncInfo`: SyncInfoResponse
	fmt.Fprintf(os.Stdout, "Response from `SynchronizationAPI.PvcSyncInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPvcSyncInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Filter using the Universally Unique Identifier (UUID) of the cluster. | 

### Return type

[**SyncInfoResponse**](SyncInfoResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PvcSyncInfoV2

> SyncInfoResponseV2 PvcSyncInfoV2(ctx).ClusterId(clusterId).Scope(scope).Execute()

Get pvcs sync info.



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
	clusterId := "d73a738f-fab3-430a-8fa3-5241493d7128" // string | Filter using the Universally Unique Identifier (UUID) of the cluster. (optional)
	scope := "scope_example" // string | the scope level of the asset. If empty - tenant scoped assets. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SynchronizationAPI.PvcSyncInfoV2(context.Background()).ClusterId(clusterId).Scope(scope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationAPI.PvcSyncInfoV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PvcSyncInfoV2`: SyncInfoResponseV2
	fmt.Fprintf(os.Stdout, "Response from `SynchronizationAPI.PvcSyncInfoV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPvcSyncInfoV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterId** | **string** | Filter using the Universally Unique Identifier (UUID) of the cluster. | 
 **scope** | **string** | the scope level of the asset. If empty - tenant scoped assets. | 

### Return type

[**SyncInfoResponseV2**](SyncInfoResponseV2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncAccessKeys

> AccessKeysSyncResponse SyncAccessKeys(ctx).SyncRequest(syncRequest).Execute()

Sync access key assets.



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
	syncRequest := *openapiclient.NewSyncRequest() // SyncRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SynchronizationAPI.SyncAccessKeys(context.Background()).SyncRequest(syncRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationAPI.SyncAccessKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncAccessKeys`: AccessKeysSyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SynchronizationAPI.SyncAccessKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncAccessKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncRequest** | [**SyncRequest**](SyncRequest.md) |  | 

### Return type

[**AccessKeysSyncResponse**](AccessKeysSyncResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncAccessKeysV2

> AccessKeysSyncResponse SyncAccessKeysV2(ctx).SyncRequestV2(syncRequestV2).Execute()

Sync access key assets.



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
	syncRequestV2 := *openapiclient.NewSyncRequestV2() // SyncRequestV2 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SynchronizationAPI.SyncAccessKeysV2(context.Background()).SyncRequestV2(syncRequestV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationAPI.SyncAccessKeysV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncAccessKeysV2`: AccessKeysSyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SynchronizationAPI.SyncAccessKeysV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncAccessKeysV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncRequestV2** | [**SyncRequestV2**](SyncRequestV2.md) |  | 

### Return type

[**AccessKeysSyncResponse**](AccessKeysSyncResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncDockerRegistry

> DockerRegistrySyncResponse SyncDockerRegistry(ctx).SyncRequest(syncRequest).Execute()

Sync docker registry credentials assets.



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
	syncRequest := *openapiclient.NewSyncRequest() // SyncRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SynchronizationAPI.SyncDockerRegistry(context.Background()).SyncRequest(syncRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationAPI.SyncDockerRegistry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncDockerRegistry`: DockerRegistrySyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SynchronizationAPI.SyncDockerRegistry`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncDockerRegistryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncRequest** | [**SyncRequest**](SyncRequest.md) |  | 

### Return type

[**DockerRegistrySyncResponse**](DockerRegistrySyncResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncDockerRegistryV2

> DockerRegistrySyncResponse SyncDockerRegistryV2(ctx).SyncRequestV2(syncRequestV2).Execute()

Sync docker registry credentials assets.



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
	syncRequestV2 := *openapiclient.NewSyncRequestV2() // SyncRequestV2 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SynchronizationAPI.SyncDockerRegistryV2(context.Background()).SyncRequestV2(syncRequestV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationAPI.SyncDockerRegistryV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncDockerRegistryV2`: DockerRegistrySyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SynchronizationAPI.SyncDockerRegistryV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncDockerRegistryV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncRequestV2** | [**SyncRequestV2**](SyncRequestV2.md) |  | 

### Return type

[**DockerRegistrySyncResponse**](DockerRegistrySyncResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncGenericSecret

> GenericSecretSyncResponse SyncGenericSecret(ctx).SyncRequestV2(syncRequestV2).Execute()

Sync generic secret assets.



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
	syncRequestV2 := *openapiclient.NewSyncRequestV2() // SyncRequestV2 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SynchronizationAPI.SyncGenericSecret(context.Background()).SyncRequestV2(syncRequestV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationAPI.SyncGenericSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncGenericSecret`: GenericSecretSyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SynchronizationAPI.SyncGenericSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncGenericSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncRequestV2** | [**SyncRequestV2**](SyncRequestV2.md) |  | 

### Return type

[**GenericSecretSyncResponse**](GenericSecretSyncResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncPasswords

> PasswordsSyncResponse SyncPasswords(ctx).SyncRequest(syncRequest).Execute()

Sync password assets.



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
	syncRequest := *openapiclient.NewSyncRequest() // SyncRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SynchronizationAPI.SyncPasswords(context.Background()).SyncRequest(syncRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationAPI.SyncPasswords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncPasswords`: PasswordsSyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SynchronizationAPI.SyncPasswords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPasswordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncRequest** | [**SyncRequest**](SyncRequest.md) |  | 

### Return type

[**PasswordsSyncResponse**](PasswordsSyncResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncPasswordsV2

> PasswordsSyncResponse SyncPasswordsV2(ctx).SyncRequestV2(syncRequestV2).Execute()

Sync password assets.



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
	syncRequestV2 := *openapiclient.NewSyncRequestV2() // SyncRequestV2 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SynchronizationAPI.SyncPasswordsV2(context.Background()).SyncRequestV2(syncRequestV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationAPI.SyncPasswordsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncPasswordsV2`: PasswordsSyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SynchronizationAPI.SyncPasswordsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPasswordsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncRequestV2** | [**SyncRequestV2**](SyncRequestV2.md) |  | 

### Return type

[**PasswordsSyncResponse**](PasswordsSyncResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncPvcs

> PVCsSyncResponse SyncPvcs(ctx).SyncRequest(syncRequest).Execute()

Sync pvc assets.



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
	syncRequest := *openapiclient.NewSyncRequest() // SyncRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SynchronizationAPI.SyncPvcs(context.Background()).SyncRequest(syncRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationAPI.SyncPvcs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncPvcs`: PVCsSyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SynchronizationAPI.SyncPvcs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPvcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncRequest** | [**SyncRequest**](SyncRequest.md) |  | 

### Return type

[**PVCsSyncResponse**](PVCsSyncResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncPvcsStatus

> SyncPvcsStatus(ctx, clusterUuid).PVCStatusRequest(pVCStatusRequest).Execute()

Sync pvc assets status.



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
	pVCStatusRequest := *openapiclient.NewPVCStatusRequest() // PVCStatusRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SynchronizationAPI.SyncPvcsStatus(context.Background(), clusterUuid).PVCStatusRequest(pVCStatusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationAPI.SyncPvcsStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterUuid** | **string** | The Universally Unique Identifier (UUID) of the cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSyncPvcsStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pVCStatusRequest** | [**PVCStatusRequest**](PVCStatusRequest.md) |  | 

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


## SyncPvcsV2

> PVCsSyncResponse SyncPvcsV2(ctx).SyncRequestV2(syncRequestV2).Execute()

Sync pvc assets.



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
	syncRequestV2 := *openapiclient.NewSyncRequestV2() // SyncRequestV2 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SynchronizationAPI.SyncPvcsV2(context.Background()).SyncRequestV2(syncRequestV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationAPI.SyncPvcsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncPvcsV2`: PVCsSyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SynchronizationAPI.SyncPvcsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncPvcsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **syncRequestV2** | [**SyncRequestV2**](SyncRequestV2.md) |  | 

### Return type

[**PVCsSyncResponse**](PVCsSyncResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncValidate

> ValidateSyncResponse SyncValidate(ctx).ValidateSyncRequest(validateSyncRequest).Execute()

Validate assets sync.



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
	validateSyncRequest := *openapiclient.NewValidateSyncRequest("ClusterId_example", "ProjectId_example", openapiclient.AssetKind("compute"), []string{"AssetNames_example"}) // ValidateSyncRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SynchronizationAPI.SyncValidate(context.Background()).ValidateSyncRequest(validateSyncRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SynchronizationAPI.SyncValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SyncValidate`: ValidateSyncResponse
	fmt.Fprintf(os.Stdout, "Response from `SynchronizationAPI.SyncValidate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSyncValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateSyncRequest** | [**ValidateSyncRequest**](ValidateSyncRequest.md) |  | 

### Return type

[**ValidateSyncResponse**](ValidateSyncResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

