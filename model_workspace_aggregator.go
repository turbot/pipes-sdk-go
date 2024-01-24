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

// WorkspaceAggregator struct for WorkspaceAggregator
type WorkspaceAggregator struct {
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
	// The unique identifier for the aggregator.
	Id string `json:"id"`
	// The unique identifier for an identity where the aggregator has been created.
	IdentityId string `json:"identity_id"`
	// The plugin name for the aggregator.
	Plugin string `json:"plugin"`
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
	WorkspaceId string `json:"workspace_id"`
}

// NewWorkspaceAggregator instantiates a new WorkspaceAggregator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceAggregator(connections []string, createdAt string, createdById string, deletedById string, handle string, id string, identityId string, plugin string, updatedById string, versionId int32, workspaceId string) *WorkspaceAggregator {
	this := WorkspaceAggregator{}
	this.Connections = connections
	this.CreatedAt = createdAt
	this.CreatedById = createdById
	this.DeletedById = deletedById
	this.Handle = handle
	this.Id = id
	this.IdentityId = identityId
	this.Plugin = plugin
	this.UpdatedById = updatedById
	this.VersionId = versionId
	this.WorkspaceId = workspaceId
	return &this
}

// NewWorkspaceAggregatorWithDefaults instantiates a new WorkspaceAggregator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceAggregatorWithDefaults() *WorkspaceAggregator {
	this := WorkspaceAggregator{}
	return &this
}

// GetConnections returns the Connections field value
func (o *WorkspaceAggregator) GetConnections() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
func (o *WorkspaceAggregator) GetConnectionsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Connections, true
}

// SetConnections sets field value
func (o *WorkspaceAggregator) SetConnections(v []string) {
	o.Connections = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *WorkspaceAggregator) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WorkspaceAggregator) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WorkspaceAggregator) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *WorkspaceAggregator) GetCreatedBy() User {
	if o == nil || o.CreatedBy == nil {
		var ret User
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceAggregator) GetCreatedByOk() (*User, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *WorkspaceAggregator) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given User and assigns it to the CreatedBy field.
func (o *WorkspaceAggregator) SetCreatedBy(v User) {
	o.CreatedBy = &v
}

// GetCreatedById returns the CreatedById field value
func (o *WorkspaceAggregator) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *WorkspaceAggregator) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *WorkspaceAggregator) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *WorkspaceAggregator) GetDeletedAt() string {
	if o == nil || o.DeletedAt == nil {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceAggregator) GetDeletedAtOk() (*string, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *WorkspaceAggregator) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *WorkspaceAggregator) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetDeletedBy returns the DeletedBy field value if set, zero value otherwise.
func (o *WorkspaceAggregator) GetDeletedBy() User {
	if o == nil || o.DeletedBy == nil {
		var ret User
		return ret
	}
	return *o.DeletedBy
}

// GetDeletedByOk returns a tuple with the DeletedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceAggregator) GetDeletedByOk() (*User, bool) {
	if o == nil || o.DeletedBy == nil {
		return nil, false
	}
	return o.DeletedBy, true
}

// HasDeletedBy returns a boolean if a field has been set.
func (o *WorkspaceAggregator) HasDeletedBy() bool {
	if o != nil && o.DeletedBy != nil {
		return true
	}

	return false
}

// SetDeletedBy gets a reference to the given User and assigns it to the DeletedBy field.
func (o *WorkspaceAggregator) SetDeletedBy(v User) {
	o.DeletedBy = &v
}

// GetDeletedById returns the DeletedById field value
func (o *WorkspaceAggregator) GetDeletedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeletedById
}

// GetDeletedByIdOk returns a tuple with the DeletedById field value
// and a boolean to check if the value has been set.
func (o *WorkspaceAggregator) GetDeletedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedById, true
}

// SetDeletedById sets field value
func (o *WorkspaceAggregator) SetDeletedById(v string) {
	o.DeletedById = v
}

// GetHandle returns the Handle field value
func (o *WorkspaceAggregator) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *WorkspaceAggregator) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value
func (o *WorkspaceAggregator) SetHandle(v string) {
	o.Handle = v
}

// GetId returns the Id field value
func (o *WorkspaceAggregator) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkspaceAggregator) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkspaceAggregator) SetId(v string) {
	o.Id = v
}

// GetIdentityId returns the IdentityId field value
func (o *WorkspaceAggregator) GetIdentityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value
// and a boolean to check if the value has been set.
func (o *WorkspaceAggregator) GetIdentityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityId, true
}

// SetIdentityId sets field value
func (o *WorkspaceAggregator) SetIdentityId(v string) {
	o.IdentityId = v
}

// GetPlugin returns the Plugin field value
func (o *WorkspaceAggregator) GetPlugin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plugin
}

// GetPluginOk returns a tuple with the Plugin field value
// and a boolean to check if the value has been set.
func (o *WorkspaceAggregator) GetPluginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plugin, true
}

// SetPlugin sets field value
func (o *WorkspaceAggregator) SetPlugin(v string) {
	o.Plugin = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkspaceAggregator) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceAggregator) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkspaceAggregator) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WorkspaceAggregator) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *WorkspaceAggregator) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceAggregator) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *WorkspaceAggregator) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *WorkspaceAggregator) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *WorkspaceAggregator) GetUpdatedBy() User {
	if o == nil || o.UpdatedBy == nil {
		var ret User
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceAggregator) GetUpdatedByOk() (*User, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *WorkspaceAggregator) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given User and assigns it to the UpdatedBy field.
func (o *WorkspaceAggregator) SetUpdatedBy(v User) {
	o.UpdatedBy = &v
}

// GetUpdatedById returns the UpdatedById field value
func (o *WorkspaceAggregator) GetUpdatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedById
}

// GetUpdatedByIdOk returns a tuple with the UpdatedById field value
// and a boolean to check if the value has been set.
func (o *WorkspaceAggregator) GetUpdatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedById, true
}

// SetUpdatedById sets field value
func (o *WorkspaceAggregator) SetUpdatedById(v string) {
	o.UpdatedById = v
}

// GetVersionId returns the VersionId field value
func (o *WorkspaceAggregator) GetVersionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *WorkspaceAggregator) GetVersionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *WorkspaceAggregator) SetVersionId(v int32) {
	o.VersionId = v
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *WorkspaceAggregator) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *WorkspaceAggregator) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *WorkspaceAggregator) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

func (o WorkspaceAggregator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["identity_id"] = o.IdentityId
	}
	if true {
		toSerialize["plugin"] = o.Plugin
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
	if true {
		toSerialize["workspace_id"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableWorkspaceAggregator struct {
	value *WorkspaceAggregator
	isSet bool
}

func (v NullableWorkspaceAggregator) Get() *WorkspaceAggregator {
	return v.value
}

func (v *NullableWorkspaceAggregator) Set(val *WorkspaceAggregator) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceAggregator) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceAggregator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceAggregator(val *WorkspaceAggregator) *NullableWorkspaceAggregator {
	return &NullableWorkspaceAggregator{value: val, isSet: true}
}

func (v NullableWorkspaceAggregator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceAggregator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
