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

// DatatankTableFreshness struct for DatatankTableFreshness
type DatatankTableFreshness struct {
	// The number of parts that are disabled.
	Disabled *int32 `json:"disabled,omitempty"`
	// The number of parts that have never successfully run.
	Error *int32 `json:"error,omitempty"`
	// The number of parts that did not have a successful refresh within the extended staleness timeframe.
	Expired *int32 `json:"expired,omitempty"`
	// The number of parts that had a successful query within the staleness timeframe.
	Fresh *int32 `json:"fresh,omitempty"`
	// The time when any of the parts was successfully refreshed most recently.
	NewestPartUpdatedAt *string `json:"newest_part_updated_at,omitempty"`
	// The time when any of the parts was successfully refreshed least recently.
	OldestPartUpdatedAt *string `json:"oldest_part_updated_at,omitempty"`
	// The number of parts that are yet to be executed.
	Pending *int32 `json:"pending,omitempty"`
	// The number of parts for which the data is in the process of being deleted.
	Removing *int32 `json:"removing,omitempty"`
	// The number of parts that had a successful query but not within the staleness timeframe.
	Stale *int32 `json:"stale,omitempty"`
	// The total number of parts in the datatank table.
	TotalParts *int32 `json:"total_parts,omitempty"`
}

// NewDatatankTableFreshness instantiates a new DatatankTableFreshness object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatatankTableFreshness() *DatatankTableFreshness {
	this := DatatankTableFreshness{}
	return &this
}

// NewDatatankTableFreshnessWithDefaults instantiates a new DatatankTableFreshness object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatatankTableFreshnessWithDefaults() *DatatankTableFreshness {
	this := DatatankTableFreshness{}
	return &this
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *DatatankTableFreshness) GetDisabled() int32 {
	if o == nil || o.Disabled == nil {
		var ret int32
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTableFreshness) GetDisabledOk() (*int32, bool) {
	if o == nil || o.Disabled == nil {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *DatatankTableFreshness) HasDisabled() bool {
	if o != nil && o.Disabled != nil {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given int32 and assigns it to the Disabled field.
func (o *DatatankTableFreshness) SetDisabled(v int32) {
	o.Disabled = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *DatatankTableFreshness) GetError() int32 {
	if o == nil || o.Error == nil {
		var ret int32
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTableFreshness) GetErrorOk() (*int32, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *DatatankTableFreshness) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given int32 and assigns it to the Error field.
func (o *DatatankTableFreshness) SetError(v int32) {
	o.Error = &v
}

// GetExpired returns the Expired field value if set, zero value otherwise.
func (o *DatatankTableFreshness) GetExpired() int32 {
	if o == nil || o.Expired == nil {
		var ret int32
		return ret
	}
	return *o.Expired
}

// GetExpiredOk returns a tuple with the Expired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTableFreshness) GetExpiredOk() (*int32, bool) {
	if o == nil || o.Expired == nil {
		return nil, false
	}
	return o.Expired, true
}

// HasExpired returns a boolean if a field has been set.
func (o *DatatankTableFreshness) HasExpired() bool {
	if o != nil && o.Expired != nil {
		return true
	}

	return false
}

// SetExpired gets a reference to the given int32 and assigns it to the Expired field.
func (o *DatatankTableFreshness) SetExpired(v int32) {
	o.Expired = &v
}

// GetFresh returns the Fresh field value if set, zero value otherwise.
func (o *DatatankTableFreshness) GetFresh() int32 {
	if o == nil || o.Fresh == nil {
		var ret int32
		return ret
	}
	return *o.Fresh
}

// GetFreshOk returns a tuple with the Fresh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTableFreshness) GetFreshOk() (*int32, bool) {
	if o == nil || o.Fresh == nil {
		return nil, false
	}
	return o.Fresh, true
}

// HasFresh returns a boolean if a field has been set.
func (o *DatatankTableFreshness) HasFresh() bool {
	if o != nil && o.Fresh != nil {
		return true
	}

	return false
}

// SetFresh gets a reference to the given int32 and assigns it to the Fresh field.
func (o *DatatankTableFreshness) SetFresh(v int32) {
	o.Fresh = &v
}

// GetNewestPartUpdatedAt returns the NewestPartUpdatedAt field value if set, zero value otherwise.
func (o *DatatankTableFreshness) GetNewestPartUpdatedAt() string {
	if o == nil || o.NewestPartUpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.NewestPartUpdatedAt
}

// GetNewestPartUpdatedAtOk returns a tuple with the NewestPartUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTableFreshness) GetNewestPartUpdatedAtOk() (*string, bool) {
	if o == nil || o.NewestPartUpdatedAt == nil {
		return nil, false
	}
	return o.NewestPartUpdatedAt, true
}

// HasNewestPartUpdatedAt returns a boolean if a field has been set.
func (o *DatatankTableFreshness) HasNewestPartUpdatedAt() bool {
	if o != nil && o.NewestPartUpdatedAt != nil {
		return true
	}

	return false
}

// SetNewestPartUpdatedAt gets a reference to the given string and assigns it to the NewestPartUpdatedAt field.
func (o *DatatankTableFreshness) SetNewestPartUpdatedAt(v string) {
	o.NewestPartUpdatedAt = &v
}

// GetOldestPartUpdatedAt returns the OldestPartUpdatedAt field value if set, zero value otherwise.
func (o *DatatankTableFreshness) GetOldestPartUpdatedAt() string {
	if o == nil || o.OldestPartUpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.OldestPartUpdatedAt
}

// GetOldestPartUpdatedAtOk returns a tuple with the OldestPartUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTableFreshness) GetOldestPartUpdatedAtOk() (*string, bool) {
	if o == nil || o.OldestPartUpdatedAt == nil {
		return nil, false
	}
	return o.OldestPartUpdatedAt, true
}

// HasOldestPartUpdatedAt returns a boolean if a field has been set.
func (o *DatatankTableFreshness) HasOldestPartUpdatedAt() bool {
	if o != nil && o.OldestPartUpdatedAt != nil {
		return true
	}

	return false
}

// SetOldestPartUpdatedAt gets a reference to the given string and assigns it to the OldestPartUpdatedAt field.
func (o *DatatankTableFreshness) SetOldestPartUpdatedAt(v string) {
	o.OldestPartUpdatedAt = &v
}

// GetPending returns the Pending field value if set, zero value otherwise.
func (o *DatatankTableFreshness) GetPending() int32 {
	if o == nil || o.Pending == nil {
		var ret int32
		return ret
	}
	return *o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTableFreshness) GetPendingOk() (*int32, bool) {
	if o == nil || o.Pending == nil {
		return nil, false
	}
	return o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *DatatankTableFreshness) HasPending() bool {
	if o != nil && o.Pending != nil {
		return true
	}

	return false
}

// SetPending gets a reference to the given int32 and assigns it to the Pending field.
func (o *DatatankTableFreshness) SetPending(v int32) {
	o.Pending = &v
}

// GetRemoving returns the Removing field value if set, zero value otherwise.
func (o *DatatankTableFreshness) GetRemoving() int32 {
	if o == nil || o.Removing == nil {
		var ret int32
		return ret
	}
	return *o.Removing
}

// GetRemovingOk returns a tuple with the Removing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTableFreshness) GetRemovingOk() (*int32, bool) {
	if o == nil || o.Removing == nil {
		return nil, false
	}
	return o.Removing, true
}

// HasRemoving returns a boolean if a field has been set.
func (o *DatatankTableFreshness) HasRemoving() bool {
	if o != nil && o.Removing != nil {
		return true
	}

	return false
}

// SetRemoving gets a reference to the given int32 and assigns it to the Removing field.
func (o *DatatankTableFreshness) SetRemoving(v int32) {
	o.Removing = &v
}

// GetStale returns the Stale field value if set, zero value otherwise.
func (o *DatatankTableFreshness) GetStale() int32 {
	if o == nil || o.Stale == nil {
		var ret int32
		return ret
	}
	return *o.Stale
}

// GetStaleOk returns a tuple with the Stale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTableFreshness) GetStaleOk() (*int32, bool) {
	if o == nil || o.Stale == nil {
		return nil, false
	}
	return o.Stale, true
}

// HasStale returns a boolean if a field has been set.
func (o *DatatankTableFreshness) HasStale() bool {
	if o != nil && o.Stale != nil {
		return true
	}

	return false
}

// SetStale gets a reference to the given int32 and assigns it to the Stale field.
func (o *DatatankTableFreshness) SetStale(v int32) {
	o.Stale = &v
}

// GetTotalParts returns the TotalParts field value if set, zero value otherwise.
func (o *DatatankTableFreshness) GetTotalParts() int32 {
	if o == nil || o.TotalParts == nil {
		var ret int32
		return ret
	}
	return *o.TotalParts
}

// GetTotalPartsOk returns a tuple with the TotalParts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTableFreshness) GetTotalPartsOk() (*int32, bool) {
	if o == nil || o.TotalParts == nil {
		return nil, false
	}
	return o.TotalParts, true
}

// HasTotalParts returns a boolean if a field has been set.
func (o *DatatankTableFreshness) HasTotalParts() bool {
	if o != nil && o.TotalParts != nil {
		return true
	}

	return false
}

// SetTotalParts gets a reference to the given int32 and assigns it to the TotalParts field.
func (o *DatatankTableFreshness) SetTotalParts(v int32) {
	o.TotalParts = &v
}

func (o DatatankTableFreshness) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Disabled != nil {
		toSerialize["disabled"] = o.Disabled
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.Expired != nil {
		toSerialize["expired"] = o.Expired
	}
	if o.Fresh != nil {
		toSerialize["fresh"] = o.Fresh
	}
	if o.NewestPartUpdatedAt != nil {
		toSerialize["newest_part_updated_at"] = o.NewestPartUpdatedAt
	}
	if o.OldestPartUpdatedAt != nil {
		toSerialize["oldest_part_updated_at"] = o.OldestPartUpdatedAt
	}
	if o.Pending != nil {
		toSerialize["pending"] = o.Pending
	}
	if o.Removing != nil {
		toSerialize["removing"] = o.Removing
	}
	if o.Stale != nil {
		toSerialize["stale"] = o.Stale
	}
	if o.TotalParts != nil {
		toSerialize["total_parts"] = o.TotalParts
	}
	return json.Marshal(toSerialize)
}

type NullableDatatankTableFreshness struct {
	value *DatatankTableFreshness
	isSet bool
}

func (v NullableDatatankTableFreshness) Get() *DatatankTableFreshness {
	return v.value
}

func (v *NullableDatatankTableFreshness) Set(val *DatatankTableFreshness) {
	v.value = val
	v.isSet = true
}

func (v NullableDatatankTableFreshness) IsSet() bool {
	return v.isSet
}

func (v *NullableDatatankTableFreshness) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatatankTableFreshness(val *DatatankTableFreshness) *NullableDatatankTableFreshness {
	return &NullableDatatankTableFreshness{value: val, isSet: true}
}

func (v NullableDatatankTableFreshness) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatatankTableFreshness) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
