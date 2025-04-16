# WorkspaceQueryResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | [**[]WorkspaceQueryResultColumn**](WorkspaceQueryResultColumn.md) |  | 
**Detail** | Pointer to **string** |  | [optional] 
**Instance** | Pointer to **string** |  | [optional] 
**Rows** | **map[string]interface{}** |  | 
**Status** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewWorkspaceQueryResult

`func NewWorkspaceQueryResult(columns []WorkspaceQueryResultColumn, rows map[string]interface{}, ) *WorkspaceQueryResult`

NewWorkspaceQueryResult instantiates a new WorkspaceQueryResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceQueryResultWithDefaults

`func NewWorkspaceQueryResultWithDefaults() *WorkspaceQueryResult`

NewWorkspaceQueryResultWithDefaults instantiates a new WorkspaceQueryResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *WorkspaceQueryResult) GetColumns() []WorkspaceQueryResultColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *WorkspaceQueryResult) GetColumnsOk() (*[]WorkspaceQueryResultColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *WorkspaceQueryResult) SetColumns(v []WorkspaceQueryResultColumn)`

SetColumns sets Columns field to given value.


### GetDetail

`func (o *WorkspaceQueryResult) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *WorkspaceQueryResult) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *WorkspaceQueryResult) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *WorkspaceQueryResult) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetInstance

`func (o *WorkspaceQueryResult) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *WorkspaceQueryResult) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *WorkspaceQueryResult) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *WorkspaceQueryResult) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetRows

`func (o *WorkspaceQueryResult) GetRows() map[string]interface{}`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *WorkspaceQueryResult) GetRowsOk() (*map[string]interface{}, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *WorkspaceQueryResult) SetRows(v map[string]interface{})`

SetRows sets Rows field to given value.


### GetStatus

`func (o *WorkspaceQueryResult) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkspaceQueryResult) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkspaceQueryResult) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkspaceQueryResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTitle

`func (o *WorkspaceQueryResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WorkspaceQueryResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WorkspaceQueryResult) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WorkspaceQueryResult) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *WorkspaceQueryResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkspaceQueryResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkspaceQueryResult) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkspaceQueryResult) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


