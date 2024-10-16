# WorkspaceMod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**Constraint** | Pointer to **string** |  | [optional] 
**ConstraintMode** | Pointer to **string** |  | [optional] 
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) | User information for the user who created this. | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**Id** | **string** |  | 
**IdentityId** | **string** | The unique identifier for an identity where the workspace mod has been install. | 
**InstalledCommit** | Pointer to **string** |  | [optional] 
**InstalledVersion** | Pointer to **string** |  | [optional] 
**IntegrationId** | Pointer to **string** |  | [optional] 
**ModId** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Pipeling** | Pointer to [**ModPipeling**](ModPipeling.md) |  | [optional] 
**State** | Pointer to [**WorkspaceModState**](WorkspaceModState.md) |  | [optional] 
**StateReason** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) | User information for the last user to update this. | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**Version** | Pointer to **string** |  | [optional] 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 
**WorkspaceId** | **string** |  | 

## Methods

### NewWorkspaceMod

`func NewWorkspaceMod(createdAt string, createdById string, id string, identityId string, updatedById string, versionId int32, workspaceId string, ) *WorkspaceMod`

NewWorkspaceMod instantiates a new WorkspaceMod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceModWithDefaults

`func NewWorkspaceModWithDefaults() *WorkspaceMod`

NewWorkspaceModWithDefaults instantiates a new WorkspaceMod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *WorkspaceMod) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *WorkspaceMod) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *WorkspaceMod) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *WorkspaceMod) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetBranch

`func (o *WorkspaceMod) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *WorkspaceMod) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *WorkspaceMod) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *WorkspaceMod) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetConstraint

`func (o *WorkspaceMod) GetConstraint() string`

GetConstraint returns the Constraint field if non-nil, zero value otherwise.

### GetConstraintOk

`func (o *WorkspaceMod) GetConstraintOk() (*string, bool)`

GetConstraintOk returns a tuple with the Constraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraint

`func (o *WorkspaceMod) SetConstraint(v string)`

SetConstraint sets Constraint field to given value.

### HasConstraint

`func (o *WorkspaceMod) HasConstraint() bool`

HasConstraint returns a boolean if a field has been set.

### GetConstraintMode

`func (o *WorkspaceMod) GetConstraintMode() string`

GetConstraintMode returns the ConstraintMode field if non-nil, zero value otherwise.

### GetConstraintModeOk

`func (o *WorkspaceMod) GetConstraintModeOk() (*string, bool)`

GetConstraintModeOk returns a tuple with the ConstraintMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintMode

`func (o *WorkspaceMod) SetConstraintMode(v string)`

SetConstraintMode sets ConstraintMode field to given value.

### HasConstraintMode

`func (o *WorkspaceMod) HasConstraintMode() bool`

HasConstraintMode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *WorkspaceMod) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkspaceMod) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkspaceMod) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *WorkspaceMod) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WorkspaceMod) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WorkspaceMod) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WorkspaceMod) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *WorkspaceMod) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *WorkspaceMod) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *WorkspaceMod) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetId

`func (o *WorkspaceMod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkspaceMod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkspaceMod) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityId

`func (o *WorkspaceMod) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *WorkspaceMod) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *WorkspaceMod) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetInstalledCommit

`func (o *WorkspaceMod) GetInstalledCommit() string`

GetInstalledCommit returns the InstalledCommit field if non-nil, zero value otherwise.

### GetInstalledCommitOk

`func (o *WorkspaceMod) GetInstalledCommitOk() (*string, bool)`

GetInstalledCommitOk returns a tuple with the InstalledCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledCommit

`func (o *WorkspaceMod) SetInstalledCommit(v string)`

SetInstalledCommit sets InstalledCommit field to given value.

### HasInstalledCommit

`func (o *WorkspaceMod) HasInstalledCommit() bool`

HasInstalledCommit returns a boolean if a field has been set.

### GetInstalledVersion

`func (o *WorkspaceMod) GetInstalledVersion() string`

GetInstalledVersion returns the InstalledVersion field if non-nil, zero value otherwise.

### GetInstalledVersionOk

`func (o *WorkspaceMod) GetInstalledVersionOk() (*string, bool)`

GetInstalledVersionOk returns a tuple with the InstalledVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstalledVersion

`func (o *WorkspaceMod) SetInstalledVersion(v string)`

SetInstalledVersion sets InstalledVersion field to given value.

### HasInstalledVersion

`func (o *WorkspaceMod) HasInstalledVersion() bool`

HasInstalledVersion returns a boolean if a field has been set.

### GetIntegrationId

`func (o *WorkspaceMod) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *WorkspaceMod) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *WorkspaceMod) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *WorkspaceMod) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetModId

`func (o *WorkspaceMod) GetModId() string`

GetModId returns the ModId field if non-nil, zero value otherwise.

### GetModIdOk

`func (o *WorkspaceMod) GetModIdOk() (*string, bool)`

GetModIdOk returns a tuple with the ModId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModId

`func (o *WorkspaceMod) SetModId(v string)`

SetModId sets ModId field to given value.

### HasModId

`func (o *WorkspaceMod) HasModId() bool`

HasModId returns a boolean if a field has been set.

### GetPath

`func (o *WorkspaceMod) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *WorkspaceMod) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *WorkspaceMod) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *WorkspaceMod) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPipeling

`func (o *WorkspaceMod) GetPipeling() ModPipeling`

GetPipeling returns the Pipeling field if non-nil, zero value otherwise.

### GetPipelingOk

`func (o *WorkspaceMod) GetPipelingOk() (*ModPipeling, bool)`

GetPipelingOk returns a tuple with the Pipeling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeling

`func (o *WorkspaceMod) SetPipeling(v ModPipeling)`

SetPipeling sets Pipeling field to given value.

### HasPipeling

`func (o *WorkspaceMod) HasPipeling() bool`

HasPipeling returns a boolean if a field has been set.

### GetState

`func (o *WorkspaceMod) GetState() WorkspaceModState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *WorkspaceMod) GetStateOk() (*WorkspaceModState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *WorkspaceMod) SetState(v WorkspaceModState)`

SetState sets State field to given value.

### HasState

`func (o *WorkspaceMod) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateReason

`func (o *WorkspaceMod) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *WorkspaceMod) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *WorkspaceMod) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *WorkspaceMod) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WorkspaceMod) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkspaceMod) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkspaceMod) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WorkspaceMod) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *WorkspaceMod) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *WorkspaceMod) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *WorkspaceMod) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *WorkspaceMod) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *WorkspaceMod) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *WorkspaceMod) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *WorkspaceMod) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetVersion

`func (o *WorkspaceMod) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkspaceMod) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkspaceMod) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *WorkspaceMod) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionId

`func (o *WorkspaceMod) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *WorkspaceMod) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *WorkspaceMod) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.


### GetWorkspaceId

`func (o *WorkspaceMod) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WorkspaceMod) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WorkspaceMod) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


