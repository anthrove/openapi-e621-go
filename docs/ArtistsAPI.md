# \ArtistsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateArtist**](ArtistsAPI.md#CreateArtist) | **Post** /artists.json | Create Artist
[**DeleteArtist**](ArtistsAPI.md#DeleteArtist) | **Delete** /artists/{idOrName}.json | Delete Artist
[**EditArtist**](ArtistsAPI.md#EditArtist) | **Patch** /artists/{idOrName}.json | Edit Artist
[**GetArtist**](ArtistsAPI.md#GetArtist) | **Get** /artists/{idOrName}.json | Get Artist
[**RevertArtist**](ArtistsAPI.md#RevertArtist) | **Put** /artists/{idOrName}/revert.json | Revert Artist
[**SearchArtists**](ArtistsAPI.md#SearchArtists) | **Get** /artists.json | Search Artists



## CreateArtist

> Artist CreateArtist(ctx).ArtistName(artistName).ArtistOtherNames(artistOtherNames).ArtistOtherNamesString(artistOtherNamesString).ArtistUrlString(artistUrlString).ArtistNotes(artistNotes).ArtistGroupName(artistGroupName).ArtistLinkedUserId(artistLinkedUserId).ArtistIsLocked(artistIsLocked).Execute()

Create Artist



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
	artistName := "artistName_example" // string | 
	artistOtherNames := []string{"Inner_example"} // []string |  (optional)
	artistOtherNamesString := "artistOtherNamesString_example" // string |  (optional)
	artistUrlString := "artistUrlString_example" // string |  (optional)
	artistNotes := "artistNotes_example" // string |  (optional)
	artistGroupName := "artistGroupName_example" // string |  (optional)
	artistLinkedUserId := float32(8.14) // float32 | Only usable for Janitor+ (optional)
	artistIsLocked := true // bool | Only usable for Janitor+ (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtistsAPI.CreateArtist(context.Background()).ArtistName(artistName).ArtistOtherNames(artistOtherNames).ArtistOtherNamesString(artistOtherNamesString).ArtistUrlString(artistUrlString).ArtistNotes(artistNotes).ArtistGroupName(artistGroupName).ArtistLinkedUserId(artistLinkedUserId).ArtistIsLocked(artistIsLocked).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtistsAPI.CreateArtist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateArtist`: Artist
	fmt.Fprintf(os.Stdout, "Response from `ArtistsAPI.CreateArtist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateArtistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artistName** | **string** |  | 
 **artistOtherNames** | **[]string** |  | 
 **artistOtherNamesString** | **string** |  | 
 **artistUrlString** | **string** |  | 
 **artistNotes** | **string** |  | 
 **artistGroupName** | **string** |  | 
 **artistLinkedUserId** | **float32** | Only usable for Janitor+ | 
 **artistIsLocked** | **bool** | Only usable for Janitor+ | 

### Return type

[**Artist**](Artist.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteArtist

> DeleteArtist(ctx, idOrName).Execute()

Delete Artist



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
	idOrName := GetArtistIdOrNameParameter(8.14) // GetArtistIdOrNameParameter | The ID or name of the artist to edit.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArtistsAPI.DeleteArtist(context.Background(), idOrName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtistsAPI.DeleteArtist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrName** | **GetArtistIdOrNameParameter** | The ID or name of the artist to edit. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArtistRequest struct via the builder pattern


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


## EditArtist

> EditArtist(ctx, idOrName).EditArtistRequest(editArtistRequest).Execute()

Edit Artist



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
	idOrName := GetArtistIdOrNameParameter(8.14) // GetArtistIdOrNameParameter | The ID or name of the artist to edit.
	editArtistRequest := *openapiclient.NewEditArtistRequest() // EditArtistRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArtistsAPI.EditArtist(context.Background(), idOrName).EditArtistRequest(editArtistRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtistsAPI.EditArtist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrName** | **GetArtistIdOrNameParameter** | The ID or name of the artist to edit. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditArtistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **editArtistRequest** | [**EditArtistRequest**](EditArtistRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/c-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtist

> GetArtist200Response GetArtist(ctx, idOrName).Execute()

Get Artist

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
	idOrName := GetArtistIdOrNameParameter(8.14) // GetArtistIdOrNameParameter | The ID or name of the artist to get.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtistsAPI.GetArtist(context.Background(), idOrName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtistsAPI.GetArtist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtist`: GetArtist200Response
	fmt.Fprintf(os.Stdout, "Response from `ArtistsAPI.GetArtist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrName** | **GetArtistIdOrNameParameter** | The ID or name of the artist to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetArtist200Response**](GetArtist200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevertArtist

> RevertArtist(ctx, idOrName).VersionId(versionId).Execute()

Revert Artist



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
	idOrName := GetArtistIdOrNameParameter(8.14) // GetArtistIdOrNameParameter | The ID or name of the artist to revert.
	versionId := float32(8.14) // float32 | The version ID to revert to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArtistsAPI.RevertArtist(context.Background(), idOrName).VersionId(versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtistsAPI.RevertArtist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrName** | **GetArtistIdOrNameParameter** | The ID or name of the artist to revert. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevertArtistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **versionId** | **float32** | The version ID to revert to. | 

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


## SearchArtists

> []SearchArtists200ResponseInner SearchArtists(ctx).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchName(searchName).SearchGroupName(searchGroupName).SearchAnyOtherNameLike(searchAnyOtherNameLike).SearchAnyNameMatches(searchAnyNameMatches).SearchAnyNameOrUrlMatches(searchAnyNameOrUrlMatches).SearchUrlMatches(searchUrlMatches).SearchCreatorName(searchCreatorName).SearchCreatorId(searchCreatorId).SearchHasTag(searchHasTag).SearchIsLinked(searchIsLinked).Execute()

Search Artists

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
	searchOrder := "searchOrder_example" // string |  (optional)
	searchName := "searchName_example" // string |  (optional)
	searchGroupName := "searchGroupName_example" // string |  (optional)
	searchAnyOtherNameLike := "searchAnyOtherNameLike_example" // string |  (optional)
	searchAnyNameMatches := "searchAnyNameMatches_example" // string |  (optional)
	searchAnyNameOrUrlMatches := "searchAnyNameOrUrlMatches_example" // string |  (optional)
	searchUrlMatches := "searchUrlMatches_example" // string |  (optional)
	searchCreatorName := "searchCreatorName_example" // string |  (optional)
	searchCreatorId := "searchCreatorId_example" // string |  (optional)
	searchHasTag := "searchHasTag_example" // string |  (optional)
	searchIsLinked := "searchIsLinked_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtistsAPI.SearchArtists(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchName(searchName).SearchGroupName(searchGroupName).SearchAnyOtherNameLike(searchAnyOtherNameLike).SearchAnyNameMatches(searchAnyNameMatches).SearchAnyNameOrUrlMatches(searchAnyNameOrUrlMatches).SearchUrlMatches(searchUrlMatches).SearchCreatorName(searchCreatorName).SearchCreatorId(searchCreatorId).SearchHasTag(searchHasTag).SearchIsLinked(searchIsLinked).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtistsAPI.SearchArtists``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchArtists`: []SearchArtists200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `ArtistsAPI.SearchArtists`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchArtistsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **float32** | Search for a specific id. | 
 **searchOrder** | **string** |  | 
 **searchName** | **string** |  | 
 **searchGroupName** | **string** |  | 
 **searchAnyOtherNameLike** | **string** |  | 
 **searchAnyNameMatches** | **string** |  | 
 **searchAnyNameOrUrlMatches** | **string** |  | 
 **searchUrlMatches** | **string** |  | 
 **searchCreatorName** | **string** |  | 
 **searchCreatorId** | **string** |  | 
 **searchHasTag** | **string** |  | 
 **searchIsLinked** | **string** |  | 

### Return type

[**[]SearchArtists200ResponseInner**](SearchArtists200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

