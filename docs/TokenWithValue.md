# TokenWithValue

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
**Token** | Pointer to **string** | Previously this was named \&quot;Token\&quot; and API expects \&quot;token\&quot; as the key, but to enable embedding of Token struct renamed it to \&quot;Value\&quot;. | [optional] 
**TokenType** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) | User information for the last user to update this. | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**UserId** | **string** |  | 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 

## Methods

### NewTokenWithValue

`func NewTokenWithValue(createdAt string, createdById string, id string, status string, updatedById string, userId string, versionId int32, ) *TokenWithValue`

NewTokenWithValue instantiates a new TokenWithValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenWithValueWithDefaults

`func NewTokenWithValueWithDefaults() *TokenWithValue`

NewTokenWithValueWithDefaults instantiates a new TokenWithValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *TokenWithValue) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TokenWithValue) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TokenWithValue) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *TokenWithValue) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TokenWithValue) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TokenWithValue) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TokenWithValue) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *TokenWithValue) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *TokenWithValue) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *TokenWithValue) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetDescription

`func (o *TokenWithValue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TokenWithValue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TokenWithValue) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TokenWithValue) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiresAt

`func (o *TokenWithValue) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *TokenWithValue) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *TokenWithValue) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *TokenWithValue) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *TokenWithValue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenWithValue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenWithValue) SetId(v string)`

SetId sets Id field to given value.


### GetLast4

`func (o *TokenWithValue) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *TokenWithValue) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *TokenWithValue) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *TokenWithValue) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### GetStatus

`func (o *TokenWithValue) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TokenWithValue) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TokenWithValue) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *TokenWithValue) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TokenWithValue) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TokenWithValue) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TokenWithValue) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetToken

`func (o *TokenWithValue) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TokenWithValue) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TokenWithValue) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TokenWithValue) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenType

`func (o *TokenWithValue) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *TokenWithValue) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *TokenWithValue) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *TokenWithValue) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TokenWithValue) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TokenWithValue) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TokenWithValue) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TokenWithValue) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *TokenWithValue) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *TokenWithValue) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *TokenWithValue) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *TokenWithValue) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *TokenWithValue) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *TokenWithValue) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *TokenWithValue) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetUserId

`func (o *TokenWithValue) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TokenWithValue) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TokenWithValue) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetVersionId

`func (o *TokenWithValue) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *TokenWithValue) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *TokenWithValue) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


