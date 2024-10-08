# \PoolVersionsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchPoolVersions**](PoolVersionsAPI.md#SearchPoolVersions) | **Get** /pool_versions.json | Search Pool Versions



## SearchPoolVersions

> SearchPoolVersions200Response SearchPoolVersions(ctx).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).SearchUpdaterId(searchUpdaterId).SearchUpdaterName(searchUpdaterName).SearchPoolId(searchPoolId).Execute()

Search Pool Versions



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
	searchIpAddr := "searchIpAddr_example" // string | Must be Admin+ to use. See [postgres' documentation](https://www.postgresql.org/docs/9.3/functions-net.html) for information on how this is parsed. Specifically, \"is contained within or equals\" (`<<=`). (optional)
	searchOrder := "searchOrder_example" // string | The order of the results. (optional)
	searchUpdaterId := float32(8.14) // float32 |  (optional)
	searchUpdaterName := "searchUpdaterName_example" // string |  (optional)
	searchPoolId := float32(8.14) // float32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolVersionsAPI.SearchPoolVersions(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).SearchUpdaterId(searchUpdaterId).SearchUpdaterName(searchUpdaterName).SearchPoolId(searchPoolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolVersionsAPI.SearchPoolVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPoolVersions`: SearchPoolVersions200Response
	fmt.Fprintf(os.Stdout, "Response from `PoolVersionsAPI.SearchPoolVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPoolVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **float32** | Search for a specific id. | 
 **searchIpAddr** | **string** | Must be Admin+ to use. See [postgres&#39; documentation](https://www.postgresql.org/docs/9.3/functions-net.html) for information on how this is parsed. Specifically, \&quot;is contained within or equals\&quot; (&#x60;&lt;&lt;&#x3D;&#x60;). | 
 **searchOrder** | **string** | The order of the results. | 
 **searchUpdaterId** | **float32** |  | 
 **searchUpdaterName** | **string** |  | 
 **searchPoolId** | **float32** |  | 

### Return type

[**SearchPoolVersions200Response**](SearchPoolVersions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

