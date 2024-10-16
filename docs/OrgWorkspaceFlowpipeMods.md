# \OrgWorkspaceFlowpipeMods

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](OrgWorkspaceFlowpipeMods.md#Get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/mod/{mod_alias} | Get organization workspace installed flowpipe mod
[**Install**](OrgWorkspaceFlowpipeMods.md#Install) | **Post** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/mod | Install flowpipe mod to organization workspace.
[**Uninstall**](OrgWorkspaceFlowpipeMods.md#Uninstall) | **Delete** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/mod/{mod_alias} | Uninstall flowpipe mod from organization workspace.
[**Update**](OrgWorkspaceFlowpipeMods.md#Update) | **Patch** /org/{org_handle}/workspace/{workspace_handle}/flowpipe/mod/{mod_alias} | Update a flowpipe mod in an organization workspace



## Get

> WorkspaceMod Get(ctx, orgHandle, workspaceHandle, modAlias).Execute()

Get organization workspace installed flowpipe mod



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
    orgHandle := "orgHandle_example" // string | The handle of an organization that owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where mod was installed.
    modAlias := "modAlias_example" // string | The mod alias or mod ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceFlowpipeMods.Get(context.Background(), orgHandle, workspaceHandle, modAlias).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceFlowpipeMods.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: WorkspaceMod
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceFlowpipeMods.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where mod was installed. | 
**modAlias** | **string** | The mod alias or mod ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorkspaceMod**](WorkspaceMod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Install

> WorkspaceMod Install(ctx, orgHandle, workspaceHandle).Request(request).Execute()

Install flowpipe mod to organization workspace.



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
    orgHandle := "orgHandle_example" // string | The handle of an organization that owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the mod will be installed.
    request := *openapiclient.NewCreateWorkspaceModRequest("Path_example") // CreateWorkspaceModRequest | The request body to install a mod in the mentioned workspace for this organization.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceFlowpipeMods.Install(context.Background(), orgHandle, workspaceHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceFlowpipeMods.Install``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Install`: WorkspaceMod
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceFlowpipeMods.Install`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the mod will be installed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**CreateWorkspaceModRequest**](CreateWorkspaceModRequest.md) | The request body to install a mod in the mentioned workspace for this organization. | 

### Return type

[**WorkspaceMod**](WorkspaceMod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Uninstall

> WorkspaceMod Uninstall(ctx, orgHandle, workspaceHandle, modAlias).Execute()

Uninstall flowpipe mod from organization workspace.



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
    orgHandle := "orgHandle_example" // string | The handle of an organization that owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the mod was installed.
    modAlias := "modAlias_example" // string | The mod alias or mod ID to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceFlowpipeMods.Uninstall(context.Background(), orgHandle, workspaceHandle, modAlias).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceFlowpipeMods.Uninstall``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Uninstall`: WorkspaceMod
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceFlowpipeMods.Uninstall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the mod was installed. | 
**modAlias** | **string** | The mod alias or mod ID to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUninstallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**WorkspaceMod**](WorkspaceMod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> WorkspaceMod Update(ctx, orgHandle, workspaceHandle, modAlias).Request(request).Execute()

Update a flowpipe mod in an organization workspace



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
    orgHandle := "orgHandle_example" // string | The handle of an organization that owns the workspace.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace where the mod will be updated.
    modAlias := "modAlias_example" // string | The mod alias or mod ID to update.
    request := *openapiclient.NewUpdateWorkspaceModRequest() // UpdateWorkspaceModRequest | The request body to update a mod for this workspace.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceFlowpipeMods.Update(context.Background(), orgHandle, workspaceHandle, modAlias).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceFlowpipeMods.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: WorkspaceMod
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceFlowpipeMods.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization that owns the workspace. | 
**workspaceHandle** | **string** | The handle of the workspace where the mod will be updated. | 
**modAlias** | **string** | The mod alias or mod ID to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **request** | [**UpdateWorkspaceModRequest**](UpdateWorkspaceModRequest.md) | The request body to update a mod for this workspace. | 

### Return type

[**WorkspaceMod**](WorkspaceMod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

