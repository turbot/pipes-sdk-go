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

// TriggerState the model 'TriggerState'
type TriggerState string

// List of TriggerState
const (
	TriggerEnabled  TriggerState = "enabled"
	TriggerDisabled TriggerState = "disabled"
)

// All allowed values of TriggerState enum
var AllowedTriggerStateEnumValues = []TriggerState{
	"enabled",
	"disabled",
}

func (v *TriggerState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TriggerState(value)
	for _, existing := range AllowedTriggerStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TriggerState", value)
}

// NewTriggerStateFromValue returns a pointer to a valid TriggerState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTriggerStateFromValue(v string) (*TriggerState, error) {
	ev := TriggerState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TriggerState: valid values are %v", v, AllowedTriggerStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TriggerState) IsValid() bool {
	for _, existing := range AllowedTriggerStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TriggerState value
func (v TriggerState) Ptr() *TriggerState {
	return &v
}

type NullableTriggerState struct {
	value *TriggerState
	isSet bool
}

func (v NullableTriggerState) Get() *TriggerState {
	return v.value
}

func (v *NullableTriggerState) Set(val *TriggerState) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerState) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerState(val *TriggerState) *NullableTriggerState {
	return &NullableTriggerState{value: val, isSet: true}
}

func (v NullableTriggerState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
