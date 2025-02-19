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

// ModTriggerState the model 'ModTriggerState'
type ModTriggerState string

// List of ModTriggerState
const (
	ModTriggerStateEnabled  ModTriggerState = "enabled"
	ModTriggerStateDisabled ModTriggerState = "disabled"
)

// All allowed values of ModTriggerState enum
var AllowedModTriggerStateEnumValues = []ModTriggerState{
	"enabled",
	"disabled",
}

func (v *ModTriggerState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModTriggerState(value)
	for _, existing := range AllowedModTriggerStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModTriggerState", value)
}

// NewModTriggerStateFromValue returns a pointer to a valid ModTriggerState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModTriggerStateFromValue(v string) (*ModTriggerState, error) {
	ev := ModTriggerState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModTriggerState: valid values are %v", v, AllowedModTriggerStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModTriggerState) IsValid() bool {
	for _, existing := range AllowedModTriggerStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ModTriggerState value
func (v ModTriggerState) Ptr() *ModTriggerState {
	return &v
}

type NullableModTriggerState struct {
	value *ModTriggerState
	isSet bool
}

func (v NullableModTriggerState) Get() *ModTriggerState {
	return v.value
}

func (v *NullableModTriggerState) Set(val *ModTriggerState) {
	v.value = val
	v.isSet = true
}

func (v NullableModTriggerState) IsSet() bool {
	return v.isSet
}

func (v *NullableModTriggerState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModTriggerState(val *ModTriggerState) *NullableModTriggerState {
	return &NullableModTriggerState{value: val, isSet: true}
}

func (v NullableModTriggerState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModTriggerState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
