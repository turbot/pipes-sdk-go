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

// ListPermissionsResponse struct for ListPermissionsResponse
type ListPermissionsResponse struct {
	Items     *[]Permission `json:"items,omitempty"`
	NextToken *string       `json:"next_token,omitempty"`
}

// NewListPermissionsResponse instantiates a new ListPermissionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPermissionsResponse() *ListPermissionsResponse {
	this := ListPermissionsResponse{}
	return &this
}

// NewListPermissionsResponseWithDefaults instantiates a new ListPermissionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPermissionsResponseWithDefaults() *ListPermissionsResponse {
	this := ListPermissionsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListPermissionsResponse) GetItems() []Permission {
	if o == nil || o.Items == nil {
		var ret []Permission
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPermissionsResponse) GetItemsOk() (*[]Permission, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListPermissionsResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Permission and assigns it to the Items field.
func (o *ListPermissionsResponse) SetItems(v []Permission) {
	o.Items = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListPermissionsResponse) GetNextToken() string {
	if o == nil || o.NextToken == nil {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPermissionsResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || o.NextToken == nil {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListPermissionsResponse) HasNextToken() bool {
	if o != nil && o.NextToken != nil {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListPermissionsResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListPermissionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.NextToken != nil {
		toSerialize["next_token"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableListPermissionsResponse struct {
	value *ListPermissionsResponse
	isSet bool
}

func (v NullableListPermissionsResponse) Get() *ListPermissionsResponse {
	return v.value
}

func (v *NullableListPermissionsResponse) Set(val *ListPermissionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListPermissionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListPermissionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPermissionsResponse(val *ListPermissionsResponse) *NullableListPermissionsResponse {
	return &NullableListPermissionsResponse{value: val, isSet: true}
}

func (v NullableListPermissionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPermissionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}