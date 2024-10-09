# \UserWorkspaceFlowpipePipelines

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Command**](UserWorkspaceFlowpipePipelines.md#Command) | **Post** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/pipeline/{pipeline_name}/command | Run user workspace Flowpipe pipeline command
[**Get**](UserWorkspaceFlowpipePipelines.md#Get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/pipeline/{pipeline_id} | Get user workspace pipeline
[**ListTriggers**](UserWorkspaceFlowpipePipelines.md#ListTriggers) | **Get** /user/{user_handle}/workspace/{workspace_handle}/flowpipe/pipeline/{pipeline_id}/trigger | Get user workspace pipeline triggers



## Command

> PipelineCommandResponse Command(ctx, userHandle, workspaceHandle, pipelineName).Request(request).Execute()

Run user workspace Flowpipe pipeline command



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
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the Flowpipe pipeline exists.
    pipelineName := "pipelineName_example" // string | Identifier of the Flowpipe pipeline on which the command will be run.
    request := *openapiclient.NewPipelineCommandRequest(map[string][]openapiclient.PipelineCommandAction{"key": "TODO"}) // PipelineCommandRequest | The request body of the pipeline command to run.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceFlowpipePipelines.Command(context.Background(), userHandle, workspaceHandle, pipelineName).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceFlowpipePipelines.Command``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Command`: PipelineCommandResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceFlowpipePipelines.Command`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user which contains the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the Flowpipe pipeline exists. | 
**pipelineName** | **string** | Identifier of the Flowpipe pipeline on which the command will be run. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**PipelineCommandRequest**](PipelineCommandRequest.md) | The request body of the pipeline command to run. | 

### Return type

[**PipelineCommandResponse**](PipelineCommandResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Pipeline Get(ctx, userHandle, workspaceHandle, pipelineId).Execute()

Get user workspace pipeline



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
    pipelineId := "pipelineId_example" // string | The id of the pipeline whose detail needs to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceFlowpipePipelines.Get(context.Background(), userHandle, workspaceHandle, pipelineId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceFlowpipePipelines.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceFlowpipePipelines.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user which contains the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the pipeline exists. | 
**pipelineId** | **string** | The id of the pipeline whose detail needs to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


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


## ListTriggers

> ListFlowpipeTriggersResponse ListTriggers(ctx, userHandle, workspaceHandle, pipelineId).Where(where).Limit(limit).NextToken(nextToken).Execute()

Get user workspace pipeline triggers



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
    pipelineId := "pipelineId_example" // string | The id of the pipeline whose triggers need to be fetched.
    where := "where_example" // string | The SQL where filter you wish to apply to this request. The filter will be parsed and sanitised and checked against the supported columns for this API. (optional)
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceFlowpipePipelines.ListTriggers(context.Background(), userHandle, workspaceHandle, pipelineId).Where(where).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceFlowpipePipelines.ListTriggers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTriggers`: ListFlowpipeTriggersResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceFlowpipePipelines.ListTriggers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user which contains the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the pipeline exists. | 
**pipelineId** | **string** | The id of the pipeline whose triggers need to be fetched. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTriggersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **where** | **string** | The SQL where filter you wish to apply to this request. The filter will be parsed and sanitised and checked against the supported columns for this API. | 
 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListFlowpipeTriggersResponse**](ListFlowpipeTriggersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

