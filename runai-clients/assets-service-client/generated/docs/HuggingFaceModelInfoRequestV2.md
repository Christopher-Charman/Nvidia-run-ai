# HuggingFaceModelInfoRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelName** | **string** | the model name | 
**HfToken** | Pointer to **NullableString** | the Hugging Face token, required only for gated models | [optional] 

## Methods

### NewHuggingFaceModelInfoRequestV2

`func NewHuggingFaceModelInfoRequestV2(modelName string, ) *HuggingFaceModelInfoRequestV2`

NewHuggingFaceModelInfoRequestV2 instantiates a new HuggingFaceModelInfoRequestV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHuggingFaceModelInfoRequestV2WithDefaults

`func NewHuggingFaceModelInfoRequestV2WithDefaults() *HuggingFaceModelInfoRequestV2`

NewHuggingFaceModelInfoRequestV2WithDefaults instantiates a new HuggingFaceModelInfoRequestV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelName

`func (o *HuggingFaceModelInfoRequestV2) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *HuggingFaceModelInfoRequestV2) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *HuggingFaceModelInfoRequestV2) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetHfToken

`func (o *HuggingFaceModelInfoRequestV2) GetHfToken() string`

GetHfToken returns the HfToken field if non-nil, zero value otherwise.

### GetHfTokenOk

`func (o *HuggingFaceModelInfoRequestV2) GetHfTokenOk() (*string, bool)`

GetHfTokenOk returns a tuple with the HfToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHfToken

`func (o *HuggingFaceModelInfoRequestV2) SetHfToken(v string)`

SetHfToken sets HfToken field to given value.

### HasHfToken

`func (o *HuggingFaceModelInfoRequestV2) HasHfToken() bool`

HasHfToken returns a boolean if a field has been set.

### SetHfTokenNil

`func (o *HuggingFaceModelInfoRequestV2) SetHfTokenNil(b bool)`

 SetHfTokenNil sets the value for HfToken to be an explicit nil

### UnsetHfToken
`func (o *HuggingFaceModelInfoRequestV2) UnsetHfToken()`

UnsetHfToken ensures that no value is present for HfToken, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


