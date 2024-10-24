# CreateIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** | The configuration for the integration. | [optional] 
**Handle** | **string** | The handle for the integration. | 
**Type** | [**IntegrationType**](IntegrationType.md) | The type of the integration, can be one of &#x60;aws&#x60;, &#x60;azure&#x60;, &#x60;gcp&#x60; or &#x60;github&#x60;. | 

## Methods

### NewCreateIntegrationRequest

`func NewCreateIntegrationRequest(handle string, type_ IntegrationType, ) *CreateIntegrationRequest`

NewCreateIntegrationRequest instantiates a new CreateIntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIntegrationRequestWithDefaults

`func NewCreateIntegrationRequestWithDefaults() *CreateIntegrationRequest`

NewCreateIntegrationRequestWithDefaults instantiates a new CreateIntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *CreateIntegrationRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateIntegrationRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateIntegrationRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateIntegrationRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetHandle

`func (o *CreateIntegrationRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *CreateIntegrationRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *CreateIntegrationRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetType

`func (o *CreateIntegrationRequest) GetType() IntegrationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateIntegrationRequest) GetTypeOk() (*IntegrationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateIntegrationRequest) SetType(v IntegrationType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


