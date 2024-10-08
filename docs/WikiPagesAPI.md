# \WikiPagesAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWikiPage**](WikiPagesAPI.md#CreateWikiPage) | **Post** /wiki_pages.json | Create Wiki Page
[**DeleteWikiPage**](WikiPagesAPI.md#DeleteWikiPage) | **Delete** /wiki_pages/{id}.json | Delete Wiki Page
[**EditWikiPage**](WikiPagesAPI.md#EditWikiPage) | **Patch** /wiki_pages/{id}.json | Edit Wiki Page
[**GetWikiPage**](WikiPagesAPI.md#GetWikiPage) | **Get** /wiki_pages/{id}.json | Get Wiki Page
[**RevertWikiPage**](WikiPagesAPI.md#RevertWikiPage) | **Put** /wiki_page/{id}/revert.json | Revert Wiki Page
[**SearchWikiPages**](WikiPagesAPI.md#SearchWikiPages) | **Get** /wiki_pages.json | Search Wiki Pages



## CreateWikiPage

> WikiPage CreateWikiPage(ctx).WikiPageTitle(wikiPageTitle).WikiPageBody(wikiPageBody).WikiPageEditReason(wikiPageEditReason).WikiPageParent(wikiPageParent).WikiPageIsLocked(wikiPageIsLocked).WikiPageIsDeleted(wikiPageIsDeleted).WikiPageSkipSecondaryValidations(wikiPageSkipSecondaryValidations).Execute()

Create Wiki Page

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
	wikiPageTitle := "wikiPageTitle_example" // string | 
	wikiPageBody := "wikiPageBody_example" // string | 
	wikiPageEditReason := "wikiPageEditReason_example" // string |  (optional)
	wikiPageParent := "wikiPageParent_example" // string | Must be Privileged+ to use. (optional)
	wikiPageIsLocked := true // bool | Must be Janitor+ to use. (optional)
	wikiPageIsDeleted := true // bool | Must be Janitor+ to use. (optional)
	wikiPageSkipSecondaryValidations := true // bool | Must be Janitor+ to use. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiPagesAPI.CreateWikiPage(context.Background()).WikiPageTitle(wikiPageTitle).WikiPageBody(wikiPageBody).WikiPageEditReason(wikiPageEditReason).WikiPageParent(wikiPageParent).WikiPageIsLocked(wikiPageIsLocked).WikiPageIsDeleted(wikiPageIsDeleted).WikiPageSkipSecondaryValidations(wikiPageSkipSecondaryValidations).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiPagesAPI.CreateWikiPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWikiPage`: WikiPage
	fmt.Fprintf(os.Stdout, "Response from `WikiPagesAPI.CreateWikiPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWikiPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **wikiPageTitle** | **string** |  | 
 **wikiPageBody** | **string** |  | 
 **wikiPageEditReason** | **string** |  | 
 **wikiPageParent** | **string** | Must be Privileged+ to use. | 
 **wikiPageIsLocked** | **bool** | Must be Janitor+ to use. | 
 **wikiPageIsDeleted** | **bool** | Must be Janitor+ to use. | 
 **wikiPageSkipSecondaryValidations** | **bool** | Must be Janitor+ to use. | 

### Return type

[**WikiPage**](WikiPage.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWikiPage

> DeleteWikiPage(ctx, id).Execute()

Delete Wiki Page



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
	id := float32(8.14) // float32 | The ID of the wiki page.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WikiPagesAPI.DeleteWikiPage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiPagesAPI.DeleteWikiPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the wiki page. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWikiPageRequest struct via the builder pattern


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


## EditWikiPage

> EditWikiPage(ctx, id).WikiPageBody(wikiPageBody).WikiPageEditReason(wikiPageEditReason).WikiPageParent(wikiPageParent).WikiPageTitle(wikiPageTitle).WikiPageIsLocked(wikiPageIsLocked).WikiPageIsDeleted(wikiPageIsDeleted).WikiPageSkipSecondaryValidations(wikiPageSkipSecondaryValidations).Execute()

Edit Wiki Page



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
	id := float32(8.14) // float32 | The ID of the wiki page.
	wikiPageBody := "wikiPageBody_example" // string |  (optional)
	wikiPageEditReason := "wikiPageEditReason_example" // string |  (optional)
	wikiPageParent := "wikiPageParent_example" // string | Must be Privileged+ to use. (optional)
	wikiPageTitle := "wikiPageTitle_example" // string | Must be Janitor+ to use. (optional)
	wikiPageIsLocked := true // bool | Must be Janitor+ to use. (optional)
	wikiPageIsDeleted := true // bool | Must be Janitor+ to use. (optional)
	wikiPageSkipSecondaryValidations := true // bool | Must be Janitor+ to use. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WikiPagesAPI.EditWikiPage(context.Background(), id).WikiPageBody(wikiPageBody).WikiPageEditReason(wikiPageEditReason).WikiPageParent(wikiPageParent).WikiPageTitle(wikiPageTitle).WikiPageIsLocked(wikiPageIsLocked).WikiPageIsDeleted(wikiPageIsDeleted).WikiPageSkipSecondaryValidations(wikiPageSkipSecondaryValidations).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiPagesAPI.EditWikiPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the wiki page. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditWikiPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **wikiPageBody** | **string** |  | 
 **wikiPageEditReason** | **string** |  | 
 **wikiPageParent** | **string** | Must be Privileged+ to use. | 
 **wikiPageTitle** | **string** | Must be Janitor+ to use. | 
 **wikiPageIsLocked** | **bool** | Must be Janitor+ to use. | 
 **wikiPageIsDeleted** | **bool** | Must be Janitor+ to use. | 
 **wikiPageSkipSecondaryValidations** | **bool** | Must be Janitor+ to use. | 

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


## GetWikiPage

> WikiPage GetWikiPage(ctx, id).Execute()

Get Wiki Page

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
	id := float32(8.14) // float32 | The ID or name of the wiki page.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiPagesAPI.GetWikiPage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiPagesAPI.GetWikiPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWikiPage`: WikiPage
	fmt.Fprintf(os.Stdout, "Response from `WikiPagesAPI.GetWikiPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID or name of the wiki page. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWikiPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**WikiPage**](WikiPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevertWikiPage

> RevertWikiPage(ctx, id).VersionId(versionId).Execute()

Revert Wiki Page

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
	id := float32(8.14) // float32 | The ID of the wiki page.
	versionId := float32(8.14) // float32 | The version ID to revert to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WikiPagesAPI.RevertWikiPage(context.Background(), id).VersionId(versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiPagesAPI.RevertWikiPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the wiki page. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevertWikiPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **versionId** | **float32** | The version ID to revert to. | 

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


## SearchWikiPages

> []WikiPage SearchWikiPages(ctx).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchTitle(searchTitle).SearchTitleMatches(searchTitleMatches).SearchBodyMatches(searchBodyMatches).SearchOtherNamesMatch(searchOtherNamesMatch).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchParent(searchParent).SearchOtherNamesPresent(searchOtherNamesPresent).SearchIsLocked(searchIsLocked).SearchIsDeleted(searchIsDeleted).Execute()

Search Wiki Pages

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
	limit := float32(8.14) // float32 | The maximum number of results to return. Between 0 and 320. (optional)
	page := float32(8.14) // float32 | The page number of results to get. Between 1 and 750. (optional)
	searchId := float32(8.14) // float32 | Search for a specific id. (optional)
	searchOrder := "searchOrder_example" // string |  (optional)
	searchTitle := "searchTitle_example" // string |  (optional)
	searchTitleMatches := "searchTitleMatches_example" // string |  (optional)
	searchBodyMatches := "searchBodyMatches_example" // string |  (optional)
	searchOtherNamesMatch := "searchOtherNamesMatch_example" // string |  (optional)
	searchCreatorId := float32(8.14) // float32 |  (optional)
	searchCreatorName := "searchCreatorName_example" // string |  (optional)
	searchParent := "searchParent_example" // string |  (optional)
	searchOtherNamesPresent := true // bool |  (optional)
	searchIsLocked := true // bool |  (optional)
	searchIsDeleted := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WikiPagesAPI.SearchWikiPages(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchTitle(searchTitle).SearchTitleMatches(searchTitleMatches).SearchBodyMatches(searchBodyMatches).SearchOtherNamesMatch(searchOtherNamesMatch).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchParent(searchParent).SearchOtherNamesPresent(searchOtherNamesPresent).SearchIsLocked(searchIsLocked).SearchIsDeleted(searchIsDeleted).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WikiPagesAPI.SearchWikiPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchWikiPages`: []WikiPage
	fmt.Fprintf(os.Stdout, "Response from `WikiPagesAPI.SearchWikiPages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchWikiPagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **float32** | Search for a specific id. | 
 **searchOrder** | **string** |  | 
 **searchTitle** | **string** |  | 
 **searchTitleMatches** | **string** |  | 
 **searchBodyMatches** | **string** |  | 
 **searchOtherNamesMatch** | **string** |  | 
 **searchCreatorId** | **float32** |  | 
 **searchCreatorName** | **string** |  | 
 **searchParent** | **string** |  | 
 **searchOtherNamesPresent** | **bool** |  | 
 **searchIsLocked** | **bool** |  | 
 **searchIsDeleted** | **bool** |  | 

### Return type

[**[]WikiPage**](WikiPage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

