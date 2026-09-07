# GetConsumptionReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**GetConsumptionReportValue**](GetConsumptionReportValue.md) |  | [optional] 
**DisplayName** | Pointer to **string** | Human-readable name for this setting | [optional] [readonly] 
**Description** | Pointer to **string** | Description of what this setting controls | [optional] [readonly] 
**Stage** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewGetConsumptionReport

`func NewGetConsumptionReport() *GetConsumptionReport`

NewGetConsumptionReport instantiates a new GetConsumptionReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConsumptionReportWithDefaults

`func NewGetConsumptionReportWithDefaults() *GetConsumptionReport`

NewGetConsumptionReportWithDefaults instantiates a new GetConsumptionReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *GetConsumptionReport) GetValue() GetConsumptionReportValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetConsumptionReport) GetValueOk() (*GetConsumptionReportValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetConsumptionReport) SetValue(v GetConsumptionReportValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetConsumptionReport) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetConsumptionReport) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetConsumptionReport) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetConsumptionReport) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetConsumptionReport) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *GetConsumptionReport) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetConsumptionReport) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetConsumptionReport) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetConsumptionReport) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStage

`func (o *GetConsumptionReport) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *GetConsumptionReport) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *GetConsumptionReport) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *GetConsumptionReport) HasStage() bool`

HasStage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


