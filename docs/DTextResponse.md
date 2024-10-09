# DTextResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Html** | **string** |  | 
**Posts** | **map[string]interface{}** |  | 

## Methods

### NewDTextResponse

`func NewDTextResponse(html string, posts map[string]interface{}, ) *DTextResponse`

NewDTextResponse instantiates a new DTextResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDTextResponseWithDefaults

`func NewDTextResponseWithDefaults() *DTextResponse`

NewDTextResponseWithDefaults instantiates a new DTextResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHtml

`func (o *DTextResponse) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *DTextResponse) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *DTextResponse) SetHtml(v string)`

SetHtml sets Html field to given value.


### GetPosts

`func (o *DTextResponse) GetPosts() map[string]interface{}`

GetPosts returns the Posts field if non-nil, zero value otherwise.

### GetPostsOk

`func (o *DTextResponse) GetPostsOk() (*map[string]interface{}, bool)`

GetPostsOk returns a tuple with the Posts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosts

`func (o *DTextResponse) SetPosts(v map[string]interface{})`

SetPosts sets Posts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


