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

// UserWorkspaceDatatankTablesService UserWorkspaceDatatankTables service
type UserWorkspaceDatatankTablesService service

type UserWorkspaceDatatankTablesApiCreateRequest struct {
	ctx             _context.Context
	ApiService      *UserWorkspaceDatatankTablesService
	userHandle      string
	workspaceHandle string
	datatankHandle  string
	request         *CreateDatatankTableRequest
}

// The request body for the workspace Datatank table to be created.
func (r UserWorkspaceDatatankTablesApiCreateRequest) Request(request CreateDatatankTableRequest) UserWorkspaceDatatankTablesApiCreateRequest {
	r.request = &request
	return r
}

func (r UserWorkspaceDatatankTablesApiCreateRequest) Execute() (DatatankTable, *_nethttp.Response, error) {
	return r.ApiService.CreateExecute(r)
}

/*
Create Create user workspace Datatank table

Creates a new user workspace tank table.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userHandle The handle of the user where we want to create the workspace Datatank table.
	@param workspaceHandle The handle of the workspace.
	@param datatankHandle The name of the Datatank.
	@return UserWorkspaceDatatankTablesApiCreateRequest
*/
func (a *UserWorkspaceDatatankTablesService) Create(ctx _context.Context, userHandle string, workspaceHandle string, datatankHandle string) UserWorkspaceDatatankTablesApiCreateRequest {
	return UserWorkspaceDatatankTablesApiCreateRequest{
		ApiService:      a,
		ctx:             ctx,
		userHandle:      userHandle,
		workspaceHandle: workspaceHandle,
		datatankHandle:  datatankHandle,
	}
}

// Execute executes the request
//
//	@return DatatankTable
func (a *UserWorkspaceDatatankTablesService) CreateExecute(r UserWorkspaceDatatankTablesApiCreateRequest) (DatatankTable, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue DatatankTable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserWorkspaceDatatankTablesService.Create")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table"
	localVarPath = strings.Replace(localVarPath, "{"+"user_handle"+"}", _neturl.PathEscape(parameterToString(r.userHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspace_handle"+"}", _neturl.PathEscape(parameterToString(r.workspaceHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"datatank_handle"+"}", _neturl.PathEscape(parameterToString(r.datatankHandle, "")), -1)

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

type UserWorkspaceDatatankTablesApiDeleteRequest struct {
	ctx               _context.Context
	ApiService        *UserWorkspaceDatatankTablesService
	userHandle        string
	workspaceHandle   string
	datatankHandle    string
	datatankTableName string
}

func (r UserWorkspaceDatatankTablesApiDeleteRequest) Execute() (DatatankTable, *_nethttp.Response, error) {
	return r.ApiService.DeleteExecute(r)
}

/*
Delete Delete user workspace Datatank table

Deletes the user workspace Datatank table.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userHandle The handle of the user.
	@param workspaceHandle The handle of the workspace.
	@param datatankHandle The name of the workspace Datatank.
	@param datatankTableName The name of the workspace Datatank table to be deleted.
	@return UserWorkspaceDatatankTablesApiDeleteRequest
*/
func (a *UserWorkspaceDatatankTablesService) Delete(ctx _context.Context, userHandle string, workspaceHandle string, datatankHandle string, datatankTableName string) UserWorkspaceDatatankTablesApiDeleteRequest {
	return UserWorkspaceDatatankTablesApiDeleteRequest{
		ApiService:        a,
		ctx:               ctx,
		userHandle:        userHandle,
		workspaceHandle:   workspaceHandle,
		datatankHandle:    datatankHandle,
		datatankTableName: datatankTableName,
	}
}

// Execute executes the request
//
//	@return DatatankTable
func (a *UserWorkspaceDatatankTablesService) DeleteExecute(r UserWorkspaceDatatankTablesApiDeleteRequest) (DatatankTable, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue DatatankTable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserWorkspaceDatatankTablesService.Delete")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_handle"+"}", _neturl.PathEscape(parameterToString(r.userHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspace_handle"+"}", _neturl.PathEscape(parameterToString(r.workspaceHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"datatank_handle"+"}", _neturl.PathEscape(parameterToString(r.datatankHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"datatank_table_name"+"}", _neturl.PathEscape(parameterToString(r.datatankTableName, "")), -1)

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

type UserWorkspaceDatatankTablesApiGetRequest struct {
	ctx               _context.Context
	ApiService        *UserWorkspaceDatatankTablesService
	userHandle        string
	workspaceHandle   string
	datatankHandle    string
	datatankTableName string
}

func (r UserWorkspaceDatatankTablesApiGetRequest) Execute() (DatatankTable, *_nethttp.Response, error) {
	return r.ApiService.GetExecute(r)
}

/*
Get Get user workspace Datatank table

Get the details for the workspace Datatank table

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userHandle The handle of the user.
	@param workspaceHandle The handle of the workspace.
	@param datatankHandle The name of the workspace Datatank.
	@param datatankTableName The name of the workspace Datatank table.
	@return UserWorkspaceDatatankTablesApiGetRequest
*/
func (a *UserWorkspaceDatatankTablesService) Get(ctx _context.Context, userHandle string, workspaceHandle string, datatankHandle string, datatankTableName string) UserWorkspaceDatatankTablesApiGetRequest {
	return UserWorkspaceDatatankTablesApiGetRequest{
		ApiService:        a,
		ctx:               ctx,
		userHandle:        userHandle,
		workspaceHandle:   workspaceHandle,
		datatankHandle:    datatankHandle,
		datatankTableName: datatankTableName,
	}
}

// Execute executes the request
//
//	@return DatatankTable
func (a *UserWorkspaceDatatankTablesService) GetExecute(r UserWorkspaceDatatankTablesApiGetRequest) (DatatankTable, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue DatatankTable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserWorkspaceDatatankTablesService.Get")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_handle"+"}", _neturl.PathEscape(parameterToString(r.userHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspace_handle"+"}", _neturl.PathEscape(parameterToString(r.workspaceHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"datatank_handle"+"}", _neturl.PathEscape(parameterToString(r.datatankHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"datatank_table_name"+"}", _neturl.PathEscape(parameterToString(r.datatankTableName, "")), -1)

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

type UserWorkspaceDatatankTablesApiListRequest struct {
	ctx             _context.Context
	ApiService      *UserWorkspaceDatatankTablesService
	userHandle      string
	workspaceHandle string
	datatankHandle  string
	limit           *int32
	nextToken       *string
}

// The max number of items to fetch per page of data, subject to a min and max of 1 and 100 respectively. If not specified will default to 25.
func (r UserWorkspaceDatatankTablesApiListRequest) Limit(limit int32) UserWorkspaceDatatankTablesApiListRequest {
	r.limit = &limit
	return r
}

// When list results are truncated, next_token will be returned, which is a cursor to fetch the next page of data. Pass next_token to the subsequent list request to fetch the next page of data.
func (r UserWorkspaceDatatankTablesApiListRequest) NextToken(nextToken string) UserWorkspaceDatatankTablesApiListRequest {
	r.nextToken = &nextToken
	return r
}

func (r UserWorkspaceDatatankTablesApiListRequest) Execute() (ListDatatankTableResponse, *_nethttp.Response, error) {
	return r.ApiService.ListExecute(r)
}

/*
List List user workspace Datatank tables

List the user workspace Datatank tables.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userHandle The handle of the user for which we want to list the workspace Datatank tables.
	@param workspaceHandle The handle of the workspace.
	@param datatankHandle The name of the Datatank.
	@return UserWorkspaceDatatankTablesApiListRequest
*/
func (a *UserWorkspaceDatatankTablesService) List(ctx _context.Context, userHandle string, workspaceHandle string, datatankHandle string) UserWorkspaceDatatankTablesApiListRequest {
	return UserWorkspaceDatatankTablesApiListRequest{
		ApiService:      a,
		ctx:             ctx,
		userHandle:      userHandle,
		workspaceHandle: workspaceHandle,
		datatankHandle:  datatankHandle,
	}
}

// Execute executes the request
//
//	@return ListDatatankTableResponse
func (a *UserWorkspaceDatatankTablesService) ListExecute(r UserWorkspaceDatatankTablesApiListRequest) (ListDatatankTableResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue ListDatatankTableResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserWorkspaceDatatankTablesService.List")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table"
	localVarPath = strings.Replace(localVarPath, "{"+"user_handle"+"}", _neturl.PathEscape(parameterToString(r.userHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspace_handle"+"}", _neturl.PathEscape(parameterToString(r.workspaceHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"datatank_handle"+"}", _neturl.PathEscape(parameterToString(r.datatankHandle, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type UserWorkspaceDatatankTablesApiUpdateRequest struct {
	ctx               _context.Context
	ApiService        *UserWorkspaceDatatankTablesService
	userHandle        string
	workspaceHandle   string
	datatankHandle    string
	datatankTableName string
	request           *UpdateDatatankTableRequest
}

// The request body to update workspace Datatank table.
func (r UserWorkspaceDatatankTablesApiUpdateRequest) Request(request UpdateDatatankTableRequest) UserWorkspaceDatatankTablesApiUpdateRequest {
	r.request = &request
	return r
}

func (r UserWorkspaceDatatankTablesApiUpdateRequest) Execute() (DatatankTable, *_nethttp.Response, error) {
	return r.ApiService.UpdateExecute(r)
}

/*
Update Update user workspace Datatank table

Update the user workspace Datatank table.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userHandle The handle of the user.
	@param workspaceHandle The handle of the workspace.
	@param datatankHandle The name of the workspace Datatank.
	@param datatankTableName The name of the workspace Datatank table to be updated.
	@return UserWorkspaceDatatankTablesApiUpdateRequest
*/
func (a *UserWorkspaceDatatankTablesService) Update(ctx _context.Context, userHandle string, workspaceHandle string, datatankHandle string, datatankTableName string) UserWorkspaceDatatankTablesApiUpdateRequest {
	return UserWorkspaceDatatankTablesApiUpdateRequest{
		ApiService:        a,
		ctx:               ctx,
		userHandle:        userHandle,
		workspaceHandle:   workspaceHandle,
		datatankHandle:    datatankHandle,
		datatankTableName: datatankTableName,
	}
}

// Execute executes the request
//
//	@return DatatankTable
func (a *UserWorkspaceDatatankTablesService) UpdateExecute(r UserWorkspaceDatatankTablesApiUpdateRequest) (DatatankTable, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod  = _nethttp.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue DatatankTable
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserWorkspaceDatatankTablesService.Update")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user/{user_handle}/workspace/{workspace_handle}/datatank/{datatank_handle}/table/{datatank_table_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"user_handle"+"}", _neturl.PathEscape(parameterToString(r.userHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"workspace_handle"+"}", _neturl.PathEscape(parameterToString(r.workspaceHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"datatank_handle"+"}", _neturl.PathEscape(parameterToString(r.datatankHandle, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"datatank_table_name"+"}", _neturl.PathEscape(parameterToString(r.datatankTableName, "")), -1)

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
