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

// CreateConnectionFolderRequest struct for CreateConnectionFolderRequest
type CreateConnectionFolderRequest struct {
	ParentId *string `json:"parent_id,omitempty"`
	Title    string  `json:"title"`
}

// NewCreateConnectionFolderRequest instantiates a new CreateConnectionFolderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateConnectionFolderRequest(title string) *CreateConnectionFolderRequest {
	this := CreateConnectionFolderRequest{}
	this.Title = title
	return &this
}

// NewCreateConnectionFolderRequestWithDefaults instantiates a new CreateConnectionFolderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateConnectionFolderRequestWithDefaults() *CreateConnectionFolderRequest {
	this := CreateConnectionFolderRequest{}
	return &this
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *CreateConnectionFolderRequest) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateConnectionFolderRequest) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *CreateConnectionFolderRequest) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *CreateConnectionFolderRequest) SetParentId(v string) {
	o.ParentId = &v
}

// GetTitle returns the Title field value
func (o *CreateConnectionFolderRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CreateConnectionFolderRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CreateConnectionFolderRequest) SetTitle(v string) {
	o.Title = v
}

func (o CreateConnectionFolderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ParentId != nil {
		toSerialize["parent_id"] = o.ParentId
	}
	if true {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableCreateConnectionFolderRequest struct {
	value *CreateConnectionFolderRequest
	isSet bool
}

func (v NullableCreateConnectionFolderRequest) Get() *CreateConnectionFolderRequest {
	return v.value
}

func (v *NullableCreateConnectionFolderRequest) Set(val *CreateConnectionFolderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateConnectionFolderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateConnectionFolderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateConnectionFolderRequest(val *CreateConnectionFolderRequest) *NullableCreateConnectionFolderRequest {
	return &NullableCreateConnectionFolderRequest{value: val, isSet: true}
}

func (v NullableCreateConnectionFolderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateConnectionFolderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}