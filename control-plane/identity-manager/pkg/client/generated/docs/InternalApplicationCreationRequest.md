# InternalApplicationCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the application. The name must be unique within the organization and can only contain lowercase alphanumeric characters and hyphens. It must start and end with a letter. | 
**Realm** | **string** |  | 

## Methods

### NewInternalApplicationCreationRequest

`func NewInternalApplicationCreationRequest(name string, realm string, ) *InternalApplicationCreationRequest`

NewInternalApplicationCreationRequest instantiates a new InternalApplicationCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalApplicationCreationRequestWithDefaults

`func NewInternalApplicationCreationRequestWithDefaults() *InternalApplicationCreationRequest`

NewInternalApplicationCreationRequestWithDefaults instantiates a new InternalApplicationCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InternalApplicationCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternalApplicationCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternalApplicationCreationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRealm

`func (o *InternalApplicationCreationRequest) GetRealm() string`

GetRealm returns the Realm field if non-nil, zero value otherwise.

### GetRealmOk

`func (o *InternalApplicationCreationRequest) GetRealmOk() (*string, bool)`

GetRealmOk returns a tuple with the Realm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealm

`func (o *InternalApplicationCreationRequest) SetRealm(v string)`

SetRealm sets Realm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


