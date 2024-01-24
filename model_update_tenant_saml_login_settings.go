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

// UpdateTenantSamlLoginSettings struct for UpdateTenantSamlLoginSettings
type UpdateTenantSamlLoginSettings struct {
	// The updated IdP public x509 certificate in PEM format.
	Certificate *string `json:"certificate,omitempty"`
	// The updated IdP issuer.
	Issuer *string `json:"issuer,omitempty"`
	// The updated IdP SSO URL.
	SsoUrl *string `json:"sso_url,omitempty"`
	// The updated state of the login method.
	State string `json:"state"`
}

// NewUpdateTenantSamlLoginSettings instantiates a new UpdateTenantSamlLoginSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateTenantSamlLoginSettings(state string) *UpdateTenantSamlLoginSettings {
	this := UpdateTenantSamlLoginSettings{}
	this.State = state
	return &this
}

// NewUpdateTenantSamlLoginSettingsWithDefaults instantiates a new UpdateTenantSamlLoginSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateTenantSamlLoginSettingsWithDefaults() *UpdateTenantSamlLoginSettings {
	this := UpdateTenantSamlLoginSettings{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *UpdateTenantSamlLoginSettings) GetCertificate() string {
	if o == nil || o.Certificate == nil {
		var ret string
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenantSamlLoginSettings) GetCertificateOk() (*string, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *UpdateTenantSamlLoginSettings) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given string and assigns it to the Certificate field.
func (o *UpdateTenantSamlLoginSettings) SetCertificate(v string) {
	o.Certificate = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *UpdateTenantSamlLoginSettings) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenantSamlLoginSettings) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *UpdateTenantSamlLoginSettings) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *UpdateTenantSamlLoginSettings) SetIssuer(v string) {
	o.Issuer = &v
}

// GetSsoUrl returns the SsoUrl field value if set, zero value otherwise.
func (o *UpdateTenantSamlLoginSettings) GetSsoUrl() string {
	if o == nil || o.SsoUrl == nil {
		var ret string
		return ret
	}
	return *o.SsoUrl
}

// GetSsoUrlOk returns a tuple with the SsoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateTenantSamlLoginSettings) GetSsoUrlOk() (*string, bool) {
	if o == nil || o.SsoUrl == nil {
		return nil, false
	}
	return o.SsoUrl, true
}

// HasSsoUrl returns a boolean if a field has been set.
func (o *UpdateTenantSamlLoginSettings) HasSsoUrl() bool {
	if o != nil && o.SsoUrl != nil {
		return true
	}

	return false
}

// SetSsoUrl gets a reference to the given string and assigns it to the SsoUrl field.
func (o *UpdateTenantSamlLoginSettings) SetSsoUrl(v string) {
	o.SsoUrl = &v
}

// GetState returns the State field value
func (o *UpdateTenantSamlLoginSettings) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *UpdateTenantSamlLoginSettings) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *UpdateTenantSamlLoginSettings) SetState(v string) {
	o.State = v
}

func (o UpdateTenantSamlLoginSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.SsoUrl != nil {
		toSerialize["sso_url"] = o.SsoUrl
	}
	if true {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateTenantSamlLoginSettings struct {
	value *UpdateTenantSamlLoginSettings
	isSet bool
}

func (v NullableUpdateTenantSamlLoginSettings) Get() *UpdateTenantSamlLoginSettings {
	return v.value
}

func (v *NullableUpdateTenantSamlLoginSettings) Set(val *UpdateTenantSamlLoginSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateTenantSamlLoginSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateTenantSamlLoginSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateTenantSamlLoginSettings(val *UpdateTenantSamlLoginSettings) *NullableUpdateTenantSamlLoginSettings {
	return &NullableUpdateTenantSamlLoginSettings{value: val, isSet: true}
}

func (v NullableUpdateTenantSamlLoginSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateTenantSamlLoginSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
