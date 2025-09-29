# ConversationArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanvasId** | Pointer to **string** |  | [optional] 
**ConversationId** | **string** |  | 
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) | User information for the user who created this. | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**Id** | **string** |  | 
**Kind** | **string** |  | 
**Language** | Pointer to **string** |  | [optional] 
**Sql** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) | User information for the last user to update this. | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**UsageStorageBytes** | Pointer to **int32** |  | [optional] 

## Methods

### NewConversationArtifact

`func NewConversationArtifact(conversationId string, createdAt string, createdById string, id string, kind string, updatedById string, ) *ConversationArtifact`

NewConversationArtifact instantiates a new ConversationArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationArtifactWithDefaults

`func NewConversationArtifactWithDefaults() *ConversationArtifact`

NewConversationArtifactWithDefaults instantiates a new ConversationArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCanvasId

`func (o *ConversationArtifact) GetCanvasId() string`

GetCanvasId returns the CanvasId field if non-nil, zero value otherwise.

### GetCanvasIdOk

`func (o *ConversationArtifact) GetCanvasIdOk() (*string, bool)`

GetCanvasIdOk returns a tuple with the CanvasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanvasId

`func (o *ConversationArtifact) SetCanvasId(v string)`

SetCanvasId sets CanvasId field to given value.

### HasCanvasId

`func (o *ConversationArtifact) HasCanvasId() bool`

HasCanvasId returns a boolean if a field has been set.

### GetConversationId

`func (o *ConversationArtifact) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *ConversationArtifact) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *ConversationArtifact) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.


### GetCreatedAt

`func (o *ConversationArtifact) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConversationArtifact) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConversationArtifact) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *ConversationArtifact) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ConversationArtifact) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ConversationArtifact) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ConversationArtifact) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *ConversationArtifact) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *ConversationArtifact) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *ConversationArtifact) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetId

`func (o *ConversationArtifact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConversationArtifact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConversationArtifact) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *ConversationArtifact) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ConversationArtifact) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ConversationArtifact) SetKind(v string)`

SetKind sets Kind field to given value.


### GetLanguage

`func (o *ConversationArtifact) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ConversationArtifact) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ConversationArtifact) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ConversationArtifact) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetSql

`func (o *ConversationArtifact) GetSql() string`

GetSql returns the Sql field if non-nil, zero value otherwise.

### GetSqlOk

`func (o *ConversationArtifact) GetSqlOk() (*string, bool)`

GetSqlOk returns a tuple with the Sql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSql

`func (o *ConversationArtifact) SetSql(v string)`

SetSql sets Sql field to given value.

### HasSql

`func (o *ConversationArtifact) HasSql() bool`

HasSql returns a boolean if a field has been set.

### GetTitle

`func (o *ConversationArtifact) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ConversationArtifact) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ConversationArtifact) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ConversationArtifact) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ConversationArtifact) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ConversationArtifact) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ConversationArtifact) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ConversationArtifact) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ConversationArtifact) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ConversationArtifact) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ConversationArtifact) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ConversationArtifact) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *ConversationArtifact) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *ConversationArtifact) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *ConversationArtifact) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetUsageStorageBytes

`func (o *ConversationArtifact) GetUsageStorageBytes() int32`

GetUsageStorageBytes returns the UsageStorageBytes field if non-nil, zero value otherwise.

### GetUsageStorageBytesOk

`func (o *ConversationArtifact) GetUsageStorageBytesOk() (*int32, bool)`

GetUsageStorageBytesOk returns a tuple with the UsageStorageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageStorageBytes

`func (o *ConversationArtifact) SetUsageStorageBytes(v int32)`

SetUsageStorageBytes sets UsageStorageBytes field to given value.

### HasUsageStorageBytes

`func (o *ConversationArtifact) HasUsageStorageBytes() bool`

HasUsageStorageBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


