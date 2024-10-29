# \UserWorkspaceFlowpipeTriggers

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Command**](UserWorkspaceFlowpipeTriggers.md#Command) | **Post** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/trigger/{trigger_name}/command | Run user workspace Flowpipe trigger command
[**Create**](UserWorkspaceFlowpipeTriggers.md#Create) | **Post** /user/{user_handle}/workspace/{workspace_handle}/trigger | Create user workspace trigger
[**Delete**](UserWorkspaceFlowpipeTriggers.md#Delete) | **Delete** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/trigger/{trigger_name} | Delete user workspace pipeline
[**Get**](UserWorkspaceFlowpipeTriggers.md#Get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/trigger/{trigger_name} | Get user workspace flowpipe trigger
[**List**](UserWorkspaceFlowpipeTriggers.md#List) | **Get** /user/{user_handle}/workspace/{workspace_handle}/trigger | List user workspace triggers
[**Update**](UserWorkspaceFlowpipeTriggers.md#Update) | **Patch** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/trigger/{trigger_name} | Update user workspace trigger



## Command

> TriggerCommandResponse Command(ctx, userHandle, workspaceHandle, triggerName).Request(request).Execute()

Run user workspace Flowpipe trigger command



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
    userHandle := "userHandle_example" // string | The handle of the user which contains the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the Flowpipe trigger exists.
    triggerName := "triggerName_example" // string | Identifier of the Flowpipe trigger on which the command will be run.
    request := *openapiclient.NewTriggerCommandRequest(map[string][]openapiclient.TriggerCommandAction{"key": "TODO"}) // TriggerCommandRequest | The request body of the trigger command to run.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceFlowpipeTriggers.Command(context.Background(), userHandle, workspaceHandle, triggerName).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceFlowpipeTriggers.Command``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Command`: TriggerCommandResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceFlowpipeTriggers.Command`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user which contains the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the Flowpipe trigger exists. | 
**triggerName** | **string** | Identifier of the Flowpipe trigger on which the command will be run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**TriggerCommandRequest**](TriggerCommandRequest.md) | The request body of the trigger command to run. | 

### Return type

[**TriggerCommandResponse**](TriggerCommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> WorkspaceModTrigger Create(ctx, userHandle, workspaceHandle).Request(request).Execute()

Create user workspace trigger



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
    userHandle := "userHandle_example" // string | The handle of the user which contains the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where we want to create the trigger.
    request := *openapiclient.NewCreateTriggerRequest(map[string]interface{}(123), "Pipeline_example", map[string][]openapiclient.PipelineFrequency{"key": "TODO"}) // CreateTriggerRequest | The request body for the trigger to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceFlowpipeTriggers.Create(context.Background(), userHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceFlowpipeTriggers.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: WorkspaceModTrigger
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceFlowpipeTriggers.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user which contains the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where we want to create the trigger. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateTriggerRequest**](CreateTriggerRequest.md) | The request body for the trigger to be created. | 

### Return type

[**WorkspaceModTrigger**](WorkspaceModTrigger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Pipeline Delete(ctx, userHandle, workspaceHandle, triggerName).Execute()

Delete user workspace pipeline



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
    userHandle := "userHandle_example" // string | The handle of the user which contains the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the pipeline exists.
    triggerName := "triggerName_example" // string | The name of the trigger which needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceFlowpipeTriggers.Delete(context.Background(), userHandle, workspaceHandle, triggerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceFlowpipeTriggers.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceFlowpipeTriggers.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user which contains the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the pipeline exists. | 
**triggerName** | **string** | The name of the trigger which needs to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Pipeline**](Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> WorkspaceModTrigger Get(ctx, userHandle, workspaceHandle, triggerName).Execute()

Get user workspace flowpipe trigger



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
    userHandle := "userHandle_example" // string | The handle of the user which contains the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the trigger exists.
    triggerName := "triggerName_example" // string | The name of the flowpipe trigger whose details needs to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceFlowpipeTriggers.Get(context.Background(), userHandle, workspaceHandle, triggerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceFlowpipeTriggers.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: WorkspaceModTrigger
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceFlowpipeTriggers.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user which contains the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the trigger exists. | 
**triggerName** | **string** | The name of the flowpipe trigger whose details needs to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorkspaceModTrigger**](WorkspaceModTrigger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListTriggersResponse List(ctx, userHandle, workspaceHandle).Where(where).Limit(limit).NextToken(nextToken).Execute()

List user workspace triggers



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
    userHandle := "userHandle_example" // string | The handle of the user for which contains the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace for which we want to list the triggers.
    where := "where_example" // string | The SQL where filter you wish to apply to this request. The filter will be parsed and sanitised and checked against the supported columns for this API. (optional)
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceFlowpipeTriggers.List(context.Background(), userHandle, workspaceHandle).Where(where).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceFlowpipeTriggers.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListTriggersResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceFlowpipeTriggers.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which contains the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace for which we want to list the triggers. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **where** | **string** | The SQL where filter you wish to apply to this request. The filter will be parsed and sanitised and checked against the supported columns for this API. | 
 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListTriggersResponse**](ListTriggersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> WorkspaceModTrigger Update(ctx, userHandle, workspaceHandle, triggerName).Request(request).Execute()

Update user workspace trigger



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
    userHandle := "userHandle_example" // string | The handle of the user which contains the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the trigger exists.
    triggerName := "triggerName_example" // string | The name of the trigger which needs to be updated.
    request := *openapiclient.NewUpdateTriggerRequest() // UpdateTriggerRequest | The request body for the trigger which needs to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceFlowpipeTriggers.Update(context.Background(), userHandle, workspaceHandle, triggerName).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceFlowpipeTriggers.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: WorkspaceModTrigger
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceFlowpipeTriggers.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user which contains the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the trigger exists. | 
**triggerName** | **string** | The name of the trigger which needs to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**UpdateTriggerRequest**](UpdateTriggerRequest.md) | The request body for the trigger which needs to be updated. | 

### Return type

[**WorkspaceModTrigger**](WorkspaceModTrigger.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

