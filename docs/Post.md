# Post

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**File** | [**PostFile**](PostFile.md) |  | 
**Preview** | [**PostPreview**](PostPreview.md) |  | 
**Sample** | [**PostSample**](PostSample.md) |  | 
**Score** | [**PostScore**](PostScore.md) |  | 
**Tags** | [**PostTags**](PostTags.md) |  | 
**LockedTags** | **[]string** |  | 
**ChangeSeq** | **float32** |  | 
**Flags** | [**PostFlags**](PostFlags.md) |  | 
**Rating** | [**Ratings**](Ratings.md) |  | 
**FavCount** | **float32** |  | 
**Sources** | **[]string** |  | 
**Pools** | **[]float32** |  | 
**Relationships** | [**PostRelationships**](PostRelationships.md) |  | 
**ApproverId** | **NullableFloat32** |  | 
**UploaderId** | **float32** |  | 
**Description** | **string** |  | 
**CommentCount** | **float32** |  | 
**IsFavorited** | **bool** |  | 
**HasNotes** | **bool** |  | 
**Duration** | **NullableFloat32** |  | 

## Methods

### NewPost

`func NewPost(id int32, createdAt time.Time, updatedAt time.Time, file PostFile, preview PostPreview, sample PostSample, score PostScore, tags PostTags, lockedTags []string, changeSeq float32, flags PostFlags, rating Ratings, favCount float32, sources []string, pools []float32, relationships PostRelationships, approverId NullableFloat32, uploaderId float32, description string, commentCount float32, isFavorited bool, hasNotes bool, duration NullableFloat32, ) *Post`

NewPost instantiates a new Post object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostWithDefaults

`func NewPostWithDefaults() *Post`

NewPostWithDefaults instantiates a new Post object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Post) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Post) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Post) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *Post) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Post) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Post) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Post) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Post) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Post) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetFile

`func (o *Post) GetFile() PostFile`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *Post) GetFileOk() (*PostFile, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *Post) SetFile(v PostFile)`

SetFile sets File field to given value.


### GetPreview

`func (o *Post) GetPreview() PostPreview`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *Post) GetPreviewOk() (*PostPreview, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *Post) SetPreview(v PostPreview)`

SetPreview sets Preview field to given value.


### GetSample

`func (o *Post) GetSample() PostSample`

GetSample returns the Sample field if non-nil, zero value otherwise.

### GetSampleOk

`func (o *Post) GetSampleOk() (*PostSample, bool)`

GetSampleOk returns a tuple with the Sample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSample

`func (o *Post) SetSample(v PostSample)`

SetSample sets Sample field to given value.


### GetScore

`func (o *Post) GetScore() PostScore`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Post) GetScoreOk() (*PostScore, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *Post) SetScore(v PostScore)`

SetScore sets Score field to given value.


### GetTags

`func (o *Post) GetTags() PostTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Post) GetTagsOk() (*PostTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Post) SetTags(v PostTags)`

SetTags sets Tags field to given value.


### GetLockedTags

`func (o *Post) GetLockedTags() []string`

GetLockedTags returns the LockedTags field if non-nil, zero value otherwise.

### GetLockedTagsOk

`func (o *Post) GetLockedTagsOk() (*[]string, bool)`

GetLockedTagsOk returns a tuple with the LockedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockedTags

`func (o *Post) SetLockedTags(v []string)`

SetLockedTags sets LockedTags field to given value.


### GetChangeSeq

`func (o *Post) GetChangeSeq() float32`

GetChangeSeq returns the ChangeSeq field if non-nil, zero value otherwise.

### GetChangeSeqOk

`func (o *Post) GetChangeSeqOk() (*float32, bool)`

GetChangeSeqOk returns a tuple with the ChangeSeq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeSeq

`func (o *Post) SetChangeSeq(v float32)`

SetChangeSeq sets ChangeSeq field to given value.


### GetFlags

`func (o *Post) GetFlags() PostFlags`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *Post) GetFlagsOk() (*PostFlags, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *Post) SetFlags(v PostFlags)`

SetFlags sets Flags field to given value.


### GetRating

`func (o *Post) GetRating() Ratings`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *Post) GetRatingOk() (*Ratings, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *Post) SetRating(v Ratings)`

SetRating sets Rating field to given value.


### GetFavCount

`func (o *Post) GetFavCount() float32`

GetFavCount returns the FavCount field if non-nil, zero value otherwise.

### GetFavCountOk

`func (o *Post) GetFavCountOk() (*float32, bool)`

GetFavCountOk returns a tuple with the FavCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavCount

`func (o *Post) SetFavCount(v float32)`

SetFavCount sets FavCount field to given value.


### GetSources

`func (o *Post) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *Post) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *Post) SetSources(v []string)`

SetSources sets Sources field to given value.


### GetPools

`func (o *Post) GetPools() []float32`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *Post) GetPoolsOk() (*[]float32, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *Post) SetPools(v []float32)`

SetPools sets Pools field to given value.


### GetRelationships

`func (o *Post) GetRelationships() PostRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *Post) GetRelationshipsOk() (*PostRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *Post) SetRelationships(v PostRelationships)`

SetRelationships sets Relationships field to given value.


### GetApproverId

`func (o *Post) GetApproverId() float32`

GetApproverId returns the ApproverId field if non-nil, zero value otherwise.

### GetApproverIdOk

`func (o *Post) GetApproverIdOk() (*float32, bool)`

GetApproverIdOk returns a tuple with the ApproverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverId

`func (o *Post) SetApproverId(v float32)`

SetApproverId sets ApproverId field to given value.


### SetApproverIdNil

`func (o *Post) SetApproverIdNil(b bool)`

 SetApproverIdNil sets the value for ApproverId to be an explicit nil

### UnsetApproverId
`func (o *Post) UnsetApproverId()`

UnsetApproverId ensures that no value is present for ApproverId, not even an explicit nil
### GetUploaderId

`func (o *Post) GetUploaderId() float32`

GetUploaderId returns the UploaderId field if non-nil, zero value otherwise.

### GetUploaderIdOk

`func (o *Post) GetUploaderIdOk() (*float32, bool)`

GetUploaderIdOk returns a tuple with the UploaderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploaderId

`func (o *Post) SetUploaderId(v float32)`

SetUploaderId sets UploaderId field to given value.


### GetDescription

`func (o *Post) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Post) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Post) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCommentCount

`func (o *Post) GetCommentCount() float32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *Post) GetCommentCountOk() (*float32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *Post) SetCommentCount(v float32)`

SetCommentCount sets CommentCount field to given value.


### GetIsFavorited

`func (o *Post) GetIsFavorited() bool`

GetIsFavorited returns the IsFavorited field if non-nil, zero value otherwise.

### GetIsFavoritedOk

`func (o *Post) GetIsFavoritedOk() (*bool, bool)`

GetIsFavoritedOk returns a tuple with the IsFavorited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFavorited

`func (o *Post) SetIsFavorited(v bool)`

SetIsFavorited sets IsFavorited field to given value.


### GetHasNotes

`func (o *Post) GetHasNotes() bool`

GetHasNotes returns the HasNotes field if non-nil, zero value otherwise.

### GetHasNotesOk

`func (o *Post) GetHasNotesOk() (*bool, bool)`

GetHasNotesOk returns a tuple with the HasNotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNotes

`func (o *Post) SetHasNotes(v bool)`

SetHasNotes sets HasNotes field to given value.


### GetDuration

`func (o *Post) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Post) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Post) SetDuration(v float32)`

SetDuration sets Duration field to given value.


### SetDurationNil

`func (o *Post) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *Post) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


