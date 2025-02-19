# TriggerCommandResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | [**TriggerCommandAction**](TriggerCommandAction.md) |  | 
**ProcessId** | Pointer to **string** |  | [optional] 

## Methods

### NewTriggerCommandResponse

`func NewTriggerCommandResponse(command TriggerCommandAction, ) *TriggerCommandResponse`

NewTriggerCommandResponse instantiates a new TriggerCommandResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerCommandResponseWithDefaults

`func NewTriggerCommandResponseWithDefaults() *TriggerCommandResponse`

NewTriggerCommandResponseWithDefaults instantiates a new TriggerCommandResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *TriggerCommandResponse) GetCommand() TriggerCommandAction`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *TriggerCommandResponse) GetCommandOk() (*TriggerCommandAction, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *TriggerCommandResponse) SetCommand(v TriggerCommandAction)`

SetCommand sets Command field to given value.


### GetProcessId

`func (o *TriggerCommandResponse) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *TriggerCommandResponse) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *TriggerCommandResponse) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *TriggerCommandResponse) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


