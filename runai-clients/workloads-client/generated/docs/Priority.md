# Priority

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the workload priority | 
**Source** | Pointer to **string** |  | [optional] 
**Name** | **string** | The name of the workload priority (e.g., Low, Medium High) | 
**Value** | **int32** | The scheduling value assigned to this priority | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**AggregatedPhase** | Pointer to **string** | Combined status of the priority across all clusters | [optional] 
**Status** | Pointer to [**[]PriorityStatus**](PriorityStatus.md) | Per-cluster status details | [optional] 
**ScopeType** | Pointer to **string** |  | [optional] 
**ScopeId** | Pointer to **string** |  | [optional] 

## Methods

### NewPriority

`func NewPriority(id string, name string, value int32, ) *Priority`

NewPriority instantiates a new Priority object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriorityWithDefaults

`func NewPriorityWithDefaults() *Priority`

NewPriorityWithDefaults instantiates a new Priority object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Priority) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Priority) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Priority) SetId(v string)`

SetId sets Id field to given value.


### GetSource

`func (o *Priority) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Priority) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Priority) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Priority) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetName

`func (o *Priority) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Priority) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Priority) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *Priority) GetValue() int32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Priority) GetValueOk() (*int32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Priority) SetValue(v int32)`

SetValue sets Value field to given value.


### GetCreatedAt

`func (o *Priority) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Priority) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Priority) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Priority) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Priority) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Priority) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Priority) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Priority) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAggregatedPhase

`func (o *Priority) GetAggregatedPhase() string`

GetAggregatedPhase returns the AggregatedPhase field if non-nil, zero value otherwise.

### GetAggregatedPhaseOk

`func (o *Priority) GetAggregatedPhaseOk() (*string, bool)`

GetAggregatedPhaseOk returns a tuple with the AggregatedPhase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatedPhase

`func (o *Priority) SetAggregatedPhase(v string)`

SetAggregatedPhase sets AggregatedPhase field to given value.

### HasAggregatedPhase

`func (o *Priority) HasAggregatedPhase() bool`

HasAggregatedPhase returns a boolean if a field has been set.

### GetStatus

`func (o *Priority) GetStatus() []PriorityStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Priority) GetStatusOk() (*[]PriorityStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Priority) SetStatus(v []PriorityStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Priority) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetScopeType

`func (o *Priority) GetScopeType() string`

GetScopeType returns the ScopeType field if non-nil, zero value otherwise.

### GetScopeTypeOk

`func (o *Priority) GetScopeTypeOk() (*string, bool)`

GetScopeTypeOk returns a tuple with the ScopeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeType

`func (o *Priority) SetScopeType(v string)`

SetScopeType sets ScopeType field to given value.

### HasScopeType

`func (o *Priority) HasScopeType() bool`

HasScopeType returns a boolean if a field has been set.

### GetScopeId

`func (o *Priority) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *Priority) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *Priority) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *Priority) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


