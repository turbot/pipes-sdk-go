# \UserWorkspaceDatatanks

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](UserWorkspaceDatatanks.md#Create) | **Post** /user/{user_handle}/workspace/{workspace_handle}/datatank | Create user workspace Datatank
[**Delete**](UserWorkspaceDatatanks.md#Delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle} | Delete user workspace Datatank table
[**Get**](UserWorkspaceDatatanks.md#Get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle} | Get user workspace Datatank
[**List**](UserWorkspaceDatatanks.md#List) | **Get** /user/{user_handle}/workspace/{workspace_handle}/datatank | List user workspace Datatank
[**Update**](UserWorkspaceDatatanks.md#Update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle} | Update user workspace Datatank



## Create

> Datatank Create(ctx, userHandle, workspaceHandle).Request(request).Execute()

Create user workspace Datatank



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
    userHandle := "userHandle_example" // string | The handle of the user where we want to create the workspace Datatank.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace.
    request := *openapiclient.NewCreateDatatankRequest("Handle_example") // CreateDatatankRequest | The request body for the workspace Datatank to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceDatatanks.Create(context.Background(), userHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceDatatanks.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Datatank
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceDatatanks.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where we want to create the workspace Datatank. | 
**workspaceHandle** | **string** | The handle of the workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateDatatankRequest**](CreateDatatankRequest.md) | The request body for the workspace Datatank to be created. | 

### Return type

[**Datatank**](Datatank.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Datatank Delete(ctx, userHandle, workspaceHandle, datatankHandle).Execute()

Delete user workspace Datatank table



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
    userHandle := "userHandle_example" // string | The handle of the user.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace.
    datatankHandle := "datatankHandle_example" // string | The name of the workspace Datatank to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceDatatanks.Delete(context.Background(), userHandle, workspaceHandle, datatankHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceDatatanks.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: Datatank
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceDatatanks.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user. | 
**workspaceHandle** | **string** | The handle of the workspace. | 
**datatankHandle** | **string** | The name of the workspace Datatank to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Datatank**](Datatank.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Datatank Get(ctx, userHandle, workspaceHandle, datatankHandle).Execute()

Get user workspace Datatank



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
    userHandle := "userHandle_example" // string | The handle of the user.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace.
    datatankHandle := "datatankHandle_example" // string | The name of the workspace Datatank.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceDatatanks.Get(context.Background(), userHandle, workspaceHandle, datatankHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceDatatanks.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Datatank
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceDatatanks.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user. | 
**workspaceHandle** | **string** | The handle of the workspace. | 
**datatankHandle** | **string** | The name of the workspace Datatank. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Datatank**](Datatank.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListDatatankResponse List(ctx, userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

List user workspace Datatank



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
    userHandle := "userHandle_example" // string | The handle of the user for which we want to list the workspace Datatank.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceDatatanks.List(context.Background(), userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceDatatanks.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListDatatankResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceDatatanks.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which we want to list the workspace Datatank. | 
**workspaceHandle** | **string** | The handle of the workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListDatatankResponse**](ListDatatankResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Datatank Update(ctx, userHandle, workspaceHandle, datatankHandle).Request(request).Execute()

Update user workspace Datatank



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
    userHandle := "userHandle_example" // string | The handle of the user.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace.
    datatankHandle := "datatankHandle_example" // string | The name of the workspace Datatank.
    request := *openapiclient.NewUpdateDatatankRequest() // UpdateDatatankRequest | The request body to update workspace Datatank.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceDatatanks.Update(context.Background(), userHandle, workspaceHandle, datatankHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceDatatanks.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Datatank
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceDatatanks.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user. | 
**workspaceHandle** | **string** | The handle of the workspace. | 
**datatankHandle** | **string** | The name of the workspace Datatank. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**UpdateDatatankRequest**](UpdateDatatankRequest.md) | The request body to update workspace Datatank. | 

### Return type

[**Datatank**](Datatank.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

