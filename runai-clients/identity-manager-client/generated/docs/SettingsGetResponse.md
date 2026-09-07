# SettingsGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRedirectSSO** | Pointer to [**AutoRedirectSsoSetting**](AutoRedirectSsoSetting.md) |  | [optional] 
**ExcludeGroupsFromTokens** | Pointer to [**ExcludeGroupsFromTokenSetting**](ExcludeGroupsFromTokenSetting.md) |  | [optional] 
**BrowserSessionTimeout** | Pointer to [**BrowserSessionTimeoutSetting**](BrowserSessionTimeoutSetting.md) |  | [optional] 
**LogoutRedirectUri** | Pointer to [**LogoutRedirectUriSetting**](LogoutRedirectUriSetting.md) |  | [optional] 

## Methods

### NewSettingsGetResponse

`func NewSettingsGetResponse() *SettingsGetResponse`

NewSettingsGetResponse instantiates a new SettingsGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsGetResponseWithDefaults

`func NewSettingsGetResponseWithDefaults() *SettingsGetResponse`

NewSettingsGetResponseWithDefaults instantiates a new SettingsGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoRedirectSSO

`func (o *SettingsGetResponse) GetAutoRedirectSSO() AutoRedirectSsoSetting`

GetAutoRedirectSSO returns the AutoRedirectSSO field if non-nil, zero value otherwise.

### GetAutoRedirectSSOOk

`func (o *SettingsGetResponse) GetAutoRedirectSSOOk() (*AutoRedirectSsoSetting, bool)`

GetAutoRedirectSSOOk returns a tuple with the AutoRedirectSSO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRedirectSSO

`func (o *SettingsGetResponse) SetAutoRedirectSSO(v AutoRedirectSsoSetting)`

SetAutoRedirectSSO sets AutoRedirectSSO field to given value.

### HasAutoRedirectSSO

`func (o *SettingsGetResponse) HasAutoRedirectSSO() bool`

HasAutoRedirectSSO returns a boolean if a field has been set.

### GetExcludeGroupsFromTokens

`func (o *SettingsGetResponse) GetExcludeGroupsFromTokens() ExcludeGroupsFromTokenSetting`

GetExcludeGroupsFromTokens returns the ExcludeGroupsFromTokens field if non-nil, zero value otherwise.

### GetExcludeGroupsFromTokensOk

`func (o *SettingsGetResponse) GetExcludeGroupsFromTokensOk() (*ExcludeGroupsFromTokenSetting, bool)`

GetExcludeGroupsFromTokensOk returns a tuple with the ExcludeGroupsFromTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeGroupsFromTokens

`func (o *SettingsGetResponse) SetExcludeGroupsFromTokens(v ExcludeGroupsFromTokenSetting)`

SetExcludeGroupsFromTokens sets ExcludeGroupsFromTokens field to given value.

### HasExcludeGroupsFromTokens

`func (o *SettingsGetResponse) HasExcludeGroupsFromTokens() bool`

HasExcludeGroupsFromTokens returns a boolean if a field has been set.

### GetBrowserSessionTimeout

`func (o *SettingsGetResponse) GetBrowserSessionTimeout() BrowserSessionTimeoutSetting`

GetBrowserSessionTimeout returns the BrowserSessionTimeout field if non-nil, zero value otherwise.

### GetBrowserSessionTimeoutOk

`func (o *SettingsGetResponse) GetBrowserSessionTimeoutOk() (*BrowserSessionTimeoutSetting, bool)`

GetBrowserSessionTimeoutOk returns a tuple with the BrowserSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserSessionTimeout

`func (o *SettingsGetResponse) SetBrowserSessionTimeout(v BrowserSessionTimeoutSetting)`

SetBrowserSessionTimeout sets BrowserSessionTimeout field to given value.

### HasBrowserSessionTimeout

`func (o *SettingsGetResponse) HasBrowserSessionTimeout() bool`

HasBrowserSessionTimeout returns a boolean if a field has been set.

### GetLogoutRedirectUri

`func (o *SettingsGetResponse) GetLogoutRedirectUri() LogoutRedirectUriSetting`

GetLogoutRedirectUri returns the LogoutRedirectUri field if non-nil, zero value otherwise.

### GetLogoutRedirectUriOk

`func (o *SettingsGetResponse) GetLogoutRedirectUriOk() (*LogoutRedirectUriSetting, bool)`

GetLogoutRedirectUriOk returns a tuple with the LogoutRedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutRedirectUri

`func (o *SettingsGetResponse) SetLogoutRedirectUri(v LogoutRedirectUriSetting)`

SetLogoutRedirectUri sets LogoutRedirectUri field to given value.

### HasLogoutRedirectUri

`func (o *SettingsGetResponse) HasLogoutRedirectUri() bool`

HasLogoutRedirectUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


