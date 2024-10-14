# UpdateNotifierRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the notifier to update | [optional] 
**Notifies** | Pointer to **map[string]interface{}** | The notify targets for the notifier to update | [optional] 
**State** | Pointer to [**NotifierState**](NotifierState.md) | The state of the notifier to update | [optional] 

## Methods

### NewUpdateNotifierRequest

`func NewUpdateNotifierRequest() *UpdateNotifierRequest`

NewUpdateNotifierRequest instantiates a new UpdateNotifierRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNotifierRequestWithDefaults

`func NewUpdateNotifierRequestWithDefaults() *UpdateNotifierRequest`

NewUpdateNotifierRequestWithDefaults instantiates a new UpdateNotifierRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNotifierRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNotifierRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNotifierRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNotifierRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotifies

`func (o *UpdateNotifierRequest) GetNotifies() map[string]interface{}`

GetNotifies returns the Notifies field if non-nil, zero value otherwise.

### GetNotifiesOk

`func (o *UpdateNotifierRequest) GetNotifiesOk() (*map[string]interface{}, bool)`

GetNotifiesOk returns a tuple with the Notifies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifies

`func (o *UpdateNotifierRequest) SetNotifies(v map[string]interface{})`

SetNotifies sets Notifies field to given value.

### HasNotifies

`func (o *UpdateNotifierRequest) HasNotifies() bool`

HasNotifies returns a boolean if a field has been set.

### GetState

`func (o *UpdateNotifierRequest) GetState() NotifierState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateNotifierRequest) GetStateOk() (*NotifierState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateNotifierRequest) SetState(v NotifierState)`

SetState sets State field to given value.

### HasState

`func (o *UpdateNotifierRequest) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


