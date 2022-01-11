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

// BankInfoApiService BankInfoApi service
type BankInfoApiService service

type ApiBankInfoListRequest struct {
	ctx _context.Context
	ApiService *BankInfoApiService
	xAccountToken *string
	accountType *string
	bankName *string
	createdAfter *time.Time
	createdBefore *time.Time
	cursor *string
	employee *string
	employeeId *string
	includeDeletedData *bool
	includeRemoteData *bool
	modifiedAfter *time.Time
	modifiedBefore *time.Time
	orderBy *string
	pageSize *int32
	remoteCreatedAt *time.Time
	remoteId *string
}

func (r ApiBankInfoListRequest) XAccountToken(xAccountToken string) ApiBankInfoListRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiBankInfoListRequest) AccountType(accountType string) ApiBankInfoListRequest {
	r.accountType = &accountType
	return r
}
func (r ApiBankInfoListRequest) BankName(bankName string) ApiBankInfoListRequest {
	r.bankName = &bankName
	return r
}
func (r ApiBankInfoListRequest) CreatedAfter(createdAfter time.Time) ApiBankInfoListRequest {
	r.createdAfter = &createdAfter
	return r
}
func (r ApiBankInfoListRequest) CreatedBefore(createdBefore time.Time) ApiBankInfoListRequest {
	r.createdBefore = &createdBefore
	return r
}
func (r ApiBankInfoListRequest) Cursor(cursor string) ApiBankInfoListRequest {
	r.cursor = &cursor
	return r
}
func (r ApiBankInfoListRequest) Employee(employee string) ApiBankInfoListRequest {
	r.employee = &employee
	return r
}
func (r ApiBankInfoListRequest) EmployeeId(employeeId string) ApiBankInfoListRequest {
	r.employeeId = &employeeId
	return r
}
func (r ApiBankInfoListRequest) IncludeDeletedData(includeDeletedData bool) ApiBankInfoListRequest {
	r.includeDeletedData = &includeDeletedData
	return r
}
func (r ApiBankInfoListRequest) IncludeRemoteData(includeRemoteData bool) ApiBankInfoListRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}
func (r ApiBankInfoListRequest) ModifiedAfter(modifiedAfter time.Time) ApiBankInfoListRequest {
	r.modifiedAfter = &modifiedAfter
	return r
}
func (r ApiBankInfoListRequest) ModifiedBefore(modifiedBefore time.Time) ApiBankInfoListRequest {
	r.modifiedBefore = &modifiedBefore
	return r
}
func (r ApiBankInfoListRequest) OrderBy(orderBy string) ApiBankInfoListRequest {
	r.orderBy = &orderBy
	return r
}
func (r ApiBankInfoListRequest) PageSize(pageSize int32) ApiBankInfoListRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiBankInfoListRequest) RemoteCreatedAt(remoteCreatedAt time.Time) ApiBankInfoListRequest {
	r.remoteCreatedAt = &remoteCreatedAt
	return r
}
func (r ApiBankInfoListRequest) RemoteId(remoteId string) ApiBankInfoListRequest {
	r.remoteId = &remoteId
	return r
}

func (r ApiBankInfoListRequest) Execute() (PaginatedBankInfoList, *_nethttp.Response, error) {
	return r.ApiService.BankInfoListExecute(r)
}

/*
 * BankInfoList Method for BankInfoList
 * Returns a list of `BankInfo` objects.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiBankInfoListRequest
 */
func (a *BankInfoApiService) BankInfoList(ctx _context.Context) ApiBankInfoListRequest {
	return ApiBankInfoListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginatedBankInfoList
 */
func (a *BankInfoApiService) BankInfoListExecute(r ApiBankInfoListRequest) (PaginatedBankInfoList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedBankInfoList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BankInfoApiService.BankInfoList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bank-info"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}

	if r.accountType != nil {
		localVarQueryParams.Add("account_type", parameterToString(*r.accountType, ""))
	}
	if r.bankName != nil {
		localVarQueryParams.Add("bank_name", parameterToString(*r.bankName, ""))
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
	if r.employee != nil {
		localVarQueryParams.Add("employee", parameterToString(*r.employee, ""))
	}
	if r.employeeId != nil {
		localVarQueryParams.Add("employee_id", parameterToString(*r.employeeId, ""))
	}
	if r.includeDeletedData != nil {
		localVarQueryParams.Add("include_deleted_data", parameterToString(*r.includeDeletedData, ""))
	}
	if r.includeRemoteData != nil {
		localVarQueryParams.Add("include_remote_data", parameterToString(*r.includeRemoteData, ""))
	}
	if r.modifiedAfter != nil {
		localVarQueryParams.Add("modified_after", parameterToString(*r.modifiedAfter, ""))
	}
	if r.modifiedBefore != nil {
		localVarQueryParams.Add("modified_before", parameterToString(*r.modifiedBefore, ""))
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("order_by", parameterToString(*r.orderBy, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.remoteCreatedAt != nil {
		localVarQueryParams.Add("remote_created_at", parameterToString(*r.remoteCreatedAt, ""))
	}
	if r.remoteId != nil {
		localVarQueryParams.Add("remote_id", parameterToString(*r.remoteId, ""))
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

type ApiBankInfoRetrieveRequest struct {
	ctx _context.Context
	ApiService *BankInfoApiService
	xAccountToken *string
	id string
	includeRemoteData *bool
}

func (r ApiBankInfoRetrieveRequest) XAccountToken(xAccountToken string) ApiBankInfoRetrieveRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiBankInfoRetrieveRequest) IncludeRemoteData(includeRemoteData bool) ApiBankInfoRetrieveRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}

func (r ApiBankInfoRetrieveRequest) Execute() (BankInfo, *_nethttp.Response, error) {
	return r.ApiService.BankInfoRetrieveExecute(r)
}

/*
 * BankInfoRetrieve Method for BankInfoRetrieve
 * Returns a `BankInfo` object with the given `id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiBankInfoRetrieveRequest
 */
func (a *BankInfoApiService) BankInfoRetrieve(ctx _context.Context, id string) ApiBankInfoRetrieveRequest {
	return ApiBankInfoRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return BankInfo
 */
func (a *BankInfoApiService) BankInfoRetrieveExecute(r ApiBankInfoRetrieveRequest) (BankInfo, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  BankInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BankInfoApiService.BankInfoRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bank-info/{id}"
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
