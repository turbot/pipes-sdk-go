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

// WorkspaceSchemaTable struct for WorkspaceSchemaTable
type WorkspaceSchemaTable struct {
	Columns     []WorkspaceSchemaTableColumn `json:"columns"`
	Description *string                      `json:"description,omitempty"`
	Name        string                       `json:"name"`
}

// NewWorkspaceSchemaTable instantiates a new WorkspaceSchemaTable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceSchemaTable(columns []WorkspaceSchemaTableColumn, name string) *WorkspaceSchemaTable {
	this := WorkspaceSchemaTable{}
	this.Columns = columns
	this.Name = name
	return &this
}

// NewWorkspaceSchemaTableWithDefaults instantiates a new WorkspaceSchemaTable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceSchemaTableWithDefaults() *WorkspaceSchemaTable {
	this := WorkspaceSchemaTable{}
	return &this
}

// GetColumns returns the Columns field value
func (o *WorkspaceSchemaTable) GetColumns() []WorkspaceSchemaTableColumn {
	if o == nil {
		var ret []WorkspaceSchemaTableColumn
		return ret
	}

	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value
// and a boolean to check if the value has been set.
func (o *WorkspaceSchemaTable) GetColumnsOk() (*[]WorkspaceSchemaTableColumn, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Columns, true
}

// SetColumns sets field value
func (o *WorkspaceSchemaTable) SetColumns(v []WorkspaceSchemaTableColumn) {
	o.Columns = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WorkspaceSchemaTable) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceSchemaTable) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WorkspaceSchemaTable) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WorkspaceSchemaTable) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *WorkspaceSchemaTable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WorkspaceSchemaTable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WorkspaceSchemaTable) SetName(v string) {
	o.Name = v
}

func (o WorkspaceSchemaTable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["columns"] = o.Columns
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableWorkspaceSchemaTable struct {
	value *WorkspaceSchemaTable
	isSet bool
}

func (v NullableWorkspaceSchemaTable) Get() *WorkspaceSchemaTable {
	return v.value
}

func (v *NullableWorkspaceSchemaTable) Set(val *WorkspaceSchemaTable) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceSchemaTable) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceSchemaTable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceSchemaTable(val *WorkspaceSchemaTable) *NullableWorkspaceSchemaTable {
	return &NullableWorkspaceSchemaTable{value: val, isSet: true}
}

func (v NullableWorkspaceSchemaTable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceSchemaTable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}