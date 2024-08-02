# \UserWorkspaceConnections

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](UserWorkspaceConnections.md#Create) | **Post** /user/{user_handle}/workspace/{workspace_handle}/connection | Create a connection for a user workspace
[**Delete**](UserWorkspaceConnections.md#Delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/connection/{connection_handle} | Delete user workspace connection association
[**Get**](UserWorkspaceConnections.md#Get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/connection/{connection_handle} | Get user workspace connection association
[**List**](UserWorkspaceConnections.md#List) | **Get** /user/{user_handle}/workspace/{workspace_handle}/connection | List connections for a user workspace
[**Test**](UserWorkspaceConnections.md#Test) | **Post** /user/{user_handle}/workspace/{workspace_handle}/connection/{connection_handle}/test | Test user workspace connection
[**Update**](UserWorkspaceConnections.md#Update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/connection/{connection_handle} | Update user workspace connection



## Create

> Connection Create(ctx, userHandle, workspaceHandle).Request(request).Execute()

Create a connection for a user workspace



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
    userHandle := "userHandle_example" // string | The handle of the user to which the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the connection will be created.
    request := *openapiclient.NewCreateConnectionRequest("Handle_example", "Plugin_example") // CreateConnectionRequest | The request body for the connection to the created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceConnections.Create(context.Background(), userHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceConnections.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Connection
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceConnections.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user to which the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace where the connection will be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateConnectionRequest**](CreateConnectionRequest.md) | The request body for the connection to the created. | 

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

> WorkspaceConnection Delete(ctx, userHandle, workspaceHandle, connectionHandle).Execute()

Delete user workspace connection association



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
    userHandle := "userHandle_example" // string | The handle of the user performing the action.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace whose association needs to be deleted.
    connectionHandle := "connectionHandle_example" // string | The handle of the connection whose association needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceConnections.Delete(context.Background(), userHandle, workspaceHandle, connectionHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceConnections.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: WorkspaceConnection
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceConnections.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user performing the action. | 
**workspaceHandle** | **string** | The handle of the workspace whose association needs to be deleted. | 
**connectionHandle** | **string** | The handle of the connection whose association needs to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorkspaceConnection**](WorkspaceConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> WorkspaceConnection Get(ctx, userHandle, workspaceHandle, connectionHandle).Execute()

Get user workspace connection association



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
    userHandle := "userHandle_example" // string | The handle of the user for which you want to get the association.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace wherethe association exist.
    connectionHandle := "connectionHandle_example" // string | The handle of the connection whose association details needs to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceConnections.Get(context.Background(), userHandle, workspaceHandle, connectionHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceConnections.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: WorkspaceConnection
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceConnections.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which you want to get the association. | 
**workspaceHandle** | **string** | The handle of the workspace wherethe association exist. | 
**connectionHandle** | **string** | The handle of the connection whose association details needs to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorkspaceConnection**](WorkspaceConnection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListWorkspaceConnectionsResponse List(ctx, userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

List connections for a user workspace



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
    userHandle := "userHandle_example" // string | The handle of the user to which the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace for which we want to list connections.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceConnections.List(context.Background(), userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceConnections.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListWorkspaceConnectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceConnections.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user to which the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace for which we want to list connections. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListWorkspaceConnectionsResponse**](ListWorkspaceConnectionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Test

> ConnectionTestResult Test(ctx, userHandle, workspaceHandle, connectionHandle).Request(request).Execute()

Test user workspace connection



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
    userHandle := "userHandle_example" // string | The handle of the user where the connection exists / intends to be created.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the connection exists.
    connectionHandle := "connectionHandle_example" // string | The handle of the connection to be tested. For connections that are not yet created, use underscore `_` as the handle, else pass the handle of the existing connection.
    request := *openapiclient.NewTestConnectionRequest("Plugin_example") // TestConnectionRequest | The request body for the connection to be tested.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceConnections.Test(context.Background(), userHandle, workspaceHandle, connectionHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceConnections.Test``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Test`: ConnectionTestResult
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceConnections.Test`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where the connection exists / intends to be created. | 
**workspaceHandle** | **string** | The handle of the workspace where the connection exists. | 
**connectionHandle** | **string** | The handle of the connection to be tested. For connections that are not yet created, use underscore &#x60;_&#x60; as the handle, else pass the handle of the existing connection. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**TestConnectionRequest**](TestConnectionRequest.md) | The request body for the connection to be tested. | 

### Return type

[**ConnectionTestResult**](ConnectionTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Connection Update(ctx, userHandle, workspaceHandle, connectionHandle).Request(request).Mode(mode).Execute()

Update user workspace connection



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
    userHandle := "userHandle_example" // string | The handle of the user where the connection to be updated exists.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the connection to be updated exists.
    connectionHandle := "connectionHandle_example" // string | The handle of the connection to update.
    request := *openapiclient.NewUpdateConnectionRequest() // UpdateConnectionRequest | The request body for the connection which needs to be updated.
    mode := "mode_example" // string | The mode of this request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceConnections.Update(context.Background(), userHandle, workspaceHandle, connectionHandle).Request(request).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceConnections.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Connection
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceConnections.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user where the connection to be updated exists. | 
**workspaceHandle** | **string** | The handle of the workspace where the connection to be updated exists. | 
**connectionHandle** | **string** | The handle of the connection to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**UpdateConnectionRequest**](UpdateConnectionRequest.md) | The request body for the connection which needs to be updated. | 
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

