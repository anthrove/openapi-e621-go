# PostReplacement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**PostId** | **float32** |  | 
**CreatorId** | **float32** |  | 
**ApproverId** | **NullableFloat32** |  | 
**FileExt** | **string** |  | 
**FileSize** | **float32** |  | 
**ImageHeight** | **float32** |  | 
**ImageWidth** | **float32** |  | 
**Md5** | **string** |  | 
**Source** | **string** |  | 
**FileName** | **string** |  | 
**Status** | **string** |  | 
**Reason** | **string** |  | 

## Methods

### NewPostReplacement

`func NewPostReplacement(id float32, createdAt time.Time, updatedAt time.Time, postId float32, creatorId float32, approverId NullableFloat32, fileExt string, fileSize float32, imageHeight float32, imageWidth float32, md5 string, source string, fileName string, status string, reason string, ) *PostReplacement`

NewPostReplacement instantiates a new PostReplacement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostReplacementWithDefaults

`func NewPostReplacementWithDefaults() *PostReplacement`

NewPostReplacementWithDefaults instantiates a new PostReplacement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostReplacement) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostReplacement) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostReplacement) SetId(v float32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *PostReplacement) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PostReplacement) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PostReplacement) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PostReplacement) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PostReplacement) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PostReplacement) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetPostId

`func (o *PostReplacement) GetPostId() float32`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *PostReplacement) GetPostIdOk() (*float32, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *PostReplacement) SetPostId(v float32)`

SetPostId sets PostId field to given value.


### GetCreatorId

`func (o *PostReplacement) GetCreatorId() float32`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *PostReplacement) GetCreatorIdOk() (*float32, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *PostReplacement) SetCreatorId(v float32)`

SetCreatorId sets CreatorId field to given value.


### GetApproverId

`func (o *PostReplacement) GetApproverId() float32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *PostReplacement) GetApproverIdOk() (*float32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *PostReplacement) SetApproverId(v float32)`

SetApproverId sets ApproverId field to given value.


### SetApproverIdNil

`func (o *PostReplacement) SetApproverIdNil(b bool)`

 SetApproverIdNil sets the value for ApproverId to be an explicit nil

### UnsetApproverId
`func (o *PostReplacement) UnsetApproverId()`

UnsetApproverId ensures that no value is present for ApproverId, not even an explicit nil
### GetFileExt

`func (o *PostReplacement) GetFileExt() string`

GetFileExt returns the FileExt field if non-nil, zero value otherwise.

### GetFileExtOk

`func (o *PostReplacement) GetFileExtOk() (*string, bool)`

GetFileExtOk returns a tuple with the FileExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExt

`func (o *PostReplacement) SetFileExt(v string)`

SetFileExt sets FileExt field to given value.


### GetFileSize

`func (o *PostReplacement) GetFileSize() float32`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *PostReplacement) GetFileSizeOk() (*float32, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *PostReplacement) SetFileSize(v float32)`

SetFileSize sets FileSize field to given value.


### GetImageHeight

`func (o *PostReplacement) GetImageHeight() float32`

GetImageHeight returns the ImageHeight field if non-nil, zero value otherwise.

### GetImageHeightOk

`func (o *PostReplacement) GetImageHeightOk() (*float32, bool)`

GetImageHeightOk returns a tuple with the ImageHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageHeight

`func (o *PostReplacement) SetImageHeight(v float32)`

SetImageHeight sets ImageHeight field to given value.


### GetImageWidth

`func (o *PostReplacement) GetImageWidth() float32`

GetImageWidth returns the ImageWidth field if non-nil, zero value otherwise.

### GetImageWidthOk

`func (o *PostReplacement) GetImageWidthOk() (*float32, bool)`

GetImageWidthOk returns a tuple with the ImageWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageWidth

`func (o *PostReplacement) SetImageWidth(v float32)`

SetImageWidth sets ImageWidth field to given value.


### GetMd5

`func (o *PostReplacement) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *PostReplacement) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *PostReplacement) SetMd5(v string)`

SetMd5 sets Md5 field to given value.


### GetSource

`func (o *PostReplacement) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PostReplacement) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PostReplacement) SetSource(v string)`

SetSource sets Source field to given value.


### GetFileName

`func (o *PostReplacement) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *PostReplacement) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *PostReplacement) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetStatus

`func (o *PostReplacement) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostReplacement) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostReplacement) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReason

`func (o *PostReplacement) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PostReplacement) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PostReplacement) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

