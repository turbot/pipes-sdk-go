# CreateAggregatorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** | The configuration for the aggregator. | [optional] 
**Connections** | Pointer to **[]string** | The connections that are a part of the aggregator. | [optional] 
**Handle** | **string** | The handle of the aggregator. | 
**Plugin** | **string** |  | 

## Methods

### NewCreateAggregatorRequest

`func NewCreateAggregatorRequest(handle string, plugin string, ) *CreateAggregatorRequest`

NewCreateAggregatorRequest instantiates a new CreateAggregatorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAggregatorRequestWithDefaults

`func NewCreateAggregatorRequestWithDefaults() *CreateAggregatorRequest`

NewCreateAggregatorRequestWithDefaults instantiates a new CreateAggregatorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *CreateAggregatorRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateAggregatorRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateAggregatorRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateAggregatorRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConnections

`func (o *CreateAggregatorRequest) GetConnections() []string`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *CreateAggregatorRequest) GetConnectionsOk() (*[]string, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *CreateAggregatorRequest) SetConnections(v []string)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *CreateAggregatorRequest) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetHandle

`func (o *CreateAggregatorRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *CreateAggregatorRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *CreateAggregatorRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetPlugin

`func (o *CreateAggregatorRequest) GetPlugin() string`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *CreateAggregatorRequest) GetPluginOk() (*string, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *CreateAggregatorRequest) SetPlugin(v string)`

SetPlugin sets Plugin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


