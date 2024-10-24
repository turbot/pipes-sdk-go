# CreateDatatankRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**DesiredState** | Pointer to [**DesiredState**](DesiredState.md) |  | [optional] 
**Handle** | **string** |  | 

## Methods

### NewCreateDatatankRequest

`func NewCreateDatatankRequest(handle string, ) *CreateDatatankRequest`

NewCreateDatatankRequest instantiates a new CreateDatatankRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDatatankRequestWithDefaults

`func NewCreateDatatankRequestWithDefaults() *CreateDatatankRequest`

NewCreateDatatankRequestWithDefaults instantiates a new CreateDatatankRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateDatatankRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDatatankRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDatatankRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateDatatankRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDesiredState

`func (o *CreateDatatankRequest) GetDesiredState() DesiredState`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *CreateDatatankRequest) GetDesiredStateOk() (*DesiredState, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *CreateDatatankRequest) SetDesiredState(v DesiredState)`

SetDesiredState sets DesiredState field to given value.

### HasDesiredState

`func (o *CreateDatatankRequest) HasDesiredState() bool`

HasDesiredState returns a boolean if a field has been set.

### GetHandle

`func (o *CreateDatatankRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *CreateDatatankRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *CreateDatatankRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


