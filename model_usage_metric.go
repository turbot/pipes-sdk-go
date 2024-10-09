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

// UsageMetric struct for UsageMetric
type UsageMetric struct {
	// The dimension for this usage metric record.
	Dimension UsageDimensionType `json:"dimension"`
	// The identity ID for this usage metric record
	IdentityId *string `json:"identity_id,omitempty"`
	// The instance type for the usage record.                                                                                                                                                       // The instance type for this usage metric record. Only populated if the usage is for a workspace.
	InstanceType *string `json:"instance_type,omitempty"`
	// The metric for this usage metric record.
	Metric UsageMetricType `json:"metric"`
	// The tenant ID for this usage metric record
	TenantId *string `json:"tenant_id,omitempty"`
	// The unit for this usage metric record.
	Unit UsageUnitType `json:"unit"`
	// The start time for this usage metric record.
	UsageDate string `json:"usage_date"`
	// The value for this usage metric record.
	Value *int64 `json:"value,omitempty"`
	// The rounded value for this usage metric record to the nearest billable unit.
	ValueRounded *int64 `json:"value_rounded,omitempty"`
	// The value for this usage metric record weighted according to any billing rules such as DB instance type.
	ValueWeighted *int64 `json:"value_weighted,omitempty"`
	// The workspace ID for this usage metric record. Only populated if the usage is for a workspace.
	WorkspaceId *string `json:"workspace_id,omitempty"`
}

// NewUsageMetric instantiates a new UsageMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageMetric(dimension UsageDimensionType, metric UsageMetricType, unit UsageUnitType, usageDate string) *UsageMetric {
	this := UsageMetric{}
	this.Dimension = dimension
	this.Metric = metric
	this.Unit = unit
	this.UsageDate = usageDate
	return &this
}

// NewUsageMetricWithDefaults instantiates a new UsageMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageMetricWithDefaults() *UsageMetric {
	this := UsageMetric{}
	return &this
}

// GetDimension returns the Dimension field value
func (o *UsageMetric) GetDimension() UsageDimensionType {
	if o == nil {
		var ret UsageDimensionType
		return ret
	}

	return o.Dimension
}

// GetDimensionOk returns a tuple with the Dimension field value
// and a boolean to check if the value has been set.
func (o *UsageMetric) GetDimensionOk() (*UsageDimensionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dimension, true
}

// SetDimension sets field value
func (o *UsageMetric) SetDimension(v UsageDimensionType) {
	o.Dimension = v
}

// GetIdentityId returns the IdentityId field value if set, zero value otherwise.
func (o *UsageMetric) GetIdentityId() string {
	if o == nil || o.IdentityId == nil {
		var ret string
		return ret
	}
	return *o.IdentityId
}

// GetIdentityIdOk returns a tuple with the IdentityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMetric) GetIdentityIdOk() (*string, bool) {
	if o == nil || o.IdentityId == nil {
		return nil, false
	}
	return o.IdentityId, true
}

// HasIdentityId returns a boolean if a field has been set.
func (o *UsageMetric) HasIdentityId() bool {
	if o != nil && o.IdentityId != nil {
		return true
	}

	return false
}

// SetIdentityId gets a reference to the given string and assigns it to the IdentityId field.
func (o *UsageMetric) SetIdentityId(v string) {
	o.IdentityId = &v
}

// GetInstanceType returns the InstanceType field value if set, zero value otherwise.
func (o *UsageMetric) GetInstanceType() string {
	if o == nil || o.InstanceType == nil {
		var ret string
		return ret
	}
	return *o.InstanceType
}

// GetInstanceTypeOk returns a tuple with the InstanceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMetric) GetInstanceTypeOk() (*string, bool) {
	if o == nil || o.InstanceType == nil {
		return nil, false
	}
	return o.InstanceType, true
}

// HasInstanceType returns a boolean if a field has been set.
func (o *UsageMetric) HasInstanceType() bool {
	if o != nil && o.InstanceType != nil {
		return true
	}

	return false
}

// SetInstanceType gets a reference to the given string and assigns it to the InstanceType field.
func (o *UsageMetric) SetInstanceType(v string) {
	o.InstanceType = &v
}

// GetMetric returns the Metric field value
func (o *UsageMetric) GetMetric() UsageMetricType {
	if o == nil {
		var ret UsageMetricType
		return ret
	}

	return o.Metric
}

// GetMetricOk returns a tuple with the Metric field value
// and a boolean to check if the value has been set.
func (o *UsageMetric) GetMetricOk() (*UsageMetricType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metric, true
}

// SetMetric sets field value
func (o *UsageMetric) SetMetric(v UsageMetricType) {
	o.Metric = v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *UsageMetric) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMetric) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *UsageMetric) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *UsageMetric) SetTenantId(v string) {
	o.TenantId = &v
}

// GetUnit returns the Unit field value
func (o *UsageMetric) GetUnit() UsageUnitType {
	if o == nil {
		var ret UsageUnitType
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *UsageMetric) GetUnitOk() (*UsageUnitType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *UsageMetric) SetUnit(v UsageUnitType) {
	o.Unit = v
}

// GetUsageDate returns the UsageDate field value
func (o *UsageMetric) GetUsageDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UsageDate
}

// GetUsageDateOk returns a tuple with the UsageDate field value
// and a boolean to check if the value has been set.
func (o *UsageMetric) GetUsageDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageDate, true
}

// SetUsageDate sets field value
func (o *UsageMetric) SetUsageDate(v string) {
	o.UsageDate = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *UsageMetric) GetValue() int64 {
	if o == nil || o.Value == nil {
		var ret int64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMetric) GetValueOk() (*int64, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *UsageMetric) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given int64 and assigns it to the Value field.
func (o *UsageMetric) SetValue(v int64) {
	o.Value = &v
}

// GetValueRounded returns the ValueRounded field value if set, zero value otherwise.
func (o *UsageMetric) GetValueRounded() int64 {
	if o == nil || o.ValueRounded == nil {
		var ret int64
		return ret
	}
	return *o.ValueRounded
}

// GetValueRoundedOk returns a tuple with the ValueRounded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMetric) GetValueRoundedOk() (*int64, bool) {
	if o == nil || o.ValueRounded == nil {
		return nil, false
	}
	return o.ValueRounded, true
}

// HasValueRounded returns a boolean if a field has been set.
func (o *UsageMetric) HasValueRounded() bool {
	if o != nil && o.ValueRounded != nil {
		return true
	}

	return false
}

// SetValueRounded gets a reference to the given int64 and assigns it to the ValueRounded field.
func (o *UsageMetric) SetValueRounded(v int64) {
	o.ValueRounded = &v
}

// GetValueWeighted returns the ValueWeighted field value if set, zero value otherwise.
func (o *UsageMetric) GetValueWeighted() int64 {
	if o == nil || o.ValueWeighted == nil {
		var ret int64
		return ret
	}
	return *o.ValueWeighted
}

// GetValueWeightedOk returns a tuple with the ValueWeighted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMetric) GetValueWeightedOk() (*int64, bool) {
	if o == nil || o.ValueWeighted == nil {
		return nil, false
	}
	return o.ValueWeighted, true
}

// HasValueWeighted returns a boolean if a field has been set.
func (o *UsageMetric) HasValueWeighted() bool {
	if o != nil && o.ValueWeighted != nil {
		return true
	}

	return false
}

// SetValueWeighted gets a reference to the given int64 and assigns it to the ValueWeighted field.
func (o *UsageMetric) SetValueWeighted(v int64) {
	o.ValueWeighted = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *UsageMetric) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMetric) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *UsageMetric) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *UsageMetric) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o UsageMetric) MarshalJSON() ([]byte, error) {
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
	if o.TenantId != nil {
		toSerialize["tenant_id"] = o.TenantId
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

type NullableUsageMetric struct {
	value *UsageMetric
	isSet bool
}

func (v NullableUsageMetric) Get() *UsageMetric {
	return v.value
}

func (v *NullableUsageMetric) Set(val *UsageMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageMetric(val *UsageMetric) *NullableUsageMetric {
	return &NullableUsageMetric{value: val, isSet: true}
}

func (v NullableUsageMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
