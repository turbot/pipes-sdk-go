# ListTriggersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ModTriggerInfo**](ModTriggerInfo.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewListTriggersResponse

`func NewListTriggersResponse() *ListTriggersResponse`

NewListTriggersResponse instantiates a new ListTriggersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTriggersResponseWithDefaults

`func NewListTriggersResponseWithDefaults() *ListTriggersResponse`

NewListTriggersResponseWithDefaults instantiates a new ListTriggersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListTriggersResponse) GetItems() []ModTriggerInfo`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListTriggersResponse) GetItemsOk() (*[]ModTriggerInfo, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListTriggersResponse) SetItems(v []ModTriggerInfo)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListTriggersResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextToken

`func (o *ListTriggersResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListTriggersResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListTriggersResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListTriggersResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


