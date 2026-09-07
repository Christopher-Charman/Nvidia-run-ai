# WorkloadTypeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | Pointer to **string** | The unique identifier of the workload category. | [optional] 
**PriorityId** | Pointer to **string** | The unique identifier of the workload priority. | [optional] 
**Name** | **string** | The unique name of the workload type. | 
**Id** | Pointer to **string** | The unique identifier of the workload type. | [optional] 
**CategoryName** | Pointer to **string** | The name of the workload category. | [optional] 
**PriorityName** | Pointer to **string** | The name of the workload priority. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time at which the type has been updated | [optional] 
**UpdatedBy** | Pointer to **string** | The recent user who updated the type | [optional] 

## Methods

### NewWorkloadTypeConfig

`func NewWorkloadTypeConfig(name string, ) *WorkloadTypeConfig`

NewWorkloadTypeConfig instantiates a new WorkloadTypeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadTypeConfigWithDefaults

`func NewWorkloadTypeConfigWithDefaults() *WorkloadTypeConfig`

NewWorkloadTypeConfigWithDefaults instantiates a new WorkloadTypeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *WorkloadTypeConfig) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *WorkloadTypeConfig) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *WorkloadTypeConfig) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *WorkloadTypeConfig) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetPriorityId

`func (o *WorkloadTypeConfig) GetPriorityId() string`

GetPriorityId returns the PriorityId field if non-nil, zero value otherwise.

### GetPriorityIdOk

`func (o *WorkloadTypeConfig) GetPriorityIdOk() (*string, bool)`

GetPriorityIdOk returns a tuple with the PriorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityId

`func (o *WorkloadTypeConfig) SetPriorityId(v string)`

SetPriorityId sets PriorityId field to given value.

### HasPriorityId

`func (o *WorkloadTypeConfig) HasPriorityId() bool`

HasPriorityId returns a boolean if a field has been set.

### GetName

`func (o *WorkloadTypeConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkloadTypeConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkloadTypeConfig) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *WorkloadTypeConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkloadTypeConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkloadTypeConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkloadTypeConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCategoryName

`func (o *WorkloadTypeConfig) GetCategoryName() string`

GetCategoryName returns the CategoryName field if non-nil, zero value otherwise.

### GetCategoryNameOk

`func (o *WorkloadTypeConfig) GetCategoryNameOk() (*string, bool)`

GetCategoryNameOk returns a tuple with the CategoryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryName

`func (o *WorkloadTypeConfig) SetCategoryName(v string)`

SetCategoryName sets CategoryName field to given value.

### HasCategoryName

`func (o *WorkloadTypeConfig) HasCategoryName() bool`

HasCategoryName returns a boolean if a field has been set.

### GetPriorityName

`func (o *WorkloadTypeConfig) GetPriorityName() string`

GetPriorityName returns the PriorityName field if non-nil, zero value otherwise.

### GetPriorityNameOk

`func (o *WorkloadTypeConfig) GetPriorityNameOk() (*string, bool)`

GetPriorityNameOk returns a tuple with the PriorityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityName

`func (o *WorkloadTypeConfig) SetPriorityName(v string)`

SetPriorityName sets PriorityName field to given value.

### HasPriorityName

`func (o *WorkloadTypeConfig) HasPriorityName() bool`

HasPriorityName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WorkloadTypeConfig) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkloadTypeConfig) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkloadTypeConfig) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WorkloadTypeConfig) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *WorkloadTypeConfig) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *WorkloadTypeConfig) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *WorkloadTypeConfig) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *WorkloadTypeConfig) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


