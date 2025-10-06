# \OrgServiceAccountTokens

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](OrgServiceAccountTokens.md#Create) | **Post** /org/{org_handle}/service_account/{service_account_identifier}/token | Create an organization service account token
[**Delete**](OrgServiceAccountTokens.md#Delete) | **Delete** /org/{org_handle}/service_account/{service_account_identifier}/token/{token_id} | Delete an organization service account token
[**Get**](OrgServiceAccountTokens.md#Get) | **Get** /org/{org_handle}/service_account/{service_account_identifier}/token/{token_id} | Get an organization service account token
[**List**](OrgServiceAccountTokens.md#List) | **Get** /org/{org_handle}/service_account/{service_account_identifier}/token | List organization service account tokens
[**Update**](OrgServiceAccountTokens.md#Update) | **Patch** /org/{org_handle}/service_account/{service_account_identifier}/token/{token_id} | Update an organization service account token



## Create

> TokenWithValue Create(ctx, orgHandle, serviceAccountIdentifier).Body(body).Execute()

Create an organization service account token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgHandle := "orgHandle_example" // string | The handle of the organization
    serviceAccountIdentifier := "serviceAccountIdentifier_example" // string | The ID of the service account
    body := *openapiclient.NewCreateTokenRequest() // CreateTokenRequest | The request body containing token details

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgServiceAccountTokens.Create(context.Background(), orgHandle, serviceAccountIdentifier).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgServiceAccountTokens.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: TokenWithValue
    fmt.Fprintf(os.Stdout, "Response from `OrgServiceAccountTokens.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization | 
**serviceAccountIdentifier** | **string** | The ID of the service account | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**CreateTokenRequest**](CreateTokenRequest.md) | The request body containing token details | 

### Return type

[**TokenWithValue**](TokenWithValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Token Delete(ctx, orgHandle, serviceAccountIdentifier, tokenId).Execute()

Delete an organization service account token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgHandle := "orgHandle_example" // string | The handle of the organization
    serviceAccountIdentifier := "serviceAccountIdentifier_example" // string | The ID of the service account
    tokenId := "tokenId_example" // string | The identifier of the token

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgServiceAccountTokens.Delete(context.Background(), orgHandle, serviceAccountIdentifier, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgServiceAccountTokens.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: Token
    fmt.Fprintf(os.Stdout, "Response from `OrgServiceAccountTokens.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization | 
**serviceAccountIdentifier** | **string** | The ID of the service account | 
**tokenId** | **string** | The identifier of the token | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Token**](Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Token Get(ctx, orgHandle, serviceAccountIdentifier, tokenId).Execute()

Get an organization service account token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgHandle := "orgHandle_example" // string | The handle of the organization
    serviceAccountIdentifier := "serviceAccountIdentifier_example" // string | The ID of the service account
    tokenId := "tokenId_example" // string | The identifier of the token

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgServiceAccountTokens.Get(context.Background(), orgHandle, serviceAccountIdentifier, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgServiceAccountTokens.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Token
    fmt.Fprintf(os.Stdout, "Response from `OrgServiceAccountTokens.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization | 
**serviceAccountIdentifier** | **string** | The ID of the service account | 
**tokenId** | **string** | The identifier of the token | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Token**](Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListTokensResponse List(ctx, orgHandle, serviceAccountIdentifier).Limit(limit).NextToken(nextToken).Execute()

List organization service account tokens



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgHandle := "orgHandle_example" // string | The handle of the organization
    serviceAccountIdentifier := "serviceAccountIdentifier_example" // string | The ID of the service account
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgServiceAccountTokens.List(context.Background(), orgHandle, serviceAccountIdentifier).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgServiceAccountTokens.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListTokensResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgServiceAccountTokens.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization | 
**serviceAccountIdentifier** | **string** | The ID of the service account | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListTokensResponse**](ListTokensResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Token Update(ctx, orgHandle, serviceAccountIdentifier, tokenId).Body(body).Execute()

Update an organization service account token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orgHandle := "orgHandle_example" // string | The handle of the organization
    serviceAccountIdentifier := "serviceAccountIdentifier_example" // string | The ID of the service account
    tokenId := "tokenId_example" // string | The identifier of the token
    body := *openapiclient.NewUpdateTokenRequest("Status_example") // UpdateTokenRequest | The request body containing the updated token details

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgServiceAccountTokens.Update(context.Background(), orgHandle, serviceAccountIdentifier, tokenId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgServiceAccountTokens.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Token
    fmt.Fprintf(os.Stdout, "Response from `OrgServiceAccountTokens.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization | 
**serviceAccountIdentifier** | **string** | The ID of the service account | 
**tokenId** | **string** | The identifier of the token | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**UpdateTokenRequest**](UpdateTokenRequest.md) | The request body containing the updated token details | 

### Return type

[**Token**](Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

