# ExportEmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | **[]string** |  | 
**Cc** | Pointer to **[]string** |  | [optional] 
**Bcc** | Pointer to **[]string** |  | [optional] 
**FromDisplayName** | Pointer to **string** |  | [optional] 
**Template** | [**Template**](Template.md) |  | 
**Attachments** | Pointer to [**[]EmailAttachment**](EmailAttachment.md) |  | [optional] 

## Methods

### NewExportEmailRequest

`func NewExportEmailRequest(to []string, template Template, ) *ExportEmailRequest`

NewExportEmailRequest instantiates a new ExportEmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportEmailRequestWithDefaults

`func NewExportEmailRequestWithDefaults() *ExportEmailRequest`

NewExportEmailRequestWithDefaults instantiates a new ExportEmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *ExportEmailRequest) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ExportEmailRequest) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ExportEmailRequest) SetTo(v []string)`

SetTo sets To field to given value.


### GetCc

`func (o *ExportEmailRequest) GetCc() []string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *ExportEmailRequest) GetCcOk() (*[]string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *ExportEmailRequest) SetCc(v []string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *ExportEmailRequest) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetBcc

`func (o *ExportEmailRequest) GetBcc() []string`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *ExportEmailRequest) GetBccOk() (*[]string, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *ExportEmailRequest) SetBcc(v []string)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *ExportEmailRequest) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetFromDisplayName

`func (o *ExportEmailRequest) GetFromDisplayName() string`

GetFromDisplayName returns the FromDisplayName field if non-nil, zero value otherwise.

### GetFromDisplayNameOk

`func (o *ExportEmailRequest) GetFromDisplayNameOk() (*string, bool)`

GetFromDisplayNameOk returns a tuple with the FromDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDisplayName

`func (o *ExportEmailRequest) SetFromDisplayName(v string)`

SetFromDisplayName sets FromDisplayName field to given value.

### HasFromDisplayName

`func (o *ExportEmailRequest) HasFromDisplayName() bool`

HasFromDisplayName returns a boolean if a field has been set.

### GetTemplate

`func (o *ExportEmailRequest) GetTemplate() Template`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ExportEmailRequest) GetTemplateOk() (*Template, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ExportEmailRequest) SetTemplate(v Template)`

SetTemplate sets Template field to given value.


### GetAttachments

`func (o *ExportEmailRequest) GetAttachments() []EmailAttachment`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *ExportEmailRequest) GetAttachmentsOk() (*[]EmailAttachment, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *ExportEmailRequest) SetAttachments(v []EmailAttachment)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *ExportEmailRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


