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

// SpProcess struct for SpProcess
type SpProcess struct {
	// The time of creation in ISO 8601 UTC.
	CreatedAt string `json:"created_at"`
	CreatedBy *User  `json:"created_by,omitempty"`
	// The ID of the user that created this.
	CreatedById string `json:"created_by_id"`
	// The unique identifier of the process.
	Id string `json:"id"`
	// The unique identifier of the identity for which the process is created.
	IdentityId *string   `json:"identity_id,omitempty"`
	Pipeline   *Pipeline `json:"pipeline,omitempty"`
	// The unique identifier of the pipeline for which the process is created.
	PipelineId *string `json:"pipeline_id,omitempty"`
	// The state of the process.
	State *string `json:"state,omitempty"`
	// The optional reason why the process is in its current state.
	StateReason *string `json:"state_reason,omitempty"`
	// The type of the process, generally denotes the activity performed e.g. workspace.create, pipeline.execute, pipeline.command.run.
	Type string `json:"type"`
	// The time of the last update in ISO 8601 UTC.
	UpdatedAt string `json:"updated_at"`
	UpdatedBy *User  `json:"updated_by,omitempty"`
	// The ID of the user that performed the last update.
	UpdatedById string `json:"updated_by_id"`
	// The usage information for this process.
	Usage *map[string]interface{} `json:"usage,omitempty"`
	// The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item.
	VersionId int32 `json:"version_id"`
	// The unique identifier of the workspace for which the process is created.
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

// NewSpProcess instantiates a new SpProcess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpProcess(createdAt string, createdById string, id string, type_ string, updatedAt string, updatedById string, versionId int32) *SpProcess {
	this := SpProcess{}
	this.CreatedAt = createdAt
	this.CreatedById = createdById
	this.Id = id
	this.Type = type_
	this.UpdatedAt = updatedAt
	this.UpdatedById = updatedById
	this.VersionId = versionId
	return &this
}

// NewSpProcessWithDefaults instantiates a new SpProcess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpProcessWithDefaults() *SpProcess {
	this := SpProcess{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *SpProcess) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SpProcess) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SpProcess) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *SpProcess) GetCreatedBy() User {
	if o == nil || o.CreatedBy == nil {
		var ret User
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpProcess) GetCreatedByOk() (*User, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *SpProcess) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given User and assigns it to the CreatedBy field.
func (o *SpProcess) SetCreatedBy(v User) {
	o.CreatedBy = &v
}

// GetCreatedById returns the CreatedById field value
func (o *SpProcess) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *SpProcess) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *SpProcess) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetId returns the Id field value
func (o *SpProcess) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SpProcess) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SpProcess) SetId(v string) {
	o.Id = v
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *SpProcess) GetIdentityId() string {
	if o == nil || o.IdentityId == nil {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpProcess) GetIdentityIdOk() (*string, bool) {
	if o == nil || o.IdentityId == nil {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *SpProcess) HasIdentityId() bool {
	if o != nil && o.IdentityId != nil {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *SpProcess) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetPipeline returns the Pipeline field value if set, zero value otherwise.
func (o *SpProcess) GetPipeline() Pipeline {
	if o == nil || o.Pipeline == nil {
		var ret Pipeline
		return ret
	}
	return *o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpProcess) GetPipelineOk() (*Pipeline, bool) {
	if o == nil || o.Pipeline == nil {
		return nil, false
	}
	return o.Pipeline, true
}

// HasPipeline returns a boolean if a field has been set.
func (o *SpProcess) HasPipeline() bool {
	if o != nil && o.Pipeline != nil {
		return true
	}

	return false
}

// SetPipeline gets a reference to the given Pipeline and assigns it to the Pipeline field.
func (o *SpProcess) SetPipeline(v Pipeline) {
	o.Pipeline = &v
}

// GetPipelineId returns the PipelineId field value if set, zero value otherwise.
func (o *SpProcess) GetPipelineId() string {
	if o == nil || o.PipelineId == nil {
		var ret string
		return ret
	}
	return *o.PipelineId
}

// GetPipelineIdOk returns a tuple with the PipelineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpProcess) GetPipelineIdOk() (*string, bool) {
	if o == nil || o.PipelineId == nil {
		return nil, false
	}
	return o.PipelineId, true
}

// HasPipelineId returns a boolean if a field has been set.
func (o *SpProcess) HasPipelineId() bool {
	if o != nil && o.PipelineId != nil {
		return true
	}

	return false
}

// SetPipelineId gets a reference to the given string and assigns it to the PipelineId field.
func (o *SpProcess) SetPipelineId(v string) {
	o.PipelineId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SpProcess) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpProcess) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SpProcess) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *SpProcess) SetState(v string) {
	o.State = &v
}

// GetStateReason returns the StateReason field value if set, zero value otherwise.
func (o *SpProcess) GetStateReason() string {
	if o == nil || o.StateReason == nil {
		var ret string
		return ret
	}
	return *o.StateReason
}

// GetStateReasonOk returns a tuple with the StateReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpProcess) GetStateReasonOk() (*string, bool) {
	if o == nil || o.StateReason == nil {
		return nil, false
	}
	return o.StateReason, true
}

// HasStateReason returns a boolean if a field has been set.
func (o *SpProcess) HasStateReason() bool {
	if o != nil && o.StateReason != nil {
		return true
	}

	return false
}

// SetStateReason gets a reference to the given string and assigns it to the StateReason field.
func (o *SpProcess) SetStateReason(v string) {
	o.StateReason = &v
}

// GetType returns the Type field value
func (o *SpProcess) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SpProcess) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SpProcess) SetType(v string) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *SpProcess) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *SpProcess) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *SpProcess) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *SpProcess) GetUpdatedBy() User {
	if o == nil || o.UpdatedBy == nil {
		var ret User
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpProcess) GetUpdatedByOk() (*User, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *SpProcess) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given User and assigns it to the UpdatedBy field.
func (o *SpProcess) SetUpdatedBy(v User) {
	o.UpdatedBy = &v
}

// GetUpdatedById returns the UpdatedById field value
func (o *SpProcess) GetUpdatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedById
}

// GetUpdatedByIdOk returns a tuple with the UpdatedById field value
// and a boolean to check if the value has been set.
func (o *SpProcess) GetUpdatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedById, true
}

// SetUpdatedById sets field value
func (o *SpProcess) SetUpdatedById(v string) {
	o.UpdatedById = v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *SpProcess) GetUsage() map[string]interface{} {
	if o == nil || o.Usage == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpProcess) GetUsageOk() (*map[string]interface{}, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *SpProcess) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given map[string]interface{} and assigns it to the Usage field.
func (o *SpProcess) SetUsage(v map[string]interface{}) {
	o.Usage = &v
}

// GetVersionId returns the VersionId field value
func (o *SpProcess) GetVersionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *SpProcess) GetVersionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *SpProcess) SetVersionId(v int32) {
	o.VersionId = v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *SpProcess) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpProcess) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *SpProcess) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *SpProcess) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o SpProcess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.IdentityId != nil {
		toSerialize["identity_id"] = o.IdentityId
	}
	if o.Pipeline != nil {
		toSerialize["pipeline"] = o.Pipeline
	}
	if o.PipelineId != nil {
		toSerialize["pipeline_id"] = o.PipelineId
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.StateReason != nil {
		toSerialize["state_reason"] = o.StateReason
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.UpdatedBy != nil {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	if true {
		toSerialize["updated_by_id"] = o.UpdatedById
	}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	if true {
		toSerialize["version_id"] = o.VersionId
	}
	if o.WorkspaceId != nil {
		toSerialize["workspace_id"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableSpProcess struct {
	value *SpProcess
	isSet bool
}

func (v NullableSpProcess) Get() *SpProcess {
	return v.value
}

func (v *NullableSpProcess) Set(val *SpProcess) {
	v.value = val
	v.isSet = true
}

func (v NullableSpProcess) IsSet() bool {
	return v.isSet
}

func (v *NullableSpProcess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpProcess(val *SpProcess) *NullableSpProcess {
	return &NullableSpProcess{value: val, isSet: true}
}

func (v NullableSpProcess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpProcess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
