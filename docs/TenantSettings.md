# TenantSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) | User information for the user who created this. | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**LoginEmail** | [**TenantLoginSettings**](TenantLoginSettings.md) | Settings related to login via Email. | 
**LoginGithub** | [**TenantLoginSettings**](TenantLoginSettings.md) | Settings related to login via Github. | 
**LoginGoogle** | [**TenantLoginSettings**](TenantLoginSettings.md) | Settings related to login via Google. | 
**LoginSaml** | [**TenantSamlLoginSettings**](TenantSamlLoginSettings.md) | Settings related to login via Okta. | 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) | User information for the last user to update this. | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**UserProvisioning** | **[]string** | The user provisioning settings for the tenant. | 
**UserProvisioningPermittedDomains** | **[]string** | The domains that new users are permitted to be provisioned from. | 
**VersionId** | **int32** | The current version of the user setting. | 
**WorkspaceSnapshotPermittedVisibility** | **[]string** | List of allowed visibility settings for snapshots. | 

## Methods

### NewTenantSettings

`func NewTenantSettings(createdAt string, createdById string, loginEmail TenantLoginSettings, loginGithub TenantLoginSettings, loginGoogle TenantLoginSettings, loginSaml TenantSamlLoginSettings, updatedById string, userProvisioning []string, userProvisioningPermittedDomains []string, versionId int32, workspaceSnapshotPermittedVisibility []string, ) *TenantSettings`

NewTenantSettings instantiates a new TenantSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantSettingsWithDefaults

`func NewTenantSettingsWithDefaults() *TenantSettings`

NewTenantSettingsWithDefaults instantiates a new TenantSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *TenantSettings) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TenantSettings) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TenantSettings) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *TenantSettings) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *TenantSettings) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *TenantSettings) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *TenantSettings) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *TenantSettings) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *TenantSettings) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *TenantSettings) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetLoginEmail

`func (o *TenantSettings) GetLoginEmail() TenantLoginSettings`

GetLoginEmail returns the LoginEmail field if non-nil, zero value otherwise.

### GetLoginEmailOk

`func (o *TenantSettings) GetLoginEmailOk() (*TenantLoginSettings, bool)`

GetLoginEmailOk returns a tuple with the LoginEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginEmail

`func (o *TenantSettings) SetLoginEmail(v TenantLoginSettings)`

SetLoginEmail sets LoginEmail field to given value.


### GetLoginGithub

`func (o *TenantSettings) GetLoginGithub() TenantLoginSettings`

GetLoginGithub returns the LoginGithub field if non-nil, zero value otherwise.

### GetLoginGithubOk

`func (o *TenantSettings) GetLoginGithubOk() (*TenantLoginSettings, bool)`

GetLoginGithubOk returns a tuple with the LoginGithub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginGithub

`func (o *TenantSettings) SetLoginGithub(v TenantLoginSettings)`

SetLoginGithub sets LoginGithub field to given value.


### GetLoginGoogle

`func (o *TenantSettings) GetLoginGoogle() TenantLoginSettings`

GetLoginGoogle returns the LoginGoogle field if non-nil, zero value otherwise.

### GetLoginGoogleOk

`func (o *TenantSettings) GetLoginGoogleOk() (*TenantLoginSettings, bool)`

GetLoginGoogleOk returns a tuple with the LoginGoogle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginGoogle

`func (o *TenantSettings) SetLoginGoogle(v TenantLoginSettings)`

SetLoginGoogle sets LoginGoogle field to given value.


### GetLoginSaml

`func (o *TenantSettings) GetLoginSaml() TenantSamlLoginSettings`

GetLoginSaml returns the LoginSaml field if non-nil, zero value otherwise.

### GetLoginSamlOk

`func (o *TenantSettings) GetLoginSamlOk() (*TenantSamlLoginSettings, bool)`

GetLoginSamlOk returns a tuple with the LoginSaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginSaml

`func (o *TenantSettings) SetLoginSaml(v TenantSamlLoginSettings)`

SetLoginSaml sets LoginSaml field to given value.


### GetUpdatedAt

`func (o *TenantSettings) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TenantSettings) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TenantSettings) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TenantSettings) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *TenantSettings) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *TenantSettings) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *TenantSettings) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *TenantSettings) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *TenantSettings) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *TenantSettings) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *TenantSettings) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetUserProvisioning

`func (o *TenantSettings) GetUserProvisioning() []string`

GetUserProvisioning returns the UserProvisioning field if non-nil, zero value otherwise.

### GetUserProvisioningOk

`func (o *TenantSettings) GetUserProvisioningOk() (*[]string, bool)`

GetUserProvisioningOk returns a tuple with the UserProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProvisioning

`func (o *TenantSettings) SetUserProvisioning(v []string)`

SetUserProvisioning sets UserProvisioning field to given value.


### GetUserProvisioningPermittedDomains

`func (o *TenantSettings) GetUserProvisioningPermittedDomains() []string`

GetUserProvisioningPermittedDomains returns the UserProvisioningPermittedDomains field if non-nil, zero value otherwise.

### GetUserProvisioningPermittedDomainsOk

`func (o *TenantSettings) GetUserProvisioningPermittedDomainsOk() (*[]string, bool)`

GetUserProvisioningPermittedDomainsOk returns a tuple with the UserProvisioningPermittedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProvisioningPermittedDomains

`func (o *TenantSettings) SetUserProvisioningPermittedDomains(v []string)`

SetUserProvisioningPermittedDomains sets UserProvisioningPermittedDomains field to given value.


### GetVersionId

`func (o *TenantSettings) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *TenantSettings) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *TenantSettings) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.


### GetWorkspaceSnapshotPermittedVisibility

`func (o *TenantSettings) GetWorkspaceSnapshotPermittedVisibility() []string`

GetWorkspaceSnapshotPermittedVisibility returns the WorkspaceSnapshotPermittedVisibility field if non-nil, zero value otherwise.

### GetWorkspaceSnapshotPermittedVisibilityOk

`func (o *TenantSettings) GetWorkspaceSnapshotPermittedVisibilityOk() (*[]string, bool)`

GetWorkspaceSnapshotPermittedVisibilityOk returns a tuple with the WorkspaceSnapshotPermittedVisibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceSnapshotPermittedVisibility

`func (o *TenantSettings) SetWorkspaceSnapshotPermittedVisibility(v []string)`

SetWorkspaceSnapshotPermittedVisibility sets WorkspaceSnapshotPermittedVisibility field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


