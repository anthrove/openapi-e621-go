# SearchUsers200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **float32** |  | 
**CreatedAt** | **time.Time** |  | 
**Name** | **string** |  | 
**Level** | **float32** |  | 
**BaseUploadLimit** | **float32** |  | 
**PostUploadCount** | **float32** |  | 
**PostUpdateCount** | **float32** |  | 
**NoteUpdateCount** | **float32** |  | 
**IsBanned** | **bool** |  | 
**CanApprovePosts** | **bool** |  | 
**CanUploadFree** | **bool** |  | 
**LevelString** | **string** |  | 
**AvatarId** | **float32** |  | 
**WikiPageVersionCount** | **int32** |  | 
**ArtistVersionCount** | **int32** |  | 
**PoolVersionCount** | **int32** |  | 
**ForumPostCount** | **int32** |  | 
**CommentCount** | **int32** |  | 
**FlagCount** | **int32** |  | 
**FavoriteCount** | **int32** |  | 
**PositiveFeedbackCount** | **int32** |  | 
**NeutralFeedbackCount** | **int32** |  | 
**NegativeFeedbackCount** | **int32** |  | 
**UploadLimit** | **int32** |  | 
**ProfileAbout** | **string** |  | 
**ProfileArtinfo** | **string** |  | 

## Methods

### NewSearchUsers200ResponseInner

`func NewSearchUsers200ResponseInner(id float32, createdAt time.Time, name string, level float32, baseUploadLimit float32, postUploadCount float32, postUpdateCount float32, noteUpdateCount float32, isBanned bool, canApprovePosts bool, canUploadFree bool, levelString string, avatarId float32, wikiPageVersionCount int32, artistVersionCount int32, poolVersionCount int32, forumPostCount int32, commentCount int32, flagCount int32, favoriteCount int32, positiveFeedbackCount int32, neutralFeedbackCount int32, negativeFeedbackCount int32, uploadLimit int32, profileAbout string, profileArtinfo string, ) *SearchUsers200ResponseInner`

NewSearchUsers200ResponseInner instantiates a new SearchUsers200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchUsers200ResponseInnerWithDefaults

`func NewSearchUsers200ResponseInnerWithDefaults() *SearchUsers200ResponseInner`

NewSearchUsers200ResponseInnerWithDefaults instantiates a new SearchUsers200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SearchUsers200ResponseInner) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchUsers200ResponseInner) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchUsers200ResponseInner) SetId(v float32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *SearchUsers200ResponseInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SearchUsers200ResponseInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SearchUsers200ResponseInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetName

`func (o *SearchUsers200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SearchUsers200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SearchUsers200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetLevel

`func (o *SearchUsers200ResponseInner) GetLevel() float32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *SearchUsers200ResponseInner) GetLevelOk() (*float32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *SearchUsers200ResponseInner) SetLevel(v float32)`

SetLevel sets Level field to given value.


### GetBaseUploadLimit

`func (o *SearchUsers200ResponseInner) GetBaseUploadLimit() float32`

GetBaseUploadLimit returns the BaseUploadLimit field if non-nil, zero value otherwise.

### GetBaseUploadLimitOk

`func (o *SearchUsers200ResponseInner) GetBaseUploadLimitOk() (*float32, bool)`

GetBaseUploadLimitOk returns a tuple with the BaseUploadLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUploadLimit

`func (o *SearchUsers200ResponseInner) SetBaseUploadLimit(v float32)`

SetBaseUploadLimit sets BaseUploadLimit field to given value.


### GetPostUploadCount

`func (o *SearchUsers200ResponseInner) GetPostUploadCount() float32`

GetPostUploadCount returns the PostUploadCount field if non-nil, zero value otherwise.

### GetPostUploadCountOk

`func (o *SearchUsers200ResponseInner) GetPostUploadCountOk() (*float32, bool)`

GetPostUploadCountOk returns a tuple with the PostUploadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostUploadCount

`func (o *SearchUsers200ResponseInner) SetPostUploadCount(v float32)`

SetPostUploadCount sets PostUploadCount field to given value.


### GetPostUpdateCount

`func (o *SearchUsers200ResponseInner) GetPostUpdateCount() float32`

GetPostUpdateCount returns the PostUpdateCount field if non-nil, zero value otherwise.

### GetPostUpdateCountOk

`func (o *SearchUsers200ResponseInner) GetPostUpdateCountOk() (*float32, bool)`

GetPostUpdateCountOk returns a tuple with the PostUpdateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostUpdateCount

`func (o *SearchUsers200ResponseInner) SetPostUpdateCount(v float32)`

SetPostUpdateCount sets PostUpdateCount field to given value.


### GetNoteUpdateCount

`func (o *SearchUsers200ResponseInner) GetNoteUpdateCount() float32`

GetNoteUpdateCount returns the NoteUpdateCount field if non-nil, zero value otherwise.

### GetNoteUpdateCountOk

`func (o *SearchUsers200ResponseInner) GetNoteUpdateCountOk() (*float32, bool)`

GetNoteUpdateCountOk returns a tuple with the NoteUpdateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteUpdateCount

`func (o *SearchUsers200ResponseInner) SetNoteUpdateCount(v float32)`

SetNoteUpdateCount sets NoteUpdateCount field to given value.


### GetIsBanned

`func (o *SearchUsers200ResponseInner) GetIsBanned() bool`

GetIsBanned returns the IsBanned field if non-nil, zero value otherwise.

### GetIsBannedOk

`func (o *SearchUsers200ResponseInner) GetIsBannedOk() (*bool, bool)`

GetIsBannedOk returns a tuple with the IsBanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBanned

`func (o *SearchUsers200ResponseInner) SetIsBanned(v bool)`

SetIsBanned sets IsBanned field to given value.


### GetCanApprovePosts

`func (o *SearchUsers200ResponseInner) GetCanApprovePosts() bool`

GetCanApprovePosts returns the CanApprovePosts field if non-nil, zero value otherwise.

### GetCanApprovePostsOk

`func (o *SearchUsers200ResponseInner) GetCanApprovePostsOk() (*bool, bool)`

GetCanApprovePostsOk returns a tuple with the CanApprovePosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanApprovePosts

`func (o *SearchUsers200ResponseInner) SetCanApprovePosts(v bool)`

SetCanApprovePosts sets CanApprovePosts field to given value.


### GetCanUploadFree

`func (o *SearchUsers200ResponseInner) GetCanUploadFree() bool`

GetCanUploadFree returns the CanUploadFree field if non-nil, zero value otherwise.

### GetCanUploadFreeOk

`func (o *SearchUsers200ResponseInner) GetCanUploadFreeOk() (*bool, bool)`

GetCanUploadFreeOk returns a tuple with the CanUploadFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUploadFree

`func (o *SearchUsers200ResponseInner) SetCanUploadFree(v bool)`

SetCanUploadFree sets CanUploadFree field to given value.


### GetLevelString

`func (o *SearchUsers200ResponseInner) GetLevelString() string`

GetLevelString returns the LevelString field if non-nil, zero value otherwise.

### GetLevelStringOk

`func (o *SearchUsers200ResponseInner) GetLevelStringOk() (*string, bool)`

GetLevelStringOk returns a tuple with the LevelString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelString

`func (o *SearchUsers200ResponseInner) SetLevelString(v string)`

SetLevelString sets LevelString field to given value.


### GetAvatarId

`func (o *SearchUsers200ResponseInner) GetAvatarId() float32`

GetAvatarId returns the AvatarId field if non-nil, zero value otherwise.

### GetAvatarIdOk

`func (o *SearchUsers200ResponseInner) GetAvatarIdOk() (*float32, bool)`

GetAvatarIdOk returns a tuple with the AvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarId

`func (o *SearchUsers200ResponseInner) SetAvatarId(v float32)`

SetAvatarId sets AvatarId field to given value.


### GetWikiPageVersionCount

`func (o *SearchUsers200ResponseInner) GetWikiPageVersionCount() int32`

GetWikiPageVersionCount returns the WikiPageVersionCount field if non-nil, zero value otherwise.

### GetWikiPageVersionCountOk

`func (o *SearchUsers200ResponseInner) GetWikiPageVersionCountOk() (*int32, bool)`

GetWikiPageVersionCountOk returns a tuple with the WikiPageVersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikiPageVersionCount

`func (o *SearchUsers200ResponseInner) SetWikiPageVersionCount(v int32)`

SetWikiPageVersionCount sets WikiPageVersionCount field to given value.


### GetArtistVersionCount

`func (o *SearchUsers200ResponseInner) GetArtistVersionCount() int32`

GetArtistVersionCount returns the ArtistVersionCount field if non-nil, zero value otherwise.

### GetArtistVersionCountOk

`func (o *SearchUsers200ResponseInner) GetArtistVersionCountOk() (*int32, bool)`

GetArtistVersionCountOk returns a tuple with the ArtistVersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistVersionCount

`func (o *SearchUsers200ResponseInner) SetArtistVersionCount(v int32)`

SetArtistVersionCount sets ArtistVersionCount field to given value.


### GetPoolVersionCount

`func (o *SearchUsers200ResponseInner) GetPoolVersionCount() int32`

GetPoolVersionCount returns the PoolVersionCount field if non-nil, zero value otherwise.

### GetPoolVersionCountOk

`func (o *SearchUsers200ResponseInner) GetPoolVersionCountOk() (*int32, bool)`

GetPoolVersionCountOk returns a tuple with the PoolVersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolVersionCount

`func (o *SearchUsers200ResponseInner) SetPoolVersionCount(v int32)`

SetPoolVersionCount sets PoolVersionCount field to given value.


### GetForumPostCount

`func (o *SearchUsers200ResponseInner) GetForumPostCount() int32`

GetForumPostCount returns the ForumPostCount field if non-nil, zero value otherwise.

### GetForumPostCountOk

`func (o *SearchUsers200ResponseInner) GetForumPostCountOk() (*int32, bool)`

GetForumPostCountOk returns a tuple with the ForumPostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumPostCount

`func (o *SearchUsers200ResponseInner) SetForumPostCount(v int32)`

SetForumPostCount sets ForumPostCount field to given value.


### GetCommentCount

`func (o *SearchUsers200ResponseInner) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *SearchUsers200ResponseInner) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *SearchUsers200ResponseInner) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.


### GetFlagCount

`func (o *SearchUsers200ResponseInner) GetFlagCount() int32`

GetFlagCount returns the FlagCount field if non-nil, zero value otherwise.

### GetFlagCountOk

`func (o *SearchUsers200ResponseInner) GetFlagCountOk() (*int32, bool)`

GetFlagCountOk returns a tuple with the FlagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagCount

`func (o *SearchUsers200ResponseInner) SetFlagCount(v int32)`

SetFlagCount sets FlagCount field to given value.


### GetFavoriteCount

`func (o *SearchUsers200ResponseInner) GetFavoriteCount() int32`

GetFavoriteCount returns the FavoriteCount field if non-nil, zero value otherwise.

### GetFavoriteCountOk

`func (o *SearchUsers200ResponseInner) GetFavoriteCountOk() (*int32, bool)`

GetFavoriteCountOk returns a tuple with the FavoriteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoriteCount

`func (o *SearchUsers200ResponseInner) SetFavoriteCount(v int32)`

SetFavoriteCount sets FavoriteCount field to given value.


### GetPositiveFeedbackCount

`func (o *SearchUsers200ResponseInner) GetPositiveFeedbackCount() int32`

GetPositiveFeedbackCount returns the PositiveFeedbackCount field if non-nil, zero value otherwise.

### GetPositiveFeedbackCountOk

`func (o *SearchUsers200ResponseInner) GetPositiveFeedbackCountOk() (*int32, bool)`

GetPositiveFeedbackCountOk returns a tuple with the PositiveFeedbackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositiveFeedbackCount

`func (o *SearchUsers200ResponseInner) SetPositiveFeedbackCount(v int32)`

SetPositiveFeedbackCount sets PositiveFeedbackCount field to given value.


### GetNeutralFeedbackCount

`func (o *SearchUsers200ResponseInner) GetNeutralFeedbackCount() int32`

GetNeutralFeedbackCount returns the NeutralFeedbackCount field if non-nil, zero value otherwise.

### GetNeutralFeedbackCountOk

`func (o *SearchUsers200ResponseInner) GetNeutralFeedbackCountOk() (*int32, bool)`

GetNeutralFeedbackCountOk returns a tuple with the NeutralFeedbackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeutralFeedbackCount

`func (o *SearchUsers200ResponseInner) SetNeutralFeedbackCount(v int32)`

SetNeutralFeedbackCount sets NeutralFeedbackCount field to given value.


### GetNegativeFeedbackCount

`func (o *SearchUsers200ResponseInner) GetNegativeFeedbackCount() int32`

GetNegativeFeedbackCount returns the NegativeFeedbackCount field if non-nil, zero value otherwise.

### GetNegativeFeedbackCountOk

`func (o *SearchUsers200ResponseInner) GetNegativeFeedbackCountOk() (*int32, bool)`

GetNegativeFeedbackCountOk returns a tuple with the NegativeFeedbackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeFeedbackCount

`func (o *SearchUsers200ResponseInner) SetNegativeFeedbackCount(v int32)`

SetNegativeFeedbackCount sets NegativeFeedbackCount field to given value.


### GetUploadLimit

`func (o *SearchUsers200ResponseInner) GetUploadLimit() int32`

GetUploadLimit returns the UploadLimit field if non-nil, zero value otherwise.

### GetUploadLimitOk

`func (o *SearchUsers200ResponseInner) GetUploadLimitOk() (*int32, bool)`

GetUploadLimitOk returns a tuple with the UploadLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadLimit

`func (o *SearchUsers200ResponseInner) SetUploadLimit(v int32)`

SetUploadLimit sets UploadLimit field to given value.


### GetProfileAbout

`func (o *SearchUsers200ResponseInner) GetProfileAbout() string`

GetProfileAbout returns the ProfileAbout field if non-nil, zero value otherwise.

### GetProfileAboutOk

`func (o *SearchUsers200ResponseInner) GetProfileAboutOk() (*string, bool)`

GetProfileAboutOk returns a tuple with the ProfileAbout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileAbout

`func (o *SearchUsers200ResponseInner) SetProfileAbout(v string)`

SetProfileAbout sets ProfileAbout field to given value.


### GetProfileArtinfo

`func (o *SearchUsers200ResponseInner) GetProfileArtinfo() string`

GetProfileArtinfo returns the ProfileArtinfo field if non-nil, zero value otherwise.

### GetProfileArtinfoOk

`func (o *SearchUsers200ResponseInner) GetProfileArtinfoOk() (*string, bool)`

GetProfileArtinfoOk returns a tuple with the ProfileArtinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileArtinfo

`func (o *SearchUsers200ResponseInner) SetProfileArtinfo(v string)`

SetProfileArtinfo sets ProfileArtinfo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


