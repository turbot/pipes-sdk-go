# WorkspaceNotifierCommandResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | [**WorkspaceNotifierCommandAction**](WorkspaceNotifierCommandAction.md) |  | 
**ProcessId** | **string** |  | 

## Methods

### NewWorkspaceNotifierCommandResponse

`func NewWorkspaceNotifierCommandResponse(command WorkspaceNotifierCommandAction, processId string, ) *WorkspaceNotifierCommandResponse`

NewWorkspaceNotifierCommandResponse instantiates a new WorkspaceNotifierCommandResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceNotifierCommandResponseWithDefaults

`func NewWorkspaceNotifierCommandResponseWithDefaults() *WorkspaceNotifierCommandResponse`

NewWorkspaceNotifierCommandResponseWithDefaults instantiates a new WorkspaceNotifierCommandResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *WorkspaceNotifierCommandResponse) GetCommand() WorkspaceNotifierCommandAction`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *WorkspaceNotifierCommandResponse) GetCommandOk() (*WorkspaceNotifierCommandAction, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *WorkspaceNotifierCommandResponse) SetCommand(v WorkspaceNotifierCommandAction)`

SetCommand sets Command field to given value.


### GetProcessId

`func (o *WorkspaceNotifierCommandResponse) GetProcessId() string`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *WorkspaceNotifierCommandResponse) GetProcessIdOk() (*string, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *WorkspaceNotifierCommandResponse) SetProcessId(v string)`

SetProcessId sets ProcessId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


