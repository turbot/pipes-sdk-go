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

// UsageUnitType the model 'UsageUnitType'
type UsageUnitType string

// List of UsageUnitType
const (
	UsageUnitTypeByte        UsageUnitType = "byte"
	UsageUnitTypeCount       UsageUnitType = "count"
	UsageUnitTypeMillisecond UsageUnitType = "millisecond"
)

// All allowed values of UsageUnitType enum
var AllowedUsageUnitTypeEnumValues = []UsageUnitType{
	"byte",
	"count",
	"millisecond",
}

func (v *UsageUnitType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UsageUnitType(value)
	for _, existing := range AllowedUsageUnitTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UsageUnitType", value)
}

// NewUsageUnitTypeFromValue returns a pointer to a valid UsageUnitType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUsageUnitTypeFromValue(v string) (*UsageUnitType, error) {
	ev := UsageUnitType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UsageUnitType: valid values are %v", v, AllowedUsageUnitTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UsageUnitType) IsValid() bool {
	for _, existing := range AllowedUsageUnitTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UsageUnitType value
func (v UsageUnitType) Ptr() *UsageUnitType {
	return &v
}

type NullableUsageUnitType struct {
	value *UsageUnitType
	isSet bool
}

func (v NullableUsageUnitType) Get() *UsageUnitType {
	return v.value
}

func (v *NullableUsageUnitType) Set(val *UsageUnitType) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageUnitType) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageUnitType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageUnitType(val *UsageUnitType) *NullableUsageUnitType {
	return &NullableUsageUnitType{value: val, isSet: true}
}

func (v NullableUsageUnitType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageUnitType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}