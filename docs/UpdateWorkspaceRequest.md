# UpdateWorkspaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbVolumeSizeBytes** | Pointer to **int64** |  | [optional] 
**DesiredState** | Pointer to [**DesiredState**](DesiredState.md) | paused is not yet supported for Workspace | [optional] 
**Handle** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to [**WorkspaceInstanceType**](WorkspaceInstanceType.md) | the updated instance type of the workspace | [optional] 
**SearchPathPrefix** | Pointer to **[]string** | Optionally update the search path prefix for the Steampipe DB. | [optional] 

## Methods

### NewUpdateWorkspaceRequest

`func NewUpdateWorkspaceRequest() *UpdateWorkspaceRequest`

NewUpdateWorkspaceRequest instantiates a new UpdateWorkspaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWorkspaceRequestWithDefaults

`func NewUpdateWorkspaceRequestWithDefaults() *UpdateWorkspaceRequest`

NewUpdateWorkspaceRequestWithDefaults instantiates a new UpdateWorkspaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbVolumeSizeBytes

`func (o *UpdateWorkspaceRequest) GetDbVolumeSizeBytes() int64`

GetDbVolumeSizeBytes returns the DbVolumeSizeBytes field if non-nil, zero value otherwise.

### GetDbVolumeSizeBytesOk

`func (o *UpdateWorkspaceRequest) GetDbVolumeSizeBytesOk() (*int64, bool)`

GetDbVolumeSizeBytesOk returns a tuple with the DbVolumeSizeBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbVolumeSizeBytes

`func (o *UpdateWorkspaceRequest) SetDbVolumeSizeBytes(v int64)`

SetDbVolumeSizeBytes sets DbVolumeSizeBytes field to given value.

### HasDbVolumeSizeBytes

`func (o *UpdateWorkspaceRequest) HasDbVolumeSizeBytes() bool`

HasDbVolumeSizeBytes returns a boolean if a field has been set.

### GetDesiredState

`func (o *UpdateWorkspaceRequest) GetDesiredState() DesiredState`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *UpdateWorkspaceRequest) GetDesiredStateOk() (*DesiredState, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *UpdateWorkspaceRequest) SetDesiredState(v DesiredState)`

SetDesiredState sets DesiredState field to given value.

### HasDesiredState

`func (o *UpdateWorkspaceRequest) HasDesiredState() bool`

HasDesiredState returns a boolean if a field has been set.

### GetHandle

`func (o *UpdateWorkspaceRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *UpdateWorkspaceRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *UpdateWorkspaceRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *UpdateWorkspaceRequest) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetInstanceType

`func (o *UpdateWorkspaceRequest) GetInstanceType() WorkspaceInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *UpdateWorkspaceRequest) GetInstanceTypeOk() (*WorkspaceInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *UpdateWorkspaceRequest) SetInstanceType(v WorkspaceInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *UpdateWorkspaceRequest) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetSearchPathPrefix

`func (o *UpdateWorkspaceRequest) GetSearchPathPrefix() []string`

GetSearchPathPrefix returns the SearchPathPrefix field if non-nil, zero value otherwise.

### GetSearchPathPrefixOk

`func (o *UpdateWorkspaceRequest) GetSearchPathPrefixOk() (*[]string, bool)`

GetSearchPathPrefixOk returns a tuple with the SearchPathPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchPathPrefix

`func (o *UpdateWorkspaceRequest) SetSearchPathPrefix(v []string)`

SetSearchPathPrefix sets SearchPathPrefix field to given value.

### HasSearchPathPrefix

`func (o *UpdateWorkspaceRequest) HasSearchPathPrefix() bool`

HasSearchPathPrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


