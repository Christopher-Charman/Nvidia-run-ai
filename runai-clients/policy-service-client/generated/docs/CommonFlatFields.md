# CommonFlatFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to **NullableString** | A command to the server as the entry point of the container running the workspace. | [optional] 
**Args** | Pointer to **NullableString** | Arguments to the command that the container running the workspace executes. | [optional] 
**Image** | Pointer to **NullableString** | Docker image name. For more information, see [Images](https://kubernetes.io/docs/concepts/containers/images). The image name is mandatory for creating a workspace. | [optional] 
**ImagePullPolicy** | Pointer to [**NullableImagePullPolicy**](ImagePullPolicy.md) |  | [optional] 
**WorkingDir** | Pointer to **NullableString** | Container&#39;s working directory. If not specified, the container runtime default will be used. This may be configured in the container image. | [optional] 
**CreateHomeDir** | Pointer to **NullableBool** | When set to &#x60;true&#x60;, creates a home directory for the container. | [optional] 
**Probes** | Pointer to [**NullableProbes**](Probes.md) |  | [optional] 
**NodeType** | Pointer to **NullableString** | Nodes (machines), or a group of nodes on which the workload will run. To use this feature, your Administrator will need to label nodes. For more information, see [Group Nodes](https://docs.run.ai/latest/admin/researcher-setup/limit-to-node-group). When using this flag with with Project-based affinity, it refines the list of allowable node groups set in the Project. For more information, see [Projects](https://docshub.run.ai/guides/platform-management/aiinitiatives/organization/projects). | [optional] 
**NodeAffinityRequired** | Pointer to [**NullableNodeAffinityRequired**](NodeAffinityRequired.md) |  | [optional] 
**PodAffinity** | Pointer to [**NullablePodAffinity**](PodAffinity.md) |  | [optional] 
**Category** | Pointer to **NullableString** | Specify the workload category assigned to the workload. Categories are used to classify and monitor different types of workloads within the NVIDIA Run:ai platform. | [optional] 
**PriorityClass** | Pointer to **NullableString** | Priority class of the workload. The default value for workspace is &#39;build&#39; and it can be changed to &#39;interactive-preemptible&#39; to allow the workload to use over-quota resources. The default value for training is &#39;train&#39; and it can be changed to &#39;build&#39; to allow the training workload to have a higher priority for in-queue scheduling and also become non-preemptive (if it&#39;s in deserved quota).. | [optional] 

## Methods

### NewCommonFlatFields

`func NewCommonFlatFields() *CommonFlatFields`

NewCommonFlatFields instantiates a new CommonFlatFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommonFlatFieldsWithDefaults

`func NewCommonFlatFieldsWithDefaults() *CommonFlatFields`

NewCommonFlatFieldsWithDefaults instantiates a new CommonFlatFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *CommonFlatFields) GetCommand() string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *CommonFlatFields) GetCommandOk() (*string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *CommonFlatFields) SetCommand(v string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *CommonFlatFields) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### SetCommandNil

`func (o *CommonFlatFields) SetCommandNil(b bool)`

 SetCommandNil sets the value for Command to be an explicit nil

### UnsetCommand
`func (o *CommonFlatFields) UnsetCommand()`

UnsetCommand ensures that no value is present for Command, not even an explicit nil
### GetArgs

`func (o *CommonFlatFields) GetArgs() string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *CommonFlatFields) GetArgsOk() (*string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *CommonFlatFields) SetArgs(v string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *CommonFlatFields) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### SetArgsNil

`func (o *CommonFlatFields) SetArgsNil(b bool)`

 SetArgsNil sets the value for Args to be an explicit nil

### UnsetArgs
`func (o *CommonFlatFields) UnsetArgs()`

UnsetArgs ensures that no value is present for Args, not even an explicit nil
### GetImage

`func (o *CommonFlatFields) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *CommonFlatFields) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *CommonFlatFields) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *CommonFlatFields) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *CommonFlatFields) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *CommonFlatFields) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil
### GetImagePullPolicy

`func (o *CommonFlatFields) GetImagePullPolicy() ImagePullPolicy`

GetImagePullPolicy returns the ImagePullPolicy field if non-nil, zero value otherwise.

### GetImagePullPolicyOk

`func (o *CommonFlatFields) GetImagePullPolicyOk() (*ImagePullPolicy, bool)`

GetImagePullPolicyOk returns a tuple with the ImagePullPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePullPolicy

`func (o *CommonFlatFields) SetImagePullPolicy(v ImagePullPolicy)`

SetImagePullPolicy sets ImagePullPolicy field to given value.

### HasImagePullPolicy

`func (o *CommonFlatFields) HasImagePullPolicy() bool`

HasImagePullPolicy returns a boolean if a field has been set.

### SetImagePullPolicyNil

`func (o *CommonFlatFields) SetImagePullPolicyNil(b bool)`

 SetImagePullPolicyNil sets the value for ImagePullPolicy to be an explicit nil

### UnsetImagePullPolicy
`func (o *CommonFlatFields) UnsetImagePullPolicy()`

UnsetImagePullPolicy ensures that no value is present for ImagePullPolicy, not even an explicit nil
### GetWorkingDir

`func (o *CommonFlatFields) GetWorkingDir() string`

GetWorkingDir returns the WorkingDir field if non-nil, zero value otherwise.

### GetWorkingDirOk

`func (o *CommonFlatFields) GetWorkingDirOk() (*string, bool)`

GetWorkingDirOk returns a tuple with the WorkingDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDir

`func (o *CommonFlatFields) SetWorkingDir(v string)`

SetWorkingDir sets WorkingDir field to given value.

### HasWorkingDir

`func (o *CommonFlatFields) HasWorkingDir() bool`

HasWorkingDir returns a boolean if a field has been set.

### SetWorkingDirNil

`func (o *CommonFlatFields) SetWorkingDirNil(b bool)`

 SetWorkingDirNil sets the value for WorkingDir to be an explicit nil

### UnsetWorkingDir
`func (o *CommonFlatFields) UnsetWorkingDir()`

UnsetWorkingDir ensures that no value is present for WorkingDir, not even an explicit nil
### GetCreateHomeDir

`func (o *CommonFlatFields) GetCreateHomeDir() bool`

GetCreateHomeDir returns the CreateHomeDir field if non-nil, zero value otherwise.

### GetCreateHomeDirOk

`func (o *CommonFlatFields) GetCreateHomeDirOk() (*bool, bool)`

GetCreateHomeDirOk returns a tuple with the CreateHomeDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateHomeDir

`func (o *CommonFlatFields) SetCreateHomeDir(v bool)`

SetCreateHomeDir sets CreateHomeDir field to given value.

### HasCreateHomeDir

`func (o *CommonFlatFields) HasCreateHomeDir() bool`

HasCreateHomeDir returns a boolean if a field has been set.

### SetCreateHomeDirNil

`func (o *CommonFlatFields) SetCreateHomeDirNil(b bool)`

 SetCreateHomeDirNil sets the value for CreateHomeDir to be an explicit nil

### UnsetCreateHomeDir
`func (o *CommonFlatFields) UnsetCreateHomeDir()`

UnsetCreateHomeDir ensures that no value is present for CreateHomeDir, not even an explicit nil
### GetProbes

`func (o *CommonFlatFields) GetProbes() Probes`

GetProbes returns the Probes field if non-nil, zero value otherwise.

### GetProbesOk

`func (o *CommonFlatFields) GetProbesOk() (*Probes, bool)`

GetProbesOk returns a tuple with the Probes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProbes

`func (o *CommonFlatFields) SetProbes(v Probes)`

SetProbes sets Probes field to given value.

### HasProbes

`func (o *CommonFlatFields) HasProbes() bool`

HasProbes returns a boolean if a field has been set.

### SetProbesNil

`func (o *CommonFlatFields) SetProbesNil(b bool)`

 SetProbesNil sets the value for Probes to be an explicit nil

### UnsetProbes
`func (o *CommonFlatFields) UnsetProbes()`

UnsetProbes ensures that no value is present for Probes, not even an explicit nil
### GetNodeType

`func (o *CommonFlatFields) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *CommonFlatFields) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *CommonFlatFields) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *CommonFlatFields) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### SetNodeTypeNil

`func (o *CommonFlatFields) SetNodeTypeNil(b bool)`

 SetNodeTypeNil sets the value for NodeType to be an explicit nil

### UnsetNodeType
`func (o *CommonFlatFields) UnsetNodeType()`

UnsetNodeType ensures that no value is present for NodeType, not even an explicit nil
### GetNodeAffinityRequired

`func (o *CommonFlatFields) GetNodeAffinityRequired() NodeAffinityRequired`

GetNodeAffinityRequired returns the NodeAffinityRequired field if non-nil, zero value otherwise.

### GetNodeAffinityRequiredOk

`func (o *CommonFlatFields) GetNodeAffinityRequiredOk() (*NodeAffinityRequired, bool)`

GetNodeAffinityRequiredOk returns a tuple with the NodeAffinityRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAffinityRequired

`func (o *CommonFlatFields) SetNodeAffinityRequired(v NodeAffinityRequired)`

SetNodeAffinityRequired sets NodeAffinityRequired field to given value.

### HasNodeAffinityRequired

`func (o *CommonFlatFields) HasNodeAffinityRequired() bool`

HasNodeAffinityRequired returns a boolean if a field has been set.

### SetNodeAffinityRequiredNil

`func (o *CommonFlatFields) SetNodeAffinityRequiredNil(b bool)`

 SetNodeAffinityRequiredNil sets the value for NodeAffinityRequired to be an explicit nil

### UnsetNodeAffinityRequired
`func (o *CommonFlatFields) UnsetNodeAffinityRequired()`

UnsetNodeAffinityRequired ensures that no value is present for NodeAffinityRequired, not even an explicit nil
### GetPodAffinity

`func (o *CommonFlatFields) GetPodAffinity() PodAffinity`

GetPodAffinity returns the PodAffinity field if non-nil, zero value otherwise.

### GetPodAffinityOk

`func (o *CommonFlatFields) GetPodAffinityOk() (*PodAffinity, bool)`

GetPodAffinityOk returns a tuple with the PodAffinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodAffinity

`func (o *CommonFlatFields) SetPodAffinity(v PodAffinity)`

SetPodAffinity sets PodAffinity field to given value.

### HasPodAffinity

`func (o *CommonFlatFields) HasPodAffinity() bool`

HasPodAffinity returns a boolean if a field has been set.

### SetPodAffinityNil

`func (o *CommonFlatFields) SetPodAffinityNil(b bool)`

 SetPodAffinityNil sets the value for PodAffinity to be an explicit nil

### UnsetPodAffinity
`func (o *CommonFlatFields) UnsetPodAffinity()`

UnsetPodAffinity ensures that no value is present for PodAffinity, not even an explicit nil
### GetCategory

`func (o *CommonFlatFields) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CommonFlatFields) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CommonFlatFields) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CommonFlatFields) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *CommonFlatFields) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *CommonFlatFields) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetPriorityClass

`func (o *CommonFlatFields) GetPriorityClass() string`

GetPriorityClass returns the PriorityClass field if non-nil, zero value otherwise.

### GetPriorityClassOk

`func (o *CommonFlatFields) GetPriorityClassOk() (*string, bool)`

GetPriorityClassOk returns a tuple with the PriorityClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityClass

`func (o *CommonFlatFields) SetPriorityClass(v string)`

SetPriorityClass sets PriorityClass field to given value.

### HasPriorityClass

`func (o *CommonFlatFields) HasPriorityClass() bool`

HasPriorityClass returns a boolean if a field has been set.

### SetPriorityClassNil

`func (o *CommonFlatFields) SetPriorityClassNil(b bool)`

 SetPriorityClassNil sets the value for PriorityClass to be an explicit nil

### UnsetPriorityClass
`func (o *CommonFlatFields) UnsetPriorityClass()`

UnsetPriorityClass ensures that no value is present for PriorityClass, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


