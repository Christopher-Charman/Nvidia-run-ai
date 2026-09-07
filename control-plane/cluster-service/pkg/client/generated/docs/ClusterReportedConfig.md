# ClusterReportedConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkloadOwnershipProtection** | Pointer to **NullableBool** | indicates if the cluster workload ownership protection is enabled | [optional] 
**SubdomainEnabled** | Pointer to **NullableBool** | indicates if the cluster subdomain is enabled | [optional] 
**IngressClass** | Pointer to **NullableString** | the ingress class that is configured for the cluster | [optional] 
**GatewayName** | Pointer to **NullableString** | the name of the Gateway API gateway resource configured for the cluster | [optional] 
**GatewayNamespace** | Pointer to **NullableString** | the namespace of the Gateway API gateway resource configured for the cluster | [optional] 

## Methods

### NewClusterReportedConfig

`func NewClusterReportedConfig() *ClusterReportedConfig`

NewClusterReportedConfig instantiates a new ClusterReportedConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterReportedConfigWithDefaults

`func NewClusterReportedConfigWithDefaults() *ClusterReportedConfig`

NewClusterReportedConfigWithDefaults instantiates a new ClusterReportedConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkloadOwnershipProtection

`func (o *ClusterReportedConfig) GetWorkloadOwnershipProtection() bool`

GetWorkloadOwnershipProtection returns the WorkloadOwnershipProtection field if non-nil, zero value otherwise.

### GetWorkloadOwnershipProtectionOk

`func (o *ClusterReportedConfig) GetWorkloadOwnershipProtectionOk() (*bool, bool)`

GetWorkloadOwnershipProtectionOk returns a tuple with the WorkloadOwnershipProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadOwnershipProtection

`func (o *ClusterReportedConfig) SetWorkloadOwnershipProtection(v bool)`

SetWorkloadOwnershipProtection sets WorkloadOwnershipProtection field to given value.

### HasWorkloadOwnershipProtection

`func (o *ClusterReportedConfig) HasWorkloadOwnershipProtection() bool`

HasWorkloadOwnershipProtection returns a boolean if a field has been set.

### SetWorkloadOwnershipProtectionNil

`func (o *ClusterReportedConfig) SetWorkloadOwnershipProtectionNil(b bool)`

 SetWorkloadOwnershipProtectionNil sets the value for WorkloadOwnershipProtection to be an explicit nil

### UnsetWorkloadOwnershipProtection
`func (o *ClusterReportedConfig) UnsetWorkloadOwnershipProtection()`

UnsetWorkloadOwnershipProtection ensures that no value is present for WorkloadOwnershipProtection, not even an explicit nil
### GetSubdomainEnabled

`func (o *ClusterReportedConfig) GetSubdomainEnabled() bool`

GetSubdomainEnabled returns the SubdomainEnabled field if non-nil, zero value otherwise.

### GetSubdomainEnabledOk

`func (o *ClusterReportedConfig) GetSubdomainEnabledOk() (*bool, bool)`

GetSubdomainEnabledOk returns a tuple with the SubdomainEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomainEnabled

`func (o *ClusterReportedConfig) SetSubdomainEnabled(v bool)`

SetSubdomainEnabled sets SubdomainEnabled field to given value.

### HasSubdomainEnabled

`func (o *ClusterReportedConfig) HasSubdomainEnabled() bool`

HasSubdomainEnabled returns a boolean if a field has been set.

### SetSubdomainEnabledNil

`func (o *ClusterReportedConfig) SetSubdomainEnabledNil(b bool)`

 SetSubdomainEnabledNil sets the value for SubdomainEnabled to be an explicit nil

### UnsetSubdomainEnabled
`func (o *ClusterReportedConfig) UnsetSubdomainEnabled()`

UnsetSubdomainEnabled ensures that no value is present for SubdomainEnabled, not even an explicit nil
### GetIngressClass

`func (o *ClusterReportedConfig) GetIngressClass() string`

GetIngressClass returns the IngressClass field if non-nil, zero value otherwise.

### GetIngressClassOk

`func (o *ClusterReportedConfig) GetIngressClassOk() (*string, bool)`

GetIngressClassOk returns a tuple with the IngressClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressClass

`func (o *ClusterReportedConfig) SetIngressClass(v string)`

SetIngressClass sets IngressClass field to given value.

### HasIngressClass

`func (o *ClusterReportedConfig) HasIngressClass() bool`

HasIngressClass returns a boolean if a field has been set.

### SetIngressClassNil

`func (o *ClusterReportedConfig) SetIngressClassNil(b bool)`

 SetIngressClassNil sets the value for IngressClass to be an explicit nil

### UnsetIngressClass
`func (o *ClusterReportedConfig) UnsetIngressClass()`

UnsetIngressClass ensures that no value is present for IngressClass, not even an explicit nil
### GetGatewayName

`func (o *ClusterReportedConfig) GetGatewayName() string`

GetGatewayName returns the GatewayName field if non-nil, zero value otherwise.

### GetGatewayNameOk

`func (o *ClusterReportedConfig) GetGatewayNameOk() (*string, bool)`

GetGatewayNameOk returns a tuple with the GatewayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayName

`func (o *ClusterReportedConfig) SetGatewayName(v string)`

SetGatewayName sets GatewayName field to given value.

### HasGatewayName

`func (o *ClusterReportedConfig) HasGatewayName() bool`

HasGatewayName returns a boolean if a field has been set.

### SetGatewayNameNil

`func (o *ClusterReportedConfig) SetGatewayNameNil(b bool)`

 SetGatewayNameNil sets the value for GatewayName to be an explicit nil

### UnsetGatewayName
`func (o *ClusterReportedConfig) UnsetGatewayName()`

UnsetGatewayName ensures that no value is present for GatewayName, not even an explicit nil
### GetGatewayNamespace

`func (o *ClusterReportedConfig) GetGatewayNamespace() string`

GetGatewayNamespace returns the GatewayNamespace field if non-nil, zero value otherwise.

### GetGatewayNamespaceOk

`func (o *ClusterReportedConfig) GetGatewayNamespaceOk() (*string, bool)`

GetGatewayNamespaceOk returns a tuple with the GatewayNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayNamespace

`func (o *ClusterReportedConfig) SetGatewayNamespace(v string)`

SetGatewayNamespace sets GatewayNamespace field to given value.

### HasGatewayNamespace

`func (o *ClusterReportedConfig) HasGatewayNamespace() bool`

HasGatewayNamespace returns a boolean if a field has been set.

### SetGatewayNamespaceNil

`func (o *ClusterReportedConfig) SetGatewayNamespaceNil(b bool)`

 SetGatewayNamespaceNil sets the value for GatewayNamespace to be an explicit nil

### UnsetGatewayNamespace
`func (o *ClusterReportedConfig) UnsetGatewayNamespace()`

UnsetGatewayNamespace ensures that no value is present for GatewayNamespace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


