# \UserWorkspaceConversationArtifacts

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserGet**](UserWorkspaceConversationArtifacts.md#UserGet) | **Get** /user/{user_handle}/workspace/{workspace_handle}/ai/conversation/{conversation_id}/artifact/{artifact_id} | Get user workspace conversation artifact metadata



## UserGet

> ConversationArtifact UserGet(ctx, userHandle, workspaceHandle, conversationId, artifactId).Execute()

Get user workspace conversation artifact metadata



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
    conversationId := "conversationId_example" // string | The ID of the conversation that produced the artifact.
    artifactId := "artifactId_example" // string | The ID of the artifact to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceConversationArtifacts.UserGet(context.Background(), userHandle, workspaceHandle, conversationId, artifactId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceConversationArtifacts.UserGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGet`: ConversationArtifact
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceConversationArtifacts.UserGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the user workspace. | 
**conversationId** | **string** | The ID of the conversation that produced the artifact. | 
**artifactId** | **string** | The ID of the artifact to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**ConversationArtifact**](ConversationArtifact.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

