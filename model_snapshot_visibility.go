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

// SnapshotVisibility the model 'SnapshotVisibility'
type SnapshotVisibility string

// List of SnapshotVisibility
const (
	SnapshotVisibilityWorkspace      SnapshotVisibility = "workspace"
	SnapshotVisibilityAnyoneWithLink SnapshotVisibility = "anyone_with_link"
)

// All allowed values of SnapshotVisibility enum
var AllowedSnapshotVisibilityEnumValues = []SnapshotVisibility{
	"workspace",
	"anyone_with_link",
}

func (v *SnapshotVisibility) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SnapshotVisibility(value)
	for _, existing := range AllowedSnapshotVisibilityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SnapshotVisibility", value)
}

// NewSnapshotVisibilityFromValue returns a pointer to a valid SnapshotVisibility
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSnapshotVisibilityFromValue(v string) (*SnapshotVisibility, error) {
	ev := SnapshotVisibility(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SnapshotVisibility: valid values are %v", v, AllowedSnapshotVisibilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SnapshotVisibility) IsValid() bool {
	for _, existing := range AllowedSnapshotVisibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SnapshotVisibility value
func (v SnapshotVisibility) Ptr() *SnapshotVisibility {
	return &v
}

type NullableSnapshotVisibility struct {
	value *SnapshotVisibility
	isSet bool
}

func (v NullableSnapshotVisibility) Get() *SnapshotVisibility {
	return v.value
}

func (v *NullableSnapshotVisibility) Set(val *SnapshotVisibility) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotVisibility) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotVisibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotVisibility(val *SnapshotVisibility) *NullableSnapshotVisibility {
	return &NullableSnapshotVisibility{value: val, isSet: true}
}

func (v NullableSnapshotVisibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotVisibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
