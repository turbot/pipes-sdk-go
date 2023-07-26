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

// WorkspaceSnapshotDataLayout struct for WorkspaceSnapshotDataLayout
type WorkspaceSnapshotDataLayout struct {
	Children  *[]WorkspaceSnapshotDataLayout `json:"children,omitempty"`
	Name      string                         `json:"name"`
	PanelType string                         `json:"panel_type"`
}

// NewWorkspaceSnapshotDataLayout instantiates a new WorkspaceSnapshotDataLayout object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceSnapshotDataLayout(name string, panelType string) *WorkspaceSnapshotDataLayout {
	this := WorkspaceSnapshotDataLayout{}
	this.Name = name
	this.PanelType = panelType
	return &this
}

// NewWorkspaceSnapshotDataLayoutWithDefaults instantiates a new WorkspaceSnapshotDataLayout object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceSnapshotDataLayoutWithDefaults() *WorkspaceSnapshotDataLayout {
	this := WorkspaceSnapshotDataLayout{}
	return &this
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *WorkspaceSnapshotDataLayout) GetChildren() []WorkspaceSnapshotDataLayout {
	if o == nil || o.Children == nil {
		var ret []WorkspaceSnapshotDataLayout
		return ret
	}
	return *o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceSnapshotDataLayout) GetChildrenOk() (*[]WorkspaceSnapshotDataLayout, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *WorkspaceSnapshotDataLayout) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []WorkspaceSnapshotDataLayout and assigns it to the Children field.
func (o *WorkspaceSnapshotDataLayout) SetChildren(v []WorkspaceSnapshotDataLayout) {
	o.Children = &v
}

// GetName returns the Name field value
func (o *WorkspaceSnapshotDataLayout) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkspaceSnapshotDataLayout) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkspaceSnapshotDataLayout) SetName(v string) {
	o.Name = v
}

// GetPanelType returns the PanelType field value
func (o *WorkspaceSnapshotDataLayout) GetPanelType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PanelType
}

// GetPanelTypeOk returns a tuple with the PanelType field value
// and a boolean to check if the value has been set.
func (o *WorkspaceSnapshotDataLayout) GetPanelTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PanelType, true
}

// SetPanelType sets field value
func (o *WorkspaceSnapshotDataLayout) SetPanelType(v string) {
	o.PanelType = v
}

func (o WorkspaceSnapshotDataLayout) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["panel_type"] = o.PanelType
	}
	return json.Marshal(toSerialize)
}

type NullableWorkspaceSnapshotDataLayout struct {
	value *WorkspaceSnapshotDataLayout
	isSet bool
}

func (v NullableWorkspaceSnapshotDataLayout) Get() *WorkspaceSnapshotDataLayout {
	return v.value
}

func (v *NullableWorkspaceSnapshotDataLayout) Set(val *WorkspaceSnapshotDataLayout) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceSnapshotDataLayout) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceSnapshotDataLayout) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceSnapshotDataLayout(val *WorkspaceSnapshotDataLayout) *NullableWorkspaceSnapshotDataLayout {
	return &NullableWorkspaceSnapshotDataLayout{value: val, isSet: true}
}

func (v NullableWorkspaceSnapshotDataLayout) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceSnapshotDataLayout) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}