# AIKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) | User information for the user who created this. | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**Id** | **string** |  | 
**IdentityId** | Pointer to **string** |  | [optional] 
**Last4** | Pointer to **string** |  | [optional] 
**Provider** | [**AIProvider**](AIProvider.md) |  | 
**TenantId** | **string** |  | 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) | User information for the last user to update this. | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 

## Methods

### NewAIKey

`func NewAIKey(createdAt string, createdById string, id string, provider AIProvider, tenantId string, updatedById string, versionId int32, ) *AIKey`

NewAIKey instantiates a new AIKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIKeyWithDefaults

`func NewAIKeyWithDefaults() *AIKey`

NewAIKeyWithDefaults instantiates a new AIKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *AIKey) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AIKey) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AIKey) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *AIKey) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AIKey) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AIKey) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AIKey) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *AIKey) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *AIKey) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *AIKey) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetId

`func (o *AIKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AIKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AIKey) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityId

`func (o *AIKey) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *AIKey) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *AIKey) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *AIKey) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetLast4

`func (o *AIKey) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *AIKey) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *AIKey) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *AIKey) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### GetProvider

`func (o *AIKey) GetProvider() AIProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AIKey) GetProviderOk() (*AIProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AIKey) SetProvider(v AIProvider)`

SetProvider sets Provider field to given value.


### GetTenantId

`func (o *AIKey) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AIKey) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AIKey) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUpdatedAt

`func (o *AIKey) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AIKey) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AIKey) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AIKey) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *AIKey) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *AIKey) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *AIKey) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *AIKey) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *AIKey) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *AIKey) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *AIKey) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetVersionId

`func (o *AIKey) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *AIKey) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *AIKey) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


