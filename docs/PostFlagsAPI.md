# \PostFlagsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePostFlag**](PostFlagsAPI.md#CreatePostFlag) | **Post** /post_flags.json | Create Post Flag
[**GetPostFlag**](PostFlagsAPI.md#GetPostFlag) | **Get** /post_flags/{id}.json | Get Post Flag
[**SearchPostFlags**](PostFlagsAPI.md#SearchPostFlags) | **Get** /post_flags.json | Search Post Flags



## CreatePostFlag

> PostFlag CreatePostFlag(ctx).PostFlagPostId(postFlagPostId).PostFlagReasonName(postFlagReasonName).PostFlagParentId(postFlagParentId).Execute()

Create Post Flag



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
	postFlagPostId := int32(56) // int32 | 
	postFlagReasonName := "postFlagReasonName_example" // string | 
	postFlagParentId := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostFlagsAPI.CreatePostFlag(context.Background()).PostFlagPostId(postFlagPostId).PostFlagReasonName(postFlagReasonName).PostFlagParentId(postFlagParentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostFlagsAPI.CreatePostFlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePostFlag`: PostFlag
	fmt.Fprintf(os.Stdout, "Response from `PostFlagsAPI.CreatePostFlag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePostFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postFlagPostId** | **int32** |  | 
 **postFlagReasonName** | **string** |  | 
 **postFlagParentId** | **int32** |  | 

### Return type

[**PostFlag**](PostFlag.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPostFlag

> PostFlag GetPostFlag(ctx, id).Execute()

Get Post Flag

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
	id := int32(56) // int32 | The ID of the post flag.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostFlagsAPI.GetPostFlag(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostFlagsAPI.GetPostFlag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPostFlag`: PostFlag
	fmt.Fprintf(os.Stdout, "Response from `PostFlagsAPI.GetPostFlag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the post flag. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPostFlagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostFlag**](PostFlag.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPostFlags

> SearchPostFlags200Response SearchPostFlags(ctx).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).SearchReasonMatches(searchReasonMatches).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchPostId(searchPostId).SearchPostTagsMatch(searchPostTagsMatch).SearchType(searchType).SearchIsResolved(searchIsResolved).Execute()

Search Post Flags



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
	searchIpAddr := "searchIpAddr_example" // string | Must be Admin+ to use. See [postgres' documentation](https://www.postgresql.org/docs/9.3/functions-net.html) for information on how this is parsed. Specifically, \"is contained within or equals\" (`<<=`). (optional)
	searchOrder := "searchOrder_example" // string |  (optional)
	searchReasonMatches := "searchReasonMatches_example" // string |  (optional)
	searchCreatorId := int32(56) // int32 |  (optional)
	searchCreatorName := "searchCreatorName_example" // string |  (optional)
	searchPostId := int32(56) // int32 |  (optional)
	searchPostTagsMatch := "searchPostTagsMatch_example" // string |  (optional)
	searchType := "searchType_example" // string |  (optional)
	searchIsResolved := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostFlagsAPI.SearchPostFlags(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).SearchReasonMatches(searchReasonMatches).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchPostId(searchPostId).SearchPostTagsMatch(searchPostTagsMatch).SearchType(searchType).SearchIsResolved(searchIsResolved).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostFlagsAPI.SearchPostFlags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPostFlags`: SearchPostFlags200Response
	fmt.Fprintf(os.Stdout, "Response from `PostFlagsAPI.SearchPostFlags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPostFlagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **int32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **int32** | Search for a specific id. | 
 **searchIpAddr** | **string** | Must be Admin+ to use. See [postgres&#39; documentation](https://www.postgresql.org/docs/9.3/functions-net.html) for information on how this is parsed. Specifically, \&quot;is contained within or equals\&quot; (&#x60;&lt;&lt;&#x3D;&#x60;). | 
 **searchOrder** | **string** |  | 
 **searchReasonMatches** | **string** |  | 
 **searchCreatorId** | **int32** |  | 
 **searchCreatorName** | **string** |  | 
 **searchPostId** | **int32** |  | 
 **searchPostTagsMatch** | **string** |  | 
 **searchType** | **string** |  | 
 **searchIsResolved** | **bool** |  | 

### Return type

[**SearchPostFlags200Response**](SearchPostFlags200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

