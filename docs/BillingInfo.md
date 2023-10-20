# BillingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CycleAnchor** | Pointer to **string** | The time when the org was created in ISO 8601 UTC. | [optional] 
**DeleteAt** | Pointer to **string** | The time when the identity will be deleted from Pipes. | [optional] 
**DeleteAtReason** | Pointer to **string** | The reason for the upcoming deletion. | [optional] 
**Status** | **string** | The billing status for the org. | 
**StatusUpdatedAt** | Pointer to **string** | When the billing status for the org was last updated. | [optional] 
**SuspendAt** | Pointer to **string** | The time when the org will be suspended due to failed billing in ISO 8601 UTC. | [optional] 
**SuspendAtReason** | Pointer to **string** | The reason for the upcoming suspension. | [optional] 

## Methods

### NewBillingInfo

`func NewBillingInfo(status string, ) *BillingInfo`

NewBillingInfo instantiates a new BillingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingInfoWithDefaults

`func NewBillingInfoWithDefaults() *BillingInfo`

NewBillingInfoWithDefaults instantiates a new BillingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCycleAnchor

`func (o *BillingInfo) GetCycleAnchor() string`

GetCycleAnchor returns the CycleAnchor field if non-nil, zero value otherwise.

### GetCycleAnchorOk

`func (o *BillingInfo) GetCycleAnchorOk() (*string, bool)`

GetCycleAnchorOk returns a tuple with the CycleAnchor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCycleAnchor

`func (o *BillingInfo) SetCycleAnchor(v string)`

SetCycleAnchor sets CycleAnchor field to given value.

### HasCycleAnchor

`func (o *BillingInfo) HasCycleAnchor() bool`

HasCycleAnchor returns a boolean if a field has been set.

### GetDeleteAt

`func (o *BillingInfo) GetDeleteAt() string`

GetDeleteAt returns the DeleteAt field if non-nil, zero value otherwise.

### GetDeleteAtOk

`func (o *BillingInfo) GetDeleteAtOk() (*string, bool)`

GetDeleteAtOk returns a tuple with the DeleteAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAt

`func (o *BillingInfo) SetDeleteAt(v string)`

SetDeleteAt sets DeleteAt field to given value.

### HasDeleteAt

`func (o *BillingInfo) HasDeleteAt() bool`

HasDeleteAt returns a boolean if a field has been set.

### GetDeleteAtReason

`func (o *BillingInfo) GetDeleteAtReason() string`

GetDeleteAtReason returns the DeleteAtReason field if non-nil, zero value otherwise.

### GetDeleteAtReasonOk

`func (o *BillingInfo) GetDeleteAtReasonOk() (*string, bool)`

GetDeleteAtReasonOk returns a tuple with the DeleteAtReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAtReason

`func (o *BillingInfo) SetDeleteAtReason(v string)`

SetDeleteAtReason sets DeleteAtReason field to given value.

### HasDeleteAtReason

`func (o *BillingInfo) HasDeleteAtReason() bool`

HasDeleteAtReason returns a boolean if a field has been set.

### GetStatus

`func (o *BillingInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillingInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillingInfo) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStatusUpdatedAt

`func (o *BillingInfo) GetStatusUpdatedAt() string`

GetStatusUpdatedAt returns the StatusUpdatedAt field if non-nil, zero value otherwise.

### GetStatusUpdatedAtOk

`func (o *BillingInfo) GetStatusUpdatedAtOk() (*string, bool)`

GetStatusUpdatedAtOk returns a tuple with the StatusUpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUpdatedAt

`func (o *BillingInfo) SetStatusUpdatedAt(v string)`

SetStatusUpdatedAt sets StatusUpdatedAt field to given value.

### HasStatusUpdatedAt

`func (o *BillingInfo) HasStatusUpdatedAt() bool`

HasStatusUpdatedAt returns a boolean if a field has been set.

### GetSuspendAt

`func (o *BillingInfo) GetSuspendAt() string`

GetSuspendAt returns the SuspendAt field if non-nil, zero value otherwise.

### GetSuspendAtOk

`func (o *BillingInfo) GetSuspendAtOk() (*string, bool)`

GetSuspendAtOk returns a tuple with the SuspendAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendAt

`func (o *BillingInfo) SetSuspendAt(v string)`

SetSuspendAt sets SuspendAt field to given value.

### HasSuspendAt

`func (o *BillingInfo) HasSuspendAt() bool`

HasSuspendAt returns a boolean if a field has been set.

### GetSuspendAtReason

`func (o *BillingInfo) GetSuspendAtReason() string`

GetSuspendAtReason returns the SuspendAtReason field if non-nil, zero value otherwise.

### GetSuspendAtReasonOk

`func (o *BillingInfo) GetSuspendAtReasonOk() (*string, bool)`

GetSuspendAtReasonOk returns a tuple with the SuspendAtReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendAtReason

`func (o *BillingInfo) SetSuspendAtReason(v string)`

SetSuspendAtReason sets SuspendAtReason field to given value.

### HasSuspendAtReason

`func (o *BillingInfo) HasSuspendAtReason() bool`

HasSuspendAtReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


