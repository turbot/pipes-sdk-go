# WorkspaceAggregatorAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregator** | Pointer to [**Aggregator**](Aggregator.md) |  | [optional] 
**AggregatorId** | **string** | The unique identifier for the aggregator. | 
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**Id** | **string** | The unique identifier for the workspace connection association. | 
**IdentityId** | **string** | The identity ID where the association exists. | 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 
**WorkspaceId** | **string** | The unique identifier for the wokspace. | 

## Methods

### NewWorkspaceAggregatorAssociation

`func NewWorkspaceAggregatorAssociation(aggregatorId string, createdAt string, createdById string, id string, identityId string, updatedById string, versionId int32, workspaceId string, ) *WorkspaceAggregatorAssociation`

NewWorkspaceAggregatorAssociation instantiates a new WorkspaceAggregatorAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceAggregatorAssociationWithDefaults

`func NewWorkspaceAggregatorAssociationWithDefaults() *WorkspaceAggregatorAssociation`

NewWorkspaceAggregatorAssociationWithDefaults instantiates a new WorkspaceAggregatorAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregator

`func (o *WorkspaceAggregatorAssociation) GetAggregator() Aggregator`

GetAggregator returns the Aggregator field if non-nil, zero value otherwise.

### GetAggregatorOk

`func (o *WorkspaceAggregatorAssociation) GetAggregatorOk() (*Aggregator, bool)`

GetAggregatorOk returns a tuple with the Aggregator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregator

`func (o *WorkspaceAggregatorAssociation) SetAggregator(v Aggregator)`

SetAggregator sets Aggregator field to given value.

### HasAggregator

`func (o *WorkspaceAggregatorAssociation) HasAggregator() bool`

HasAggregator returns a boolean if a field has been set.

### GetAggregatorId

`func (o *WorkspaceAggregatorAssociation) GetAggregatorId() string`

GetAggregatorId returns the AggregatorId field if non-nil, zero value otherwise.

### GetAggregatorIdOk

`func (o *WorkspaceAggregatorAssociation) GetAggregatorIdOk() (*string, bool)`

GetAggregatorIdOk returns a tuple with the AggregatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregatorId

`func (o *WorkspaceAggregatorAssociation) SetAggregatorId(v string)`

SetAggregatorId sets AggregatorId field to given value.


### GetCreatedAt

`func (o *WorkspaceAggregatorAssociation) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkspaceAggregatorAssociation) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkspaceAggregatorAssociation) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *WorkspaceAggregatorAssociation) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *WorkspaceAggregatorAssociation) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *WorkspaceAggregatorAssociation) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *WorkspaceAggregatorAssociation) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *WorkspaceAggregatorAssociation) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *WorkspaceAggregatorAssociation) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *WorkspaceAggregatorAssociation) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetId

`func (o *WorkspaceAggregatorAssociation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkspaceAggregatorAssociation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkspaceAggregatorAssociation) SetId(v string)`

SetId sets Id field to given value.


### GetIdentityId

`func (o *WorkspaceAggregatorAssociation) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *WorkspaceAggregatorAssociation) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *WorkspaceAggregatorAssociation) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.


### GetUpdatedAt

`func (o *WorkspaceAggregatorAssociation) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkspaceAggregatorAssociation) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkspaceAggregatorAssociation) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *WorkspaceAggregatorAssociation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *WorkspaceAggregatorAssociation) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *WorkspaceAggregatorAssociation) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *WorkspaceAggregatorAssociation) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *WorkspaceAggregatorAssociation) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *WorkspaceAggregatorAssociation) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *WorkspaceAggregatorAssociation) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *WorkspaceAggregatorAssociation) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetVersionId

`func (o *WorkspaceAggregatorAssociation) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *WorkspaceAggregatorAssociation) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *WorkspaceAggregatorAssociation) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.


### GetWorkspaceId

`func (o *WorkspaceAggregatorAssociation) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *WorkspaceAggregatorAssociation) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *WorkspaceAggregatorAssociation) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


