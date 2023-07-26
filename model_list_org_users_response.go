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

// ListOrgUsersResponse struct for ListOrgUsersResponse
type ListOrgUsersResponse struct {
	Items     *[]OrgUser `json:"items,omitempty"`
	NextToken *string    `json:"next_token,omitempty"`
}

// NewListOrgUsersResponse instantiates a new ListOrgUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOrgUsersResponse() *ListOrgUsersResponse {
	this := ListOrgUsersResponse{}
	return &this
}

// NewListOrgUsersResponseWithDefaults instantiates a new ListOrgUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOrgUsersResponseWithDefaults() *ListOrgUsersResponse {
	this := ListOrgUsersResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListOrgUsersResponse) GetItems() []OrgUser {
	if o == nil || o.Items == nil {
		var ret []OrgUser
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrgUsersResponse) GetItemsOk() (*[]OrgUser, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListOrgUsersResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []OrgUser and assigns it to the Items field.
func (o *ListOrgUsersResponse) SetItems(v []OrgUser) {
	o.Items = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListOrgUsersResponse) GetNextToken() string {
	if o == nil || o.NextToken == nil {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOrgUsersResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || o.NextToken == nil {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListOrgUsersResponse) HasNextToken() bool {
	if o != nil && o.NextToken != nil {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListOrgUsersResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListOrgUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.NextToken != nil {
		toSerialize["next_token"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableListOrgUsersResponse struct {
	value *ListOrgUsersResponse
	isSet bool
}

func (v NullableListOrgUsersResponse) Get() *ListOrgUsersResponse {
	return v.value
}

func (v *NullableListOrgUsersResponse) Set(val *ListOrgUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrgUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrgUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrgUsersResponse(val *ListOrgUsersResponse) *NullableListOrgUsersResponse {
	return &NullableListOrgUsersResponse{value: val, isSet: true}
}

func (v NullableListOrgUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrgUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
