# TriggerCommandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **map[string]interface{}** |  | [optional] 
**Command** | [**TriggerCommandAction**](TriggerCommandAction.md) |  | 

## Methods

### NewTriggerCommandRequest

`func NewTriggerCommandRequest(command TriggerCommandAction, ) *TriggerCommandRequest`

NewTriggerCommandRequest instantiates a new TriggerCommandRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerCommandRequestWithDefaults

`func NewTriggerCommandRequestWithDefaults() *TriggerCommandRequest`

NewTriggerCommandRequestWithDefaults instantiates a new TriggerCommandRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *TriggerCommandRequest) GetArgs() map[string]interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *TriggerCommandRequest) GetArgsOk() (*map[string]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *TriggerCommandRequest) SetArgs(v map[string]interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *TriggerCommandRequest) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetCommand

`func (o *TriggerCommandRequest) GetCommand() TriggerCommandAction`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *TriggerCommandRequest) GetCommandOk() (*TriggerCommandAction, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *TriggerCommandRequest) SetCommand(v TriggerCommandAction)`

SetCommand sets Command field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


