# UpdateTriggerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **map[string]interface{}** | A map of arguments to be passed to be trigger. | [optional] 
**Description** | Pointer to **string** | Optional description for the trigger. | [optional] 
**Name** | Pointer to **string** | The name of the trigger. | [optional] 
**Schedule** | Pointer to [**PipelineFrequency**](PipelineFrequency.md) | The schedule at which the trigger will run. | [optional] 
**State** | Pointer to [**TriggerState**](TriggerState.md) | The desired state of the trigger. | [optional] 
**Title** | Pointer to **string** | Optional title for the trigger. | [optional] 

## Methods

### NewUpdateTriggerRequest

`func NewUpdateTriggerRequest() *UpdateTriggerRequest`

NewUpdateTriggerRequest instantiates a new UpdateTriggerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTriggerRequestWithDefaults

`func NewUpdateTriggerRequestWithDefaults() *UpdateTriggerRequest`

NewUpdateTriggerRequestWithDefaults instantiates a new UpdateTriggerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *UpdateTriggerRequest) GetArgs() map[string]interface{}`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *UpdateTriggerRequest) GetArgsOk() (*map[string]interface{}, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *UpdateTriggerRequest) SetArgs(v map[string]interface{})`

SetArgs sets Args field to given value.

### HasArgs

`func (o *UpdateTriggerRequest) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateTriggerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateTriggerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateTriggerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateTriggerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateTriggerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTriggerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTriggerRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTriggerRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSchedule

`func (o *UpdateTriggerRequest) GetSchedule() PipelineFrequency`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *UpdateTriggerRequest) GetScheduleOk() (*PipelineFrequency, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *UpdateTriggerRequest) SetSchedule(v PipelineFrequency)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *UpdateTriggerRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetState

`func (o *UpdateTriggerRequest) GetState() TriggerState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateTriggerRequest) GetStateOk() (*TriggerState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateTriggerRequest) SetState(v TriggerState)`

SetState sets State field to given value.

### HasState

`func (o *UpdateTriggerRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTitle

`func (o *UpdateTriggerRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateTriggerRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateTriggerRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateTriggerRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


