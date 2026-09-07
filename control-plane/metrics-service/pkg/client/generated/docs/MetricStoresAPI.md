# \MetricStoresAPI

All URIs are relative to *https://app.run.ai*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMetricStore**](MetricStoresAPI.md#CreateMetricStore) | **Post** /internal/metric-stores | Create/Save metric store for tenant
[**GetMetricStores**](MetricStoresAPI.md#GetMetricStores) | **Get** /internal/metric-stores | Get tenant metric stores



## CreateMetricStore

> TenantMetricStore CreateMetricStore(ctx).TenantMetricStore(tenantMetricStore).Execute()

Create/Save metric store for tenant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/control-plane/metrics-service/pkg/client"
)

func main() {
	tenantMetricStore := *openapiclient.NewTenantMetricStore(int32(123)) // TenantMetricStore | metrics store

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricStoresAPI.CreateMetricStore(context.Background()).TenantMetricStore(tenantMetricStore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricStoresAPI.CreateMetricStore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMetricStore`: TenantMetricStore
	fmt.Fprintf(os.Stdout, "Response from `MetricStoresAPI.CreateMetricStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMetricStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tenantMetricStore** | [**TenantMetricStore**](TenantMetricStore.md) | metrics store | 

### Return type

[**TenantMetricStore**](TenantMetricStore.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricStores

> []TenantMetricStore GetMetricStores(ctx).OverrideTenantId(overrideTenantId).Execute()

Get tenant metric stores



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/dsx-ai-factory/run-ai/control-plane/metrics-service/pkg/client"
)

func main() {
	overrideTenantId := int32(123) // int32 | The tenant Id of the current request, overrides the tenantId in the header. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MetricStoresAPI.GetMetricStores(context.Background()).OverrideTenantId(overrideTenantId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MetricStoresAPI.GetMetricStores``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMetricStores`: []TenantMetricStore
	fmt.Fprintf(os.Stdout, "Response from `MetricStoresAPI.GetMetricStores`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricStoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **overrideTenantId** | **int32** | The tenant Id of the current request, overrides the tenantId in the header. | 

### Return type

[**[]TenantMetricStore**](TenantMetricStore.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

