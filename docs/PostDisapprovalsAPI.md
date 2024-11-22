# \PostDisapprovalsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePostDisapproval**](PostDisapprovalsAPI.md#CreatePostDisapproval) | **Post** /moderator/post/disapprovals.json | Create Post Disapproval
[**SearchPostDisapprovals**](PostDisapprovalsAPI.md#SearchPostDisapprovals) | **Get** /moderator/post/disapprovals.json | Search Post Disapprovals



## CreatePostDisapproval

> PostDisapproval CreatePostDisapproval(ctx).PostDisapprovalPostId(postDisapprovalPostId).PostDisapprovalReason(postDisapprovalReason).PostDisapprovalMessage(postDisapprovalMessage).Execute()

Create Post Disapproval



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
	postDisapprovalPostId := int32(56) // int32 | 
	postDisapprovalReason := "postDisapprovalReason_example" // string | 
	postDisapprovalMessage := "postDisapprovalMessage_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostDisapprovalsAPI.CreatePostDisapproval(context.Background()).PostDisapprovalPostId(postDisapprovalPostId).PostDisapprovalReason(postDisapprovalReason).PostDisapprovalMessage(postDisapprovalMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostDisapprovalsAPI.CreatePostDisapproval``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePostDisapproval`: PostDisapproval
	fmt.Fprintf(os.Stdout, "Response from `PostDisapprovalsAPI.CreatePostDisapproval`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePostDisapprovalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postDisapprovalPostId** | **int32** |  | 
 **postDisapprovalReason** | **string** |  | 
 **postDisapprovalMessage** | **string** |  | 

### Return type

[**PostDisapproval**](PostDisapproval.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPostDisapprovals

> SearchPostDisapprovals200Response SearchPostDisapprovals(ctx).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchPostId(searchPostId).SearchMessage(searchMessage).SearchPostTagsMatch(searchPostTagsMatch).SearchReason(searchReason).SearchHasMessage(searchHasMessage).Execute()

Search Post Disapprovals



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
	searchCreatorId := int32(56) // int32 |  (optional)
	searchCreatorName := "searchCreatorName_example" // string |  (optional)
	searchPostId := int32(56) // int32 |  (optional)
	searchMessage := "searchMessage_example" // string |  (optional)
	searchPostTagsMatch := "searchPostTagsMatch_example" // string |  (optional)
	searchReason := "searchReason_example" // string |  (optional)
	searchHasMessage := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostDisapprovalsAPI.SearchPostDisapprovals(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchPostId(searchPostId).SearchMessage(searchMessage).SearchPostTagsMatch(searchPostTagsMatch).SearchReason(searchReason).SearchHasMessage(searchHasMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostDisapprovalsAPI.SearchPostDisapprovals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPostDisapprovals`: SearchPostDisapprovals200Response
	fmt.Fprintf(os.Stdout, "Response from `PostDisapprovalsAPI.SearchPostDisapprovals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPostDisapprovalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **int32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **int32** | Search for a specific id. | 
 **searchOrder** | **string** |  | 
 **searchCreatorId** | **int32** |  | 
 **searchCreatorName** | **string** |  | 
 **searchPostId** | **int32** |  | 
 **searchMessage** | **string** |  | 
 **searchPostTagsMatch** | **string** |  | 
 **searchReason** | **string** |  | 
 **searchHasMessage** | **bool** |  | 

### Return type

[**SearchPostDisapprovals200Response**](SearchPostDisapprovals200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

