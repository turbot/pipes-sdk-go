# UpdateOrgRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** |  | [optional] 
**Handle** | Pointer to **string** |  | [optional] 
**TokenMinIssuedAt** | Pointer to **string** | The time which user and temporary auth tokens must be issued after. | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**UsageComputeAction** | Pointer to **string** |  | [optional] 
**UsageComputeThreshold** | Pointer to **int64** |  | [optional] 
**UsageStorageAction** | Pointer to **string** |  | [optional] 
**UsageStorageThreshold** | Pointer to **int64** |  | [optional] 
**UsageUserAction** | Pointer to **string** |  | [optional] 
**UsageUserThreshold** | Pointer to **int64** |  | [optional] 

## Methods

### NewUpdateOrgRequest

`func NewUpdateOrgRequest() *UpdateOrgRequest`

NewUpdateOrgRequest instantiates a new UpdateOrgRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrgRequestWithDefaults

`func NewUpdateOrgRequestWithDefaults() *UpdateOrgRequest`

NewUpdateOrgRequestWithDefaults instantiates a new UpdateOrgRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *UpdateOrgRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateOrgRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateOrgRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateOrgRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHandle

`func (o *UpdateOrgRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *UpdateOrgRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *UpdateOrgRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *UpdateOrgRequest) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetTokenMinIssuedAt

`func (o *UpdateOrgRequest) GetTokenMinIssuedAt() string`

GetTokenMinIssuedAt returns the TokenMinIssuedAt field if non-nil, zero value otherwise.

### GetTokenMinIssuedAtOk

`func (o *UpdateOrgRequest) GetTokenMinIssuedAtOk() (*string, bool)`

GetTokenMinIssuedAtOk returns a tuple with the TokenMinIssuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenMinIssuedAt

`func (o *UpdateOrgRequest) SetTokenMinIssuedAt(v string)`

SetTokenMinIssuedAt sets TokenMinIssuedAt field to given value.

### HasTokenMinIssuedAt

`func (o *UpdateOrgRequest) HasTokenMinIssuedAt() bool`

HasTokenMinIssuedAt returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateOrgRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateOrgRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateOrgRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateOrgRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsageComputeAction

`func (o *UpdateOrgRequest) GetUsageComputeAction() string`

GetUsageComputeAction returns the UsageComputeAction field if non-nil, zero value otherwise.

### GetUsageComputeActionOk

`func (o *UpdateOrgRequest) GetUsageComputeActionOk() (*string, bool)`

GetUsageComputeActionOk returns a tuple with the UsageComputeAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageComputeAction

`func (o *UpdateOrgRequest) SetUsageComputeAction(v string)`

SetUsageComputeAction sets UsageComputeAction field to given value.

### HasUsageComputeAction

`func (o *UpdateOrgRequest) HasUsageComputeAction() bool`

HasUsageComputeAction returns a boolean if a field has been set.

### GetUsageComputeThreshold

`func (o *UpdateOrgRequest) GetUsageComputeThreshold() int64`

GetUsageComputeThreshold returns the UsageComputeThreshold field if non-nil, zero value otherwise.

### GetUsageComputeThresholdOk

`func (o *UpdateOrgRequest) GetUsageComputeThresholdOk() (*int64, bool)`

GetUsageComputeThresholdOk returns a tuple with the UsageComputeThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageComputeThreshold

`func (o *UpdateOrgRequest) SetUsageComputeThreshold(v int64)`

SetUsageComputeThreshold sets UsageComputeThreshold field to given value.

### HasUsageComputeThreshold

`func (o *UpdateOrgRequest) HasUsageComputeThreshold() bool`

HasUsageComputeThreshold returns a boolean if a field has been set.

### GetUsageStorageAction

`func (o *UpdateOrgRequest) GetUsageStorageAction() string`

GetUsageStorageAction returns the UsageStorageAction field if non-nil, zero value otherwise.

### GetUsageStorageActionOk

`func (o *UpdateOrgRequest) GetUsageStorageActionOk() (*string, bool)`

GetUsageStorageActionOk returns a tuple with the UsageStorageAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageStorageAction

`func (o *UpdateOrgRequest) SetUsageStorageAction(v string)`

SetUsageStorageAction sets UsageStorageAction field to given value.

### HasUsageStorageAction

`func (o *UpdateOrgRequest) HasUsageStorageAction() bool`

HasUsageStorageAction returns a boolean if a field has been set.

### GetUsageStorageThreshold

`func (o *UpdateOrgRequest) GetUsageStorageThreshold() int64`

GetUsageStorageThreshold returns the UsageStorageThreshold field if non-nil, zero value otherwise.

### GetUsageStorageThresholdOk

`func (o *UpdateOrgRequest) GetUsageStorageThresholdOk() (*int64, bool)`

GetUsageStorageThresholdOk returns a tuple with the UsageStorageThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageStorageThreshold

`func (o *UpdateOrgRequest) SetUsageStorageThreshold(v int64)`

SetUsageStorageThreshold sets UsageStorageThreshold field to given value.

### HasUsageStorageThreshold

`func (o *UpdateOrgRequest) HasUsageStorageThreshold() bool`

HasUsageStorageThreshold returns a boolean if a field has been set.

### GetUsageUserAction

`func (o *UpdateOrgRequest) GetUsageUserAction() string`

GetUsageUserAction returns the UsageUserAction field if non-nil, zero value otherwise.

### GetUsageUserActionOk

`func (o *UpdateOrgRequest) GetUsageUserActionOk() (*string, bool)`

GetUsageUserActionOk returns a tuple with the UsageUserAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageUserAction

`func (o *UpdateOrgRequest) SetUsageUserAction(v string)`

SetUsageUserAction sets UsageUserAction field to given value.

### HasUsageUserAction

`func (o *UpdateOrgRequest) HasUsageUserAction() bool`

HasUsageUserAction returns a boolean if a field has been set.

### GetUsageUserThreshold

`func (o *UpdateOrgRequest) GetUsageUserThreshold() int64`

GetUsageUserThreshold returns the UsageUserThreshold field if non-nil, zero value otherwise.

### GetUsageUserThresholdOk

`func (o *UpdateOrgRequest) GetUsageUserThresholdOk() (*int64, bool)`

GetUsageUserThresholdOk returns a tuple with the UsageUserThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageUserThreshold

`func (o *UpdateOrgRequest) SetUsageUserThreshold(v int64)`

SetUsageUserThreshold sets UsageUserThreshold field to given value.

### HasUsageUserThreshold

`func (o *UpdateOrgRequest) HasUsageUserThreshold() bool`

HasUsageUserThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


