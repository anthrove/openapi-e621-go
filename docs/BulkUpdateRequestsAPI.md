# \BulkUpdateRequestsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveBulkUpdateRequest**](BulkUpdateRequestsAPI.md#ApproveBulkUpdateRequest) | **Post** /bulk_update_request/{id}/approve.json | Approve Bulk Update Request
[**CreateBulkUpdateRequest**](BulkUpdateRequestsAPI.md#CreateBulkUpdateRequest) | **Post** /bulk_update_requests.json | Create Bulk Update Request
[**EditBulkUpdateRequest**](BulkUpdateRequestsAPI.md#EditBulkUpdateRequest) | **Patch** /bulk_update_request/{id}.json | Edit Bulk Update Request
[**GetBulkUpdateRequest**](BulkUpdateRequestsAPI.md#GetBulkUpdateRequest) | **Get** /bulk_update_request/{id}.json | Get Bulk Update Request
[**RejectBulkUpdateRequest**](BulkUpdateRequestsAPI.md#RejectBulkUpdateRequest) | **Delete** /bulk_update_request/{id}.json | Reject Bulk Update Request
[**SearchBulkUpdateRequests**](BulkUpdateRequestsAPI.md#SearchBulkUpdateRequests) | **Get** /bulk_update_requests.json | Search Bulk Update Requests



## ApproveBulkUpdateRequest

> ApproveBulkUpdateRequest(ctx, id).Execute()

Approve Bulk Update Request



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
	id := float32(8.14) // float32 | The ID of the bulk update request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BulkUpdateRequestsAPI.ApproveBulkUpdateRequest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkUpdateRequestsAPI.ApproveBulkUpdateRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the bulk update request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveBulkUpdateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBulkUpdateRequest

> BulkUpdateRequest CreateBulkUpdateRequest(ctx).BulkUpdateRequestScript(bulkUpdateRequestScript).BulkUpdateRequestTitle(bulkUpdateRequestTitle).BulkUpdateRequestReason(bulkUpdateRequestReason).BulkUpdateRequestForumTopicId(bulkUpdateRequestForumTopicId).BulkUpdateRequestSkipForum(bulkUpdateRequestSkipForum).Execute()

Create Bulk Update Request

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
	bulkUpdateRequestScript := "bulkUpdateRequestScript_example" // string | 
	bulkUpdateRequestTitle := "bulkUpdateRequestTitle_example" // string | 
	bulkUpdateRequestReason := "bulkUpdateRequestReason_example" // string | 
	bulkUpdateRequestForumTopicId := float32(8.14) // float32 |  (optional)
	bulkUpdateRequestSkipForum := true // bool | Only usable for Admin+ (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkUpdateRequestsAPI.CreateBulkUpdateRequest(context.Background()).BulkUpdateRequestScript(bulkUpdateRequestScript).BulkUpdateRequestTitle(bulkUpdateRequestTitle).BulkUpdateRequestReason(bulkUpdateRequestReason).BulkUpdateRequestForumTopicId(bulkUpdateRequestForumTopicId).BulkUpdateRequestSkipForum(bulkUpdateRequestSkipForum).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkUpdateRequestsAPI.CreateBulkUpdateRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBulkUpdateRequest`: BulkUpdateRequest
	fmt.Fprintf(os.Stdout, "Response from `BulkUpdateRequestsAPI.CreateBulkUpdateRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBulkUpdateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkUpdateRequestScript** | **string** |  | 
 **bulkUpdateRequestTitle** | **string** |  | 
 **bulkUpdateRequestReason** | **string** |  | 
 **bulkUpdateRequestForumTopicId** | **float32** |  | 
 **bulkUpdateRequestSkipForum** | **bool** | Only usable for Admin+ | 

### Return type

[**BulkUpdateRequest**](BulkUpdateRequest.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditBulkUpdateRequest

> EditBulkUpdateRequest(ctx, id).BulkUpdateRequestScript(bulkUpdateRequestScript).BulkUpdateRequestForumTopicId(bulkUpdateRequestForumTopicId).BulkUpdateRequestForumPostId(bulkUpdateRequestForumPostId).Execute()

Edit Bulk Update Request



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
	id := float32(8.14) // float32 | The ID of the bulk update request.
	bulkUpdateRequestScript := "bulkUpdateRequestScript_example" // string |  (optional)
	bulkUpdateRequestForumTopicId := "bulkUpdateRequestForumTopicId_example" // string | You must be Admin+. (optional)
	bulkUpdateRequestForumPostId := "bulkUpdateRequestForumPostId_example" // string | You must be Admin+. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BulkUpdateRequestsAPI.EditBulkUpdateRequest(context.Background(), id).BulkUpdateRequestScript(bulkUpdateRequestScript).BulkUpdateRequestForumTopicId(bulkUpdateRequestForumTopicId).BulkUpdateRequestForumPostId(bulkUpdateRequestForumPostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkUpdateRequestsAPI.EditBulkUpdateRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the bulk update request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditBulkUpdateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkUpdateRequestScript** | **string** |  | 
 **bulkUpdateRequestForumTopicId** | **string** | You must be Admin+. | 
 **bulkUpdateRequestForumPostId** | **string** | You must be Admin+. | 

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


## GetBulkUpdateRequest

> BulkUpdateRequest GetBulkUpdateRequest(ctx, id).Execute()

Get Bulk Update Request

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
	id := float32(8.14) // float32 | The ID of the bulk update request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkUpdateRequestsAPI.GetBulkUpdateRequest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkUpdateRequestsAPI.GetBulkUpdateRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBulkUpdateRequest`: BulkUpdateRequest
	fmt.Fprintf(os.Stdout, "Response from `BulkUpdateRequestsAPI.GetBulkUpdateRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the bulk update request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBulkUpdateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BulkUpdateRequest**](BulkUpdateRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectBulkUpdateRequest

> RejectBulkUpdateRequest(ctx, id).Execute()

Reject Bulk Update Request



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
	id := float32(8.14) // float32 | The ID of the bulk update request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BulkUpdateRequestsAPI.RejectBulkUpdateRequest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkUpdateRequestsAPI.RejectBulkUpdateRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the bulk update request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectBulkUpdateRequestRequest struct via the builder pattern


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


## SearchBulkUpdateRequests

> SearchBulkUpdateRequests200Response SearchBulkUpdateRequests(ctx).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchUserId(searchUserId).SearchUserName(searchUserName).SearchApproverId(searchApproverId).SearchApproverName(searchApproverName).SearchForumTopicId(searchForumTopicId).SearchForumPostId(searchForumPostId).SearchStatus(searchStatus).SearchTitleMatches(searchTitleMatches).SearchScriptMatches(searchScriptMatches).Execute()

Search Bulk Update Requests



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
	searchUserId := float32(8.14) // float32 |  (optional)
	searchUserName := "searchUserName_example" // string |  (optional)
	searchApproverId := float32(8.14) // float32 |  (optional)
	searchApproverName := "searchApproverName_example" // string |  (optional)
	searchForumTopicId := float32(8.14) // float32 |  (optional)
	searchForumPostId := float32(8.14) // float32 |  (optional)
	searchStatus := "searchStatus_example" // string |  (optional)
	searchTitleMatches := "searchTitleMatches_example" // string |  (optional)
	searchScriptMatches := "searchScriptMatches_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BulkUpdateRequestsAPI.SearchBulkUpdateRequests(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchUserId(searchUserId).SearchUserName(searchUserName).SearchApproverId(searchApproverId).SearchApproverName(searchApproverName).SearchForumTopicId(searchForumTopicId).SearchForumPostId(searchForumPostId).SearchStatus(searchStatus).SearchTitleMatches(searchTitleMatches).SearchScriptMatches(searchScriptMatches).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BulkUpdateRequestsAPI.SearchBulkUpdateRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchBulkUpdateRequests`: SearchBulkUpdateRequests200Response
	fmt.Fprintf(os.Stdout, "Response from `BulkUpdateRequestsAPI.SearchBulkUpdateRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchBulkUpdateRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **float32** | Search for a specific id. | 
 **searchOrder** | **string** |  | 
 **searchUserId** | **float32** |  | 
 **searchUserName** | **string** |  | 
 **searchApproverId** | **float32** |  | 
 **searchApproverName** | **string** |  | 
 **searchForumTopicId** | **float32** |  | 
 **searchForumPostId** | **float32** |  | 
 **searchStatus** | **string** |  | 
 **searchTitleMatches** | **string** |  | 
 **searchScriptMatches** | **string** |  | 

### Return type

[**SearchBulkUpdateRequests200Response**](SearchBulkUpdateRequests200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

