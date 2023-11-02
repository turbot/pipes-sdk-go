# DatatankPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | The time of creation in ISO 8601 UTC. | 
**CreatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**CreatedById** | **string** | The ID of the user that created this. | 
**DatatankTable** | Pointer to [**DatatankTable**](DatatankTable.md) |  | [optional] 
**DatatankTableId** | **string** |  | 
**DeletedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**DeletedBy** | Pointer to [**User**](User.md) |  | [optional] 
**DeletedById** | **string** | The ID of the user that performed the deletion. | 
**DesiredState** | **string** |  | 
**Frequency** | [**PipelineFrequency**](PipelineFrequency.md) |  | 
**FreshnessState** | **string** |  | 
**Id** | **string** |  | 
**LastErrorAt** | Pointer to **string** |  | [optional] 
**LastErrorProcess** | Pointer to [**SpProcess**](SpProcess.md) |  | [optional] 
**LastErrorProcessId** | Pointer to **string** |  | [optional] 
**LastSuccessfulUpdateAt** | Pointer to **string** |  | [optional] 
**LastSuccessfulUpdateProcess** | Pointer to [**SpProcess**](SpProcess.md) |  | [optional] 
**LastSuccessfulUpdateProcessId** | Pointer to **string** |  | [optional] 
**LastUpdateAttemptAt** | Pointer to **string** |  | [optional] 
**LastUpdateAttemptProcess** | Pointer to [**SpProcess**](SpProcess.md) |  | [optional] 
**LastUpdateAttemptProcessId** | Pointer to **string** |  | [optional] 
**MigratingFreshnessState** | Pointer to **string** |  | [optional] 
**MigratingLastErrorAt** | Pointer to **string** |  | [optional] 
**MigratingLastErrorProcessId** | Pointer to **string** |  | [optional] 
**MigratingLastSuccessfulUpdateAt** | Pointer to **string** |  | [optional] 
**MigratingLastSuccessfulUpdateProcess** | Pointer to [**SpProcess**](SpProcess.md) |  | [optional] 
**MigratingLastSuccessfulUpdateProcessId** | Pointer to **string** |  | [optional] 
**MigratingLastUpdateAttemptAt** | Pointer to **string** |  | [optional] 
**MigratingLastUpdateAttemptProcess** | Pointer to [**SpProcess**](SpProcess.md) |  | [optional] 
**MigratingLastUpdateAttemptProcessId** | Pointer to **string** |  | [optional] 
**PartitionCandidate** | Pointer to **string** |  | [optional] 
**PartitionCurrent** | **string** |  | 
**Pipeline** | Pointer to [**Pipeline**](Pipeline.md) |  | [optional] 
**PipelineId** | **string** |  | 
**SourceConnection** | Pointer to [**WorkspaceSchema**](WorkspaceSchema.md) |  | [optional] 
**SourceConnectionId** | Pointer to **string** |  | [optional] 
**State** | **string** |  | 
**StateReason** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **string** | The time of the last update in ISO 8601 UTC. | [optional] 
**UpdatedBy** | Pointer to [**User**](User.md) |  | [optional] 
**UpdatedById** | **string** | The ID of the user that performed the last update. | 
**VersionId** | **int32** | The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item. | 

## Methods

### NewDatatankPart

`func NewDatatankPart(createdAt string, createdById string, datatankTableId string, deletedById string, desiredState string, frequency PipelineFrequency, freshnessState string, id string, partitionCurrent string, pipelineId string, state string, updatedById string, versionId int32, ) *DatatankPart`

NewDatatankPart instantiates a new DatatankPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatatankPartWithDefaults

`func NewDatatankPartWithDefaults() *DatatankPart`

NewDatatankPartWithDefaults instantiates a new DatatankPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *DatatankPart) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DatatankPart) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DatatankPart) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *DatatankPart) GetCreatedBy() User`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DatatankPart) GetCreatedByOk() (*User, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DatatankPart) SetCreatedBy(v User)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *DatatankPart) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedById

`func (o *DatatankPart) GetCreatedById() string`

GetCreatedById returns the CreatedById field if non-nil, zero value otherwise.

### GetCreatedByIdOk

`func (o *DatatankPart) GetCreatedByIdOk() (*string, bool)`

GetCreatedByIdOk returns a tuple with the CreatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedById

`func (o *DatatankPart) SetCreatedById(v string)`

SetCreatedById sets CreatedById field to given value.


### GetDatatankTable

`func (o *DatatankPart) GetDatatankTable() DatatankTable`

GetDatatankTable returns the DatatankTable field if non-nil, zero value otherwise.

### GetDatatankTableOk

`func (o *DatatankPart) GetDatatankTableOk() (*DatatankTable, bool)`

GetDatatankTableOk returns a tuple with the DatatankTable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatatankTable

`func (o *DatatankPart) SetDatatankTable(v DatatankTable)`

SetDatatankTable sets DatatankTable field to given value.

### HasDatatankTable

`func (o *DatatankPart) HasDatatankTable() bool`

HasDatatankTable returns a boolean if a field has been set.

### GetDatatankTableId

`func (o *DatatankPart) GetDatatankTableId() string`

GetDatatankTableId returns the DatatankTableId field if non-nil, zero value otherwise.

### GetDatatankTableIdOk

`func (o *DatatankPart) GetDatatankTableIdOk() (*string, bool)`

GetDatatankTableIdOk returns a tuple with the DatatankTableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatatankTableId

`func (o *DatatankPart) SetDatatankTableId(v string)`

SetDatatankTableId sets DatatankTableId field to given value.


### GetDeletedAt

`func (o *DatatankPart) GetDeletedAt() string`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *DatatankPart) GetDeletedAtOk() (*string, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *DatatankPart) SetDeletedAt(v string)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *DatatankPart) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### GetDeletedBy

`func (o *DatatankPart) GetDeletedBy() User`

GetDeletedBy returns the DeletedBy field if non-nil, zero value otherwise.

### GetDeletedByOk

`func (o *DatatankPart) GetDeletedByOk() (*User, bool)`

GetDeletedByOk returns a tuple with the DeletedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedBy

`func (o *DatatankPart) SetDeletedBy(v User)`

SetDeletedBy sets DeletedBy field to given value.

### HasDeletedBy

`func (o *DatatankPart) HasDeletedBy() bool`

HasDeletedBy returns a boolean if a field has been set.

### GetDeletedById

`func (o *DatatankPart) GetDeletedById() string`

GetDeletedById returns the DeletedById field if non-nil, zero value otherwise.

### GetDeletedByIdOk

`func (o *DatatankPart) GetDeletedByIdOk() (*string, bool)`

GetDeletedByIdOk returns a tuple with the DeletedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedById

`func (o *DatatankPart) SetDeletedById(v string)`

SetDeletedById sets DeletedById field to given value.


### GetDesiredState

`func (o *DatatankPart) GetDesiredState() string`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *DatatankPart) GetDesiredStateOk() (*string, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *DatatankPart) SetDesiredState(v string)`

SetDesiredState sets DesiredState field to given value.


### GetFrequency

`func (o *DatatankPart) GetFrequency() PipelineFrequency`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *DatatankPart) GetFrequencyOk() (*PipelineFrequency, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *DatatankPart) SetFrequency(v PipelineFrequency)`

SetFrequency sets Frequency field to given value.


### GetFreshnessState

`func (o *DatatankPart) GetFreshnessState() string`

GetFreshnessState returns the FreshnessState field if non-nil, zero value otherwise.

### GetFreshnessStateOk

`func (o *DatatankPart) GetFreshnessStateOk() (*string, bool)`

GetFreshnessStateOk returns a tuple with the FreshnessState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreshnessState

`func (o *DatatankPart) SetFreshnessState(v string)`

SetFreshnessState sets FreshnessState field to given value.


### GetId

`func (o *DatatankPart) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatatankPart) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatatankPart) SetId(v string)`

SetId sets Id field to given value.


### GetLastErrorAt

`func (o *DatatankPart) GetLastErrorAt() string`

GetLastErrorAt returns the LastErrorAt field if non-nil, zero value otherwise.

### GetLastErrorAtOk

`func (o *DatatankPart) GetLastErrorAtOk() (*string, bool)`

GetLastErrorAtOk returns a tuple with the LastErrorAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorAt

`func (o *DatatankPart) SetLastErrorAt(v string)`

SetLastErrorAt sets LastErrorAt field to given value.

### HasLastErrorAt

`func (o *DatatankPart) HasLastErrorAt() bool`

HasLastErrorAt returns a boolean if a field has been set.

### GetLastErrorProcess

`func (o *DatatankPart) GetLastErrorProcess() SpProcess`

GetLastErrorProcess returns the LastErrorProcess field if non-nil, zero value otherwise.

### GetLastErrorProcessOk

`func (o *DatatankPart) GetLastErrorProcessOk() (*SpProcess, bool)`

GetLastErrorProcessOk returns a tuple with the LastErrorProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorProcess

`func (o *DatatankPart) SetLastErrorProcess(v SpProcess)`

SetLastErrorProcess sets LastErrorProcess field to given value.

### HasLastErrorProcess

`func (o *DatatankPart) HasLastErrorProcess() bool`

HasLastErrorProcess returns a boolean if a field has been set.

### GetLastErrorProcessId

`func (o *DatatankPart) GetLastErrorProcessId() string`

GetLastErrorProcessId returns the LastErrorProcessId field if non-nil, zero value otherwise.

### GetLastErrorProcessIdOk

`func (o *DatatankPart) GetLastErrorProcessIdOk() (*string, bool)`

GetLastErrorProcessIdOk returns a tuple with the LastErrorProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorProcessId

`func (o *DatatankPart) SetLastErrorProcessId(v string)`

SetLastErrorProcessId sets LastErrorProcessId field to given value.

### HasLastErrorProcessId

`func (o *DatatankPart) HasLastErrorProcessId() bool`

HasLastErrorProcessId returns a boolean if a field has been set.

### GetLastSuccessfulUpdateAt

`func (o *DatatankPart) GetLastSuccessfulUpdateAt() string`

GetLastSuccessfulUpdateAt returns the LastSuccessfulUpdateAt field if non-nil, zero value otherwise.

### GetLastSuccessfulUpdateAtOk

`func (o *DatatankPart) GetLastSuccessfulUpdateAtOk() (*string, bool)`

GetLastSuccessfulUpdateAtOk returns a tuple with the LastSuccessfulUpdateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulUpdateAt

`func (o *DatatankPart) SetLastSuccessfulUpdateAt(v string)`

SetLastSuccessfulUpdateAt sets LastSuccessfulUpdateAt field to given value.

### HasLastSuccessfulUpdateAt

`func (o *DatatankPart) HasLastSuccessfulUpdateAt() bool`

HasLastSuccessfulUpdateAt returns a boolean if a field has been set.

### GetLastSuccessfulUpdateProcess

`func (o *DatatankPart) GetLastSuccessfulUpdateProcess() SpProcess`

GetLastSuccessfulUpdateProcess returns the LastSuccessfulUpdateProcess field if non-nil, zero value otherwise.

### GetLastSuccessfulUpdateProcessOk

`func (o *DatatankPart) GetLastSuccessfulUpdateProcessOk() (*SpProcess, bool)`

GetLastSuccessfulUpdateProcessOk returns a tuple with the LastSuccessfulUpdateProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulUpdateProcess

`func (o *DatatankPart) SetLastSuccessfulUpdateProcess(v SpProcess)`

SetLastSuccessfulUpdateProcess sets LastSuccessfulUpdateProcess field to given value.

### HasLastSuccessfulUpdateProcess

`func (o *DatatankPart) HasLastSuccessfulUpdateProcess() bool`

HasLastSuccessfulUpdateProcess returns a boolean if a field has been set.

### GetLastSuccessfulUpdateProcessId

`func (o *DatatankPart) GetLastSuccessfulUpdateProcessId() string`

GetLastSuccessfulUpdateProcessId returns the LastSuccessfulUpdateProcessId field if non-nil, zero value otherwise.

### GetLastSuccessfulUpdateProcessIdOk

`func (o *DatatankPart) GetLastSuccessfulUpdateProcessIdOk() (*string, bool)`

GetLastSuccessfulUpdateProcessIdOk returns a tuple with the LastSuccessfulUpdateProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulUpdateProcessId

`func (o *DatatankPart) SetLastSuccessfulUpdateProcessId(v string)`

SetLastSuccessfulUpdateProcessId sets LastSuccessfulUpdateProcessId field to given value.

### HasLastSuccessfulUpdateProcessId

`func (o *DatatankPart) HasLastSuccessfulUpdateProcessId() bool`

HasLastSuccessfulUpdateProcessId returns a boolean if a field has been set.

### GetLastUpdateAttemptAt

`func (o *DatatankPart) GetLastUpdateAttemptAt() string`

GetLastUpdateAttemptAt returns the LastUpdateAttemptAt field if non-nil, zero value otherwise.

### GetLastUpdateAttemptAtOk

`func (o *DatatankPart) GetLastUpdateAttemptAtOk() (*string, bool)`

GetLastUpdateAttemptAtOk returns a tuple with the LastUpdateAttemptAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateAttemptAt

`func (o *DatatankPart) SetLastUpdateAttemptAt(v string)`

SetLastUpdateAttemptAt sets LastUpdateAttemptAt field to given value.

### HasLastUpdateAttemptAt

`func (o *DatatankPart) HasLastUpdateAttemptAt() bool`

HasLastUpdateAttemptAt returns a boolean if a field has been set.

### GetLastUpdateAttemptProcess

`func (o *DatatankPart) GetLastUpdateAttemptProcess() SpProcess`

GetLastUpdateAttemptProcess returns the LastUpdateAttemptProcess field if non-nil, zero value otherwise.

### GetLastUpdateAttemptProcessOk

`func (o *DatatankPart) GetLastUpdateAttemptProcessOk() (*SpProcess, bool)`

GetLastUpdateAttemptProcessOk returns a tuple with the LastUpdateAttemptProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateAttemptProcess

`func (o *DatatankPart) SetLastUpdateAttemptProcess(v SpProcess)`

SetLastUpdateAttemptProcess sets LastUpdateAttemptProcess field to given value.

### HasLastUpdateAttemptProcess

`func (o *DatatankPart) HasLastUpdateAttemptProcess() bool`

HasLastUpdateAttemptProcess returns a boolean if a field has been set.

### GetLastUpdateAttemptProcessId

`func (o *DatatankPart) GetLastUpdateAttemptProcessId() string`

GetLastUpdateAttemptProcessId returns the LastUpdateAttemptProcessId field if non-nil, zero value otherwise.

### GetLastUpdateAttemptProcessIdOk

`func (o *DatatankPart) GetLastUpdateAttemptProcessIdOk() (*string, bool)`

GetLastUpdateAttemptProcessIdOk returns a tuple with the LastUpdateAttemptProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateAttemptProcessId

`func (o *DatatankPart) SetLastUpdateAttemptProcessId(v string)`

SetLastUpdateAttemptProcessId sets LastUpdateAttemptProcessId field to given value.

### HasLastUpdateAttemptProcessId

`func (o *DatatankPart) HasLastUpdateAttemptProcessId() bool`

HasLastUpdateAttemptProcessId returns a boolean if a field has been set.

### GetMigratingFreshnessState

`func (o *DatatankPart) GetMigratingFreshnessState() string`

GetMigratingFreshnessState returns the MigratingFreshnessState field if non-nil, zero value otherwise.

### GetMigratingFreshnessStateOk

`func (o *DatatankPart) GetMigratingFreshnessStateOk() (*string, bool)`

GetMigratingFreshnessStateOk returns a tuple with the MigratingFreshnessState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratingFreshnessState

`func (o *DatatankPart) SetMigratingFreshnessState(v string)`

SetMigratingFreshnessState sets MigratingFreshnessState field to given value.

### HasMigratingFreshnessState

`func (o *DatatankPart) HasMigratingFreshnessState() bool`

HasMigratingFreshnessState returns a boolean if a field has been set.

### GetMigratingLastErrorAt

`func (o *DatatankPart) GetMigratingLastErrorAt() string`

GetMigratingLastErrorAt returns the MigratingLastErrorAt field if non-nil, zero value otherwise.

### GetMigratingLastErrorAtOk

`func (o *DatatankPart) GetMigratingLastErrorAtOk() (*string, bool)`

GetMigratingLastErrorAtOk returns a tuple with the MigratingLastErrorAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratingLastErrorAt

`func (o *DatatankPart) SetMigratingLastErrorAt(v string)`

SetMigratingLastErrorAt sets MigratingLastErrorAt field to given value.

### HasMigratingLastErrorAt

`func (o *DatatankPart) HasMigratingLastErrorAt() bool`

HasMigratingLastErrorAt returns a boolean if a field has been set.

### GetMigratingLastErrorProcessId

`func (o *DatatankPart) GetMigratingLastErrorProcessId() string`

GetMigratingLastErrorProcessId returns the MigratingLastErrorProcessId field if non-nil, zero value otherwise.

### GetMigratingLastErrorProcessIdOk

`func (o *DatatankPart) GetMigratingLastErrorProcessIdOk() (*string, bool)`

GetMigratingLastErrorProcessIdOk returns a tuple with the MigratingLastErrorProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratingLastErrorProcessId

`func (o *DatatankPart) SetMigratingLastErrorProcessId(v string)`

SetMigratingLastErrorProcessId sets MigratingLastErrorProcessId field to given value.

### HasMigratingLastErrorProcessId

`func (o *DatatankPart) HasMigratingLastErrorProcessId() bool`

HasMigratingLastErrorProcessId returns a boolean if a field has been set.

### GetMigratingLastSuccessfulUpdateAt

`func (o *DatatankPart) GetMigratingLastSuccessfulUpdateAt() string`

GetMigratingLastSuccessfulUpdateAt returns the MigratingLastSuccessfulUpdateAt field if non-nil, zero value otherwise.

### GetMigratingLastSuccessfulUpdateAtOk

`func (o *DatatankPart) GetMigratingLastSuccessfulUpdateAtOk() (*string, bool)`

GetMigratingLastSuccessfulUpdateAtOk returns a tuple with the MigratingLastSuccessfulUpdateAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratingLastSuccessfulUpdateAt

`func (o *DatatankPart) SetMigratingLastSuccessfulUpdateAt(v string)`

SetMigratingLastSuccessfulUpdateAt sets MigratingLastSuccessfulUpdateAt field to given value.

### HasMigratingLastSuccessfulUpdateAt

`func (o *DatatankPart) HasMigratingLastSuccessfulUpdateAt() bool`

HasMigratingLastSuccessfulUpdateAt returns a boolean if a field has been set.

### GetMigratingLastSuccessfulUpdateProcess

`func (o *DatatankPart) GetMigratingLastSuccessfulUpdateProcess() SpProcess`

GetMigratingLastSuccessfulUpdateProcess returns the MigratingLastSuccessfulUpdateProcess field if non-nil, zero value otherwise.

### GetMigratingLastSuccessfulUpdateProcessOk

`func (o *DatatankPart) GetMigratingLastSuccessfulUpdateProcessOk() (*SpProcess, bool)`

GetMigratingLastSuccessfulUpdateProcessOk returns a tuple with the MigratingLastSuccessfulUpdateProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratingLastSuccessfulUpdateProcess

`func (o *DatatankPart) SetMigratingLastSuccessfulUpdateProcess(v SpProcess)`

SetMigratingLastSuccessfulUpdateProcess sets MigratingLastSuccessfulUpdateProcess field to given value.

### HasMigratingLastSuccessfulUpdateProcess

`func (o *DatatankPart) HasMigratingLastSuccessfulUpdateProcess() bool`

HasMigratingLastSuccessfulUpdateProcess returns a boolean if a field has been set.

### GetMigratingLastSuccessfulUpdateProcessId

`func (o *DatatankPart) GetMigratingLastSuccessfulUpdateProcessId() string`

GetMigratingLastSuccessfulUpdateProcessId returns the MigratingLastSuccessfulUpdateProcessId field if non-nil, zero value otherwise.

### GetMigratingLastSuccessfulUpdateProcessIdOk

`func (o *DatatankPart) GetMigratingLastSuccessfulUpdateProcessIdOk() (*string, bool)`

GetMigratingLastSuccessfulUpdateProcessIdOk returns a tuple with the MigratingLastSuccessfulUpdateProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratingLastSuccessfulUpdateProcessId

`func (o *DatatankPart) SetMigratingLastSuccessfulUpdateProcessId(v string)`

SetMigratingLastSuccessfulUpdateProcessId sets MigratingLastSuccessfulUpdateProcessId field to given value.

### HasMigratingLastSuccessfulUpdateProcessId

`func (o *DatatankPart) HasMigratingLastSuccessfulUpdateProcessId() bool`

HasMigratingLastSuccessfulUpdateProcessId returns a boolean if a field has been set.

### GetMigratingLastUpdateAttemptAt

`func (o *DatatankPart) GetMigratingLastUpdateAttemptAt() string`

GetMigratingLastUpdateAttemptAt returns the MigratingLastUpdateAttemptAt field if non-nil, zero value otherwise.

### GetMigratingLastUpdateAttemptAtOk

`func (o *DatatankPart) GetMigratingLastUpdateAttemptAtOk() (*string, bool)`

GetMigratingLastUpdateAttemptAtOk returns a tuple with the MigratingLastUpdateAttemptAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratingLastUpdateAttemptAt

`func (o *DatatankPart) SetMigratingLastUpdateAttemptAt(v string)`

SetMigratingLastUpdateAttemptAt sets MigratingLastUpdateAttemptAt field to given value.

### HasMigratingLastUpdateAttemptAt

`func (o *DatatankPart) HasMigratingLastUpdateAttemptAt() bool`

HasMigratingLastUpdateAttemptAt returns a boolean if a field has been set.

### GetMigratingLastUpdateAttemptProcess

`func (o *DatatankPart) GetMigratingLastUpdateAttemptProcess() SpProcess`

GetMigratingLastUpdateAttemptProcess returns the MigratingLastUpdateAttemptProcess field if non-nil, zero value otherwise.

### GetMigratingLastUpdateAttemptProcessOk

`func (o *DatatankPart) GetMigratingLastUpdateAttemptProcessOk() (*SpProcess, bool)`

GetMigratingLastUpdateAttemptProcessOk returns a tuple with the MigratingLastUpdateAttemptProcess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratingLastUpdateAttemptProcess

`func (o *DatatankPart) SetMigratingLastUpdateAttemptProcess(v SpProcess)`

SetMigratingLastUpdateAttemptProcess sets MigratingLastUpdateAttemptProcess field to given value.

### HasMigratingLastUpdateAttemptProcess

`func (o *DatatankPart) HasMigratingLastUpdateAttemptProcess() bool`

HasMigratingLastUpdateAttemptProcess returns a boolean if a field has been set.

### GetMigratingLastUpdateAttemptProcessId

`func (o *DatatankPart) GetMigratingLastUpdateAttemptProcessId() string`

GetMigratingLastUpdateAttemptProcessId returns the MigratingLastUpdateAttemptProcessId field if non-nil, zero value otherwise.

### GetMigratingLastUpdateAttemptProcessIdOk

`func (o *DatatankPart) GetMigratingLastUpdateAttemptProcessIdOk() (*string, bool)`

GetMigratingLastUpdateAttemptProcessIdOk returns a tuple with the MigratingLastUpdateAttemptProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratingLastUpdateAttemptProcessId

`func (o *DatatankPart) SetMigratingLastUpdateAttemptProcessId(v string)`

SetMigratingLastUpdateAttemptProcessId sets MigratingLastUpdateAttemptProcessId field to given value.

### HasMigratingLastUpdateAttemptProcessId

`func (o *DatatankPart) HasMigratingLastUpdateAttemptProcessId() bool`

HasMigratingLastUpdateAttemptProcessId returns a boolean if a field has been set.

### GetPartitionCandidate

`func (o *DatatankPart) GetPartitionCandidate() string`

GetPartitionCandidate returns the PartitionCandidate field if non-nil, zero value otherwise.

### GetPartitionCandidateOk

`func (o *DatatankPart) GetPartitionCandidateOk() (*string, bool)`

GetPartitionCandidateOk returns a tuple with the PartitionCandidate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionCandidate

`func (o *DatatankPart) SetPartitionCandidate(v string)`

SetPartitionCandidate sets PartitionCandidate field to given value.

### HasPartitionCandidate

`func (o *DatatankPart) HasPartitionCandidate() bool`

HasPartitionCandidate returns a boolean if a field has been set.

### GetPartitionCurrent

`func (o *DatatankPart) GetPartitionCurrent() string`

GetPartitionCurrent returns the PartitionCurrent field if non-nil, zero value otherwise.

### GetPartitionCurrentOk

`func (o *DatatankPart) GetPartitionCurrentOk() (*string, bool)`

GetPartitionCurrentOk returns a tuple with the PartitionCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionCurrent

`func (o *DatatankPart) SetPartitionCurrent(v string)`

SetPartitionCurrent sets PartitionCurrent field to given value.


### GetPipeline

`func (o *DatatankPart) GetPipeline() Pipeline`

GetPipeline returns the Pipeline field if non-nil, zero value otherwise.

### GetPipelineOk

`func (o *DatatankPart) GetPipelineOk() (*Pipeline, bool)`

GetPipelineOk returns a tuple with the Pipeline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipeline

`func (o *DatatankPart) SetPipeline(v Pipeline)`

SetPipeline sets Pipeline field to given value.

### HasPipeline

`func (o *DatatankPart) HasPipeline() bool`

HasPipeline returns a boolean if a field has been set.

### GetPipelineId

`func (o *DatatankPart) GetPipelineId() string`

GetPipelineId returns the PipelineId field if non-nil, zero value otherwise.

### GetPipelineIdOk

`func (o *DatatankPart) GetPipelineIdOk() (*string, bool)`

GetPipelineIdOk returns a tuple with the PipelineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPipelineId

`func (o *DatatankPart) SetPipelineId(v string)`

SetPipelineId sets PipelineId field to given value.


### GetSourceConnection

`func (o *DatatankPart) GetSourceConnection() WorkspaceSchema`

GetSourceConnection returns the SourceConnection field if non-nil, zero value otherwise.

### GetSourceConnectionOk

`func (o *DatatankPart) GetSourceConnectionOk() (*WorkspaceSchema, bool)`

GetSourceConnectionOk returns a tuple with the SourceConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConnection

`func (o *DatatankPart) SetSourceConnection(v WorkspaceSchema)`

SetSourceConnection sets SourceConnection field to given value.

### HasSourceConnection

`func (o *DatatankPart) HasSourceConnection() bool`

HasSourceConnection returns a boolean if a field has been set.

### GetSourceConnectionId

`func (o *DatatankPart) GetSourceConnectionId() string`

GetSourceConnectionId returns the SourceConnectionId field if non-nil, zero value otherwise.

### GetSourceConnectionIdOk

`func (o *DatatankPart) GetSourceConnectionIdOk() (*string, bool)`

GetSourceConnectionIdOk returns a tuple with the SourceConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceConnectionId

`func (o *DatatankPart) SetSourceConnectionId(v string)`

SetSourceConnectionId sets SourceConnectionId field to given value.

### HasSourceConnectionId

`func (o *DatatankPart) HasSourceConnectionId() bool`

HasSourceConnectionId returns a boolean if a field has been set.

### GetState

`func (o *DatatankPart) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DatatankPart) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DatatankPart) SetState(v string)`

SetState sets State field to given value.


### GetStateReason

`func (o *DatatankPart) GetStateReason() string`

GetStateReason returns the StateReason field if non-nil, zero value otherwise.

### GetStateReasonOk

`func (o *DatatankPart) GetStateReasonOk() (*string, bool)`

GetStateReasonOk returns a tuple with the StateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateReason

`func (o *DatatankPart) SetStateReason(v string)`

SetStateReason sets StateReason field to given value.

### HasStateReason

`func (o *DatatankPart) HasStateReason() bool`

HasStateReason returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DatatankPart) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DatatankPart) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DatatankPart) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DatatankPart) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *DatatankPart) GetUpdatedBy() User`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *DatatankPart) GetUpdatedByOk() (*User, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *DatatankPart) SetUpdatedBy(v User)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *DatatankPart) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetUpdatedById

`func (o *DatatankPart) GetUpdatedById() string`

GetUpdatedById returns the UpdatedById field if non-nil, zero value otherwise.

### GetUpdatedByIdOk

`func (o *DatatankPart) GetUpdatedByIdOk() (*string, bool)`

GetUpdatedByIdOk returns a tuple with the UpdatedById field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedById

`func (o *DatatankPart) SetUpdatedById(v string)`

SetUpdatedById sets UpdatedById field to given value.


### GetVersionId

`func (o *DatatankPart) GetVersionId() int32`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *DatatankPart) GetVersionIdOk() (*int32, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *DatatankPart) SetVersionId(v int32)`

SetVersionId sets VersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


