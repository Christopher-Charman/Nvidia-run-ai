# DistributedStorageFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pvc** | Pointer to [**[]PvcInstance**](PvcInstance.md) | Set of pvc persistent volume claims to use in the workload. | [optional] 
**ConfigMapVolume** | Pointer to [**[]ConfigMapInstance**](ConfigMapInstance.md) | Set of config map volumes to use in the workload | [optional] 
**SecretVolume** | Pointer to [**[]SecretInstance**](SecretInstance.md) | Set of secret volumes to use in the workload | [optional] 

## Methods

### NewDistributedStorageFields

`func NewDistributedStorageFields() *DistributedStorageFields`

NewDistributedStorageFields instantiates a new DistributedStorageFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDistributedStorageFieldsWithDefaults

`func NewDistributedStorageFieldsWithDefaults() *DistributedStorageFields`

NewDistributedStorageFieldsWithDefaults instantiates a new DistributedStorageFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPvc

`func (o *DistributedStorageFields) GetPvc() []PvcInstance`

GetPvc returns the Pvc field if non-nil, zero value otherwise.

### GetPvcOk

`func (o *DistributedStorageFields) GetPvcOk() (*[]PvcInstance, bool)`

GetPvcOk returns a tuple with the Pvc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvc

`func (o *DistributedStorageFields) SetPvc(v []PvcInstance)`

SetPvc sets Pvc field to given value.

### HasPvc

`func (o *DistributedStorageFields) HasPvc() bool`

HasPvc returns a boolean if a field has been set.

### SetPvcNil

`func (o *DistributedStorageFields) SetPvcNil(b bool)`

 SetPvcNil sets the value for Pvc to be an explicit nil

### UnsetPvc
`func (o *DistributedStorageFields) UnsetPvc()`

UnsetPvc ensures that no value is present for Pvc, not even an explicit nil
### GetConfigMapVolume

`func (o *DistributedStorageFields) GetConfigMapVolume() []ConfigMapInstance`

GetConfigMapVolume returns the ConfigMapVolume field if non-nil, zero value otherwise.

### GetConfigMapVolumeOk

`func (o *DistributedStorageFields) GetConfigMapVolumeOk() (*[]ConfigMapInstance, bool)`

GetConfigMapVolumeOk returns a tuple with the ConfigMapVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMapVolume

`func (o *DistributedStorageFields) SetConfigMapVolume(v []ConfigMapInstance)`

SetConfigMapVolume sets ConfigMapVolume field to given value.

### HasConfigMapVolume

`func (o *DistributedStorageFields) HasConfigMapVolume() bool`

HasConfigMapVolume returns a boolean if a field has been set.

### SetConfigMapVolumeNil

`func (o *DistributedStorageFields) SetConfigMapVolumeNil(b bool)`

 SetConfigMapVolumeNil sets the value for ConfigMapVolume to be an explicit nil

### UnsetConfigMapVolume
`func (o *DistributedStorageFields) UnsetConfigMapVolume()`

UnsetConfigMapVolume ensures that no value is present for ConfigMapVolume, not even an explicit nil
### GetSecretVolume

`func (o *DistributedStorageFields) GetSecretVolume() []SecretInstance`

GetSecretVolume returns the SecretVolume field if non-nil, zero value otherwise.

### GetSecretVolumeOk

`func (o *DistributedStorageFields) GetSecretVolumeOk() (*[]SecretInstance, bool)`

GetSecretVolumeOk returns a tuple with the SecretVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretVolume

`func (o *DistributedStorageFields) SetSecretVolume(v []SecretInstance)`

SetSecretVolume sets SecretVolume field to given value.

### HasSecretVolume

`func (o *DistributedStorageFields) HasSecretVolume() bool`

HasSecretVolume returns a boolean if a field has been set.

### SetSecretVolumeNil

`func (o *DistributedStorageFields) SetSecretVolumeNil(b bool)`

 SetSecretVolumeNil sets the value for SecretVolume to be an explicit nil

### UnsetSecretVolume
`func (o *DistributedStorageFields) UnsetSecretVolume()`

UnsetSecretVolume ensures that no value is present for SecretVolume, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


