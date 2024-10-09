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

// BillingSubscription struct for BillingSubscription
type BillingSubscription struct {
	CancelAt               string                     `json:"cancel_at"`
	CancelAtPeriodEnd      string                     `json:"cancel_at_period_end"`
	CanceledAt             string                     `json:"canceled_at"`
	CreatedAt              string                     `json:"created_at"`
	Currency               string                     `json:"currency"`
	CurrentPeriodEnd       string                     `json:"current_period_end"`
	CurrentPeriodStart     string                     `json:"current_period_start"`
	DefaultPaymentMethodId *string                    `json:"default_payment_method_id,omitempty"`
	Id                     string                     `json:"id"`
	Items                  *[]BillingSubscriptionItem `json:"items,omitempty"`
	Status                 *string                    `json:"status,omitempty"`
}

// NewBillingSubscription instantiates a new BillingSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingSubscription(cancelAt string, cancelAtPeriodEnd string, canceledAt string, createdAt string, currency string, currentPeriodEnd string, currentPeriodStart string, id string) *BillingSubscription {
	this := BillingSubscription{}
	this.CancelAt = cancelAt
	this.CancelAtPeriodEnd = cancelAtPeriodEnd
	this.CanceledAt = canceledAt
	this.CreatedAt = createdAt
	this.Currency = currency
	this.CurrentPeriodEnd = currentPeriodEnd
	this.CurrentPeriodStart = currentPeriodStart
	this.Id = id
	return &this
}

// NewBillingSubscriptionWithDefaults instantiates a new BillingSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingSubscriptionWithDefaults() *BillingSubscription {
	this := BillingSubscription{}
	return &this
}

// GetCancelAt returns the CancelAt field value
func (o *BillingSubscription) GetCancelAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CancelAt
}

// GetCancelAtOk returns a tuple with the CancelAt field value
// and a boolean to check if the value has been set.
func (o *BillingSubscription) GetCancelAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CancelAt, true
}

// SetCancelAt sets field value
func (o *BillingSubscription) SetCancelAt(v string) {
	o.CancelAt = v
}

// GetCancelAtPeriodEnd returns the CancelAtPeriodEnd field value
func (o *BillingSubscription) GetCancelAtPeriodEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CancelAtPeriodEnd
}

// GetCancelAtPeriodEndOk returns a tuple with the CancelAtPeriodEnd field value
// and a boolean to check if the value has been set.
func (o *BillingSubscription) GetCancelAtPeriodEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CancelAtPeriodEnd, true
}

// SetCancelAtPeriodEnd sets field value
func (o *BillingSubscription) SetCancelAtPeriodEnd(v string) {
	o.CancelAtPeriodEnd = v
}

// GetCanceledAt returns the CanceledAt field value
func (o *BillingSubscription) GetCanceledAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CanceledAt
}

// GetCanceledAtOk returns a tuple with the CanceledAt field value
// and a boolean to check if the value has been set.
func (o *BillingSubscription) GetCanceledAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanceledAt, true
}

// SetCanceledAt sets field value
func (o *BillingSubscription) SetCanceledAt(v string) {
	o.CanceledAt = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *BillingSubscription) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *BillingSubscription) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *BillingSubscription) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCurrency returns the Currency field value
func (o *BillingSubscription) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *BillingSubscription) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *BillingSubscription) SetCurrency(v string) {
	o.Currency = v
}

// GetCurrentPeriodEnd returns the CurrentPeriodEnd field value
func (o *BillingSubscription) GetCurrentPeriodEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentPeriodEnd
}

// GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field value
// and a boolean to check if the value has been set.
func (o *BillingSubscription) GetCurrentPeriodEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPeriodEnd, true
}

// SetCurrentPeriodEnd sets field value
func (o *BillingSubscription) SetCurrentPeriodEnd(v string) {
	o.CurrentPeriodEnd = v
}

// GetCurrentPeriodStart returns the CurrentPeriodStart field value
func (o *BillingSubscription) GetCurrentPeriodStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrentPeriodStart
}

// GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field value
// and a boolean to check if the value has been set.
func (o *BillingSubscription) GetCurrentPeriodStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPeriodStart, true
}

// SetCurrentPeriodStart sets field value
func (o *BillingSubscription) SetCurrentPeriodStart(v string) {
	o.CurrentPeriodStart = v
}

// GetDefaultPaymentMethodId returns the DefaultPaymentMethodId field value if set, zero value otherwise.
func (o *BillingSubscription) GetDefaultPaymentMethodId() string {
	if o == nil || o.DefaultPaymentMethodId == nil {
		var ret string
		return ret
	}
	return *o.DefaultPaymentMethodId
}

// GetDefaultPaymentMethodIdOk returns a tuple with the DefaultPaymentMethodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingSubscription) GetDefaultPaymentMethodIdOk() (*string, bool) {
	if o == nil || o.DefaultPaymentMethodId == nil {
		return nil, false
	}
	return o.DefaultPaymentMethodId, true
}

// HasDefaultPaymentMethodId returns a boolean if a field has been set.
func (o *BillingSubscription) HasDefaultPaymentMethodId() bool {
	if o != nil && o.DefaultPaymentMethodId != nil {
		return true
	}

	return false
}

// SetDefaultPaymentMethodId gets a reference to the given string and assigns it to the DefaultPaymentMethodId field.
func (o *BillingSubscription) SetDefaultPaymentMethodId(v string) {
	o.DefaultPaymentMethodId = &v
}

// GetId returns the Id field value
func (o *BillingSubscription) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BillingSubscription) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BillingSubscription) SetId(v string) {
	o.Id = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *BillingSubscription) GetItems() []BillingSubscriptionItem {
	if o == nil || o.Items == nil {
		var ret []BillingSubscriptionItem
		return ret
	}
	return *o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingSubscription) GetItemsOk() (*[]BillingSubscriptionItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *BillingSubscription) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []BillingSubscriptionItem and assigns it to the Items field.
func (o *BillingSubscription) SetItems(v []BillingSubscriptionItem) {
	o.Items = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BillingSubscription) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingSubscription) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BillingSubscription) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BillingSubscription) SetStatus(v string) {
	o.Status = &v
}

func (o BillingSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cancel_at"] = o.CancelAt
	}
	if true {
		toSerialize["cancel_at_period_end"] = o.CancelAtPeriodEnd
	}
	if true {
		toSerialize["canceled_at"] = o.CanceledAt
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["currency"] = o.Currency
	}
	if true {
		toSerialize["current_period_end"] = o.CurrentPeriodEnd
	}
	if true {
		toSerialize["current_period_start"] = o.CurrentPeriodStart
	}
	if o.DefaultPaymentMethodId != nil {
		toSerialize["default_payment_method_id"] = o.DefaultPaymentMethodId
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableBillingSubscription struct {
	value *BillingSubscription
	isSet bool
}

func (v NullableBillingSubscription) Get() *BillingSubscription {
	return v.value
}

func (v *NullableBillingSubscription) Set(val *BillingSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingSubscription(val *BillingSubscription) *NullableBillingSubscription {
	return &NullableBillingSubscription{value: val, isSet: true}
}

func (v NullableBillingSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
