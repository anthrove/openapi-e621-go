# IQDBPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**UpScore** | **int32** |  | 
**DownScore** | **int32** |  | 
**Score** | **int32** |  | 
**Source** | **string** |  | 
**Md5** | **string** |  | 
**Rating** | [**Ratings**](Ratings.md) |  | 
**IsNoteLocked** | **bool** |  | 
**IsRatingLocked** | **bool** |  | 
**IsStatusLocked** | **bool** |  | 
**IsPending** | **bool** |  | 
**IsFlagged** | **bool** |  | 
**IsDeleted** | **bool** |  | 
**UploaderId** | **int32** |  | 
**ApproverId** | **int32** |  | 
**LastNotedAt** | **NullableTime** |  | 
**LastCommentBumpedAt** | **NullableTime** |  | 
**FavCount** | **int32** |  | 
**TagString** | **string** |  | 
**TagCount** | **int32** |  | 
**TagCountGeneral** | **int32** |  | 
**TagCountArtist** | **int32** |  | 
**TagCountCharacter** | **int32** |  | 
**TagCountCopyright** | **int32** |  | 
**FileExt** | **string** |  | 
**FileSize** | **int64** |  | 
**ImageWidth** | **int32** |  | 
**ImageHeight** | **int32** |  | 
**ParentId** | **NullableInt32** |  | 
**HasChildren** | **bool** |  | 
**LastCommentedAt** | **NullableTime** |  | 
**HasActiveChildren** | **bool** |  | 
**BitFlags** | **float32** |  | 
**TagCountMeta** | **int32** |  | 
**LockedTags** | **NullableString** |  | 
**TagCountSpecies** | **int32** |  | 
**TagCountInvalid** | **int32** |  | 
**Description** | **string** |  | 
**CommentCount** | **int32** |  | 
**ChangeSeq** | **int32** |  | 
**TagCountLore** | **int32** |  | 
**BgColor** | **NullableString** |  | 
**GeneratedSamples** | **[]string** |  | 
**Duration** | **NullableString** |  | 
**IsCommentDisabled** | **bool** |  | 
**IsCommentLocked** | **bool** |  | 
**HasLarge** | **bool** |  | 
**HasVisibleChildren** | **bool** |  | 
**ChildrenIds** | **NullableString** |  | 
**PoolIds** | **[]int32** |  | 
**IsFavorited** | **bool** |  | 
**FileUrl** | Pointer to **string** |  | [optional] 
**LargeFileUrl** | Pointer to **string** |  | [optional] 
**PreviewFileUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewIQDBPost

`func NewIQDBPost(id int32, createdAt time.Time, updatedAt time.Time, upScore int32, downScore int32, score int32, source string, md5 string, rating Ratings, isNoteLocked bool, isRatingLocked bool, isStatusLocked bool, isPending bool, isFlagged bool, isDeleted bool, uploaderId int32, approverId int32, lastNotedAt NullableTime, lastCommentBumpedAt NullableTime, favCount int32, tagString string, tagCount int32, tagCountGeneral int32, tagCountArtist int32, tagCountCharacter int32, tagCountCopyright int32, fileExt string, fileSize int64, imageWidth int32, imageHeight int32, parentId NullableInt32, hasChildren bool, lastCommentedAt NullableTime, hasActiveChildren bool, bitFlags float32, tagCountMeta int32, lockedTags NullableString, tagCountSpecies int32, tagCountInvalid int32, description string, commentCount int32, changeSeq int32, tagCountLore int32, bgColor NullableString, generatedSamples []string, duration NullableString, isCommentDisabled bool, isCommentLocked bool, hasLarge bool, hasVisibleChildren bool, childrenIds NullableString, poolIds []int32, isFavorited bool, ) *IQDBPost`

NewIQDBPost instantiates a new IQDBPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIQDBPostWithDefaults

`func NewIQDBPostWithDefaults() *IQDBPost`

NewIQDBPostWithDefaults instantiates a new IQDBPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IQDBPost) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IQDBPost) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IQDBPost) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *IQDBPost) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IQDBPost) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IQDBPost) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *IQDBPost) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IQDBPost) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IQDBPost) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUpScore

`func (o *IQDBPost) GetUpScore() int32`

GetUpScore returns the UpScore field if non-nil, zero value otherwise.

### GetUpScoreOk

`func (o *IQDBPost) GetUpScoreOk() (*int32, bool)`

GetUpScoreOk returns a tuple with the UpScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpScore

`func (o *IQDBPost) SetUpScore(v int32)`

SetUpScore sets UpScore field to given value.


### GetDownScore

`func (o *IQDBPost) GetDownScore() int32`

GetDownScore returns the DownScore field if non-nil, zero value otherwise.

### GetDownScoreOk

`func (o *IQDBPost) GetDownScoreOk() (*int32, bool)`

GetDownScoreOk returns a tuple with the DownScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownScore

`func (o *IQDBPost) SetDownScore(v int32)`

SetDownScore sets DownScore field to given value.


### GetScore

`func (o *IQDBPost) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *IQDBPost) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *IQDBPost) SetScore(v int32)`

SetScore sets Score field to given value.


### GetSource

`func (o *IQDBPost) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *IQDBPost) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *IQDBPost) SetSource(v string)`

SetSource sets Source field to given value.


### GetMd5

`func (o *IQDBPost) GetMd5() string`

GetMd5 returns the Md5 field if non-nil, zero value otherwise.

### GetMd5Ok

`func (o *IQDBPost) GetMd5Ok() (*string, bool)`

GetMd5Ok returns a tuple with the Md5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5

`func (o *IQDBPost) SetMd5(v string)`

SetMd5 sets Md5 field to given value.


### GetRating

`func (o *IQDBPost) GetRating() Ratings`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *IQDBPost) GetRatingOk() (*Ratings, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *IQDBPost) SetRating(v Ratings)`

SetRating sets Rating field to given value.


### GetIsNoteLocked

`func (o *IQDBPost) GetIsNoteLocked() bool`

GetIsNoteLocked returns the IsNoteLocked field if non-nil, zero value otherwise.

### GetIsNoteLockedOk

`func (o *IQDBPost) GetIsNoteLockedOk() (*bool, bool)`

GetIsNoteLockedOk returns a tuple with the IsNoteLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNoteLocked

`func (o *IQDBPost) SetIsNoteLocked(v bool)`

SetIsNoteLocked sets IsNoteLocked field to given value.


### GetIsRatingLocked

`func (o *IQDBPost) GetIsRatingLocked() bool`

GetIsRatingLocked returns the IsRatingLocked field if non-nil, zero value otherwise.

### GetIsRatingLockedOk

`func (o *IQDBPost) GetIsRatingLockedOk() (*bool, bool)`

GetIsRatingLockedOk returns a tuple with the IsRatingLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRatingLocked

`func (o *IQDBPost) SetIsRatingLocked(v bool)`

SetIsRatingLocked sets IsRatingLocked field to given value.


### GetIsStatusLocked

`func (o *IQDBPost) GetIsStatusLocked() bool`

GetIsStatusLocked returns the IsStatusLocked field if non-nil, zero value otherwise.

### GetIsStatusLockedOk

`func (o *IQDBPost) GetIsStatusLockedOk() (*bool, bool)`

GetIsStatusLockedOk returns a tuple with the IsStatusLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStatusLocked

`func (o *IQDBPost) SetIsStatusLocked(v bool)`

SetIsStatusLocked sets IsStatusLocked field to given value.


### GetIsPending

`func (o *IQDBPost) GetIsPending() bool`

GetIsPending returns the IsPending field if non-nil, zero value otherwise.

### GetIsPendingOk

`func (o *IQDBPost) GetIsPendingOk() (*bool, bool)`

GetIsPendingOk returns a tuple with the IsPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPending

`func (o *IQDBPost) SetIsPending(v bool)`

SetIsPending sets IsPending field to given value.


### GetIsFlagged

`func (o *IQDBPost) GetIsFlagged() bool`

GetIsFlagged returns the IsFlagged field if non-nil, zero value otherwise.

### GetIsFlaggedOk

`func (o *IQDBPost) GetIsFlaggedOk() (*bool, bool)`

GetIsFlaggedOk returns a tuple with the IsFlagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFlagged

`func (o *IQDBPost) SetIsFlagged(v bool)`

SetIsFlagged sets IsFlagged field to given value.


### GetIsDeleted

`func (o *IQDBPost) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *IQDBPost) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *IQDBPost) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.


### GetUploaderId

`func (o *IQDBPost) GetUploaderId() int32`

GetUploaderId returns the UploaderId field if non-nil, zero value otherwise.

### GetUploaderIdOk

`func (o *IQDBPost) GetUploaderIdOk() (*int32, bool)`

GetUploaderIdOk returns a tuple with the UploaderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaderId

`func (o *IQDBPost) SetUploaderId(v int32)`

SetUploaderId sets UploaderId field to given value.


### GetApproverId

`func (o *IQDBPost) GetApproverId() int32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *IQDBPost) GetApproverIdOk() (*int32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *IQDBPost) SetApproverId(v int32)`

SetApproverId sets ApproverId field to given value.


### GetLastNotedAt

`func (o *IQDBPost) GetLastNotedAt() time.Time`

GetLastNotedAt returns the LastNotedAt field if non-nil, zero value otherwise.

### GetLastNotedAtOk

`func (o *IQDBPost) GetLastNotedAtOk() (*time.Time, bool)`

GetLastNotedAtOk returns a tuple with the LastNotedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNotedAt

`func (o *IQDBPost) SetLastNotedAt(v time.Time)`

SetLastNotedAt sets LastNotedAt field to given value.


### SetLastNotedAtNil

`func (o *IQDBPost) SetLastNotedAtNil(b bool)`

 SetLastNotedAtNil sets the value for LastNotedAt to be an explicit nil

### UnsetLastNotedAt
`func (o *IQDBPost) UnsetLastNotedAt()`

UnsetLastNotedAt ensures that no value is present for LastNotedAt, not even an explicit nil
### GetLastCommentBumpedAt

`func (o *IQDBPost) GetLastCommentBumpedAt() time.Time`

GetLastCommentBumpedAt returns the LastCommentBumpedAt field if non-nil, zero value otherwise.

### GetLastCommentBumpedAtOk

`func (o *IQDBPost) GetLastCommentBumpedAtOk() (*time.Time, bool)`

GetLastCommentBumpedAtOk returns a tuple with the LastCommentBumpedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommentBumpedAt

`func (o *IQDBPost) SetLastCommentBumpedAt(v time.Time)`

SetLastCommentBumpedAt sets LastCommentBumpedAt field to given value.


### SetLastCommentBumpedAtNil

`func (o *IQDBPost) SetLastCommentBumpedAtNil(b bool)`

 SetLastCommentBumpedAtNil sets the value for LastCommentBumpedAt to be an explicit nil

### UnsetLastCommentBumpedAt
`func (o *IQDBPost) UnsetLastCommentBumpedAt()`

UnsetLastCommentBumpedAt ensures that no value is present for LastCommentBumpedAt, not even an explicit nil
### GetFavCount

`func (o *IQDBPost) GetFavCount() int32`

GetFavCount returns the FavCount field if non-nil, zero value otherwise.

### GetFavCountOk

`func (o *IQDBPost) GetFavCountOk() (*int32, bool)`

GetFavCountOk returns a tuple with the FavCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavCount

`func (o *IQDBPost) SetFavCount(v int32)`

SetFavCount sets FavCount field to given value.


### GetTagString

`func (o *IQDBPost) GetTagString() string`

GetTagString returns the TagString field if non-nil, zero value otherwise.

### GetTagStringOk

`func (o *IQDBPost) GetTagStringOk() (*string, bool)`

GetTagStringOk returns a tuple with the TagString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagString

`func (o *IQDBPost) SetTagString(v string)`

SetTagString sets TagString field to given value.


### GetTagCount

`func (o *IQDBPost) GetTagCount() int32`

GetTagCount returns the TagCount field if non-nil, zero value otherwise.

### GetTagCountOk

`func (o *IQDBPost) GetTagCountOk() (*int32, bool)`

GetTagCountOk returns a tuple with the TagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCount

`func (o *IQDBPost) SetTagCount(v int32)`

SetTagCount sets TagCount field to given value.


### GetTagCountGeneral

`func (o *IQDBPost) GetTagCountGeneral() int32`

GetTagCountGeneral returns the TagCountGeneral field if non-nil, zero value otherwise.

### GetTagCountGeneralOk

`func (o *IQDBPost) GetTagCountGeneralOk() (*int32, bool)`

GetTagCountGeneralOk returns a tuple with the TagCountGeneral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCountGeneral

`func (o *IQDBPost) SetTagCountGeneral(v int32)`

SetTagCountGeneral sets TagCountGeneral field to given value.


### GetTagCountArtist

`func (o *IQDBPost) GetTagCountArtist() int32`

GetTagCountArtist returns the TagCountArtist field if non-nil, zero value otherwise.

### GetTagCountArtistOk

`func (o *IQDBPost) GetTagCountArtistOk() (*int32, bool)`

GetTagCountArtistOk returns a tuple with the TagCountArtist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCountArtist

`func (o *IQDBPost) SetTagCountArtist(v int32)`

SetTagCountArtist sets TagCountArtist field to given value.


### GetTagCountCharacter

`func (o *IQDBPost) GetTagCountCharacter() int32`

GetTagCountCharacter returns the TagCountCharacter field if non-nil, zero value otherwise.

### GetTagCountCharacterOk

`func (o *IQDBPost) GetTagCountCharacterOk() (*int32, bool)`

GetTagCountCharacterOk returns a tuple with the TagCountCharacter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCountCharacter

`func (o *IQDBPost) SetTagCountCharacter(v int32)`

SetTagCountCharacter sets TagCountCharacter field to given value.


### GetTagCountCopyright

`func (o *IQDBPost) GetTagCountCopyright() int32`

GetTagCountCopyright returns the TagCountCopyright field if non-nil, zero value otherwise.

### GetTagCountCopyrightOk

`func (o *IQDBPost) GetTagCountCopyrightOk() (*int32, bool)`

GetTagCountCopyrightOk returns a tuple with the TagCountCopyright field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCountCopyright

`func (o *IQDBPost) SetTagCountCopyright(v int32)`

SetTagCountCopyright sets TagCountCopyright field to given value.


### GetFileExt

`func (o *IQDBPost) GetFileExt() string`

GetFileExt returns the FileExt field if non-nil, zero value otherwise.

### GetFileExtOk

`func (o *IQDBPost) GetFileExtOk() (*string, bool)`

GetFileExtOk returns a tuple with the FileExt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExt

`func (o *IQDBPost) SetFileExt(v string)`

SetFileExt sets FileExt field to given value.


### GetFileSize

`func (o *IQDBPost) GetFileSize() int64`

GetFileSize returns the FileSize field if non-nil, zero value otherwise.

### GetFileSizeOk

`func (o *IQDBPost) GetFileSizeOk() (*int64, bool)`

GetFileSizeOk returns a tuple with the FileSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSize

`func (o *IQDBPost) SetFileSize(v int64)`

SetFileSize sets FileSize field to given value.


### GetImageWidth

`func (o *IQDBPost) GetImageWidth() int32`

GetImageWidth returns the ImageWidth field if non-nil, zero value otherwise.

### GetImageWidthOk

`func (o *IQDBPost) GetImageWidthOk() (*int32, bool)`

GetImageWidthOk returns a tuple with the ImageWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageWidth

`func (o *IQDBPost) SetImageWidth(v int32)`

SetImageWidth sets ImageWidth field to given value.


### GetImageHeight

`func (o *IQDBPost) GetImageHeight() int32`

GetImageHeight returns the ImageHeight field if non-nil, zero value otherwise.

### GetImageHeightOk

`func (o *IQDBPost) GetImageHeightOk() (*int32, bool)`

GetImageHeightOk returns a tuple with the ImageHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageHeight

`func (o *IQDBPost) SetImageHeight(v int32)`

SetImageHeight sets ImageHeight field to given value.


### GetParentId

`func (o *IQDBPost) GetParentId() int32`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *IQDBPost) GetParentIdOk() (*int32, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *IQDBPost) SetParentId(v int32)`

SetParentId sets ParentId field to given value.


### SetParentIdNil

`func (o *IQDBPost) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *IQDBPost) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetHasChildren

`func (o *IQDBPost) GetHasChildren() bool`

GetHasChildren returns the HasChildren field if non-nil, zero value otherwise.

### GetHasChildrenOk

`func (o *IQDBPost) GetHasChildrenOk() (*bool, bool)`

GetHasChildrenOk returns a tuple with the HasChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasChildren

`func (o *IQDBPost) SetHasChildren(v bool)`

SetHasChildren sets HasChildren field to given value.


### GetLastCommentedAt

`func (o *IQDBPost) GetLastCommentedAt() time.Time`

GetLastCommentedAt returns the LastCommentedAt field if non-nil, zero value otherwise.

### GetLastCommentedAtOk

`func (o *IQDBPost) GetLastCommentedAtOk() (*time.Time, bool)`

GetLastCommentedAtOk returns a tuple with the LastCommentedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommentedAt

`func (o *IQDBPost) SetLastCommentedAt(v time.Time)`

SetLastCommentedAt sets LastCommentedAt field to given value.


### SetLastCommentedAtNil

`func (o *IQDBPost) SetLastCommentedAtNil(b bool)`

 SetLastCommentedAtNil sets the value for LastCommentedAt to be an explicit nil

### UnsetLastCommentedAt
`func (o *IQDBPost) UnsetLastCommentedAt()`

UnsetLastCommentedAt ensures that no value is present for LastCommentedAt, not even an explicit nil
### GetHasActiveChildren

`func (o *IQDBPost) GetHasActiveChildren() bool`

GetHasActiveChildren returns the HasActiveChildren field if non-nil, zero value otherwise.

### GetHasActiveChildrenOk

`func (o *IQDBPost) GetHasActiveChildrenOk() (*bool, bool)`

GetHasActiveChildrenOk returns a tuple with the HasActiveChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasActiveChildren

`func (o *IQDBPost) SetHasActiveChildren(v bool)`

SetHasActiveChildren sets HasActiveChildren field to given value.


### GetBitFlags

`func (o *IQDBPost) GetBitFlags() float32`

GetBitFlags returns the BitFlags field if non-nil, zero value otherwise.

### GetBitFlagsOk

`func (o *IQDBPost) GetBitFlagsOk() (*float32, bool)`

GetBitFlagsOk returns a tuple with the BitFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitFlags

`func (o *IQDBPost) SetBitFlags(v float32)`

SetBitFlags sets BitFlags field to given value.


### GetTagCountMeta

`func (o *IQDBPost) GetTagCountMeta() int32`

GetTagCountMeta returns the TagCountMeta field if non-nil, zero value otherwise.

### GetTagCountMetaOk

`func (o *IQDBPost) GetTagCountMetaOk() (*int32, bool)`

GetTagCountMetaOk returns a tuple with the TagCountMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCountMeta

`func (o *IQDBPost) SetTagCountMeta(v int32)`

SetTagCountMeta sets TagCountMeta field to given value.


### GetLockedTags

`func (o *IQDBPost) GetLockedTags() string`

GetLockedTags returns the LockedTags field if non-nil, zero value otherwise.

### GetLockedTagsOk

`func (o *IQDBPost) GetLockedTagsOk() (*string, bool)`

GetLockedTagsOk returns a tuple with the LockedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedTags

`func (o *IQDBPost) SetLockedTags(v string)`

SetLockedTags sets LockedTags field to given value.


### SetLockedTagsNil

`func (o *IQDBPost) SetLockedTagsNil(b bool)`

 SetLockedTagsNil sets the value for LockedTags to be an explicit nil

### UnsetLockedTags
`func (o *IQDBPost) UnsetLockedTags()`

UnsetLockedTags ensures that no value is present for LockedTags, not even an explicit nil
### GetTagCountSpecies

`func (o *IQDBPost) GetTagCountSpecies() int32`

GetTagCountSpecies returns the TagCountSpecies field if non-nil, zero value otherwise.

### GetTagCountSpeciesOk

`func (o *IQDBPost) GetTagCountSpeciesOk() (*int32, bool)`

GetTagCountSpeciesOk returns a tuple with the TagCountSpecies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCountSpecies

`func (o *IQDBPost) SetTagCountSpecies(v int32)`

SetTagCountSpecies sets TagCountSpecies field to given value.


### GetTagCountInvalid

`func (o *IQDBPost) GetTagCountInvalid() int32`

GetTagCountInvalid returns the TagCountInvalid field if non-nil, zero value otherwise.

### GetTagCountInvalidOk

`func (o *IQDBPost) GetTagCountInvalidOk() (*int32, bool)`

GetTagCountInvalidOk returns a tuple with the TagCountInvalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCountInvalid

`func (o *IQDBPost) SetTagCountInvalid(v int32)`

SetTagCountInvalid sets TagCountInvalid field to given value.


### GetDescription

`func (o *IQDBPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IQDBPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IQDBPost) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCommentCount

`func (o *IQDBPost) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *IQDBPost) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *IQDBPost) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.


### GetChangeSeq

`func (o *IQDBPost) GetChangeSeq() int32`

GetChangeSeq returns the ChangeSeq field if non-nil, zero value otherwise.

### GetChangeSeqOk

`func (o *IQDBPost) GetChangeSeqOk() (*int32, bool)`

GetChangeSeqOk returns a tuple with the ChangeSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeSeq

`func (o *IQDBPost) SetChangeSeq(v int32)`

SetChangeSeq sets ChangeSeq field to given value.


### GetTagCountLore

`func (o *IQDBPost) GetTagCountLore() int32`

GetTagCountLore returns the TagCountLore field if non-nil, zero value otherwise.

### GetTagCountLoreOk

`func (o *IQDBPost) GetTagCountLoreOk() (*int32, bool)`

GetTagCountLoreOk returns a tuple with the TagCountLore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagCountLore

`func (o *IQDBPost) SetTagCountLore(v int32)`

SetTagCountLore sets TagCountLore field to given value.


### GetBgColor

`func (o *IQDBPost) GetBgColor() string`

GetBgColor returns the BgColor field if non-nil, zero value otherwise.

### GetBgColorOk

`func (o *IQDBPost) GetBgColorOk() (*string, bool)`

GetBgColorOk returns a tuple with the BgColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgColor

`func (o *IQDBPost) SetBgColor(v string)`

SetBgColor sets BgColor field to given value.


### SetBgColorNil

`func (o *IQDBPost) SetBgColorNil(b bool)`

 SetBgColorNil sets the value for BgColor to be an explicit nil

### UnsetBgColor
`func (o *IQDBPost) UnsetBgColor()`

UnsetBgColor ensures that no value is present for BgColor, not even an explicit nil
### GetGeneratedSamples

`func (o *IQDBPost) GetGeneratedSamples() []string`

GetGeneratedSamples returns the GeneratedSamples field if non-nil, zero value otherwise.

### GetGeneratedSamplesOk

`func (o *IQDBPost) GetGeneratedSamplesOk() (*[]string, bool)`

GetGeneratedSamplesOk returns a tuple with the GeneratedSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedSamples

`func (o *IQDBPost) SetGeneratedSamples(v []string)`

SetGeneratedSamples sets GeneratedSamples field to given value.


### GetDuration

`func (o *IQDBPost) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *IQDBPost) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *IQDBPost) SetDuration(v string)`

SetDuration sets Duration field to given value.


### SetDurationNil

`func (o *IQDBPost) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *IQDBPost) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetIsCommentDisabled

`func (o *IQDBPost) GetIsCommentDisabled() bool`

GetIsCommentDisabled returns the IsCommentDisabled field if non-nil, zero value otherwise.

### GetIsCommentDisabledOk

`func (o *IQDBPost) GetIsCommentDisabledOk() (*bool, bool)`

GetIsCommentDisabledOk returns a tuple with the IsCommentDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCommentDisabled

`func (o *IQDBPost) SetIsCommentDisabled(v bool)`

SetIsCommentDisabled sets IsCommentDisabled field to given value.


### GetIsCommentLocked

`func (o *IQDBPost) GetIsCommentLocked() bool`

GetIsCommentLocked returns the IsCommentLocked field if non-nil, zero value otherwise.

### GetIsCommentLockedOk

`func (o *IQDBPost) GetIsCommentLockedOk() (*bool, bool)`

GetIsCommentLockedOk returns a tuple with the IsCommentLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCommentLocked

`func (o *IQDBPost) SetIsCommentLocked(v bool)`

SetIsCommentLocked sets IsCommentLocked field to given value.


### GetHasLarge

`func (o *IQDBPost) GetHasLarge() bool`

GetHasLarge returns the HasLarge field if non-nil, zero value otherwise.

### GetHasLargeOk

`func (o *IQDBPost) GetHasLargeOk() (*bool, bool)`

GetHasLargeOk returns a tuple with the HasLarge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLarge

`func (o *IQDBPost) SetHasLarge(v bool)`

SetHasLarge sets HasLarge field to given value.


### GetHasVisibleChildren

`func (o *IQDBPost) GetHasVisibleChildren() bool`

GetHasVisibleChildren returns the HasVisibleChildren field if non-nil, zero value otherwise.

### GetHasVisibleChildrenOk

`func (o *IQDBPost) GetHasVisibleChildrenOk() (*bool, bool)`

GetHasVisibleChildrenOk returns a tuple with the HasVisibleChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVisibleChildren

`func (o *IQDBPost) SetHasVisibleChildren(v bool)`

SetHasVisibleChildren sets HasVisibleChildren field to given value.


### GetChildrenIds

`func (o *IQDBPost) GetChildrenIds() string`

GetChildrenIds returns the ChildrenIds field if non-nil, zero value otherwise.

### GetChildrenIdsOk

`func (o *IQDBPost) GetChildrenIdsOk() (*string, bool)`

GetChildrenIdsOk returns a tuple with the ChildrenIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildrenIds

`func (o *IQDBPost) SetChildrenIds(v string)`

SetChildrenIds sets ChildrenIds field to given value.


### SetChildrenIdsNil

`func (o *IQDBPost) SetChildrenIdsNil(b bool)`

 SetChildrenIdsNil sets the value for ChildrenIds to be an explicit nil

### UnsetChildrenIds
`func (o *IQDBPost) UnsetChildrenIds()`

UnsetChildrenIds ensures that no value is present for ChildrenIds, not even an explicit nil
### GetPoolIds

`func (o *IQDBPost) GetPoolIds() []int32`

GetPoolIds returns the PoolIds field if non-nil, zero value otherwise.

### GetPoolIdsOk

`func (o *IQDBPost) GetPoolIdsOk() (*[]int32, bool)`

GetPoolIdsOk returns a tuple with the PoolIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolIds

`func (o *IQDBPost) SetPoolIds(v []int32)`

SetPoolIds sets PoolIds field to given value.


### GetIsFavorited

`func (o *IQDBPost) GetIsFavorited() bool`

GetIsFavorited returns the IsFavorited field if non-nil, zero value otherwise.

### GetIsFavoritedOk

`func (o *IQDBPost) GetIsFavoritedOk() (*bool, bool)`

GetIsFavoritedOk returns a tuple with the IsFavorited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorited

`func (o *IQDBPost) SetIsFavorited(v bool)`

SetIsFavorited sets IsFavorited field to given value.


### GetFileUrl

`func (o *IQDBPost) GetFileUrl() string`

GetFileUrl returns the FileUrl field if non-nil, zero value otherwise.

### GetFileUrlOk

`func (o *IQDBPost) GetFileUrlOk() (*string, bool)`

GetFileUrlOk returns a tuple with the FileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileUrl

`func (o *IQDBPost) SetFileUrl(v string)`

SetFileUrl sets FileUrl field to given value.

### HasFileUrl

`func (o *IQDBPost) HasFileUrl() bool`

HasFileUrl returns a boolean if a field has been set.

### GetLargeFileUrl

`func (o *IQDBPost) GetLargeFileUrl() string`

GetLargeFileUrl returns the LargeFileUrl field if non-nil, zero value otherwise.

### GetLargeFileUrlOk

`func (o *IQDBPost) GetLargeFileUrlOk() (*string, bool)`

GetLargeFileUrlOk returns a tuple with the LargeFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLargeFileUrl

`func (o *IQDBPost) SetLargeFileUrl(v string)`

SetLargeFileUrl sets LargeFileUrl field to given value.

### HasLargeFileUrl

`func (o *IQDBPost) HasLargeFileUrl() bool`

HasLargeFileUrl returns a boolean if a field has been set.

### GetPreviewFileUrl

`func (o *IQDBPost) GetPreviewFileUrl() string`

GetPreviewFileUrl returns the PreviewFileUrl field if non-nil, zero value otherwise.

### GetPreviewFileUrlOk

`func (o *IQDBPost) GetPreviewFileUrlOk() (*string, bool)`

GetPreviewFileUrlOk returns a tuple with the PreviewFileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewFileUrl

`func (o *IQDBPost) SetPreviewFileUrl(v string)`

SetPreviewFileUrl sets PreviewFileUrl field to given value.

### HasPreviewFileUrl

`func (o *IQDBPost) HasPreviewFileUrl() bool`

HasPreviewFileUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


