# \OrgWorkspaceDatatankParts

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Command**](OrgWorkspaceDatatankParts.md#Command) | **Post** /org/{org_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name}/part/{datatank_part_id}/command | Run org workspace Datatank table partition command
[**Get**](OrgWorkspaceDatatankParts.md#Get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name}/table/{datatank_part_id} | Get org workspace Datatank table partition
[**List**](OrgWorkspaceDatatankParts.md#List) | **Get** /org/{org_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name}/part | List org workspace Datatank partitions
[**Update**](OrgWorkspaceDatatankParts.md#Update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name}/part/{datatank_part_id} | Update org workspace Datatank table partition



## Command

> DatatankPart Command(ctx, orgHandle, workspaceHandle, datatankHandle, datatankTableName, datatankPartId).Execute()

Run org workspace Datatank table partition command



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
    orgHandle := "orgHandle_example" // string | The handle of the org for which we want to list the workspace Datatank.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace.
    datatankHandle := "datatankHandle_example" // string | The name of the workspace Datatank.
    datatankTableName := "datatankTableName_example" // string | The name of the workspace Datatank table.
    datatankPartId := "datatankPartId_example" // string | The id of the workspace Datatank table partition.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceDatatankParts.Command(context.Background(), orgHandle, workspaceHandle, datatankHandle, datatankTableName, datatankPartId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceDatatankParts.Command``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Command`: DatatankPart
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceDatatankParts.Command`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which we want to list the workspace Datatank. | 
**workspaceHandle** | **string** | The handle of the workspace. | 
**datatankHandle** | **string** | The name of the workspace Datatank. | 
**datatankTableName** | **string** | The name of the workspace Datatank table. | 
**datatankPartId** | **string** | The id of the workspace Datatank table partition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommandRequest struct via the builder pattern


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


## Get

> DatatankPart Get(ctx, orgHandle, workspaceHandle, datatankHandle, datatankTableName, datatankPartId).Execute()

Get org workspace Datatank table partition



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
    orgHandle := "orgHandle_example" // string | The handle of the org.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace.
    datatankHandle := "datatankHandle_example" // string | The name of the workspace Datatank.
    datatankTableName := "datatankTableName_example" // string | The name of the workspace Datatank table.
    datatankPartId := "datatankPartId_example" // string | The id of the workspace Datatank table partition.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceDatatankParts.Get(context.Background(), orgHandle, workspaceHandle, datatankHandle, datatankTableName, datatankPartId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceDatatankParts.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: DatatankPart
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceDatatankParts.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org. | 
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

> ListDatatankPartResponse List(ctx, orgHandle, workspaceHandle, datatankHandle, datatankTableName).Limit(limit).NextToken(nextToken).Execute()

List org workspace Datatank partitions



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
    orgHandle := "orgHandle_example" // string | The handle of the org.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace.
    datatankHandle := "datatankHandle_example" // string | The name of the Datatank.
    datatankTableName := "datatankTableName_example" // string | The name of the workspace Datatank table.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceDatatankParts.List(context.Background(), orgHandle, workspaceHandle, datatankHandle, datatankTableName).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceDatatankParts.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListDatatankPartResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceDatatankParts.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org. | 
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

> DatatankPart Update(ctx, orgHandle, workspaceHandle, datatankHandle, datatankTableName, datatankPartId).Request(request).Execute()

Update org workspace Datatank table partition



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
    orgHandle := "orgHandle_example" // string | The handle of the org for which we want to list the workspace Datatank.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace.
    datatankHandle := "datatankHandle_example" // string | The name of the workspace Datatank.
    datatankTableName := "datatankTableName_example" // string | The name of the workspace Datatank table.
    datatankPartId := "datatankPartId_example" // string | The id of the workspace Datatank table partition to be updated.
    request := *openapiclient.NewUpdateDatatankPartRequest(map[string][]openapiclient.DesiredState{"key": "TODO"}) // UpdateDatatankPartRequest | The request body to update workspace Datatank table partition.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceDatatankParts.Update(context.Background(), orgHandle, workspaceHandle, datatankHandle, datatankTableName, datatankPartId).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceDatatankParts.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: DatatankPart
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceDatatankParts.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which we want to list the workspace Datatank. | 
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

