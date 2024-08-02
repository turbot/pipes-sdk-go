# \OrgWorkspaceConnectionFolders

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](OrgWorkspaceConnectionFolders.md#Create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/connection_folder | Create org workspace connection folder
[**Delete**](OrgWorkspaceConnectionFolders.md#Delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/connection_folder/{folder_id} | Delete org workspace connection folder
[**Get**](OrgWorkspaceConnectionFolders.md#Get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/connection_folder/{folder_id} | Get org workspace connection folder
[**List**](OrgWorkspaceConnectionFolders.md#List) | **Get** /org/{org_handle}/workspace/{workspace_handle}/connection_folder | List org workspace connection folders
[**Update**](OrgWorkspaceConnectionFolders.md#Update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/connection_folder/{folder_id} | Update org workspace connection folder



## Create

> Connection Create(ctx, orgHandle, workspaceHandle).Request(request).Mode(mode).Execute()

Create org workspace connection folder



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
    orgHandle := "orgHandle_example" // string | The handle of the org where we want to create the connection folder.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the folder is to be created.
    request := *openapiclient.NewCreateConnectionFolderRequest("Title_example") // CreateConnectionFolderRequest | The request body for the connection to be created.
    mode := "mode_example" // string | The mode of this request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceConnectionFolders.Create(context.Background(), orgHandle, workspaceHandle).Request(request).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceConnectionFolders.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Connection
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceConnectionFolders.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org where we want to create the connection folder. | 
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

> Connection Delete(ctx, orgHandle, workspaceHandle, folderId).Execute()

Delete org workspace connection folder



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
    orgHandle := "orgHandle_example" // string | The handle of the org where the folder to be deleted exists.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the folder to be deleted exists.
    folderId := "folderId_example" // string | The id of the connection folder to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceConnectionFolders.Delete(context.Background(), orgHandle, workspaceHandle, folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceConnectionFolders.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: Connection
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceConnectionFolders.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org where the folder to be deleted exists. | 
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

> Connection Get(ctx, orgHandle, workspaceHandle, folderId).Execute()

Get org workspace connection folder



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
    orgHandle := "orgHandle_example" // string | The handle of the org where the folder exists.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the folder exists.
    folderId := "folderId_example" // string | The id of the connection folder to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceConnectionFolders.Get(context.Background(), orgHandle, workspaceHandle, folderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceConnectionFolders.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Connection
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceConnectionFolders.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org where the folder exists. | 
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

> ListConnectionsResponse List(ctx, orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

List org workspace connection folders



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
    orgHandle := "orgHandle_example" // string | The handle of the org where we want to create the connection folder.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the folder is to be created.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceConnectionFolders.List(context.Background(), orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceConnectionFolders.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListConnectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceConnectionFolders.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org where we want to create the connection folder. | 
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

> Connection Update(ctx, orgHandle, workspaceHandle, folderId).Request(request).Mode(mode).Execute()

Update org workspace connection folder



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
    orgHandle := "orgHandle_example" // string | The handle of the org where the folder to be updated exists.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the folder to be updated exists.
    folderId := "folderId_example" // string | The id of the connection folder to be updated.
    request := *openapiclient.NewUpdateConnectionFolderRequest() // UpdateConnectionFolderRequest | The request body for the connection which needs to be updated.
    mode := "mode_example" // string | The mode of this request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceConnectionFolders.Update(context.Background(), orgHandle, workspaceHandle, folderId).Request(request).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceConnectionFolders.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Connection
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceConnectionFolders.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org where the folder to be updated exists. | 
**workspaceHandle** | **string** | The handle of the workspace where the folder to be updated exists. | 
**folderId** | **string** | The id of the connection folder to be updated. | 

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

