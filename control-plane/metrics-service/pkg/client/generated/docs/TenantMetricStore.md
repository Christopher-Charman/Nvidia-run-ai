# TenantMetricStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **int32** | The id of the tenant. | 
**DatasourceId** | Pointer to **string** | Id of datasource | [optional] 
**MetricStore** | Pointer to [**MetricStore**](MetricStore.md) |  | [optional] 

## Methods

### NewTenantMetricStore

`func NewTenantMetricStore(tenantId int32, ) *TenantMetricStore`

NewTenantMetricStore instantiates a new TenantMetricStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantMetricStoreWithDefaults

`func NewTenantMetricStoreWithDefaults() *TenantMetricStore`

NewTenantMetricStoreWithDefaults instantiates a new TenantMetricStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *TenantMetricStore) GetTenantId() int32`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantMetricStore) GetTenantIdOk() (*int32, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantMetricStore) SetTenantId(v int32)`

SetTenantId sets TenantId field to given value.


### GetDatasourceId

`func (o *TenantMetricStore) GetDatasourceId() string`

GetDatasourceId returns the DatasourceId field if non-nil, zero value otherwise.

### GetDatasourceIdOk

`func (o *TenantMetricStore) GetDatasourceIdOk() (*string, bool)`

GetDatasourceIdOk returns a tuple with the DatasourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceId

`func (o *TenantMetricStore) SetDatasourceId(v string)`

SetDatasourceId sets DatasourceId field to given value.

### HasDatasourceId

`func (o *TenantMetricStore) HasDatasourceId() bool`

HasDatasourceId returns a boolean if a field has been set.

### GetMetricStore

`func (o *TenantMetricStore) GetMetricStore() MetricStore`

GetMetricStore returns the MetricStore field if non-nil, zero value otherwise.

### GetMetricStoreOk

`func (o *TenantMetricStore) GetMetricStoreOk() (*MetricStore, bool)`

GetMetricStoreOk returns a tuple with the MetricStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricStore

`func (o *TenantMetricStore) SetMetricStore(v MetricStore)`

SetMetricStore sets MetricStore field to given value.

### HasMetricStore

`func (o *TenantMetricStore) HasMetricStore() bool`

HasMetricStore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


