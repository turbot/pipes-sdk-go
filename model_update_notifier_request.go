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

// UpdateNotifierRequest struct for UpdateNotifierRequest
type UpdateNotifierRequest struct {
	// The name of the notifier to update
	Name *string `json:"name,omitempty"`
	// The notify targets for the notifier to update
	Notifies *[]map[string]interface{} `json:"notifies,omitempty"`
	// The state of the notifier to update
	State *NotifierState `json:"state,omitempty"`
}

// NewUpdateNotifierRequest instantiates a new UpdateNotifierRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateNotifierRequest() *UpdateNotifierRequest {
	this := UpdateNotifierRequest{}
	return &this
}

// NewUpdateNotifierRequestWithDefaults instantiates a new UpdateNotifierRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateNotifierRequestWithDefaults() *UpdateNotifierRequest {
	this := UpdateNotifierRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateNotifierRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNotifierRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateNotifierRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateNotifierRequest) SetName(v string) {
	o.Name = &v
}

// GetNotifies returns the Notifies field value if set, zero value otherwise.
func (o *UpdateNotifierRequest) GetNotifies() []map[string]interface{} {
	if o == nil || o.Notifies == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Notifies
}

// GetNotifiesOk returns a tuple with the Notifies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNotifierRequest) GetNotifiesOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Notifies == nil {
		return nil, false
	}
	return o.Notifies, true
}

// HasNotifies returns a boolean if a field has been set.
func (o *UpdateNotifierRequest) HasNotifies() bool {
	if o != nil && o.Notifies != nil {
		return true
	}

	return false
}

// SetNotifies gets a reference to the given []map[string]interface{} and assigns it to the Notifies field.
func (o *UpdateNotifierRequest) SetNotifies(v []map[string]interface{}) {
	o.Notifies = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *UpdateNotifierRequest) GetState() NotifierState {
	if o == nil || o.State == nil {
		var ret NotifierState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateNotifierRequest) GetStateOk() (*NotifierState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *UpdateNotifierRequest) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given NotifierState and assigns it to the State field.
func (o *UpdateNotifierRequest) SetState(v NotifierState) {
	o.State = &v
}

func (o UpdateNotifierRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Notifies != nil {
		toSerialize["notifies"] = o.Notifies
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateNotifierRequest struct {
	value *UpdateNotifierRequest
	isSet bool
}

func (v NullableUpdateNotifierRequest) Get() *UpdateNotifierRequest {
	return v.value
}

func (v *NullableUpdateNotifierRequest) Set(val *UpdateNotifierRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNotifierRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNotifierRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNotifierRequest(val *UpdateNotifierRequest) *NullableUpdateNotifierRequest {
	return &NullableUpdateNotifierRequest{value: val, isSet: true}
}

func (v NullableUpdateNotifierRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNotifierRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
