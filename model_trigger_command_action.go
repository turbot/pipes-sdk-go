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

// TriggerCommandAction the model 'TriggerCommandAction'
type TriggerCommandAction string

// List of TriggerCommandAction
const (
	TriggerCommandRun TriggerCommandAction = "run"
)

// All allowed values of TriggerCommandAction enum
var AllowedTriggerCommandActionEnumValues = []TriggerCommandAction{
	"run",
}

func (v *TriggerCommandAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TriggerCommandAction(value)
	for _, existing := range AllowedTriggerCommandActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TriggerCommandAction", value)
}

// NewTriggerCommandActionFromValue returns a pointer to a valid TriggerCommandAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTriggerCommandActionFromValue(v string) (*TriggerCommandAction, error) {
	ev := TriggerCommandAction(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TriggerCommandAction: valid values are %v", v, AllowedTriggerCommandActionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TriggerCommandAction) IsValid() bool {
	for _, existing := range AllowedTriggerCommandActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TriggerCommandAction value
func (v TriggerCommandAction) Ptr() *TriggerCommandAction {
	return &v
}

type NullableTriggerCommandAction struct {
	value *TriggerCommandAction
	isSet bool
}

func (v NullableTriggerCommandAction) Get() *TriggerCommandAction {
	return v.value
}

func (v *NullableTriggerCommandAction) Set(val *TriggerCommandAction) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerCommandAction) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerCommandAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerCommandAction(val *TriggerCommandAction) *NullableTriggerCommandAction {
	return &NullableTriggerCommandAction{value: val, isSet: true}
}

func (v NullableTriggerCommandAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerCommandAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}