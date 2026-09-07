# DepartmentGetRanking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the department | 
**Id** | **string** | The id of the department | 
**AllocationValue** | **NullableFloat64** | The allocation value of the department | 
**UtilizationValue** | **NullableFloat64** | The utilization value of the department | 
**MemoryUtilizationValue** | Pointer to **NullableFloat64** | The memory utilization value of the project | [optional] 

## Methods

### NewDepartmentGetRanking

`func NewDepartmentGetRanking(name string, id string, allocationValue NullableFloat64, utilizationValue NullableFloat64, ) *DepartmentGetRanking`

NewDepartmentGetRanking instantiates a new DepartmentGetRanking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepartmentGetRankingWithDefaults

`func NewDepartmentGetRankingWithDefaults() *DepartmentGetRanking`

NewDepartmentGetRankingWithDefaults instantiates a new DepartmentGetRanking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DepartmentGetRanking) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DepartmentGetRanking) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DepartmentGetRanking) SetName(v string)`

SetName sets Name field to given value.


### GetId

`func (o *DepartmentGetRanking) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DepartmentGetRanking) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DepartmentGetRanking) SetId(v string)`

SetId sets Id field to given value.


### GetAllocationValue

`func (o *DepartmentGetRanking) GetAllocationValue() float64`

GetAllocationValue returns the AllocationValue field if non-nil, zero value otherwise.

### GetAllocationValueOk

`func (o *DepartmentGetRanking) GetAllocationValueOk() (*float64, bool)`

GetAllocationValueOk returns a tuple with the AllocationValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationValue

`func (o *DepartmentGetRanking) SetAllocationValue(v float64)`

SetAllocationValue sets AllocationValue field to given value.


### SetAllocationValueNil

`func (o *DepartmentGetRanking) SetAllocationValueNil(b bool)`

 SetAllocationValueNil sets the value for AllocationValue to be an explicit nil

### UnsetAllocationValue
`func (o *DepartmentGetRanking) UnsetAllocationValue()`

UnsetAllocationValue ensures that no value is present for AllocationValue, not even an explicit nil
### GetUtilizationValue

`func (o *DepartmentGetRanking) GetUtilizationValue() float64`

GetUtilizationValue returns the UtilizationValue field if non-nil, zero value otherwise.

### GetUtilizationValueOk

`func (o *DepartmentGetRanking) GetUtilizationValueOk() (*float64, bool)`

GetUtilizationValueOk returns a tuple with the UtilizationValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtilizationValue

`func (o *DepartmentGetRanking) SetUtilizationValue(v float64)`

SetUtilizationValue sets UtilizationValue field to given value.


### SetUtilizationValueNil

`func (o *DepartmentGetRanking) SetUtilizationValueNil(b bool)`

 SetUtilizationValueNil sets the value for UtilizationValue to be an explicit nil

### UnsetUtilizationValue
`func (o *DepartmentGetRanking) UnsetUtilizationValue()`

UnsetUtilizationValue ensures that no value is present for UtilizationValue, not even an explicit nil
### GetMemoryUtilizationValue

`func (o *DepartmentGetRanking) GetMemoryUtilizationValue() float64`

GetMemoryUtilizationValue returns the MemoryUtilizationValue field if non-nil, zero value otherwise.

### GetMemoryUtilizationValueOk

`func (o *DepartmentGetRanking) GetMemoryUtilizationValueOk() (*float64, bool)`

GetMemoryUtilizationValueOk returns a tuple with the MemoryUtilizationValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUtilizationValue

`func (o *DepartmentGetRanking) SetMemoryUtilizationValue(v float64)`

SetMemoryUtilizationValue sets MemoryUtilizationValue field to given value.

### HasMemoryUtilizationValue

`func (o *DepartmentGetRanking) HasMemoryUtilizationValue() bool`

HasMemoryUtilizationValue returns a boolean if a field has been set.

### SetMemoryUtilizationValueNil

`func (o *DepartmentGetRanking) SetMemoryUtilizationValueNil(b bool)`

 SetMemoryUtilizationValueNil sets the value for MemoryUtilizationValue to be an explicit nil

### UnsetMemoryUtilizationValue
`func (o *DepartmentGetRanking) UnsetMemoryUtilizationValue()`

UnsetMemoryUtilizationValue ensures that no value is present for MemoryUtilizationValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


