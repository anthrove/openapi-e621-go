# \AvoidPostingEntriesAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAvoidPosting**](AvoidPostingEntriesAPI.md#CreateAvoidPosting) | **Post** /avoid_postings.json | Create Avoid Posting Entry
[**DeleteAvoidPosting**](AvoidPostingEntriesAPI.md#DeleteAvoidPosting) | **Put** /avoid_postings/{idOrArtistName}/delete.json | Delete Avoid Posting Entry
[**DestroyAvoidPosting**](AvoidPostingEntriesAPI.md#DestroyAvoidPosting) | **Delete** /avoid_postings/{idOrArtistName}.json | Destroy Avoid Posting Entry
[**EditAvoidPosting**](AvoidPostingEntriesAPI.md#EditAvoidPosting) | **Patch** /avoid_postings/{idOrArtistName}.json | Edit Avoid Posting Entry
[**GetAvoidPosting**](AvoidPostingEntriesAPI.md#GetAvoidPosting) | **Get** /avoid_postings/{idOrArtistName}.json | Get Avoid Posting Entry
[**SearchAvoidPostings**](AvoidPostingEntriesAPI.md#SearchAvoidPostings) | **Get** /avoid_postings.json | Search Avoid Posting Entries
[**UndeleteAvoidPosting**](AvoidPostingEntriesAPI.md#UndeleteAvoidPosting) | **Put** /avoid_postings/{idOrArtistName}/undelete.json | Undelete Avoid Posting Entry



## CreateAvoidPosting

> AvoidPosting CreateAvoidPosting(ctx).AvoidPostingDetails(avoidPostingDetails).AvoidPostingStaffNotes(avoidPostingStaffNotes).AvoidPostingIsActive(avoidPostingIsActive).AvoidPostingArtistAttributesId(avoidPostingArtistAttributesId).AvoidPostingArtistAttributesName(avoidPostingArtistAttributesName).AvoidPostingArtistAttributesOtherNamesString(avoidPostingArtistAttributesOtherNamesString).AvoidPostingArtistAttributesOtherNames(avoidPostingArtistAttributesOtherNames).AvoidPostingArtistAttributesGroupName(avoidPostingArtistAttributesGroupName).AvoidPostingArtistAttributesLinkedUserId(avoidPostingArtistAttributesLinkedUserId).Execute()

Create Avoid Posting Entry



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
	avoidPostingDetails := "avoidPostingDetails_example" // string |  (optional)
	avoidPostingStaffNotes := "avoidPostingStaffNotes_example" // string |  (optional)
	avoidPostingIsActive := true // bool |  (optional)
	avoidPostingArtistAttributesId := float32(8.14) // float32 |  (optional)
	avoidPostingArtistAttributesName := "avoidPostingArtistAttributesName_example" // string | If provided and the artist does not exist, an artist will be created. (optional)
	avoidPostingArtistAttributesOtherNamesString := "avoidPostingArtistAttributesOtherNamesString_example" // string |  (optional)
	avoidPostingArtistAttributesOtherNames := []string{"Inner_example"} // []string |  (optional)
	avoidPostingArtistAttributesGroupName := "avoidPostingArtistAttributesGroupName_example" // string |  (optional)
	avoidPostingArtistAttributesLinkedUserId := float32(8.14) // float32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvoidPostingEntriesAPI.CreateAvoidPosting(context.Background()).AvoidPostingDetails(avoidPostingDetails).AvoidPostingStaffNotes(avoidPostingStaffNotes).AvoidPostingIsActive(avoidPostingIsActive).AvoidPostingArtistAttributesId(avoidPostingArtistAttributesId).AvoidPostingArtistAttributesName(avoidPostingArtistAttributesName).AvoidPostingArtistAttributesOtherNamesString(avoidPostingArtistAttributesOtherNamesString).AvoidPostingArtistAttributesOtherNames(avoidPostingArtistAttributesOtherNames).AvoidPostingArtistAttributesGroupName(avoidPostingArtistAttributesGroupName).AvoidPostingArtistAttributesLinkedUserId(avoidPostingArtistAttributesLinkedUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvoidPostingEntriesAPI.CreateAvoidPosting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAvoidPosting`: AvoidPosting
	fmt.Fprintf(os.Stdout, "Response from `AvoidPostingEntriesAPI.CreateAvoidPosting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAvoidPostingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **avoidPostingDetails** | **string** |  | 
 **avoidPostingStaffNotes** | **string** |  | 
 **avoidPostingIsActive** | **bool** |  | 
 **avoidPostingArtistAttributesId** | **float32** |  | 
 **avoidPostingArtistAttributesName** | **string** | If provided and the artist does not exist, an artist will be created. | 
 **avoidPostingArtistAttributesOtherNamesString** | **string** |  | 
 **avoidPostingArtistAttributesOtherNames** | **[]string** |  | 
 **avoidPostingArtistAttributesGroupName** | **string** |  | 
 **avoidPostingArtistAttributesLinkedUserId** | **float32** |  | 

### Return type

[**AvoidPosting**](AvoidPosting.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAvoidPosting

> DeleteAvoidPosting(ctx, idOrArtistName).Execute()

Delete Avoid Posting Entry



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
	idOrArtistName := "idOrArtistName_example" // string | The ID of the avoid posting entry, or the name of the artist.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AvoidPostingEntriesAPI.DeleteAvoidPosting(context.Background(), idOrArtistName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvoidPostingEntriesAPI.DeleteAvoidPosting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrArtistName** | **string** | The ID of the avoid posting entry, or the name of the artist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAvoidPostingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroyAvoidPosting

> DestroyAvoidPosting(ctx, idOrArtistName).Execute()

Destroy Avoid Posting Entry



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
	idOrArtistName := "idOrArtistName_example" // string | The ID of the avoid posting entry, or the name of the artist.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AvoidPostingEntriesAPI.DestroyAvoidPosting(context.Background(), idOrArtistName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvoidPostingEntriesAPI.DestroyAvoidPosting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrArtistName** | **string** | The ID of the avoid posting entry, or the name of the artist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyAvoidPostingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditAvoidPosting

> AvoidPosting EditAvoidPosting(ctx, idOrArtistName).AvoidPostingDetails(avoidPostingDetails).AvoidPostingStaffNotes(avoidPostingStaffNotes).AvoidPostingIsActive(avoidPostingIsActive).AvoidPostingArtistAttributesName(avoidPostingArtistAttributesName).AvoidPostingArtistAttributesOtherNamesString(avoidPostingArtistAttributesOtherNamesString).AvoidPostingArtistAttributesOtherNames(avoidPostingArtistAttributesOtherNames).AvoidPostingArtistAttributesGroupName(avoidPostingArtistAttributesGroupName).AvoidPostingArtistAttributesLinkedUserId(avoidPostingArtistAttributesLinkedUserId).Execute()

Edit Avoid Posting Entry



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
	idOrArtistName := "idOrArtistName_example" // string | The ID of the avoid posting entry, or the name of the artist.
	avoidPostingDetails := "avoidPostingDetails_example" // string |  (optional)
	avoidPostingStaffNotes := "avoidPostingStaffNotes_example" // string |  (optional)
	avoidPostingIsActive := true // bool |  (optional)
	avoidPostingArtistAttributesName := "avoidPostingArtistAttributesName_example" // string |  (optional)
	avoidPostingArtistAttributesOtherNamesString := "avoidPostingArtistAttributesOtherNamesString_example" // string |  (optional)
	avoidPostingArtistAttributesOtherNames := []string{"Inner_example"} // []string |  (optional)
	avoidPostingArtistAttributesGroupName := "avoidPostingArtistAttributesGroupName_example" // string |  (optional)
	avoidPostingArtistAttributesLinkedUserId := float32(8.14) // float32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvoidPostingEntriesAPI.EditAvoidPosting(context.Background(), idOrArtistName).AvoidPostingDetails(avoidPostingDetails).AvoidPostingStaffNotes(avoidPostingStaffNotes).AvoidPostingIsActive(avoidPostingIsActive).AvoidPostingArtistAttributesName(avoidPostingArtistAttributesName).AvoidPostingArtistAttributesOtherNamesString(avoidPostingArtistAttributesOtherNamesString).AvoidPostingArtistAttributesOtherNames(avoidPostingArtistAttributesOtherNames).AvoidPostingArtistAttributesGroupName(avoidPostingArtistAttributesGroupName).AvoidPostingArtistAttributesLinkedUserId(avoidPostingArtistAttributesLinkedUserId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvoidPostingEntriesAPI.EditAvoidPosting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditAvoidPosting`: AvoidPosting
	fmt.Fprintf(os.Stdout, "Response from `AvoidPostingEntriesAPI.EditAvoidPosting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrArtistName** | **string** | The ID of the avoid posting entry, or the name of the artist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAvoidPostingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **avoidPostingDetails** | **string** |  | 
 **avoidPostingStaffNotes** | **string** |  | 
 **avoidPostingIsActive** | **bool** |  | 
 **avoidPostingArtistAttributesName** | **string** |  | 
 **avoidPostingArtistAttributesOtherNamesString** | **string** |  | 
 **avoidPostingArtistAttributesOtherNames** | **[]string** |  | 
 **avoidPostingArtistAttributesGroupName** | **string** |  | 
 **avoidPostingArtistAttributesLinkedUserId** | **float32** |  | 

### Return type

[**AvoidPosting**](AvoidPosting.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvoidPosting

> AvoidPosting GetAvoidPosting(ctx, idOrArtistName).Execute()

Get Avoid Posting Entry

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
	idOrArtistName := "idOrArtistName_example" // string | The ID of the avoid posting entry, or the name of the artist.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvoidPostingEntriesAPI.GetAvoidPosting(context.Background(), idOrArtistName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvoidPostingEntriesAPI.GetAvoidPosting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvoidPosting`: AvoidPosting
	fmt.Fprintf(os.Stdout, "Response from `AvoidPostingEntriesAPI.GetAvoidPosting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrArtistName** | **string** | The ID of the avoid posting entry, or the name of the artist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvoidPostingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AvoidPosting**](AvoidPosting.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAvoidPostings

> SearchAvoidPostings200Response SearchAvoidPostings(ctx).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).SearchCreatorName(searchCreatorName).SearchCreatorId(searchCreatorId).SearchAnyNameMatches(searchAnyNameMatches).SearchArtistName(searchArtistName).SearchArtistId(searchArtistId).SearchAnyOtherNameMatches(searchAnyOtherNameMatches).SearchDetails(searchDetails).SearchStaffNotes(searchStaffNotes).SearchIsActive(searchIsActive).Execute()

Search Avoid Posting Entries



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
	limit := float32(8.14) // float32 | The maximum number of results to return. Between 0 and 320. (optional)
	page := float32(8.14) // float32 | The page number of results to get. Between 1 and 750. (optional)
	searchId := float32(8.14) // float32 | Search for a specific id. (optional)
	searchIpAddr := "searchIpAddr_example" // string | Must be Admin+ to use. See [postgres' documentation](https://www.postgresql.org/docs/9.3/functions-net.html) for information on how this is parsed. Specifically, \"is contained within or equals\" (`<<=`). (optional)
	searchOrder := "searchOrder_example" // string |  (optional)
	searchCreatorName := "searchCreatorName_example" // string |  (optional)
	searchCreatorId := "searchCreatorId_example" // string |  (optional)
	searchAnyNameMatches := "searchAnyNameMatches_example" // string | Any name matching. (optional)
	searchArtistName := "searchArtistName_example" // string | The artist name of the avoid posting entry. (optional)
	searchArtistId := "searchArtistId_example" // string | The artist id for the avoid posting entry. (optional)
	searchAnyOtherNameMatches := "searchAnyOtherNameMatches_example" // string | Any other name matching. (optional)
	searchDetails := "searchDetails_example" // string | The details of the avoid posting entry. (optional)
	searchStaffNotes := "searchStaffNotes_example" // string | The staff notes on the avoid posting entry. Must be Janitor+ to use. (optional)
	searchIsActive := true // bool | If the avoid posting entry is active. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvoidPostingEntriesAPI.SearchAvoidPostings(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).SearchCreatorName(searchCreatorName).SearchCreatorId(searchCreatorId).SearchAnyNameMatches(searchAnyNameMatches).SearchArtistName(searchArtistName).SearchArtistId(searchArtistId).SearchAnyOtherNameMatches(searchAnyOtherNameMatches).SearchDetails(searchDetails).SearchStaffNotes(searchStaffNotes).SearchIsActive(searchIsActive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvoidPostingEntriesAPI.SearchAvoidPostings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAvoidPostings`: SearchAvoidPostings200Response
	fmt.Fprintf(os.Stdout, "Response from `AvoidPostingEntriesAPI.SearchAvoidPostings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAvoidPostingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **float32** | Search for a specific id. | 
 **searchIpAddr** | **string** | Must be Admin+ to use. See [postgres&#39; documentation](https://www.postgresql.org/docs/9.3/functions-net.html) for information on how this is parsed. Specifically, \&quot;is contained within or equals\&quot; (&#x60;&lt;&lt;&#x3D;&#x60;). | 
 **searchOrder** | **string** |  | 
 **searchCreatorName** | **string** |  | 
 **searchCreatorId** | **string** |  | 
 **searchAnyNameMatches** | **string** | Any name matching. | 
 **searchArtistName** | **string** | The artist name of the avoid posting entry. | 
 **searchArtistId** | **string** | The artist id for the avoid posting entry. | 
 **searchAnyOtherNameMatches** | **string** | Any other name matching. | 
 **searchDetails** | **string** | The details of the avoid posting entry. | 
 **searchStaffNotes** | **string** | The staff notes on the avoid posting entry. Must be Janitor+ to use. | 
 **searchIsActive** | **bool** | If the avoid posting entry is active. | 

### Return type

[**SearchAvoidPostings200Response**](SearchAvoidPostings200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UndeleteAvoidPosting

> UndeleteAvoidPosting(ctx, idOrArtistName).Execute()

Undelete Avoid Posting Entry



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
	idOrArtistName := "idOrArtistName_example" // string | The ID of the avoid posting entry, or the name of the artist.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AvoidPostingEntriesAPI.UndeleteAvoidPosting(context.Background(), idOrArtistName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvoidPostingEntriesAPI.UndeleteAvoidPosting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrArtistName** | **string** | The ID of the avoid posting entry, or the name of the artist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUndeleteAvoidPostingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

