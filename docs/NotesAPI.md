# \NotesAPI

All URIs are relative to *https://e621.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNote**](NotesAPI.md#CreateNote) | **Post** /notes.json | Create Note
[**DeleteNote**](NotesAPI.md#DeleteNote) | **Delete** /notes/{id}.json | Delete Note
[**EditNote**](NotesAPI.md#EditNote) | **Patch** /notes/{id}.json | Edit Note
[**GetNote**](NotesAPI.md#GetNote) | **Get** /notes/{id}.json | Get Note
[**RevertNote**](NotesAPI.md#RevertNote) | **Post** /notes/{id}/revert.json | Revert Note
[**SearchNotes**](NotesAPI.md#SearchNotes) | **Get** /notes.json | Search Notes



## CreateNote

> CreateNote201Response CreateNote(ctx).NotePostId(notePostId).NoteX(noteX).NoteY(noteY).NoteWidth(noteWidth).NoteHeight(noteHeight).NoteBody(noteBody).NoteHtmlId(noteHtmlId).Execute()

Create Note

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
	notePostId := float32(8.14) // float32 | 
	noteX := float32(8.14) // float32 | 
	noteY := float32(8.14) // float32 | 
	noteWidth := float32(8.14) // float32 | 
	noteHeight := float32(8.14) // float32 | 
	noteBody := "noteBody_example" // string | 
	noteHtmlId := "noteHtmlId_example" // string | Passthrough, used in frontend. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.CreateNote(context.Background()).NotePostId(notePostId).NoteX(noteX).NoteY(noteY).NoteWidth(noteWidth).NoteHeight(noteHeight).NoteBody(noteBody).NoteHtmlId(noteHtmlId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.CreateNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNote`: CreateNote201Response
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.CreateNote`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **notePostId** | **float32** |  | 
 **noteX** | **float32** |  | 
 **noteY** | **float32** |  | 
 **noteWidth** | **float32** |  | 
 **noteHeight** | **float32** |  | 
 **noteBody** | **string** |  | 
 **noteHtmlId** | **string** | Passthrough, used in frontend. | 

### Return type

[**CreateNote201Response**](CreateNote201Response.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNote

> DeleteNote(ctx, id).Execute()

Delete Note

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
	id := float32(8.14) // float32 | The ID of the note.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotesAPI.DeleteNote(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.DeleteNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the note. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNoteRequest struct via the builder pattern


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


## EditNote

> EditNote(ctx, id).NoteX(noteX).NoteY(noteY).NoteWidth(noteWidth).NoteHeight(noteHeight).NoteBody(noteBody).Execute()

Edit Note

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
	id := float32(8.14) // float32 | The ID of the note.
	noteX := float32(8.14) // float32 |  (optional)
	noteY := float32(8.14) // float32 |  (optional)
	noteWidth := float32(8.14) // float32 |  (optional)
	noteHeight := float32(8.14) // float32 |  (optional)
	noteBody := "noteBody_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotesAPI.EditNote(context.Background(), id).NoteX(noteX).NoteY(noteY).NoteWidth(noteWidth).NoteHeight(noteHeight).NoteBody(noteBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.EditNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the note. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **noteX** | **float32** |  | 
 **noteY** | **float32** |  | 
 **noteWidth** | **float32** |  | 
 **noteHeight** | **float32** |  | 
 **noteBody** | **string** |  | 

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


## GetNote

> Note GetNote(ctx, id).Execute()

Get Note

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
	id := float32(8.14) // float32 | The ID of the note.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.GetNote(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.GetNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNote`: Note
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.GetNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the note. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Note**](Note.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevertNote

> RevertNote(ctx, id).VersionId(versionId).Execute()

Revert Note

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
	id := float32(8.14) // float32 | The ID of the note.
	versionId := float32(8.14) // float32 | The version ID to revert to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NotesAPI.RevertNote(context.Background(), id).VersionId(versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.RevertNote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **float32** | The ID of the note. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevertNoteRequest struct via the builder pattern


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


## SearchNotes

> SearchNotes200Response SearchNotes(ctx).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchBodyMatches(searchBodyMatches).SearchIsActive(searchIsActive).SearchPostId(searchPostId).SearchPostTagsMatch(searchPostTagsMatch).SearchPostNoteUpdaterId(searchPostNoteUpdaterId).SearchPostNoteUpdaterName(searchPostNoteUpdaterName).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).Execute()

Search Notes



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
	searchBodyMatches := "searchBodyMatches_example" // string |  (optional)
	searchIsActive := true // bool |  (optional)
	searchPostId := float32(8.14) // float32 |  (optional)
	searchPostTagsMatch := "searchPostTagsMatch_example" // string |  (optional)
	searchPostNoteUpdaterId := float32(8.14) // float32 |  (optional)
	searchPostNoteUpdaterName := "searchPostNoteUpdaterName_example" // string |  (optional)
	searchCreatorId := float32(8.14) // float32 |  (optional)
	searchCreatorName := "searchCreatorName_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NotesAPI.SearchNotes(context.Background()).Limit(limit).Page(page).SearchId(searchId).SearchOrder(searchOrder).SearchBodyMatches(searchBodyMatches).SearchIsActive(searchIsActive).SearchPostId(searchPostId).SearchPostTagsMatch(searchPostTagsMatch).SearchPostNoteUpdaterId(searchPostNoteUpdaterId).SearchPostNoteUpdaterName(searchPostNoteUpdaterName).SearchCreatorId(searchCreatorId).SearchCreatorName(searchCreatorName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NotesAPI.SearchNotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchNotes`: SearchNotes200Response
	fmt.Fprintf(os.Stdout, "Response from `NotesAPI.SearchNotes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchNotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of results to return. Between 0 and 320. | 
 **page** | **int32** | The page number of results to get. Between 1 and 750. | 
 **searchId** | **int32** | Search for a specific id. | 
 **searchOrder** | **string** |  | 
 **searchBodyMatches** | **string** |  | 
 **searchIsActive** | **bool** |  | 
 **searchPostId** | **float32** |  | 
 **searchPostTagsMatch** | **string** |  | 
 **searchPostNoteUpdaterId** | **float32** |  | 
 **searchPostNoteUpdaterName** | **string** |  | 
 **searchCreatorId** | **float32** |  | 
 **searchCreatorName** | **string** |  | 

### Return type

[**SearchNotes200Response**](SearchNotes200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

