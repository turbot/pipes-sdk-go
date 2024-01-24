# UpdateTenantSamlLoginSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** | The updated IdP public x509 certificate in PEM format. | [optional] 
**Issuer** | Pointer to **string** | The updated IdP issuer. | [optional] 
**SsoUrl** | Pointer to **string** | The updated IdP SSO URL. | [optional] 
**State** | **string** | The updated state of the login method. | 

## Methods

### NewUpdateTenantSamlLoginSettings

`func NewUpdateTenantSamlLoginSettings(state string, ) *UpdateTenantSamlLoginSettings`

NewUpdateTenantSamlLoginSettings instantiates a new UpdateTenantSamlLoginSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTenantSamlLoginSettingsWithDefaults

`func NewUpdateTenantSamlLoginSettingsWithDefaults() *UpdateTenantSamlLoginSettings`

NewUpdateTenantSamlLoginSettingsWithDefaults instantiates a new UpdateTenantSamlLoginSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *UpdateTenantSamlLoginSettings) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *UpdateTenantSamlLoginSettings) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *UpdateTenantSamlLoginSettings) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *UpdateTenantSamlLoginSettings) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetIssuer

`func (o *UpdateTenantSamlLoginSettings) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *UpdateTenantSamlLoginSettings) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *UpdateTenantSamlLoginSettings) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *UpdateTenantSamlLoginSettings) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSsoUrl

`func (o *UpdateTenantSamlLoginSettings) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *UpdateTenantSamlLoginSettings) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *UpdateTenantSamlLoginSettings) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.

### HasSsoUrl

`func (o *UpdateTenantSamlLoginSettings) HasSsoUrl() bool`

HasSsoUrl returns a boolean if a field has been set.

### GetState

`func (o *UpdateTenantSamlLoginSettings) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateTenantSamlLoginSettings) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateTenantSamlLoginSettings) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


