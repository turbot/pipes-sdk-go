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

// WorkspaceState the model 'WorkspaceState'
type WorkspaceState string

// List of WorkspaceState
const (
	WorkspacePending      WorkspaceState = "pending"
	WorkspaceRunning      WorkspaceState = "running"
	WorkspaceShuttingDown WorkspaceState = "shutting-down"
	WorkspaceTerminated   WorkspaceState = "terminated"
	WorkspaceStopping     WorkspaceState = "stopping"
	WorkspaceStopped      WorkspaceState = "stopped"
	WorkspaceCreating     WorkspaceState = "creating"
	WorkspaceEnabling     WorkspaceState = "enabling"
	WorkspaceEnabled      WorkspaceState = "enabled"
	WorkspacePausing      WorkspaceState = "pausing"
	WorkspacePaused       WorkspaceState = "paused"
	WorkspaceDisabling    WorkspaceState = "disabling"
	WorkspaceDisabled     WorkspaceState = "disabled"
)

// All allowed values of WorkspaceState enum
var AllowedWorkspaceStateEnumValues = []WorkspaceState{
	"pending",
	"running",
	"shutting-down",
	"terminated",
	"stopping",
	"stopped",
	"creating",
	"enabling",
	"enabled",
	"pausing",
	"paused",
	"disabling",
	"disabled",
}

func (v *WorkspaceState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WorkspaceState(value)
	for _, existing := range AllowedWorkspaceStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WorkspaceState", value)
}

// NewWorkspaceStateFromValue returns a pointer to a valid WorkspaceState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWorkspaceStateFromValue(v string) (*WorkspaceState, error) {
	ev := WorkspaceState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WorkspaceState: valid values are %v", v, AllowedWorkspaceStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WorkspaceState) IsValid() bool {
	for _, existing := range AllowedWorkspaceStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WorkspaceState value
func (v WorkspaceState) Ptr() *WorkspaceState {
	return &v
}

type NullableWorkspaceState struct {
	value *WorkspaceState
	isSet bool
}

func (v NullableWorkspaceState) Get() *WorkspaceState {
	return v.value
}

func (v *NullableWorkspaceState) Set(val *WorkspaceState) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceState) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceState(val *WorkspaceState) *NullableWorkspaceState {
	return &NullableWorkspaceState{value: val, isSet: true}
}

func (v NullableWorkspaceState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
