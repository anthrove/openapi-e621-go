# \AvoidPostingVersionsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchAvoidPostingVersions**](AvoidPostingVersionsAPI.md#SearchAvoidPostingVersions) | **Get** /avoid_posting_versions.json | Search Avoid Posting Versions



## SearchAvoidPostingVersions

> SearchAvoidPostingVersions200Response SearchAvoidPostingVersions(ctx).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).SearchUpdaterName(searchUpdaterName).SearchUpdaterId(searchUpdaterId).SearchAnyNameMatches(searchAnyNameMatches).SearchArtistName(searchArtistName).SearchArtistId(searchArtistId).SearchAnyOtherNameMatches(searchAnyOtherNameMatches).SearchGroupName(searchGroupName).SearchIsActive(searchIsActive).Execute()

Search Avoid Posting Versions

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	limit := float32(8.14) // float32 | The maximum number of results to return. Between 0 and 320. (optional)
	page := float32(8.14) // float32 | The page number of results to get. Between 1 and 750. (optional)
	searchId := float32(8.14) // float32 | Search for a specific id. (optional)
	searchIpAddr := "searchIpAddr_example" // string | Must be Admin+ to use. See [postgres' documentation](https://www.postgresql.org/docs/9.3/functions-net.html) for information on how this is parsed. Specifically, \"is contained within or equals\" (`<<=`). (optional)
	searchOrder := "searchOrder_example" // string | The order of the results. (optional)
	searchUpdaterName := "searchUpdaterName_example" // string | The name of the updater of the avoid posting entry. (optional)
	searchUpdaterId := "searchUpdaterId_example" // string | The ID of the updater of the avoid posting entry. (optional)
	searchAnyNameMatches := "searchAnyNameMatches_example" // string | Any name matching. (optional)
	searchArtistName := "searchArtistName_example" // string | The artist name of the avoid posting entry. (optional)
	searchArtistId := "searchArtistId_example" // string | The artist id for the avoid posting entry. (optional)
	searchAnyOtherNameMatches := "searchAnyOtherNameMatches_example" // string | Any other name matching. (optional)
	searchGroupName := "searchGroupName_example" // string | Any other name matching. (optional)
	searchIsActive := true // bool | If the avoid posting entry is active. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvoidPostingVersionsAPI.SearchAvoidPostingVersions(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).SearchUpdaterName(searchUpdaterName).SearchUpdaterId(searchUpdaterId).SearchAnyNameMatches(searchAnyNameMatches).SearchArtistName(searchArtistName).SearchArtistId(searchArtistId).SearchAnyOtherNameMatches(searchAnyOtherNameMatches).SearchGroupName(searchGroupName).SearchIsActive(searchIsActive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvoidPostingVersionsAPI.SearchAvoidPostingVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAvoidPostingVersions`: SearchAvoidPostingVersions200Response
	fmt.Fprintf(os.Stdout, "Response from `AvoidPostingVersionsAPI.SearchAvoidPostingVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAvoidPostingVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **float32** | Search for a specific id. | 
 **searchIpAddr** | **string** | Must be Admin+ to use. See [postgres&#39; documentation](https://www.postgresql.org/docs/9.3/functions-net.html) for information on how this is parsed. Specifically, \&quot;is contained within or equals\&quot; (&#x60;&lt;&lt;&#x3D;&#x60;). | 
 **searchOrder** | **string** | The order of the results. | 
 **searchUpdaterName** | **string** | The name of the updater of the avoid posting entry. | 
 **searchUpdaterId** | **string** | The ID of the updater of the avoid posting entry. | 
 **searchAnyNameMatches** | **string** | Any name matching. | 
 **searchArtistName** | **string** | The artist name of the avoid posting entry. | 
 **searchArtistId** | **string** | The artist id for the avoid posting entry. | 
 **searchAnyOtherNameMatches** | **string** | Any other name matching. | 
 **searchGroupName** | **string** | Any other name matching. | 
 **searchIsActive** | **bool** | If the avoid posting entry is active. | 

### Return type

[**SearchAvoidPostingVersions200Response**](SearchAvoidPostingVersions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

