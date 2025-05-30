# UpdateConnectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**ConfigSource** | Pointer to [**ConnectionConfigSource**](ConnectionConfigSource.md) |  | [optional] 
**CredentialSource** | Pointer to [**ConnectionCredentialSource**](ConnectionCredentialSource.md) |  | [optional] 
**Handle** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateConnectionRequest

`func NewUpdateConnectionRequest() *UpdateConnectionRequest`

NewUpdateConnectionRequest instantiates a new UpdateConnectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateConnectionRequestWithDefaults

`func NewUpdateConnectionRequestWithDefaults() *UpdateConnectionRequest`

NewUpdateConnectionRequestWithDefaults instantiates a new UpdateConnectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *UpdateConnectionRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateConnectionRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateConnectionRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateConnectionRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConfigSource

`func (o *UpdateConnectionRequest) GetConfigSource() ConnectionConfigSource`

GetConfigSource returns the ConfigSource field if non-nil, zero value otherwise.

### GetConfigSourceOk

`func (o *UpdateConnectionRequest) GetConfigSourceOk() (*ConnectionConfigSource, bool)`

GetConfigSourceOk returns a tuple with the ConfigSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigSource

`func (o *UpdateConnectionRequest) SetConfigSource(v ConnectionConfigSource)`

SetConfigSource sets ConfigSource field to given value.

### HasConfigSource

`func (o *UpdateConnectionRequest) HasConfigSource() bool`

HasConfigSource returns a boolean if a field has been set.

### GetCredentialSource

`func (o *UpdateConnectionRequest) GetCredentialSource() ConnectionCredentialSource`

GetCredentialSource returns the CredentialSource field if non-nil, zero value otherwise.

### GetCredentialSourceOk

`func (o *UpdateConnectionRequest) GetCredentialSourceOk() (*ConnectionCredentialSource, bool)`

GetCredentialSourceOk returns a tuple with the CredentialSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialSource

`func (o *UpdateConnectionRequest) SetCredentialSource(v ConnectionCredentialSource)`

SetCredentialSource sets CredentialSource field to given value.

### HasCredentialSource

`func (o *UpdateConnectionRequest) HasCredentialSource() bool`

HasCredentialSource returns a boolean if a field has been set.

### GetHandle

`func (o *UpdateConnectionRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *UpdateConnectionRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *UpdateConnectionRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *UpdateConnectionRequest) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetParentId

`func (o *UpdateConnectionRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *UpdateConnectionRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *UpdateConnectionRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *UpdateConnectionRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


