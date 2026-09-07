# HuggingFaceModelInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelName** | **string** | the model name | 
**PipelineTag** | Pointer to **NullableString** | the pipeline tag of the model | [optional] 
**Architectures** | Pointer to **[]string** | the list of model architectures | [optional] 
**VllmSupportingVersions** | Pointer to **[]string** | the list of vLLM versions that support this model&#39;s architecture | [optional] 
**MinMemoryRequirementsMb** | Pointer to **NullableInt32** | the minimum memory requirements, in MB, for using this model for inference (won&#39;t be returned if it can&#39;t be computed for the provided model) | [optional] 

## Methods

### NewHuggingFaceModelInfo

`func NewHuggingFaceModelInfo(modelName string, ) *HuggingFaceModelInfo`

NewHuggingFaceModelInfo instantiates a new HuggingFaceModelInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHuggingFaceModelInfoWithDefaults

`func NewHuggingFaceModelInfoWithDefaults() *HuggingFaceModelInfo`

NewHuggingFaceModelInfoWithDefaults instantiates a new HuggingFaceModelInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelName

`func (o *HuggingFaceModelInfo) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *HuggingFaceModelInfo) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *HuggingFaceModelInfo) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetPipelineTag

`func (o *HuggingFaceModelInfo) GetPipelineTag() string`

GetPipelineTag returns the PipelineTag field if non-nil, zero value otherwise.

### GetPipelineTagOk

`func (o *HuggingFaceModelInfo) GetPipelineTagOk() (*string, bool)`

GetPipelineTagOk returns a tuple with the PipelineTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineTag

`func (o *HuggingFaceModelInfo) SetPipelineTag(v string)`

SetPipelineTag sets PipelineTag field to given value.

### HasPipelineTag

`func (o *HuggingFaceModelInfo) HasPipelineTag() bool`

HasPipelineTag returns a boolean if a field has been set.

### SetPipelineTagNil

`func (o *HuggingFaceModelInfo) SetPipelineTagNil(b bool)`

 SetPipelineTagNil sets the value for PipelineTag to be an explicit nil

### UnsetPipelineTag
`func (o *HuggingFaceModelInfo) UnsetPipelineTag()`

UnsetPipelineTag ensures that no value is present for PipelineTag, not even an explicit nil
### GetArchitectures

`func (o *HuggingFaceModelInfo) GetArchitectures() []string`

GetArchitectures returns the Architectures field if non-nil, zero value otherwise.

### GetArchitecturesOk

`func (o *HuggingFaceModelInfo) GetArchitecturesOk() (*[]string, bool)`

GetArchitecturesOk returns a tuple with the Architectures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitectures

`func (o *HuggingFaceModelInfo) SetArchitectures(v []string)`

SetArchitectures sets Architectures field to given value.

### HasArchitectures

`func (o *HuggingFaceModelInfo) HasArchitectures() bool`

HasArchitectures returns a boolean if a field has been set.

### SetArchitecturesNil

`func (o *HuggingFaceModelInfo) SetArchitecturesNil(b bool)`

 SetArchitecturesNil sets the value for Architectures to be an explicit nil

### UnsetArchitectures
`func (o *HuggingFaceModelInfo) UnsetArchitectures()`

UnsetArchitectures ensures that no value is present for Architectures, not even an explicit nil
### GetVllmSupportingVersions

`func (o *HuggingFaceModelInfo) GetVllmSupportingVersions() []string`

GetVllmSupportingVersions returns the VllmSupportingVersions field if non-nil, zero value otherwise.

### GetVllmSupportingVersionsOk

`func (o *HuggingFaceModelInfo) GetVllmSupportingVersionsOk() (*[]string, bool)`

GetVllmSupportingVersionsOk returns a tuple with the VllmSupportingVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVllmSupportingVersions

`func (o *HuggingFaceModelInfo) SetVllmSupportingVersions(v []string)`

SetVllmSupportingVersions sets VllmSupportingVersions field to given value.

### HasVllmSupportingVersions

`func (o *HuggingFaceModelInfo) HasVllmSupportingVersions() bool`

HasVllmSupportingVersions returns a boolean if a field has been set.

### SetVllmSupportingVersionsNil

`func (o *HuggingFaceModelInfo) SetVllmSupportingVersionsNil(b bool)`

 SetVllmSupportingVersionsNil sets the value for VllmSupportingVersions to be an explicit nil

### UnsetVllmSupportingVersions
`func (o *HuggingFaceModelInfo) UnsetVllmSupportingVersions()`

UnsetVllmSupportingVersions ensures that no value is present for VllmSupportingVersions, not even an explicit nil
### GetMinMemoryRequirementsMb

`func (o *HuggingFaceModelInfo) GetMinMemoryRequirementsMb() int32`

GetMinMemoryRequirementsMb returns the MinMemoryRequirementsMb field if non-nil, zero value otherwise.

### GetMinMemoryRequirementsMbOk

`func (o *HuggingFaceModelInfo) GetMinMemoryRequirementsMbOk() (*int32, bool)`

GetMinMemoryRequirementsMbOk returns a tuple with the MinMemoryRequirementsMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemoryRequirementsMb

`func (o *HuggingFaceModelInfo) SetMinMemoryRequirementsMb(v int32)`

SetMinMemoryRequirementsMb sets MinMemoryRequirementsMb field to given value.

### HasMinMemoryRequirementsMb

`func (o *HuggingFaceModelInfo) HasMinMemoryRequirementsMb() bool`

HasMinMemoryRequirementsMb returns a boolean if a field has been set.

### SetMinMemoryRequirementsMbNil

`func (o *HuggingFaceModelInfo) SetMinMemoryRequirementsMbNil(b bool)`

 SetMinMemoryRequirementsMbNil sets the value for MinMemoryRequirementsMb to be an explicit nil

### UnsetMinMemoryRequirementsMb
`func (o *HuggingFaceModelInfo) UnsetMinMemoryRequirementsMb()`

UnsetMinMemoryRequirementsMb ensures that no value is present for MinMemoryRequirementsMb, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


