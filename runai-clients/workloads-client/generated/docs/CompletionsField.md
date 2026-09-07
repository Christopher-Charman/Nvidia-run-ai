# CompletionsField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completions** | Pointer to **NullableInt32** | Used with Hyperparameter Optimization. Specifies the number of successful pods the job should reach to be completed. The Job will be marked as successful once the specified amount of pods has been reached (applicable to standard training only). | [optional] 

## Methods

### NewCompletionsField

`func NewCompletionsField() *CompletionsField`

NewCompletionsField instantiates a new CompletionsField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompletionsFieldWithDefaults

`func NewCompletionsFieldWithDefaults() *CompletionsField`

NewCompletionsFieldWithDefaults instantiates a new CompletionsField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletions

`func (o *CompletionsField) GetCompletions() int32`

GetCompletions returns the Completions field if non-nil, zero value otherwise.

### GetCompletionsOk

`func (o *CompletionsField) GetCompletionsOk() (*int32, bool)`

GetCompletionsOk returns a tuple with the Completions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletions

`func (o *CompletionsField) SetCompletions(v int32)`

SetCompletions sets Completions field to given value.

### HasCompletions

`func (o *CompletionsField) HasCompletions() bool`

HasCompletions returns a boolean if a field has been set.

### SetCompletionsNil

`func (o *CompletionsField) SetCompletionsNil(b bool)`

 SetCompletionsNil sets the value for Completions to be an explicit nil

### UnsetCompletions
`func (o *CompletionsField) UnsetCompletions()`

UnsetCompletions ensures that no value is present for Completions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


