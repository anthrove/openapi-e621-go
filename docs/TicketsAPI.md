# \TicketsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClaimTicket**](TicketsAPI.md#ClaimTicket) | **Post** /tickets/{id}/claim.json | Claim Ticket
[**EditTicket**](TicketsAPI.md#EditTicket) | **Patch** /tickets/{id}.json | Edit Ticket
[**GetTicket**](TicketsAPI.md#GetTicket) | **Get** /tickets/{id}.json | Get Ticket
[**SearchTickets**](TicketsAPI.md#SearchTickets) | **Get** /tickets.json | Search Tickets
[**UnclaimTicket**](TicketsAPI.md#UnclaimTicket) | **Post** /tickets/{id}/unclaim.json | Unclaim Ticket



## ClaimTicket

> Ticket ClaimTicket(ctx, id).Execute()

Claim Ticket



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anthrove/openapi-e621-go"
)

func main() {
	id := float32(8.14) // float32 | The ID of the ticket.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketsAPI.ClaimTicket(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketsAPI.ClaimTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ClaimTicket`: Ticket
	fmt.Fprintf(os.Stdout, "Response from `TicketsAPI.ClaimTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the ticket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiClaimTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Ticket**](Ticket.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditTicket

> EditTicket(ctx, id).TicketResponse(ticketResponse).TicketStatus(ticketStatus).TicketRecordType(ticketRecordType).TicketSendUpdateDmail(ticketSendUpdateDmail).Execute()

Edit Ticket



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anthrove/openapi-e621-go"
)

func main() {
	id := float32(8.14) // float32 | The ID of the ticket.
	ticketResponse := "ticketResponse_example" // string | 
	ticketStatus := "ticketStatus_example" // string |  (optional)
	ticketRecordType := openapiclient.WarningTypes("warning") // WarningTypes |  (optional)
	ticketSendUpdateDmail := true // bool | An update dmail will always be sent when the status is changed. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TicketsAPI.EditTicket(context.Background(), id).TicketResponse(ticketResponse).TicketStatus(ticketStatus).TicketRecordType(ticketRecordType).TicketSendUpdateDmail(ticketSendUpdateDmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketsAPI.EditTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the ticket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ticketResponse** | **string** |  | 
 **ticketStatus** | **string** |  | 
 **ticketRecordType** | [**WarningTypes**](WarningTypes.md) |  | 
 **ticketSendUpdateDmail** | **bool** | An update dmail will always be sent when the status is changed. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTicket

> Ticket GetTicket(ctx, id).Execute()

Get Ticket



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anthrove/openapi-e621-go"
)

func main() {
	id := float32(8.14) // float32 | The ID of the ticket.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketsAPI.GetTicket(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketsAPI.GetTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTicket`: Ticket
	fmt.Fprintf(os.Stdout, "Response from `TicketsAPI.GetTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the ticket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Ticket**](Ticket.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchTickets

> SearchTickets200Response SearchTickets(ctx).Limit(limit).Page(page).SearchId(searchId).SearchCreatorName(searchCreatorName).SearchCreatorId(searchCreatorId).SearchClaimantName(searchClaimantName).SearchClaimantId(searchClaimantId).SearchAccusedName(searchAccusedName).SearchAccusedId(searchAccusedId).SearchQtype(searchQtype).SearchReason(searchReason).SearchStatus(searchStatus).Execute()

Search Tickets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anthrove/openapi-e621-go"
)

func main() {
	limit := int32(56) // int32 | The maximum number of results to return. Between 0 and 320. (optional)
	page := int32(56) // int32 | The page number of results to get. Between 1 and 750. (optional)
	searchId := int32(56) // int32 | Search for a specific id. (optional)
	searchCreatorName := "searchCreatorName_example" // string | You must be Moderator+. (optional)
	searchCreatorId := float32(8.14) // float32 | You must be Moderator+ unless providing your own id. (optional)
	searchClaimantName := "searchClaimantName_example" // string | You must be Moderator+. (optional)
	searchClaimantId := float32(8.14) // float32 | You must be Moderator+. (optional)
	searchAccusedName := "searchAccusedName_example" // string | You must be Moderator+. (optional)
	searchAccusedId := float32(8.14) // float32 | You must be Moderator+. (optional)
	searchQtype := openapiclient.TicketTypes("blip") // TicketTypes |  (optional)
	searchReason := "searchReason_example" // string | You must be Moderator+. (optional)
	searchStatus := "searchStatus_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketsAPI.SearchTickets(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchCreatorName(searchCreatorName).SearchCreatorId(searchCreatorId).SearchClaimantName(searchClaimantName).SearchClaimantId(searchClaimantId).SearchAccusedName(searchAccusedName).SearchAccusedId(searchAccusedId).SearchQtype(searchQtype).SearchReason(searchReason).SearchStatus(searchStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketsAPI.SearchTickets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchTickets`: SearchTickets200Response
	fmt.Fprintf(os.Stdout, "Response from `TicketsAPI.SearchTickets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchTicketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **int32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **int32** | Search for a specific id. | 
 **searchCreatorName** | **string** | You must be Moderator+. | 
 **searchCreatorId** | **float32** | You must be Moderator+ unless providing your own id. | 
 **searchClaimantName** | **string** | You must be Moderator+. | 
 **searchClaimantId** | **float32** | You must be Moderator+. | 
 **searchAccusedName** | **string** | You must be Moderator+. | 
 **searchAccusedId** | **float32** | You must be Moderator+. | 
 **searchQtype** | [**TicketTypes**](TicketTypes.md) |  | 
 **searchReason** | **string** | You must be Moderator+. | 
 **searchStatus** | **string** |  | 

### Return type

[**SearchTickets200Response**](SearchTickets200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnclaimTicket

> Ticket UnclaimTicket(ctx, id).Execute()

Unclaim Ticket



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/anthrove/openapi-e621-go"
)

func main() {
	id := float32(8.14) // float32 | The ID of the ticket.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TicketsAPI.UnclaimTicket(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TicketsAPI.UnclaimTicket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnclaimTicket`: Ticket
	fmt.Fprintf(os.Stdout, "Response from `TicketsAPI.UnclaimTicket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the ticket. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnclaimTicketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Ticket**](Ticket.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

