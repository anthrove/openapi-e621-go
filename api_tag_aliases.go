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

// TagAliasesAPIService TagAliasesAPI service
type TagAliasesAPIService service

type TagAliasesAPIApproveTagAliasRequest struct {
	ctx        context.Context
	ApiService *TagAliasesAPIService
	id         float32
}

func (r TagAliasesAPIApproveTagAliasRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApproveTagAliasExecute(r)
}

/*
ApproveTagAlias Approve Tag Alias

You must be Admin+.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the tag alias.
	@return TagAliasesAPIApproveTagAliasRequest
*/
func (a *TagAliasesAPIService) ApproveTagAlias(ctx context.Context, id float32) TagAliasesAPIApproveTagAliasRequest {
	return TagAliasesAPIApproveTagAliasRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *TagAliasesAPIService) ApproveTagAliasExecute(r TagAliasesAPIApproveTagAliasRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagAliasesAPIService.ApproveTagAlias")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tag_aliases/{id}/approve.json"
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

type TagAliasesAPICreateTagAliasRequest struct {
	ctx                    context.Context
	ApiService             *TagAliasesAPIService
	tagAliasAntecedentName *string
	tagAliasConsequentName *string
	tagAliasReason         *string
	tagAliasSkipForum      *bool
}

func (r TagAliasesAPICreateTagAliasRequest) TagAliasAntecedentName(tagAliasAntecedentName string) TagAliasesAPICreateTagAliasRequest {
	r.tagAliasAntecedentName = &tagAliasAntecedentName
	return r
}

func (r TagAliasesAPICreateTagAliasRequest) TagAliasConsequentName(tagAliasConsequentName string) TagAliasesAPICreateTagAliasRequest {
	r.tagAliasConsequentName = &tagAliasConsequentName
	return r
}

func (r TagAliasesAPICreateTagAliasRequest) TagAliasReason(tagAliasReason string) TagAliasesAPICreateTagAliasRequest {
	r.tagAliasReason = &tagAliasReason
	return r
}

// Must be Admin+.
func (r TagAliasesAPICreateTagAliasRequest) TagAliasSkipForum(tagAliasSkipForum bool) TagAliasesAPICreateTagAliasRequest {
	r.tagAliasSkipForum = &tagAliasSkipForum
	return r
}

func (r TagAliasesAPICreateTagAliasRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateTagAliasExecute(r)
}

/*
CreateTagAlias Create Tag Alias

Errors will result in a 406 with no information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TagAliasesAPICreateTagAliasRequest
*/
func (a *TagAliasesAPIService) CreateTagAlias(ctx context.Context) TagAliasesAPICreateTagAliasRequest {
	return TagAliasesAPICreateTagAliasRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *TagAliasesAPIService) CreateTagAliasExecute(r TagAliasesAPICreateTagAliasRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagAliasesAPIService.CreateTagAlias")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tag_alias_requests.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tagAliasAntecedentName == nil {
		return nil, reportError("tagAliasAntecedentName is required and must be specified")
	}
	if r.tagAliasConsequentName == nil {
		return nil, reportError("tagAliasConsequentName is required and must be specified")
	}
	if r.tagAliasReason == nil {
		return nil, reportError("tagAliasReason is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "tag_alias[antecedent_name]", r.tagAliasAntecedentName, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "tag_alias[consequent_name]", r.tagAliasConsequentName, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "tag_alias[reason]", r.tagAliasReason, "")
	if r.tagAliasSkipForum != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "tag_alias[skip_forum]", r.tagAliasSkipForum, "")
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

type TagAliasesAPIEditTagAliasRequest struct {
	ctx                    context.Context
	ApiService             *TagAliasesAPIService
	id                     float32
	tagAliasAntecedentName *string
	tagAliasConsequentName *string
	tagAliasForumTopicId   *string
}

func (r TagAliasesAPIEditTagAliasRequest) TagAliasAntecedentName(tagAliasAntecedentName string) TagAliasesAPIEditTagAliasRequest {
	r.tagAliasAntecedentName = &tagAliasAntecedentName
	return r
}

func (r TagAliasesAPIEditTagAliasRequest) TagAliasConsequentName(tagAliasConsequentName string) TagAliasesAPIEditTagAliasRequest {
	r.tagAliasConsequentName = &tagAliasConsequentName
	return r
}

func (r TagAliasesAPIEditTagAliasRequest) TagAliasForumTopicId(tagAliasForumTopicId string) TagAliasesAPIEditTagAliasRequest {
	r.tagAliasForumTopicId = &tagAliasForumTopicId
	return r
}

func (r TagAliasesAPIEditTagAliasRequest) Execute() (*http.Response, error) {
	return r.ApiService.EditTagAliasExecute(r)
}

/*
EditTagAlias Edit Tag Alias

You must be Admin+.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the tag alias.
	@return TagAliasesAPIEditTagAliasRequest
*/
func (a *TagAliasesAPIService) EditTagAlias(ctx context.Context, id float32) TagAliasesAPIEditTagAliasRequest {
	return TagAliasesAPIEditTagAliasRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *TagAliasesAPIService) EditTagAliasExecute(r TagAliasesAPIEditTagAliasRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagAliasesAPIService.EditTagAlias")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tag_aliases/{id}.json"
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
	if r.tagAliasAntecedentName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "tag_alias[antecedent_name]", r.tagAliasAntecedentName, "")
	}
	if r.tagAliasConsequentName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "tag_alias[consequent_name]", r.tagAliasConsequentName, "")
	}
	if r.tagAliasForumTopicId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "tag_alias[forum_topic_id]", r.tagAliasForumTopicId, "")
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

type TagAliasesAPIGetTagAliasRequest struct {
	ctx        context.Context
	ApiService *TagAliasesAPIService
	id         float32
}

func (r TagAliasesAPIGetTagAliasRequest) Execute() (*TagAlias, *http.Response, error) {
	return r.ApiService.GetTagAliasExecute(r)
}

/*
GetTagAlias Get Tag Alias

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the tag alias.
	@return TagAliasesAPIGetTagAliasRequest
*/
func (a *TagAliasesAPIService) GetTagAlias(ctx context.Context, id float32) TagAliasesAPIGetTagAliasRequest {
	return TagAliasesAPIGetTagAliasRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return TagAlias
func (a *TagAliasesAPIService) GetTagAliasExecute(r TagAliasesAPIGetTagAliasRequest) (*TagAlias, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *TagAlias
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagAliasesAPIService.GetTagAlias")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tag_aliases/{id}.json"
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

type TagAliasesAPIRejectTagAliasRequest struct {
	ctx        context.Context
	ApiService *TagAliasesAPIService
	id         float32
}

func (r TagAliasesAPIRejectTagAliasRequest) Execute() (*http.Response, error) {
	return r.ApiService.RejectTagAliasExecute(r)
}

/*
RejectTagAlias Reject Tag Alias

You must be the creator of the request (if pending), or Admin+.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the tag alias.
	@return TagAliasesAPIRejectTagAliasRequest
*/
func (a *TagAliasesAPIService) RejectTagAlias(ctx context.Context, id float32) TagAliasesAPIRejectTagAliasRequest {
	return TagAliasesAPIRejectTagAliasRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *TagAliasesAPIService) RejectTagAliasExecute(r TagAliasesAPIRejectTagAliasRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagAliasesAPIService.RejectTagAlias")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tag_aliases/{id}.json"
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

type TagAliasesAPISearchTagAliasesRequest struct {
	ctx                         context.Context
	ApiService                  *TagAliasesAPIService
	limit                       *int32
	page                        *int32
	searchId                    *int32
	searchOrder                 *string
	searchNameMatches           *string
	searchAntecedentName        *string
	searchConsequentName        *string
	searchStatus                *TagRequestStatuses
	searchAntecedentTagCategory *TagCategories
	searchConsequentTagCategory *TagCategories
	searchCreatorId             *float32
	searchCreatorName           *string
	searchApproverId            *float32
	searchApproverName          *string
}

// The maximum number of results to return. Between 0 and 320.
func (r TagAliasesAPISearchTagAliasesRequest) Limit(limit int32) TagAliasesAPISearchTagAliasesRequest {
	r.limit = &limit
	return r
}

// The page number of results to get. Between 1 and 750.
func (r TagAliasesAPISearchTagAliasesRequest) Page(page int32) TagAliasesAPISearchTagAliasesRequest {
	r.page = &page
	return r
}

// Search for a specific id.
func (r TagAliasesAPISearchTagAliasesRequest) SearchId(searchId int32) TagAliasesAPISearchTagAliasesRequest {
	r.searchId = &searchId
	return r
}

func (r TagAliasesAPISearchTagAliasesRequest) SearchOrder(searchOrder string) TagAliasesAPISearchTagAliasesRequest {
	r.searchOrder = &searchOrder
	return r
}

func (r TagAliasesAPISearchTagAliasesRequest) SearchNameMatches(searchNameMatches string) TagAliasesAPISearchTagAliasesRequest {
	r.searchNameMatches = &searchNameMatches
	return r
}

func (r TagAliasesAPISearchTagAliasesRequest) SearchAntecedentName(searchAntecedentName string) TagAliasesAPISearchTagAliasesRequest {
	r.searchAntecedentName = &searchAntecedentName
	return r
}

func (r TagAliasesAPISearchTagAliasesRequest) SearchConsequentName(searchConsequentName string) TagAliasesAPISearchTagAliasesRequest {
	r.searchConsequentName = &searchConsequentName
	return r
}

func (r TagAliasesAPISearchTagAliasesRequest) SearchStatus(searchStatus TagRequestStatuses) TagAliasesAPISearchTagAliasesRequest {
	r.searchStatus = &searchStatus
	return r
}

func (r TagAliasesAPISearchTagAliasesRequest) SearchAntecedentTagCategory(searchAntecedentTagCategory TagCategories) TagAliasesAPISearchTagAliasesRequest {
	r.searchAntecedentTagCategory = &searchAntecedentTagCategory
	return r
}

func (r TagAliasesAPISearchTagAliasesRequest) SearchConsequentTagCategory(searchConsequentTagCategory TagCategories) TagAliasesAPISearchTagAliasesRequest {
	r.searchConsequentTagCategory = &searchConsequentTagCategory
	return r
}

func (r TagAliasesAPISearchTagAliasesRequest) SearchCreatorId(searchCreatorId float32) TagAliasesAPISearchTagAliasesRequest {
	r.searchCreatorId = &searchCreatorId
	return r
}

func (r TagAliasesAPISearchTagAliasesRequest) SearchCreatorName(searchCreatorName string) TagAliasesAPISearchTagAliasesRequest {
	r.searchCreatorName = &searchCreatorName
	return r
}

func (r TagAliasesAPISearchTagAliasesRequest) SearchApproverId(searchApproverId float32) TagAliasesAPISearchTagAliasesRequest {
	r.searchApproverId = &searchApproverId
	return r
}

func (r TagAliasesAPISearchTagAliasesRequest) SearchApproverName(searchApproverName string) TagAliasesAPISearchTagAliasesRequest {
	r.searchApproverName = &searchApproverName
	return r
}

func (r TagAliasesAPISearchTagAliasesRequest) Execute() (*SearchTagAliases200Response, *http.Response, error) {
	return r.ApiService.SearchTagAliasesExecute(r)
}

/*
SearchTagAliases Search Tag Aliases

When no results are found, an object with an `tag_aliases` key is returned.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TagAliasesAPISearchTagAliasesRequest
*/
func (a *TagAliasesAPIService) SearchTagAliases(ctx context.Context) TagAliasesAPISearchTagAliasesRequest {
	return TagAliasesAPISearchTagAliasesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SearchTagAliases200Response
func (a *TagAliasesAPIService) SearchTagAliasesExecute(r TagAliasesAPISearchTagAliasesRequest) (*SearchTagAliases200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SearchTagAliases200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagAliasesAPIService.SearchTagAliases")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tag_aliases.json"

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
