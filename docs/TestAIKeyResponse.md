# TestAIKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**AIKeyTestStatus**](AIKeyTestStatus.md) |  | [optional] 

## Methods

### NewTestAIKeyResponse

`func NewTestAIKeyResponse() *TestAIKeyResponse`

NewTestAIKeyResponse instantiates a new TestAIKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestAIKeyResponseWithDefaults

`func NewTestAIKeyResponseWithDefaults() *TestAIKeyResponse`

NewTestAIKeyResponseWithDefaults instantiates a new TestAIKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *TestAIKeyResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *TestAIKeyResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *TestAIKeyResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *TestAIKeyResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *TestAIKeyResponse) GetStatus() AIKeyTestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestAIKeyResponse) GetStatusOk() (*AIKeyTestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestAIKeyResponse) SetStatus(v AIKeyTestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TestAIKeyResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


