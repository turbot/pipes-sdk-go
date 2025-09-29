# UpdateServiceAccountUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the service account user. | [optional] 
**Title** | Pointer to **string** | The title of the service account user. | [optional] 
**TokenMinIssuedAt** | Pointer to **string** | The time which user and temporary auth tokens must be issued after. | [optional] 

## Methods

### NewUpdateServiceAccountUserRequest

`func NewUpdateServiceAccountUserRequest() *UpdateServiceAccountUserRequest`

NewUpdateServiceAccountUserRequest instantiates a new UpdateServiceAccountUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServiceAccountUserRequestWithDefaults

`func NewUpdateServiceAccountUserRequestWithDefaults() *UpdateServiceAccountUserRequest`

NewUpdateServiceAccountUserRequestWithDefaults instantiates a new UpdateServiceAccountUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateServiceAccountUserRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateServiceAccountUserRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateServiceAccountUserRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateServiceAccountUserRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTitle

`func (o *UpdateServiceAccountUserRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateServiceAccountUserRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateServiceAccountUserRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateServiceAccountUserRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTokenMinIssuedAt

`func (o *UpdateServiceAccountUserRequest) GetTokenMinIssuedAt() string`

GetTokenMinIssuedAt returns the TokenMinIssuedAt field if non-nil, zero value otherwise.

### GetTokenMinIssuedAtOk

`func (o *UpdateServiceAccountUserRequest) GetTokenMinIssuedAtOk() (*string, bool)`

GetTokenMinIssuedAtOk returns a tuple with the TokenMinIssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMinIssuedAt

`func (o *UpdateServiceAccountUserRequest) SetTokenMinIssuedAt(v string)`

SetTokenMinIssuedAt sets TokenMinIssuedAt field to given value.

### HasTokenMinIssuedAt

`func (o *UpdateServiceAccountUserRequest) HasTokenMinIssuedAt() bool`

HasTokenMinIssuedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


