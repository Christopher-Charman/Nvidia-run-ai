# EmailAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** |  | 
**Filename** | **string** |  | 
**ContentType** | Pointer to **string** |  | [optional] 

## Methods

### NewEmailAttachment

`func NewEmailAttachment(content string, filename string, ) *EmailAttachment`

NewEmailAttachment instantiates a new EmailAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailAttachmentWithDefaults

`func NewEmailAttachmentWithDefaults() *EmailAttachment`

NewEmailAttachmentWithDefaults instantiates a new EmailAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *EmailAttachment) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *EmailAttachment) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *EmailAttachment) SetContent(v string)`

SetContent sets Content field to given value.


### GetFilename

`func (o *EmailAttachment) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *EmailAttachment) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *EmailAttachment) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetContentType

`func (o *EmailAttachment) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *EmailAttachment) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *EmailAttachment) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *EmailAttachment) HasContentType() bool`

HasContentType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


