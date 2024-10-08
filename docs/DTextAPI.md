# \DTextAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PreviewDText**](DTextAPI.md#PreviewDText) | **Post** /dtext_preview.json | Preview DText



## PreviewDText

> DTextResponse PreviewDText(ctx).PreviewDTextRequest(previewDTextRequest).Execute()

Preview DText



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
	previewDTextRequest := *openapiclient.NewPreviewDTextRequest("Body_example") // PreviewDTextRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DTextAPI.PreviewDText(context.Background()).PreviewDTextRequest(previewDTextRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DTextAPI.PreviewDText``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreviewDText`: DTextResponse
	fmt.Fprintf(os.Stdout, "Response from `DTextAPI.PreviewDText`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPreviewDTextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **previewDTextRequest** | [**PreviewDTextRequest**](PreviewDTextRequest.md) |  | 

### Return type

[**DTextResponse**](DTextResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

