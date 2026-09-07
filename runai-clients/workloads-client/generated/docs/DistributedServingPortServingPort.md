# DistributedServingPortServingPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **NullableInt32** | The port that the container running the distributed inference service exposes (mandatory). | [optional] 
**Protocol** | Pointer to [**NullableDistributedServingPortProtocol**](DistributedServingPortProtocol.md) |  | [optional] [default to DISTRIBUTEDSERVINGPORTPROTOCOL_HTTP]
**AuthorizationType** | Pointer to [**NullableDistributedServingPortAccessAuthorizationTypeEnum**](DistributedServingPortAccessAuthorizationTypeEnum.md) |  | [optional] 
**AuthorizedUsers** | Pointer to **[]string** | List of users that are allowed to access the URL. Note that authorizedUsers and authorizedGroups are mutually exclusive. | [optional] 
**AuthorizedGroups** | Pointer to **[]string** | List of groups that are allowed to access the URL. Note that authorizedUsers and authorizedGroups are mutually exclusive. | [optional] 
**ExposeExternally** | Pointer to **NullableBool** | Indicates whether the inference serving endpoint should be accessible outside the cluster. If set to true, the endpoint will be exposed externally. | [optional] [default to true]

## Methods

### NewDistributedServingPortServingPort

`func NewDistributedServingPortServingPort() *DistributedServingPortServingPort`

NewDistributedServingPortServingPort instantiates a new DistributedServingPortServingPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributedServingPortServingPortWithDefaults

`func NewDistributedServingPortServingPortWithDefaults() *DistributedServingPortServingPort`

NewDistributedServingPortServingPortWithDefaults instantiates a new DistributedServingPortServingPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *DistributedServingPortServingPort) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *DistributedServingPortServingPort) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *DistributedServingPortServingPort) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *DistributedServingPortServingPort) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *DistributedServingPortServingPort) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *DistributedServingPortServingPort) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetProtocol

`func (o *DistributedServingPortServingPort) GetProtocol() DistributedServingPortProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *DistributedServingPortServingPort) GetProtocolOk() (*DistributedServingPortProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *DistributedServingPortServingPort) SetProtocol(v DistributedServingPortProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *DistributedServingPortServingPort) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *DistributedServingPortServingPort) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *DistributedServingPortServingPort) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetAuthorizationType

`func (o *DistributedServingPortServingPort) GetAuthorizationType() DistributedServingPortAccessAuthorizationTypeEnum`

GetAuthorizationType returns the AuthorizationType field if non-nil, zero value otherwise.

### GetAuthorizationTypeOk

`func (o *DistributedServingPortServingPort) GetAuthorizationTypeOk() (*DistributedServingPortAccessAuthorizationTypeEnum, bool)`

GetAuthorizationTypeOk returns a tuple with the AuthorizationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationType

`func (o *DistributedServingPortServingPort) SetAuthorizationType(v DistributedServingPortAccessAuthorizationTypeEnum)`

SetAuthorizationType sets AuthorizationType field to given value.

### HasAuthorizationType

`func (o *DistributedServingPortServingPort) HasAuthorizationType() bool`

HasAuthorizationType returns a boolean if a field has been set.

### SetAuthorizationTypeNil

`func (o *DistributedServingPortServingPort) SetAuthorizationTypeNil(b bool)`

 SetAuthorizationTypeNil sets the value for AuthorizationType to be an explicit nil

### UnsetAuthorizationType
`func (o *DistributedServingPortServingPort) UnsetAuthorizationType()`

UnsetAuthorizationType ensures that no value is present for AuthorizationType, not even an explicit nil
### GetAuthorizedUsers

`func (o *DistributedServingPortServingPort) GetAuthorizedUsers() []string`

GetAuthorizedUsers returns the AuthorizedUsers field if non-nil, zero value otherwise.

### GetAuthorizedUsersOk

`func (o *DistributedServingPortServingPort) GetAuthorizedUsersOk() (*[]string, bool)`

GetAuthorizedUsersOk returns a tuple with the AuthorizedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedUsers

`func (o *DistributedServingPortServingPort) SetAuthorizedUsers(v []string)`

SetAuthorizedUsers sets AuthorizedUsers field to given value.

### HasAuthorizedUsers

`func (o *DistributedServingPortServingPort) HasAuthorizedUsers() bool`

HasAuthorizedUsers returns a boolean if a field has been set.

### SetAuthorizedUsersNil

`func (o *DistributedServingPortServingPort) SetAuthorizedUsersNil(b bool)`

 SetAuthorizedUsersNil sets the value for AuthorizedUsers to be an explicit nil

### UnsetAuthorizedUsers
`func (o *DistributedServingPortServingPort) UnsetAuthorizedUsers()`

UnsetAuthorizedUsers ensures that no value is present for AuthorizedUsers, not even an explicit nil
### GetAuthorizedGroups

`func (o *DistributedServingPortServingPort) GetAuthorizedGroups() []string`

GetAuthorizedGroups returns the AuthorizedGroups field if non-nil, zero value otherwise.

### GetAuthorizedGroupsOk

`func (o *DistributedServingPortServingPort) GetAuthorizedGroupsOk() (*[]string, bool)`

GetAuthorizedGroupsOk returns a tuple with the AuthorizedGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedGroups

`func (o *DistributedServingPortServingPort) SetAuthorizedGroups(v []string)`

SetAuthorizedGroups sets AuthorizedGroups field to given value.

### HasAuthorizedGroups

`func (o *DistributedServingPortServingPort) HasAuthorizedGroups() bool`

HasAuthorizedGroups returns a boolean if a field has been set.

### SetAuthorizedGroupsNil

`func (o *DistributedServingPortServingPort) SetAuthorizedGroupsNil(b bool)`

 SetAuthorizedGroupsNil sets the value for AuthorizedGroups to be an explicit nil

### UnsetAuthorizedGroups
`func (o *DistributedServingPortServingPort) UnsetAuthorizedGroups()`

UnsetAuthorizedGroups ensures that no value is present for AuthorizedGroups, not even an explicit nil
### GetExposeExternally

`func (o *DistributedServingPortServingPort) GetExposeExternally() bool`

GetExposeExternally returns the ExposeExternally field if non-nil, zero value otherwise.

### GetExposeExternallyOk

`func (o *DistributedServingPortServingPort) GetExposeExternallyOk() (*bool, bool)`

GetExposeExternallyOk returns a tuple with the ExposeExternally field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeExternally

`func (o *DistributedServingPortServingPort) SetExposeExternally(v bool)`

SetExposeExternally sets ExposeExternally field to given value.

### HasExposeExternally

`func (o *DistributedServingPortServingPort) HasExposeExternally() bool`

HasExposeExternally returns a boolean if a field has been set.

### SetExposeExternallyNil

`func (o *DistributedServingPortServingPort) SetExposeExternallyNil(b bool)`

 SetExposeExternallyNil sets the value for ExposeExternally to be an explicit nil

### UnsetExposeExternally
`func (o *DistributedServingPortServingPort) UnsetExposeExternally()`

UnsetExposeExternally ensures that no value is present for ExposeExternally, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


