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

// UpdateTenantUserRequest struct for UpdateTenantUserRequest
type UpdateTenantUserRequest struct {
	Role *string `json:"role,omitempty"`
}

// NewUpdateTenantUserRequest instantiates a new UpdateTenantUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTenantUserRequest() *UpdateTenantUserRequest {
	this := UpdateTenantUserRequest{}
	return &this
}

// NewUpdateTenantUserRequestWithDefaults instantiates a new UpdateTenantUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTenantUserRequestWithDefaults() *UpdateTenantUserRequest {
	this := UpdateTenantUserRequest{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *UpdateTenantUserRequest) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenantUserRequest) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *UpdateTenantUserRequest) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *UpdateTenantUserRequest) SetRole(v string) {
	o.Role = &v
}

func (o UpdateTenantUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateTenantUserRequest struct {
	value *UpdateTenantUserRequest
	isSet bool
}

func (v NullableUpdateTenantUserRequest) Get() *UpdateTenantUserRequest {
	return v.value
}

func (v *NullableUpdateTenantUserRequest) Set(val *UpdateTenantUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTenantUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTenantUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTenantUserRequest(val *UpdateTenantUserRequest) *NullableUpdateTenantUserRequest {
	return &NullableUpdateTenantUserRequest{value: val, isSet: true}
}

func (v NullableUpdateTenantUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTenantUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
