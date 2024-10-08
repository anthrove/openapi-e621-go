# \NoteVersionsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchNoteVersions**](NoteVersionsAPI.md#SearchNoteVersions) | **Get** /note_versions.json | Search Note Versions



## SearchNoteVersions

> SearchNoteVersions200Response SearchNoteVersions(ctx).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).SearchUpdaterId(searchUpdaterId).SearchUpdaterName(searchUpdaterName).SearchPostId(searchPostId).SearchNoteId(searchNoteId).SearchIsActive(searchIsActive).SearchBodyMatches(searchBodyMatches).Execute()

Search Note Versions



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
	searchOrder := "searchOrder_example" // string | The order of the results. (optional)
	searchUpdaterId := float32(8.14) // float32 |  (optional)
	searchUpdaterName := "searchUpdaterName_example" // string |  (optional)
	searchPostId := float32(8.14) // float32 |  (optional)
	searchNoteId := float32(8.14) // float32 |  (optional)
	searchIsActive := true // bool |  (optional)
	searchBodyMatches := "searchBodyMatches_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NoteVersionsAPI.SearchNoteVersions(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).SearchUpdaterId(searchUpdaterId).SearchUpdaterName(searchUpdaterName).SearchPostId(searchPostId).SearchNoteId(searchNoteId).SearchIsActive(searchIsActive).SearchBodyMatches(searchBodyMatches).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NoteVersionsAPI.SearchNoteVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchNoteVersions`: SearchNoteVersions200Response
	fmt.Fprintf(os.Stdout, "Response from `NoteVersionsAPI.SearchNoteVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchNoteVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **float32** | Search for a specific id. | 
 **searchIpAddr** | **string** | Must be Admin+ to use. See [postgres&#39; documentation](https://www.postgresql.org/docs/9.3/functions-net.html) for information on how this is parsed. Specifically, \&quot;is contained within or equals\&quot; (&#x60;&lt;&lt;&#x3D;&#x60;). | 
 **searchOrder** | **string** | The order of the results. | 
 **searchUpdaterId** | **float32** |  | 
 **searchUpdaterName** | **string** |  | 
 **searchPostId** | **float32** |  | 
 **searchNoteId** | **float32** |  | 
 **searchIsActive** | **bool** |  | 
 **searchBodyMatches** | **string** |  | 

### Return type

[**SearchNoteVersions200Response**](SearchNoteVersions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

