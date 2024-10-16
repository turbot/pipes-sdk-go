# Notifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) | User information for the user who created this. | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**DeletedAt** | Pointer to **string** | The time the item was deleted in ISO 8601 UTC. | [optional] 
**DeletedBy** | Pointer to [**User**](User.md) | User information for the user that performed the deletion. | [optional] 
**DeletedById** | **string** | The ID of the user that performed the deletion. | 
**Id** | **string** | The unique identifier for the notifier. | 
**IdentityId** | Pointer to **string** | The unique identifier for an identity where the notifier has been created. | [optional] 
**Name** | **string** | The unique handle of the notifier, unique up and down the resource hierarchy. | 
**Notifies** | Pointer to **[]map[string]interface{}** | The configured notify targets for this notifier | [optional] 
**Precedence** | Pointer to [**NotifierPrecedence**](NotifierPrecedence.md) | The precedence of this notifier | [optional] 
**State** | Pointer to [**NotifierState**](NotifierState.md) | The state of the notifier | [optional] 
**StateReason** | Pointer to **string** | The reason for the state of the notifier | [optional] 
**TenantId** | **string** | The unique identifier for the tenant where the notifier has been created. | 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) | User information for the last user to update this. | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 
**WorkspaceId** | Pointer to **string** | The unique identifier for the workspace where the notifier has been created. | [optional] 

## Methods

### NewNotifier

`func NewNotifier(createdAt string, createdById string, deletedById string, id string, name string, tenantId string, updatedById string, versionId int32, ) *Notifier`

NewNotifier instantiates a new Notifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifierWithDefaults

`func NewNotifierWithDefaults() *Notifier`

NewNotifierWithDefaults instantiates a new Notifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Notifier) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Notifier) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Notifier) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Notifier) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Notifier) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Notifier) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Notifier) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *Notifier) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *Notifier) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *Notifier) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetDeletedAt

`func (o *Notifier) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *Notifier) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *Notifier) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *Notifier) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *Notifier) GetDeletedBy() User`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *Notifier) GetDeletedByOk() (*User, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *Notifier) SetDeletedBy(v User)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *Notifier) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedById

`func (o *Notifier) GetDeletedById() string`

GetDeletedById returns the DeletedById field if non-nil, zero value otherwise.

### GetDeletedByIdOk

`func (o *Notifier) GetDeletedByIdOk() (*string, bool)`

GetDeletedByIdOk returns a tuple with the DeletedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedById

`func (o *Notifier) SetDeletedById(v string)`

SetDeletedById sets DeletedById field to given value.


### GetId

`func (o *Notifier) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Notifier) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Notifier) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityId

`func (o *Notifier) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *Notifier) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *Notifier) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *Notifier) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetName

`func (o *Notifier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Notifier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Notifier) SetName(v string)`

SetName sets Name field to given value.


### GetNotifies

`func (o *Notifier) GetNotifies() []map[string]interface{}`

GetNotifies returns the Notifies field if non-nil, zero value otherwise.

### GetNotifiesOk

`func (o *Notifier) GetNotifiesOk() (*[]map[string]interface{}, bool)`

GetNotifiesOk returns a tuple with the Notifies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifies

`func (o *Notifier) SetNotifies(v []map[string]interface{})`

SetNotifies sets Notifies field to given value.

### HasNotifies

`func (o *Notifier) HasNotifies() bool`

HasNotifies returns a boolean if a field has been set.

### GetPrecedence

`func (o *Notifier) GetPrecedence() NotifierPrecedence`

GetPrecedence returns the Precedence field if non-nil, zero value otherwise.

### GetPrecedenceOk

`func (o *Notifier) GetPrecedenceOk() (*NotifierPrecedence, bool)`

GetPrecedenceOk returns a tuple with the Precedence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrecedence

`func (o *Notifier) SetPrecedence(v NotifierPrecedence)`

SetPrecedence sets Precedence field to given value.

### HasPrecedence

`func (o *Notifier) HasPrecedence() bool`

HasPrecedence returns a boolean if a field has been set.

### GetState

`func (o *Notifier) GetState() NotifierState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Notifier) GetStateOk() (*NotifierState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Notifier) SetState(v NotifierState)`

SetState sets State field to given value.

### HasState

`func (o *Notifier) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateReason

`func (o *Notifier) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *Notifier) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *Notifier) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *Notifier) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### GetTenantId

`func (o *Notifier) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Notifier) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Notifier) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUpdatedAt

`func (o *Notifier) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Notifier) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Notifier) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Notifier) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *Notifier) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *Notifier) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *Notifier) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *Notifier) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *Notifier) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *Notifier) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *Notifier) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetVersionId

`func (o *Notifier) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *Notifier) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *Notifier) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.


### GetWorkspaceId

`func (o *Notifier) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *Notifier) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *Notifier) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *Notifier) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


