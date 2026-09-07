# ParallelismField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parallelism** | Pointer to **NullableInt32** | Used with Hyperparameter Optimization. Specifies the maximum number of pods the workload should run at any given time (applicable to standard training only). | [optional] 

## Methods

### NewParallelismField

`func NewParallelismField() *ParallelismField`

NewParallelismField instantiates a new ParallelismField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParallelismFieldWithDefaults

`func NewParallelismFieldWithDefaults() *ParallelismField`

NewParallelismFieldWithDefaults instantiates a new ParallelismField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParallelism

`func (o *ParallelismField) GetParallelism() int32`

GetParallelism returns the Parallelism field if non-nil, zero value otherwise.

### GetParallelismOk

`func (o *ParallelismField) GetParallelismOk() (*int32, bool)`

GetParallelismOk returns a tuple with the Parallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelism

`func (o *ParallelismField) SetParallelism(v int32)`

SetParallelism sets Parallelism field to given value.

### HasParallelism

`func (o *ParallelismField) HasParallelism() bool`

HasParallelism returns a boolean if a field has been set.

### SetParallelismNil

`func (o *ParallelismField) SetParallelismNil(b bool)`

 SetParallelismNil sets the value for Parallelism to be an explicit nil

### UnsetParallelism
`func (o *ParallelismField) UnsetParallelism()`

UnsetParallelism ensures that no value is present for Parallelism, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


