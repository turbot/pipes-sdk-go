# UpdateTenantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Handle** | Pointer to **string** |  | [optional] 
**UsageComputeAction** | Pointer to [**IdentityUsageThresholdAction**](IdentityUsageThresholdAction.md) |  | [optional] 
**UsageComputeThreshold** | Pointer to **int64** |  | [optional] 
**UsageStorageAction** | Pointer to [**IdentityUsageThresholdAction**](IdentityUsageThresholdAction.md) |  | [optional] 
**UsageStorageThreshold** | Pointer to **int64** |  | [optional] 
**UsageUserAction** | Pointer to [**IdentityUsageThresholdAction**](IdentityUsageThresholdAction.md) |  | [optional] 
**UsageUserThreshold** | Pointer to **int64** |  | [optional] 

## Methods

### NewUpdateTenantRequest

`func NewUpdateTenantRequest() *UpdateTenantRequest`

NewUpdateTenantRequest instantiates a new UpdateTenantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTenantRequestWithDefaults

`func NewUpdateTenantRequestWithDefaults() *UpdateTenantRequest`

NewUpdateTenantRequestWithDefaults instantiates a new UpdateTenantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *UpdateTenantRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateTenantRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateTenantRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateTenantRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHandle

`func (o *UpdateTenantRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *UpdateTenantRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *UpdateTenantRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *UpdateTenantRequest) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetUsageComputeAction

`func (o *UpdateTenantRequest) GetUsageComputeAction() IdentityUsageThresholdAction`

GetUsageComputeAction returns the UsageComputeAction field if non-nil, zero value otherwise.

### GetUsageComputeActionOk

`func (o *UpdateTenantRequest) GetUsageComputeActionOk() (*IdentityUsageThresholdAction, bool)`

GetUsageComputeActionOk returns a tuple with the UsageComputeAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageComputeAction

`func (o *UpdateTenantRequest) SetUsageComputeAction(v IdentityUsageThresholdAction)`

SetUsageComputeAction sets UsageComputeAction field to given value.

### HasUsageComputeAction

`func (o *UpdateTenantRequest) HasUsageComputeAction() bool`

HasUsageComputeAction returns a boolean if a field has been set.

### GetUsageComputeThreshold

`func (o *UpdateTenantRequest) GetUsageComputeThreshold() int64`

GetUsageComputeThreshold returns the UsageComputeThreshold field if non-nil, zero value otherwise.

### GetUsageComputeThresholdOk

`func (o *UpdateTenantRequest) GetUsageComputeThresholdOk() (*int64, bool)`

GetUsageComputeThresholdOk returns a tuple with the UsageComputeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageComputeThreshold

`func (o *UpdateTenantRequest) SetUsageComputeThreshold(v int64)`

SetUsageComputeThreshold sets UsageComputeThreshold field to given value.

### HasUsageComputeThreshold

`func (o *UpdateTenantRequest) HasUsageComputeThreshold() bool`

HasUsageComputeThreshold returns a boolean if a field has been set.

### GetUsageStorageAction

`func (o *UpdateTenantRequest) GetUsageStorageAction() IdentityUsageThresholdAction`

GetUsageStorageAction returns the UsageStorageAction field if non-nil, zero value otherwise.

### GetUsageStorageActionOk

`func (o *UpdateTenantRequest) GetUsageStorageActionOk() (*IdentityUsageThresholdAction, bool)`

GetUsageStorageActionOk returns a tuple with the UsageStorageAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageStorageAction

`func (o *UpdateTenantRequest) SetUsageStorageAction(v IdentityUsageThresholdAction)`

SetUsageStorageAction sets UsageStorageAction field to given value.

### HasUsageStorageAction

`func (o *UpdateTenantRequest) HasUsageStorageAction() bool`

HasUsageStorageAction returns a boolean if a field has been set.

### GetUsageStorageThreshold

`func (o *UpdateTenantRequest) GetUsageStorageThreshold() int64`

GetUsageStorageThreshold returns the UsageStorageThreshold field if non-nil, zero value otherwise.

### GetUsageStorageThresholdOk

`func (o *UpdateTenantRequest) GetUsageStorageThresholdOk() (*int64, bool)`

GetUsageStorageThresholdOk returns a tuple with the UsageStorageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageStorageThreshold

`func (o *UpdateTenantRequest) SetUsageStorageThreshold(v int64)`

SetUsageStorageThreshold sets UsageStorageThreshold field to given value.

### HasUsageStorageThreshold

`func (o *UpdateTenantRequest) HasUsageStorageThreshold() bool`

HasUsageStorageThreshold returns a boolean if a field has been set.

### GetUsageUserAction

`func (o *UpdateTenantRequest) GetUsageUserAction() IdentityUsageThresholdAction`

GetUsageUserAction returns the UsageUserAction field if non-nil, zero value otherwise.

### GetUsageUserActionOk

`func (o *UpdateTenantRequest) GetUsageUserActionOk() (*IdentityUsageThresholdAction, bool)`

GetUsageUserActionOk returns a tuple with the UsageUserAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageUserAction

`func (o *UpdateTenantRequest) SetUsageUserAction(v IdentityUsageThresholdAction)`

SetUsageUserAction sets UsageUserAction field to given value.

### HasUsageUserAction

`func (o *UpdateTenantRequest) HasUsageUserAction() bool`

HasUsageUserAction returns a boolean if a field has been set.

### GetUsageUserThreshold

`func (o *UpdateTenantRequest) GetUsageUserThreshold() int64`

GetUsageUserThreshold returns the UsageUserThreshold field if non-nil, zero value otherwise.

### GetUsageUserThresholdOk

`func (o *UpdateTenantRequest) GetUsageUserThresholdOk() (*int64, bool)`

GetUsageUserThresholdOk returns a tuple with the UsageUserThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageUserThreshold

`func (o *UpdateTenantRequest) SetUsageUserThreshold(v int64)`

SetUsageUserThreshold sets UsageUserThreshold field to given value.

### HasUsageUserThreshold

`func (o *UpdateTenantRequest) HasUsageUserThreshold() bool`

HasUsageUserThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


