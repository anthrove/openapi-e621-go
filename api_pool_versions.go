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


// PoolVersionsAPIService PoolVersionsAPI service
type PoolVersionsAPIService service

type ApiSearchPoolVersionsRequest struct {
	ctx context.Context
	ApiService *PoolVersionsAPIService
	limit *float32
	page *float32
	searchId *float32
	searchIpAddr *string
	searchOrder *string
	searchUpdaterId *float32
	searchUpdaterName *string
	searchPoolId *float32
}

// The maximum number of results to return. Between 0 and 320.
func (r ApiSearchPoolVersionsRequest) Limit(limit float32) ApiSearchPoolVersionsRequest {
	r.limit = &limit
	return r
}

// The page number of results to get. Between 1 and 750.
func (r ApiSearchPoolVersionsRequest) Page(page float32) ApiSearchPoolVersionsRequest {
	r.page = &page
	return r
}

// Search for a specific id.
func (r ApiSearchPoolVersionsRequest) SearchId(searchId float32) ApiSearchPoolVersionsRequest {
	r.searchId = &searchId
	return r
}

// Must be Admin+ to use. See [postgres&#39; documentation](https://www.postgresql.org/docs/9.3/functions-net.html) for information on how this is parsed. Specifically, \&quot;is contained within or equals\&quot; (&#x60;&lt;&lt;&#x3D;&#x60;).
func (r ApiSearchPoolVersionsRequest) SearchIpAddr(searchIpAddr string) ApiSearchPoolVersionsRequest {
	r.searchIpAddr = &searchIpAddr
	return r
}

// The order of the results.
func (r ApiSearchPoolVersionsRequest) SearchOrder(searchOrder string) ApiSearchPoolVersionsRequest {
	r.searchOrder = &searchOrder
	return r
}

func (r ApiSearchPoolVersionsRequest) SearchUpdaterId(searchUpdaterId float32) ApiSearchPoolVersionsRequest {
	r.searchUpdaterId = &searchUpdaterId
	return r
}

func (r ApiSearchPoolVersionsRequest) SearchUpdaterName(searchUpdaterName string) ApiSearchPoolVersionsRequest {
	r.searchUpdaterName = &searchUpdaterName
	return r
}

func (r ApiSearchPoolVersionsRequest) SearchPoolId(searchPoolId float32) ApiSearchPoolVersionsRequest {
	r.searchPoolId = &searchPoolId
	return r
}

func (r ApiSearchPoolVersionsRequest) Execute() (*SearchPoolVersions200Response, *http.Response, error) {
	return r.ApiService.SearchPoolVersionsExecute(r)
}

/*
SearchPoolVersions Search Pool Versions

When no results are found, an object with an `pool_versions` key is returned.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchPoolVersionsRequest
*/
func (a *PoolVersionsAPIService) SearchPoolVersions(ctx context.Context) ApiSearchPoolVersionsRequest {
	return ApiSearchPoolVersionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SearchPoolVersions200Response
func (a *PoolVersionsAPIService) SearchPoolVersionsExecute(r ApiSearchPoolVersionsRequest) (*SearchPoolVersions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SearchPoolVersions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoolVersionsAPIService.SearchPoolVersions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pool_versions.json"

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
	if r.searchIpAddr != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[ip_addr]", r.searchIpAddr, "")
	}
	if r.searchOrder != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[order]", r.searchOrder, "")
	}
	if r.searchUpdaterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[updater_id]", r.searchUpdaterId, "")
	}
	if r.searchUpdaterName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[updater_name]", r.searchUpdaterName, "")
	}
	if r.searchPoolId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[pool_id]", r.searchPoolId, "")
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
