# \ForumPostVotesAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateForumPostVote**](ForumPostVotesAPI.md#CreateForumPostVote) | **Post** /forum_posts/{id}/votes.json | Create Forum Post Vote
[**DeleteForumPostVote**](ForumPostVotesAPI.md#DeleteForumPostVote) | **Delete** /forum_posts/{id}/votes.json | Delete Forum Post Vote



## CreateForumPostVote

> ForumPostVote CreateForumPostVote(ctx, id).ForumPostVoteScore(forumPostVoteScore).Execute()

Create Forum Post Vote

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
	id := "id_example" // string | The ID of the forum post.
	forumPostVoteScore := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ForumPostVotesAPI.CreateForumPostVote(context.Background(), id).ForumPostVoteScore(forumPostVoteScore).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForumPostVotesAPI.CreateForumPostVote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateForumPostVote`: ForumPostVote
	fmt.Fprintf(os.Stdout, "Response from `ForumPostVotesAPI.CreateForumPostVote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the forum post. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateForumPostVoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forumPostVoteScore** | **float32** |  | 

### Return type

[**ForumPostVote**](ForumPostVote.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteForumPostVote

> DeleteForumPostVote(ctx, id).Execute()

Delete Forum Post Vote

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
	id := "id_example" // string | The ID of the forum post.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ForumPostVotesAPI.DeleteForumPostVote(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForumPostVotesAPI.DeleteForumPostVote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the forum post. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteForumPostVoteRequest struct via the builder pattern


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

