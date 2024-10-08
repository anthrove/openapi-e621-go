# \WikiPageVersionsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWikiPageVersion**](WikiPageVersionsAPI.md#GetWikiPageVersion) | **Get** /wiki_page_versions/{id}.json | Get Wiki Page Version
[**SearchWikiPageVersions**](WikiPageVersionsAPI.md#SearchWikiPageVersions) | **Get** /wiki_page_versions.json | Search Wiki Page Versions



## GetWikiPageVersion

> WikiPageVersion GetWikiPageVersion(ctx, id).Execute()

Get Wiki Page Version

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
	id := float32(8.14) // float32 | The ID of the wiki page version.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiPageVersionsAPI.GetWikiPageVersion(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiPageVersionsAPI.GetWikiPageVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWikiPageVersion`: WikiPageVersion
	fmt.Fprintf(os.Stdout, "Response from `WikiPageVersionsAPI.GetWikiPageVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the wiki page version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWikiPageVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WikiPageVersion**](WikiPageVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchWikiPageVersions

> SearchWikiPageVersions200Response SearchWikiPageVersions(ctx).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).SearchUpdaterId(searchUpdaterId).SearchUpdaterName(searchUpdaterName).SearchWikiPageId(searchWikiPageId).SearchTitle(searchTitle).SearchBody(searchBody).SearchIsLocked(searchIsLocked).SearchIsDeleted(searchIsDeleted).Execute()

Search Wiki Page Versions



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
	searchOrder := "searchOrder_example" // string |  (optional)
	searchUpdaterId := float32(8.14) // float32 |  (optional)
	searchUpdaterName := "searchUpdaterName_example" // string |  (optional)
	searchWikiPageId := float32(8.14) // float32 |  (optional)
	searchTitle := "searchTitle_example" // string |  (optional)
	searchBody := "searchBody_example" // string |  (optional)
	searchIsLocked := true // bool |  (optional)
	searchIsDeleted := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiPageVersionsAPI.SearchWikiPageVersions(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).SearchUpdaterId(searchUpdaterId).SearchUpdaterName(searchUpdaterName).SearchWikiPageId(searchWikiPageId).SearchTitle(searchTitle).SearchBody(searchBody).SearchIsLocked(searchIsLocked).SearchIsDeleted(searchIsDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiPageVersionsAPI.SearchWikiPageVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchWikiPageVersions`: SearchWikiPageVersions200Response
	fmt.Fprintf(os.Stdout, "Response from `WikiPageVersionsAPI.SearchWikiPageVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchWikiPageVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **float32** | Search for a specific id. | 
 **searchIpAddr** | **string** | Must be Admin+ to use. See [postgres&#39; documentation](https://www.postgresql.org/docs/9.3/functions-net.html) for information on how this is parsed. Specifically, \&quot;is contained within or equals\&quot; (&#x60;&lt;&lt;&#x3D;&#x60;). | 
 **searchOrder** | **string** |  | 
 **searchUpdaterId** | **float32** |  | 
 **searchUpdaterName** | **string** |  | 
 **searchWikiPageId** | **float32** |  | 
 **searchTitle** | **string** |  | 
 **searchBody** | **string** |  | 
 **searchIsLocked** | **bool** |  | 
 **searchIsDeleted** | **bool** |  | 

### Return type

[**SearchWikiPageVersions200Response**](SearchWikiPageVersions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

