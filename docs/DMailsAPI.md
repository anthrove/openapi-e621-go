# \DMailsAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDMail**](DMailsAPI.md#DeleteDMail) | **Delete** /dmails/{id}.json | Delete DMail
[**GetDMail**](DMailsAPI.md#GetDMail) | **Get** /dmails/{id}.json | Get DMail
[**MarkAllDMailsAsUnread**](DMailsAPI.md#MarkAllDMailsAsUnread) | **Put** /dmails/mark_all_as_unread.json | Mark All DMails As Unread
[**MarkDMailAsRead**](DMailsAPI.md#MarkDMailAsRead) | **Put** /dmails/{id}/mark_as_read.json | Mark DMail As Read
[**MarkDMailAsUnread**](DMailsAPI.md#MarkDMailAsUnread) | **Put** /dmails/{id}/mark_as_unread.json | Mark DMail As Unread
[**SearchDMails**](DMailsAPI.md#SearchDMails) | **Get** /dmails.json | Search DMails
[**UpdateUserDmailFilter**](DMailsAPI.md#UpdateUserDmailFilter) | **Patch** /maintenance/user/dmail_filter.json | Update User Dmail Filter



## DeleteDMail

> DeleteDMail(ctx, id).Execute()

Delete DMail



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
	id := int32(56) // int32 | The ID of the dmail.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DMailsAPI.DeleteDMail(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DMailsAPI.DeleteDMail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the dmail. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDMailRequest struct via the builder pattern


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


## GetDMail

> DMail GetDMail(ctx, id).Execute()

Get DMail



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
	id := int32(56) // int32 | The ID of the dmail.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DMailsAPI.GetDMail(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DMailsAPI.GetDMail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDMail`: DMail
	fmt.Fprintf(os.Stdout, "Response from `DMailsAPI.GetDMail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the dmail. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDMailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DMail**](DMail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkAllDMailsAsUnread

> MarkAllDMailsAsUnread(ctx).Execute()

Mark All DMails As Unread

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
	r, err := apiClient.DMailsAPI.MarkAllDMailsAsUnread(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DMailsAPI.MarkAllDMailsAsUnread``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMarkAllDMailsAsUnreadRequest struct via the builder pattern


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


## MarkDMailAsRead

> MarkDMailAsRead(ctx, id).Execute()

Mark DMail As Read

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
	id := int32(56) // int32 | The ID of the dmail.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DMailsAPI.MarkDMailAsRead(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DMailsAPI.MarkDMailAsRead``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the dmail. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkDMailAsReadRequest struct via the builder pattern


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


## MarkDMailAsUnread

> MarkDMailAsUnread(ctx, id).Execute()

Mark DMail As Unread

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
	id := int32(56) // int32 | The ID of the dmail.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DMailsAPI.MarkDMailAsUnread(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DMailsAPI.MarkDMailAsUnread``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The ID of the dmail. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkDMailAsUnreadRequest struct via the builder pattern


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


## SearchDMails

> SearchDMails200Response SearchDMails(ctx).Limit(limit).Page(page).SearchId(searchId).SearchTitleMatches(searchTitleMatches).SearchMessageMatches(searchMessageMatches).SearchToName(searchToName).SearchToId(searchToId).SearchFromName(searchFromName).SearchFromId(searchFromId).SearchIsRead(searchIsRead).SearchIsDeleted(searchIsDeleted).SearchRead(searchRead).Execute()

Search DMails



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
	searchTitleMatches := "searchTitleMatches_example" // string |  (optional)
	searchMessageMatches := "searchMessageMatches_example" // string |  (optional)
	searchToName := "searchToName_example" // string |  (optional)
	searchToId := int32(56) // int32 |  (optional)
	searchFromName := "searchFromName_example" // string |  (optional)
	searchFromId := int32(56) // int32 |  (optional)
	searchIsRead := true // bool |  (optional)
	searchIsDeleted := true // bool |  (optional)
	searchRead := true // bool |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DMailsAPI.SearchDMails(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchTitleMatches(searchTitleMatches).SearchMessageMatches(searchMessageMatches).SearchToName(searchToName).SearchToId(searchToId).SearchFromName(searchFromName).SearchFromId(searchFromId).SearchIsRead(searchIsRead).SearchIsDeleted(searchIsDeleted).SearchRead(searchRead).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DMailsAPI.SearchDMails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchDMails`: SearchDMails200Response
	fmt.Fprintf(os.Stdout, "Response from `DMailsAPI.SearchDMails`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchDMailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **int32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **int32** | Search for a specific id. | 
 **searchTitleMatches** | **string** |  | 
 **searchMessageMatches** | **string** |  | 
 **searchToName** | **string** |  | 
 **searchToId** | **int32** |  | 
 **searchFromName** | **string** |  | 
 **searchFromId** | **int32** |  | 
 **searchIsRead** | **bool** |  | 
 **searchIsDeleted** | **bool** |  | 
 **searchRead** | **bool** |  | 

### Return type

[**SearchDMails200Response**](SearchDMails200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserDmailFilter

> UpdateUserDmailFilter(ctx).DmailId(dmailId).DmailFilterWords(dmailFilterWords).Execute()

Update User Dmail Filter

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
	dmailId := int32(56) // int32 | Due to the odd way this route works, a dmail is REQUIRED to edit your dmail filter. You must be the owner of the dmail.
	dmailFilterWords := "dmailFilterWords_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DMailsAPI.UpdateUserDmailFilter(context.Background()).DmailId(dmailId).DmailFilterWords(dmailFilterWords).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DMailsAPI.UpdateUserDmailFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserDmailFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dmailId** | **int32** | Due to the odd way this route works, a dmail is REQUIRED to edit your dmail filter. You must be the owner of the dmail. | 
 **dmailFilterWords** | **string** |  | 

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

