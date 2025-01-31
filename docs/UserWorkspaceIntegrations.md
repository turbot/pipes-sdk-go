# \UserWorkspaceIntegrations

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Get**](UserWorkspaceIntegrations.md#Get) | **Get** /user/{user_handle}/workspace/{workspace_handle}/integration/{integration_handle} | Get user workspace integration
[**List**](UserWorkspaceIntegrations.md#List) | **Get** /user/{user_handle}/workspace/{workspace_handle}/integration | List user workspace integrations
[**ListGithubRepositories**](UserWorkspaceIntegrations.md#ListGithubRepositories) | **Get** /user/{user_handle}/workspace/{workspace_handle}/integration/{integration_handle}/github/repository | List user workspace integration github respositories
[**ListGithubRepositoryHeads**](UserWorkspaceIntegrations.md#ListGithubRepositoryHeads) | **Get** /user/{user_handle}/workspace/{workspace_handle}/integration/{integration_handle}/github/repository/{repository_owner}/{repository_name}/head | List user workspace integration github respository heads
[**ListGitlabProjectBranches**](UserWorkspaceIntegrations.md#ListGitlabProjectBranches) | **Get** /user/{user_handle}/workspace/{workspace_handle}/integration/{integration_handle}/gitlab/project/{project_id}/branch | List user workspace integration gitlab project branches
[**ListGitlabProjects**](UserWorkspaceIntegrations.md#ListGitlabProjects) | **Get** /user/{user_handle}/workspace/{workspace_handle}/integration/{integration_handle}/gitlab/project | List user workspace integration gitlab projects



## Get

> Integration Get(ctx, userHandle, workspaceHandle, integrationHandle).Execute()

Get user workspace integration



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
    userHandle := "userHandle_example" // string | The handle of the user to which the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace for which we want to get the integration details.
    integrationHandle := "integrationHandle_example" // string | The handle of the integration whose details need to be retrieved.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceIntegrations.Get(context.Background(), userHandle, workspaceHandle, integrationHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceIntegrations.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Integration
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceIntegrations.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user to which the workspace belongs to. | 
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

> ListIntegrationsResponse List(ctx, userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()

List user workspace integrations



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
    userHandle := "userHandle_example" // string | The handle of the user to which the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace for whom you want to list available integrations.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceIntegrations.List(context.Background(), userHandle, workspaceHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceIntegrations.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListIntegrationsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceIntegrations.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user to which the workspace belongs to. | 
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

> ListGithubRepositoriesResponse ListGithubRepositories(ctx, userHandle, workspaceHandle, integrationHandle).Limit(limit).NextToken(nextToken).Execute()

List user workspace integration github respositories



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
    userHandle := "userHandle_example" // string | The handle of the user to which the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace for whom you want to list available integrations.
    integrationHandle := "integrationHandle_example" // string | The handle of the integration for which permitted repositories will be listed.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceIntegrations.ListGithubRepositories(context.Background(), userHandle, workspaceHandle, integrationHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceIntegrations.ListGithubRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGithubRepositories`: ListGithubRepositoriesResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceIntegrations.ListGithubRepositories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user to which the workspace belongs to. | 
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

> ListGithubRepositoryBranchesResponse ListGithubRepositoryHeads(ctx, userHandle, workspaceHandle, integrationHandle, repositoryOwner, repositoryName).Limit(limit).NextToken(nextToken).Execute()

List user workspace integration github respository heads



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
    userHandle := "userHandle_example" // string | The handle of the user to which the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace for whom you want to list available integrations.
    integrationHandle := "integrationHandle_example" // string | The handle of the integration for which heads will be listed.
    repositoryOwner := "repositoryOwner_example" // string | The owner of the repository.
    repositoryName := "repositoryName_example" // string | The name of the repository for which heads will be listed.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceIntegrations.ListGithubRepositoryHeads(context.Background(), userHandle, workspaceHandle, integrationHandle, repositoryOwner, repositoryName).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceIntegrations.ListGithubRepositoryHeads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGithubRepositoryHeads`: ListGithubRepositoryBranchesResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceIntegrations.ListGithubRepositoryHeads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user to which the workspace belongs to. | 
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


## ListGitlabProjectBranches

> ListGitlabProjectBranchesResponse ListGitlabProjectBranches(ctx, userHandle, workspaceHandle, integrationHandle, projectId).Limit(limit).NextToken(nextToken).Execute()

List user workspace integration gitlab project branches



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
    userHandle := "userHandle_example" // string | The handle of the user to which the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace for whom you want to list available integrations.
    integrationHandle := "integrationHandle_example" // string | The handle of the integration for which heads will be listed.
    projectId := "projectId_example" // string | The id of the project including the namespace and the name.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceIntegrations.ListGitlabProjectBranches(context.Background(), userHandle, workspaceHandle, integrationHandle, projectId).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceIntegrations.ListGitlabProjectBranches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGitlabProjectBranches`: ListGitlabProjectBranchesResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceIntegrations.ListGitlabProjectBranches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user to which the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace for whom you want to list available integrations. | 
**integrationHandle** | **string** | The handle of the integration for which heads will be listed. | 
**projectId** | **string** | The id of the project including the namespace and the name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGitlabProjectBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListGitlabProjectBranchesResponse**](ListGitlabProjectBranchesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGitlabProjects

> ListGitlabProjectsResponse ListGitlabProjects(ctx, userHandle, workspaceHandle, integrationHandle).Limit(limit).NextToken(nextToken).Execute()

List user workspace integration gitlab projects



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
    userHandle := "userHandle_example" // string | The handle of the user to which the workspace belongs to.
    workspaceHandle := "workspaceHandle_example" // string | The handle of the workspace for whom you want to list available integrations.
    integrationHandle := "integrationHandle_example" // string | The handle of the integration for which permitted repositories will be listed.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserWorkspaceIntegrations.ListGitlabProjects(context.Background(), userHandle, workspaceHandle, integrationHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserWorkspaceIntegrations.ListGitlabProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGitlabProjects`: ListGitlabProjectsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserWorkspaceIntegrations.ListGitlabProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user to which the workspace belongs to. | 
**workspaceHandle** | **string** | The handle of the workspace for whom you want to list available integrations. | 
**integrationHandle** | **string** | The handle of the integration for which permitted repositories will be listed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGitlabProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListGitlabProjectsResponse**](ListGitlabProjectsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

