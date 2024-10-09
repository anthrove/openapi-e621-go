/*
E621

OpenAPI definition for E621's API. You can find the source [here](https://github.com/DonovanDMC/E621OpenAPI)<br> This document is intended to compliment E621's existing [API Documentation](https://e621.net/help/api).<br> Note if E621's api is under attack and/or cloudflare protections are enabled, the \"Try it out\" buttons here will not work.<br> If they are not working, you can check this [Unofficial Status Page](https://status.e621.ws).

API version: d69c34e
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// UserNameChangeRequestsAPIService UserNameChangeRequestsAPI service
type UserNameChangeRequestsAPIService service

type UserNameChangeRequestsAPICreateUserNameChangeRequestRequest struct {
	ctx                               context.Context
	ApiService                        *UserNameChangeRequestsAPIService
	userNameChangeRequestDesiredName  *string
	userNameChangeRequestChangeReason *string
}

func (r UserNameChangeRequestsAPICreateUserNameChangeRequestRequest) UserNameChangeRequestDesiredName(userNameChangeRequestDesiredName string) UserNameChangeRequestsAPICreateUserNameChangeRequestRequest {
	r.userNameChangeRequestDesiredName = &userNameChangeRequestDesiredName
	return r
}

func (r UserNameChangeRequestsAPICreateUserNameChangeRequestRequest) UserNameChangeRequestChangeReason(userNameChangeRequestChangeReason string) UserNameChangeRequestsAPICreateUserNameChangeRequestRequest {
	r.userNameChangeRequestChangeReason = &userNameChangeRequestChangeReason
	return r
}

func (r UserNameChangeRequestsAPICreateUserNameChangeRequestRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateUserNameChangeRequestExecute(r)
}

/*
CreateUserNameChangeRequest Create User Name Change Request

You must be Moderator+.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return UserNameChangeRequestsAPICreateUserNameChangeRequestRequest
*/
func (a *UserNameChangeRequestsAPIService) CreateUserNameChangeRequest(ctx context.Context) UserNameChangeRequestsAPICreateUserNameChangeRequestRequest {
	return UserNameChangeRequestsAPICreateUserNameChangeRequestRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *UserNameChangeRequestsAPIService) CreateUserNameChangeRequestExecute(r UserNameChangeRequestsAPICreateUserNameChangeRequestRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserNameChangeRequestsAPIService.CreateUserNameChangeRequest")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user_name_change_requests.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userNameChangeRequestDesiredName == nil {
		return nil, reportError("userNameChangeRequestDesiredName is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

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
	parameterAddToHeaderOrQuery(localVarFormParams, "user_name_change_request[desired_name]", r.userNameChangeRequestDesiredName, "")
	if r.userNameChangeRequestChangeReason != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user_name_change_request[change_reason]", r.userNameChangeRequestChangeReason, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type UserNameChangeRequestsAPIGetUserNameChangeRequestRequest struct {
	ctx        context.Context
	ApiService *UserNameChangeRequestsAPIService
	id         float32
}

func (r UserNameChangeRequestsAPIGetUserNameChangeRequestRequest) Execute() (*UserNameChangeRequest, *http.Response, error) {
	return r.ApiService.GetUserNameChangeRequestExecute(r)
}

/*
GetUserNameChangeRequest Get User Name Change Request

You must be the creator of the request or Moderator+.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the name change request.
	@return UserNameChangeRequestsAPIGetUserNameChangeRequestRequest
*/
func (a *UserNameChangeRequestsAPIService) GetUserNameChangeRequest(ctx context.Context, id float32) UserNameChangeRequestsAPIGetUserNameChangeRequestRequest {
	return UserNameChangeRequestsAPIGetUserNameChangeRequestRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return UserNameChangeRequest
func (a *UserNameChangeRequestsAPIService) GetUserNameChangeRequestExecute(r UserNameChangeRequestsAPIGetUserNameChangeRequestRequest) (*UserNameChangeRequest, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UserNameChangeRequest
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserNameChangeRequestsAPIService.GetUserNameChangeRequest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user_name_change_requests/{id}.json"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UserNameChangeRequestsAPISearchUserNameChangeRequestsRequest struct {
	ctx                context.Context
	ApiService         *UserNameChangeRequestsAPIService
	limit              *int32
	page               *int32
	searchId           *int32
	searchOrder        *string
	searchCurrentId    *float32
	searchCurrentName  *string
	searchOriginalName *string
	searchDesiredName  *string
}

// The maximum number of results to return. Between 0 and 320.
func (r UserNameChangeRequestsAPISearchUserNameChangeRequestsRequest) Limit(limit int32) UserNameChangeRequestsAPISearchUserNameChangeRequestsRequest {
	r.limit = &limit
	return r
}

// The page number of results to get. Between 1 and 750.
func (r UserNameChangeRequestsAPISearchUserNameChangeRequestsRequest) Page(page int32) UserNameChangeRequestsAPISearchUserNameChangeRequestsRequest {
	r.page = &page
	return r
}

// Search for a specific id.
func (r UserNameChangeRequestsAPISearchUserNameChangeRequestsRequest) SearchId(searchId int32) UserNameChangeRequestsAPISearchUserNameChangeRequestsRequest {
	r.searchId = &searchId
	return r
}

func (r UserNameChangeRequestsAPISearchUserNameChangeRequestsRequest) SearchOrder(searchOrder string) UserNameChangeRequestsAPISearchUserNameChangeRequestsRequest {
	r.searchOrder = &searchOrder
	return r
}

func (r UserNameChangeRequestsAPISearchUserNameChangeRequestsRequest) SearchCurrentId(searchCurrentId float32) UserNameChangeRequestsAPISearchUserNameChangeRequestsRequest {
	r.searchCurrentId = &searchCurrentId
	return r
}

func (r UserNameChangeRequestsAPISearchUserNameChangeRequestsRequest) SearchCurrentName(searchCurrentName string) UserNameChangeRequestsAPISearchUserNameChangeRequestsRequest {
	r.searchCurrentName = &searchCurrentName
	return r
}

func (r UserNameChangeRequestsAPISearchUserNameChangeRequestsRequest) SearchOriginalName(searchOriginalName string) UserNameChangeRequestsAPISearchUserNameChangeRequestsRequest {
	r.searchOriginalName = &searchOriginalName
	return r
}

func (r UserNameChangeRequestsAPISearchUserNameChangeRequestsRequest) SearchDesiredName(searchDesiredName string) UserNameChangeRequestsAPISearchUserNameChangeRequestsRequest {
	r.searchDesiredName = &searchDesiredName
	return r
}

func (r UserNameChangeRequestsAPISearchUserNameChangeRequestsRequest) Execute() (*SearchUserNameChangeRequests200Response, *http.Response, error) {
	return r.ApiService.SearchUserNameChangeRequestsExecute(r)
}

/*
SearchUserNameChangeRequests Search User Name Change Requests

You must be Moderator+. When no results are found, an object with an `user_name_change_requests` key is returned.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return UserNameChangeRequestsAPISearchUserNameChangeRequestsRequest
*/
func (a *UserNameChangeRequestsAPIService) SearchUserNameChangeRequests(ctx context.Context) UserNameChangeRequestsAPISearchUserNameChangeRequestsRequest {
	return UserNameChangeRequestsAPISearchUserNameChangeRequestsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SearchUserNameChangeRequests200Response
func (a *UserNameChangeRequestsAPIService) SearchUserNameChangeRequestsExecute(r UserNameChangeRequestsAPISearchUserNameChangeRequestsRequest) (*SearchUserNameChangeRequests200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SearchUserNameChangeRequests200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserNameChangeRequestsAPIService.SearchUserNameChangeRequests")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user_name_change_requests.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.searchId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[id]", r.searchId, "")
	}
	if r.searchOrder != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[order]", r.searchOrder, "")
	}
	if r.searchCurrentId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[current_id]", r.searchCurrentId, "")
	}
	if r.searchCurrentName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[current_name]", r.searchCurrentName, "")
	}
	if r.searchOriginalName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[original_name]", r.searchOriginalName, "")
	}
	if r.searchDesiredName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[desired_name]", r.searchDesiredName, "")
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
