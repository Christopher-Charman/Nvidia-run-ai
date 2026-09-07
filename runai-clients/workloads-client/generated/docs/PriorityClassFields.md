# PriorityClassFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriorityClass** | Pointer to **NullableString** | Priority class of the workload. The default value for workspace is &#39;build&#39; and it can be changed to &#39;interactive-preemptible&#39; to allow the workload to use over-quota resources. The default value for training is &#39;train&#39; and it can be changed to &#39;build&#39; to allow the training workload to have a higher priority for in-queue scheduling and also become non-preemptive (if it&#39;s in deserved quota).. | [optional] 

## Methods

### NewPriorityClassFields

`func NewPriorityClassFields() *PriorityClassFields`

NewPriorityClassFields instantiates a new PriorityClassFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriorityClassFieldsWithDefaults

`func NewPriorityClassFieldsWithDefaults() *PriorityClassFields`

NewPriorityClassFieldsWithDefaults instantiates a new PriorityClassFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriorityClass

`func (o *PriorityClassFields) GetPriorityClass() string`

GetPriorityClass returns the PriorityClass field if non-nil, zero value otherwise.

### GetPriorityClassOk

`func (o *PriorityClassFields) GetPriorityClassOk() (*string, bool)`

GetPriorityClassOk returns a tuple with the PriorityClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityClass

`func (o *PriorityClassFields) SetPriorityClass(v string)`

SetPriorityClass sets PriorityClass field to given value.

### HasPriorityClass

`func (o *PriorityClassFields) HasPriorityClass() bool`

HasPriorityClass returns a boolean if a field has been set.

### SetPriorityClassNil

`func (o *PriorityClassFields) SetPriorityClassNil(b bool)`

 SetPriorityClassNil sets the value for PriorityClass to be an explicit nil

### UnsetPriorityClass
`func (o *PriorityClassFields) UnsetPriorityClass()`

UnsetPriorityClass ensures that no value is present for PriorityClass, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


