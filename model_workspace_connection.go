/*
Turbot Pipes

Turbot Pipes is an intelligence, automation & security platform built specifically for DevOps.

API version: {{OPEN_API_VERSION}}
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipes

import (
	"encoding/json"
)

// WorkspaceConnection struct for WorkspaceConnection
type WorkspaceConnection struct {
	Association *WorkspaceConnectionAssociation    `json:"association,omitempty"`
	Config      *map[string]interface{} `json:"config,omitempty"`
	// The source of the configuration that the conection will use. One of `self` or `integration`.
	ConfigSource *string `json:"config_source,omitempty"`
	// The time of creation in ISO 8601 UTC.
	CreatedAt string `json:"created_at"`
	CreatedBy *User  `json:"created_by,omitempty"`
	// The ID of the user that created this.
	CreatedById string `json:"created_by_id"`
	// The source of the credential that the conection will use. One of `self` or `integration`.
	CredentialSource *string `json:"credential_source,omitempty"`
	// The time the item was deleted in ISO 8601 UTC.
	DeletedAt *string `json:"deleted_at,omitempty"`
	DeletedBy *User   `json:"deleted_by,omitempty"`
	// The ID of the user that performed the deletion.
	DeletedById string `json:"deleted_by_id"`
	// The handle name of the  connection.
	Handle *string `json:"handle,omitempty"`
	// The dynamically-generated handle for the connection. Only populated if this is a discovered connection.
	HandleDynamic *string `json:"handle_dynamic,omitempty"`
	// The handle mode for the connection.
	HandleMode *string `json:"handle_mode,omitempty"`
	// The unique identifier for the connection.
	Id string `json:"id"`
	// The unique identifier for an identity where the connection has been created.
	IdentityId  *string      `json:"identity_id,omitempty"`
	Integration *Integration `json:"integration,omitempty"`
	// The source identifier for this connection. Only populated if its a connection thats been discovered by an integration.
	IntegrationResourceIdentifier *string `json:"integration_resource_identifier,omitempty"`
	// A friendly resource name for the connection. Only populated if its a connection thats been doscovered by an integration.
	IntegrationResourceName *string `json:"integration_resource_name,omitempty"`
	// The source path for this connection. Only populated if its a connection thats been discovered by an integration.
	IntegrationResourcePath *string `json:"integration_resource_path,omitempty"`
	// The source type for this connection. Only populated if its a connection thats been discovered by an integration.
	IntegrationResourceType *string `json:"integration_resource_type,omitempty"`
	// The ID of the aggregator that manages this connection. Only populated if this is a discovered connection.
	ManagedById *string `json:"managed_by_id,omitempty"`
	// The id of the entity where the connection is stored. Can be either tenant, identity, workspace or connection-folder.
	ParentId string `json:"parent_id"`
	// The plugin name for the connection.
	Plugin *string `json:"plugin,omitempty"`
	// The plugin version for the connection.
	PluginVersion *string `json:"plugin_version,omitempty"`
	// The unique identifier for the tenant where the connection has been created.
	TenantId string `json:"tenant_id"`
	// The title of the connection. Only populated when the connection is of type connection folder.
	Title *string `json:"title,omitempty"`
	// The trunk for the connection.
	Trunk *[]ConnectionTrunkItem `json:"trunk,omitempty"`
	// Type of connection i.e aggregator or connection.
	Type *string `json:"type,omitempty"`
	// The time of the last update in ISO 8601 UTC.
	UpdatedAt *string `json:"updated_at,omitempty"`
	UpdatedBy *User   `json:"updated_by,omitempty"`
	// The ID of the user that performed the last update.
	UpdatedById string `json:"updated_by_id"`
	// The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item.
	VersionId int32 `json:"version_id"`
	// The unique identifier for the workspace where the connection has been created.
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

// NewWorkspaceConnection instantiates a new WorkspaceConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceConnection(createdAt string, createdById string, deletedById string, id string, parentId string, tenantId string, updatedById string, versionId int32) *WorkspaceConnection {
	this := WorkspaceConnection{}
	this.CreatedAt = createdAt
	this.CreatedById = createdById
	this.DeletedById = deletedById
	this.Id = id
	this.ParentId = parentId
	this.TenantId = tenantId
	this.UpdatedById = updatedById
	this.VersionId = versionId
	return &this
}

// NewWorkspaceConnectionWithDefaults instantiates a new WorkspaceConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceConnectionWithDefaults() *WorkspaceConnection {
	this := WorkspaceConnection{}
	return &this
}

// GetAssociation returns the Association field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetAssociation() WorkspaceConnectionAssociation {
	if o == nil || o.Association == nil {
		var ret WorkspaceConnectionAssociation
		return ret
	}
	return *o.Association
}

// GetAssociationOk returns a tuple with the Association field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetAssociationOk() (*WorkspaceConnectionAssociation, bool) {
	if o == nil || o.Association == nil {
		return nil, false
	}
	return o.Association, true
}

// HasAssociation returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasAssociation() bool {
	if o != nil && o.Association != nil {
		return true
	}

	return false
}

// SetAssociation gets a reference to the given WorkspaceConnectionAssociation and assigns it to the Association field.
func (o *WorkspaceConnection) SetAssociation(v WorkspaceConnectionAssociation) {
	o.Association = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetConfig() map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *WorkspaceConnection) SetConfig(v map[string]interface{}) {
	o.Config = &v
}

// GetConfigSource returns the ConfigSource field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetConfigSource() string {
	if o == nil || o.ConfigSource == nil {
		var ret string
		return ret
	}
	return *o.ConfigSource
}

// GetConfigSourceOk returns a tuple with the ConfigSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetConfigSourceOk() (*string, bool) {
	if o == nil || o.ConfigSource == nil {
		return nil, false
	}
	return o.ConfigSource, true
}

// HasConfigSource returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasConfigSource() bool {
	if o != nil && o.ConfigSource != nil {
		return true
	}

	return false
}

// SetConfigSource gets a reference to the given string and assigns it to the ConfigSource field.
func (o *WorkspaceConnection) SetConfigSource(v string) {
	o.ConfigSource = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *WorkspaceConnection) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WorkspaceConnection) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetCreatedBy() User {
	if o == nil || o.CreatedBy == nil {
		var ret User
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetCreatedByOk() (*User, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given User and assigns it to the CreatedBy field.
func (o *WorkspaceConnection) SetCreatedBy(v User) {
	o.CreatedBy = &v
}

// GetCreatedById returns the CreatedById field value
func (o *WorkspaceConnection) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *WorkspaceConnection) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetCredentialSource returns the CredentialSource field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetCredentialSource() string {
	if o == nil || o.CredentialSource == nil {
		var ret string
		return ret
	}
	return *o.CredentialSource
}

// GetCredentialSourceOk returns a tuple with the CredentialSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetCredentialSourceOk() (*string, bool) {
	if o == nil || o.CredentialSource == nil {
		return nil, false
	}
	return o.CredentialSource, true
}

// HasCredentialSource returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasCredentialSource() bool {
	if o != nil && o.CredentialSource != nil {
		return true
	}

	return false
}

// SetCredentialSource gets a reference to the given string and assigns it to the CredentialSource field.
func (o *WorkspaceConnection) SetCredentialSource(v string) {
	o.CredentialSource = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetDeletedAt() string {
	if o == nil || o.DeletedAt == nil {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetDeletedAtOk() (*string, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *WorkspaceConnection) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetDeletedBy returns the DeletedBy field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetDeletedBy() User {
	if o == nil || o.DeletedBy == nil {
		var ret User
		return ret
	}
	return *o.DeletedBy
}

// GetDeletedByOk returns a tuple with the DeletedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetDeletedByOk() (*User, bool) {
	if o == nil || o.DeletedBy == nil {
		return nil, false
	}
	return o.DeletedBy, true
}

// HasDeletedBy returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasDeletedBy() bool {
	if o != nil && o.DeletedBy != nil {
		return true
	}

	return false
}

// SetDeletedBy gets a reference to the given User and assigns it to the DeletedBy field.
func (o *WorkspaceConnection) SetDeletedBy(v User) {
	o.DeletedBy = &v
}

// GetDeletedById returns the DeletedById field value
func (o *WorkspaceConnection) GetDeletedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeletedById
}

// GetDeletedByIdOk returns a tuple with the DeletedById field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetDeletedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedById, true
}

// SetDeletedById sets field value
func (o *WorkspaceConnection) SetDeletedById(v string) {
	o.DeletedById = v
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasHandle() bool {
	if o != nil && o.Handle != nil {
		return true
	}

	return false
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *WorkspaceConnection) SetHandle(v string) {
	o.Handle = &v
}

// GetHandleDynamic returns the HandleDynamic field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetHandleDynamic() string {
	if o == nil || o.HandleDynamic == nil {
		var ret string
		return ret
	}
	return *o.HandleDynamic
}

// GetHandleDynamicOk returns a tuple with the HandleDynamic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetHandleDynamicOk() (*string, bool) {
	if o == nil || o.HandleDynamic == nil {
		return nil, false
	}
	return o.HandleDynamic, true
}

// HasHandleDynamic returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasHandleDynamic() bool {
	if o != nil && o.HandleDynamic != nil {
		return true
	}

	return false
}

// SetHandleDynamic gets a reference to the given string and assigns it to the HandleDynamic field.
func (o *WorkspaceConnection) SetHandleDynamic(v string) {
	o.HandleDynamic = &v
}

// GetHandleMode returns the HandleMode field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetHandleMode() string {
	if o == nil || o.HandleMode == nil {
		var ret string
		return ret
	}
	return *o.HandleMode
}

// GetHandleModeOk returns a tuple with the HandleMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetHandleModeOk() (*string, bool) {
	if o == nil || o.HandleMode == nil {
		return nil, false
	}
	return o.HandleMode, true
}

// HasHandleMode returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasHandleMode() bool {
	if o != nil && o.HandleMode != nil {
		return true
	}

	return false
}

// SetHandleMode gets a reference to the given string and assigns it to the HandleMode field.
func (o *WorkspaceConnection) SetHandleMode(v string) {
	o.HandleMode = &v
}

// GetId returns the Id field value
func (o *WorkspaceConnection) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkspaceConnection) SetId(v string) {
	o.Id = v
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetIdentityId() string {
	if o == nil || o.IdentityId == nil {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetIdentityIdOk() (*string, bool) {
	if o == nil || o.IdentityId == nil {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasIdentityId() bool {
	if o != nil && o.IdentityId != nil {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *WorkspaceConnection) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetIntegration() Integration {
	if o == nil || o.Integration == nil {
		var ret Integration
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetIntegrationOk() (*Integration, bool) {
	if o == nil || o.Integration == nil {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasIntegration() bool {
	if o != nil && o.Integration != nil {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given Integration and assigns it to the Integration field.
func (o *WorkspaceConnection) SetIntegration(v Integration) {
	o.Integration = &v
}

// GetIntegrationResourceIdentifier returns the IntegrationResourceIdentifier field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetIntegrationResourceIdentifier() string {
	if o == nil || o.IntegrationResourceIdentifier == nil {
		var ret string
		return ret
	}
	return *o.IntegrationResourceIdentifier
}

// GetIntegrationResourceIdentifierOk returns a tuple with the IntegrationResourceIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetIntegrationResourceIdentifierOk() (*string, bool) {
	if o == nil || o.IntegrationResourceIdentifier == nil {
		return nil, false
	}
	return o.IntegrationResourceIdentifier, true
}

// HasIntegrationResourceIdentifier returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasIntegrationResourceIdentifier() bool {
	if o != nil && o.IntegrationResourceIdentifier != nil {
		return true
	}

	return false
}

// SetIntegrationResourceIdentifier gets a reference to the given string and assigns it to the IntegrationResourceIdentifier field.
func (o *WorkspaceConnection) SetIntegrationResourceIdentifier(v string) {
	o.IntegrationResourceIdentifier = &v
}

// GetIntegrationResourceName returns the IntegrationResourceName field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetIntegrationResourceName() string {
	if o == nil || o.IntegrationResourceName == nil {
		var ret string
		return ret
	}
	return *o.IntegrationResourceName
}

// GetIntegrationResourceNameOk returns a tuple with the IntegrationResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetIntegrationResourceNameOk() (*string, bool) {
	if o == nil || o.IntegrationResourceName == nil {
		return nil, false
	}
	return o.IntegrationResourceName, true
}

// HasIntegrationResourceName returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasIntegrationResourceName() bool {
	if o != nil && o.IntegrationResourceName != nil {
		return true
	}

	return false
}

// SetIntegrationResourceName gets a reference to the given string and assigns it to the IntegrationResourceName field.
func (o *WorkspaceConnection) SetIntegrationResourceName(v string) {
	o.IntegrationResourceName = &v
}

// GetIntegrationResourcePath returns the IntegrationResourcePath field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetIntegrationResourcePath() string {
	if o == nil || o.IntegrationResourcePath == nil {
		var ret string
		return ret
	}
	return *o.IntegrationResourcePath
}

// GetIntegrationResourcePathOk returns a tuple with the IntegrationResourcePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetIntegrationResourcePathOk() (*string, bool) {
	if o == nil || o.IntegrationResourcePath == nil {
		return nil, false
	}
	return o.IntegrationResourcePath, true
}

// HasIntegrationResourcePath returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasIntegrationResourcePath() bool {
	if o != nil && o.IntegrationResourcePath != nil {
		return true
	}

	return false
}

// SetIntegrationResourcePath gets a reference to the given string and assigns it to the IntegrationResourcePath field.
func (o *WorkspaceConnection) SetIntegrationResourcePath(v string) {
	o.IntegrationResourcePath = &v
}

// GetIntegrationResourceType returns the IntegrationResourceType field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetIntegrationResourceType() string {
	if o == nil || o.IntegrationResourceType == nil {
		var ret string
		return ret
	}
	return *o.IntegrationResourceType
}

// GetIntegrationResourceTypeOk returns a tuple with the IntegrationResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetIntegrationResourceTypeOk() (*string, bool) {
	if o == nil || o.IntegrationResourceType == nil {
		return nil, false
	}
	return o.IntegrationResourceType, true
}

// HasIntegrationResourceType returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasIntegrationResourceType() bool {
	if o != nil && o.IntegrationResourceType != nil {
		return true
	}

	return false
}

// SetIntegrationResourceType gets a reference to the given string and assigns it to the IntegrationResourceType field.
func (o *WorkspaceConnection) SetIntegrationResourceType(v string) {
	o.IntegrationResourceType = &v
}

// GetManagedById returns the ManagedById field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetManagedById() string {
	if o == nil || o.ManagedById == nil {
		var ret string
		return ret
	}
	return *o.ManagedById
}

// GetManagedByIdOk returns a tuple with the ManagedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetManagedByIdOk() (*string, bool) {
	if o == nil || o.ManagedById == nil {
		return nil, false
	}
	return o.ManagedById, true
}

// HasManagedById returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasManagedById() bool {
	if o != nil && o.ManagedById != nil {
		return true
	}

	return false
}

// SetManagedById gets a reference to the given string and assigns it to the ManagedById field.
func (o *WorkspaceConnection) SetManagedById(v string) {
	o.ManagedById = &v
}

// GetParentId returns the ParentId field value
func (o *WorkspaceConnection) GetParentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetParentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentId, true
}

// SetParentId sets field value
func (o *WorkspaceConnection) SetParentId(v string) {
	o.ParentId = v
}

// GetPlugin returns the Plugin field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetPlugin() string {
	if o == nil || o.Plugin == nil {
		var ret string
		return ret
	}
	return *o.Plugin
}

// GetPluginOk returns a tuple with the Plugin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetPluginOk() (*string, bool) {
	if o == nil || o.Plugin == nil {
		return nil, false
	}
	return o.Plugin, true
}

// HasPlugin returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasPlugin() bool {
	if o != nil && o.Plugin != nil {
		return true
	}

	return false
}

// SetPlugin gets a reference to the given string and assigns it to the Plugin field.
func (o *WorkspaceConnection) SetPlugin(v string) {
	o.Plugin = &v
}

// GetPluginVersion returns the PluginVersion field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetPluginVersion() string {
	if o == nil || o.PluginVersion == nil {
		var ret string
		return ret
	}
	return *o.PluginVersion
}

// GetPluginVersionOk returns a tuple with the PluginVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetPluginVersionOk() (*string, bool) {
	if o == nil || o.PluginVersion == nil {
		return nil, false
	}
	return o.PluginVersion, true
}

// HasPluginVersion returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasPluginVersion() bool {
	if o != nil && o.PluginVersion != nil {
		return true
	}

	return false
}

// SetPluginVersion gets a reference to the given string and assigns it to the PluginVersion field.
func (o *WorkspaceConnection) SetPluginVersion(v string) {
	o.PluginVersion = &v
}

// GetTenantId returns the TenantId field value
func (o *WorkspaceConnection) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *WorkspaceConnection) SetTenantId(v string) {
	o.TenantId = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *WorkspaceConnection) SetTitle(v string) {
	o.Title = &v
}

// GetTrunk returns the Trunk field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetTrunk() []ConnectionTrunkItem {
	if o == nil || o.Trunk == nil {
		var ret []ConnectionTrunkItem
		return ret
	}
	return *o.Trunk
}

// GetTrunkOk returns a tuple with the Trunk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetTrunkOk() (*[]ConnectionTrunkItem, bool) {
	if o == nil || o.Trunk == nil {
		return nil, false
	}
	return o.Trunk, true
}

// HasTrunk returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasTrunk() bool {
	if o != nil && o.Trunk != nil {
		return true
	}

	return false
}

// SetTrunk gets a reference to the given []ConnectionTrunkItem and assigns it to the Trunk field.
func (o *WorkspaceConnection) SetTrunk(v []ConnectionTrunkItem) {
	o.Trunk = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkspaceConnection) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *WorkspaceConnection) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetUpdatedBy() User {
	if o == nil || o.UpdatedBy == nil {
		var ret User
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetUpdatedByOk() (*User, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given User and assigns it to the UpdatedBy field.
func (o *WorkspaceConnection) SetUpdatedBy(v User) {
	o.UpdatedBy = &v
}

// GetUpdatedById returns the UpdatedById field value
func (o *WorkspaceConnection) GetUpdatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedById
}

// GetUpdatedByIdOk returns a tuple with the UpdatedById field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetUpdatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedById, true
}

// SetUpdatedById sets field value
func (o *WorkspaceConnection) SetUpdatedById(v string) {
	o.UpdatedById = v
}

// GetVersionId returns the VersionId field value
func (o *WorkspaceConnection) GetVersionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetVersionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *WorkspaceConnection) SetVersionId(v int32) {
	o.VersionId = v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *WorkspaceConnection) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnection) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *WorkspaceConnection) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *WorkspaceConnection) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o WorkspaceConnection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Association != nil {
		toSerialize["association"] = o.Association
	}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if o.ConfigSource != nil {
		toSerialize["config_source"] = o.ConfigSource
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if true {
		toSerialize["created_by_id"] = o.CreatedById
	}
	if o.CredentialSource != nil {
		toSerialize["credential_source"] = o.CredentialSource
	}
	if o.DeletedAt != nil {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if o.DeletedBy != nil {
		toSerialize["deleted_by"] = o.DeletedBy
	}
	if true {
		toSerialize["deleted_by_id"] = o.DeletedById
	}
	if o.Handle != nil {
		toSerialize["handle"] = o.Handle
	}
	if o.HandleDynamic != nil {
		toSerialize["handle_dynamic"] = o.HandleDynamic
	}
	if o.HandleMode != nil {
		toSerialize["handle_mode"] = o.HandleMode
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.IdentityId != nil {
		toSerialize["identity_id"] = o.IdentityId
	}
	if o.Integration != nil {
		toSerialize["integration"] = o.Integration
	}
	if o.IntegrationResourceIdentifier != nil {
		toSerialize["integration_resource_identifier"] = o.IntegrationResourceIdentifier
	}
	if o.IntegrationResourceName != nil {
		toSerialize["integration_resource_name"] = o.IntegrationResourceName
	}
	if o.IntegrationResourcePath != nil {
		toSerialize["integration_resource_path"] = o.IntegrationResourcePath
	}
	if o.IntegrationResourceType != nil {
		toSerialize["integration_resource_type"] = o.IntegrationResourceType
	}
	if o.ManagedById != nil {
		toSerialize["managed_by_id"] = o.ManagedById
	}
	if true {
		toSerialize["parent_id"] = o.ParentId
	}
	if o.Plugin != nil {
		toSerialize["plugin"] = o.Plugin
	}
	if o.PluginVersion != nil {
		toSerialize["plugin_version"] = o.PluginVersion
	}
	if true {
		toSerialize["tenant_id"] = o.TenantId
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Trunk != nil {
		toSerialize["trunk"] = o.Trunk
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.UpdatedBy != nil {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	if true {
		toSerialize["updated_by_id"] = o.UpdatedById
	}
	if true {
		toSerialize["version_id"] = o.VersionId
	}
	if o.WorkspaceId != nil {
		toSerialize["workspace_id"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableWorkspaceConnection struct {
	value *WorkspaceConnection
	isSet bool
}

func (v NullableWorkspaceConnection) Get() *WorkspaceConnection {
	return v.value
}

func (v *NullableWorkspaceConnection) Set(val *WorkspaceConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceConnection(val *WorkspaceConnection) *NullableWorkspaceConnection {
	return &NullableWorkspaceConnection{value: val, isSet: true}
}

func (v NullableWorkspaceConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
