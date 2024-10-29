# \UserNotifiers

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](UserNotifiers.md#Create) | **Post** /user/{user_handle}/notifier | Create user notifier
[**Delete**](UserNotifiers.md#Delete) | **Delete** /user/{user_handle}/notifier/{notifier_name} | Delete user notifier
[**Get**](UserNotifiers.md#Get) | **Get** /user/{user_handle}/notifier/{notifier_name} | Get user notifier
[**List**](UserNotifiers.md#List) | **Get** /user/{user_handle}/notifier | List user notifiers
[**Update**](UserNotifiers.md#Update) | **Patch** /user/{user_handle}/notifier/{notifier_name} | Update user notifier



## Create

> Notifier Create(ctx, userHandle).Request(request).Execute()

Create user notifier



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
    userHandle := "userHandle_example" // string | The handle of the user for which the notifier is to be created.
    request := *openapiclient.NewCreateNotifierRequest("Name_example", []map[string]interface{}{map[string]interface{}(123)}) // CreateNotifierRequest | The request body for the notifier to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserNotifiers.Create(context.Background(), userHandle).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserNotifiers.Create``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Create`: Notifier
    fmt.Fprintf(os.Stdout, "Response from `UserNotifiers.Create`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which the notifier is to be created. | 

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

> Notifier Delete(ctx, userHandle, notifierName).Execute()

Delete user notifier



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
    userHandle := "userHandle_example" // string | The handle of the user for which the notifier is to be deleted.
    notifierName := "notifierName_example" // string | The name of the notifier to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserNotifiers.Delete(context.Background(), userHandle, notifierName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserNotifiers.Delete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Delete`: Notifier
    fmt.Fprintf(os.Stdout, "Response from `UserNotifiers.Delete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which the notifier is to be deleted. | 
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

> Notifier Get(ctx, userHandle, notifierName).Execute()

Get user notifier



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
    userHandle := "userHandle_example" // string | The handle of the user for which the notifier is to be retrieved.
    notifierName := "notifierName_example" // string | The name of the notifier to be retrieved.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserNotifiers.Get(context.Background(), userHandle, notifierName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserNotifiers.Get``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Get`: Notifier
    fmt.Fprintf(os.Stdout, "Response from `UserNotifiers.Get`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which the notifier is to be retrieved. | 
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

> ListNotifiersResponse List(ctx, userHandle).Limit(limit).NextToken(nextToken).Where(where).Execute()

List user notifiers



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
    userHandle := "userHandle_example" // string | The handle of the user for which the notifiers are to be listed.
    limit := int32(56) // int32 | The maximum number of items to return. (optional)
    nextToken := "nextToken_example" // string | The token to retrieve the next page. (optional)
    where := "where_example" // string | The filter to apply to the list. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserNotifiers.List(context.Background(), userHandle).Limit(limit).NextToken(nextToken).Where(where).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserNotifiers.List``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `List`: ListNotifiersResponse
    fmt.Fprintf(os.Stdout, "Response from `UserNotifiers.List`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which the notifiers are to be listed. | 

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

> Notifier Update(ctx, userHandle, notifierName).Request(request).Execute()

Update user notifier



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
    userHandle := "userHandle_example" // string | The handle of the user for which the notifier is to be updated.
    notifierName := "notifierName_example" // string | The name of the notifier to be updated.
    request := *openapiclient.NewUpdateNotifierRequest() // UpdateNotifierRequest | The request body for the notifier to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserNotifiers.Update(context.Background(), userHandle, notifierName).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserNotifiers.Update``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Update`: Notifier
    fmt.Fprintf(os.Stdout, "Response from `UserNotifiers.Update`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which the notifier is to be updated. | 
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

