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

// ListConstraintsResponse struct for ListConstraintsResponse
type ListConstraintsResponse struct {
	Items     *[]Constraint `json:"items,omitempty"`
	NextToken *string       `json:"next_token,omitempty"`
}

// NewListConstraintsResponse instantiates a new ListConstraintsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConstraintsResponse() *ListConstraintsResponse {
	this := ListConstraintsResponse{}
	return &this
}

// NewListConstraintsResponseWithDefaults instantiates a new ListConstraintsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConstraintsResponseWithDefaults() *ListConstraintsResponse {
	this := ListConstraintsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListConstraintsResponse) GetItems() []Constraint {
	if o == nil || o.Items == nil {
		var ret []Constraint
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConstraintsResponse) GetItemsOk() (*[]Constraint, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListConstraintsResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Constraint and assigns it to the Items field.
func (o *ListConstraintsResponse) SetItems(v []Constraint) {
	o.Items = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListConstraintsResponse) GetNextToken() string {
	if o == nil || o.NextToken == nil {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConstraintsResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || o.NextToken == nil {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListConstraintsResponse) HasNextToken() bool {
	if o != nil && o.NextToken != nil {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListConstraintsResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListConstraintsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.NextToken != nil {
		toSerialize["next_token"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableListConstraintsResponse struct {
	value *ListConstraintsResponse
	isSet bool
}

func (v NullableListConstraintsResponse) Get() *ListConstraintsResponse {
	return v.value
}

func (v *NullableListConstraintsResponse) Set(val *ListConstraintsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListConstraintsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListConstraintsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConstraintsResponse(val *ListConstraintsResponse) *NullableListConstraintsResponse {
	return &NullableListConstraintsResponse{value: val, isSet: true}
}

func (v NullableListConstraintsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConstraintsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
