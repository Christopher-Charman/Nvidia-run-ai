# PostTenantEulaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The email of the user | 
**FileHash** | Pointer to **NullableString** | Signed eula hash | [optional] 

## Methods

### NewPostTenantEulaRequest

`func NewPostTenantEulaRequest(email string, ) *PostTenantEulaRequest`

NewPostTenantEulaRequest instantiates a new PostTenantEulaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTenantEulaRequestWithDefaults

`func NewPostTenantEulaRequestWithDefaults() *PostTenantEulaRequest`

NewPostTenantEulaRequestWithDefaults instantiates a new PostTenantEulaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *PostTenantEulaRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PostTenantEulaRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PostTenantEulaRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFileHash

`func (o *PostTenantEulaRequest) GetFileHash() string`

GetFileHash returns the FileHash field if non-nil, zero value otherwise.

### GetFileHashOk

`func (o *PostTenantEulaRequest) GetFileHashOk() (*string, bool)`

GetFileHashOk returns a tuple with the FileHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHash

`func (o *PostTenantEulaRequest) SetFileHash(v string)`

SetFileHash sets FileHash field to given value.

### HasFileHash

`func (o *PostTenantEulaRequest) HasFileHash() bool`

HasFileHash returns a boolean if a field has been set.

### SetFileHashNil

`func (o *PostTenantEulaRequest) SetFileHashNil(b bool)`

 SetFileHashNil sets the value for FileHash to be an explicit nil

### UnsetFileHash
`func (o *PostTenantEulaRequest) UnsetFileHash()`

UnsetFileHash ensures that no value is present for FileHash, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


