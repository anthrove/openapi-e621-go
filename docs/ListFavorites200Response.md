# ListFavorites200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Posts** | [**[]Post**](Post.md) |  | 

## Methods

### NewListFavorites200Response

`func NewListFavorites200Response(posts []Post, ) *ListFavorites200Response`

NewListFavorites200Response instantiates a new ListFavorites200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFavorites200ResponseWithDefaults

`func NewListFavorites200ResponseWithDefaults() *ListFavorites200Response`

NewListFavorites200ResponseWithDefaults instantiates a new ListFavorites200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosts

`func (o *ListFavorites200Response) GetPosts() []Post`

GetPosts returns the Posts field if non-nil, zero value otherwise.

### GetPostsOk

`func (o *ListFavorites200Response) GetPostsOk() (*[]Post, bool)`

GetPostsOk returns a tuple with the Posts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosts

`func (o *ListFavorites200Response) SetPosts(v []Post)`

SetPosts sets Posts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


