# WorkspaceModPipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) | User information for the user who created this. | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**Description** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**LastProcess** | Pointer to [**SpProcess**](SpProcess.md) | Information for the last process of the pipeline. | [optional] 
**LastProcessId** | Pointer to **string** | The id of the last process that was run for the pipeline. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Params** | Pointer to **map[string]interface{}** |  | [optional] 
**Steps** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Triggers** | Pointer to **map[string]interface{}** | The list of triggers which the pipeline is associated with. | [optional] 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) | User information for the last user to update this. | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 
**WorkspaceMod** | Pointer to [**WorkspaceMod**](WorkspaceMod.md) |  | [optional] 
**WorkspaceModId** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkspaceModPipeline

`func NewWorkspaceModPipeline(createdAt string, createdById string, updatedById string, versionId int32, ) *WorkspaceModPipeline`

NewWorkspaceModPipeline instantiates a new WorkspaceModPipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceModPipelineWithDefaults

`func NewWorkspaceModPipelineWithDefaults() *WorkspaceModPipeline`

NewWorkspaceModPipelineWithDefaults instantiates a new WorkspaceModPipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *WorkspaceModPipeline) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkspaceModPipeline) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkspaceModPipeline) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *WorkspaceModPipeline) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WorkspaceModPipeline) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WorkspaceModPipeline) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WorkspaceModPipeline) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *WorkspaceModPipeline) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *WorkspaceModPipeline) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *WorkspaceModPipeline) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetDescription

`func (o *WorkspaceModPipeline) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkspaceModPipeline) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkspaceModPipeline) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkspaceModPipeline) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *WorkspaceModPipeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkspaceModPipeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkspaceModPipeline) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkspaceModPipeline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastProcess

`func (o *WorkspaceModPipeline) GetLastProcess() SpProcess`

GetLastProcess returns the LastProcess field if non-nil, zero value otherwise.

### GetLastProcessOk

`func (o *WorkspaceModPipeline) GetLastProcessOk() (*SpProcess, bool)`

GetLastProcessOk returns a tuple with the LastProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProcess

`func (o *WorkspaceModPipeline) SetLastProcess(v SpProcess)`

SetLastProcess sets LastProcess field to given value.

### HasLastProcess

`func (o *WorkspaceModPipeline) HasLastProcess() bool`

HasLastProcess returns a boolean if a field has been set.

### GetLastProcessId

`func (o *WorkspaceModPipeline) GetLastProcessId() string`

GetLastProcessId returns the LastProcessId field if non-nil, zero value otherwise.

### GetLastProcessIdOk

`func (o *WorkspaceModPipeline) GetLastProcessIdOk() (*string, bool)`

GetLastProcessIdOk returns a tuple with the LastProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProcessId

`func (o *WorkspaceModPipeline) SetLastProcessId(v string)`

SetLastProcessId sets LastProcessId field to given value.

### HasLastProcessId

`func (o *WorkspaceModPipeline) HasLastProcessId() bool`

HasLastProcessId returns a boolean if a field has been set.

### GetName

`func (o *WorkspaceModPipeline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceModPipeline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceModPipeline) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkspaceModPipeline) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParams

`func (o *WorkspaceModPipeline) GetParams() map[string]interface{}`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *WorkspaceModPipeline) GetParamsOk() (*map[string]interface{}, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *WorkspaceModPipeline) SetParams(v map[string]interface{})`

SetParams sets Params field to given value.

### HasParams

`func (o *WorkspaceModPipeline) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetSteps

`func (o *WorkspaceModPipeline) GetSteps() map[string]interface{}`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *WorkspaceModPipeline) GetStepsOk() (*map[string]interface{}, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *WorkspaceModPipeline) SetSteps(v map[string]interface{})`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *WorkspaceModPipeline) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetTags

`func (o *WorkspaceModPipeline) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *WorkspaceModPipeline) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *WorkspaceModPipeline) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *WorkspaceModPipeline) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTitle

`func (o *WorkspaceModPipeline) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkspaceModPipeline) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkspaceModPipeline) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkspaceModPipeline) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTriggers

`func (o *WorkspaceModPipeline) GetTriggers() map[string]interface{}`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *WorkspaceModPipeline) GetTriggersOk() (*map[string]interface{}, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *WorkspaceModPipeline) SetTriggers(v map[string]interface{})`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *WorkspaceModPipeline) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WorkspaceModPipeline) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkspaceModPipeline) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkspaceModPipeline) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WorkspaceModPipeline) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *WorkspaceModPipeline) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *WorkspaceModPipeline) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *WorkspaceModPipeline) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *WorkspaceModPipeline) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *WorkspaceModPipeline) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *WorkspaceModPipeline) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *WorkspaceModPipeline) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetVersionId

`func (o *WorkspaceModPipeline) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *WorkspaceModPipeline) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *WorkspaceModPipeline) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.


### GetWorkspaceMod

`func (o *WorkspaceModPipeline) GetWorkspaceMod() WorkspaceMod`

GetWorkspaceMod returns the WorkspaceMod field if non-nil, zero value otherwise.

### GetWorkspaceModOk

`func (o *WorkspaceModPipeline) GetWorkspaceModOk() (*WorkspaceMod, bool)`

GetWorkspaceModOk returns a tuple with the WorkspaceMod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceMod

`func (o *WorkspaceModPipeline) SetWorkspaceMod(v WorkspaceMod)`

SetWorkspaceMod sets WorkspaceMod field to given value.

### HasWorkspaceMod

`func (o *WorkspaceModPipeline) HasWorkspaceMod() bool`

HasWorkspaceMod returns a boolean if a field has been set.

### GetWorkspaceModId

`func (o *WorkspaceModPipeline) GetWorkspaceModId() string`

GetWorkspaceModId returns the WorkspaceModId field if non-nil, zero value otherwise.

### GetWorkspaceModIdOk

`func (o *WorkspaceModPipeline) GetWorkspaceModIdOk() (*string, bool)`

GetWorkspaceModIdOk returns a tuple with the WorkspaceModId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceModId

`func (o *WorkspaceModPipeline) SetWorkspaceModId(v string)`

SetWorkspaceModId sets WorkspaceModId field to given value.

### HasWorkspaceModId

`func (o *WorkspaceModPipeline) HasWorkspaceModId() bool`

HasWorkspaceModId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


