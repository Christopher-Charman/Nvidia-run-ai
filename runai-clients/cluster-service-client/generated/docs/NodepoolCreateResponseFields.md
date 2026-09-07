# NodepoolCreateResponseFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**LabelKey** | **string** | Label key for associated nodes to the Node Pool (with value as in labelValue) | 
**LabelValue** | **string** | Label value for associated nodes to the Node Pool (with key as in labelKey) | 
**ClusterId** | **string** |  | 
**OverProvisioningRatio** | Pointer to **NullableInt32** |  | [optional] [default to 1]
**PlacementStrategy** | Pointer to [**NodepoolCreateResponseFieldsPlacementStrategy**](NodepoolCreateResponseFieldsPlacementStrategy.md) |  | [optional] 
**GpuNetworkAccelerationLabelKey** | Pointer to **string** | Label key by which to determine GPUNetworkAccelerationDetection nodes | [optional] 
**GpuNetworkAccelerationDetection** | Pointer to [**NullableNodepoolGPUNetworkAccelerationDetection**](NodepoolGPUNetworkAccelerationDetection.md) |  | [optional] [default to NODEPOOLGPUNETWORKACCELERATIONDETECTION_AUTO]
**GpuResourceOptimization** | Pointer to [**NullableGPUResourceOptimization**](GPUResourceOptimization.md) |  | [optional] 

## Methods

### NewNodepoolCreateResponseFields

`func NewNodepoolCreateResponseFields(name string, labelKey string, labelValue string, clusterId string, ) *NodepoolCreateResponseFields`

NewNodepoolCreateResponseFields instantiates a new NodepoolCreateResponseFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodepoolCreateResponseFieldsWithDefaults

`func NewNodepoolCreateResponseFieldsWithDefaults() *NodepoolCreateResponseFields`

NewNodepoolCreateResponseFieldsWithDefaults instantiates a new NodepoolCreateResponseFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NodepoolCreateResponseFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodepoolCreateResponseFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodepoolCreateResponseFields) SetName(v string)`

SetName sets Name field to given value.


### GetLabelKey

`func (o *NodepoolCreateResponseFields) GetLabelKey() string`

GetLabelKey returns the LabelKey field if non-nil, zero value otherwise.

### GetLabelKeyOk

`func (o *NodepoolCreateResponseFields) GetLabelKeyOk() (*string, bool)`

GetLabelKeyOk returns a tuple with the LabelKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelKey

`func (o *NodepoolCreateResponseFields) SetLabelKey(v string)`

SetLabelKey sets LabelKey field to given value.


### GetLabelValue

`func (o *NodepoolCreateResponseFields) GetLabelValue() string`

GetLabelValue returns the LabelValue field if non-nil, zero value otherwise.

### GetLabelValueOk

`func (o *NodepoolCreateResponseFields) GetLabelValueOk() (*string, bool)`

GetLabelValueOk returns a tuple with the LabelValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelValue

`func (o *NodepoolCreateResponseFields) SetLabelValue(v string)`

SetLabelValue sets LabelValue field to given value.


### GetClusterId

`func (o *NodepoolCreateResponseFields) GetClusterId() string`

GetClusterId returns the ClusterId field if non-nil, zero value otherwise.

### GetClusterIdOk

`func (o *NodepoolCreateResponseFields) GetClusterIdOk() (*string, bool)`

GetClusterIdOk returns a tuple with the ClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterId

`func (o *NodepoolCreateResponseFields) SetClusterId(v string)`

SetClusterId sets ClusterId field to given value.


### GetOverProvisioningRatio

`func (o *NodepoolCreateResponseFields) GetOverProvisioningRatio() int32`

GetOverProvisioningRatio returns the OverProvisioningRatio field if non-nil, zero value otherwise.

### GetOverProvisioningRatioOk

`func (o *NodepoolCreateResponseFields) GetOverProvisioningRatioOk() (*int32, bool)`

GetOverProvisioningRatioOk returns a tuple with the OverProvisioningRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverProvisioningRatio

`func (o *NodepoolCreateResponseFields) SetOverProvisioningRatio(v int32)`

SetOverProvisioningRatio sets OverProvisioningRatio field to given value.

### HasOverProvisioningRatio

`func (o *NodepoolCreateResponseFields) HasOverProvisioningRatio() bool`

HasOverProvisioningRatio returns a boolean if a field has been set.

### SetOverProvisioningRatioNil

`func (o *NodepoolCreateResponseFields) SetOverProvisioningRatioNil(b bool)`

 SetOverProvisioningRatioNil sets the value for OverProvisioningRatio to be an explicit nil

### UnsetOverProvisioningRatio
`func (o *NodepoolCreateResponseFields) UnsetOverProvisioningRatio()`

UnsetOverProvisioningRatio ensures that no value is present for OverProvisioningRatio, not even an explicit nil
### GetPlacementStrategy

`func (o *NodepoolCreateResponseFields) GetPlacementStrategy() NodepoolCreateResponseFieldsPlacementStrategy`

GetPlacementStrategy returns the PlacementStrategy field if non-nil, zero value otherwise.

### GetPlacementStrategyOk

`func (o *NodepoolCreateResponseFields) GetPlacementStrategyOk() (*NodepoolCreateResponseFieldsPlacementStrategy, bool)`

GetPlacementStrategyOk returns a tuple with the PlacementStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementStrategy

`func (o *NodepoolCreateResponseFields) SetPlacementStrategy(v NodepoolCreateResponseFieldsPlacementStrategy)`

SetPlacementStrategy sets PlacementStrategy field to given value.

### HasPlacementStrategy

`func (o *NodepoolCreateResponseFields) HasPlacementStrategy() bool`

HasPlacementStrategy returns a boolean if a field has been set.

### GetGpuNetworkAccelerationLabelKey

`func (o *NodepoolCreateResponseFields) GetGpuNetworkAccelerationLabelKey() string`

GetGpuNetworkAccelerationLabelKey returns the GpuNetworkAccelerationLabelKey field if non-nil, zero value otherwise.

### GetGpuNetworkAccelerationLabelKeyOk

`func (o *NodepoolCreateResponseFields) GetGpuNetworkAccelerationLabelKeyOk() (*string, bool)`

GetGpuNetworkAccelerationLabelKeyOk returns a tuple with the GpuNetworkAccelerationLabelKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuNetworkAccelerationLabelKey

`func (o *NodepoolCreateResponseFields) SetGpuNetworkAccelerationLabelKey(v string)`

SetGpuNetworkAccelerationLabelKey sets GpuNetworkAccelerationLabelKey field to given value.

### HasGpuNetworkAccelerationLabelKey

`func (o *NodepoolCreateResponseFields) HasGpuNetworkAccelerationLabelKey() bool`

HasGpuNetworkAccelerationLabelKey returns a boolean if a field has been set.

### GetGpuNetworkAccelerationDetection

`func (o *NodepoolCreateResponseFields) GetGpuNetworkAccelerationDetection() NodepoolGPUNetworkAccelerationDetection`

GetGpuNetworkAccelerationDetection returns the GpuNetworkAccelerationDetection field if non-nil, zero value otherwise.

### GetGpuNetworkAccelerationDetectionOk

`func (o *NodepoolCreateResponseFields) GetGpuNetworkAccelerationDetectionOk() (*NodepoolGPUNetworkAccelerationDetection, bool)`

GetGpuNetworkAccelerationDetectionOk returns a tuple with the GpuNetworkAccelerationDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuNetworkAccelerationDetection

`func (o *NodepoolCreateResponseFields) SetGpuNetworkAccelerationDetection(v NodepoolGPUNetworkAccelerationDetection)`

SetGpuNetworkAccelerationDetection sets GpuNetworkAccelerationDetection field to given value.

### HasGpuNetworkAccelerationDetection

`func (o *NodepoolCreateResponseFields) HasGpuNetworkAccelerationDetection() bool`

HasGpuNetworkAccelerationDetection returns a boolean if a field has been set.

### SetGpuNetworkAccelerationDetectionNil

`func (o *NodepoolCreateResponseFields) SetGpuNetworkAccelerationDetectionNil(b bool)`

 SetGpuNetworkAccelerationDetectionNil sets the value for GpuNetworkAccelerationDetection to be an explicit nil

### UnsetGpuNetworkAccelerationDetection
`func (o *NodepoolCreateResponseFields) UnsetGpuNetworkAccelerationDetection()`

UnsetGpuNetworkAccelerationDetection ensures that no value is present for GpuNetworkAccelerationDetection, not even an explicit nil
### GetGpuResourceOptimization

`func (o *NodepoolCreateResponseFields) GetGpuResourceOptimization() GPUResourceOptimization`

GetGpuResourceOptimization returns the GpuResourceOptimization field if non-nil, zero value otherwise.

### GetGpuResourceOptimizationOk

`func (o *NodepoolCreateResponseFields) GetGpuResourceOptimizationOk() (*GPUResourceOptimization, bool)`

GetGpuResourceOptimizationOk returns a tuple with the GpuResourceOptimization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpuResourceOptimization

`func (o *NodepoolCreateResponseFields) SetGpuResourceOptimization(v GPUResourceOptimization)`

SetGpuResourceOptimization sets GpuResourceOptimization field to given value.

### HasGpuResourceOptimization

`func (o *NodepoolCreateResponseFields) HasGpuResourceOptimization() bool`

HasGpuResourceOptimization returns a boolean if a field has been set.

### SetGpuResourceOptimizationNil

`func (o *NodepoolCreateResponseFields) SetGpuResourceOptimizationNil(b bool)`

 SetGpuResourceOptimizationNil sets the value for GpuResourceOptimization to be an explicit nil

### UnsetGpuResourceOptimization
`func (o *NodepoolCreateResponseFields) UnsetGpuResourceOptimization()`

UnsetGpuResourceOptimization ensures that no value is present for GpuResourceOptimization, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


