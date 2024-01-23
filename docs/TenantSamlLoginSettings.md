# TenantSamlLoginSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** | The IdP public x509 certificate in PEM format. | [optional] 
**Issuer** | Pointer to **string** | The IdP issuer. | [optional] 
**SsoUrl** | Pointer to **string** | The IdP SSO URL. | [optional] 
**State** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewTenantSamlLoginSettings

`func NewTenantSamlLoginSettings(state string, type_ string, ) *TenantSamlLoginSettings`

NewTenantSamlLoginSettings instantiates a new TenantSamlLoginSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantSamlLoginSettingsWithDefaults

`func NewTenantSamlLoginSettingsWithDefaults() *TenantSamlLoginSettings`

NewTenantSamlLoginSettingsWithDefaults instantiates a new TenantSamlLoginSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *TenantSamlLoginSettings) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *TenantSamlLoginSettings) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *TenantSamlLoginSettings) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *TenantSamlLoginSettings) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetIssuer

`func (o *TenantSamlLoginSettings) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *TenantSamlLoginSettings) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *TenantSamlLoginSettings) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *TenantSamlLoginSettings) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSsoUrl

`func (o *TenantSamlLoginSettings) GetSsoUrl() string`

GetSsoUrl returns the SsoUrl field if non-nil, zero value otherwise.

### GetSsoUrlOk

`func (o *TenantSamlLoginSettings) GetSsoUrlOk() (*string, bool)`

GetSsoUrlOk returns a tuple with the SsoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoUrl

`func (o *TenantSamlLoginSettings) SetSsoUrl(v string)`

SetSsoUrl sets SsoUrl field to given value.

### HasSsoUrl

`func (o *TenantSamlLoginSettings) HasSsoUrl() bool`

HasSsoUrl returns a boolean if a field has been set.

### GetState

`func (o *TenantSamlLoginSettings) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TenantSamlLoginSettings) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TenantSamlLoginSettings) SetState(v string)`

SetState sets State field to given value.


### GetType

`func (o *TenantSamlLoginSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TenantSamlLoginSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TenantSamlLoginSettings) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


