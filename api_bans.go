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

// BansAPIService BansAPI service
type BansAPIService service

type BansAPIGetBanRequest struct {
	ctx        context.Context
	ApiService *BansAPIService
	id         float32
}

func (r BansAPIGetBanRequest) Execute() (*Ban, *http.Response, error) {
	return r.ApiService.GetBanExecute(r)
}

/*
GetBan Get Ban

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the ban to get.
	@return BansAPIGetBanRequest
*/
func (a *BansAPIService) GetBan(ctx context.Context, id float32) BansAPIGetBanRequest {
	return BansAPIGetBanRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return Ban
func (a *BansAPIService) GetBanExecute(r BansAPIGetBanRequest) (*Ban, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Ban
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BansAPIService.GetBan")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bans/{id}.json"
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

type BansAPISearchBansRequest struct {
	ctx                 context.Context
	ApiService          *BansAPIService
	limit               *int32
	page                *int32
	searchId            *int32
	searchOrder         *string
	searchBannerId      *string
	searchBannerName    *string
	searchUserId        *string
	searchUserName      *string
	searchReasonMatches *string
	searchExpired       *string
}

// The maximum number of results to return. Between 0 and 320.
func (r BansAPISearchBansRequest) Limit(limit int32) BansAPISearchBansRequest {
	r.limit = &limit
	return r
}

// The page number of results to get. Between 1 and 750.
func (r BansAPISearchBansRequest) Page(page int32) BansAPISearchBansRequest {
	r.page = &page
	return r
}

// Search for a specific id.
func (r BansAPISearchBansRequest) SearchId(searchId int32) BansAPISearchBansRequest {
	r.searchId = &searchId
	return r
}

// The order of the results.
func (r BansAPISearchBansRequest) SearchOrder(searchOrder string) BansAPISearchBansRequest {
	r.searchOrder = &searchOrder
	return r
}

// The ID of the banner.
func (r BansAPISearchBansRequest) SearchBannerId(searchBannerId string) BansAPISearchBansRequest {
	r.searchBannerId = &searchBannerId
	return r
}

// The name of banner.
func (r BansAPISearchBansRequest) SearchBannerName(searchBannerName string) BansAPISearchBansRequest {
	r.searchBannerName = &searchBannerName
	return r
}

// The ID of the banned user.
func (r BansAPISearchBansRequest) SearchUserId(searchUserId string) BansAPISearchBansRequest {
	r.searchUserId = &searchUserId
	return r
}

// The name of the banned user.
func (r BansAPISearchBansRequest) SearchUserName(searchUserName string) BansAPISearchBansRequest {
	r.searchUserName = &searchUserName
	return r
}

// The reason of the ban.
func (r BansAPISearchBansRequest) SearchReasonMatches(searchReasonMatches string) BansAPISearchBansRequest {
	r.searchReasonMatches = &searchReasonMatches
	return r
}

// If the ban is expired.
func (r BansAPISearchBansRequest) SearchExpired(searchExpired string) BansAPISearchBansRequest {
	r.searchExpired = &searchExpired
	return r
}

func (r BansAPISearchBansRequest) Execute() (*SearchBans200Response, *http.Response, error) {
	return r.ApiService.SearchBansExecute(r)
}

/*
SearchBans Search Bans

When no results are found, an object with an `bans` key is returned.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return BansAPISearchBansRequest
*/
func (a *BansAPIService) SearchBans(ctx context.Context) BansAPISearchBansRequest {
	return BansAPISearchBansRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SearchBans200Response
func (a *BansAPIService) SearchBansExecute(r BansAPISearchBansRequest) (*SearchBans200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SearchBans200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BansAPIService.SearchBans")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bans.json"

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
	if r.searchBannerId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[banner_id]", r.searchBannerId, "")
	}
	if r.searchBannerName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[banner_name]", r.searchBannerName, "")
	}
	if r.searchUserId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[user_id]", r.searchUserId, "")
	}
	if r.searchUserName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[user_name]", r.searchUserName, "")
	}
	if r.searchReasonMatches != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[reason_matches]", r.searchReasonMatches, "")
	}
	if r.searchExpired != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[expired]", r.searchExpired, "")
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