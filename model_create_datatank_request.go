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

// CreateDatatankRequest struct for CreateDatatankRequest
type CreateDatatankRequest struct {
	Description *string `json:"description,omitempty"`
	Handle      string  `json:"handle"`
}

// NewCreateDatatankRequest instantiates a new CreateDatatankRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDatatankRequest(handle string) *CreateDatatankRequest {
	this := CreateDatatankRequest{}
	this.Handle = handle
	return &this
}

// NewCreateDatatankRequestWithDefaults instantiates a new CreateDatatankRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDatatankRequestWithDefaults() *CreateDatatankRequest {
	this := CreateDatatankRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateDatatankRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatatankRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateDatatankRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateDatatankRequest) SetDescription(v string) {
	o.Description = &v
}

// GetHandle returns the Handle field value
func (o *CreateDatatankRequest) GetHandle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Handle
}

// GetHandleOk returns a tuple with the Handle field value
// and a boolean to check if the value has been set.
func (o *CreateDatatankRequest) GetHandleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Handle, true
}

// SetHandle sets field value
func (o *CreateDatatankRequest) SetHandle(v string) {
	o.Handle = v
}

func (o CreateDatatankRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["handle"] = o.Handle
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDatatankRequest struct {
	value *CreateDatatankRequest
	isSet bool
}

func (v NullableCreateDatatankRequest) Get() *CreateDatatankRequest {
	return v.value
}

func (v *NullableCreateDatatankRequest) Set(val *CreateDatatankRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDatatankRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDatatankRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDatatankRequest(val *CreateDatatankRequest) *NullableCreateDatatankRequest {
	return &NullableCreateDatatankRequest{value: val, isSet: true}
}

func (v NullableCreateDatatankRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDatatankRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}