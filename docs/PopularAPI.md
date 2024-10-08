# \PopularAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListPopular**](PopularAPI.md#ListPopular) | **Get** /popular.json | List Most Upvoted Posts



## ListPopular

> ListFavorites200Response ListPopular(ctx).Limit(limit).Date(date).Scale(scale).Execute()

List Most Upvoted Posts

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	limit := float32(8.14) // float32 | The maximum number of results to return. Between 0 and 320. (optional)
	date := time.Now() // string | The date to list popular uploads for. Only The day, month, and year are considered. (optional)
	scale := "scale_example" // string | The scale of the results, in relation to `date`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PopularAPI.ListPopular(context.Background()).Limit(limit).Date(date).Scale(scale).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PopularAPI.ListPopular``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPopular`: ListFavorites200Response
	fmt.Fprintf(os.Stdout, "Response from `PopularAPI.ListPopular`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPopularRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **date** | **string** | The date to list popular uploads for. Only The day, month, and year are considered. | 
 **scale** | **string** | The scale of the results, in relation to &#x60;date&#x60;. | 

### Return type

[**ListFavorites200Response**](ListFavorites200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

