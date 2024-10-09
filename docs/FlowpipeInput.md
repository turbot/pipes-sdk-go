# FlowpipeInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) | User information for the user who created this. | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**DeletedAt** | Pointer to **string** | The time the item was deleted in ISO 8601 UTC. | [optional] 
**DeletedBy** | Pointer to [**User**](User.md) | User information for the user that performed the deletion. | [optional] 
**DeletedById** | **string** | The ID of the user that performed the deletion. | 
**Id** | **string** | The unique ID of the Flowpipe input. | 
**IdentityId** | Pointer to **string** | The unique ID of the identity that owns the Flowpipe input. | [optional] 
**Inputs** | Pointer to **map[string]interface{}** | The form data for the Flowpipe input. | [optional] 
**Message** | Pointer to **string** | The message for the Flowpipe message step. | [optional] 
**Notifier** | Pointer to [**Notifier**](Notifier.md) | The details of the notifier this flowpipe input uses. | [optional] 
**NotifierId** | Pointer to **string** | The name of the Flowpipe notifier to use for this input. | [optional] 
**Process** | Pointer to [**SpProcess**](SpProcess.md) | The details of the process associated with this flowpipe input. | [optional] 
**ProcessId** | Pointer to **string** | The unique ID of the Pipes process for the Flowpipe run that owns this input. | [optional] 
**RandomId** | Pointer to **string** | The Flowpipe execution ID for this Flowpipe input. | [optional] 
**State** | Pointer to [**FlowpipeInputState**](FlowpipeInputState.md) | The state of the Flowpipe input. | [optional] 
**StateReason** | Pointer to **string** | The reason for the state of the Flowpipe input. | [optional] 
**StepExecutionId** | Pointer to **string** | The Flowpipe step execution ID for this Flowpipe input. | [optional] 
**StepType** | Pointer to [**FlowpipeInputStepType**](FlowpipeInputStepType.md) | The type of step for the Flowpipe step input/message. | [optional] 
**TenantId** | **string** | The unique ID of the tenant that owns the Flowpipe input. | 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) | User information for the last user to update this. | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 
**WorkspaceId** | Pointer to **string** | The unique ID of the workspace that owns the Flowpipe input. | [optional] 

## Methods

### NewFlowpipeInput

`func NewFlowpipeInput(createdAt string, createdById string, deletedById string, id string, tenantId string, updatedById string, versionId int32, ) *FlowpipeInput`

NewFlowpipeInput instantiates a new FlowpipeInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFlowpipeInputWithDefaults

`func NewFlowpipeInputWithDefaults() *FlowpipeInput`

NewFlowpipeInputWithDefaults instantiates a new FlowpipeInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FlowpipeInput) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FlowpipeInput) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FlowpipeInput) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *FlowpipeInput) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *FlowpipeInput) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *FlowpipeInput) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *FlowpipeInput) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *FlowpipeInput) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *FlowpipeInput) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *FlowpipeInput) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetDeletedAt

`func (o *FlowpipeInput) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *FlowpipeInput) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *FlowpipeInput) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *FlowpipeInput) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *FlowpipeInput) GetDeletedBy() User`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *FlowpipeInput) GetDeletedByOk() (*User, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *FlowpipeInput) SetDeletedBy(v User)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *FlowpipeInput) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedById

`func (o *FlowpipeInput) GetDeletedById() string`

GetDeletedById returns the DeletedById field if non-nil, zero value otherwise.

### GetDeletedByIdOk

`func (o *FlowpipeInput) GetDeletedByIdOk() (*string, bool)`

GetDeletedByIdOk returns a tuple with the DeletedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedById

`func (o *FlowpipeInput) SetDeletedById(v string)`

SetDeletedById sets DeletedById field to given value.


### GetId

`func (o *FlowpipeInput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FlowpipeInput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FlowpipeInput) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityId

`func (o *FlowpipeInput) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *FlowpipeInput) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *FlowpipeInput) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *FlowpipeInput) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetInputs

`func (o *FlowpipeInput) GetInputs() map[string]interface{}`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *FlowpipeInput) GetInputsOk() (*map[string]interface{}, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *FlowpipeInput) SetInputs(v map[string]interface{})`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *FlowpipeInput) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetMessage

`func (o *FlowpipeInput) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FlowpipeInput) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FlowpipeInput) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FlowpipeInput) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNotifier

`func (o *FlowpipeInput) GetNotifier() Notifier`

GetNotifier returns the Notifier field if non-nil, zero value otherwise.

### GetNotifierOk

`func (o *FlowpipeInput) GetNotifierOk() (*Notifier, bool)`

GetNotifierOk returns a tuple with the Notifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifier

`func (o *FlowpipeInput) SetNotifier(v Notifier)`

SetNotifier sets Notifier field to given value.

### HasNotifier

`func (o *FlowpipeInput) HasNotifier() bool`

HasNotifier returns a boolean if a field has been set.

### GetNotifierId

`func (o *FlowpipeInput) GetNotifierId() string`

GetNotifierId returns the NotifierId field if non-nil, zero value otherwise.

### GetNotifierIdOk

`func (o *FlowpipeInput) GetNotifierIdOk() (*string, bool)`

GetNotifierIdOk returns a tuple with the NotifierId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierId

`func (o *FlowpipeInput) SetNotifierId(v string)`

SetNotifierId sets NotifierId field to given value.

### HasNotifierId

`func (o *FlowpipeInput) HasNotifierId() bool`

HasNotifierId returns a boolean if a field has been set.

### GetProcess

`func (o *FlowpipeInput) GetProcess() SpProcess`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *FlowpipeInput) GetProcessOk() (*SpProcess, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcess

`func (o *FlowpipeInput) SetProcess(v SpProcess)`

SetProcess sets Process field to given value.

### HasProcess

`func (o *FlowpipeInput) HasProcess() bool`

HasProcess returns a boolean if a field has been set.

### GetProcessId

`func (o *FlowpipeInput) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *FlowpipeInput) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *FlowpipeInput) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *FlowpipeInput) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### GetRandomId

`func (o *FlowpipeInput) GetRandomId() string`

GetRandomId returns the RandomId field if non-nil, zero value otherwise.

### GetRandomIdOk

`func (o *FlowpipeInput) GetRandomIdOk() (*string, bool)`

GetRandomIdOk returns a tuple with the RandomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRandomId

`func (o *FlowpipeInput) SetRandomId(v string)`

SetRandomId sets RandomId field to given value.

### HasRandomId

`func (o *FlowpipeInput) HasRandomId() bool`

HasRandomId returns a boolean if a field has been set.

### GetState

`func (o *FlowpipeInput) GetState() FlowpipeInputState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *FlowpipeInput) GetStateOk() (*FlowpipeInputState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *FlowpipeInput) SetState(v FlowpipeInputState)`

SetState sets State field to given value.

### HasState

`func (o *FlowpipeInput) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateReason

`func (o *FlowpipeInput) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *FlowpipeInput) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *FlowpipeInput) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *FlowpipeInput) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### GetStepExecutionId

`func (o *FlowpipeInput) GetStepExecutionId() string`

GetStepExecutionId returns the StepExecutionId field if non-nil, zero value otherwise.

### GetStepExecutionIdOk

`func (o *FlowpipeInput) GetStepExecutionIdOk() (*string, bool)`

GetStepExecutionIdOk returns a tuple with the StepExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepExecutionId

`func (o *FlowpipeInput) SetStepExecutionId(v string)`

SetStepExecutionId sets StepExecutionId field to given value.

### HasStepExecutionId

`func (o *FlowpipeInput) HasStepExecutionId() bool`

HasStepExecutionId returns a boolean if a field has been set.

### GetStepType

`func (o *FlowpipeInput) GetStepType() FlowpipeInputStepType`

GetStepType returns the StepType field if non-nil, zero value otherwise.

### GetStepTypeOk

`func (o *FlowpipeInput) GetStepTypeOk() (*FlowpipeInputStepType, bool)`

GetStepTypeOk returns a tuple with the StepType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepType

`func (o *FlowpipeInput) SetStepType(v FlowpipeInputStepType)`

SetStepType sets StepType field to given value.

### HasStepType

`func (o *FlowpipeInput) HasStepType() bool`

HasStepType returns a boolean if a field has been set.

### GetTenantId

`func (o *FlowpipeInput) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *FlowpipeInput) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *FlowpipeInput) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUpdatedAt

`func (o *FlowpipeInput) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FlowpipeInput) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FlowpipeInput) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FlowpipeInput) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *FlowpipeInput) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *FlowpipeInput) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *FlowpipeInput) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *FlowpipeInput) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *FlowpipeInput) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *FlowpipeInput) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *FlowpipeInput) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetVersionId

`func (o *FlowpipeInput) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *FlowpipeInput) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *FlowpipeInput) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.


### GetWorkspaceId

`func (o *FlowpipeInput) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *FlowpipeInput) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *FlowpipeInput) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *FlowpipeInput) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


