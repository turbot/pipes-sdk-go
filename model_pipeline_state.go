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

// PipelineState the model 'PipelineState'
type PipelineState string

// List of PipelineState
const (
	PipelineEnabling  PipelineState = "enabling"
	PipelineEnabled   PipelineState = "enabled"
	PipelinePausing   PipelineState = "pausing"
	PipelinePaused    PipelineState = "paused"
	PipelineDisabling PipelineState = "disabling"
	PipelineDisabled  PipelineState = "disabled"
	PipelineDeleted   PipelineState = "deleted"
)

// All allowed values of PipelineState enum
var AllowedPipelineStateEnumValues = []PipelineState{
	"enabling",
	"enabled",
	"pausing",
	"paused",
	"disabling",
	"disabled",
	"deleted",
}

func (v *PipelineState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PipelineState(value)
	for _, existing := range AllowedPipelineStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PipelineState", value)
}

// NewPipelineStateFromValue returns a pointer to a valid PipelineState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPipelineStateFromValue(v string) (*PipelineState, error) {
	ev := PipelineState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PipelineState: valid values are %v", v, AllowedPipelineStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PipelineState) IsValid() bool {
	for _, existing := range AllowedPipelineStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PipelineState value
func (v PipelineState) Ptr() *PipelineState {
	return &v
}

type NullablePipelineState struct {
	value *PipelineState
	isSet bool
}

func (v NullablePipelineState) Get() *PipelineState {
	return v.value
}

func (v *NullablePipelineState) Set(val *PipelineState) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineState) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineState(val *PipelineState) *NullablePipelineState {
	return &NullablePipelineState{value: val, isSet: true}
}

func (v NullablePipelineState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
