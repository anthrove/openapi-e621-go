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
	"os"
	"strings"
)

// MascotsAPIService MascotsAPI service
type MascotsAPIService service

type MascotsAPICreateMascotRequest struct {
	ctx                     context.Context
	ApiService              *MascotsAPIService
	mascotMascotFile        *os.File
	mascotDisplayName       *string
	mascotBackgroundColor   *string
	mascotArtistUrl         *string
	mascotArtistName        *string
	mascotAvailableOnString *string
	mascotActive            *bool
	mascotHideAnonymous     *bool
}

func (r MascotsAPICreateMascotRequest) MascotMascotFile(mascotMascotFile *os.File) MascotsAPICreateMascotRequest {
	r.mascotMascotFile = mascotMascotFile
	return r
}

func (r MascotsAPICreateMascotRequest) MascotDisplayName(mascotDisplayName string) MascotsAPICreateMascotRequest {
	r.mascotDisplayName = &mascotDisplayName
	return r
}

func (r MascotsAPICreateMascotRequest) MascotBackgroundColor(mascotBackgroundColor string) MascotsAPICreateMascotRequest {
	r.mascotBackgroundColor = &mascotBackgroundColor
	return r
}

func (r MascotsAPICreateMascotRequest) MascotArtistUrl(mascotArtistUrl string) MascotsAPICreateMascotRequest {
	r.mascotArtistUrl = &mascotArtistUrl
	return r
}

func (r MascotsAPICreateMascotRequest) MascotArtistName(mascotArtistName string) MascotsAPICreateMascotRequest {
	r.mascotArtistName = &mascotArtistName
	return r
}

// Comma separated site names.
func (r MascotsAPICreateMascotRequest) MascotAvailableOnString(mascotAvailableOnString string) MascotsAPICreateMascotRequest {
	r.mascotAvailableOnString = &mascotAvailableOnString
	return r
}

func (r MascotsAPICreateMascotRequest) MascotActive(mascotActive bool) MascotsAPICreateMascotRequest {
	r.mascotActive = &mascotActive
	return r
}

func (r MascotsAPICreateMascotRequest) MascotHideAnonymous(mascotHideAnonymous bool) MascotsAPICreateMascotRequest {
	r.mascotHideAnonymous = &mascotHideAnonymous
	return r
}

func (r MascotsAPICreateMascotRequest) Execute() (*Mascot, *http.Response, error) {
	return r.ApiService.CreateMascotExecute(r)
}

/*
CreateMascot Create Mascot

You must be Admin+.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return MascotsAPICreateMascotRequest
*/
func (a *MascotsAPIService) CreateMascot(ctx context.Context) MascotsAPICreateMascotRequest {
	return MascotsAPICreateMascotRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return Mascot
func (a *MascotsAPIService) CreateMascotExecute(r MascotsAPICreateMascotRequest) (*Mascot, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Mascot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MascotsAPIService.CreateMascot")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mascots.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.mascotMascotFile == nil {
		return localVarReturnValue, nil, reportError("mascotMascotFile is required and must be specified")
	}
	if r.mascotDisplayName == nil {
		return localVarReturnValue, nil, reportError("mascotDisplayName is required and must be specified")
	}
	if r.mascotBackgroundColor == nil {
		return localVarReturnValue, nil, reportError("mascotBackgroundColor is required and must be specified")
	}
	if r.mascotArtistUrl == nil {
		return localVarReturnValue, nil, reportError("mascotArtistUrl is required and must be specified")
	}
	if r.mascotArtistName == nil {
		return localVarReturnValue, nil, reportError("mascotArtistName is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	var mascotMascotFileLocalVarFormFileName string
	var mascotMascotFileLocalVarFileName string
	var mascotMascotFileLocalVarFileBytes []byte

	mascotMascotFileLocalVarFormFileName = "mascot[mascot_file]"
	mascotMascotFileLocalVarFile := r.mascotMascotFile

	if mascotMascotFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(mascotMascotFileLocalVarFile)

		mascotMascotFileLocalVarFileBytes = fbs
		mascotMascotFileLocalVarFileName = mascotMascotFileLocalVarFile.Name()
		mascotMascotFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: mascotMascotFileLocalVarFileBytes, fileName: mascotMascotFileLocalVarFileName, formFileName: mascotMascotFileLocalVarFormFileName})
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "mascot[display_name]", r.mascotDisplayName, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "mascot[background_color]", r.mascotBackgroundColor, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "mascot[artist_url]", r.mascotArtistUrl, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "mascot[artist_name]", r.mascotArtistName, "")
	if r.mascotAvailableOnString != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "mascot[available_on_string]", r.mascotAvailableOnString, "")
	}
	if r.mascotActive != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "mascot[active]", r.mascotActive, "")
	}
	if r.mascotHideAnonymous != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "mascot[hide_anonymous]", r.mascotHideAnonymous, "")
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

type MascotsAPIDeleteMascotRequest struct {
	ctx        context.Context
	ApiService *MascotsAPIService
	id         int32
}

func (r MascotsAPIDeleteMascotRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteMascotExecute(r)
}

/*
DeleteMascot Delete Mascot

You must be Admin+.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the mascot.
	@return MascotsAPIDeleteMascotRequest
*/
func (a *MascotsAPIService) DeleteMascot(ctx context.Context, id int32) MascotsAPIDeleteMascotRequest {
	return MascotsAPIDeleteMascotRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *MascotsAPIService) DeleteMascotExecute(r MascotsAPIDeleteMascotRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MascotsAPIService.DeleteMascot")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mascots/{id}.json"
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

type MascotsAPIEditMascotRequest struct {
	ctx                     context.Context
	ApiService              *MascotsAPIService
	id                      int32
	mascotMascotFile        *os.File
	mascotDisplayName       *string
	mascotBackgroundColor   *string
	mascotArtistUrl         *string
	mascotArtistName        *string
	mascotAvailableOnString *string
	mascotActive            *bool
	mascotHideAnonymous     *bool
}

func (r MascotsAPIEditMascotRequest) MascotMascotFile(mascotMascotFile *os.File) MascotsAPIEditMascotRequest {
	r.mascotMascotFile = mascotMascotFile
	return r
}

func (r MascotsAPIEditMascotRequest) MascotDisplayName(mascotDisplayName string) MascotsAPIEditMascotRequest {
	r.mascotDisplayName = &mascotDisplayName
	return r
}

func (r MascotsAPIEditMascotRequest) MascotBackgroundColor(mascotBackgroundColor string) MascotsAPIEditMascotRequest {
	r.mascotBackgroundColor = &mascotBackgroundColor
	return r
}

func (r MascotsAPIEditMascotRequest) MascotArtistUrl(mascotArtistUrl string) MascotsAPIEditMascotRequest {
	r.mascotArtistUrl = &mascotArtistUrl
	return r
}

func (r MascotsAPIEditMascotRequest) MascotArtistName(mascotArtistName string) MascotsAPIEditMascotRequest {
	r.mascotArtistName = &mascotArtistName
	return r
}

// Comma separated site names.
func (r MascotsAPIEditMascotRequest) MascotAvailableOnString(mascotAvailableOnString string) MascotsAPIEditMascotRequest {
	r.mascotAvailableOnString = &mascotAvailableOnString
	return r
}

func (r MascotsAPIEditMascotRequest) MascotActive(mascotActive bool) MascotsAPIEditMascotRequest {
	r.mascotActive = &mascotActive
	return r
}

func (r MascotsAPIEditMascotRequest) MascotHideAnonymous(mascotHideAnonymous bool) MascotsAPIEditMascotRequest {
	r.mascotHideAnonymous = &mascotHideAnonymous
	return r
}

func (r MascotsAPIEditMascotRequest) Execute() (*http.Response, error) {
	return r.ApiService.EditMascotExecute(r)
}

/*
EditMascot Edit Mascot

You must be Admin+.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the mascot.
	@return MascotsAPIEditMascotRequest
*/
func (a *MascotsAPIService) EditMascot(ctx context.Context, id int32) MascotsAPIEditMascotRequest {
	return MascotsAPIEditMascotRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *MascotsAPIService) EditMascotExecute(r MascotsAPIEditMascotRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MascotsAPIService.EditMascot")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mascots/{id}.json"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	var mascotMascotFileLocalVarFormFileName string
	var mascotMascotFileLocalVarFileName string
	var mascotMascotFileLocalVarFileBytes []byte

	mascotMascotFileLocalVarFormFileName = "mascot[mascot_file]"
	mascotMascotFileLocalVarFile := r.mascotMascotFile

	if mascotMascotFileLocalVarFile != nil {
		fbs, _ := io.ReadAll(mascotMascotFileLocalVarFile)

		mascotMascotFileLocalVarFileBytes = fbs
		mascotMascotFileLocalVarFileName = mascotMascotFileLocalVarFile.Name()
		mascotMascotFileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: mascotMascotFileLocalVarFileBytes, fileName: mascotMascotFileLocalVarFileName, formFileName: mascotMascotFileLocalVarFormFileName})
	}
	if r.mascotDisplayName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "mascot[display_name]", r.mascotDisplayName, "")
	}
	if r.mascotBackgroundColor != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "mascot[background_color]", r.mascotBackgroundColor, "")
	}
	if r.mascotArtistUrl != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "mascot[artist_url]", r.mascotArtistUrl, "")
	}
	if r.mascotArtistName != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "mascot[artist_name]", r.mascotArtistName, "")
	}
	if r.mascotAvailableOnString != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "mascot[available_on_string]", r.mascotAvailableOnString, "")
	}
	if r.mascotActive != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "mascot[active]", r.mascotActive, "")
	}
	if r.mascotHideAnonymous != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "mascot[hide_anonymous]", r.mascotHideAnonymous, "")
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

type MascotsAPISearchMascotsRequest struct {
	ctx        context.Context
	ApiService *MascotsAPIService
	limit      *int32
	page       *int32
	searchId   *int32
}

// The maximum number of results to return. Between 0 and 320.
func (r MascotsAPISearchMascotsRequest) Limit(limit int32) MascotsAPISearchMascotsRequest {
	r.limit = &limit
	return r
}

// The page number of results to get. Between 1 and 750.
func (r MascotsAPISearchMascotsRequest) Page(page int32) MascotsAPISearchMascotsRequest {
	r.page = &page
	return r
}

// Search for a specific id.
func (r MascotsAPISearchMascotsRequest) SearchId(searchId int32) MascotsAPISearchMascotsRequest {
	r.searchId = &searchId
	return r
}

func (r MascotsAPISearchMascotsRequest) Execute() (*SearchMascots200Response, *http.Response, error) {
	return r.ApiService.SearchMascotsExecute(r)
}

/*
SearchMascots Search Mascots

When no results are found, an object with an `mascots` key is returned.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return MascotsAPISearchMascotsRequest
*/
func (a *MascotsAPIService) SearchMascots(ctx context.Context) MascotsAPISearchMascotsRequest {
	return MascotsAPISearchMascotsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SearchMascots200Response
func (a *MascotsAPIService) SearchMascotsExecute(r MascotsAPISearchMascotsRequest) (*SearchMascots200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SearchMascots200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MascotsAPIService.SearchMascots")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mascots.json"

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
