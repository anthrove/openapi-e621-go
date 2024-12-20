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

// UserFeedbacksAPIService UserFeedbacksAPI service
type UserFeedbacksAPIService service

type UserFeedbacksAPICreateUserFeedbackRequest struct {
	ctx                  context.Context
	ApiService           *UserFeedbacksAPIService
	userFeedbackBody     *string
	userFeedbackCategory *FeedbackCategories
	userFeedbackUserId   *int32
	userFeedbackUserName *string
}

func (r UserFeedbacksAPICreateUserFeedbackRequest) UserFeedbackBody(userFeedbackBody string) UserFeedbacksAPICreateUserFeedbackRequest {
	r.userFeedbackBody = &userFeedbackBody
	return r
}

func (r UserFeedbacksAPICreateUserFeedbackRequest) UserFeedbackCategory(userFeedbackCategory FeedbackCategories) UserFeedbacksAPICreateUserFeedbackRequest {
	r.userFeedbackCategory = &userFeedbackCategory
	return r
}

func (r UserFeedbacksAPICreateUserFeedbackRequest) UserFeedbackUserId(userFeedbackUserId int32) UserFeedbacksAPICreateUserFeedbackRequest {
	r.userFeedbackUserId = &userFeedbackUserId
	return r
}

func (r UserFeedbacksAPICreateUserFeedbackRequest) UserFeedbackUserName(userFeedbackUserName string) UserFeedbacksAPICreateUserFeedbackRequest {
	r.userFeedbackUserName = &userFeedbackUserName
	return r
}

func (r UserFeedbacksAPICreateUserFeedbackRequest) Execute() (*UserFeedback, *http.Response, error) {
	return r.ApiService.CreateUserFeedbackExecute(r)
}

/*
CreateUserFeedback Create User Feedback

You must be Moderator+.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return UserFeedbacksAPICreateUserFeedbackRequest
*/
func (a *UserFeedbacksAPIService) CreateUserFeedback(ctx context.Context) UserFeedbacksAPICreateUserFeedbackRequest {
	return UserFeedbacksAPICreateUserFeedbackRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return UserFeedback
func (a *UserFeedbacksAPIService) CreateUserFeedbackExecute(r UserFeedbacksAPICreateUserFeedbackRequest) (*UserFeedback, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UserFeedback
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserFeedbacksAPIService.CreateUserFeedback")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user_feedbacks.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.userFeedbackBody == nil {
		return localVarReturnValue, nil, reportError("userFeedbackBody is required and must be specified")
	}
	if r.userFeedbackCategory == nil {
		return localVarReturnValue, nil, reportError("userFeedbackCategory is required and must be specified")
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
	if r.userFeedbackUserId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user_feedback[user_id]", r.userFeedbackUserId, "")
	}
	if r.userFeedbackUserName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user_feedback[user_name]", r.userFeedbackUserName, "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "user_feedback[body]", r.userFeedbackBody, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "user_feedback[category]", r.userFeedbackCategory, "")
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
		if localVarHTTPResponse.StatusCode == 422 {
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

type UserFeedbacksAPIDeleteUserFeedbackRequest struct {
	ctx        context.Context
	ApiService *UserFeedbacksAPIService
	id         int32
}

func (r UserFeedbacksAPIDeleteUserFeedbackRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteUserFeedbackExecute(r)
}

/*
DeleteUserFeedback Delete User Feedback

You must be Moderator+.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the feedback.
	@return UserFeedbacksAPIDeleteUserFeedbackRequest
*/
func (a *UserFeedbacksAPIService) DeleteUserFeedback(ctx context.Context, id int32) UserFeedbacksAPIDeleteUserFeedbackRequest {
	return UserFeedbacksAPIDeleteUserFeedbackRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *UserFeedbacksAPIService) DeleteUserFeedbackExecute(r UserFeedbacksAPIDeleteUserFeedbackRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserFeedbacksAPIService.DeleteUserFeedback")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user_feedbacks/{id}/delete.json"
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type UserFeedbacksAPIDestroyUserFeedbackRequest struct {
	ctx        context.Context
	ApiService *UserFeedbacksAPIService
	id         int32
}

func (r UserFeedbacksAPIDestroyUserFeedbackRequest) Execute() (*http.Response, error) {
	return r.ApiService.DestroyUserFeedbackExecute(r)
}

/*
DestroyUserFeedback Destroy User Feedback

You must be Admin+, or the creator and Moderator+.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the feedback.
	@return UserFeedbacksAPIDestroyUserFeedbackRequest
*/
func (a *UserFeedbacksAPIService) DestroyUserFeedback(ctx context.Context, id int32) UserFeedbacksAPIDestroyUserFeedbackRequest {
	return UserFeedbacksAPIDestroyUserFeedbackRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *UserFeedbacksAPIService) DestroyUserFeedbackExecute(r UserFeedbacksAPIDestroyUserFeedbackRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserFeedbacksAPIService.DestroyUserFeedback")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user_feedbacks/{id}.json"
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type UserFeedbacksAPIEditUserFeedbackRequest struct {
	ctx                  context.Context
	ApiService           *UserFeedbacksAPIService
	id                   int32
	userFeedbackBody     *string
	userFeedbackCategory *FeedbackCategories
}

func (r UserFeedbacksAPIEditUserFeedbackRequest) UserFeedbackBody(userFeedbackBody string) UserFeedbacksAPIEditUserFeedbackRequest {
	r.userFeedbackBody = &userFeedbackBody
	return r
}

func (r UserFeedbacksAPIEditUserFeedbackRequest) UserFeedbackCategory(userFeedbackCategory FeedbackCategories) UserFeedbacksAPIEditUserFeedbackRequest {
	r.userFeedbackCategory = &userFeedbackCategory
	return r
}

func (r UserFeedbacksAPIEditUserFeedbackRequest) Execute() (*http.Response, error) {
	return r.ApiService.EditUserFeedbackExecute(r)
}

/*
EditUserFeedback Edit User Feedback

You must be Moderator+.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the feedback.
	@return UserFeedbacksAPIEditUserFeedbackRequest
*/
func (a *UserFeedbacksAPIService) EditUserFeedback(ctx context.Context, id int32) UserFeedbacksAPIEditUserFeedbackRequest {
	return UserFeedbacksAPIEditUserFeedbackRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *UserFeedbacksAPIService) EditUserFeedbackExecute(r UserFeedbacksAPIEditUserFeedbackRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserFeedbacksAPIService.EditUserFeedback")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user_feedbacks/{id}.json"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.userFeedbackBody != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user_feedback[body]", r.userFeedbackBody, "")
	}
	if r.userFeedbackCategory != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user_feedback[category]", r.userFeedbackCategory, "")
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type UserFeedbacksAPIGetUserFeedbackRequest struct {
	ctx        context.Context
	ApiService *UserFeedbacksAPIService
	id         int32
}

func (r UserFeedbacksAPIGetUserFeedbackRequest) Execute() (*UserFeedback, *http.Response, error) {
	return r.ApiService.GetUserFeedbackExecute(r)
}

/*
GetUserFeedback Get User Feedback

You must be Moderator+ if the feedback is deleted.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the feedback.
	@return UserFeedbacksAPIGetUserFeedbackRequest
*/
func (a *UserFeedbacksAPIService) GetUserFeedback(ctx context.Context, id int32) UserFeedbacksAPIGetUserFeedbackRequest {
	return UserFeedbacksAPIGetUserFeedbackRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return UserFeedback
func (a *UserFeedbacksAPIService) GetUserFeedbackExecute(r UserFeedbacksAPIGetUserFeedbackRequest) (*UserFeedback, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UserFeedback
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserFeedbacksAPIService.GetUserFeedback")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user_feedbacks/{id}.json"
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

type UserFeedbacksAPISearchUserFeedbacksRequest struct {
	ctx               context.Context
	ApiService        *UserFeedbacksAPIService
	limit             *int32
	page              *int32
	searchId          *int32
	searchOrder       *string
	searchDeleted     *string
	searchBodyMatches *string
	searchUserId      *int32
	searchUserName    *string
	searchCreatorId   *int32
	searchCreatorName *string
	searchCategory    *FeedbackCategories
}

// The maximum number of results to return. Between 0 and 320.
func (r UserFeedbacksAPISearchUserFeedbacksRequest) Limit(limit int32) UserFeedbacksAPISearchUserFeedbacksRequest {
	r.limit = &limit
	return r
}

// The page number of results to get. Between 1 and 750.
func (r UserFeedbacksAPISearchUserFeedbacksRequest) Page(page int32) UserFeedbacksAPISearchUserFeedbacksRequest {
	r.page = &page
	return r
}

// Search for a specific id.
func (r UserFeedbacksAPISearchUserFeedbacksRequest) SearchId(searchId int32) UserFeedbacksAPISearchUserFeedbacksRequest {
	r.searchId = &searchId
	return r
}

func (r UserFeedbacksAPISearchUserFeedbacksRequest) SearchOrder(searchOrder string) UserFeedbacksAPISearchUserFeedbacksRequest {
	r.searchOrder = &searchOrder
	return r
}

// You must be Moderator+.
func (r UserFeedbacksAPISearchUserFeedbacksRequest) SearchDeleted(searchDeleted string) UserFeedbacksAPISearchUserFeedbacksRequest {
	r.searchDeleted = &searchDeleted
	return r
}

func (r UserFeedbacksAPISearchUserFeedbacksRequest) SearchBodyMatches(searchBodyMatches string) UserFeedbacksAPISearchUserFeedbacksRequest {
	r.searchBodyMatches = &searchBodyMatches
	return r
}

func (r UserFeedbacksAPISearchUserFeedbacksRequest) SearchUserId(searchUserId int32) UserFeedbacksAPISearchUserFeedbacksRequest {
	r.searchUserId = &searchUserId
	return r
}

func (r UserFeedbacksAPISearchUserFeedbacksRequest) SearchUserName(searchUserName string) UserFeedbacksAPISearchUserFeedbacksRequest {
	r.searchUserName = &searchUserName
	return r
}

func (r UserFeedbacksAPISearchUserFeedbacksRequest) SearchCreatorId(searchCreatorId int32) UserFeedbacksAPISearchUserFeedbacksRequest {
	r.searchCreatorId = &searchCreatorId
	return r
}

func (r UserFeedbacksAPISearchUserFeedbacksRequest) SearchCreatorName(searchCreatorName string) UserFeedbacksAPISearchUserFeedbacksRequest {
	r.searchCreatorName = &searchCreatorName
	return r
}

func (r UserFeedbacksAPISearchUserFeedbacksRequest) SearchCategory(searchCategory FeedbackCategories) UserFeedbacksAPISearchUserFeedbacksRequest {
	r.searchCategory = &searchCategory
	return r
}

func (r UserFeedbacksAPISearchUserFeedbacksRequest) Execute() (*SearchUserFeedbacks200Response, *http.Response, error) {
	return r.ApiService.SearchUserFeedbacksExecute(r)
}

/*
SearchUserFeedbacks Search User Feedbacks

When no results are found, an object with an `user_feedbacks` key is returned.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return UserFeedbacksAPISearchUserFeedbacksRequest
*/
func (a *UserFeedbacksAPIService) SearchUserFeedbacks(ctx context.Context) UserFeedbacksAPISearchUserFeedbacksRequest {
	return UserFeedbacksAPISearchUserFeedbacksRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SearchUserFeedbacks200Response
func (a *UserFeedbacksAPIService) SearchUserFeedbacksExecute(r UserFeedbacksAPISearchUserFeedbacksRequest) (*SearchUserFeedbacks200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SearchUserFeedbacks200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserFeedbacksAPIService.SearchUserFeedbacks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user_feedbacks.json"

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
	if r.searchDeleted != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[deleted]", r.searchDeleted, "")
	}
	if r.searchBodyMatches != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[body_matches]", r.searchBodyMatches, "")
	}
	if r.searchUserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[user_id]", r.searchUserId, "")
	}
	if r.searchUserName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[user_name]", r.searchUserName, "")
	}
	if r.searchCreatorId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[creator_id]", r.searchCreatorId, "")
	}
	if r.searchCreatorName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[creator_name]", r.searchCreatorName, "")
	}
	if r.searchCategory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[category]", r.searchCategory, "")
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

type UserFeedbacksAPIUndeleteUserFeedbackRequest struct {
	ctx        context.Context
	ApiService *UserFeedbacksAPIService
	id         int32
}

func (r UserFeedbacksAPIUndeleteUserFeedbackRequest) Execute() (*http.Response, error) {
	return r.ApiService.UndeleteUserFeedbackExecute(r)
}

/*
UndeleteUserFeedback Undelete User Feedback

You must be Moderator+.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the feedback.
	@return UserFeedbacksAPIUndeleteUserFeedbackRequest
*/
func (a *UserFeedbacksAPIService) UndeleteUserFeedback(ctx context.Context, id int32) UserFeedbacksAPIUndeleteUserFeedbackRequest {
	return UserFeedbacksAPIUndeleteUserFeedbackRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *UserFeedbacksAPIService) UndeleteUserFeedbackExecute(r UserFeedbacksAPIUndeleteUserFeedbackRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserFeedbacksAPIService.UndeleteUserFeedback")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/user_feedbacks/{id}/undelete.json"
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
		if localVarHTTPResponse.StatusCode == 404 {
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
