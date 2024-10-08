# \PostSetsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPostsToPostSet**](PostSetsAPI.md#AddPostsToPostSet) | **Post** /post_sets/{id}/add_posts.json | Add Posts To Post Set
[**CreatePostSet**](PostSetsAPI.md#CreatePostSet) | **Post** /post_sets.json | Create Post Set
[**DeletePostSet**](PostSetsAPI.md#DeletePostSet) | **Delete** /post_sets/{id}.json | Delete Post Set
[**EditPostSet**](PostSetsAPI.md#EditPostSet) | **Patch** /post_sets/{id}.json | Edit Post Set
[**GetPostSet**](PostSetsAPI.md#GetPostSet) | **Get** /post_sets/{id}.json | Get Post Set
[**ListPostSetsForSelect**](PostSetsAPI.md#ListPostSetsForSelect) | **Get** /post_sets/for_select.json | List Post Sets For Select
[**RemovePostsFromPostSet**](PostSetsAPI.md#RemovePostsFromPostSet) | **Post** /post_sets/{id}/remove_posts.json | Remove Posts From Post Set
[**SearchPostSets**](PostSetsAPI.md#SearchPostSets) | **Get** /post_sets.json | Search Post Sets
[**UpdatePostSetPosts**](PostSetsAPI.md#UpdatePostSetPosts) | **Post** /post_sets/{id}/update_posts.json | Update Post Set Posts



## AddPostsToPostSet

> PostSet AddPostsToPostSet(ctx, id).PostIds(postIds).Execute()

Add Posts To Post Set



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
	id := float32(8.14) // float32 | The ID of the post set.
	postIds := []float32{float32(123)} // []float32 | post_ids[]=1&post_ids[]=2

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostSetsAPI.AddPostsToPostSet(context.Background(), id).PostIds(postIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostSetsAPI.AddPostsToPostSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddPostsToPostSet`: PostSet
	fmt.Fprintf(os.Stdout, "Response from `PostSetsAPI.AddPostsToPostSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the post set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPostsToPostSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postIds** | **[]float32** | post_ids[]&#x3D;1&amp;post_ids[]&#x3D;2 | 

### Return type

[**PostSet**](PostSet.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePostSet

> PostSet CreatePostSet(ctx).PostSetName(postSetName).PostSetShortname(postSetShortname).PostSetDescription(postSetDescription).PostSetIsPublic(postSetIsPublic).PostSetRansferOnDelete(postSetRansferOnDelete).Execute()

Create Post Set

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
	postSetName := "postSetName_example" // string | 
	postSetShortname := "postSetShortname_example" // string | 
	postSetDescription := "postSetDescription_example" // string |  (optional)
	postSetIsPublic := true // bool |  (optional)
	postSetRansferOnDelete := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostSetsAPI.CreatePostSet(context.Background()).PostSetName(postSetName).PostSetShortname(postSetShortname).PostSetDescription(postSetDescription).PostSetIsPublic(postSetIsPublic).PostSetRansferOnDelete(postSetRansferOnDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostSetsAPI.CreatePostSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePostSet`: PostSet
	fmt.Fprintf(os.Stdout, "Response from `PostSetsAPI.CreatePostSet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePostSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postSetName** | **string** |  | 
 **postSetShortname** | **string** |  | 
 **postSetDescription** | **string** |  | 
 **postSetIsPublic** | **bool** |  | 
 **postSetRansferOnDelete** | **bool** |  | 

### Return type

[**PostSet**](PostSet.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePostSet

> DeletePostSet(ctx, id).Execute()

Delete Post Set



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
	id := float32(8.14) // float32 | The ID of the post set.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PostSetsAPI.DeletePostSet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostSetsAPI.DeletePostSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the post set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePostSetRequest struct via the builder pattern


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


## EditPostSet

> EditPostSet(ctx, id).PostSetName(postSetName).PostSetShortname(postSetShortname).PostSetDescription(postSetDescription).PostSetIsPublic(postSetIsPublic).PostSetTransferOnDelete(postSetTransferOnDelete).Execute()

Edit Post Set



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
	id := float32(8.14) // float32 | The ID of the post sets.
	postSetName := "postSetName_example" // string |  (optional)
	postSetShortname := "postSetShortname_example" // string |  (optional)
	postSetDescription := "postSetDescription_example" // string |  (optional)
	postSetIsPublic := true // bool |  (optional)
	postSetTransferOnDelete := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PostSetsAPI.EditPostSet(context.Background(), id).PostSetName(postSetName).PostSetShortname(postSetShortname).PostSetDescription(postSetDescription).PostSetIsPublic(postSetIsPublic).PostSetTransferOnDelete(postSetTransferOnDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostSetsAPI.EditPostSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the post sets. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditPostSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postSetName** | **string** |  | 
 **postSetShortname** | **string** |  | 
 **postSetDescription** | **string** |  | 
 **postSetIsPublic** | **bool** |  | 
 **postSetTransferOnDelete** | **bool** |  | 

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


## GetPostSet

> PostSet GetPostSet(ctx, id).Execute()

Get Post Set



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
	id := float32(8.14) // float32 | The ID of the post set.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostSetsAPI.GetPostSet(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostSetsAPI.GetPostSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPostSet`: PostSet
	fmt.Fprintf(os.Stdout, "Response from `PostSetsAPI.GetPostSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the post set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPostSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostSet**](PostSet.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPostSetsForSelect

> ListPostSetsForSelect200Response ListPostSetsForSelect(ctx).Execute()

List Post Sets For Select



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostSetsAPI.ListPostSetsForSelect(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostSetsAPI.ListPostSetsForSelect``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPostSetsForSelect`: ListPostSetsForSelect200Response
	fmt.Fprintf(os.Stdout, "Response from `PostSetsAPI.ListPostSetsForSelect`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPostSetsForSelectRequest struct via the builder pattern


### Return type

[**ListPostSetsForSelect200Response**](ListPostSetsForSelect200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePostsFromPostSet

> PostSet RemovePostsFromPostSet(ctx, id).PostIds(postIds).Execute()

Remove Posts From Post Set



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
	id := float32(8.14) // float32 | The ID of the post set.
	postIds := []float32{float32(123)} // []float32 | post_ids[]=1&post_ids[]=2

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostSetsAPI.RemovePostsFromPostSet(context.Background(), id).PostIds(postIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostSetsAPI.RemovePostsFromPostSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemovePostsFromPostSet`: PostSet
	fmt.Fprintf(os.Stdout, "Response from `PostSetsAPI.RemovePostsFromPostSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the post set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePostsFromPostSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postIds** | **[]float32** | post_ids[]&#x3D;1&amp;post_ids[]&#x3D;2 | 

### Return type

[**PostSet**](PostSet.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchPostSets

> SearchPostSets200Response SearchPostSets(ctx).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchName(searchName).SearchShortname(searchShortname).SearchIsPublic(searchIsPublic).SearchPostId(searchPostId).SearchMaintainerId(searchMaintainerId).Execute()

Search Post Sets



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
	searchOrder := "searchOrder_example" // string |  (optional)
	searchCreatorId := float32(8.14) // float32 |  (optional)
	searchCreatorName := "searchCreatorName_example" // string |  (optional)
	searchName := "searchName_example" // string |  (optional)
	searchShortname := "searchShortname_example" // string |  (optional)
	searchIsPublic := true // bool | You must be Moderator+. (optional)
	searchPostId := float32(8.14) // float32 |  (optional)
	searchMaintainerId := float32(8.14) // float32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostSetsAPI.SearchPostSets(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchName(searchName).SearchShortname(searchShortname).SearchIsPublic(searchIsPublic).SearchPostId(searchPostId).SearchMaintainerId(searchMaintainerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostSetsAPI.SearchPostSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPostSets`: SearchPostSets200Response
	fmt.Fprintf(os.Stdout, "Response from `PostSetsAPI.SearchPostSets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPostSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **float32** | Search for a specific id. | 
 **searchOrder** | **string** |  | 
 **searchCreatorId** | **float32** |  | 
 **searchCreatorName** | **string** |  | 
 **searchName** | **string** |  | 
 **searchShortname** | **string** |  | 
 **searchIsPublic** | **bool** | You must be Moderator+. | 
 **searchPostId** | **float32** |  | 
 **searchMaintainerId** | **float32** |  | 

### Return type

[**SearchPostSets200Response**](SearchPostSets200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePostSetPosts

> PostSet UpdatePostSetPosts(ctx, id).PostSetPostIdsString(postSetPostIdsString).Execute()

Update Post Set Posts



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
	id := float32(8.14) // float32 | The ID of the post set.
	postSetPostIdsString := "postSetPostIdsString_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostSetsAPI.UpdatePostSetPosts(context.Background(), id).PostSetPostIdsString(postSetPostIdsString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostSetsAPI.UpdatePostSetPosts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePostSetPosts`: PostSet
	fmt.Fprintf(os.Stdout, "Response from `PostSetsAPI.UpdatePostSetPosts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the post set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePostSetPostsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postSetPostIdsString** | **string** |  | 

### Return type

[**PostSet**](PostSet.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

