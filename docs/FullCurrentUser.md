# FullCurrentUser

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
**AvatarId** | **int32** |  | 
**ArtistVersionCount** | **int32** |  | 
**CommentCount** | **int32** |  | 
**FavoritesCount** | Pointer to **int32** |  | [optional] 
**FlagCount** | **int32** |  | 
**ForumPostCount** | **int32** |  | 
**NegativeFeedbackCount** | **int32** |  | 
**NeutralFeedbackCount** | **int32** |  | 
**PoolVersionCount** | **int32** |  | 
**PositiveFeedbackCount** | **int32** |  | 
**ProfileAbout** | **string** |  | 
**ProfileArtinfo** | **string** |  | 
**UploadLimit** | **int32** |  | 
**WikiPageVersionCount** | **int32** |  | 
**BlacklistUsers** | **bool** |  | 
**DescriptionCollapsedInitially** | **bool** |  | 
**HideComments** | **bool** |  | 
**ShowHiddenComments** | **bool** |  | 
**ShowPostStatistics** | **bool** |  | 
**ReceiveEmailNotifications** | **bool** |  | 
**EnableKeyboardNavigation** | **bool** |  | 
**EnablePrivacyMode** | **bool** |  | 
**StyleUsernames** | **bool** |  | 
**EnableAutoComplete** | **bool** |  | 
**DisableCroppedThumbnails** | **bool** |  | 
**EnableSafeMode** | **bool** |  | 
**DisableResponsiveMode** | **bool** |  | 
**NoFlagging** | **bool** |  | 
**DisableUserDmails** | **bool** |  | 
**EnableCompactUploader** | **bool** |  | 
**ReplacementsBeta** | **bool** |  | 
**UpdatedAt** | **time.Time** |  | 
**Email** | **string** |  | 
**LastLoggedInAt** | **time.Time** |  | 
**LastForumReadAt** | **time.Time** |  | 
**RecentTags** | **string** |  | 
**CommentThreshold** | **float32** |  | 
**DefaultImageSizedefaultImageSize** | Pointer to **string** |  | [optional] 
**FavoriteTags** | **string** |  | 
**BlacklistedTags** | **string** |  | 
**TimeZone** | **string** |  | 
**PerPage** | **int32** |  | 
**CustomStyle** | **string** |  | 
**FavoriteCount** | **int32** |  | 
**ApiRegenMultiplier** | **float32** |  | 
**ApiBurstLimit** | **float32** |  | 
**RemainingApiLimit** | **float32** |  | 
**StatementTimeout** | **float32** |  | 
**FavoriteLimit** | **int32** |  | 
**TagQueryLimit** | **int32** |  | 
**HasMail** | **bool** |  | 

## Methods

### NewFullCurrentUser

`func NewFullCurrentUser(id int32, createdAt time.Time, name string, level float32, baseUploadLimit int32, postUploadCount int32, postUpdateCount int32, noteUpdateCount int32, isBanned bool, canApprovePosts bool, canUploadFree bool, levelString string, avatarId int32, artistVersionCount int32, commentCount int32, flagCount int32, forumPostCount int32, negativeFeedbackCount int32, neutralFeedbackCount int32, poolVersionCount int32, positiveFeedbackCount int32, profileAbout string, profileArtinfo string, uploadLimit int32, wikiPageVersionCount int32, blacklistUsers bool, descriptionCollapsedInitially bool, hideComments bool, showHiddenComments bool, showPostStatistics bool, receiveEmailNotifications bool, enableKeyboardNavigation bool, enablePrivacyMode bool, styleUsernames bool, enableAutoComplete bool, disableCroppedThumbnails bool, enableSafeMode bool, disableResponsiveMode bool, noFlagging bool, disableUserDmails bool, enableCompactUploader bool, replacementsBeta bool, updatedAt time.Time, email string, lastLoggedInAt time.Time, lastForumReadAt time.Time, recentTags string, commentThreshold float32, favoriteTags string, blacklistedTags string, timeZone string, perPage int32, customStyle string, favoriteCount int32, apiRegenMultiplier float32, apiBurstLimit float32, remainingApiLimit float32, statementTimeout float32, favoriteLimit int32, tagQueryLimit int32, hasMail bool, ) *FullCurrentUser`

NewFullCurrentUser instantiates a new FullCurrentUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFullCurrentUserWithDefaults

`func NewFullCurrentUserWithDefaults() *FullCurrentUser`

NewFullCurrentUserWithDefaults instantiates a new FullCurrentUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FullCurrentUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FullCurrentUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FullCurrentUser) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *FullCurrentUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FullCurrentUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FullCurrentUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetName

`func (o *FullCurrentUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FullCurrentUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FullCurrentUser) SetName(v string)`

SetName sets Name field to given value.


### GetLevel

`func (o *FullCurrentUser) GetLevel() float32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *FullCurrentUser) GetLevelOk() (*float32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *FullCurrentUser) SetLevel(v float32)`

SetLevel sets Level field to given value.


### GetBaseUploadLimit

`func (o *FullCurrentUser) GetBaseUploadLimit() int32`

GetBaseUploadLimit returns the BaseUploadLimit field if non-nil, zero value otherwise.

### GetBaseUploadLimitOk

`func (o *FullCurrentUser) GetBaseUploadLimitOk() (*int32, bool)`

GetBaseUploadLimitOk returns a tuple with the BaseUploadLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUploadLimit

`func (o *FullCurrentUser) SetBaseUploadLimit(v int32)`

SetBaseUploadLimit sets BaseUploadLimit field to given value.


### GetPostUploadCount

`func (o *FullCurrentUser) GetPostUploadCount() int32`

GetPostUploadCount returns the PostUploadCount field if non-nil, zero value otherwise.

### GetPostUploadCountOk

`func (o *FullCurrentUser) GetPostUploadCountOk() (*int32, bool)`

GetPostUploadCountOk returns a tuple with the PostUploadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostUploadCount

`func (o *FullCurrentUser) SetPostUploadCount(v int32)`

SetPostUploadCount sets PostUploadCount field to given value.


### GetPostUpdateCount

`func (o *FullCurrentUser) GetPostUpdateCount() int32`

GetPostUpdateCount returns the PostUpdateCount field if non-nil, zero value otherwise.

### GetPostUpdateCountOk

`func (o *FullCurrentUser) GetPostUpdateCountOk() (*int32, bool)`

GetPostUpdateCountOk returns a tuple with the PostUpdateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostUpdateCount

`func (o *FullCurrentUser) SetPostUpdateCount(v int32)`

SetPostUpdateCount sets PostUpdateCount field to given value.


### GetNoteUpdateCount

`func (o *FullCurrentUser) GetNoteUpdateCount() int32`

GetNoteUpdateCount returns the NoteUpdateCount field if non-nil, zero value otherwise.

### GetNoteUpdateCountOk

`func (o *FullCurrentUser) GetNoteUpdateCountOk() (*int32, bool)`

GetNoteUpdateCountOk returns a tuple with the NoteUpdateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteUpdateCount

`func (o *FullCurrentUser) SetNoteUpdateCount(v int32)`

SetNoteUpdateCount sets NoteUpdateCount field to given value.


### GetIsBanned

`func (o *FullCurrentUser) GetIsBanned() bool`

GetIsBanned returns the IsBanned field if non-nil, zero value otherwise.

### GetIsBannedOk

`func (o *FullCurrentUser) GetIsBannedOk() (*bool, bool)`

GetIsBannedOk returns a tuple with the IsBanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBanned

`func (o *FullCurrentUser) SetIsBanned(v bool)`

SetIsBanned sets IsBanned field to given value.


### GetCanApprovePosts

`func (o *FullCurrentUser) GetCanApprovePosts() bool`

GetCanApprovePosts returns the CanApprovePosts field if non-nil, zero value otherwise.

### GetCanApprovePostsOk

`func (o *FullCurrentUser) GetCanApprovePostsOk() (*bool, bool)`

GetCanApprovePostsOk returns a tuple with the CanApprovePosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanApprovePosts

`func (o *FullCurrentUser) SetCanApprovePosts(v bool)`

SetCanApprovePosts sets CanApprovePosts field to given value.


### GetCanUploadFree

`func (o *FullCurrentUser) GetCanUploadFree() bool`

GetCanUploadFree returns the CanUploadFree field if non-nil, zero value otherwise.

### GetCanUploadFreeOk

`func (o *FullCurrentUser) GetCanUploadFreeOk() (*bool, bool)`

GetCanUploadFreeOk returns a tuple with the CanUploadFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUploadFree

`func (o *FullCurrentUser) SetCanUploadFree(v bool)`

SetCanUploadFree sets CanUploadFree field to given value.


### GetLevelString

`func (o *FullCurrentUser) GetLevelString() string`

GetLevelString returns the LevelString field if non-nil, zero value otherwise.

### GetLevelStringOk

`func (o *FullCurrentUser) GetLevelStringOk() (*string, bool)`

GetLevelStringOk returns a tuple with the LevelString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelString

`func (o *FullCurrentUser) SetLevelString(v string)`

SetLevelString sets LevelString field to given value.


### GetAvatarId

`func (o *FullCurrentUser) GetAvatarId() int32`

GetAvatarId returns the AvatarId field if non-nil, zero value otherwise.

### GetAvatarIdOk

`func (o *FullCurrentUser) GetAvatarIdOk() (*int32, bool)`

GetAvatarIdOk returns a tuple with the AvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarId

`func (o *FullCurrentUser) SetAvatarId(v int32)`

SetAvatarId sets AvatarId field to given value.


### GetArtistVersionCount

`func (o *FullCurrentUser) GetArtistVersionCount() int32`

GetArtistVersionCount returns the ArtistVersionCount field if non-nil, zero value otherwise.

### GetArtistVersionCountOk

`func (o *FullCurrentUser) GetArtistVersionCountOk() (*int32, bool)`

GetArtistVersionCountOk returns a tuple with the ArtistVersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistVersionCount

`func (o *FullCurrentUser) SetArtistVersionCount(v int32)`

SetArtistVersionCount sets ArtistVersionCount field to given value.


### GetCommentCount

`func (o *FullCurrentUser) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *FullCurrentUser) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *FullCurrentUser) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.


### GetFavoritesCount

`func (o *FullCurrentUser) GetFavoritesCount() int32`

GetFavoritesCount returns the FavoritesCount field if non-nil, zero value otherwise.

### GetFavoritesCountOk

`func (o *FullCurrentUser) GetFavoritesCountOk() (*int32, bool)`

GetFavoritesCountOk returns a tuple with the FavoritesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoritesCount

`func (o *FullCurrentUser) SetFavoritesCount(v int32)`

SetFavoritesCount sets FavoritesCount field to given value.

### HasFavoritesCount

`func (o *FullCurrentUser) HasFavoritesCount() bool`

HasFavoritesCount returns a boolean if a field has been set.

### GetFlagCount

`func (o *FullCurrentUser) GetFlagCount() int32`

GetFlagCount returns the FlagCount field if non-nil, zero value otherwise.

### GetFlagCountOk

`func (o *FullCurrentUser) GetFlagCountOk() (*int32, bool)`

GetFlagCountOk returns a tuple with the FlagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagCount

`func (o *FullCurrentUser) SetFlagCount(v int32)`

SetFlagCount sets FlagCount field to given value.


### GetForumPostCount

`func (o *FullCurrentUser) GetForumPostCount() int32`

GetForumPostCount returns the ForumPostCount field if non-nil, zero value otherwise.

### GetForumPostCountOk

`func (o *FullCurrentUser) GetForumPostCountOk() (*int32, bool)`

GetForumPostCountOk returns a tuple with the ForumPostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumPostCount

`func (o *FullCurrentUser) SetForumPostCount(v int32)`

SetForumPostCount sets ForumPostCount field to given value.


### GetNegativeFeedbackCount

`func (o *FullCurrentUser) GetNegativeFeedbackCount() int32`

GetNegativeFeedbackCount returns the NegativeFeedbackCount field if non-nil, zero value otherwise.

### GetNegativeFeedbackCountOk

`func (o *FullCurrentUser) GetNegativeFeedbackCountOk() (*int32, bool)`

GetNegativeFeedbackCountOk returns a tuple with the NegativeFeedbackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeFeedbackCount

`func (o *FullCurrentUser) SetNegativeFeedbackCount(v int32)`

SetNegativeFeedbackCount sets NegativeFeedbackCount field to given value.


### GetNeutralFeedbackCount

`func (o *FullCurrentUser) GetNeutralFeedbackCount() int32`

GetNeutralFeedbackCount returns the NeutralFeedbackCount field if non-nil, zero value otherwise.

### GetNeutralFeedbackCountOk

`func (o *FullCurrentUser) GetNeutralFeedbackCountOk() (*int32, bool)`

GetNeutralFeedbackCountOk returns a tuple with the NeutralFeedbackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeutralFeedbackCount

`func (o *FullCurrentUser) SetNeutralFeedbackCount(v int32)`

SetNeutralFeedbackCount sets NeutralFeedbackCount field to given value.


### GetPoolVersionCount

`func (o *FullCurrentUser) GetPoolVersionCount() int32`

GetPoolVersionCount returns the PoolVersionCount field if non-nil, zero value otherwise.

### GetPoolVersionCountOk

`func (o *FullCurrentUser) GetPoolVersionCountOk() (*int32, bool)`

GetPoolVersionCountOk returns a tuple with the PoolVersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolVersionCount

`func (o *FullCurrentUser) SetPoolVersionCount(v int32)`

SetPoolVersionCount sets PoolVersionCount field to given value.


### GetPositiveFeedbackCount

`func (o *FullCurrentUser) GetPositiveFeedbackCount() int32`

GetPositiveFeedbackCount returns the PositiveFeedbackCount field if non-nil, zero value otherwise.

### GetPositiveFeedbackCountOk

`func (o *FullCurrentUser) GetPositiveFeedbackCountOk() (*int32, bool)`

GetPositiveFeedbackCountOk returns a tuple with the PositiveFeedbackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositiveFeedbackCount

`func (o *FullCurrentUser) SetPositiveFeedbackCount(v int32)`

SetPositiveFeedbackCount sets PositiveFeedbackCount field to given value.


### GetProfileAbout

`func (o *FullCurrentUser) GetProfileAbout() string`

GetProfileAbout returns the ProfileAbout field if non-nil, zero value otherwise.

### GetProfileAboutOk

`func (o *FullCurrentUser) GetProfileAboutOk() (*string, bool)`

GetProfileAboutOk returns a tuple with the ProfileAbout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileAbout

`func (o *FullCurrentUser) SetProfileAbout(v string)`

SetProfileAbout sets ProfileAbout field to given value.


### GetProfileArtinfo

`func (o *FullCurrentUser) GetProfileArtinfo() string`

GetProfileArtinfo returns the ProfileArtinfo field if non-nil, zero value otherwise.

### GetProfileArtinfoOk

`func (o *FullCurrentUser) GetProfileArtinfoOk() (*string, bool)`

GetProfileArtinfoOk returns a tuple with the ProfileArtinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileArtinfo

`func (o *FullCurrentUser) SetProfileArtinfo(v string)`

SetProfileArtinfo sets ProfileArtinfo field to given value.


### GetUploadLimit

`func (o *FullCurrentUser) GetUploadLimit() int32`

GetUploadLimit returns the UploadLimit field if non-nil, zero value otherwise.

### GetUploadLimitOk

`func (o *FullCurrentUser) GetUploadLimitOk() (*int32, bool)`

GetUploadLimitOk returns a tuple with the UploadLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadLimit

`func (o *FullCurrentUser) SetUploadLimit(v int32)`

SetUploadLimit sets UploadLimit field to given value.


### GetWikiPageVersionCount

`func (o *FullCurrentUser) GetWikiPageVersionCount() int32`

GetWikiPageVersionCount returns the WikiPageVersionCount field if non-nil, zero value otherwise.

### GetWikiPageVersionCountOk

`func (o *FullCurrentUser) GetWikiPageVersionCountOk() (*int32, bool)`

GetWikiPageVersionCountOk returns a tuple with the WikiPageVersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikiPageVersionCount

`func (o *FullCurrentUser) SetWikiPageVersionCount(v int32)`

SetWikiPageVersionCount sets WikiPageVersionCount field to given value.


### GetBlacklistUsers

`func (o *FullCurrentUser) GetBlacklistUsers() bool`

GetBlacklistUsers returns the BlacklistUsers field if non-nil, zero value otherwise.

### GetBlacklistUsersOk

`func (o *FullCurrentUser) GetBlacklistUsersOk() (*bool, bool)`

GetBlacklistUsersOk returns a tuple with the BlacklistUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistUsers

`func (o *FullCurrentUser) SetBlacklistUsers(v bool)`

SetBlacklistUsers sets BlacklistUsers field to given value.


### GetDescriptionCollapsedInitially

`func (o *FullCurrentUser) GetDescriptionCollapsedInitially() bool`

GetDescriptionCollapsedInitially returns the DescriptionCollapsedInitially field if non-nil, zero value otherwise.

### GetDescriptionCollapsedInitiallyOk

`func (o *FullCurrentUser) GetDescriptionCollapsedInitiallyOk() (*bool, bool)`

GetDescriptionCollapsedInitiallyOk returns a tuple with the DescriptionCollapsedInitially field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionCollapsedInitially

`func (o *FullCurrentUser) SetDescriptionCollapsedInitially(v bool)`

SetDescriptionCollapsedInitially sets DescriptionCollapsedInitially field to given value.


### GetHideComments

`func (o *FullCurrentUser) GetHideComments() bool`

GetHideComments returns the HideComments field if non-nil, zero value otherwise.

### GetHideCommentsOk

`func (o *FullCurrentUser) GetHideCommentsOk() (*bool, bool)`

GetHideCommentsOk returns a tuple with the HideComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideComments

`func (o *FullCurrentUser) SetHideComments(v bool)`

SetHideComments sets HideComments field to given value.


### GetShowHiddenComments

`func (o *FullCurrentUser) GetShowHiddenComments() bool`

GetShowHiddenComments returns the ShowHiddenComments field if non-nil, zero value otherwise.

### GetShowHiddenCommentsOk

`func (o *FullCurrentUser) GetShowHiddenCommentsOk() (*bool, bool)`

GetShowHiddenCommentsOk returns a tuple with the ShowHiddenComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHiddenComments

`func (o *FullCurrentUser) SetShowHiddenComments(v bool)`

SetShowHiddenComments sets ShowHiddenComments field to given value.


### GetShowPostStatistics

`func (o *FullCurrentUser) GetShowPostStatistics() bool`

GetShowPostStatistics returns the ShowPostStatistics field if non-nil, zero value otherwise.

### GetShowPostStatisticsOk

`func (o *FullCurrentUser) GetShowPostStatisticsOk() (*bool, bool)`

GetShowPostStatisticsOk returns a tuple with the ShowPostStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPostStatistics

`func (o *FullCurrentUser) SetShowPostStatistics(v bool)`

SetShowPostStatistics sets ShowPostStatistics field to given value.


### GetReceiveEmailNotifications

`func (o *FullCurrentUser) GetReceiveEmailNotifications() bool`

GetReceiveEmailNotifications returns the ReceiveEmailNotifications field if non-nil, zero value otherwise.

### GetReceiveEmailNotificationsOk

`func (o *FullCurrentUser) GetReceiveEmailNotificationsOk() (*bool, bool)`

GetReceiveEmailNotificationsOk returns a tuple with the ReceiveEmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveEmailNotifications

`func (o *FullCurrentUser) SetReceiveEmailNotifications(v bool)`

SetReceiveEmailNotifications sets ReceiveEmailNotifications field to given value.


### GetEnableKeyboardNavigation

`func (o *FullCurrentUser) GetEnableKeyboardNavigation() bool`

GetEnableKeyboardNavigation returns the EnableKeyboardNavigation field if non-nil, zero value otherwise.

### GetEnableKeyboardNavigationOk

`func (o *FullCurrentUser) GetEnableKeyboardNavigationOk() (*bool, bool)`

GetEnableKeyboardNavigationOk returns a tuple with the EnableKeyboardNavigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableKeyboardNavigation

`func (o *FullCurrentUser) SetEnableKeyboardNavigation(v bool)`

SetEnableKeyboardNavigation sets EnableKeyboardNavigation field to given value.


### GetEnablePrivacyMode

`func (o *FullCurrentUser) GetEnablePrivacyMode() bool`

GetEnablePrivacyMode returns the EnablePrivacyMode field if non-nil, zero value otherwise.

### GetEnablePrivacyModeOk

`func (o *FullCurrentUser) GetEnablePrivacyModeOk() (*bool, bool)`

GetEnablePrivacyModeOk returns a tuple with the EnablePrivacyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePrivacyMode

`func (o *FullCurrentUser) SetEnablePrivacyMode(v bool)`

SetEnablePrivacyMode sets EnablePrivacyMode field to given value.


### GetStyleUsernames

`func (o *FullCurrentUser) GetStyleUsernames() bool`

GetStyleUsernames returns the StyleUsernames field if non-nil, zero value otherwise.

### GetStyleUsernamesOk

`func (o *FullCurrentUser) GetStyleUsernamesOk() (*bool, bool)`

GetStyleUsernamesOk returns a tuple with the StyleUsernames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyleUsernames

`func (o *FullCurrentUser) SetStyleUsernames(v bool)`

SetStyleUsernames sets StyleUsernames field to given value.


### GetEnableAutoComplete

`func (o *FullCurrentUser) GetEnableAutoComplete() bool`

GetEnableAutoComplete returns the EnableAutoComplete field if non-nil, zero value otherwise.

### GetEnableAutoCompleteOk

`func (o *FullCurrentUser) GetEnableAutoCompleteOk() (*bool, bool)`

GetEnableAutoCompleteOk returns a tuple with the EnableAutoComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoComplete

`func (o *FullCurrentUser) SetEnableAutoComplete(v bool)`

SetEnableAutoComplete sets EnableAutoComplete field to given value.


### GetDisableCroppedThumbnails

`func (o *FullCurrentUser) GetDisableCroppedThumbnails() bool`

GetDisableCroppedThumbnails returns the DisableCroppedThumbnails field if non-nil, zero value otherwise.

### GetDisableCroppedThumbnailsOk

`func (o *FullCurrentUser) GetDisableCroppedThumbnailsOk() (*bool, bool)`

GetDisableCroppedThumbnailsOk returns a tuple with the DisableCroppedThumbnails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableCroppedThumbnails

`func (o *FullCurrentUser) SetDisableCroppedThumbnails(v bool)`

SetDisableCroppedThumbnails sets DisableCroppedThumbnails field to given value.


### GetEnableSafeMode

`func (o *FullCurrentUser) GetEnableSafeMode() bool`

GetEnableSafeMode returns the EnableSafeMode field if non-nil, zero value otherwise.

### GetEnableSafeModeOk

`func (o *FullCurrentUser) GetEnableSafeModeOk() (*bool, bool)`

GetEnableSafeModeOk returns a tuple with the EnableSafeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSafeMode

`func (o *FullCurrentUser) SetEnableSafeMode(v bool)`

SetEnableSafeMode sets EnableSafeMode field to given value.


### GetDisableResponsiveMode

`func (o *FullCurrentUser) GetDisableResponsiveMode() bool`

GetDisableResponsiveMode returns the DisableResponsiveMode field if non-nil, zero value otherwise.

### GetDisableResponsiveModeOk

`func (o *FullCurrentUser) GetDisableResponsiveModeOk() (*bool, bool)`

GetDisableResponsiveModeOk returns a tuple with the DisableResponsiveMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableResponsiveMode

`func (o *FullCurrentUser) SetDisableResponsiveMode(v bool)`

SetDisableResponsiveMode sets DisableResponsiveMode field to given value.


### GetNoFlagging

`func (o *FullCurrentUser) GetNoFlagging() bool`

GetNoFlagging returns the NoFlagging field if non-nil, zero value otherwise.

### GetNoFlaggingOk

`func (o *FullCurrentUser) GetNoFlaggingOk() (*bool, bool)`

GetNoFlaggingOk returns a tuple with the NoFlagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoFlagging

`func (o *FullCurrentUser) SetNoFlagging(v bool)`

SetNoFlagging sets NoFlagging field to given value.


### GetDisableUserDmails

`func (o *FullCurrentUser) GetDisableUserDmails() bool`

GetDisableUserDmails returns the DisableUserDmails field if non-nil, zero value otherwise.

### GetDisableUserDmailsOk

`func (o *FullCurrentUser) GetDisableUserDmailsOk() (*bool, bool)`

GetDisableUserDmailsOk returns a tuple with the DisableUserDmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableUserDmails

`func (o *FullCurrentUser) SetDisableUserDmails(v bool)`

SetDisableUserDmails sets DisableUserDmails field to given value.


### GetEnableCompactUploader

`func (o *FullCurrentUser) GetEnableCompactUploader() bool`

GetEnableCompactUploader returns the EnableCompactUploader field if non-nil, zero value otherwise.

### GetEnableCompactUploaderOk

`func (o *FullCurrentUser) GetEnableCompactUploaderOk() (*bool, bool)`

GetEnableCompactUploaderOk returns a tuple with the EnableCompactUploader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCompactUploader

`func (o *FullCurrentUser) SetEnableCompactUploader(v bool)`

SetEnableCompactUploader sets EnableCompactUploader field to given value.


### GetReplacementsBeta

`func (o *FullCurrentUser) GetReplacementsBeta() bool`

GetReplacementsBeta returns the ReplacementsBeta field if non-nil, zero value otherwise.

### GetReplacementsBetaOk

`func (o *FullCurrentUser) GetReplacementsBetaOk() (*bool, bool)`

GetReplacementsBetaOk returns a tuple with the ReplacementsBeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementsBeta

`func (o *FullCurrentUser) SetReplacementsBeta(v bool)`

SetReplacementsBeta sets ReplacementsBeta field to given value.


### GetUpdatedAt

`func (o *FullCurrentUser) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FullCurrentUser) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FullCurrentUser) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetEmail

`func (o *FullCurrentUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FullCurrentUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FullCurrentUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetLastLoggedInAt

`func (o *FullCurrentUser) GetLastLoggedInAt() time.Time`

GetLastLoggedInAt returns the LastLoggedInAt field if non-nil, zero value otherwise.

### GetLastLoggedInAtOk

`func (o *FullCurrentUser) GetLastLoggedInAtOk() (*time.Time, bool)`

GetLastLoggedInAtOk returns a tuple with the LastLoggedInAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoggedInAt

`func (o *FullCurrentUser) SetLastLoggedInAt(v time.Time)`

SetLastLoggedInAt sets LastLoggedInAt field to given value.


### GetLastForumReadAt

`func (o *FullCurrentUser) GetLastForumReadAt() time.Time`

GetLastForumReadAt returns the LastForumReadAt field if non-nil, zero value otherwise.

### GetLastForumReadAtOk

`func (o *FullCurrentUser) GetLastForumReadAtOk() (*time.Time, bool)`

GetLastForumReadAtOk returns a tuple with the LastForumReadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastForumReadAt

`func (o *FullCurrentUser) SetLastForumReadAt(v time.Time)`

SetLastForumReadAt sets LastForumReadAt field to given value.


### GetRecentTags

`func (o *FullCurrentUser) GetRecentTags() string`

GetRecentTags returns the RecentTags field if non-nil, zero value otherwise.

### GetRecentTagsOk

`func (o *FullCurrentUser) GetRecentTagsOk() (*string, bool)`

GetRecentTagsOk returns a tuple with the RecentTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentTags

`func (o *FullCurrentUser) SetRecentTags(v string)`

SetRecentTags sets RecentTags field to given value.


### GetCommentThreshold

`func (o *FullCurrentUser) GetCommentThreshold() float32`

GetCommentThreshold returns the CommentThreshold field if non-nil, zero value otherwise.

### GetCommentThresholdOk

`func (o *FullCurrentUser) GetCommentThresholdOk() (*float32, bool)`

GetCommentThresholdOk returns a tuple with the CommentThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentThreshold

`func (o *FullCurrentUser) SetCommentThreshold(v float32)`

SetCommentThreshold sets CommentThreshold field to given value.


### GetDefaultImageSizedefaultImageSize

`func (o *FullCurrentUser) GetDefaultImageSizedefaultImageSize() string`

GetDefaultImageSizedefaultImageSize returns the DefaultImageSizedefaultImageSize field if non-nil, zero value otherwise.

### GetDefaultImageSizedefaultImageSizeOk

`func (o *FullCurrentUser) GetDefaultImageSizedefaultImageSizeOk() (*string, bool)`

GetDefaultImageSizedefaultImageSizeOk returns a tuple with the DefaultImageSizedefaultImageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImageSizedefaultImageSize

`func (o *FullCurrentUser) SetDefaultImageSizedefaultImageSize(v string)`

SetDefaultImageSizedefaultImageSize sets DefaultImageSizedefaultImageSize field to given value.

### HasDefaultImageSizedefaultImageSize

`func (o *FullCurrentUser) HasDefaultImageSizedefaultImageSize() bool`

HasDefaultImageSizedefaultImageSize returns a boolean if a field has been set.

### GetFavoriteTags

`func (o *FullCurrentUser) GetFavoriteTags() string`

GetFavoriteTags returns the FavoriteTags field if non-nil, zero value otherwise.

### GetFavoriteTagsOk

`func (o *FullCurrentUser) GetFavoriteTagsOk() (*string, bool)`

GetFavoriteTagsOk returns a tuple with the FavoriteTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoriteTags

`func (o *FullCurrentUser) SetFavoriteTags(v string)`

SetFavoriteTags sets FavoriteTags field to given value.


### GetBlacklistedTags

`func (o *FullCurrentUser) GetBlacklistedTags() string`

GetBlacklistedTags returns the BlacklistedTags field if non-nil, zero value otherwise.

### GetBlacklistedTagsOk

`func (o *FullCurrentUser) GetBlacklistedTagsOk() (*string, bool)`

GetBlacklistedTagsOk returns a tuple with the BlacklistedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistedTags

`func (o *FullCurrentUser) SetBlacklistedTags(v string)`

SetBlacklistedTags sets BlacklistedTags field to given value.


### GetTimeZone

`func (o *FullCurrentUser) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *FullCurrentUser) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *FullCurrentUser) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.


### GetPerPage

`func (o *FullCurrentUser) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *FullCurrentUser) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *FullCurrentUser) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.


### GetCustomStyle

`func (o *FullCurrentUser) GetCustomStyle() string`

GetCustomStyle returns the CustomStyle field if non-nil, zero value otherwise.

### GetCustomStyleOk

`func (o *FullCurrentUser) GetCustomStyleOk() (*string, bool)`

GetCustomStyleOk returns a tuple with the CustomStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStyle

`func (o *FullCurrentUser) SetCustomStyle(v string)`

SetCustomStyle sets CustomStyle field to given value.


### GetFavoriteCount

`func (o *FullCurrentUser) GetFavoriteCount() int32`

GetFavoriteCount returns the FavoriteCount field if non-nil, zero value otherwise.

### GetFavoriteCountOk

`func (o *FullCurrentUser) GetFavoriteCountOk() (*int32, bool)`

GetFavoriteCountOk returns a tuple with the FavoriteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoriteCount

`func (o *FullCurrentUser) SetFavoriteCount(v int32)`

SetFavoriteCount sets FavoriteCount field to given value.


### GetApiRegenMultiplier

`func (o *FullCurrentUser) GetApiRegenMultiplier() float32`

GetApiRegenMultiplier returns the ApiRegenMultiplier field if non-nil, zero value otherwise.

### GetApiRegenMultiplierOk

`func (o *FullCurrentUser) GetApiRegenMultiplierOk() (*float32, bool)`

GetApiRegenMultiplierOk returns a tuple with the ApiRegenMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiRegenMultiplier

`func (o *FullCurrentUser) SetApiRegenMultiplier(v float32)`

SetApiRegenMultiplier sets ApiRegenMultiplier field to given value.


### GetApiBurstLimit

`func (o *FullCurrentUser) GetApiBurstLimit() float32`

GetApiBurstLimit returns the ApiBurstLimit field if non-nil, zero value otherwise.

### GetApiBurstLimitOk

`func (o *FullCurrentUser) GetApiBurstLimitOk() (*float32, bool)`

GetApiBurstLimitOk returns a tuple with the ApiBurstLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBurstLimit

`func (o *FullCurrentUser) SetApiBurstLimit(v float32)`

SetApiBurstLimit sets ApiBurstLimit field to given value.


### GetRemainingApiLimit

`func (o *FullCurrentUser) GetRemainingApiLimit() float32`

GetRemainingApiLimit returns the RemainingApiLimit field if non-nil, zero value otherwise.

### GetRemainingApiLimitOk

`func (o *FullCurrentUser) GetRemainingApiLimitOk() (*float32, bool)`

GetRemainingApiLimitOk returns a tuple with the RemainingApiLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingApiLimit

`func (o *FullCurrentUser) SetRemainingApiLimit(v float32)`

SetRemainingApiLimit sets RemainingApiLimit field to given value.


### GetStatementTimeout

`func (o *FullCurrentUser) GetStatementTimeout() float32`

GetStatementTimeout returns the StatementTimeout field if non-nil, zero value otherwise.

### GetStatementTimeoutOk

`func (o *FullCurrentUser) GetStatementTimeoutOk() (*float32, bool)`

GetStatementTimeoutOk returns a tuple with the StatementTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementTimeout

`func (o *FullCurrentUser) SetStatementTimeout(v float32)`

SetStatementTimeout sets StatementTimeout field to given value.


### GetFavoriteLimit

`func (o *FullCurrentUser) GetFavoriteLimit() int32`

GetFavoriteLimit returns the FavoriteLimit field if non-nil, zero value otherwise.

### GetFavoriteLimitOk

`func (o *FullCurrentUser) GetFavoriteLimitOk() (*int32, bool)`

GetFavoriteLimitOk returns a tuple with the FavoriteLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoriteLimit

`func (o *FullCurrentUser) SetFavoriteLimit(v int32)`

SetFavoriteLimit sets FavoriteLimit field to given value.


### GetTagQueryLimit

`func (o *FullCurrentUser) GetTagQueryLimit() int32`

GetTagQueryLimit returns the TagQueryLimit field if non-nil, zero value otherwise.

### GetTagQueryLimitOk

`func (o *FullCurrentUser) GetTagQueryLimitOk() (*int32, bool)`

GetTagQueryLimitOk returns a tuple with the TagQueryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagQueryLimit

`func (o *FullCurrentUser) SetTagQueryLimit(v int32)`

SetTagQueryLimit sets TagQueryLimit field to given value.


### GetHasMail

`func (o *FullCurrentUser) GetHasMail() bool`

GetHasMail returns the HasMail field if non-nil, zero value otherwise.

### GetHasMailOk

`func (o *FullCurrentUser) GetHasMailOk() (*bool, bool)`

GetHasMailOk returns a tuple with the HasMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMail

`func (o *FullCurrentUser) SetHasMail(v bool)`

SetHasMail sets HasMail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


