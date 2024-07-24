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

// Aggregator struct for Aggregator
type Aggregator struct {
	Config *map[string]interface{} `json:"config,omitempty"`
	// The plugin or connection configuration.
	Connections []string `json:"connections"`
	// The time of creation in ISO 8601 UTC.
	CreatedAt string `json:"created_at"`
	CreatedBy *User  `json:"created_by,omitempty"`
	// The ID of the user that created this.
	CreatedById string `json:"created_by_id"`
	// The time the item was deleted in ISO 8601 UTC.
	DeletedAt *string `json:"deleted_at,omitempty"`
	DeletedBy *User   `json:"deleted_by,omitempty"`
	// The ID of the user that performed the deletion.
	DeletedById string `json:"deleted_by_id"`
	// The handle name of the aggregator.
	Handle string `json:"handle"`
	// The dynamically-generated handle for the aggregator. Only populated if this is a discovered aggregator.
	HandleDynamic *string `json:"handle_dynamic,omitempty"`
	// The handle mode for the aggregator.
	HandleMode *string `json:"handle_mode,omitempty"`
	// The handle prefix to use for aggregators and connections discovered by this aggregator. Only populated if this is a dynamic aggregator.
	HandlePrefix *string `json:"handle_prefix,omitempty"`
	// The unique identifier for the aggregator.
	Id string `json:"id"`
	// The unique identifier for an identity where the aggregator has been created.
	IdentityId  *string      `json:"identity_id,omitempty"`
	Integration *Integration `json:"integration,omitempty"`
	// The source identifier for this aggregator. Only populated if its a aggregator thats been discovered by an integration.
	IntegrationResourceIdentifier *string `json:"integration_resource_identifier,omitempty"`
	// A friendly resource name for the aggregator. Only populated if its a aggregator thats been doscovered by an integration.
	IntegrationResourceName *string `json:"integration_resource_name,omitempty"`
	// The source path for this aggregator. Only populated if its a aggregator thats been discovered by an integration.
	IntegrationResourcePath *string `json:"integration_resource_path,omitempty"`
	// The source type for this aggregator. Only populated if its a aggregator thats been discovered by an integration.
	IntegrationResourceType *string `json:"integration_resource_type,omitempty"`
	// The ID of the aggregator that manages this aggregator. Only populated if this is a discovered aggregator.
	ManagedById *string `json:"managed_by_id,omitempty"`
	// The plugin name for the aggregator.
	Plugin string `json:"plugin"`
	// The unique identifier for the tenant where this aggregator is created.
	TenantId string `json:"tenant_id"`
	// Type of connection i.e aggregator or connection.
	Type *string `json:"type,omitempty"`
	// The time of the last update in ISO 8601 UTC.
	UpdatedAt *string `json:"updated_at,omitempty"`
	UpdatedBy *User   `json:"updated_by,omitempty"`
	// The ID of the user that performed the last update.
	UpdatedById string `json:"updated_by_id"`
	// The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item.
	VersionId int32 `json:"version_id"`
	// The unique identifier for the workspace.
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

// NewAggregator instantiates a new Aggregator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAggregator(connections []string, createdAt string, createdById string, deletedById string, handle string, id string, plugin string, tenantId string, updatedById string, versionId int32) *Aggregator {
	this := Aggregator{}
	this.Connections = connections
	this.CreatedAt = createdAt
	this.CreatedById = createdById
	this.DeletedById = deletedById
	this.Handle = handle
	this.Id = id
	this.Plugin = plugin
	this.TenantId = tenantId
	this.UpdatedById = updatedById
	this.VersionId = versionId
	return &this
}

// NewAggregatorWithDefaults instantiates a new Aggregator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAggregatorWithDefaults() *Aggregator {
	this := Aggregator{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Aggregator) GetConfig() map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregator) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Aggregator) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *Aggregator) SetConfig(v map[string]interface{}) {
	o.Config = &v
}

// GetConnections returns the Connections field value
func (o *Aggregator) GetConnections() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
func (o *Aggregator) GetConnectionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connections, true
}

// SetConnections sets field value
func (o *Aggregator) SetConnections(v []string) {
	o.Connections = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Aggregator) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Aggregator) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Aggregator) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *Aggregator) GetCreatedBy() User {
	if o == nil || o.CreatedBy == nil {
		var ret User
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregator) GetCreatedByOk() (*User, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *Aggregator) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given User and assigns it to the CreatedBy field.
func (o *Aggregator) SetCreatedBy(v User) {
	o.CreatedBy = &v
}

// GetCreatedById returns the CreatedById field value
func (o *Aggregator) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *Aggregator) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *Aggregator) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *Aggregator) GetDeletedAt() string {
	if o == nil || o.DeletedAt == nil {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregator) GetDeletedAtOk() (*string, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *Aggregator) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *Aggregator) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetDeletedBy returns the DeletedBy field value if set, zero value otherwise.
func (o *Aggregator) GetDeletedBy() User {
	if o == nil || o.DeletedBy == nil {
		var ret User
		return ret
	}
	return *o.DeletedBy
}

// GetDeletedByOk returns a tuple with the DeletedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregator) GetDeletedByOk() (*User, bool) {
	if o == nil || o.DeletedBy == nil {
		return nil, false
	}
	return o.DeletedBy, true
}

// HasDeletedBy returns a boolean if a field has been set.
func (o *Aggregator) HasDeletedBy() bool {
	if o != nil && o.DeletedBy != nil {
		return true
	}

	return false
}

// SetDeletedBy gets a reference to the given User and assigns it to the DeletedBy field.
func (o *Aggregator) SetDeletedBy(v User) {
	o.DeletedBy = &v
}

// GetDeletedById returns the DeletedById field value
func (o *Aggregator) GetDeletedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeletedById
}

// GetDeletedByIdOk returns a tuple with the DeletedById field value
// and a boolean to check if the value has been set.
func (o *Aggregator) GetDeletedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedById, true
}

// SetDeletedById sets field value
func (o *Aggregator) SetDeletedById(v string) {
	o.DeletedById = v
}

// GetHandle returns the Handle field value
func (o *Aggregator) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *Aggregator) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value
func (o *Aggregator) SetHandle(v string) {
	o.Handle = v
}

// GetHandleDynamic returns the HandleDynamic field value if set, zero value otherwise.
func (o *Aggregator) GetHandleDynamic() string {
	if o == nil || o.HandleDynamic == nil {
		var ret string
		return ret
	}
	return *o.HandleDynamic
}

// GetHandleDynamicOk returns a tuple with the HandleDynamic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregator) GetHandleDynamicOk() (*string, bool) {
	if o == nil || o.HandleDynamic == nil {
		return nil, false
	}
	return o.HandleDynamic, true
}

// HasHandleDynamic returns a boolean if a field has been set.
func (o *Aggregator) HasHandleDynamic() bool {
	if o != nil && o.HandleDynamic != nil {
		return true
	}

	return false
}

// SetHandleDynamic gets a reference to the given string and assigns it to the HandleDynamic field.
func (o *Aggregator) SetHandleDynamic(v string) {
	o.HandleDynamic = &v
}

// GetHandleMode returns the HandleMode field value if set, zero value otherwise.
func (o *Aggregator) GetHandleMode() string {
	if o == nil || o.HandleMode == nil {
		var ret string
		return ret
	}
	return *o.HandleMode
}

// GetHandleModeOk returns a tuple with the HandleMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregator) GetHandleModeOk() (*string, bool) {
	if o == nil || o.HandleMode == nil {
		return nil, false
	}
	return o.HandleMode, true
}

// HasHandleMode returns a boolean if a field has been set.
func (o *Aggregator) HasHandleMode() bool {
	if o != nil && o.HandleMode != nil {
		return true
	}

	return false
}

// SetHandleMode gets a reference to the given string and assigns it to the HandleMode field.
func (o *Aggregator) SetHandleMode(v string) {
	o.HandleMode = &v
}

// GetHandlePrefix returns the HandlePrefix field value if set, zero value otherwise.
func (o *Aggregator) GetHandlePrefix() string {
	if o == nil || o.HandlePrefix == nil {
		var ret string
		return ret
	}
	return *o.HandlePrefix
}

// GetHandlePrefixOk returns a tuple with the HandlePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregator) GetHandlePrefixOk() (*string, bool) {
	if o == nil || o.HandlePrefix == nil {
		return nil, false
	}
	return o.HandlePrefix, true
}

// HasHandlePrefix returns a boolean if a field has been set.
func (o *Aggregator) HasHandlePrefix() bool {
	if o != nil && o.HandlePrefix != nil {
		return true
	}

	return false
}

// SetHandlePrefix gets a reference to the given string and assigns it to the HandlePrefix field.
func (o *Aggregator) SetHandlePrefix(v string) {
	o.HandlePrefix = &v
}

// GetId returns the Id field value
func (o *Aggregator) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Aggregator) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Aggregator) SetId(v string) {
	o.Id = v
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *Aggregator) GetIdentityId() string {
	if o == nil || o.IdentityId == nil {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregator) GetIdentityIdOk() (*string, bool) {
	if o == nil || o.IdentityId == nil {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *Aggregator) HasIdentityId() bool {
	if o != nil && o.IdentityId != nil {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *Aggregator) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *Aggregator) GetIntegration() Integration {
	if o == nil || o.Integration == nil {
		var ret Integration
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregator) GetIntegrationOk() (*Integration, bool) {
	if o == nil || o.Integration == nil {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *Aggregator) HasIntegration() bool {
	if o != nil && o.Integration != nil {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given Integration and assigns it to the Integration field.
func (o *Aggregator) SetIntegration(v Integration) {
	o.Integration = &v
}

// GetIntegrationResourceIdentifier returns the IntegrationResourceIdentifier field value if set, zero value otherwise.
func (o *Aggregator) GetIntegrationResourceIdentifier() string {
	if o == nil || o.IntegrationResourceIdentifier == nil {
		var ret string
		return ret
	}
	return *o.IntegrationResourceIdentifier
}

// GetIntegrationResourceIdentifierOk returns a tuple with the IntegrationResourceIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregator) GetIntegrationResourceIdentifierOk() (*string, bool) {
	if o == nil || o.IntegrationResourceIdentifier == nil {
		return nil, false
	}
	return o.IntegrationResourceIdentifier, true
}

// HasIntegrationResourceIdentifier returns a boolean if a field has been set.
func (o *Aggregator) HasIntegrationResourceIdentifier() bool {
	if o != nil && o.IntegrationResourceIdentifier != nil {
		return true
	}

	return false
}

// SetIntegrationResourceIdentifier gets a reference to the given string and assigns it to the IntegrationResourceIdentifier field.
func (o *Aggregator) SetIntegrationResourceIdentifier(v string) {
	o.IntegrationResourceIdentifier = &v
}

// GetIntegrationResourceName returns the IntegrationResourceName field value if set, zero value otherwise.
func (o *Aggregator) GetIntegrationResourceName() string {
	if o == nil || o.IntegrationResourceName == nil {
		var ret string
		return ret
	}
	return *o.IntegrationResourceName
}

// GetIntegrationResourceNameOk returns a tuple with the IntegrationResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregator) GetIntegrationResourceNameOk() (*string, bool) {
	if o == nil || o.IntegrationResourceName == nil {
		return nil, false
	}
	return o.IntegrationResourceName, true
}

// HasIntegrationResourceName returns a boolean if a field has been set.
func (o *Aggregator) HasIntegrationResourceName() bool {
	if o != nil && o.IntegrationResourceName != nil {
		return true
	}

	return false
}

// SetIntegrationResourceName gets a reference to the given string and assigns it to the IntegrationResourceName field.
func (o *Aggregator) SetIntegrationResourceName(v string) {
	o.IntegrationResourceName = &v
}

// GetIntegrationResourcePath returns the IntegrationResourcePath field value if set, zero value otherwise.
func (o *Aggregator) GetIntegrationResourcePath() string {
	if o == nil || o.IntegrationResourcePath == nil {
		var ret string
		return ret
	}
	return *o.IntegrationResourcePath
}

// GetIntegrationResourcePathOk returns a tuple with the IntegrationResourcePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregator) GetIntegrationResourcePathOk() (*string, bool) {
	if o == nil || o.IntegrationResourcePath == nil {
		return nil, false
	}
	return o.IntegrationResourcePath, true
}

// HasIntegrationResourcePath returns a boolean if a field has been set.
func (o *Aggregator) HasIntegrationResourcePath() bool {
	if o != nil && o.IntegrationResourcePath != nil {
		return true
	}

	return false
}

// SetIntegrationResourcePath gets a reference to the given string and assigns it to the IntegrationResourcePath field.
func (o *Aggregator) SetIntegrationResourcePath(v string) {
	o.IntegrationResourcePath = &v
}

// GetIntegrationResourceType returns the IntegrationResourceType field value if set, zero value otherwise.
func (o *Aggregator) GetIntegrationResourceType() string {
	if o == nil || o.IntegrationResourceType == nil {
		var ret string
		return ret
	}
	return *o.IntegrationResourceType
}

// GetIntegrationResourceTypeOk returns a tuple with the IntegrationResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregator) GetIntegrationResourceTypeOk() (*string, bool) {
	if o == nil || o.IntegrationResourceType == nil {
		return nil, false
	}
	return o.IntegrationResourceType, true
}

// HasIntegrationResourceType returns a boolean if a field has been set.
func (o *Aggregator) HasIntegrationResourceType() bool {
	if o != nil && o.IntegrationResourceType != nil {
		return true
	}

	return false
}

// SetIntegrationResourceType gets a reference to the given string and assigns it to the IntegrationResourceType field.
func (o *Aggregator) SetIntegrationResourceType(v string) {
	o.IntegrationResourceType = &v
}

// GetManagedById returns the ManagedById field value if set, zero value otherwise.
func (o *Aggregator) GetManagedById() string {
	if o == nil || o.ManagedById == nil {
		var ret string
		return ret
	}
	return *o.ManagedById
}

// GetManagedByIdOk returns a tuple with the ManagedById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregator) GetManagedByIdOk() (*string, bool) {
	if o == nil || o.ManagedById == nil {
		return nil, false
	}
	return o.ManagedById, true
}

// HasManagedById returns a boolean if a field has been set.
func (o *Aggregator) HasManagedById() bool {
	if o != nil && o.ManagedById != nil {
		return true
	}

	return false
}

// SetManagedById gets a reference to the given string and assigns it to the ManagedById field.
func (o *Aggregator) SetManagedById(v string) {
	o.ManagedById = &v
}

// GetPlugin returns the Plugin field value
func (o *Aggregator) GetPlugin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plugin
}

// GetPluginOk returns a tuple with the Plugin field value
// and a boolean to check if the value has been set.
func (o *Aggregator) GetPluginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plugin, true
}

// SetPlugin sets field value
func (o *Aggregator) SetPlugin(v string) {
	o.Plugin = v
}

// GetTenantId returns the TenantId field value
func (o *Aggregator) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *Aggregator) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *Aggregator) SetTenantId(v string) {
	o.TenantId = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Aggregator) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregator) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Aggregator) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Aggregator) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Aggregator) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregator) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Aggregator) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Aggregator) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *Aggregator) GetUpdatedBy() User {
	if o == nil || o.UpdatedBy == nil {
		var ret User
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregator) GetUpdatedByOk() (*User, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *Aggregator) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given User and assigns it to the UpdatedBy field.
func (o *Aggregator) SetUpdatedBy(v User) {
	o.UpdatedBy = &v
}

// GetUpdatedById returns the UpdatedById field value
func (o *Aggregator) GetUpdatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedById
}

// GetUpdatedByIdOk returns a tuple with the UpdatedById field value
// and a boolean to check if the value has been set.
func (o *Aggregator) GetUpdatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedById, true
}

// SetUpdatedById sets field value
func (o *Aggregator) SetUpdatedById(v string) {
	o.UpdatedById = v
}

// GetVersionId returns the VersionId field value
func (o *Aggregator) GetVersionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *Aggregator) GetVersionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *Aggregator) SetVersionId(v int32) {
	o.VersionId = v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *Aggregator) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Aggregator) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *Aggregator) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *Aggregator) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o Aggregator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if true {
		toSerialize["connections"] = o.Connections
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
	if o.DeletedAt != nil {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if o.DeletedBy != nil {
		toSerialize["deleted_by"] = o.DeletedBy
	}
	if true {
		toSerialize["deleted_by_id"] = o.DeletedById
	}
	if true {
		toSerialize["handle"] = o.Handle
	}
	if o.HandleDynamic != nil {
		toSerialize["handle_dynamic"] = o.HandleDynamic
	}
	if o.HandleMode != nil {
		toSerialize["handle_mode"] = o.HandleMode
	}
	if o.HandlePrefix != nil {
		toSerialize["handle_prefix"] = o.HandlePrefix
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
		toSerialize["plugin"] = o.Plugin
	}
	if true {
		toSerialize["tenant_id"] = o.TenantId
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

type NullableAggregator struct {
	value *Aggregator
	isSet bool
}

func (v NullableAggregator) Get() *Aggregator {
	return v.value
}

func (v *NullableAggregator) Set(val *Aggregator) {
	v.value = val
	v.isSet = true
}

func (v NullableAggregator) IsSet() bool {
	return v.isSet
}

func (v *NullableAggregator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAggregator(val *Aggregator) *NullableAggregator {
	return &NullableAggregator{value: val, isSet: true}
}

func (v NullableAggregator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAggregator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}