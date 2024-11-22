# \IQDBAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**QueryIQDBGet**](IQDBAPI.md#QueryIQDBGet) | **Get** /iqdb_queries.json | Query IQDB
[**QueryIQDPost**](IQDBAPI.md#QueryIQDPost) | **Post** /iqdb_queries.json | Query IQDB



## QueryIQDBGet

> QueryIQDBGet200Response QueryIQDBGet(ctx).SearchScoreCutoff(searchScoreCutoff).SearchUrl(searchUrl).SearchPostId(searchPostId).SearchHash(searchHash).Execute()

Query IQDB

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
	searchScoreCutoff := float32(8.14) // float32 |  (optional)
	searchUrl := "searchUrl_example" // string |  (optional)
	searchPostId := int32(56) // int32 |  (optional)
	searchHash := "searchHash_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IQDBAPI.QueryIQDBGet(context.Background()).SearchScoreCutoff(searchScoreCutoff).SearchUrl(searchUrl).SearchPostId(searchPostId).SearchHash(searchHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IQDBAPI.QueryIQDBGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryIQDBGet`: QueryIQDBGet200Response
	fmt.Fprintf(os.Stdout, "Response from `IQDBAPI.QueryIQDBGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryIQDBGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchScoreCutoff** | **float32** |  | 
 **searchUrl** | **string** |  | 
 **searchPostId** | **int32** |  | 
 **searchHash** | **string** |  | 

### Return type

[**QueryIQDBGet200Response**](QueryIQDBGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryIQDPost

> QueryIQDBGet200Response QueryIQDPost(ctx).SearchFile(searchFile).SearchScoreCutoff(searchScoreCutoff).SearchUrl(searchUrl).SearchPostId(searchPostId).SearchHash(searchHash).Execute()

Query IQDB

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
	searchFile := os.NewFile(1234, "some_file") // *os.File |  (optional)
	searchScoreCutoff := float32(8.14) // float32 |  (optional)
	searchUrl := "searchUrl_example" // string |  (optional)
	searchPostId := "searchPostId_example" // string |  (optional)
	searchHash := "searchHash_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IQDBAPI.QueryIQDPost(context.Background()).SearchFile(searchFile).SearchScoreCutoff(searchScoreCutoff).SearchUrl(searchUrl).SearchPostId(searchPostId).SearchHash(searchHash).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IQDBAPI.QueryIQDPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `QueryIQDPost`: QueryIQDBGet200Response
	fmt.Fprintf(os.Stdout, "Response from `IQDBAPI.QueryIQDPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryIQDPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchFile** | ***os.File** |  | 
 **searchScoreCutoff** | **float32** |  | 
 **searchUrl** | **string** |  | 
 **searchPostId** | **string** |  | 
 **searchHash** | **string** |  | 

### Return type

[**QueryIQDBGet200Response**](QueryIQDBGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data, application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

