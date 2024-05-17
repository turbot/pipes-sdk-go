# Connection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
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

## Methods

### NewConnection

`func NewConnection(createdAt string, createdById string, deletedById string, handle string, id string, tenantId string, updatedById string, versionId int32, ) *Connection`

NewConnection instantiates a new Connection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionWithDefaults

`func NewConnectionWithDefaults() *Connection`

NewConnectionWithDefaults instantiates a new Connection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *Connection) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Connection) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Connection) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Connection) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Connection) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Connection) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Connection) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Connection) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Connection) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Connection) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Connection) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *Connection) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *Connection) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *Connection) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetDeletedAt

`func (o *Connection) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Connection) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Connection) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Connection) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *Connection) GetDeletedBy() User`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *Connection) GetDeletedByOk() (*User, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *Connection) SetDeletedBy(v User)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *Connection) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedById

`func (o *Connection) GetDeletedById() string`

GetDeletedById returns the DeletedById field if non-nil, zero value otherwise.

### GetDeletedByIdOk

`func (o *Connection) GetDeletedByIdOk() (*string, bool)`

GetDeletedByIdOk returns a tuple with the DeletedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedById

`func (o *Connection) SetDeletedById(v string)`

SetDeletedById sets DeletedById field to given value.


### GetHandle

`func (o *Connection) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *Connection) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *Connection) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetHandleDynamic

`func (o *Connection) GetHandleDynamic() string`

GetHandleDynamic returns the HandleDynamic field if non-nil, zero value otherwise.

### GetHandleDynamicOk

`func (o *Connection) GetHandleDynamicOk() (*string, bool)`

GetHandleDynamicOk returns a tuple with the HandleDynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandleDynamic

`func (o *Connection) SetHandleDynamic(v string)`

SetHandleDynamic sets HandleDynamic field to given value.

### HasHandleDynamic

`func (o *Connection) HasHandleDynamic() bool`

HasHandleDynamic returns a boolean if a field has been set.

### GetHandleMode

`func (o *Connection) GetHandleMode() string`

GetHandleMode returns the HandleMode field if non-nil, zero value otherwise.

### GetHandleModeOk

`func (o *Connection) GetHandleModeOk() (*string, bool)`

GetHandleModeOk returns a tuple with the HandleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandleMode

`func (o *Connection) SetHandleMode(v string)`

SetHandleMode sets HandleMode field to given value.

### HasHandleMode

`func (o *Connection) HasHandleMode() bool`

HasHandleMode returns a boolean if a field has been set.

### GetId

`func (o *Connection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Connection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Connection) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityId

`func (o *Connection) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *Connection) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *Connection) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *Connection) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetIntegrationResourceId

`func (o *Connection) GetIntegrationResourceId() string`

GetIntegrationResourceId returns the IntegrationResourceId field if non-nil, zero value otherwise.

### GetIntegrationResourceIdOk

`func (o *Connection) GetIntegrationResourceIdOk() (*string, bool)`

GetIntegrationResourceIdOk returns a tuple with the IntegrationResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationResourceId

`func (o *Connection) SetIntegrationResourceId(v string)`

SetIntegrationResourceId sets IntegrationResourceId field to given value.

### HasIntegrationResourceId

`func (o *Connection) HasIntegrationResourceId() bool`

HasIntegrationResourceId returns a boolean if a field has been set.

### GetManagedById

`func (o *Connection) GetManagedById() string`

GetManagedById returns the ManagedById field if non-nil, zero value otherwise.

### GetManagedByIdOk

`func (o *Connection) GetManagedByIdOk() (*string, bool)`

GetManagedByIdOk returns a tuple with the ManagedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedById

`func (o *Connection) SetManagedById(v string)`

SetManagedById sets ManagedById field to given value.

### HasManagedById

`func (o *Connection) HasManagedById() bool`

HasManagedById returns a boolean if a field has been set.

### GetPlugin

`func (o *Connection) GetPlugin() string`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *Connection) GetPluginOk() (*string, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *Connection) SetPlugin(v string)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *Connection) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.

### GetPluginVersion

`func (o *Connection) GetPluginVersion() string`

GetPluginVersion returns the PluginVersion field if non-nil, zero value otherwise.

### GetPluginVersionOk

`func (o *Connection) GetPluginVersionOk() (*string, bool)`

GetPluginVersionOk returns a tuple with the PluginVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginVersion

`func (o *Connection) SetPluginVersion(v string)`

SetPluginVersion sets PluginVersion field to given value.

### HasPluginVersion

`func (o *Connection) HasPluginVersion() bool`

HasPluginVersion returns a boolean if a field has been set.

### GetSourceIdentifier

`func (o *Connection) GetSourceIdentifier() string`

GetSourceIdentifier returns the SourceIdentifier field if non-nil, zero value otherwise.

### GetSourceIdentifierOk

`func (o *Connection) GetSourceIdentifierOk() (*string, bool)`

GetSourceIdentifierOk returns a tuple with the SourceIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceIdentifier

`func (o *Connection) SetSourceIdentifier(v string)`

SetSourceIdentifier sets SourceIdentifier field to given value.

### HasSourceIdentifier

`func (o *Connection) HasSourceIdentifier() bool`

HasSourceIdentifier returns a boolean if a field has been set.

### GetTenantId

`func (o *Connection) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Connection) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Connection) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetType

`func (o *Connection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Connection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Connection) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Connection) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Connection) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Connection) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Connection) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Connection) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Connection) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Connection) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Connection) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Connection) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *Connection) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *Connection) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *Connection) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetVersionId

`func (o *Connection) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *Connection) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *Connection) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


