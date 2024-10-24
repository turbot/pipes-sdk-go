# \TenantBilling

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBillingIntent**](TenantBilling.md#CreateBillingIntent) | **Post** /tenant/{tenant_handle}/billing/payment_method/intent | Create tenant Stripe billing intent
[**CreateTenantSubscription**](TenantBilling.md#CreateTenantSubscription) | **Post** /tenant/{tenant_handle}/billing/subscription | Create a new subscription
[**DeleteTenantBillingPaymentMethod**](TenantBilling.md#DeleteTenantBillingPaymentMethod) | **Delete** /tenant/{tenant_handle}/billing/payment_method/{payment_method_id} | Delete tenant billing payment method.
[**DeleteTenantBillingSubscription**](TenantBilling.md#DeleteTenantBillingSubscription) | **Delete** /tenant/{tenant_handle}/billing/subscription/{subscription_id} | Delete tenant subscription
[**GetPlan**](TenantBilling.md#GetPlan) | **Get** /tenant/{tenant_handle}/billing/plan | Get tenant billing plan.
[**GetTenantBillingInvoice**](TenantBilling.md#GetTenantBillingInvoice) | **Get** /tenant/{tenant_handle}/billing/invoice/{invoice_id} | Get a tenant invoice
[**GetTenantBillingPaymentMethod**](TenantBilling.md#GetTenantBillingPaymentMethod) | **Get** /tenant/{tenant_handle}/billing/payment_method/{payment_method_id} | Get tenant billing payment method.
[**GetTenantBillingSubscription**](TenantBilling.md#GetTenantBillingSubscription) | **Get** /tenant/{tenant_handle}/billing/subscription/{subscription_id} | Get tenant subscription
[**GetTenantBillingUpcomingInvoice**](TenantBilling.md#GetTenantBillingUpcomingInvoice) | **Get** /tenant/{tenant_handle}/billing/stripe/invoice/upcoming | Get upcoming invoice
[**ListTenantBillingInvoices**](TenantBilling.md#ListTenantBillingInvoices) | **Get** /tenant/{tenant_handle}/billing/invoice | List tenant invoices
[**ListTenantBillingPaymentMethod**](TenantBilling.md#ListTenantBillingPaymentMethod) | **Get** /tenant/{tenant_handle}/billing/payment_method | List tenant billing payment methods.
[**ListTenantSubscription**](TenantBilling.md#ListTenantSubscription) | **Get** /tenant/{tenant_handle}/billing/subscription | List tenant subscriptions
[**UpdateTenantBillingPaymentMethod**](TenantBilling.md#UpdateTenantBillingPaymentMethod) | **Patch** /tenant/{tenant_handle}/billing/payment_method/{payment_method_id} | Update tenant billing payment method.
[**UpdateTenantBillingSubscription**](TenantBilling.md#UpdateTenantBillingSubscription) | **Patch** /tenant/{tenant_handle}/billing/subscription/{subscription_id} | Update tenant subscription



## CreateBillingIntent

> BillingSetupIntent CreateBillingIntent(ctx, tenantHandle).Execute()

Create tenant Stripe billing intent



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
    tenantHandle := "tenantHandle_example" // string | The handle of an organization where the billing intent will be created

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantBilling.CreateBillingIntent(context.Background(), tenantHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantBilling.CreateBillingIntent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBillingIntent`: BillingSetupIntent
    fmt.Fprintf(os.Stdout, "Response from `TenantBilling.CreateBillingIntent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantHandle** | **string** | The handle of an organization where the billing intent will be created | 

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


## CreateTenantSubscription

> BillingSubscription CreateTenantSubscription(ctx, tenantHandle).Execute()

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
    tenantHandle := "tenantHandle_example" // string | The handle of the tenant where the subscription will be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantBilling.CreateTenantSubscription(context.Background(), tenantHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantBilling.CreateTenantSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTenantSubscription`: BillingSubscription
    fmt.Fprintf(os.Stdout, "Response from `TenantBilling.CreateTenantSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantHandle** | **string** | The handle of the tenant where the subscription will be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantSubscriptionRequest struct via the builder pattern


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


## DeleteTenantBillingPaymentMethod

> PaymentMethod DeleteTenantBillingPaymentMethod(ctx, tenantHandle, paymentMethodId).Execute()

Delete tenant billing payment method.



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
    tenantHandle := "tenantHandle_example" // string | The handle of the tenant to delete the payment method for
    paymentMethodId := "paymentMethodId_example" // string | The org payment method id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantBilling.DeleteTenantBillingPaymentMethod(context.Background(), tenantHandle, paymentMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantBilling.DeleteTenantBillingPaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTenantBillingPaymentMethod`: PaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `TenantBilling.DeleteTenantBillingPaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantHandle** | **string** | The handle of the tenant to delete the payment method for | 
**paymentMethodId** | **string** | The org payment method id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantBillingPaymentMethodRequest struct via the builder pattern


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


## DeleteTenantBillingSubscription

> BillingSubscription DeleteTenantBillingSubscription(ctx, tenantHandle, subscriptionId).Execute()

Delete tenant subscription



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
    tenantHandle := "tenantHandle_example" // string | The handle of the tenant to delete the subscription for.
    subscriptionId := "subscriptionId_example" // string | The tenant subscription id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantBilling.DeleteTenantBillingSubscription(context.Background(), tenantHandle, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantBilling.DeleteTenantBillingSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteTenantBillingSubscription`: BillingSubscription
    fmt.Fprintf(os.Stdout, "Response from `TenantBilling.DeleteTenantBillingSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantHandle** | **string** | The handle of the tenant to delete the subscription for. | 
**subscriptionId** | **string** | The tenant subscription id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantBillingSubscriptionRequest struct via the builder pattern


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


## GetPlan

> TenantPlan GetPlan(ctx, tenantHandle).Execute()

Get tenant billing plan.



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
    tenantHandle := "tenantHandle_example" // string | Specify the handle of the tenant to retrieve the billing plan for.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantBilling.GetPlan(context.Background(), tenantHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantBilling.GetPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlan`: TenantPlan
    fmt.Fprintf(os.Stdout, "Response from `TenantBilling.GetPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantHandle** | **string** | Specify the handle of the tenant to retrieve the billing plan for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TenantPlan**](TenantPlan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantBillingInvoice

> BillingInvoice GetTenantBillingInvoice(ctx, tenantHandle, invoiceId).Execute()

Get a tenant invoice



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
    tenantHandle := "tenantHandle_example" // string | The handle of an organization
    invoiceId := "invoiceId_example" // string | The org invoice id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantBilling.GetTenantBillingInvoice(context.Background(), tenantHandle, invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantBilling.GetTenantBillingInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantBillingInvoice`: BillingInvoice
    fmt.Fprintf(os.Stdout, "Response from `TenantBilling.GetTenantBillingInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantHandle** | **string** | The handle of an organization | 
**invoiceId** | **string** | The org invoice id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantBillingInvoiceRequest struct via the builder pattern


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


## GetTenantBillingPaymentMethod

> PaymentMethod GetTenantBillingPaymentMethod(ctx, tenantHandle, paymentMethodId).Execute()

Get tenant billing payment method.



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
    tenantHandle := "tenantHandle_example" // string | The handle of the tenant to get the payment method for
    paymentMethodId := "paymentMethodId_example" // string | The org payment method id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantBilling.GetTenantBillingPaymentMethod(context.Background(), tenantHandle, paymentMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantBilling.GetTenantBillingPaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantBillingPaymentMethod`: PaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `TenantBilling.GetTenantBillingPaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantHandle** | **string** | The handle of the tenant to get the payment method for | 
**paymentMethodId** | **string** | The org payment method id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantBillingPaymentMethodRequest struct via the builder pattern


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


## GetTenantBillingSubscription

> BillingSubscription GetTenantBillingSubscription(ctx, tenantHandle, subscriptionId).Execute()

Get tenant subscription



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
    tenantHandle := "tenantHandle_example" // string | The handle of a the tenant to get the subscription for.
    subscriptionId := "subscriptionId_example" // string | The tenant subscription id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantBilling.GetTenantBillingSubscription(context.Background(), tenantHandle, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantBilling.GetTenantBillingSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantBillingSubscription`: BillingSubscription
    fmt.Fprintf(os.Stdout, "Response from `TenantBilling.GetTenantBillingSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantHandle** | **string** | The handle of a the tenant to get the subscription for. | 
**subscriptionId** | **string** | The tenant subscription id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantBillingSubscriptionRequest struct via the builder pattern


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


## GetTenantBillingUpcomingInvoice

> BillingInvoice GetTenantBillingUpcomingInvoice(ctx, tenantHandle).Execute()

Get upcoming invoice



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
    tenantHandle := "tenantHandle_example" // string | The handle of the tenant

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantBilling.GetTenantBillingUpcomingInvoice(context.Background(), tenantHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantBilling.GetTenantBillingUpcomingInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantBillingUpcomingInvoice`: BillingInvoice
    fmt.Fprintf(os.Stdout, "Response from `TenantBilling.GetTenantBillingUpcomingInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantHandle** | **string** | The handle of the tenant | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantBillingUpcomingInvoiceRequest struct via the builder pattern


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


## ListTenantBillingInvoices

> ListBillingInvoicesResponse ListTenantBillingInvoices(ctx, tenantHandle).Limit(limit).NextToken(nextToken).Execute()

List tenant invoices



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
    tenantHandle := "tenantHandle_example" // string | The handle of the tenant to list invoices for
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantBilling.ListTenantBillingInvoices(context.Background(), tenantHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantBilling.ListTenantBillingInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTenantBillingInvoices`: ListBillingInvoicesResponse
    fmt.Fprintf(os.Stdout, "Response from `TenantBilling.ListTenantBillingInvoices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantHandle** | **string** | The handle of the tenant to list invoices for | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTenantBillingInvoicesRequest struct via the builder pattern


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


## ListTenantBillingPaymentMethod

> ListBillingPaymentMethodsResponse ListTenantBillingPaymentMethod(ctx, tenantHandle).Execute()

List tenant billing payment methods.



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
    tenantHandle := "tenantHandle_example" // string | The handle of the tenant to list payment methods for

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantBilling.ListTenantBillingPaymentMethod(context.Background(), tenantHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantBilling.ListTenantBillingPaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTenantBillingPaymentMethod`: ListBillingPaymentMethodsResponse
    fmt.Fprintf(os.Stdout, "Response from `TenantBilling.ListTenantBillingPaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantHandle** | **string** | The handle of the tenant to list payment methods for | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTenantBillingPaymentMethodRequest struct via the builder pattern


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


## ListTenantSubscription

> ListSubscriptionsResponse ListTenantSubscription(ctx, tenantHandle).Execute()

List tenant subscriptions



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
    tenantHandle := "tenantHandle_example" // string | The handle of the tenant to list subscriptions for.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantBilling.ListTenantSubscription(context.Background(), tenantHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantBilling.ListTenantSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTenantSubscription`: ListSubscriptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `TenantBilling.ListTenantSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantHandle** | **string** | The handle of the tenant to list subscriptions for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTenantSubscriptionRequest struct via the builder pattern


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


## UpdateTenantBillingPaymentMethod

> PaymentMethod UpdateTenantBillingPaymentMethod(ctx, tenantHandle, paymentMethodId).Request(request).Execute()

Update tenant billing payment method.



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
    tenantHandle := "tenantHandle_example" // string | The handle of the tenant to update the payment method for
    paymentMethodId := "paymentMethodId_example" // string | The org payment method id.
    request := *openapiclient.NewUpdateBillingPaymentMethodRequest() // UpdateBillingPaymentMethodRequest | The request body for the tenant billing payment method update.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantBilling.UpdateTenantBillingPaymentMethod(context.Background(), tenantHandle, paymentMethodId).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantBilling.UpdateTenantBillingPaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTenantBillingPaymentMethod`: PaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `TenantBilling.UpdateTenantBillingPaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantHandle** | **string** | The handle of the tenant to update the payment method for | 
**paymentMethodId** | **string** | The org payment method id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantBillingPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**UpdateBillingPaymentMethodRequest**](UpdateBillingPaymentMethodRequest.md) | The request body for the tenant billing payment method update. | 

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


## UpdateTenantBillingSubscription

> BillingSubscription UpdateTenantBillingSubscription(ctx, tenantHandle, subscriptionId).Request(request).Execute()

Update tenant subscription



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
    tenantHandle := "tenantHandle_example" // string | The handle of the tenant to update the subscription for.
    subscriptionId := "subscriptionId_example" // string | The tenant subscription id.
    request := *openapiclient.NewUpdateBillingSubscriptionRequest() // UpdateBillingSubscriptionRequest | Request body to update the tenant subscription

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TenantBilling.UpdateTenantBillingSubscription(context.Background(), tenantHandle, subscriptionId).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TenantBilling.UpdateTenantBillingSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTenantBillingSubscription`: BillingSubscription
    fmt.Fprintf(os.Stdout, "Response from `TenantBilling.UpdateTenantBillingSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantHandle** | **string** | The handle of the tenant to update the subscription for. | 
**subscriptionId** | **string** | The tenant subscription id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantBillingSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**UpdateBillingSubscriptionRequest**](UpdateBillingSubscriptionRequest.md) | Request body to update the tenant subscription | 

### Return type

[**BillingSubscription**](BillingSubscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

