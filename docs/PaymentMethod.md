# PaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Card** | Pointer to [**Card**](Card.md) |  | [optional] 
**Default** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**UsBankAccount** | Pointer to [**USBankAccount**](USBankAccount.md) |  | [optional] 

## Methods

### NewPaymentMethod

`func NewPaymentMethod(type_ string, ) *PaymentMethod`

NewPaymentMethod instantiates a new PaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodWithDefaults

`func NewPaymentMethodWithDefaults() *PaymentMethod`

NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCard

`func (o *PaymentMethod) GetCard() Card`

GetCard returns the Card field if non-nil, zero value otherwise.

### GetCardOk

`func (o *PaymentMethod) GetCardOk() (*Card, bool)`

GetCardOk returns a tuple with the Card field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCard

`func (o *PaymentMethod) SetCard(v Card)`

SetCard sets Card field to given value.

### HasCard

`func (o *PaymentMethod) HasCard() bool`

HasCard returns a boolean if a field has been set.

### GetDefault

`func (o *PaymentMethod) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *PaymentMethod) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *PaymentMethod) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *PaymentMethod) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetId

`func (o *PaymentMethod) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethod) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethod) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentMethod) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *PaymentMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethod) SetType(v string)`

SetType sets Type field to given value.


### GetUsBankAccount

`func (o *PaymentMethod) GetUsBankAccount() USBankAccount`

GetUsBankAccount returns the UsBankAccount field if non-nil, zero value otherwise.

### GetUsBankAccountOk

`func (o *PaymentMethod) GetUsBankAccountOk() (*USBankAccount, bool)`

GetUsBankAccountOk returns a tuple with the UsBankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsBankAccount

`func (o *PaymentMethod) SetUsBankAccount(v USBankAccount)`

SetUsBankAccount sets UsBankAccount field to given value.

### HasUsBankAccount

`func (o *PaymentMethod) HasUsBankAccount() bool`

HasUsBankAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


