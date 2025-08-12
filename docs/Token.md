# Token

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) | User information for the user who created this. | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**Description** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**Last4** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**TokenType** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) | User information for the last user to update this. | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**UserId** | **string** |  | 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 

## Methods

### NewToken

`func NewToken(createdAt string, createdById string, id string, status string, updatedById string, userId string, versionId int32, ) *Token`

NewToken instantiates a new Token object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWithDefaults

`func NewTokenWithDefaults() *Token`

NewTokenWithDefaults instantiates a new Token object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Token) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Token) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Token) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Token) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Token) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Token) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Token) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *Token) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *Token) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *Token) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetDescription

`func (o *Token) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Token) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Token) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Token) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiresAt

`func (o *Token) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Token) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Token) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *Token) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *Token) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Token) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Token) SetId(v string)`

SetId sets Id field to given value.


### GetLast4

`func (o *Token) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *Token) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *Token) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *Token) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### GetStatus

`func (o *Token) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Token) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Token) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *Token) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Token) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Token) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Token) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTokenType

`func (o *Token) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *Token) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *Token) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *Token) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Token) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Token) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Token) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Token) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Token) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Token) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Token) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Token) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *Token) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *Token) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *Token) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetUserId

`func (o *Token) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Token) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Token) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetVersionId

`func (o *Token) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *Token) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *Token) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


