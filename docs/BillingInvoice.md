# BillingInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttemptCount** | Pointer to **int32** |  | [optional] 
**Attempted** | Pointer to **bool** |  | [optional] 
**BillingReason** | Pointer to **string** |  | [optional] 
**CreatedAt** | **string** |  | 
**Currency** | Pointer to **string** |  | [optional] 
**HostedInvoiceUrl** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InvoicePdf** | Pointer to **string** |  | [optional] 
**NextPaymentAttempt** | **string** |  | 
**PeriodEnd** | **string** |  | 
**PeriodStart** | **string** |  | 
**ReceiptNumber** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewBillingInvoice

`func NewBillingInvoice(createdAt string, nextPaymentAttempt string, periodEnd string, periodStart string, ) *BillingInvoice`

NewBillingInvoice instantiates a new BillingInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInvoiceWithDefaults

`func NewBillingInvoiceWithDefaults() *BillingInvoice`

NewBillingInvoiceWithDefaults instantiates a new BillingInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttemptCount

`func (o *BillingInvoice) GetAttemptCount() int32`

GetAttemptCount returns the AttemptCount field if non-nil, zero value otherwise.

### GetAttemptCountOk

`func (o *BillingInvoice) GetAttemptCountOk() (*int32, bool)`

GetAttemptCountOk returns a tuple with the AttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptCount

`func (o *BillingInvoice) SetAttemptCount(v int32)`

SetAttemptCount sets AttemptCount field to given value.

### HasAttemptCount

`func (o *BillingInvoice) HasAttemptCount() bool`

HasAttemptCount returns a boolean if a field has been set.

### GetAttempted

`func (o *BillingInvoice) GetAttempted() bool`

GetAttempted returns the Attempted field if non-nil, zero value otherwise.

### GetAttemptedOk

`func (o *BillingInvoice) GetAttemptedOk() (*bool, bool)`

GetAttemptedOk returns a tuple with the Attempted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempted

`func (o *BillingInvoice) SetAttempted(v bool)`

SetAttempted sets Attempted field to given value.

### HasAttempted

`func (o *BillingInvoice) HasAttempted() bool`

HasAttempted returns a boolean if a field has been set.

### GetBillingReason

`func (o *BillingInvoice) GetBillingReason() string`

GetBillingReason returns the BillingReason field if non-nil, zero value otherwise.

### GetBillingReasonOk

`func (o *BillingInvoice) GetBillingReasonOk() (*string, bool)`

GetBillingReasonOk returns a tuple with the BillingReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingReason

`func (o *BillingInvoice) SetBillingReason(v string)`

SetBillingReason sets BillingReason field to given value.

### HasBillingReason

`func (o *BillingInvoice) HasBillingReason() bool`

HasBillingReason returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BillingInvoice) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BillingInvoice) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BillingInvoice) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCurrency

`func (o *BillingInvoice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillingInvoice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillingInvoice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillingInvoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetHostedInvoiceUrl

`func (o *BillingInvoice) GetHostedInvoiceUrl() string`

GetHostedInvoiceUrl returns the HostedInvoiceUrl field if non-nil, zero value otherwise.

### GetHostedInvoiceUrlOk

`func (o *BillingInvoice) GetHostedInvoiceUrlOk() (*string, bool)`

GetHostedInvoiceUrlOk returns a tuple with the HostedInvoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostedInvoiceUrl

`func (o *BillingInvoice) SetHostedInvoiceUrl(v string)`

SetHostedInvoiceUrl sets HostedInvoiceUrl field to given value.

### HasHostedInvoiceUrl

`func (o *BillingInvoice) HasHostedInvoiceUrl() bool`

HasHostedInvoiceUrl returns a boolean if a field has been set.

### GetId

`func (o *BillingInvoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BillingInvoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BillingInvoice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BillingInvoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInvoicePdf

`func (o *BillingInvoice) GetInvoicePdf() string`

GetInvoicePdf returns the InvoicePdf field if non-nil, zero value otherwise.

### GetInvoicePdfOk

`func (o *BillingInvoice) GetInvoicePdfOk() (*string, bool)`

GetInvoicePdfOk returns a tuple with the InvoicePdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoicePdf

`func (o *BillingInvoice) SetInvoicePdf(v string)`

SetInvoicePdf sets InvoicePdf field to given value.

### HasInvoicePdf

`func (o *BillingInvoice) HasInvoicePdf() bool`

HasInvoicePdf returns a boolean if a field has been set.

### GetNextPaymentAttempt

`func (o *BillingInvoice) GetNextPaymentAttempt() string`

GetNextPaymentAttempt returns the NextPaymentAttempt field if non-nil, zero value otherwise.

### GetNextPaymentAttemptOk

`func (o *BillingInvoice) GetNextPaymentAttemptOk() (*string, bool)`

GetNextPaymentAttemptOk returns a tuple with the NextPaymentAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPaymentAttempt

`func (o *BillingInvoice) SetNextPaymentAttempt(v string)`

SetNextPaymentAttempt sets NextPaymentAttempt field to given value.


### GetPeriodEnd

`func (o *BillingInvoice) GetPeriodEnd() string`

GetPeriodEnd returns the PeriodEnd field if non-nil, zero value otherwise.

### GetPeriodEndOk

`func (o *BillingInvoice) GetPeriodEndOk() (*string, bool)`

GetPeriodEndOk returns a tuple with the PeriodEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodEnd

`func (o *BillingInvoice) SetPeriodEnd(v string)`

SetPeriodEnd sets PeriodEnd field to given value.


### GetPeriodStart

`func (o *BillingInvoice) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *BillingInvoice) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *BillingInvoice) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.


### GetReceiptNumber

`func (o *BillingInvoice) GetReceiptNumber() string`

GetReceiptNumber returns the ReceiptNumber field if non-nil, zero value otherwise.

### GetReceiptNumberOk

`func (o *BillingInvoice) GetReceiptNumberOk() (*string, bool)`

GetReceiptNumberOk returns a tuple with the ReceiptNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiptNumber

`func (o *BillingInvoice) SetReceiptNumber(v string)`

SetReceiptNumber sets ReceiptNumber field to given value.

### HasReceiptNumber

`func (o *BillingInvoice) HasReceiptNumber() bool`

HasReceiptNumber returns a boolean if a field has been set.

### GetStatus

`func (o *BillingInvoice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingInvoice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingInvoice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillingInvoice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotal

`func (o *BillingInvoice) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BillingInvoice) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BillingInvoice) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BillingInvoice) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


