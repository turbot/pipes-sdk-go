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

// WorkspaceSnapshotData struct for WorkspaceSnapshotData
type WorkspaceSnapshotData struct {
	// The time the dashboard execution ended.
	EndTime string `json:"end_time"`
	// Any inputs and their values used in this snapshot.
	Inputs *map[string]interface{} `json:"inputs,omitempty"`
	// The dashboard layout tree for this snapshot.
	Layout WorkspaceSnapshotDataLayout `json:"layout"`
	// Any additional metadata for this snapshot.
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	// The map of panels and their data for this snapshot.
	Panels map[string]interface{} `json:"panels"`
	// The schema version of this snapshot.
	SchemaVersion string `json:"schema_version"`
	// The time the dashboard execution started.
	StartTime string `json:"start_time"`
	// Any mod variables in scope for this snapshot.
	Variables *map[string]interface{} `json:"variables,omitempty"`
}

// NewWorkspaceSnapshotData instantiates a new WorkspaceSnapshotData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceSnapshotData(endTime string, layout WorkspaceSnapshotDataLayout, panels map[string]interface{}, schemaVersion string, startTime string) *WorkspaceSnapshotData {
	this := WorkspaceSnapshotData{}
	this.EndTime = endTime
	this.Layout = layout
	this.Panels = panels
	this.SchemaVersion = schemaVersion
	this.StartTime = startTime
	return &this
}

// NewWorkspaceSnapshotDataWithDefaults instantiates a new WorkspaceSnapshotData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceSnapshotDataWithDefaults() *WorkspaceSnapshotData {
	this := WorkspaceSnapshotData{}
	return &this
}

// GetEndTime returns the EndTime field value
func (o *WorkspaceSnapshotData) GetEndTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *WorkspaceSnapshotData) GetEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *WorkspaceSnapshotData) SetEndTime(v string) {
	o.EndTime = v
}

// GetInputs returns the Inputs field value if set, zero value otherwise.
func (o *WorkspaceSnapshotData) GetInputs() map[string]interface{} {
	if o == nil || o.Inputs == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceSnapshotData) GetInputsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Inputs == nil {
		return nil, false
	}
	return o.Inputs, true
}

// HasInputs returns a boolean if a field has been set.
func (o *WorkspaceSnapshotData) HasInputs() bool {
	if o != nil && o.Inputs != nil {
		return true
	}

	return false
}

// SetInputs gets a reference to the given map[string]interface{} and assigns it to the Inputs field.
func (o *WorkspaceSnapshotData) SetInputs(v map[string]interface{}) {
	o.Inputs = &v
}

// GetLayout returns the Layout field value
func (o *WorkspaceSnapshotData) GetLayout() WorkspaceSnapshotDataLayout {
	if o == nil {
		var ret WorkspaceSnapshotDataLayout
		return ret
	}

	return o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value
// and a boolean to check if the value has been set.
func (o *WorkspaceSnapshotData) GetLayoutOk() (*WorkspaceSnapshotDataLayout, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Layout, true
}

// SetLayout sets field value
func (o *WorkspaceSnapshotData) SetLayout(v WorkspaceSnapshotDataLayout) {
	o.Layout = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *WorkspaceSnapshotData) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceSnapshotData) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *WorkspaceSnapshotData) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *WorkspaceSnapshotData) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetPanels returns the Panels field value
func (o *WorkspaceSnapshotData) GetPanels() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Panels
}

// GetPanelsOk returns a tuple with the Panels field value
// and a boolean to check if the value has been set.
func (o *WorkspaceSnapshotData) GetPanelsOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Panels, true
}

// SetPanels sets field value
func (o *WorkspaceSnapshotData) SetPanels(v map[string]interface{}) {
	o.Panels = v
}

// GetSchemaVersion returns the SchemaVersion field value
func (o *WorkspaceSnapshotData) GetSchemaVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value
// and a boolean to check if the value has been set.
func (o *WorkspaceSnapshotData) GetSchemaVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaVersion, true
}

// SetSchemaVersion sets field value
func (o *WorkspaceSnapshotData) SetSchemaVersion(v string) {
	o.SchemaVersion = v
}

// GetStartTime returns the StartTime field value
func (o *WorkspaceSnapshotData) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *WorkspaceSnapshotData) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *WorkspaceSnapshotData) SetStartTime(v string) {
	o.StartTime = v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *WorkspaceSnapshotData) GetVariables() map[string]interface{} {
	if o == nil || o.Variables == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceSnapshotData) GetVariablesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Variables == nil {
		return nil, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *WorkspaceSnapshotData) HasVariables() bool {
	if o != nil && o.Variables != nil {
		return true
	}

	return false
}

// SetVariables gets a reference to the given map[string]interface{} and assigns it to the Variables field.
func (o *WorkspaceSnapshotData) SetVariables(v map[string]interface{}) {
	o.Variables = &v
}

func (o WorkspaceSnapshotData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["end_time"] = o.EndTime
	}
	if o.Inputs != nil {
		toSerialize["inputs"] = o.Inputs
	}
	if true {
		toSerialize["layout"] = o.Layout
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["panels"] = o.Panels
	}
	if true {
		toSerialize["schema_version"] = o.SchemaVersion
	}
	if true {
		toSerialize["start_time"] = o.StartTime
	}
	if o.Variables != nil {
		toSerialize["variables"] = o.Variables
	}
	return json.Marshal(toSerialize)
}

type NullableWorkspaceSnapshotData struct {
	value *WorkspaceSnapshotData
	isSet bool
}

func (v NullableWorkspaceSnapshotData) Get() *WorkspaceSnapshotData {
	return v.value
}

func (v *NullableWorkspaceSnapshotData) Set(val *WorkspaceSnapshotData) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceSnapshotData) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceSnapshotData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceSnapshotData(val *WorkspaceSnapshotData) *NullableWorkspaceSnapshotData {
	return &NullableWorkspaceSnapshotData{value: val, isSet: true}
}

func (v NullableWorkspaceSnapshotData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceSnapshotData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
