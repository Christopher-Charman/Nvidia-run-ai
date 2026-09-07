# DistributedServingPortAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationType** | Pointer to [**NullableDistributedServingPortAccessAuthorizationTypeEnum**](DistributedServingPortAccessAuthorizationTypeEnum.md) |  | [optional] 
**AuthorizedUsers** | Pointer to **[]string** | List of users that are allowed to access the URL. Note that authorizedUsers and authorizedGroups are mutually exclusive. | [optional] 
**AuthorizedGroups** | Pointer to **[]string** | List of groups that are allowed to access the URL. Note that authorizedUsers and authorizedGroups are mutually exclusive. | [optional] 
**ExposeExternally** | Pointer to **NullableBool** | Indicates whether the inference serving endpoint should be accessible outside the cluster. If set to true, the endpoint will be exposed externally. | [optional] [default to true]

## Methods

### NewDistributedServingPortAccess

`func NewDistributedServingPortAccess() *DistributedServingPortAccess`

NewDistributedServingPortAccess instantiates a new DistributedServingPortAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributedServingPortAccessWithDefaults

`func NewDistributedServingPortAccessWithDefaults() *DistributedServingPortAccess`

NewDistributedServingPortAccessWithDefaults instantiates a new DistributedServingPortAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationType

`func (o *DistributedServingPortAccess) GetAuthorizationType() DistributedServingPortAccessAuthorizationTypeEnum`

GetAuthorizationType returns the AuthorizationType field if non-nil, zero value otherwise.

### GetAuthorizationTypeOk

`func (o *DistributedServingPortAccess) GetAuthorizationTypeOk() (*DistributedServingPortAccessAuthorizationTypeEnum, bool)`

GetAuthorizationTypeOk returns a tuple with the AuthorizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationType

`func (o *DistributedServingPortAccess) SetAuthorizationType(v DistributedServingPortAccessAuthorizationTypeEnum)`

SetAuthorizationType sets AuthorizationType field to given value.

### HasAuthorizationType

`func (o *DistributedServingPortAccess) HasAuthorizationType() bool`

HasAuthorizationType returns a boolean if a field has been set.

### SetAuthorizationTypeNil

`func (o *DistributedServingPortAccess) SetAuthorizationTypeNil(b bool)`

 SetAuthorizationTypeNil sets the value for AuthorizationType to be an explicit nil

### UnsetAuthorizationType
`func (o *DistributedServingPortAccess) UnsetAuthorizationType()`

UnsetAuthorizationType ensures that no value is present for AuthorizationType, not even an explicit nil
### GetAuthorizedUsers

`func (o *DistributedServingPortAccess) GetAuthorizedUsers() []string`

GetAuthorizedUsers returns the AuthorizedUsers field if non-nil, zero value otherwise.

### GetAuthorizedUsersOk

`func (o *DistributedServingPortAccess) GetAuthorizedUsersOk() (*[]string, bool)`

GetAuthorizedUsersOk returns a tuple with the AuthorizedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedUsers

`func (o *DistributedServingPortAccess) SetAuthorizedUsers(v []string)`

SetAuthorizedUsers sets AuthorizedUsers field to given value.

### HasAuthorizedUsers

`func (o *DistributedServingPortAccess) HasAuthorizedUsers() bool`

HasAuthorizedUsers returns a boolean if a field has been set.

### SetAuthorizedUsersNil

`func (o *DistributedServingPortAccess) SetAuthorizedUsersNil(b bool)`

 SetAuthorizedUsersNil sets the value for AuthorizedUsers to be an explicit nil

### UnsetAuthorizedUsers
`func (o *DistributedServingPortAccess) UnsetAuthorizedUsers()`

UnsetAuthorizedUsers ensures that no value is present for AuthorizedUsers, not even an explicit nil
### GetAuthorizedGroups

`func (o *DistributedServingPortAccess) GetAuthorizedGroups() []string`

GetAuthorizedGroups returns the AuthorizedGroups field if non-nil, zero value otherwise.

### GetAuthorizedGroupsOk

`func (o *DistributedServingPortAccess) GetAuthorizedGroupsOk() (*[]string, bool)`

GetAuthorizedGroupsOk returns a tuple with the AuthorizedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedGroups

`func (o *DistributedServingPortAccess) SetAuthorizedGroups(v []string)`

SetAuthorizedGroups sets AuthorizedGroups field to given value.

### HasAuthorizedGroups

`func (o *DistributedServingPortAccess) HasAuthorizedGroups() bool`

HasAuthorizedGroups returns a boolean if a field has been set.

### SetAuthorizedGroupsNil

`func (o *DistributedServingPortAccess) SetAuthorizedGroupsNil(b bool)`

 SetAuthorizedGroupsNil sets the value for AuthorizedGroups to be an explicit nil

### UnsetAuthorizedGroups
`func (o *DistributedServingPortAccess) UnsetAuthorizedGroups()`

UnsetAuthorizedGroups ensures that no value is present for AuthorizedGroups, not even an explicit nil
### GetExposeExternally

`func (o *DistributedServingPortAccess) GetExposeExternally() bool`

GetExposeExternally returns the ExposeExternally field if non-nil, zero value otherwise.

### GetExposeExternallyOk

`func (o *DistributedServingPortAccess) GetExposeExternallyOk() (*bool, bool)`

GetExposeExternallyOk returns a tuple with the ExposeExternally field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeExternally

`func (o *DistributedServingPortAccess) SetExposeExternally(v bool)`

SetExposeExternally sets ExposeExternally field to given value.

### HasExposeExternally

`func (o *DistributedServingPortAccess) HasExposeExternally() bool`

HasExposeExternally returns a boolean if a field has been set.

### SetExposeExternallyNil

`func (o *DistributedServingPortAccess) SetExposeExternallyNil(b bool)`

 SetExposeExternallyNil sets the value for ExposeExternally to be an explicit nil

### UnsetExposeExternally
`func (o *DistributedServingPortAccess) UnsetExposeExternally()`

UnsetExposeExternally ensures that no value is present for ExposeExternally, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


