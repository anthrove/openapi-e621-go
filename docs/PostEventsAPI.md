# \PostEventsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchPostEvents**](PostEventsAPI.md#SearchPostEvents) | **Get** /post_events.json | Search Post Events



## SearchPostEvents

> SearchPostEvents200Response SearchPostEvents(ctx).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchPostId(searchPostId).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchAction(searchAction).Execute()

Search Post Events

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
	searchPostId := float32(8.14) // float32 |  (optional)
	searchCreatorId := float32(8.14) // float32 |  (optional)
	searchCreatorName := "searchCreatorName_example" // string |  (optional)
	searchAction := openapiclient.PostEventActions("deleted") // PostEventActions |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PostEventsAPI.SearchPostEvents(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchPostId(searchPostId).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchAction(searchAction).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PostEventsAPI.SearchPostEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPostEvents`: SearchPostEvents200Response
	fmt.Fprintf(os.Stdout, "Response from `PostEventsAPI.SearchPostEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPostEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **int32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **int32** | Search for a specific id. | 
 **searchOrder** | **string** |  | 
 **searchPostId** | **float32** |  | 
 **searchCreatorId** | **float32** |  | 
 **searchCreatorName** | **string** |  | 
 **searchAction** | [**PostEventActions**](PostEventActions.md) |  | 

### Return type

[**SearchPostEvents200Response**](SearchPostEvents200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

