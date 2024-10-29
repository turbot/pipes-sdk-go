# ModTriggerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) | User information for the user who created this. | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ModId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Params** | Pointer to **map[string]interface{}** |  | [optional] 
**Pipelines** | Pointer to **map[string]interface{}** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**ModTriggerState**](ModTriggerState.md) |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) | User information for the last user to update this. | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 

## Methods

### NewModTriggerInfo

`func NewModTriggerInfo(createdAt string, createdById string, updatedById string, versionId int32, ) *ModTriggerInfo`

NewModTriggerInfo instantiates a new ModTriggerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModTriggerInfoWithDefaults

`func NewModTriggerInfoWithDefaults() *ModTriggerInfo`

NewModTriggerInfoWithDefaults instantiates a new ModTriggerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ModTriggerInfo) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ModTriggerInfo) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ModTriggerInfo) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *ModTriggerInfo) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ModTriggerInfo) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ModTriggerInfo) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ModTriggerInfo) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *ModTriggerInfo) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *ModTriggerInfo) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *ModTriggerInfo) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetDescription

`func (o *ModTriggerInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModTriggerInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModTriggerInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModTriggerInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ModTriggerInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModTriggerInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModTriggerInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModTriggerInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModId

`func (o *ModTriggerInfo) GetModId() string`

GetModId returns the ModId field if non-nil, zero value otherwise.

### GetModIdOk

`func (o *ModTriggerInfo) GetModIdOk() (*string, bool)`

GetModIdOk returns a tuple with the ModId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModId

`func (o *ModTriggerInfo) SetModId(v string)`

SetModId sets ModId field to given value.

### HasModId

`func (o *ModTriggerInfo) HasModId() bool`

HasModId returns a boolean if a field has been set.

### GetName

`func (o *ModTriggerInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModTriggerInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModTriggerInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModTriggerInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParams

`func (o *ModTriggerInfo) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ModTriggerInfo) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ModTriggerInfo) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *ModTriggerInfo) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetPipelines

`func (o *ModTriggerInfo) GetPipelines() map[string]interface{}`

GetPipelines returns the Pipelines field if non-nil, zero value otherwise.

### GetPipelinesOk

`func (o *ModTriggerInfo) GetPipelinesOk() (*map[string]interface{}, bool)`

GetPipelinesOk returns a tuple with the Pipelines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelines

`func (o *ModTriggerInfo) SetPipelines(v map[string]interface{})`

SetPipelines sets Pipelines field to given value.

### HasPipelines

`func (o *ModTriggerInfo) HasPipelines() bool`

HasPipelines returns a boolean if a field has been set.

### GetQuery

`func (o *ModTriggerInfo) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ModTriggerInfo) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ModTriggerInfo) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ModTriggerInfo) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSchedule

`func (o *ModTriggerInfo) GetSchedule() string`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ModTriggerInfo) GetScheduleOk() (*string, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ModTriggerInfo) SetSchedule(v string)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ModTriggerInfo) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetState

`func (o *ModTriggerInfo) GetState() ModTriggerState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ModTriggerInfo) GetStateOk() (*ModTriggerState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ModTriggerInfo) SetState(v ModTriggerState)`

SetState sets State field to given value.

### HasState

`func (o *ModTriggerInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *ModTriggerInfo) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModTriggerInfo) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModTriggerInfo) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ModTriggerInfo) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTitle

`func (o *ModTriggerInfo) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ModTriggerInfo) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ModTriggerInfo) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ModTriggerInfo) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *ModTriggerInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModTriggerInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModTriggerInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModTriggerInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ModTriggerInfo) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ModTriggerInfo) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ModTriggerInfo) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ModTriggerInfo) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ModTriggerInfo) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ModTriggerInfo) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ModTriggerInfo) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ModTriggerInfo) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *ModTriggerInfo) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *ModTriggerInfo) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *ModTriggerInfo) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetVersionId

`func (o *ModTriggerInfo) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *ModTriggerInfo) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *ModTriggerInfo) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


