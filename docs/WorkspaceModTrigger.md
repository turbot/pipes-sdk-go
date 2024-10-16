# WorkspaceModTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **map[string]interface{}** |  | [optional] 
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) | User information for the user who created this. | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**DefaultSchedule** | Pointer to **string** |  | [optional] 
**DefaultState** | Pointer to [**TriggerState**](TriggerState.md) |  | [optional] 
**DeletedAt** | Pointer to **string** | The time the item was deleted in ISO 8601 UTC. | [optional] 
**DeletedBy** | Pointer to [**User**](User.md) | User information for the user that performed the deletion. | [optional] 
**DeletedById** | **string** | The ID of the user that performed the deletion. | 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Identity** | Pointer to [**Identity**](Identity.md) |  | [optional] 
**IdentityId** | Pointer to **string** |  | [optional] 
**LastProcess** | Pointer to [**SpProcess**](SpProcess.md) | Information for the last process of the pipeline. | [optional] 
**LastProcessId** | Pointer to **string** | The id of the last process that was run for the pipeline. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NextRunAt** | Pointer to **string** | The time when the trigger is next scheduled to run in ISO 8601 UTC. | [optional] 
**Pipelines** | Pointer to **map[string]interface{}** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to [**PipelineFrequency**](PipelineFrequency.md) |  | [optional] 
**State** | Pointer to [**TriggerState**](TriggerState.md) |  | [optional] 
**StateReason** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) | User information for the last user to update this. | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 
**Workspace** | Pointer to [**Workspace**](Workspace.md) |  | [optional] 
**WorkspaceId** | Pointer to **string** |  | [optional] 
**WorkspaceMod** | Pointer to [**WorkspaceMod**](WorkspaceMod.md) |  | [optional] 
**WorkspaceModId** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkspaceModTrigger

`func NewWorkspaceModTrigger(createdAt string, createdById string, deletedById string, updatedById string, versionId int32, ) *WorkspaceModTrigger`

NewWorkspaceModTrigger instantiates a new WorkspaceModTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceModTriggerWithDefaults

`func NewWorkspaceModTriggerWithDefaults() *WorkspaceModTrigger`

NewWorkspaceModTriggerWithDefaults instantiates a new WorkspaceModTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *WorkspaceModTrigger) GetArgs() map[string]interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *WorkspaceModTrigger) GetArgsOk() (*map[string]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *WorkspaceModTrigger) SetArgs(v map[string]interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *WorkspaceModTrigger) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WorkspaceModTrigger) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkspaceModTrigger) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkspaceModTrigger) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *WorkspaceModTrigger) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WorkspaceModTrigger) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WorkspaceModTrigger) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WorkspaceModTrigger) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *WorkspaceModTrigger) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *WorkspaceModTrigger) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *WorkspaceModTrigger) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetDefaultSchedule

`func (o *WorkspaceModTrigger) GetDefaultSchedule() string`

GetDefaultSchedule returns the DefaultSchedule field if non-nil, zero value otherwise.

### GetDefaultScheduleOk

`func (o *WorkspaceModTrigger) GetDefaultScheduleOk() (*string, bool)`

GetDefaultScheduleOk returns a tuple with the DefaultSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSchedule

`func (o *WorkspaceModTrigger) SetDefaultSchedule(v string)`

SetDefaultSchedule sets DefaultSchedule field to given value.

### HasDefaultSchedule

`func (o *WorkspaceModTrigger) HasDefaultSchedule() bool`

HasDefaultSchedule returns a boolean if a field has been set.

### GetDefaultState

`func (o *WorkspaceModTrigger) GetDefaultState() TriggerState`

GetDefaultState returns the DefaultState field if non-nil, zero value otherwise.

### GetDefaultStateOk

`func (o *WorkspaceModTrigger) GetDefaultStateOk() (*TriggerState, bool)`

GetDefaultStateOk returns a tuple with the DefaultState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultState

`func (o *WorkspaceModTrigger) SetDefaultState(v TriggerState)`

SetDefaultState sets DefaultState field to given value.

### HasDefaultState

`func (o *WorkspaceModTrigger) HasDefaultState() bool`

HasDefaultState returns a boolean if a field has been set.

### GetDeletedAt

`func (o *WorkspaceModTrigger) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *WorkspaceModTrigger) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *WorkspaceModTrigger) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *WorkspaceModTrigger) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *WorkspaceModTrigger) GetDeletedBy() User`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *WorkspaceModTrigger) GetDeletedByOk() (*User, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *WorkspaceModTrigger) SetDeletedBy(v User)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *WorkspaceModTrigger) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedById

`func (o *WorkspaceModTrigger) GetDeletedById() string`

GetDeletedById returns the DeletedById field if non-nil, zero value otherwise.

### GetDeletedByIdOk

`func (o *WorkspaceModTrigger) GetDeletedByIdOk() (*string, bool)`

GetDeletedByIdOk returns a tuple with the DeletedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedById

`func (o *WorkspaceModTrigger) SetDeletedById(v string)`

SetDeletedById sets DeletedById field to given value.


### GetDescription

`func (o *WorkspaceModTrigger) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkspaceModTrigger) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkspaceModTrigger) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkspaceModTrigger) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *WorkspaceModTrigger) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkspaceModTrigger) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkspaceModTrigger) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkspaceModTrigger) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentity

`func (o *WorkspaceModTrigger) GetIdentity() Identity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *WorkspaceModTrigger) GetIdentityOk() (*Identity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *WorkspaceModTrigger) SetIdentity(v Identity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *WorkspaceModTrigger) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetIdentityId

`func (o *WorkspaceModTrigger) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *WorkspaceModTrigger) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *WorkspaceModTrigger) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *WorkspaceModTrigger) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetLastProcess

`func (o *WorkspaceModTrigger) GetLastProcess() SpProcess`

GetLastProcess returns the LastProcess field if non-nil, zero value otherwise.

### GetLastProcessOk

`func (o *WorkspaceModTrigger) GetLastProcessOk() (*SpProcess, bool)`

GetLastProcessOk returns a tuple with the LastProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProcess

`func (o *WorkspaceModTrigger) SetLastProcess(v SpProcess)`

SetLastProcess sets LastProcess field to given value.

### HasLastProcess

`func (o *WorkspaceModTrigger) HasLastProcess() bool`

HasLastProcess returns a boolean if a field has been set.

### GetLastProcessId

`func (o *WorkspaceModTrigger) GetLastProcessId() string`

GetLastProcessId returns the LastProcessId field if non-nil, zero value otherwise.

### GetLastProcessIdOk

`func (o *WorkspaceModTrigger) GetLastProcessIdOk() (*string, bool)`

GetLastProcessIdOk returns a tuple with the LastProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProcessId

`func (o *WorkspaceModTrigger) SetLastProcessId(v string)`

SetLastProcessId sets LastProcessId field to given value.

### HasLastProcessId

`func (o *WorkspaceModTrigger) HasLastProcessId() bool`

HasLastProcessId returns a boolean if a field has been set.

### GetName

`func (o *WorkspaceModTrigger) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceModTrigger) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceModTrigger) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkspaceModTrigger) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextRunAt

`func (o *WorkspaceModTrigger) GetNextRunAt() string`

GetNextRunAt returns the NextRunAt field if non-nil, zero value otherwise.

### GetNextRunAtOk

`func (o *WorkspaceModTrigger) GetNextRunAtOk() (*string, bool)`

GetNextRunAtOk returns a tuple with the NextRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunAt

`func (o *WorkspaceModTrigger) SetNextRunAt(v string)`

SetNextRunAt sets NextRunAt field to given value.

### HasNextRunAt

`func (o *WorkspaceModTrigger) HasNextRunAt() bool`

HasNextRunAt returns a boolean if a field has been set.

### GetPipelines

`func (o *WorkspaceModTrigger) GetPipelines() map[string]interface{}`

GetPipelines returns the Pipelines field if non-nil, zero value otherwise.

### GetPipelinesOk

`func (o *WorkspaceModTrigger) GetPipelinesOk() (*map[string]interface{}, bool)`

GetPipelinesOk returns a tuple with the Pipelines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelines

`func (o *WorkspaceModTrigger) SetPipelines(v map[string]interface{})`

SetPipelines sets Pipelines field to given value.

### HasPipelines

`func (o *WorkspaceModTrigger) HasPipelines() bool`

HasPipelines returns a boolean if a field has been set.

### GetQuery

`func (o *WorkspaceModTrigger) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *WorkspaceModTrigger) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *WorkspaceModTrigger) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *WorkspaceModTrigger) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSchedule

`func (o *WorkspaceModTrigger) GetSchedule() PipelineFrequency`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *WorkspaceModTrigger) GetScheduleOk() (*PipelineFrequency, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *WorkspaceModTrigger) SetSchedule(v PipelineFrequency)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *WorkspaceModTrigger) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetState

`func (o *WorkspaceModTrigger) GetState() TriggerState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WorkspaceModTrigger) GetStateOk() (*TriggerState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WorkspaceModTrigger) SetState(v TriggerState)`

SetState sets State field to given value.

### HasState

`func (o *WorkspaceModTrigger) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateReason

`func (o *WorkspaceModTrigger) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *WorkspaceModTrigger) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *WorkspaceModTrigger) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *WorkspaceModTrigger) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### GetTags

`func (o *WorkspaceModTrigger) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkspaceModTrigger) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkspaceModTrigger) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *WorkspaceModTrigger) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTitle

`func (o *WorkspaceModTrigger) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkspaceModTrigger) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkspaceModTrigger) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkspaceModTrigger) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *WorkspaceModTrigger) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkspaceModTrigger) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkspaceModTrigger) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkspaceModTrigger) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WorkspaceModTrigger) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkspaceModTrigger) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkspaceModTrigger) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WorkspaceModTrigger) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *WorkspaceModTrigger) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *WorkspaceModTrigger) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *WorkspaceModTrigger) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *WorkspaceModTrigger) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *WorkspaceModTrigger) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *WorkspaceModTrigger) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *WorkspaceModTrigger) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetVersionId

`func (o *WorkspaceModTrigger) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *WorkspaceModTrigger) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *WorkspaceModTrigger) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.


### GetWorkspace

`func (o *WorkspaceModTrigger) GetWorkspace() Workspace`

GetWorkspace returns the Workspace field if non-nil, zero value otherwise.

### GetWorkspaceOk

`func (o *WorkspaceModTrigger) GetWorkspaceOk() (*Workspace, bool)`

GetWorkspaceOk returns a tuple with the Workspace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspace

`func (o *WorkspaceModTrigger) SetWorkspace(v Workspace)`

SetWorkspace sets Workspace field to given value.

### HasWorkspace

`func (o *WorkspaceModTrigger) HasWorkspace() bool`

HasWorkspace returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *WorkspaceModTrigger) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WorkspaceModTrigger) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WorkspaceModTrigger) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *WorkspaceModTrigger) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### GetWorkspaceMod

`func (o *WorkspaceModTrigger) GetWorkspaceMod() WorkspaceMod`

GetWorkspaceMod returns the WorkspaceMod field if non-nil, zero value otherwise.

### GetWorkspaceModOk

`func (o *WorkspaceModTrigger) GetWorkspaceModOk() (*WorkspaceMod, bool)`

GetWorkspaceModOk returns a tuple with the WorkspaceMod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceMod

`func (o *WorkspaceModTrigger) SetWorkspaceMod(v WorkspaceMod)`

SetWorkspaceMod sets WorkspaceMod field to given value.

### HasWorkspaceMod

`func (o *WorkspaceModTrigger) HasWorkspaceMod() bool`

HasWorkspaceMod returns a boolean if a field has been set.

### GetWorkspaceModId

`func (o *WorkspaceModTrigger) GetWorkspaceModId() string`

GetWorkspaceModId returns the WorkspaceModId field if non-nil, zero value otherwise.

### GetWorkspaceModIdOk

`func (o *WorkspaceModTrigger) GetWorkspaceModIdOk() (*string, bool)`

GetWorkspaceModIdOk returns a tuple with the WorkspaceModId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceModId

`func (o *WorkspaceModTrigger) SetWorkspaceModId(v string)`

SetWorkspaceModId sets WorkspaceModId field to given value.

### HasWorkspaceModId

`func (o *WorkspaceModTrigger) HasWorkspaceModId() bool`

HasWorkspaceModId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


