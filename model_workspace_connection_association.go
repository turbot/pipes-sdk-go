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

// WorkspaceConnectionAssociation struct for WorkspaceConnectionAssociation
type WorkspaceConnectionAssociation struct {
	// Additional information about the connection.
	Connection *Connection `json:"connection,omitempty"`
	// The unique identifier for the connection.
	ConnectionId string `json:"connection_id"`
	// The time of creation in ISO 8601 UTC.
	CreatedAt string `json:"created_at"`
	// User information for the user who created this.
	CreatedBy *User `json:"created_by,omitempty"`
	// The ID of the user that created this.
	CreatedById string `json:"created_by_id"`
	// The unique identifier for the workspace connection association.
	Id string `json:"id"`
	// The identity ID where the association exists.
	IdentityId string `json:"identity_id"`
	// The time of the last update in ISO 8601 UTC.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// User information for the last user to update this.
	UpdatedBy *User `json:"updated_by,omitempty"`
	// The ID of the user that performed the last update.
	UpdatedById string `json:"updated_by_id"`
	// The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item.
	VersionId int32 `json:"version_id"`
	// The unique identifier for the wokspace.
	WorkspaceId string `json:"workspace_id"`
}

// NewWorkspaceConnectionAssociation instantiates a new WorkspaceConnectionAssociation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceConnectionAssociation(connectionId string, createdAt string, createdById string, id string, identityId string, updatedById string, versionId int32, workspaceId string) *WorkspaceConnectionAssociation {
	this := WorkspaceConnectionAssociation{}
	this.ConnectionId = connectionId
	this.CreatedAt = createdAt
	this.CreatedById = createdById
	this.Id = id
	this.IdentityId = identityId
	this.UpdatedById = updatedById
	this.VersionId = versionId
	this.WorkspaceId = workspaceId
	return &this
}

// NewWorkspaceConnectionAssociationWithDefaults instantiates a new WorkspaceConnectionAssociation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceConnectionAssociationWithDefaults() *WorkspaceConnectionAssociation {
	this := WorkspaceConnectionAssociation{}
	return &this
}

// GetConnection returns the Connection field value if set, zero value otherwise.
func (o *WorkspaceConnectionAssociation) GetConnection() Connection {
	if o == nil || o.Connection == nil {
		var ret Connection
		return ret
	}
	return *o.Connection
}

// GetConnectionOk returns a tuple with the Connection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnectionAssociation) GetConnectionOk() (*Connection, bool) {
	if o == nil || o.Connection == nil {
		return nil, false
	}
	return o.Connection, true
}

// HasConnection returns a boolean if a field has been set.
func (o *WorkspaceConnectionAssociation) HasConnection() bool {
	if o != nil && o.Connection != nil {
		return true
	}

	return false
}

// SetConnection gets a reference to the given Connection and assigns it to the Connection field.
func (o *WorkspaceConnectionAssociation) SetConnection(v Connection) {
	o.Connection = &v
}

// GetConnectionId returns the ConnectionId field value
func (o *WorkspaceConnectionAssociation) GetConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConnectionAssociation) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionId, true
}

// SetConnectionId sets field value
func (o *WorkspaceConnectionAssociation) SetConnectionId(v string) {
	o.ConnectionId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *WorkspaceConnectionAssociation) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConnectionAssociation) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WorkspaceConnectionAssociation) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *WorkspaceConnectionAssociation) GetCreatedBy() User {
	if o == nil || o.CreatedBy == nil {
		var ret User
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnectionAssociation) GetCreatedByOk() (*User, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *WorkspaceConnectionAssociation) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given User and assigns it to the CreatedBy field.
func (o *WorkspaceConnectionAssociation) SetCreatedBy(v User) {
	o.CreatedBy = &v
}

// GetCreatedById returns the CreatedById field value
func (o *WorkspaceConnectionAssociation) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConnectionAssociation) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *WorkspaceConnectionAssociation) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetId returns the Id field value
func (o *WorkspaceConnectionAssociation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConnectionAssociation) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WorkspaceConnectionAssociation) SetId(v string) {
	o.Id = v
}

// GetIdentityId returns the IdentityId field value
func (o *WorkspaceConnectionAssociation) GetIdentityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConnectionAssociation) GetIdentityIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdentityId, true
}

// SetIdentityId sets field value
func (o *WorkspaceConnectionAssociation) SetIdentityId(v string) {
	o.IdentityId = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *WorkspaceConnectionAssociation) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnectionAssociation) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *WorkspaceConnectionAssociation) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *WorkspaceConnectionAssociation) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *WorkspaceConnectionAssociation) GetUpdatedBy() User {
	if o == nil || o.UpdatedBy == nil {
		var ret User
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceConnectionAssociation) GetUpdatedByOk() (*User, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *WorkspaceConnectionAssociation) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given User and assigns it to the UpdatedBy field.
func (o *WorkspaceConnectionAssociation) SetUpdatedBy(v User) {
	o.UpdatedBy = &v
}

// GetUpdatedById returns the UpdatedById field value
func (o *WorkspaceConnectionAssociation) GetUpdatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedById
}

// GetUpdatedByIdOk returns a tuple with the UpdatedById field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConnectionAssociation) GetUpdatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedById, true
}

// SetUpdatedById sets field value
func (o *WorkspaceConnectionAssociation) SetUpdatedById(v string) {
	o.UpdatedById = v
}

// GetVersionId returns the VersionId field value
func (o *WorkspaceConnectionAssociation) GetVersionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConnectionAssociation) GetVersionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *WorkspaceConnectionAssociation) SetVersionId(v int32) {
	o.VersionId = v
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *WorkspaceConnectionAssociation) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *WorkspaceConnectionAssociation) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *WorkspaceConnectionAssociation) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

func (o WorkspaceConnectionAssociation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Connection != nil {
		toSerialize["connection"] = o.Connection
	}
	if true {
		toSerialize["connection_id"] = o.ConnectionId
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
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["identity_id"] = o.IdentityId
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

type NullableWorkspaceConnectionAssociation struct {
	value *WorkspaceConnectionAssociation
	isSet bool
}

func (v NullableWorkspaceConnectionAssociation) Get() *WorkspaceConnectionAssociation {
	return v.value
}

func (v *NullableWorkspaceConnectionAssociation) Set(val *WorkspaceConnectionAssociation) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceConnectionAssociation) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceConnectionAssociation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceConnectionAssociation(val *WorkspaceConnectionAssociation) *NullableWorkspaceConnectionAssociation {
	return &NullableWorkspaceConnectionAssociation{value: val, isSet: true}
}

func (v NullableWorkspaceConnectionAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceConnectionAssociation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
