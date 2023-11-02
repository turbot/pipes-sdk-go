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

// UpdateOrgRequest struct for UpdateOrgRequest
type UpdateOrgRequest struct {
	DisplayName *string `json:"display_name,omitempty"`
	Handle      *string `json:"handle,omitempty"`
	// The time which user and temporary auth tokens must be issued after.
	TokenMinIssuedAt      *string `json:"token_min_issued_at,omitempty"`
	Url                   *string `json:"url,omitempty"`
	UsageComputeAction    *string `json:"usage_compute_action,omitempty"`
	UsageComputeThreshold *int64  `json:"usage_compute_threshold,omitempty"`
	UsageStorageAction    *string `json:"usage_storage_action,omitempty"`
	UsageStorageThreshold *int64  `json:"usage_storage_threshold,omitempty"`
	UsageUserAction       *string `json:"usage_user_action,omitempty"`
	UsageUserThreshold    *int64  `json:"usage_user_threshold,omitempty"`
}

// NewUpdateOrgRequest instantiates a new UpdateOrgRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOrgRequest() *UpdateOrgRequest {
	this := UpdateOrgRequest{}
	return &this
}

// NewUpdateOrgRequestWithDefaults instantiates a new UpdateOrgRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOrgRequestWithDefaults() *UpdateOrgRequest {
	this := UpdateOrgRequest{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *UpdateOrgRequest) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrgRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *UpdateOrgRequest) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *UpdateOrgRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetHandle returns the Handle field value if set, zero value otherwise.
func (o *UpdateOrgRequest) GetHandle() string {
	if o == nil || o.Handle == nil {
		var ret string
		return ret
	}
	return *o.Handle
}

// GetHandleOk returns a tuple with the Handle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrgRequest) GetHandleOk() (*string, bool) {
	if o == nil || o.Handle == nil {
		return nil, false
	}
	return o.Handle, true
}

// HasHandle returns a boolean if a field has been set.
func (o *UpdateOrgRequest) HasHandle() bool {
	if o != nil && o.Handle != nil {
		return true
	}

	return false
}

// SetHandle gets a reference to the given string and assigns it to the Handle field.
func (o *UpdateOrgRequest) SetHandle(v string) {
	o.Handle = &v
}

// GetTokenMinIssuedAt returns the TokenMinIssuedAt field value if set, zero value otherwise.
func (o *UpdateOrgRequest) GetTokenMinIssuedAt() string {
	if o == nil || o.TokenMinIssuedAt == nil {
		var ret string
		return ret
	}
	return *o.TokenMinIssuedAt
}

// GetTokenMinIssuedAtOk returns a tuple with the TokenMinIssuedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrgRequest) GetTokenMinIssuedAtOk() (*string, bool) {
	if o == nil || o.TokenMinIssuedAt == nil {
		return nil, false
	}
	return o.TokenMinIssuedAt, true
}

// HasTokenMinIssuedAt returns a boolean if a field has been set.
func (o *UpdateOrgRequest) HasTokenMinIssuedAt() bool {
	if o != nil && o.TokenMinIssuedAt != nil {
		return true
	}

	return false
}

// SetTokenMinIssuedAt gets a reference to the given string and assigns it to the TokenMinIssuedAt field.
func (o *UpdateOrgRequest) SetTokenMinIssuedAt(v string) {
	o.TokenMinIssuedAt = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *UpdateOrgRequest) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrgRequest) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *UpdateOrgRequest) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *UpdateOrgRequest) SetUrl(v string) {
	o.Url = &v
}

// GetUsageComputeAction returns the UsageComputeAction field value if set, zero value otherwise.
func (o *UpdateOrgRequest) GetUsageComputeAction() string {
	if o == nil || o.UsageComputeAction == nil {
		var ret string
		return ret
	}
	return *o.UsageComputeAction
}

// GetUsageComputeActionOk returns a tuple with the UsageComputeAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrgRequest) GetUsageComputeActionOk() (*string, bool) {
	if o == nil || o.UsageComputeAction == nil {
		return nil, false
	}
	return o.UsageComputeAction, true
}

// HasUsageComputeAction returns a boolean if a field has been set.
func (o *UpdateOrgRequest) HasUsageComputeAction() bool {
	if o != nil && o.UsageComputeAction != nil {
		return true
	}

	return false
}

// SetUsageComputeAction gets a reference to the given string and assigns it to the UsageComputeAction field.
func (o *UpdateOrgRequest) SetUsageComputeAction(v string) {
	o.UsageComputeAction = &v
}

// GetUsageComputeThreshold returns the UsageComputeThreshold field value if set, zero value otherwise.
func (o *UpdateOrgRequest) GetUsageComputeThreshold() int64 {
	if o == nil || o.UsageComputeThreshold == nil {
		var ret int64
		return ret
	}
	return *o.UsageComputeThreshold
}

// GetUsageComputeThresholdOk returns a tuple with the UsageComputeThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrgRequest) GetUsageComputeThresholdOk() (*int64, bool) {
	if o == nil || o.UsageComputeThreshold == nil {
		return nil, false
	}
	return o.UsageComputeThreshold, true
}

// HasUsageComputeThreshold returns a boolean if a field has been set.
func (o *UpdateOrgRequest) HasUsageComputeThreshold() bool {
	if o != nil && o.UsageComputeThreshold != nil {
		return true
	}

	return false
}

// SetUsageComputeThreshold gets a reference to the given int64 and assigns it to the UsageComputeThreshold field.
func (o *UpdateOrgRequest) SetUsageComputeThreshold(v int64) {
	o.UsageComputeThreshold = &v
}

// GetUsageStorageAction returns the UsageStorageAction field value if set, zero value otherwise.
func (o *UpdateOrgRequest) GetUsageStorageAction() string {
	if o == nil || o.UsageStorageAction == nil {
		var ret string
		return ret
	}
	return *o.UsageStorageAction
}

// GetUsageStorageActionOk returns a tuple with the UsageStorageAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrgRequest) GetUsageStorageActionOk() (*string, bool) {
	if o == nil || o.UsageStorageAction == nil {
		return nil, false
	}
	return o.UsageStorageAction, true
}

// HasUsageStorageAction returns a boolean if a field has been set.
func (o *UpdateOrgRequest) HasUsageStorageAction() bool {
	if o != nil && o.UsageStorageAction != nil {
		return true
	}

	return false
}

// SetUsageStorageAction gets a reference to the given string and assigns it to the UsageStorageAction field.
func (o *UpdateOrgRequest) SetUsageStorageAction(v string) {
	o.UsageStorageAction = &v
}

// GetUsageStorageThreshold returns the UsageStorageThreshold field value if set, zero value otherwise.
func (o *UpdateOrgRequest) GetUsageStorageThreshold() int64 {
	if o == nil || o.UsageStorageThreshold == nil {
		var ret int64
		return ret
	}
	return *o.UsageStorageThreshold
}

// GetUsageStorageThresholdOk returns a tuple with the UsageStorageThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrgRequest) GetUsageStorageThresholdOk() (*int64, bool) {
	if o == nil || o.UsageStorageThreshold == nil {
		return nil, false
	}
	return o.UsageStorageThreshold, true
}

// HasUsageStorageThreshold returns a boolean if a field has been set.
func (o *UpdateOrgRequest) HasUsageStorageThreshold() bool {
	if o != nil && o.UsageStorageThreshold != nil {
		return true
	}

	return false
}

// SetUsageStorageThreshold gets a reference to the given int64 and assigns it to the UsageStorageThreshold field.
func (o *UpdateOrgRequest) SetUsageStorageThreshold(v int64) {
	o.UsageStorageThreshold = &v
}

// GetUsageUserAction returns the UsageUserAction field value if set, zero value otherwise.
func (o *UpdateOrgRequest) GetUsageUserAction() string {
	if o == nil || o.UsageUserAction == nil {
		var ret string
		return ret
	}
	return *o.UsageUserAction
}

// GetUsageUserActionOk returns a tuple with the UsageUserAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrgRequest) GetUsageUserActionOk() (*string, bool) {
	if o == nil || o.UsageUserAction == nil {
		return nil, false
	}
	return o.UsageUserAction, true
}

// HasUsageUserAction returns a boolean if a field has been set.
func (o *UpdateOrgRequest) HasUsageUserAction() bool {
	if o != nil && o.UsageUserAction != nil {
		return true
	}

	return false
}

// SetUsageUserAction gets a reference to the given string and assigns it to the UsageUserAction field.
func (o *UpdateOrgRequest) SetUsageUserAction(v string) {
	o.UsageUserAction = &v
}

// GetUsageUserThreshold returns the UsageUserThreshold field value if set, zero value otherwise.
func (o *UpdateOrgRequest) GetUsageUserThreshold() int64 {
	if o == nil || o.UsageUserThreshold == nil {
		var ret int64
		return ret
	}
	return *o.UsageUserThreshold
}

// GetUsageUserThresholdOk returns a tuple with the UsageUserThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOrgRequest) GetUsageUserThresholdOk() (*int64, bool) {
	if o == nil || o.UsageUserThreshold == nil {
		return nil, false
	}
	return o.UsageUserThreshold, true
}

// HasUsageUserThreshold returns a boolean if a field has been set.
func (o *UpdateOrgRequest) HasUsageUserThreshold() bool {
	if o != nil && o.UsageUserThreshold != nil {
		return true
	}

	return false
}

// SetUsageUserThreshold gets a reference to the given int64 and assigns it to the UsageUserThreshold field.
func (o *UpdateOrgRequest) SetUsageUserThreshold(v int64) {
	o.UsageUserThreshold = &v
}

func (o UpdateOrgRequest) MarshalJSON() ([]byte, error) {
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
	if o.UsageUserAction != nil {
		toSerialize["usage_user_action"] = o.UsageUserAction
	}
	if o.UsageUserThreshold != nil {
		toSerialize["usage_user_threshold"] = o.UsageUserThreshold
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateOrgRequest struct {
	value *UpdateOrgRequest
	isSet bool
}

func (v NullableUpdateOrgRequest) Get() *UpdateOrgRequest {
	return v.value
}

func (v *NullableUpdateOrgRequest) Set(val *UpdateOrgRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOrgRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOrgRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOrgRequest(val *UpdateOrgRequest) *NullableUpdateOrgRequest {
	return &NullableUpdateOrgRequest{value: val, isSet: true}
}

func (v NullableUpdateOrgRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOrgRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
