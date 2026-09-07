# ExtendedMetricStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**DatasourceId** | **string** | Id of datasource | 
**Read** | [**MetricStoreEndpoint**](MetricStoreEndpoint.md) |  | 
**Write** | [**MetricStoreEndpoint**](MetricStoreEndpoint.md) |  | 

## Methods

### NewExtendedMetricStore

`func NewExtendedMetricStore(type_ string, datasourceId string, read MetricStoreEndpoint, write MetricStoreEndpoint, ) *ExtendedMetricStore`

NewExtendedMetricStore instantiates a new ExtendedMetricStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendedMetricStoreWithDefaults

`func NewExtendedMetricStoreWithDefaults() *ExtendedMetricStore`

NewExtendedMetricStoreWithDefaults instantiates a new ExtendedMetricStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExtendedMetricStore) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExtendedMetricStore) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExtendedMetricStore) SetType(v string)`

SetType sets Type field to given value.


### GetDatasourceId

`func (o *ExtendedMetricStore) GetDatasourceId() string`

GetDatasourceId returns the DatasourceId field if non-nil, zero value otherwise.

### GetDatasourceIdOk

`func (o *ExtendedMetricStore) GetDatasourceIdOk() (*string, bool)`

GetDatasourceIdOk returns a tuple with the DatasourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasourceId

`func (o *ExtendedMetricStore) SetDatasourceId(v string)`

SetDatasourceId sets DatasourceId field to given value.


### GetRead

`func (o *ExtendedMetricStore) GetRead() MetricStoreEndpoint`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *ExtendedMetricStore) GetReadOk() (*MetricStoreEndpoint, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *ExtendedMetricStore) SetRead(v MetricStoreEndpoint)`

SetRead sets Read field to given value.


### GetWrite

`func (o *ExtendedMetricStore) GetWrite() MetricStoreEndpoint`

GetWrite returns the Write field if non-nil, zero value otherwise.

### GetWriteOk

`func (o *ExtendedMetricStore) GetWriteOk() (*MetricStoreEndpoint, bool)`

GetWriteOk returns a tuple with the Write field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrite

`func (o *ExtendedMetricStore) SetWrite(v MetricStoreEndpoint)`

SetWrite sets Write field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


