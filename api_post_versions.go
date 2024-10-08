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
)


// PostVersionsAPIService PostVersionsAPI service
type PostVersionsAPIService service

type ApiSearchPostVersionsRequest struct {
	ctx context.Context
	ApiService *PostVersionsAPIService
	limit *float32
	page *float32
	searchId *float32
	searchUpdaterName *string
	searchUpdaterId *float32
	searchPostId *float32
	searchStartId *float32
	searchRating *Ratings
	searchRatingChanged *string
	searchParentId *float32
	searchParentIdChanged *bool
	searchTags *string
	searchTagsRemoved *string
	searchTagsAdded *string
	searchLockedTags *string
	searchLockedTagsRemoved *string
	searchLockedTagsAdded *string
	searchReason *string
	searchDescription *string
	searchDescriptionChanged *bool
	searchSourceChanged *bool
	searchUploads *string
}

// The maximum number of results to return. Between 0 and 320.
func (r ApiSearchPostVersionsRequest) Limit(limit float32) ApiSearchPostVersionsRequest {
	r.limit = &limit
	return r
}

// The page number of results to get. Between 1 and 750. Note that for post versions specifically, you can only go through the 10,000 most recent results with page numbers.
func (r ApiSearchPostVersionsRequest) Page(page float32) ApiSearchPostVersionsRequest {
	r.page = &page
	return r
}

// Search for a specific id.
func (r ApiSearchPostVersionsRequest) SearchId(searchId float32) ApiSearchPostVersionsRequest {
	r.searchId = &searchId
	return r
}

func (r ApiSearchPostVersionsRequest) SearchUpdaterName(searchUpdaterName string) ApiSearchPostVersionsRequest {
	r.searchUpdaterName = &searchUpdaterName
	return r
}

func (r ApiSearchPostVersionsRequest) SearchUpdaterId(searchUpdaterId float32) ApiSearchPostVersionsRequest {
	r.searchUpdaterId = &searchUpdaterId
	return r
}

func (r ApiSearchPostVersionsRequest) SearchPostId(searchPostId float32) ApiSearchPostVersionsRequest {
	r.searchPostId = &searchPostId
	return r
}

func (r ApiSearchPostVersionsRequest) SearchStartId(searchStartId float32) ApiSearchPostVersionsRequest {
	r.searchStartId = &searchStartId
	return r
}

func (r ApiSearchPostVersionsRequest) SearchRating(searchRating Ratings) ApiSearchPostVersionsRequest {
	r.searchRating = &searchRating
	return r
}

func (r ApiSearchPostVersionsRequest) SearchRatingChanged(searchRatingChanged string) ApiSearchPostVersionsRequest {
	r.searchRatingChanged = &searchRatingChanged
	return r
}

func (r ApiSearchPostVersionsRequest) SearchParentId(searchParentId float32) ApiSearchPostVersionsRequest {
	r.searchParentId = &searchParentId
	return r
}

func (r ApiSearchPostVersionsRequest) SearchParentIdChanged(searchParentIdChanged bool) ApiSearchPostVersionsRequest {
	r.searchParentIdChanged = &searchParentIdChanged
	return r
}

func (r ApiSearchPostVersionsRequest) SearchTags(searchTags string) ApiSearchPostVersionsRequest {
	r.searchTags = &searchTags
	return r
}

func (r ApiSearchPostVersionsRequest) SearchTagsRemoved(searchTagsRemoved string) ApiSearchPostVersionsRequest {
	r.searchTagsRemoved = &searchTagsRemoved
	return r
}

func (r ApiSearchPostVersionsRequest) SearchTagsAdded(searchTagsAdded string) ApiSearchPostVersionsRequest {
	r.searchTagsAdded = &searchTagsAdded
	return r
}

func (r ApiSearchPostVersionsRequest) SearchLockedTags(searchLockedTags string) ApiSearchPostVersionsRequest {
	r.searchLockedTags = &searchLockedTags
	return r
}

func (r ApiSearchPostVersionsRequest) SearchLockedTagsRemoved(searchLockedTagsRemoved string) ApiSearchPostVersionsRequest {
	r.searchLockedTagsRemoved = &searchLockedTagsRemoved
	return r
}

func (r ApiSearchPostVersionsRequest) SearchLockedTagsAdded(searchLockedTagsAdded string) ApiSearchPostVersionsRequest {
	r.searchLockedTagsAdded = &searchLockedTagsAdded
	return r
}

func (r ApiSearchPostVersionsRequest) SearchReason(searchReason string) ApiSearchPostVersionsRequest {
	r.searchReason = &searchReason
	return r
}

func (r ApiSearchPostVersionsRequest) SearchDescription(searchDescription string) ApiSearchPostVersionsRequest {
	r.searchDescription = &searchDescription
	return r
}

func (r ApiSearchPostVersionsRequest) SearchDescriptionChanged(searchDescriptionChanged bool) ApiSearchPostVersionsRequest {
	r.searchDescriptionChanged = &searchDescriptionChanged
	return r
}

func (r ApiSearchPostVersionsRequest) SearchSourceChanged(searchSourceChanged bool) ApiSearchPostVersionsRequest {
	r.searchSourceChanged = &searchSourceChanged
	return r
}

func (r ApiSearchPostVersionsRequest) SearchUploads(searchUploads string) ApiSearchPostVersionsRequest {
	r.searchUploads = &searchUploads
	return r
}

func (r ApiSearchPostVersionsRequest) Execute() ([]PostVersion, *http.Response, error) {
	return r.ApiService.SearchPostVersionsExecute(r)
}

/*
SearchPostVersions Search Post Versions

Errors if no results are found.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchPostVersionsRequest
*/
func (a *PostVersionsAPIService) SearchPostVersions(ctx context.Context) ApiSearchPostVersionsRequest {
	return ApiSearchPostVersionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []PostVersion
func (a *PostVersionsAPIService) SearchPostVersionsExecute(r ApiSearchPostVersionsRequest) ([]PostVersion, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []PostVersion
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostVersionsAPIService.SearchPostVersions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/post_versions.json"

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
	if r.searchUpdaterName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[updater_name]", r.searchUpdaterName, "")
	}
	if r.searchUpdaterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[updater_id]", r.searchUpdaterId, "")
	}
	if r.searchPostId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[post_id]", r.searchPostId, "")
	}
	if r.searchStartId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[start_id]", r.searchStartId, "")
	}
	if r.searchRating != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[rating]", r.searchRating, "")
	}
	if r.searchRatingChanged != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[rating_changed]", r.searchRatingChanged, "")
	}
	if r.searchParentId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[parent_id]", r.searchParentId, "")
	}
	if r.searchParentIdChanged != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[parent_id_changed]", r.searchParentIdChanged, "")
	}
	if r.searchTags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[tags]", r.searchTags, "")
	}
	if r.searchTagsRemoved != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[tags_removed]", r.searchTagsRemoved, "")
	}
	if r.searchTagsAdded != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[tags_added]", r.searchTagsAdded, "")
	}
	if r.searchLockedTags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[locked_tags]", r.searchLockedTags, "")
	}
	if r.searchLockedTagsRemoved != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[locked_tags_removed]", r.searchLockedTagsRemoved, "")
	}
	if r.searchLockedTagsAdded != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[locked_tags_added]", r.searchLockedTagsAdded, "")
	}
	if r.searchReason != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[reason]", r.searchReason, "")
	}
	if r.searchDescription != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[description]", r.searchDescription, "")
	}
	if r.searchDescriptionChanged != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[description_changed]", r.searchDescriptionChanged, "")
	}
	if r.searchSourceChanged != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[source_changed]", r.searchSourceChanged, "")
	}
	if r.searchUploads != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[uploads]", r.searchUploads, "")
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