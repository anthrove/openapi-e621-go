# \TagsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CorrectTag**](TagsAPI.md#CorrectTag) | **Post** /tags/{id}/correction.json | Correct Tag
[**EditTag**](TagsAPI.md#EditTag) | **Patch** /tags/{id}.json | Edit Tag
[**GetTag**](TagsAPI.md#GetTag) | **Get** /tags/{id}.json | Get Tag
[**PreviewTags**](TagsAPI.md#PreviewTags) | **Get** /tags/preview.json | Preview Tags
[**SearchTags**](TagsAPI.md#SearchTags) | **Get** /tags.json | Search Tags



## CorrectTag

> CorrectTag(ctx, id).Commit(commit).Execute()

Correct Tag



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
	id := float32(8.14) // float32 | The ID of the tag.
	commit := "commit_example" // string | If not set, nothing will happen.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagsAPI.CorrectTag(context.Background(), id).Commit(commit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.CorrectTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the tag. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCorrectTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commit** | **string** | If not set, nothing will happen. | 

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


## EditTag

> EditTag(ctx, id).TagCategory(tagCategory).TagIsLocked(tagIsLocked).Execute()

Edit Tag



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
	id := float32(8.14) // float32 | The ID of the tag.
	tagCategory := openapiclient.TagCategories(0) // TagCategories |  (optional)
	tagIsLocked := true // bool | Must be Admin+. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagsAPI.EditTag(context.Background(), id).TagCategory(tagCategory).TagIsLocked(tagIsLocked).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.EditTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the tag. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagCategory** | [**TagCategories**](TagCategories.md) |  | 
 **tagIsLocked** | **bool** | Must be Admin+. | 

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


## GetTag

> Tag GetTag(ctx, id).Execute()

Get Tag

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
	id := GetArtistIdOrNameParameter(8.14) // GetArtistIdOrNameParameter | The ID or name of the tag.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.GetTag(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTag`: Tag
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **GetArtistIdOrNameParameter** | The ID or name of the tag. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tag**](Tag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewTags

> []TagPreview PreviewTags(ctx).Tags(tags).Execute()

Preview Tags



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
	tags := "tags_example" // string | The tags to preview, space separated.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.PreviewTags(context.Background()).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.PreviewTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreviewTags`: []TagPreview
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.PreviewTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | **string** | The tags to preview, space separated. | 

### Return type

[**[]TagPreview**](TagPreview.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchTags

> SearchTags200Response SearchTags(ctx).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchFuzzyNameMatches(searchFuzzyNameMatches).SearchNameMatches(searchNameMatches).SearchName(searchName).SearchCategory(searchCategory).SearchHideEmpty(searchHideEmpty).SearchHasWiki(searchHasWiki).SearchHasArtist(searchHasArtist).Execute()

Search Tags



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
	searchFuzzyNameMatches := "searchFuzzyNameMatches_example" // string |  (optional)
	searchNameMatches := "searchNameMatches_example" // string |  (optional)
	searchName := "searchName_example" // string |  (optional)
	searchCategory := openapiclient.TagCategories(0) // TagCategories |  (optional)
	searchHideEmpty := true // bool |  (optional)
	searchHasWiki := true // bool |  (optional)
	searchHasArtist := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.SearchTags(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchFuzzyNameMatches(searchFuzzyNameMatches).SearchNameMatches(searchNameMatches).SearchName(searchName).SearchCategory(searchCategory).SearchHideEmpty(searchHideEmpty).SearchHasWiki(searchHasWiki).SearchHasArtist(searchHasArtist).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.SearchTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchTags`: SearchTags200Response
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.SearchTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **int32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **int32** | Search for a specific id. | 
 **searchOrder** | **string** |  | 
 **searchFuzzyNameMatches** | **string** |  | 
 **searchNameMatches** | **string** |  | 
 **searchName** | **string** |  | 
 **searchCategory** | [**TagCategories**](TagCategories.md) |  | 
 **searchHideEmpty** | **bool** |  | 
 **searchHasWiki** | **bool** |  | 
 **searchHasArtist** | **bool** |  | 

### Return type

[**SearchTags200Response**](SearchTags200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

