# \UsersAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditCurrentUser**](UsersAPI.md#EditCurrentUser) | **Patch** /users/{id}.json | Edit Current User
[**FixUserCounts**](UsersAPI.md#FixUserCounts) | **Post** /maintenance/user/count_fixes.json | Fix User Counts
[**GetCurrentUser**](UsersAPI.md#GetCurrentUser) | **Get** /users/upload_limit.json | Get Current User
[**GetUser**](UsersAPI.md#GetUser) | **Get** /users/{id}.json | Get User
[**SearchUsers**](UsersAPI.md#SearchUsers) | **Get** /users.json | Search Users



## EditCurrentUser

> EditCurrentUser(ctx, id).UserCommentThreshold(userCommentThreshold).UserDefaultImageSize(userDefaultImageSize).UserFavoriteTags(userFavoriteTags).UserBlacklistedTags(userBlacklistedTags).UserTimeZone(userTimeZone).UserPerPage(userPerPage).UserCustomStyle(userCustomStyle).UserDescriptionCollapsedInitially(userDescriptionCollapsedInitially).UserHideComments(userHideComments).UserReceiveEmailNotifications(userReceiveEmailNotifications).UserEnableKeyboardNavigation(userEnableKeyboardNavigation).UserEnablePrivacyMode(userEnablePrivacyMode).UserDisableUserDmails(userDisableUserDmails).UserBlacklistUsers(userBlacklistUsers).UserShowPostStatistics(userShowPostStatistics).UserStyleUsernames(userStyleUsernames).UserShowHiddenComments(userShowHiddenComments).UserEnableAutocomplete(userEnableAutocomplete).UserDisableCroppedThumbnails(userDisableCroppedThumbnails).UserEnableSafeMode(userEnableSafeMode).UserDisableResponsiveMode(userDisableResponsiveMode).UserDmailFilterAttributesId(userDmailFilterAttributesId).UserDmailFilterAttributesWords(userDmailFilterAttributesWords).UserProfileAbout(userProfileAbout).UserProfileArtinfo(userProfileArtinfo).UserAvatarId(userAvatarId).UserEnableCompactUploader(userEnableCompactUploader).Execute()

Edit Current User

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
	id := float32(8.14) // float32 | The ID of the user. The actual value is ignored, but something must be supplied.
	userCommentThreshold := float32(8.14) // float32 |  (optional)
	userDefaultImageSize := "userDefaultImageSize_example" // string |  (optional)
	userFavoriteTags := "userFavoriteTags_example" // string |  (optional)
	userBlacklistedTags := "userBlacklistedTags_example" // string |  (optional)
	userTimeZone := "userTimeZone_example" // string | https://en.wikipedia.org/wiki/List_of_tz_database_time_zones (optional)
	userPerPage := float32(8.14) // float32 |  (optional)
	userCustomStyle := "userCustomStyle_example" // string |  (optional)
	userDescriptionCollapsedInitially := true // bool |  (optional)
	userHideComments := true // bool |  (optional)
	userReceiveEmailNotifications := true // bool |  (optional)
	userEnableKeyboardNavigation := true // bool |  (optional)
	userEnablePrivacyMode := true // bool |  (optional)
	userDisableUserDmails := true // bool |  (optional)
	userBlacklistUsers := true // bool |  (optional)
	userShowPostStatistics := true // bool |  (optional)
	userStyleUsernames := true // bool |  (optional)
	userShowHiddenComments := true // bool |  (optional)
	userEnableAutocomplete := true // bool |  (optional)
	userDisableCroppedThumbnails := true // bool |  (optional)
	userEnableSafeMode := true // bool |  (optional)
	userDisableResponsiveMode := true // bool |  (optional)
	userDmailFilterAttributesId := float32(8.14) // float32 |  (optional)
	userDmailFilterAttributesWords := "userDmailFilterAttributesWords_example" // string |  (optional)
	userProfileAbout := "userProfileAbout_example" // string |  (optional)
	userProfileArtinfo := "userProfileArtinfo_example" // string |  (optional)
	userAvatarId := float32(8.14) // float32 |  (optional)
	userEnableCompactUploader := true // bool | You must have uploaded at least 10 posts. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.EditCurrentUser(context.Background(), id).UserCommentThreshold(userCommentThreshold).UserDefaultImageSize(userDefaultImageSize).UserFavoriteTags(userFavoriteTags).UserBlacklistedTags(userBlacklistedTags).UserTimeZone(userTimeZone).UserPerPage(userPerPage).UserCustomStyle(userCustomStyle).UserDescriptionCollapsedInitially(userDescriptionCollapsedInitially).UserHideComments(userHideComments).UserReceiveEmailNotifications(userReceiveEmailNotifications).UserEnableKeyboardNavigation(userEnableKeyboardNavigation).UserEnablePrivacyMode(userEnablePrivacyMode).UserDisableUserDmails(userDisableUserDmails).UserBlacklistUsers(userBlacklistUsers).UserShowPostStatistics(userShowPostStatistics).UserStyleUsernames(userStyleUsernames).UserShowHiddenComments(userShowHiddenComments).UserEnableAutocomplete(userEnableAutocomplete).UserDisableCroppedThumbnails(userDisableCroppedThumbnails).UserEnableSafeMode(userEnableSafeMode).UserDisableResponsiveMode(userDisableResponsiveMode).UserDmailFilterAttributesId(userDmailFilterAttributesId).UserDmailFilterAttributesWords(userDmailFilterAttributesWords).UserProfileAbout(userProfileAbout).UserProfileArtinfo(userProfileArtinfo).UserAvatarId(userAvatarId).UserEnableCompactUploader(userEnableCompactUploader).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.EditCurrentUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the user. The actual value is ignored, but something must be supplied. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditCurrentUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userCommentThreshold** | **float32** |  | 
 **userDefaultImageSize** | **string** |  | 
 **userFavoriteTags** | **string** |  | 
 **userBlacklistedTags** | **string** |  | 
 **userTimeZone** | **string** | https://en.wikipedia.org/wiki/List_of_tz_database_time_zones | 
 **userPerPage** | **float32** |  | 
 **userCustomStyle** | **string** |  | 
 **userDescriptionCollapsedInitially** | **bool** |  | 
 **userHideComments** | **bool** |  | 
 **userReceiveEmailNotifications** | **bool** |  | 
 **userEnableKeyboardNavigation** | **bool** |  | 
 **userEnablePrivacyMode** | **bool** |  | 
 **userDisableUserDmails** | **bool** |  | 
 **userBlacklistUsers** | **bool** |  | 
 **userShowPostStatistics** | **bool** |  | 
 **userStyleUsernames** | **bool** |  | 
 **userShowHiddenComments** | **bool** |  | 
 **userEnableAutocomplete** | **bool** |  | 
 **userDisableCroppedThumbnails** | **bool** |  | 
 **userEnableSafeMode** | **bool** |  | 
 **userDisableResponsiveMode** | **bool** |  | 
 **userDmailFilterAttributesId** | **float32** |  | 
 **userDmailFilterAttributesWords** | **string** |  | 
 **userProfileAbout** | **string** |  | 
 **userProfileArtinfo** | **string** |  | 
 **userAvatarId** | **float32** |  | 
 **userEnableCompactUploader** | **bool** | You must have uploaded at least 10 posts. | 

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


## FixUserCounts

> FixUserCounts(ctx).Execute()

Fix User Counts

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.FixUserCounts(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.FixUserCounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFixUserCountsRequest struct via the builder pattern


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


## GetCurrentUser

> CurrentUser GetCurrentUser(ctx).Execute()

Get Current User



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetCurrentUser(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetCurrentUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCurrentUser`: CurrentUser
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetCurrentUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserRequest struct via the builder pattern


### Return type

[**CurrentUser**](CurrentUser.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> GetUser200Response GetUser(ctx, id).Execute()

Get User

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
	id := float32(8.14) // float32 | The ID of the user.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.GetUser(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUser`: GetUser200Response
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetUser200Response**](GetUser200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchUsers

> []SearchUsers200ResponseInner SearchUsers(ctx).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).SearchNameMatches(searchNameMatches).SearchAboutMe(searchAboutMe).SearchAvatarId(searchAvatarId).SearchLevel(searchLevel).SearchMinLevel(searchMinLevel).SearchMaxLevel(searchMaxLevel).SearchCanUploadFree(searchCanUploadFree).SearchCanApprovePosts(searchCanApprovePosts).SearchEmailMatches(searchEmailMatches).Execute()

Search Users

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
	searchNameMatches := "searchNameMatches_example" // string |  (optional)
	searchAboutMe := "searchAboutMe_example" // string |  (optional)
	searchAvatarId := float32(8.14) // float32 |  (optional)
	searchLevel := float32(8.14) // float32 |  (optional)
	searchMinLevel := float32(8.14) // float32 |  (optional)
	searchMaxLevel := float32(8.14) // float32 |  (optional)
	searchCanUploadFree := true // bool |  (optional)
	searchCanApprovePosts := true // bool |  (optional)
	searchEmailMatches := "searchEmailMatches_example" // string | You must be Admin+. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.SearchUsers(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchIpAddr(searchIpAddr).SearchOrder(searchOrder).SearchNameMatches(searchNameMatches).SearchAboutMe(searchAboutMe).SearchAvatarId(searchAvatarId).SearchLevel(searchLevel).SearchMinLevel(searchMinLevel).SearchMaxLevel(searchMaxLevel).SearchCanUploadFree(searchCanUploadFree).SearchCanApprovePosts(searchCanApprovePosts).SearchEmailMatches(searchEmailMatches).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.SearchUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchUsers`: []SearchUsers200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.SearchUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **float32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **float32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **float32** | Search for a specific id. | 
 **searchIpAddr** | **string** | Must be Admin+ to use. See [postgres&#39; documentation](https://www.postgresql.org/docs/9.3/functions-net.html) for information on how this is parsed. Specifically, \&quot;is contained within or equals\&quot; (&#x60;&lt;&lt;&#x3D;&#x60;). | 
 **searchOrder** | **string** |  | 
 **searchNameMatches** | **string** |  | 
 **searchAboutMe** | **string** |  | 
 **searchAvatarId** | **float32** |  | 
 **searchLevel** | **float32** |  | 
 **searchMinLevel** | **float32** |  | 
 **searchMaxLevel** | **float32** |  | 
 **searchCanUploadFree** | **bool** |  | 
 **searchCanApprovePosts** | **bool** |  | 
 **searchEmailMatches** | **string** | You must be Admin+. | 

### Return type

[**[]SearchUsers200ResponseInner**](SearchUsers200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

