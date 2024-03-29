# IdentityUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimension** | **string** | The dimension for this usage record. | 
**IdentityId** | Pointer to **string** | The identity ID for this usage record | [optional] 
**InstanceType** | Pointer to **string** | The instance type for the usage record.                                                                                                                                                       // The instance type for this usage record. Only populated if the usage is for a workspace. | [optional] 
**Metric** | **string** | The metric for this usage record. | 
**TenantId** | Pointer to **string** | The tenant ID for this usage record | [optional] 
**Unit** | **string** | The unit for this usage record. | 
**UsageDate** | **string** | The start time for this usage record. | 
**Value** | Pointer to **int64** | The value for this usage record. | [optional] 
**ValueRounded** | Pointer to **int64** | The rounded value for this usage record to the nearest billable unit. | [optional] 
**ValueWeighted** | Pointer to **int64** | The value for this usage record weighted according to any billing rules such as DB instance type. | [optional] 
**WorkspaceId** | Pointer to **string** | The workspace ID for this usage record. Only populated if the usage is for a workspace. | [optional] 

## Methods

### NewIdentityUsage

`func NewIdentityUsage(dimension string, metric string, unit string, usageDate string, ) *IdentityUsage`

NewIdentityUsage instantiates a new IdentityUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityUsageWithDefaults

`func NewIdentityUsageWithDefaults() *IdentityUsage`

NewIdentityUsageWithDefaults instantiates a new IdentityUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimension

`func (o *IdentityUsage) GetDimension() string`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *IdentityUsage) GetDimensionOk() (*string, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *IdentityUsage) SetDimension(v string)`

SetDimension sets Dimension field to given value.


### GetIdentityId

`func (o *IdentityUsage) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *IdentityUsage) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *IdentityUsage) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *IdentityUsage) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetInstanceType

`func (o *IdentityUsage) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *IdentityUsage) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *IdentityUsage) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *IdentityUsage) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetMetric

`func (o *IdentityUsage) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *IdentityUsage) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *IdentityUsage) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetTenantId

`func (o *IdentityUsage) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *IdentityUsage) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *IdentityUsage) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *IdentityUsage) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUnit

`func (o *IdentityUsage) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *IdentityUsage) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *IdentityUsage) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetUsageDate

`func (o *IdentityUsage) GetUsageDate() string`

GetUsageDate returns the UsageDate field if non-nil, zero value otherwise.

### GetUsageDateOk

`func (o *IdentityUsage) GetUsageDateOk() (*string, bool)`

GetUsageDateOk returns a tuple with the UsageDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageDate

`func (o *IdentityUsage) SetUsageDate(v string)`

SetUsageDate sets UsageDate field to given value.


### GetValue

`func (o *IdentityUsage) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IdentityUsage) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IdentityUsage) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *IdentityUsage) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueRounded

`func (o *IdentityUsage) GetValueRounded() int64`

GetValueRounded returns the ValueRounded field if non-nil, zero value otherwise.

### GetValueRoundedOk

`func (o *IdentityUsage) GetValueRoundedOk() (*int64, bool)`

GetValueRoundedOk returns a tuple with the ValueRounded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueRounded

`func (o *IdentityUsage) SetValueRounded(v int64)`

SetValueRounded sets ValueRounded field to given value.

### HasValueRounded

`func (o *IdentityUsage) HasValueRounded() bool`

HasValueRounded returns a boolean if a field has been set.

### GetValueWeighted

`func (o *IdentityUsage) GetValueWeighted() int64`

GetValueWeighted returns the ValueWeighted field if non-nil, zero value otherwise.

### GetValueWeightedOk

`func (o *IdentityUsage) GetValueWeightedOk() (*int64, bool)`

GetValueWeightedOk returns a tuple with the ValueWeighted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueWeighted

`func (o *IdentityUsage) SetValueWeighted(v int64)`

SetValueWeighted sets ValueWeighted field to given value.

### HasValueWeighted

`func (o *IdentityUsage) HasValueWeighted() bool`

HasValueWeighted returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *IdentityUsage) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *IdentityUsage) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *IdentityUsage) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *IdentityUsage) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


