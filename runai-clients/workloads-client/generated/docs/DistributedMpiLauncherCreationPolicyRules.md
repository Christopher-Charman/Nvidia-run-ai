# DistributedMpiLauncherCreationPolicyRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceOfRule** | Pointer to [**NullableSourceOfRule**](SourceOfRule.md) |  | [optional] 
**Required** | Pointer to **NullableBool** | Whether the field is mandatory, default to false. | [optional] 
**CanEdit** | Pointer to **NullableBool** | Whether the value of the field is editable, default to true | [optional] 
**Options** | Pointer to [**[]DistributedMpiLauncherCreationPolicyOptionsOptionsInner**](DistributedMpiLauncherCreationPolicyOptionsOptionsInner.md) |  | [optional] 

## Methods

### NewDistributedMpiLauncherCreationPolicyRules

`func NewDistributedMpiLauncherCreationPolicyRules() *DistributedMpiLauncherCreationPolicyRules`

NewDistributedMpiLauncherCreationPolicyRules instantiates a new DistributedMpiLauncherCreationPolicyRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributedMpiLauncherCreationPolicyRulesWithDefaults

`func NewDistributedMpiLauncherCreationPolicyRulesWithDefaults() *DistributedMpiLauncherCreationPolicyRules`

NewDistributedMpiLauncherCreationPolicyRulesWithDefaults instantiates a new DistributedMpiLauncherCreationPolicyRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceOfRule

`func (o *DistributedMpiLauncherCreationPolicyRules) GetSourceOfRule() SourceOfRule`

GetSourceOfRule returns the SourceOfRule field if non-nil, zero value otherwise.

### GetSourceOfRuleOk

`func (o *DistributedMpiLauncherCreationPolicyRules) GetSourceOfRuleOk() (*SourceOfRule, bool)`

GetSourceOfRuleOk returns a tuple with the SourceOfRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceOfRule

`func (o *DistributedMpiLauncherCreationPolicyRules) SetSourceOfRule(v SourceOfRule)`

SetSourceOfRule sets SourceOfRule field to given value.

### HasSourceOfRule

`func (o *DistributedMpiLauncherCreationPolicyRules) HasSourceOfRule() bool`

HasSourceOfRule returns a boolean if a field has been set.

### SetSourceOfRuleNil

`func (o *DistributedMpiLauncherCreationPolicyRules) SetSourceOfRuleNil(b bool)`

 SetSourceOfRuleNil sets the value for SourceOfRule to be an explicit nil

### UnsetSourceOfRule
`func (o *DistributedMpiLauncherCreationPolicyRules) UnsetSourceOfRule()`

UnsetSourceOfRule ensures that no value is present for SourceOfRule, not even an explicit nil
### GetRequired

`func (o *DistributedMpiLauncherCreationPolicyRules) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *DistributedMpiLauncherCreationPolicyRules) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *DistributedMpiLauncherCreationPolicyRules) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *DistributedMpiLauncherCreationPolicyRules) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *DistributedMpiLauncherCreationPolicyRules) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *DistributedMpiLauncherCreationPolicyRules) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil
### GetCanEdit

`func (o *DistributedMpiLauncherCreationPolicyRules) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *DistributedMpiLauncherCreationPolicyRules) GetCanEditOk() (*bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEdit

`func (o *DistributedMpiLauncherCreationPolicyRules) SetCanEdit(v bool)`

SetCanEdit sets CanEdit field to given value.

### HasCanEdit

`func (o *DistributedMpiLauncherCreationPolicyRules) HasCanEdit() bool`

HasCanEdit returns a boolean if a field has been set.

### SetCanEditNil

`func (o *DistributedMpiLauncherCreationPolicyRules) SetCanEditNil(b bool)`

 SetCanEditNil sets the value for CanEdit to be an explicit nil

### UnsetCanEdit
`func (o *DistributedMpiLauncherCreationPolicyRules) UnsetCanEdit()`

UnsetCanEdit ensures that no value is present for CanEdit, not even an explicit nil
### GetOptions

`func (o *DistributedMpiLauncherCreationPolicyRules) GetOptions() []DistributedMpiLauncherCreationPolicyOptionsOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DistributedMpiLauncherCreationPolicyRules) GetOptionsOk() (*[]DistributedMpiLauncherCreationPolicyOptionsOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DistributedMpiLauncherCreationPolicyRules) SetOptions(v []DistributedMpiLauncherCreationPolicyOptionsOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *DistributedMpiLauncherCreationPolicyRules) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *DistributedMpiLauncherCreationPolicyRules) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *DistributedMpiLauncherCreationPolicyRules) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


