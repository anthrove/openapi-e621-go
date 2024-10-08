# \NewsUpdatesAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNewsUpdate**](NewsUpdatesAPI.md#CreateNewsUpdate) | **Post** /news_updates.json | Create News Update
[**DeleteNewsUpdate**](NewsUpdatesAPI.md#DeleteNewsUpdate) | **Delete** /news_updates/{id}.json | Delete News Update
[**EditNewsUpdate**](NewsUpdatesAPI.md#EditNewsUpdate) | **Patch** /news_updates/{id}.json | Edit News Update
[**ListNewsUpdates**](NewsUpdatesAPI.md#ListNewsUpdates) | **Get** /news_updates.json | List News Updates



## CreateNewsUpdate

> NewsUpdate CreateNewsUpdate(ctx).NewsUpdateMessage(newsUpdateMessage).Execute()

Create News Update



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
	newsUpdateMessage := "newsUpdateMessage_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewsUpdatesAPI.CreateNewsUpdate(context.Background()).NewsUpdateMessage(newsUpdateMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsUpdatesAPI.CreateNewsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNewsUpdate`: NewsUpdate
	fmt.Fprintf(os.Stdout, "Response from `NewsUpdatesAPI.CreateNewsUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNewsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newsUpdateMessage** | **string** |  | 

### Return type

[**NewsUpdate**](NewsUpdate.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNewsUpdate

> DeleteNewsUpdate(ctx, id).Execute()

Delete News Update



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
	id := float32(8.14) // float32 | The ID of the news update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NewsUpdatesAPI.DeleteNewsUpdate(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsUpdatesAPI.DeleteNewsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the news update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNewsUpdateRequest struct via the builder pattern


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


## EditNewsUpdate

> EditNewsUpdate(ctx, id).NewsUpdateMessage(newsUpdateMessage).Execute()

Edit News Update



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
	id := float32(8.14) // float32 | The ID of the news update.
	newsUpdateMessage := "newsUpdateMessage_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NewsUpdatesAPI.EditNewsUpdate(context.Background(), id).NewsUpdateMessage(newsUpdateMessage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsUpdatesAPI.EditNewsUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the news update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditNewsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newsUpdateMessage** | **string** |  | 

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


## ListNewsUpdates

> ListNewsUpdates200Response ListNewsUpdates(ctx).Limit(limit).Page(page).Execute()

List News Updates



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NewsUpdatesAPI.ListNewsUpdates(context.Background()).Limit(limit).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NewsUpdatesAPI.ListNewsUpdates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNewsUpdates`: ListNewsUpdates200Response
	fmt.Fprintf(os.Stdout, "Response from `NewsUpdatesAPI.ListNewsUpdates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNewsUpdatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. | 

### Return type

[**ListNewsUpdates200Response**](ListNewsUpdates200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

