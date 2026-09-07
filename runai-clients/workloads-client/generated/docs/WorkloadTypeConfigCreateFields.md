# WorkloadTypeConfigCreateFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | Pointer to **string** | The unique identifier of the workload category. | [optional] 
**PriorityId** | Pointer to **string** | The unique identifier of the workload priority. | [optional] 
**Name** | **string** | The unique name of the workload type. | 

## Methods

### NewWorkloadTypeConfigCreateFields

`func NewWorkloadTypeConfigCreateFields(name string, ) *WorkloadTypeConfigCreateFields`

NewWorkloadTypeConfigCreateFields instantiates a new WorkloadTypeConfigCreateFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkloadTypeConfigCreateFieldsWithDefaults

`func NewWorkloadTypeConfigCreateFieldsWithDefaults() *WorkloadTypeConfigCreateFields`

NewWorkloadTypeConfigCreateFieldsWithDefaults instantiates a new WorkloadTypeConfigCreateFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *WorkloadTypeConfigCreateFields) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *WorkloadTypeConfigCreateFields) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *WorkloadTypeConfigCreateFields) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *WorkloadTypeConfigCreateFields) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetPriorityId

`func (o *WorkloadTypeConfigCreateFields) GetPriorityId() string`

GetPriorityId returns the PriorityId field if non-nil, zero value otherwise.

### GetPriorityIdOk

`func (o *WorkloadTypeConfigCreateFields) GetPriorityIdOk() (*string, bool)`

GetPriorityIdOk returns a tuple with the PriorityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityId

`func (o *WorkloadTypeConfigCreateFields) SetPriorityId(v string)`

SetPriorityId sets PriorityId field to given value.

### HasPriorityId

`func (o *WorkloadTypeConfigCreateFields) HasPriorityId() bool`

HasPriorityId returns a boolean if a field has been set.

### GetName

`func (o *WorkloadTypeConfigCreateFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkloadTypeConfigCreateFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkloadTypeConfigCreateFields) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


