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

// NotesAPIService NotesAPI service
type NotesAPIService service

type NotesAPICreateNoteRequest struct {
	ctx        context.Context
	ApiService *NotesAPIService
	notePostId *float32
	noteX      *float32
	noteY      *float32
	noteWidth  *float32
	noteHeight *float32
	noteBody   *string
	noteHtmlId *string
}

func (r NotesAPICreateNoteRequest) NotePostId(notePostId float32) NotesAPICreateNoteRequest {
	r.notePostId = &notePostId
	return r
}

func (r NotesAPICreateNoteRequest) NoteX(noteX float32) NotesAPICreateNoteRequest {
	r.noteX = &noteX
	return r
}

func (r NotesAPICreateNoteRequest) NoteY(noteY float32) NotesAPICreateNoteRequest {
	r.noteY = &noteY
	return r
}

func (r NotesAPICreateNoteRequest) NoteWidth(noteWidth float32) NotesAPICreateNoteRequest {
	r.noteWidth = &noteWidth
	return r
}

func (r NotesAPICreateNoteRequest) NoteHeight(noteHeight float32) NotesAPICreateNoteRequest {
	r.noteHeight = &noteHeight
	return r
}

func (r NotesAPICreateNoteRequest) NoteBody(noteBody string) NotesAPICreateNoteRequest {
	r.noteBody = &noteBody
	return r
}

// Passthrough, used in frontend.
func (r NotesAPICreateNoteRequest) NoteHtmlId(noteHtmlId string) NotesAPICreateNoteRequest {
	r.noteHtmlId = &noteHtmlId
	return r
}

func (r NotesAPICreateNoteRequest) Execute() (*CreateNote201Response, *http.Response, error) {
	return r.ApiService.CreateNoteExecute(r)
}

/*
CreateNote Create Note

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return NotesAPICreateNoteRequest
*/
func (a *NotesAPIService) CreateNote(ctx context.Context) NotesAPICreateNoteRequest {
	return NotesAPICreateNoteRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateNote201Response
func (a *NotesAPIService) CreateNoteExecute(r NotesAPICreateNoteRequest) (*CreateNote201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateNote201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotesAPIService.CreateNote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notes.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.notePostId == nil {
		return localVarReturnValue, nil, reportError("notePostId is required and must be specified")
	}
	if r.noteX == nil {
		return localVarReturnValue, nil, reportError("noteX is required and must be specified")
	}
	if r.noteY == nil {
		return localVarReturnValue, nil, reportError("noteY is required and must be specified")
	}
	if r.noteWidth == nil {
		return localVarReturnValue, nil, reportError("noteWidth is required and must be specified")
	}
	if r.noteHeight == nil {
		return localVarReturnValue, nil, reportError("noteHeight is required and must be specified")
	}
	if r.noteBody == nil {
		return localVarReturnValue, nil, reportError("noteBody is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "note[post_id]", r.notePostId, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "note[x]", r.noteX, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "note[y]", r.noteY, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "note[width]", r.noteWidth, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "note[height]", r.noteHeight, "")
	parameterAddToHeaderOrQuery(localVarFormParams, "note[body]", r.noteBody, "")
	if r.noteHtmlId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "note[html_id]", r.noteHtmlId, "")
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

type NotesAPIDeleteNoteRequest struct {
	ctx        context.Context
	ApiService *NotesAPIService
	id         float32
}

func (r NotesAPIDeleteNoteRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteNoteExecute(r)
}

/*
DeleteNote Delete Note

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the note.
	@return NotesAPIDeleteNoteRequest
*/
func (a *NotesAPIService) DeleteNote(ctx context.Context, id float32) NotesAPIDeleteNoteRequest {
	return NotesAPIDeleteNoteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *NotesAPIService) DeleteNoteExecute(r NotesAPIDeleteNoteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotesAPIService.DeleteNote")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notes/{id}.json"
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

type NotesAPIEditNoteRequest struct {
	ctx        context.Context
	ApiService *NotesAPIService
	id         float32
	noteX      *float32
	noteY      *float32
	noteWidth  *float32
	noteHeight *float32
	noteBody   *string
}

func (r NotesAPIEditNoteRequest) NoteX(noteX float32) NotesAPIEditNoteRequest {
	r.noteX = &noteX
	return r
}

func (r NotesAPIEditNoteRequest) NoteY(noteY float32) NotesAPIEditNoteRequest {
	r.noteY = &noteY
	return r
}

func (r NotesAPIEditNoteRequest) NoteWidth(noteWidth float32) NotesAPIEditNoteRequest {
	r.noteWidth = &noteWidth
	return r
}

func (r NotesAPIEditNoteRequest) NoteHeight(noteHeight float32) NotesAPIEditNoteRequest {
	r.noteHeight = &noteHeight
	return r
}

func (r NotesAPIEditNoteRequest) NoteBody(noteBody string) NotesAPIEditNoteRequest {
	r.noteBody = &noteBody
	return r
}

func (r NotesAPIEditNoteRequest) Execute() (*http.Response, error) {
	return r.ApiService.EditNoteExecute(r)
}

/*
EditNote Edit Note

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the note.
	@return NotesAPIEditNoteRequest
*/
func (a *NotesAPIService) EditNote(ctx context.Context, id float32) NotesAPIEditNoteRequest {
	return NotesAPIEditNoteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *NotesAPIService) EditNoteExecute(r NotesAPIEditNoteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotesAPIService.EditNote")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notes/{id}.json"
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
	if r.noteX != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "note[x]", r.noteX, "")
	}
	if r.noteY != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "note[y]", r.noteY, "")
	}
	if r.noteWidth != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "note[width]", r.noteWidth, "")
	}
	if r.noteHeight != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "note[height]", r.noteHeight, "")
	}
	if r.noteBody != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "note[body]", r.noteBody, "")
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

type NotesAPIGetNoteRequest struct {
	ctx        context.Context
	ApiService *NotesAPIService
	id         float32
}

func (r NotesAPIGetNoteRequest) Execute() (*Note, *http.Response, error) {
	return r.ApiService.GetNoteExecute(r)
}

/*
GetNote Get Note

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the note.
	@return NotesAPIGetNoteRequest
*/
func (a *NotesAPIService) GetNote(ctx context.Context, id float32) NotesAPIGetNoteRequest {
	return NotesAPIGetNoteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return Note
func (a *NotesAPIService) GetNoteExecute(r NotesAPIGetNoteRequest) (*Note, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Note
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotesAPIService.GetNote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notes/{id}.json"
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

type NotesAPIRevertNoteRequest struct {
	ctx        context.Context
	ApiService *NotesAPIService
	id         float32
	versionId  *float32
}

// The version ID to revert to.
func (r NotesAPIRevertNoteRequest) VersionId(versionId float32) NotesAPIRevertNoteRequest {
	r.versionId = &versionId
	return r
}

func (r NotesAPIRevertNoteRequest) Execute() (*http.Response, error) {
	return r.ApiService.RevertNoteExecute(r)
}

/*
RevertNote Revert Note

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the note.
	@return NotesAPIRevertNoteRequest
*/
func (a *NotesAPIService) RevertNote(ctx context.Context, id float32) NotesAPIRevertNoteRequest {
	return NotesAPIRevertNoteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *NotesAPIService) RevertNoteExecute(r NotesAPIRevertNoteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotesAPIService.RevertNote")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notes/{id}/revert.json"
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

type NotesAPISearchNotesRequest struct {
	ctx                       context.Context
	ApiService                *NotesAPIService
	limit                     *int32
	page                      *int32
	searchId                  *int32
	searchOrder               *string
	searchBodyMatches         *string
	searchIsActive            *bool
	searchPostId              *float32
	searchPostTagsMatch       *string
	searchPostNoteUpdaterId   *float32
	searchPostNoteUpdaterName *string
	searchCreatorId           *float32
	searchCreatorName         *string
}

// The maximum number of results to return. Between 0 and 320.
func (r NotesAPISearchNotesRequest) Limit(limit int32) NotesAPISearchNotesRequest {
	r.limit = &limit
	return r
}

// The page number of results to get. Between 1 and 750.
func (r NotesAPISearchNotesRequest) Page(page int32) NotesAPISearchNotesRequest {
	r.page = &page
	return r
}

// Search for a specific id.
func (r NotesAPISearchNotesRequest) SearchId(searchId int32) NotesAPISearchNotesRequest {
	r.searchId = &searchId
	return r
}

func (r NotesAPISearchNotesRequest) SearchOrder(searchOrder string) NotesAPISearchNotesRequest {
	r.searchOrder = &searchOrder
	return r
}

func (r NotesAPISearchNotesRequest) SearchBodyMatches(searchBodyMatches string) NotesAPISearchNotesRequest {
	r.searchBodyMatches = &searchBodyMatches
	return r
}

func (r NotesAPISearchNotesRequest) SearchIsActive(searchIsActive bool) NotesAPISearchNotesRequest {
	r.searchIsActive = &searchIsActive
	return r
}

func (r NotesAPISearchNotesRequest) SearchPostId(searchPostId float32) NotesAPISearchNotesRequest {
	r.searchPostId = &searchPostId
	return r
}

func (r NotesAPISearchNotesRequest) SearchPostTagsMatch(searchPostTagsMatch string) NotesAPISearchNotesRequest {
	r.searchPostTagsMatch = &searchPostTagsMatch
	return r
}

func (r NotesAPISearchNotesRequest) SearchPostNoteUpdaterId(searchPostNoteUpdaterId float32) NotesAPISearchNotesRequest {
	r.searchPostNoteUpdaterId = &searchPostNoteUpdaterId
	return r
}

func (r NotesAPISearchNotesRequest) SearchPostNoteUpdaterName(searchPostNoteUpdaterName string) NotesAPISearchNotesRequest {
	r.searchPostNoteUpdaterName = &searchPostNoteUpdaterName
	return r
}

func (r NotesAPISearchNotesRequest) SearchCreatorId(searchCreatorId float32) NotesAPISearchNotesRequest {
	r.searchCreatorId = &searchCreatorId
	return r
}

func (r NotesAPISearchNotesRequest) SearchCreatorName(searchCreatorName string) NotesAPISearchNotesRequest {
	r.searchCreatorName = &searchCreatorName
	return r
}

func (r NotesAPISearchNotesRequest) Execute() (*SearchNotes200Response, *http.Response, error) {
	return r.ApiService.SearchNotesExecute(r)
}

/*
SearchNotes Search Notes

When no results are found, an object with an `notes` key is returned.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return NotesAPISearchNotesRequest
*/
func (a *NotesAPIService) SearchNotes(ctx context.Context) NotesAPISearchNotesRequest {
	return NotesAPISearchNotesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SearchNotes200Response
func (a *NotesAPIService) SearchNotesExecute(r NotesAPISearchNotesRequest) (*SearchNotes200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SearchNotes200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NotesAPIService.SearchNotes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notes.json"

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
	if r.searchBodyMatches != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[body_matches]", r.searchBodyMatches, "")
	}
	if r.searchIsActive != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[is_active]", r.searchIsActive, "")
	}
	if r.searchPostId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[post_id]", r.searchPostId, "")
	}
	if r.searchPostTagsMatch != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[post_tags_match]", r.searchPostTagsMatch, "")
	}
	if r.searchPostNoteUpdaterId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[post_note_updater_id]", r.searchPostNoteUpdaterId, "")
	}
	if r.searchPostNoteUpdaterName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[post_note_updater_name]", r.searchPostNoteUpdaterName, "")
	}
	if r.searchCreatorId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[creator_id]", r.searchCreatorId, "")
	}
	if r.searchCreatorName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[creator_name]", r.searchCreatorName, "")
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
