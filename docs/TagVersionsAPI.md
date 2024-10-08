# \TagVersionsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchTagVersions**](TagVersionsAPI.md#SearchTagVersions) | **Get** /tag_type_versions.json | Search Tag Versions



## SearchTagVersions

> SearchTagVersions200Response SearchTagVersions(ctx).Limit(limit).Page(page).SearchId(searchId).SearchTag(searchTag).SearchUserId(searchUserId).SearchUserName(searchUserName).Execute()

Search Tag Versions



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
	searchTag := "searchTag_example" // string |  (optional)
	searchUserId := "searchUserId_example" // string |  (optional)
	searchUserName := "searchUserName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagVersionsAPI.SearchTagVersions(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchTag(searchTag).SearchUserId(searchUserId).SearchUserName(searchUserName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagVersionsAPI.SearchTagVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchTagVersions`: SearchTagVersions200Response
	fmt.Fprintf(os.Stdout, "Response from `TagVersionsAPI.SearchTagVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchTagVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **float32** | Search for a specific id. | 
 **searchTag** | **string** |  | 
 **searchUserId** | **string** |  | 
 **searchUserName** | **string** |  | 

### Return type

[**SearchTagVersions200Response**](SearchTagVersions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

