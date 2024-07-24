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

// ConnectionTrunkItem struct for ConnectionTrunkItem
type ConnectionTrunkItem struct {
	Id    *string `json:"id,omitempty"`
	Title *string `json:"title,omitempty"`
	Type  *string `json:"type,omitempty"`
}

// NewConnectionTrunkItem instantiates a new ConnectionTrunkItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionTrunkItem() *ConnectionTrunkItem {
	this := ConnectionTrunkItem{}
	return &this
}

// NewConnectionTrunkItemWithDefaults instantiates a new ConnectionTrunkItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionTrunkItemWithDefaults() *ConnectionTrunkItem {
	this := ConnectionTrunkItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConnectionTrunkItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionTrunkItem) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConnectionTrunkItem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConnectionTrunkItem) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ConnectionTrunkItem) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionTrunkItem) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ConnectionTrunkItem) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ConnectionTrunkItem) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ConnectionTrunkItem) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionTrunkItem) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ConnectionTrunkItem) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ConnectionTrunkItem) SetType(v string) {
	o.Type = &v
}

func (o ConnectionTrunkItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionTrunkItem struct {
	value *ConnectionTrunkItem
	isSet bool
}

func (v NullableConnectionTrunkItem) Get() *ConnectionTrunkItem {
	return v.value
}

func (v *NullableConnectionTrunkItem) Set(val *ConnectionTrunkItem) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionTrunkItem) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionTrunkItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionTrunkItem(val *ConnectionTrunkItem) *NullableConnectionTrunkItem {
	return &NullableConnectionTrunkItem{value: val, isSet: true}
}

func (v NullableConnectionTrunkItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionTrunkItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
