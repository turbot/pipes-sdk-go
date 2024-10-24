# USBankAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountHolderType** | Pointer to **string** |  | [optional] 
**Last4** | Pointer to **string** |  | [optional] 

## Methods

### NewUSBankAccount

`func NewUSBankAccount() *USBankAccount`

NewUSBankAccount instantiates a new USBankAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUSBankAccountWithDefaults

`func NewUSBankAccountWithDefaults() *USBankAccount`

NewUSBankAccountWithDefaults instantiates a new USBankAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountHolderType

`func (o *USBankAccount) GetAccountHolderType() string`

GetAccountHolderType returns the AccountHolderType field if non-nil, zero value otherwise.

### GetAccountHolderTypeOk

`func (o *USBankAccount) GetAccountHolderTypeOk() (*string, bool)`

GetAccountHolderTypeOk returns a tuple with the AccountHolderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderType

`func (o *USBankAccount) SetAccountHolderType(v string)`

SetAccountHolderType sets AccountHolderType field to given value.

### HasAccountHolderType

`func (o *USBankAccount) HasAccountHolderType() bool`

HasAccountHolderType returns a boolean if a field has been set.

### GetLast4

`func (o *USBankAccount) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *USBankAccount) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *USBankAccount) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *USBankAccount) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


