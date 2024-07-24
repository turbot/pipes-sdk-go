/*
Turbot Pipes

Turbot Pipes is an intelligence, automation & security platform built specifically for DevOps.

API version: {{OPEN_API_VERSION}}
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipes

import (
	"encoding/json"
)

// CreateWorkspaceNotificationRequest struct for CreateWorkspaceNotificationRequest
type CreateWorkspaceNotificationRequest struct {
	Events []string                      `json:"events"`
	Sender *NotificationRecipientRequest `json:"sender,omitempty"`
	Title  string                        `json:"title"`
}

// NewCreateWorkspaceNotificationRequest instantiates a new CreateWorkspaceNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWorkspaceNotificationRequest(events []string, title string) *CreateWorkspaceNotificationRequest {
	this := CreateWorkspaceNotificationRequest{}
	this.Events = events
	this.Title = title
	return &this
}

// NewCreateWorkspaceNotificationRequestWithDefaults instantiates a new CreateWorkspaceNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWorkspaceNotificationRequestWithDefaults() *CreateWorkspaceNotificationRequest {
	this := CreateWorkspaceNotificationRequest{}
	return &this
}

// GetEvents returns the Events field value
func (o *CreateWorkspaceNotificationRequest) GetEvents() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceNotificationRequest) GetEventsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Events, true
}

// SetEvents sets field value
func (o *CreateWorkspaceNotificationRequest) SetEvents(v []string) {
	o.Events = v
}

// GetSender returns the Sender field value if set, zero value otherwise.
func (o *CreateWorkspaceNotificationRequest) GetSender() NotificationRecipientRequest {
	if o == nil || o.Sender == nil {
		var ret NotificationRecipientRequest
		return ret
	}
	return *o.Sender
}

// GetSenderOk returns a tuple with the Sender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceNotificationRequest) GetSenderOk() (*NotificationRecipientRequest, bool) {
	if o == nil || o.Sender == nil {
		return nil, false
	}
	return o.Sender, true
}

// HasSender returns a boolean if a field has been set.
func (o *CreateWorkspaceNotificationRequest) HasSender() bool {
	if o != nil && o.Sender != nil {
		return true
	}

	return false
}

// SetSender gets a reference to the given NotificationRecipientRequest and assigns it to the Sender field.
func (o *CreateWorkspaceNotificationRequest) SetSender(v NotificationRecipientRequest) {
	o.Sender = &v
}

// GetTitle returns the Title field value
func (o *CreateWorkspaceNotificationRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceNotificationRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CreateWorkspaceNotificationRequest) SetTitle(v string) {
	o.Title = v
}

func (o CreateWorkspaceNotificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["events"] = o.Events
	}
	if o.Sender != nil {
		toSerialize["sender"] = o.Sender
	}
	if true {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableCreateWorkspaceNotificationRequest struct {
	value *CreateWorkspaceNotificationRequest
	isSet bool
}

func (v NullableCreateWorkspaceNotificationRequest) Get() *CreateWorkspaceNotificationRequest {
	return v.value
}

func (v *NullableCreateWorkspaceNotificationRequest) Set(val *CreateWorkspaceNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWorkspaceNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWorkspaceNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWorkspaceNotificationRequest(val *CreateWorkspaceNotificationRequest) *NullableCreateWorkspaceNotificationRequest {
	return &NullableCreateWorkspaceNotificationRequest{value: val, isSet: true}
}

func (v NullableCreateWorkspaceNotificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWorkspaceNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}