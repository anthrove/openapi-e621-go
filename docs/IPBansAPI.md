# \IPBansAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIPBan**](IPBansAPI.md#CreateIPBan) | **Post** /ip_bans.json | Create IP Ban
[**DeleteIPBan**](IPBansAPI.md#DeleteIPBan) | **Delete** /ip_bans/{id}.json | Delete IP Ban
[**ListIPBans**](IPBansAPI.md#ListIPBans) | **Get** /ip_bans.json | List IP Bans



## CreateIPBan

> IPBan CreateIPBan(ctx).IpBanIpAddr(ipBanIpAddr).IpBanReason(ipBanReason).Execute()

Create IP Ban



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
	ipBanIpAddr := "ipBanIpAddr_example" // string | 
	ipBanReason := "ipBanReason_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPBansAPI.CreateIPBan(context.Background()).IpBanIpAddr(ipBanIpAddr).IpBanReason(ipBanReason).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPBansAPI.CreateIPBan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIPBan`: IPBan
	fmt.Fprintf(os.Stdout, "Response from `IPBansAPI.CreateIPBan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIPBanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ipBanIpAddr** | **string** |  | 
 **ipBanReason** | **string** |  | 

### Return type

[**IPBan**](IPBan.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIPBan

> DeleteIPBan(ctx, id).Execute()

Delete IP Ban



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
	id := float32(8.14) // float32 | The ID of the ip ban.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IPBansAPI.DeleteIPBan(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPBansAPI.DeleteIPBan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the ip ban. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIPBanRequest struct via the builder pattern


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


## ListIPBans

> ListIPBans200Response ListIPBans(ctx).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).SearchBannerId(searchBannerId).SearchBannerName(searchBannerName).SearchReason(searchReason).Execute()

List IP Bans



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
	searchIpAddr := "searchIpAddr_example" // string | Must be Admin+ to use. See [postgres' documentation](https://www.postgresql.org/docs/9.3/functions-net.html) for information on how this is parsed. Specifically, \"is contained within or equals\" (`<<=`). (optional)
	searchOrder := "searchOrder_example" // string |  (optional)
	searchBannerId := float32(8.14) // float32 |  (optional)
	searchBannerName := "searchBannerName_example" // string |  (optional)
	searchReason := "searchReason_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPBansAPI.ListIPBans(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).SearchBannerId(searchBannerId).SearchBannerName(searchBannerName).SearchReason(searchReason).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPBansAPI.ListIPBans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListIPBans`: ListIPBans200Response
	fmt.Fprintf(os.Stdout, "Response from `IPBansAPI.ListIPBans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIPBansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **float32** | Search for a specific id. | 
 **searchIpAddr** | **string** | Must be Admin+ to use. See [postgres&#39; documentation](https://www.postgresql.org/docs/9.3/functions-net.html) for information on how this is parsed. Specifically, \&quot;is contained within or equals\&quot; (&#x60;&lt;&lt;&#x3D;&#x60;). | 
 **searchOrder** | **string** |  | 
 **searchBannerId** | **float32** |  | 
 **searchBannerName** | **string** |  | 
 **searchReason** | **string** |  | 

### Return type

[**ListIPBans200Response**](ListIPBans200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

