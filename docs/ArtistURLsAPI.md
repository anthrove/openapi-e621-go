# \ArtistURLsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchArtistUrls**](ArtistURLsAPI.md#SearchArtistUrls) | **Get** /artist_urls.json | Search Artist URLs



## SearchArtistUrls

> []SearchArtistUrls200ResponseInner SearchArtistUrls(ctx).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchArtistName(searchArtistName).SearchArtistId(searchArtistId).SearchIsActive(searchIsActive).SearchUrl(searchUrl).SearchNormalizedUrl(searchNormalizedUrl).SearchArtist(searchArtist).SearchUrlMatches(searchUrlMatches).SearchNormalizedUrlMatches(searchNormalizedUrlMatches).Execute()

Search Artist URLs

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
	searchOrder := "searchOrder_example" // string | The order of the results. (optional)
	searchArtistName := "searchArtistName_example" // string | The name of the artist. (optional)
	searchArtistId := "searchArtistId_example" // string | The id of the artist. (optional)
	searchIsActive := true // bool | If the artist url is active. (optional)
	searchUrl := "searchUrl_example" // string | The url. (optional)
	searchNormalizedUrl := "searchNormalizedUrl_example" // string | The normalized url. (http:, trailing `/`) (optional)
	searchArtist := map[string]interface{}{ ... } // map[string]interface{} | Legacy nested search for artist. Supports the same parameters as /artists.json. (optional)
	searchUrlMatches := "searchUrlMatches_example" // string | Legacy name for `search[url]`. (optional)
	searchNormalizedUrlMatches := "searchNormalizedUrlMatches_example" // string | Legacy name for `search[normalized_url]`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtistURLsAPI.SearchArtistUrls(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchArtistName(searchArtistName).SearchArtistId(searchArtistId).SearchIsActive(searchIsActive).SearchUrl(searchUrl).SearchNormalizedUrl(searchNormalizedUrl).SearchArtist(searchArtist).SearchUrlMatches(searchUrlMatches).SearchNormalizedUrlMatches(searchNormalizedUrlMatches).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtistURLsAPI.SearchArtistUrls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchArtistUrls`: []SearchArtistUrls200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ArtistURLsAPI.SearchArtistUrls`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchArtistUrlsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **float32** | Search for a specific id. | 
 **searchOrder** | **string** | The order of the results. | 
 **searchArtistName** | **string** | The name of the artist. | 
 **searchArtistId** | **string** | The id of the artist. | 
 **searchIsActive** | **bool** | If the artist url is active. | 
 **searchUrl** | **string** | The url. | 
 **searchNormalizedUrl** | **string** | The normalized url. (http:, trailing &#x60;/&#x60;) | 
 **searchArtist** | [**map[string]interface{}**](map[string]interface{}.md) | Legacy nested search for artist. Supports the same parameters as /artists.json. | 
 **searchUrlMatches** | **string** | Legacy name for &#x60;search[url]&#x60;. | 
 **searchNormalizedUrlMatches** | **string** | Legacy name for &#x60;search[normalized_url]&#x60;. | 

### Return type

[**[]SearchArtistUrls200ResponseInner**](SearchArtistUrls200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

