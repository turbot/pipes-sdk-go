# \OrgWorkspaceConversations

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgChat**](OrgWorkspaceConversations.md#OrgChat) | **Post** /org/{org_handle}/workspace/{workspace_handle}/ai/chat | Chat in org workspace
[**OrgDelete**](OrgWorkspaceConversations.md#OrgDelete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/ai/conversation/{conversation_id} | Delete org workspace conversation
[**OrgGet**](OrgWorkspaceConversations.md#OrgGet) | **Get** /org/{org_handle}/workspace/{workspace_handle}/ai/conversation/{conversation_id} | Get org workspace conversation
[**OrgList**](OrgWorkspaceConversations.md#OrgList) | **Get** /org/{org_handle}/workspace/{workspace_handle}/ai/conversation | List org workspace conversations
[**OrgUpdate**](OrgWorkspaceConversations.md#OrgUpdate) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/ai/conversation/{conversation_id} | Update org workspace conversation



## OrgChat

> Conversation OrgChat(ctx, orgHandle, workspaceHandle).Body(body).Execute()

Chat in org workspace



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
    orgHandle := "orgHandle_example" // string | The handle of the organization that owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the org workspace.
    body := *openapiclient.NewChatRequest() // ChatRequest | The chat request payload containing the message, model and optional conversation_id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceConversations.OrgChat(context.Background(), orgHandle, workspaceHandle).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceConversations.OrgChat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgChat`: Conversation
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceConversations.OrgChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the org workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgChatRequest struct via the builder pattern


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


## OrgDelete

> Conversation OrgDelete(ctx, orgHandle, workspaceHandle, conversationId).Execute()

Delete org workspace conversation



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
    orgHandle := "orgHandle_example" // string | The handle of the organization that owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the org workspace.
    conversationId := "conversationId_example" // string | The ID of the conversation to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceConversations.OrgDelete(context.Background(), orgHandle, workspaceHandle, conversationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceConversations.OrgDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgDelete`: Conversation
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceConversations.OrgDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the org workspace. | 
**conversationId** | **string** | The ID of the conversation to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgDeleteRequest struct via the builder pattern


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


## OrgGet

> Conversation OrgGet(ctx, orgHandle, workspaceHandle, conversationId).Execute()

Get org workspace conversation



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
    orgHandle := "orgHandle_example" // string | The handle of the organization that owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the org workspace.
    conversationId := "conversationId_example" // string | The ID of the conversation to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceConversations.OrgGet(context.Background(), orgHandle, workspaceHandle, conversationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceConversations.OrgGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgGet`: Conversation
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceConversations.OrgGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the org workspace. | 
**conversationId** | **string** | The ID of the conversation to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgGetRequest struct via the builder pattern


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


## OrgList

> ListConversationsResponse OrgList(ctx, orgHandle, workspaceHandle).Where(where).Limit(limit).NextToken(nextToken).Execute()

List org workspace conversations



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
    orgHandle := "orgHandle_example" // string | The handle of the organization that owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the org workspace.
    where := "where_example" // string | Optional filter expression for title. (optional)
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned. Pass next_token to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceConversations.OrgList(context.Background(), orgHandle, workspaceHandle).Where(where).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceConversations.OrgList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgList`: ListConversationsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceConversations.OrgList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the org workspace. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgListRequest struct via the builder pattern


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


## OrgUpdate

> Conversation OrgUpdate(ctx, orgHandle, workspaceHandle, conversationId).Body(body).Execute()

Update org workspace conversation



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
    orgHandle := "orgHandle_example" // string | The handle of the organization that owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the org workspace.
    conversationId := "conversationId_example" // string | The ID of the conversation to update.
    body := *openapiclient.NewUpdateConversationRequest() // UpdateConversationRequest | Fields to update on the conversation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceConversations.OrgUpdate(context.Background(), orgHandle, workspaceHandle, conversationId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceConversations.OrgUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgUpdate`: Conversation
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceConversations.OrgUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the org workspace. | 
**conversationId** | **string** | The ID of the conversation to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgUpdateRequest struct via the builder pattern


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

