# UsageMetric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimension** | [**UsageDimensionType**](UsageDimensionType.md) | The dimension for this usage metric record. | 
**IdentityId** | Pointer to **string** | The identity ID for this usage metric record | [optional] 
**InstanceType** | Pointer to **string** | The instance type for the usage record.                                                                                                                                                       // The instance type for this usage metric record. Only populated if the usage is for a workspace. | [optional] 
**Metric** | [**UsageMetricType**](UsageMetricType.md) | The metric for this usage metric record. | 
**TenantId** | Pointer to **string** | The tenant ID for this usage metric record | [optional] 
**Unit** | [**UsageUnitType**](UsageUnitType.md) | The unit for this usage metric record. | 
**UsageDate** | **string** | The start time for this usage metric record. | 
**Value** | Pointer to **int64** | The value for this usage metric record. | [optional] 
**ValueRounded** | Pointer to **int64** | The rounded value for this usage metric record to the nearest billable unit. | [optional] 
**ValueWeighted** | Pointer to **int64** | The value for this usage metric record weighted according to any billing rules such as DB instance type. | [optional] 
**WorkspaceId** | Pointer to **string** | The workspace ID for this usage metric record. Only populated if the usage is for a workspace. | [optional] 

## Methods

### NewUsageMetric

`func NewUsageMetric(dimension UsageDimensionType, metric UsageMetricType, unit UsageUnitType, usageDate string, ) *UsageMetric`

NewUsageMetric instantiates a new UsageMetric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageMetricWithDefaults

`func NewUsageMetricWithDefaults() *UsageMetric`

NewUsageMetricWithDefaults instantiates a new UsageMetric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimension

`func (o *UsageMetric) GetDimension() UsageDimensionType`

GetDimension returns the Dimension field if non-nil, zero value otherwise.

### GetDimensionOk

`func (o *UsageMetric) GetDimensionOk() (*UsageDimensionType, bool)`

GetDimensionOk returns a tuple with the Dimension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimension

`func (o *UsageMetric) SetDimension(v UsageDimensionType)`

SetDimension sets Dimension field to given value.


### GetIdentityId

`func (o *UsageMetric) GetIdentityId() string`

GetIdentityId returns the IdentityId field if non-nil, zero value otherwise.

### GetIdentityIdOk

`func (o *UsageMetric) GetIdentityIdOk() (*string, bool)`

GetIdentityIdOk returns a tuple with the IdentityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityId

`func (o *UsageMetric) SetIdentityId(v string)`

SetIdentityId sets IdentityId field to given value.

### HasIdentityId

`func (o *UsageMetric) HasIdentityId() bool`

HasIdentityId returns a boolean if a field has been set.

### GetInstanceType

`func (o *UsageMetric) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *UsageMetric) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *UsageMetric) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *UsageMetric) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetMetric

`func (o *UsageMetric) GetMetric() UsageMetricType`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *UsageMetric) GetMetricOk() (*UsageMetricType, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *UsageMetric) SetMetric(v UsageMetricType)`

SetMetric sets Metric field to given value.


### GetTenantId

`func (o *UsageMetric) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UsageMetric) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UsageMetric) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *UsageMetric) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUnit

`func (o *UsageMetric) GetUnit() UsageUnitType`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *UsageMetric) GetUnitOk() (*UsageUnitType, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *UsageMetric) SetUnit(v UsageUnitType)`

SetUnit sets Unit field to given value.


### GetUsageDate

`func (o *UsageMetric) GetUsageDate() string`

GetUsageDate returns the UsageDate field if non-nil, zero value otherwise.

### GetUsageDateOk

`func (o *UsageMetric) GetUsageDateOk() (*string, bool)`

GetUsageDateOk returns a tuple with the UsageDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageDate

`func (o *UsageMetric) SetUsageDate(v string)`

SetUsageDate sets UsageDate field to given value.


### GetValue

`func (o *UsageMetric) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UsageMetric) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UsageMetric) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *UsageMetric) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetValueRounded

`func (o *UsageMetric) GetValueRounded() int64`

GetValueRounded returns the ValueRounded field if non-nil, zero value otherwise.

### GetValueRoundedOk

`func (o *UsageMetric) GetValueRoundedOk() (*int64, bool)`

GetValueRoundedOk returns a tuple with the ValueRounded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueRounded

`func (o *UsageMetric) SetValueRounded(v int64)`

SetValueRounded sets ValueRounded field to given value.

### HasValueRounded

`func (o *UsageMetric) HasValueRounded() bool`

HasValueRounded returns a boolean if a field has been set.

### GetValueWeighted

`func (o *UsageMetric) GetValueWeighted() int64`

GetValueWeighted returns the ValueWeighted field if non-nil, zero value otherwise.

### GetValueWeightedOk

`func (o *UsageMetric) GetValueWeightedOk() (*int64, bool)`

GetValueWeightedOk returns a tuple with the ValueWeighted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueWeighted

`func (o *UsageMetric) SetValueWeighted(v int64)`

SetValueWeighted sets ValueWeighted field to given value.

### HasValueWeighted

`func (o *UsageMetric) HasValueWeighted() bool`

HasValueWeighted returns a boolean if a field has been set.

### GetWorkspaceId

`func (o *UsageMetric) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *UsageMetric) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *UsageMetric) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *UsageMetric) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


