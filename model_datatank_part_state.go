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

// DatatankPartState the model 'DatatankPartState'
type DatatankPartState string

// List of DatatankPartState
const (
	DatatankPartCreating DatatankPartState = "creating"
	DatatankPartEnabled  DatatankPartState = "enabled"
	DatatankPartPaused   DatatankPartState = "paused"
	DatatankPartDisabled DatatankPartState = "disabled"
	DatatankPartDeleted  DatatankPartState = "deleted"
	DatatankPartFailed   DatatankPartState = "failed"
)

// All allowed values of DatatankPartState enum
var AllowedDatatankPartStateEnumValues = []DatatankPartState{
	"creating",
	"enabled",
	"paused",
	"disabled",
	"deleted",
	"failed",
}

func (v *DatatankPartState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DatatankPartState(value)
	for _, existing := range AllowedDatatankPartStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DatatankPartState", value)
}

// NewDatatankPartStateFromValue returns a pointer to a valid DatatankPartState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDatatankPartStateFromValue(v string) (*DatatankPartState, error) {
	ev := DatatankPartState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DatatankPartState: valid values are %v", v, AllowedDatatankPartStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DatatankPartState) IsValid() bool {
	for _, existing := range AllowedDatatankPartStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatatankPartState value
func (v DatatankPartState) Ptr() *DatatankPartState {
	return &v
}

type NullableDatatankPartState struct {
	value *DatatankPartState
	isSet bool
}

func (v NullableDatatankPartState) Get() *DatatankPartState {
	return v.value
}

func (v *NullableDatatankPartState) Set(val *DatatankPartState) {
	v.value = val
	v.isSet = true
}

func (v NullableDatatankPartState) IsSet() bool {
	return v.isSet
}

func (v *NullableDatatankPartState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatatankPartState(val *DatatankPartState) *NullableDatatankPartState {
	return &NullableDatatankPartState{value: val, isSet: true}
}

func (v NullableDatatankPartState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatatankPartState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
