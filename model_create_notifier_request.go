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

// CreateNotifierRequest struct for CreateNotifierRequest
type CreateNotifierRequest struct {
	// The name of the notifier to create
	Name string `json:"name"`
	// The notify targets for the notifier to create
	Notifies map[string]interface{} `json:"notifies"`
	// The state of the notifier to create
	State *NotifierState `json:"state,omitempty"`
}

// NewCreateNotifierRequest instantiates a new CreateNotifierRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNotifierRequest(name string, notifies map[string]interface{}) *CreateNotifierRequest {
	this := CreateNotifierRequest{}
	this.Name = name
	this.Notifies = notifies
	return &this
}

// NewCreateNotifierRequestWithDefaults instantiates a new CreateNotifierRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNotifierRequestWithDefaults() *CreateNotifierRequest {
	this := CreateNotifierRequest{}
	return &this
}

// GetName returns the Name field value
func (o *CreateNotifierRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateNotifierRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateNotifierRequest) SetName(v string) {
	o.Name = v
}

// GetNotifies returns the Notifies field value
func (o *CreateNotifierRequest) GetNotifies() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Notifies
}

// GetNotifiesOk returns a tuple with the Notifies field value
// and a boolean to check if the value has been set.
func (o *CreateNotifierRequest) GetNotifiesOk() (*map[string]interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Notifies, true
}

// SetNotifies sets field value
func (o *CreateNotifierRequest) SetNotifies(v map[string]interface{}) {
	o.Notifies = v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *CreateNotifierRequest) GetState() NotifierState {
	if o == nil || o.State == nil {
		var ret NotifierState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNotifierRequest) GetStateOk() (*NotifierState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *CreateNotifierRequest) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given NotifierState and assigns it to the State field.
func (o *CreateNotifierRequest) SetState(v NotifierState) {
	o.State = &v
}

func (o CreateNotifierRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["notifies"] = o.Notifies
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableCreateNotifierRequest struct {
	value *CreateNotifierRequest
	isSet bool
}

func (v NullableCreateNotifierRequest) Get() *CreateNotifierRequest {
	return v.value
}

func (v *NullableCreateNotifierRequest) Set(val *CreateNotifierRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNotifierRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNotifierRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNotifierRequest(val *CreateNotifierRequest) *NullableCreateNotifierRequest {
	return &NullableCreateNotifierRequest{value: val, isSet: true}
}

func (v NullableCreateNotifierRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNotifierRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
