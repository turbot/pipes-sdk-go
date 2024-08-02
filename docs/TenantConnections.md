# \TenantConnections

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](TenantConnections.md#Create) | **Post** /connection | Create tenant connection
[**CreatePermission**](TenantConnections.md#CreatePermission) | **Post** /connection/{connection_handle}/permission | Create tenant connection permission
[**Delete**](TenantConnections.md#Delete) | **Delete** /connection/{connection_handle} | Delete tenant connection
[**DeletePermission**](TenantConnections.md#DeletePermission) | **Delete** /connection/{connection_handle}/permission/{permission_id} | Delete tenant connection permission
[**Get**](TenantConnections.md#Get) | **Get** /connection/{connection_handle} | Get tenant connection
[**GetPermission**](TenantConnections.md#GetPermission) | **Get** /connection/{connection_handle}/permission/{permission_id} | Get tenant connection permission
[**List**](TenantConnections.md#List) | **Get** /connection | List tenant connections
[**ListPermission**](TenantConnections.md#ListPermission) | **Get** /connection/{connection_handle}/permission | List tenant connection permissions
[**ListWorkspaces**](TenantConnections.md#ListWorkspaces) | **Get** /connection/{connection_handle}/workspace | List tenant connection workspaces
[**Test**](TenantConnections.md#Test) | **Post** /connection/{connection_handle}/test | Test tenant connection
[**Update**](TenantConnections.md#Update) | **Patch** /connection/{connection_handle} | Update tenant connection
[**UpdatePermission**](TenantConnections.md#UpdatePermission) | **Patch** /connection/{connection_handle}/permission/{permission_id} | Update tenant connection permission



## Create

> Connection Create(ctx).Request(request).Mode(mode).Execute()

Create tenant connection



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
    request := *openapiclient.NewCreateConnectionRequest("Handle_example", "Plugin_example") // CreateConnectionRequest | The request body for the connection to be created.
    mode := "mode_example" // string | The mode of this request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantConnections.Create(context.Background()).Request(request).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantConnections.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Connection
    fmt.Fprintf(os.Stdout, "Response from `TenantConnections.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**CreateConnectionRequest**](CreateConnectionRequest.md) | The request body for the connection to be created. | 
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


## CreatePermission

> Permission CreatePermission(ctx, connectionHandle).Request(request).Execute()

Create tenant connection permission



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
    connectionHandle := "connectionHandle_example" // string | The handle of the connection for which we want to list permissions.
    request := *openapiclient.NewCreatePermissionRequest() // CreatePermissionRequest | The request body for the connection permission to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantConnections.CreatePermission(context.Background(), connectionHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantConnections.CreatePermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePermission`: Permission
    fmt.Fprintf(os.Stdout, "Response from `TenantConnections.CreatePermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionHandle** | **string** | The handle of the connection for which we want to list permissions. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**CreatePermissionRequest**](CreatePermissionRequest.md) | The request body for the connection permission to be created. | 

### Return type

[**Permission**](Permission.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Connection Delete(ctx, connectionHandle).Execute()

Delete tenant connection



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
    connectionHandle := "connectionHandle_example" // string | The handle of the connection that needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantConnections.Delete(context.Background(), connectionHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantConnections.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: Connection
    fmt.Fprintf(os.Stdout, "Response from `TenantConnections.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionHandle** | **string** | The handle of the connection that needs to be deleted. | 

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


## DeletePermission

> Permission DeletePermission(ctx, connectionHandle, permissionId).Execute()

Delete tenant connection permission



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
    connectionHandle := "connectionHandle_example" // string | The handle of the connection for which we want to delete permission for.
    permissionId := "permissionId_example" // string | The id of the permission which we want to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantConnections.DeletePermission(context.Background(), connectionHandle, permissionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantConnections.DeletePermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePermission`: Permission
    fmt.Fprintf(os.Stdout, "Response from `TenantConnections.DeletePermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionHandle** | **string** | The handle of the connection for which we want to delete permission for. | 
**permissionId** | **string** | The id of the permission which we want to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Permission**](Permission.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Connection Get(ctx, connectionHandle).Execute()

Get tenant connection



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
    connectionHandle := "connectionHandle_example" // string | The handle of the connection to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantConnections.Get(context.Background(), connectionHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantConnections.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Connection
    fmt.Fprintf(os.Stdout, "Response from `TenantConnections.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionHandle** | **string** | The handle of the connection to retrieve. | 

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


## GetPermission

> Permission GetPermission(ctx, connectionHandle, permissionId).Execute()

Get tenant connection permission



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
    connectionHandle := "connectionHandle_example" // string | The handle of the connection for which we want to get permission for.
    permissionId := "permissionId_example" // string | The id of the permission which we want to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantConnections.GetPermission(context.Background(), connectionHandle, permissionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantConnections.GetPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPermission`: Permission
    fmt.Fprintf(os.Stdout, "Response from `TenantConnections.GetPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionHandle** | **string** | The handle of the connection for which we want to get permission for. | 
**permissionId** | **string** | The id of the permission which we want to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Permission**](Permission.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListConnectionsResponse List(ctx).Limit(limit).NextToken(nextToken).Execute()

List tenant connections



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
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantConnections.List(context.Background()).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantConnections.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListConnectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `TenantConnections.List`: %v\n", resp)
}
```

### Path Parameters



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


## ListPermission

> ListPermissionsResponse ListPermission(ctx, connectionHandle).Limit(limit).NextToken(nextToken).Execute()

List tenant connection permissions



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
    connectionHandle := "connectionHandle_example" // string | The handle of the connection for which we want to list permissions.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantConnections.ListPermission(context.Background(), connectionHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantConnections.ListPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPermission`: ListPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `TenantConnections.ListPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionHandle** | **string** | The handle of the connection for which we want to list permissions. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListPermissionsResponse**](ListPermissionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspaces

> ListWorkspaceConnectionAssociationsResponse ListWorkspaces(ctx, connectionHandle).Limit(limit).NextToken(nextToken).Execute()

List tenant connection workspaces



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
    connectionHandle := "connectionHandle_example" // string | The handle of the connection for which we want to list workspaces.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantConnections.ListWorkspaces(context.Background(), connectionHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantConnections.ListWorkspaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkspaces`: ListWorkspaceConnectionAssociationsResponse
    fmt.Fprintf(os.Stdout, "Response from `TenantConnections.ListWorkspaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionHandle** | **string** | The handle of the connection for which we want to list workspaces. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListWorkspaceConnectionAssociationsResponse**](ListWorkspaceConnectionAssociationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Test

> ConnectionTestResult Test(ctx, connectionHandle).Request(request).Execute()

Test tenant connection



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
    connectionHandle := "connectionHandle_example" // string | The handle of the connection to be tested. For connections that are not yet created, use underscore `_` as the handle, else pass the handle of the existing connection.
    request := *openapiclient.NewTestConnectionRequest("Plugin_example") // TestConnectionRequest | The request body for the connection to be tested.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantConnections.Test(context.Background(), connectionHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantConnections.Test``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Test`: ConnectionTestResult
    fmt.Fprintf(os.Stdout, "Response from `TenantConnections.Test`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

> Connection Update(ctx, connectionHandle).Request(request).Mode(mode).Execute()

Update tenant connection



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
    connectionHandle := "connectionHandle_example" // string | The handle of the connection which needs to be updated.
    request := *openapiclient.NewUpdateConnectionRequest() // UpdateConnectionRequest | The request body for the connection which needs to be updated.
    mode := "mode_example" // string | The mode of this request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantConnections.Update(context.Background(), connectionHandle).Request(request).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantConnections.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Connection
    fmt.Fprintf(os.Stdout, "Response from `TenantConnections.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionHandle** | **string** | The handle of the connection which needs to be updated. | 

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


## UpdatePermission

> Permission UpdatePermission(ctx, connectionHandle, permissionId).Request(request).Execute()

Update tenant connection permission



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
    connectionHandle := "connectionHandle_example" // string | The handle of the connection for which we want to update permission for.
    permissionId := "permissionId_example" // string | The id of the permission which we want to update.
    request := *openapiclient.NewUpdatePermissionRequest() // UpdatePermissionRequest | The request body for the connection permission to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantConnections.UpdatePermission(context.Background(), connectionHandle, permissionId).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantConnections.UpdatePermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePermission`: Permission
    fmt.Fprintf(os.Stdout, "Response from `TenantConnections.UpdatePermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectionHandle** | **string** | The handle of the connection for which we want to update permission for. | 
**permissionId** | **string** | The id of the permission which we want to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**UpdatePermissionRequest**](UpdatePermissionRequest.md) | The request body for the connection permission to be updated. | 

### Return type

[**Permission**](Permission.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

