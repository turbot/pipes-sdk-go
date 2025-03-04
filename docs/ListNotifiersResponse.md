# ListNotifiersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Notifier**](Notifier.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewListNotifiersResponse

`func NewListNotifiersResponse() *ListNotifiersResponse`

NewListNotifiersResponse instantiates a new ListNotifiersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNotifiersResponseWithDefaults

`func NewListNotifiersResponseWithDefaults() *ListNotifiersResponse`

NewListNotifiersResponseWithDefaults instantiates a new ListNotifiersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListNotifiersResponse) GetItems() []Notifier`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListNotifiersResponse) GetItemsOk() (*[]Notifier, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListNotifiersResponse) SetItems(v []Notifier)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListNotifiersResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextToken

`func (o *ListNotifiersResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListNotifiersResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListNotifiersResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListNotifiersResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


