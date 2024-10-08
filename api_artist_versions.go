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


// ArtistVersionsAPIService ArtistVersionsAPI service
type ArtistVersionsAPIService service

type ApiSearchArtistVersionsRequest struct {
	ctx context.Context
	ApiService *ArtistVersionsAPIService
	limit *float32
	page *float32
	searchId *float32
	searchIpAddr *string
	searchOrder *string
	searchName *string
	searchArtistId *string
	searchUpdaterName *string
	searchUpdaterId *string
}

// The maximum number of results to return. Between 0 and 320.
func (r ApiSearchArtistVersionsRequest) Limit(limit float32) ApiSearchArtistVersionsRequest {
	r.limit = &limit
	return r
}

// The page number of results to get. Between 1 and 750.
func (r ApiSearchArtistVersionsRequest) Page(page float32) ApiSearchArtistVersionsRequest {
	r.page = &page
	return r
}

// Search for a specific id.
func (r ApiSearchArtistVersionsRequest) SearchId(searchId float32) ApiSearchArtistVersionsRequest {
	r.searchId = &searchId
	return r
}

// Must be Admin+ to use. See [postgres&#39; documentation](https://www.postgresql.org/docs/9.3/functions-net.html) for information on how this is parsed. Specifically, \&quot;is contained within or equals\&quot; (&#x60;&lt;&lt;&#x3D;&#x60;).
func (r ApiSearchArtistVersionsRequest) SearchIpAddr(searchIpAddr string) ApiSearchArtistVersionsRequest {
	r.searchIpAddr = &searchIpAddr
	return r
}

// The order of the results.
func (r ApiSearchArtistVersionsRequest) SearchOrder(searchOrder string) ApiSearchArtistVersionsRequest {
	r.searchOrder = &searchOrder
	return r
}

// The name of the artist.
func (r ApiSearchArtistVersionsRequest) SearchName(searchName string) ApiSearchArtistVersionsRequest {
	r.searchName = &searchName
	return r
}

// The id of the artist.
func (r ApiSearchArtistVersionsRequest) SearchArtistId(searchArtistId string) ApiSearchArtistVersionsRequest {
	r.searchArtistId = &searchArtistId
	return r
}

// The name of the user that updated the artist.
func (r ApiSearchArtistVersionsRequest) SearchUpdaterName(searchUpdaterName string) ApiSearchArtistVersionsRequest {
	r.searchUpdaterName = &searchUpdaterName
	return r
}

// The id of the user that updated the artist.
func (r ApiSearchArtistVersionsRequest) SearchUpdaterId(searchUpdaterId string) ApiSearchArtistVersionsRequest {
	r.searchUpdaterId = &searchUpdaterId
	return r
}

func (r ApiSearchArtistVersionsRequest) Execute() (*SearchArtistVersions200Response, *http.Response, error) {
	return r.ApiService.SearchArtistVersionsExecute(r)
}

/*
SearchArtistVersions Search Artist Versions

When no results are found, an object with an `artist_versions` key is returned.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchArtistVersionsRequest
*/
func (a *ArtistVersionsAPIService) SearchArtistVersions(ctx context.Context) ApiSearchArtistVersionsRequest {
	return ApiSearchArtistVersionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SearchArtistVersions200Response
func (a *ArtistVersionsAPIService) SearchArtistVersionsExecute(r ApiSearchArtistVersionsRequest) (*SearchArtistVersions200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SearchArtistVersions200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ArtistVersionsAPIService.SearchArtistVersions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/artist_versions.json"

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
	if r.searchName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[name]", r.searchName, "")
	}
	if r.searchArtistId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[artist_id]", r.searchArtistId, "")
	}
	if r.searchUpdaterName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[updater_name]", r.searchUpdaterName, "")
	}
	if r.searchUpdaterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[updater_id]", r.searchUpdaterId, "")
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
