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

// DeleteAvatarResponse struct for DeleteAvatarResponse
type DeleteAvatarResponse struct {
	// The URL of the deleted avatar.
	AvatarUrl *string `json:"avatar_url,omitempty"`
}

// NewDeleteAvatarResponse instantiates a new DeleteAvatarResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteAvatarResponse() *DeleteAvatarResponse {
	this := DeleteAvatarResponse{}
	return &this
}

// NewDeleteAvatarResponseWithDefaults instantiates a new DeleteAvatarResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteAvatarResponseWithDefaults() *DeleteAvatarResponse {
	this := DeleteAvatarResponse{}
	return &this
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *DeleteAvatarResponse) GetAvatarUrl() string {
	if o == nil || o.AvatarUrl == nil {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteAvatarResponse) GetAvatarUrlOk() (*string, bool) {
	if o == nil || o.AvatarUrl == nil {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *DeleteAvatarResponse) HasAvatarUrl() bool {
	if o != nil && o.AvatarUrl != nil {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *DeleteAvatarResponse) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

func (o DeleteAvatarResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AvatarUrl != nil {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteAvatarResponse struct {
	value *DeleteAvatarResponse
	isSet bool
}

func (v NullableDeleteAvatarResponse) Get() *DeleteAvatarResponse {
	return v.value
}

func (v *NullableDeleteAvatarResponse) Set(val *DeleteAvatarResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteAvatarResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteAvatarResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteAvatarResponse(val *DeleteAvatarResponse) *NullableDeleteAvatarResponse {
	return &NullableDeleteAvatarResponse{value: val, isSet: true}
}

func (v NullableDeleteAvatarResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteAvatarResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
