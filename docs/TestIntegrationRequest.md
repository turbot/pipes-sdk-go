# TestIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Type** | Pointer to [**IntegrationType**](IntegrationType.md) |  | [optional] 

## Methods

### NewTestIntegrationRequest

`func NewTestIntegrationRequest() *TestIntegrationRequest`

NewTestIntegrationRequest instantiates a new TestIntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestIntegrationRequestWithDefaults

`func NewTestIntegrationRequestWithDefaults() *TestIntegrationRequest`

NewTestIntegrationRequestWithDefaults instantiates a new TestIntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *TestIntegrationRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *TestIntegrationRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *TestIntegrationRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *TestIntegrationRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetType

`func (o *TestIntegrationRequest) GetType() IntegrationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TestIntegrationRequest) GetTypeOk() (*IntegrationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TestIntegrationRequest) SetType(v IntegrationType)`

SetType sets Type field to given value.

### HasType

`func (o *TestIntegrationRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


