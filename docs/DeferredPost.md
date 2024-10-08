# DeferredPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**Flags** | **string** |  | 
**Tags** | **string** |  | 
**Rating** | [**Ratings**](Ratings.md) |  | 
**FileExt** | **string** |  | 
**Width** | **float32** |  | 
**Height** | **float32** |  | 
**Size** | **float32** |  | 
**CreatedAt** | **time.Time** |  | 
**Uploader** | **string** |  | 
**UploaderId** | **float32** |  | 
**Score** | **float32** |  | 
**FavCount** | **float32** |  | 
**IsFavorited** | **bool** |  | 
**Pools** | **[]float32** |  | 
**Md5** | **string** |  | 
**PreviewUrl** | **NullableString** |  | 
**LargeUrl** | **NullableString** |  | 
**FileUrl** | **NullableString** |  | 
**PreviewWidth** | **float32** |  | 
**PreviewHeight** | **float32** |  | 

## Methods

### NewDeferredPost

`func NewDeferredPost(id float32, flags string, tags string, rating Ratings, fileExt string, width float32, height float32, size float32, createdAt time.Time, uploader string, uploaderId float32, score float32, favCount float32, isFavorited bool, pools []float32, md5 string, previewUrl NullableString, largeUrl NullableString, fileUrl NullableString, previewWidth float32, previewHeight float32, ) *DeferredPost`

NewDeferredPost instantiates a new DeferredPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeferredPostWithDefaults

`func NewDeferredPostWithDefaults() *DeferredPost`

NewDeferredPostWithDefaults instantiates a new DeferredPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeferredPost) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeferredPost) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeferredPost) SetId(v float32)`

SetId sets Id field to given value.


### GetFlags

`func (o *DeferredPost) GetFlags() string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *DeferredPost) GetFlagsOk() (*string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *DeferredPost) SetFlags(v string)`

SetFlags sets Flags field to given value.


### GetTags

`func (o *DeferredPost) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeferredPost) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeferredPost) SetTags(v string)`

SetTags sets Tags field to given value.


### GetRating

`func (o *DeferredPost) GetRating() Ratings`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *DeferredPost) GetRatingOk() (*Ratings, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *DeferredPost) SetRating(v Ratings)`

SetRating sets Rating field to given value.


### GetFileExt

`func (o *DeferredPost) GetFileExt() string`

GetFileExt returns the FileExt field if non-nil, zero value otherwise.

### GetFileExtOk

`func (o *DeferredPost) GetFileExtOk() (*string, bool)`

GetFileExtOk returns a tuple with the FileExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExt

`func (o *DeferredPost) SetFileExt(v string)`

SetFileExt sets FileExt field to given value.


### GetWidth

`func (o *DeferredPost) GetWidth() float32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *DeferredPost) GetWidthOk() (*float32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *DeferredPost) SetWidth(v float32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *DeferredPost) GetHeight() float32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *DeferredPost) GetHeightOk() (*float32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *DeferredPost) SetHeight(v float32)`

SetHeight sets Height field to given value.


### GetSize

`func (o *DeferredPost) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DeferredPost) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DeferredPost) SetSize(v float32)`

SetSize sets Size field to given value.


### GetCreatedAt

`func (o *DeferredPost) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DeferredPost) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DeferredPost) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUploader

`func (o *DeferredPost) GetUploader() string`

GetUploader returns the Uploader field if non-nil, zero value otherwise.

### GetUploaderOk

`func (o *DeferredPost) GetUploaderOk() (*string, bool)`

GetUploaderOk returns a tuple with the Uploader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploader

`func (o *DeferredPost) SetUploader(v string)`

SetUploader sets Uploader field to given value.


### GetUploaderId

`func (o *DeferredPost) GetUploaderId() float32`

GetUploaderId returns the UploaderId field if non-nil, zero value otherwise.

### GetUploaderIdOk

`func (o *DeferredPost) GetUploaderIdOk() (*float32, bool)`

GetUploaderIdOk returns a tuple with the UploaderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaderId

`func (o *DeferredPost) SetUploaderId(v float32)`

SetUploaderId sets UploaderId field to given value.


### GetScore

`func (o *DeferredPost) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *DeferredPost) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *DeferredPost) SetScore(v float32)`

SetScore sets Score field to given value.


### GetFavCount

`func (o *DeferredPost) GetFavCount() float32`

GetFavCount returns the FavCount field if non-nil, zero value otherwise.

### GetFavCountOk

`func (o *DeferredPost) GetFavCountOk() (*float32, bool)`

GetFavCountOk returns a tuple with the FavCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavCount

`func (o *DeferredPost) SetFavCount(v float32)`

SetFavCount sets FavCount field to given value.


### GetIsFavorited

`func (o *DeferredPost) GetIsFavorited() bool`

GetIsFavorited returns the IsFavorited field if non-nil, zero value otherwise.

### GetIsFavoritedOk

`func (o *DeferredPost) GetIsFavoritedOk() (*bool, bool)`

GetIsFavoritedOk returns a tuple with the IsFavorited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorited

`func (o *DeferredPost) SetIsFavorited(v bool)`

SetIsFavorited sets IsFavorited field to given value.


### GetPools

`func (o *DeferredPost) GetPools() []float32`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *DeferredPost) GetPoolsOk() (*[]float32, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *DeferredPost) SetPools(v []float32)`

SetPools sets Pools field to given value.


### GetMd5

`func (o *DeferredPost) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *DeferredPost) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *DeferredPost) SetMd5(v string)`

SetMd5 sets Md5 field to given value.


### GetPreviewUrl

`func (o *DeferredPost) GetPreviewUrl() string`

GetPreviewUrl returns the PreviewUrl field if non-nil, zero value otherwise.

### GetPreviewUrlOk

`func (o *DeferredPost) GetPreviewUrlOk() (*string, bool)`

GetPreviewUrlOk returns a tuple with the PreviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewUrl

`func (o *DeferredPost) SetPreviewUrl(v string)`

SetPreviewUrl sets PreviewUrl field to given value.


### SetPreviewUrlNil

`func (o *DeferredPost) SetPreviewUrlNil(b bool)`

 SetPreviewUrlNil sets the value for PreviewUrl to be an explicit nil

### UnsetPreviewUrl
`func (o *DeferredPost) UnsetPreviewUrl()`

UnsetPreviewUrl ensures that no value is present for PreviewUrl, not even an explicit nil
### GetLargeUrl

`func (o *DeferredPost) GetLargeUrl() string`

GetLargeUrl returns the LargeUrl field if non-nil, zero value otherwise.

### GetLargeUrlOk

`func (o *DeferredPost) GetLargeUrlOk() (*string, bool)`

GetLargeUrlOk returns a tuple with the LargeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeUrl

`func (o *DeferredPost) SetLargeUrl(v string)`

SetLargeUrl sets LargeUrl field to given value.


### SetLargeUrlNil

`func (o *DeferredPost) SetLargeUrlNil(b bool)`

 SetLargeUrlNil sets the value for LargeUrl to be an explicit nil

### UnsetLargeUrl
`func (o *DeferredPost) UnsetLargeUrl()`

UnsetLargeUrl ensures that no value is present for LargeUrl, not even an explicit nil
### GetFileUrl

`func (o *DeferredPost) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *DeferredPost) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *DeferredPost) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.


### SetFileUrlNil

`func (o *DeferredPost) SetFileUrlNil(b bool)`

 SetFileUrlNil sets the value for FileUrl to be an explicit nil

### UnsetFileUrl
`func (o *DeferredPost) UnsetFileUrl()`

UnsetFileUrl ensures that no value is present for FileUrl, not even an explicit nil
### GetPreviewWidth

`func (o *DeferredPost) GetPreviewWidth() float32`

GetPreviewWidth returns the PreviewWidth field if non-nil, zero value otherwise.

### GetPreviewWidthOk

`func (o *DeferredPost) GetPreviewWidthOk() (*float32, bool)`

GetPreviewWidthOk returns a tuple with the PreviewWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewWidth

`func (o *DeferredPost) SetPreviewWidth(v float32)`

SetPreviewWidth sets PreviewWidth field to given value.


### GetPreviewHeight

`func (o *DeferredPost) GetPreviewHeight() float32`

GetPreviewHeight returns the PreviewHeight field if non-nil, zero value otherwise.

### GetPreviewHeightOk

`func (o *DeferredPost) GetPreviewHeightOk() (*float32, bool)`

GetPreviewHeightOk returns a tuple with the PreviewHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewHeight

`func (o *DeferredPost) SetPreviewHeight(v float32)`

SetPreviewHeight sets PreviewHeight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


