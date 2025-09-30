# \ServiceAccounts

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](ServiceAccounts.md#Create) | **Post** /org/{org_handle}/service_account | Create a service account at the organization level
[**Create_0**](ServiceAccounts.md#Create_0) | **Post** /service_account | Create a tenant service account
[**Delete**](ServiceAccounts.md#Delete) | **Delete** /org/{org_handle}/service_account/{service_account_identifier} | Delete an organization service account
[**Delete_0**](ServiceAccounts.md#Delete_0) | **Delete** /service_account/{service_account_identifier} | Delete a tenant service account
[**OrgCreate**](ServiceAccounts.md#OrgCreate) | **Post** /org/{org_handle}/service_account/{service_account_identifier}/token | Create an organization service account token
[**OrgDelete**](ServiceAccounts.md#OrgDelete) | **Delete** /org/{org_handle}/service_account/{service_account_identifier}/token/{token_id} | Delete an organization service account token
[**OrgGet**](ServiceAccounts.md#OrgGet) | **Get** /org/{org_handle}/service_account/{service_account_identifier} | Get organization service account
[**OrgGet_0**](ServiceAccounts.md#OrgGet_0) | **Get** /org/{org_handle}/service_account/{service_account_identifier}/token/{token_id} | Get an organization service account token
[**OrgList**](ServiceAccounts.md#OrgList) | **Get** /org/{org_handle}/service_account | List organization service accounts
[**OrgListAuditLogs**](ServiceAccounts.md#OrgListAuditLogs) | **Get** /org/{org_handle}/service_account/{service_account_identifier}/audit_log | List audit logs for an org service account
[**OrgList_0**](ServiceAccounts.md#OrgList_0) | **Get** /org/{org_handle}/service_account/{service_account_identifier}/token | List organization service account tokens
[**OrgUpdate**](ServiceAccounts.md#OrgUpdate) | **Patch** /org/{org_handle}/service_account/{service_account_identifier}/token/{token_id} | Update an organization service account token
[**TenantCreate**](ServiceAccounts.md#TenantCreate) | **Post** /service_account/{service_account_identifier}/token | Create a tenant service account token
[**TenantDelete**](ServiceAccounts.md#TenantDelete) | **Delete** /service_account/{service_account_identifier}/token/{token_id} | Delete a tenant service account token
[**TenantGet**](ServiceAccounts.md#TenantGet) | **Get** /service_account/{service_account_identifier} | Get tenant service account
[**TenantGet_0**](ServiceAccounts.md#TenantGet_0) | **Get** /service_account/{service_account_identifier}/token/{token_id} | Get a tenant service account token
[**TenantList**](ServiceAccounts.md#TenantList) | **Get** /service_account | List tenant service accounts
[**TenantListAuditLogs**](ServiceAccounts.md#TenantListAuditLogs) | **Get** /service_account/{service_account_identifier}/audit_log | List audit logs for a tenant service account
[**TenantList_0**](ServiceAccounts.md#TenantList_0) | **Get** /service_account/{service_account_identifier}/token | List tenant service account tokens
[**TenantUpdate**](ServiceAccounts.md#TenantUpdate) | **Patch** /service_account/{service_account_identifier}/token/{token_id} | Update a tenant service account token
[**Update**](ServiceAccounts.md#Update) | **Patch** /org/{org_handle}/service_account/{service_account_identifier} | Update an organization service account
[**Update_0**](ServiceAccounts.md#Update_0) | **Patch** /service_account/{service_account_identifier} | Update a tenant service account



## Create

> User Create(ctx, orgHandle).Body(body).Execute()

Create a service account at the organization level



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
    orgHandle := "orgHandle_example" // string | The handle of the organization where the service account will be created
    body := *openapiclient.NewCreateServiceAccountUserRequest() // CreateServiceAccountUserRequest | The request body containing service account details

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAccounts.Create(context.Background(), orgHandle).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccounts.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: User
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccounts.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization where the service account will be created | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateServiceAccountUserRequest**](CreateServiceAccountUserRequest.md) | The request body containing service account details | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create_0

> User Create_0(ctx).Body(body).Execute()

Create a tenant service account



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
    body := *openapiclient.NewCreateServiceAccountUserRequest() // CreateServiceAccountUserRequest | The request body containing service account details

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAccounts.Create_0(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccounts.Create_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create_0`: User
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccounts.Create_0`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreate_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateServiceAccountUserRequest**](CreateServiceAccountUserRequest.md) | The request body containing service account details | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Delete(ctx, orgHandle, serviceAccountIdentifier).Execute()

Delete an organization service account



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
    serviceAccountIdentifier := "serviceAccountIdentifier_example" // string | The ID of the service account to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAccounts.Delete(context.Background(), orgHandle, serviceAccountIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccounts.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization | 
**serviceAccountIdentifier** | **string** | The ID of the service account to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete_0

> Delete_0(ctx, serviceAccountIdentifier).Execute()

Delete a tenant service account



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
    serviceAccountIdentifier := "serviceAccountIdentifier_example" // string | The ID of the service account to delete

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAccounts.Delete_0(context.Background(), serviceAccountIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccounts.Delete_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountIdentifier** | **string** | The ID of the service account to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDelete_2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgCreate

> TokenWithValue OrgCreate(ctx, orgHandle, serviceAccountIdentifier).Body(body).Execute()

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
    resp, r, err := api_client.ServiceAccounts.OrgCreate(context.Background(), orgHandle, serviceAccountIdentifier).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccounts.OrgCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgCreate`: TokenWithValue
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccounts.OrgCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization | 
**serviceAccountIdentifier** | **string** | The ID of the service account | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgCreateRequest struct via the builder pattern


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


## OrgDelete

> Token OrgDelete(ctx, orgHandle, serviceAccountIdentifier, tokenId).Execute()

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
    resp, r, err := api_client.ServiceAccounts.OrgDelete(context.Background(), orgHandle, serviceAccountIdentifier, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccounts.OrgDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgDelete`: Token
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccounts.OrgDelete`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiOrgDeleteRequest struct via the builder pattern


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


## OrgGet

> User OrgGet(ctx, orgHandle, serviceAccountIdentifier).Execute()

Get organization service account



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAccounts.OrgGet(context.Background(), orgHandle, serviceAccountIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccounts.OrgGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgGet`: User
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccounts.OrgGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization | 
**serviceAccountIdentifier** | **string** | The ID of the service account | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgGet_0

> Token OrgGet_0(ctx, orgHandle, serviceAccountIdentifier, tokenId).Execute()

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
    resp, r, err := api_client.ServiceAccounts.OrgGet_0(context.Background(), orgHandle, serviceAccountIdentifier, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccounts.OrgGet_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgGet_0`: Token
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccounts.OrgGet_0`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiOrgGet_3Request struct via the builder pattern


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


## OrgList

> ListUsersResponse OrgList(ctx, orgHandle).Limit(limit).NextToken(nextToken).Execute()

List organization service accounts



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
    limit := int32(56) // int32 | Maximum number of items to return. Default 100. Max 1000. (optional)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAccounts.OrgList(context.Background(), orgHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccounts.OrgList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgList`: ListUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccounts.OrgList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of items to return. Default 100. Max 1000. | 
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListUsersResponse**](ListUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgListAuditLogs

> ListAuditLogsResponse OrgListAuditLogs(ctx, orgHandle, serviceAccountIdentifier).Limit(limit).NextToken(nextToken).Execute()

List audit logs for an org service account



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
    resp, r, err := api_client.ServiceAccounts.OrgListAuditLogs(context.Background(), orgHandle, serviceAccountIdentifier).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccounts.OrgListAuditLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgListAuditLogs`: ListAuditLogsResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccounts.OrgListAuditLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization | 
**serviceAccountIdentifier** | **string** | The ID of the service account | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgListAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListAuditLogsResponse**](ListAuditLogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgList_0

> ListTokensResponse OrgList_0(ctx, orgHandle, serviceAccountIdentifier).Limit(limit).NextToken(nextToken).Execute()

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
    resp, r, err := api_client.ServiceAccounts.OrgList_0(context.Background(), orgHandle, serviceAccountIdentifier).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccounts.OrgList_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgList_0`: ListTokensResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccounts.OrgList_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization | 
**serviceAccountIdentifier** | **string** | The ID of the service account | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgList_4Request struct via the builder pattern


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


## OrgUpdate

> Token OrgUpdate(ctx, orgHandle, serviceAccountIdentifier, tokenId).Body(body).Execute()

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
    resp, r, err := api_client.ServiceAccounts.OrgUpdate(context.Background(), orgHandle, serviceAccountIdentifier, tokenId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccounts.OrgUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgUpdate`: Token
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccounts.OrgUpdate`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiOrgUpdateRequest struct via the builder pattern


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


## TenantCreate

> TokenWithValue TenantCreate(ctx, serviceAccountIdentifier).Body(body).Execute()

Create a tenant service account token



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
    serviceAccountIdentifier := "serviceAccountIdentifier_example" // string | The ID of the service account
    body := *openapiclient.NewCreateTokenRequest() // CreateTokenRequest | The request body containing token details

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAccounts.TenantCreate(context.Background(), serviceAccountIdentifier).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccounts.TenantCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantCreate`: TokenWithValue
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccounts.TenantCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountIdentifier** | **string** | The ID of the service account | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantCreateRequest struct via the builder pattern


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


## TenantDelete

> Token TenantDelete(ctx, serviceAccountIdentifier, tokenId).Execute()

Delete a tenant service account token



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
    serviceAccountIdentifier := "serviceAccountIdentifier_example" // string | The ID of the service account
    tokenId := "tokenId_example" // string | The identifier of the token

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAccounts.TenantDelete(context.Background(), serviceAccountIdentifier, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccounts.TenantDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantDelete`: Token
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccounts.TenantDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountIdentifier** | **string** | The ID of the service account | 
**tokenId** | **string** | The identifier of the token | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantDeleteRequest struct via the builder pattern


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


## TenantGet

> User TenantGet(ctx, serviceAccountIdentifier).Execute()

Get tenant service account



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
    serviceAccountIdentifier := "serviceAccountIdentifier_example" // string | The ID of the service account

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAccounts.TenantGet(context.Background(), serviceAccountIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccounts.TenantGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantGet`: User
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccounts.TenantGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountIdentifier** | **string** | The ID of the service account | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantGet_0

> Token TenantGet_0(ctx, serviceAccountIdentifier, tokenId).Execute()

Get a tenant service account token



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
    serviceAccountIdentifier := "serviceAccountIdentifier_example" // string | The ID of the service account
    tokenId := "tokenId_example" // string | The identifier of the token

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAccounts.TenantGet_0(context.Background(), serviceAccountIdentifier, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccounts.TenantGet_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantGet_0`: Token
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccounts.TenantGet_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountIdentifier** | **string** | The ID of the service account | 
**tokenId** | **string** | The identifier of the token | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantGet_5Request struct via the builder pattern


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


## TenantList

> ListUsersResponse TenantList(ctx).Limit(limit).NextToken(nextToken).Execute()

List tenant service accounts



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
    limit := int32(56) // int32 | Maximum number of items to return. Default 100. Max 1000. (optional)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAccounts.TenantList(context.Background()).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccounts.TenantList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantList`: ListUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccounts.TenantList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items to return. Default 100. Max 1000. | 
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListUsersResponse**](ListUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantListAuditLogs

> ListAuditLogsResponse TenantListAuditLogs(ctx, serviceAccountIdentifier).Limit(limit).NextToken(nextToken).Execute()

List audit logs for a tenant service account



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
    serviceAccountIdentifier := "serviceAccountIdentifier_example" // string | The ID of the service account
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAccounts.TenantListAuditLogs(context.Background(), serviceAccountIdentifier).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccounts.TenantListAuditLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantListAuditLogs`: ListAuditLogsResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccounts.TenantListAuditLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountIdentifier** | **string** | The ID of the service account | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantListAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListAuditLogsResponse**](ListAuditLogsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantList_0

> ListTokensResponse TenantList_0(ctx, serviceAccountIdentifier).Limit(limit).NextToken(nextToken).Execute()

List tenant service account tokens



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
    serviceAccountIdentifier := "serviceAccountIdentifier_example" // string | The ID of the service account
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAccounts.TenantList_0(context.Background(), serviceAccountIdentifier).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccounts.TenantList_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantList_0`: ListTokensResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccounts.TenantList_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountIdentifier** | **string** | The ID of the service account | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantList_6Request struct via the builder pattern


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


## TenantUpdate

> Token TenantUpdate(ctx, serviceAccountIdentifier, tokenId).Body(body).Execute()

Update a tenant service account token



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
    serviceAccountIdentifier := "serviceAccountIdentifier_example" // string | The ID of the service account
    tokenId := "tokenId_example" // string | The identifier of the token
    body := *openapiclient.NewUpdateTokenRequest("Status_example") // UpdateTokenRequest | The request body containing the updated token details

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAccounts.TenantUpdate(context.Background(), serviceAccountIdentifier, tokenId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccounts.TenantUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantUpdate`: Token
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccounts.TenantUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountIdentifier** | **string** | The ID of the service account | 
**tokenId** | **string** | The identifier of the token | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantUpdateRequest struct via the builder pattern


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


## Update

> User Update(ctx, orgHandle, serviceAccountIdentifier).Body(body).Execute()

Update an organization service account



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
    serviceAccountIdentifier := "serviceAccountIdentifier_example" // string | The ID of the service account to update
    body := *openapiclient.NewUpdateServiceAccountUserRequest() // UpdateServiceAccountUserRequest | The request body containing the updated service account details

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAccounts.Update(context.Background(), orgHandle, serviceAccountIdentifier).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccounts.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: User
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccounts.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization | 
**serviceAccountIdentifier** | **string** | The ID of the service account to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateServiceAccountUserRequest**](UpdateServiceAccountUserRequest.md) | The request body containing the updated service account details | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update_0

> User Update_0(ctx, serviceAccountIdentifier).Body(body).Execute()

Update a tenant service account



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
    serviceAccountIdentifier := "serviceAccountIdentifier_example" // string | The ID of the service account to update
    body := *openapiclient.NewUpdateServiceAccountUserRequest() // UpdateServiceAccountUserRequest | The request body containing the updated service account details

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAccounts.Update_0(context.Background(), serviceAccountIdentifier).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAccounts.Update_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update_0`: User
    fmt.Fprintf(os.Stdout, "Response from `ServiceAccounts.Update_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceAccountIdentifier** | **string** | The ID of the service account to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdate_7Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateServiceAccountUserRequest**](UpdateServiceAccountUserRequest.md) | The request body containing the updated service account details | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

