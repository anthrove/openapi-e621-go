# \PostVersionsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchPostVersions**](PostVersionsAPI.md#SearchPostVersions) | **Get** /post_versions.json | Search Post Versions



## SearchPostVersions

> []PostVersion SearchPostVersions(ctx).Limit(limit).Page(page).SearchId(searchId).SearchUpdaterName(searchUpdaterName).SearchUpdaterId(searchUpdaterId).SearchPostId(searchPostId).SearchStartId(searchStartId).SearchRating(searchRating).SearchRatingChanged(searchRatingChanged).SearchParentId(searchParentId).SearchParentIdChanged(searchParentIdChanged).SearchTags(searchTags).SearchTagsRemoved(searchTagsRemoved).SearchTagsAdded(searchTagsAdded).SearchLockedTags(searchLockedTags).SearchLockedTagsRemoved(searchLockedTagsRemoved).SearchLockedTagsAdded(searchLockedTagsAdded).SearchReason(searchReason).SearchDescription(searchDescription).SearchDescriptionChanged(searchDescriptionChanged).SearchSourceChanged(searchSourceChanged).SearchUploads(searchUploads).Execute()

Search Post Versions



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
	page := float32(8.14) // float32 | The page number of results to get. Between 1 and 750. Note that for post versions specifically, you can only go through the 10,000 most recent results with page numbers. (optional)
	searchId := int32(56) // int32 | Search for a specific id. (optional)
	searchUpdaterName := "searchUpdaterName_example" // string |  (optional)
	searchUpdaterId := float32(8.14) // float32 |  (optional)
	searchPostId := float32(8.14) // float32 |  (optional)
	searchStartId := float32(8.14) // float32 |  (optional)
	searchRating := openapiclient.Ratings("s") // Ratings |  (optional)
	searchRatingChanged := "searchRatingChanged_example" // string |  (optional)
	searchParentId := float32(8.14) // float32 |  (optional)
	searchParentIdChanged := true // bool |  (optional)
	searchTags := "searchTags_example" // string |  (optional)
	searchTagsRemoved := "searchTagsRemoved_example" // string |  (optional)
	searchTagsAdded := "searchTagsAdded_example" // string |  (optional)
	searchLockedTags := "searchLockedTags_example" // string |  (optional)
	searchLockedTagsRemoved := "searchLockedTagsRemoved_example" // string |  (optional)
	searchLockedTagsAdded := "searchLockedTagsAdded_example" // string |  (optional)
	searchReason := "searchReason_example" // string |  (optional)
	searchDescription := "searchDescription_example" // string |  (optional)
	searchDescriptionChanged := true // bool |  (optional)
	searchSourceChanged := true // bool |  (optional)
	searchUploads := "searchUploads_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostVersionsAPI.SearchPostVersions(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchUpdaterName(searchUpdaterName).SearchUpdaterId(searchUpdaterId).SearchPostId(searchPostId).SearchStartId(searchStartId).SearchRating(searchRating).SearchRatingChanged(searchRatingChanged).SearchParentId(searchParentId).SearchParentIdChanged(searchParentIdChanged).SearchTags(searchTags).SearchTagsRemoved(searchTagsRemoved).SearchTagsAdded(searchTagsAdded).SearchLockedTags(searchLockedTags).SearchLockedTagsRemoved(searchLockedTagsRemoved).SearchLockedTagsAdded(searchLockedTagsAdded).SearchReason(searchReason).SearchDescription(searchDescription).SearchDescriptionChanged(searchDescriptionChanged).SearchSourceChanged(searchSourceChanged).SearchUploads(searchUploads).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostVersionsAPI.SearchPostVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPostVersions`: []PostVersion
	fmt.Fprintf(os.Stdout, "Response from `PostVersionsAPI.SearchPostVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPostVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. Note that for post versions specifically, you can only go through the 10,000 most recent results with page numbers. | 
 **searchId** | **int32** | Search for a specific id. | 
 **searchUpdaterName** | **string** |  | 
 **searchUpdaterId** | **float32** |  | 
 **searchPostId** | **float32** |  | 
 **searchStartId** | **float32** |  | 
 **searchRating** | [**Ratings**](Ratings.md) |  | 
 **searchRatingChanged** | **string** |  | 
 **searchParentId** | **float32** |  | 
 **searchParentIdChanged** | **bool** |  | 
 **searchTags** | **string** |  | 
 **searchTagsRemoved** | **string** |  | 
 **searchTagsAdded** | **string** |  | 
 **searchLockedTags** | **string** |  | 
 **searchLockedTagsRemoved** | **string** |  | 
 **searchLockedTagsAdded** | **string** |  | 
 **searchReason** | **string** |  | 
 **searchDescription** | **string** |  | 
 **searchDescriptionChanged** | **bool** |  | 
 **searchSourceChanged** | **bool** |  | 
 **searchUploads** | **string** |  | 

### Return type

[**[]PostVersion**](PostVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

