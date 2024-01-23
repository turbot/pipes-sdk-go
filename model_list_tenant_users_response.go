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

// ListTenantUsersResponse struct for ListTenantUsersResponse
type ListTenantUsersResponse struct {
	Items     *[]TenantUser `json:"items,omitempty"`
	NextToken *string       `json:"next_token,omitempty"`
}

// NewListTenantUsersResponse instantiates a new ListTenantUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTenantUsersResponse() *ListTenantUsersResponse {
	this := ListTenantUsersResponse{}
	return &this
}

// NewListTenantUsersResponseWithDefaults instantiates a new ListTenantUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTenantUsersResponseWithDefaults() *ListTenantUsersResponse {
	this := ListTenantUsersResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListTenantUsersResponse) GetItems() []TenantUser {
	if o == nil || o.Items == nil {
		var ret []TenantUser
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTenantUsersResponse) GetItemsOk() (*[]TenantUser, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListTenantUsersResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []TenantUser and assigns it to the Items field.
func (o *ListTenantUsersResponse) SetItems(v []TenantUser) {
	o.Items = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListTenantUsersResponse) GetNextToken() string {
	if o == nil || o.NextToken == nil {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTenantUsersResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || o.NextToken == nil {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListTenantUsersResponse) HasNextToken() bool {
	if o != nil && o.NextToken != nil {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListTenantUsersResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListTenantUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.NextToken != nil {
		toSerialize["next_token"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableListTenantUsersResponse struct {
	value *ListTenantUsersResponse
	isSet bool
}

func (v NullableListTenantUsersResponse) Get() *ListTenantUsersResponse {
	return v.value
}

func (v *NullableListTenantUsersResponse) Set(val *ListTenantUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListTenantUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListTenantUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTenantUsersResponse(val *ListTenantUsersResponse) *NullableListTenantUsersResponse {
	return &NullableListTenantUsersResponse{value: val, isSet: true}
}

func (v NullableListTenantUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTenantUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}