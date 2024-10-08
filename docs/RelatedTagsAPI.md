# \RelatedTagsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListBulkRelatedTags**](RelatedTagsAPI.md#ListBulkRelatedTags) | **Post** /related_tag/bulk.json | List Bulk Related Tags



## ListBulkRelatedTags

> map[string]interface{} ListBulkRelatedTags(ctx).ListBulkRelatedTagsRequest(listBulkRelatedTagsRequest).Execute()

List Bulk Related Tags

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
	listBulkRelatedTagsRequest := *openapiclient.NewListBulkRelatedTagsRequest() // ListBulkRelatedTagsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RelatedTagsAPI.ListBulkRelatedTags(context.Background()).ListBulkRelatedTagsRequest(listBulkRelatedTagsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RelatedTagsAPI.ListBulkRelatedTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListBulkRelatedTags`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `RelatedTagsAPI.ListBulkRelatedTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBulkRelatedTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listBulkRelatedTagsRequest** | [**ListBulkRelatedTagsRequest**](ListBulkRelatedTagsRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

