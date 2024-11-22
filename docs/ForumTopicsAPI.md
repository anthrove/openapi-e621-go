# \ForumTopicsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateForumTopic**](ForumTopicsAPI.md#CreateForumTopic) | **Post** /forum_topics.json | Create Forum Topic
[**DeleteForumTopic**](ForumTopicsAPI.md#DeleteForumTopic) | **Delete** /forum_topics/{id}.json | Delete Forum Topic
[**EditForumTopic**](ForumTopicsAPI.md#EditForumTopic) | **Patch** /forum_topics/{id}.json | Edit Forum Topic
[**GetForumTopic**](ForumTopicsAPI.md#GetForumTopic) | **Get** /forum_topics/{id}.json | Get Forum Forum Topic
[**HideForumTopic**](ForumTopicsAPI.md#HideForumTopic) | **Post** /forum_topics/{id}/hide.json | Hide Forum Topic
[**MarkAllForumTopicsAsRead**](ForumTopicsAPI.md#MarkAllForumTopicsAsRead) | **Post** /forum_topics/mark_all_as_read.json | Mark All Forum Topics As Read
[**SearchForumTopics**](ForumTopicsAPI.md#SearchForumTopics) | **Get** /forum_topics.json | Search Forum Topics
[**SubscribeForumTopic**](ForumTopicsAPI.md#SubscribeForumTopic) | **Post** /forum_topics/{id}/subscribe.json | Subscribe To Forum Topic
[**UnhideForumTopic**](ForumTopicsAPI.md#UnhideForumTopic) | **Post** /forum_topics/{id}/unhide.json | Unhide Forum Topic
[**UnsubscribeForumTopic**](ForumTopicsAPI.md#UnsubscribeForumTopic) | **Post** /forum_topics/{id}/unsubscribe.json | Unsubscribe From Forum Topic



## CreateForumTopic

> ForumTopic CreateForumTopic(ctx).ForumTopicTitle(forumTopicTitle).ForumTopicCategoryId(forumTopicCategoryId).ForumTopicOriginalPostAttributesId(forumTopicOriginalPostAttributesId).ForumTopicOriginalPostAttributesBody(forumTopicOriginalPostAttributesBody).ForumTopicIsSticky(forumTopicIsSticky).ForumTopicIsLocked(forumTopicIsLocked).Execute()

Create Forum Topic

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
	forumTopicTitle := "forumTopicTitle_example" // string | 
	forumTopicCategoryId := int32(56) // int32 | 
	forumTopicOriginalPostAttributesId := int32(56) // int32 | Forum post ID. Mutually exclusive with body, one must be provided. (optional)
	forumTopicOriginalPostAttributesBody := "forumTopicOriginalPostAttributesBody_example" // string | First forum post body. Mutually exclusive with id, one must be provided. (optional)
	forumTopicIsSticky := true // bool | You must be Moderator+. (optional)
	forumTopicIsLocked := true // bool | You must be Moderator+. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ForumTopicsAPI.CreateForumTopic(context.Background()).ForumTopicTitle(forumTopicTitle).ForumTopicCategoryId(forumTopicCategoryId).ForumTopicOriginalPostAttributesId(forumTopicOriginalPostAttributesId).ForumTopicOriginalPostAttributesBody(forumTopicOriginalPostAttributesBody).ForumTopicIsSticky(forumTopicIsSticky).ForumTopicIsLocked(forumTopicIsLocked).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForumTopicsAPI.CreateForumTopic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateForumTopic`: ForumTopic
	fmt.Fprintf(os.Stdout, "Response from `ForumTopicsAPI.CreateForumTopic`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateForumTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **forumTopicTitle** | **string** |  | 
 **forumTopicCategoryId** | **int32** |  | 
 **forumTopicOriginalPostAttributesId** | **int32** | Forum post ID. Mutually exclusive with body, one must be provided. | 
 **forumTopicOriginalPostAttributesBody** | **string** | First forum post body. Mutually exclusive with id, one must be provided. | 
 **forumTopicIsSticky** | **bool** | You must be Moderator+. | 
 **forumTopicIsLocked** | **bool** | You must be Moderator+. | 

### Return type

[**ForumTopic**](ForumTopic.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteForumTopic

> DeleteForumTopic(ctx, id).Execute()

Delete Forum Topic



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
	id := int32(56) // int32 | The ID of the forum topic.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ForumTopicsAPI.DeleteForumTopic(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForumTopicsAPI.DeleteForumTopic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the forum topic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteForumTopicRequest struct via the builder pattern


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


## EditForumTopic

> EditForumTopic(ctx, id).ForumTopicTitle(forumTopicTitle).ForumTopicCategoryId(forumTopicCategoryId).ForumTopicOriginalPostAttributesId(forumTopicOriginalPostAttributesId).ForumTopicOriginalPostAttributesBody(forumTopicOriginalPostAttributesBody).ForumTopicIsSticky(forumTopicIsSticky).ForumTopicIsLocked(forumTopicIsLocked).Execute()

Edit Forum Topic

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
	id := int32(56) // int32 | The ID of the forum topic.
	forumTopicTitle := "forumTopicTitle_example" // string |  (optional)
	forumTopicCategoryId := int32(56) // int32 |  (optional)
	forumTopicOriginalPostAttributesId := int32(56) // int32 | Forum post ID. Silently ignored (optional)
	forumTopicOriginalPostAttributesBody := "forumTopicOriginalPostAttributesBody_example" // string | First forum post body. (optional)
	forumTopicIsSticky := true // bool | You must be Moderator+. (optional)
	forumTopicIsLocked := true // bool | You must be Moderator+. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ForumTopicsAPI.EditForumTopic(context.Background(), id).ForumTopicTitle(forumTopicTitle).ForumTopicCategoryId(forumTopicCategoryId).ForumTopicOriginalPostAttributesId(forumTopicOriginalPostAttributesId).ForumTopicOriginalPostAttributesBody(forumTopicOriginalPostAttributesBody).ForumTopicIsSticky(forumTopicIsSticky).ForumTopicIsLocked(forumTopicIsLocked).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForumTopicsAPI.EditForumTopic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the forum topic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditForumTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **forumTopicTitle** | **string** |  | 
 **forumTopicCategoryId** | **int32** |  | 
 **forumTopicOriginalPostAttributesId** | **int32** | Forum post ID. Silently ignored | 
 **forumTopicOriginalPostAttributesBody** | **string** | First forum post body. | 
 **forumTopicIsSticky** | **bool** | You must be Moderator+. | 
 **forumTopicIsLocked** | **bool** | You must be Moderator+. | 

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


## GetForumTopic

> ForumTopic GetForumTopic(ctx, id).Execute()

Get Forum Forum Topic



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
	id := int32(56) // int32 | The ID of the forum topic.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ForumTopicsAPI.GetForumTopic(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForumTopicsAPI.GetForumTopic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetForumTopic`: ForumTopic
	fmt.Fprintf(os.Stdout, "Response from `ForumTopicsAPI.GetForumTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the forum topic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetForumTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ForumTopic**](ForumTopic.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HideForumTopic

> ForumTopic HideForumTopic(ctx, id).Execute()

Hide Forum Topic



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
	id := int32(56) // int32 | The ID of the forum topic.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ForumTopicsAPI.HideForumTopic(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForumTopicsAPI.HideForumTopic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HideForumTopic`: ForumTopic
	fmt.Fprintf(os.Stdout, "Response from `ForumTopicsAPI.HideForumTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the forum topic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHideForumTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ForumTopic**](ForumTopic.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkAllForumTopicsAsRead

> MarkAllForumTopicsAsRead(ctx).MarkAllForumTopicsAsReadRequest(markAllForumTopicsAsReadRequest).Execute()

Mark All Forum Topics As Read

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
	markAllForumTopicsAsReadRequest := *openapiclient.NewMarkAllForumTopicsAsReadRequest(int32(123)) // MarkAllForumTopicsAsReadRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ForumTopicsAPI.MarkAllForumTopicsAsRead(context.Background()).MarkAllForumTopicsAsReadRequest(markAllForumTopicsAsReadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForumTopicsAPI.MarkAllForumTopicsAsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkAllForumTopicsAsReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **markAllForumTopicsAsReadRequest** | [**MarkAllForumTopicsAsReadRequest**](MarkAllForumTopicsAsReadRequest.md) |  | 

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


## SearchForumTopics

> []ForumTopic SearchForumTopics(ctx).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchTitle(searchTitle).SearchTitleMatches(searchTitleMatches).SearchCategoryId(searchCategoryId).SearchIsSticky(searchIsSticky).SearchIsLocked(searchIsLocked).SearchIsHidden(searchIsHidden).Execute()

Search Forum Topics

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
	searchTitle := "searchTitle_example" // string |  (optional)
	searchTitleMatches := "searchTitleMatches_example" // string |  (optional)
	searchCategoryId := int32(56) // int32 |  (optional)
	searchIsSticky := true // bool |  (optional)
	searchIsLocked := true // bool |  (optional)
	searchIsHidden := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ForumTopicsAPI.SearchForumTopics(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchTitle(searchTitle).SearchTitleMatches(searchTitleMatches).SearchCategoryId(searchCategoryId).SearchIsSticky(searchIsSticky).SearchIsLocked(searchIsLocked).SearchIsHidden(searchIsHidden).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForumTopicsAPI.SearchForumTopics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchForumTopics`: []ForumTopic
	fmt.Fprintf(os.Stdout, "Response from `ForumTopicsAPI.SearchForumTopics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchForumTopicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **int32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **int32** | Search for a specific id. | 
 **searchOrder** | **string** |  | 
 **searchTitle** | **string** |  | 
 **searchTitleMatches** | **string** |  | 
 **searchCategoryId** | **int32** |  | 
 **searchIsSticky** | **bool** |  | 
 **searchIsLocked** | **bool** |  | 
 **searchIsHidden** | **bool** |  | 

### Return type

[**[]ForumTopic**](ForumTopic.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscribeForumTopic

> ForumTopic SubscribeForumTopic(ctx, id).Execute()

Subscribe To Forum Topic

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
	id := int32(56) // int32 | The ID of the forum topic.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ForumTopicsAPI.SubscribeForumTopic(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForumTopicsAPI.SubscribeForumTopic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscribeForumTopic`: ForumTopic
	fmt.Fprintf(os.Stdout, "Response from `ForumTopicsAPI.SubscribeForumTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the forum topic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeForumTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ForumTopic**](ForumTopic.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnhideForumTopic

> ForumTopic UnhideForumTopic(ctx, id).Execute()

Unhide Forum Topic



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
	id := int32(56) // int32 | The ID of the forum topic.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ForumTopicsAPI.UnhideForumTopic(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForumTopicsAPI.UnhideForumTopic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnhideForumTopic`: ForumTopic
	fmt.Fprintf(os.Stdout, "Response from `ForumTopicsAPI.UnhideForumTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the forum topic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnhideForumTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ForumTopic**](ForumTopic.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsubscribeForumTopic

> ForumTopic UnsubscribeForumTopic(ctx, id).Execute()

Unsubscribe From Forum Topic

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
	id := int32(56) // int32 | The ID of the forum topic.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ForumTopicsAPI.UnsubscribeForumTopic(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ForumTopicsAPI.UnsubscribeForumTopic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnsubscribeForumTopic`: ForumTopic
	fmt.Fprintf(os.Stdout, "Response from `ForumTopicsAPI.UnsubscribeForumTopic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the forum topic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsubscribeForumTopicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ForumTopic**](ForumTopic.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

