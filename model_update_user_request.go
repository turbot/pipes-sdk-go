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

// UpdateUserRequest struct for UpdateUserRequest
type UpdateUserRequest struct {
	DisplayName           *string   `json:"display_name,omitempty"`
	Handle                *string   `json:"handle,omitempty"`
	TokenMinIssuedAt      *JSONTime `json:"token_min_issued_at,omitempty"`
	Url                   *string   `json:"url,omitempty"`
	UsageComputeAction    *string   `json:"usage_compute_action,omitempty"`
	UsageComputeThreshold *int32    `json:"usage_compute_threshold,omitempty"`
	UsageStorageAction    *string   `json:"usage_storage_action,omitempty"`
	UsageStorageThreshold *int32    `json:"usage_storage_threshold,omitempty"`
}

// NewUpdateUserRequest instantiates a new UpdateUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserRequest() *UpdateUserRequest {
	this := UpdateUserRequest{}
	return &this
}

// NewUpdateUserRequestWithDefaults instantiates a new UpdateUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserRequestWithDefaults() *UpdateUserRequest {
	this := UpdateUserRequest{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UpdateUserRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasHandle() bool {
	if o != nil && o.Handle != nil {
		return true
	}

	return false
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *UpdateUserRequest) SetHandle(v string) {
	o.Handle = &v
}

// GetTokenMinIssuedAt returns the TokenMinIssuedAt field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetTokenMinIssuedAt() JSONTime {
	if o == nil || o.TokenMinIssuedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.TokenMinIssuedAt
}

// GetTokenMinIssuedAtOk returns a tuple with the TokenMinIssuedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetTokenMinIssuedAtOk() (*JSONTime, bool) {
	if o == nil || o.TokenMinIssuedAt == nil {
		return nil, false
	}
	return o.TokenMinIssuedAt, true
}

// HasTokenMinIssuedAt returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasTokenMinIssuedAt() bool {
	if o != nil && o.TokenMinIssuedAt != nil {
		return true
	}

	return false
}

// SetTokenMinIssuedAt gets a reference to the given JSONTime and assigns it to the TokenMinIssuedAt field.
func (o *UpdateUserRequest) SetTokenMinIssuedAt(v JSONTime) {
	o.TokenMinIssuedAt = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UpdateUserRequest) SetUrl(v string) {
	o.Url = &v
}

// GetUsageComputeAction returns the UsageComputeAction field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetUsageComputeAction() string {
	if o == nil || o.UsageComputeAction == nil {
		var ret string
		return ret
	}
	return *o.UsageComputeAction
}

// GetUsageComputeActionOk returns a tuple with the UsageComputeAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetUsageComputeActionOk() (*string, bool) {
	if o == nil || o.UsageComputeAction == nil {
		return nil, false
	}
	return o.UsageComputeAction, true
}

// HasUsageComputeAction returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasUsageComputeAction() bool {
	if o != nil && o.UsageComputeAction != nil {
		return true
	}

	return false
}

// SetUsageComputeAction gets a reference to the given string and assigns it to the UsageComputeAction field.
func (o *UpdateUserRequest) SetUsageComputeAction(v string) {
	o.UsageComputeAction = &v
}

// GetUsageComputeThreshold returns the UsageComputeThreshold field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetUsageComputeThreshold() int32 {
	if o == nil || o.UsageComputeThreshold == nil {
		var ret int32
		return ret
	}
	return *o.UsageComputeThreshold
}

// GetUsageComputeThresholdOk returns a tuple with the UsageComputeThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetUsageComputeThresholdOk() (*int32, bool) {
	if o == nil || o.UsageComputeThreshold == nil {
		return nil, false
	}
	return o.UsageComputeThreshold, true
}

// HasUsageComputeThreshold returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasUsageComputeThreshold() bool {
	if o != nil && o.UsageComputeThreshold != nil {
		return true
	}

	return false
}

// SetUsageComputeThreshold gets a reference to the given int32 and assigns it to the UsageComputeThreshold field.
func (o *UpdateUserRequest) SetUsageComputeThreshold(v int32) {
	o.UsageComputeThreshold = &v
}

// GetUsageStorageAction returns the UsageStorageAction field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetUsageStorageAction() string {
	if o == nil || o.UsageStorageAction == nil {
		var ret string
		return ret
	}
	return *o.UsageStorageAction
}

// GetUsageStorageActionOk returns a tuple with the UsageStorageAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetUsageStorageActionOk() (*string, bool) {
	if o == nil || o.UsageStorageAction == nil {
		return nil, false
	}
	return o.UsageStorageAction, true
}

// HasUsageStorageAction returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasUsageStorageAction() bool {
	if o != nil && o.UsageStorageAction != nil {
		return true
	}

	return false
}

// SetUsageStorageAction gets a reference to the given string and assigns it to the UsageStorageAction field.
func (o *UpdateUserRequest) SetUsageStorageAction(v string) {
	o.UsageStorageAction = &v
}

// GetUsageStorageThreshold returns the UsageStorageThreshold field value if set, zero value otherwise.
func (o *UpdateUserRequest) GetUsageStorageThreshold() int32 {
	if o == nil || o.UsageStorageThreshold == nil {
		var ret int32
		return ret
	}
	return *o.UsageStorageThreshold
}

// GetUsageStorageThresholdOk returns a tuple with the UsageStorageThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserRequest) GetUsageStorageThresholdOk() (*int32, bool) {
	if o == nil || o.UsageStorageThreshold == nil {
		return nil, false
	}
	return o.UsageStorageThreshold, true
}

// HasUsageStorageThreshold returns a boolean if a field has been set.
func (o *UpdateUserRequest) HasUsageStorageThreshold() bool {
	if o != nil && o.UsageStorageThreshold != nil {
		return true
	}

	return false
}

// SetUsageStorageThreshold gets a reference to the given int32 and assigns it to the UsageStorageThreshold field.
func (o *UpdateUserRequest) SetUsageStorageThreshold(v int32) {
	o.UsageStorageThreshold = &v
}

func (o UpdateUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Handle != nil {
		toSerialize["handle"] = o.Handle
	}
	if o.TokenMinIssuedAt != nil {
		toSerialize["token_min_issued_at"] = o.TokenMinIssuedAt
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.UsageComputeAction != nil {
		toSerialize["usage_compute_action"] = o.UsageComputeAction
	}
	if o.UsageComputeThreshold != nil {
		toSerialize["usage_compute_threshold"] = o.UsageComputeThreshold
	}
	if o.UsageStorageAction != nil {
		toSerialize["usage_storage_action"] = o.UsageStorageAction
	}
	if o.UsageStorageThreshold != nil {
		toSerialize["usage_storage_threshold"] = o.UsageStorageThreshold
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateUserRequest struct {
	value *UpdateUserRequest
	isSet bool
}

func (v NullableUpdateUserRequest) Get() *UpdateUserRequest {
	return v.value
}

func (v *NullableUpdateUserRequest) Set(val *UpdateUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserRequest(val *UpdateUserRequest) *NullableUpdateUserRequest {
	return &NullableUpdateUserRequest{value: val, isSet: true}
}

func (v NullableUpdateUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
