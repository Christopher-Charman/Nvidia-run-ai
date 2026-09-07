# AutoRedirectSsoSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **NullableBool** |  | [optional] 
**IdpAlias** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAutoRedirectSsoSetting

`func NewAutoRedirectSsoSetting() *AutoRedirectSsoSetting`

NewAutoRedirectSsoSetting instantiates a new AutoRedirectSsoSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoRedirectSsoSettingWithDefaults

`func NewAutoRedirectSsoSettingWithDefaults() *AutoRedirectSsoSetting`

NewAutoRedirectSsoSettingWithDefaults instantiates a new AutoRedirectSsoSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AutoRedirectSsoSetting) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AutoRedirectSsoSetting) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AutoRedirectSsoSetting) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AutoRedirectSsoSetting) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *AutoRedirectSsoSetting) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *AutoRedirectSsoSetting) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetIdpAlias

`func (o *AutoRedirectSsoSetting) GetIdpAlias() string`

GetIdpAlias returns the IdpAlias field if non-nil, zero value otherwise.

### GetIdpAliasOk

`func (o *AutoRedirectSsoSetting) GetIdpAliasOk() (*string, bool)`

GetIdpAliasOk returns a tuple with the IdpAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpAlias

`func (o *AutoRedirectSsoSetting) SetIdpAlias(v string)`

SetIdpAlias sets IdpAlias field to given value.

### HasIdpAlias

`func (o *AutoRedirectSsoSetting) HasIdpAlias() bool`

HasIdpAlias returns a boolean if a field has been set.

### SetIdpAliasNil

`func (o *AutoRedirectSsoSetting) SetIdpAliasNil(b bool)`

 SetIdpAliasNil sets the value for IdpAlias to be an explicit nil

### UnsetIdpAlias
`func (o *AutoRedirectSsoSetting) UnsetIdpAlias()`

UnsetIdpAlias ensures that no value is present for IdpAlias, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


