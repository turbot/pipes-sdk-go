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

// ListFlowpipePipelinesResponse struct for ListFlowpipePipelinesResponse
type ListFlowpipePipelinesResponse struct {
	Items     *[]WorkspaceModPipeline `json:"items,omitempty"`
	NextToken *string                 `json:"next_token,omitempty"`
}

// NewListFlowpipePipelinesResponse instantiates a new ListFlowpipePipelinesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListFlowpipePipelinesResponse() *ListFlowpipePipelinesResponse {
	this := ListFlowpipePipelinesResponse{}
	return &this
}

// NewListFlowpipePipelinesResponseWithDefaults instantiates a new ListFlowpipePipelinesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListFlowpipePipelinesResponseWithDefaults() *ListFlowpipePipelinesResponse {
	this := ListFlowpipePipelinesResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListFlowpipePipelinesResponse) GetItems() []WorkspaceModPipeline {
	if o == nil || o.Items == nil {
		var ret []WorkspaceModPipeline
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFlowpipePipelinesResponse) GetItemsOk() (*[]WorkspaceModPipeline, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListFlowpipePipelinesResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []WorkspaceModPipeline and assigns it to the Items field.
func (o *ListFlowpipePipelinesResponse) SetItems(v []WorkspaceModPipeline) {
	o.Items = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListFlowpipePipelinesResponse) GetNextToken() string {
	if o == nil || o.NextToken == nil {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListFlowpipePipelinesResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || o.NextToken == nil {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListFlowpipePipelinesResponse) HasNextToken() bool {
	if o != nil && o.NextToken != nil {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListFlowpipePipelinesResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListFlowpipePipelinesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.NextToken != nil {
		toSerialize["next_token"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableListFlowpipePipelinesResponse struct {
	value *ListFlowpipePipelinesResponse
	isSet bool
}

func (v NullableListFlowpipePipelinesResponse) Get() *ListFlowpipePipelinesResponse {
	return v.value
}

func (v *NullableListFlowpipePipelinesResponse) Set(val *ListFlowpipePipelinesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListFlowpipePipelinesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListFlowpipePipelinesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListFlowpipePipelinesResponse(val *ListFlowpipePipelinesResponse) *NullableListFlowpipePipelinesResponse {
	return &NullableListFlowpipePipelinesResponse{value: val, isSet: true}
}

func (v NullableListFlowpipePipelinesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListFlowpipePipelinesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}