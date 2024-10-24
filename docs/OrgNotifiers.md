# \OrgNotifiers

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](OrgNotifiers.md#Create) | **Post** /org/{org_handle}/notifier | Create org notifier
[**Delete**](OrgNotifiers.md#Delete) | **Delete** /org/{org_handle}/notifier/{notifier_name} | Delete org notifier
[**Get**](OrgNotifiers.md#Get) | **Get** /org/{org_handle}/notifier/{notifier_name} | Get org notifier
[**List**](OrgNotifiers.md#List) | **Get** /org/{org_handle}/notifier | List org notifiers
[**Update**](OrgNotifiers.md#Update) | **Patch** /org/{org_handle}/notifier/{notifier_name} | Update org notifier



## Create

> Notifier Create(ctx, orgHandle).Request(request).Execute()

Create org notifier



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
    orgHandle := "orgHandle_example" // string | The handle of the org for which the notifier is to be created.
    request := *openapiclient.NewCreateNotifierRequest("Name_example", []map[string]interface{}{map[string]interface{}(123)}) // CreateNotifierRequest | The request body for the notifier to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgNotifiers.Create(context.Background(), orgHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgNotifiers.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Notifier
    fmt.Fprintf(os.Stdout, "Response from `OrgNotifiers.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which the notifier is to be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**CreateNotifierRequest**](CreateNotifierRequest.md) | The request body for the notifier to be created. | 

### Return type

[**Notifier**](Notifier.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Delete

> Notifier Delete(ctx, orgHandle, notifierName).Execute()

Delete org notifier



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
    orgHandle := "orgHandle_example" // string | The handle of the org for which the notifier is to be deleted.
    notifierName := "notifierName_example" // string | The name of the notifier to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgNotifiers.Delete(context.Background(), orgHandle, notifierName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgNotifiers.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: Notifier
    fmt.Fprintf(os.Stdout, "Response from `OrgNotifiers.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which the notifier is to be deleted. | 
**notifierName** | **string** | The name of the notifier to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Notifier**](Notifier.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Get

> Notifier Get(ctx, orgHandle, notifierName).Execute()

Get org notifier



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
    orgHandle := "orgHandle_example" // string | The handle of the org for which the notifier is to be retrieved.
    notifierName := "notifierName_example" // string | The name of the notifier to be retrieved.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgNotifiers.Get(context.Background(), orgHandle, notifierName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgNotifiers.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Notifier
    fmt.Fprintf(os.Stdout, "Response from `OrgNotifiers.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which the notifier is to be retrieved. | 
**notifierName** | **string** | The name of the notifier to be retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Notifier**](Notifier.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## List

> ListNotifiersResponse List(ctx, orgHandle).Limit(limit).NextToken(nextToken).Where(where).Execute()

List org notifiers



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
    orgHandle := "orgHandle_example" // string | The handle of the org for which the notifiers are to be listed.
    limit := int32(56) // int32 | The maximum number of items to return. (optional)
    nextToken := "nextToken_example" // string | The token to retrieve the next page. (optional)
    where := "where_example" // string | The filter to apply to the list. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgNotifiers.List(context.Background(), orgHandle).Limit(limit).NextToken(nextToken).Where(where).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgNotifiers.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListNotifiersResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgNotifiers.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which the notifiers are to be listed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of items to return. | 
 **nextToken** | **string** | The token to retrieve the next page. | 
 **where** | **string** | The filter to apply to the list. | 

### Return type

[**ListNotifiersResponse**](ListNotifiersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Update

> Notifier Update(ctx, orgHandle, notifierName).Request(request).Execute()

Update org notifier



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
    orgHandle := "orgHandle_example" // string | The handle of the org for which the notifier is to be updated.
    notifierName := "notifierName_example" // string | The name of the notifier to be updated.
    request := *openapiclient.NewUpdateNotifierRequest() // UpdateNotifierRequest | The request body for the notifier to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgNotifiers.Update(context.Background(), orgHandle, notifierName).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgNotifiers.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Notifier
    fmt.Fprintf(os.Stdout, "Response from `OrgNotifiers.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the org for which the notifier is to be updated. | 
**notifierName** | **string** | The name of the notifier to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**UpdateNotifierRequest**](UpdateNotifierRequest.md) | The request body for the notifier to be updated. | 

### Return type

[**Notifier**](Notifier.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

