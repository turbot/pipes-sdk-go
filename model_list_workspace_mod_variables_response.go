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

// ListWorkspaceModVariablesResponse struct for ListWorkspaceModVariablesResponse
type ListWorkspaceModVariablesResponse struct {
	Items     *[]WorkspaceModVariable `json:"items,omitempty"`
	NextToken *string                 `json:"next_token,omitempty"`
}

// NewListWorkspaceModVariablesResponse instantiates a new ListWorkspaceModVariablesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWorkspaceModVariablesResponse() *ListWorkspaceModVariablesResponse {
	this := ListWorkspaceModVariablesResponse{}
	return &this
}

// NewListWorkspaceModVariablesResponseWithDefaults instantiates a new ListWorkspaceModVariablesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWorkspaceModVariablesResponseWithDefaults() *ListWorkspaceModVariablesResponse {
	this := ListWorkspaceModVariablesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListWorkspaceModVariablesResponse) GetItems() []WorkspaceModVariable {
	if o == nil || o.Items == nil {
		var ret []WorkspaceModVariable
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorkspaceModVariablesResponse) GetItemsOk() (*[]WorkspaceModVariable, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListWorkspaceModVariablesResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []WorkspaceModVariable and assigns it to the Items field.
func (o *ListWorkspaceModVariablesResponse) SetItems(v []WorkspaceModVariable) {
	o.Items = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListWorkspaceModVariablesResponse) GetNextToken() string {
	if o == nil || o.NextToken == nil {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorkspaceModVariablesResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || o.NextToken == nil {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListWorkspaceModVariablesResponse) HasNextToken() bool {
	if o != nil && o.NextToken != nil {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListWorkspaceModVariablesResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListWorkspaceModVariablesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.NextToken != nil {
		toSerialize["next_token"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableListWorkspaceModVariablesResponse struct {
	value *ListWorkspaceModVariablesResponse
	isSet bool
}

func (v NullableListWorkspaceModVariablesResponse) Get() *ListWorkspaceModVariablesResponse {
	return v.value
}

func (v *NullableListWorkspaceModVariablesResponse) Set(val *ListWorkspaceModVariablesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListWorkspaceModVariablesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListWorkspaceModVariablesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWorkspaceModVariablesResponse(val *ListWorkspaceModVariablesResponse) *NullableListWorkspaceModVariablesResponse {
	return &NullableListWorkspaceModVariablesResponse{value: val, isSet: true}
}

func (v NullableListWorkspaceModVariablesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWorkspaceModVariablesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}