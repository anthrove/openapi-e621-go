# \TagAliasesAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApproveTagAlias**](TagAliasesAPI.md#ApproveTagAlias) | **Post** /tag_aliases/{id}/approve.json | Approve Tag Alias
[**CreateTagAlias**](TagAliasesAPI.md#CreateTagAlias) | **Post** /tag_alias_requests.json | Create Tag Alias
[**EditTagAlias**](TagAliasesAPI.md#EditTagAlias) | **Patch** /tag_aliases/{id}.json | Edit Tag Alias
[**GetTagAlias**](TagAliasesAPI.md#GetTagAlias) | **Get** /tag_aliases/{id}.json | Get Tag Alias
[**RejectTagAlias**](TagAliasesAPI.md#RejectTagAlias) | **Delete** /tag_aliases/{id}.json | Reject Tag Alias
[**SearchTagAliases**](TagAliasesAPI.md#SearchTagAliases) | **Get** /tag_aliases.json | Search Tag Aliases



## ApproveTagAlias

> ApproveTagAlias(ctx, id).Execute()

Approve Tag Alias



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
	id := float32(8.14) // float32 | The ID of the tag alias.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagAliasesAPI.ApproveTagAlias(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAliasesAPI.ApproveTagAlias``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the tag alias. | 

### Other Parameters

Other parameters are passed through a pointer to a apiApproveTagAliasRequest struct via the builder pattern


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


## CreateTagAlias

> CreateTagAlias(ctx).TagAliasAntecedentName(tagAliasAntecedentName).TagAliasConsequentName(tagAliasConsequentName).TagAliasReason(tagAliasReason).TagAliasSkipForum(tagAliasSkipForum).Execute()

Create Tag Alias



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
	tagAliasAntecedentName := "tagAliasAntecedentName_example" // string | 
	tagAliasConsequentName := "tagAliasConsequentName_example" // string | 
	tagAliasReason := "tagAliasReason_example" // string | 
	tagAliasSkipForum := true // bool | Must be Admin+. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagAliasesAPI.CreateTagAlias(context.Background()).TagAliasAntecedentName(tagAliasAntecedentName).TagAliasConsequentName(tagAliasConsequentName).TagAliasReason(tagAliasReason).TagAliasSkipForum(tagAliasSkipForum).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAliasesAPI.CreateTagAlias``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagAliasAntecedentName** | **string** |  | 
 **tagAliasConsequentName** | **string** |  | 
 **tagAliasReason** | **string** |  | 
 **tagAliasSkipForum** | **bool** | Must be Admin+. | 

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


## EditTagAlias

> EditTagAlias(ctx, id).TagAliasAntecedentName(tagAliasAntecedentName).TagAliasConsequentName(tagAliasConsequentName).TagAliasForumTopicId(tagAliasForumTopicId).Execute()

Edit Tag Alias



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
	id := float32(8.14) // float32 | The ID of the tag alias.
	tagAliasAntecedentName := "tagAliasAntecedentName_example" // string |  (optional)
	tagAliasConsequentName := "tagAliasConsequentName_example" // string |  (optional)
	tagAliasForumTopicId := "tagAliasForumTopicId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagAliasesAPI.EditTagAlias(context.Background(), id).TagAliasAntecedentName(tagAliasAntecedentName).TagAliasConsequentName(tagAliasConsequentName).TagAliasForumTopicId(tagAliasForumTopicId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAliasesAPI.EditTagAlias``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the tag alias. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditTagAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagAliasAntecedentName** | **string** |  | 
 **tagAliasConsequentName** | **string** |  | 
 **tagAliasForumTopicId** | **string** |  | 

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


## GetTagAlias

> TagAlias GetTagAlias(ctx, id).Execute()

Get Tag Alias

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
	id := float32(8.14) // float32 | The ID of the tag alias.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagAliasesAPI.GetTagAlias(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAliasesAPI.GetTagAlias``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTagAlias`: TagAlias
	fmt.Fprintf(os.Stdout, "Response from `TagAliasesAPI.GetTagAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the tag alias. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TagAlias**](TagAlias.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RejectTagAlias

> RejectTagAlias(ctx, id).Execute()

Reject Tag Alias



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
	id := float32(8.14) // float32 | The ID of the tag alias.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TagAliasesAPI.RejectTagAlias(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAliasesAPI.RejectTagAlias``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the tag alias. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRejectTagAliasRequest struct via the builder pattern


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


## SearchTagAliases

> SearchTagAliases200Response SearchTagAliases(ctx).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchNameMatches(searchNameMatches).SearchAntecedentName(searchAntecedentName).SearchConsequentName(searchConsequentName).SearchStatus(searchStatus).SearchAntecedentTagCategory(searchAntecedentTagCategory).SearchConsequentTagCategory(searchConsequentTagCategory).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchApproverId(searchApproverId).SearchApproverName(searchApproverName).Execute()

Search Tag Aliases



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
	resp, r, err := apiClient.TagAliasesAPI.SearchTagAliases(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchNameMatches(searchNameMatches).SearchAntecedentName(searchAntecedentName).SearchConsequentName(searchConsequentName).SearchStatus(searchStatus).SearchAntecedentTagCategory(searchAntecedentTagCategory).SearchConsequentTagCategory(searchConsequentTagCategory).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchApproverId(searchApproverId).SearchApproverName(searchApproverName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagAliasesAPI.SearchTagAliases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchTagAliases`: SearchTagAliases200Response
	fmt.Fprintf(os.Stdout, "Response from `TagAliasesAPI.SearchTagAliases`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchTagAliasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **float32** | Search for a specific id. | 
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

[**SearchTagAliases200Response**](SearchTagAliases200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

