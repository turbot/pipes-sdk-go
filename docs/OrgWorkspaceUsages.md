# \OrgWorkspaceUsages

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**List**](OrgWorkspaceUsages.md#List) | **Get** /org/{org_handle}/workspace/{workspace_handle}/usage | List org workspace usage



## List

> ListUsageMetricsResponse List(ctx, orgHandle, workspaceHandle).Where(where).Limit(limit).NextToken(nextToken).Execute()

List org workspace usage



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
    orgHandle := "orgHandle_example" // string | The handle of the organization where we want to create the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where you want to list the associations.
    where := "where_example" // string | The SQL where filter you wish to apply to this request. The filter will be parsed and sanitised and checked against the supported columns for this API. (optional)
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceUsages.List(context.Background(), orgHandle, workspaceHandle).Where(where).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceUsages.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListUsageMetricsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceUsages.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization where we want to create the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where you want to list the associations. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **where** | **string** | The SQL where filter you wish to apply to this request. The filter will be parsed and sanitised and checked against the supported columns for this API. | 
 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListUsageMetricsResponse**](ListUsageMetricsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

