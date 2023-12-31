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

// ListDatatankPartResponse struct for ListDatatankPartResponse
type ListDatatankPartResponse struct {
	Items     *[]DatatankPart `json:"items,omitempty"`
	NextToken *string         `json:"next_token,omitempty"`
}

// NewListDatatankPartResponse instantiates a new ListDatatankPartResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListDatatankPartResponse() *ListDatatankPartResponse {
	this := ListDatatankPartResponse{}
	return &this
}

// NewListDatatankPartResponseWithDefaults instantiates a new ListDatatankPartResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListDatatankPartResponseWithDefaults() *ListDatatankPartResponse {
	this := ListDatatankPartResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListDatatankPartResponse) GetItems() []DatatankPart {
	if o == nil || o.Items == nil {
		var ret []DatatankPart
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDatatankPartResponse) GetItemsOk() (*[]DatatankPart, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListDatatankPartResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []DatatankPart and assigns it to the Items field.
func (o *ListDatatankPartResponse) SetItems(v []DatatankPart) {
	o.Items = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListDatatankPartResponse) GetNextToken() string {
	if o == nil || o.NextToken == nil {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListDatatankPartResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || o.NextToken == nil {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListDatatankPartResponse) HasNextToken() bool {
	if o != nil && o.NextToken != nil {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListDatatankPartResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListDatatankPartResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.NextToken != nil {
		toSerialize["next_token"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableListDatatankPartResponse struct {
	value *ListDatatankPartResponse
	isSet bool
}

func (v NullableListDatatankPartResponse) Get() *ListDatatankPartResponse {
	return v.value
}

func (v *NullableListDatatankPartResponse) Set(val *ListDatatankPartResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListDatatankPartResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListDatatankPartResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListDatatankPartResponse(val *ListDatatankPartResponse) *NullableListDatatankPartResponse {
	return &NullableListDatatankPartResponse{value: val, isSet: true}
}

func (v NullableListDatatankPartResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListDatatankPartResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
