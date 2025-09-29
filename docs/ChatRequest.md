# ChatRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversationId** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **string** | TODO: should provider be inferred from the model? | [optional] 

## Methods

### NewChatRequest

`func NewChatRequest() *ChatRequest`

NewChatRequest instantiates a new ChatRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatRequestWithDefaults

`func NewChatRequestWithDefaults() *ChatRequest`

NewChatRequestWithDefaults instantiates a new ChatRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversationId

`func (o *ChatRequest) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *ChatRequest) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *ChatRequest) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *ChatRequest) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### GetMessage

`func (o *ChatRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ChatRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ChatRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ChatRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetModel

`func (o *ChatRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChatRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChatRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ChatRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetProvider

`func (o *ChatRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ChatRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ChatRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ChatRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


