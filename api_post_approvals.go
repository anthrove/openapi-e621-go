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
)

// PostApprovalsAPIService PostApprovalsAPI service
type PostApprovalsAPIService service

type PostApprovalsAPIApprovePostRequest struct {
	ctx        context.Context
	ApiService *PostApprovalsAPIService
	postId     *float32
}

func (r PostApprovalsAPIApprovePostRequest) PostId(postId float32) PostApprovalsAPIApprovePostRequest {
	r.postId = &postId
	return r
}

func (r PostApprovalsAPIApprovePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApprovePostExecute(r)
}

/*
ApprovePost Approve Post

You must have the "Approve Posts" permission.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PostApprovalsAPIApprovePostRequest
*/
func (a *PostApprovalsAPIService) ApprovePost(ctx context.Context) PostApprovalsAPIApprovePostRequest {
	return PostApprovalsAPIApprovePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *PostApprovalsAPIService) ApprovePostExecute(r PostApprovalsAPIApprovePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostApprovalsAPIService.ApprovePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/moderator/post/approval.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.postId == nil {
		return nil, reportError("postId is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "post_id", r.postId, "")
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

type PostApprovalsAPISearchPostApprovalsRequest struct {
	ctx                 context.Context
	ApiService          *PostApprovalsAPIService
	limit               *int32
	page                *int32
	searchId            *int32
	searchOrder         *string
	searchPostTagsMatch *string
	searchUserId        *float32
	searchUserName      *string
	searchPostId        *float32
}

// The maximum number of results to return. Between 0 and 320.
func (r PostApprovalsAPISearchPostApprovalsRequest) Limit(limit int32) PostApprovalsAPISearchPostApprovalsRequest {
	r.limit = &limit
	return r
}

// The page number of results to get. Between 1 and 750.
func (r PostApprovalsAPISearchPostApprovalsRequest) Page(page int32) PostApprovalsAPISearchPostApprovalsRequest {
	r.page = &page
	return r
}

// Search for a specific id.
func (r PostApprovalsAPISearchPostApprovalsRequest) SearchId(searchId int32) PostApprovalsAPISearchPostApprovalsRequest {
	r.searchId = &searchId
	return r
}

func (r PostApprovalsAPISearchPostApprovalsRequest) SearchOrder(searchOrder string) PostApprovalsAPISearchPostApprovalsRequest {
	r.searchOrder = &searchOrder
	return r
}

func (r PostApprovalsAPISearchPostApprovalsRequest) SearchPostTagsMatch(searchPostTagsMatch string) PostApprovalsAPISearchPostApprovalsRequest {
	r.searchPostTagsMatch = &searchPostTagsMatch
	return r
}

func (r PostApprovalsAPISearchPostApprovalsRequest) SearchUserId(searchUserId float32) PostApprovalsAPISearchPostApprovalsRequest {
	r.searchUserId = &searchUserId
	return r
}

func (r PostApprovalsAPISearchPostApprovalsRequest) SearchUserName(searchUserName string) PostApprovalsAPISearchPostApprovalsRequest {
	r.searchUserName = &searchUserName
	return r
}

func (r PostApprovalsAPISearchPostApprovalsRequest) SearchPostId(searchPostId float32) PostApprovalsAPISearchPostApprovalsRequest {
	r.searchPostId = &searchPostId
	return r
}

func (r PostApprovalsAPISearchPostApprovalsRequest) Execute() (*SearchPostApprovals200Response, *http.Response, error) {
	return r.ApiService.SearchPostApprovalsExecute(r)
}

/*
SearchPostApprovals Search Post Approvals

When no results are found, an object with an `post_approvals` key is returned.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PostApprovalsAPISearchPostApprovalsRequest
*/
func (a *PostApprovalsAPIService) SearchPostApprovals(ctx context.Context) PostApprovalsAPISearchPostApprovalsRequest {
	return PostApprovalsAPISearchPostApprovalsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SearchPostApprovals200Response
func (a *PostApprovalsAPIService) SearchPostApprovalsExecute(r PostApprovalsAPISearchPostApprovalsRequest) (*SearchPostApprovals200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SearchPostApprovals200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostApprovalsAPIService.SearchPostApprovals")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/post_approvals.json"

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
	if r.searchPostTagsMatch != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[post_tags_match]", r.searchPostTagsMatch, "")
	}
	if r.searchUserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[user_id]", r.searchUserId, "")
	}
	if r.searchUserName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[user_name]", r.searchUserName, "")
	}
	if r.searchPostId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[post_id]", r.searchPostId, "")
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

type PostApprovalsAPIUnapprovePostRequest struct {
	ctx        context.Context
	ApiService *PostApprovalsAPIService
	postId     *float32
}

func (r PostApprovalsAPIUnapprovePostRequest) PostId(postId float32) PostApprovalsAPIUnapprovePostRequest {
	r.postId = &postId
	return r
}

func (r PostApprovalsAPIUnapprovePostRequest) Execute() (*http.Response, error) {
	return r.ApiService.UnapprovePostExecute(r)
}

/*
UnapprovePost Unapprove Post

You must have the "Approve Posts" permission. The response does not differ for success or failure.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PostApprovalsAPIUnapprovePostRequest
*/
func (a *PostApprovalsAPIService) UnapprovePost(ctx context.Context) PostApprovalsAPIUnapprovePostRequest {
	return PostApprovalsAPIUnapprovePostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *PostApprovalsAPIService) UnapprovePostExecute(r PostApprovalsAPIUnapprovePostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostApprovalsAPIService.UnapprovePost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/moderator/post/approval.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.postId == nil {
		return nil, reportError("postId is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "post_id", r.postId, "")
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
