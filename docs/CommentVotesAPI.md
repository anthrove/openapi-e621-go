# \CommentVotesAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCommentVote**](CommentVotesAPI.md#CreateCommentVote) | **Post** /comments/{id}/votes.json | Create Comment Vote
[**DeleteCommentVote**](CommentVotesAPI.md#DeleteCommentVote) | **Delete** /comments/{id}/votes.json | Delete Comment Vote
[**DeleteCommentVotes**](CommentVotesAPI.md#DeleteCommentVotes) | **Post** /comment_votes/delete.json | Delete Comment Vote
[**LockCommentVotes**](CommentVotesAPI.md#LockCommentVotes) | **Post** /comment_votes/lock.json | Lock Comment Votes



## CreateCommentVote

> CreateCommentVote200Response CreateCommentVote(ctx, id).Score(score).NoUnvote(noUnvote).Execute()

Create Comment Vote

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
	id := "id_example" // string | The ID of the comment.
	score := float32(8.14) // float32 | 
	noUnvote := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentVotesAPI.CreateCommentVote(context.Background(), id).Score(score).NoUnvote(noUnvote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentVotesAPI.CreateCommentVote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCommentVote`: CreateCommentVote200Response
	fmt.Fprintf(os.Stdout, "Response from `CommentVotesAPI.CreateCommentVote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCommentVoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **score** | **float32** |  | 
 **noUnvote** | **bool** |  | 

### Return type

[**CreateCommentVote200Response**](CreateCommentVote200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommentVote

> DeleteCommentVote(ctx, id).Execute()

Delete Comment Vote

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
	id := "id_example" // string | The ID of the comment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommentVotesAPI.DeleteCommentVote(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentVotesAPI.DeleteCommentVote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommentVoteRequest struct via the builder pattern


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


## DeleteCommentVotes

> DeleteCommentVotes(ctx).Ids(ids).Execute()

Delete Comment Vote



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
	ids := "ids_example" // string | The IDs of the comment votes, comma separated.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommentVotesAPI.DeleteCommentVotes(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentVotesAPI.DeleteCommentVotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommentVotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | The IDs of the comment votes, comma separated. | 

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


## LockCommentVotes

> LockCommentVotes(ctx).Ids(ids).Execute()

Lock Comment Votes



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
	ids := "ids_example" // string | The IDs of the comment votes, comma separated.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CommentVotesAPI.LockCommentVotes(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentVotesAPI.LockCommentVotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLockCommentVotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | The IDs of the comment votes, comma separated. | 

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

