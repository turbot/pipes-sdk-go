# CreateAIKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** |  | 
**Provider** | [**AIProvider**](AIProvider.md) |  | 

## Methods

### NewCreateAIKeyRequest

`func NewCreateAIKeyRequest(apiKey string, provider AIProvider, ) *CreateAIKeyRequest`

NewCreateAIKeyRequest instantiates a new CreateAIKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAIKeyRequestWithDefaults

`func NewCreateAIKeyRequestWithDefaults() *CreateAIKeyRequest`

NewCreateAIKeyRequestWithDefaults instantiates a new CreateAIKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *CreateAIKeyRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *CreateAIKeyRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *CreateAIKeyRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetProvider

`func (o *CreateAIKeyRequest) GetProvider() AIProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CreateAIKeyRequest) GetProviderOk() (*AIProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CreateAIKeyRequest) SetProvider(v AIProvider)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


