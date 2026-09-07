# MetricStoreEndpointAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** |  | 
**Password** | **string** |  | 
**Type** | Pointer to [**MetricStoreEndpointAuthType**](MetricStoreEndpointAuthType.md) |  | [optional] 

## Methods

### NewMetricStoreEndpointAuth

`func NewMetricStoreEndpointAuth(username string, password string, ) *MetricStoreEndpointAuth`

NewMetricStoreEndpointAuth instantiates a new MetricStoreEndpointAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricStoreEndpointAuthWithDefaults

`func NewMetricStoreEndpointAuthWithDefaults() *MetricStoreEndpointAuth`

NewMetricStoreEndpointAuthWithDefaults instantiates a new MetricStoreEndpointAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *MetricStoreEndpointAuth) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *MetricStoreEndpointAuth) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *MetricStoreEndpointAuth) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *MetricStoreEndpointAuth) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MetricStoreEndpointAuth) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MetricStoreEndpointAuth) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetType

`func (o *MetricStoreEndpointAuth) GetType() MetricStoreEndpointAuthType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricStoreEndpointAuth) GetTypeOk() (*MetricStoreEndpointAuthType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricStoreEndpointAuth) SetType(v MetricStoreEndpointAuthType)`

SetType sets Type field to given value.

### HasType

`func (o *MetricStoreEndpointAuth) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


