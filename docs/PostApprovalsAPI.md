# \PostApprovalsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApprovePost**](PostApprovalsAPI.md#ApprovePost) | **Post** /moderator/post/approval.json | Approve Post
[**SearchPostApprovals**](PostApprovalsAPI.md#SearchPostApprovals) | **Get** /post_approvals.json | Search Post Approvals
[**UnapprovePost**](PostApprovalsAPI.md#UnapprovePost) | **Delete** /moderator/post/approval.json | Unapprove Post



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
	r, err := apiClient.PostApprovalsAPI.ApprovePost(context.Background()).PostId(postId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostApprovalsAPI.ApprovePost``: %v\n", err)
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


## SearchPostApprovals

> SearchPostApprovals200Response SearchPostApprovals(ctx).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchPostTagsMatch(searchPostTagsMatch).SearchUserId(searchUserId).SearchUserName(searchUserName).SearchPostId(searchPostId).Execute()

Search Post Approvals



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
	searchPostTagsMatch := "searchPostTagsMatch_example" // string |  (optional)
	searchUserId := float32(8.14) // float32 |  (optional)
	searchUserName := "searchUserName_example" // string |  (optional)
	searchPostId := float32(8.14) // float32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostApprovalsAPI.SearchPostApprovals(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchPostTagsMatch(searchPostTagsMatch).SearchUserId(searchUserId).SearchUserName(searchUserName).SearchPostId(searchPostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostApprovalsAPI.SearchPostApprovals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPostApprovals`: SearchPostApprovals200Response
	fmt.Fprintf(os.Stdout, "Response from `PostApprovalsAPI.SearchPostApprovals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPostApprovalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **int32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **int32** | Search for a specific id. | 
 **searchOrder** | **string** |  | 
 **searchPostTagsMatch** | **string** |  | 
 **searchUserId** | **float32** |  | 
 **searchUserName** | **string** |  | 
 **searchPostId** | **float32** |  | 

### Return type

[**SearchPostApprovals200Response**](SearchPostApprovals200Response.md)

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
	r, err := apiClient.PostApprovalsAPI.UnapprovePost(context.Background()).PostId(postId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostApprovalsAPI.UnapprovePost``: %v\n", err)
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

