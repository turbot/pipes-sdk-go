# WorkspaceAggregator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Association** | Pointer to [**WorkspaceAggregatorAssociation**](WorkspaceAggregatorAssociation.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Connections** | **[]string** | The plugin or connection configuration. | 
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**DeletedAt** | Pointer to **string** | The time the item was deleted in ISO 8601 UTC. | [optional] 
**DeletedBy** | Pointer to [**User**](User.md) |  | [optional] 
**DeletedById** | **string** | The ID of the user that performed the deletion. | 
**Handle** | **string** | The handle name of the aggregator. | 
**HandleDynamic** | Pointer to **string** | The dynamically-generated handle for the aggregator. Only populated if this is a discovered aggregator. | [optional] 
**HandleMode** | Pointer to **string** | The handle mode for the aggregator. | [optional] 
**HandlePrefix** | Pointer to **string** | The handle prefix to use for aggregators and connections discovered by this aggregator. Only populated if this is a dynamic aggregator. | [optional] 
**Id** | **string** | The unique identifier for the aggregator. | 
**IdentityId** | Pointer to **string** | The unique identifier for an identity where the aggregator has been created. | [optional] 
**Integration** | Pointer to [**Integration**](Integration.md) |  | [optional] 
**IntegrationResourceIdentifier** | Pointer to **string** | The source identifier for this aggregator. Only populated if its a aggregator thats been discovered by an integration. | [optional] 
**IntegrationResourceName** | Pointer to **string** | A friendly resource name for the aggregator. Only populated if its a aggregator thats been doscovered by an integration. | [optional] 
**IntegrationResourcePath** | Pointer to **string** | The source path for this aggregator. Only populated if its a aggregator thats been discovered by an integration. | [optional] 
**IntegrationResourceType** | Pointer to **string** | The source type for this aggregator. Only populated if its a aggregator thats been discovered by an integration. | [optional] 
**ManagedById** | Pointer to **string** | The ID of the aggregator that manages this aggregator. Only populated if this is a discovered aggregator. | [optional] 
**Plugin** | **string** | The plugin name for the aggregator. | 
**TenantId** | **string** | The unique identifier for the tenant where this aggregator is created. | 
**Type** | Pointer to **string** | Type of connection i.e aggregator or connection. | [optional] 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 
**WorkspaceId** | Pointer to **string** | The unique identifier for the workspace. | [optional] 

## Methods

### NewWorkspaceAggregator

`func NewWorkspaceAggregator(connections []string, createdAt string, createdById string, deletedById string, handle string, id string, plugin string, tenantId string, updatedById string, versionId int32, ) *WorkspaceAggregator`

NewWorkspaceAggregator instantiates a new WorkspaceAggregator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceAggregatorWithDefaults

`func NewWorkspaceAggregatorWithDefaults() *WorkspaceAggregator`

NewWorkspaceAggregatorWithDefaults instantiates a new WorkspaceAggregator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociation

`func (o *WorkspaceAggregator) GetAssociation() WorkspaceAggregatorAssociation`

GetAssociation returns the Association field if non-nil, zero value otherwise.

### GetAssociationOk

`func (o *WorkspaceAggregator) GetAssociationOk() (*WorkspaceAggregatorAssociation, bool)`

GetAssociationOk returns a tuple with the Association field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociation

`func (o *WorkspaceAggregator) SetAssociation(v WorkspaceAggregatorAssociation)`

SetAssociation sets Association field to given value.

### HasAssociation

`func (o *WorkspaceAggregator) HasAssociation() bool`

HasAssociation returns a boolean if a field has been set.

### GetConfig

`func (o *WorkspaceAggregator) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *WorkspaceAggregator) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *WorkspaceAggregator) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *WorkspaceAggregator) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConnections

`func (o *WorkspaceAggregator) GetConnections() []string`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *WorkspaceAggregator) GetConnectionsOk() (*[]string, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *WorkspaceAggregator) SetConnections(v []string)`

SetConnections sets Connections field to given value.


### GetCreatedAt

`func (o *WorkspaceAggregator) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkspaceAggregator) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkspaceAggregator) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *WorkspaceAggregator) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WorkspaceAggregator) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WorkspaceAggregator) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WorkspaceAggregator) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *WorkspaceAggregator) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *WorkspaceAggregator) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *WorkspaceAggregator) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetDeletedAt

`func (o *WorkspaceAggregator) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *WorkspaceAggregator) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *WorkspaceAggregator) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *WorkspaceAggregator) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *WorkspaceAggregator) GetDeletedBy() User`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *WorkspaceAggregator) GetDeletedByOk() (*User, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *WorkspaceAggregator) SetDeletedBy(v User)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *WorkspaceAggregator) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedById

`func (o *WorkspaceAggregator) GetDeletedById() string`

GetDeletedById returns the DeletedById field if non-nil, zero value otherwise.

### GetDeletedByIdOk

`func (o *WorkspaceAggregator) GetDeletedByIdOk() (*string, bool)`

GetDeletedByIdOk returns a tuple with the DeletedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedById

`func (o *WorkspaceAggregator) SetDeletedById(v string)`

SetDeletedById sets DeletedById field to given value.


### GetHandle

`func (o *WorkspaceAggregator) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *WorkspaceAggregator) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *WorkspaceAggregator) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetHandleDynamic

`func (o *WorkspaceAggregator) GetHandleDynamic() string`

GetHandleDynamic returns the HandleDynamic field if non-nil, zero value otherwise.

### GetHandleDynamicOk

`func (o *WorkspaceAggregator) GetHandleDynamicOk() (*string, bool)`

GetHandleDynamicOk returns a tuple with the HandleDynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandleDynamic

`func (o *WorkspaceAggregator) SetHandleDynamic(v string)`

SetHandleDynamic sets HandleDynamic field to given value.

### HasHandleDynamic

`func (o *WorkspaceAggregator) HasHandleDynamic() bool`

HasHandleDynamic returns a boolean if a field has been set.

### GetHandleMode

`func (o *WorkspaceAggregator) GetHandleMode() string`

GetHandleMode returns the HandleMode field if non-nil, zero value otherwise.

### GetHandleModeOk

`func (o *WorkspaceAggregator) GetHandleModeOk() (*string, bool)`

GetHandleModeOk returns a tuple with the HandleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandleMode

`func (o *WorkspaceAggregator) SetHandleMode(v string)`

SetHandleMode sets HandleMode field to given value.

### HasHandleMode

`func (o *WorkspaceAggregator) HasHandleMode() bool`

HasHandleMode returns a boolean if a field has been set.

### GetHandlePrefix

`func (o *WorkspaceAggregator) GetHandlePrefix() string`

GetHandlePrefix returns the HandlePrefix field if non-nil, zero value otherwise.

### GetHandlePrefixOk

`func (o *WorkspaceAggregator) GetHandlePrefixOk() (*string, bool)`

GetHandlePrefixOk returns a tuple with the HandlePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlePrefix

`func (o *WorkspaceAggregator) SetHandlePrefix(v string)`

SetHandlePrefix sets HandlePrefix field to given value.

### HasHandlePrefix

`func (o *WorkspaceAggregator) HasHandlePrefix() bool`

HasHandlePrefix returns a boolean if a field has been set.

### GetId

`func (o *WorkspaceAggregator) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkspaceAggregator) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkspaceAggregator) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityId

`func (o *WorkspaceAggregator) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *WorkspaceAggregator) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *WorkspaceAggregator) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *WorkspaceAggregator) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetIntegration

`func (o *WorkspaceAggregator) GetIntegration() Integration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *WorkspaceAggregator) GetIntegrationOk() (*Integration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *WorkspaceAggregator) SetIntegration(v Integration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *WorkspaceAggregator) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetIntegrationResourceIdentifier

`func (o *WorkspaceAggregator) GetIntegrationResourceIdentifier() string`

GetIntegrationResourceIdentifier returns the IntegrationResourceIdentifier field if non-nil, zero value otherwise.

### GetIntegrationResourceIdentifierOk

`func (o *WorkspaceAggregator) GetIntegrationResourceIdentifierOk() (*string, bool)`

GetIntegrationResourceIdentifierOk returns a tuple with the IntegrationResourceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationResourceIdentifier

`func (o *WorkspaceAggregator) SetIntegrationResourceIdentifier(v string)`

SetIntegrationResourceIdentifier sets IntegrationResourceIdentifier field to given value.

### HasIntegrationResourceIdentifier

`func (o *WorkspaceAggregator) HasIntegrationResourceIdentifier() bool`

HasIntegrationResourceIdentifier returns a boolean if a field has been set.

### GetIntegrationResourceName

`func (o *WorkspaceAggregator) GetIntegrationResourceName() string`

GetIntegrationResourceName returns the IntegrationResourceName field if non-nil, zero value otherwise.

### GetIntegrationResourceNameOk

`func (o *WorkspaceAggregator) GetIntegrationResourceNameOk() (*string, bool)`

GetIntegrationResourceNameOk returns a tuple with the IntegrationResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationResourceName

`func (o *WorkspaceAggregator) SetIntegrationResourceName(v string)`

SetIntegrationResourceName sets IntegrationResourceName field to given value.

### HasIntegrationResourceName

`func (o *WorkspaceAggregator) HasIntegrationResourceName() bool`

HasIntegrationResourceName returns a boolean if a field has been set.

### GetIntegrationResourcePath

`func (o *WorkspaceAggregator) GetIntegrationResourcePath() string`

GetIntegrationResourcePath returns the IntegrationResourcePath field if non-nil, zero value otherwise.

### GetIntegrationResourcePathOk

`func (o *WorkspaceAggregator) GetIntegrationResourcePathOk() (*string, bool)`

GetIntegrationResourcePathOk returns a tuple with the IntegrationResourcePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationResourcePath

`func (o *WorkspaceAggregator) SetIntegrationResourcePath(v string)`

SetIntegrationResourcePath sets IntegrationResourcePath field to given value.

### HasIntegrationResourcePath

`func (o *WorkspaceAggregator) HasIntegrationResourcePath() bool`

HasIntegrationResourcePath returns a boolean if a field has been set.

### GetIntegrationResourceType

`func (o *WorkspaceAggregator) GetIntegrationResourceType() string`

GetIntegrationResourceType returns the IntegrationResourceType field if non-nil, zero value otherwise.

### GetIntegrationResourceTypeOk

`func (o *WorkspaceAggregator) GetIntegrationResourceTypeOk() (*string, bool)`

GetIntegrationResourceTypeOk returns a tuple with the IntegrationResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationResourceType

`func (o *WorkspaceAggregator) SetIntegrationResourceType(v string)`

SetIntegrationResourceType sets IntegrationResourceType field to given value.

### HasIntegrationResourceType

`func (o *WorkspaceAggregator) HasIntegrationResourceType() bool`

HasIntegrationResourceType returns a boolean if a field has been set.

### GetManagedById

`func (o *WorkspaceAggregator) GetManagedById() string`

GetManagedById returns the ManagedById field if non-nil, zero value otherwise.

### GetManagedByIdOk

`func (o *WorkspaceAggregator) GetManagedByIdOk() (*string, bool)`

GetManagedByIdOk returns a tuple with the ManagedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedById

`func (o *WorkspaceAggregator) SetManagedById(v string)`

SetManagedById sets ManagedById field to given value.

### HasManagedById

`func (o *WorkspaceAggregator) HasManagedById() bool`

HasManagedById returns a boolean if a field has been set.

### GetPlugin

`func (o *WorkspaceAggregator) GetPlugin() string`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *WorkspaceAggregator) GetPluginOk() (*string, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *WorkspaceAggregator) SetPlugin(v string)`

SetPlugin sets Plugin field to given value.


### GetTenantId

`func (o *WorkspaceAggregator) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *WorkspaceAggregator) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *WorkspaceAggregator) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetType

`func (o *WorkspaceAggregator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkspaceAggregator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkspaceAggregator) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkspaceAggregator) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *WorkspaceAggregator) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkspaceAggregator) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkspaceAggregator) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WorkspaceAggregator) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *WorkspaceAggregator) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *WorkspaceAggregator) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *WorkspaceAggregator) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *WorkspaceAggregator) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *WorkspaceAggregator) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *WorkspaceAggregator) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *WorkspaceAggregator) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetVersionId

`func (o *WorkspaceAggregator) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *WorkspaceAggregator) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *WorkspaceAggregator) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.


### GetWorkspaceId

`func (o *WorkspaceAggregator) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WorkspaceAggregator) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WorkspaceAggregator) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *WorkspaceAggregator) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


