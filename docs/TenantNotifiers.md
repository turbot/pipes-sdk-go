# \TenantNotifiers

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](TenantNotifiers.md#Create) | **Post** /notifier | Create tenant notifier
[**Delete**](TenantNotifiers.md#Delete) | **Delete** /notifier/{notifier_name} | Delete tenant notifier
[**Get**](TenantNotifiers.md#Get) | **Get** /notifier/{notifier_name} | Get tenant notifier
[**List**](TenantNotifiers.md#List) | **Get** /notifier | List tenant notifiers
[**Update**](TenantNotifiers.md#Update) | **Patch** /notifier/{notifier_name} | Update tenant notifier



## Create

> Notifier Create(ctx).Request(request).Execute()

Create tenant notifier



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
    request := *openapiclient.NewCreateNotifierRequest("Name_example", []map[string]interface{}{map[string]interface{}(123)}) // CreateNotifierRequest | The request body for the notifier to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantNotifiers.Create(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantNotifiers.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Notifier
    fmt.Fprintf(os.Stdout, "Response from `TenantNotifiers.Create`: %v\n", resp)
}
```

### Path Parameters



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

> Notifier Delete(ctx, notifierName).Execute()

Delete tenant notifier



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
    notifierName := "notifierName_example" // string | The name of the notifier to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantNotifiers.Delete(context.Background(), notifierName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantNotifiers.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: Notifier
    fmt.Fprintf(os.Stdout, "Response from `TenantNotifiers.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

> Notifier Get(ctx, notifierName).Execute()

Get tenant notifier



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
    notifierName := "notifierName_example" // string | The name of the notifier to be retrieved.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantNotifiers.Get(context.Background(), notifierName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantNotifiers.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Notifier
    fmt.Fprintf(os.Stdout, "Response from `TenantNotifiers.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

> ListNotifiersResponse List(ctx).Limit(limit).NextToken(nextToken).Where(where).Execute()

List tenant notifiers



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
    limit := int32(56) // int32 | The maximum number of items to return. (optional)
    nextToken := "nextToken_example" // string | The token to retrieve the next page. (optional)
    where := "where_example" // string | The filter to apply to the list. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantNotifiers.List(context.Background()).Limit(limit).NextToken(nextToken).Where(where).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantNotifiers.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListNotifiersResponse
    fmt.Fprintf(os.Stdout, "Response from `TenantNotifiers.List`: %v\n", resp)
}
```

### Path Parameters



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

> Notifier Update(ctx, notifierName).Request(request).Execute()

Update tenant notifier



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
    notifierName := "notifierName_example" // string | The name of the notifier to be updated.
    request := *openapiclient.NewUpdateNotifierRequest() // UpdateNotifierRequest | The request body for the notifier to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantNotifiers.Update(context.Background(), notifierName).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantNotifiers.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Notifier
    fmt.Fprintf(os.Stdout, "Response from `TenantNotifiers.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
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

