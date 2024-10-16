# \OrgWorkspaceFlowpipeInputs

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](OrgWorkspaceFlowpipeInputs.md#Create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/input | Create org workspace flowpipe input
[**Get**](OrgWorkspaceFlowpipeInputs.md#Get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/form/{step_execution_id}/{random_id} | Get org workspace flowpipe form
[**Get_0**](OrgWorkspaceFlowpipeInputs.md#Get_0) | **Get** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/input/{input_id} | Get org workspace flowpipe input
[**List**](OrgWorkspaceFlowpipeInputs.md#List) | **Get** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/input | List org workspace flowpipe inputs
[**Post**](OrgWorkspaceFlowpipeInputs.md#Post) | **Post** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/form/{step_execution_id}/{random_id}/submit | Post org workspace flowpipe form response



## Create

> FlowpipeInput Create(ctx, orgHandle, workspaceHandle).Request(request).Execute()

Create org workspace flowpipe input



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
    orgHandle := "orgHandle_example" // string | The handle of the org to which the workspace belongs.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace to create the Flowpipe input against.
    request := *openapiclient.NewCreateFlowpipeInputRequest("ExecutionId_example", "NotifierName_example", "StepExecutionId_example", map[string][]openapiclient.FlowpipeInputStepType{"key": "TODO"}) // CreateFlowpipeInputRequest | The request body for the Flowpipe input.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceFlowpipeInputs.Create(context.Background(), orgHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceFlowpipeInputs.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: FlowpipeInput
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceFlowpipeInputs.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org to which the workspace belongs. | 
**workspaceHandle** | **string** | The handle of the workspace to create the Flowpipe input against. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateFlowpipeInputRequest**](CreateFlowpipeInputRequest.md) | The request body for the Flowpipe input. | 

### Return type

[**FlowpipeInput**](FlowpipeInput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> FlowpipeInput Get(ctx, orgHandle, workspaceHandle, stepExecutionId, randomId).Execute()

Get org workspace flowpipe form



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
    orgHandle := "orgHandle_example" // string | The handle of the org to which the workspace belongs.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace to which the Flowpipe input belongs.
    stepExecutionId := "stepExecutionId_example" // string | The execution ID of the step to which the Flowpipe input belongs.
    randomId := "randomId_example" // string | The random ID used to verify the input response.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceFlowpipeInputs.Get(context.Background(), orgHandle, workspaceHandle, stepExecutionId, randomId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceFlowpipeInputs.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: FlowpipeInput
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceFlowpipeInputs.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org to which the workspace belongs. | 
**workspaceHandle** | **string** | The handle of the workspace to which the Flowpipe input belongs. | 
**stepExecutionId** | **string** | The execution ID of the step to which the Flowpipe input belongs. | 
**randomId** | **string** | The random ID used to verify the input response. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**FlowpipeInput**](FlowpipeInput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get_0

> FlowpipeInput Get_0(ctx, orgHandle, workspaceHandle, inputId).Execute()

Get org workspace flowpipe input



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
    orgHandle := "orgHandle_example" // string | The handle of the org to which the workspace belongs.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace to which the Flowpipe input belongs.
    inputId := "inputId_example" // string | The ID of the Flowpipe input. Optionally can pass the step execution ID in lieu of the input ID.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceFlowpipeInputs.Get_0(context.Background(), orgHandle, workspaceHandle, inputId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceFlowpipeInputs.Get_0``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get_0`: FlowpipeInput
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceFlowpipeInputs.Get_0`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org to which the workspace belongs. | 
**workspaceHandle** | **string** | The handle of the workspace to which the Flowpipe input belongs. | 
**inputId** | **string** | The ID of the Flowpipe input. Optionally can pass the step execution ID in lieu of the input ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGet_1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**FlowpipeInput**](FlowpipeInput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListFlowpipeInputsResponse List(ctx, orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

List org workspace flowpipe inputs



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
    orgHandle := "orgHandle_example" // string | The handle of the org to which the workspace belongs.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace to which the Flowpipe inputs belong.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceFlowpipeInputs.List(context.Background(), orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceFlowpipeInputs.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListFlowpipeInputsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceFlowpipeInputs.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org to which the workspace belongs. | 
**workspaceHandle** | **string** | The handle of the workspace to which the Flowpipe inputs belong. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListFlowpipeInputsResponse**](ListFlowpipeInputsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> FlowpipeInput Post(ctx, orgHandle, workspaceHandle, stepExecutionId, randomId).Request(request).Execute()

Post org workspace flowpipe form response



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
    orgHandle := "orgHandle_example" // string | The handle of the org to which the workspace belongs.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace to which the Flowpipe input belongs.
    stepExecutionId := "stepExecutionId_example" // string | The execution ID of the step to which the Flowpipe input belongs.
    randomId := "randomId_example" // string | The random ID used to verify the input response.
    request := map[string]map[string]interface{}{"key": map[string]interface{}(123)} // map[string]map[string]interface{} | The request body for the response to the Flowpipe input form.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceFlowpipeInputs.Post(context.Background(), orgHandle, workspaceHandle, stepExecutionId, randomId).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceFlowpipeInputs.Post``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Post`: FlowpipeInput
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceFlowpipeInputs.Post`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org to which the workspace belongs. | 
**workspaceHandle** | **string** | The handle of the workspace to which the Flowpipe input belongs. | 
**stepExecutionId** | **string** | The execution ID of the step to which the Flowpipe input belongs. | 
**randomId** | **string** | The random ID used to verify the input response. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **request** | **map[string]map[string]interface{}** | The request body for the response to the Flowpipe input form. | 

### Return type

[**FlowpipeInput**](FlowpipeInput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

