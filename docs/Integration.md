# Integration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** | The configuration for the integration. | [optional] 
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) | User information for the user who created this. | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**DeletedAt** | Pointer to **string** | The time the item was deleted in ISO 8601 UTC. | [optional] 
**DeletedBy** | Pointer to [**User**](User.md) | User information for the user that performed the deletion. | [optional] 
**DeletedById** | **string** | The ID of the user that performed the deletion. | 
**GithubInstallationId** | **int32** | The GitHub installation ID for this integration. | 
**Handle** | **string** | The handle for the integration. | 
**Id** | **string** | The unique identifier for the integration. | 
**IdentityId** | Pointer to **string** | The unique identifier for an identity where this integration is created. | [optional] 
**Pipeline** | Pointer to [**Pipeline**](Pipeline.md) | Additional information about the pipeline the integration belongs to. | [optional] 
**PipelineId** | Pointer to **string** | The pipeline ID for this integration. | [optional] 
**State** | Pointer to [**IntegrationState**](IntegrationState.md) | The state of the integration. | [optional] 
**StateReason** | Pointer to **string** | The reason for the state of the integration. | [optional] 
**TenantId** | **string** | The unique identifier for the tenant where this integration is created. | 
**Type** | Pointer to [**IntegrationType**](IntegrationType.md) | The type of the integration. | [optional] 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) | User information for the last user to update this. | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 

## Methods

### NewIntegration

`func NewIntegration(createdAt string, createdById string, deletedById string, githubInstallationId int32, handle string, id string, tenantId string, updatedById string, versionId int32, ) *Integration`

NewIntegration instantiates a new Integration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationWithDefaults

`func NewIntegrationWithDefaults() *Integration`

NewIntegrationWithDefaults instantiates a new Integration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *Integration) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Integration) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Integration) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Integration) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Integration) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Integration) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Integration) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Integration) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Integration) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Integration) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Integration) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *Integration) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *Integration) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *Integration) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetDeletedAt

`func (o *Integration) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Integration) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Integration) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Integration) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *Integration) GetDeletedBy() User`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *Integration) GetDeletedByOk() (*User, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *Integration) SetDeletedBy(v User)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *Integration) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedById

`func (o *Integration) GetDeletedById() string`

GetDeletedById returns the DeletedById field if non-nil, zero value otherwise.

### GetDeletedByIdOk

`func (o *Integration) GetDeletedByIdOk() (*string, bool)`

GetDeletedByIdOk returns a tuple with the DeletedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedById

`func (o *Integration) SetDeletedById(v string)`

SetDeletedById sets DeletedById field to given value.


### GetGithubInstallationId

`func (o *Integration) GetGithubInstallationId() int32`

GetGithubInstallationId returns the GithubInstallationId field if non-nil, zero value otherwise.

### GetGithubInstallationIdOk

`func (o *Integration) GetGithubInstallationIdOk() (*int32, bool)`

GetGithubInstallationIdOk returns a tuple with the GithubInstallationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubInstallationId

`func (o *Integration) SetGithubInstallationId(v int32)`

SetGithubInstallationId sets GithubInstallationId field to given value.


### GetHandle

`func (o *Integration) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *Integration) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *Integration) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetId

`func (o *Integration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Integration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Integration) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityId

`func (o *Integration) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *Integration) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *Integration) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *Integration) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetPipeline

`func (o *Integration) GetPipeline() Pipeline`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *Integration) GetPipelineOk() (*Pipeline, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *Integration) SetPipeline(v Pipeline)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *Integration) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### GetPipelineId

`func (o *Integration) GetPipelineId() string`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *Integration) GetPipelineIdOk() (*string, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *Integration) SetPipelineId(v string)`

SetPipelineId sets PipelineId field to given value.

### HasPipelineId

`func (o *Integration) HasPipelineId() bool`

HasPipelineId returns a boolean if a field has been set.

### GetState

`func (o *Integration) GetState() IntegrationState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Integration) GetStateOk() (*IntegrationState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Integration) SetState(v IntegrationState)`

SetState sets State field to given value.

### HasState

`func (o *Integration) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateReason

`func (o *Integration) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *Integration) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *Integration) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *Integration) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### GetTenantId

`func (o *Integration) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Integration) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Integration) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetType

`func (o *Integration) GetType() IntegrationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Integration) GetTypeOk() (*IntegrationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Integration) SetType(v IntegrationType)`

SetType sets Type field to given value.

### HasType

`func (o *Integration) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Integration) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Integration) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Integration) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Integration) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Integration) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Integration) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Integration) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Integration) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *Integration) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *Integration) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *Integration) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetVersionId

`func (o *Integration) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *Integration) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *Integration) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


