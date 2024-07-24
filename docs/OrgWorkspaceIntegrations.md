# \OrgWorkspaceIntegrations

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](OrgWorkspaceIntegrations.md#Get) | **Get** /org/{org_handle}/workspace/{workspace_handle}/integration/{integration_handle} | Get org workspace integration
[**List**](OrgWorkspaceIntegrations.md#List) | **Get** /org/{org_handle}/workspace/{workspace_handle}/integration | List org workspace integrations
[**ListGithubRepositories**](OrgWorkspaceIntegrations.md#ListGithubRepositories) | **Get** /org/{org_handle}/workspace/{workspace_handle}/integration/{integration_handle}/github/repository | List org workspace integration github respositories
[**ListGithubRepositoryHeads**](OrgWorkspaceIntegrations.md#ListGithubRepositoryHeads) | **Get** /org/{org_handle}/workspace/{workspace_handle}/integration/{integration_handle}/github/repository/{repository_owner}/{repository_name}/head | List org workspace integration github respository heads



## Get

> Integration Get(ctx, orgHandle, workspaceHandle, integrationHandle).Execute()

Get org workspace integration



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
    orgHandle := "orgHandle_example" // string | The handle of the org to which the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace for which we want to get the integration details.
    integrationHandle := "integrationHandle_example" // string | The handle of the integration whose details need to be retrieved.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceIntegrations.Get(context.Background(), orgHandle, workspaceHandle, integrationHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceIntegrations.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Integration
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceIntegrations.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org to which the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace for which we want to get the integration details. | 
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


## List

> ListIntegrationsResponse List(ctx, orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

List org workspace integrations



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
    orgHandle := "orgHandle_example" // string | The handle of the organization to which the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace for whom you want to list available integrations.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceIntegrations.List(context.Background(), orgHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceIntegrations.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListIntegrationsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceIntegrations.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization to which the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace for whom you want to list available integrations. | 

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


## ListGithubRepositories

> ListGithubRepositoriesResponse ListGithubRepositories(ctx, orgHandle, workspaceHandle, integrationHandle).Limit(limit).NextToken(nextToken).Execute()

List org workspace integration github respositories



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
    orgHandle := "orgHandle_example" // string | The handle of the organization to which the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace for whom you want to list available integrations.
    integrationHandle := "integrationHandle_example" // string | The handle of the integration for which permitted repositories will be listed.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceIntegrations.ListGithubRepositories(context.Background(), orgHandle, workspaceHandle, integrationHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceIntegrations.ListGithubRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGithubRepositories`: ListGithubRepositoriesResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceIntegrations.ListGithubRepositories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization to which the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace for whom you want to list available integrations. | 
**integrationHandle** | **string** | The handle of the integration for which permitted repositories will be listed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGithubRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListGithubRepositoriesResponse**](ListGithubRepositoriesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGithubRepositoryHeads

> ListGithubRepositoryBranchesResponse ListGithubRepositoryHeads(ctx, orgHandle, workspaceHandle, integrationHandle, repositoryOwner, repositoryName).Limit(limit).NextToken(nextToken).Execute()

List org workspace integration github respository heads



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
    orgHandle := "orgHandle_example" // string | The handle of the organization to which the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace for whom you want to list available integrations.
    integrationHandle := "integrationHandle_example" // string | The handle of the integration for which heads will be listed.
    repositoryOwner := "repositoryOwner_example" // string | The owner of the repository.
    repositoryName := "repositoryName_example" // string | The name of the repository for which heads will be listed.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgWorkspaceIntegrations.ListGithubRepositoryHeads(context.Background(), orgHandle, workspaceHandle, integrationHandle, repositoryOwner, repositoryName).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgWorkspaceIntegrations.ListGithubRepositoryHeads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGithubRepositoryHeads`: ListGithubRepositoryBranchesResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgWorkspaceIntegrations.ListGithubRepositoryHeads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization to which the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace for whom you want to list available integrations. | 
**integrationHandle** | **string** | The handle of the integration for which heads will be listed. | 
**repositoryOwner** | **string** | The owner of the repository. | 
**repositoryName** | **string** | The name of the repository for which heads will be listed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGithubRepositoryHeadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListGithubRepositoryBranchesResponse**](ListGithubRepositoryBranchesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

