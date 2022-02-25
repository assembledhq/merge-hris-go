/*
 * Merge HRIS API
 *
 * The unified API for building rich integrations with multiple HR Information System platforms.
 *
 * API version: 1.0
 * Contact: hello@merge.dev
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merge_hris_client

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
	"time"
)

// Linger please
var (
	_ _context.Context
)

// EmployeesApiService EmployeesApi service
type EmployeesApiService service

type ApiEmployeesCreateRequest struct {
	ctx _context.Context
	ApiService *EmployeesApiService
	xAccountToken *string
	employeeEndpointRequest *EmployeeEndpointRequest
	isDebugMode *bool
	runAsync *bool
}

func (r ApiEmployeesCreateRequest) XAccountToken(xAccountToken string) ApiEmployeesCreateRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiEmployeesCreateRequest) EmployeeEndpointRequest(employeeEndpointRequest EmployeeEndpointRequest) ApiEmployeesCreateRequest {
	r.employeeEndpointRequest = &employeeEndpointRequest
	return r
}
func (r ApiEmployeesCreateRequest) IsDebugMode(isDebugMode bool) ApiEmployeesCreateRequest {
	r.isDebugMode = &isDebugMode
	return r
}
func (r ApiEmployeesCreateRequest) RunAsync(runAsync bool) ApiEmployeesCreateRequest {
	r.runAsync = &runAsync
	return r
}

func (r ApiEmployeesCreateRequest) Execute() (EmployeeResponse, *_nethttp.Response, error) {
	return r.ApiService.EmployeesCreateExecute(r)
}

/*
 * EmployeesCreate Method for EmployeesCreate
 * Creates an `Employee` object with the given values.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiEmployeesCreateRequest
 */
func (a *EmployeesApiService) EmployeesCreate(ctx _context.Context) ApiEmployeesCreateRequest {
	return ApiEmployeesCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return EmployeeResponse
 */
func (a *EmployeesApiService) EmployeesCreateExecute(r ApiEmployeesCreateRequest) (EmployeeResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  EmployeeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmployeesApiService.EmployeesCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/employees"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}
	if r.employeeEndpointRequest == nil {
		return localVarReturnValue, nil, reportError("employeeEndpointRequest is required and must be specified")
	}

	if r.isDebugMode != nil {
		localVarQueryParams.Add("is_debug_mode", parameterToString(*r.isDebugMode, ""))
	}
	if r.runAsync != nil {
		localVarQueryParams.Add("run_async", parameterToString(*r.runAsync, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

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
	localVarHeaderParams["X-Account-Token"] = parameterToString(*r.xAccountToken, "")
	// body params
	localVarPostBody = r.employeeEndpointRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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

type ApiEmployeesIgnoreCreateRequest struct {
	ctx _context.Context
	ApiService *EmployeesApiService
	modelId string
	ignoreCommonModelRequest *IgnoreCommonModelRequest
}

func (r ApiEmployeesIgnoreCreateRequest) IgnoreCommonModelRequest(ignoreCommonModelRequest IgnoreCommonModelRequest) ApiEmployeesIgnoreCreateRequest {
	r.ignoreCommonModelRequest = &ignoreCommonModelRequest
	return r
}

func (r ApiEmployeesIgnoreCreateRequest) Execute() (IgnoreCommonModel, *_nethttp.Response, error) {
	return r.ApiService.EmployeesIgnoreCreateExecute(r)
}

/*
 * EmployeesIgnoreCreate Method for EmployeesIgnoreCreate
 * Ignores a specific row based on the `model_id` in the url. These records will have their properties set to null, and will not be updated in future syncs. The "reason" and "message" fields in the request body will be stored for audit purposes.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param modelId
 * @return ApiEmployeesIgnoreCreateRequest
 */
func (a *EmployeesApiService) EmployeesIgnoreCreate(ctx _context.Context, modelId string) ApiEmployeesIgnoreCreateRequest {
	return ApiEmployeesIgnoreCreateRequest{
		ApiService: a,
		ctx: ctx,
		modelId: modelId,
	}
}

/*
 * Execute executes the request
 * @return IgnoreCommonModel
 */
func (a *EmployeesApiService) EmployeesIgnoreCreateExecute(r ApiEmployeesIgnoreCreateRequest) (IgnoreCommonModel, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  IgnoreCommonModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmployeesApiService.EmployeesIgnoreCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/employees/ignore/{model_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"model_id"+"}", _neturl.PathEscape(parameterToString(r.modelId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.ignoreCommonModelRequest == nil {
		return localVarReturnValue, nil, reportError("ignoreCommonModelRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

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
	localVarPostBody = r.ignoreCommonModelRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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

type ApiEmployeesListRequest struct {
	ctx _context.Context
	ApiService *EmployeesApiService
	xAccountToken *string
	companyId *string
	createdAfter *time.Time
	createdBefore *time.Time
	cursor *string
	includeDeletedData *bool
	includeRemoteData *bool
	includeSensitiveFields *bool
	managerId *string
	modifiedAfter *time.Time
	modifiedBefore *time.Time
	pageSize *int32
	payGroupId *string
	personalEmail *string
	remoteId *string
	teamId *string
	workEmail *string
	workLocationId *string
}

func (r ApiEmployeesListRequest) XAccountToken(xAccountToken string) ApiEmployeesListRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiEmployeesListRequest) CompanyId(companyId string) ApiEmployeesListRequest {
	r.companyId = &companyId
	return r
}
func (r ApiEmployeesListRequest) CreatedAfter(createdAfter time.Time) ApiEmployeesListRequest {
	r.createdAfter = &createdAfter
	return r
}
func (r ApiEmployeesListRequest) CreatedBefore(createdBefore time.Time) ApiEmployeesListRequest {
	r.createdBefore = &createdBefore
	return r
}
func (r ApiEmployeesListRequest) Cursor(cursor string) ApiEmployeesListRequest {
	r.cursor = &cursor
	return r
}
func (r ApiEmployeesListRequest) IncludeDeletedData(includeDeletedData bool) ApiEmployeesListRequest {
	r.includeDeletedData = &includeDeletedData
	return r
}
func (r ApiEmployeesListRequest) IncludeRemoteData(includeRemoteData bool) ApiEmployeesListRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}
func (r ApiEmployeesListRequest) IncludeSensitiveFields(includeSensitiveFields bool) ApiEmployeesListRequest {
	r.includeSensitiveFields = &includeSensitiveFields
	return r
}
func (r ApiEmployeesListRequest) ManagerId(managerId string) ApiEmployeesListRequest {
	r.managerId = &managerId
	return r
}
func (r ApiEmployeesListRequest) ModifiedAfter(modifiedAfter time.Time) ApiEmployeesListRequest {
	r.modifiedAfter = &modifiedAfter
	return r
}
func (r ApiEmployeesListRequest) ModifiedBefore(modifiedBefore time.Time) ApiEmployeesListRequest {
	r.modifiedBefore = &modifiedBefore
	return r
}
func (r ApiEmployeesListRequest) PageSize(pageSize int32) ApiEmployeesListRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiEmployeesListRequest) PayGroupId(payGroupId string) ApiEmployeesListRequest {
	r.payGroupId = &payGroupId
	return r
}
func (r ApiEmployeesListRequest) PersonalEmail(personalEmail string) ApiEmployeesListRequest {
	r.personalEmail = &personalEmail
	return r
}
func (r ApiEmployeesListRequest) RemoteId(remoteId string) ApiEmployeesListRequest {
	r.remoteId = &remoteId
	return r
}
func (r ApiEmployeesListRequest) TeamId(teamId string) ApiEmployeesListRequest {
	r.teamId = &teamId
	return r
}
func (r ApiEmployeesListRequest) WorkEmail(workEmail string) ApiEmployeesListRequest {
	r.workEmail = &workEmail
	return r
}
func (r ApiEmployeesListRequest) WorkLocationId(workLocationId string) ApiEmployeesListRequest {
	r.workLocationId = &workLocationId
	return r
}

func (r ApiEmployeesListRequest) Execute() (PaginatedEmployeeList, *_nethttp.Response, error) {
	return r.ApiService.EmployeesListExecute(r)
}

/*
 * EmployeesList Method for EmployeesList
 * Returns a list of `Employee` objects.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiEmployeesListRequest
 */
func (a *EmployeesApiService) EmployeesList(ctx _context.Context) ApiEmployeesListRequest {
	return ApiEmployeesListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginatedEmployeeList
 */
func (a *EmployeesApiService) EmployeesListExecute(r ApiEmployeesListRequest) (PaginatedEmployeeList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedEmployeeList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmployeesApiService.EmployeesList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/employees"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}

	if r.companyId != nil {
		localVarQueryParams.Add("company_id", parameterToString(*r.companyId, ""))
	}
	if r.createdAfter != nil {
		localVarQueryParams.Add("created_after", parameterToString(*r.createdAfter, ""))
	}
	if r.createdBefore != nil {
		localVarQueryParams.Add("created_before", parameterToString(*r.createdBefore, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.includeDeletedData != nil {
		localVarQueryParams.Add("include_deleted_data", parameterToString(*r.includeDeletedData, ""))
	}
	if r.includeRemoteData != nil {
		localVarQueryParams.Add("include_remote_data", parameterToString(*r.includeRemoteData, ""))
	}
	if r.includeSensitiveFields != nil {
		localVarQueryParams.Add("include_sensitive_fields", parameterToString(*r.includeSensitiveFields, ""))
	}
	if r.managerId != nil {
		localVarQueryParams.Add("manager_id", parameterToString(*r.managerId, ""))
	}
	if r.modifiedAfter != nil {
		localVarQueryParams.Add("modified_after", parameterToString(*r.modifiedAfter, ""))
	}
	if r.modifiedBefore != nil {
		localVarQueryParams.Add("modified_before", parameterToString(*r.modifiedBefore, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.payGroupId != nil {
		localVarQueryParams.Add("pay_group_id", parameterToString(*r.payGroupId, ""))
	}
	if r.personalEmail != nil {
		localVarQueryParams.Add("personal_email", parameterToString(*r.personalEmail, ""))
	}
	if r.remoteId != nil {
		localVarQueryParams.Add("remote_id", parameterToString(*r.remoteId, ""))
	}
	if r.teamId != nil {
		localVarQueryParams.Add("team_id", parameterToString(*r.teamId, ""))
	}
	if r.workEmail != nil {
		localVarQueryParams.Add("work_email", parameterToString(*r.workEmail, ""))
	}
	if r.workLocationId != nil {
		localVarQueryParams.Add("work_location_id", parameterToString(*r.workLocationId, ""))
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
	localVarHeaderParams["X-Account-Token"] = parameterToString(*r.xAccountToken, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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

type ApiEmployeesRetrieveRequest struct {
	ctx _context.Context
	ApiService *EmployeesApiService
	xAccountToken *string
	id string
	includeRemoteData *bool
	includeSensitiveFields *bool
}

func (r ApiEmployeesRetrieveRequest) XAccountToken(xAccountToken string) ApiEmployeesRetrieveRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiEmployeesRetrieveRequest) IncludeRemoteData(includeRemoteData bool) ApiEmployeesRetrieveRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}
func (r ApiEmployeesRetrieveRequest) IncludeSensitiveFields(includeSensitiveFields bool) ApiEmployeesRetrieveRequest {
	r.includeSensitiveFields = &includeSensitiveFields
	return r
}

func (r ApiEmployeesRetrieveRequest) Execute() (Employee, *_nethttp.Response, error) {
	return r.ApiService.EmployeesRetrieveExecute(r)
}

/*
 * EmployeesRetrieve Method for EmployeesRetrieve
 * Returns an `Employee` object with the given `id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiEmployeesRetrieveRequest
 */
func (a *EmployeesApiService) EmployeesRetrieve(ctx _context.Context, id string) ApiEmployeesRetrieveRequest {
	return ApiEmployeesRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return Employee
 */
func (a *EmployeesApiService) EmployeesRetrieveExecute(r ApiEmployeesRetrieveRequest) (Employee, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Employee
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EmployeesApiService.EmployeesRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/employees/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}

	if r.includeRemoteData != nil {
		localVarQueryParams.Add("include_remote_data", parameterToString(*r.includeRemoteData, ""))
	}
	if r.includeSensitiveFields != nil {
		localVarQueryParams.Add("include_sensitive_fields", parameterToString(*r.includeSensitiveFields, ""))
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
	localVarHeaderParams["X-Account-Token"] = parameterToString(*r.xAccountToken, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
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
