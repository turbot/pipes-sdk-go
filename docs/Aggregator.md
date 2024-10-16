# Aggregator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** | The aggregator configuration. | [optional] 
**Connections** | **[]string** | The plugin or connection configuration. | 
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) | User information for the user who created this. | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**DeletedAt** | Pointer to **string** | The time the item was deleted in ISO 8601 UTC. | [optional] 
**DeletedBy** | Pointer to [**User**](User.md) | User information for the user that performed the deletion. | [optional] 
**DeletedById** | **string** | The ID of the user that performed the deletion. | 
**Handle** | **string** | The handle name of the aggregator. | 
**HandleDynamic** | Pointer to **string** | The dynamically-generated handle for the aggregator. Only populated if this is a discovered aggregator. | [optional] 
**HandleMode** | Pointer to [**ConnectionHandleMode**](ConnectionHandleMode.md) | The handle mode for the aggregator. | [optional] 
**HandlePrefix** | Pointer to **string** | The handle prefix to use for aggregators and connections discovered by this aggregator. Only populated if this is a dynamic aggregator. | [optional] 
**Id** | **string** | The unique identifier for the aggregator. | 
**IdentityId** | Pointer to **string** | The unique identifier for an identity where the aggregator has been created. | [optional] 
**Integration** | Pointer to [**Integration**](Integration.md) | Details of the integration that manages this aggregator. | [optional] 
**IntegrationResourceIdentifier** | Pointer to **string** | The source identifier for this aggregator. Only populated if its a aggregator thats been discovered by an integration. | [optional] 
**IntegrationResourceName** | Pointer to **string** | A friendly resource name for the aggregator. Only populated if its a aggregator thats been doscovered by an integration. | [optional] 
**IntegrationResourcePath** | Pointer to **string** | The source path for this aggregator. Only populated if its a aggregator thats been discovered by an integration. | [optional] 
**IntegrationResourceType** | Pointer to **string** | The source type for this aggregator. Only populated if its a aggregator thats been discovered by an integration. | [optional] 
**ManagedById** | Pointer to **string** | The ID of the aggregator that manages this aggregator. Only populated if this is a discovered aggregator. | [optional] 
**Plugin** | **string** | The plugin name for the aggregator. | 
**TenantId** | **string** | The unique identifier for the tenant where this aggregator is created. | 
**Type** | Pointer to **string** | Type of connection i.e aggregator or connection. | [optional] 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) | User information for the last user to update this. | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 
**WorkspaceId** | Pointer to **string** | The unique identifier for the workspace. | [optional] 

## Methods

### NewAggregator

`func NewAggregator(connections []string, createdAt string, createdById string, deletedById string, handle string, id string, plugin string, tenantId string, updatedById string, versionId int32, ) *Aggregator`

NewAggregator instantiates a new Aggregator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregatorWithDefaults

`func NewAggregatorWithDefaults() *Aggregator`

NewAggregatorWithDefaults instantiates a new Aggregator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *Aggregator) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Aggregator) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Aggregator) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Aggregator) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConnections

`func (o *Aggregator) GetConnections() []string`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *Aggregator) GetConnectionsOk() (*[]string, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *Aggregator) SetConnections(v []string)`

SetConnections sets Connections field to given value.


### GetCreatedAt

`func (o *Aggregator) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Aggregator) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Aggregator) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Aggregator) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Aggregator) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Aggregator) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Aggregator) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *Aggregator) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *Aggregator) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *Aggregator) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetDeletedAt

`func (o *Aggregator) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Aggregator) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Aggregator) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Aggregator) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *Aggregator) GetDeletedBy() User`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *Aggregator) GetDeletedByOk() (*User, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *Aggregator) SetDeletedBy(v User)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *Aggregator) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedById

`func (o *Aggregator) GetDeletedById() string`

GetDeletedById returns the DeletedById field if non-nil, zero value otherwise.

### GetDeletedByIdOk

`func (o *Aggregator) GetDeletedByIdOk() (*string, bool)`

GetDeletedByIdOk returns a tuple with the DeletedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedById

`func (o *Aggregator) SetDeletedById(v string)`

SetDeletedById sets DeletedById field to given value.


### GetHandle

`func (o *Aggregator) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *Aggregator) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *Aggregator) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetHandleDynamic

`func (o *Aggregator) GetHandleDynamic() string`

GetHandleDynamic returns the HandleDynamic field if non-nil, zero value otherwise.

### GetHandleDynamicOk

`func (o *Aggregator) GetHandleDynamicOk() (*string, bool)`

GetHandleDynamicOk returns a tuple with the HandleDynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandleDynamic

`func (o *Aggregator) SetHandleDynamic(v string)`

SetHandleDynamic sets HandleDynamic field to given value.

### HasHandleDynamic

`func (o *Aggregator) HasHandleDynamic() bool`

HasHandleDynamic returns a boolean if a field has been set.

### GetHandleMode

`func (o *Aggregator) GetHandleMode() ConnectionHandleMode`

GetHandleMode returns the HandleMode field if non-nil, zero value otherwise.

### GetHandleModeOk

`func (o *Aggregator) GetHandleModeOk() (*ConnectionHandleMode, bool)`

GetHandleModeOk returns a tuple with the HandleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandleMode

`func (o *Aggregator) SetHandleMode(v ConnectionHandleMode)`

SetHandleMode sets HandleMode field to given value.

### HasHandleMode

`func (o *Aggregator) HasHandleMode() bool`

HasHandleMode returns a boolean if a field has been set.

### GetHandlePrefix

`func (o *Aggregator) GetHandlePrefix() string`

GetHandlePrefix returns the HandlePrefix field if non-nil, zero value otherwise.

### GetHandlePrefixOk

`func (o *Aggregator) GetHandlePrefixOk() (*string, bool)`

GetHandlePrefixOk returns a tuple with the HandlePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlePrefix

`func (o *Aggregator) SetHandlePrefix(v string)`

SetHandlePrefix sets HandlePrefix field to given value.

### HasHandlePrefix

`func (o *Aggregator) HasHandlePrefix() bool`

HasHandlePrefix returns a boolean if a field has been set.

### GetId

`func (o *Aggregator) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Aggregator) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Aggregator) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityId

`func (o *Aggregator) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *Aggregator) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *Aggregator) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *Aggregator) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetIntegration

`func (o *Aggregator) GetIntegration() Integration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *Aggregator) GetIntegrationOk() (*Integration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *Aggregator) SetIntegration(v Integration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *Aggregator) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetIntegrationResourceIdentifier

`func (o *Aggregator) GetIntegrationResourceIdentifier() string`

GetIntegrationResourceIdentifier returns the IntegrationResourceIdentifier field if non-nil, zero value otherwise.

### GetIntegrationResourceIdentifierOk

`func (o *Aggregator) GetIntegrationResourceIdentifierOk() (*string, bool)`

GetIntegrationResourceIdentifierOk returns a tuple with the IntegrationResourceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationResourceIdentifier

`func (o *Aggregator) SetIntegrationResourceIdentifier(v string)`

SetIntegrationResourceIdentifier sets IntegrationResourceIdentifier field to given value.

### HasIntegrationResourceIdentifier

`func (o *Aggregator) HasIntegrationResourceIdentifier() bool`

HasIntegrationResourceIdentifier returns a boolean if a field has been set.

### GetIntegrationResourceName

`func (o *Aggregator) GetIntegrationResourceName() string`

GetIntegrationResourceName returns the IntegrationResourceName field if non-nil, zero value otherwise.

### GetIntegrationResourceNameOk

`func (o *Aggregator) GetIntegrationResourceNameOk() (*string, bool)`

GetIntegrationResourceNameOk returns a tuple with the IntegrationResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationResourceName

`func (o *Aggregator) SetIntegrationResourceName(v string)`

SetIntegrationResourceName sets IntegrationResourceName field to given value.

### HasIntegrationResourceName

`func (o *Aggregator) HasIntegrationResourceName() bool`

HasIntegrationResourceName returns a boolean if a field has been set.

### GetIntegrationResourcePath

`func (o *Aggregator) GetIntegrationResourcePath() string`

GetIntegrationResourcePath returns the IntegrationResourcePath field if non-nil, zero value otherwise.

### GetIntegrationResourcePathOk

`func (o *Aggregator) GetIntegrationResourcePathOk() (*string, bool)`

GetIntegrationResourcePathOk returns a tuple with the IntegrationResourcePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationResourcePath

`func (o *Aggregator) SetIntegrationResourcePath(v string)`

SetIntegrationResourcePath sets IntegrationResourcePath field to given value.

### HasIntegrationResourcePath

`func (o *Aggregator) HasIntegrationResourcePath() bool`

HasIntegrationResourcePath returns a boolean if a field has been set.

### GetIntegrationResourceType

`func (o *Aggregator) GetIntegrationResourceType() string`

GetIntegrationResourceType returns the IntegrationResourceType field if non-nil, zero value otherwise.

### GetIntegrationResourceTypeOk

`func (o *Aggregator) GetIntegrationResourceTypeOk() (*string, bool)`

GetIntegrationResourceTypeOk returns a tuple with the IntegrationResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationResourceType

`func (o *Aggregator) SetIntegrationResourceType(v string)`

SetIntegrationResourceType sets IntegrationResourceType field to given value.

### HasIntegrationResourceType

`func (o *Aggregator) HasIntegrationResourceType() bool`

HasIntegrationResourceType returns a boolean if a field has been set.

### GetManagedById

`func (o *Aggregator) GetManagedById() string`

GetManagedById returns the ManagedById field if non-nil, zero value otherwise.

### GetManagedByIdOk

`func (o *Aggregator) GetManagedByIdOk() (*string, bool)`

GetManagedByIdOk returns a tuple with the ManagedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedById

`func (o *Aggregator) SetManagedById(v string)`

SetManagedById sets ManagedById field to given value.

### HasManagedById

`func (o *Aggregator) HasManagedById() bool`

HasManagedById returns a boolean if a field has been set.

### GetPlugin

`func (o *Aggregator) GetPlugin() string`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *Aggregator) GetPluginOk() (*string, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *Aggregator) SetPlugin(v string)`

SetPlugin sets Plugin field to given value.


### GetTenantId

`func (o *Aggregator) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Aggregator) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Aggregator) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetType

`func (o *Aggregator) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Aggregator) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Aggregator) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Aggregator) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Aggregator) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Aggregator) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Aggregator) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Aggregator) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Aggregator) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Aggregator) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Aggregator) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Aggregator) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *Aggregator) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *Aggregator) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *Aggregator) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetVersionId

`func (o *Aggregator) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *Aggregator) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *Aggregator) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.


### GetWorkspaceId

`func (o *Aggregator) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *Aggregator) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *Aggregator) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *Aggregator) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


