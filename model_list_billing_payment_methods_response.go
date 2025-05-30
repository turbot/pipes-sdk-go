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

// ListBillingPaymentMethodsResponse struct for ListBillingPaymentMethodsResponse
type ListBillingPaymentMethodsResponse struct {
	Items     *[]PaymentMethod `json:"items,omitempty"`
	NextToken *string          `json:"next_token,omitempty"`
}

// NewListBillingPaymentMethodsResponse instantiates a new ListBillingPaymentMethodsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListBillingPaymentMethodsResponse() *ListBillingPaymentMethodsResponse {
	this := ListBillingPaymentMethodsResponse{}
	return &this
}

// NewListBillingPaymentMethodsResponseWithDefaults instantiates a new ListBillingPaymentMethodsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListBillingPaymentMethodsResponseWithDefaults() *ListBillingPaymentMethodsResponse {
	this := ListBillingPaymentMethodsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListBillingPaymentMethodsResponse) GetItems() []PaymentMethod {
	if o == nil || o.Items == nil {
		var ret []PaymentMethod
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBillingPaymentMethodsResponse) GetItemsOk() (*[]PaymentMethod, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListBillingPaymentMethodsResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []PaymentMethod and assigns it to the Items field.
func (o *ListBillingPaymentMethodsResponse) SetItems(v []PaymentMethod) {
	o.Items = &v
}

// GetNextToken returns the NextToken field value if set, zero value otherwise.
func (o *ListBillingPaymentMethodsResponse) GetNextToken() string {
	if o == nil || o.NextToken == nil {
		var ret string
		return ret
	}
	return *o.NextToken
}

// GetNextTokenOk returns a tuple with the NextToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBillingPaymentMethodsResponse) GetNextTokenOk() (*string, bool) {
	if o == nil || o.NextToken == nil {
		return nil, false
	}
	return o.NextToken, true
}

// HasNextToken returns a boolean if a field has been set.
func (o *ListBillingPaymentMethodsResponse) HasNextToken() bool {
	if o != nil && o.NextToken != nil {
		return true
	}

	return false
}

// SetNextToken gets a reference to the given string and assigns it to the NextToken field.
func (o *ListBillingPaymentMethodsResponse) SetNextToken(v string) {
	o.NextToken = &v
}

func (o ListBillingPaymentMethodsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.NextToken != nil {
		toSerialize["next_token"] = o.NextToken
	}
	return json.Marshal(toSerialize)
}

type NullableListBillingPaymentMethodsResponse struct {
	value *ListBillingPaymentMethodsResponse
	isSet bool
}

func (v NullableListBillingPaymentMethodsResponse) Get() *ListBillingPaymentMethodsResponse {
	return v.value
}

func (v *NullableListBillingPaymentMethodsResponse) Set(val *ListBillingPaymentMethodsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListBillingPaymentMethodsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListBillingPaymentMethodsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListBillingPaymentMethodsResponse(val *ListBillingPaymentMethodsResponse) *NullableListBillingPaymentMethodsResponse {
	return &NullableListBillingPaymentMethodsResponse{value: val, isSet: true}
}

func (v NullableListBillingPaymentMethodsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListBillingPaymentMethodsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
