# SpProcess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | Pointer to **string** | The unique identifier of the connection for which the process is created. | [optional] 
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) | User information for the user who created this. | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**FlowpipePipelineId** | Pointer to **string** | The unique identifier of the flowpipe pipeline for which the process is created. | [optional] 
**Id** | **string** | The unique identifier of the process. | 
**IdentityId** | Pointer to **string** | The unique identifier of the identity for which the process is created. | [optional] 
**Pipe** | **string** | The pipe for which the process has been created. | 
**Pipeline** | Pointer to [**Pipeline**](Pipeline.md) | The current details of the pipeline for which the process is created. | [optional] 
**PipelineId** | Pointer to **string** | The unique identifier of the pipeline for which the process is created. | [optional] 
**State** | Pointer to [**ProcessState**](ProcessState.md) | The state of the process. | [optional] 
**StateReason** | Pointer to **string** | The optional reason why the process is in its current state. | [optional] 
**TailpipePartitionId** | Pointer to **string** | The unique identifier of the tailpipe partition for which the process is created. | [optional] 
**TenantId** | Pointer to **string** | The unique identifier of the tenant for which the process is created. | [optional] 
**Trigger** | Pointer to [**WorkspaceModTrigger**](WorkspaceModTrigger.md) | The current details of the trigger for which the process is created. | [optional] 
**TriggerId** | Pointer to **string** | The unique identifier of the trigger for which the process is created. | [optional] 
**Type** | **string** | The type of the process, generally denotes the activity performed e.g. workspace.create, pipeline.execute, pipeline.command.run. | 
**UpdatedAt** | **string** | The time of the last update in ISO 8601 UTC. | 
**UpdatedBy** | Pointer to [**User**](User.md) | User information for the last user to update this. | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**Usage** | Pointer to **map[string]interface{}** | The usage information for this process. | [optional] 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 
**WorkspaceId** | Pointer to **string** | The unique identifier of the workspace for which the process is created. | [optional] 

## Methods

### NewSpProcess

`func NewSpProcess(createdAt string, createdById string, id string, pipe string, type_ string, updatedAt string, updatedById string, versionId int32, ) *SpProcess`

NewSpProcess instantiates a new SpProcess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpProcessWithDefaults

`func NewSpProcessWithDefaults() *SpProcess`

NewSpProcessWithDefaults instantiates a new SpProcess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *SpProcess) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *SpProcess) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *SpProcess) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *SpProcess) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SpProcess) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SpProcess) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SpProcess) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *SpProcess) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SpProcess) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SpProcess) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SpProcess) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *SpProcess) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *SpProcess) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *SpProcess) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetFlowpipePipelineId

`func (o *SpProcess) GetFlowpipePipelineId() string`

GetFlowpipePipelineId returns the FlowpipePipelineId field if non-nil, zero value otherwise.

### GetFlowpipePipelineIdOk

`func (o *SpProcess) GetFlowpipePipelineIdOk() (*string, bool)`

GetFlowpipePipelineIdOk returns a tuple with the FlowpipePipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowpipePipelineId

`func (o *SpProcess) SetFlowpipePipelineId(v string)`

SetFlowpipePipelineId sets FlowpipePipelineId field to given value.

### HasFlowpipePipelineId

`func (o *SpProcess) HasFlowpipePipelineId() bool`

HasFlowpipePipelineId returns a boolean if a field has been set.

### GetId

`func (o *SpProcess) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpProcess) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpProcess) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityId

`func (o *SpProcess) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *SpProcess) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *SpProcess) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *SpProcess) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetPipe

`func (o *SpProcess) GetPipe() string`

GetPipe returns the Pipe field if non-nil, zero value otherwise.

### GetPipeOk

`func (o *SpProcess) GetPipeOk() (*string, bool)`

GetPipeOk returns a tuple with the Pipe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipe

`func (o *SpProcess) SetPipe(v string)`

SetPipe sets Pipe field to given value.


### GetPipeline

`func (o *SpProcess) GetPipeline() Pipeline`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *SpProcess) GetPipelineOk() (*Pipeline, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *SpProcess) SetPipeline(v Pipeline)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *SpProcess) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### GetPipelineId

`func (o *SpProcess) GetPipelineId() string`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *SpProcess) GetPipelineIdOk() (*string, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *SpProcess) SetPipelineId(v string)`

SetPipelineId sets PipelineId field to given value.

### HasPipelineId

`func (o *SpProcess) HasPipelineId() bool`

HasPipelineId returns a boolean if a field has been set.

### GetState

`func (o *SpProcess) GetState() ProcessState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *SpProcess) GetStateOk() (*ProcessState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *SpProcess) SetState(v ProcessState)`

SetState sets State field to given value.

### HasState

`func (o *SpProcess) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateReason

`func (o *SpProcess) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *SpProcess) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *SpProcess) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *SpProcess) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### GetTailpipePartitionId

`func (o *SpProcess) GetTailpipePartitionId() string`

GetTailpipePartitionId returns the TailpipePartitionId field if non-nil, zero value otherwise.

### GetTailpipePartitionIdOk

`func (o *SpProcess) GetTailpipePartitionIdOk() (*string, bool)`

GetTailpipePartitionIdOk returns a tuple with the TailpipePartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTailpipePartitionId

`func (o *SpProcess) SetTailpipePartitionId(v string)`

SetTailpipePartitionId sets TailpipePartitionId field to given value.

### HasTailpipePartitionId

`func (o *SpProcess) HasTailpipePartitionId() bool`

HasTailpipePartitionId returns a boolean if a field has been set.

### GetTenantId

`func (o *SpProcess) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SpProcess) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SpProcess) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SpProcess) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTrigger

`func (o *SpProcess) GetTrigger() WorkspaceModTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *SpProcess) GetTriggerOk() (*WorkspaceModTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *SpProcess) SetTrigger(v WorkspaceModTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *SpProcess) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetTriggerId

`func (o *SpProcess) GetTriggerId() string`

GetTriggerId returns the TriggerId field if non-nil, zero value otherwise.

### GetTriggerIdOk

`func (o *SpProcess) GetTriggerIdOk() (*string, bool)`

GetTriggerIdOk returns a tuple with the TriggerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerId

`func (o *SpProcess) SetTriggerId(v string)`

SetTriggerId sets TriggerId field to given value.

### HasTriggerId

`func (o *SpProcess) HasTriggerId() bool`

HasTriggerId returns a boolean if a field has been set.

### GetType

`func (o *SpProcess) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SpProcess) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SpProcess) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *SpProcess) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SpProcess) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SpProcess) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUpdatedBy

`func (o *SpProcess) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *SpProcess) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *SpProcess) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *SpProcess) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *SpProcess) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *SpProcess) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *SpProcess) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetUsage

`func (o *SpProcess) GetUsage() map[string]interface{}`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *SpProcess) GetUsageOk() (*map[string]interface{}, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *SpProcess) SetUsage(v map[string]interface{})`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *SpProcess) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetVersionId

`func (o *SpProcess) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *SpProcess) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *SpProcess) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.


### GetWorkspaceId

`func (o *SpProcess) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *SpProcess) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *SpProcess) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *SpProcess) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


