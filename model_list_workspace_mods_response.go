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

// ListWorkspaceModsResponse struct for ListWorkspaceModsResponse
type ListWorkspaceModsResponse struct {
	Items     *[]WorkspaceMod `json:"items,omitempty"`
	NextToken *string         `json:"next_token,omitempty"`
}

// NewListWorkspaceModsResponse instantiates a new ListWorkspaceModsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWorkspaceModsResponse() *ListWorkspaceModsResponse {
	this := ListWorkspaceModsResponse{}
	return &this
}

// NewListWorkspaceModsResponseWithDefaults instantiates a new ListWorkspaceModsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWorkspaceModsResponseWithDefaults() *ListWorkspaceModsResponse {
	this := ListWorkspaceModsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListWorkspaceModsResponse) GetItems() []WorkspaceMod {
	if o == nil || o.Items == nil {
		var ret []WorkspaceMod
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorkspaceModsResponse) GetItemsOk() (*[]WorkspaceMod, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListWorkspaceModsResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []WorkspaceMod and assigns it to the Items field.
func (o *ListWorkspaceModsResponse) SetItems(v []WorkspaceMod) {
	o.Items = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListWorkspaceModsResponse) GetNextToken() string {
	if o == nil || o.NextToken == nil {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorkspaceModsResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || o.NextToken == nil {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListWorkspaceModsResponse) HasNextToken() bool {
	if o != nil && o.NextToken != nil {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListWorkspaceModsResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListWorkspaceModsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.NextToken != nil {
		toSerialize["next_token"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableListWorkspaceModsResponse struct {
	value *ListWorkspaceModsResponse
	isSet bool
}

func (v NullableListWorkspaceModsResponse) Get() *ListWorkspaceModsResponse {
	return v.value
}

func (v *NullableListWorkspaceModsResponse) Set(val *ListWorkspaceModsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListWorkspaceModsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListWorkspaceModsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWorkspaceModsResponse(val *ListWorkspaceModsResponse) *NullableListWorkspaceModsResponse {
	return &NullableListWorkspaceModsResponse{value: val, isSet: true}
}

func (v NullableListWorkspaceModsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWorkspaceModsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
