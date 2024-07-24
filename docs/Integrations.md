# \Integrations

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfirmGithubIntegration**](Integrations.md#ConfirmGithubIntegration) | **Get** /integration/github/confirm | Confirm GitHub integration installation



## ConfirmGithubIntegration

> ConfirmGithubIntegration(ctx).State(state).InstallationId(installationId).Execute()

Confirm GitHub integration installation



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
    state := "state_example" // string | The base64 encoded JSON state for the request state.
    installationId := "installationId_example" // string | The GitHub installation id which needs to be confirmed.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.Integrations.ConfirmGithubIntegration(context.Background()).State(state).InstallationId(installationId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Integrations.ConfirmGithubIntegration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfirmGithubIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** | The base64 encoded JSON state for the request state. | 
 **installationId** | **string** | The GitHub installation id which needs to be confirmed. | 

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

