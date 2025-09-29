# CreateTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Expiration** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateTokenRequest

`func NewCreateTokenRequest() *CreateTokenRequest`

NewCreateTokenRequest instantiates a new CreateTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTokenRequestWithDefaults

`func NewCreateTokenRequestWithDefaults() *CreateTokenRequest`

NewCreateTokenRequestWithDefaults instantiates a new CreateTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateTokenRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTokenRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTokenRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTokenRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiration

`func (o *CreateTokenRequest) GetExpiration() int32`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *CreateTokenRequest) GetExpirationOk() (*int32, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *CreateTokenRequest) SetExpiration(v int32)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *CreateTokenRequest) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetTitle

`func (o *CreateTokenRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateTokenRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateTokenRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateTokenRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


