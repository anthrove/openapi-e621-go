# \MascotsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMascot**](MascotsAPI.md#CreateMascot) | **Post** /mascots.json | Create Mascot
[**DeleteMascot**](MascotsAPI.md#DeleteMascot) | **Delete** /mascots/{id}.json | Delete Mascot
[**EditMascot**](MascotsAPI.md#EditMascot) | **Patch** /mascots/{id}.json | Edit Mascot
[**SearchMascots**](MascotsAPI.md#SearchMascots) | **Get** /mascots.json | Search Mascots



## CreateMascot

> Mascot CreateMascot(ctx).MascotMascotFile(mascotMascotFile).MascotDisplayName(mascotDisplayName).MascotBackgroundColor(mascotBackgroundColor).MascotArtistUrl(mascotArtistUrl).MascotArtistName(mascotArtistName).MascotAvailableOnString(mascotAvailableOnString).MascotActive(mascotActive).MascotHideAnonymous(mascotHideAnonymous).Execute()

Create Mascot



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
	mascotMascotFile := os.NewFile(1234, "some_file") // *os.File | 
	mascotDisplayName := "mascotDisplayName_example" // string | 
	mascotBackgroundColor := "mascotBackgroundColor_example" // string | 
	mascotArtistUrl := "mascotArtistUrl_example" // string | 
	mascotArtistName := "mascotArtistName_example" // string | 
	mascotAvailableOnString := "mascotAvailableOnString_example" // string | Comma separated site names. (optional)
	mascotActive := true // bool |  (optional)
	mascotHideAnonymous := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MascotsAPI.CreateMascot(context.Background()).MascotMascotFile(mascotMascotFile).MascotDisplayName(mascotDisplayName).MascotBackgroundColor(mascotBackgroundColor).MascotArtistUrl(mascotArtistUrl).MascotArtistName(mascotArtistName).MascotAvailableOnString(mascotAvailableOnString).MascotActive(mascotActive).MascotHideAnonymous(mascotHideAnonymous).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MascotsAPI.CreateMascot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMascot`: Mascot
	fmt.Fprintf(os.Stdout, "Response from `MascotsAPI.CreateMascot`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMascotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mascotMascotFile** | ***os.File** |  | 
 **mascotDisplayName** | **string** |  | 
 **mascotBackgroundColor** | **string** |  | 
 **mascotArtistUrl** | **string** |  | 
 **mascotArtistName** | **string** |  | 
 **mascotAvailableOnString** | **string** | Comma separated site names. | 
 **mascotActive** | **bool** |  | 
 **mascotHideAnonymous** | **bool** |  | 

### Return type

[**Mascot**](Mascot.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMascot

> DeleteMascot(ctx, id).Execute()

Delete Mascot



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
	id := float32(8.14) // float32 | The ID of the mascot.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MascotsAPI.DeleteMascot(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MascotsAPI.DeleteMascot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the mascot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMascotRequest struct via the builder pattern


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


## EditMascot

> EditMascot(ctx, id).MascotMascotFile(mascotMascotFile).MascotDisplayName(mascotDisplayName).MascotBackgroundColor(mascotBackgroundColor).MascotArtistUrl(mascotArtistUrl).MascotArtistName(mascotArtistName).MascotAvailableOnString(mascotAvailableOnString).MascotActive(mascotActive).MascotHideAnonymous(mascotHideAnonymous).Execute()

Edit Mascot



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
	id := float32(8.14) // float32 | The ID of the mascot.
	mascotMascotFile := os.NewFile(1234, "some_file") // *os.File |  (optional)
	mascotDisplayName := "mascotDisplayName_example" // string |  (optional)
	mascotBackgroundColor := "mascotBackgroundColor_example" // string |  (optional)
	mascotArtistUrl := "mascotArtistUrl_example" // string |  (optional)
	mascotArtistName := "mascotArtistName_example" // string |  (optional)
	mascotAvailableOnString := "mascotAvailableOnString_example" // string | Comma separated site names. (optional)
	mascotActive := true // bool |  (optional)
	mascotHideAnonymous := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MascotsAPI.EditMascot(context.Background(), id).MascotMascotFile(mascotMascotFile).MascotDisplayName(mascotDisplayName).MascotBackgroundColor(mascotBackgroundColor).MascotArtistUrl(mascotArtistUrl).MascotArtistName(mascotArtistName).MascotAvailableOnString(mascotAvailableOnString).MascotActive(mascotActive).MascotHideAnonymous(mascotHideAnonymous).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MascotsAPI.EditMascot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the mascot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditMascotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mascotMascotFile** | ***os.File** |  | 
 **mascotDisplayName** | **string** |  | 
 **mascotBackgroundColor** | **string** |  | 
 **mascotArtistUrl** | **string** |  | 
 **mascotArtistName** | **string** |  | 
 **mascotAvailableOnString** | **string** | Comma separated site names. | 
 **mascotActive** | **bool** |  | 
 **mascotHideAnonymous** | **bool** |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchMascots

> SearchMascots200Response SearchMascots(ctx).Limit(limit).Page(page).SearchId(searchId).Execute()

Search Mascots



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MascotsAPI.SearchMascots(context.Background()).Limit(limit).Page(page).SearchId(searchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MascotsAPI.SearchMascots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchMascots`: SearchMascots200Response
	fmt.Fprintf(os.Stdout, "Response from `MascotsAPI.SearchMascots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchMascotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **float32** | Search for a specific id. | 

### Return type

[**SearchMascots200Response**](SearchMascots200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

