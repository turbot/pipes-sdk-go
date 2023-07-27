/*
Turbot Pipes

Turbot Pipes is an intelligence, automation & security platform built specifically for DevOps.

API version: {{OPEN_API_VERSION}}
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipes

import (
	"encoding/json"
)

// PipelineFrequency struct for PipelineFrequency
type PipelineFrequency struct {
	// The schedule for the pipeline if the type is interval or cron.
	Schedule *string `json:"schedule,omitempty"`
	// The type of frequency. Can be interval, cron or manual.
	Type string `json:"type"`
}

// NewPipelineFrequency instantiates a new PipelineFrequency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineFrequency(type_ string) *PipelineFrequency {
	this := PipelineFrequency{}
	this.Type = type_
	return &this
}

// NewPipelineFrequencyWithDefaults instantiates a new PipelineFrequency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineFrequencyWithDefaults() *PipelineFrequency {
	this := PipelineFrequency{}
	return &this
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *PipelineFrequency) GetSchedule() string {
	if o == nil || o.Schedule == nil {
		var ret string
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineFrequency) GetScheduleOk() (*string, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *PipelineFrequency) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given string and assigns it to the Schedule field.
func (o *PipelineFrequency) SetSchedule(v string) {
	o.Schedule = &v
}

// GetType returns the Type field value
func (o *PipelineFrequency) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PipelineFrequency) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PipelineFrequency) SetType(v string) {
	o.Type = v
}

func (o PipelineFrequency) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineFrequency struct {
	value *PipelineFrequency
	isSet bool
}

func (v NullablePipelineFrequency) Get() *PipelineFrequency {
	return v.value
}

func (v *NullablePipelineFrequency) Set(val *PipelineFrequency) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineFrequency) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineFrequency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineFrequency(val *PipelineFrequency) *NullablePipelineFrequency {
	return &NullablePipelineFrequency{value: val, isSet: true}
}

func (v NullablePipelineFrequency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineFrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
