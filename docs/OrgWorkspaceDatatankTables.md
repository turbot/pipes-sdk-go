# \OrgWorkspaceDatatankTables

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](OrgWorkspaceDatatankTables.md#Create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table | Create org workspace Datatank table
[**Delete**](OrgWorkspaceDatatankTables.md#Delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name} | Delete org workspace Datatank table
[**Get**](OrgWorkspaceDatatankTables.md#Get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name} | Get org workspace Datatank table
[**List**](OrgWorkspaceDatatankTables.md#List) | **Get** /org/{org_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table | List org workspace Datatank tables
[**Update**](OrgWorkspaceDatatankTables.md#Update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name} | Update org workspace Datatank table



## Create

> DatatankTable Create(ctx, orgHandle, workspaceHandle, datatankHandle).Request(request).Execute()

Create org workspace Datatank table



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
    orgHandle := "orgHandle_example" // string | The handle of the org where we want to create the workspace Datatank table.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace.
    datatankHandle := "datatankHandle_example" // string | The name of the Datatank.
    request := *openapiclient.NewCreateDatatankTableRequest("Name_example", "Type_example") // CreateDatatankTableRequest | The request body for the workspace Datatank table to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceDatatankTables.Create(context.Background(), orgHandle, workspaceHandle, datatankHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceDatatankTables.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: DatatankTable
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceDatatankTables.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org where we want to create the workspace Datatank table. | 
**workspaceHandle** | **string** | The handle of the workspace. | 
**datatankHandle** | **string** | The name of the Datatank. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**CreateDatatankTableRequest**](CreateDatatankTableRequest.md) | The request body for the workspace Datatank table to be created. | 

### Return type

[**DatatankTable**](DatatankTable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> DatatankTable Delete(ctx, orgHandle, workspaceHandle, datatankHandle, datatankTableName).Execute()

Delete org workspace Datatank table



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
    datatankHandle := "datatankHandle_example" // string | The name of the workspace Datatank to be deleted.
    datatankTableName := "datatankTableName_example" // string | The name of the workspace Datatank table.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceDatatankTables.Delete(context.Background(), orgHandle, workspaceHandle, datatankHandle, datatankTableName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceDatatankTables.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: DatatankTable
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceDatatankTables.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org. | 
**workspaceHandle** | **string** | The handle of the workspace. | 
**datatankHandle** | **string** | The name of the workspace Datatank to be deleted. | 
**datatankTableName** | **string** | The name of the workspace Datatank table. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**DatatankTable**](DatatankTable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> DatatankTable Get(ctx, orgHandle, workspaceHandle, datatankHandle, datatankTableName).Execute()

Get org workspace Datatank table



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
    orgHandle := "orgHandle_example" // string | The handle of the org where we want to create the workspace Datatank table.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace.
    datatankHandle := "datatankHandle_example" // string | The name of the workspace Datatank.
    datatankTableName := "datatankTableName_example" // string | The name of the workspace Datatank table.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceDatatankTables.Get(context.Background(), orgHandle, workspaceHandle, datatankHandle, datatankTableName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceDatatankTables.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: DatatankTable
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceDatatankTables.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org where we want to create the workspace Datatank table. | 
**workspaceHandle** | **string** | The handle of the workspace. | 
**datatankHandle** | **string** | The name of the workspace Datatank. | 
**datatankTableName** | **string** | The name of the workspace Datatank table. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**DatatankTable**](DatatankTable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListDatatankTableResponse List(ctx, orgHandle, workspaceHandle, datatankHandle).Limit(limit).NextToken(nextToken).Execute()

List org workspace Datatank tables



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
    orgHandle := "orgHandle_example" // string | The handle of the user for which we want to list the workspace Datatank tables.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace.
    datatankHandle := "datatankHandle_example" // string | The name of the Datatank.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceDatatankTables.List(context.Background(), orgHandle, workspaceHandle, datatankHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceDatatankTables.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListDatatankTableResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceDatatankTables.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the user for which we want to list the workspace Datatank tables. | 
**workspaceHandle** | **string** | The handle of the workspace. | 
**datatankHandle** | **string** | The name of the Datatank. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListDatatankTableResponse**](ListDatatankTableResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> DatatankTable Update(ctx, orgHandle, workspaceHandle, datatankHandle, datatankTableName).Request(request).Execute()

Update org workspace Datatank table



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
    datatankTableName := "datatankTableName_example" // string | The name of the workspace Datatank table to be updated.
    request := *openapiclient.NewUpdateDatatankTableRequest() // UpdateDatatankTableRequest | The request body to update workspace Datatank table.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceDatatankTables.Update(context.Background(), orgHandle, workspaceHandle, datatankHandle, datatankTableName).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceDatatankTables.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: DatatankTable
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceDatatankTables.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which we want to list the workspace Datatank. | 
**workspaceHandle** | **string** | The handle of the workspace. | 
**datatankHandle** | **string** | The name of the workspace Datatank. | 
**datatankTableName** | **string** | The name of the workspace Datatank table to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **request** | [**UpdateDatatankTableRequest**](UpdateDatatankTableRequest.md) | The request body to update workspace Datatank table. | 

### Return type

[**DatatankTable**](DatatankTable.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

