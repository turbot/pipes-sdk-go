# \UserBilling

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBillingIntent**](UserBilling.md#CreateBillingIntent) | **Post** /user/{user_handle}/billing/payment_method/intent | Create Stripe billing intent
[**CreateUserSubscription**](UserBilling.md#CreateUserSubscription) | **Post** /user/{user_handle}/billing/subscription | Create a new subscription
[**DeleteUserBillingPaymentMethod**](UserBilling.md#DeleteUserBillingPaymentMethod) | **Delete** /user/{user_handle}/billing/payment_method/{payment_method_id} | Delete user billing payment method.
[**DeleteUserBillingSubscription**](UserBilling.md#DeleteUserBillingSubscription) | **Delete** /user/{user_handle}/billing/subscription/{subscription_id} | Delete user subscription
[**GetUserBillingInvoice**](UserBilling.md#GetUserBillingInvoice) | **Get** /user/{user_handle}/billing/invoice/{invoice_id} | Get a user invoice
[**GetUserBillingPaymentMethod**](UserBilling.md#GetUserBillingPaymentMethod) | **Get** /user/{user_handle}/billing/payment_method/{payment_method_id} | Get user billing payment method.
[**GetUserBillingSubscription**](UserBilling.md#GetUserBillingSubscription) | **Get** /user/{user_handle}/billing/subscription/{subscription_id} | Get user subscription
[**GetUserBillingUpcomingInvoice**](UserBilling.md#GetUserBillingUpcomingInvoice) | **Get** /user/{user_handle}/billing/stripe/invoice/upcoming | Get upcoming user invoice
[**ListUserBillingInvoices**](UserBilling.md#ListUserBillingInvoices) | **Get** /user/{user_handle}/billing/invoice | List user invoices
[**ListUserBillingPaymentMethod**](UserBilling.md#ListUserBillingPaymentMethod) | **Get** /user/{user_handle}/billing/payment_method | List user billing payment methods.
[**ListUserSubscription**](UserBilling.md#ListUserSubscription) | **Get** /user/{user_handle}/billing/subscription | List user subscriptions
[**UpdateUserBillingPaymentMethod**](UserBilling.md#UpdateUserBillingPaymentMethod) | **Patch** /user/{user_handle}/billing/payment_method/{payment_method_id} | Update user billing payment method.



## CreateBillingIntent

> BillingSetupIntent CreateBillingIntent(ctx, userHandle).Execute()

Create Stripe billing intent



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
    userHandle := "userHandle_example" // string | The handle of a user where the billing intent will be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserBilling.CreateBillingIntent(context.Background(), userHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserBilling.CreateBillingIntent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBillingIntent`: BillingSetupIntent
    fmt.Fprintf(os.Stdout, "Response from `UserBilling.CreateBillingIntent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of a user where the billing intent will be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBillingIntentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillingSetupIntent**](BillingSetupIntent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserSubscription

> BillingSubscription CreateUserSubscription(ctx, userHandle).Execute()

Create a new subscription



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
    userHandle := "userHandle_example" // string | The handle of a user where the subscription will be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserBilling.CreateUserSubscription(context.Background(), userHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserBilling.CreateUserSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserSubscription`: BillingSubscription
    fmt.Fprintf(os.Stdout, "Response from `UserBilling.CreateUserSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of a user where the subscription will be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillingSubscription**](BillingSubscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserBillingPaymentMethod

> PaymentMethod DeleteUserBillingPaymentMethod(ctx, userHandle, paymentMethodId).Execute()

Delete user billing payment method.



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
    userHandle := "userHandle_example" // string | The handle of the user to delete the payment method for.
    paymentMethodId := "paymentMethodId_example" // string | The user payment method id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserBilling.DeleteUserBillingPaymentMethod(context.Background(), userHandle, paymentMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserBilling.DeleteUserBillingPaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUserBillingPaymentMethod`: PaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `UserBilling.DeleteUserBillingPaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user to delete the payment method for. | 
**paymentMethodId** | **string** | The user payment method id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserBillingPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaymentMethod**](PaymentMethod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserBillingSubscription

> BillingSubscription DeleteUserBillingSubscription(ctx, userHandle, subscriptionId).Execute()

Delete user subscription



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
    userHandle := "userHandle_example" // string | The handle of a user to delete the subscription for.
    subscriptionId := "subscriptionId_example" // string | The user subscription id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserBilling.DeleteUserBillingSubscription(context.Background(), userHandle, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserBilling.DeleteUserBillingSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUserBillingSubscription`: BillingSubscription
    fmt.Fprintf(os.Stdout, "Response from `UserBilling.DeleteUserBillingSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of a user to delete the subscription for. | 
**subscriptionId** | **string** | The user subscription id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserBillingSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BillingSubscription**](BillingSubscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingInvoice

> BillingInvoice GetUserBillingInvoice(ctx, userHandle, invoiceId).Execute()

Get a user invoice



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
    userHandle := "userHandle_example" // string | Specify the handle of a user to retrieve an invoice for.
    invoiceId := "invoiceId_example" // string | The user invoice id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserBilling.GetUserBillingInvoice(context.Background(), userHandle, invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserBilling.GetUserBillingInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserBillingInvoice`: BillingInvoice
    fmt.Fprintf(os.Stdout, "Response from `UserBilling.GetUserBillingInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | Specify the handle of a user to retrieve an invoice for. | 
**invoiceId** | **string** | The user invoice id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BillingInvoice**](BillingInvoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingPaymentMethod

> PaymentMethod GetUserBillingPaymentMethod(ctx, userHandle, paymentMethodId).Execute()

Get user billing payment method.



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
    userHandle := "userHandle_example" // string | The handle of the user to get the payment method for.
    paymentMethodId := "paymentMethodId_example" // string | The user payment method id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserBilling.GetUserBillingPaymentMethod(context.Background(), userHandle, paymentMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserBilling.GetUserBillingPaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserBillingPaymentMethod`: PaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `UserBilling.GetUserBillingPaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user to get the payment method for. | 
**paymentMethodId** | **string** | The user payment method id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PaymentMethod**](PaymentMethod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingSubscription

> BillingSubscription GetUserBillingSubscription(ctx, userHandle, subscriptionId).Execute()

Get user subscription



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
    userHandle := "userHandle_example" // string | The handle of a user.
    subscriptionId := "subscriptionId_example" // string | The user subscription id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserBilling.GetUserBillingSubscription(context.Background(), userHandle, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserBilling.GetUserBillingSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserBillingSubscription`: BillingSubscription
    fmt.Fprintf(os.Stdout, "Response from `UserBilling.GetUserBillingSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of a user. | 
**subscriptionId** | **string** | The user subscription id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BillingSubscription**](BillingSubscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBillingUpcomingInvoice

> BillingInvoice GetUserBillingUpcomingInvoice(ctx, userHandle).Execute()

Get upcoming user invoice



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
    userHandle := "userHandle_example" // string | The handle of the user to get the upcoming invoice for.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserBilling.GetUserBillingUpcomingInvoice(context.Background(), userHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserBilling.GetUserBillingUpcomingInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserBillingUpcomingInvoice`: BillingInvoice
    fmt.Fprintf(os.Stdout, "Response from `UserBilling.GetUserBillingUpcomingInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user to get the upcoming invoice for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBillingUpcomingInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillingInvoice**](BillingInvoice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserBillingInvoices

> ListBillingInvoicesResponse ListUserBillingInvoices(ctx, userHandle).Limit(limit).NextToken(nextToken).Execute()

List user invoices



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
    userHandle := "userHandle_example" // string | The handle of a user to list invoices for.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserBilling.ListUserBillingInvoices(context.Background(), userHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserBilling.ListUserBillingInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserBillingInvoices`: ListBillingInvoicesResponse
    fmt.Fprintf(os.Stdout, "Response from `UserBilling.ListUserBillingInvoices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of a user to list invoices for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserBillingInvoicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. | 

### Return type

[**ListBillingInvoicesResponse**](ListBillingInvoicesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserBillingPaymentMethod

> ListBillingPaymentMethodsResponse ListUserBillingPaymentMethod(ctx, userHandle).Execute()

List user billing payment methods.



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
    userHandle := "userHandle_example" // string | The handle of the user to list the payment methods for.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserBilling.ListUserBillingPaymentMethod(context.Background(), userHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserBilling.ListUserBillingPaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserBillingPaymentMethod`: ListBillingPaymentMethodsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserBilling.ListUserBillingPaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user to list the payment methods for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserBillingPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListBillingPaymentMethodsResponse**](ListBillingPaymentMethodsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserSubscription

> ListSubscriptionsResponse ListUserSubscription(ctx, userHandle).Execute()

List user subscriptions



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
    userHandle := "userHandle_example" // string | The handle of the user to retrieve the subscriptions for.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserBilling.ListUserSubscription(context.Background(), userHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserBilling.ListUserSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserSubscription`: ListSubscriptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserBilling.ListUserSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user to retrieve the subscriptions for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListSubscriptionsResponse**](ListSubscriptionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserBillingPaymentMethod

> PaymentMethod UpdateUserBillingPaymentMethod(ctx, userHandle, paymentMethodId).Request(request).Execute()

Update user billing payment method.



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
    userHandle := "userHandle_example" // string | The handle of the user to update the payment method for.
    paymentMethodId := "paymentMethodId_example" // string | The user payment method id.
    request := *openapiclient.NewUpdateBillingPaymentMethodRequest() // UpdateBillingPaymentMethodRequest | The request body for the user billing payment method update.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserBilling.UpdateUserBillingPaymentMethod(context.Background(), userHandle, paymentMethodId).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserBilling.UpdateUserBillingPaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserBillingPaymentMethod`: PaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `UserBilling.UpdateUserBillingPaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user to update the payment method for. | 
**paymentMethodId** | **string** | The user payment method id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserBillingPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**UpdateBillingPaymentMethodRequest**](UpdateBillingPaymentMethodRequest.md) | The request body for the user billing payment method update. | 

### Return type

[**PaymentMethod**](PaymentMethod.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

