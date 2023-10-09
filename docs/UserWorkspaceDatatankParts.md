# \UserWorkspaceDatatankParts

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Command**](UserWorkspaceDatatankParts.md#Command) | **Post** /user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name}/part/{datatank_part_id}/command | Run user workspace Datatank table partition command
[**Get**](UserWorkspaceDatatankParts.md#Get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name}/table/{datatank_part_id} | Get user workspace Datatank table partition
[**List**](UserWorkspaceDatatankParts.md#List) | **Get** /user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name}/part | List user workspace Datatank partitions
[**Update**](UserWorkspaceDatatankParts.md#Update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name}/part/{datatank_part_id} | Update user workspace Datatank table partition



## Command

> DatatankPart Command(ctx, userHandle, workspaceHandle, datatankHandle, datatankTableName, datatankPartId).Request(request).Execute()

Run user workspace Datatank table partition command



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
    datatankTableName := "datatankTableName_example" // string | The name of the workspace Datatank table.
    datatankPartId := "datatankPartId_example" // string | The id of the workspace Datatank table partition.
    request := *openapiclient.NewCmdDatatankPartRequest("Command_example") // CmdDatatankPartRequest | The request body for the datatank part command.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceDatatankParts.Command(context.Background(), userHandle, workspaceHandle, datatankHandle, datatankTableName, datatankPartId).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceDatatankParts.Command``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Command`: DatatankPart
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceDatatankParts.Command`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user. | 
**workspaceHandle** | **string** | The handle of the workspace. | 
**datatankHandle** | **string** | The name of the workspace Datatank. | 
**datatankTableName** | **string** | The name of the workspace Datatank table. | 
**datatankPartId** | **string** | The id of the workspace Datatank table partition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **request** | [**CmdDatatankPartRequest**](CmdDatatankPartRequest.md) | The request body for the datatank part command. | 

### Return type

[**DatatankPart**](DatatankPart.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> DatatankPart Get(ctx, userHandle, workspaceHandle, datatankHandle, datatankTableName, datatankPartId).Execute()

Get user workspace Datatank table partition



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
    datatankTableName := "datatankTableName_example" // string | The name of the workspace Datatank table.
    datatankPartId := "datatankPartId_example" // string | The id of the workspace Datatank table partition.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceDatatankParts.Get(context.Background(), userHandle, workspaceHandle, datatankHandle, datatankTableName, datatankPartId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceDatatankParts.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: DatatankPart
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceDatatankParts.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user. | 
**workspaceHandle** | **string** | The handle of the workspace. | 
**datatankHandle** | **string** | The name of the workspace Datatank. | 
**datatankTableName** | **string** | The name of the workspace Datatank table. | 
**datatankPartId** | **string** | The id of the workspace Datatank table partition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**DatatankPart**](DatatankPart.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListDatatankPartResponse List(ctx, userHandle, workspaceHandle, datatankHandle, datatankTableName).Limit(limit).NextToken(nextToken).Execute()

List user workspace Datatank partitions



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
    datatankHandle := "datatankHandle_example" // string | The name of the Datatank.
    datatankTableName := "datatankTableName_example" // string | The name of the workspace Datatank table.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceDatatankParts.List(context.Background(), userHandle, workspaceHandle, datatankHandle, datatankTableName).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceDatatankParts.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListDatatankPartResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceDatatankParts.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which we want to list the workspace Datatank. | 
**workspaceHandle** | **string** | The handle of the workspace. | 
**datatankHandle** | **string** | The name of the Datatank. | 
**datatankTableName** | **string** | The name of the workspace Datatank table. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListDatatankPartResponse**](ListDatatankPartResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> DatatankPart Update(ctx, userHandle, workspaceHandle, datatankHandle, datatankTableName, datatankPartId).Request(request).Execute()

Update user workspace Datatank table partition



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
    datatankTableName := "datatankTableName_example" // string | The name of the workspace Datatank table.
    datatankPartId := "datatankPartId_example" // string | The id of the workspace Datatank table partition to be updated.
    request := *openapiclient.NewUpdateDatatankPartRequest("DesiredState_example") // UpdateDatatankPartRequest | The request body to update workspace Datatank table partition.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceDatatankParts.Update(context.Background(), userHandle, workspaceHandle, datatankHandle, datatankTableName, datatankPartId).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceDatatankParts.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: DatatankPart
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceDatatankParts.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user. | 
**workspaceHandle** | **string** | The handle of the workspace. | 
**datatankHandle** | **string** | The name of the workspace Datatank. | 
**datatankTableName** | **string** | The name of the workspace Datatank table. | 
**datatankPartId** | **string** | The id of the workspace Datatank table partition to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **request** | [**UpdateDatatankPartRequest**](UpdateDatatankPartRequest.md) | The request body to update workspace Datatank table partition. | 

### Return type

[**DatatankPart**](DatatankPart.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

