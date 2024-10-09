# \ForumPostsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateForumPost**](ForumPostsAPI.md#CreateForumPost) | **Post** /forum_posts.json | Create Forum Post
[**DeleteForumPost**](ForumPostsAPI.md#DeleteForumPost) | **Delete** /forum_posts/{id}.json | Delete Forum Post
[**EditForumPost**](ForumPostsAPI.md#EditForumPost) | **Patch** /forum_posts/{id}.json | Edit Forum Post
[**GetForumPost**](ForumPostsAPI.md#GetForumPost) | **Get** /forum_posts/{id}.json | Get Forum Post
[**HideForumPost**](ForumPostsAPI.md#HideForumPost) | **Post** /forum_posts/{id}/hide.json | Hide Forum Post
[**MarkForumPost**](ForumPostsAPI.md#MarkForumPost) | **Post** /forum_posts/{id}/warning.json | Mark Forum Post
[**SearchForumPosts**](ForumPostsAPI.md#SearchForumPosts) | **Get** /forum_posts.json | Search Forum Posts
[**UnhideForumPost**](ForumPostsAPI.md#UnhideForumPost) | **Post** /forum_posts/{id}/unhide.json | Unhide Forum Post



## CreateForumPost

> ForumPost CreateForumPost(ctx).ForumPostBody(forumPostBody).ForumPostTopicId(forumPostTopicId).Execute()

Create Forum Post

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
	forumPostBody := "forumPostBody_example" // string | 
	forumPostTopicId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ForumPostsAPI.CreateForumPost(context.Background()).ForumPostBody(forumPostBody).ForumPostTopicId(forumPostTopicId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForumPostsAPI.CreateForumPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateForumPost`: ForumPost
	fmt.Fprintf(os.Stdout, "Response from `ForumPostsAPI.CreateForumPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateForumPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forumPostBody** | **string** |  | 
 **forumPostTopicId** | **float32** |  | 

### Return type

[**ForumPost**](ForumPost.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteForumPost

> DeleteForumPost(ctx, id).Execute()

Delete Forum Post



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
	id := float32(8.14) // float32 | The ID of the forum post.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ForumPostsAPI.DeleteForumPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForumPostsAPI.DeleteForumPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the forum post. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteForumPostRequest struct via the builder pattern


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


## EditForumPost

> EditForumPost(ctx, id).ForumPostBody(forumPostBody).Execute()

Edit Forum Post



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
	id := float32(8.14) // float32 | The ID of the forum post.
	forumPostBody := "forumPostBody_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ForumPostsAPI.EditForumPost(context.Background(), id).ForumPostBody(forumPostBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForumPostsAPI.EditForumPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the forum post. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditForumPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forumPostBody** | **string** |  | 

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


## GetForumPost

> ForumPost GetForumPost(ctx, id).Execute()

Get Forum Post



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
	id := float32(8.14) // float32 | The ID of the forum post.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ForumPostsAPI.GetForumPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForumPostsAPI.GetForumPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetForumPost`: ForumPost
	fmt.Fprintf(os.Stdout, "Response from `ForumPostsAPI.GetForumPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the forum post. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetForumPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ForumPost**](ForumPost.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HideForumPost

> ForumPost HideForumPost(ctx, id).Execute()

Hide Forum Post



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
	id := float32(8.14) // float32 | The ID of the forum post.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ForumPostsAPI.HideForumPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForumPostsAPI.HideForumPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HideForumPost`: ForumPost
	fmt.Fprintf(os.Stdout, "Response from `ForumPostsAPI.HideForumPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the forum post. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHideForumPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ForumPost**](ForumPost.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkForumPost

> DTextResponse MarkForumPost(ctx, id).MarkBlipRequest(markBlipRequest).Execute()

Mark Forum Post



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
	id := float32(8.14) // float32 | The ID of the forum post.
	markBlipRequest := *openapiclient.NewMarkBlipRequest("RecordType_example") // MarkBlipRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ForumPostsAPI.MarkForumPost(context.Background(), id).MarkBlipRequest(markBlipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForumPostsAPI.MarkForumPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkForumPost`: DTextResponse
	fmt.Fprintf(os.Stdout, "Response from `ForumPostsAPI.MarkForumPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the forum post. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkForumPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **markBlipRequest** | [**MarkBlipRequest**](MarkBlipRequest.md) |  | 

### Return type

[**DTextResponse**](DTextResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchForumPosts

> SearchForumPosts200Response SearchForumPosts(ctx).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchTopicId(searchTopicId).SearchTopicTitleMatches(searchTopicTitleMatches).SearchBodyMatches(searchBodyMatches).SearchTopicCategoryId(searchTopicCategoryId).SearchIsHidden(searchIsHidden).Execute()

Search Forum Posts



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
	searchCreatorId := float32(8.14) // float32 |  (optional)
	searchCreatorName := "searchCreatorName_example" // string |  (optional)
	searchTopicId := float32(8.14) // float32 |  (optional)
	searchTopicTitleMatches := "searchTopicTitleMatches_example" // string |  (optional)
	searchBodyMatches := "searchBodyMatches_example" // string |  (optional)
	searchTopicCategoryId := float32(8.14) // float32 |  (optional)
	searchIsHidden := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ForumPostsAPI.SearchForumPosts(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchTopicId(searchTopicId).SearchTopicTitleMatches(searchTopicTitleMatches).SearchBodyMatches(searchBodyMatches).SearchTopicCategoryId(searchTopicCategoryId).SearchIsHidden(searchIsHidden).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForumPostsAPI.SearchForumPosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchForumPosts`: SearchForumPosts200Response
	fmt.Fprintf(os.Stdout, "Response from `ForumPostsAPI.SearchForumPosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchForumPostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **int32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **int32** | Search for a specific id. | 
 **searchOrder** | **string** |  | 
 **searchCreatorId** | **float32** |  | 
 **searchCreatorName** | **string** |  | 
 **searchTopicId** | **float32** |  | 
 **searchTopicTitleMatches** | **string** |  | 
 **searchBodyMatches** | **string** |  | 
 **searchTopicCategoryId** | **float32** |  | 
 **searchIsHidden** | **bool** |  | 

### Return type

[**SearchForumPosts200Response**](SearchForumPosts200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnhideForumPost

> ForumPost UnhideForumPost(ctx, id).Execute()

Unhide Forum Post



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
	id := float32(8.14) // float32 | The ID of the forum post.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ForumPostsAPI.UnhideForumPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForumPostsAPI.UnhideForumPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnhideForumPost`: ForumPost
	fmt.Fprintf(os.Stdout, "Response from `ForumPostsAPI.UnhideForumPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the forum post. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnhideForumPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ForumPost**](ForumPost.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

