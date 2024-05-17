# WorkspaceConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Association** | Pointer to [**WorkspaceConnectionAssociation**](WorkspaceConnectionAssociation.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**ConnectionLevel** | **string** | The level at which the connection exists, can be wither &#39;identity&#39; or &#39;workspace&#39;. | 
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**DeletedAt** | Pointer to **string** | The time the item was deleted in ISO 8601 UTC. | [optional] 
**DeletedBy** | Pointer to [**User**](User.md) |  | [optional] 
**DeletedById** | **string** | The ID of the user that performed the deletion. | 
**Handle** | **string** | The handle name of the  connection. | 
**HandleDynamic** | Pointer to **string** | The dynamically-generated handle for the connection. Only populated if this is a discovered connection. | [optional] 
**HandleMode** | Pointer to **string** | The handle mode for the connection. | [optional] 
**Id** | **string** | The unique identifier for the connection. | 
**IdentityId** | Pointer to **string** | The unique identifier for an identity where the connection has been created. | [optional] 
**IntegrationResourceId** | Pointer to **string** | The integration resource ID for the connection. | [optional] 
**ManagedById** | Pointer to **string** | The ID of the aggregator that manages this connection. Only populated if this is a discovered connection. | [optional] 
**Plugin** | Pointer to **string** | The plugin name for the connection. | [optional] 
**PluginVersion** | Pointer to **string** | The plugin version for the connection. | [optional] 
**SourceIdentifier** | Pointer to **string** | The source identifier for this connection. Only populated if this is a discovered connection. | [optional] 
**TenantId** | **string** | The unique identifier for the tenant where the connection has been created. | 
**Type** | Pointer to **string** | Type of connection i.e aggregator or connection. | [optional] 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 
**WorkspaceId** | **string** | The unique identifier for the workspace. | 

## Methods

### NewWorkspaceConnection

`func NewWorkspaceConnection(connectionLevel string, createdAt string, createdById string, deletedById string, handle string, id string, tenantId string, updatedById string, versionId int32, workspaceId string, ) *WorkspaceConnection`

NewWorkspaceConnection instantiates a new WorkspaceConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceConnectionWithDefaults

`func NewWorkspaceConnectionWithDefaults() *WorkspaceConnection`

NewWorkspaceConnectionWithDefaults instantiates a new WorkspaceConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociation

`func (o *WorkspaceConnection) GetAssociation() WorkspaceConnectionAssociation`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *WorkspaceConnection) GetAssociationOk() (*WorkspaceConnectionAssociation, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *WorkspaceConnection) SetAssociation(v WorkspaceConnectionAssociation)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *WorkspaceConnection) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetConfig

`func (o *WorkspaceConnection) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *WorkspaceConnection) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *WorkspaceConnection) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *WorkspaceConnection) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConnectionLevel

`func (o *WorkspaceConnection) GetConnectionLevel() string`

GetConnectionLevel returns the ConnectionLevel field if non-nil, zero value otherwise.

### GetConnectionLevelOk

`func (o *WorkspaceConnection) GetConnectionLevelOk() (*string, bool)`

GetConnectionLevelOk returns a tuple with the ConnectionLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionLevel

`func (o *WorkspaceConnection) SetConnectionLevel(v string)`

SetConnectionLevel sets ConnectionLevel field to given value.


### GetCreatedAt

`func (o *WorkspaceConnection) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkspaceConnection) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkspaceConnection) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *WorkspaceConnection) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WorkspaceConnection) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WorkspaceConnection) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WorkspaceConnection) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *WorkspaceConnection) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *WorkspaceConnection) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *WorkspaceConnection) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetDeletedAt

`func (o *WorkspaceConnection) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *WorkspaceConnection) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *WorkspaceConnection) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *WorkspaceConnection) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *WorkspaceConnection) GetDeletedBy() User`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *WorkspaceConnection) GetDeletedByOk() (*User, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *WorkspaceConnection) SetDeletedBy(v User)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *WorkspaceConnection) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedById

`func (o *WorkspaceConnection) GetDeletedById() string`

GetDeletedById returns the DeletedById field if non-nil, zero value otherwise.

### GetDeletedByIdOk

`func (o *WorkspaceConnection) GetDeletedByIdOk() (*string, bool)`

GetDeletedByIdOk returns a tuple with the DeletedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedById

`func (o *WorkspaceConnection) SetDeletedById(v string)`

SetDeletedById sets DeletedById field to given value.


### GetHandle

`func (o *WorkspaceConnection) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *WorkspaceConnection) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *WorkspaceConnection) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetHandleDynamic

`func (o *WorkspaceConnection) GetHandleDynamic() string`

GetHandleDynamic returns the HandleDynamic field if non-nil, zero value otherwise.

### GetHandleDynamicOk

`func (o *WorkspaceConnection) GetHandleDynamicOk() (*string, bool)`

GetHandleDynamicOk returns a tuple with the HandleDynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandleDynamic

`func (o *WorkspaceConnection) SetHandleDynamic(v string)`

SetHandleDynamic sets HandleDynamic field to given value.

### HasHandleDynamic

`func (o *WorkspaceConnection) HasHandleDynamic() bool`

HasHandleDynamic returns a boolean if a field has been set.

### GetHandleMode

`func (o *WorkspaceConnection) GetHandleMode() string`

GetHandleMode returns the HandleMode field if non-nil, zero value otherwise.

### GetHandleModeOk

`func (o *WorkspaceConnection) GetHandleModeOk() (*string, bool)`

GetHandleModeOk returns a tuple with the HandleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandleMode

`func (o *WorkspaceConnection) SetHandleMode(v string)`

SetHandleMode sets HandleMode field to given value.

### HasHandleMode

`func (o *WorkspaceConnection) HasHandleMode() bool`

HasHandleMode returns a boolean if a field has been set.

### GetId

`func (o *WorkspaceConnection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkspaceConnection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkspaceConnection) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityId

`func (o *WorkspaceConnection) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *WorkspaceConnection) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *WorkspaceConnection) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *WorkspaceConnection) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetIntegrationResourceId

`func (o *WorkspaceConnection) GetIntegrationResourceId() string`

GetIntegrationResourceId returns the IntegrationResourceId field if non-nil, zero value otherwise.

### GetIntegrationResourceIdOk

`func (o *WorkspaceConnection) GetIntegrationResourceIdOk() (*string, bool)`

GetIntegrationResourceIdOk returns a tuple with the IntegrationResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationResourceId

`func (o *WorkspaceConnection) SetIntegrationResourceId(v string)`

SetIntegrationResourceId sets IntegrationResourceId field to given value.

### HasIntegrationResourceId

`func (o *WorkspaceConnection) HasIntegrationResourceId() bool`

HasIntegrationResourceId returns a boolean if a field has been set.

### GetManagedById

`func (o *WorkspaceConnection) GetManagedById() string`

GetManagedById returns the ManagedById field if non-nil, zero value otherwise.

### GetManagedByIdOk

`func (o *WorkspaceConnection) GetManagedByIdOk() (*string, bool)`

GetManagedByIdOk returns a tuple with the ManagedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedById

`func (o *WorkspaceConnection) SetManagedById(v string)`

SetManagedById sets ManagedById field to given value.

### HasManagedById

`func (o *WorkspaceConnection) HasManagedById() bool`

HasManagedById returns a boolean if a field has been set.

### GetPlugin

`func (o *WorkspaceConnection) GetPlugin() string`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *WorkspaceConnection) GetPluginOk() (*string, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *WorkspaceConnection) SetPlugin(v string)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *WorkspaceConnection) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.

### GetPluginVersion

`func (o *WorkspaceConnection) GetPluginVersion() string`

GetPluginVersion returns the PluginVersion field if non-nil, zero value otherwise.

### GetPluginVersionOk

`func (o *WorkspaceConnection) GetPluginVersionOk() (*string, bool)`

GetPluginVersionOk returns a tuple with the PluginVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginVersion

`func (o *WorkspaceConnection) SetPluginVersion(v string)`

SetPluginVersion sets PluginVersion field to given value.

### HasPluginVersion

`func (o *WorkspaceConnection) HasPluginVersion() bool`

HasPluginVersion returns a boolean if a field has been set.

### GetSourceIdentifier

`func (o *WorkspaceConnection) GetSourceIdentifier() string`

GetSourceIdentifier returns the SourceIdentifier field if non-nil, zero value otherwise.

### GetSourceIdentifierOk

`func (o *WorkspaceConnection) GetSourceIdentifierOk() (*string, bool)`

GetSourceIdentifierOk returns a tuple with the SourceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIdentifier

`func (o *WorkspaceConnection) SetSourceIdentifier(v string)`

SetSourceIdentifier sets SourceIdentifier field to given value.

### HasSourceIdentifier

`func (o *WorkspaceConnection) HasSourceIdentifier() bool`

HasSourceIdentifier returns a boolean if a field has been set.

### GetTenantId

`func (o *WorkspaceConnection) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *WorkspaceConnection) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *WorkspaceConnection) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetType

`func (o *WorkspaceConnection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkspaceConnection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkspaceConnection) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkspaceConnection) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WorkspaceConnection) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkspaceConnection) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkspaceConnection) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WorkspaceConnection) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *WorkspaceConnection) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *WorkspaceConnection) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *WorkspaceConnection) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *WorkspaceConnection) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *WorkspaceConnection) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *WorkspaceConnection) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *WorkspaceConnection) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetVersionId

`func (o *WorkspaceConnection) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *WorkspaceConnection) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *WorkspaceConnection) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.


### GetWorkspaceId

`func (o *WorkspaceConnection) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WorkspaceConnection) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WorkspaceConnection) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


