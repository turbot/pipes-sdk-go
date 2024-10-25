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

// ModTriggerInfo struct for ModTriggerInfo
type ModTriggerInfo struct {
	// The time of creation in ISO 8601 UTC.
	CreatedAt string `json:"created_at"`
	// User information for the user who created this.
	CreatedBy *User `json:"created_by,omitempty"`
	// The ID of the user that created this.
	CreatedById string                  `json:"created_by_id"`
	Description *string                 `json:"description,omitempty"`
	Id          *string                 `json:"id,omitempty"`
	ModId       *string                 `json:"mod_id,omitempty"`
	Name        *string                 `json:"name,omitempty"`
	Pipelines   *map[string]interface{} `json:"pipelines,omitempty"`
	Query       *string                 `json:"query,omitempty"`
	Schedule    *string                 `json:"schedule,omitempty"`
	State       *ModTriggerState        `json:"state,omitempty"`
	Tags        *map[string]interface{} `json:"tags,omitempty"`
	Title       *string                 `json:"title,omitempty"`
	Type        *string                 `json:"type,omitempty"`
	// The time of the last update in ISO 8601 UTC.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// User information for the last user to update this.
	UpdatedBy *User `json:"updated_by,omitempty"`
	// The ID of the user that performed the last update.
	UpdatedById string `json:"updated_by_id"`
	// The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item.
	VersionId int32 `json:"version_id"`
}

// NewModTriggerInfo instantiates a new ModTriggerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModTriggerInfo(createdAt string, createdById string, updatedById string, versionId int32) *ModTriggerInfo {
	this := ModTriggerInfo{}
	this.CreatedAt = createdAt
	this.CreatedById = createdById
	this.UpdatedById = updatedById
	this.VersionId = versionId
	return &this
}

// NewModTriggerInfoWithDefaults instantiates a new ModTriggerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModTriggerInfoWithDefaults() *ModTriggerInfo {
	this := ModTriggerInfo{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *ModTriggerInfo) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ModTriggerInfo) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ModTriggerInfo) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ModTriggerInfo) GetCreatedBy() User {
	if o == nil || o.CreatedBy == nil {
		var ret User
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModTriggerInfo) GetCreatedByOk() (*User, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ModTriggerInfo) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given User and assigns it to the CreatedBy field.
func (o *ModTriggerInfo) SetCreatedBy(v User) {
	o.CreatedBy = &v
}

// GetCreatedById returns the CreatedById field value
func (o *ModTriggerInfo) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *ModTriggerInfo) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *ModTriggerInfo) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModTriggerInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModTriggerInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModTriggerInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModTriggerInfo) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModTriggerInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModTriggerInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModTriggerInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModTriggerInfo) SetId(v string) {
	o.Id = &v
}

// GetModId returns the ModId field value if set, zero value otherwise.
func (o *ModTriggerInfo) GetModId() string {
	if o == nil || o.ModId == nil {
		var ret string
		return ret
	}
	return *o.ModId
}

// GetModIdOk returns a tuple with the ModId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModTriggerInfo) GetModIdOk() (*string, bool) {
	if o == nil || o.ModId == nil {
		return nil, false
	}
	return o.ModId, true
}

// HasModId returns a boolean if a field has been set.
func (o *ModTriggerInfo) HasModId() bool {
	if o != nil && o.ModId != nil {
		return true
	}

	return false
}

// SetModId gets a reference to the given string and assigns it to the ModId field.
func (o *ModTriggerInfo) SetModId(v string) {
	o.ModId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModTriggerInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModTriggerInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModTriggerInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModTriggerInfo) SetName(v string) {
	o.Name = &v
}

// GetPipelines returns the Pipelines field value if set, zero value otherwise.
func (o *ModTriggerInfo) GetPipelines() map[string]interface{} {
	if o == nil || o.Pipelines == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Pipelines
}

// GetPipelinesOk returns a tuple with the Pipelines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModTriggerInfo) GetPipelinesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Pipelines == nil {
		return nil, false
	}
	return o.Pipelines, true
}

// HasPipelines returns a boolean if a field has been set.
func (o *ModTriggerInfo) HasPipelines() bool {
	if o != nil && o.Pipelines != nil {
		return true
	}

	return false
}

// SetPipelines gets a reference to the given map[string]interface{} and assigns it to the Pipelines field.
func (o *ModTriggerInfo) SetPipelines(v map[string]interface{}) {
	o.Pipelines = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ModTriggerInfo) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModTriggerInfo) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ModTriggerInfo) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *ModTriggerInfo) SetQuery(v string) {
	o.Query = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *ModTriggerInfo) GetSchedule() string {
	if o == nil || o.Schedule == nil {
		var ret string
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModTriggerInfo) GetScheduleOk() (*string, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *ModTriggerInfo) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given string and assigns it to the Schedule field.
func (o *ModTriggerInfo) SetSchedule(v string) {
	o.Schedule = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ModTriggerInfo) GetState() ModTriggerState {
	if o == nil || o.State == nil {
		var ret ModTriggerState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModTriggerInfo) GetStateOk() (*ModTriggerState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ModTriggerInfo) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given ModTriggerState and assigns it to the State field.
func (o *ModTriggerInfo) SetState(v ModTriggerState) {
	o.State = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ModTriggerInfo) GetTags() map[string]interface{} {
	if o == nil || o.Tags == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModTriggerInfo) GetTagsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ModTriggerInfo) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *ModTriggerInfo) SetTags(v map[string]interface{}) {
	o.Tags = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ModTriggerInfo) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModTriggerInfo) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ModTriggerInfo) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ModTriggerInfo) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ModTriggerInfo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModTriggerInfo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ModTriggerInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ModTriggerInfo) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ModTriggerInfo) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModTriggerInfo) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ModTriggerInfo) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ModTriggerInfo) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *ModTriggerInfo) GetUpdatedBy() User {
	if o == nil || o.UpdatedBy == nil {
		var ret User
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModTriggerInfo) GetUpdatedByOk() (*User, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *ModTriggerInfo) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given User and assigns it to the UpdatedBy field.
func (o *ModTriggerInfo) SetUpdatedBy(v User) {
	o.UpdatedBy = &v
}

// GetUpdatedById returns the UpdatedById field value
func (o *ModTriggerInfo) GetUpdatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedById
}

// GetUpdatedByIdOk returns a tuple with the UpdatedById field value
// and a boolean to check if the value has been set.
func (o *ModTriggerInfo) GetUpdatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedById, true
}

// SetUpdatedById sets field value
func (o *ModTriggerInfo) SetUpdatedById(v string) {
	o.UpdatedById = v
}

// GetVersionId returns the VersionId field value
func (o *ModTriggerInfo) GetVersionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *ModTriggerInfo) GetVersionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *ModTriggerInfo) SetVersionId(v int32) {
	o.VersionId = v
}

func (o ModTriggerInfo) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ModId != nil {
		toSerialize["mod_id"] = o.ModId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Pipelines != nil {
		toSerialize["pipelines"] = o.Pipelines
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
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
	return json.Marshal(toSerialize)
}

type NullableModTriggerInfo struct {
	value *ModTriggerInfo
	isSet bool
}

func (v NullableModTriggerInfo) Get() *ModTriggerInfo {
	return v.value
}

func (v *NullableModTriggerInfo) Set(val *ModTriggerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModTriggerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModTriggerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModTriggerInfo(val *ModTriggerInfo) *NullableModTriggerInfo {
	return &NullableModTriggerInfo{value: val, isSet: true}
}

func (v NullableModTriggerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModTriggerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
