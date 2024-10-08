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


// TagsAPIService TagsAPI service
type TagsAPIService service

type ApiCorrectTagRequest struct {
	ctx context.Context
	ApiService *TagsAPIService
	id float32
	commit *string
}

// If not set, nothing will happen.
func (r ApiCorrectTagRequest) Commit(commit string) ApiCorrectTagRequest {
	r.commit = &commit
	return r
}

func (r ApiCorrectTagRequest) Execute() (*http.Response, error) {
	return r.ApiService.CorrectTagExecute(r)
}

/*
CorrectTag Correct Tag

You must be Janitor+. `commit=Fix` must be set.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the tag.
 @return ApiCorrectTagRequest
*/
func (a *TagsAPIService) CorrectTag(ctx context.Context, id float32) ApiCorrectTagRequest {
	return ApiCorrectTagRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *TagsAPIService) CorrectTagExecute(r ApiCorrectTagRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsAPIService.CorrectTag")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tags/{id}/correction.json"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.commit == nil {
		return nil, reportError("commit is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "commit", r.commit, "")
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

type ApiEditTagRequest struct {
	ctx context.Context
	ApiService *TagsAPIService
	id float32
	tagCategory *TagCategories
	tagIsLocked *bool
}

func (r ApiEditTagRequest) TagCategory(tagCategory TagCategories) ApiEditTagRequest {
	r.tagCategory = &tagCategory
	return r
}

// Must be Admin+.
func (r ApiEditTagRequest) TagIsLocked(tagIsLocked bool) ApiEditTagRequest {
	r.tagIsLocked = &tagIsLocked
	return r
}

func (r ApiEditTagRequest) Execute() (*http.Response, error) {
	return r.ApiService.EditTagExecute(r)
}

/*
EditTag Edit Tag

Must be Admin+ if the tag is locked or post count is >100.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the tag.
 @return ApiEditTagRequest
*/
func (a *TagsAPIService) EditTag(ctx context.Context, id float32) ApiEditTagRequest {
	return ApiEditTagRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *TagsAPIService) EditTagExecute(r ApiEditTagRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsAPIService.EditTag")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tags/{id}.json"
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
	if r.tagCategory != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "tag[category]", r.tagCategory, "")
	}
	if r.tagIsLocked != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "tag[is_locked]", r.tagIsLocked, "")
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

type ApiGetTagRequest struct {
	ctx context.Context
	ApiService *TagsAPIService
	id GetArtistIdOrNameParameter
}

func (r ApiGetTagRequest) Execute() (*Tag, *http.Response, error) {
	return r.ApiService.GetTagExecute(r)
}

/*
GetTag Get Tag

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID or name of the tag.
 @return ApiGetTagRequest
*/
func (a *TagsAPIService) GetTag(ctx context.Context, id GetArtistIdOrNameParameter) ApiGetTagRequest {
	return ApiGetTagRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Tag
func (a *TagsAPIService) GetTagExecute(r ApiGetTagRequest) (*Tag, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Tag
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsAPIService.GetTag")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tags/{id}.json"
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

type ApiPreviewTagsRequest struct {
	ctx context.Context
	ApiService *TagsAPIService
	tags *string
}

// The tags to preview, space separated.
func (r ApiPreviewTagsRequest) Tags(tags string) ApiPreviewTagsRequest {
	r.tags = &tags
	return r
}

func (r ApiPreviewTagsRequest) Execute() ([]TagPreview, *http.Response, error) {
	return r.ApiService.PreviewTagsExecute(r)
}

/*
PreviewTags Preview Tags

Note while this route does not require auth, without auth it requires a CSRF token. For that reason it has been marked as requiring auth.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiPreviewTagsRequest
*/
func (a *TagsAPIService) PreviewTags(ctx context.Context) ApiPreviewTagsRequest {
	return ApiPreviewTagsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []TagPreview
func (a *TagsAPIService) PreviewTagsExecute(r ApiPreviewTagsRequest) ([]TagPreview, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []TagPreview
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsAPIService.PreviewTags")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tags/preview.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.tags == nil {
		return localVarReturnValue, nil, reportError("tags is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "tags", r.tags, "")
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

type ApiSearchTagsRequest struct {
	ctx context.Context
	ApiService *TagsAPIService
	limit *float32
	page *float32
	searchId *float32
	searchOrder *string
	searchFuzzyNameMatches *string
	searchNameMatches *string
	searchName *string
	searchCategory *TagCategories
	searchHideEmpty *bool
	searchHasWiki *bool
	searchHasArtist *bool
}

// The maximum number of results to return. Between 0 and 320.
func (r ApiSearchTagsRequest) Limit(limit float32) ApiSearchTagsRequest {
	r.limit = &limit
	return r
}

// The page number of results to get. Between 1 and 750.
func (r ApiSearchTagsRequest) Page(page float32) ApiSearchTagsRequest {
	r.page = &page
	return r
}

// Search for a specific id.
func (r ApiSearchTagsRequest) SearchId(searchId float32) ApiSearchTagsRequest {
	r.searchId = &searchId
	return r
}

func (r ApiSearchTagsRequest) SearchOrder(searchOrder string) ApiSearchTagsRequest {
	r.searchOrder = &searchOrder
	return r
}

func (r ApiSearchTagsRequest) SearchFuzzyNameMatches(searchFuzzyNameMatches string) ApiSearchTagsRequest {
	r.searchFuzzyNameMatches = &searchFuzzyNameMatches
	return r
}

func (r ApiSearchTagsRequest) SearchNameMatches(searchNameMatches string) ApiSearchTagsRequest {
	r.searchNameMatches = &searchNameMatches
	return r
}

func (r ApiSearchTagsRequest) SearchName(searchName string) ApiSearchTagsRequest {
	r.searchName = &searchName
	return r
}

func (r ApiSearchTagsRequest) SearchCategory(searchCategory TagCategories) ApiSearchTagsRequest {
	r.searchCategory = &searchCategory
	return r
}

func (r ApiSearchTagsRequest) SearchHideEmpty(searchHideEmpty bool) ApiSearchTagsRequest {
	r.searchHideEmpty = &searchHideEmpty
	return r
}

func (r ApiSearchTagsRequest) SearchHasWiki(searchHasWiki bool) ApiSearchTagsRequest {
	r.searchHasWiki = &searchHasWiki
	return r
}

func (r ApiSearchTagsRequest) SearchHasArtist(searchHasArtist bool) ApiSearchTagsRequest {
	r.searchHasArtist = &searchHasArtist
	return r
}

func (r ApiSearchTagsRequest) Execute() (*SearchTags200Response, *http.Response, error) {
	return r.ApiService.SearchTagsExecute(r)
}

/*
SearchTags Search Tags

When no results are found, an object with an `tags` key is returned.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchTagsRequest
*/
func (a *TagsAPIService) SearchTags(ctx context.Context) ApiSearchTagsRequest {
	return ApiSearchTagsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SearchTags200Response
func (a *TagsAPIService) SearchTagsExecute(r ApiSearchTagsRequest) (*SearchTags200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SearchTags200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TagsAPIService.SearchTags")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tags.json"

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
	if r.searchFuzzyNameMatches != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[fuzzy_name_matches]", r.searchFuzzyNameMatches, "")
	}
	if r.searchNameMatches != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[name_matches]", r.searchNameMatches, "")
	}
	if r.searchName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[name]", r.searchName, "")
	}
	if r.searchCategory != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[category]", r.searchCategory, "")
	}
	if r.searchHideEmpty != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[hide_empty]", r.searchHideEmpty, "")
	}
	if r.searchHasWiki != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[has_wiki]", r.searchHasWiki, "")
	}
	if r.searchHasArtist != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[has_artist]", r.searchHasArtist, "")
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