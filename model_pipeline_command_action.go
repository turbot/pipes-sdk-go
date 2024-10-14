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

// PipelineCommandAction the model 'PipelineCommandAction'
type PipelineCommandAction string

// List of PipelineCommandAction
const (
	PipelineCommandRun PipelineCommandAction = "run"
)

// All allowed values of PipelineCommandAction enum
var AllowedPipelineCommandActionEnumValues = []PipelineCommandAction{
	"run",
}

func (v *PipelineCommandAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PipelineCommandAction(value)
	for _, existing := range AllowedPipelineCommandActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PipelineCommandAction", value)
}

// NewPipelineCommandActionFromValue returns a pointer to a valid PipelineCommandAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPipelineCommandActionFromValue(v string) (*PipelineCommandAction, error) {
	ev := PipelineCommandAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PipelineCommandAction: valid values are %v", v, AllowedPipelineCommandActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PipelineCommandAction) IsValid() bool {
	for _, existing := range AllowedPipelineCommandActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PipelineCommandAction value
func (v PipelineCommandAction) Ptr() *PipelineCommandAction {
	return &v
}

type NullablePipelineCommandAction struct {
	value *PipelineCommandAction
	isSet bool
}

func (v NullablePipelineCommandAction) Get() *PipelineCommandAction {
	return v.value
}

func (v *NullablePipelineCommandAction) Set(val *PipelineCommandAction) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineCommandAction) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineCommandAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineCommandAction(val *PipelineCommandAction) *NullablePipelineCommandAction {
	return &NullablePipelineCommandAction{value: val, isSet: true}
}

func (v NullablePipelineCommandAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineCommandAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}