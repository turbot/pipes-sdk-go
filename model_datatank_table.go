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

// DatatankTable struct for DatatankTable
type DatatankTable struct {
	// The time of creation in ISO 8601 UTC.
	CreatedAt string `json:"created_at"`
	// User information for the user who created this.
	CreatedBy *User `json:"created_by,omitempty"`
	// The ID of the user that created this.
	CreatedById string    `json:"created_by_id"`
	Datatank    *Datatank `json:"datatank,omitempty"`
	DatatankId  string    `json:"datatank_id"`
	// The time the item was deleted in ISO 8601 UTC.
	DeletedAt *string `json:"deleted_at,omitempty"`
	// User information for the user that performed the deletion.
	DeletedBy *User `json:"deleted_by,omitempty"`
	// The ID of the user that performed the deletion.
	DeletedById        string                  `json:"deleted_by_id"`
	Description        *string                 `json:"description,omitempty"`
	DesiredState       DesiredState            `json:"desired_state"`
	Frequency          PipelineFrequency       `json:"frequency"`
	Freshness          *DatatankTableFreshness `json:"freshness,omitempty"`
	Id                 string                  `json:"id"`
	MigratingFreshness *DatatankTableFreshness `json:"migrating_freshness,omitempty"`
	MigratingName      *string                 `json:"migrating_name,omitempty"`
	Name               string                  `json:"name"`
	PartPer            *string                 `json:"part_per,omitempty"`
	SourceQuery        *string                 `json:"source_query,omitempty"`
	SourceSchema       *string                 `json:"source_schema,omitempty"`
	SourceTable        *string                 `json:"source_table,omitempty"`
	State              DatatankTableState      `json:"state"`
	StateReason        *string                 `json:"state_reason,omitempty"`
	Type               string                  `json:"type"`
	// The time of the last update in ISO 8601 UTC.
	UpdatedAt *string `json:"updated_at,omitempty"`
	// User information for the last user to update this.
	UpdatedBy *User `json:"updated_by,omitempty"`
	// The ID of the user that performed the last update.
	UpdatedById string `json:"updated_by_id"`
	// The version ID of this item. Pass this version ID via an If-Match header when performing mutation operations on the item.
	VersionId int32 `json:"version_id"`
}

// NewDatatankTable instantiates a new DatatankTable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatatankTable(createdAt string, createdById string, datatankId string, deletedById string, desiredState DesiredState, frequency PipelineFrequency, id string, name string, state DatatankTableState, type_ string, updatedById string, versionId int32) *DatatankTable {
	this := DatatankTable{}
	this.CreatedAt = createdAt
	this.CreatedById = createdById
	this.DatatankId = datatankId
	this.DeletedById = deletedById
	this.DesiredState = desiredState
	this.Frequency = frequency
	this.Id = id
	this.Name = name
	this.State = state
	this.Type = type_
	this.UpdatedById = updatedById
	this.VersionId = versionId
	return &this
}

// NewDatatankTableWithDefaults instantiates a new DatatankTable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatatankTableWithDefaults() *DatatankTable {
	this := DatatankTable{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *DatatankTable) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *DatatankTable) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *DatatankTable) GetCreatedBy() User {
	if o == nil || o.CreatedBy == nil {
		var ret User
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetCreatedByOk() (*User, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *DatatankTable) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given User and assigns it to the CreatedBy field.
func (o *DatatankTable) SetCreatedBy(v User) {
	o.CreatedBy = &v
}

// GetCreatedById returns the CreatedById field value
func (o *DatatankTable) GetCreatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedById
}

// GetCreatedByIdOk returns a tuple with the CreatedById field value
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetCreatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedById, true
}

// SetCreatedById sets field value
func (o *DatatankTable) SetCreatedById(v string) {
	o.CreatedById = v
}

// GetDatatank returns the Datatank field value if set, zero value otherwise.
func (o *DatatankTable) GetDatatank() Datatank {
	if o == nil || o.Datatank == nil {
		var ret Datatank
		return ret
	}
	return *o.Datatank
}

// GetDatatankOk returns a tuple with the Datatank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetDatatankOk() (*Datatank, bool) {
	if o == nil || o.Datatank == nil {
		return nil, false
	}
	return o.Datatank, true
}

// HasDatatank returns a boolean if a field has been set.
func (o *DatatankTable) HasDatatank() bool {
	if o != nil && o.Datatank != nil {
		return true
	}

	return false
}

// SetDatatank gets a reference to the given Datatank and assigns it to the Datatank field.
func (o *DatatankTable) SetDatatank(v Datatank) {
	o.Datatank = &v
}

// GetDatatankId returns the DatatankId field value
func (o *DatatankTable) GetDatatankId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatatankId
}

// GetDatatankIdOk returns a tuple with the DatatankId field value
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetDatatankIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatatankId, true
}

// SetDatatankId sets field value
func (o *DatatankTable) SetDatatankId(v string) {
	o.DatatankId = v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise.
func (o *DatatankTable) GetDeletedAt() string {
	if o == nil || o.DeletedAt == nil {
		var ret string
		return ret
	}
	return *o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetDeletedAtOk() (*string, bool) {
	if o == nil || o.DeletedAt == nil {
		return nil, false
	}
	return o.DeletedAt, true
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *DatatankTable) HasDeletedAt() bool {
	if o != nil && o.DeletedAt != nil {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given string and assigns it to the DeletedAt field.
func (o *DatatankTable) SetDeletedAt(v string) {
	o.DeletedAt = &v
}

// GetDeletedBy returns the DeletedBy field value if set, zero value otherwise.
func (o *DatatankTable) GetDeletedBy() User {
	if o == nil || o.DeletedBy == nil {
		var ret User
		return ret
	}
	return *o.DeletedBy
}

// GetDeletedByOk returns a tuple with the DeletedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetDeletedByOk() (*User, bool) {
	if o == nil || o.DeletedBy == nil {
		return nil, false
	}
	return o.DeletedBy, true
}

// HasDeletedBy returns a boolean if a field has been set.
func (o *DatatankTable) HasDeletedBy() bool {
	if o != nil && o.DeletedBy != nil {
		return true
	}

	return false
}

// SetDeletedBy gets a reference to the given User and assigns it to the DeletedBy field.
func (o *DatatankTable) SetDeletedBy(v User) {
	o.DeletedBy = &v
}

// GetDeletedById returns the DeletedById field value
func (o *DatatankTable) GetDeletedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeletedById
}

// GetDeletedByIdOk returns a tuple with the DeletedById field value
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetDeletedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedById, true
}

// SetDeletedById sets field value
func (o *DatatankTable) SetDeletedById(v string) {
	o.DeletedById = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DatatankTable) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DatatankTable) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DatatankTable) SetDescription(v string) {
	o.Description = &v
}

// GetDesiredState returns the DesiredState field value
func (o *DatatankTable) GetDesiredState() DesiredState {
	if o == nil {
		var ret DesiredState
		return ret
	}

	return o.DesiredState
}

// GetDesiredStateOk returns a tuple with the DesiredState field value
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetDesiredStateOk() (*DesiredState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DesiredState, true
}

// SetDesiredState sets field value
func (o *DatatankTable) SetDesiredState(v DesiredState) {
	o.DesiredState = v
}

// GetFrequency returns the Frequency field value
func (o *DatatankTable) GetFrequency() PipelineFrequency {
	if o == nil {
		var ret PipelineFrequency
		return ret
	}

	return o.Frequency
}

// GetFrequencyOk returns a tuple with the Frequency field value
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetFrequencyOk() (*PipelineFrequency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Frequency, true
}

// SetFrequency sets field value
func (o *DatatankTable) SetFrequency(v PipelineFrequency) {
	o.Frequency = v
}

// GetFreshness returns the Freshness field value if set, zero value otherwise.
func (o *DatatankTable) GetFreshness() DatatankTableFreshness {
	if o == nil || o.Freshness == nil {
		var ret DatatankTableFreshness
		return ret
	}
	return *o.Freshness
}

// GetFreshnessOk returns a tuple with the Freshness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetFreshnessOk() (*DatatankTableFreshness, bool) {
	if o == nil || o.Freshness == nil {
		return nil, false
	}
	return o.Freshness, true
}

// HasFreshness returns a boolean if a field has been set.
func (o *DatatankTable) HasFreshness() bool {
	if o != nil && o.Freshness != nil {
		return true
	}

	return false
}

// SetFreshness gets a reference to the given DatatankTableFreshness and assigns it to the Freshness field.
func (o *DatatankTable) SetFreshness(v DatatankTableFreshness) {
	o.Freshness = &v
}

// GetId returns the Id field value
func (o *DatatankTable) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DatatankTable) SetId(v string) {
	o.Id = v
}

// GetMigratingFreshness returns the MigratingFreshness field value if set, zero value otherwise.
func (o *DatatankTable) GetMigratingFreshness() DatatankTableFreshness {
	if o == nil || o.MigratingFreshness == nil {
		var ret DatatankTableFreshness
		return ret
	}
	return *o.MigratingFreshness
}

// GetMigratingFreshnessOk returns a tuple with the MigratingFreshness field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetMigratingFreshnessOk() (*DatatankTableFreshness, bool) {
	if o == nil || o.MigratingFreshness == nil {
		return nil, false
	}
	return o.MigratingFreshness, true
}

// HasMigratingFreshness returns a boolean if a field has been set.
func (o *DatatankTable) HasMigratingFreshness() bool {
	if o != nil && o.MigratingFreshness != nil {
		return true
	}

	return false
}

// SetMigratingFreshness gets a reference to the given DatatankTableFreshness and assigns it to the MigratingFreshness field.
func (o *DatatankTable) SetMigratingFreshness(v DatatankTableFreshness) {
	o.MigratingFreshness = &v
}

// GetMigratingName returns the MigratingName field value if set, zero value otherwise.
func (o *DatatankTable) GetMigratingName() string {
	if o == nil || o.MigratingName == nil {
		var ret string
		return ret
	}
	return *o.MigratingName
}

// GetMigratingNameOk returns a tuple with the MigratingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetMigratingNameOk() (*string, bool) {
	if o == nil || o.MigratingName == nil {
		return nil, false
	}
	return o.MigratingName, true
}

// HasMigratingName returns a boolean if a field has been set.
func (o *DatatankTable) HasMigratingName() bool {
	if o != nil && o.MigratingName != nil {
		return true
	}

	return false
}

// SetMigratingName gets a reference to the given string and assigns it to the MigratingName field.
func (o *DatatankTable) SetMigratingName(v string) {
	o.MigratingName = &v
}

// GetName returns the Name field value
func (o *DatatankTable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DatatankTable) SetName(v string) {
	o.Name = v
}

// GetPartPer returns the PartPer field value if set, zero value otherwise.
func (o *DatatankTable) GetPartPer() string {
	if o == nil || o.PartPer == nil {
		var ret string
		return ret
	}
	return *o.PartPer
}

// GetPartPerOk returns a tuple with the PartPer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetPartPerOk() (*string, bool) {
	if o == nil || o.PartPer == nil {
		return nil, false
	}
	return o.PartPer, true
}

// HasPartPer returns a boolean if a field has been set.
func (o *DatatankTable) HasPartPer() bool {
	if o != nil && o.PartPer != nil {
		return true
	}

	return false
}

// SetPartPer gets a reference to the given string and assigns it to the PartPer field.
func (o *DatatankTable) SetPartPer(v string) {
	o.PartPer = &v
}

// GetSourceQuery returns the SourceQuery field value if set, zero value otherwise.
func (o *DatatankTable) GetSourceQuery() string {
	if o == nil || o.SourceQuery == nil {
		var ret string
		return ret
	}
	return *o.SourceQuery
}

// GetSourceQueryOk returns a tuple with the SourceQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetSourceQueryOk() (*string, bool) {
	if o == nil || o.SourceQuery == nil {
		return nil, false
	}
	return o.SourceQuery, true
}

// HasSourceQuery returns a boolean if a field has been set.
func (o *DatatankTable) HasSourceQuery() bool {
	if o != nil && o.SourceQuery != nil {
		return true
	}

	return false
}

// SetSourceQuery gets a reference to the given string and assigns it to the SourceQuery field.
func (o *DatatankTable) SetSourceQuery(v string) {
	o.SourceQuery = &v
}

// GetSourceSchema returns the SourceSchema field value if set, zero value otherwise.
func (o *DatatankTable) GetSourceSchema() string {
	if o == nil || o.SourceSchema == nil {
		var ret string
		return ret
	}
	return *o.SourceSchema
}

// GetSourceSchemaOk returns a tuple with the SourceSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetSourceSchemaOk() (*string, bool) {
	if o == nil || o.SourceSchema == nil {
		return nil, false
	}
	return o.SourceSchema, true
}

// HasSourceSchema returns a boolean if a field has been set.
func (o *DatatankTable) HasSourceSchema() bool {
	if o != nil && o.SourceSchema != nil {
		return true
	}

	return false
}

// SetSourceSchema gets a reference to the given string and assigns it to the SourceSchema field.
func (o *DatatankTable) SetSourceSchema(v string) {
	o.SourceSchema = &v
}

// GetSourceTable returns the SourceTable field value if set, zero value otherwise.
func (o *DatatankTable) GetSourceTable() string {
	if o == nil || o.SourceTable == nil {
		var ret string
		return ret
	}
	return *o.SourceTable
}

// GetSourceTableOk returns a tuple with the SourceTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetSourceTableOk() (*string, bool) {
	if o == nil || o.SourceTable == nil {
		return nil, false
	}
	return o.SourceTable, true
}

// HasSourceTable returns a boolean if a field has been set.
func (o *DatatankTable) HasSourceTable() bool {
	if o != nil && o.SourceTable != nil {
		return true
	}

	return false
}

// SetSourceTable gets a reference to the given string and assigns it to the SourceTable field.
func (o *DatatankTable) SetSourceTable(v string) {
	o.SourceTable = &v
}

// GetState returns the State field value
func (o *DatatankTable) GetState() DatatankTableState {
	if o == nil {
		var ret DatatankTableState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetStateOk() (*DatatankTableState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *DatatankTable) SetState(v DatatankTableState) {
	o.State = v
}

// GetStateReason returns the StateReason field value if set, zero value otherwise.
func (o *DatatankTable) GetStateReason() string {
	if o == nil || o.StateReason == nil {
		var ret string
		return ret
	}
	return *o.StateReason
}

// GetStateReasonOk returns a tuple with the StateReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetStateReasonOk() (*string, bool) {
	if o == nil || o.StateReason == nil {
		return nil, false
	}
	return o.StateReason, true
}

// HasStateReason returns a boolean if a field has been set.
func (o *DatatankTable) HasStateReason() bool {
	if o != nil && o.StateReason != nil {
		return true
	}

	return false
}

// SetStateReason gets a reference to the given string and assigns it to the StateReason field.
func (o *DatatankTable) SetStateReason(v string) {
	o.StateReason = &v
}

// GetType returns the Type field value
func (o *DatatankTable) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DatatankTable) SetType(v string) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DatatankTable) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DatatankTable) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *DatatankTable) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetUpdatedBy returns the UpdatedBy field value if set, zero value otherwise.
func (o *DatatankTable) GetUpdatedBy() User {
	if o == nil || o.UpdatedBy == nil {
		var ret User
		return ret
	}
	return *o.UpdatedBy
}

// GetUpdatedByOk returns a tuple with the UpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetUpdatedByOk() (*User, bool) {
	if o == nil || o.UpdatedBy == nil {
		return nil, false
	}
	return o.UpdatedBy, true
}

// HasUpdatedBy returns a boolean if a field has been set.
func (o *DatatankTable) HasUpdatedBy() bool {
	if o != nil && o.UpdatedBy != nil {
		return true
	}

	return false
}

// SetUpdatedBy gets a reference to the given User and assigns it to the UpdatedBy field.
func (o *DatatankTable) SetUpdatedBy(v User) {
	o.UpdatedBy = &v
}

// GetUpdatedById returns the UpdatedById field value
func (o *DatatankTable) GetUpdatedById() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedById
}

// GetUpdatedByIdOk returns a tuple with the UpdatedById field value
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetUpdatedByIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedById, true
}

// SetUpdatedById sets field value
func (o *DatatankTable) SetUpdatedById(v string) {
	o.UpdatedById = v
}

// GetVersionId returns the VersionId field value
func (o *DatatankTable) GetVersionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *DatatankTable) GetVersionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *DatatankTable) SetVersionId(v int32) {
	o.VersionId = v
}

func (o DatatankTable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if true {
		toSerialize["created_by_id"] = o.CreatedById
	}
	if o.Datatank != nil {
		toSerialize["datatank"] = o.Datatank
	}
	if true {
		toSerialize["datatank_id"] = o.DatatankId
	}
	if o.DeletedAt != nil {
		toSerialize["deleted_at"] = o.DeletedAt
	}
	if o.DeletedBy != nil {
		toSerialize["deleted_by"] = o.DeletedBy
	}
	if true {
		toSerialize["deleted_by_id"] = o.DeletedById
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["desired_state"] = o.DesiredState
	}
	if true {
		toSerialize["frequency"] = o.Frequency
	}
	if o.Freshness != nil {
		toSerialize["freshness"] = o.Freshness
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.MigratingFreshness != nil {
		toSerialize["migrating_freshness"] = o.MigratingFreshness
	}
	if o.MigratingName != nil {
		toSerialize["migrating_name"] = o.MigratingName
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.PartPer != nil {
		toSerialize["part_per"] = o.PartPer
	}
	if o.SourceQuery != nil {
		toSerialize["source_query"] = o.SourceQuery
	}
	if o.SourceSchema != nil {
		toSerialize["source_schema"] = o.SourceSchema
	}
	if o.SourceTable != nil {
		toSerialize["source_table"] = o.SourceTable
	}
	if true {
		toSerialize["state"] = o.State
	}
	if o.StateReason != nil {
		toSerialize["state_reason"] = o.StateReason
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.UpdatedBy != nil {
		toSerialize["updated_by"] = o.UpdatedBy
	}
	if true {
		toSerialize["updated_by_id"] = o.UpdatedById
	}
	if true {
		toSerialize["version_id"] = o.VersionId
	}
	return json.Marshal(toSerialize)
}

type NullableDatatankTable struct {
	value *DatatankTable
	isSet bool
}

func (v NullableDatatankTable) Get() *DatatankTable {
	return v.value
}

func (v *NullableDatatankTable) Set(val *DatatankTable) {
	v.value = val
	v.isSet = true
}

func (v NullableDatatankTable) IsSet() bool {
	return v.isSet
}

func (v *NullableDatatankTable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatatankTable(val *DatatankTable) *NullableDatatankTable {
	return &NullableDatatankTable{value: val, isSet: true}
}

func (v NullableDatatankTable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatatankTable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
