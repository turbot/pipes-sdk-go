# UpdateIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** | The configuration for the integration. | [optional] 
**Handle** | Pointer to **string** | The handle for the integration. | [optional] 
**State** | Pointer to [**IntegrationState**](IntegrationState.md) | The state of the integration. | [optional] 

## Methods

### NewUpdateIntegrationRequest

`func NewUpdateIntegrationRequest() *UpdateIntegrationRequest`

NewUpdateIntegrationRequest instantiates a new UpdateIntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIntegrationRequestWithDefaults

`func NewUpdateIntegrationRequestWithDefaults() *UpdateIntegrationRequest`

NewUpdateIntegrationRequestWithDefaults instantiates a new UpdateIntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *UpdateIntegrationRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateIntegrationRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateIntegrationRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateIntegrationRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetHandle

`func (o *UpdateIntegrationRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *UpdateIntegrationRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *UpdateIntegrationRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *UpdateIntegrationRequest) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetState

`func (o *UpdateIntegrationRequest) GetState() IntegrationState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateIntegrationRequest) GetStateOk() (*IntegrationState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateIntegrationRequest) SetState(v IntegrationState)`

SetState sets State field to given value.

### HasState

`func (o *UpdateIntegrationRequest) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


