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

// IdentityUsage struct for IdentityUsage
type IdentityUsage struct {
	// The dimension for this usage record.
	Dimension string `json:"dimension"`
	// The identity ID for this usage record
	IdentityId *string `json:"identity_id,omitempty"`
	// The instance type for the usage record.                                                                                                                                                       // The instance type for this usage record. Only populated if the usage is for a workspace.
	InstanceType *string `json:"instance_type,omitempty"`
	// The metric for this usage record.
	Metric string `json:"metric"`
	// The unit for this usage record.
	Unit string `json:"unit"`
	// The start time for this usage record.
	UsageDate string `json:"usage_date"`
	// The value for this usage record.
	Value *int64 `json:"value,omitempty"`
	// The rounded value for this usage record to the nearest billable unit.
	ValueRounded *int64 `json:"value_rounded,omitempty"`
	// The value for this usage record weighted according to any billing rules such as DB instance type.
	ValueWeighted *int64 `json:"value_weighted,omitempty"`
	// The workspace ID for this usage record. Only populated if the usage is for a workspace.
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

// NewIdentityUsage instantiates a new IdentityUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityUsage(dimension string, metric string, unit string, usageDate string) *IdentityUsage {
	this := IdentityUsage{}
	this.Dimension = dimension
	this.Metric = metric
	this.Unit = unit
	this.UsageDate = usageDate
	return &this
}

// NewIdentityUsageWithDefaults instantiates a new IdentityUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityUsageWithDefaults() *IdentityUsage {
	this := IdentityUsage{}
	return &this
}

// GetDimension returns the Dimension field value
func (o *IdentityUsage) GetDimension() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *IdentityUsage) GetDimensionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value
func (o *IdentityUsage) SetDimension(v string) {
	o.Dimension = v
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *IdentityUsage) GetIdentityId() string {
	if o == nil || o.IdentityId == nil {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUsage) GetIdentityIdOk() (*string, bool) {
	if o == nil || o.IdentityId == nil {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *IdentityUsage) HasIdentityId() bool {
	if o != nil && o.IdentityId != nil {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *IdentityUsage) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *IdentityUsage) GetInstanceType() string {
	if o == nil || o.InstanceType == nil {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUsage) GetInstanceTypeOk() (*string, bool) {
	if o == nil || o.InstanceType == nil {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *IdentityUsage) HasInstanceType() bool {
	if o != nil && o.InstanceType != nil {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *IdentityUsage) SetInstanceType(v string) {
	o.InstanceType = &v
}

// GetMetric returns the Metric field value
func (o *IdentityUsage) GetMetric() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *IdentityUsage) GetMetricOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value
func (o *IdentityUsage) SetMetric(v string) {
	o.Metric = v
}

// GetUnit returns the Unit field value
func (o *IdentityUsage) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *IdentityUsage) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *IdentityUsage) SetUnit(v string) {
	o.Unit = v
}

// GetUsageDate returns the UsageDate field value
func (o *IdentityUsage) GetUsageDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UsageDate
}

// GetUsageDateOk returns a tuple with the UsageDate field value
// and a boolean to check if the value has been set.
func (o *IdentityUsage) GetUsageDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageDate, true
}

// SetUsageDate sets field value
func (o *IdentityUsage) SetUsageDate(v string) {
	o.UsageDate = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IdentityUsage) GetValue() int64 {
	if o == nil || o.Value == nil {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUsage) GetValueOk() (*int64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IdentityUsage) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *IdentityUsage) SetValue(v int64) {
	o.Value = &v
}

// GetValueRounded returns the ValueRounded field value if set, zero value otherwise.
func (o *IdentityUsage) GetValueRounded() int64 {
	if o == nil || o.ValueRounded == nil {
		var ret int64
		return ret
	}
	return *o.ValueRounded
}

// GetValueRoundedOk returns a tuple with the ValueRounded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUsage) GetValueRoundedOk() (*int64, bool) {
	if o == nil || o.ValueRounded == nil {
		return nil, false
	}
	return o.ValueRounded, true
}

// HasValueRounded returns a boolean if a field has been set.
func (o *IdentityUsage) HasValueRounded() bool {
	if o != nil && o.ValueRounded != nil {
		return true
	}

	return false
}

// SetValueRounded gets a reference to the given int64 and assigns it to the ValueRounded field.
func (o *IdentityUsage) SetValueRounded(v int64) {
	o.ValueRounded = &v
}

// GetValueWeighted returns the ValueWeighted field value if set, zero value otherwise.
func (o *IdentityUsage) GetValueWeighted() int64 {
	if o == nil || o.ValueWeighted == nil {
		var ret int64
		return ret
	}
	return *o.ValueWeighted
}

// GetValueWeightedOk returns a tuple with the ValueWeighted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUsage) GetValueWeightedOk() (*int64, bool) {
	if o == nil || o.ValueWeighted == nil {
		return nil, false
	}
	return o.ValueWeighted, true
}

// HasValueWeighted returns a boolean if a field has been set.
func (o *IdentityUsage) HasValueWeighted() bool {
	if o != nil && o.ValueWeighted != nil {
		return true
	}

	return false
}

// SetValueWeighted gets a reference to the given int64 and assigns it to the ValueWeighted field.
func (o *IdentityUsage) SetValueWeighted(v int64) {
	o.ValueWeighted = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *IdentityUsage) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityUsage) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *IdentityUsage) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *IdentityUsage) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o IdentityUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dimension"] = o.Dimension
	}
	if o.IdentityId != nil {
		toSerialize["identity_id"] = o.IdentityId
	}
	if o.InstanceType != nil {
		toSerialize["instance_type"] = o.InstanceType
	}
	if true {
		toSerialize["metric"] = o.Metric
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	if true {
		toSerialize["usage_date"] = o.UsageDate
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.ValueRounded != nil {
		toSerialize["value_rounded"] = o.ValueRounded
	}
	if o.ValueWeighted != nil {
		toSerialize["value_weighted"] = o.ValueWeighted
	}
	if o.WorkspaceId != nil {
		toSerialize["workspace_id"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableIdentityUsage struct {
	value *IdentityUsage
	isSet bool
}

func (v NullableIdentityUsage) Get() *IdentityUsage {
	return v.value
}

func (v *NullableIdentityUsage) Set(val *IdentityUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityUsage(val *IdentityUsage) *NullableIdentityUsage {
	return &NullableIdentityUsage{value: val, isSet: true}
}

func (v NullableIdentityUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
