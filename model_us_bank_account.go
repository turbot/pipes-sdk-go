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

// USBankAccount struct for USBankAccount
type USBankAccount struct {
	AccountHolderType *string `json:"account_holder_type,omitempty"`
	Last4             *string `json:"last4,omitempty"`
}

// NewUSBankAccount instantiates a new USBankAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUSBankAccount() *USBankAccount {
	this := USBankAccount{}
	return &this
}

// NewUSBankAccountWithDefaults instantiates a new USBankAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUSBankAccountWithDefaults() *USBankAccount {
	this := USBankAccount{}
	return &this
}

// GetAccountHolderType returns the AccountHolderType field value if set, zero value otherwise.
func (o *USBankAccount) GetAccountHolderType() string {
	if o == nil || o.AccountHolderType == nil {
		var ret string
		return ret
	}
	return *o.AccountHolderType
}

// GetAccountHolderTypeOk returns a tuple with the AccountHolderType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *USBankAccount) GetAccountHolderTypeOk() (*string, bool) {
	if o == nil || o.AccountHolderType == nil {
		return nil, false
	}
	return o.AccountHolderType, true
}

// HasAccountHolderType returns a boolean if a field has been set.
func (o *USBankAccount) HasAccountHolderType() bool {
	if o != nil && o.AccountHolderType != nil {
		return true
	}

	return false
}

// SetAccountHolderType gets a reference to the given string and assigns it to the AccountHolderType field.
func (o *USBankAccount) SetAccountHolderType(v string) {
	o.AccountHolderType = &v
}

// GetLast4 returns the Last4 field value if set, zero value otherwise.
func (o *USBankAccount) GetLast4() string {
	if o == nil || o.Last4 == nil {
		var ret string
		return ret
	}
	return *o.Last4
}

// GetLast4Ok returns a tuple with the Last4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *USBankAccount) GetLast4Ok() (*string, bool) {
	if o == nil || o.Last4 == nil {
		return nil, false
	}
	return o.Last4, true
}

// HasLast4 returns a boolean if a field has been set.
func (o *USBankAccount) HasLast4() bool {
	if o != nil && o.Last4 != nil {
		return true
	}

	return false
}

// SetLast4 gets a reference to the given string and assigns it to the Last4 field.
func (o *USBankAccount) SetLast4(v string) {
	o.Last4 = &v
}

func (o USBankAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountHolderType != nil {
		toSerialize["account_holder_type"] = o.AccountHolderType
	}
	if o.Last4 != nil {
		toSerialize["last4"] = o.Last4
	}
	return json.Marshal(toSerialize)
}

type NullableUSBankAccount struct {
	value *USBankAccount
	isSet bool
}

func (v NullableUSBankAccount) Get() *USBankAccount {
	return v.value
}

func (v *NullableUSBankAccount) Set(val *USBankAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableUSBankAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableUSBankAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUSBankAccount(val *USBankAccount) *NullableUSBankAccount {
	return &NullableUSBankAccount{value: val, isSet: true}
}

func (v NullableUSBankAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUSBankAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
