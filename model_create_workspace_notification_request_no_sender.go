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

// CreateWorkspaceNotificationRequestNoSender struct for CreateWorkspaceNotificationRequestNoSender
type CreateWorkspaceNotificationRequestNoSender struct {
	Events []string `json:"events"`
	Title  string   `json:"title"`
}

// NewCreateWorkspaceNotificationRequestNoSender instantiates a new CreateWorkspaceNotificationRequestNoSender object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWorkspaceNotificationRequestNoSender(events []string, title string) *CreateWorkspaceNotificationRequestNoSender {
	this := CreateWorkspaceNotificationRequestNoSender{}
	this.Events = events
	this.Title = title
	return &this
}

// NewCreateWorkspaceNotificationRequestNoSenderWithDefaults instantiates a new CreateWorkspaceNotificationRequestNoSender object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWorkspaceNotificationRequestNoSenderWithDefaults() *CreateWorkspaceNotificationRequestNoSender {
	this := CreateWorkspaceNotificationRequestNoSender{}
	return &this
}

// GetEvents returns the Events field value
func (o *CreateWorkspaceNotificationRequestNoSender) GetEvents() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceNotificationRequestNoSender) GetEventsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Events, true
}

// SetEvents sets field value
func (o *CreateWorkspaceNotificationRequestNoSender) SetEvents(v []string) {
	o.Events = v
}

// GetTitle returns the Title field value
func (o *CreateWorkspaceNotificationRequestNoSender) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *CreateWorkspaceNotificationRequestNoSender) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *CreateWorkspaceNotificationRequestNoSender) SetTitle(v string) {
	o.Title = v
}

func (o CreateWorkspaceNotificationRequestNoSender) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["events"] = o.Events
	}
	if true {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableCreateWorkspaceNotificationRequestNoSender struct {
	value *CreateWorkspaceNotificationRequestNoSender
	isSet bool
}

func (v NullableCreateWorkspaceNotificationRequestNoSender) Get() *CreateWorkspaceNotificationRequestNoSender {
	return v.value
}

func (v *NullableCreateWorkspaceNotificationRequestNoSender) Set(val *CreateWorkspaceNotificationRequestNoSender) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWorkspaceNotificationRequestNoSender) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWorkspaceNotificationRequestNoSender) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWorkspaceNotificationRequestNoSender(val *CreateWorkspaceNotificationRequestNoSender) *NullableCreateWorkspaceNotificationRequestNoSender {
	return &NullableCreateWorkspaceNotificationRequestNoSender{value: val, isSet: true}
}

func (v NullableCreateWorkspaceNotificationRequestNoSender) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWorkspaceNotificationRequestNoSender) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
