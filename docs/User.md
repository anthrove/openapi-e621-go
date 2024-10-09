# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**Name** | **string** |  | 
**Level** | **float32** |  | 
**BaseUploadLimit** | **int32** |  | 
**PostUploadCount** | **int32** |  | 
**PostUpdateCount** | **int32** |  | 
**NoteUpdateCount** | **int32** |  | 
**IsBanned** | **bool** |  | 
**CanApprovePosts** | **bool** |  | 
**CanUploadFree** | **bool** |  | 
**LevelString** | **string** |  | 
**AvatarId** | **NullableInt32** |  | 
**ArtistVersionCount** | Pointer to **int32** |  | [optional] 
**CommentCount** | Pointer to **int32** |  | [optional] 
**FavoriteCount** | Pointer to **int32** |  | [optional] 
**FlagCount** | Pointer to **float32** |  | [optional] 
**ForumPostCount** | Pointer to **int32** |  | [optional] 
**NegativeFeedbackCount** | Pointer to **int32** |  | [optional] 
**NeutralFeedbackCount** | Pointer to **int32** |  | [optional] 
**PoolVersionCount** | Pointer to **int32** |  | [optional] 
**PositiveFeedbackCount** | Pointer to **int32** |  | [optional] 
**ProfileAbout** | Pointer to **string** |  | [optional] 
**ProfileArtinfo** | Pointer to **string** |  | [optional] 
**UploadLimit** | Pointer to **int32** |  | [optional] 
**WikiPageVersionCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewUser

`func NewUser(id int32, createdAt time.Time, name string, level float32, baseUploadLimit int32, postUploadCount int32, postUpdateCount int32, noteUpdateCount int32, isBanned bool, canApprovePosts bool, canUploadFree bool, levelString string, avatarId NullableInt32, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *User) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.


### GetLevel

`func (o *User) GetLevel() float32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *User) GetLevelOk() (*float32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *User) SetLevel(v float32)`

SetLevel sets Level field to given value.


### GetBaseUploadLimit

`func (o *User) GetBaseUploadLimit() int32`

GetBaseUploadLimit returns the BaseUploadLimit field if non-nil, zero value otherwise.

### GetBaseUploadLimitOk

`func (o *User) GetBaseUploadLimitOk() (*int32, bool)`

GetBaseUploadLimitOk returns a tuple with the BaseUploadLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUploadLimit

`func (o *User) SetBaseUploadLimit(v int32)`

SetBaseUploadLimit sets BaseUploadLimit field to given value.


### GetPostUploadCount

`func (o *User) GetPostUploadCount() int32`

GetPostUploadCount returns the PostUploadCount field if non-nil, zero value otherwise.

### GetPostUploadCountOk

`func (o *User) GetPostUploadCountOk() (*int32, bool)`

GetPostUploadCountOk returns a tuple with the PostUploadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostUploadCount

`func (o *User) SetPostUploadCount(v int32)`

SetPostUploadCount sets PostUploadCount field to given value.


### GetPostUpdateCount

`func (o *User) GetPostUpdateCount() int32`

GetPostUpdateCount returns the PostUpdateCount field if non-nil, zero value otherwise.

### GetPostUpdateCountOk

`func (o *User) GetPostUpdateCountOk() (*int32, bool)`

GetPostUpdateCountOk returns a tuple with the PostUpdateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostUpdateCount

`func (o *User) SetPostUpdateCount(v int32)`

SetPostUpdateCount sets PostUpdateCount field to given value.


### GetNoteUpdateCount

`func (o *User) GetNoteUpdateCount() int32`

GetNoteUpdateCount returns the NoteUpdateCount field if non-nil, zero value otherwise.

### GetNoteUpdateCountOk

`func (o *User) GetNoteUpdateCountOk() (*int32, bool)`

GetNoteUpdateCountOk returns a tuple with the NoteUpdateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteUpdateCount

`func (o *User) SetNoteUpdateCount(v int32)`

SetNoteUpdateCount sets NoteUpdateCount field to given value.


### GetIsBanned

`func (o *User) GetIsBanned() bool`

GetIsBanned returns the IsBanned field if non-nil, zero value otherwise.

### GetIsBannedOk

`func (o *User) GetIsBannedOk() (*bool, bool)`

GetIsBannedOk returns a tuple with the IsBanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBanned

`func (o *User) SetIsBanned(v bool)`

SetIsBanned sets IsBanned field to given value.


### GetCanApprovePosts

`func (o *User) GetCanApprovePosts() bool`

GetCanApprovePosts returns the CanApprovePosts field if non-nil, zero value otherwise.

### GetCanApprovePostsOk

`func (o *User) GetCanApprovePostsOk() (*bool, bool)`

GetCanApprovePostsOk returns a tuple with the CanApprovePosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanApprovePosts

`func (o *User) SetCanApprovePosts(v bool)`

SetCanApprovePosts sets CanApprovePosts field to given value.


### GetCanUploadFree

`func (o *User) GetCanUploadFree() bool`

GetCanUploadFree returns the CanUploadFree field if non-nil, zero value otherwise.

### GetCanUploadFreeOk

`func (o *User) GetCanUploadFreeOk() (*bool, bool)`

GetCanUploadFreeOk returns a tuple with the CanUploadFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUploadFree

`func (o *User) SetCanUploadFree(v bool)`

SetCanUploadFree sets CanUploadFree field to given value.


### GetLevelString

`func (o *User) GetLevelString() string`

GetLevelString returns the LevelString field if non-nil, zero value otherwise.

### GetLevelStringOk

`func (o *User) GetLevelStringOk() (*string, bool)`

GetLevelStringOk returns a tuple with the LevelString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelString

`func (o *User) SetLevelString(v string)`

SetLevelString sets LevelString field to given value.


### GetAvatarId

`func (o *User) GetAvatarId() int32`

GetAvatarId returns the AvatarId field if non-nil, zero value otherwise.

### GetAvatarIdOk

`func (o *User) GetAvatarIdOk() (*int32, bool)`

GetAvatarIdOk returns a tuple with the AvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarId

`func (o *User) SetAvatarId(v int32)`

SetAvatarId sets AvatarId field to given value.


### SetAvatarIdNil

`func (o *User) SetAvatarIdNil(b bool)`

 SetAvatarIdNil sets the value for AvatarId to be an explicit nil

### UnsetAvatarId
`func (o *User) UnsetAvatarId()`

UnsetAvatarId ensures that no value is present for AvatarId, not even an explicit nil
### GetArtistVersionCount

`func (o *User) GetArtistVersionCount() int32`

GetArtistVersionCount returns the ArtistVersionCount field if non-nil, zero value otherwise.

### GetArtistVersionCountOk

`func (o *User) GetArtistVersionCountOk() (*int32, bool)`

GetArtistVersionCountOk returns a tuple with the ArtistVersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistVersionCount

`func (o *User) SetArtistVersionCount(v int32)`

SetArtistVersionCount sets ArtistVersionCount field to given value.

### HasArtistVersionCount

`func (o *User) HasArtistVersionCount() bool`

HasArtistVersionCount returns a boolean if a field has been set.

### GetCommentCount

`func (o *User) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *User) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *User) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.

### HasCommentCount

`func (o *User) HasCommentCount() bool`

HasCommentCount returns a boolean if a field has been set.

### GetFavoriteCount

`func (o *User) GetFavoriteCount() int32`

GetFavoriteCount returns the FavoriteCount field if non-nil, zero value otherwise.

### GetFavoriteCountOk

`func (o *User) GetFavoriteCountOk() (*int32, bool)`

GetFavoriteCountOk returns a tuple with the FavoriteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoriteCount

`func (o *User) SetFavoriteCount(v int32)`

SetFavoriteCount sets FavoriteCount field to given value.

### HasFavoriteCount

`func (o *User) HasFavoriteCount() bool`

HasFavoriteCount returns a boolean if a field has been set.

### GetFlagCount

`func (o *User) GetFlagCount() float32`

GetFlagCount returns the FlagCount field if non-nil, zero value otherwise.

### GetFlagCountOk

`func (o *User) GetFlagCountOk() (*float32, bool)`

GetFlagCountOk returns a tuple with the FlagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagCount

`func (o *User) SetFlagCount(v float32)`

SetFlagCount sets FlagCount field to given value.

### HasFlagCount

`func (o *User) HasFlagCount() bool`

HasFlagCount returns a boolean if a field has been set.

### GetForumPostCount

`func (o *User) GetForumPostCount() int32`

GetForumPostCount returns the ForumPostCount field if non-nil, zero value otherwise.

### GetForumPostCountOk

`func (o *User) GetForumPostCountOk() (*int32, bool)`

GetForumPostCountOk returns a tuple with the ForumPostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumPostCount

`func (o *User) SetForumPostCount(v int32)`

SetForumPostCount sets ForumPostCount field to given value.

### HasForumPostCount

`func (o *User) HasForumPostCount() bool`

HasForumPostCount returns a boolean if a field has been set.

### GetNegativeFeedbackCount

`func (o *User) GetNegativeFeedbackCount() int32`

GetNegativeFeedbackCount returns the NegativeFeedbackCount field if non-nil, zero value otherwise.

### GetNegativeFeedbackCountOk

`func (o *User) GetNegativeFeedbackCountOk() (*int32, bool)`

GetNegativeFeedbackCountOk returns a tuple with the NegativeFeedbackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeFeedbackCount

`func (o *User) SetNegativeFeedbackCount(v int32)`

SetNegativeFeedbackCount sets NegativeFeedbackCount field to given value.

### HasNegativeFeedbackCount

`func (o *User) HasNegativeFeedbackCount() bool`

HasNegativeFeedbackCount returns a boolean if a field has been set.

### GetNeutralFeedbackCount

`func (o *User) GetNeutralFeedbackCount() int32`

GetNeutralFeedbackCount returns the NeutralFeedbackCount field if non-nil, zero value otherwise.

### GetNeutralFeedbackCountOk

`func (o *User) GetNeutralFeedbackCountOk() (*int32, bool)`

GetNeutralFeedbackCountOk returns a tuple with the NeutralFeedbackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeutralFeedbackCount

`func (o *User) SetNeutralFeedbackCount(v int32)`

SetNeutralFeedbackCount sets NeutralFeedbackCount field to given value.

### HasNeutralFeedbackCount

`func (o *User) HasNeutralFeedbackCount() bool`

HasNeutralFeedbackCount returns a boolean if a field has been set.

### GetPoolVersionCount

`func (o *User) GetPoolVersionCount() int32`

GetPoolVersionCount returns the PoolVersionCount field if non-nil, zero value otherwise.

### GetPoolVersionCountOk

`func (o *User) GetPoolVersionCountOk() (*int32, bool)`

GetPoolVersionCountOk returns a tuple with the PoolVersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolVersionCount

`func (o *User) SetPoolVersionCount(v int32)`

SetPoolVersionCount sets PoolVersionCount field to given value.

### HasPoolVersionCount

`func (o *User) HasPoolVersionCount() bool`

HasPoolVersionCount returns a boolean if a field has been set.

### GetPositiveFeedbackCount

`func (o *User) GetPositiveFeedbackCount() int32`

GetPositiveFeedbackCount returns the PositiveFeedbackCount field if non-nil, zero value otherwise.

### GetPositiveFeedbackCountOk

`func (o *User) GetPositiveFeedbackCountOk() (*int32, bool)`

GetPositiveFeedbackCountOk returns a tuple with the PositiveFeedbackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositiveFeedbackCount

`func (o *User) SetPositiveFeedbackCount(v int32)`

SetPositiveFeedbackCount sets PositiveFeedbackCount field to given value.

### HasPositiveFeedbackCount

`func (o *User) HasPositiveFeedbackCount() bool`

HasPositiveFeedbackCount returns a boolean if a field has been set.

### GetProfileAbout

`func (o *User) GetProfileAbout() string`

GetProfileAbout returns the ProfileAbout field if non-nil, zero value otherwise.

### GetProfileAboutOk

`func (o *User) GetProfileAboutOk() (*string, bool)`

GetProfileAboutOk returns a tuple with the ProfileAbout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileAbout

`func (o *User) SetProfileAbout(v string)`

SetProfileAbout sets ProfileAbout field to given value.

### HasProfileAbout

`func (o *User) HasProfileAbout() bool`

HasProfileAbout returns a boolean if a field has been set.

### GetProfileArtinfo

`func (o *User) GetProfileArtinfo() string`

GetProfileArtinfo returns the ProfileArtinfo field if non-nil, zero value otherwise.

### GetProfileArtinfoOk

`func (o *User) GetProfileArtinfoOk() (*string, bool)`

GetProfileArtinfoOk returns a tuple with the ProfileArtinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileArtinfo

`func (o *User) SetProfileArtinfo(v string)`

SetProfileArtinfo sets ProfileArtinfo field to given value.

### HasProfileArtinfo

`func (o *User) HasProfileArtinfo() bool`

HasProfileArtinfo returns a boolean if a field has been set.

### GetUploadLimit

`func (o *User) GetUploadLimit() int32`

GetUploadLimit returns the UploadLimit field if non-nil, zero value otherwise.

### GetUploadLimitOk

`func (o *User) GetUploadLimitOk() (*int32, bool)`

GetUploadLimitOk returns a tuple with the UploadLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadLimit

`func (o *User) SetUploadLimit(v int32)`

SetUploadLimit sets UploadLimit field to given value.

### HasUploadLimit

`func (o *User) HasUploadLimit() bool`

HasUploadLimit returns a boolean if a field has been set.

### GetWikiPageVersionCount

`func (o *User) GetWikiPageVersionCount() int32`

GetWikiPageVersionCount returns the WikiPageVersionCount field if non-nil, zero value otherwise.

### GetWikiPageVersionCountOk

`func (o *User) GetWikiPageVersionCountOk() (*int32, bool)`

GetWikiPageVersionCountOk returns a tuple with the WikiPageVersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikiPageVersionCount

`func (o *User) SetWikiPageVersionCount(v int32)`

SetWikiPageVersionCount sets WikiPageVersionCount field to given value.

### HasWikiPageVersionCount

`func (o *User) HasWikiPageVersionCount() bool`

HasWikiPageVersionCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


