# CreateNotifierRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the notifier to create | 
**Notifies** | **map[string]interface{}** | The notify targets for the notifier to create | 
**State** | Pointer to [**NotifierState**](NotifierState.md) | The state of the notifier to create | [optional] 

## Methods

### NewCreateNotifierRequest

`func NewCreateNotifierRequest(name string, notifies map[string]interface{}, ) *CreateNotifierRequest`

NewCreateNotifierRequest instantiates a new CreateNotifierRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNotifierRequestWithDefaults

`func NewCreateNotifierRequestWithDefaults() *CreateNotifierRequest`

NewCreateNotifierRequestWithDefaults instantiates a new CreateNotifierRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateNotifierRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateNotifierRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateNotifierRequest) SetName(v string)`

SetName sets Name field to given value.


### GetNotifies

`func (o *CreateNotifierRequest) GetNotifies() map[string]interface{}`

GetNotifies returns the Notifies field if non-nil, zero value otherwise.

### GetNotifiesOk

`func (o *CreateNotifierRequest) GetNotifiesOk() (*map[string]interface{}, bool)`

GetNotifiesOk returns a tuple with the Notifies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifies

`func (o *CreateNotifierRequest) SetNotifies(v map[string]interface{})`

SetNotifies sets Notifies field to given value.


### GetState

`func (o *CreateNotifierRequest) GetState() NotifierState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateNotifierRequest) GetStateOk() (*NotifierState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateNotifierRequest) SetState(v NotifierState)`

SetState sets State field to given value.

### HasState

`func (o *CreateNotifierRequest) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


