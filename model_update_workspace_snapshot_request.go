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

// UpdateWorkspaceSnapshotRequest struct for UpdateWorkspaceSnapshotRequest
type UpdateWorkspaceSnapshotRequest struct {
	Tags interface{} `json:"tags,omitempty"`
	// The updated title for the snapshot.
	Title *string `json:"title,omitempty"`
	// The updated visibility for the snapshot.
	Visibility *SnapshotVisibility `json:"visibility,omitempty"`
}

// NewUpdateWorkspaceSnapshotRequest instantiates a new UpdateWorkspaceSnapshotRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateWorkspaceSnapshotRequest() *UpdateWorkspaceSnapshotRequest {
	this := UpdateWorkspaceSnapshotRequest{}
	return &this
}

// NewUpdateWorkspaceSnapshotRequestWithDefaults instantiates a new UpdateWorkspaceSnapshotRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateWorkspaceSnapshotRequestWithDefaults() *UpdateWorkspaceSnapshotRequest {
	this := UpdateWorkspaceSnapshotRequest{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateWorkspaceSnapshotRequest) GetTags() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateWorkspaceSnapshotRequest) GetTagsOk() (*interface{}, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateWorkspaceSnapshotRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given interface{} and assigns it to the Tags field.
func (o *UpdateWorkspaceSnapshotRequest) SetTags(v interface{}) {
	o.Tags = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UpdateWorkspaceSnapshotRequest) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWorkspaceSnapshotRequest) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UpdateWorkspaceSnapshotRequest) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UpdateWorkspaceSnapshotRequest) SetTitle(v string) {
	o.Title = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *UpdateWorkspaceSnapshotRequest) GetVisibility() SnapshotVisibility {
	if o == nil || o.Visibility == nil {
		var ret SnapshotVisibility
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateWorkspaceSnapshotRequest) GetVisibilityOk() (*SnapshotVisibility, bool) {
	if o == nil || o.Visibility == nil {
		return nil, false
	}
	return o.Visibility, true
}

// HasVisibility returns a boolean if a field has been set.
func (o *UpdateWorkspaceSnapshotRequest) HasVisibility() bool {
	if o != nil && o.Visibility != nil {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given SnapshotVisibility and assigns it to the Visibility field.
func (o *UpdateWorkspaceSnapshotRequest) SetVisibility(v SnapshotVisibility) {
	o.Visibility = &v
}

func (o UpdateWorkspaceSnapshotRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Visibility != nil {
		toSerialize["visibility"] = o.Visibility
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateWorkspaceSnapshotRequest struct {
	value *UpdateWorkspaceSnapshotRequest
	isSet bool
}

func (v NullableUpdateWorkspaceSnapshotRequest) Get() *UpdateWorkspaceSnapshotRequest {
	return v.value
}

func (v *NullableUpdateWorkspaceSnapshotRequest) Set(val *UpdateWorkspaceSnapshotRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateWorkspaceSnapshotRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateWorkspaceSnapshotRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateWorkspaceSnapshotRequest(val *UpdateWorkspaceSnapshotRequest) *NullableUpdateWorkspaceSnapshotRequest {
	return &NullableUpdateWorkspaceSnapshotRequest{value: val, isSet: true}
}

func (v NullableUpdateWorkspaceSnapshotRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateWorkspaceSnapshotRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
