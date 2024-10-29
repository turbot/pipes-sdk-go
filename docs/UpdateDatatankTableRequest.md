# UpdateDatatankTableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DesiredState** | Pointer to [**DesiredState**](DesiredState.md) |  | [optional] 
**Frequency** | Pointer to [**PipelineFrequency**](PipelineFrequency.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PartPer** | Pointer to **string** |  | [optional] 
**SourceQuery** | Pointer to **string** |  | [optional] 
**SourceSchema** | Pointer to **string** |  | [optional] 
**SourceTable** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateDatatankTableRequest

`func NewUpdateDatatankTableRequest() *UpdateDatatankTableRequest`

NewUpdateDatatankTableRequest instantiates a new UpdateDatatankTableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDatatankTableRequestWithDefaults

`func NewUpdateDatatankTableRequestWithDefaults() *UpdateDatatankTableRequest`

NewUpdateDatatankTableRequestWithDefaults instantiates a new UpdateDatatankTableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateDatatankTableRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateDatatankTableRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateDatatankTableRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateDatatankTableRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDesiredState

`func (o *UpdateDatatankTableRequest) GetDesiredState() DesiredState`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *UpdateDatatankTableRequest) GetDesiredStateOk() (*DesiredState, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *UpdateDatatankTableRequest) SetDesiredState(v DesiredState)`

SetDesiredState sets DesiredState field to given value.

### HasDesiredState

`func (o *UpdateDatatankTableRequest) HasDesiredState() bool`

HasDesiredState returns a boolean if a field has been set.

### GetFrequency

`func (o *UpdateDatatankTableRequest) GetFrequency() PipelineFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *UpdateDatatankTableRequest) GetFrequencyOk() (*PipelineFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *UpdateDatatankTableRequest) SetFrequency(v PipelineFrequency)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *UpdateDatatankTableRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetName

`func (o *UpdateDatatankTableRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDatatankTableRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDatatankTableRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDatatankTableRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPartPer

`func (o *UpdateDatatankTableRequest) GetPartPer() string`

GetPartPer returns the PartPer field if non-nil, zero value otherwise.

### GetPartPerOk

`func (o *UpdateDatatankTableRequest) GetPartPerOk() (*string, bool)`

GetPartPerOk returns a tuple with the PartPer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartPer

`func (o *UpdateDatatankTableRequest) SetPartPer(v string)`

SetPartPer sets PartPer field to given value.

### HasPartPer

`func (o *UpdateDatatankTableRequest) HasPartPer() bool`

HasPartPer returns a boolean if a field has been set.

### GetSourceQuery

`func (o *UpdateDatatankTableRequest) GetSourceQuery() string`

GetSourceQuery returns the SourceQuery field if non-nil, zero value otherwise.

### GetSourceQueryOk

`func (o *UpdateDatatankTableRequest) GetSourceQueryOk() (*string, bool)`

GetSourceQueryOk returns a tuple with the SourceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceQuery

`func (o *UpdateDatatankTableRequest) SetSourceQuery(v string)`

SetSourceQuery sets SourceQuery field to given value.

### HasSourceQuery

`func (o *UpdateDatatankTableRequest) HasSourceQuery() bool`

HasSourceQuery returns a boolean if a field has been set.

### GetSourceSchema

`func (o *UpdateDatatankTableRequest) GetSourceSchema() string`

GetSourceSchema returns the SourceSchema field if non-nil, zero value otherwise.

### GetSourceSchemaOk

`func (o *UpdateDatatankTableRequest) GetSourceSchemaOk() (*string, bool)`

GetSourceSchemaOk returns a tuple with the SourceSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSchema

`func (o *UpdateDatatankTableRequest) SetSourceSchema(v string)`

SetSourceSchema sets SourceSchema field to given value.

### HasSourceSchema

`func (o *UpdateDatatankTableRequest) HasSourceSchema() bool`

HasSourceSchema returns a boolean if a field has been set.

### GetSourceTable

`func (o *UpdateDatatankTableRequest) GetSourceTable() string`

GetSourceTable returns the SourceTable field if non-nil, zero value otherwise.

### GetSourceTableOk

`func (o *UpdateDatatankTableRequest) GetSourceTableOk() (*string, bool)`

GetSourceTableOk returns a tuple with the SourceTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTable

`func (o *UpdateDatatankTableRequest) SetSourceTable(v string)`

SetSourceTable sets SourceTable field to given value.

### HasSourceTable

`func (o *UpdateDatatankTableRequest) HasSourceTable() bool`

HasSourceTable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


