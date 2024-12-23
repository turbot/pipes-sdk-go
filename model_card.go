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

// Card struct for Card
type Card struct {
	Brand    *string `json:"brand,omitempty"`
	Country  *string `json:"country,omitempty"`
	ExpMonth *int32  `json:"exp_month,omitempty"`
	ExpYear  *int32  `json:"exp_year,omitempty"`
	Funding  *string `json:"funding,omitempty"`
	Last4    *string `json:"last4,omitempty"`
}

// NewCard instantiates a new Card object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCard() *Card {
	this := Card{}
	return &this
}

// NewCardWithDefaults instantiates a new Card object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardWithDefaults() *Card {
	this := Card{}
	return &this
}

// GetBrand returns the Brand field value if set, zero value otherwise.
func (o *Card) GetBrand() string {
	if o == nil || o.Brand == nil {
		var ret string
		return ret
	}
	return *o.Brand
}

// GetBrandOk returns a tuple with the Brand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetBrandOk() (*string, bool) {
	if o == nil || o.Brand == nil {
		return nil, false
	}
	return o.Brand, true
}

// HasBrand returns a boolean if a field has been set.
func (o *Card) HasBrand() bool {
	if o != nil && o.Brand != nil {
		return true
	}

	return false
}

// SetBrand gets a reference to the given string and assigns it to the Brand field.
func (o *Card) SetBrand(v string) {
	o.Brand = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *Card) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *Card) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *Card) SetCountry(v string) {
	o.Country = &v
}

// GetExpMonth returns the ExpMonth field value if set, zero value otherwise.
func (o *Card) GetExpMonth() int32 {
	if o == nil || o.ExpMonth == nil {
		var ret int32
		return ret
	}
	return *o.ExpMonth
}

// GetExpMonthOk returns a tuple with the ExpMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetExpMonthOk() (*int32, bool) {
	if o == nil || o.ExpMonth == nil {
		return nil, false
	}
	return o.ExpMonth, true
}

// HasExpMonth returns a boolean if a field has been set.
func (o *Card) HasExpMonth() bool {
	if o != nil && o.ExpMonth != nil {
		return true
	}

	return false
}

// SetExpMonth gets a reference to the given int32 and assigns it to the ExpMonth field.
func (o *Card) SetExpMonth(v int32) {
	o.ExpMonth = &v
}

// GetExpYear returns the ExpYear field value if set, zero value otherwise.
func (o *Card) GetExpYear() int32 {
	if o == nil || o.ExpYear == nil {
		var ret int32
		return ret
	}
	return *o.ExpYear
}

// GetExpYearOk returns a tuple with the ExpYear field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetExpYearOk() (*int32, bool) {
	if o == nil || o.ExpYear == nil {
		return nil, false
	}
	return o.ExpYear, true
}

// HasExpYear returns a boolean if a field has been set.
func (o *Card) HasExpYear() bool {
	if o != nil && o.ExpYear != nil {
		return true
	}

	return false
}

// SetExpYear gets a reference to the given int32 and assigns it to the ExpYear field.
func (o *Card) SetExpYear(v int32) {
	o.ExpYear = &v
}

// GetFunding returns the Funding field value if set, zero value otherwise.
func (o *Card) GetFunding() string {
	if o == nil || o.Funding == nil {
		var ret string
		return ret
	}
	return *o.Funding
}

// GetFundingOk returns a tuple with the Funding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetFundingOk() (*string, bool) {
	if o == nil || o.Funding == nil {
		return nil, false
	}
	return o.Funding, true
}

// HasFunding returns a boolean if a field has been set.
func (o *Card) HasFunding() bool {
	if o != nil && o.Funding != nil {
		return true
	}

	return false
}

// SetFunding gets a reference to the given string and assigns it to the Funding field.
func (o *Card) SetFunding(v string) {
	o.Funding = &v
}

// GetLast4 returns the Last4 field value if set, zero value otherwise.
func (o *Card) GetLast4() string {
	if o == nil || o.Last4 == nil {
		var ret string
		return ret
	}
	return *o.Last4
}

// GetLast4Ok returns a tuple with the Last4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Card) GetLast4Ok() (*string, bool) {
	if o == nil || o.Last4 == nil {
		return nil, false
	}
	return o.Last4, true
}

// HasLast4 returns a boolean if a field has been set.
func (o *Card) HasLast4() bool {
	if o != nil && o.Last4 != nil {
		return true
	}

	return false
}

// SetLast4 gets a reference to the given string and assigns it to the Last4 field.
func (o *Card) SetLast4(v string) {
	o.Last4 = &v
}

func (o Card) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Brand != nil {
		toSerialize["brand"] = o.Brand
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.ExpMonth != nil {
		toSerialize["exp_month"] = o.ExpMonth
	}
	if o.ExpYear != nil {
		toSerialize["exp_year"] = o.ExpYear
	}
	if o.Funding != nil {
		toSerialize["funding"] = o.Funding
	}
	if o.Last4 != nil {
		toSerialize["last4"] = o.Last4
	}
	return json.Marshal(toSerialize)
}

type NullableCard struct {
	value *Card
	isSet bool
}

func (v NullableCard) Get() *Card {
	return v.value
}

func (v *NullableCard) Set(val *Card) {
	v.value = val
	v.isSet = true
}

func (v NullableCard) IsSet() bool {
	return v.isSet
}

func (v *NullableCard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCard(val *Card) *NullableCard {
	return &NullableCard{value: val, isSet: true}
}

func (v NullableCard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
