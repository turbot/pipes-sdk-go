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

// Quota struct for Quota
type Quota struct {
	// Remaining limit
	Available *int32 `json:"available,omitempty"`
	// Max limit
	Limit *int32 `json:"limit,omitempty"`
	// Exhausted limit
	Used *int32 `json:"used,omitempty"`
}

// NewQuota instantiates a new Quota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQuota() *Quota {
	this := Quota{}
	return &this
}

// NewQuotaWithDefaults instantiates a new Quota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQuotaWithDefaults() *Quota {
	this := Quota{}
	return &this
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *Quota) GetAvailable() int32 {
	if o == nil || o.Available == nil {
		var ret int32
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quota) GetAvailableOk() (*int32, bool) {
	if o == nil || o.Available == nil {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *Quota) HasAvailable() bool {
	if o != nil && o.Available != nil {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given int32 and assigns it to the Available field.
func (o *Quota) SetAvailable(v int32) {
	o.Available = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *Quota) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quota) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *Quota) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *Quota) SetLimit(v int32) {
	o.Limit = &v
}

// GetUsed returns the Used field value if set, zero value otherwise.
func (o *Quota) GetUsed() int32 {
	if o == nil || o.Used == nil {
		var ret int32
		return ret
	}
	return *o.Used
}

// GetUsedOk returns a tuple with the Used field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Quota) GetUsedOk() (*int32, bool) {
	if o == nil || o.Used == nil {
		return nil, false
	}
	return o.Used, true
}

// HasUsed returns a boolean if a field has been set.
func (o *Quota) HasUsed() bool {
	if o != nil && o.Used != nil {
		return true
	}

	return false
}

// SetUsed gets a reference to the given int32 and assigns it to the Used field.
func (o *Quota) SetUsed(v int32) {
	o.Used = &v
}

func (o Quota) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Available != nil {
		toSerialize["available"] = o.Available
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Used != nil {
		toSerialize["used"] = o.Used
	}
	return json.Marshal(toSerialize)
}

type NullableQuota struct {
	value *Quota
	isSet bool
}

func (v NullableQuota) Get() *Quota {
	return v.value
}

func (v *NullableQuota) Set(val *Quota) {
	v.value = val
	v.isSet = true
}

func (v NullableQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuota(val *Quota) *NullableQuota {
	return &NullableQuota{value: val, isSet: true}
}

func (v NullableQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
