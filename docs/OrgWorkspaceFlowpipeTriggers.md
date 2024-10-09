# \OrgWorkspaceFlowpipeTriggers

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Command**](OrgWorkspaceFlowpipeTriggers.md#Command) | **Post** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/trigger/{trigger_name}/command | Run organization workspace Flowpipe trigger command
[**Create**](OrgWorkspaceFlowpipeTriggers.md#Create) | **Post** /org/{org_handle}/workspace/{workspace_handle}/trigger | Create org workspace trigger
[**Delete**](OrgWorkspaceFlowpipeTriggers.md#Delete) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/trigger/{trigger_name} | Delete org workspace pipeline
[**Get**](OrgWorkspaceFlowpipeTriggers.md#Get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/trigger/{trigger_name} | Get org workspace flowpipe trigger
[**Update**](OrgWorkspaceFlowpipeTriggers.md#Update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/trigger/{trigger_name} | Update org workspace trigger



## Command

> TriggerCommandResponse Command(ctx, orgHandle, workspaceHandle, triggerName).Request(request).Execute()

Run organization workspace Flowpipe trigger command



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
    orgHandle := "orgHandle_example" // string | The handle of the organization which contains the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the Flowpipe trigger exists.
    triggerName := "triggerName_example" // string | Identifier of the Flowpipe trigger on which the command will be run.
    request := *openapiclient.NewTriggerCommandRequest(map[string][]openapiclient.TriggerCommandAction{"key": "TODO"}) // TriggerCommandRequest | The request body of the trigger command to run.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceFlowpipeTriggers.Command(context.Background(), orgHandle, workspaceHandle, triggerName).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceFlowpipeTriggers.Command``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Command`: TriggerCommandResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceFlowpipeTriggers.Command`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization which contains the workspace. | 
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

> WorkspaceModTrigger Create(ctx, orgHandle, workspaceHandle).Request(request).Execute()

Create org workspace trigger



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
    orgHandle := "orgHandle_example" // string | The handle of the org which contains the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where we want to create the trigger.
    request := *openapiclient.NewCreateTriggerRequest(map[string]interface{}(123), "Pipeline_example", map[string][]openapiclient.PipelineFrequency{"key": "TODO"}) // CreateTriggerRequest | The request body for the trigger to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceFlowpipeTriggers.Create(context.Background(), orgHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceFlowpipeTriggers.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: WorkspaceModTrigger
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceFlowpipeTriggers.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org which contains the workspace. | 
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

> Pipeline Delete(ctx, orgHandle, workspaceHandle, triggerName).Execute()

Delete org workspace pipeline



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
    orgHandle := "orgHandle_example" // string | The handle of the organization which contains the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the pipeline exists.
    triggerName := "triggerName_example" // string | The name of the trigger which needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceFlowpipeTriggers.Delete(context.Background(), orgHandle, workspaceHandle, triggerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceFlowpipeTriggers.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceFlowpipeTriggers.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization which contains the workspace. | 
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

> WorkspaceModTrigger Get(ctx, orgHandle, workspaceHandle, triggerName).Execute()

Get org workspace flowpipe trigger



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
    orgHandle := "orgHandle_example" // string | The handle of the org which contains the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the trigger exists.
    triggerName := "triggerName_example" // string | The name of the flowpipe trigger whose details needs to be fetched.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceFlowpipeTriggers.Get(context.Background(), orgHandle, workspaceHandle, triggerName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceFlowpipeTriggers.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: WorkspaceModTrigger
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceFlowpipeTriggers.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org which contains the workspace. | 
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


## Update

> WorkspaceModTrigger Update(ctx, orgHandle, workspaceHandle, triggerName).Request(request).Execute()

Update org workspace trigger



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
    orgHandle := "orgHandle_example" // string | The handle of the org which contains the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the trigger exists.
    triggerName := "triggerName_example" // string | The name of the trigger which needs to be updated.
    request := *openapiclient.NewUpdateTriggerRequest() // UpdateTriggerRequest | The request body for the trigger which needs to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceFlowpipeTriggers.Update(context.Background(), orgHandle, workspaceHandle, triggerName).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceFlowpipeTriggers.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: WorkspaceModTrigger
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceFlowpipeTriggers.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org which contains the workspace. | 
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

