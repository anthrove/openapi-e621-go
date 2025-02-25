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

// TicketsAPIService TicketsAPI service
type TicketsAPIService service

type TicketsAPIClaimTicketRequest struct {
	ctx        context.Context
	ApiService *TicketsAPIService
	id         int32
}

func (r TicketsAPIClaimTicketRequest) Execute() (*Ticket, *http.Response, error) {
	return r.ApiService.ClaimTicketExecute(r)
}

/*
ClaimTicket Claim Ticket

You must be Moderator+. Errors are quietly swallowed and shown as notices in html.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the ticket.
	@return TicketsAPIClaimTicketRequest
*/
func (a *TicketsAPIService) ClaimTicket(ctx context.Context, id int32) TicketsAPIClaimTicketRequest {
	return TicketsAPIClaimTicketRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return Ticket
func (a *TicketsAPIService) ClaimTicketExecute(r TicketsAPIClaimTicketRequest) (*Ticket, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Ticket
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

type TicketsAPIEditTicketRequest struct {
	ctx                   context.Context
	ApiService            *TicketsAPIService
	id                    int32
	ticketResponse        *string
	ticketStatus          *string
	ticketRecordType      *WarningTypes
	ticketSendUpdateDmail *bool
}

func (r TicketsAPIEditTicketRequest) TicketResponse(ticketResponse string) TicketsAPIEditTicketRequest {
	r.ticketResponse = &ticketResponse
	return r
}

func (r TicketsAPIEditTicketRequest) TicketStatus(ticketStatus string) TicketsAPIEditTicketRequest {
	r.ticketStatus = &ticketStatus
	return r
}

func (r TicketsAPIEditTicketRequest) TicketRecordType(ticketRecordType WarningTypes) TicketsAPIEditTicketRequest {
	r.ticketRecordType = &ticketRecordType
	return r
}

// An update dmail will always be sent when the status is changed.
func (r TicketsAPIEditTicketRequest) TicketSendUpdateDmail(ticketSendUpdateDmail bool) TicketsAPIEditTicketRequest {
	r.ticketSendUpdateDmail = &ticketSendUpdateDmail
	return r
}

func (r TicketsAPIEditTicketRequest) Execute() (*http.Response, error) {
	return r.ApiService.EditTicketExecute(r)
}

/*
EditTicket Edit Ticket

You must be Moderator+.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the ticket.
	@return TicketsAPIEditTicketRequest
*/
func (a *TicketsAPIService) EditTicket(ctx context.Context, id int32) TicketsAPIEditTicketRequest {
	return TicketsAPIEditTicketRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *TicketsAPIService) EditTicketExecute(r TicketsAPIEditTicketRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
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

type TicketsAPIGetTicketRequest struct {
	ctx        context.Context
	ApiService *TicketsAPIService
	id         int32
}

func (r TicketsAPIGetTicketRequest) Execute() (*Ticket, *http.Response, error) {
	return r.ApiService.GetTicketExecute(r)
}

/*
GetTicket Get Ticket

You must be Janitor+ to see tickets you did not create.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the ticket.
	@return TicketsAPIGetTicketRequest
*/
func (a *TicketsAPIService) GetTicket(ctx context.Context, id int32) TicketsAPIGetTicketRequest {
	return TicketsAPIGetTicketRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return Ticket
func (a *TicketsAPIService) GetTicketExecute(r TicketsAPIGetTicketRequest) (*Ticket, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Ticket
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

type TicketsAPISearchTicketsRequest struct {
	ctx                context.Context
	ApiService         *TicketsAPIService
	limit              *int32
	page               *int32
	searchId           *int32
	searchCreatorName  *string
	searchCreatorId    *int32
	searchClaimantName *string
	searchClaimantId   *int32
	searchAccusedName  *string
	searchAccusedId    *int32
	searchQtype        *TicketTypes
	searchReason       *string
	searchStatus       *string
}

// The maximum number of results to return. Between 0 and 320.
func (r TicketsAPISearchTicketsRequest) Limit(limit int32) TicketsAPISearchTicketsRequest {
	r.limit = &limit
	return r
}

// The page number of results to get. Between 1 and 750.
func (r TicketsAPISearchTicketsRequest) Page(page int32) TicketsAPISearchTicketsRequest {
	r.page = &page
	return r
}

// Search for a specific id.
func (r TicketsAPISearchTicketsRequest) SearchId(searchId int32) TicketsAPISearchTicketsRequest {
	r.searchId = &searchId
	return r
}

// You must be Moderator+.
func (r TicketsAPISearchTicketsRequest) SearchCreatorName(searchCreatorName string) TicketsAPISearchTicketsRequest {
	r.searchCreatorName = &searchCreatorName
	return r
}

// You must be Moderator+ unless providing your own id.
func (r TicketsAPISearchTicketsRequest) SearchCreatorId(searchCreatorId int32) TicketsAPISearchTicketsRequest {
	r.searchCreatorId = &searchCreatorId
	return r
}

// You must be Moderator+.
func (r TicketsAPISearchTicketsRequest) SearchClaimantName(searchClaimantName string) TicketsAPISearchTicketsRequest {
	r.searchClaimantName = &searchClaimantName
	return r
}

// You must be Moderator+.
func (r TicketsAPISearchTicketsRequest) SearchClaimantId(searchClaimantId int32) TicketsAPISearchTicketsRequest {
	r.searchClaimantId = &searchClaimantId
	return r
}

// You must be Moderator+.
func (r TicketsAPISearchTicketsRequest) SearchAccusedName(searchAccusedName string) TicketsAPISearchTicketsRequest {
	r.searchAccusedName = &searchAccusedName
	return r
}

// You must be Moderator+.
func (r TicketsAPISearchTicketsRequest) SearchAccusedId(searchAccusedId int32) TicketsAPISearchTicketsRequest {
	r.searchAccusedId = &searchAccusedId
	return r
}

func (r TicketsAPISearchTicketsRequest) SearchQtype(searchQtype TicketTypes) TicketsAPISearchTicketsRequest {
	r.searchQtype = &searchQtype
	return r
}

// You must be Moderator+.
func (r TicketsAPISearchTicketsRequest) SearchReason(searchReason string) TicketsAPISearchTicketsRequest {
	r.searchReason = &searchReason
	return r
}

func (r TicketsAPISearchTicketsRequest) SearchStatus(searchStatus string) TicketsAPISearchTicketsRequest {
	r.searchStatus = &searchStatus
	return r
}

func (r TicketsAPISearchTicketsRequest) Execute() (*SearchTickets200Response, *http.Response, error) {
	return r.ApiService.SearchTicketsExecute(r)
}

/*
SearchTickets Search Tickets

You must be Janitor+ to see tickets you did not create. When no results are found, an object with an `tickets` key is returned.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TicketsAPISearchTicketsRequest
*/
func (a *TicketsAPIService) SearchTickets(ctx context.Context) TicketsAPISearchTicketsRequest {
	return TicketsAPISearchTicketsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return SearchTickets200Response
func (a *TicketsAPIService) SearchTicketsExecute(r TicketsAPISearchTicketsRequest) (*SearchTickets200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SearchTickets200Response
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

type TicketsAPIUnclaimTicketRequest struct {
	ctx        context.Context
	ApiService *TicketsAPIService
	id         int32
}

func (r TicketsAPIUnclaimTicketRequest) Execute() (*Ticket, *http.Response, error) {
	return r.ApiService.UnclaimTicketExecute(r)
}

/*
UnclaimTicket Unclaim Ticket

You must be Moderator+. Errors are quietly swallowed and shown as notices in html.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the ticket.
	@return TicketsAPIUnclaimTicketRequest
*/
func (a *TicketsAPIService) UnclaimTicket(ctx context.Context, id int32) TicketsAPIUnclaimTicketRequest {
	return TicketsAPIUnclaimTicketRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return Ticket
func (a *TicketsAPIService) UnclaimTicketExecute(r TicketsAPIUnclaimTicketRequest) (*Ticket, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *Ticket
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
