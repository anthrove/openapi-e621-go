/*
E621

OpenAPI definition for E621's API. You can find the source [here](https://github.com/DonovanDMC/E621OpenAPI)<br> This document is intended to compliment E621's existing [API Documentation](https://e621.net/help/api).<br> Note if E621's api is under attack and/or cloudflare protections are enabled, the \"Try it out\" buttons here will not work.<br> If they are not working, you can check this [Unofficial Status Page](https://status.e621.ws). 

API version: d69c34e
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi-e621-go

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// TagImplicationsAPIService TagImplicationsAPI service
type TagImplicationsAPIService service

type ApiApproveTagImplicationRequest struct {
	ctx context.Context
	ApiService *TagImplicationsAPIService
	id float32
}

func (r ApiApproveTagImplicationRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApproveTagImplicationExecute(r)
}

/*
ApproveTagImplication Approve Tag Implication

You must be Admin+.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the tag implication.
 @return ApiApproveTagImplicationRequest
*/
func (a *TagImplicationsAPIService) ApproveTagImplication(ctx context.Context, id float32) ApiApproveTagImplicationRequest {
	return ApiApproveTagImplicationRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *TagImplicationsAPIService) ApproveTagImplicationExecute(r ApiApproveTagImplicationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagImplicationsAPIService.ApproveTagImplication")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tag_implications/{id}/approve.json"
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

type ApiCreateTagImplicationRequest struct {
	ctx context.Context
	ApiService *TagImplicationsAPIService
	tagImplicationAntecedentName *string
	tagImplicationConsequentName *string
	tagImplicationReason *string
	tagImplicationSkipForum *bool
}

func (r ApiCreateTagImplicationRequest) TagImplicationAntecedentName(tagImplicationAntecedentName string) ApiCreateTagImplicationRequest {
	r.tagImplicationAntecedentName = &tagImplicationAntecedentName
	return r
}

func (r ApiCreateTagImplicationRequest) TagImplicationConsequentName(tagImplicationConsequentName string) ApiCreateTagImplicationRequest {
	r.tagImplicationConsequentName = &tagImplicationConsequentName
	return r
}

func (r ApiCreateTagImplicationRequest) TagImplicationReason(tagImplicationReason string) ApiCreateTagImplicationRequest {
	r.tagImplicationReason = &tagImplicationReason
	return r
}

// Must be Admin+.
func (r ApiCreateTagImplicationRequest) TagImplicationSkipForum(tagImplicationSkipForum bool) ApiCreateTagImplicationRequest {
	r.tagImplicationSkipForum = &tagImplicationSkipForum
	return r
}

func (r ApiCreateTagImplicationRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateTagImplicationExecute(r)
}

/*
CreateTagImplication Create Tag Implication

Errors will result in a 406 with no information.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateTagImplicationRequest
*/
func (a *TagImplicationsAPIService) CreateTagImplication(ctx context.Context) ApiCreateTagImplicationRequest {
	return ApiCreateTagImplicationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *TagImplicationsAPIService) CreateTagImplicationExecute(r ApiCreateTagImplicationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagImplicationsAPIService.CreateTagImplication")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tag_implication_requests.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tagImplicationAntecedentName == nil {
		return nil, reportError("tagImplicationAntecedentName is required and must be specified")
	}
	if r.tagImplicationConsequentName == nil {
		return nil, reportError("tagImplicationConsequentName is required and must be specified")
	}
	if r.tagImplicationReason == nil {
		return nil, reportError("tagImplicationReason is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "tag_implication[antecedent_name]", r.tagImplicationAntecedentName, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "tag_implication[consequent_name]", r.tagImplicationConsequentName, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "tag_implication[reason]", r.tagImplicationReason, "")
	if r.tagImplicationSkipForum != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "tag_implication[skip_forum]", r.tagImplicationSkipForum, "")
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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiEditTagImplicationRequest struct {
	ctx context.Context
	ApiService *TagImplicationsAPIService
	id float32
	tagImplicationAntecedentName *string
	tagImplicationConsequentName *string
	tagImplicationForumTopicId *string
}

func (r ApiEditTagImplicationRequest) TagImplicationAntecedentName(tagImplicationAntecedentName string) ApiEditTagImplicationRequest {
	r.tagImplicationAntecedentName = &tagImplicationAntecedentName
	return r
}

func (r ApiEditTagImplicationRequest) TagImplicationConsequentName(tagImplicationConsequentName string) ApiEditTagImplicationRequest {
	r.tagImplicationConsequentName = &tagImplicationConsequentName
	return r
}

func (r ApiEditTagImplicationRequest) TagImplicationForumTopicId(tagImplicationForumTopicId string) ApiEditTagImplicationRequest {
	r.tagImplicationForumTopicId = &tagImplicationForumTopicId
	return r
}

func (r ApiEditTagImplicationRequest) Execute() (*http.Response, error) {
	return r.ApiService.EditTagImplicationExecute(r)
}

/*
EditTagImplication Edit Tag Implication

You must be Admin+.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the tag implication.
 @return ApiEditTagImplicationRequest
*/
func (a *TagImplicationsAPIService) EditTagImplication(ctx context.Context, id float32) ApiEditTagImplicationRequest {
	return ApiEditTagImplicationRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *TagImplicationsAPIService) EditTagImplicationExecute(r ApiEditTagImplicationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagImplicationsAPIService.EditTagImplication")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tag_implications/{id}.json"
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
	if r.tagImplicationAntecedentName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "tag_implication[antecedent_name]", r.tagImplicationAntecedentName, "")
	}
	if r.tagImplicationConsequentName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "tag_implication[consequent_name]", r.tagImplicationConsequentName, "")
	}
	if r.tagImplicationForumTopicId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "tag_implication[forum_topic_id]", r.tagImplicationForumTopicId, "")
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

type ApiGetTagImplicationRequest struct {
	ctx context.Context
	ApiService *TagImplicationsAPIService
	id float32
}

func (r ApiGetTagImplicationRequest) Execute() (*TagImplication, *http.Response, error) {
	return r.ApiService.GetTagImplicationExecute(r)
}

/*
GetTagImplication Get Tag Implication

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the tag implication.
 @return ApiGetTagImplicationRequest
*/
func (a *TagImplicationsAPIService) GetTagImplication(ctx context.Context, id float32) ApiGetTagImplicationRequest {
	return ApiGetTagImplicationRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return TagImplication
func (a *TagImplicationsAPIService) GetTagImplicationExecute(r ApiGetTagImplicationRequest) (*TagImplication, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TagImplication
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagImplicationsAPIService.GetTagImplication")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tag_implications/{id}.json"
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

type ApiRejectTagImplicationRequest struct {
	ctx context.Context
	ApiService *TagImplicationsAPIService
	id float32
}

func (r ApiRejectTagImplicationRequest) Execute() (*http.Response, error) {
	return r.ApiService.RejectTagImplicationExecute(r)
}

/*
RejectTagImplication Reject Tag Implication

You must be the creator of the request (if pending), or Admin+.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the tag implication.
 @return ApiRejectTagImplicationRequest
*/
func (a *TagImplicationsAPIService) RejectTagImplication(ctx context.Context, id float32) ApiRejectTagImplicationRequest {
	return ApiRejectTagImplicationRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *TagImplicationsAPIService) RejectTagImplicationExecute(r ApiRejectTagImplicationRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagImplicationsAPIService.RejectTagImplication")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tag_implications/{id}.json"
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

type ApiSearchTagImplicationsRequest struct {
	ctx context.Context
	ApiService *TagImplicationsAPIService
	limit *float32
	page *float32
	searchId *float32
	searchOrder *string
	searchNameMatches *string
	searchAntecedentName *string
	searchConsequentName *string
	searchStatus *TagRequestStatuses
	searchAntecedentTagCategory *TagCategories
	searchConsequentTagCategory *TagCategories
	searchCreatorId *float32
	searchCreatorName *string
	searchApproverId *float32
	searchApproverName *string
}

// The maximum number of results to return. Between 0 and 320.
func (r ApiSearchTagImplicationsRequest) Limit(limit float32) ApiSearchTagImplicationsRequest {
	r.limit = &limit
	return r
}

// The page number of results to get. Between 1 and 750.
func (r ApiSearchTagImplicationsRequest) Page(page float32) ApiSearchTagImplicationsRequest {
	r.page = &page
	return r
}

// Search for a specific id.
func (r ApiSearchTagImplicationsRequest) SearchId(searchId float32) ApiSearchTagImplicationsRequest {
	r.searchId = &searchId
	return r
}

func (r ApiSearchTagImplicationsRequest) SearchOrder(searchOrder string) ApiSearchTagImplicationsRequest {
	r.searchOrder = &searchOrder
	return r
}

func (r ApiSearchTagImplicationsRequest) SearchNameMatches(searchNameMatches string) ApiSearchTagImplicationsRequest {
	r.searchNameMatches = &searchNameMatches
	return r
}

func (r ApiSearchTagImplicationsRequest) SearchAntecedentName(searchAntecedentName string) ApiSearchTagImplicationsRequest {
	r.searchAntecedentName = &searchAntecedentName
	return r
}

func (r ApiSearchTagImplicationsRequest) SearchConsequentName(searchConsequentName string) ApiSearchTagImplicationsRequest {
	r.searchConsequentName = &searchConsequentName
	return r
}

func (r ApiSearchTagImplicationsRequest) SearchStatus(searchStatus TagRequestStatuses) ApiSearchTagImplicationsRequest {
	r.searchStatus = &searchStatus
	return r
}

func (r ApiSearchTagImplicationsRequest) SearchAntecedentTagCategory(searchAntecedentTagCategory TagCategories) ApiSearchTagImplicationsRequest {
	r.searchAntecedentTagCategory = &searchAntecedentTagCategory
	return r
}

func (r ApiSearchTagImplicationsRequest) SearchConsequentTagCategory(searchConsequentTagCategory TagCategories) ApiSearchTagImplicationsRequest {
	r.searchConsequentTagCategory = &searchConsequentTagCategory
	return r
}

func (r ApiSearchTagImplicationsRequest) SearchCreatorId(searchCreatorId float32) ApiSearchTagImplicationsRequest {
	r.searchCreatorId = &searchCreatorId
	return r
}

func (r ApiSearchTagImplicationsRequest) SearchCreatorName(searchCreatorName string) ApiSearchTagImplicationsRequest {
	r.searchCreatorName = &searchCreatorName
	return r
}

func (r ApiSearchTagImplicationsRequest) SearchApproverId(searchApproverId float32) ApiSearchTagImplicationsRequest {
	r.searchApproverId = &searchApproverId
	return r
}

func (r ApiSearchTagImplicationsRequest) SearchApproverName(searchApproverName string) ApiSearchTagImplicationsRequest {
	r.searchApproverName = &searchApproverName
	return r
}

func (r ApiSearchTagImplicationsRequest) Execute() (*SearchTagImplications200Response, *http.Response, error) {
	return r.ApiService.SearchTagImplicationsExecute(r)
}

/*
SearchTagImplications Search Tag Implications

When no results are found, an object with an `tag_implications` key is returned.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchTagImplicationsRequest
*/
func (a *TagImplicationsAPIService) SearchTagImplications(ctx context.Context) ApiSearchTagImplicationsRequest {
	return ApiSearchTagImplicationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SearchTagImplications200Response
func (a *TagImplicationsAPIService) SearchTagImplicationsExecute(r ApiSearchTagImplicationsRequest) (*SearchTagImplications200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SearchTagImplications200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagImplicationsAPIService.SearchTagImplications")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tag_implications.json"

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
	if r.searchNameMatches != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[name_matches]", r.searchNameMatches, "")
	}
	if r.searchAntecedentName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[antecedent_name]", r.searchAntecedentName, "")
	}
	if r.searchConsequentName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[consequent_name]", r.searchConsequentName, "")
	}
	if r.searchStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[status]", r.searchStatus, "")
	}
	if r.searchAntecedentTagCategory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[antecedent_tag_category]", r.searchAntecedentTagCategory, "")
	}
	if r.searchConsequentTagCategory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[consequent_tag_category]", r.searchConsequentTagCategory, "")
	}
	if r.searchCreatorId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[creator_id]", r.searchCreatorId, "")
	}
	if r.searchCreatorName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[creator_name]", r.searchCreatorName, "")
	}
	if r.searchApproverId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[approver_id]", r.searchApproverId, "")
	}
	if r.searchApproverName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[approver_name]", r.searchApproverName, "")
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
