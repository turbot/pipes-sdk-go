/*
Turbot Pipes

Turbot Pipes is an intelligence, automation & security platform built specifically for DevOps.

API version: {{OPEN_API_VERSION}}
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipes

import (
	"encoding/json"
	"fmt"
)

// WorkspaceSchemaState the model 'WorkspaceSchemaState'
type WorkspaceSchemaState string

// List of WorkspaceSchemaState
const (
	WorkspaceSchemaStateGranted  WorkspaceSchemaState = "granted"
	WorkspaceSchemaStateDirect   WorkspaceSchemaState = "direct"
	WorkspaceSchemaStateIndirect WorkspaceSchemaState = "indirect"
)

// All allowed values of WorkspaceSchemaState enum
var AllowedWorkspaceSchemaStateEnumValues = []WorkspaceSchemaState{
	"granted",
	"direct",
	"indirect",
}

func (v *WorkspaceSchemaState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WorkspaceSchemaState(value)
	for _, existing := range AllowedWorkspaceSchemaStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WorkspaceSchemaState", value)
}

// NewWorkspaceSchemaStateFromValue returns a pointer to a valid WorkspaceSchemaState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWorkspaceSchemaStateFromValue(v string) (*WorkspaceSchemaState, error) {
	ev := WorkspaceSchemaState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WorkspaceSchemaState: valid values are %v", v, AllowedWorkspaceSchemaStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WorkspaceSchemaState) IsValid() bool {
	for _, existing := range AllowedWorkspaceSchemaStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WorkspaceSchemaState value
func (v WorkspaceSchemaState) Ptr() *WorkspaceSchemaState {
	return &v
}

type NullableWorkspaceSchemaState struct {
	value *WorkspaceSchemaState
	isSet bool
}

func (v NullableWorkspaceSchemaState) Get() *WorkspaceSchemaState {
	return v.value
}

func (v *NullableWorkspaceSchemaState) Set(val *WorkspaceSchemaState) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceSchemaState) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceSchemaState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceSchemaState(val *WorkspaceSchemaState) *NullableWorkspaceSchemaState {
	return &NullableWorkspaceSchemaState{value: val, isSet: true}
}

func (v NullableWorkspaceSchemaState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceSchemaState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}