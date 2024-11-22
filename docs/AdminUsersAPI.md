# \AdminUsersAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminEditUser**](AdminUsersAPI.md#AdminEditUser) | **Patch** /admins/users/{id}.json | Admin Edit User
[**GetAltList**](AdminUsersAPI.md#GetAltList) | **Get** /admin/users/alt_list.json | Get Alt List



## AdminEditUser

> AdminEditUser(ctx, id).UserVerified(userVerified).UserLevel(userLevel).UserName(userName).UserProfileAbout(userProfileAbout).UserProfileArtinfo(userProfileArtinfo).UserBaseUploadLimit(userBaseUploadLimit).UserEnablePrivacyMode(userEnablePrivacyMode).UserEmail(userEmail).UserCanApprovePosts(userCanApprovePosts).UserCanUploadFree(userCanUploadFree).UserNoFlagging(userNoFlagging).UserReplacementsBeta(userReplacementsBeta).Execute()

Admin Edit User



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
	id := int32(56) // int32 | The ID of the user.
	userVerified := true // bool | Must have the bd staff user flag to use. (optional)
	userLevel := int32(56) // int32 | Must have the bd staff user flag to promote to Admin+. (optional)
	userName := "userName_example" // string |  (optional)
	userProfileAbout := "userProfileAbout_example" // string |  (optional)
	userProfileArtinfo := "userProfileArtinfo_example" // string |  (optional)
	userBaseUploadLimit := float32(8.14) // float32 |  (optional)
	userEnablePrivacyMode := true // bool |  (optional)
	userEmail := "userEmail_example" // string | Must have the bd staff user flag to use. (optional)
	userCanApprovePosts := true // bool |  (optional)
	userCanUploadFree := true // bool |  (optional)
	userNoFlagging := true // bool |  (optional)
	userReplacementsBeta := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AdminUsersAPI.AdminEditUser(context.Background(), id).UserVerified(userVerified).UserLevel(userLevel).UserName(userName).UserProfileAbout(userProfileAbout).UserProfileArtinfo(userProfileArtinfo).UserBaseUploadLimit(userBaseUploadLimit).UserEnablePrivacyMode(userEnablePrivacyMode).UserEmail(userEmail).UserCanApprovePosts(userCanApprovePosts).UserCanUploadFree(userCanUploadFree).UserNoFlagging(userNoFlagging).UserReplacementsBeta(userReplacementsBeta).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminUsersAPI.AdminEditUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminEditUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userVerified** | **bool** | Must have the bd staff user flag to use. | 
 **userLevel** | **int32** | Must have the bd staff user flag to promote to Admin+. | 
 **userName** | **string** |  | 
 **userProfileAbout** | **string** |  | 
 **userProfileArtinfo** | **string** |  | 
 **userBaseUploadLimit** | **float32** |  | 
 **userEnablePrivacyMode** | **bool** |  | 
 **userEmail** | **string** | Must have the bd staff user flag to use. | 
 **userCanApprovePosts** | **bool** |  | 
 **userCanUploadFree** | **bool** |  | 
 **userNoFlagging** | **bool** |  | 
 **userReplacementsBeta** | **bool** |  | 

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


## GetAltList

> []GetAltList200ResponseInner GetAltList(ctx).Page(page).Execute()

Get Alt List



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
	page := int32(56) // int32 | The page number of results to get. Between 1 and 9999. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AdminUsersAPI.GetAltList(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AdminUsersAPI.GetAltList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAltList`: []GetAltList200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `AdminUsersAPI.GetAltList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAltListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page number of results to get. Between 1 and 9999. | 

### Return type

[**[]GetAltList200ResponseInner**](GetAltList200ResponseInner.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

