# DatatankTableFreshness

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **int32** | The number of parts that are disabled. | [optional] 
**Error** | Pointer to **int32** | The number of parts that have never successfully run. | [optional] 
**Expired** | Pointer to **int32** | The number of parts that did not have a successful refresh within the extended staleness timeframe. | [optional] 
**Fresh** | Pointer to **int32** | The number of parts that had a successful query within the staleness timeframe. | [optional] 
**NewestPartUpdatedAt** | Pointer to [**JSONTime**](JSONTime.md) |  | [optional] 
**OldestPartUpdatedAt** | Pointer to [**JSONTime**](JSONTime.md) |  | [optional] 
**Pending** | Pointer to **int32** | The number of parts that are yet to be executed. | [optional] 
**Removing** | Pointer to **int32** | The number of parts for which the data is in the process of being deleted. | [optional] 
**Stale** | Pointer to **int32** | The number of parts that had a successful query but not within the staleness timeframe. | [optional] 
**TotalParts** | Pointer to **int32** | The total number of parts in the datatank table. | [optional] 

## Methods

### NewDatatankTableFreshness

`func NewDatatankTableFreshness() *DatatankTableFreshness`

NewDatatankTableFreshness instantiates a new DatatankTableFreshness object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatatankTableFreshnessWithDefaults

`func NewDatatankTableFreshnessWithDefaults() *DatatankTableFreshness`

NewDatatankTableFreshnessWithDefaults instantiates a new DatatankTableFreshness object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *DatatankTableFreshness) GetDisabled() int32`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *DatatankTableFreshness) GetDisabledOk() (*int32, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *DatatankTableFreshness) SetDisabled(v int32)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *DatatankTableFreshness) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetError

`func (o *DatatankTableFreshness) GetError() int32`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *DatatankTableFreshness) GetErrorOk() (*int32, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *DatatankTableFreshness) SetError(v int32)`

SetError sets Error field to given value.

### HasError

`func (o *DatatankTableFreshness) HasError() bool`

HasError returns a boolean if a field has been set.

### GetExpired

`func (o *DatatankTableFreshness) GetExpired() int32`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *DatatankTableFreshness) GetExpiredOk() (*int32, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *DatatankTableFreshness) SetExpired(v int32)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *DatatankTableFreshness) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetFresh

`func (o *DatatankTableFreshness) GetFresh() int32`

GetFresh returns the Fresh field if non-nil, zero value otherwise.

### GetFreshOk

`func (o *DatatankTableFreshness) GetFreshOk() (*int32, bool)`

GetFreshOk returns a tuple with the Fresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFresh

`func (o *DatatankTableFreshness) SetFresh(v int32)`

SetFresh sets Fresh field to given value.

### HasFresh

`func (o *DatatankTableFreshness) HasFresh() bool`

HasFresh returns a boolean if a field has been set.

### GetNewestPartUpdatedAt

`func (o *DatatankTableFreshness) GetNewestPartUpdatedAt() JSONTime`

GetNewestPartUpdatedAt returns the NewestPartUpdatedAt field if non-nil, zero value otherwise.

### GetNewestPartUpdatedAtOk

`func (o *DatatankTableFreshness) GetNewestPartUpdatedAtOk() (*JSONTime, bool)`

GetNewestPartUpdatedAtOk returns a tuple with the NewestPartUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewestPartUpdatedAt

`func (o *DatatankTableFreshness) SetNewestPartUpdatedAt(v JSONTime)`

SetNewestPartUpdatedAt sets NewestPartUpdatedAt field to given value.

### HasNewestPartUpdatedAt

`func (o *DatatankTableFreshness) HasNewestPartUpdatedAt() bool`

HasNewestPartUpdatedAt returns a boolean if a field has been set.

### GetOldestPartUpdatedAt

`func (o *DatatankTableFreshness) GetOldestPartUpdatedAt() JSONTime`

GetOldestPartUpdatedAt returns the OldestPartUpdatedAt field if non-nil, zero value otherwise.

### GetOldestPartUpdatedAtOk

`func (o *DatatankTableFreshness) GetOldestPartUpdatedAtOk() (*JSONTime, bool)`

GetOldestPartUpdatedAtOk returns a tuple with the OldestPartUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldestPartUpdatedAt

`func (o *DatatankTableFreshness) SetOldestPartUpdatedAt(v JSONTime)`

SetOldestPartUpdatedAt sets OldestPartUpdatedAt field to given value.

### HasOldestPartUpdatedAt

`func (o *DatatankTableFreshness) HasOldestPartUpdatedAt() bool`

HasOldestPartUpdatedAt returns a boolean if a field has been set.

### GetPending

`func (o *DatatankTableFreshness) GetPending() int32`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *DatatankTableFreshness) GetPendingOk() (*int32, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *DatatankTableFreshness) SetPending(v int32)`

SetPending sets Pending field to given value.

### HasPending

`func (o *DatatankTableFreshness) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetRemoving

`func (o *DatatankTableFreshness) GetRemoving() int32`

GetRemoving returns the Removing field if non-nil, zero value otherwise.

### GetRemovingOk

`func (o *DatatankTableFreshness) GetRemovingOk() (*int32, bool)`

GetRemovingOk returns a tuple with the Removing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoving

`func (o *DatatankTableFreshness) SetRemoving(v int32)`

SetRemoving sets Removing field to given value.

### HasRemoving

`func (o *DatatankTableFreshness) HasRemoving() bool`

HasRemoving returns a boolean if a field has been set.

### GetStale

`func (o *DatatankTableFreshness) GetStale() int32`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *DatatankTableFreshness) GetStaleOk() (*int32, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *DatatankTableFreshness) SetStale(v int32)`

SetStale sets Stale field to given value.

### HasStale

`func (o *DatatankTableFreshness) HasStale() bool`

HasStale returns a boolean if a field has been set.

### GetTotalParts

`func (o *DatatankTableFreshness) GetTotalParts() int32`

GetTotalParts returns the TotalParts field if non-nil, zero value otherwise.

### GetTotalPartsOk

`func (o *DatatankTableFreshness) GetTotalPartsOk() (*int32, bool)`

GetTotalPartsOk returns a tuple with the TotalParts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalParts

`func (o *DatatankTableFreshness) SetTotalParts(v int32)`

SetTotalParts sets TotalParts field to given value.

### HasTotalParts

`func (o *DatatankTableFreshness) HasTotalParts() bool`

HasTotalParts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


