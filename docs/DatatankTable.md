# DatatankTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**Datatank** | Pointer to [**Datatank**](Datatank.md) |  | [optional] 
**DatatankId** | **string** |  | 
**DeletedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**DeletedBy** | Pointer to [**User**](User.md) |  | [optional] 
**DeletedById** | **string** | The ID of the user that performed the deletion. | 
**Description** | Pointer to **string** |  | [optional] 
**DesiredState** | **string** |  | 
**Frequency** | [**PipelineFrequency**](PipelineFrequency.md) |  | 
**Freshness** | Pointer to [**DatatankTableFreshness**](DatatankTableFreshness.md) |  | [optional] 
**Id** | **string** |  | 
**MigratingFreshness** | Pointer to [**DatatankTableFreshness**](DatatankTableFreshness.md) |  | [optional] 
**MigratingName** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**PartPer** | Pointer to **string** |  | [optional] 
**SourceQuery** | Pointer to **string** |  | [optional] 
**SourceSchema** | Pointer to **string** |  | [optional] 
**SourceTable** | Pointer to **string** |  | [optional] 
**State** | **string** |  | 
**StateReason** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 

## Methods

### NewDatatankTable

`func NewDatatankTable(createdAt string, createdById string, datatankId string, deletedById string, desiredState string, frequency PipelineFrequency, id string, name string, state string, type_ string, updatedById string, versionId int32, ) *DatatankTable`

NewDatatankTable instantiates a new DatatankTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatatankTableWithDefaults

`func NewDatatankTableWithDefaults() *DatatankTable`

NewDatatankTableWithDefaults instantiates a new DatatankTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DatatankTable) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DatatankTable) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DatatankTable) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *DatatankTable) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DatatankTable) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DatatankTable) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DatatankTable) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *DatatankTable) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *DatatankTable) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *DatatankTable) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetDatatank

`func (o *DatatankTable) GetDatatank() Datatank`

GetDatatank returns the Datatank field if non-nil, zero value otherwise.

### GetDatatankOk

`func (o *DatatankTable) GetDatatankOk() (*Datatank, bool)`

GetDatatankOk returns a tuple with the Datatank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatatank

`func (o *DatatankTable) SetDatatank(v Datatank)`

SetDatatank sets Datatank field to given value.

### HasDatatank

`func (o *DatatankTable) HasDatatank() bool`

HasDatatank returns a boolean if a field has been set.

### GetDatatankId

`func (o *DatatankTable) GetDatatankId() string`

GetDatatankId returns the DatatankId field if non-nil, zero value otherwise.

### GetDatatankIdOk

`func (o *DatatankTable) GetDatatankIdOk() (*string, bool)`

GetDatatankIdOk returns a tuple with the DatatankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatatankId

`func (o *DatatankTable) SetDatatankId(v string)`

SetDatatankId sets DatatankId field to given value.


### GetDeletedAt

`func (o *DatatankTable) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DatatankTable) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DatatankTable) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DatatankTable) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *DatatankTable) GetDeletedBy() User`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *DatatankTable) GetDeletedByOk() (*User, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *DatatankTable) SetDeletedBy(v User)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *DatatankTable) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedById

`func (o *DatatankTable) GetDeletedById() string`

GetDeletedById returns the DeletedById field if non-nil, zero value otherwise.

### GetDeletedByIdOk

`func (o *DatatankTable) GetDeletedByIdOk() (*string, bool)`

GetDeletedByIdOk returns a tuple with the DeletedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedById

`func (o *DatatankTable) SetDeletedById(v string)`

SetDeletedById sets DeletedById field to given value.


### GetDescription

`func (o *DatatankTable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DatatankTable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DatatankTable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DatatankTable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDesiredState

`func (o *DatatankTable) GetDesiredState() string`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *DatatankTable) GetDesiredStateOk() (*string, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *DatatankTable) SetDesiredState(v string)`

SetDesiredState sets DesiredState field to given value.


### GetFrequency

`func (o *DatatankTable) GetFrequency() PipelineFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *DatatankTable) GetFrequencyOk() (*PipelineFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *DatatankTable) SetFrequency(v PipelineFrequency)`

SetFrequency sets Frequency field to given value.


### GetFreshness

`func (o *DatatankTable) GetFreshness() DatatankTableFreshness`

GetFreshness returns the Freshness field if non-nil, zero value otherwise.

### GetFreshnessOk

`func (o *DatatankTable) GetFreshnessOk() (*DatatankTableFreshness, bool)`

GetFreshnessOk returns a tuple with the Freshness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreshness

`func (o *DatatankTable) SetFreshness(v DatatankTableFreshness)`

SetFreshness sets Freshness field to given value.

### HasFreshness

`func (o *DatatankTable) HasFreshness() bool`

HasFreshness returns a boolean if a field has been set.

### GetId

`func (o *DatatankTable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatatankTable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatatankTable) SetId(v string)`

SetId sets Id field to given value.


### GetMigratingFreshness

`func (o *DatatankTable) GetMigratingFreshness() DatatankTableFreshness`

GetMigratingFreshness returns the MigratingFreshness field if non-nil, zero value otherwise.

### GetMigratingFreshnessOk

`func (o *DatatankTable) GetMigratingFreshnessOk() (*DatatankTableFreshness, bool)`

GetMigratingFreshnessOk returns a tuple with the MigratingFreshness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratingFreshness

`func (o *DatatankTable) SetMigratingFreshness(v DatatankTableFreshness)`

SetMigratingFreshness sets MigratingFreshness field to given value.

### HasMigratingFreshness

`func (o *DatatankTable) HasMigratingFreshness() bool`

HasMigratingFreshness returns a boolean if a field has been set.

### GetMigratingName

`func (o *DatatankTable) GetMigratingName() string`

GetMigratingName returns the MigratingName field if non-nil, zero value otherwise.

### GetMigratingNameOk

`func (o *DatatankTable) GetMigratingNameOk() (*string, bool)`

GetMigratingNameOk returns a tuple with the MigratingName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratingName

`func (o *DatatankTable) SetMigratingName(v string)`

SetMigratingName sets MigratingName field to given value.

### HasMigratingName

`func (o *DatatankTable) HasMigratingName() bool`

HasMigratingName returns a boolean if a field has been set.

### GetName

`func (o *DatatankTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatatankTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatatankTable) SetName(v string)`

SetName sets Name field to given value.


### GetPartPer

`func (o *DatatankTable) GetPartPer() string`

GetPartPer returns the PartPer field if non-nil, zero value otherwise.

### GetPartPerOk

`func (o *DatatankTable) GetPartPerOk() (*string, bool)`

GetPartPerOk returns a tuple with the PartPer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartPer

`func (o *DatatankTable) SetPartPer(v string)`

SetPartPer sets PartPer field to given value.

### HasPartPer

`func (o *DatatankTable) HasPartPer() bool`

HasPartPer returns a boolean if a field has been set.

### GetSourceQuery

`func (o *DatatankTable) GetSourceQuery() string`

GetSourceQuery returns the SourceQuery field if non-nil, zero value otherwise.

### GetSourceQueryOk

`func (o *DatatankTable) GetSourceQueryOk() (*string, bool)`

GetSourceQueryOk returns a tuple with the SourceQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceQuery

`func (o *DatatankTable) SetSourceQuery(v string)`

SetSourceQuery sets SourceQuery field to given value.

### HasSourceQuery

`func (o *DatatankTable) HasSourceQuery() bool`

HasSourceQuery returns a boolean if a field has been set.

### GetSourceSchema

`func (o *DatatankTable) GetSourceSchema() string`

GetSourceSchema returns the SourceSchema field if non-nil, zero value otherwise.

### GetSourceSchemaOk

`func (o *DatatankTable) GetSourceSchemaOk() (*string, bool)`

GetSourceSchemaOk returns a tuple with the SourceSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSchema

`func (o *DatatankTable) SetSourceSchema(v string)`

SetSourceSchema sets SourceSchema field to given value.

### HasSourceSchema

`func (o *DatatankTable) HasSourceSchema() bool`

HasSourceSchema returns a boolean if a field has been set.

### GetSourceTable

`func (o *DatatankTable) GetSourceTable() string`

GetSourceTable returns the SourceTable field if non-nil, zero value otherwise.

### GetSourceTableOk

`func (o *DatatankTable) GetSourceTableOk() (*string, bool)`

GetSourceTableOk returns a tuple with the SourceTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTable

`func (o *DatatankTable) SetSourceTable(v string)`

SetSourceTable sets SourceTable field to given value.

### HasSourceTable

`func (o *DatatankTable) HasSourceTable() bool`

HasSourceTable returns a boolean if a field has been set.

### GetState

`func (o *DatatankTable) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DatatankTable) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DatatankTable) SetState(v string)`

SetState sets State field to given value.


### GetStateReason

`func (o *DatatankTable) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *DatatankTable) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *DatatankTable) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *DatatankTable) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### GetType

`func (o *DatatankTable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatatankTable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatatankTable) SetType(v string)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *DatatankTable) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DatatankTable) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DatatankTable) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DatatankTable) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DatatankTable) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DatatankTable) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DatatankTable) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DatatankTable) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *DatatankTable) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *DatatankTable) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *DatatankTable) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetVersionId

`func (o *DatatankTable) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *DatatankTable) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *DatatankTable) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


