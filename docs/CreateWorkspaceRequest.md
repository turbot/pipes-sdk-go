# CreateWorkspaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbVolumeSizeBytes** | Pointer to **int64** |  | [optional] 
**DesiredState** | Pointer to **string** | paused is not yet supported for Workspace | [optional] 
**Handle** | **string** |  | 
**InstanceType** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateWorkspaceRequest

`func NewCreateWorkspaceRequest(handle string, ) *CreateWorkspaceRequest`

NewCreateWorkspaceRequest instantiates a new CreateWorkspaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkspaceRequestWithDefaults

`func NewCreateWorkspaceRequestWithDefaults() *CreateWorkspaceRequest`

NewCreateWorkspaceRequestWithDefaults instantiates a new CreateWorkspaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbVolumeSizeBytes

`func (o *CreateWorkspaceRequest) GetDbVolumeSizeBytes() int64`

GetDbVolumeSizeBytes returns the DbVolumeSizeBytes field if non-nil, zero value otherwise.

### GetDbVolumeSizeBytesOk

`func (o *CreateWorkspaceRequest) GetDbVolumeSizeBytesOk() (*int64, bool)`

GetDbVolumeSizeBytesOk returns a tuple with the DbVolumeSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbVolumeSizeBytes

`func (o *CreateWorkspaceRequest) SetDbVolumeSizeBytes(v int64)`

SetDbVolumeSizeBytes sets DbVolumeSizeBytes field to given value.

### HasDbVolumeSizeBytes

`func (o *CreateWorkspaceRequest) HasDbVolumeSizeBytes() bool`

HasDbVolumeSizeBytes returns a boolean if a field has been set.

### GetDesiredState

`func (o *CreateWorkspaceRequest) GetDesiredState() string`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *CreateWorkspaceRequest) GetDesiredStateOk() (*string, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *CreateWorkspaceRequest) SetDesiredState(v string)`

SetDesiredState sets DesiredState field to given value.

### HasDesiredState

`func (o *CreateWorkspaceRequest) HasDesiredState() bool`

HasDesiredState returns a boolean if a field has been set.

### GetHandle

`func (o *CreateWorkspaceRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *CreateWorkspaceRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *CreateWorkspaceRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetInstanceType

`func (o *CreateWorkspaceRequest) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *CreateWorkspaceRequest) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *CreateWorkspaceRequest) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *CreateWorkspaceRequest) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


