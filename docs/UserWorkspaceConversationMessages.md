# \UserWorkspaceConversationMessages

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserList**](UserWorkspaceConversationMessages.md#UserList) | **Get** /user/{user_handle}/workspace/{workspace_handle}/ai/conversation/{conversation_id}/message | List user workspace conversation messages



## UserList

> ListConversationMessagesResponse UserList(ctx, userHandle, workspaceHandle, conversationId).Limit(limit).NextToken(nextToken).Execute()

List user workspace conversation messages



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
    conversationId := "conversationId_example" // string | The ID of the conversation whose messages are being listed.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned. Pass next_token to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceConversationMessages.UserList(context.Background(), userHandle, workspaceHandle, conversationId).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceConversationMessages.UserList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserList`: ListConversationMessagesResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceConversationMessages.UserList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the user workspace. | 
**conversationId** | **string** | The ID of the conversation whose messages are being listed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned. Pass next_token to fetch the next page of data. | 

### Return type

[**ListConversationMessagesResponse**](ListConversationMessagesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

