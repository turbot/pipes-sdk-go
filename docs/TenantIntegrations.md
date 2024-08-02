# \TenantIntegrations

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Command**](TenantIntegrations.md#Command) | **Post** /integration/{integration_handle}/command | Run tenant integration command
[**Create**](TenantIntegrations.md#Create) | **Post** /integration | Create tenant integration
[**Delete**](TenantIntegrations.md#Delete) | **Delete** /integration/{integration_handle} | Delete tenant integration
[**Get**](TenantIntegrations.md#Get) | **Get** /integration/{integration_handle} | Get tenant integration
[**InstallGithubIntegration**](TenantIntegrations.md#InstallGithubIntegration) | **Get** /integration/{integration_handle}/github/install | Install GitHub integration on a custom tenant
[**List**](TenantIntegrations.md#List) | **Get** /integration | List tenant integrations
[**Test**](TenantIntegrations.md#Test) | **Post** /integration/{integration_handle}/test | Test custom tenant integration
[**Update**](TenantIntegrations.md#Update) | **Patch** /integration/{integration_handle} | Update tenant integration



## Command

> IntegrationCommandRequest Command(ctx, integrationHandle).Request(request).Execute()

Run tenant integration command



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
    integrationHandle := "integrationHandle_example" // string | The handle of the integration for which we want to run the command.
    request := *openapiclient.NewWorkspaceCommandRequest("Command_example") // WorkspaceCommandRequest | The request body for the workspace command.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantIntegrations.Command(context.Background(), integrationHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantIntegrations.Command``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Command`: IntegrationCommandRequest
    fmt.Fprintf(os.Stdout, "Response from `TenantIntegrations.Command`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationHandle** | **string** | The handle of the integration for which we want to run the command. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCommandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**WorkspaceCommandRequest**](WorkspaceCommandRequest.md) | The request body for the workspace command. | 

### Return type

[**IntegrationCommandRequest**](IntegrationCommandRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Create

> Integration Create(ctx).Request(request).Execute()

Create tenant integration



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
    request := *openapiclient.NewCreateIntegrationRequest("Handle_example", "Type_example") // CreateIntegrationRequest | The request body for the integration to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantIntegrations.Create(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantIntegrations.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Integration
    fmt.Fprintf(os.Stdout, "Response from `TenantIntegrations.Create`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**CreateIntegrationRequest**](CreateIntegrationRequest.md) | The request body for the integration to be created. | 

### Return type

[**Integration**](Integration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Integration Delete(ctx, integrationHandle).Execute()

Delete tenant integration



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
    integrationHandle := "integrationHandle_example" // string | The handle of the integration which needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantIntegrations.Delete(context.Background(), integrationHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantIntegrations.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: Integration
    fmt.Fprintf(os.Stdout, "Response from `TenantIntegrations.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationHandle** | **string** | The handle of the integration which needs to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Integration**](Integration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Integration Get(ctx, integrationHandle).Execute()

Get tenant integration



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
    integrationHandle := "integrationHandle_example" // string | The handle of the integration whose details need to be retrieved.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantIntegrations.Get(context.Background(), integrationHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantIntegrations.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Integration
    fmt.Fprintf(os.Stdout, "Response from `TenantIntegrations.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationHandle** | **string** | The handle of the integration whose details need to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Integration**](Integration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstallGithubIntegration

> InstallGithubIntegration(ctx, integrationHandle).Execute()

Install GitHub integration on a custom tenant



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
    integrationHandle := "integrationHandle_example" // string | The handle of the github integration which needs to be installed.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantIntegrations.InstallGithubIntegration(context.Background(), integrationHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantIntegrations.InstallGithubIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationHandle** | **string** | The handle of the github integration which needs to be installed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiInstallGithubIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListIntegrationsResponse List(ctx).Limit(limit).NextToken(nextToken).Execute()

List tenant integrations



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
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantIntegrations.List(context.Background()).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantIntegrations.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListIntegrationsResponse
    fmt.Fprintf(os.Stdout, "Response from `TenantIntegrations.List`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListIntegrationsResponse**](ListIntegrationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Test

> Integration Test(ctx, integrationHandle).Execute()

Test custom tenant integration



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
    integrationHandle := "integrationHandle_example" // string | The handle of the integration to be tested. For integrations that are not yet created, use underscore `_` as the handle, else pass the handle of the existing integration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantIntegrations.Test(context.Background(), integrationHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantIntegrations.Test``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Test`: Integration
    fmt.Fprintf(os.Stdout, "Response from `TenantIntegrations.Test`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationHandle** | **string** | The handle of the integration to be tested. For integrations that are not yet created, use underscore &#x60;_&#x60; as the handle, else pass the handle of the existing integration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Integration**](Integration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Integration Update(ctx, integrationHandle).Request(request).Execute()

Update tenant integration



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
    integrationHandle := "integrationHandle_example" // string | The handle of the integration whose details need to be updated.
    request := *openapiclient.NewUpdateIntegrationRequest() // UpdateIntegrationRequest | The request body for the integration update.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantIntegrations.Update(context.Background(), integrationHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantIntegrations.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Integration
    fmt.Fprintf(os.Stdout, "Response from `TenantIntegrations.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationHandle** | **string** | The handle of the integration whose details need to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**UpdateIntegrationRequest**](UpdateIntegrationRequest.md) | The request body for the integration update. | 

### Return type

[**Integration**](Integration.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

