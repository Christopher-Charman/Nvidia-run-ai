# NodepoolDeprecatedFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OverProvisioningRatio** | Pointer to **NullableInt32** | Deprecated top-level field kept ONLY in the GET response so older runai-cluster agents that still read them continue to work. | [optional] [default to 1]
**PlacementStrategy** | Pointer to [**NodepoolDeprecatedPlacementStrategy**](NodepoolDeprecatedPlacementStrategy.md) |  | [optional] 

## Methods

### NewNodepoolDeprecatedFields

`func NewNodepoolDeprecatedFields() *NodepoolDeprecatedFields`

NewNodepoolDeprecatedFields instantiates a new NodepoolDeprecatedFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodepoolDeprecatedFieldsWithDefaults

`func NewNodepoolDeprecatedFieldsWithDefaults() *NodepoolDeprecatedFields`

NewNodepoolDeprecatedFieldsWithDefaults instantiates a new NodepoolDeprecatedFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverProvisioningRatio

`func (o *NodepoolDeprecatedFields) GetOverProvisioningRatio() int32`

GetOverProvisioningRatio returns the OverProvisioningRatio field if non-nil, zero value otherwise.

### GetOverProvisioningRatioOk

`func (o *NodepoolDeprecatedFields) GetOverProvisioningRatioOk() (*int32, bool)`

GetOverProvisioningRatioOk returns a tuple with the OverProvisioningRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverProvisioningRatio

`func (o *NodepoolDeprecatedFields) SetOverProvisioningRatio(v int32)`

SetOverProvisioningRatio sets OverProvisioningRatio field to given value.

### HasOverProvisioningRatio

`func (o *NodepoolDeprecatedFields) HasOverProvisioningRatio() bool`

HasOverProvisioningRatio returns a boolean if a field has been set.

### SetOverProvisioningRatioNil

`func (o *NodepoolDeprecatedFields) SetOverProvisioningRatioNil(b bool)`

 SetOverProvisioningRatioNil sets the value for OverProvisioningRatio to be an explicit nil

### UnsetOverProvisioningRatio
`func (o *NodepoolDeprecatedFields) UnsetOverProvisioningRatio()`

UnsetOverProvisioningRatio ensures that no value is present for OverProvisioningRatio, not even an explicit nil
### GetPlacementStrategy

`func (o *NodepoolDeprecatedFields) GetPlacementStrategy() NodepoolDeprecatedPlacementStrategy`

GetPlacementStrategy returns the PlacementStrategy field if non-nil, zero value otherwise.

### GetPlacementStrategyOk

`func (o *NodepoolDeprecatedFields) GetPlacementStrategyOk() (*NodepoolDeprecatedPlacementStrategy, bool)`

GetPlacementStrategyOk returns a tuple with the PlacementStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementStrategy

`func (o *NodepoolDeprecatedFields) SetPlacementStrategy(v NodepoolDeprecatedPlacementStrategy)`

SetPlacementStrategy sets PlacementStrategy field to given value.

### HasPlacementStrategy

`func (o *NodepoolDeprecatedFields) HasPlacementStrategy() bool`

HasPlacementStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


