# \TenantMembers

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmInvite**](TenantMembers.md#ConfirmInvite) | **Get** /tenant/{tenant_handle}/member/invite/confirm | Confirm tenant member invite
[**Delete**](TenantMembers.md#Delete) | **Delete** /tenant/{tenant_handle}/member/{user_handle} | Delete tenant member
[**DeleteInvite**](TenantMembers.md#DeleteInvite) | **Delete** /tenant/{tenant_handle}/member/invite | Delete tenant member invite
[**Get**](TenantMembers.md#Get) | **Get** /tenant/{tenant_handle}/member/{user_handle} | Get tenant member
[**Invite**](TenantMembers.md#Invite) | **Post** /tenant/{tenant_handle}/member/invite | Invite tenant member
[**List**](TenantMembers.md#List) | **Get** /tenant/{tenant_handle}/member | List Tenant Members
[**Update**](TenantMembers.md#Update) | **Patch** /tenant/{tenant_handle}/member/{user_handle} | Update tenant member



## ConfirmInvite

> ConfirmInvite(ctx, tenantHandle).T(t).Execute()

Confirm tenant member invite



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
    tenantHandle := "tenantHandle_example" // string | Specify the handle of an tenant where the member needs to be invited.
    t := "t_example" // string | Specify the token.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantMembers.ConfirmInvite(context.Background(), tenantHandle).T(t).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantMembers.ConfirmInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantHandle** | **string** | Specify the handle of an tenant where the member needs to be invited. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t** | **string** | Specify the token. | 

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


## Delete

> TenantUser Delete(ctx, tenantHandle, userHandle).Execute()

Delete tenant member



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
    tenantHandle := "tenantHandle_example" // string | Specify the handle of the tenant where the member exists.
    userHandle := "userHandle_example" // string | Specify the handle of the user which need to be removed.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantMembers.Delete(context.Background(), tenantHandle, userHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantMembers.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: TenantUser
    fmt.Fprintf(os.Stdout, "Response from `TenantMembers.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantHandle** | **string** | Specify the handle of the tenant where the member exists. | 
**userHandle** | **string** | Specify the handle of the user which need to be removed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TenantUser**](TenantUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInvite

> DeleteInvite(ctx, tenantHandle).T(t).Execute()

Delete tenant member invite



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
    tenantHandle := "tenantHandle_example" // string | Specify the tenant handle.
    t := "t_example" // string | Specify the token to be rejected.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantMembers.DeleteInvite(context.Background(), tenantHandle).T(t).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantMembers.DeleteInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantHandle** | **string** | Specify the tenant handle. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t** | **string** | Specify the token to be rejected. | 

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


## Get

> TenantUser Get(ctx, tenantHandle, userHandle).Execute()

Get tenant member



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
    tenantHandle := "tenantHandle_example" // string | Specify the tenant handle where the member is associated.
    userHandle := "userHandle_example" // string | Specify the handle of the user whose information you want to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantMembers.Get(context.Background(), tenantHandle, userHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantMembers.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: TenantUser
    fmt.Fprintf(os.Stdout, "Response from `TenantMembers.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantHandle** | **string** | Specify the tenant handle where the member is associated. | 
**userHandle** | **string** | Specify the handle of the user whose information you want to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TenantUser**](TenantUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Invite

> TenantUser Invite(ctx, tenantHandle).Request(request).Execute()

Invite tenant member



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
    tenantHandle := "tenantHandle_example" // string | Specify the handle of an tenant where the member need to be invited.
    request := *openapiclient.NewInviteTenantUserRequest("Email_example", "Role_example") // InviteTenantUserRequest | The request body to invite a member to a tenant.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantMembers.Invite(context.Background(), tenantHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantMembers.Invite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Invite`: TenantUser
    fmt.Fprintf(os.Stdout, "Response from `TenantMembers.Invite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantHandle** | **string** | Specify the handle of an tenant where the member need to be invited. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**InviteTenantUserRequest**](InviteTenantUserRequest.md) | The request body to invite a member to a tenant. | 

### Return type

[**TenantUser**](TenantUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListTenantUsersResponse List(ctx, tenantHandle).Limit(limit).NextToken(nextToken).Execute()

List Tenant Members



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
    tenantHandle := "tenantHandle_example" // string | Specify the tenant handle.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantMembers.List(context.Background(), tenantHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantMembers.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListTenantUsersResponse
    fmt.Fprintf(os.Stdout, "Response from `TenantMembers.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantHandle** | **string** | Specify the tenant handle. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListTenantUsersResponse**](ListTenantUsersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> TenantUser Update(ctx, tenantHandle, userHandle).Request(request).Execute()

Update tenant member



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
    tenantHandle := "tenantHandle_example" // string | Specify the handle of the tenant where the member exists.
    userHandle := "userHandle_example" // string | Specify the handle of the user whose role need to be updated.
    request := *openapiclient.NewUpdateTenantUserRequest() // UpdateTenantUserRequest | The request body for the member.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantMembers.Update(context.Background(), tenantHandle, userHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantMembers.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: TenantUser
    fmt.Fprintf(os.Stdout, "Response from `TenantMembers.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantHandle** | **string** | Specify the handle of the tenant where the member exists. | 
**userHandle** | **string** | Specify the handle of the user whose role need to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**UpdateTenantUserRequest**](UpdateTenantUserRequest.md) | The request body for the member. | 

### Return type

[**TenantUser**](TenantUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

