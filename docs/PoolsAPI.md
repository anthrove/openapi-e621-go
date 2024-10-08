# \PoolsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPostToPool**](PoolsAPI.md#AddPostToPool) | **Post** /pool_element.js | Add Post To Pool
[**CreatePool**](PoolsAPI.md#CreatePool) | **Post** /pools.json | Create Pool
[**DeletePool**](PoolsAPI.md#DeletePool) | **Delete** /pools/{id}.json | Delete Pool
[**EditPool**](PoolsAPI.md#EditPool) | **Patch** /pools/{id}.json | Edit Pool
[**GetPool**](PoolsAPI.md#GetPool) | **Get** /pools/{id}.json | Get Pool
[**RemovePostFromPool**](PoolsAPI.md#RemovePostFromPool) | **Delete** /pool_element.json | Remove Post From Pool
[**RevertPool**](PoolsAPI.md#RevertPool) | **Put** /pools/{id}/revert.json | Revert Pool
[**SearchPools**](PoolsAPI.md#SearchPools) | **Get** /pools.json | Search Pools



## AddPostToPool

> AddPostToPool(ctx).PostId(postId).PoolId(poolId).PoolName(poolName).Execute()

Add Post To Pool



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
	postId := float32(8.14) // float32 | 
	poolId := float32(8.14) // float32 | Mutually exclusive with pool_name. (optional)
	poolName := "poolName_example" // string | Mutually exclusive with pool_id. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PoolsAPI.AddPostToPool(context.Background()).PostId(postId).PoolId(poolId).PoolName(poolName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.AddPostToPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddPostToPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postId** | **float32** |  | 
 **poolId** | **float32** | Mutually exclusive with pool_name. | 
 **poolName** | **string** | Mutually exclusive with pool_id. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePool

> Pool CreatePool(ctx).PoolName(poolName).PoolDescription(poolDescription).PoolCategory(poolCategory).IpoolSActive(ipoolSActive).PoolPostIdsString(poolPostIdsString).PoolPostIds(poolPostIds).Execute()

Create Pool

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
	poolName := "poolName_example" // string | 
	poolDescription := "poolDescription_example" // string |  (optional)
	poolCategory := openapiclient.PoolCategories("collection") // PoolCategories |  (optional)
	ipoolSActive := true // bool |  (optional)
	poolPostIdsString := "poolPostIdsString_example" // string | Space separated list of post IDs. Mutually exclusive with post_ids. (optional)
	poolPostIds := []float32{float32(123)} // []float32 | Array of post IDs. Mutually exclusive with post_ids_string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.CreatePool(context.Background()).PoolName(poolName).PoolDescription(poolDescription).PoolCategory(poolCategory).IpoolSActive(ipoolSActive).PoolPostIdsString(poolPostIdsString).PoolPostIds(poolPostIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.CreatePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePool`: Pool
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.CreatePool`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolName** | **string** |  | 
 **poolDescription** | **string** |  | 
 **poolCategory** | [**PoolCategories**](PoolCategories.md) |  | 
 **ipoolSActive** | **bool** |  | 
 **poolPostIdsString** | **string** | Space separated list of post IDs. Mutually exclusive with post_ids. | 
 **poolPostIds** | **[]float32** | Array of post IDs. Mutually exclusive with post_ids_string. | 

### Return type

[**Pool**](Pool.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePool

> DeletePool(ctx, id).Execute()

Delete Pool



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
	id := float32(8.14) // float32 | The ID of the pool.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PoolsAPI.DeletePool(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.DeletePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePoolRequest struct via the builder pattern


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


## EditPool

> EditPool(ctx, id).PoolName(poolName).PoolDescription(poolDescription).PoolIsActive(poolIsActive).PoolCategory(poolCategory).PoolPostIdsString(poolPostIdsString).PoolPostIds(poolPostIds).Execute()

Edit Pool

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
	id := float32(8.14) // float32 | The ID of the pool.
	poolName := "poolName_example" // string |  (optional)
	poolDescription := "poolDescription_example" // string |  (optional)
	poolIsActive := true // bool |  (optional)
	poolCategory := openapiclient.PoolCategories("collection") // PoolCategories | If the pool has more than 30 posts, you must be Janitor+. (optional)
	poolPostIdsString := "poolPostIdsString_example" // string | Space separated list of post IDs. Mutually exclusive with post_ids. (optional)
	poolPostIds := []float32{float32(123)} // []float32 | Array of post IDs. Mutually exclusive with post_ids_string. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PoolsAPI.EditPool(context.Background(), id).PoolName(poolName).PoolDescription(poolDescription).PoolIsActive(poolIsActive).PoolCategory(poolCategory).PoolPostIdsString(poolPostIdsString).PoolPostIds(poolPostIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.EditPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **poolName** | **string** |  | 
 **poolDescription** | **string** |  | 
 **poolIsActive** | **bool** |  | 
 **poolCategory** | [**PoolCategories**](PoolCategories.md) | If the pool has more than 30 posts, you must be Janitor+. | 
 **poolPostIdsString** | **string** | Space separated list of post IDs. Mutually exclusive with post_ids. | 
 **poolPostIds** | **[]float32** | Array of post IDs. Mutually exclusive with post_ids_string. | 

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


## GetPool

> Pool GetPool(ctx, id).Execute()

Get Pool

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
	id := float32(8.14) // float32 | The ID of the pool.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.GetPool(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.GetPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPool`: Pool
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.GetPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Pool**](Pool.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePostFromPool

> RemovePostFromPool(ctx).PoolId(poolId).PostId(postId).Execute()

Remove Post From Pool

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
	poolId := float32(8.14) // float32 | 
	postId := float32(8.14) // float32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PoolsAPI.RemovePostFromPool(context.Background()).PoolId(poolId).PostId(postId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.RemovePostFromPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemovePostFromPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolId** | **float32** |  | 
 **postId** | **float32** |  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevertPool

> RevertPool(ctx, id).VersionId(versionId).Execute()

Revert Pool

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
	id := float32(8.14) // float32 | The ID of the pool.
	versionId := float32(8.14) // float32 | The version ID to revert to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PoolsAPI.RevertPool(context.Background(), id).VersionId(versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.RevertPool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the pool. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevertPoolRequest struct via the builder pattern


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


## SearchPools

> []Pool SearchPools(ctx).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchNameMatches(searchNameMatches).SearchDescriptionMatches(searchDescriptionMatches).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchCategory(searchCategory).SearchIsActive(searchIsActive).Execute()

Search Pools

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
	searchId := float32(8.14) // float32 | Search for a specific id. (optional)
	searchOrder := "searchOrder_example" // string |  (optional)
	searchNameMatches := "searchNameMatches_example" // string |  (optional)
	searchDescriptionMatches := "searchDescriptionMatches_example" // string |  (optional)
	searchCreatorId := float32(8.14) // float32 |  (optional)
	searchCreatorName := "searchCreatorName_example" // string |  (optional)
	searchCategory := openapiclient.PoolCategories("collection") // PoolCategories |  (optional)
	searchIsActive := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PoolsAPI.SearchPools(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchNameMatches(searchNameMatches).SearchDescriptionMatches(searchDescriptionMatches).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchCategory(searchCategory).SearchIsActive(searchIsActive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PoolsAPI.SearchPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPools`: []Pool
	fmt.Fprintf(os.Stdout, "Response from `PoolsAPI.SearchPools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **float32** | Search for a specific id. | 
 **searchOrder** | **string** |  | 
 **searchNameMatches** | **string** |  | 
 **searchDescriptionMatches** | **string** |  | 
 **searchCreatorId** | **float32** |  | 
 **searchCreatorName** | **string** |  | 
 **searchCategory** | [**PoolCategories**](PoolCategories.md) |  | 
 **searchIsActive** | **bool** |  | 

### Return type

[**[]Pool**](Pool.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

