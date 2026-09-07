# SuperSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServingPort** | Pointer to [**NullableSuperSetServingPort**](SuperSetServingPort.md) |  | [optional] 
**RestartPolicy** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSuperSet

`func NewSuperSet() *SuperSet`

NewSuperSet instantiates a new SuperSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuperSetWithDefaults

`func NewSuperSetWithDefaults() *SuperSet`

NewSuperSetWithDefaults instantiates a new SuperSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServingPort

`func (o *SuperSet) GetServingPort() SuperSetServingPort`

GetServingPort returns the ServingPort field if non-nil, zero value otherwise.

### GetServingPortOk

`func (o *SuperSet) GetServingPortOk() (*SuperSetServingPort, bool)`

GetServingPortOk returns a tuple with the ServingPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingPort

`func (o *SuperSet) SetServingPort(v SuperSetServingPort)`

SetServingPort sets ServingPort field to given value.

### HasServingPort

`func (o *SuperSet) HasServingPort() bool`

HasServingPort returns a boolean if a field has been set.

### SetServingPortNil

`func (o *SuperSet) SetServingPortNil(b bool)`

 SetServingPortNil sets the value for ServingPort to be an explicit nil

### UnsetServingPort
`func (o *SuperSet) UnsetServingPort()`

UnsetServingPort ensures that no value is present for ServingPort, not even an explicit nil
### GetRestartPolicy

`func (o *SuperSet) GetRestartPolicy() string`

GetRestartPolicy returns the RestartPolicy field if non-nil, zero value otherwise.

### GetRestartPolicyOk

`func (o *SuperSet) GetRestartPolicyOk() (*string, bool)`

GetRestartPolicyOk returns a tuple with the RestartPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestartPolicy

`func (o *SuperSet) SetRestartPolicy(v string)`

SetRestartPolicy sets RestartPolicy field to given value.

### HasRestartPolicy

`func (o *SuperSet) HasRestartPolicy() bool`

HasRestartPolicy returns a boolean if a field has been set.

### SetRestartPolicyNil

`func (o *SuperSet) SetRestartPolicyNil(b bool)`

 SetRestartPolicyNil sets the value for RestartPolicy to be an explicit nil

### UnsetRestartPolicy
`func (o *SuperSet) UnsetRestartPolicy()`

UnsetRestartPolicy ensures that no value is present for RestartPolicy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


