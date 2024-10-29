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

// FlowpipeInputStepType the model 'FlowpipeInputStepType'
type FlowpipeInputStepType string

// List of FlowpipeInputStepType
const (
	FlowpipeInputStepTypeInput   FlowpipeInputStepType = "input"
	FlowpipeInputStepTypeMessage FlowpipeInputStepType = "message"
)

// All allowed values of FlowpipeInputStepType enum
var AllowedFlowpipeInputStepTypeEnumValues = []FlowpipeInputStepType{
	"input",
	"message",
}

func (v *FlowpipeInputStepType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FlowpipeInputStepType(value)
	for _, existing := range AllowedFlowpipeInputStepTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FlowpipeInputStepType", value)
}

// NewFlowpipeInputStepTypeFromValue returns a pointer to a valid FlowpipeInputStepType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFlowpipeInputStepTypeFromValue(v string) (*FlowpipeInputStepType, error) {
	ev := FlowpipeInputStepType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FlowpipeInputStepType: valid values are %v", v, AllowedFlowpipeInputStepTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FlowpipeInputStepType) IsValid() bool {
	for _, existing := range AllowedFlowpipeInputStepTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FlowpipeInputStepType value
func (v FlowpipeInputStepType) Ptr() *FlowpipeInputStepType {
	return &v
}

type NullableFlowpipeInputStepType struct {
	value *FlowpipeInputStepType
	isSet bool
}

func (v NullableFlowpipeInputStepType) Get() *FlowpipeInputStepType {
	return v.value
}

func (v *NullableFlowpipeInputStepType) Set(val *FlowpipeInputStepType) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowpipeInputStepType) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowpipeInputStepType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowpipeInputStepType(val *FlowpipeInputStepType) *NullableFlowpipeInputStepType {
	return &NullableFlowpipeInputStepType{value: val, isSet: true}
}

func (v NullableFlowpipeInputStepType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowpipeInputStepType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
