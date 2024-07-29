# \OrgConnectionFolders

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](OrgConnectionFolders.md#Create) | **Post** /org/{org_handle}/connection_folder | Create org connection folder
[**CreatePermission**](OrgConnectionFolders.md#CreatePermission) | **Post** /org/{org_handle}/connection_folder/{folder_id}/permission | Create org connection folder permission
[**Delete**](OrgConnectionFolders.md#Delete) | **Delete** /org/{org_handle}/connection_folder/{folder_id} | Delete org connection folder
[**DeletePermission**](OrgConnectionFolders.md#DeletePermission) | **Delete** /org/{org_handle}/connection_folder/{folder_id}/permission/{permission_id} | Delete org connection folder permission
[**Get**](OrgConnectionFolders.md#Get) | **Get** /org/{org_handle}/connection_folder/{folder_id} | Get org connection folder
[**GetPermission**](OrgConnectionFolders.md#GetPermission) | **Get** /org/{org_handle}/connection_folder/{folder_id}/permission/{permission_id} | Get org connection folder permission
[**List**](OrgConnectionFolders.md#List) | **Get** /org/{org_handle}/connection_folder | List org connection folders
[**ListPermission**](OrgConnectionFolders.md#ListPermission) | **Get** /org/{org_handle}/connection_folder/{folder_id}/permission | List org connection folder permissions
[**ListWorkspaces**](OrgConnectionFolders.md#ListWorkspaces) | **Get** /org/{org_handle}/connection_folder/{folder_id}/workspace | List org connection folder workspaces
[**Update**](OrgConnectionFolders.md#Update) | **Patch** /org/{org_handle}/connection_folder/{folder_id} | Update org connection folder
[**UpdatePermission**](OrgConnectionFolders.md#UpdatePermission) | **Patch** /org/{org_handle}/connection_folder/{folder_id}/permission/{permission_id} | Update org connection folder permission



## Create

> Connection Create(ctx, orgHandle).Request(request).Mode(mode).Execute()

Create org connection folder



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
    orgHandle := "orgHandle_example" // string | The handle of the org for which we want to create a connection folder.
    request := *openapiclient.NewCreateConnectionFolderRequest("Title_example") // CreateConnectionFolderRequest | The request body for the connection folder to be created.
    mode := "mode_example" // string | The mode of this request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnectionFolders.Create(context.Background(), orgHandle).Request(request).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnectionFolders.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Connection
    fmt.Fprintf(os.Stdout, "Response from `OrgConnectionFolders.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which we want to create a connection folder. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**CreateConnectionFolderRequest**](CreateConnectionFolderRequest.md) | The request body for the connection folder to be created. | 
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

> Permission CreatePermission(ctx, orgHandle, folderId).Request(request).Execute()

Create org connection folder permission



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
    orgHandle := "orgHandle_example" // string | The handle of the org for which we want to create connection folder permissions.
    folderId := "folderId_example" // string | The ID of the connection folder for which we want to create permissions.
    request := *openapiclient.NewCreatePermissionRequest() // CreatePermissionRequest | The request body for the connection folder permission to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnectionFolders.CreatePermission(context.Background(), orgHandle, folderId).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnectionFolders.CreatePermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePermission`: Permission
    fmt.Fprintf(os.Stdout, "Response from `OrgConnectionFolders.CreatePermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which we want to create connection folder permissions. | 
**folderId** | **string** | The ID of the connection folder for which we want to create permissions. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreatePermissionRequest**](CreatePermissionRequest.md) | The request body for the connection folder permission to be created. | 

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

> Connection Delete(ctx, orgHandle, folderId).Execute()

Delete org connection folder



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
    orgHandle := "orgHandle_example" // string | The handle of the org for which we want to delete the connection folder.
    folderId := "folderId_example" // string | The ID of the connection folder that needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnectionFolders.Delete(context.Background(), orgHandle, folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnectionFolders.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: Connection
    fmt.Fprintf(os.Stdout, "Response from `OrgConnectionFolders.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which we want to delete the connection folder. | 
**folderId** | **string** | The ID of the connection folder that needs to be deleted. | 

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

> Permission DeletePermission(ctx, orgHandle, folderId, permissionId).Execute()

Delete org connection folder permission



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
    orgHandle := "orgHandle_example" // string | The handle of the org for which we want to delete connection folder permissions.
    folderId := "folderId_example" // string | The ID of the connection folder for which we want to delete permission.
    permissionId := "permissionId_example" // string | The id of the permission which we want to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnectionFolders.DeletePermission(context.Background(), orgHandle, folderId, permissionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnectionFolders.DeletePermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePermission`: Permission
    fmt.Fprintf(os.Stdout, "Response from `OrgConnectionFolders.DeletePermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which we want to delete connection folder permissions. | 
**folderId** | **string** | The ID of the connection folder for which we want to delete permission. | 
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

> Connection Get(ctx, orgHandle, folderId).Execute()

Get org connection folder



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
    orgHandle := "orgHandle_example" // string | The handle of the org for which we want to retrieve the connection folder.
    folderId := "folderId_example" // string | The ID of the connection folder to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnectionFolders.Get(context.Background(), orgHandle, folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnectionFolders.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Connection
    fmt.Fprintf(os.Stdout, "Response from `OrgConnectionFolders.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which we want to retrieve the connection folder. | 
**folderId** | **string** | The ID of the connection folder to retrieve. | 

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

> Permission GetPermission(ctx, orgHandle, folderId, permissionId).Execute()

Get org connection folder permission



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
    orgHandle := "orgHandle_example" // string | The handle of the org for which we want to get connection folder permissions.
    folderId := "folderId_example" // string | The ID of the connection folder for which we want to get permission.
    permissionId := "permissionId_example" // string | The id of the permission which we want to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnectionFolders.GetPermission(context.Background(), orgHandle, folderId, permissionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnectionFolders.GetPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPermission`: Permission
    fmt.Fprintf(os.Stdout, "Response from `OrgConnectionFolders.GetPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which we want to get connection folder permissions. | 
**folderId** | **string** | The ID of the connection folder for which we want to get permission. | 
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

> ListConnectionsResponse List(ctx, orgHandle).Limit(limit).NextToken(nextToken).Execute()

List org connection folders



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
    orgHandle := "orgHandle_example" // string | The handle of the org for which we want to list connection folders.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnectionFolders.List(context.Background(), orgHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnectionFolders.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListConnectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgConnectionFolders.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which we want to list connection folders. | 

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

> ListPermissionsResponse ListPermission(ctx, orgHandle, folderId).Limit(limit).NextToken(nextToken).Execute()

List org connection folder permissions



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
    orgHandle := "orgHandle_example" // string | The handle of the org for which we want to list connection folder permissions.
    folderId := "folderId_example" // string | The ID of the connection folder for which we want to list permissions.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnectionFolders.ListPermission(context.Background(), orgHandle, folderId).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnectionFolders.ListPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPermission`: ListPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgConnectionFolders.ListPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which we want to list connection folder permissions. | 
**folderId** | **string** | The ID of the connection folder for which we want to list permissions. | 

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

> ListWorkspaceConnectionAssociationsResponse ListWorkspaces(ctx, orgHandle, folderId).Limit(limit).NextToken(nextToken).Execute()

List org connection folder workspaces



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
    orgHandle := "orgHandle_example" // string | The handle of the org for which we want to list connection folder workspace.
    folderId := "folderId_example" // string | The ID of the connection folder for which we want to list workspaces.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnectionFolders.ListWorkspaces(context.Background(), orgHandle, folderId).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnectionFolders.ListWorkspaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkspaces`: ListWorkspaceConnectionAssociationsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgConnectionFolders.ListWorkspaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which we want to list connection folder workspace. | 
**folderId** | **string** | The ID of the connection folder for which we want to list workspaces. | 

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


## Update

> Connection Update(ctx, orgHandle, folderId).Request(request).Mode(mode).Execute()

Update org connection folder



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
    orgHandle := "orgHandle_example" // string | The handle of the org for which we want to update the connection folder.
    folderId := "folderId_example" // string | The ID of the connection folder which needs to be updated.
    request := *openapiclient.NewUpdateConnectionFolderRequest() // UpdateConnectionFolderRequest | The request body for the connection which needs to be updated.
    mode := "mode_example" // string | The mode of this request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnectionFolders.Update(context.Background(), orgHandle, folderId).Request(request).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnectionFolders.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Connection
    fmt.Fprintf(os.Stdout, "Response from `OrgConnectionFolders.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which we want to update the connection folder. | 
**folderId** | **string** | The ID of the connection folder which needs to be updated. | 

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


## UpdatePermission

> Permission UpdatePermission(ctx, orgHandle, folderId, permissionId).Request(request).Execute()

Update org connection folder permission



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
    orgHandle := "orgHandle_example" // string | The handle of the org for which we want to update connection permissions.
    folderId := "folderId_example" // string | The ID of the connection folder for which we want to update permission.
    permissionId := "permissionId_example" // string | The id of the permission which we want to update.
    request := *openapiclient.NewUpdatePermissionRequest() // UpdatePermissionRequest | The request body for the connection permission to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgConnectionFolders.UpdatePermission(context.Background(), orgHandle, folderId, permissionId).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgConnectionFolders.UpdatePermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePermission`: Permission
    fmt.Fprintf(os.Stdout, "Response from `OrgConnectionFolders.UpdatePermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which we want to update connection permissions. | 
**folderId** | **string** | The ID of the connection folder for which we want to update permission. | 
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

