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

// ListIntegrationsResponse struct for ListIntegrationsResponse
type ListIntegrationsResponse struct {
	Items     *[]Integration `json:"items,omitempty"`
	NextToken *string        `json:"next_token,omitempty"`
}

// NewListIntegrationsResponse instantiates a new ListIntegrationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListIntegrationsResponse() *ListIntegrationsResponse {
	this := ListIntegrationsResponse{}
	return &this
}

// NewListIntegrationsResponseWithDefaults instantiates a new ListIntegrationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListIntegrationsResponseWithDefaults() *ListIntegrationsResponse {
	this := ListIntegrationsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListIntegrationsResponse) GetItems() []Integration {
	if o == nil || o.Items == nil {
		var ret []Integration
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrationsResponse) GetItemsOk() (*[]Integration, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListIntegrationsResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Integration and assigns it to the Items field.
func (o *ListIntegrationsResponse) SetItems(v []Integration) {
	o.Items = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListIntegrationsResponse) GetNextToken() string {
	if o == nil || o.NextToken == nil {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListIntegrationsResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || o.NextToken == nil {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListIntegrationsResponse) HasNextToken() bool {
	if o != nil && o.NextToken != nil {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListIntegrationsResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListIntegrationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.NextToken != nil {
		toSerialize["next_token"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableListIntegrationsResponse struct {
	value *ListIntegrationsResponse
	isSet bool
}

func (v NullableListIntegrationsResponse) Get() *ListIntegrationsResponse {
	return v.value
}

func (v *NullableListIntegrationsResponse) Set(val *ListIntegrationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListIntegrationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListIntegrationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListIntegrationsResponse(val *ListIntegrationsResponse) *NullableListIntegrationsResponse {
	return &NullableListIntegrationsResponse{value: val, isSet: true}
}

func (v NullableListIntegrationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListIntegrationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}