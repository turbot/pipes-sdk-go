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

// ListWorkspaceConnectionsResponse struct for ListWorkspaceConnectionsResponse
type ListWorkspaceConnectionsResponse struct {
	Items     *[]WorkspaceConnection `json:"items,omitempty"`
	NextToken *string                `json:"next_token,omitempty"`
}

// NewListWorkspaceConnectionsResponse instantiates a new ListWorkspaceConnectionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWorkspaceConnectionsResponse() *ListWorkspaceConnectionsResponse {
	this := ListWorkspaceConnectionsResponse{}
	return &this
}

// NewListWorkspaceConnectionsResponseWithDefaults instantiates a new ListWorkspaceConnectionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWorkspaceConnectionsResponseWithDefaults() *ListWorkspaceConnectionsResponse {
	this := ListWorkspaceConnectionsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListWorkspaceConnectionsResponse) GetItems() []WorkspaceConnection {
	if o == nil || o.Items == nil {
		var ret []WorkspaceConnection
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorkspaceConnectionsResponse) GetItemsOk() (*[]WorkspaceConnection, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListWorkspaceConnectionsResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []WorkspaceConnection and assigns it to the Items field.
func (o *ListWorkspaceConnectionsResponse) SetItems(v []WorkspaceConnection) {
	o.Items = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListWorkspaceConnectionsResponse) GetNextToken() string {
	if o == nil || o.NextToken == nil {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListWorkspaceConnectionsResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || o.NextToken == nil {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListWorkspaceConnectionsResponse) HasNextToken() bool {
	if o != nil && o.NextToken != nil {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListWorkspaceConnectionsResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListWorkspaceConnectionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.NextToken != nil {
		toSerialize["next_token"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableListWorkspaceConnectionsResponse struct {
	value *ListWorkspaceConnectionsResponse
	isSet bool
}

func (v NullableListWorkspaceConnectionsResponse) Get() *ListWorkspaceConnectionsResponse {
	return v.value
}

func (v *NullableListWorkspaceConnectionsResponse) Set(val *ListWorkspaceConnectionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListWorkspaceConnectionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListWorkspaceConnectionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWorkspaceConnectionsResponse(val *ListWorkspaceConnectionsResponse) *NullableListWorkspaceConnectionsResponse {
	return &NullableListWorkspaceConnectionsResponse{value: val, isSet: true}
}

func (v NullableListWorkspaceConnectionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWorkspaceConnectionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
