# \OrgIntegrations

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Command**](OrgIntegrations.md#Command) | **Post** /org/{org_handle}/integration/{integration_handle}/command | Run org integration command
[**Create**](OrgIntegrations.md#Create) | **Post** /org/{org_handle}/integration | Create org integration
[**Delete**](OrgIntegrations.md#Delete) | **Delete** /org/{org_handle}/integration/{integration_handle} | Delete org integration
[**Get**](OrgIntegrations.md#Get) | **Get** /org/{org_handle}/integration/{integration_handle} | Get org integration
[**InstallGithubIntegration**](OrgIntegrations.md#InstallGithubIntegration) | **Get** /org/{org_handle}/integration/{integration_handle}/github/install | Install GitHub integration for an org
[**List**](OrgIntegrations.md#List) | **Get** /org/{org_handle}/integration | List org integrations
[**Test**](OrgIntegrations.md#Test) | **Post** /org/{org_handle}/integration/{integration_handle}/test | Test org integration
[**Update**](OrgIntegrations.md#Update) | **Patch** /org/{org_handle}/integration/{integration_handle} | Update org integration



## Command

> IntegrationCommandRequest Command(ctx, orgHandle, integrationHandle).Request(request).Execute()

Run org integration command



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
    orgHandle := "orgHandle_example" // string | The handle of the organization to which the integration belongs to.
    integrationHandle := "integrationHandle_example" // string | The handle of the integration for which we want to run the command.
    request := *openapiclient.NewWorkspaceCommandRequest("Command_example") // WorkspaceCommandRequest | The request body for the workspace command.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgIntegrations.Command(context.Background(), orgHandle, integrationHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgIntegrations.Command``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Command`: IntegrationCommandRequest
    fmt.Fprintf(os.Stdout, "Response from `OrgIntegrations.Command`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization to which the integration belongs to. | 
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

> Integration Create(ctx, orgHandle).Request(request).Execute()

Create org integration



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
    orgHandle := "orgHandle_example" // string | The handle of the org for which you want to create an integration.
    request := *openapiclient.NewCreateIntegrationRequest("Handle_example", "Type_example") // CreateIntegrationRequest | The request body for the integration to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgIntegrations.Create(context.Background(), orgHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgIntegrations.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Integration
    fmt.Fprintf(os.Stdout, "Response from `OrgIntegrations.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which you want to create an integration. | 

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

> Integration Delete(ctx, orgHandle, integrationHandle).Execute()

Delete org integration



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
    orgHandle := "orgHandle_example" // string | The handle of the organization on which the integration is configured.
    integrationHandle := "integrationHandle_example" // string | The handle of the integration which needs to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgIntegrations.Delete(context.Background(), orgHandle, integrationHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgIntegrations.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: Integration
    fmt.Fprintf(os.Stdout, "Response from `OrgIntegrations.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization on which the integration is configured. | 
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

> Integration Get(ctx, orgHandle, integrationHandle).Execute()

Get org integration



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
    orgHandle := "orgHandle_example" // string | The handle of the org to which the integration belongs to.
    integrationHandle := "integrationHandle_example" // string | The handle of the integration whose details need to be retrieved.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgIntegrations.Get(context.Background(), orgHandle, integrationHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgIntegrations.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Integration
    fmt.Fprintf(os.Stdout, "Response from `OrgIntegrations.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org to which the integration belongs to. | 
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

> InstallGithubIntegration(ctx, orgHandle, integrationHandle).Execute()

Install GitHub integration for an org



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
    orgHandle := "orgHandle_example" // string | The handle of the organization to which the integration belongs to.
    integrationHandle := "integrationHandle_example" // string | The handle of the github integration which needs to be installed.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgIntegrations.InstallGithubIntegration(context.Background(), orgHandle, integrationHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgIntegrations.InstallGithubIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization to which the integration belongs to. | 
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

> ListIntegrationsResponse List(ctx, orgHandle).Limit(limit).NextToken(nextToken).Execute()

List org integrations



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
    orgHandle := "orgHandle_example" // string | The handle of the org for which you want to list integrations.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgIntegrations.List(context.Background(), orgHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgIntegrations.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListIntegrationsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgIntegrations.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which you want to list integrations. | 

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

> Integration Test(ctx, orgHandle, integrationHandle).Execute()

Test org integration



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
    orgHandle := "orgHandle_example" // string | The handle of the org to which the integration belongs to.
    integrationHandle := "integrationHandle_example" // string | The handle of the integration to be tested. For integrations that are not yet created, use underscore `_` as the handle, else pass the handle of the existing integration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgIntegrations.Test(context.Background(), orgHandle, integrationHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgIntegrations.Test``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Test`: Integration
    fmt.Fprintf(os.Stdout, "Response from `OrgIntegrations.Test`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org to which the integration belongs to. | 
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

> Integration Update(ctx, orgHandle, integrationHandle).Request(request).Execute()

Update org integration



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
    orgHandle := "orgHandle_example" // string | The handle of the org on which the integration is configured.
    integrationHandle := "integrationHandle_example" // string | The handle of the integration whose details need to be updated.
    request := *openapiclient.NewUpdateIntegrationRequest() // UpdateIntegrationRequest | The request body for the integration update.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgIntegrations.Update(context.Background(), orgHandle, integrationHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgIntegrations.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Integration
    fmt.Fprintf(os.Stdout, "Response from `OrgIntegrations.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org on which the integration is configured. | 
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

