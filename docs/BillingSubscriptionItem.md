# BillingSubscriptionItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingScheme** | **string** | Describes how to compute the price per period. Either per_unit or tiered. per_unit indicates that the fixed amount (specified in unit_amount or unit_amount_decimal) will be charged per unit in quantity (for prices with usage_type&#x3D;licensed), or per unit of total usage (for prices with usage_type&#x3D;metered). tiered indicates that the unit pricing will be computed using a tiering strategy as defined using the tiers and tiers_mode attributes. | 
**Id** | **string** |  | 
**PricingId** | **string** |  | 
**Quantity** | **int32** |  | 
**UnitAmount** | **int32** | The unit amount in cents to be charged, represented as a whole integer if possible. Only set if billing_scheme&#x3D;per_unit. | 

## Methods

### NewBillingSubscriptionItem

`func NewBillingSubscriptionItem(billingScheme string, id string, pricingId string, quantity int32, unitAmount int32, ) *BillingSubscriptionItem`

NewBillingSubscriptionItem instantiates a new BillingSubscriptionItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingSubscriptionItemWithDefaults

`func NewBillingSubscriptionItemWithDefaults() *BillingSubscriptionItem`

NewBillingSubscriptionItemWithDefaults instantiates a new BillingSubscriptionItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingScheme

`func (o *BillingSubscriptionItem) GetBillingScheme() string`

GetBillingScheme returns the BillingScheme field if non-nil, zero value otherwise.

### GetBillingSchemeOk

`func (o *BillingSubscriptionItem) GetBillingSchemeOk() (*string, bool)`

GetBillingSchemeOk returns a tuple with the BillingScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingScheme

`func (o *BillingSubscriptionItem) SetBillingScheme(v string)`

SetBillingScheme sets BillingScheme field to given value.


### GetId

`func (o *BillingSubscriptionItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingSubscriptionItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingSubscriptionItem) SetId(v string)`

SetId sets Id field to given value.


### GetPricingId

`func (o *BillingSubscriptionItem) GetPricingId() string`

GetPricingId returns the PricingId field if non-nil, zero value otherwise.

### GetPricingIdOk

`func (o *BillingSubscriptionItem) GetPricingIdOk() (*string, bool)`

GetPricingIdOk returns a tuple with the PricingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingId

`func (o *BillingSubscriptionItem) SetPricingId(v string)`

SetPricingId sets PricingId field to given value.


### GetQuantity

`func (o *BillingSubscriptionItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *BillingSubscriptionItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *BillingSubscriptionItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetUnitAmount

`func (o *BillingSubscriptionItem) GetUnitAmount() int32`

GetUnitAmount returns the UnitAmount field if non-nil, zero value otherwise.

### GetUnitAmountOk

`func (o *BillingSubscriptionItem) GetUnitAmountOk() (*int32, bool)`

GetUnitAmountOk returns a tuple with the UnitAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitAmount

`func (o *BillingSubscriptionItem) SetUnitAmount(v int32)`

SetUnitAmount sets UnitAmount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


