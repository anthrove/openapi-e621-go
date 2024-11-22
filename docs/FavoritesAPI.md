# \FavoritesAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFavorite**](FavoritesAPI.md#AddFavorite) | **Post** /favorites.json | Add Favorite
[**ListFavorites**](FavoritesAPI.md#ListFavorites) | **Get** /favorites.json | List Favorites
[**RemoveFavorite**](FavoritesAPI.md#RemoveFavorite) | **Delete** /favorites/{id}.json | Remove Favorite



## AddFavorite

> AddFavorite201Response AddFavorite(ctx).AddFavoriteRequest(addFavoriteRequest).Execute()

Add Favorite

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
	addFavoriteRequest := *openapiclient.NewAddFavoriteRequest(int32(123)) // AddFavoriteRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FavoritesAPI.AddFavorite(context.Background()).AddFavoriteRequest(addFavoriteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.AddFavorite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddFavorite`: AddFavorite201Response
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.AddFavorite`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddFavoriteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addFavoriteRequest** | [**AddFavoriteRequest**](AddFavoriteRequest.md) |  | 

### Return type

[**AddFavorite201Response**](AddFavorite201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFavorites

> ListFavorites200Response ListFavorites(ctx).Limit(limit).Page(page).UserId(userId).Execute()

List Favorites

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
	userId := int32(56) // int32 | You must be the user or Moderator+ if the user has their favorites hidden. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FavoritesAPI.ListFavorites(context.Background()).Limit(limit).Page(page).UserId(userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.ListFavorites``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFavorites`: ListFavorites200Response
	fmt.Fprintf(os.Stdout, "Response from `FavoritesAPI.ListFavorites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFavoritesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **int32** | The page number of results to get. Between 1 and 750. | 
 **userId** | **int32** | You must be the user or Moderator+ if the user has their favorites hidden. | 

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


## RemoveFavorite

> RemoveFavorite(ctx, id).Execute()

Remove Favorite

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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FavoritesAPI.RemoveFavorite(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FavoritesAPI.RemoveFavorite``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFavoriteRequest struct via the builder pattern


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

