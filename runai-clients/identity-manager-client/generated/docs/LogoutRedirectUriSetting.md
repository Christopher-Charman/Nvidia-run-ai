# LogoutRedirectUriSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uri** | Pointer to **NullableString** | The URI where users will be redirected after logging out of the application | [optional] 

## Methods

### NewLogoutRedirectUriSetting

`func NewLogoutRedirectUriSetting() *LogoutRedirectUriSetting`

NewLogoutRedirectUriSetting instantiates a new LogoutRedirectUriSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogoutRedirectUriSettingWithDefaults

`func NewLogoutRedirectUriSettingWithDefaults() *LogoutRedirectUriSetting`

NewLogoutRedirectUriSettingWithDefaults instantiates a new LogoutRedirectUriSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUri

`func (o *LogoutRedirectUriSetting) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *LogoutRedirectUriSetting) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *LogoutRedirectUriSetting) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *LogoutRedirectUriSetting) HasUri() bool`

HasUri returns a boolean if a field has been set.

### SetUriNil

`func (o *LogoutRedirectUriSetting) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *LogoutRedirectUriSetting) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


