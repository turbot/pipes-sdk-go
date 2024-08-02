# UpdateAggregatorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Connections** | Pointer to **[]string** | The connections that are a part of the aggregator. | [optional] 
**Handle** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateAggregatorRequest

`func NewUpdateAggregatorRequest() *UpdateAggregatorRequest`

NewUpdateAggregatorRequest instantiates a new UpdateAggregatorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAggregatorRequestWithDefaults

`func NewUpdateAggregatorRequestWithDefaults() *UpdateAggregatorRequest`

NewUpdateAggregatorRequestWithDefaults instantiates a new UpdateAggregatorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *UpdateAggregatorRequest) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateAggregatorRequest) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateAggregatorRequest) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateAggregatorRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConnections

`func (o *UpdateAggregatorRequest) GetConnections() []string`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *UpdateAggregatorRequest) GetConnectionsOk() (*[]string, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *UpdateAggregatorRequest) SetConnections(v []string)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *UpdateAggregatorRequest) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetHandle

`func (o *UpdateAggregatorRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *UpdateAggregatorRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *UpdateAggregatorRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *UpdateAggregatorRequest) HasHandle() bool`

HasHandle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


