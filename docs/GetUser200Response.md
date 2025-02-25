# GetUser200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**CreatedAt** | **time.Time** |  | 
**Name** | **string** |  | 
**Level** | **int32** |  | 
**BaseUploadLimit** | **int32** |  | 
**PostUploadCount** | **int32** |  | 
**PostUpdateCount** | **int32** |  | 
**NoteUpdateCount** | **int32** |  | 
**IsBanned** | **bool** |  | 
**CanApprovePosts** | **bool** |  | 
**CanUploadFree** | **bool** |  | 
**LevelString** | **string** |  | 
**AvatarId** | **int32** |  | 
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
**LastLoggedInAt** | **string** |  | 
**LastForumReadAt** | **string** |  | 
**RecentTags** | **string** |  | 
**CommentThreshold** | **int32** |  | 
**DefaultImageSize** | **string** |  | 
**FavoriteTags** | **string** |  | 
**BlacklistedTags** | **string** |  | 
**TimeZone** | **string** |  | 
**PerPage** | **int32** |  | 
**CustomStyle** | **string** |  | 
**ApiRegenMultiplier** | **int32** |  | 
**ApiBurstLimit** | **int32** |  | 
**RemainingApiLimit** | **int32** |  | 
**StatementTimeout** | **int32** |  | 
**FavoriteLimit** | **int32** |  | 
**TagQueryLimit** | **int32** |  | 
**HasMail** | **bool** |  | 

## Methods

### NewGetUser200Response

`func NewGetUser200Response(id int32, createdAt time.Time, name string, level int32, baseUploadLimit int32, postUploadCount int32, postUpdateCount int32, noteUpdateCount int32, isBanned bool, canApprovePosts bool, canUploadFree bool, levelString string, avatarId int32, wikiPageVersionCount int32, artistVersionCount int32, poolVersionCount int32, forumPostCount int32, commentCount int32, flagCount int32, favoriteCount int32, positiveFeedbackCount int32, neutralFeedbackCount int32, negativeFeedbackCount int32, uploadLimit int32, profileAbout string, profileArtinfo string, blacklistUsers bool, descriptionCollapsedInitially bool, hideComments bool, showHiddenComments bool, showPostStatistics bool, receiveEmailNotifications bool, enableKeyboardNavigation bool, enablePrivacyMode bool, styleUsernames bool, enableAutoComplete bool, disableCroppedThumbnails bool, enableSafeMode bool, disableResponsiveMode bool, noFlagging bool, disableUserDmails bool, enableCompactUploader bool, replacementsBeta bool, updatedAt time.Time, email string, lastLoggedInAt string, lastForumReadAt string, recentTags string, commentThreshold int32, defaultImageSize string, favoriteTags string, blacklistedTags string, timeZone string, perPage int32, customStyle string, apiRegenMultiplier int32, apiBurstLimit int32, remainingApiLimit int32, statementTimeout int32, favoriteLimit int32, tagQueryLimit int32, hasMail bool, ) *GetUser200Response`

NewGetUser200Response instantiates a new GetUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUser200ResponseWithDefaults

`func NewGetUser200ResponseWithDefaults() *GetUser200Response`

NewGetUser200ResponseWithDefaults instantiates a new GetUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetUser200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetUser200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetUser200Response) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *GetUser200Response) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetUser200Response) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetUser200Response) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetName

`func (o *GetUser200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetUser200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetUser200Response) SetName(v string)`

SetName sets Name field to given value.


### GetLevel

`func (o *GetUser200Response) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *GetUser200Response) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *GetUser200Response) SetLevel(v int32)`

SetLevel sets Level field to given value.


### GetBaseUploadLimit

`func (o *GetUser200Response) GetBaseUploadLimit() int32`

GetBaseUploadLimit returns the BaseUploadLimit field if non-nil, zero value otherwise.

### GetBaseUploadLimitOk

`func (o *GetUser200Response) GetBaseUploadLimitOk() (*int32, bool)`

GetBaseUploadLimitOk returns a tuple with the BaseUploadLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUploadLimit

`func (o *GetUser200Response) SetBaseUploadLimit(v int32)`

SetBaseUploadLimit sets BaseUploadLimit field to given value.


### GetPostUploadCount

`func (o *GetUser200Response) GetPostUploadCount() int32`

GetPostUploadCount returns the PostUploadCount field if non-nil, zero value otherwise.

### GetPostUploadCountOk

`func (o *GetUser200Response) GetPostUploadCountOk() (*int32, bool)`

GetPostUploadCountOk returns a tuple with the PostUploadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostUploadCount

`func (o *GetUser200Response) SetPostUploadCount(v int32)`

SetPostUploadCount sets PostUploadCount field to given value.


### GetPostUpdateCount

`func (o *GetUser200Response) GetPostUpdateCount() int32`

GetPostUpdateCount returns the PostUpdateCount field if non-nil, zero value otherwise.

### GetPostUpdateCountOk

`func (o *GetUser200Response) GetPostUpdateCountOk() (*int32, bool)`

GetPostUpdateCountOk returns a tuple with the PostUpdateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostUpdateCount

`func (o *GetUser200Response) SetPostUpdateCount(v int32)`

SetPostUpdateCount sets PostUpdateCount field to given value.


### GetNoteUpdateCount

`func (o *GetUser200Response) GetNoteUpdateCount() int32`

GetNoteUpdateCount returns the NoteUpdateCount field if non-nil, zero value otherwise.

### GetNoteUpdateCountOk

`func (o *GetUser200Response) GetNoteUpdateCountOk() (*int32, bool)`

GetNoteUpdateCountOk returns a tuple with the NoteUpdateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteUpdateCount

`func (o *GetUser200Response) SetNoteUpdateCount(v int32)`

SetNoteUpdateCount sets NoteUpdateCount field to given value.


### GetIsBanned

`func (o *GetUser200Response) GetIsBanned() bool`

GetIsBanned returns the IsBanned field if non-nil, zero value otherwise.

### GetIsBannedOk

`func (o *GetUser200Response) GetIsBannedOk() (*bool, bool)`

GetIsBannedOk returns a tuple with the IsBanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBanned

`func (o *GetUser200Response) SetIsBanned(v bool)`

SetIsBanned sets IsBanned field to given value.


### GetCanApprovePosts

`func (o *GetUser200Response) GetCanApprovePosts() bool`

GetCanApprovePosts returns the CanApprovePosts field if non-nil, zero value otherwise.

### GetCanApprovePostsOk

`func (o *GetUser200Response) GetCanApprovePostsOk() (*bool, bool)`

GetCanApprovePostsOk returns a tuple with the CanApprovePosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanApprovePosts

`func (o *GetUser200Response) SetCanApprovePosts(v bool)`

SetCanApprovePosts sets CanApprovePosts field to given value.


### GetCanUploadFree

`func (o *GetUser200Response) GetCanUploadFree() bool`

GetCanUploadFree returns the CanUploadFree field if non-nil, zero value otherwise.

### GetCanUploadFreeOk

`func (o *GetUser200Response) GetCanUploadFreeOk() (*bool, bool)`

GetCanUploadFreeOk returns a tuple with the CanUploadFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUploadFree

`func (o *GetUser200Response) SetCanUploadFree(v bool)`

SetCanUploadFree sets CanUploadFree field to given value.


### GetLevelString

`func (o *GetUser200Response) GetLevelString() string`

GetLevelString returns the LevelString field if non-nil, zero value otherwise.

### GetLevelStringOk

`func (o *GetUser200Response) GetLevelStringOk() (*string, bool)`

GetLevelStringOk returns a tuple with the LevelString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelString

`func (o *GetUser200Response) SetLevelString(v string)`

SetLevelString sets LevelString field to given value.


### GetAvatarId

`func (o *GetUser200Response) GetAvatarId() int32`

GetAvatarId returns the AvatarId field if non-nil, zero value otherwise.

### GetAvatarIdOk

`func (o *GetUser200Response) GetAvatarIdOk() (*int32, bool)`

GetAvatarIdOk returns a tuple with the AvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarId

`func (o *GetUser200Response) SetAvatarId(v int32)`

SetAvatarId sets AvatarId field to given value.


### GetWikiPageVersionCount

`func (o *GetUser200Response) GetWikiPageVersionCount() int32`

GetWikiPageVersionCount returns the WikiPageVersionCount field if non-nil, zero value otherwise.

### GetWikiPageVersionCountOk

`func (o *GetUser200Response) GetWikiPageVersionCountOk() (*int32, bool)`

GetWikiPageVersionCountOk returns a tuple with the WikiPageVersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWikiPageVersionCount

`func (o *GetUser200Response) SetWikiPageVersionCount(v int32)`

SetWikiPageVersionCount sets WikiPageVersionCount field to given value.


### GetArtistVersionCount

`func (o *GetUser200Response) GetArtistVersionCount() int32`

GetArtistVersionCount returns the ArtistVersionCount field if non-nil, zero value otherwise.

### GetArtistVersionCountOk

`func (o *GetUser200Response) GetArtistVersionCountOk() (*int32, bool)`

GetArtistVersionCountOk returns a tuple with the ArtistVersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtistVersionCount

`func (o *GetUser200Response) SetArtistVersionCount(v int32)`

SetArtistVersionCount sets ArtistVersionCount field to given value.


### GetPoolVersionCount

`func (o *GetUser200Response) GetPoolVersionCount() int32`

GetPoolVersionCount returns the PoolVersionCount field if non-nil, zero value otherwise.

### GetPoolVersionCountOk

`func (o *GetUser200Response) GetPoolVersionCountOk() (*int32, bool)`

GetPoolVersionCountOk returns a tuple with the PoolVersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolVersionCount

`func (o *GetUser200Response) SetPoolVersionCount(v int32)`

SetPoolVersionCount sets PoolVersionCount field to given value.


### GetForumPostCount

`func (o *GetUser200Response) GetForumPostCount() int32`

GetForumPostCount returns the ForumPostCount field if non-nil, zero value otherwise.

### GetForumPostCountOk

`func (o *GetUser200Response) GetForumPostCountOk() (*int32, bool)`

GetForumPostCountOk returns a tuple with the ForumPostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumPostCount

`func (o *GetUser200Response) SetForumPostCount(v int32)`

SetForumPostCount sets ForumPostCount field to given value.


### GetCommentCount

`func (o *GetUser200Response) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *GetUser200Response) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *GetUser200Response) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.


### GetFlagCount

`func (o *GetUser200Response) GetFlagCount() int32`

GetFlagCount returns the FlagCount field if non-nil, zero value otherwise.

### GetFlagCountOk

`func (o *GetUser200Response) GetFlagCountOk() (*int32, bool)`

GetFlagCountOk returns a tuple with the FlagCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagCount

`func (o *GetUser200Response) SetFlagCount(v int32)`

SetFlagCount sets FlagCount field to given value.


### GetFavoriteCount

`func (o *GetUser200Response) GetFavoriteCount() int32`

GetFavoriteCount returns the FavoriteCount field if non-nil, zero value otherwise.

### GetFavoriteCountOk

`func (o *GetUser200Response) GetFavoriteCountOk() (*int32, bool)`

GetFavoriteCountOk returns a tuple with the FavoriteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoriteCount

`func (o *GetUser200Response) SetFavoriteCount(v int32)`

SetFavoriteCount sets FavoriteCount field to given value.


### GetPositiveFeedbackCount

`func (o *GetUser200Response) GetPositiveFeedbackCount() int32`

GetPositiveFeedbackCount returns the PositiveFeedbackCount field if non-nil, zero value otherwise.

### GetPositiveFeedbackCountOk

`func (o *GetUser200Response) GetPositiveFeedbackCountOk() (*int32, bool)`

GetPositiveFeedbackCountOk returns a tuple with the PositiveFeedbackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositiveFeedbackCount

`func (o *GetUser200Response) SetPositiveFeedbackCount(v int32)`

SetPositiveFeedbackCount sets PositiveFeedbackCount field to given value.


### GetNeutralFeedbackCount

`func (o *GetUser200Response) GetNeutralFeedbackCount() int32`

GetNeutralFeedbackCount returns the NeutralFeedbackCount field if non-nil, zero value otherwise.

### GetNeutralFeedbackCountOk

`func (o *GetUser200Response) GetNeutralFeedbackCountOk() (*int32, bool)`

GetNeutralFeedbackCountOk returns a tuple with the NeutralFeedbackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeutralFeedbackCount

`func (o *GetUser200Response) SetNeutralFeedbackCount(v int32)`

SetNeutralFeedbackCount sets NeutralFeedbackCount field to given value.


### GetNegativeFeedbackCount

`func (o *GetUser200Response) GetNegativeFeedbackCount() int32`

GetNegativeFeedbackCount returns the NegativeFeedbackCount field if non-nil, zero value otherwise.

### GetNegativeFeedbackCountOk

`func (o *GetUser200Response) GetNegativeFeedbackCountOk() (*int32, bool)`

GetNegativeFeedbackCountOk returns a tuple with the NegativeFeedbackCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativeFeedbackCount

`func (o *GetUser200Response) SetNegativeFeedbackCount(v int32)`

SetNegativeFeedbackCount sets NegativeFeedbackCount field to given value.


### GetUploadLimit

`func (o *GetUser200Response) GetUploadLimit() int32`

GetUploadLimit returns the UploadLimit field if non-nil, zero value otherwise.

### GetUploadLimitOk

`func (o *GetUser200Response) GetUploadLimitOk() (*int32, bool)`

GetUploadLimitOk returns a tuple with the UploadLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadLimit

`func (o *GetUser200Response) SetUploadLimit(v int32)`

SetUploadLimit sets UploadLimit field to given value.


### GetProfileAbout

`func (o *GetUser200Response) GetProfileAbout() string`

GetProfileAbout returns the ProfileAbout field if non-nil, zero value otherwise.

### GetProfileAboutOk

`func (o *GetUser200Response) GetProfileAboutOk() (*string, bool)`

GetProfileAboutOk returns a tuple with the ProfileAbout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileAbout

`func (o *GetUser200Response) SetProfileAbout(v string)`

SetProfileAbout sets ProfileAbout field to given value.


### GetProfileArtinfo

`func (o *GetUser200Response) GetProfileArtinfo() string`

GetProfileArtinfo returns the ProfileArtinfo field if non-nil, zero value otherwise.

### GetProfileArtinfoOk

`func (o *GetUser200Response) GetProfileArtinfoOk() (*string, bool)`

GetProfileArtinfoOk returns a tuple with the ProfileArtinfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileArtinfo

`func (o *GetUser200Response) SetProfileArtinfo(v string)`

SetProfileArtinfo sets ProfileArtinfo field to given value.


### GetBlacklistUsers

`func (o *GetUser200Response) GetBlacklistUsers() bool`

GetBlacklistUsers returns the BlacklistUsers field if non-nil, zero value otherwise.

### GetBlacklistUsersOk

`func (o *GetUser200Response) GetBlacklistUsersOk() (*bool, bool)`

GetBlacklistUsersOk returns a tuple with the BlacklistUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistUsers

`func (o *GetUser200Response) SetBlacklistUsers(v bool)`

SetBlacklistUsers sets BlacklistUsers field to given value.


### GetDescriptionCollapsedInitially

`func (o *GetUser200Response) GetDescriptionCollapsedInitially() bool`

GetDescriptionCollapsedInitially returns the DescriptionCollapsedInitially field if non-nil, zero value otherwise.

### GetDescriptionCollapsedInitiallyOk

`func (o *GetUser200Response) GetDescriptionCollapsedInitiallyOk() (*bool, bool)`

GetDescriptionCollapsedInitiallyOk returns a tuple with the DescriptionCollapsedInitially field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionCollapsedInitially

`func (o *GetUser200Response) SetDescriptionCollapsedInitially(v bool)`

SetDescriptionCollapsedInitially sets DescriptionCollapsedInitially field to given value.


### GetHideComments

`func (o *GetUser200Response) GetHideComments() bool`

GetHideComments returns the HideComments field if non-nil, zero value otherwise.

### GetHideCommentsOk

`func (o *GetUser200Response) GetHideCommentsOk() (*bool, bool)`

GetHideCommentsOk returns a tuple with the HideComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideComments

`func (o *GetUser200Response) SetHideComments(v bool)`

SetHideComments sets HideComments field to given value.


### GetShowHiddenComments

`func (o *GetUser200Response) GetShowHiddenComments() bool`

GetShowHiddenComments returns the ShowHiddenComments field if non-nil, zero value otherwise.

### GetShowHiddenCommentsOk

`func (o *GetUser200Response) GetShowHiddenCommentsOk() (*bool, bool)`

GetShowHiddenCommentsOk returns a tuple with the ShowHiddenComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHiddenComments

`func (o *GetUser200Response) SetShowHiddenComments(v bool)`

SetShowHiddenComments sets ShowHiddenComments field to given value.


### GetShowPostStatistics

`func (o *GetUser200Response) GetShowPostStatistics() bool`

GetShowPostStatistics returns the ShowPostStatistics field if non-nil, zero value otherwise.

### GetShowPostStatisticsOk

`func (o *GetUser200Response) GetShowPostStatisticsOk() (*bool, bool)`

GetShowPostStatisticsOk returns a tuple with the ShowPostStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPostStatistics

`func (o *GetUser200Response) SetShowPostStatistics(v bool)`

SetShowPostStatistics sets ShowPostStatistics field to given value.


### GetReceiveEmailNotifications

`func (o *GetUser200Response) GetReceiveEmailNotifications() bool`

GetReceiveEmailNotifications returns the ReceiveEmailNotifications field if non-nil, zero value otherwise.

### GetReceiveEmailNotificationsOk

`func (o *GetUser200Response) GetReceiveEmailNotificationsOk() (*bool, bool)`

GetReceiveEmailNotificationsOk returns a tuple with the ReceiveEmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveEmailNotifications

`func (o *GetUser200Response) SetReceiveEmailNotifications(v bool)`

SetReceiveEmailNotifications sets ReceiveEmailNotifications field to given value.


### GetEnableKeyboardNavigation

`func (o *GetUser200Response) GetEnableKeyboardNavigation() bool`

GetEnableKeyboardNavigation returns the EnableKeyboardNavigation field if non-nil, zero value otherwise.

### GetEnableKeyboardNavigationOk

`func (o *GetUser200Response) GetEnableKeyboardNavigationOk() (*bool, bool)`

GetEnableKeyboardNavigationOk returns a tuple with the EnableKeyboardNavigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableKeyboardNavigation

`func (o *GetUser200Response) SetEnableKeyboardNavigation(v bool)`

SetEnableKeyboardNavigation sets EnableKeyboardNavigation field to given value.


### GetEnablePrivacyMode

`func (o *GetUser200Response) GetEnablePrivacyMode() bool`

GetEnablePrivacyMode returns the EnablePrivacyMode field if non-nil, zero value otherwise.

### GetEnablePrivacyModeOk

`func (o *GetUser200Response) GetEnablePrivacyModeOk() (*bool, bool)`

GetEnablePrivacyModeOk returns a tuple with the EnablePrivacyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePrivacyMode

`func (o *GetUser200Response) SetEnablePrivacyMode(v bool)`

SetEnablePrivacyMode sets EnablePrivacyMode field to given value.


### GetStyleUsernames

`func (o *GetUser200Response) GetStyleUsernames() bool`

GetStyleUsernames returns the StyleUsernames field if non-nil, zero value otherwise.

### GetStyleUsernamesOk

`func (o *GetUser200Response) GetStyleUsernamesOk() (*bool, bool)`

GetStyleUsernamesOk returns a tuple with the StyleUsernames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyleUsernames

`func (o *GetUser200Response) SetStyleUsernames(v bool)`

SetStyleUsernames sets StyleUsernames field to given value.


### GetEnableAutoComplete

`func (o *GetUser200Response) GetEnableAutoComplete() bool`

GetEnableAutoComplete returns the EnableAutoComplete field if non-nil, zero value otherwise.

### GetEnableAutoCompleteOk

`func (o *GetUser200Response) GetEnableAutoCompleteOk() (*bool, bool)`

GetEnableAutoCompleteOk returns a tuple with the EnableAutoComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoComplete

`func (o *GetUser200Response) SetEnableAutoComplete(v bool)`

SetEnableAutoComplete sets EnableAutoComplete field to given value.


### GetDisableCroppedThumbnails

`func (o *GetUser200Response) GetDisableCroppedThumbnails() bool`

GetDisableCroppedThumbnails returns the DisableCroppedThumbnails field if non-nil, zero value otherwise.

### GetDisableCroppedThumbnailsOk

`func (o *GetUser200Response) GetDisableCroppedThumbnailsOk() (*bool, bool)`

GetDisableCroppedThumbnailsOk returns a tuple with the DisableCroppedThumbnails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableCroppedThumbnails

`func (o *GetUser200Response) SetDisableCroppedThumbnails(v bool)`

SetDisableCroppedThumbnails sets DisableCroppedThumbnails field to given value.


### GetEnableSafeMode

`func (o *GetUser200Response) GetEnableSafeMode() bool`

GetEnableSafeMode returns the EnableSafeMode field if non-nil, zero value otherwise.

### GetEnableSafeModeOk

`func (o *GetUser200Response) GetEnableSafeModeOk() (*bool, bool)`

GetEnableSafeModeOk returns a tuple with the EnableSafeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSafeMode

`func (o *GetUser200Response) SetEnableSafeMode(v bool)`

SetEnableSafeMode sets EnableSafeMode field to given value.


### GetDisableResponsiveMode

`func (o *GetUser200Response) GetDisableResponsiveMode() bool`

GetDisableResponsiveMode returns the DisableResponsiveMode field if non-nil, zero value otherwise.

### GetDisableResponsiveModeOk

`func (o *GetUser200Response) GetDisableResponsiveModeOk() (*bool, bool)`

GetDisableResponsiveModeOk returns a tuple with the DisableResponsiveMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableResponsiveMode

`func (o *GetUser200Response) SetDisableResponsiveMode(v bool)`

SetDisableResponsiveMode sets DisableResponsiveMode field to given value.


### GetNoFlagging

`func (o *GetUser200Response) GetNoFlagging() bool`

GetNoFlagging returns the NoFlagging field if non-nil, zero value otherwise.

### GetNoFlaggingOk

`func (o *GetUser200Response) GetNoFlaggingOk() (*bool, bool)`

GetNoFlaggingOk returns a tuple with the NoFlagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoFlagging

`func (o *GetUser200Response) SetNoFlagging(v bool)`

SetNoFlagging sets NoFlagging field to given value.


### GetDisableUserDmails

`func (o *GetUser200Response) GetDisableUserDmails() bool`

GetDisableUserDmails returns the DisableUserDmails field if non-nil, zero value otherwise.

### GetDisableUserDmailsOk

`func (o *GetUser200Response) GetDisableUserDmailsOk() (*bool, bool)`

GetDisableUserDmailsOk returns a tuple with the DisableUserDmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableUserDmails

`func (o *GetUser200Response) SetDisableUserDmails(v bool)`

SetDisableUserDmails sets DisableUserDmails field to given value.


### GetEnableCompactUploader

`func (o *GetUser200Response) GetEnableCompactUploader() bool`

GetEnableCompactUploader returns the EnableCompactUploader field if non-nil, zero value otherwise.

### GetEnableCompactUploaderOk

`func (o *GetUser200Response) GetEnableCompactUploaderOk() (*bool, bool)`

GetEnableCompactUploaderOk returns a tuple with the EnableCompactUploader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCompactUploader

`func (o *GetUser200Response) SetEnableCompactUploader(v bool)`

SetEnableCompactUploader sets EnableCompactUploader field to given value.


### GetReplacementsBeta

`func (o *GetUser200Response) GetReplacementsBeta() bool`

GetReplacementsBeta returns the ReplacementsBeta field if non-nil, zero value otherwise.

### GetReplacementsBetaOk

`func (o *GetUser200Response) GetReplacementsBetaOk() (*bool, bool)`

GetReplacementsBetaOk returns a tuple with the ReplacementsBeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementsBeta

`func (o *GetUser200Response) SetReplacementsBeta(v bool)`

SetReplacementsBeta sets ReplacementsBeta field to given value.


### GetUpdatedAt

`func (o *GetUser200Response) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetUser200Response) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetUser200Response) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetEmail

`func (o *GetUser200Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetUser200Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetUser200Response) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetLastLoggedInAt

`func (o *GetUser200Response) GetLastLoggedInAt() string`

GetLastLoggedInAt returns the LastLoggedInAt field if non-nil, zero value otherwise.

### GetLastLoggedInAtOk

`func (o *GetUser200Response) GetLastLoggedInAtOk() (*string, bool)`

GetLastLoggedInAtOk returns a tuple with the LastLoggedInAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoggedInAt

`func (o *GetUser200Response) SetLastLoggedInAt(v string)`

SetLastLoggedInAt sets LastLoggedInAt field to given value.


### GetLastForumReadAt

`func (o *GetUser200Response) GetLastForumReadAt() string`

GetLastForumReadAt returns the LastForumReadAt field if non-nil, zero value otherwise.

### GetLastForumReadAtOk

`func (o *GetUser200Response) GetLastForumReadAtOk() (*string, bool)`

GetLastForumReadAtOk returns a tuple with the LastForumReadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastForumReadAt

`func (o *GetUser200Response) SetLastForumReadAt(v string)`

SetLastForumReadAt sets LastForumReadAt field to given value.


### GetRecentTags

`func (o *GetUser200Response) GetRecentTags() string`

GetRecentTags returns the RecentTags field if non-nil, zero value otherwise.

### GetRecentTagsOk

`func (o *GetUser200Response) GetRecentTagsOk() (*string, bool)`

GetRecentTagsOk returns a tuple with the RecentTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentTags

`func (o *GetUser200Response) SetRecentTags(v string)`

SetRecentTags sets RecentTags field to given value.


### GetCommentThreshold

`func (o *GetUser200Response) GetCommentThreshold() int32`

GetCommentThreshold returns the CommentThreshold field if non-nil, zero value otherwise.

### GetCommentThresholdOk

`func (o *GetUser200Response) GetCommentThresholdOk() (*int32, bool)`

GetCommentThresholdOk returns a tuple with the CommentThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentThreshold

`func (o *GetUser200Response) SetCommentThreshold(v int32)`

SetCommentThreshold sets CommentThreshold field to given value.


### GetDefaultImageSize

`func (o *GetUser200Response) GetDefaultImageSize() string`

GetDefaultImageSize returns the DefaultImageSize field if non-nil, zero value otherwise.

### GetDefaultImageSizeOk

`func (o *GetUser200Response) GetDefaultImageSizeOk() (*string, bool)`

GetDefaultImageSizeOk returns a tuple with the DefaultImageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImageSize

`func (o *GetUser200Response) SetDefaultImageSize(v string)`

SetDefaultImageSize sets DefaultImageSize field to given value.


### GetFavoriteTags

`func (o *GetUser200Response) GetFavoriteTags() string`

GetFavoriteTags returns the FavoriteTags field if non-nil, zero value otherwise.

### GetFavoriteTagsOk

`func (o *GetUser200Response) GetFavoriteTagsOk() (*string, bool)`

GetFavoriteTagsOk returns a tuple with the FavoriteTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoriteTags

`func (o *GetUser200Response) SetFavoriteTags(v string)`

SetFavoriteTags sets FavoriteTags field to given value.


### GetBlacklistedTags

`func (o *GetUser200Response) GetBlacklistedTags() string`

GetBlacklistedTags returns the BlacklistedTags field if non-nil, zero value otherwise.

### GetBlacklistedTagsOk

`func (o *GetUser200Response) GetBlacklistedTagsOk() (*string, bool)`

GetBlacklistedTagsOk returns a tuple with the BlacklistedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistedTags

`func (o *GetUser200Response) SetBlacklistedTags(v string)`

SetBlacklistedTags sets BlacklistedTags field to given value.


### GetTimeZone

`func (o *GetUser200Response) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *GetUser200Response) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *GetUser200Response) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.


### GetPerPage

`func (o *GetUser200Response) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *GetUser200Response) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *GetUser200Response) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.


### GetCustomStyle

`func (o *GetUser200Response) GetCustomStyle() string`

GetCustomStyle returns the CustomStyle field if non-nil, zero value otherwise.

### GetCustomStyleOk

`func (o *GetUser200Response) GetCustomStyleOk() (*string, bool)`

GetCustomStyleOk returns a tuple with the CustomStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStyle

`func (o *GetUser200Response) SetCustomStyle(v string)`

SetCustomStyle sets CustomStyle field to given value.


### GetApiRegenMultiplier

`func (o *GetUser200Response) GetApiRegenMultiplier() int32`

GetApiRegenMultiplier returns the ApiRegenMultiplier field if non-nil, zero value otherwise.

### GetApiRegenMultiplierOk

`func (o *GetUser200Response) GetApiRegenMultiplierOk() (*int32, bool)`

GetApiRegenMultiplierOk returns a tuple with the ApiRegenMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiRegenMultiplier

`func (o *GetUser200Response) SetApiRegenMultiplier(v int32)`

SetApiRegenMultiplier sets ApiRegenMultiplier field to given value.


### GetApiBurstLimit

`func (o *GetUser200Response) GetApiBurstLimit() int32`

GetApiBurstLimit returns the ApiBurstLimit field if non-nil, zero value otherwise.

### GetApiBurstLimitOk

`func (o *GetUser200Response) GetApiBurstLimitOk() (*int32, bool)`

GetApiBurstLimitOk returns a tuple with the ApiBurstLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBurstLimit

`func (o *GetUser200Response) SetApiBurstLimit(v int32)`

SetApiBurstLimit sets ApiBurstLimit field to given value.


### GetRemainingApiLimit

`func (o *GetUser200Response) GetRemainingApiLimit() int32`

GetRemainingApiLimit returns the RemainingApiLimit field if non-nil, zero value otherwise.

### GetRemainingApiLimitOk

`func (o *GetUser200Response) GetRemainingApiLimitOk() (*int32, bool)`

GetRemainingApiLimitOk returns a tuple with the RemainingApiLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingApiLimit

`func (o *GetUser200Response) SetRemainingApiLimit(v int32)`

SetRemainingApiLimit sets RemainingApiLimit field to given value.


### GetStatementTimeout

`func (o *GetUser200Response) GetStatementTimeout() int32`

GetStatementTimeout returns the StatementTimeout field if non-nil, zero value otherwise.

### GetStatementTimeoutOk

`func (o *GetUser200Response) GetStatementTimeoutOk() (*int32, bool)`

GetStatementTimeoutOk returns a tuple with the StatementTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementTimeout

`func (o *GetUser200Response) SetStatementTimeout(v int32)`

SetStatementTimeout sets StatementTimeout field to given value.


### GetFavoriteLimit

`func (o *GetUser200Response) GetFavoriteLimit() int32`

GetFavoriteLimit returns the FavoriteLimit field if non-nil, zero value otherwise.

### GetFavoriteLimitOk

`func (o *GetUser200Response) GetFavoriteLimitOk() (*int32, bool)`

GetFavoriteLimitOk returns a tuple with the FavoriteLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoriteLimit

`func (o *GetUser200Response) SetFavoriteLimit(v int32)`

SetFavoriteLimit sets FavoriteLimit field to given value.


### GetTagQueryLimit

`func (o *GetUser200Response) GetTagQueryLimit() int32`

GetTagQueryLimit returns the TagQueryLimit field if non-nil, zero value otherwise.

### GetTagQueryLimitOk

`func (o *GetUser200Response) GetTagQueryLimitOk() (*int32, bool)`

GetTagQueryLimitOk returns a tuple with the TagQueryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagQueryLimit

`func (o *GetUser200Response) SetTagQueryLimit(v int32)`

SetTagQueryLimit sets TagQueryLimit field to given value.


### GetHasMail

`func (o *GetUser200Response) GetHasMail() bool`

GetHasMail returns the HasMail field if non-nil, zero value otherwise.

### GetHasMailOk

`func (o *GetUser200Response) GetHasMailOk() (*bool, bool)`

GetHasMailOk returns a tuple with the HasMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMail

`func (o *GetUser200Response) SetHasMail(v bool)`

SetHasMail sets HasMail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


