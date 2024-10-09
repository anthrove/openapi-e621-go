# \BlipsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBlip**](BlipsAPI.md#CreateBlip) | **Post** /blips.json | Create Blip
[**DeleteBlip**](BlipsAPI.md#DeleteBlip) | **Delete** /blips/{id}.json | Delete Blip
[**EditBlip**](BlipsAPI.md#EditBlip) | **Patch** /blips/{id}.json | Edit Blip
[**GetBlip**](BlipsAPI.md#GetBlip) | **Get** /blips/{id}.json | Get Blip
[**HideBlip**](BlipsAPI.md#HideBlip) | **Post** /blips/{id}/hide.json | Hide Blip
[**MarkBlip**](BlipsAPI.md#MarkBlip) | **Post** /blips/{id}/warning.json | Mark Blip
[**SearchBlips**](BlipsAPI.md#SearchBlips) | **Get** /blips.json | Search Blips
[**UnhideBlip**](BlipsAPI.md#UnhideBlip) | **Post** /blips/{id}/unhide.json | Unhide Blip



## CreateBlip

> Blip CreateBlip(ctx).BlipBody(blipBody).BlipResponseTo(blipResponseTo).Execute()

Create Blip

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
	blipBody := "blipBody_example" // string | 
	blipResponseTo := float32(8.14) // float32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlipsAPI.CreateBlip(context.Background()).BlipBody(blipBody).BlipResponseTo(blipResponseTo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlipsAPI.CreateBlip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBlip`: Blip
	fmt.Fprintf(os.Stdout, "Response from `BlipsAPI.CreateBlip`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blipBody** | **string** |  | 
 **blipResponseTo** | **float32** |  | 

### Return type

[**Blip**](Blip.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBlip

> DeleteBlip(ctx, id).Execute()

Delete Blip



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
	id := float32(8.14) // float32 | The ID of the blip.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BlipsAPI.DeleteBlip(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlipsAPI.DeleteBlip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the blip. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBlipRequest struct via the builder pattern


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


## EditBlip

> Blip EditBlip(ctx, id).BlipBody(blipBody).Execute()

Edit Blip



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
	id := float32(8.14) // float32 | The ID of the blip.
	blipBody := "blipBody_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlipsAPI.EditBlip(context.Background(), id).BlipBody(blipBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlipsAPI.EditBlip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditBlip`: Blip
	fmt.Fprintf(os.Stdout, "Response from `BlipsAPI.EditBlip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the blip. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditBlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blipBody** | **string** |  | 

### Return type

[**Blip**](Blip.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBlip

> []Blip GetBlip(ctx, id).Execute()

Get Blip

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
	id := float32(8.14) // float32 | The ID of the blip to get.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlipsAPI.GetBlip(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlipsAPI.GetBlip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBlip`: []Blip
	fmt.Fprintf(os.Stdout, "Response from `BlipsAPI.GetBlip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the blip to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Blip**](Blip.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HideBlip

> Blip HideBlip(ctx, id).Execute()

Hide Blip



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
	id := float32(8.14) // float32 | The ID of the blip.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlipsAPI.HideBlip(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlipsAPI.HideBlip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HideBlip`: Blip
	fmt.Fprintf(os.Stdout, "Response from `BlipsAPI.HideBlip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the blip. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHideBlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Blip**](Blip.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkBlip

> DTextResponse MarkBlip(ctx, id).MarkBlipRequest(markBlipRequest).Execute()

Mark Blip



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
	id := float32(8.14) // float32 | The ID of the blip.
	markBlipRequest := *openapiclient.NewMarkBlipRequest("RecordType_example") // MarkBlipRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlipsAPI.MarkBlip(context.Background(), id).MarkBlipRequest(markBlipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlipsAPI.MarkBlip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkBlip`: DTextResponse
	fmt.Fprintf(os.Stdout, "Response from `BlipsAPI.MarkBlip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the blip. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkBlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **markBlipRequest** | [**MarkBlipRequest**](MarkBlipRequest.md) |  | 

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


## SearchBlips

> SearchBlips200Response SearchBlips(ctx).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchBodyMatches(searchBodyMatches).SearchResponseTo(searchResponseTo).Execute()

Search Blips



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
	searchId := int32(56) // int32 | Search for a specific id. (optional)
	searchIpAddr := "searchIpAddr_example" // string | Must be Admin+ to use. See [postgres' documentation](https://www.postgresql.org/docs/9.3/functions-net.html) for information on how this is parsed. Specifically, \"is contained within or equals\" (`<<=`). (optional)
	searchOrder := "searchOrder_example" // string |  (optional)
	searchCreatorId := float32(8.14) // float32 |  (optional)
	searchCreatorName := "searchCreatorName_example" // string |  (optional)
	searchBodyMatches := "searchBodyMatches_example" // string |  (optional)
	searchResponseTo := float32(8.14) // float32 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlipsAPI.SearchBlips(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchBodyMatches(searchBodyMatches).SearchResponseTo(searchResponseTo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlipsAPI.SearchBlips``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchBlips`: SearchBlips200Response
	fmt.Fprintf(os.Stdout, "Response from `BlipsAPI.SearchBlips`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchBlipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **int32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **int32** | Search for a specific id. | 
 **searchIpAddr** | **string** | Must be Admin+ to use. See [postgres&#39; documentation](https://www.postgresql.org/docs/9.3/functions-net.html) for information on how this is parsed. Specifically, \&quot;is contained within or equals\&quot; (&#x60;&lt;&lt;&#x3D;&#x60;). | 
 **searchOrder** | **string** |  | 
 **searchCreatorId** | **float32** |  | 
 **searchCreatorName** | **string** |  | 
 **searchBodyMatches** | **string** |  | 
 **searchResponseTo** | **float32** |  | 

### Return type

[**SearchBlips200Response**](SearchBlips200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnhideBlip

> Blip UnhideBlip(ctx, id).Execute()

Unhide Blip



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
	id := float32(8.14) // float32 | The ID of the blip.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BlipsAPI.UnhideBlip(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BlipsAPI.UnhideBlip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnhideBlip`: Blip
	fmt.Fprintf(os.Stdout, "Response from `BlipsAPI.UnhideBlip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the blip. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnhideBlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Blip**](Blip.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

