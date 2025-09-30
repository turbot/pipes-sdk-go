# \AIKeys

All URIs are relative to *https://pipes.turbot.com/api/v0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrgCreate**](AIKeys.md#OrgCreate) | **Post** /org/{org_handle}/ai/key | Create org AI key
[**OrgDelete**](AIKeys.md#OrgDelete) | **Delete** /org/{org_handle}/ai/key/{provider} | Delete org AI key
[**OrgGet**](AIKeys.md#OrgGet) | **Get** /org/{org_handle}/ai/key/{provider} | Get org AI key
[**OrgList**](AIKeys.md#OrgList) | **Get** /org/{org_handle}/ai/key | List org AI keys
[**OrgTest**](AIKeys.md#OrgTest) | **Post** /org/{org_handle}/ai/key/{provider}/test | Test org AI key
[**OrgUpdate**](AIKeys.md#OrgUpdate) | **Patch** /org/{org_handle}/ai/key/{provider} | Update org AI key
[**TenantCreate**](AIKeys.md#TenantCreate) | **Post** /ai/key | Create tenant AI key
[**TenantDelete**](AIKeys.md#TenantDelete) | **Delete** /ai/key/{provider} | Delete tenant AI key
[**TenantGet**](AIKeys.md#TenantGet) | **Get** /ai/key/{provider} | Get tenant AI key
[**TenantList**](AIKeys.md#TenantList) | **Get** /ai/key | List tenant AI keys
[**TenantTest**](AIKeys.md#TenantTest) | **Post** /ai/key/{provider}/test | Test tenant AI key
[**TenantUpdate**](AIKeys.md#TenantUpdate) | **Patch** /ai/key/{provider} | Update tenant AI key
[**UserCreate**](AIKeys.md#UserCreate) | **Post** /user/{user_handle}/ai/key | Create user AI key
[**UserDelete**](AIKeys.md#UserDelete) | **Delete** /user/{user_handle}/ai/key/{provider} | Delete user AI key
[**UserGet**](AIKeys.md#UserGet) | **Get** /user/{user_handle}/ai/key/{provider} | Get user AI key
[**UserList**](AIKeys.md#UserList) | **Get** /user/{user_handle}/ai/key | List user AI keys
[**UserTest**](AIKeys.md#UserTest) | **Post** /user/{user_handle}/ai/key/{provider}/test | Test user AI key
[**UserUpdate**](AIKeys.md#UserUpdate) | **Patch** /user/{user_handle}/ai/key/{provider} | Update user AI key



## OrgCreate

> AIKey OrgCreate(ctx, orgHandle).Body(body).Execute()

Create org AI key



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
    orgHandle := "orgHandle_example" // string | The handle of the organization for which the AI key is being created.
    body := *openapiclient.NewCreateAIKeyRequest("ApiKey_example", map[string][]openapiclient.AIProvider{"key": "TODO"}) // CreateAIKeyRequest | The request body containing the AI provider and API key to save.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AIKeys.OrgCreate(context.Background(), orgHandle).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIKeys.OrgCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgCreate`: AIKey
    fmt.Fprintf(os.Stdout, "Response from `AIKeys.OrgCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization for which the AI key is being created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateAIKeyRequest**](CreateAIKeyRequest.md) | The request body containing the AI provider and API key to save. | 

### Return type

[**AIKey**](AIKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgDelete

> AIKey OrgDelete(ctx, orgHandle, provider).Execute()

Delete org AI key



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
    orgHandle := "orgHandle_example" // string | The handle of the organization that owns the AI key.
    provider := "provider_example" // string | The AI provider for which to delete the key.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AIKeys.OrgDelete(context.Background(), orgHandle, provider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIKeys.OrgDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgDelete`: AIKey
    fmt.Fprintf(os.Stdout, "Response from `AIKeys.OrgDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization that owns the AI key. | 
**provider** | **string** | The AI provider for which to delete the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AIKey**](AIKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgGet

> AIKey OrgGet(ctx, orgHandle, provider).Execute()

Get org AI key



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
    orgHandle := "orgHandle_example" // string | The handle of the organization that owns the AI key.
    provider := "provider_example" // string | The AI provider for which to retrieve the key.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AIKeys.OrgGet(context.Background(), orgHandle, provider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIKeys.OrgGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgGet`: AIKey
    fmt.Fprintf(os.Stdout, "Response from `AIKeys.OrgGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization that owns the AI key. | 
**provider** | **string** | The AI provider for which to retrieve the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AIKey**](AIKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgList

> ListAIKeysResponse OrgList(ctx, orgHandle).Limit(limit).NextToken(nextToken).Execute()

List org AI keys



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
    orgHandle := "orgHandle_example" // string | The handle of the organization for which to list AI keys.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned. Pass next_token to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AIKeys.OrgList(context.Background(), orgHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIKeys.OrgList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgList`: ListAIKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `AIKeys.OrgList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization for which to list AI keys. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned. Pass next_token to fetch the next page of data. | 

### Return type

[**ListAIKeysResponse**](ListAIKeysResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgTest

> TestAIKeyResponse OrgTest(ctx, orgHandle, provider).Body(body).Execute()

Test org AI key



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
    orgHandle := "orgHandle_example" // string | The handle of the organization that owns the AI key.
    provider := "provider_example" // string | The AI provider for which to test the key.
    body := *openapiclient.NewTestAIKeyRequest() // TestAIKeyRequest | Optional request body containing an API key to test instead of the saved one. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AIKeys.OrgTest(context.Background(), orgHandle, provider).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIKeys.OrgTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgTest`: TestAIKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `AIKeys.OrgTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization that owns the AI key. | 
**provider** | **string** | The AI provider for which to test the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**TestAIKeyRequest**](TestAIKeyRequest.md) | Optional request body containing an API key to test instead of the saved one. | 

### Return type

[**TestAIKeyResponse**](TestAIKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OrgUpdate

> AIKey OrgUpdate(ctx, orgHandle, provider).Body(body).Execute()

Update org AI key



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
    orgHandle := "orgHandle_example" // string | The handle of the organization that owns the AI key.
    provider := "provider_example" // string | The AI provider for which to update the key.
    body := *openapiclient.NewUpdateAIKeyRequest("ApiKey_example") // UpdateAIKeyRequest | The request body containing the updated API key value.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AIKeys.OrgUpdate(context.Background(), orgHandle, provider).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIKeys.OrgUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `OrgUpdate`: AIKey
    fmt.Fprintf(os.Stdout, "Response from `AIKeys.OrgUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orgHandle** | **string** | The handle of the organization that owns the AI key. | 
**provider** | **string** | The AI provider for which to update the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrgUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateAIKeyRequest**](UpdateAIKeyRequest.md) | The request body containing the updated API key value. | 

### Return type

[**AIKey**](AIKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantCreate

> AIKey TenantCreate(ctx).Body(body).Execute()

Create tenant AI key



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
    body := *openapiclient.NewCreateAIKeyRequest("ApiKey_example", map[string][]openapiclient.AIProvider{"key": "TODO"}) // CreateAIKeyRequest | The request body containing the AI provider and API key to save.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AIKeys.TenantCreate(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIKeys.TenantCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantCreate`: AIKey
    fmt.Fprintf(os.Stdout, "Response from `AIKeys.TenantCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateAIKeyRequest**](CreateAIKeyRequest.md) | The request body containing the AI provider and API key to save. | 

### Return type

[**AIKey**](AIKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantDelete

> AIKey TenantDelete(ctx, provider).Execute()

Delete tenant AI key



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
    provider := "provider_example" // string | The AI provider for which to delete the key.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AIKeys.TenantDelete(context.Background(), provider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIKeys.TenantDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantDelete`: AIKey
    fmt.Fprintf(os.Stdout, "Response from `AIKeys.TenantDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | The AI provider for which to delete the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AIKey**](AIKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantGet

> AIKey TenantGet(ctx, provider).Execute()

Get tenant AI key



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
    provider := "provider_example" // string | The AI provider for which to retrieve the key.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AIKeys.TenantGet(context.Background(), provider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIKeys.TenantGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantGet`: AIKey
    fmt.Fprintf(os.Stdout, "Response from `AIKeys.TenantGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | The AI provider for which to retrieve the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AIKey**](AIKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantList

> ListAIKeysResponse TenantList(ctx).Limit(limit).NextToken(nextToken).Execute()

List tenant AI keys



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
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned. Pass next_token to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AIKeys.TenantList(context.Background()).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIKeys.TenantList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantList`: ListAIKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `AIKeys.TenantList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTenantListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned. Pass next_token to fetch the next page of data. | 

### Return type

[**ListAIKeysResponse**](ListAIKeysResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantTest

> TestAIKeyResponse TenantTest(ctx, provider).Body(body).Execute()

Test tenant AI key



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
    provider := "provider_example" // string | The AI provider for which to test the key.
    body := *openapiclient.NewTestAIKeyRequest() // TestAIKeyRequest | Optional request body containing an API key to test instead of the saved one. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AIKeys.TenantTest(context.Background(), provider).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIKeys.TenantTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantTest`: TestAIKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `AIKeys.TenantTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | The AI provider for which to test the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**TestAIKeyRequest**](TestAIKeyRequest.md) | Optional request body containing an API key to test instead of the saved one. | 

### Return type

[**TestAIKeyResponse**](TestAIKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TenantUpdate

> AIKey TenantUpdate(ctx, provider).Body(body).Execute()

Update tenant AI key



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
    provider := "provider_example" // string | The AI provider for which to update the key.
    body := *openapiclient.NewUpdateAIKeyRequest("ApiKey_example") // UpdateAIKeyRequest | The request body containing the updated API key value.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AIKeys.TenantUpdate(context.Background(), provider).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIKeys.TenantUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TenantUpdate`: AIKey
    fmt.Fprintf(os.Stdout, "Response from `AIKeys.TenantUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | The AI provider for which to update the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTenantUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**UpdateAIKeyRequest**](UpdateAIKeyRequest.md) | The request body containing the updated API key value. | 

### Return type

[**AIKey**](AIKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserCreate

> AIKey UserCreate(ctx, userHandle).Body(body).Execute()

Create user AI key



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
    userHandle := "userHandle_example" // string | The handle of the user for which the AI key is being created.
    body := *openapiclient.NewCreateAIKeyRequest("ApiKey_example", map[string][]openapiclient.AIProvider{"key": "TODO"}) // CreateAIKeyRequest | The request body containing the AI provider and API key to save.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AIKeys.UserCreate(context.Background(), userHandle).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIKeys.UserCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserCreate`: AIKey
    fmt.Fprintf(os.Stdout, "Response from `AIKeys.UserCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which the AI key is being created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateAIKeyRequest**](CreateAIKeyRequest.md) | The request body containing the AI provider and API key to save. | 

### Return type

[**AIKey**](AIKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserDelete

> AIKey UserDelete(ctx, userHandle, provider).Execute()

Delete user AI key



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
    userHandle := "userHandle_example" // string | The handle of the user that owns the AI key.
    provider := "provider_example" // string | The AI provider for which to delete the key.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AIKeys.UserDelete(context.Background(), userHandle, provider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIKeys.UserDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserDelete`: AIKey
    fmt.Fprintf(os.Stdout, "Response from `AIKeys.UserDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that owns the AI key. | 
**provider** | **string** | The AI provider for which to delete the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AIKey**](AIKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserGet

> AIKey UserGet(ctx, userHandle, provider).Execute()

Get user AI key



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
    userHandle := "userHandle_example" // string | The handle of the user that owns the AI key.
    provider := "provider_example" // string | The AI provider for which to retrieve the key.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AIKeys.UserGet(context.Background(), userHandle, provider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIKeys.UserGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserGet`: AIKey
    fmt.Fprintf(os.Stdout, "Response from `AIKeys.UserGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that owns the AI key. | 
**provider** | **string** | The AI provider for which to retrieve the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AIKey**](AIKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserList

> ListAIKeysResponse UserList(ctx, userHandle).Limit(limit).NextToken(nextToken).Execute()

List user AI keys



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
    userHandle := "userHandle_example" // string | The handle of the user for which to list AI keys.
    limit := int32(56) // int32 | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. (optional) (default to 25)
    nextToken := "nextToken_example" // string | When list results are truncated, next_token will be returned. Pass next_token to fetch the next page of data. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AIKeys.UserList(context.Background(), userHandle).Limit(limit).NextToken(nextToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIKeys.UserList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserList`: ListAIKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `AIKeys.UserList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user for which to list AI keys. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25. | [default to 25]
 **nextToken** | **string** | When list results are truncated, next_token will be returned. Pass next_token to fetch the next page of data. | 

### Return type

[**ListAIKeysResponse**](ListAIKeysResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserTest

> TestAIKeyResponse UserTest(ctx, userHandle, provider).Body(body).Execute()

Test user AI key



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
    userHandle := "userHandle_example" // string | The handle of the user that owns the AI key.
    provider := "provider_example" // string | The AI provider for which to test the key.
    body := *openapiclient.NewTestAIKeyRequest() // TestAIKeyRequest | Optional request body containing an API key to test instead of the saved one. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AIKeys.UserTest(context.Background(), userHandle, provider).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIKeys.UserTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserTest`: TestAIKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `AIKeys.UserTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that owns the AI key. | 
**provider** | **string** | The AI provider for which to test the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**TestAIKeyRequest**](TestAIKeyRequest.md) | Optional request body containing an API key to test instead of the saved one. | 

### Return type

[**TestAIKeyResponse**](TestAIKeyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UserUpdate

> AIKey UserUpdate(ctx, userHandle, provider).Body(body).Execute()

Update user AI key



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
    userHandle := "userHandle_example" // string | The handle of the user that owns the AI key.
    provider := "provider_example" // string | The AI provider for which to update the key.
    body := *openapiclient.NewUpdateAIKeyRequest("ApiKey_example") // UpdateAIKeyRequest | The request body containing the updated API key value.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AIKeys.UserUpdate(context.Background(), userHandle, provider).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIKeys.UserUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserUpdate`: AIKey
    fmt.Fprintf(os.Stdout, "Response from `AIKeys.UserUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userHandle** | **string** | The handle of the user that owns the AI key. | 
**provider** | **string** | The AI provider for which to update the key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**UpdateAIKeyRequest**](UpdateAIKeyRequest.md) | The request body containing the updated API key value. | 

### Return type

[**AIKey**](AIKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

