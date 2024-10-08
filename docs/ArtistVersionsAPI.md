# \ArtistVersionsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchArtistVersions**](ArtistVersionsAPI.md#SearchArtistVersions) | **Get** /artist_versions.json | Search Artist Versions



## SearchArtistVersions

> SearchArtistVersions200Response SearchArtistVersions(ctx).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).SearchName(searchName).SearchArtistId(searchArtistId).SearchUpdaterName(searchUpdaterName).SearchUpdaterId(searchUpdaterId).Execute()

Search Artist Versions



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
	searchName := "searchName_example" // string | The name of the artist. (optional)
	searchArtistId := "searchArtistId_example" // string | The id of the artist. (optional)
	searchUpdaterName := "searchUpdaterName_example" // string | The name of the user that updated the artist. (optional)
	searchUpdaterId := "searchUpdaterId_example" // string | The id of the user that updated the artist. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtistVersionsAPI.SearchArtistVersions(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).SearchName(searchName).SearchArtistId(searchArtistId).SearchUpdaterName(searchUpdaterName).SearchUpdaterId(searchUpdaterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtistVersionsAPI.SearchArtistVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchArtistVersions`: SearchArtistVersions200Response
	fmt.Fprintf(os.Stdout, "Response from `ArtistVersionsAPI.SearchArtistVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchArtistVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **float32** | Search for a specific id. | 
 **searchIpAddr** | **string** | Must be Admin+ to use. See [postgres&#39; documentation](https://www.postgresql.org/docs/9.3/functions-net.html) for information on how this is parsed. Specifically, \&quot;is contained within or equals\&quot; (&#x60;&lt;&lt;&#x3D;&#x60;). | 
 **searchOrder** | **string** | The order of the results. | 
 **searchName** | **string** | The name of the artist. | 
 **searchArtistId** | **string** | The id of the artist. | 
 **searchUpdaterName** | **string** | The name of the user that updated the artist. | 
 **searchUpdaterId** | **string** | The id of the user that updated the artist. | 

### Return type

[**SearchArtistVersions200Response**](SearchArtistVersions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

