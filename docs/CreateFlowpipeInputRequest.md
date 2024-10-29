# CreateFlowpipeInputRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutionId** | **string** |  | 
**Inputs** | Pointer to **map[string]interface{}** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**NotifierName** | **string** |  | 
**Overrides** | Pointer to **map[string]interface{}** |  | [optional] 
**PipelineExecutionId** | Pointer to **string** |  | [optional] 
**StepExecutionId** | **string** |  | 
**StepType** | [**FlowpipeInputStepType**](FlowpipeInputStepType.md) |  | 

## Methods

### NewCreateFlowpipeInputRequest

`func NewCreateFlowpipeInputRequest(executionId string, notifierName string, stepExecutionId string, stepType FlowpipeInputStepType, ) *CreateFlowpipeInputRequest`

NewCreateFlowpipeInputRequest instantiates a new CreateFlowpipeInputRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFlowpipeInputRequestWithDefaults

`func NewCreateFlowpipeInputRequestWithDefaults() *CreateFlowpipeInputRequest`

NewCreateFlowpipeInputRequestWithDefaults instantiates a new CreateFlowpipeInputRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutionId

`func (o *CreateFlowpipeInputRequest) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *CreateFlowpipeInputRequest) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *CreateFlowpipeInputRequest) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.


### GetInputs

`func (o *CreateFlowpipeInputRequest) GetInputs() map[string]interface{}`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *CreateFlowpipeInputRequest) GetInputsOk() (*map[string]interface{}, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *CreateFlowpipeInputRequest) SetInputs(v map[string]interface{})`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *CreateFlowpipeInputRequest) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetMessage

`func (o *CreateFlowpipeInputRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateFlowpipeInputRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateFlowpipeInputRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateFlowpipeInputRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNotifierName

`func (o *CreateFlowpipeInputRequest) GetNotifierName() string`

GetNotifierName returns the NotifierName field if non-nil, zero value otherwise.

### GetNotifierNameOk

`func (o *CreateFlowpipeInputRequest) GetNotifierNameOk() (*string, bool)`

GetNotifierNameOk returns a tuple with the NotifierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifierName

`func (o *CreateFlowpipeInputRequest) SetNotifierName(v string)`

SetNotifierName sets NotifierName field to given value.


### GetOverrides

`func (o *CreateFlowpipeInputRequest) GetOverrides() map[string]interface{}`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *CreateFlowpipeInputRequest) GetOverridesOk() (*map[string]interface{}, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *CreateFlowpipeInputRequest) SetOverrides(v map[string]interface{})`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *CreateFlowpipeInputRequest) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.

### GetPipelineExecutionId

`func (o *CreateFlowpipeInputRequest) GetPipelineExecutionId() string`

GetPipelineExecutionId returns the PipelineExecutionId field if non-nil, zero value otherwise.

### GetPipelineExecutionIdOk

`func (o *CreateFlowpipeInputRequest) GetPipelineExecutionIdOk() (*string, bool)`

GetPipelineExecutionIdOk returns a tuple with the PipelineExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineExecutionId

`func (o *CreateFlowpipeInputRequest) SetPipelineExecutionId(v string)`

SetPipelineExecutionId sets PipelineExecutionId field to given value.

### HasPipelineExecutionId

`func (o *CreateFlowpipeInputRequest) HasPipelineExecutionId() bool`

HasPipelineExecutionId returns a boolean if a field has been set.

### GetStepExecutionId

`func (o *CreateFlowpipeInputRequest) GetStepExecutionId() string`

GetStepExecutionId returns the StepExecutionId field if non-nil, zero value otherwise.

### GetStepExecutionIdOk

`func (o *CreateFlowpipeInputRequest) GetStepExecutionIdOk() (*string, bool)`

GetStepExecutionIdOk returns a tuple with the StepExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepExecutionId

`func (o *CreateFlowpipeInputRequest) SetStepExecutionId(v string)`

SetStepExecutionId sets StepExecutionId field to given value.


### GetStepType

`func (o *CreateFlowpipeInputRequest) GetStepType() FlowpipeInputStepType`

GetStepType returns the StepType field if non-nil, zero value otherwise.

### GetStepTypeOk

`func (o *CreateFlowpipeInputRequest) GetStepTypeOk() (*FlowpipeInputStepType, bool)`

GetStepTypeOk returns a tuple with the StepType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepType

`func (o *CreateFlowpipeInputRequest) SetStepType(v FlowpipeInputStepType)`

SetStepType sets StepType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


