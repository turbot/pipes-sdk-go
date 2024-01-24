# Tenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarUrl** | Pointer to **string** |  | [optional] 
**Billing** | Pointer to [**BillingInfo**](BillingInfo.md) |  | [optional] 
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**DeletedAt** | Pointer to **string** | The time the item was deleted in ISO 8601 UTC. | [optional] 
**DeletedBy** | Pointer to [**User**](User.md) |  | [optional] 
**DeletedById** | **string** | The ID of the user that performed the deletion. | 
**DisplayName** | Pointer to **string** | The display name of the tenant. | [optional] 
**Handle** | **string** | The handle name of the tenant. | 
**Id** | **string** | The unique identifier of the tenant. | 
**State** | Pointer to **string** | The state of the tenant. | [optional] 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**UsageComputeAction** | Pointer to **string** |  | [optional] 
**UsageComputeThreshold** | Pointer to **int64** |  | [optional] 
**UsageStorageAction** | Pointer to **string** |  | [optional] 
**UsageStorageThreshold** | Pointer to **int64** |  | [optional] 
**UsageUserAction** | Pointer to **string** |  | [optional] 
**UsageUserThreshold** | Pointer to **int32** |  | [optional] 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 

## Methods

### NewTenant

`func NewTenant(createdAt string, createdById string, deletedById string, handle string, id string, updatedById string, versionId int32, ) *Tenant`

NewTenant instantiates a new Tenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantWithDefaults

`func NewTenantWithDefaults() *Tenant`

NewTenantWithDefaults instantiates a new Tenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarUrl

`func (o *Tenant) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *Tenant) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *Tenant) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *Tenant) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetBilling

`func (o *Tenant) GetBilling() BillingInfo`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *Tenant) GetBillingOk() (*BillingInfo, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *Tenant) SetBilling(v BillingInfo)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *Tenant) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Tenant) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Tenant) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Tenant) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Tenant) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Tenant) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Tenant) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Tenant) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *Tenant) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *Tenant) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *Tenant) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetDeletedAt

`func (o *Tenant) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Tenant) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Tenant) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Tenant) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *Tenant) GetDeletedBy() User`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *Tenant) GetDeletedByOk() (*User, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *Tenant) SetDeletedBy(v User)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *Tenant) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedById

`func (o *Tenant) GetDeletedById() string`

GetDeletedById returns the DeletedById field if non-nil, zero value otherwise.

### GetDeletedByIdOk

`func (o *Tenant) GetDeletedByIdOk() (*string, bool)`

GetDeletedByIdOk returns a tuple with the DeletedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedById

`func (o *Tenant) SetDeletedById(v string)`

SetDeletedById sets DeletedById field to given value.


### GetDisplayName

`func (o *Tenant) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Tenant) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Tenant) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Tenant) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHandle

`func (o *Tenant) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *Tenant) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *Tenant) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetId

`func (o *Tenant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tenant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tenant) SetId(v string)`

SetId sets Id field to given value.


### GetState

`func (o *Tenant) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Tenant) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Tenant) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Tenant) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Tenant) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Tenant) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Tenant) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Tenant) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Tenant) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Tenant) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Tenant) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Tenant) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *Tenant) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *Tenant) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *Tenant) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetUsageComputeAction

`func (o *Tenant) GetUsageComputeAction() string`

GetUsageComputeAction returns the UsageComputeAction field if non-nil, zero value otherwise.

### GetUsageComputeActionOk

`func (o *Tenant) GetUsageComputeActionOk() (*string, bool)`

GetUsageComputeActionOk returns a tuple with the UsageComputeAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageComputeAction

`func (o *Tenant) SetUsageComputeAction(v string)`

SetUsageComputeAction sets UsageComputeAction field to given value.

### HasUsageComputeAction

`func (o *Tenant) HasUsageComputeAction() bool`

HasUsageComputeAction returns a boolean if a field has been set.

### GetUsageComputeThreshold

`func (o *Tenant) GetUsageComputeThreshold() int64`

GetUsageComputeThreshold returns the UsageComputeThreshold field if non-nil, zero value otherwise.

### GetUsageComputeThresholdOk

`func (o *Tenant) GetUsageComputeThresholdOk() (*int64, bool)`

GetUsageComputeThresholdOk returns a tuple with the UsageComputeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageComputeThreshold

`func (o *Tenant) SetUsageComputeThreshold(v int64)`

SetUsageComputeThreshold sets UsageComputeThreshold field to given value.

### HasUsageComputeThreshold

`func (o *Tenant) HasUsageComputeThreshold() bool`

HasUsageComputeThreshold returns a boolean if a field has been set.

### GetUsageStorageAction

`func (o *Tenant) GetUsageStorageAction() string`

GetUsageStorageAction returns the UsageStorageAction field if non-nil, zero value otherwise.

### GetUsageStorageActionOk

`func (o *Tenant) GetUsageStorageActionOk() (*string, bool)`

GetUsageStorageActionOk returns a tuple with the UsageStorageAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageStorageAction

`func (o *Tenant) SetUsageStorageAction(v string)`

SetUsageStorageAction sets UsageStorageAction field to given value.

### HasUsageStorageAction

`func (o *Tenant) HasUsageStorageAction() bool`

HasUsageStorageAction returns a boolean if a field has been set.

### GetUsageStorageThreshold

`func (o *Tenant) GetUsageStorageThreshold() int64`

GetUsageStorageThreshold returns the UsageStorageThreshold field if non-nil, zero value otherwise.

### GetUsageStorageThresholdOk

`func (o *Tenant) GetUsageStorageThresholdOk() (*int64, bool)`

GetUsageStorageThresholdOk returns a tuple with the UsageStorageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageStorageThreshold

`func (o *Tenant) SetUsageStorageThreshold(v int64)`

SetUsageStorageThreshold sets UsageStorageThreshold field to given value.

### HasUsageStorageThreshold

`func (o *Tenant) HasUsageStorageThreshold() bool`

HasUsageStorageThreshold returns a boolean if a field has been set.

### GetUsageUserAction

`func (o *Tenant) GetUsageUserAction() string`

GetUsageUserAction returns the UsageUserAction field if non-nil, zero value otherwise.

### GetUsageUserActionOk

`func (o *Tenant) GetUsageUserActionOk() (*string, bool)`

GetUsageUserActionOk returns a tuple with the UsageUserAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageUserAction

`func (o *Tenant) SetUsageUserAction(v string)`

SetUsageUserAction sets UsageUserAction field to given value.

### HasUsageUserAction

`func (o *Tenant) HasUsageUserAction() bool`

HasUsageUserAction returns a boolean if a field has been set.

### GetUsageUserThreshold

`func (o *Tenant) GetUsageUserThreshold() int32`

GetUsageUserThreshold returns the UsageUserThreshold field if non-nil, zero value otherwise.

### GetUsageUserThresholdOk

`func (o *Tenant) GetUsageUserThresholdOk() (*int32, bool)`

GetUsageUserThresholdOk returns a tuple with the UsageUserThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageUserThreshold

`func (o *Tenant) SetUsageUserThreshold(v int32)`

SetUsageUserThreshold sets UsageUserThreshold field to given value.

### HasUsageUserThreshold

`func (o *Tenant) HasUsageUserThreshold() bool`

HasUsageUserThreshold returns a boolean if a field has been set.

### GetVersionId

`func (o *Tenant) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *Tenant) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *Tenant) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


