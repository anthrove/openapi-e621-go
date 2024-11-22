# \UserFeedbacksAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserFeedback**](UserFeedbacksAPI.md#CreateUserFeedback) | **Post** /user_feedbacks.json | Create User Feedback
[**DeleteUserFeedback**](UserFeedbacksAPI.md#DeleteUserFeedback) | **Put** /user_feedbacks/{id}/delete.json | Delete User Feedback
[**DestroyUserFeedback**](UserFeedbacksAPI.md#DestroyUserFeedback) | **Delete** /user_feedbacks/{id}.json | Destroy User Feedback
[**EditUserFeedback**](UserFeedbacksAPI.md#EditUserFeedback) | **Patch** /user_feedbacks/{id}.json | Edit User Feedback
[**GetUserFeedback**](UserFeedbacksAPI.md#GetUserFeedback) | **Get** /user_feedbacks/{id}.json | Get User Feedback
[**SearchUserFeedbacks**](UserFeedbacksAPI.md#SearchUserFeedbacks) | **Get** /user_feedbacks.json | Search User Feedbacks
[**UndeleteUserFeedback**](UserFeedbacksAPI.md#UndeleteUserFeedback) | **Put** /user_feedbacks/{id}/undelete.json | Undelete User Feedback



## CreateUserFeedback

> UserFeedback CreateUserFeedback(ctx).UserFeedbackBody(userFeedbackBody).UserFeedbackCategory(userFeedbackCategory).UserFeedbackUserId(userFeedbackUserId).UserFeedbackUserName(userFeedbackUserName).Execute()

Create User Feedback



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
	userFeedbackBody := "userFeedbackBody_example" // string | 
	userFeedbackCategory := openapiclient.FeedbackCategories("negative") // FeedbackCategories | 
	userFeedbackUserId := int32(56) // int32 |  (optional)
	userFeedbackUserName := "userFeedbackUserName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserFeedbacksAPI.CreateUserFeedback(context.Background()).UserFeedbackBody(userFeedbackBody).UserFeedbackCategory(userFeedbackCategory).UserFeedbackUserId(userFeedbackUserId).UserFeedbackUserName(userFeedbackUserName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserFeedbacksAPI.CreateUserFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserFeedback`: UserFeedback
	fmt.Fprintf(os.Stdout, "Response from `UserFeedbacksAPI.CreateUserFeedback`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userFeedbackBody** | **string** |  | 
 **userFeedbackCategory** | [**FeedbackCategories**](FeedbackCategories.md) |  | 
 **userFeedbackUserId** | **int32** |  | 
 **userFeedbackUserName** | **string** |  | 

### Return type

[**UserFeedback**](UserFeedback.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserFeedback

> DeleteUserFeedback(ctx, id).Execute()

Delete User Feedback



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
	id := int32(56) // int32 | The ID of the feedback.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserFeedbacksAPI.DeleteUserFeedback(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserFeedbacksAPI.DeleteUserFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the feedback. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserFeedbackRequest struct via the builder pattern


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


## DestroyUserFeedback

> DestroyUserFeedback(ctx, id).Execute()

Destroy User Feedback



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
	id := int32(56) // int32 | The ID of the feedback.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserFeedbacksAPI.DestroyUserFeedback(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserFeedbacksAPI.DestroyUserFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the feedback. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroyUserFeedbackRequest struct via the builder pattern


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


## EditUserFeedback

> EditUserFeedback(ctx, id).UserFeedbackBody(userFeedbackBody).UserFeedbackCategory(userFeedbackCategory).Execute()

Edit User Feedback



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
	id := int32(56) // int32 | The ID of the feedback.
	userFeedbackBody := "userFeedbackBody_example" // string |  (optional)
	userFeedbackCategory := openapiclient.FeedbackCategories("negative") // FeedbackCategories |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserFeedbacksAPI.EditUserFeedback(context.Background(), id).UserFeedbackBody(userFeedbackBody).UserFeedbackCategory(userFeedbackCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserFeedbacksAPI.EditUserFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the feedback. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditUserFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userFeedbackBody** | **string** |  | 
 **userFeedbackCategory** | [**FeedbackCategories**](FeedbackCategories.md) |  | 

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


## GetUserFeedback

> UserFeedback GetUserFeedback(ctx, id).Execute()

Get User Feedback



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
	id := int32(56) // int32 | The ID of the feedback.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserFeedbacksAPI.GetUserFeedback(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserFeedbacksAPI.GetUserFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserFeedback`: UserFeedback
	fmt.Fprintf(os.Stdout, "Response from `UserFeedbacksAPI.GetUserFeedback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the feedback. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserFeedback**](UserFeedback.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUserFeedbacks

> SearchUserFeedbacks200Response SearchUserFeedbacks(ctx).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchDeleted(searchDeleted).SearchBodyMatches(searchBodyMatches).SearchUserId(searchUserId).SearchUserName(searchUserName).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchCategory(searchCategory).Execute()

Search User Feedbacks



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
	searchOrder := "searchOrder_example" // string |  (optional)
	searchDeleted := "searchDeleted_example" // string | You must be Moderator+. (optional)
	searchBodyMatches := "searchBodyMatches_example" // string |  (optional)
	searchUserId := int32(56) // int32 |  (optional)
	searchUserName := "searchUserName_example" // string |  (optional)
	searchCreatorId := int32(56) // int32 |  (optional)
	searchCreatorName := "searchCreatorName_example" // string |  (optional)
	searchCategory := openapiclient.FeedbackCategories("negative") // FeedbackCategories |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserFeedbacksAPI.SearchUserFeedbacks(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchDeleted(searchDeleted).SearchBodyMatches(searchBodyMatches).SearchUserId(searchUserId).SearchUserName(searchUserName).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).SearchCategory(searchCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserFeedbacksAPI.SearchUserFeedbacks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUserFeedbacks`: SearchUserFeedbacks200Response
	fmt.Fprintf(os.Stdout, "Response from `UserFeedbacksAPI.SearchUserFeedbacks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUserFeedbacksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **int32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **int32** | Search for a specific id. | 
 **searchOrder** | **string** |  | 
 **searchDeleted** | **string** | You must be Moderator+. | 
 **searchBodyMatches** | **string** |  | 
 **searchUserId** | **int32** |  | 
 **searchUserName** | **string** |  | 
 **searchCreatorId** | **int32** |  | 
 **searchCreatorName** | **string** |  | 
 **searchCategory** | [**FeedbackCategories**](FeedbackCategories.md) |  | 

### Return type

[**SearchUserFeedbacks200Response**](SearchUserFeedbacks200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UndeleteUserFeedback

> UndeleteUserFeedback(ctx, id).Execute()

Undelete User Feedback



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
	id := int32(56) // int32 | The ID of the feedback.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserFeedbacksAPI.UndeleteUserFeedback(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserFeedbacksAPI.UndeleteUserFeedback``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the feedback. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUndeleteUserFeedbackRequest struct via the builder pattern


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

