# \EmailBlacklistsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmailBlacklist**](EmailBlacklistsAPI.md#CreateEmailBlacklist) | **Post** /email_blacklists.json | Create Email Blacklist
[**DeleteEmailBlacklist**](EmailBlacklistsAPI.md#DeleteEmailBlacklist) | **Delete** /email_blacklists/{id}.json | Delete Email Blacklist
[**SearchEmailBlacklists**](EmailBlacklistsAPI.md#SearchEmailBlacklists) | **Get** /email_blacklists.json | Search Email Blacklists



## CreateEmailBlacklist

> EmailBlacklist CreateEmailBlacklist(ctx).EmailBlacklistDomain(emailBlacklistDomain).EmailBlacklistReason(emailBlacklistReason).Execute()

Create Email Blacklist



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
	emailBlacklistDomain := "emailBlacklistDomain_example" // string | 
	emailBlacklistReason := "emailBlacklistReason_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailBlacklistsAPI.CreateEmailBlacklist(context.Background()).EmailBlacklistDomain(emailBlacklistDomain).EmailBlacklistReason(emailBlacklistReason).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailBlacklistsAPI.CreateEmailBlacklist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEmailBlacklist`: EmailBlacklist
	fmt.Fprintf(os.Stdout, "Response from `EmailBlacklistsAPI.CreateEmailBlacklist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmailBlacklistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **emailBlacklistDomain** | **string** |  | 
 **emailBlacklistReason** | **string** |  | 

### Return type

[**EmailBlacklist**](EmailBlacklist.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEmailBlacklist

> DeleteEmailBlacklist(ctx, id).Execute()

Delete Email Blacklist



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
	id := float32(8.14) // float32 | The ID of the email blacklist.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EmailBlacklistsAPI.DeleteEmailBlacklist(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailBlacklistsAPI.DeleteEmailBlacklist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the email blacklist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmailBlacklistRequest struct via the builder pattern


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


## SearchEmailBlacklists

> SearchEmailBlacklists200Response SearchEmailBlacklists(ctx).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchDomain(searchDomain).SearchReason(searchReason).Execute()

Search Email Blacklists



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
	searchDomain := "searchDomain_example" // string |  (optional)
	searchReason := "searchReason_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EmailBlacklistsAPI.SearchEmailBlacklists(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchDomain(searchDomain).SearchReason(searchReason).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EmailBlacklistsAPI.SearchEmailBlacklists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchEmailBlacklists`: SearchEmailBlacklists200Response
	fmt.Fprintf(os.Stdout, "Response from `EmailBlacklistsAPI.SearchEmailBlacklists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchEmailBlacklistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **float32** | Search for a specific id. | 
 **searchOrder** | **string** |  | 
 **searchDomain** | **string** |  | 
 **searchReason** | **string** |  | 

### Return type

[**SearchEmailBlacklists200Response**](SearchEmailBlacklists200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

