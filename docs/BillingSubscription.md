# BillingSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelAt** | **string** |  | 
**CancelAtPeriodEnd** | **string** |  | 
**CanceledAt** | **string** |  | 
**CreatedAt** | **string** |  | 
**Currency** | **string** |  | 
**CurrentPeriodEnd** | **string** |  | 
**CurrentPeriodStart** | **string** |  | 
**DefaultPaymentMethodId** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Items** | Pointer to [**[]BillingSubscriptionItem**](BillingSubscriptionItem.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingSubscription

`func NewBillingSubscription(cancelAt string, cancelAtPeriodEnd string, canceledAt string, createdAt string, currency string, currentPeriodEnd string, currentPeriodStart string, id string, ) *BillingSubscription`

NewBillingSubscription instantiates a new BillingSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingSubscriptionWithDefaults

`func NewBillingSubscriptionWithDefaults() *BillingSubscription`

NewBillingSubscriptionWithDefaults instantiates a new BillingSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelAt

`func (o *BillingSubscription) GetCancelAt() string`

GetCancelAt returns the CancelAt field if non-nil, zero value otherwise.

### GetCancelAtOk

`func (o *BillingSubscription) GetCancelAtOk() (*string, bool)`

GetCancelAtOk returns a tuple with the CancelAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAt

`func (o *BillingSubscription) SetCancelAt(v string)`

SetCancelAt sets CancelAt field to given value.


### GetCancelAtPeriodEnd

`func (o *BillingSubscription) GetCancelAtPeriodEnd() string`

GetCancelAtPeriodEnd returns the CancelAtPeriodEnd field if non-nil, zero value otherwise.

### GetCancelAtPeriodEndOk

`func (o *BillingSubscription) GetCancelAtPeriodEndOk() (*string, bool)`

GetCancelAtPeriodEndOk returns a tuple with the CancelAtPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAtPeriodEnd

`func (o *BillingSubscription) SetCancelAtPeriodEnd(v string)`

SetCancelAtPeriodEnd sets CancelAtPeriodEnd field to given value.


### GetCanceledAt

`func (o *BillingSubscription) GetCanceledAt() string`

GetCanceledAt returns the CanceledAt field if non-nil, zero value otherwise.

### GetCanceledAtOk

`func (o *BillingSubscription) GetCanceledAtOk() (*string, bool)`

GetCanceledAtOk returns a tuple with the CanceledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanceledAt

`func (o *BillingSubscription) SetCanceledAt(v string)`

SetCanceledAt sets CanceledAt field to given value.


### GetCreatedAt

`func (o *BillingSubscription) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingSubscription) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingSubscription) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrency

`func (o *BillingSubscription) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingSubscription) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingSubscription) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCurrentPeriodEnd

`func (o *BillingSubscription) GetCurrentPeriodEnd() string`

GetCurrentPeriodEnd returns the CurrentPeriodEnd field if non-nil, zero value otherwise.

### GetCurrentPeriodEndOk

`func (o *BillingSubscription) GetCurrentPeriodEndOk() (*string, bool)`

GetCurrentPeriodEndOk returns a tuple with the CurrentPeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodEnd

`func (o *BillingSubscription) SetCurrentPeriodEnd(v string)`

SetCurrentPeriodEnd sets CurrentPeriodEnd field to given value.


### GetCurrentPeriodStart

`func (o *BillingSubscription) GetCurrentPeriodStart() string`

GetCurrentPeriodStart returns the CurrentPeriodStart field if non-nil, zero value otherwise.

### GetCurrentPeriodStartOk

`func (o *BillingSubscription) GetCurrentPeriodStartOk() (*string, bool)`

GetCurrentPeriodStartOk returns a tuple with the CurrentPeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodStart

`func (o *BillingSubscription) SetCurrentPeriodStart(v string)`

SetCurrentPeriodStart sets CurrentPeriodStart field to given value.


### GetDefaultPaymentMethodId

`func (o *BillingSubscription) GetDefaultPaymentMethodId() string`

GetDefaultPaymentMethodId returns the DefaultPaymentMethodId field if non-nil, zero value otherwise.

### GetDefaultPaymentMethodIdOk

`func (o *BillingSubscription) GetDefaultPaymentMethodIdOk() (*string, bool)`

GetDefaultPaymentMethodIdOk returns a tuple with the DefaultPaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPaymentMethodId

`func (o *BillingSubscription) SetDefaultPaymentMethodId(v string)`

SetDefaultPaymentMethodId sets DefaultPaymentMethodId field to given value.

### HasDefaultPaymentMethodId

`func (o *BillingSubscription) HasDefaultPaymentMethodId() bool`

HasDefaultPaymentMethodId returns a boolean if a field has been set.

### GetId

`func (o *BillingSubscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingSubscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingSubscription) SetId(v string)`

SetId sets Id field to given value.


### GetItems

`func (o *BillingSubscription) GetItems() []BillingSubscriptionItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *BillingSubscription) GetItemsOk() (*[]BillingSubscriptionItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *BillingSubscription) SetItems(v []BillingSubscriptionItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *BillingSubscription) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetStatus

`func (o *BillingSubscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingSubscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingSubscription) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillingSubscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


