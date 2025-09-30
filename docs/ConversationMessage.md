# ConversationMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactId** | Pointer to **string** |  | [optional] 
**ArtifactType** | Pointer to **string** |  | [optional] 
**CanvasId** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **map[string]interface{}** |  | [optional] 
**ConversationId** | Pointer to **string** |  | [optional] 
**CreatedAt** | **string** |  | 
**CreatedById** | **string** |  | 
**Id** | Pointer to **string** |  | [optional] 
**MessageType** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Param** | Pointer to **map[string]interface{}** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**ThinkingId** | Pointer to **string** |  | [optional] 
**ToolCallId** | Pointer to **string** |  | [optional] 

## Methods

### NewConversationMessage

`func NewConversationMessage(createdAt string, createdById string, ) *ConversationMessage`

NewConversationMessage instantiates a new ConversationMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationMessageWithDefaults

`func NewConversationMessageWithDefaults() *ConversationMessage`

NewConversationMessageWithDefaults instantiates a new ConversationMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactId

`func (o *ConversationMessage) GetArtifactId() string`

GetArtifactId returns the ArtifactId field if non-nil, zero value otherwise.

### GetArtifactIdOk

`func (o *ConversationMessage) GetArtifactIdOk() (*string, bool)`

GetArtifactIdOk returns a tuple with the ArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactId

`func (o *ConversationMessage) SetArtifactId(v string)`

SetArtifactId sets ArtifactId field to given value.

### HasArtifactId

`func (o *ConversationMessage) HasArtifactId() bool`

HasArtifactId returns a boolean if a field has been set.

### GetArtifactType

`func (o *ConversationMessage) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *ConversationMessage) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *ConversationMessage) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.

### HasArtifactType

`func (o *ConversationMessage) HasArtifactType() bool`

HasArtifactType returns a boolean if a field has been set.

### GetCanvasId

`func (o *ConversationMessage) GetCanvasId() string`

GetCanvasId returns the CanvasId field if non-nil, zero value otherwise.

### GetCanvasIdOk

`func (o *ConversationMessage) GetCanvasIdOk() (*string, bool)`

GetCanvasIdOk returns a tuple with the CanvasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanvasId

`func (o *ConversationMessage) SetCanvasId(v string)`

SetCanvasId sets CanvasId field to given value.

### HasCanvasId

`func (o *ConversationMessage) HasCanvasId() bool`

HasCanvasId returns a boolean if a field has been set.

### GetContent

`func (o *ConversationMessage) GetContent() map[string]interface{}`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ConversationMessage) GetContentOk() (*map[string]interface{}, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ConversationMessage) SetContent(v map[string]interface{})`

SetContent sets Content field to given value.

### HasContent

`func (o *ConversationMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetConversationId

`func (o *ConversationMessage) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *ConversationMessage) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *ConversationMessage) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *ConversationMessage) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ConversationMessage) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConversationMessage) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConversationMessage) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedById

`func (o *ConversationMessage) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *ConversationMessage) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *ConversationMessage) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetId

`func (o *ConversationMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConversationMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConversationMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConversationMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessageType

`func (o *ConversationMessage) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *ConversationMessage) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *ConversationMessage) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.

### HasMessageType

`func (o *ConversationMessage) HasMessageType() bool`

HasMessageType returns a boolean if a field has been set.

### GetMetadata

`func (o *ConversationMessage) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ConversationMessage) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ConversationMessage) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ConversationMessage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetParam

`func (o *ConversationMessage) GetParam() map[string]interface{}`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *ConversationMessage) GetParamOk() (*map[string]interface{}, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *ConversationMessage) SetParam(v map[string]interface{})`

SetParam sets Param field to given value.

### HasParam

`func (o *ConversationMessage) HasParam() bool`

HasParam returns a boolean if a field has been set.

### GetRole

`func (o *ConversationMessage) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ConversationMessage) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ConversationMessage) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ConversationMessage) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetThinkingId

`func (o *ConversationMessage) GetThinkingId() string`

GetThinkingId returns the ThinkingId field if non-nil, zero value otherwise.

### GetThinkingIdOk

`func (o *ConversationMessage) GetThinkingIdOk() (*string, bool)`

GetThinkingIdOk returns a tuple with the ThinkingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinkingId

`func (o *ConversationMessage) SetThinkingId(v string)`

SetThinkingId sets ThinkingId field to given value.

### HasThinkingId

`func (o *ConversationMessage) HasThinkingId() bool`

HasThinkingId returns a boolean if a field has been set.

### GetToolCallId

`func (o *ConversationMessage) GetToolCallId() string`

GetToolCallId returns the ToolCallId field if non-nil, zero value otherwise.

### GetToolCallIdOk

`func (o *ConversationMessage) GetToolCallIdOk() (*string, bool)`

GetToolCallIdOk returns a tuple with the ToolCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCallId

`func (o *ConversationMessage) SetToolCallId(v string)`

SetToolCallId sets ToolCallId field to given value.

### HasToolCallId

`func (o *ConversationMessage) HasToolCallId() bool`

HasToolCallId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


