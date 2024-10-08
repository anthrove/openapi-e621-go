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


// TicketsAPIService TicketsAPI service
type TicketsAPIService service

type ApiClaimTicketRequest struct {
	ctx context.Context
	ApiService *TicketsAPIService
	id float32
}

func (r ApiClaimTicketRequest) Execute() (*Ticket, *http.Response, error) {
	return r.ApiService.ClaimTicketExecute(r)
}

/*
ClaimTicket Claim Ticket

You must be Moderator+. Errors are quietly swallowed and shown as notices in html.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the ticket.
 @return ApiClaimTicketRequest
*/
func (a *TicketsAPIService) ClaimTicket(ctx context.Context, id float32) ApiClaimTicketRequest {
	return ApiClaimTicketRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Ticket
func (a *TicketsAPIService) ClaimTicketExecute(r ApiClaimTicketRequest) (*Ticket, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ticket
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TicketsAPIService.ClaimTicket")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tickets/{id}/claim.json"
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

type ApiEditTicketRequest struct {
	ctx context.Context
	ApiService *TicketsAPIService
	id float32
	ticketResponse *string
	ticketStatus *string
	ticketRecordType *WarningTypes
	ticketSendUpdateDmail *bool
}

func (r ApiEditTicketRequest) TicketResponse(ticketResponse string) ApiEditTicketRequest {
	r.ticketResponse = &ticketResponse
	return r
}

func (r ApiEditTicketRequest) TicketStatus(ticketStatus string) ApiEditTicketRequest {
	r.ticketStatus = &ticketStatus
	return r
}

func (r ApiEditTicketRequest) TicketRecordType(ticketRecordType WarningTypes) ApiEditTicketRequest {
	r.ticketRecordType = &ticketRecordType
	return r
}

// An update dmail will always be sent when the status is changed.
func (r ApiEditTicketRequest) TicketSendUpdateDmail(ticketSendUpdateDmail bool) ApiEditTicketRequest {
	r.ticketSendUpdateDmail = &ticketSendUpdateDmail
	return r
}

func (r ApiEditTicketRequest) Execute() (*http.Response, error) {
	return r.ApiService.EditTicketExecute(r)
}

/*
EditTicket Edit Ticket

You must be Moderator+.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the ticket.
 @return ApiEditTicketRequest
*/
func (a *TicketsAPIService) EditTicket(ctx context.Context, id float32) ApiEditTicketRequest {
	return ApiEditTicketRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *TicketsAPIService) EditTicketExecute(r ApiEditTicketRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TicketsAPIService.EditTicket")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tickets/{id}.json"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ticketResponse == nil {
		return nil, reportError("ticketResponse is required and must be specified")
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
	if r.ticketStatus != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "ticket[status]", r.ticketStatus, "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "ticket[response]", r.ticketResponse, "")
	if r.ticketRecordType != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "ticket[record_type]", r.ticketRecordType, "")
	}
	if r.ticketSendUpdateDmail != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "ticket[send_update_dmail]", r.ticketSendUpdateDmail, "")
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

type ApiGetTicketRequest struct {
	ctx context.Context
	ApiService *TicketsAPIService
	id float32
}

func (r ApiGetTicketRequest) Execute() (*Ticket, *http.Response, error) {
	return r.ApiService.GetTicketExecute(r)
}

/*
GetTicket Get Ticket

You must be Janitor+ to see tickets you did not create.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the ticket.
 @return ApiGetTicketRequest
*/
func (a *TicketsAPIService) GetTicket(ctx context.Context, id float32) ApiGetTicketRequest {
	return ApiGetTicketRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Ticket
func (a *TicketsAPIService) GetTicketExecute(r ApiGetTicketRequest) (*Ticket, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ticket
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TicketsAPIService.GetTicket")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tickets/{id}.json"
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

type ApiSearchTicketsRequest struct {
	ctx context.Context
	ApiService *TicketsAPIService
	limit *float32
	page *float32
	searchId *float32
	searchCreatorName *string
	searchCreatorId *float32
	searchClaimantName *string
	searchClaimantId *float32
	searchAccusedName *string
	searchAccusedId *float32
	searchQtype *TicketTypes
	searchReason *string
	searchStatus *string
}

// The maximum number of results to return. Between 0 and 320.
func (r ApiSearchTicketsRequest) Limit(limit float32) ApiSearchTicketsRequest {
	r.limit = &limit
	return r
}

// The page number of results to get. Between 1 and 750.
func (r ApiSearchTicketsRequest) Page(page float32) ApiSearchTicketsRequest {
	r.page = &page
	return r
}

// Search for a specific id.
func (r ApiSearchTicketsRequest) SearchId(searchId float32) ApiSearchTicketsRequest {
	r.searchId = &searchId
	return r
}

// You must be Moderator+.
func (r ApiSearchTicketsRequest) SearchCreatorName(searchCreatorName string) ApiSearchTicketsRequest {
	r.searchCreatorName = &searchCreatorName
	return r
}

// You must be Moderator+ unless providing your own id.
func (r ApiSearchTicketsRequest) SearchCreatorId(searchCreatorId float32) ApiSearchTicketsRequest {
	r.searchCreatorId = &searchCreatorId
	return r
}

// You must be Moderator+.
func (r ApiSearchTicketsRequest) SearchClaimantName(searchClaimantName string) ApiSearchTicketsRequest {
	r.searchClaimantName = &searchClaimantName
	return r
}

// You must be Moderator+.
func (r ApiSearchTicketsRequest) SearchClaimantId(searchClaimantId float32) ApiSearchTicketsRequest {
	r.searchClaimantId = &searchClaimantId
	return r
}

// You must be Moderator+.
func (r ApiSearchTicketsRequest) SearchAccusedName(searchAccusedName string) ApiSearchTicketsRequest {
	r.searchAccusedName = &searchAccusedName
	return r
}

// You must be Moderator+.
func (r ApiSearchTicketsRequest) SearchAccusedId(searchAccusedId float32) ApiSearchTicketsRequest {
	r.searchAccusedId = &searchAccusedId
	return r
}

func (r ApiSearchTicketsRequest) SearchQtype(searchQtype TicketTypes) ApiSearchTicketsRequest {
	r.searchQtype = &searchQtype
	return r
}

// You must be Moderator+.
func (r ApiSearchTicketsRequest) SearchReason(searchReason string) ApiSearchTicketsRequest {
	r.searchReason = &searchReason
	return r
}

func (r ApiSearchTicketsRequest) SearchStatus(searchStatus string) ApiSearchTicketsRequest {
	r.searchStatus = &searchStatus
	return r
}

func (r ApiSearchTicketsRequest) Execute() (*SearchTickets200Response, *http.Response, error) {
	return r.ApiService.SearchTicketsExecute(r)
}

/*
SearchTickets Search Tickets

You must be Janitor+ to see tickets you did not create. When no results are found, an object with an `tickets` key is returned.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchTicketsRequest
*/
func (a *TicketsAPIService) SearchTickets(ctx context.Context) ApiSearchTicketsRequest {
	return ApiSearchTicketsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SearchTickets200Response
func (a *TicketsAPIService) SearchTicketsExecute(r ApiSearchTicketsRequest) (*SearchTickets200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SearchTickets200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TicketsAPIService.SearchTickets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tickets.json"

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
	if r.searchCreatorName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[creator_name]", r.searchCreatorName, "")
	}
	if r.searchCreatorId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[creator_id]", r.searchCreatorId, "")
	}
	if r.searchClaimantName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[claimant_name]", r.searchClaimantName, "")
	}
	if r.searchClaimantId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[claimant_id]", r.searchClaimantId, "")
	}
	if r.searchAccusedName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[accused_name]", r.searchAccusedName, "")
	}
	if r.searchAccusedId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[accused_id]", r.searchAccusedId, "")
	}
	if r.searchQtype != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[qtype]", r.searchQtype, "")
	}
	if r.searchReason != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[reason]", r.searchReason, "")
	}
	if r.searchStatus != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[status]", r.searchStatus, "")
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
		if localVarHTTPResponse.StatusCode == 403 {
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

type ApiUnclaimTicketRequest struct {
	ctx context.Context
	ApiService *TicketsAPIService
	id float32
}

func (r ApiUnclaimTicketRequest) Execute() (*Ticket, *http.Response, error) {
	return r.ApiService.UnclaimTicketExecute(r)
}

/*
UnclaimTicket Unclaim Ticket

You must be Moderator+. Errors are quietly swallowed and shown as notices in html.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the ticket.
 @return ApiUnclaimTicketRequest
*/
func (a *TicketsAPIService) UnclaimTicket(ctx context.Context, id float32) ApiUnclaimTicketRequest {
	return ApiUnclaimTicketRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Ticket
func (a *TicketsAPIService) UnclaimTicketExecute(r ApiUnclaimTicketRequest) (*Ticket, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ticket
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TicketsAPIService.UnclaimTicket")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tickets/{id}/unclaim.json"
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