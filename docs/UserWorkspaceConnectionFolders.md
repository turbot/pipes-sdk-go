# \UserWorkspaceConnectionFolders

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](UserWorkspaceConnectionFolders.md#Create) | **Post** /user/{user_handle}/workspace/{workspace_handle}/connection_folder | Create user workspace connection folder
[**Delete**](UserWorkspaceConnectionFolders.md#Delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/connection_folder/{folder_id} | Delete user workspace connection folder
[**Get**](UserWorkspaceConnectionFolders.md#Get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/connection_folder/{folder_id} | Get user workspace connection folder
[**List**](UserWorkspaceConnectionFolders.md#List) | **Get** /user/{user_handle}/workspace/{workspace_handle}/connection_folder | List user workspace connections folders
[**Update**](UserWorkspaceConnectionFolders.md#Update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/connection_folder/{folder_id} | Update user workspace connection folder



## Create

> Connection Create(ctx, userHandle, workspaceHandle).Request(request).Mode(mode).Execute()

Create user workspace connection folder



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
    userHandle := "userHandle_example" // string | The handle of the user where we want to create the connection folder.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the folder is to be created.
    request := *openapiclient.NewCreateConnectionFolderRequest("Title_example") // CreateConnectionFolderRequest | The request body for the connection to be created.
    mode := "mode_example" // string | The mode of this request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceConnectionFolders.Create(context.Background(), userHandle, workspaceHandle).Request(request).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceConnectionFolders.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Connection
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceConnectionFolders.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where we want to create the connection folder. | 
**workspaceHandle** | **string** | The handle of the workspace where the folder is to be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateConnectionFolderRequest**](CreateConnectionFolderRequest.md) | The request body for the connection to be created. | 
 **mode** | **string** | The mode of this request | 

### Return type

[**Connection**](Connection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Connection Delete(ctx, userHandle, workspaceHandle, folderId).Execute()

Delete user workspace connection folder



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
    userHandle := "userHandle_example" // string | The handle of the user where the folder to be deleted exists.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the folder to be deleted exists.
    folderId := "folderId_example" // string | The id of the connection folder to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceConnectionFolders.Delete(context.Background(), userHandle, workspaceHandle, folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceConnectionFolders.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: Connection
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceConnectionFolders.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where the folder to be deleted exists. | 
**workspaceHandle** | **string** | The handle of the workspace where the folder to be deleted exists. | 
**folderId** | **string** | The id of the connection folder to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Connection**](Connection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Connection Get(ctx, userHandle, workspaceHandle, folderId).Execute()

Get user workspace connection folder



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
    userHandle := "userHandle_example" // string | The handle of the user where the folder exists.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the folder exists.
    folderId := "folderId_example" // string | The id of the connection folder to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceConnectionFolders.Get(context.Background(), userHandle, workspaceHandle, folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceConnectionFolders.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Connection
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceConnectionFolders.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where the folder exists. | 
**workspaceHandle** | **string** | The handle of the workspace where the folder exists. | 
**folderId** | **string** | The id of the connection folder to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Connection**](Connection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListConnectionsResponse List(ctx, userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

List user workspace connections folders



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
    userHandle := "userHandle_example" // string | The handle of the user where we want to create the connection folder.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the folder is to be created.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceConnectionFolders.List(context.Background(), userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceConnectionFolders.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListConnectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceConnectionFolders.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where we want to create the connection folder. | 
**workspaceHandle** | **string** | The handle of the workspace where the folder is to be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListConnectionsResponse**](ListConnectionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Connection Update(ctx, userHandle, workspaceHandle, folderId).Request(request).Mode(mode).Execute()

Update user workspace connection folder



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
    userHandle := "userHandle_example" // string | The handle of the user where the folder to be updated exists.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the folder to be updated exists.
    folderId := "folderId_example" // string | The id of the connection folder to update.
    request := *openapiclient.NewUpdateConnectionFolderRequest() // UpdateConnectionFolderRequest | The request body for the connection which needs to be updated.
    mode := "mode_example" // string | The mode of this request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceConnectionFolders.Update(context.Background(), userHandle, workspaceHandle, folderId).Request(request).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceConnectionFolders.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Connection
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceConnectionFolders.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where the folder to be updated exists. | 
**workspaceHandle** | **string** | The handle of the workspace where the folder to be updated exists. | 
**folderId** | **string** | The id of the connection folder to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**UpdateConnectionFolderRequest**](UpdateConnectionFolderRequest.md) | The request body for the connection which needs to be updated. | 
 **mode** | **string** | The mode of this request | 

### Return type

[**Connection**](Connection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

