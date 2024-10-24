# WorkspaceProcessCommandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | [**WorkspaceProcessCommandAction**](WorkspaceProcessCommandAction.md) |  | 

## Methods

### NewWorkspaceProcessCommandRequest

`func NewWorkspaceProcessCommandRequest(command WorkspaceProcessCommandAction, ) *WorkspaceProcessCommandRequest`

NewWorkspaceProcessCommandRequest instantiates a new WorkspaceProcessCommandRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceProcessCommandRequestWithDefaults

`func NewWorkspaceProcessCommandRequestWithDefaults() *WorkspaceProcessCommandRequest`

NewWorkspaceProcessCommandRequestWithDefaults instantiates a new WorkspaceProcessCommandRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *WorkspaceProcessCommandRequest) GetCommand() WorkspaceProcessCommandAction`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *WorkspaceProcessCommandRequest) GetCommandOk() (*WorkspaceProcessCommandAction, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *WorkspaceProcessCommandRequest) SetCommand(v WorkspaceProcessCommandAction)`

SetCommand sets Command field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


