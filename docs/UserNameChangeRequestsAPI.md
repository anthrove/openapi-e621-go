# \UserNameChangeRequestsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserNameChangeRequest**](UserNameChangeRequestsAPI.md#CreateUserNameChangeRequest) | **Post** /user_name_change_requests.json | Create User Name Change Request
[**GetUserNameChangeRequest**](UserNameChangeRequestsAPI.md#GetUserNameChangeRequest) | **Get** /user_name_change_requests/{id}.json | Get User Name Change Request
[**SearchUserNameChangeRequests**](UserNameChangeRequestsAPI.md#SearchUserNameChangeRequests) | **Get** /user_name_change_requests.json | Search User Name Change Requests



## CreateUserNameChangeRequest

> CreateUserNameChangeRequest(ctx).UserNameChangeRequestDesiredName(userNameChangeRequestDesiredName).UserNameChangeRequestChangeReason(userNameChangeRequestChangeReason).Execute()

Create User Name Change Request



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
	userNameChangeRequestDesiredName := "userNameChangeRequestDesiredName_example" // string | 
	userNameChangeRequestChangeReason := "userNameChangeRequestChangeReason_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserNameChangeRequestsAPI.CreateUserNameChangeRequest(context.Background()).UserNameChangeRequestDesiredName(userNameChangeRequestDesiredName).UserNameChangeRequestChangeReason(userNameChangeRequestChangeReason).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserNameChangeRequestsAPI.CreateUserNameChangeRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserNameChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userNameChangeRequestDesiredName** | **string** |  | 
 **userNameChangeRequestChangeReason** | **string** |  | 

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


## GetUserNameChangeRequest

> UserNameChangeRequest GetUserNameChangeRequest(ctx, id).Execute()

Get User Name Change Request



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
	id := float32(8.14) // float32 | The ID of the name change request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserNameChangeRequestsAPI.GetUserNameChangeRequest(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserNameChangeRequestsAPI.GetUserNameChangeRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserNameChangeRequest`: UserNameChangeRequest
	fmt.Fprintf(os.Stdout, "Response from `UserNameChangeRequestsAPI.GetUserNameChangeRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the name change request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserNameChangeRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserNameChangeRequest**](UserNameChangeRequest.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUserNameChangeRequests

> SearchUserNameChangeRequests200Response SearchUserNameChangeRequests(ctx).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchCurrentId(searchCurrentId).SearchCurrentName(searchCurrentName).SearchOriginalName(searchOriginalName).SearchDesiredName(searchDesiredName).Execute()

Search User Name Change Requests



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
	searchCurrentId := float32(8.14) // float32 |  (optional)
	searchCurrentName := "searchCurrentName_example" // string |  (optional)
	searchOriginalName := "searchOriginalName_example" // string |  (optional)
	searchDesiredName := "searchDesiredName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserNameChangeRequestsAPI.SearchUserNameChangeRequests(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchCurrentId(searchCurrentId).SearchCurrentName(searchCurrentName).SearchOriginalName(searchOriginalName).SearchDesiredName(searchDesiredName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserNameChangeRequestsAPI.SearchUserNameChangeRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUserNameChangeRequests`: SearchUserNameChangeRequests200Response
	fmt.Fprintf(os.Stdout, "Response from `UserNameChangeRequestsAPI.SearchUserNameChangeRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUserNameChangeRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **float32** | Search for a specific id. | 
 **searchOrder** | **string** |  | 
 **searchCurrentId** | **float32** |  | 
 **searchCurrentName** | **string** |  | 
 **searchOriginalName** | **string** |  | 
 **searchDesiredName** | **string** |  | 

### Return type

[**SearchUserNameChangeRequests200Response**](SearchUserNameChangeRequests200Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

