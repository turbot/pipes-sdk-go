# WorkspaceCommandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | [**WorkspaceCommandAction**](WorkspaceCommandAction.md) |  | 

## Methods

### NewWorkspaceCommandRequest

`func NewWorkspaceCommandRequest(command WorkspaceCommandAction, ) *WorkspaceCommandRequest`

NewWorkspaceCommandRequest instantiates a new WorkspaceCommandRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceCommandRequestWithDefaults

`func NewWorkspaceCommandRequestWithDefaults() *WorkspaceCommandRequest`

NewWorkspaceCommandRequestWithDefaults instantiates a new WorkspaceCommandRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *WorkspaceCommandRequest) GetCommand() WorkspaceCommandAction`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *WorkspaceCommandRequest) GetCommandOk() (*WorkspaceCommandAction, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *WorkspaceCommandRequest) SetCommand(v WorkspaceCommandAction)`

SetCommand sets Command field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


