# CreateDatatankTableRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DesiredState** | Pointer to [**DesiredState**](DesiredState.md) |  | [optional] 
**Frequency** | Pointer to [**PipelineFrequency**](PipelineFrequency.md) |  | [optional] 
**Name** | **string** |  | 
**PartPer** | Pointer to **string** |  | [optional] 
**SourceQuery** | Pointer to **string** |  | [optional] 
**SourceSchema** | Pointer to **string** |  | [optional] 
**SourceTable** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewCreateDatatankTableRequest

`func NewCreateDatatankTableRequest(name string, type_ string, ) *CreateDatatankTableRequest`

NewCreateDatatankTableRequest instantiates a new CreateDatatankTableRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDatatankTableRequestWithDefaults

`func NewCreateDatatankTableRequestWithDefaults() *CreateDatatankTableRequest`

NewCreateDatatankTableRequestWithDefaults instantiates a new CreateDatatankTableRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateDatatankTableRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDatatankTableRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDatatankTableRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateDatatankTableRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDesiredState

`func (o *CreateDatatankTableRequest) GetDesiredState() DesiredState`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *CreateDatatankTableRequest) GetDesiredStateOk() (*DesiredState, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *CreateDatatankTableRequest) SetDesiredState(v DesiredState)`

SetDesiredState sets DesiredState field to given value.

### HasDesiredState

`func (o *CreateDatatankTableRequest) HasDesiredState() bool`

HasDesiredState returns a boolean if a field has been set.

### GetFrequency

`func (o *CreateDatatankTableRequest) GetFrequency() PipelineFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *CreateDatatankTableRequest) GetFrequencyOk() (*PipelineFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *CreateDatatankTableRequest) SetFrequency(v PipelineFrequency)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *CreateDatatankTableRequest) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetName

`func (o *CreateDatatankTableRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDatatankTableRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDatatankTableRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPartPer

`func (o *CreateDatatankTableRequest) GetPartPer() string`

GetPartPer returns the PartPer field if non-nil, zero value otherwise.

### GetPartPerOk

`func (o *CreateDatatankTableRequest) GetPartPerOk() (*string, bool)`

GetPartPerOk returns a tuple with the PartPer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartPer

`func (o *CreateDatatankTableRequest) SetPartPer(v string)`

SetPartPer sets PartPer field to given value.

### HasPartPer

`func (o *CreateDatatankTableRequest) HasPartPer() bool`

HasPartPer returns a boolean if a field has been set.

### GetSourceQuery

`func (o *CreateDatatankTableRequest) GetSourceQuery() string`

GetSourceQuery returns the SourceQuery field if non-nil, zero value otherwise.

### GetSourceQueryOk

`func (o *CreateDatatankTableRequest) GetSourceQueryOk() (*string, bool)`

GetSourceQueryOk returns a tuple with the SourceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceQuery

`func (o *CreateDatatankTableRequest) SetSourceQuery(v string)`

SetSourceQuery sets SourceQuery field to given value.

### HasSourceQuery

`func (o *CreateDatatankTableRequest) HasSourceQuery() bool`

HasSourceQuery returns a boolean if a field has been set.

### GetSourceSchema

`func (o *CreateDatatankTableRequest) GetSourceSchema() string`

GetSourceSchema returns the SourceSchema field if non-nil, zero value otherwise.

### GetSourceSchemaOk

`func (o *CreateDatatankTableRequest) GetSourceSchemaOk() (*string, bool)`

GetSourceSchemaOk returns a tuple with the SourceSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSchema

`func (o *CreateDatatankTableRequest) SetSourceSchema(v string)`

SetSourceSchema sets SourceSchema field to given value.

### HasSourceSchema

`func (o *CreateDatatankTableRequest) HasSourceSchema() bool`

HasSourceSchema returns a boolean if a field has been set.

### GetSourceTable

`func (o *CreateDatatankTableRequest) GetSourceTable() string`

GetSourceTable returns the SourceTable field if non-nil, zero value otherwise.

### GetSourceTableOk

`func (o *CreateDatatankTableRequest) GetSourceTableOk() (*string, bool)`

GetSourceTableOk returns a tuple with the SourceTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTable

`func (o *CreateDatatankTableRequest) SetSourceTable(v string)`

SetSourceTable sets SourceTable field to given value.

### HasSourceTable

`func (o *CreateDatatankTableRequest) HasSourceTable() bool`

HasSourceTable returns a boolean if a field has been set.

### GetType

`func (o *CreateDatatankTableRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateDatatankTableRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateDatatankTableRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


