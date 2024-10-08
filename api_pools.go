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


// PoolsAPIService PoolsAPI service
type PoolsAPIService service

type ApiAddPostToPoolRequest struct {
	ctx context.Context
	ApiService *PoolsAPIService
	postId *float32
	poolId *float32
	poolName *string
}

func (r ApiAddPostToPoolRequest) PostId(postId float32) ApiAddPostToPoolRequest {
	r.postId = &postId
	return r
}

// Mutually exclusive with pool_name.
func (r ApiAddPostToPoolRequest) PoolId(poolId float32) ApiAddPostToPoolRequest {
	r.poolId = &poolId
	return r
}

// Mutually exclusive with pool_id.
func (r ApiAddPostToPoolRequest) PoolName(poolName string) ApiAddPostToPoolRequest {
	r.poolName = &poolName
	return r
}

func (r ApiAddPostToPoolRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddPostToPoolExecute(r)
}

/*
AddPostToPool Add Post To Pool

Note that the extension is JS, this route will not return JSON.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAddPostToPoolRequest
*/
func (a *PoolsAPIService) AddPostToPool(ctx context.Context) ApiAddPostToPoolRequest {
	return ApiAddPostToPoolRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PoolsAPIService) AddPostToPoolExecute(r ApiAddPostToPoolRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoolsAPIService.AddPostToPool")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pool_element.js"

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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.poolId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "pool_id", r.poolId, "")
	}
	if r.poolName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "pool_name", r.poolName, "")
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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiCreatePoolRequest struct {
	ctx context.Context
	ApiService *PoolsAPIService
	poolName *string
	poolDescription *string
	poolCategory *PoolCategories
	ipoolSActive *bool
	poolPostIdsString *string
	poolPostIds *[]float32
}

func (r ApiCreatePoolRequest) PoolName(poolName string) ApiCreatePoolRequest {
	r.poolName = &poolName
	return r
}

func (r ApiCreatePoolRequest) PoolDescription(poolDescription string) ApiCreatePoolRequest {
	r.poolDescription = &poolDescription
	return r
}

func (r ApiCreatePoolRequest) PoolCategory(poolCategory PoolCategories) ApiCreatePoolRequest {
	r.poolCategory = &poolCategory
	return r
}

func (r ApiCreatePoolRequest) IpoolSActive(ipoolSActive bool) ApiCreatePoolRequest {
	r.ipoolSActive = &ipoolSActive
	return r
}

// Space separated list of post IDs. Mutually exclusive with post_ids.
func (r ApiCreatePoolRequest) PoolPostIdsString(poolPostIdsString string) ApiCreatePoolRequest {
	r.poolPostIdsString = &poolPostIdsString
	return r
}

// Array of post IDs. Mutually exclusive with post_ids_string.
func (r ApiCreatePoolRequest) PoolPostIds(poolPostIds []float32) ApiCreatePoolRequest {
	r.poolPostIds = &poolPostIds
	return r
}

func (r ApiCreatePoolRequest) Execute() (*Pool, *http.Response, error) {
	return r.ApiService.CreatePoolExecute(r)
}

/*
CreatePool Create Pool

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreatePoolRequest
*/
func (a *PoolsAPIService) CreatePool(ctx context.Context) ApiCreatePoolRequest {
	return ApiCreatePoolRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Pool
func (a *PoolsAPIService) CreatePoolExecute(r ApiCreatePoolRequest) (*Pool, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Pool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoolsAPIService.CreatePool")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pools.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.poolName == nil {
		return localVarReturnValue, nil, reportError("poolName is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "pool[name]", r.poolName, "")
	if r.poolDescription != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "pool[description]", r.poolDescription, "")
	}
	if r.poolCategory != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "pool[category]", r.poolCategory, "")
	}
	if r.ipoolSActive != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "ipool[s_active]", r.ipoolSActive, "")
	}
	if r.poolPostIdsString != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "pool[post_ids_string]", r.poolPostIdsString, "")
	}
	if r.poolPostIds != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "pool[post_ids]", r.poolPostIds, "csv")
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

type ApiDeletePoolRequest struct {
	ctx context.Context
	ApiService *PoolsAPIService
	id float32
}

func (r ApiDeletePoolRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeletePoolExecute(r)
}

/*
DeletePool Delete Pool

You must be Janitor+.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the pool.
 @return ApiDeletePoolRequest
*/
func (a *PoolsAPIService) DeletePool(ctx context.Context, id float32) ApiDeletePoolRequest {
	return ApiDeletePoolRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *PoolsAPIService) DeletePoolExecute(r ApiDeletePoolRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoolsAPIService.DeletePool")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pools/{id}.json"
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

type ApiEditPoolRequest struct {
	ctx context.Context
	ApiService *PoolsAPIService
	id float32
	poolName *string
	poolDescription *string
	poolIsActive *bool
	poolCategory *PoolCategories
	poolPostIdsString *string
	poolPostIds *[]float32
}

func (r ApiEditPoolRequest) PoolName(poolName string) ApiEditPoolRequest {
	r.poolName = &poolName
	return r
}

func (r ApiEditPoolRequest) PoolDescription(poolDescription string) ApiEditPoolRequest {
	r.poolDescription = &poolDescription
	return r
}

func (r ApiEditPoolRequest) PoolIsActive(poolIsActive bool) ApiEditPoolRequest {
	r.poolIsActive = &poolIsActive
	return r
}

// If the pool has more than 30 posts, you must be Janitor+.
func (r ApiEditPoolRequest) PoolCategory(poolCategory PoolCategories) ApiEditPoolRequest {
	r.poolCategory = &poolCategory
	return r
}

// Space separated list of post IDs. Mutually exclusive with post_ids.
func (r ApiEditPoolRequest) PoolPostIdsString(poolPostIdsString string) ApiEditPoolRequest {
	r.poolPostIdsString = &poolPostIdsString
	return r
}

// Array of post IDs. Mutually exclusive with post_ids_string.
func (r ApiEditPoolRequest) PoolPostIds(poolPostIds []float32) ApiEditPoolRequest {
	r.poolPostIds = &poolPostIds
	return r
}

func (r ApiEditPoolRequest) Execute() (*http.Response, error) {
	return r.ApiService.EditPoolExecute(r)
}

/*
EditPool Edit Pool

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the pool.
 @return ApiEditPoolRequest
*/
func (a *PoolsAPIService) EditPool(ctx context.Context, id float32) ApiEditPoolRequest {
	return ApiEditPoolRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *PoolsAPIService) EditPoolExecute(r ApiEditPoolRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoolsAPIService.EditPool")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pools/{id}.json"
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
	if r.poolName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "pool[name]", r.poolName, "")
	}
	if r.poolDescription != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "pool[description]", r.poolDescription, "")
	}
	if r.poolIsActive != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "pool[is_active]", r.poolIsActive, "")
	}
	if r.poolCategory != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "pool[category]", r.poolCategory, "")
	}
	if r.poolPostIdsString != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "pool[post_ids_string]", r.poolPostIdsString, "")
	}
	if r.poolPostIds != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "pool[post_ids]", r.poolPostIds, "csv")
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

type ApiGetPoolRequest struct {
	ctx context.Context
	ApiService *PoolsAPIService
	id float32
}

func (r ApiGetPoolRequest) Execute() (*Pool, *http.Response, error) {
	return r.ApiService.GetPoolExecute(r)
}

/*
GetPool Get Pool

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the pool.
 @return ApiGetPoolRequest
*/
func (a *PoolsAPIService) GetPool(ctx context.Context, id float32) ApiGetPoolRequest {
	return ApiGetPoolRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Pool
func (a *PoolsAPIService) GetPoolExecute(r ApiGetPoolRequest) (*Pool, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Pool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoolsAPIService.GetPool")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pools/{id}.json"
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

type ApiRemovePostFromPoolRequest struct {
	ctx context.Context
	ApiService *PoolsAPIService
	poolId *float32
	postId *float32
}

func (r ApiRemovePostFromPoolRequest) PoolId(poolId float32) ApiRemovePostFromPoolRequest {
	r.poolId = &poolId
	return r
}

func (r ApiRemovePostFromPoolRequest) PostId(postId float32) ApiRemovePostFromPoolRequest {
	r.postId = &postId
	return r
}

func (r ApiRemovePostFromPoolRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemovePostFromPoolExecute(r)
}

/*
RemovePostFromPool Remove Post From Pool

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRemovePostFromPoolRequest
*/
func (a *PoolsAPIService) RemovePostFromPool(ctx context.Context) ApiRemovePostFromPoolRequest {
	return ApiRemovePostFromPoolRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *PoolsAPIService) RemovePostFromPoolExecute(r ApiRemovePostFromPoolRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoolsAPIService.RemovePostFromPool")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pool_element.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.poolId == nil {
		return nil, reportError("poolId is required and must be specified")
	}
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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "pool_id", r.poolId, "")
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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiRevertPoolRequest struct {
	ctx context.Context
	ApiService *PoolsAPIService
	id float32
	versionId *float32
}

// The version ID to revert to.
func (r ApiRevertPoolRequest) VersionId(versionId float32) ApiRevertPoolRequest {
	r.versionId = &versionId
	return r
}

func (r ApiRevertPoolRequest) Execute() (*http.Response, error) {
	return r.ApiService.RevertPoolExecute(r)
}

/*
RevertPool Revert Pool

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the pool.
 @return ApiRevertPoolRequest
*/
func (a *PoolsAPIService) RevertPool(ctx context.Context, id float32) ApiRevertPoolRequest {
	return ApiRevertPoolRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *PoolsAPIService) RevertPoolExecute(r ApiRevertPoolRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoolsAPIService.RevertPool")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pools/{id}/revert.json"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.versionId == nil {
		return nil, reportError("versionId is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "version_id", r.versionId, "")
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

type ApiSearchPoolsRequest struct {
	ctx context.Context
	ApiService *PoolsAPIService
	limit *float32
	page *float32
	searchId *float32
	searchOrder *string
	searchNameMatches *string
	searchDescriptionMatches *string
	searchCreatorId *float32
	searchCreatorName *string
	searchCategory *PoolCategories
	searchIsActive *bool
}

// The maximum number of results to return. Between 0 and 320.
func (r ApiSearchPoolsRequest) Limit(limit float32) ApiSearchPoolsRequest {
	r.limit = &limit
	return r
}

// The page number of results to get. Between 1 and 750.
func (r ApiSearchPoolsRequest) Page(page float32) ApiSearchPoolsRequest {
	r.page = &page
	return r
}

// Search for a specific id.
func (r ApiSearchPoolsRequest) SearchId(searchId float32) ApiSearchPoolsRequest {
	r.searchId = &searchId
	return r
}

func (r ApiSearchPoolsRequest) SearchOrder(searchOrder string) ApiSearchPoolsRequest {
	r.searchOrder = &searchOrder
	return r
}

func (r ApiSearchPoolsRequest) SearchNameMatches(searchNameMatches string) ApiSearchPoolsRequest {
	r.searchNameMatches = &searchNameMatches
	return r
}

func (r ApiSearchPoolsRequest) SearchDescriptionMatches(searchDescriptionMatches string) ApiSearchPoolsRequest {
	r.searchDescriptionMatches = &searchDescriptionMatches
	return r
}

func (r ApiSearchPoolsRequest) SearchCreatorId(searchCreatorId float32) ApiSearchPoolsRequest {
	r.searchCreatorId = &searchCreatorId
	return r
}

func (r ApiSearchPoolsRequest) SearchCreatorName(searchCreatorName string) ApiSearchPoolsRequest {
	r.searchCreatorName = &searchCreatorName
	return r
}

func (r ApiSearchPoolsRequest) SearchCategory(searchCategory PoolCategories) ApiSearchPoolsRequest {
	r.searchCategory = &searchCategory
	return r
}

func (r ApiSearchPoolsRequest) SearchIsActive(searchIsActive bool) ApiSearchPoolsRequest {
	r.searchIsActive = &searchIsActive
	return r
}

func (r ApiSearchPoolsRequest) Execute() ([]Pool, *http.Response, error) {
	return r.ApiService.SearchPoolsExecute(r)
}

/*
SearchPools Search Pools

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchPoolsRequest
*/
func (a *PoolsAPIService) SearchPools(ctx context.Context) ApiSearchPoolsRequest {
	return ApiSearchPoolsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Pool
func (a *PoolsAPIService) SearchPoolsExecute(r ApiSearchPoolsRequest) ([]Pool, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Pool
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PoolsAPIService.SearchPools")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/pools.json"

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
	if r.searchDescriptionMatches != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[description_matches]", r.searchDescriptionMatches, "")
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
	if r.searchIsActive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[is_active]", r.searchIsActive, "")
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
