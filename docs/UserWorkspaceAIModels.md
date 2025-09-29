# \UserWorkspaceAIModels

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserList**](UserWorkspaceAIModels.md#UserList) | **Get** /user/{user_handle}/workspace/{workspace_handle}/ai/model | List user workspace AI models



## UserList

> ListModelsResponse UserList(ctx, userHandle, workspaceHandle).Execute()

List user workspace AI models



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
    userHandle := "userHandle_example" // string | The handle of the user that owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the user workspace.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceAIModels.UserList(context.Background(), userHandle, workspaceHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceAIModels.UserList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserList`: ListModelsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceAIModels.UserList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the user workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ListModelsResponse**](ListModelsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

