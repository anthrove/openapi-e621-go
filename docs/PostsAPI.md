# \PostsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApprovePost**](PostsAPI.md#ApprovePost) | **Post** /moderator/post/approval.json | Approve Post
[**CopyNotesToPost**](PostsAPI.md#CopyNotesToPost) | **Put** /posts/{id}/copy_notes.json | Copy Notes To Post
[**DeletePost**](PostsAPI.md#DeletePost) | **Post** /moderator/post/posts/{id}/delete.json | Delete Post
[**EditPost**](PostsAPI.md#EditPost) | **Patch** /posts/{id}.json | Edit Post
[**ExpungePost**](PostsAPI.md#ExpungePost) | **Post** /moderator/post/posts/{id}/expunge.json | Expunge Post
[**GetPost**](PostsAPI.md#GetPost) | **Get** /posts/{id}.json | Get Post
[**GetPostInSequence**](PostsAPI.md#GetPostInSequence) | **Get** /posts/{id}/show_seq.json | Get Post In Sequence
[**GetRandomPost**](PostsAPI.md#GetRandomPost) | **Get** /posts/random.json | Get Random Post
[**MarkPostAsTranslated**](PostsAPI.md#MarkPostAsTranslated) | **Post** /posts/{id}/mark_as_translated.json | Mark Post As Translated
[**MovePostFavorites**](PostsAPI.md#MovePostFavorites) | **Post** /moderator/post/posts/{id}/move_favorites.json | Move Post Favorites
[**RegeneratePostThumbnails**](PostsAPI.md#RegeneratePostThumbnails) | **Post** /moderator/post/posts/{id}/regenerate_thumbnails.json | Regenerate Post Thumbnails
[**RegeneratePostVideos**](PostsAPI.md#RegeneratePostVideos) | **Post** /moderator/post/posts/{id}/regenerate_videos.json | Regenerate Post Videos
[**RevertPost**](PostsAPI.md#RevertPost) | **Post** /posts/{id}/revert.json | Revert Post
[**SearchPosts**](PostsAPI.md#SearchPosts) | **Get** /posts.json | Search Posts
[**UnapprovePost**](PostsAPI.md#UnapprovePost) | **Delete** /moderator/post/approval.json | Unapprove Post
[**UndeletePost**](PostsAPI.md#UndeletePost) | **Post** /moderator/post/posts/{id}/undelete.json | Undelete Post
[**UnflagPost**](PostsAPI.md#UnflagPost) | **Delete** /posts/{id}/flag.json | Unflag Post
[**UpdatePostIqdb**](PostsAPI.md#UpdatePostIqdb) | **Get** /posts/{id}/update_iqdb.json | Update Post IQDB
[**UploadPost**](PostsAPI.md#UploadPost) | **Post** /uploads.json | Upload Post



## ApprovePost

> ApprovePost(ctx).PostId(postId).Execute()

Approve Post



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
	postId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PostsAPI.ApprovePost(context.Background()).PostId(postId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.ApprovePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApprovePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postId** | **float32** |  | 

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


## CopyNotesToPost

> CopyNotesToPost(ctx, id).CopyNotesToPostRequest(copyNotesToPostRequest).Execute()

Copy Notes To Post

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
	id := float32(8.14) // float32 | The ID of the post.
	copyNotesToPostRequest := *openapiclient.NewCopyNotesToPostRequest(float32(123)) // CopyNotesToPostRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PostsAPI.CopyNotesToPost(context.Background(), id).CopyNotesToPostRequest(copyNotesToPostRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.CopyNotesToPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the post. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyNotesToPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **copyNotesToPostRequest** | [**CopyNotesToPostRequest**](CopyNotesToPostRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePost

> DeletePost(ctx, id).Reason(reason).MoveFavorites(moveFavorites).CopySources(copySources).CopyTags(copyTags).Commit(commit).Execute()

Delete Post



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
	id := float32(8.14) // float32 | The ID of the post.
	reason := "reason_example" // string | If the post does not have an active flag, this is required. (optional)
	moveFavorites := true // bool | Move favorites to parent. (optional)
	copySources := true // bool | Copy sources to parent. (optional)
	copyTags := true // bool | Copy tags to parent. (optional)
	commit := "commit_example" // string | If not set, nothing will happen. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PostsAPI.DeletePost(context.Background(), id).Reason(reason).MoveFavorites(moveFavorites).CopySources(copySources).CopyTags(copyTags).Commit(commit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.DeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the post. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | **string** | If the post does not have an active flag, this is required. | 
 **moveFavorites** | **bool** | Move favorites to parent. | 
 **copySources** | **bool** | Copy sources to parent. | 
 **copyTags** | **bool** | Copy tags to parent. | 
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


## EditPost

> Post EditPost(ctx, id).PostTagString(postTagString).PostOldTagString(postOldTagString).PostTagStringDiff(postTagStringDiff).PostSourceDiff(postSourceDiff).PostSource(postSource).PostOldSource(postOldSource).PostParentId(postParentId).PostOldParentId(postOldParentId).PostDescription(postDescription).PostOldDescription(postOldDescription).PostRating(postRating).PostOldRating(postOldRating).PostEditReason(postEditReason).PostIsRatingLocked(postIsRatingLocked).PostIsNoteLocked(postIsNoteLocked).PostBgColor(postBgColor).PostIsCommentLocked(postIsCommentLocked).PostIsStatusLocked(postIsStatusLocked).PostLockedTags(postLockedTags).PostHideFromAnonymous(postHideFromAnonymous).PostHideFromSearchEngines(postHideFromSearchEngines).Execute()

Edit Post



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
	id := float32(8.14) // float32 | The ID of the post.
	postTagString := "postTagString_example" // string | Replaces all tags on the post. (optional)
	postOldTagString := "postOldTagString_example" // string | The tag string before your edits, used to reconcile conflicts. (optional)
	postTagStringDiff := "postTagStringDiff_example" // string | Tags with a minus are removed, else they are added. Mutually exclusive with tag_string. (optional)
	postSourceDiff := "postSourceDiff_example" // string | Sources with a minus are removed, else they are added. It is not possible to add inactive sources through this. Mutually exclusive with source. (optional)
	postSource := "postSource_example" // string | Replaces all sources on the post. (optional)
	postOldSource := "postOldSource_example" // string | The sources before your edits, used to reconcile conflicts. (optional)
	postParentId := float32(8.14) // float32 |  (optional)
	postOldParentId := float32(8.14) // float32 |  (optional)
	postDescription := "postDescription_example" // string |  (optional)
	postOldDescription := "postOldDescription_example" // string |  (optional)
	postRating := openapiclient.Ratings("s") // Ratings |  (optional)
	postOldRating := openapiclient.Ratings("s") // Ratings |  (optional)
	postEditReason := "postEditReason_example" // string |  (optional)
	postIsRatingLocked := true // bool | You must be Privileged+. (optional)
	postIsNoteLocked := true // bool | You must be Janitor+. (optional)
	postBgColor := "postBgColor_example" // string | You must be Janitor+. (optional)
	postIsCommentLocked := true // bool | You must be Admin+. (optional)
	postIsStatusLocked := true // bool | You must be Admin+. (optional)
	postLockedTags := "postLockedTags_example" // string | You must be Admin+. (optional)
	postHideFromAnonymous := true // bool | You must be Admin+. (optional)
	postHideFromSearchEngines := true // bool | You must be Admin+. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.EditPost(context.Background(), id).PostTagString(postTagString).PostOldTagString(postOldTagString).PostTagStringDiff(postTagStringDiff).PostSourceDiff(postSourceDiff).PostSource(postSource).PostOldSource(postOldSource).PostParentId(postParentId).PostOldParentId(postOldParentId).PostDescription(postDescription).PostOldDescription(postOldDescription).PostRating(postRating).PostOldRating(postOldRating).PostEditReason(postEditReason).PostIsRatingLocked(postIsRatingLocked).PostIsNoteLocked(postIsNoteLocked).PostBgColor(postBgColor).PostIsCommentLocked(postIsCommentLocked).PostIsStatusLocked(postIsStatusLocked).PostLockedTags(postLockedTags).PostHideFromAnonymous(postHideFromAnonymous).PostHideFromSearchEngines(postHideFromSearchEngines).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.EditPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditPost`: Post
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.EditPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the post. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postTagString** | **string** | Replaces all tags on the post. | 
 **postOldTagString** | **string** | The tag string before your edits, used to reconcile conflicts. | 
 **postTagStringDiff** | **string** | Tags with a minus are removed, else they are added. Mutually exclusive with tag_string. | 
 **postSourceDiff** | **string** | Sources with a minus are removed, else they are added. It is not possible to add inactive sources through this. Mutually exclusive with source. | 
 **postSource** | **string** | Replaces all sources on the post. | 
 **postOldSource** | **string** | The sources before your edits, used to reconcile conflicts. | 
 **postParentId** | **float32** |  | 
 **postOldParentId** | **float32** |  | 
 **postDescription** | **string** |  | 
 **postOldDescription** | **string** |  | 
 **postRating** | [**Ratings**](Ratings.md) |  | 
 **postOldRating** | [**Ratings**](Ratings.md) |  | 
 **postEditReason** | **string** |  | 
 **postIsRatingLocked** | **bool** | You must be Privileged+. | 
 **postIsNoteLocked** | **bool** | You must be Janitor+. | 
 **postBgColor** | **string** | You must be Janitor+. | 
 **postIsCommentLocked** | **bool** | You must be Admin+. | 
 **postIsStatusLocked** | **bool** | You must be Admin+. | 
 **postLockedTags** | **string** | You must be Admin+. | 
 **postHideFromAnonymous** | **bool** | You must be Admin+. | 
 **postHideFromSearchEngines** | **bool** | You must be Admin+. | 

### Return type

[**Post**](Post.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpungePost

> AddFavorite201Response ExpungePost(ctx, id).Reason(reason).Execute()

Expunge Post



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
	id := float32(8.14) // float32 | The ID of the post.
	reason := "reason_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.ExpungePost(context.Background(), id).Reason(reason).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.ExpungePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExpungePost`: AddFavorite201Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.ExpungePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the post. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpungePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reason** | **string** |  | 

### Return type

[**AddFavorite201Response**](AddFavorite201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPost

> AddFavorite201Response GetPost(ctx, id).Execute()

Get Post

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
	id := float32(8.14) // float32 | The ID of the post.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.GetPost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.GetPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPost`: AddFavorite201Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.GetPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the post. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddFavorite201Response**](AddFavorite201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPostInSequence

> AddFavorite201Response GetPostInSequence(ctx, id).Seq(seq).Execute()

Get Post In Sequence

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
	id := float32(8.14) // float32 | The ID of the post.
	seq := "seq_example" // string | The direction to move in the sequence. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.GetPostInSequence(context.Background(), id).Seq(seq).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.GetPostInSequence``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPostInSequence`: AddFavorite201Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.GetPostInSequence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the post. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPostInSequenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **seq** | **string** | The direction to move in the sequence. | 

### Return type

[**AddFavorite201Response**](AddFavorite201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRandomPost

> AddFavorite201Response GetRandomPost(ctx).Tags(tags).Execute()

Get Random Post

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
	tags := "tags_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.GetRandomPost(context.Background()).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.GetRandomPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRandomPost`: AddFavorite201Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.GetRandomPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRandomPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tags** | **string** |  | 

### Return type

[**AddFavorite201Response**](AddFavorite201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkPostAsTranslated

> AddFavorite201Response MarkPostAsTranslated(ctx, id).MarkPostAsTranslatedRequest(markPostAsTranslatedRequest).Execute()

Mark Post As Translated



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
	id := float32(8.14) // float32 | The ID of the post.
	markPostAsTranslatedRequest := *openapiclient.NewMarkPostAsTranslatedRequest() // MarkPostAsTranslatedRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.MarkPostAsTranslated(context.Background(), id).MarkPostAsTranslatedRequest(markPostAsTranslatedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.MarkPostAsTranslated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkPostAsTranslated`: AddFavorite201Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.MarkPostAsTranslated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the post. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkPostAsTranslatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **markPostAsTranslatedRequest** | [**MarkPostAsTranslatedRequest**](MarkPostAsTranslatedRequest.md) |  | 

### Return type

[**AddFavorite201Response**](AddFavorite201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MovePostFavorites

> MovePostFavorites(ctx, id).Commit(commit).Execute()

Move Post Favorites



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
	id := float32(8.14) // float32 | The ID of the post.
	commit := "commit_example" // string | If not set, nothing will happen.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PostsAPI.MovePostFavorites(context.Background(), id).Commit(commit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.MovePostFavorites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the post. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMovePostFavoritesRequest struct via the builder pattern


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


## RegeneratePostThumbnails

> AddFavorite201Response RegeneratePostThumbnails(ctx, id).Execute()

Regenerate Post Thumbnails



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
	id := float32(8.14) // float32 | The ID of the post.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.RegeneratePostThumbnails(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.RegeneratePostThumbnails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegeneratePostThumbnails`: AddFavorite201Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.RegeneratePostThumbnails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the post. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegeneratePostThumbnailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddFavorite201Response**](AddFavorite201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegeneratePostVideos

> RegeneratePostVideos(ctx, id).Execute()

Regenerate Post Videos



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
	id := float32(8.14) // float32 | The ID of the post.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PostsAPI.RegeneratePostVideos(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.RegeneratePostVideos``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the post. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegeneratePostVideosRequest struct via the builder pattern


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


## RevertPost

> RevertPost(ctx, id).VersionId(versionId).Execute()

Revert Post

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
	id := float32(8.14) // float32 | The ID of the post.
	versionId := float32(8.14) // float32 | The version ID to revert to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PostsAPI.RevertPost(context.Background(), id).VersionId(versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.RevertPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the post. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevertPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **versionId** | **float32** | The version ID to revert to. | 

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


## SearchPosts

> ListFavorites200Response SearchPosts(ctx).Limit(limit).Page(page).Tags(tags).Md5(md5).Random(random).Execute()

Search Posts

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
	tags := "tags_example" // string |  (optional)
	md5 := "md5_example" // string |  (optional)
	random := "random_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.SearchPosts(context.Background()).Limit(limit).Page(page).Tags(tags).Md5(md5).Random(random).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.SearchPosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPosts`: ListFavorites200Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.SearchPosts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **int32** | The page number of results to get. Between 1 and 750. | 
 **tags** | **string** |  | 
 **md5** | **string** |  | 
 **random** | **string** |  | 

### Return type

[**ListFavorites200Response**](ListFavorites200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnapprovePost

> UnapprovePost(ctx).PostId(postId).Execute()

Unapprove Post



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
	postId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PostsAPI.UnapprovePost(context.Background()).PostId(postId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.UnapprovePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnapprovePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postId** | **float32** |  | 

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


## UndeletePost

> AddFavorite201Response UndeletePost(ctx, id).Execute()

Undelete Post



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
	id := float32(8.14) // float32 | The ID of the post.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.UndeletePost(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.UndeletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UndeletePost`: AddFavorite201Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.UndeletePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the post. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUndeletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddFavorite201Response**](AddFavorite201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnflagPost

> UnflagPost(ctx, id).Approval(approval).Execute()

Unflag Post



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
	id := float32(8.14) // float32 | The ID of the post.
	approval := "approval_example" // string | Approves the post if set to \\\"approve\\\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PostsAPI.UnflagPost(context.Background(), id).Approval(approval).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.UnflagPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the post. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnflagPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **approval** | **string** | Approves the post if set to \\\&quot;approve\\\&quot;. | 

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


## UpdatePostIqdb

> AddFavorite201Response UpdatePostIqdb(ctx, id).Execute()

Update Post IQDB



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
	id := float32(8.14) // float32 | The ID of the post.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.UpdatePostIqdb(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.UpdatePostIqdb``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePostIqdb`: AddFavorite201Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.UpdatePostIqdb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the post. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePostIqdbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AddFavorite201Response**](AddFavorite201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

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
	uploadParentId := float32(8.14) // float32 |  (optional)
	uploadDescription := "uploadDescription_example" // string |  (optional)
	uploadAsPending := true // bool | Must have the \\\"Unrestricted Uploads\\\" permission. (optional)
	uploadLockedRating := true // bool | Must be Privileged+ to use. (optional)
	uploadLockedTags := "uploadLockedTags_example" // string | Must be Admin+ to use. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostsAPI.UploadPost(context.Background()).UploadTagString(uploadTagString).UploadRating(uploadRating).UploadFile(uploadFile).UploadDirectUrl(uploadDirectUrl).UploadSource(uploadSource).UploadParentId(uploadParentId).UploadDescription(uploadDescription).UploadAsPending(uploadAsPending).UploadLockedRating(uploadLockedRating).UploadLockedTags(uploadLockedTags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostsAPI.UploadPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadPost`: UploadPost200Response
	fmt.Fprintf(os.Stdout, "Response from `PostsAPI.UploadPost`: %v\n", resp)
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
 **uploadParentId** | **float32** |  | 
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

