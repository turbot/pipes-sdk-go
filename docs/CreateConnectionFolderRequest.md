# CreateConnectionFolderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentId** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 

## Methods

### NewCreateConnectionFolderRequest

`func NewCreateConnectionFolderRequest(title string, ) *CreateConnectionFolderRequest`

NewCreateConnectionFolderRequest instantiates a new CreateConnectionFolderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConnectionFolderRequestWithDefaults

`func NewCreateConnectionFolderRequestWithDefaults() *CreateConnectionFolderRequest`

NewCreateConnectionFolderRequestWithDefaults instantiates a new CreateConnectionFolderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentId

`func (o *CreateConnectionFolderRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CreateConnectionFolderRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CreateConnectionFolderRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CreateConnectionFolderRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetTitle

`func (o *CreateConnectionFolderRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateConnectionFolderRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateConnectionFolderRequest) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


