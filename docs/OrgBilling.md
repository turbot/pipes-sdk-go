# \OrgBilling

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBillingIntent**](OrgBilling.md#CreateBillingIntent) | **Post** /org/{org_handle}/billing/payment_method/intent | Create Stripe billing intent
[**CreateOrgSubscription**](OrgBilling.md#CreateOrgSubscription) | **Post** /org/{org_handle}/billing/subscription | Create a new subscription
[**DeleteOrgBillingPaymentMethod**](OrgBilling.md#DeleteOrgBillingPaymentMethod) | **Delete** /org/{org_handle}/billing/payment_method/{payment_method_id} | Delete org billing payment method.
[**DeleteOrgBillingSubscription**](OrgBilling.md#DeleteOrgBillingSubscription) | **Delete** /org/{org_handle}/billing/subscription/{subscription_id} | Delete org subscription
[**GetOrgBillingInvoice**](OrgBilling.md#GetOrgBillingInvoice) | **Get** /org/{org_handle}/billing/invoice/{invoice_id} | Get an organization invoice
[**GetOrgBillingPaymentMethod**](OrgBilling.md#GetOrgBillingPaymentMethod) | **Get** /org/{org_handle}/billing/payment_method/{payment_method_id} | Get org billing payment method.
[**GetOrgBillingSubscription**](OrgBilling.md#GetOrgBillingSubscription) | **Get** /org/{org_handle}/billing/subscription/{subscription_id} | Get org subscription
[**GetOrgBillingUpcomingInvoice**](OrgBilling.md#GetOrgBillingUpcomingInvoice) | **Get** /org/{org_handle}/billing/stripe/invoice/upcoming | Get upcoming invoice
[**ListOrgBillingInvoices**](OrgBilling.md#ListOrgBillingInvoices) | **Get** /org/{org_handle}/billing/invoice | List org invoices
[**ListOrgBillingPaymentMethod**](OrgBilling.md#ListOrgBillingPaymentMethod) | **Get** /org/{org_handle}/billing/payment_method | List org billing payment methods.
[**ListOrgSubscription**](OrgBilling.md#ListOrgSubscription) | **Get** /org/{org_handle}/billing/subscription | List org subscriptions
[**UpdateOrgBillingPaymentMethod**](OrgBilling.md#UpdateOrgBillingPaymentMethod) | **Patch** /org/{org_handle}/billing/payment_method/{payment_method_id} | Update org billing payment method.
[**UpdateOrgBillingSubscription**](OrgBilling.md#UpdateOrgBillingSubscription) | **Patch** /org/{org_handle}/billing/subscription/{subscription_id} | Update org subscription



## CreateBillingIntent

> BillingSetupIntent CreateBillingIntent(ctx, orgHandle).Execute()

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
    orgHandle := "orgHandle_example" // string | The handle of an organization where the billing intent will be created

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgBilling.CreateBillingIntent(context.Background(), orgHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgBilling.CreateBillingIntent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBillingIntent`: BillingSetupIntent
    fmt.Fprintf(os.Stdout, "Response from `OrgBilling.CreateBillingIntent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization where the billing intent will be created | 

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


## CreateOrgSubscription

> BillingSubscription CreateOrgSubscription(ctx, orgHandle).Execute()

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
    orgHandle := "orgHandle_example" // string | The handle of the organization where the subscription will be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgBilling.CreateOrgSubscription(context.Background(), orgHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgBilling.CreateOrgSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrgSubscription`: BillingSubscription
    fmt.Fprintf(os.Stdout, "Response from `OrgBilling.CreateOrgSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization where the subscription will be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgSubscriptionRequest struct via the builder pattern


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


## DeleteOrgBillingPaymentMethod

> PaymentMethod DeleteOrgBillingPaymentMethod(ctx, orgHandle, paymentMethodId).Execute()

Delete org billing payment method.



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
    orgHandle := "orgHandle_example" // string | The handle of an organization
    paymentMethodId := "paymentMethodId_example" // string | The org payment method id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgBilling.DeleteOrgBillingPaymentMethod(context.Background(), orgHandle, paymentMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgBilling.DeleteOrgBillingPaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrgBillingPaymentMethod`: PaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `OrgBilling.DeleteOrgBillingPaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization | 
**paymentMethodId** | **string** | The org payment method id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgBillingPaymentMethodRequest struct via the builder pattern


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


## DeleteOrgBillingSubscription

> BillingSubscription DeleteOrgBillingSubscription(ctx, orgHandle, subscriptionId).Execute()

Delete org subscription



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
    orgHandle := "orgHandle_example" // string | The handle of an organization.
    subscriptionId := "subscriptionId_example" // string | The org subscription id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgBilling.DeleteOrgBillingSubscription(context.Background(), orgHandle, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgBilling.DeleteOrgBillingSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteOrgBillingSubscription`: BillingSubscription
    fmt.Fprintf(os.Stdout, "Response from `OrgBilling.DeleteOrgBillingSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization. | 
**subscriptionId** | **string** | The org subscription id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgBillingSubscriptionRequest struct via the builder pattern


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


## GetOrgBillingInvoice

> BillingInvoice GetOrgBillingInvoice(ctx, orgHandle, invoiceId).Execute()

Get an organization invoice



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
    orgHandle := "orgHandle_example" // string | The handle of an organization
    invoiceId := "invoiceId_example" // string | The org invoice id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgBilling.GetOrgBillingInvoice(context.Background(), orgHandle, invoiceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgBilling.GetOrgBillingInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgBillingInvoice`: BillingInvoice
    fmt.Fprintf(os.Stdout, "Response from `OrgBilling.GetOrgBillingInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization | 
**invoiceId** | **string** | The org invoice id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgBillingInvoiceRequest struct via the builder pattern


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


## GetOrgBillingPaymentMethod

> PaymentMethod GetOrgBillingPaymentMethod(ctx, orgHandle, paymentMethodId).Execute()

Get org billing payment method.



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
    orgHandle := "orgHandle_example" // string | The handle of an organization
    paymentMethodId := "paymentMethodId_example" // string | The org payment method id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgBilling.GetOrgBillingPaymentMethod(context.Background(), orgHandle, paymentMethodId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgBilling.GetOrgBillingPaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgBillingPaymentMethod`: PaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `OrgBilling.GetOrgBillingPaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization | 
**paymentMethodId** | **string** | The org payment method id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgBillingPaymentMethodRequest struct via the builder pattern


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


## GetOrgBillingSubscription

> BillingSubscription GetOrgBillingSubscription(ctx, orgHandle, subscriptionId).Execute()

Get org subscription



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
    orgHandle := "orgHandle_example" // string | The handle of an organization.
    subscriptionId := "subscriptionId_example" // string | The org subscription id.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgBilling.GetOrgBillingSubscription(context.Background(), orgHandle, subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgBilling.GetOrgBillingSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgBillingSubscription`: BillingSubscription
    fmt.Fprintf(os.Stdout, "Response from `OrgBilling.GetOrgBillingSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization. | 
**subscriptionId** | **string** | The org subscription id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgBillingSubscriptionRequest struct via the builder pattern


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


## GetOrgBillingUpcomingInvoice

> BillingInvoice GetOrgBillingUpcomingInvoice(ctx, orgHandle).Execute()

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
    orgHandle := "orgHandle_example" // string | The handle of an organization

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgBilling.GetOrgBillingUpcomingInvoice(context.Background(), orgHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgBilling.GetOrgBillingUpcomingInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrgBillingUpcomingInvoice`: BillingInvoice
    fmt.Fprintf(os.Stdout, "Response from `OrgBilling.GetOrgBillingUpcomingInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgBillingUpcomingInvoiceRequest struct via the builder pattern


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


## ListOrgBillingInvoices

> ListBillingInvoicesResponse ListOrgBillingInvoices(ctx, orgHandle).Limit(limit).NextToken(nextToken).Execute()

List org invoices



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
    orgHandle := "orgHandle_example" // string | The handle of an organization
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgBilling.ListOrgBillingInvoices(context.Background(), orgHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgBilling.ListOrgBillingInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgBillingInvoices`: ListBillingInvoicesResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgBilling.ListOrgBillingInvoices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgBillingInvoicesRequest struct via the builder pattern


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


## ListOrgBillingPaymentMethod

> ListBillingPaymentMethodsResponse ListOrgBillingPaymentMethod(ctx, orgHandle).Execute()

List org billing payment methods.



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
    orgHandle := "orgHandle_example" // string | The handle of an organization

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgBilling.ListOrgBillingPaymentMethod(context.Background(), orgHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgBilling.ListOrgBillingPaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgBillingPaymentMethod`: ListBillingPaymentMethodsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgBilling.ListOrgBillingPaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgBillingPaymentMethodRequest struct via the builder pattern


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


## ListOrgSubscription

> ListSubscriptionsResponse ListOrgSubscription(ctx, orgHandle).Execute()

List org subscriptions



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
    orgHandle := "orgHandle_example" // string | The handle of an organization

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgBilling.ListOrgSubscription(context.Background(), orgHandle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgBilling.ListOrgSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrgSubscription`: ListSubscriptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrgBilling.ListOrgSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgSubscriptionRequest struct via the builder pattern


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


## UpdateOrgBillingPaymentMethod

> PaymentMethod UpdateOrgBillingPaymentMethod(ctx, orgHandle, paymentMethodId).Request(request).Execute()

Update org billing payment method.



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
    orgHandle := "orgHandle_example" // string | The handle of an organization
    paymentMethodId := "paymentMethodId_example" // string | The org payment method id.
    request := *openapiclient.NewUpdateBillingPaymentMethodRequest() // UpdateBillingPaymentMethodRequest | The request body for the org billing payment method update.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgBilling.UpdateOrgBillingPaymentMethod(context.Background(), orgHandle, paymentMethodId).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgBilling.UpdateOrgBillingPaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrgBillingPaymentMethod`: PaymentMethod
    fmt.Fprintf(os.Stdout, "Response from `OrgBilling.UpdateOrgBillingPaymentMethod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization | 
**paymentMethodId** | **string** | The org payment method id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgBillingPaymentMethodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**UpdateBillingPaymentMethodRequest**](UpdateBillingPaymentMethodRequest.md) | The request body for the org billing payment method update. | 

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


## UpdateOrgBillingSubscription

> BillingSubscription UpdateOrgBillingSubscription(ctx, orgHandle, subscriptionId).Request(request).Execute()

Update org subscription



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
    orgHandle := "orgHandle_example" // string | The handle of an organization.
    subscriptionId := "subscriptionId_example" // string | The org subscription id.
    request := *openapiclient.NewUpdateBillingSubscriptionRequest() // UpdateBillingSubscriptionRequest | Request body to update org subscription

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.OrgBilling.UpdateOrgBillingSubscription(context.Background(), orgHandle, subscriptionId).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrgBilling.UpdateOrgBillingSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrgBillingSubscription`: BillingSubscription
    fmt.Fprintf(os.Stdout, "Response from `OrgBilling.UpdateOrgBillingSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of an organization. | 
**subscriptionId** | **string** | The org subscription id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgBillingSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**UpdateBillingSubscriptionRequest**](UpdateBillingSubscriptionRequest.md) | Request body to update org subscription | 

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

