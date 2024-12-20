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

// WikiPageVersionsAPIService WikiPageVersionsAPI service
type WikiPageVersionsAPIService service

type WikiPageVersionsAPIGetWikiPageVersionRequest struct {
	ctx        context.Context
	ApiService *WikiPageVersionsAPIService
	id         int32
}

func (r WikiPageVersionsAPIGetWikiPageVersionRequest) Execute() (*WikiPageVersion, *http.Response, error) {
	return r.ApiService.GetWikiPageVersionExecute(r)
}

/*
GetWikiPageVersion Get Wiki Page Version

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the wiki page version.
	@return WikiPageVersionsAPIGetWikiPageVersionRequest
*/
func (a *WikiPageVersionsAPIService) GetWikiPageVersion(ctx context.Context, id int32) WikiPageVersionsAPIGetWikiPageVersionRequest {
	return WikiPageVersionsAPIGetWikiPageVersionRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return WikiPageVersion
func (a *WikiPageVersionsAPIService) GetWikiPageVersionExecute(r WikiPageVersionsAPIGetWikiPageVersionRequest) (*WikiPageVersion, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *WikiPageVersion
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WikiPageVersionsAPIService.GetWikiPageVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wiki_page_versions/{id}.json"
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

type WikiPageVersionsAPISearchWikiPageVersionsRequest struct {
	ctx               context.Context
	ApiService        *WikiPageVersionsAPIService
	limit             *int32
	page              *int32
	searchId          *int32
	searchIpAddr      *string
	searchOrder       *string
	searchUpdaterId   *int32
	searchUpdaterName *string
	searchWikiPageId  *int32
	searchTitle       *string
	searchBody        *string
	searchIsLocked    *bool
	searchIsDeleted   *bool
}

// The maximum number of results to return. Between 0 and 320.
func (r WikiPageVersionsAPISearchWikiPageVersionsRequest) Limit(limit int32) WikiPageVersionsAPISearchWikiPageVersionsRequest {
	r.limit = &limit
	return r
}

// The page number of results to get. Between 1 and 750.
func (r WikiPageVersionsAPISearchWikiPageVersionsRequest) Page(page int32) WikiPageVersionsAPISearchWikiPageVersionsRequest {
	r.page = &page
	return r
}

// Search for a specific id.
func (r WikiPageVersionsAPISearchWikiPageVersionsRequest) SearchId(searchId int32) WikiPageVersionsAPISearchWikiPageVersionsRequest {
	r.searchId = &searchId
	return r
}

// Must be Admin+ to use. See [postgres&#39; documentation](https://www.postgresql.org/docs/9.3/functions-net.html) for information on how this is parsed. Specifically, \&quot;is contained within or equals\&quot; (&#x60;&lt;&lt;&#x3D;&#x60;).
func (r WikiPageVersionsAPISearchWikiPageVersionsRequest) SearchIpAddr(searchIpAddr string) WikiPageVersionsAPISearchWikiPageVersionsRequest {
	r.searchIpAddr = &searchIpAddr
	return r
}

func (r WikiPageVersionsAPISearchWikiPageVersionsRequest) SearchOrder(searchOrder string) WikiPageVersionsAPISearchWikiPageVersionsRequest {
	r.searchOrder = &searchOrder
	return r
}

func (r WikiPageVersionsAPISearchWikiPageVersionsRequest) SearchUpdaterId(searchUpdaterId int32) WikiPageVersionsAPISearchWikiPageVersionsRequest {
	r.searchUpdaterId = &searchUpdaterId
	return r
}

func (r WikiPageVersionsAPISearchWikiPageVersionsRequest) SearchUpdaterName(searchUpdaterName string) WikiPageVersionsAPISearchWikiPageVersionsRequest {
	r.searchUpdaterName = &searchUpdaterName
	return r
}

func (r WikiPageVersionsAPISearchWikiPageVersionsRequest) SearchWikiPageId(searchWikiPageId int32) WikiPageVersionsAPISearchWikiPageVersionsRequest {
	r.searchWikiPageId = &searchWikiPageId
	return r
}

func (r WikiPageVersionsAPISearchWikiPageVersionsRequest) SearchTitle(searchTitle string) WikiPageVersionsAPISearchWikiPageVersionsRequest {
	r.searchTitle = &searchTitle
	return r
}

func (r WikiPageVersionsAPISearchWikiPageVersionsRequest) SearchBody(searchBody string) WikiPageVersionsAPISearchWikiPageVersionsRequest {
	r.searchBody = &searchBody
	return r
}

func (r WikiPageVersionsAPISearchWikiPageVersionsRequest) SearchIsLocked(searchIsLocked bool) WikiPageVersionsAPISearchWikiPageVersionsRequest {
	r.searchIsLocked = &searchIsLocked
	return r
}

func (r WikiPageVersionsAPISearchWikiPageVersionsRequest) SearchIsDeleted(searchIsDeleted bool) WikiPageVersionsAPISearchWikiPageVersionsRequest {
	r.searchIsDeleted = &searchIsDeleted
	return r
}

func (r WikiPageVersionsAPISearchWikiPageVersionsRequest) Execute() (*SearchWikiPageVersions200Response, *http.Response, error) {
	return r.ApiService.SearchWikiPageVersionsExecute(r)
}

/*
SearchWikiPageVersions Search Wiki Page Versions

When no results are found, an object with an `wiki_page_versions` key is returned.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return WikiPageVersionsAPISearchWikiPageVersionsRequest
*/
func (a *WikiPageVersionsAPIService) SearchWikiPageVersions(ctx context.Context) WikiPageVersionsAPISearchWikiPageVersionsRequest {
	return WikiPageVersionsAPISearchWikiPageVersionsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SearchWikiPageVersions200Response
func (a *WikiPageVersionsAPIService) SearchWikiPageVersionsExecute(r WikiPageVersionsAPISearchWikiPageVersionsRequest) (*SearchWikiPageVersions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SearchWikiPageVersions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WikiPageVersionsAPIService.SearchWikiPageVersions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/wiki_page_versions.json"

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
	if r.searchWikiPageId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[wiki_page_id]", r.searchWikiPageId, "")
	}
	if r.searchTitle != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[title]", r.searchTitle, "")
	}
	if r.searchBody != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[body]", r.searchBody, "")
	}
	if r.searchIsLocked != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[is_locked]", r.searchIsLocked, "")
	}
	if r.searchIsDeleted != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[is_deleted]", r.searchIsDeleted, "")
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
