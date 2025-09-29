# CreateServiceAccountUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the service account user. | [optional] 
**Title** | Pointer to **string** | The title of the service account user. | [optional] 

## Methods

### NewCreateServiceAccountUserRequest

`func NewCreateServiceAccountUserRequest() *CreateServiceAccountUserRequest`

NewCreateServiceAccountUserRequest instantiates a new CreateServiceAccountUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceAccountUserRequestWithDefaults

`func NewCreateServiceAccountUserRequestWithDefaults() *CreateServiceAccountUserRequest`

NewCreateServiceAccountUserRequestWithDefaults instantiates a new CreateServiceAccountUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateServiceAccountUserRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateServiceAccountUserRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateServiceAccountUserRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateServiceAccountUserRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTitle

`func (o *CreateServiceAccountUserRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateServiceAccountUserRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateServiceAccountUserRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateServiceAccountUserRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


