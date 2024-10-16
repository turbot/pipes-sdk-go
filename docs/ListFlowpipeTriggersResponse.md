# ListFlowpipeTriggersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]WorkspaceModTrigger**](WorkspaceModTrigger.md) |  | [optional] 
**NextToken** | Pointer to **string** |  | [optional] 

## Methods

### NewListFlowpipeTriggersResponse

`func NewListFlowpipeTriggersResponse() *ListFlowpipeTriggersResponse`

NewListFlowpipeTriggersResponse instantiates a new ListFlowpipeTriggersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFlowpipeTriggersResponseWithDefaults

`func NewListFlowpipeTriggersResponseWithDefaults() *ListFlowpipeTriggersResponse`

NewListFlowpipeTriggersResponseWithDefaults instantiates a new ListFlowpipeTriggersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListFlowpipeTriggersResponse) GetItems() []WorkspaceModTrigger`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListFlowpipeTriggersResponse) GetItemsOk() (*[]WorkspaceModTrigger, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListFlowpipeTriggersResponse) SetItems(v []WorkspaceModTrigger)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListFlowpipeTriggersResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetNextToken

`func (o *ListFlowpipeTriggersResponse) GetNextToken() string`

GetNextToken returns the NextToken field if non-nil, zero value otherwise.

### GetNextTokenOk

`func (o *ListFlowpipeTriggersResponse) GetNextTokenOk() (*string, bool)`

GetNextTokenOk returns a tuple with the NextToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextToken

`func (o *ListFlowpipeTriggersResponse) SetNextToken(v string)`

SetNextToken sets NextToken field to given value.

### HasNextToken

`func (o *ListFlowpipeTriggersResponse) HasNextToken() bool`

HasNextToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


