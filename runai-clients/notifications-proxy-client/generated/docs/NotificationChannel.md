# NotificationChannel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**NotificationChannelType**](NotificationChannelType.md) |  | [default to NOTIFICATIONCHANNELTYPE_EMAIL]
**Config** | [**NotificationChannelConfig**](NotificationChannelConfig.md) |  | 
**Name** | **string** |  | 

## Methods

### NewNotificationChannel

`func NewNotificationChannel(type_ NotificationChannelType, config NotificationChannelConfig, name string, ) *NotificationChannel`

NewNotificationChannel instantiates a new NotificationChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationChannelWithDefaults

`func NewNotificationChannelWithDefaults() *NotificationChannel`

NewNotificationChannelWithDefaults instantiates a new NotificationChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationChannel) GetType() NotificationChannelType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationChannel) GetTypeOk() (*NotificationChannelType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationChannel) SetType(v NotificationChannelType)`

SetType sets Type field to given value.


### GetConfig

`func (o *NotificationChannel) GetConfig() NotificationChannelConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NotificationChannel) GetConfigOk() (*NotificationChannelConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NotificationChannel) SetConfig(v NotificationChannelConfig)`

SetConfig sets Config field to given value.


### GetName

`func (o *NotificationChannel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationChannel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationChannel) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


