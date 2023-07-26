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

// TemporaryTokenRequest struct for TemporaryTokenRequest
type TemporaryTokenRequest struct {
	ClientIp  string  `json:"client_ip"`
	Code      string  `json:"code"`
	CreatedAt string  `json:"created_at"`
	Id        string  `json:"id"`
	State     string  `json:"state"`
	Token     *string `json:"token,omitempty"`
	// The time of the last update in ISO 8601 UTC.
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewTemporaryTokenRequest instantiates a new TemporaryTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemporaryTokenRequest(clientIp string, code string, createdAt string, id string, state string) *TemporaryTokenRequest {
	this := TemporaryTokenRequest{}
	this.ClientIp = clientIp
	this.Code = code
	this.CreatedAt = createdAt
	this.Id = id
	this.State = state
	return &this
}

// NewTemporaryTokenRequestWithDefaults instantiates a new TemporaryTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemporaryTokenRequestWithDefaults() *TemporaryTokenRequest {
	this := TemporaryTokenRequest{}
	return &this
}

// GetClientIp returns the ClientIp field value
func (o *TemporaryTokenRequest) GetClientIp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientIp
}

// GetClientIpOk returns a tuple with the ClientIp field value
// and a boolean to check if the value has been set.
func (o *TemporaryTokenRequest) GetClientIpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientIp, true
}

// SetClientIp sets field value
func (o *TemporaryTokenRequest) SetClientIp(v string) {
	o.ClientIp = v
}

// GetCode returns the Code field value
func (o *TemporaryTokenRequest) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *TemporaryTokenRequest) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *TemporaryTokenRequest) SetCode(v string) {
	o.Code = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TemporaryTokenRequest) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TemporaryTokenRequest) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TemporaryTokenRequest) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetId returns the Id field value
func (o *TemporaryTokenRequest) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TemporaryTokenRequest) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TemporaryTokenRequest) SetId(v string) {
	o.Id = v
}

// GetState returns the State field value
func (o *TemporaryTokenRequest) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *TemporaryTokenRequest) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *TemporaryTokenRequest) SetState(v string) {
	o.State = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *TemporaryTokenRequest) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemporaryTokenRequest) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *TemporaryTokenRequest) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *TemporaryTokenRequest) SetToken(v string) {
	o.Token = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *TemporaryTokenRequest) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemporaryTokenRequest) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TemporaryTokenRequest) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *TemporaryTokenRequest) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o TemporaryTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["client_ip"] = o.ClientIp
	}
	if true {
		toSerialize["code"] = o.Code
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["state"] = o.State
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableTemporaryTokenRequest struct {
	value *TemporaryTokenRequest
	isSet bool
}

func (v NullableTemporaryTokenRequest) Get() *TemporaryTokenRequest {
	return v.value
}

func (v *NullableTemporaryTokenRequest) Set(val *TemporaryTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTemporaryTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTemporaryTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemporaryTokenRequest(val *TemporaryTokenRequest) *NullableTemporaryTokenRequest {
	return &NullableTemporaryTokenRequest{value: val, isSet: true}
}

func (v NullableTemporaryTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemporaryTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
