# \PostReplacementsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApprovePostReplacement**](PostReplacementsAPI.md#ApprovePostReplacement) | **Put** /post_replacements/{id}/approve.json | Approve Post Replacement
[**CreatePostReplacement**](PostReplacementsAPI.md#CreatePostReplacement) | **Post** /post_replacements.json | Create Post Replacement
[**DeletePostReplacement**](PostReplacementsAPI.md#DeletePostReplacement) | **Delete** /post_replacements/{id}.json | Delete Post Replacement
[**PromotePostReplacement**](PostReplacementsAPI.md#PromotePostReplacement) | **Post** /post_replacements/{id}/promote.json | Promote Post Replacement
[**RejectPostReplacement**](PostReplacementsAPI.md#RejectPostReplacement) | **Put** /post_replacements/{id}/reject.json | Reject Post Replacement
[**SearchPostReplacements**](PostReplacementsAPI.md#SearchPostReplacements) | **Get** /post_replacements.json | Search Post Replacements
[**TogglePostReplacementPenalty**](PostReplacementsAPI.md#TogglePostReplacementPenalty) | **Put** /post_replacements/{id}/toggle_penalize.json | Toggle Post Replacement Penalty



## ApprovePostReplacement

> ApprovePostReplacement(ctx, id).Execute()

Approve Post Replacement



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
	id := int32(56) // int32 | The ID of the post replacement.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PostReplacementsAPI.ApprovePostReplacement(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostReplacementsAPI.ApprovePostReplacement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the post replacement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApprovePostReplacementRequest struct via the builder pattern


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


## CreatePostReplacement

> CreatePostReplacement200Response CreatePostReplacement(ctx).PostReplacementReason(postReplacementReason).PostReplacementReplacementFile(postReplacementReplacementFile).PostReplacementReplacementUrl(postReplacementReplacementUrl).PostReplacementSource(postReplacementSource).PostReplacementAsPending(postReplacementAsPending).Execute()

Create Post Replacement

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
	postReplacementReason := "postReplacementReason_example" // string | 
	postReplacementReplacementFile := os.NewFile(1234, "some_file") // *os.File | Mutually exclusive with replacement_url. (optional)
	postReplacementReplacementUrl := "postReplacementReplacementUrl_example" // string | Mutually exclusive with replacement_file. (optional)
	postReplacementSource := "postReplacementSource_example" // string |  (optional)
	postReplacementAsPending := true // bool | You must be Janitor+. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostReplacementsAPI.CreatePostReplacement(context.Background()).PostReplacementReason(postReplacementReason).PostReplacementReplacementFile(postReplacementReplacementFile).PostReplacementReplacementUrl(postReplacementReplacementUrl).PostReplacementSource(postReplacementSource).PostReplacementAsPending(postReplacementAsPending).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostReplacementsAPI.CreatePostReplacement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePostReplacement`: CreatePostReplacement200Response
	fmt.Fprintf(os.Stdout, "Response from `PostReplacementsAPI.CreatePostReplacement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePostReplacementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postReplacementReason** | **string** |  | 
 **postReplacementReplacementFile** | ***os.File** | Mutually exclusive with replacement_url. | 
 **postReplacementReplacementUrl** | **string** | Mutually exclusive with replacement_file. | 
 **postReplacementSource** | **string** |  | 
 **postReplacementAsPending** | **bool** | You must be Janitor+. | 

### Return type

[**CreatePostReplacement200Response**](CreatePostReplacement200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePostReplacement

> DeletePostReplacement(ctx, id).Execute()

Delete Post Replacement



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
	id := int32(56) // int32 | The ID of the post replacement.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PostReplacementsAPI.DeletePostReplacement(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostReplacementsAPI.DeletePostReplacement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the post replacement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePostReplacementRequest struct via the builder pattern


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


## PromotePostReplacement

> AddFavorite201Response PromotePostReplacement(ctx, id).Execute()

Promote Post Replacement



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
	id := int32(56) // int32 | The ID of the post replacement.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostReplacementsAPI.PromotePostReplacement(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostReplacementsAPI.PromotePostReplacement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PromotePostReplacement`: AddFavorite201Response
	fmt.Fprintf(os.Stdout, "Response from `PostReplacementsAPI.PromotePostReplacement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the post replacement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPromotePostReplacementRequest struct via the builder pattern


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


## RejectPostReplacement

> RejectPostReplacement(ctx, id).RejectPostReplacementRequest(rejectPostReplacementRequest).Execute()

Reject Post Replacement



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
	id := int32(56) // int32 | The ID of the post replacement.
	rejectPostReplacementRequest := *openapiclient.NewRejectPostReplacementRequest() // RejectPostReplacementRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PostReplacementsAPI.RejectPostReplacement(context.Background(), id).RejectPostReplacementRequest(rejectPostReplacementRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostReplacementsAPI.RejectPostReplacement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the post replacement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectPostReplacementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rejectPostReplacementRequest** | [**RejectPostReplacementRequest**](RejectPostReplacementRequest.md) |  | 

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


## SearchPostReplacements

> SearchPostReplacements200Response SearchPostReplacements(ctx).Limit(limit).Page(page).SearchId(searchId).SearchFileExt(searchFileExt).SearchMd5(searchMd5).SearchStatus(searchStatus).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchApproverId(searchApproverId).SearchApproverName(searchApproverName).SearchRejectorId(searchRejectorId).SearchRejectorName(searchRejectorName).SearchUploaderNameOnApprove(searchUploaderNameOnApprove).SearchUploaderIdOnApprove(searchUploaderIdOnApprove).Execute()

Search Post Replacements



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
	searchFileExt := "searchFileExt_example" // string |  (optional)
	searchMd5 := "searchMd5_example" // string |  (optional)
	searchStatus := "searchStatus_example" // string |  (optional)
	searchCreatorId := int32(56) // int32 |  (optional)
	searchCreatorName := "searchCreatorName_example" // string |  (optional)
	searchApproverId := int32(56) // int32 |  (optional)
	searchApproverName := "searchApproverName_example" // string |  (optional)
	searchRejectorId := int32(56) // int32 |  (optional)
	searchRejectorName := "searchRejectorName_example" // string |  (optional)
	searchUploaderNameOnApprove := "searchUploaderNameOnApprove_example" // string |  (optional)
	searchUploaderIdOnApprove := int32(56) // int32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostReplacementsAPI.SearchPostReplacements(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchFileExt(searchFileExt).SearchMd5(searchMd5).SearchStatus(searchStatus).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchApproverId(searchApproverId).SearchApproverName(searchApproverName).SearchRejectorId(searchRejectorId).SearchRejectorName(searchRejectorName).SearchUploaderNameOnApprove(searchUploaderNameOnApprove).SearchUploaderIdOnApprove(searchUploaderIdOnApprove).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostReplacementsAPI.SearchPostReplacements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPostReplacements`: SearchPostReplacements200Response
	fmt.Fprintf(os.Stdout, "Response from `PostReplacementsAPI.SearchPostReplacements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPostReplacementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **int32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **int32** | Search for a specific id. | 
 **searchFileExt** | **string** |  | 
 **searchMd5** | **string** |  | 
 **searchStatus** | **string** |  | 
 **searchCreatorId** | **int32** |  | 
 **searchCreatorName** | **string** |  | 
 **searchApproverId** | **int32** |  | 
 **searchApproverName** | **string** |  | 
 **searchRejectorId** | **int32** |  | 
 **searchRejectorName** | **string** |  | 
 **searchUploaderNameOnApprove** | **string** |  | 
 **searchUploaderIdOnApprove** | **int32** |  | 

### Return type

[**SearchPostReplacements200Response**](SearchPostReplacements200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TogglePostReplacementPenalty

> TogglePostReplacementPenalty(ctx, id).Execute()

Toggle Post Replacement Penalty



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
	id := int32(56) // int32 | The ID of the post replacement.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PostReplacementsAPI.TogglePostReplacementPenalty(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostReplacementsAPI.TogglePostReplacementPenalty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the post replacement. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTogglePostReplacementPenaltyRequest struct via the builder pattern


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

