# \CommentsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateComment**](CommentsAPI.md#CreateComment) | **Post** /comments.json | Create Comment
[**DeleteComment**](CommentsAPI.md#DeleteComment) | **Delete** /comments/{id}.json | Delete Comment
[**EditComment**](CommentsAPI.md#EditComment) | **Patch** /comments/{id}.json | Edit Comment
[**GetComment**](CommentsAPI.md#GetComment) | **Get** /comments/{id}.json | Get Comment
[**HideComment**](CommentsAPI.md#HideComment) | **Post** /comments/{id}/hide.json | Hide Comment
[**MarkComment**](CommentsAPI.md#MarkComment) | **Post** /comments/{id}/warning.json | Mark Comment
[**SearchComments**](CommentsAPI.md#SearchComments) | **Get** /comments.json | Search Comments
[**UnhideComment**](CommentsAPI.md#UnhideComment) | **Post** /comments/{id}/unhide.json | Unhide Comment



## CreateComment

> Comment CreateComment(ctx).CommentBody(commentBody).CommentPostId(commentPostId).CommentDoNotBumpPost(commentDoNotBumpPost).CommentIsSticky(commentIsSticky).CommentIsHidden(commentIsHidden).Execute()

Create Comment

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
	commentBody := "commentBody_example" // string | 
	commentPostId := float32(8.14) // float32 | 
	commentDoNotBumpPost := true // bool |  (optional)
	commentIsSticky := true // bool | Only usable for Janitor+ (optional)
	commentIsHidden := true // bool | Only usable for Moderator+ (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentsAPI.CreateComment(context.Background()).CommentBody(commentBody).CommentPostId(commentPostId).CommentDoNotBumpPost(commentDoNotBumpPost).CommentIsSticky(commentIsSticky).CommentIsHidden(commentIsHidden).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.CreateComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateComment`: Comment
	fmt.Fprintf(os.Stdout, "Response from `CommentsAPI.CreateComment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commentBody** | **string** |  | 
 **commentPostId** | **float32** |  | 
 **commentDoNotBumpPost** | **bool** |  | 
 **commentIsSticky** | **bool** | Only usable for Janitor+ | 
 **commentIsHidden** | **bool** | Only usable for Moderator+ | 

### Return type

[**Comment**](Comment.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteComment

> DeleteComment(ctx, id).Execute()

Delete Comment



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
	id := float32(8.14) // float32 | The ID of the comment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommentsAPI.DeleteComment(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.DeleteComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommentRequest struct via the builder pattern


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


## EditComment

> EditComment(ctx, id).CommentBody(commentBody).CommentIsSticky(commentIsSticky).CommentIsHidden(commentIsHidden).Execute()

Edit Comment



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
	id := float32(8.14) // float32 | The ID of the comment.
	commentBody := "commentBody_example" // string |  (optional)
	commentIsSticky := true // bool | Only usable for Janitor+ (optional)
	commentIsHidden := true // bool | Only usable for Moderator+ (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommentsAPI.EditComment(context.Background(), id).CommentBody(commentBody).CommentIsSticky(commentIsSticky).CommentIsHidden(commentIsHidden).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.EditComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **commentBody** | **string** |  | 
 **commentIsSticky** | **bool** | Only usable for Janitor+ | 
 **commentIsHidden** | **bool** | Only usable for Moderator+ | 

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


## GetComment

> Comment GetComment(ctx, id).Execute()

Get Comment



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
	id := float32(8.14) // float32 | The ID of the comment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentsAPI.GetComment(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.GetComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComment`: Comment
	fmt.Fprintf(os.Stdout, "Response from `CommentsAPI.GetComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Comment**](Comment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HideComment

> Comment HideComment(ctx, id).Execute()

Hide Comment



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
	id := float32(8.14) // float32 | The ID of the comment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentsAPI.HideComment(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.HideComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HideComment`: Comment
	fmt.Fprintf(os.Stdout, "Response from `CommentsAPI.HideComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHideCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Comment**](Comment.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkComment

> DTextResponse MarkComment(ctx, id).MarkBlipRequest(markBlipRequest).Execute()

Mark Comment



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
	id := float32(8.14) // float32 | The ID of the comment.
	markBlipRequest := *openapiclient.NewMarkBlipRequest("RecordType_example") // MarkBlipRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentsAPI.MarkComment(context.Background(), id).MarkBlipRequest(markBlipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.MarkComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkComment`: DTextResponse
	fmt.Fprintf(os.Stdout, "Response from `CommentsAPI.MarkComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkCommentRequest struct via the builder pattern


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


## SearchComments

> SearchComments200Response SearchComments(ctx).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).GroupBy(groupBy).SearchBodyMatches(searchBodyMatches).SearchPostId(searchPostId).SearchPostTagsMatch(searchPostTagsMatch).SearchPostNoteUpdaterName(searchPostNoteUpdaterName).SearchPostNoteUpdaterId(searchPostNoteUpdaterId).SearchCreatorName(searchCreatorName).SearchCreatorId(searchCreatorId).SearchIsSticky(searchIsSticky).SearchIsHidden(searchIsHidden).SearchDoNotBumpPost(searchDoNotBumpPost).Execute()

Search Comments



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
	groupBy := "groupBy_example" // string |  (optional)
	searchBodyMatches := "searchBodyMatches_example" // string |  (optional)
	searchPostId := float32(8.14) // float32 | Accepts a comma separated list. (optional)
	searchPostTagsMatch := "searchPostTagsMatch_example" // string |  (optional)
	searchPostNoteUpdaterName := "searchPostNoteUpdaterName_example" // string |  (optional)
	searchPostNoteUpdaterId := float32(8.14) // float32 |  (optional)
	searchCreatorName := "searchCreatorName_example" // string |  (optional)
	searchCreatorId := float32(8.14) // float32 |  (optional)
	searchIsSticky := true // bool |  (optional)
	searchIsHidden := true // bool | Only usable by Moderator+ (optional)
	searchDoNotBumpPost := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentsAPI.SearchComments(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).GroupBy(groupBy).SearchBodyMatches(searchBodyMatches).SearchPostId(searchPostId).SearchPostTagsMatch(searchPostTagsMatch).SearchPostNoteUpdaterName(searchPostNoteUpdaterName).SearchPostNoteUpdaterId(searchPostNoteUpdaterId).SearchCreatorName(searchCreatorName).SearchCreatorId(searchCreatorId).SearchIsSticky(searchIsSticky).SearchIsHidden(searchIsHidden).SearchDoNotBumpPost(searchDoNotBumpPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.SearchComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchComments`: SearchComments200Response
	fmt.Fprintf(os.Stdout, "Response from `CommentsAPI.SearchComments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **float32** | Search for a specific id. | 
 **searchIpAddr** | **string** | Must be Admin+ to use. See [postgres&#39; documentation](https://www.postgresql.org/docs/9.3/functions-net.html) for information on how this is parsed. Specifically, \&quot;is contained within or equals\&quot; (&#x60;&lt;&lt;&#x3D;&#x60;). | 
 **searchOrder** | **string** |  | 
 **groupBy** | **string** |  | 
 **searchBodyMatches** | **string** |  | 
 **searchPostId** | **float32** | Accepts a comma separated list. | 
 **searchPostTagsMatch** | **string** |  | 
 **searchPostNoteUpdaterName** | **string** |  | 
 **searchPostNoteUpdaterId** | **float32** |  | 
 **searchCreatorName** | **string** |  | 
 **searchCreatorId** | **float32** |  | 
 **searchIsSticky** | **bool** |  | 
 **searchIsHidden** | **bool** | Only usable by Moderator+ | 
 **searchDoNotBumpPost** | **bool** |  | 

### Return type

[**SearchComments200Response**](SearchComments200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnhideComment

> Comment UnhideComment(ctx, id).Execute()

Unhide Comment



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
	id := float32(8.14) // float32 | The ID of the comment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentsAPI.UnhideComment(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentsAPI.UnhideComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnhideComment`: Comment
	fmt.Fprintf(os.Stdout, "Response from `CommentsAPI.UnhideComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnhideCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Comment**](Comment.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

