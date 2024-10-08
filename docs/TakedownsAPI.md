# \TakedownsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPostsToTakedownByIds**](TakedownsAPI.md#AddPostsToTakedownByIds) | **Post** /takedowns/{id}/add_by_ids.json | Add Posts To Takedown By IDs
[**AddPostsToTakedownByTags**](TakedownsAPI.md#AddPostsToTakedownByTags) | **Post** /takedowns/{id}/add_by_tags.json | Add Posts To Takedown By Tags
[**CountMatchingPosts**](TakedownsAPI.md#CountMatchingPosts) | **Post** /takedowns/{id}/count_matching_posts.json | Count Matching Posts
[**CreateTakedown**](TakedownsAPI.md#CreateTakedown) | **Post** /takedowns.json | Create Takedown
[**DeleteTakedown**](TakedownsAPI.md#DeleteTakedown) | **Delete** /takedowns/{id}.json | Delete Takedown
[**EditTakedown**](TakedownsAPI.md#EditTakedown) | **Patch** /takedowns/{id}.json | Edit Takedown
[**GetTakedown**](TakedownsAPI.md#GetTakedown) | **Get** /takedowns/{id}.json | Get Takedown
[**RemovePostsFromTakedownByIds**](TakedownsAPI.md#RemovePostsFromTakedownByIds) | **Post** /takedowns/{id}/remove_by_ids.json | Remove Posts From Takedown By IDs
[**SearchTakedowns**](TakedownsAPI.md#SearchTakedowns) | **Get** /takedowns.json | Search Takedowns



## AddPostsToTakedownByIds

> AddPostsToTakedownByIds200Response AddPostsToTakedownByIds(ctx, id).PostIds(postIds).Execute()

Add Posts To Takedown By IDs



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
	id := float32(8.14) // float32 | The ID of the takedown.
	postIds := "postIds_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TakedownsAPI.AddPostsToTakedownByIds(context.Background(), id).PostIds(postIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TakedownsAPI.AddPostsToTakedownByIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPostsToTakedownByIds`: AddPostsToTakedownByIds200Response
	fmt.Fprintf(os.Stdout, "Response from `TakedownsAPI.AddPostsToTakedownByIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the takedown. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPostsToTakedownByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postIds** | **string** |  | 

### Return type

[**AddPostsToTakedownByIds200Response**](AddPostsToTakedownByIds200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddPostsToTakedownByTags

> AddPostsToTakedownByIds200Response AddPostsToTakedownByTags(ctx, id).PostTags(postTags).Execute()

Add Posts To Takedown By Tags



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
	id := float32(8.14) // float32 | The ID of the takedown.
	postTags := "postTags_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TakedownsAPI.AddPostsToTakedownByTags(context.Background(), id).PostTags(postTags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TakedownsAPI.AddPostsToTakedownByTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPostsToTakedownByTags`: AddPostsToTakedownByIds200Response
	fmt.Fprintf(os.Stdout, "Response from `TakedownsAPI.AddPostsToTakedownByTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the takedown. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPostsToTakedownByTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postTags** | **string** |  | 

### Return type

[**AddPostsToTakedownByIds200Response**](AddPostsToTakedownByIds200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountMatchingPosts

> CountMatchingPosts200Response CountMatchingPosts(ctx, id).PostTags(postTags).Execute()

Count Matching Posts



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
	id := float32(8.14) // float32 | The ID of the takedown.
	postTags := "postTags_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TakedownsAPI.CountMatchingPosts(context.Background(), id).PostTags(postTags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TakedownsAPI.CountMatchingPosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountMatchingPosts`: CountMatchingPosts200Response
	fmt.Fprintf(os.Stdout, "Response from `TakedownsAPI.CountMatchingPosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the takedown. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountMatchingPostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postTags** | **string** |  | 

### Return type

[**CountMatchingPosts200Response**](CountMatchingPosts200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTakedown

> Takedown CreateTakedown(ctx).TakedownEmail(takedownEmail).TakedownReason(takedownReason).TakedownSource(takedownSource).TakedownInstructions(takedownInstructions).TakedownPostIds(takedownPostIds).TakedownReasonHidden(takedownReasonHidden).TakedownNotes(takedownNotes).TakedownDelPostIds(takedownDelPostIds).TakedownStatus(takedownStatus).Execute()

Create Takedown

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
	takedownEmail := "takedownEmail_example" // string | 
	takedownReason := "takedownReason_example" // string | 
	takedownSource := "takedownSource_example" // string |  (optional)
	takedownInstructions := "takedownInstructions_example" // string |  (optional)
	takedownPostIds := []float32{float32(123)} // []float32 | takedown[post_ids][]=1&takedown[post_ids][]=2 (optional)
	takedownReasonHidden := true // bool |  (optional)
	takedownNotes := "takedownNotes_example" // string | Must have the bd staff user flag to use. (optional)
	takedownDelPostIds := []float32{float32(123)} // []float32 | Must have the bd staff user flag to use. takedown[del_post_ids][]=1&takedown[del_post_ids][]=2  (optional)
	takedownStatus := "takedownStatus_example" // string | Must have the bd staff user flag to use. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TakedownsAPI.CreateTakedown(context.Background()).TakedownEmail(takedownEmail).TakedownReason(takedownReason).TakedownSource(takedownSource).TakedownInstructions(takedownInstructions).TakedownPostIds(takedownPostIds).TakedownReasonHidden(takedownReasonHidden).TakedownNotes(takedownNotes).TakedownDelPostIds(takedownDelPostIds).TakedownStatus(takedownStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TakedownsAPI.CreateTakedown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTakedown`: Takedown
	fmt.Fprintf(os.Stdout, "Response from `TakedownsAPI.CreateTakedown`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTakedownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **takedownEmail** | **string** |  | 
 **takedownReason** | **string** |  | 
 **takedownSource** | **string** |  | 
 **takedownInstructions** | **string** |  | 
 **takedownPostIds** | **[]float32** | takedown[post_ids][]&#x3D;1&amp;takedown[post_ids][]&#x3D;2 | 
 **takedownReasonHidden** | **bool** |  | 
 **takedownNotes** | **string** | Must have the bd staff user flag to use. | 
 **takedownDelPostIds** | **[]float32** | Must have the bd staff user flag to use. takedown[del_post_ids][]&#x3D;1&amp;takedown[del_post_ids][]&#x3D;2  | 
 **takedownStatus** | **string** | Must have the bd staff user flag to use. | 

### Return type

[**Takedown**](Takedown.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTakedown

> DeleteTakedown(ctx, id).Execute()

Delete Takedown



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
	id := float32(8.14) // float32 | The ID of the takedown.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TakedownsAPI.DeleteTakedown(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TakedownsAPI.DeleteTakedown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the takedown. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTakedownRequest struct via the builder pattern


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


## EditTakedown

> EditTakedown(ctx, id).TakedownNotes(takedownNotes).TakedownReasonHidden(takedownReasonHidden).TakedownPosts(takedownPosts).ProcessTakedown(processTakedown).DeleteReason(deleteReason).Execute()

Edit Takedown



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
	id := float32(8.14) // float32 | The ID of the takedown.
	takedownNotes := "takedownNotes_example" // string |  (optional)
	takedownReasonHidden := true // bool |  (optional)
	takedownPosts := "takedownPosts_example" // string |  (optional)
	processTakedown := true // bool | If not truthy, the takedown will be denied. (optional)
	deleteReason := "deleteReason_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TakedownsAPI.EditTakedown(context.Background(), id).TakedownNotes(takedownNotes).TakedownReasonHidden(takedownReasonHidden).TakedownPosts(takedownPosts).ProcessTakedown(processTakedown).DeleteReason(deleteReason).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TakedownsAPI.EditTakedown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the takedown. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditTakedownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **takedownNotes** | **string** |  | 
 **takedownReasonHidden** | **bool** |  | 
 **takedownPosts** | **string** |  | 
 **processTakedown** | **bool** | If not truthy, the takedown will be denied. | 
 **deleteReason** | **string** |  | 

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


## GetTakedown

> Takedown GetTakedown(ctx, id).Execute()

Get Takedown

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
	id := float32(8.14) // float32 | The ID of the takedown.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TakedownsAPI.GetTakedown(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TakedownsAPI.GetTakedown``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTakedown`: Takedown
	fmt.Fprintf(os.Stdout, "Response from `TakedownsAPI.GetTakedown`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the takedown. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTakedownRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Takedown**](Takedown.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePostsFromTakedownByIds

> RemovePostsFromTakedownByIds(ctx, id).PostIds(postIds).Execute()

Remove Posts From Takedown By IDs



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
	id := float32(8.14) // float32 | The ID of the takedown.
	postIds := "postIds_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TakedownsAPI.RemovePostsFromTakedownByIds(context.Background(), id).PostIds(postIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TakedownsAPI.RemovePostsFromTakedownByIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the takedown. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePostsFromTakedownByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postIds** | **string** |  | 

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


## SearchTakedowns

> SearchTakedowns200Response SearchTakedowns(ctx).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).SearchStatus(searchStatus).SearchSource(searchSource).SearchReason(searchReason).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchReasonHidden(searchReasonHidden).SearchInstructions(searchInstructions).SearchPostId(searchPostId).SearchNotes(searchNotes).SearchEmail(searchEmail).SearchVericode(searchVericode).Execute()

Search Takedowns



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
	searchStatus := "searchStatus_example" // string |  (optional)
	searchSource := "searchSource_example" // string | Must be Moderator+ to use. (optional)
	searchReason := "searchReason_example" // string | Must be Moderator+ to use. (optional)
	searchCreatorId := float32(8.14) // float32 | Must be Moderator+ to use. (optional)
	searchCreatorName := "searchCreatorName_example" // string | Must be Moderator+ to use. (optional)
	searchReasonHidden := true // bool | Must be Moderator+ to use. (optional)
	searchInstructions := "searchInstructions_example" // string | Must be Moderator+ to use. (optional)
	searchPostId := float32(8.14) // float32 | Must be Moderator+ to use. (optional)
	searchNotes := "searchNotes_example" // string | Must be Moderator+ to use. (optional)
	searchEmail := "searchEmail_example" // string | Must be Admin+ to use. (optional)
	searchVericode := "searchVericode_example" // string | Must be Admin+ to use. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TakedownsAPI.SearchTakedowns(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).SearchStatus(searchStatus).SearchSource(searchSource).SearchReason(searchReason).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchReasonHidden(searchReasonHidden).SearchInstructions(searchInstructions).SearchPostId(searchPostId).SearchNotes(searchNotes).SearchEmail(searchEmail).SearchVericode(searchVericode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TakedownsAPI.SearchTakedowns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchTakedowns`: SearchTakedowns200Response
	fmt.Fprintf(os.Stdout, "Response from `TakedownsAPI.SearchTakedowns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchTakedownsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **float32** | Search for a specific id. | 
 **searchIpAddr** | **string** | Must be Admin+ to use. See [postgres&#39; documentation](https://www.postgresql.org/docs/9.3/functions-net.html) for information on how this is parsed. Specifically, \&quot;is contained within or equals\&quot; (&#x60;&lt;&lt;&#x3D;&#x60;). | 
 **searchOrder** | **string** |  | 
 **searchStatus** | **string** |  | 
 **searchSource** | **string** | Must be Moderator+ to use. | 
 **searchReason** | **string** | Must be Moderator+ to use. | 
 **searchCreatorId** | **float32** | Must be Moderator+ to use. | 
 **searchCreatorName** | **string** | Must be Moderator+ to use. | 
 **searchReasonHidden** | **bool** | Must be Moderator+ to use. | 
 **searchInstructions** | **string** | Must be Moderator+ to use. | 
 **searchPostId** | **float32** | Must be Moderator+ to use. | 
 **searchNotes** | **string** | Must be Moderator+ to use. | 
 **searchEmail** | **string** | Must be Admin+ to use. | 
 **searchVericode** | **string** | Must be Admin+ to use. | 

### Return type

[**SearchTakedowns200Response**](SearchTakedowns200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

