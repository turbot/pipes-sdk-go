# \UserWorkspaceConversations

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserChat**](UserWorkspaceConversations.md#UserChat) | **Post** /user/{user_handle}/workspace/{workspace_handle}/ai/chat | Chat in user workspace
[**UserDelete**](UserWorkspaceConversations.md#UserDelete) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/ai/conversation/{conversation_id} | Delete user workspace conversation
[**UserGet**](UserWorkspaceConversations.md#UserGet) | **Get** /user/{user_handle}/workspace/{workspace_handle}/ai/conversation/{conversation_id} | Get user workspace conversation
[**UserList**](UserWorkspaceConversations.md#UserList) | **Get** /user/{user_handle}/workspace/{workspace_handle}/ai/conversation | List user workspace conversations
[**UserUpdate**](UserWorkspaceConversations.md#UserUpdate) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/ai/conversation/{conversation_id} | Update user workspace conversation



## UserChat

> Conversation UserChat(ctx, userHandle, workspaceHandle).Body(body).Execute()

Chat in user workspace



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
    body := *openapiclient.NewChatRequest() // ChatRequest | The chat request payload containing the message, model and optional conversation_id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceConversations.UserChat(context.Background(), userHandle, workspaceHandle).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceConversations.UserChat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserChat`: Conversation
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceConversations.UserChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the user workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ChatRequest**](ChatRequest.md) | The chat request payload containing the message, model and optional conversation_id. | 

### Return type

[**Conversation**](Conversation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserDelete

> Conversation UserDelete(ctx, userHandle, workspaceHandle, conversationId).Execute()

Delete user workspace conversation



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
    conversationId := "conversationId_example" // string | The ID of the conversation to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceConversations.UserDelete(context.Background(), userHandle, workspaceHandle, conversationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceConversations.UserDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserDelete`: Conversation
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceConversations.UserDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the user workspace. | 
**conversationId** | **string** | The ID of the conversation to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Conversation**](Conversation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGet

> Conversation UserGet(ctx, userHandle, workspaceHandle, conversationId).Execute()

Get user workspace conversation



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
    conversationId := "conversationId_example" // string | The ID of the conversation to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceConversations.UserGet(context.Background(), userHandle, workspaceHandle, conversationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceConversations.UserGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGet`: Conversation
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceConversations.UserGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the user workspace. | 
**conversationId** | **string** | The ID of the conversation to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Conversation**](Conversation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserList

> ListConversationsResponse UserList(ctx, userHandle, workspaceHandle).Where(where).Limit(limit).NextToken(nextToken).Execute()

List user workspace conversations



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
    where := "where_example" // string | Optional filter expression for title. (optional)
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned. Pass next_token to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceConversations.UserList(context.Background(), userHandle, workspaceHandle).Where(where).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceConversations.UserList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserList`: ListConversationsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceConversations.UserList`: %v\n", resp)
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


 **where** | **string** | Optional filter expression for title. | 
 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned. Pass next_token to fetch the next page of data. | 

### Return type

[**ListConversationsResponse**](ListConversationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUpdate

> Conversation UserUpdate(ctx, userHandle, workspaceHandle, conversationId).Body(body).Execute()

Update user workspace conversation



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
    conversationId := "conversationId_example" // string | The ID of the conversation to update.
    body := *openapiclient.NewUpdateConversationRequest() // UpdateConversationRequest | Fields to update on the conversation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceConversations.UserUpdate(context.Background(), userHandle, workspaceHandle, conversationId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceConversations.UserUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserUpdate`: Conversation
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceConversations.UserUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the user workspace. | 
**conversationId** | **string** | The ID of the conversation to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**UpdateConversationRequest**](UpdateConversationRequest.md) | Fields to update on the conversation. | 

### Return type

[**Conversation**](Conversation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

