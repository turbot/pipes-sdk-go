# CreateTriggerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | **map[string]interface{}** | A map of arguments to be passed to be trigger. | 
**Description** | Pointer to **string** | Optional description for the trigger. | [optional] 
**Name** | Pointer to **string** | The name of the trigger, has to be unique in a workspace. | [optional] 
**Pipeline** | **string** | The name / identifier of the pipeline which needs to be triggered. | 
**Schedule** | [**PipelineFrequency**](PipelineFrequency.md) | The schedule at which the trigger will run. | 
**State** | Pointer to [**TriggerState**](TriggerState.md) | The state of the trigger. | [optional] 
**Title** | Pointer to **string** | Optional title for the trigger. | [optional] 

## Methods

### NewCreateTriggerRequest

`func NewCreateTriggerRequest(args map[string]interface{}, pipeline string, schedule PipelineFrequency, ) *CreateTriggerRequest`

NewCreateTriggerRequest instantiates a new CreateTriggerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTriggerRequestWithDefaults

`func NewCreateTriggerRequestWithDefaults() *CreateTriggerRequest`

NewCreateTriggerRequestWithDefaults instantiates a new CreateTriggerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *CreateTriggerRequest) GetArgs() map[string]interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *CreateTriggerRequest) GetArgsOk() (*map[string]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *CreateTriggerRequest) SetArgs(v map[string]interface{})`

SetArgs sets Args field to given value.


### GetDescription

`func (o *CreateTriggerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTriggerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTriggerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTriggerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CreateTriggerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTriggerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTriggerRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateTriggerRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPipeline

`func (o *CreateTriggerRequest) GetPipeline() string`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *CreateTriggerRequest) GetPipelineOk() (*string, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *CreateTriggerRequest) SetPipeline(v string)`

SetPipeline sets Pipeline field to given value.


### GetSchedule

`func (o *CreateTriggerRequest) GetSchedule() PipelineFrequency`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CreateTriggerRequest) GetScheduleOk() (*PipelineFrequency, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CreateTriggerRequest) SetSchedule(v PipelineFrequency)`

SetSchedule sets Schedule field to given value.


### GetState

`func (o *CreateTriggerRequest) GetState() TriggerState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateTriggerRequest) GetStateOk() (*TriggerState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateTriggerRequest) SetState(v TriggerState)`

SetState sets State field to given value.

### HasState

`func (o *CreateTriggerRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTitle

`func (o *CreateTriggerRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateTriggerRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateTriggerRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateTriggerRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


