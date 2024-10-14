# \UserWorkspaceNotifiers

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](UserWorkspaceNotifiers.md#Create) | **Post** /user/{user_handle}/workspace/{workspace_handle}/notifier | Create user workspace notifier
[**Delete**](UserWorkspaceNotifiers.md#Delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/notifier/{notifier_name} | Delete user workspace notifier
[**Get**](UserWorkspaceNotifiers.md#Get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/notifier/{notifier_name} | Get user workspace notifier
[**List**](UserWorkspaceNotifiers.md#List) | **Get** /user/{user_handle}/workspace/{workspace_handle}/notifier | List user workspace notifiers
[**Post**](UserWorkspaceNotifiers.md#Post) | **Post** /user/{user_handle}/workspace/{workspace_handle}/notifier/{notifier_name}/command | Post user workspace notifier commands
[**Update**](UserWorkspaceNotifiers.md#Update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/notifier/{notifier_name} | Update user workspace notifier



## Create

> Notifier Create(ctx, userHandle, workspaceHandle).Request(request).Execute()

Create user workspace notifier



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
    userHandle := "userHandle_example" // string | The handle of the user for which the notifier is to be created.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace for which the notifier is to be created.
    request := map[string]interface{}{"key": "TODO"} // map[string]interface{} | The request body for the notifier to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceNotifiers.Create(context.Background(), userHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceNotifiers.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Notifier
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceNotifiers.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which the notifier is to be created. | 
**workspaceHandle** | **string** | The handle of the workspace for which the notifier is to be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | **map[string]interface{}** | The request body for the notifier to be created. | 

### Return type

[**Notifier**](Notifier.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Notifier Delete(ctx, userHandle, workspaceHandle, notifierName).Execute()

Delete user workspace notifier



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
    userHandle := "userHandle_example" // string | The handle of the user for which the notifier is to be deleted.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace for which the notifier is to be deleted.
    notifierName := "notifierName_example" // string | The name of the notifier to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceNotifiers.Delete(context.Background(), userHandle, workspaceHandle, notifierName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceNotifiers.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: Notifier
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceNotifiers.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which the notifier is to be deleted. | 
**workspaceHandle** | **string** | The handle of the workspace for which the notifier is to be deleted. | 
**notifierName** | **string** | The name of the notifier to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Notifier**](Notifier.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Notifier Get(ctx, userHandle, workspaceHandle, notifierName).Execute()

Get user workspace notifier



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
    userHandle := "userHandle_example" // string | The handle of the user for which the notifier is to be retrieved.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace for which the notifier is to be retrieved.
    notifierName := "notifierName_example" // string | The name of the notifier to be retrieved.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceNotifiers.Get(context.Background(), userHandle, workspaceHandle, notifierName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceNotifiers.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Notifier
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceNotifiers.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which the notifier is to be retrieved. | 
**workspaceHandle** | **string** | The handle of the workspace for which the notifier is to be retrieved. | 
**notifierName** | **string** | The name of the notifier to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Notifier**](Notifier.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListNotifiersResponse List(ctx, userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Where(where).Execute()

List user workspace notifiers



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
    userHandle := "userHandle_example" // string | The handle of the user for which the notifiers are to be listed.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace for which the notifiers are to be listed.
    limit := int32(56) // int32 | The maximum number of items to return. (optional)
    nextToken := "nextToken_example" // string | The token to retrieve the next page. (optional)
    where := "where_example" // string | The filter to apply to the list. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceNotifiers.List(context.Background(), userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Where(where).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceNotifiers.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListNotifiersResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceNotifiers.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which the notifiers are to be listed. | 
**workspaceHandle** | **string** | The handle of the workspace for which the notifiers are to be listed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The maximum number of items to return. | 
 **nextToken** | **string** | The token to retrieve the next page. | 
 **where** | **string** | The filter to apply to the list. | 

### Return type

[**ListNotifiersResponse**](ListNotifiersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> WorkspaceNotifierCommandResponse Post(ctx, userHandle, workspaceHandle, notifierName).Request(request).Execute()

Post user workspace notifier commands



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
    userHandle := "userHandle_example" // string | The handle of the user for which the notifier is to be updated.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace for which the notifier is to be updated.
    notifierName := "notifierName_example" // string | The name of the notifier for which the command is to be posted.
    request := *openapiclient.NewWorkspaceNotifierCommandRequest(map[string][]openapiclient.WorkspaceNotifierCommandAction{"key": "TODO"}) // WorkspaceNotifierCommandRequest | The request body for the notifier command to be posted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceNotifiers.Post(context.Background(), userHandle, workspaceHandle, notifierName).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceNotifiers.Post``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Post`: WorkspaceNotifierCommandResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceNotifiers.Post`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which the notifier is to be updated. | 
**workspaceHandle** | **string** | The handle of the workspace for which the notifier is to be updated. | 
**notifierName** | **string** | The name of the notifier for which the command is to be posted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**WorkspaceNotifierCommandRequest**](WorkspaceNotifierCommandRequest.md) | The request body for the notifier command to be posted. | 

### Return type

[**WorkspaceNotifierCommandResponse**](WorkspaceNotifierCommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Notifier Update(ctx, userHandle, workspaceHandle, notifierName).Request(request).Execute()

Update user workspace notifier



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
    userHandle := "userHandle_example" // string | The handle of the user for which the notifier is to be updated.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace for which the notifier is to be updated.
    notifierName := "notifierName_example" // string | The name of the notifier to be updated.
    request := map[string]interface{}{"key": "TODO"} // map[string]interface{} | The request body for the notifier to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceNotifiers.Update(context.Background(), userHandle, workspaceHandle, notifierName).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceNotifiers.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Notifier
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceNotifiers.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which the notifier is to be updated. | 
**workspaceHandle** | **string** | The handle of the workspace for which the notifier is to be updated. | 
**notifierName** | **string** | The name of the notifier to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | **map[string]interface{}** | The request body for the notifier to be updated. | 

### Return type

[**Notifier**](Notifier.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

