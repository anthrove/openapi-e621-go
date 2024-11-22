# \UploadsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchUploads**](UploadsAPI.md#SearchUploads) | **Get** /uploads.json | Search Uploads
[**UploadPost**](UploadsAPI.md#UploadPost) | **Post** /uploads.json | Upload Post



## SearchUploads

> SearchUploads200Response SearchUploads(ctx).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchUploaderId(searchUploaderId).SearchUploaderName(searchUploaderName).SearchSource(searchSource).SearchSourceMatches(searchSourceMatches).SearchRating(searchRating).SearchParentId(searchParentId).SearchPostId(searchPostId).SearchHasPost(searchHasPost).SearchPostTagsMatch(searchPostTagsMatch).SearchStatus(searchStatus).SearchBacktrace(searchBacktrace).SearchTagString(searchTagString).Execute()

Search Uploads



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
	searchOrder := "searchOrder_example" // string |  (optional)
	searchUploaderId := int32(56) // int32 |  (optional)
	searchUploaderName := "searchUploaderName_example" // string |  (optional)
	searchSource := "searchSource_example" // string |  (optional)
	searchSourceMatches := "searchSourceMatches_example" // string |  (optional)
	searchRating := openapiclient.Ratings("s") // Ratings |  (optional)
	searchParentId := int32(56) // int32 |  (optional)
	searchPostId := int32(56) // int32 |  (optional)
	searchHasPost := true // bool |  (optional)
	searchPostTagsMatch := "searchPostTagsMatch_example" // string |  (optional)
	searchStatus := "searchStatus_example" // string |  (optional)
	searchBacktrace := "searchBacktrace_example" // string |  (optional)
	searchTagString := "searchTagString_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadsAPI.SearchUploads(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchUploaderId(searchUploaderId).SearchUploaderName(searchUploaderName).SearchSource(searchSource).SearchSourceMatches(searchSourceMatches).SearchRating(searchRating).SearchParentId(searchParentId).SearchPostId(searchPostId).SearchHasPost(searchHasPost).SearchPostTagsMatch(searchPostTagsMatch).SearchStatus(searchStatus).SearchBacktrace(searchBacktrace).SearchTagString(searchTagString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadsAPI.SearchUploads``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUploads`: SearchUploads200Response
	fmt.Fprintf(os.Stdout, "Response from `UploadsAPI.SearchUploads`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUploadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **int32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **int32** | Search for a specific id. | 
 **searchOrder** | **string** |  | 
 **searchUploaderId** | **int32** |  | 
 **searchUploaderName** | **string** |  | 
 **searchSource** | **string** |  | 
 **searchSourceMatches** | **string** |  | 
 **searchRating** | [**Ratings**](Ratings.md) |  | 
 **searchParentId** | **int32** |  | 
 **searchPostId** | **int32** |  | 
 **searchHasPost** | **bool** |  | 
 **searchPostTagsMatch** | **string** |  | 
 **searchStatus** | **string** |  | 
 **searchBacktrace** | **string** |  | 
 **searchTagString** | **string** |  | 

### Return type

[**SearchUploads200Response**](SearchUploads200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadPost

> UploadPost200Response UploadPost(ctx).UploadTagString(uploadTagString).UploadRating(uploadRating).UploadFile(uploadFile).UploadDirectUrl(uploadDirectUrl).UploadSource(uploadSource).UploadParentId(uploadParentId).UploadDescription(uploadDescription).UploadAsPending(uploadAsPending).UploadLockedRating(uploadLockedRating).UploadLockedTags(uploadLockedTags).Execute()

Upload Post

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
	uploadTagString := "uploadTagString_example" // string | 
	uploadRating := openapiclient.Ratings("s") // Ratings | 
	uploadFile := os.NewFile(1234, "some_file") // *os.File | Mutually exclusive with direct_url. (optional)
	uploadDirectUrl := "uploadDirectUrl_example" // string | Mutually exclusive with file. (optional)
	uploadSource := "uploadSource_example" // string |  (optional)
	uploadParentId := int32(56) // int32 |  (optional)
	uploadDescription := "uploadDescription_example" // string |  (optional)
	uploadAsPending := true // bool | Must have the \\\"Unrestricted Uploads\\\" permission. (optional)
	uploadLockedRating := true // bool | Must be Privileged+ to use. (optional)
	uploadLockedTags := "uploadLockedTags_example" // string | Must be Admin+ to use. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadsAPI.UploadPost(context.Background()).UploadTagString(uploadTagString).UploadRating(uploadRating).UploadFile(uploadFile).UploadDirectUrl(uploadDirectUrl).UploadSource(uploadSource).UploadParentId(uploadParentId).UploadDescription(uploadDescription).UploadAsPending(uploadAsPending).UploadLockedRating(uploadLockedRating).UploadLockedTags(uploadLockedTags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadsAPI.UploadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadPost`: UploadPost200Response
	fmt.Fprintf(os.Stdout, "Response from `UploadsAPI.UploadPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadTagString** | **string** |  | 
 **uploadRating** | [**Ratings**](Ratings.md) |  | 
 **uploadFile** | ***os.File** | Mutually exclusive with direct_url. | 
 **uploadDirectUrl** | **string** | Mutually exclusive with file. | 
 **uploadSource** | **string** |  | 
 **uploadParentId** | **int32** |  | 
 **uploadDescription** | **string** |  | 
 **uploadAsPending** | **bool** | Must have the \\\&quot;Unrestricted Uploads\\\&quot; permission. | 
 **uploadLockedRating** | **bool** | Must be Privileged+ to use. | 
 **uploadLockedTags** | **string** | Must be Admin+ to use. | 

### Return type

[**UploadPost200Response**](UploadPost200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

