# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarUrl** | Pointer to **string** |  | [optional] 
**Billing** | Pointer to [**BillingInfo**](BillingInfo.md) |  | [optional] 
**CreatedAt** | **string** | The user created time. | 
**DisplayName** | Pointer to **string** |  | [optional] 
**Handle** | **string** | The handle name of a user. | 
**Id** | **string** | The unique identifier of a user. | 
**PreviewAccessMode** | Pointer to **string** |  | [optional] 
**Status** | **string** | The status of the user i.e accepted or pending. | 
**TokenMinIssuedAt** | Pointer to **string** | The time which user and temporary auth tokens must be issued after. | [optional] 
**UpdatedAt** | Pointer to **string** | The user updated time. | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**UsageComputeAction** | Pointer to **string** |  | [optional] 
**UsageComputeThreshold** | Pointer to **int64** |  | [optional] 
**UsageStorageAction** | Pointer to **string** |  | [optional] 
**UsageStorageThreshold** | Pointer to **int64** |  | [optional] 
**VersionId** | **int32** | The current version of a user. | 

## Methods

### NewUser

`func NewUser(createdAt string, handle string, id string, status string, versionId int32, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarUrl

`func (o *User) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *User) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *User) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *User) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetBilling

`func (o *User) GetBilling() BillingInfo`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *User) GetBillingOk() (*BillingInfo, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *User) SetBilling(v BillingInfo)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *User) HasBilling() bool`

HasBilling returns a boolean if a field has been set.

### GetCreatedAt

`func (o *User) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDisplayName

`func (o *User) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *User) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *User) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *User) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHandle

`func (o *User) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *User) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *User) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.


### GetPreviewAccessMode

`func (o *User) GetPreviewAccessMode() string`

GetPreviewAccessMode returns the PreviewAccessMode field if non-nil, zero value otherwise.

### GetPreviewAccessModeOk

`func (o *User) GetPreviewAccessModeOk() (*string, bool)`

GetPreviewAccessModeOk returns a tuple with the PreviewAccessMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewAccessMode

`func (o *User) SetPreviewAccessMode(v string)`

SetPreviewAccessMode sets PreviewAccessMode field to given value.

### HasPreviewAccessMode

`func (o *User) HasPreviewAccessMode() bool`

HasPreviewAccessMode returns a boolean if a field has been set.

### GetStatus

`func (o *User) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *User) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *User) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTokenMinIssuedAt

`func (o *User) GetTokenMinIssuedAt() string`

GetTokenMinIssuedAt returns the TokenMinIssuedAt field if non-nil, zero value otherwise.

### GetTokenMinIssuedAtOk

`func (o *User) GetTokenMinIssuedAtOk() (*string, bool)`

GetTokenMinIssuedAtOk returns a tuple with the TokenMinIssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMinIssuedAt

`func (o *User) SetTokenMinIssuedAt(v string)`

SetTokenMinIssuedAt sets TokenMinIssuedAt field to given value.

### HasTokenMinIssuedAt

`func (o *User) HasTokenMinIssuedAt() bool`

HasTokenMinIssuedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *User) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *User) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *User) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *User) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUrl

`func (o *User) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *User) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *User) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *User) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsageComputeAction

`func (o *User) GetUsageComputeAction() string`

GetUsageComputeAction returns the UsageComputeAction field if non-nil, zero value otherwise.

### GetUsageComputeActionOk

`func (o *User) GetUsageComputeActionOk() (*string, bool)`

GetUsageComputeActionOk returns a tuple with the UsageComputeAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageComputeAction

`func (o *User) SetUsageComputeAction(v string)`

SetUsageComputeAction sets UsageComputeAction field to given value.

### HasUsageComputeAction

`func (o *User) HasUsageComputeAction() bool`

HasUsageComputeAction returns a boolean if a field has been set.

### GetUsageComputeThreshold

`func (o *User) GetUsageComputeThreshold() int64`

GetUsageComputeThreshold returns the UsageComputeThreshold field if non-nil, zero value otherwise.

### GetUsageComputeThresholdOk

`func (o *User) GetUsageComputeThresholdOk() (*int64, bool)`

GetUsageComputeThresholdOk returns a tuple with the UsageComputeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageComputeThreshold

`func (o *User) SetUsageComputeThreshold(v int64)`

SetUsageComputeThreshold sets UsageComputeThreshold field to given value.

### HasUsageComputeThreshold

`func (o *User) HasUsageComputeThreshold() bool`

HasUsageComputeThreshold returns a boolean if a field has been set.

### GetUsageStorageAction

`func (o *User) GetUsageStorageAction() string`

GetUsageStorageAction returns the UsageStorageAction field if non-nil, zero value otherwise.

### GetUsageStorageActionOk

`func (o *User) GetUsageStorageActionOk() (*string, bool)`

GetUsageStorageActionOk returns a tuple with the UsageStorageAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageStorageAction

`func (o *User) SetUsageStorageAction(v string)`

SetUsageStorageAction sets UsageStorageAction field to given value.

### HasUsageStorageAction

`func (o *User) HasUsageStorageAction() bool`

HasUsageStorageAction returns a boolean if a field has been set.

### GetUsageStorageThreshold

`func (o *User) GetUsageStorageThreshold() int64`

GetUsageStorageThreshold returns the UsageStorageThreshold field if non-nil, zero value otherwise.

### GetUsageStorageThresholdOk

`func (o *User) GetUsageStorageThresholdOk() (*int64, bool)`

GetUsageStorageThresholdOk returns a tuple with the UsageStorageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageStorageThreshold

`func (o *User) SetUsageStorageThreshold(v int64)`

SetUsageStorageThreshold sets UsageStorageThreshold field to given value.

### HasUsageStorageThreshold

`func (o *User) HasUsageStorageThreshold() bool`

HasUsageStorageThreshold returns a boolean if a field has been set.

### GetVersionId

`func (o *User) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *User) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *User) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


