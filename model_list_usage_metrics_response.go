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

// ListUsageMetricsResponse struct for ListUsageMetricsResponse
type ListUsageMetricsResponse struct {
	Items     *[]UsageMetric `json:"items,omitempty"`
	NextToken *string        `json:"next_token,omitempty"`
}

// NewListUsageMetricsResponse instantiates a new ListUsageMetricsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUsageMetricsResponse() *ListUsageMetricsResponse {
	this := ListUsageMetricsResponse{}
	return &this
}

// NewListUsageMetricsResponseWithDefaults instantiates a new ListUsageMetricsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUsageMetricsResponseWithDefaults() *ListUsageMetricsResponse {
	this := ListUsageMetricsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListUsageMetricsResponse) GetItems() []UsageMetric {
	if o == nil || o.Items == nil {
		var ret []UsageMetric
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsageMetricsResponse) GetItemsOk() (*[]UsageMetric, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListUsageMetricsResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []UsageMetric and assigns it to the Items field.
func (o *ListUsageMetricsResponse) SetItems(v []UsageMetric) {
	o.Items = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListUsageMetricsResponse) GetNextToken() string {
	if o == nil || o.NextToken == nil {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsageMetricsResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || o.NextToken == nil {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListUsageMetricsResponse) HasNextToken() bool {
	if o != nil && o.NextToken != nil {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListUsageMetricsResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListUsageMetricsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.NextToken != nil {
		toSerialize["next_token"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableListUsageMetricsResponse struct {
	value *ListUsageMetricsResponse
	isSet bool
}

func (v NullableListUsageMetricsResponse) Get() *ListUsageMetricsResponse {
	return v.value
}

func (v *NullableListUsageMetricsResponse) Set(val *ListUsageMetricsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListUsageMetricsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListUsageMetricsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUsageMetricsResponse(val *ListUsageMetricsResponse) *NullableListUsageMetricsResponse {
	return &NullableListUsageMetricsResponse{value: val, isSet: true}
}

func (v NullableListUsageMetricsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUsageMetricsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}