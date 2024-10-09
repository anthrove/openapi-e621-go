# Upload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Source** | **string** |  | 
**Rating** | [**Ratings**](Ratings.md) |  | 
**UploaderId** | **int32** |  | 
**TagString** | **string** |  | 
**Status** | **string** | Note: The \&quot;error\&quot; status will be proceeded by an error, ex: \&quot;error: RuntimeError - No file or source URL provided\&quot;  | 
**Backtrace** | **NullableString** |  | 
**PostId** | **NullableFloat32** |  | 
**Md5Confirmation** | **interface{}** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**ParentId** | **NullableFloat32** |  | 
**Md5** | **NullableString** |  | 
**FileExt** | [**GetArtistIdOrNameParameter**](GetArtistIdOrNameParameter.md) |  | 
**FileSize** | **NullableFloat32** |  | 
**ImageWidth** | **NullableFloat32** |  | 
**ImageHeight** | **NullableFloat32** |  | 
**Description** | **string** |  | 
**UploaderName** | **string** |  | 

## Methods

### NewUpload

`func NewUpload(id int32, source string, rating Ratings, uploaderId int32, tagString string, status string, backtrace NullableString, postId NullableFloat32, md5Confirmation interface{}, createdAt time.Time, updatedAt time.Time, parentId NullableFloat32, md5 NullableString, fileExt GetArtistIdOrNameParameter, fileSize NullableFloat32, imageWidth NullableFloat32, imageHeight NullableFloat32, description string, uploaderName string, ) *Upload`

NewUpload instantiates a new Upload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadWithDefaults

`func NewUploadWithDefaults() *Upload`

NewUploadWithDefaults instantiates a new Upload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Upload) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Upload) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Upload) SetId(v int32)`

SetId sets Id field to given value.


### GetSource

`func (o *Upload) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Upload) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Upload) SetSource(v string)`

SetSource sets Source field to given value.


### GetRating

`func (o *Upload) GetRating() Ratings`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *Upload) GetRatingOk() (*Ratings, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *Upload) SetRating(v Ratings)`

SetRating sets Rating field to given value.


### GetUploaderId

`func (o *Upload) GetUploaderId() int32`

GetUploaderId returns the UploaderId field if non-nil, zero value otherwise.

### GetUploaderIdOk

`func (o *Upload) GetUploaderIdOk() (*int32, bool)`

GetUploaderIdOk returns a tuple with the UploaderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaderId

`func (o *Upload) SetUploaderId(v int32)`

SetUploaderId sets UploaderId field to given value.


### GetTagString

`func (o *Upload) GetTagString() string`

GetTagString returns the TagString field if non-nil, zero value otherwise.

### GetTagStringOk

`func (o *Upload) GetTagStringOk() (*string, bool)`

GetTagStringOk returns a tuple with the TagString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagString

`func (o *Upload) SetTagString(v string)`

SetTagString sets TagString field to given value.


### GetStatus

`func (o *Upload) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Upload) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Upload) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetBacktrace

`func (o *Upload) GetBacktrace() string`

GetBacktrace returns the Backtrace field if non-nil, zero value otherwise.

### GetBacktraceOk

`func (o *Upload) GetBacktraceOk() (*string, bool)`

GetBacktraceOk returns a tuple with the Backtrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBacktrace

`func (o *Upload) SetBacktrace(v string)`

SetBacktrace sets Backtrace field to given value.


### SetBacktraceNil

`func (o *Upload) SetBacktraceNil(b bool)`

 SetBacktraceNil sets the value for Backtrace to be an explicit nil

### UnsetBacktrace
`func (o *Upload) UnsetBacktrace()`

UnsetBacktrace ensures that no value is present for Backtrace, not even an explicit nil
### GetPostId

`func (o *Upload) GetPostId() float32`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *Upload) GetPostIdOk() (*float32, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *Upload) SetPostId(v float32)`

SetPostId sets PostId field to given value.


### SetPostIdNil

`func (o *Upload) SetPostIdNil(b bool)`

 SetPostIdNil sets the value for PostId to be an explicit nil

### UnsetPostId
`func (o *Upload) UnsetPostId()`

UnsetPostId ensures that no value is present for PostId, not even an explicit nil
### GetMd5Confirmation

`func (o *Upload) GetMd5Confirmation() interface{}`

GetMd5Confirmation returns the Md5Confirmation field if non-nil, zero value otherwise.

### GetMd5ConfirmationOk

`func (o *Upload) GetMd5ConfirmationOk() (*interface{}, bool)`

GetMd5ConfirmationOk returns a tuple with the Md5Confirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5Confirmation

`func (o *Upload) SetMd5Confirmation(v interface{})`

SetMd5Confirmation sets Md5Confirmation field to given value.


### SetMd5ConfirmationNil

`func (o *Upload) SetMd5ConfirmationNil(b bool)`

 SetMd5ConfirmationNil sets the value for Md5Confirmation to be an explicit nil

### UnsetMd5Confirmation
`func (o *Upload) UnsetMd5Confirmation()`

UnsetMd5Confirmation ensures that no value is present for Md5Confirmation, not even an explicit nil
### GetCreatedAt

`func (o *Upload) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Upload) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Upload) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Upload) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Upload) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Upload) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetParentId

`func (o *Upload) GetParentId() float32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Upload) GetParentIdOk() (*float32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Upload) SetParentId(v float32)`

SetParentId sets ParentId field to given value.


### SetParentIdNil

`func (o *Upload) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *Upload) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetMd5

`func (o *Upload) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *Upload) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *Upload) SetMd5(v string)`

SetMd5 sets Md5 field to given value.


### SetMd5Nil

`func (o *Upload) SetMd5Nil(b bool)`

 SetMd5Nil sets the value for Md5 to be an explicit nil

### UnsetMd5
`func (o *Upload) UnsetMd5()`

UnsetMd5 ensures that no value is present for Md5, not even an explicit nil
### GetFileExt

`func (o *Upload) GetFileExt() GetArtistIdOrNameParameter`

GetFileExt returns the FileExt field if non-nil, zero value otherwise.

### GetFileExtOk

`func (o *Upload) GetFileExtOk() (*GetArtistIdOrNameParameter, bool)`

GetFileExtOk returns a tuple with the FileExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExt

`func (o *Upload) SetFileExt(v GetArtistIdOrNameParameter)`

SetFileExt sets FileExt field to given value.


### GetFileSize

`func (o *Upload) GetFileSize() float32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *Upload) GetFileSizeOk() (*float32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *Upload) SetFileSize(v float32)`

SetFileSize sets FileSize field to given value.


### SetFileSizeNil

`func (o *Upload) SetFileSizeNil(b bool)`

 SetFileSizeNil sets the value for FileSize to be an explicit nil

### UnsetFileSize
`func (o *Upload) UnsetFileSize()`

UnsetFileSize ensures that no value is present for FileSize, not even an explicit nil
### GetImageWidth

`func (o *Upload) GetImageWidth() float32`

GetImageWidth returns the ImageWidth field if non-nil, zero value otherwise.

### GetImageWidthOk

`func (o *Upload) GetImageWidthOk() (*float32, bool)`

GetImageWidthOk returns a tuple with the ImageWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageWidth

`func (o *Upload) SetImageWidth(v float32)`

SetImageWidth sets ImageWidth field to given value.


### SetImageWidthNil

`func (o *Upload) SetImageWidthNil(b bool)`

 SetImageWidthNil sets the value for ImageWidth to be an explicit nil

### UnsetImageWidth
`func (o *Upload) UnsetImageWidth()`

UnsetImageWidth ensures that no value is present for ImageWidth, not even an explicit nil
### GetImageHeight

`func (o *Upload) GetImageHeight() float32`

GetImageHeight returns the ImageHeight field if non-nil, zero value otherwise.

### GetImageHeightOk

`func (o *Upload) GetImageHeightOk() (*float32, bool)`

GetImageHeightOk returns a tuple with the ImageHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageHeight

`func (o *Upload) SetImageHeight(v float32)`

SetImageHeight sets ImageHeight field to given value.


### SetImageHeightNil

`func (o *Upload) SetImageHeightNil(b bool)`

 SetImageHeightNil sets the value for ImageHeight to be an explicit nil

### UnsetImageHeight
`func (o *Upload) UnsetImageHeight()`

UnsetImageHeight ensures that no value is present for ImageHeight, not even an explicit nil
### GetDescription

`func (o *Upload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Upload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Upload) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetUploaderName

`func (o *Upload) GetUploaderName() string`

GetUploaderName returns the UploaderName field if non-nil, zero value otherwise.

### GetUploaderNameOk

`func (o *Upload) GetUploaderNameOk() (*string, bool)`

GetUploaderNameOk returns a tuple with the UploaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaderName

`func (o *Upload) SetUploaderName(v string)`

SetUploaderName sets UploaderName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


