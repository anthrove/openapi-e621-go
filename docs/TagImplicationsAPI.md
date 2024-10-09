# \TagImplicationsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveTagImplication**](TagImplicationsAPI.md#ApproveTagImplication) | **Post** /tag_implications/{id}/approve.json | Approve Tag Implication
[**CreateTagImplication**](TagImplicationsAPI.md#CreateTagImplication) | **Post** /tag_implication_requests.json | Create Tag Implication
[**EditTagImplication**](TagImplicationsAPI.md#EditTagImplication) | **Patch** /tag_implications/{id}.json | Edit Tag Implication
[**GetTagImplication**](TagImplicationsAPI.md#GetTagImplication) | **Get** /tag_implications/{id}.json | Get Tag Implication
[**RejectTagImplication**](TagImplicationsAPI.md#RejectTagImplication) | **Delete** /tag_implications/{id}.json | Reject Tag Implication
[**SearchTagImplications**](TagImplicationsAPI.md#SearchTagImplications) | **Get** /tag_implications.json | Search Tag Implications



## ApproveTagImplication

> ApproveTagImplication(ctx, id).Execute()

Approve Tag Implication



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
	id := float32(8.14) // float32 | The ID of the tag implication.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagImplicationsAPI.ApproveTagImplication(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagImplicationsAPI.ApproveTagImplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the tag implication. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveTagImplicationRequest struct via the builder pattern


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


## CreateTagImplication

> CreateTagImplication(ctx).TagImplicationAntecedentName(tagImplicationAntecedentName).TagImplicationConsequentName(tagImplicationConsequentName).TagImplicationReason(tagImplicationReason).TagImplicationSkipForum(tagImplicationSkipForum).Execute()

Create Tag Implication



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
	tagImplicationAntecedentName := "tagImplicationAntecedentName_example" // string | 
	tagImplicationConsequentName := "tagImplicationConsequentName_example" // string | 
	tagImplicationReason := "tagImplicationReason_example" // string | 
	tagImplicationSkipForum := true // bool | Must be Admin+. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagImplicationsAPI.CreateTagImplication(context.Background()).TagImplicationAntecedentName(tagImplicationAntecedentName).TagImplicationConsequentName(tagImplicationConsequentName).TagImplicationReason(tagImplicationReason).TagImplicationSkipForum(tagImplicationSkipForum).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagImplicationsAPI.CreateTagImplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagImplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagImplicationAntecedentName** | **string** |  | 
 **tagImplicationConsequentName** | **string** |  | 
 **tagImplicationReason** | **string** |  | 
 **tagImplicationSkipForum** | **bool** | Must be Admin+. | 

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


## EditTagImplication

> EditTagImplication(ctx, id).TagImplicationAntecedentName(tagImplicationAntecedentName).TagImplicationConsequentName(tagImplicationConsequentName).TagImplicationForumTopicId(tagImplicationForumTopicId).Execute()

Edit Tag Implication



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
	id := float32(8.14) // float32 | The ID of the tag implication.
	tagImplicationAntecedentName := "tagImplicationAntecedentName_example" // string |  (optional)
	tagImplicationConsequentName := "tagImplicationConsequentName_example" // string |  (optional)
	tagImplicationForumTopicId := "tagImplicationForumTopicId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagImplicationsAPI.EditTagImplication(context.Background(), id).TagImplicationAntecedentName(tagImplicationAntecedentName).TagImplicationConsequentName(tagImplicationConsequentName).TagImplicationForumTopicId(tagImplicationForumTopicId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagImplicationsAPI.EditTagImplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the tag implication. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditTagImplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagImplicationAntecedentName** | **string** |  | 
 **tagImplicationConsequentName** | **string** |  | 
 **tagImplicationForumTopicId** | **string** |  | 

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


## GetTagImplication

> TagImplication GetTagImplication(ctx, id).Execute()

Get Tag Implication

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
	id := float32(8.14) // float32 | The ID of the tag implication.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagImplicationsAPI.GetTagImplication(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagImplicationsAPI.GetTagImplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTagImplication`: TagImplication
	fmt.Fprintf(os.Stdout, "Response from `TagImplicationsAPI.GetTagImplication`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the tag implication. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagImplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TagImplication**](TagImplication.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectTagImplication

> RejectTagImplication(ctx, id).Execute()

Reject Tag Implication



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
	id := float32(8.14) // float32 | The ID of the tag implication.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagImplicationsAPI.RejectTagImplication(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagImplicationsAPI.RejectTagImplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the tag implication. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectTagImplicationRequest struct via the builder pattern


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


## SearchTagImplications

> SearchTagImplications200Response SearchTagImplications(ctx).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchNameMatches(searchNameMatches).SearchAntecedentName(searchAntecedentName).SearchConsequentName(searchConsequentName).SearchStatus(searchStatus).SearchAntecedentTagCategory(searchAntecedentTagCategory).SearchConsequentTagCategory(searchConsequentTagCategory).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchApproverId(searchApproverId).SearchApproverName(searchApproverName).Execute()

Search Tag Implications



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
	searchNameMatches := "searchNameMatches_example" // string |  (optional)
	searchAntecedentName := "searchAntecedentName_example" // string |  (optional)
	searchConsequentName := "searchConsequentName_example" // string |  (optional)
	searchStatus := openapiclient.TagRequestStatuses("active") // TagRequestStatuses |  (optional)
	searchAntecedentTagCategory := openapiclient.TagCategories(0) // TagCategories |  (optional)
	searchConsequentTagCategory := openapiclient.TagCategories(0) // TagCategories |  (optional)
	searchCreatorId := float32(8.14) // float32 |  (optional)
	searchCreatorName := "searchCreatorName_example" // string |  (optional)
	searchApproverId := float32(8.14) // float32 |  (optional)
	searchApproverName := "searchApproverName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagImplicationsAPI.SearchTagImplications(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchNameMatches(searchNameMatches).SearchAntecedentName(searchAntecedentName).SearchConsequentName(searchConsequentName).SearchStatus(searchStatus).SearchAntecedentTagCategory(searchAntecedentTagCategory).SearchConsequentTagCategory(searchConsequentTagCategory).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchApproverId(searchApproverId).SearchApproverName(searchApproverName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagImplicationsAPI.SearchTagImplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchTagImplications`: SearchTagImplications200Response
	fmt.Fprintf(os.Stdout, "Response from `TagImplicationsAPI.SearchTagImplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchTagImplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **int32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **int32** | Search for a specific id. | 
 **searchOrder** | **string** |  | 
 **searchNameMatches** | **string** |  | 
 **searchAntecedentName** | **string** |  | 
 **searchConsequentName** | **string** |  | 
 **searchStatus** | [**TagRequestStatuses**](TagRequestStatuses.md) |  | 
 **searchAntecedentTagCategory** | [**TagCategories**](TagCategories.md) |  | 
 **searchConsequentTagCategory** | [**TagCategories**](TagCategories.md) |  | 
 **searchCreatorId** | **float32** |  | 
 **searchCreatorName** | **string** |  | 
 **searchApproverId** | **float32** |  | 
 **searchApproverName** | **string** |  | 

### Return type

[**SearchTagImplications200Response**](SearchTagImplications200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

