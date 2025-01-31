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
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// OrgWorkspaceFlowpipePipelinesService OrgWorkspaceFlowpipePipelines service
type OrgWorkspaceFlowpipePipelinesService service

type OrgWorkspaceFlowpipePipelinesApiCommandRequest struct {
	ctx             _context.Context
	ApiService      *OrgWorkspaceFlowpipePipelinesService
	orgHandle       string
	workspaceHandle string
	pipelineName    string
	request         *PipelineCommandRequest
}

// The request body of the pipeline command to run.
func (r OrgWorkspaceFlowpipePipelinesApiCommandRequest) Request(request PipelineCommandRequest) OrgWorkspaceFlowpipePipelinesApiCommandRequest {
	r.request = &request
	return r
}

func (r OrgWorkspaceFlowpipePipelinesApiCommandRequest) Execute() (PipelineCommandResponse, *_nethttp.Response, error) {
	return r.ApiService.CommandExecute(r)
}

/*
Command Run organization workspace Flowpipe pipeline command

Run a command on a pipeline in a workspace belonging to an organization.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgHandle The handle of the organization which contains the workspace.
	@param workspaceHandle The handle of the workspace where the Flowpipe pipeline exists.
	@param pipelineName Identifier of the Flowpipe pipeline on which the command will be run.
	@return OrgWorkspaceFlowpipePipelinesApiCommandRequest
*/
func (a *OrgWorkspaceFlowpipePipelinesService) Command(ctx _context.Context, orgHandle string, workspaceHandle string, pipelineName string) OrgWorkspaceFlowpipePipelinesApiCommandRequest {
	return OrgWorkspaceFlowpipePipelinesApiCommandRequest{
		ApiService:      a,
		ctx:             ctx,
		orgHandle:       orgHandle,
		workspaceHandle: workspaceHandle,
		pipelineName:    pipelineName,
	}
}

// Execute executes the request
//
//	@return PipelineCommandResponse
func (a *OrgWorkspaceFlowpipePipelinesService) CommandExecute(r OrgWorkspaceFlowpipePipelinesApiCommandRequest) (PipelineCommandResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue PipelineCommandResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrgWorkspaceFlowpipePipelinesService.Command")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/org/{org_handle}/workspace/{workspace_handle}/flowpipe/pipeline/{pipeline_name}/command"
	localVarPath = strings.Replace(localVarPath, "{"+"org_handle"+"}", _neturl.PathEscape(parameterToString(r.orgHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspace_handle"+"}", _neturl.PathEscape(parameterToString(r.workspaceHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipeline_name"+"}", _neturl.PathEscape(parameterToString(r.pipelineName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.request == nil {
		return localVarReturnValue, nil, reportError("request is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.request
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 402 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OrgWorkspaceFlowpipePipelinesApiGetRequest struct {
	ctx             _context.Context
	ApiService      *OrgWorkspaceFlowpipePipelinesService
	orgHandle       string
	workspaceHandle string
	pipelineId      string
}

func (r OrgWorkspaceFlowpipePipelinesApiGetRequest) Execute() (WorkspaceModPipeline, *_nethttp.Response, error) {
	return r.ApiService.GetExecute(r)
}

/*
Get Get org workspace pipeline

Get the details of a pipeline for a workspace of an organization.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgHandle The handle of the organization which contains the workspace.
	@param workspaceHandle The handle of the workspace where the pipeline exists.
	@param pipelineId The id of the pipeline whose detail needs to be fetched.
	@return OrgWorkspaceFlowpipePipelinesApiGetRequest
*/
func (a *OrgWorkspaceFlowpipePipelinesService) Get(ctx _context.Context, orgHandle string, workspaceHandle string, pipelineId string) OrgWorkspaceFlowpipePipelinesApiGetRequest {
	return OrgWorkspaceFlowpipePipelinesApiGetRequest{
		ApiService:      a,
		ctx:             ctx,
		orgHandle:       orgHandle,
		workspaceHandle: workspaceHandle,
		pipelineId:      pipelineId,
	}
}

// Execute executes the request
//
//	@return WorkspaceModPipeline
func (a *OrgWorkspaceFlowpipePipelinesService) GetExecute(r OrgWorkspaceFlowpipePipelinesApiGetRequest) (WorkspaceModPipeline, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue WorkspaceModPipeline
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrgWorkspaceFlowpipePipelinesService.Get")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/org/{org_handle}/workspace/{workspace_handle}/flowpipe/pipeline/{pipeline_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"org_handle"+"}", _neturl.PathEscape(parameterToString(r.orgHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspace_handle"+"}", _neturl.PathEscape(parameterToString(r.workspaceHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipeline_id"+"}", _neturl.PathEscape(parameterToString(r.pipelineId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OrgWorkspaceFlowpipePipelinesApiListRequest struct {
	ctx             _context.Context
	ApiService      *OrgWorkspaceFlowpipePipelinesService
	orgHandle       string
	workspaceHandle string
	where           *string
	limit           *int32
	nextToken       *string
}

// The SQL where filter you wish to apply to this request. The filter will be parsed and sanitised and checked against the supported columns for this API.
func (r OrgWorkspaceFlowpipePipelinesApiListRequest) Where(where string) OrgWorkspaceFlowpipePipelinesApiListRequest {
	r.where = &where
	return r
}

// The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25.
func (r OrgWorkspaceFlowpipePipelinesApiListRequest) Limit(limit int32) OrgWorkspaceFlowpipePipelinesApiListRequest {
	r.limit = &limit
	return r
}

// When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data.
func (r OrgWorkspaceFlowpipePipelinesApiListRequest) NextToken(nextToken string) OrgWorkspaceFlowpipePipelinesApiListRequest {
	r.nextToken = &nextToken
	return r
}

func (r OrgWorkspaceFlowpipePipelinesApiListRequest) Execute() (ListFlowpipePipelinesResponse, *_nethttp.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List List organization workspace pipelines

Lists the pipelines for an organization workspace.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgHandle The handle of the organization which contains the workspace.
	@param workspaceHandle The handle of the workspace for which we want to list the pipelines.
	@return OrgWorkspaceFlowpipePipelinesApiListRequest
*/
func (a *OrgWorkspaceFlowpipePipelinesService) List(ctx _context.Context, orgHandle string, workspaceHandle string) OrgWorkspaceFlowpipePipelinesApiListRequest {
	return OrgWorkspaceFlowpipePipelinesApiListRequest{
		ApiService:      a,
		ctx:             ctx,
		orgHandle:       orgHandle,
		workspaceHandle: workspaceHandle,
	}
}

// Execute executes the request
//
//	@return ListFlowpipePipelinesResponse
func (a *OrgWorkspaceFlowpipePipelinesService) ListExecute(r OrgWorkspaceFlowpipePipelinesApiListRequest) (ListFlowpipePipelinesResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue ListFlowpipePipelinesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrgWorkspaceFlowpipePipelinesService.List")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/org/{org_handle}/workspace/{workspace_handle}/flowpipe/pipeline"
	localVarPath = strings.Replace(localVarPath, "{"+"org_handle"+"}", _neturl.PathEscape(parameterToString(r.orgHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspace_handle"+"}", _neturl.PathEscape(parameterToString(r.workspaceHandle, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.nextToken != nil {
		localVarQueryParams.Add("next_token", parameterToString(*r.nextToken, ""))
	}
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
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OrgWorkspaceFlowpipePipelinesApiListTriggersRequest struct {
	ctx             _context.Context
	ApiService      *OrgWorkspaceFlowpipePipelinesService
	orgHandle       string
	workspaceHandle string
	pipelineId      string
	where           *string
	limit           *int32
	nextToken       *string
}

// The SQL where filter you wish to apply to this request. The filter will be parsed and sanitised and checked against the supported columns for this API.
func (r OrgWorkspaceFlowpipePipelinesApiListTriggersRequest) Where(where string) OrgWorkspaceFlowpipePipelinesApiListTriggersRequest {
	r.where = &where
	return r
}

// The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25.
func (r OrgWorkspaceFlowpipePipelinesApiListTriggersRequest) Limit(limit int32) OrgWorkspaceFlowpipePipelinesApiListTriggersRequest {
	r.limit = &limit
	return r
}

// When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data.
func (r OrgWorkspaceFlowpipePipelinesApiListTriggersRequest) NextToken(nextToken string) OrgWorkspaceFlowpipePipelinesApiListTriggersRequest {
	r.nextToken = &nextToken
	return r
}

func (r OrgWorkspaceFlowpipePipelinesApiListTriggersRequest) Execute() (ListFlowpipeTriggersResponse, *_nethttp.Response, error) {
	return r.ApiService.ListTriggersExecute(r)
}

/*
ListTriggers Get org workspace pipeline triggers

Get the list of triggers that a pipeline is associated with.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgHandle The handle of the org which contains the workspace.
	@param workspaceHandle The handle of the workspace where the pipeline exists.
	@param pipelineId The id of the pipeline whose triggers need to be fetched.
	@return OrgWorkspaceFlowpipePipelinesApiListTriggersRequest
*/
func (a *OrgWorkspaceFlowpipePipelinesService) ListTriggers(ctx _context.Context, orgHandle string, workspaceHandle string, pipelineId string) OrgWorkspaceFlowpipePipelinesApiListTriggersRequest {
	return OrgWorkspaceFlowpipePipelinesApiListTriggersRequest{
		ApiService:      a,
		ctx:             ctx,
		orgHandle:       orgHandle,
		workspaceHandle: workspaceHandle,
		pipelineId:      pipelineId,
	}
}

// Execute executes the request
//
//	@return ListFlowpipeTriggersResponse
func (a *OrgWorkspaceFlowpipePipelinesService) ListTriggersExecute(r OrgWorkspaceFlowpipePipelinesApiListTriggersRequest) (ListFlowpipeTriggersResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue ListFlowpipeTriggersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrgWorkspaceFlowpipePipelinesService.ListTriggers")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/org/{org_handle}/workspace/{workspace_handle}/flowpipe/pipeline/{pipeline_id}/trigger"
	localVarPath = strings.Replace(localVarPath, "{"+"org_handle"+"}", _neturl.PathEscape(parameterToString(r.orgHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspace_handle"+"}", _neturl.PathEscape(parameterToString(r.workspaceHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pipeline_id"+"}", _neturl.PathEscape(parameterToString(r.pipelineId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.where != nil {
		localVarQueryParams.Add("where", parameterToString(*r.where, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.nextToken != nil {
		localVarQueryParams.Add("next_token", parameterToString(*r.nextToken, ""))
	}
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
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorModel
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
