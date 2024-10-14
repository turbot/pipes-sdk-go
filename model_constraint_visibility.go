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

// ConstraintVisibility the model 'ConstraintVisibility'
type ConstraintVisibility string

// List of ConstraintVisibility
const (
	ConstraintVisibilityPublic  ConstraintVisibility = "public"
	ConstraintVisibilityPrivate ConstraintVisibility = "private"
)

// All allowed values of ConstraintVisibility enum
var AllowedConstraintVisibilityEnumValues = []ConstraintVisibility{
	"public",
	"private",
}

func (v *ConstraintVisibility) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConstraintVisibility(value)
	for _, existing := range AllowedConstraintVisibilityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConstraintVisibility", value)
}

// NewConstraintVisibilityFromValue returns a pointer to a valid ConstraintVisibility
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConstraintVisibilityFromValue(v string) (*ConstraintVisibility, error) {
	ev := ConstraintVisibility(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConstraintVisibility: valid values are %v", v, AllowedConstraintVisibilityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConstraintVisibility) IsValid() bool {
	for _, existing := range AllowedConstraintVisibilityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConstraintVisibility value
func (v ConstraintVisibility) Ptr() *ConstraintVisibility {
	return &v
}

type NullableConstraintVisibility struct {
	value *ConstraintVisibility
	isSet bool
}

func (v NullableConstraintVisibility) Get() *ConstraintVisibility {
	return v.value
}

func (v *NullableConstraintVisibility) Set(val *ConstraintVisibility) {
	v.value = val
	v.isSet = true
}

func (v NullableConstraintVisibility) IsSet() bool {
	return v.isSet
}

func (v *NullableConstraintVisibility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConstraintVisibility(val *ConstraintVisibility) *NullableConstraintVisibility {
	return &NullableConstraintVisibility{value: val, isSet: true}
}

func (v NullableConstraintVisibility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConstraintVisibility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}