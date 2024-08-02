# UpdatePermissionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentityHandle** | Pointer to **string** |  | [optional] 
**TenantHandle** | Pointer to **string** |  | [optional] 
**WorkspaceHandle** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdatePermissionRequest

`func NewUpdatePermissionRequest() *UpdatePermissionRequest`

NewUpdatePermissionRequest instantiates a new UpdatePermissionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePermissionRequestWithDefaults

`func NewUpdatePermissionRequestWithDefaults() *UpdatePermissionRequest`

NewUpdatePermissionRequestWithDefaults instantiates a new UpdatePermissionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentityHandle

`func (o *UpdatePermissionRequest) GetIdentityHandle() string`

GetIdentityHandle returns the IdentityHandle field if non-nil, zero value otherwise.

### GetIdentityHandleOk

`func (o *UpdatePermissionRequest) GetIdentityHandleOk() (*string, bool)`

GetIdentityHandleOk returns a tuple with the IdentityHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityHandle

`func (o *UpdatePermissionRequest) SetIdentityHandle(v string)`

SetIdentityHandle sets IdentityHandle field to given value.

### HasIdentityHandle

`func (o *UpdatePermissionRequest) HasIdentityHandle() bool`

HasIdentityHandle returns a boolean if a field has been set.

### GetTenantHandle

`func (o *UpdatePermissionRequest) GetTenantHandle() string`

GetTenantHandle returns the TenantHandle field if non-nil, zero value otherwise.

### GetTenantHandleOk

`func (o *UpdatePermissionRequest) GetTenantHandleOk() (*string, bool)`

GetTenantHandleOk returns a tuple with the TenantHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantHandle

`func (o *UpdatePermissionRequest) SetTenantHandle(v string)`

SetTenantHandle sets TenantHandle field to given value.

### HasTenantHandle

`func (o *UpdatePermissionRequest) HasTenantHandle() bool`

HasTenantHandle returns a boolean if a field has been set.

### GetWorkspaceHandle

`func (o *UpdatePermissionRequest) GetWorkspaceHandle() string`

GetWorkspaceHandle returns the WorkspaceHandle field if non-nil, zero value otherwise.

### GetWorkspaceHandleOk

`func (o *UpdatePermissionRequest) GetWorkspaceHandleOk() (*string, bool)`

GetWorkspaceHandleOk returns a tuple with the WorkspaceHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceHandle

`func (o *UpdatePermissionRequest) SetWorkspaceHandle(v string)`

SetWorkspaceHandle sets WorkspaceHandle field to given value.

### HasWorkspaceHandle

`func (o *UpdatePermissionRequest) HasWorkspaceHandle() bool`

HasWorkspaceHandle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


