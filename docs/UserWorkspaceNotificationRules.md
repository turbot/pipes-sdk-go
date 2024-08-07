# \UserWorkspaceNotificationRules

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](UserWorkspaceNotificationRules.md#Create) | **Post** /user/{user_handle}/workspace/{workspace_handle}/notification_rule | Create user workspace notification rule
[**Delete**](UserWorkspaceNotificationRules.md#Delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/notification_rule/{notification_rule_id} | Delete user workspace notification rule
[**Get**](UserWorkspaceNotificationRules.md#Get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/notification_rule/{notification_rule_id} | Get user workspace notification rule
[**List**](UserWorkspaceNotificationRules.md#List) | **Get** /user/{user_handle}/workspace/{workspace_handle}/notification_rule | List user workspace notification rules



## Create

> NotificationRule Create(ctx, userHandle, workspaceHandle).Request(request).Execute()

Create user workspace notification rule



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
    userHandle := "userHandle_example" // string | Specify the handle of the user who owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the notification rule will be created.
    request := *openapiclient.NewCreateWorkspaceNotificationRequestNoSender([]string{"Events_example"}, "Title_example") // CreateWorkspaceNotificationRequestNoSender | The request body to create a notification rule for this user.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceNotificationRules.Create(context.Background(), userHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceNotificationRules.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: NotificationRule
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceNotificationRules.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | Specify the handle of the user who owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the notification rule will be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateWorkspaceNotificationRequestNoSender**](CreateWorkspaceNotificationRequestNoSender.md) | The request body to create a notification rule for this user. | 

### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> NotificationRule Delete(ctx, userHandle, workspaceHandle, notificationRuleId).Execute()

Delete user workspace notification rule



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
    userHandle := "userHandle_example" // string | Specify the handle of the user who owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the notification exists.
    notificationRuleId := "notificationRuleId_example" // string | The notification rule id to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceNotificationRules.Delete(context.Background(), userHandle, workspaceHandle, notificationRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceNotificationRules.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: NotificationRule
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceNotificationRules.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | Specify the handle of the user who owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the notification exists. | 
**notificationRuleId** | **string** | The notification rule id to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> NotificationRule Get(ctx, userHandle, workspaceHandle, notificationRuleId).Execute()

Get user workspace notification rule



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
    userHandle := "userHandle_example" // string | Specify the handle of the user who owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | Provide the handle of the workspace where the notification exists.
    notificationRuleId := "notificationRuleId_example" // string | The notification rule id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceNotificationRules.Get(context.Background(), userHandle, workspaceHandle, notificationRuleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceNotificationRules.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: NotificationRule
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceNotificationRules.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | Specify the handle of the user who owns the workspace. | 
**workspaceHandle** | **string** | Provide the handle of the workspace where the notification exists. | 
**notificationRuleId** | **string** | The notification rule id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListNotificationRulesResponse List(ctx, userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

List user workspace notification rules



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
    userHandle := "userHandle_example" // string | Specify the handle of the user who owns the workspace
    workspaceHandle := "workspaceHandle_example" // string | Provide the handle of the workspace where the notification exists.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceNotificationRules.List(context.Background(), userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceNotificationRules.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListNotificationRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceNotificationRules.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | Specify the handle of the user who owns the workspace | 
**workspaceHandle** | **string** | Provide the handle of the workspace where the notification exists. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListNotificationRulesResponse**](ListNotificationRulesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

