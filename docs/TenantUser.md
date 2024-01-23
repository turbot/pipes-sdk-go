# TenantUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**Email** | **string** | The email of the user. | 
**Id** | **string** | The unique identifier of the tenant member. | 
**Role** | **string** | The role of the tenant user. | 
**Status** | **string** | The status of the tenant member i.e invited or accepted. | 
**Tenant** | Pointer to [**Tenant**](Tenant.md) |  | [optional] 
**TenantId** | **string** | The identifier of the tenant. | 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**User** | Pointer to [**User**](User.md) |  | [optional] 
**UserId** | **string** | The identifier of the user that belongs to the tenant. | 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 

## Methods

### NewTenantUser

`func NewTenantUser(createdAt string, createdById string, email string, id string, role string, status string, tenantId string, updatedById string, userId string, versionId int32, ) *TenantUser`

NewTenantUser instantiates a new TenantUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantUserWithDefaults

`func NewTenantUserWithDefaults() *TenantUser`

NewTenantUserWithDefaults instantiates a new TenantUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *TenantUser) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TenantUser) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TenantUser) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *TenantUser) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TenantUser) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TenantUser) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TenantUser) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *TenantUser) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *TenantUser) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *TenantUser) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetEmail

`func (o *TenantUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TenantUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TenantUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetId

`func (o *TenantUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantUser) SetId(v string)`

SetId sets Id field to given value.


### GetRole

`func (o *TenantUser) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *TenantUser) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *TenantUser) SetRole(v string)`

SetRole sets Role field to given value.


### GetStatus

`func (o *TenantUser) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TenantUser) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TenantUser) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTenant

`func (o *TenantUser) GetTenant() Tenant`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *TenantUser) GetTenantOk() (*Tenant, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *TenantUser) SetTenant(v Tenant)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *TenantUser) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetTenantId

`func (o *TenantUser) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantUser) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantUser) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUpdatedAt

`func (o *TenantUser) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TenantUser) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TenantUser) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TenantUser) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *TenantUser) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *TenantUser) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *TenantUser) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *TenantUser) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *TenantUser) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *TenantUser) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *TenantUser) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetUser

`func (o *TenantUser) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TenantUser) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TenantUser) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *TenantUser) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserId

`func (o *TenantUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TenantUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TenantUser) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetVersionId

`func (o *TenantUser) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *TenantUser) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *TenantUser) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


