# Org

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarUrl** | Pointer to **string** | The avatar of an org. | [optional] 
**Billing** | Pointer to [**BillingInfo**](BillingInfo.md) | Struct to store billing information | [optional] 
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) | User information for the user who created this. | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**DeletedAt** | Pointer to **string** | The time the item was deleted in ISO 8601 UTC. | [optional] 
**DeletedBy** | Pointer to [**User**](User.md) | User information for the user that performed the deletion. | [optional] 
**DeletedById** | **string** | The ID of the user that performed the deletion. | 
**DisplayName** | Pointer to **string** | The display name of an org. | [optional] 
**Handle** | **string** | The handle name for an org. | 
**Id** | **string** | The unique identifier for an org. | 
**TenantId** | **string** | The unique identifier of the tenant the org belongs to. | 
**TokenMinIssuedAt** | Pointer to **string** | The time which user and temporary auth tokens must be issued after. | [optional] 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) | User information for the last user to update this. | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**Url** | Pointer to **string** |  | [optional] 
**UsageComputeAction** | Pointer to [**IdentityUsageThresholdAction**](IdentityUsageThresholdAction.md) |  | [optional] 
**UsageComputeThreshold** | Pointer to **int64** |  | [optional] 
**UsageStorageAction** | Pointer to [**IdentityUsageThresholdAction**](IdentityUsageThresholdAction.md) |  | [optional] 
**UsageStorageThreshold** | Pointer to **int64** |  | [optional] 
**UsageUserAction** | Pointer to [**IdentityUsageThresholdAction**](IdentityUsageThresholdAction.md) |  | [optional] 
**UsageUserThreshold** | Pointer to **int64** |  | [optional] 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 

## Methods

### NewOrg

`func NewOrg(createdAt string, createdById string, deletedById string, handle string, id string, tenantId string, updatedById string, versionId int32, ) *Org`

NewOrg instantiates a new Org object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgWithDefaults

`func NewOrgWithDefaults() *Org`

NewOrgWithDefaults instantiates a new Org object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarUrl

`func (o *Org) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *Org) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *Org) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *Org) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetBilling

`func (o *Org) GetBilling() BillingInfo`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *Org) GetBillingOk() (*BillingInfo, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *Org) SetBilling(v BillingInfo)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *Org) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Org) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Org) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Org) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Org) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Org) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Org) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Org) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *Org) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *Org) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *Org) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetDeletedAt

`func (o *Org) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Org) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Org) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Org) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *Org) GetDeletedBy() User`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *Org) GetDeletedByOk() (*User, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *Org) SetDeletedBy(v User)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *Org) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedById

`func (o *Org) GetDeletedById() string`

GetDeletedById returns the DeletedById field if non-nil, zero value otherwise.

### GetDeletedByIdOk

`func (o *Org) GetDeletedByIdOk() (*string, bool)`

GetDeletedByIdOk returns a tuple with the DeletedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedById

`func (o *Org) SetDeletedById(v string)`

SetDeletedById sets DeletedById field to given value.


### GetDisplayName

`func (o *Org) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Org) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Org) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Org) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHandle

`func (o *Org) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *Org) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *Org) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetId

`func (o *Org) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Org) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Org) SetId(v string)`

SetId sets Id field to given value.


### GetTenantId

`func (o *Org) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Org) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Org) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetTokenMinIssuedAt

`func (o *Org) GetTokenMinIssuedAt() string`

GetTokenMinIssuedAt returns the TokenMinIssuedAt field if non-nil, zero value otherwise.

### GetTokenMinIssuedAtOk

`func (o *Org) GetTokenMinIssuedAtOk() (*string, bool)`

GetTokenMinIssuedAtOk returns a tuple with the TokenMinIssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMinIssuedAt

`func (o *Org) SetTokenMinIssuedAt(v string)`

SetTokenMinIssuedAt sets TokenMinIssuedAt field to given value.

### HasTokenMinIssuedAt

`func (o *Org) HasTokenMinIssuedAt() bool`

HasTokenMinIssuedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Org) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Org) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Org) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Org) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Org) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Org) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Org) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Org) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *Org) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *Org) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *Org) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetUrl

`func (o *Org) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Org) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Org) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Org) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsageComputeAction

`func (o *Org) GetUsageComputeAction() IdentityUsageThresholdAction`

GetUsageComputeAction returns the UsageComputeAction field if non-nil, zero value otherwise.

### GetUsageComputeActionOk

`func (o *Org) GetUsageComputeActionOk() (*IdentityUsageThresholdAction, bool)`

GetUsageComputeActionOk returns a tuple with the UsageComputeAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageComputeAction

`func (o *Org) SetUsageComputeAction(v IdentityUsageThresholdAction)`

SetUsageComputeAction sets UsageComputeAction field to given value.

### HasUsageComputeAction

`func (o *Org) HasUsageComputeAction() bool`

HasUsageComputeAction returns a boolean if a field has been set.

### GetUsageComputeThreshold

`func (o *Org) GetUsageComputeThreshold() int64`

GetUsageComputeThreshold returns the UsageComputeThreshold field if non-nil, zero value otherwise.

### GetUsageComputeThresholdOk

`func (o *Org) GetUsageComputeThresholdOk() (*int64, bool)`

GetUsageComputeThresholdOk returns a tuple with the UsageComputeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageComputeThreshold

`func (o *Org) SetUsageComputeThreshold(v int64)`

SetUsageComputeThreshold sets UsageComputeThreshold field to given value.

### HasUsageComputeThreshold

`func (o *Org) HasUsageComputeThreshold() bool`

HasUsageComputeThreshold returns a boolean if a field has been set.

### GetUsageStorageAction

`func (o *Org) GetUsageStorageAction() IdentityUsageThresholdAction`

GetUsageStorageAction returns the UsageStorageAction field if non-nil, zero value otherwise.

### GetUsageStorageActionOk

`func (o *Org) GetUsageStorageActionOk() (*IdentityUsageThresholdAction, bool)`

GetUsageStorageActionOk returns a tuple with the UsageStorageAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageStorageAction

`func (o *Org) SetUsageStorageAction(v IdentityUsageThresholdAction)`

SetUsageStorageAction sets UsageStorageAction field to given value.

### HasUsageStorageAction

`func (o *Org) HasUsageStorageAction() bool`

HasUsageStorageAction returns a boolean if a field has been set.

### GetUsageStorageThreshold

`func (o *Org) GetUsageStorageThreshold() int64`

GetUsageStorageThreshold returns the UsageStorageThreshold field if non-nil, zero value otherwise.

### GetUsageStorageThresholdOk

`func (o *Org) GetUsageStorageThresholdOk() (*int64, bool)`

GetUsageStorageThresholdOk returns a tuple with the UsageStorageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageStorageThreshold

`func (o *Org) SetUsageStorageThreshold(v int64)`

SetUsageStorageThreshold sets UsageStorageThreshold field to given value.

### HasUsageStorageThreshold

`func (o *Org) HasUsageStorageThreshold() bool`

HasUsageStorageThreshold returns a boolean if a field has been set.

### GetUsageUserAction

`func (o *Org) GetUsageUserAction() IdentityUsageThresholdAction`

GetUsageUserAction returns the UsageUserAction field if non-nil, zero value otherwise.

### GetUsageUserActionOk

`func (o *Org) GetUsageUserActionOk() (*IdentityUsageThresholdAction, bool)`

GetUsageUserActionOk returns a tuple with the UsageUserAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageUserAction

`func (o *Org) SetUsageUserAction(v IdentityUsageThresholdAction)`

SetUsageUserAction sets UsageUserAction field to given value.

### HasUsageUserAction

`func (o *Org) HasUsageUserAction() bool`

HasUsageUserAction returns a boolean if a field has been set.

### GetUsageUserThreshold

`func (o *Org) GetUsageUserThreshold() int64`

GetUsageUserThreshold returns the UsageUserThreshold field if non-nil, zero value otherwise.

### GetUsageUserThresholdOk

`func (o *Org) GetUsageUserThresholdOk() (*int64, bool)`

GetUsageUserThresholdOk returns a tuple with the UsageUserThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageUserThreshold

`func (o *Org) SetUsageUserThreshold(v int64)`

SetUsageUserThreshold sets UsageUserThreshold field to given value.

### HasUsageUserThreshold

`func (o *Org) HasUsageUserThreshold() bool`

HasUsageUserThreshold returns a boolean if a field has been set.

### GetVersionId

`func (o *Org) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *Org) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *Org) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


