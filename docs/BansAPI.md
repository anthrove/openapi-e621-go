# \BansAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBan**](BansAPI.md#GetBan) | **Get** /bans/{id}.json | Get Ban
[**SearchBans**](BansAPI.md#SearchBans) | **Get** /bans.json | Search Bans



## GetBan

> Ban GetBan(ctx, id).Execute()

Get Ban

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
	id := float32(8.14) // float32 | The ID of the ban to get.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BansAPI.GetBan(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BansAPI.GetBan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBan`: Ban
	fmt.Fprintf(os.Stdout, "Response from `BansAPI.GetBan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the ban to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Ban**](Ban.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchBans

> SearchBans200Response SearchBans(ctx).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchBannerId(searchBannerId).SearchBannerName(searchBannerName).SearchUserId(searchUserId).SearchUserName(searchUserName).SearchReasonMatches(searchReasonMatches).SearchExpired(searchExpired).Execute()

Search Bans



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
	searchOrder := "searchOrder_example" // string | The order of the results. (optional)
	searchBannerId := "searchBannerId_example" // string | The ID of the banner. (optional)
	searchBannerName := "searchBannerName_example" // string | The name of banner. (optional)
	searchUserId := "searchUserId_example" // string | The ID of the banned user. (optional)
	searchUserName := "searchUserName_example" // string | The name of the banned user. (optional)
	searchReasonMatches := "searchReasonMatches_example" // string | The reason of the ban. (optional)
	searchExpired := "searchExpired_example" // string | If the ban is expired. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BansAPI.SearchBans(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchBannerId(searchBannerId).SearchBannerName(searchBannerName).SearchUserId(searchUserId).SearchUserName(searchUserName).SearchReasonMatches(searchReasonMatches).SearchExpired(searchExpired).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BansAPI.SearchBans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchBans`: SearchBans200Response
	fmt.Fprintf(os.Stdout, "Response from `BansAPI.SearchBans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchBansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **float32** | Search for a specific id. | 
 **searchOrder** | **string** | The order of the results. | 
 **searchBannerId** | **string** | The ID of the banner. | 
 **searchBannerName** | **string** | The name of banner. | 
 **searchUserId** | **string** | The ID of the banned user. | 
 **searchUserName** | **string** | The name of the banned user. | 
 **searchReasonMatches** | **string** | The reason of the ban. | 
 **searchExpired** | **string** | If the ban is expired. | 

### Return type

[**SearchBans200Response**](SearchBans200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

