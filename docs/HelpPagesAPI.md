# \HelpPagesAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHelpPage**](HelpPagesAPI.md#CreateHelpPage) | **Post** /help.json | Create Help Page
[**DeleteHelpPage**](HelpPagesAPI.md#DeleteHelpPage) | **Delete** /help/{id}.json | Delete Help Page
[**EditHelpPage**](HelpPagesAPI.md#EditHelpPage) | **Patch** /help/{id}.json | Edit Help Page
[**GetHelpPage**](HelpPagesAPI.md#GetHelpPage) | **Get** /help/{id}.json | Get Help Page
[**ListHelpPages**](HelpPagesAPI.md#ListHelpPages) | **Get** /help.json | List Help Pages



## CreateHelpPage

> Help CreateHelpPage(ctx).HelpPageName(helpPageName).HelpPageWikiPage(helpPageWikiPage).HelpPageRelated(helpPageRelated).HelpPageTitle(helpPageTitle).Execute()

Create Help Page



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
	helpPageName := "helpPageName_example" // string | 
	helpPageWikiPage := "helpPageWikiPage_example" // string | 
	helpPageRelated := "helpPageRelated_example" // string | Separate with a comma followed by a space. (optional)
	helpPageTitle := "helpPageTitle_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelpPagesAPI.CreateHelpPage(context.Background()).HelpPageName(helpPageName).HelpPageWikiPage(helpPageWikiPage).HelpPageRelated(helpPageRelated).HelpPageTitle(helpPageTitle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelpPagesAPI.CreateHelpPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHelpPage`: Help
	fmt.Fprintf(os.Stdout, "Response from `HelpPagesAPI.CreateHelpPage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHelpPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **helpPageName** | **string** |  | 
 **helpPageWikiPage** | **string** |  | 
 **helpPageRelated** | **string** | Separate with a comma followed by a space. | 
 **helpPageTitle** | **string** |  | 

### Return type

[**Help**](Help.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHelpPage

> DeleteHelpPage(ctx, id).Execute()

Delete Help Page



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
	id := float32(8.14) // float32 | The ID of the help page.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HelpPagesAPI.DeleteHelpPage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelpPagesAPI.DeleteHelpPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the help page. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHelpPageRequest struct via the builder pattern


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


## EditHelpPage

> EditHelpPage(ctx, id).HelpPageName(helpPageName).HelpPageWikiPage(helpPageWikiPage).HelpPageRelated(helpPageRelated).HelpPageTitle(helpPageTitle).Execute()

Edit Help Page



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
	id := float32(8.14) // float32 | The ID of the help page.
	helpPageName := "helpPageName_example" // string |  (optional)
	helpPageWikiPage := "helpPageWikiPage_example" // string |  (optional)
	helpPageRelated := "helpPageRelated_example" // string | Separate with a comma followed by a space. (optional)
	helpPageTitle := "helpPageTitle_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HelpPagesAPI.EditHelpPage(context.Background(), id).HelpPageName(helpPageName).HelpPageWikiPage(helpPageWikiPage).HelpPageRelated(helpPageRelated).HelpPageTitle(helpPageTitle).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelpPagesAPI.EditHelpPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the help page. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditHelpPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **helpPageName** | **string** |  | 
 **helpPageWikiPage** | **string** |  | 
 **helpPageRelated** | **string** | Separate with a comma followed by a space. | 
 **helpPageTitle** | **string** |  | 

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


## GetHelpPage

> Help GetHelpPage(ctx, id).Execute()

Get Help Page

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
	id := GetArtistIdOrNameParameter(8.14) // GetArtistIdOrNameParameter | The ID or name of the help page.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HelpPagesAPI.GetHelpPage(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelpPagesAPI.GetHelpPage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHelpPage`: Help
	fmt.Fprintf(os.Stdout, "Response from `HelpPagesAPI.GetHelpPage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **GetArtistIdOrNameParameter** | The ID or name of the help page. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHelpPageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Help**](Help.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHelpPages

> []Help ListHelpPages(ctx).Execute()

List Help Pages



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
	resp, r, err := apiClient.HelpPagesAPI.ListHelpPages(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HelpPagesAPI.ListHelpPages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHelpPages`: []Help
	fmt.Fprintf(os.Stdout, "Response from `HelpPagesAPI.ListHelpPages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListHelpPagesRequest struct via the builder pattern


### Return type

[**[]Help**](Help.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

