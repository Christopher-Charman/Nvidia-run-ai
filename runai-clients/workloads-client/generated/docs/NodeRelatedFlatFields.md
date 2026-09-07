# NodeRelatedFlatFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodePools** | Pointer to **[]string** | A prioritized list of node pools for the scheduler to run the workspace on. The scheduler will always try to use the first node pool before moving to the next one if the first is not available. | [optional] 

## Methods

### NewNodeRelatedFlatFields

`func NewNodeRelatedFlatFields() *NodeRelatedFlatFields`

NewNodeRelatedFlatFields instantiates a new NodeRelatedFlatFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeRelatedFlatFieldsWithDefaults

`func NewNodeRelatedFlatFieldsWithDefaults() *NodeRelatedFlatFields`

NewNodeRelatedFlatFieldsWithDefaults instantiates a new NodeRelatedFlatFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodePools

`func (o *NodeRelatedFlatFields) GetNodePools() []string`

GetNodePools returns the NodePools field if non-nil, zero value otherwise.

### GetNodePoolsOk

`func (o *NodeRelatedFlatFields) GetNodePoolsOk() (*[]string, bool)`

GetNodePoolsOk returns a tuple with the NodePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePools

`func (o *NodeRelatedFlatFields) SetNodePools(v []string)`

SetNodePools sets NodePools field to given value.

### HasNodePools

`func (o *NodeRelatedFlatFields) HasNodePools() bool`

HasNodePools returns a boolean if a field has been set.

### SetNodePoolsNil

`func (o *NodeRelatedFlatFields) SetNodePoolsNil(b bool)`

 SetNodePoolsNil sets the value for NodePools to be an explicit nil

### UnsetNodePools
`func (o *NodeRelatedFlatFields) UnsetNodePools()`

UnsetNodePools ensures that no value is present for NodePools, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


