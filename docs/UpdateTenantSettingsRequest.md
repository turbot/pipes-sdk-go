# UpdateTenantSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CliSessionTimeout** | Pointer to **int32** | The CLI session timeout in seconds. | [optional] 
**ConsoleSessionTimeout** | Pointer to **int32** | The console (browser) session timeout in seconds. | [optional] 
**LoginEmail** | Pointer to [**UpdateTenantLoginSettings**](UpdateTenantLoginSettings.md) | Settings related to login via Email. | [optional] 
**LoginGithub** | Pointer to [**UpdateTenantLoginSettings**](UpdateTenantLoginSettings.md) | Settings related to login via Github. | [optional] 
**LoginGoogle** | Pointer to [**UpdateTenantLoginSettings**](UpdateTenantLoginSettings.md) | Settings related to login via Google. | [optional] 
**LoginSaml** | Pointer to [**UpdateTenantSamlLoginSettings**](UpdateTenantSamlLoginSettings.md) | LoginMicrosoft      *TenantLoginSettings     &#x60;json:\&quot;login_microsoft\&quot; binding:\&quot;omitempty\&quot;&#x60;                                  // Settings related to login via Microsoft. | [optional] 
**MaxTokenExpiration** | Pointer to **int32** | The maximum token expiration in seconds. (Default: 0 indicates we allow non-expiring tokens) | [optional] 
**PersonalWorkspaces** | Pointer to [**TenantPersonalWorkspaces**](TenantPersonalWorkspaces.md) | Settings related to user personal workspaces. | [optional] 
**PostgresEndpoint** | Pointer to [**PostgresEndpointState**](PostgresEndpointState.md) | Settings related to user personal workspaces. | [optional] 
**UserProvisioning** | Pointer to **[]string** | The user provisioning settings for the tenant. | [optional] 
**UserProvisioningPermittedDomains** | Pointer to **[]string** | The domains that new users are permitted to be provisioned from. | [optional] 
**WorkspaceSnapshotPermittedVisibility** | Pointer to **[]string** | Settings related to visibility of snapshots in the tenant. | [optional] 

## Methods

### NewUpdateTenantSettingsRequest

`func NewUpdateTenantSettingsRequest() *UpdateTenantSettingsRequest`

NewUpdateTenantSettingsRequest instantiates a new UpdateTenantSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTenantSettingsRequestWithDefaults

`func NewUpdateTenantSettingsRequestWithDefaults() *UpdateTenantSettingsRequest`

NewUpdateTenantSettingsRequestWithDefaults instantiates a new UpdateTenantSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCliSessionTimeout

`func (o *UpdateTenantSettingsRequest) GetCliSessionTimeout() int32`

GetCliSessionTimeout returns the CliSessionTimeout field if non-nil, zero value otherwise.

### GetCliSessionTimeoutOk

`func (o *UpdateTenantSettingsRequest) GetCliSessionTimeoutOk() (*int32, bool)`

GetCliSessionTimeoutOk returns a tuple with the CliSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliSessionTimeout

`func (o *UpdateTenantSettingsRequest) SetCliSessionTimeout(v int32)`

SetCliSessionTimeout sets CliSessionTimeout field to given value.

### HasCliSessionTimeout

`func (o *UpdateTenantSettingsRequest) HasCliSessionTimeout() bool`

HasCliSessionTimeout returns a boolean if a field has been set.

### GetConsoleSessionTimeout

`func (o *UpdateTenantSettingsRequest) GetConsoleSessionTimeout() int32`

GetConsoleSessionTimeout returns the ConsoleSessionTimeout field if non-nil, zero value otherwise.

### GetConsoleSessionTimeoutOk

`func (o *UpdateTenantSettingsRequest) GetConsoleSessionTimeoutOk() (*int32, bool)`

GetConsoleSessionTimeoutOk returns a tuple with the ConsoleSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsoleSessionTimeout

`func (o *UpdateTenantSettingsRequest) SetConsoleSessionTimeout(v int32)`

SetConsoleSessionTimeout sets ConsoleSessionTimeout field to given value.

### HasConsoleSessionTimeout

`func (o *UpdateTenantSettingsRequest) HasConsoleSessionTimeout() bool`

HasConsoleSessionTimeout returns a boolean if a field has been set.

### GetLoginEmail

`func (o *UpdateTenantSettingsRequest) GetLoginEmail() UpdateTenantLoginSettings`

GetLoginEmail returns the LoginEmail field if non-nil, zero value otherwise.

### GetLoginEmailOk

`func (o *UpdateTenantSettingsRequest) GetLoginEmailOk() (*UpdateTenantLoginSettings, bool)`

GetLoginEmailOk returns a tuple with the LoginEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginEmail

`func (o *UpdateTenantSettingsRequest) SetLoginEmail(v UpdateTenantLoginSettings)`

SetLoginEmail sets LoginEmail field to given value.

### HasLoginEmail

`func (o *UpdateTenantSettingsRequest) HasLoginEmail() bool`

HasLoginEmail returns a boolean if a field has been set.

### GetLoginGithub

`func (o *UpdateTenantSettingsRequest) GetLoginGithub() UpdateTenantLoginSettings`

GetLoginGithub returns the LoginGithub field if non-nil, zero value otherwise.

### GetLoginGithubOk

`func (o *UpdateTenantSettingsRequest) GetLoginGithubOk() (*UpdateTenantLoginSettings, bool)`

GetLoginGithubOk returns a tuple with the LoginGithub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginGithub

`func (o *UpdateTenantSettingsRequest) SetLoginGithub(v UpdateTenantLoginSettings)`

SetLoginGithub sets LoginGithub field to given value.

### HasLoginGithub

`func (o *UpdateTenantSettingsRequest) HasLoginGithub() bool`

HasLoginGithub returns a boolean if a field has been set.

### GetLoginGoogle

`func (o *UpdateTenantSettingsRequest) GetLoginGoogle() UpdateTenantLoginSettings`

GetLoginGoogle returns the LoginGoogle field if non-nil, zero value otherwise.

### GetLoginGoogleOk

`func (o *UpdateTenantSettingsRequest) GetLoginGoogleOk() (*UpdateTenantLoginSettings, bool)`

GetLoginGoogleOk returns a tuple with the LoginGoogle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginGoogle

`func (o *UpdateTenantSettingsRequest) SetLoginGoogle(v UpdateTenantLoginSettings)`

SetLoginGoogle sets LoginGoogle field to given value.

### HasLoginGoogle

`func (o *UpdateTenantSettingsRequest) HasLoginGoogle() bool`

HasLoginGoogle returns a boolean if a field has been set.

### GetLoginSaml

`func (o *UpdateTenantSettingsRequest) GetLoginSaml() UpdateTenantSamlLoginSettings`

GetLoginSaml returns the LoginSaml field if non-nil, zero value otherwise.

### GetLoginSamlOk

`func (o *UpdateTenantSettingsRequest) GetLoginSamlOk() (*UpdateTenantSamlLoginSettings, bool)`

GetLoginSamlOk returns a tuple with the LoginSaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginSaml

`func (o *UpdateTenantSettingsRequest) SetLoginSaml(v UpdateTenantSamlLoginSettings)`

SetLoginSaml sets LoginSaml field to given value.

### HasLoginSaml

`func (o *UpdateTenantSettingsRequest) HasLoginSaml() bool`

HasLoginSaml returns a boolean if a field has been set.

### GetMaxTokenExpiration

`func (o *UpdateTenantSettingsRequest) GetMaxTokenExpiration() int32`

GetMaxTokenExpiration returns the MaxTokenExpiration field if non-nil, zero value otherwise.

### GetMaxTokenExpirationOk

`func (o *UpdateTenantSettingsRequest) GetMaxTokenExpirationOk() (*int32, bool)`

GetMaxTokenExpirationOk returns a tuple with the MaxTokenExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokenExpiration

`func (o *UpdateTenantSettingsRequest) SetMaxTokenExpiration(v int32)`

SetMaxTokenExpiration sets MaxTokenExpiration field to given value.

### HasMaxTokenExpiration

`func (o *UpdateTenantSettingsRequest) HasMaxTokenExpiration() bool`

HasMaxTokenExpiration returns a boolean if a field has been set.

### GetPersonalWorkspaces

`func (o *UpdateTenantSettingsRequest) GetPersonalWorkspaces() TenantPersonalWorkspaces`

GetPersonalWorkspaces returns the PersonalWorkspaces field if non-nil, zero value otherwise.

### GetPersonalWorkspacesOk

`func (o *UpdateTenantSettingsRequest) GetPersonalWorkspacesOk() (*TenantPersonalWorkspaces, bool)`

GetPersonalWorkspacesOk returns a tuple with the PersonalWorkspaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalWorkspaces

`func (o *UpdateTenantSettingsRequest) SetPersonalWorkspaces(v TenantPersonalWorkspaces)`

SetPersonalWorkspaces sets PersonalWorkspaces field to given value.

### HasPersonalWorkspaces

`func (o *UpdateTenantSettingsRequest) HasPersonalWorkspaces() bool`

HasPersonalWorkspaces returns a boolean if a field has been set.

### GetPostgresEndpoint

`func (o *UpdateTenantSettingsRequest) GetPostgresEndpoint() PostgresEndpointState`

GetPostgresEndpoint returns the PostgresEndpoint field if non-nil, zero value otherwise.

### GetPostgresEndpointOk

`func (o *UpdateTenantSettingsRequest) GetPostgresEndpointOk() (*PostgresEndpointState, bool)`

GetPostgresEndpointOk returns a tuple with the PostgresEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresEndpoint

`func (o *UpdateTenantSettingsRequest) SetPostgresEndpoint(v PostgresEndpointState)`

SetPostgresEndpoint sets PostgresEndpoint field to given value.

### HasPostgresEndpoint

`func (o *UpdateTenantSettingsRequest) HasPostgresEndpoint() bool`

HasPostgresEndpoint returns a boolean if a field has been set.

### GetUserProvisioning

`func (o *UpdateTenantSettingsRequest) GetUserProvisioning() []string`

GetUserProvisioning returns the UserProvisioning field if non-nil, zero value otherwise.

### GetUserProvisioningOk

`func (o *UpdateTenantSettingsRequest) GetUserProvisioningOk() (*[]string, bool)`

GetUserProvisioningOk returns a tuple with the UserProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProvisioning

`func (o *UpdateTenantSettingsRequest) SetUserProvisioning(v []string)`

SetUserProvisioning sets UserProvisioning field to given value.

### HasUserProvisioning

`func (o *UpdateTenantSettingsRequest) HasUserProvisioning() bool`

HasUserProvisioning returns a boolean if a field has been set.

### GetUserProvisioningPermittedDomains

`func (o *UpdateTenantSettingsRequest) GetUserProvisioningPermittedDomains() []string`

GetUserProvisioningPermittedDomains returns the UserProvisioningPermittedDomains field if non-nil, zero value otherwise.

### GetUserProvisioningPermittedDomainsOk

`func (o *UpdateTenantSettingsRequest) GetUserProvisioningPermittedDomainsOk() (*[]string, bool)`

GetUserProvisioningPermittedDomainsOk returns a tuple with the UserProvisioningPermittedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProvisioningPermittedDomains

`func (o *UpdateTenantSettingsRequest) SetUserProvisioningPermittedDomains(v []string)`

SetUserProvisioningPermittedDomains sets UserProvisioningPermittedDomains field to given value.

### HasUserProvisioningPermittedDomains

`func (o *UpdateTenantSettingsRequest) HasUserProvisioningPermittedDomains() bool`

HasUserProvisioningPermittedDomains returns a boolean if a field has been set.

### GetWorkspaceSnapshotPermittedVisibility

`func (o *UpdateTenantSettingsRequest) GetWorkspaceSnapshotPermittedVisibility() []string`

GetWorkspaceSnapshotPermittedVisibility returns the WorkspaceSnapshotPermittedVisibility field if non-nil, zero value otherwise.

### GetWorkspaceSnapshotPermittedVisibilityOk

`func (o *UpdateTenantSettingsRequest) GetWorkspaceSnapshotPermittedVisibilityOk() (*[]string, bool)`

GetWorkspaceSnapshotPermittedVisibilityOk returns a tuple with the WorkspaceSnapshotPermittedVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSnapshotPermittedVisibility

`func (o *UpdateTenantSettingsRequest) SetWorkspaceSnapshotPermittedVisibility(v []string)`

SetWorkspaceSnapshotPermittedVisibility sets WorkspaceSnapshotPermittedVisibility field to given value.

### HasWorkspaceSnapshotPermittedVisibility

`func (o *UpdateTenantSettingsRequest) HasWorkspaceSnapshotPermittedVisibility() bool`

HasWorkspaceSnapshotPermittedVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


