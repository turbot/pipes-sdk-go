/*
Turbot Pipes

Turbot Pipes is an intelligence, automation & security platform built specifically for DevOps.

API version: {{OPEN_API_VERSION}}
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipes

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// IntegrationsService Integrations service
type IntegrationsService service

type IntegrationsApiConfirmGithubIntegrationRequest struct {
	ctx            _context.Context
	ApiService     *IntegrationsService
	state          *string
	installationId *string
}

// The base64 encoded JSON state for the request state.
func (r IntegrationsApiConfirmGithubIntegrationRequest) State(state string) IntegrationsApiConfirmGithubIntegrationRequest {
	r.state = &state
	return r
}

// The GitHub installation id which needs to be confirmed.
func (r IntegrationsApiConfirmGithubIntegrationRequest) InstallationId(installationId string) IntegrationsApiConfirmGithubIntegrationRequest {
	r.installationId = &installationId
	return r
}

func (r IntegrationsApiConfirmGithubIntegrationRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.ConfirmGithubIntegrationExecute(r)
}

/*
ConfirmGithubIntegration Confirm GitHub integration installation

Confirm the installation of a GitHub integration.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return IntegrationsApiConfirmGithubIntegrationRequest
*/
func (a *IntegrationsService) ConfirmGithubIntegration(ctx _context.Context) IntegrationsApiConfirmGithubIntegrationRequest {
	return IntegrationsApiConfirmGithubIntegrationRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *IntegrationsService) ConfirmGithubIntegrationExecute(r IntegrationsApiConfirmGithubIntegrationRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationsService.ConfirmGithubIntegration")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/integration/github/confirm"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.state == nil {
		return nil, reportError("state is required and must be specified")
	}
	if r.installationId == nil {
		return nil, reportError("installationId is required and must be specified")
	}

	localVarQueryParams.Add("state", parameterToString(*r.state, ""))
	localVarQueryParams.Add("installation_id", parameterToString(*r.installationId, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 307 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
