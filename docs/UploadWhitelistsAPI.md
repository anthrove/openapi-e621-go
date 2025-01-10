# \UploadWhitelistsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckIfUrlIsAllowed**](UploadWhitelistsAPI.md#CheckIfUrlIsAllowed) | **Get** /upload_whitelists/{id}/is_allowed.json | Check If URL Is Allowed
[**CreateUploadWhitelist**](UploadWhitelistsAPI.md#CreateUploadWhitelist) | **Post** /upload_whitelists.json | Create Upload Whitelist
[**DeleteUploadWhitelist**](UploadWhitelistsAPI.md#DeleteUploadWhitelist) | **Delete** /upload_whitelists/{id}.json | Delete Upload Whitelist
[**EditUploadWhitelist**](UploadWhitelistsAPI.md#EditUploadWhitelist) | **Patch** /upload_whitelists/{id}.json | Edit Upload Whitelist
[**SearchUploadWhitelists**](UploadWhitelistsAPI.md#SearchUploadWhitelists) | **Get** /upload_whitelists.json | Search Upload Whitelists



## CheckIfUrlIsAllowed

> CheckIfUrlIsAllowed200Response CheckIfUrlIsAllowed(ctx, id).Url(url).Execute()

Check If URL Is Allowed

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
	id := int32(56) // int32 | The ID of the upload whitelist.
	url := "url_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadWhitelistsAPI.CheckIfUrlIsAllowed(context.Background(), id).Url(url).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadWhitelistsAPI.CheckIfUrlIsAllowed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckIfUrlIsAllowed`: CheckIfUrlIsAllowed200Response
	fmt.Fprintf(os.Stdout, "Response from `UploadWhitelistsAPI.CheckIfUrlIsAllowed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the upload whitelist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckIfUrlIsAllowedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **url** | **string** |  | 

### Return type

[**CheckIfUrlIsAllowed200Response**](CheckIfUrlIsAllowed200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUploadWhitelist

> UploadWhitelist CreateUploadWhitelist(ctx).UploadWhitelistAllowed(uploadWhitelistAllowed).UploadWhitelistPattern(uploadWhitelistPattern).UploadWhitelistReason(uploadWhitelistReason).UploadWhitelistNote(uploadWhitelistNote).UploadWhitelistHidden(uploadWhitelistHidden).Execute()

Create Upload Whitelist



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
	uploadWhitelistAllowed := "uploadWhitelistAllowed_example" // string | 
	uploadWhitelistPattern := "uploadWhitelistPattern_example" // string | 
	uploadWhitelistReason := "uploadWhitelistReason_example" // string |  (optional)
	uploadWhitelistNote := "uploadWhitelistNote_example" // string |  (optional)
	uploadWhitelistHidden := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadWhitelistsAPI.CreateUploadWhitelist(context.Background()).UploadWhitelistAllowed(uploadWhitelistAllowed).UploadWhitelistPattern(uploadWhitelistPattern).UploadWhitelistReason(uploadWhitelistReason).UploadWhitelistNote(uploadWhitelistNote).UploadWhitelistHidden(uploadWhitelistHidden).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadWhitelistsAPI.CreateUploadWhitelist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUploadWhitelist`: UploadWhitelist
	fmt.Fprintf(os.Stdout, "Response from `UploadWhitelistsAPI.CreateUploadWhitelist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUploadWhitelistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **uploadWhitelistAllowed** | **string** |  | 
 **uploadWhitelistPattern** | **string** |  | 
 **uploadWhitelistReason** | **string** |  | 
 **uploadWhitelistNote** | **string** |  | 
 **uploadWhitelistHidden** | **bool** |  | 

### Return type

[**UploadWhitelist**](UploadWhitelist.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUploadWhitelist

> DeleteUploadWhitelist(ctx, id).Execute()

Delete Upload Whitelist



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
	id := int32(56) // int32 | The ID of the upload whitelist.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UploadWhitelistsAPI.DeleteUploadWhitelist(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadWhitelistsAPI.DeleteUploadWhitelist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the upload whitelist. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUploadWhitelistRequest struct via the builder pattern


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


## EditUploadWhitelist

> EditUploadWhitelist(ctx, id).UploadWhitelistAllowed(uploadWhitelistAllowed).UploadWhitelistPattern(uploadWhitelistPattern).UploadWhitelistReason(uploadWhitelistReason).UploadWhitelistNote(uploadWhitelistNote).UploadWhitelistHidden(uploadWhitelistHidden).Execute()

Edit Upload Whitelist

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
	id := int32(56) // int32 | The ID of the upload whitelist entry.
	uploadWhitelistAllowed := "uploadWhitelistAllowed_example" // string |  (optional)
	uploadWhitelistPattern := "uploadWhitelistPattern_example" // string |  (optional)
	uploadWhitelistReason := "uploadWhitelistReason_example" // string |  (optional)
	uploadWhitelistNote := "uploadWhitelistNote_example" // string |  (optional)
	uploadWhitelistHidden := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UploadWhitelistsAPI.EditUploadWhitelist(context.Background(), id).UploadWhitelistAllowed(uploadWhitelistAllowed).UploadWhitelistPattern(uploadWhitelistPattern).UploadWhitelistReason(uploadWhitelistReason).UploadWhitelistNote(uploadWhitelistNote).UploadWhitelistHidden(uploadWhitelistHidden).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadWhitelistsAPI.EditUploadWhitelist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the upload whitelist entry. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditUploadWhitelistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uploadWhitelistAllowed** | **string** |  | 
 **uploadWhitelistPattern** | **string** |  | 
 **uploadWhitelistReason** | **string** |  | 
 **uploadWhitelistNote** | **string** |  | 
 **uploadWhitelistHidden** | **bool** |  | 

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


## SearchUploadWhitelists

> SearchUploadWhitelists200Response SearchUploadWhitelists(ctx).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchPattern(searchPattern).SearchNote(searchNote).Execute()

Search Upload Whitelists



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
	searchPattern := "searchPattern_example" // string |  (optional)
	searchNote := "searchNote_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UploadWhitelistsAPI.SearchUploadWhitelists(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchPattern(searchPattern).SearchNote(searchNote).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UploadWhitelistsAPI.SearchUploadWhitelists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUploadWhitelists`: SearchUploadWhitelists200Response
	fmt.Fprintf(os.Stdout, "Response from `UploadWhitelistsAPI.SearchUploadWhitelists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUploadWhitelistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **int32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **int32** | Search for a specific id. | 
 **searchOrder** | **string** |  | 
 **searchPattern** | **string** |  | 
 **searchNote** | **string** |  | 

### Return type

[**SearchUploadWhitelists200Response**](SearchUploadWhitelists200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

