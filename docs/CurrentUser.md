# CurrentUser

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
**FavoriteCount** | **int32** |  | 
**ApiRegenMultiplier** | **int32** |  | 
**ApiBurstLimit** | **int32** |  | 
**RemainingApiLimit** | **int32** |  | 
**StatementTimeout** | **int32** |  | 
**FavoriteLimit** | **int32** |  | 
**TagQueryLimit** | **int32** |  | 
**HasMail** | **bool** |  | 

## Methods

### NewCurrentUser

`func NewCurrentUser(id int32, createdAt time.Time, name string, level int32, baseUploadLimit int32, postUploadCount int32, postUpdateCount int32, noteUpdateCount int32, isBanned bool, canApprovePosts bool, canUploadFree bool, levelString string, avatarId int32, blacklistUsers bool, descriptionCollapsedInitially bool, hideComments bool, showHiddenComments bool, showPostStatistics bool, receiveEmailNotifications bool, enableKeyboardNavigation bool, enablePrivacyMode bool, styleUsernames bool, enableAutoComplete bool, disableCroppedThumbnails bool, enableSafeMode bool, disableResponsiveMode bool, noFlagging bool, disableUserDmails bool, enableCompactUploader bool, replacementsBeta bool, updatedAt time.Time, email string, lastLoggedInAt string, lastForumReadAt string, recentTags string, commentThreshold int32, defaultImageSize string, favoriteTags string, blacklistedTags string, timeZone string, perPage int32, customStyle string, favoriteCount int32, apiRegenMultiplier int32, apiBurstLimit int32, remainingApiLimit int32, statementTimeout int32, favoriteLimit int32, tagQueryLimit int32, hasMail bool, ) *CurrentUser`

NewCurrentUser instantiates a new CurrentUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentUserWithDefaults

`func NewCurrentUserWithDefaults() *CurrentUser`

NewCurrentUserWithDefaults instantiates a new CurrentUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CurrentUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CurrentUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CurrentUser) SetId(v int32)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *CurrentUser) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CurrentUser) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CurrentUser) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetName

`func (o *CurrentUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CurrentUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CurrentUser) SetName(v string)`

SetName sets Name field to given value.


### GetLevel

`func (o *CurrentUser) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CurrentUser) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CurrentUser) SetLevel(v int32)`

SetLevel sets Level field to given value.


### GetBaseUploadLimit

`func (o *CurrentUser) GetBaseUploadLimit() int32`

GetBaseUploadLimit returns the BaseUploadLimit field if non-nil, zero value otherwise.

### GetBaseUploadLimitOk

`func (o *CurrentUser) GetBaseUploadLimitOk() (*int32, bool)`

GetBaseUploadLimitOk returns a tuple with the BaseUploadLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUploadLimit

`func (o *CurrentUser) SetBaseUploadLimit(v int32)`

SetBaseUploadLimit sets BaseUploadLimit field to given value.


### GetPostUploadCount

`func (o *CurrentUser) GetPostUploadCount() int32`

GetPostUploadCount returns the PostUploadCount field if non-nil, zero value otherwise.

### GetPostUploadCountOk

`func (o *CurrentUser) GetPostUploadCountOk() (*int32, bool)`

GetPostUploadCountOk returns a tuple with the PostUploadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostUploadCount

`func (o *CurrentUser) SetPostUploadCount(v int32)`

SetPostUploadCount sets PostUploadCount field to given value.


### GetPostUpdateCount

`func (o *CurrentUser) GetPostUpdateCount() int32`

GetPostUpdateCount returns the PostUpdateCount field if non-nil, zero value otherwise.

### GetPostUpdateCountOk

`func (o *CurrentUser) GetPostUpdateCountOk() (*int32, bool)`

GetPostUpdateCountOk returns a tuple with the PostUpdateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostUpdateCount

`func (o *CurrentUser) SetPostUpdateCount(v int32)`

SetPostUpdateCount sets PostUpdateCount field to given value.


### GetNoteUpdateCount

`func (o *CurrentUser) GetNoteUpdateCount() int32`

GetNoteUpdateCount returns the NoteUpdateCount field if non-nil, zero value otherwise.

### GetNoteUpdateCountOk

`func (o *CurrentUser) GetNoteUpdateCountOk() (*int32, bool)`

GetNoteUpdateCountOk returns a tuple with the NoteUpdateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoteUpdateCount

`func (o *CurrentUser) SetNoteUpdateCount(v int32)`

SetNoteUpdateCount sets NoteUpdateCount field to given value.


### GetIsBanned

`func (o *CurrentUser) GetIsBanned() bool`

GetIsBanned returns the IsBanned field if non-nil, zero value otherwise.

### GetIsBannedOk

`func (o *CurrentUser) GetIsBannedOk() (*bool, bool)`

GetIsBannedOk returns a tuple with the IsBanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBanned

`func (o *CurrentUser) SetIsBanned(v bool)`

SetIsBanned sets IsBanned field to given value.


### GetCanApprovePosts

`func (o *CurrentUser) GetCanApprovePosts() bool`

GetCanApprovePosts returns the CanApprovePosts field if non-nil, zero value otherwise.

### GetCanApprovePostsOk

`func (o *CurrentUser) GetCanApprovePostsOk() (*bool, bool)`

GetCanApprovePostsOk returns a tuple with the CanApprovePosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanApprovePosts

`func (o *CurrentUser) SetCanApprovePosts(v bool)`

SetCanApprovePosts sets CanApprovePosts field to given value.


### GetCanUploadFree

`func (o *CurrentUser) GetCanUploadFree() bool`

GetCanUploadFree returns the CanUploadFree field if non-nil, zero value otherwise.

### GetCanUploadFreeOk

`func (o *CurrentUser) GetCanUploadFreeOk() (*bool, bool)`

GetCanUploadFreeOk returns a tuple with the CanUploadFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUploadFree

`func (o *CurrentUser) SetCanUploadFree(v bool)`

SetCanUploadFree sets CanUploadFree field to given value.


### GetLevelString

`func (o *CurrentUser) GetLevelString() string`

GetLevelString returns the LevelString field if non-nil, zero value otherwise.

### GetLevelStringOk

`func (o *CurrentUser) GetLevelStringOk() (*string, bool)`

GetLevelStringOk returns a tuple with the LevelString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelString

`func (o *CurrentUser) SetLevelString(v string)`

SetLevelString sets LevelString field to given value.


### GetAvatarId

`func (o *CurrentUser) GetAvatarId() int32`

GetAvatarId returns the AvatarId field if non-nil, zero value otherwise.

### GetAvatarIdOk

`func (o *CurrentUser) GetAvatarIdOk() (*int32, bool)`

GetAvatarIdOk returns a tuple with the AvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarId

`func (o *CurrentUser) SetAvatarId(v int32)`

SetAvatarId sets AvatarId field to given value.


### GetBlacklistUsers

`func (o *CurrentUser) GetBlacklistUsers() bool`

GetBlacklistUsers returns the BlacklistUsers field if non-nil, zero value otherwise.

### GetBlacklistUsersOk

`func (o *CurrentUser) GetBlacklistUsersOk() (*bool, bool)`

GetBlacklistUsersOk returns a tuple with the BlacklistUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistUsers

`func (o *CurrentUser) SetBlacklistUsers(v bool)`

SetBlacklistUsers sets BlacklistUsers field to given value.


### GetDescriptionCollapsedInitially

`func (o *CurrentUser) GetDescriptionCollapsedInitially() bool`

GetDescriptionCollapsedInitially returns the DescriptionCollapsedInitially field if non-nil, zero value otherwise.

### GetDescriptionCollapsedInitiallyOk

`func (o *CurrentUser) GetDescriptionCollapsedInitiallyOk() (*bool, bool)`

GetDescriptionCollapsedInitiallyOk returns a tuple with the DescriptionCollapsedInitially field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionCollapsedInitially

`func (o *CurrentUser) SetDescriptionCollapsedInitially(v bool)`

SetDescriptionCollapsedInitially sets DescriptionCollapsedInitially field to given value.


### GetHideComments

`func (o *CurrentUser) GetHideComments() bool`

GetHideComments returns the HideComments field if non-nil, zero value otherwise.

### GetHideCommentsOk

`func (o *CurrentUser) GetHideCommentsOk() (*bool, bool)`

GetHideCommentsOk returns a tuple with the HideComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideComments

`func (o *CurrentUser) SetHideComments(v bool)`

SetHideComments sets HideComments field to given value.


### GetShowHiddenComments

`func (o *CurrentUser) GetShowHiddenComments() bool`

GetShowHiddenComments returns the ShowHiddenComments field if non-nil, zero value otherwise.

### GetShowHiddenCommentsOk

`func (o *CurrentUser) GetShowHiddenCommentsOk() (*bool, bool)`

GetShowHiddenCommentsOk returns a tuple with the ShowHiddenComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowHiddenComments

`func (o *CurrentUser) SetShowHiddenComments(v bool)`

SetShowHiddenComments sets ShowHiddenComments field to given value.


### GetShowPostStatistics

`func (o *CurrentUser) GetShowPostStatistics() bool`

GetShowPostStatistics returns the ShowPostStatistics field if non-nil, zero value otherwise.

### GetShowPostStatisticsOk

`func (o *CurrentUser) GetShowPostStatisticsOk() (*bool, bool)`

GetShowPostStatisticsOk returns a tuple with the ShowPostStatistics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPostStatistics

`func (o *CurrentUser) SetShowPostStatistics(v bool)`

SetShowPostStatistics sets ShowPostStatistics field to given value.


### GetReceiveEmailNotifications

`func (o *CurrentUser) GetReceiveEmailNotifications() bool`

GetReceiveEmailNotifications returns the ReceiveEmailNotifications field if non-nil, zero value otherwise.

### GetReceiveEmailNotificationsOk

`func (o *CurrentUser) GetReceiveEmailNotificationsOk() (*bool, bool)`

GetReceiveEmailNotificationsOk returns a tuple with the ReceiveEmailNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveEmailNotifications

`func (o *CurrentUser) SetReceiveEmailNotifications(v bool)`

SetReceiveEmailNotifications sets ReceiveEmailNotifications field to given value.


### GetEnableKeyboardNavigation

`func (o *CurrentUser) GetEnableKeyboardNavigation() bool`

GetEnableKeyboardNavigation returns the EnableKeyboardNavigation field if non-nil, zero value otherwise.

### GetEnableKeyboardNavigationOk

`func (o *CurrentUser) GetEnableKeyboardNavigationOk() (*bool, bool)`

GetEnableKeyboardNavigationOk returns a tuple with the EnableKeyboardNavigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableKeyboardNavigation

`func (o *CurrentUser) SetEnableKeyboardNavigation(v bool)`

SetEnableKeyboardNavigation sets EnableKeyboardNavigation field to given value.


### GetEnablePrivacyMode

`func (o *CurrentUser) GetEnablePrivacyMode() bool`

GetEnablePrivacyMode returns the EnablePrivacyMode field if non-nil, zero value otherwise.

### GetEnablePrivacyModeOk

`func (o *CurrentUser) GetEnablePrivacyModeOk() (*bool, bool)`

GetEnablePrivacyModeOk returns a tuple with the EnablePrivacyMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePrivacyMode

`func (o *CurrentUser) SetEnablePrivacyMode(v bool)`

SetEnablePrivacyMode sets EnablePrivacyMode field to given value.


### GetStyleUsernames

`func (o *CurrentUser) GetStyleUsernames() bool`

GetStyleUsernames returns the StyleUsernames field if non-nil, zero value otherwise.

### GetStyleUsernamesOk

`func (o *CurrentUser) GetStyleUsernamesOk() (*bool, bool)`

GetStyleUsernamesOk returns a tuple with the StyleUsernames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyleUsernames

`func (o *CurrentUser) SetStyleUsernames(v bool)`

SetStyleUsernames sets StyleUsernames field to given value.


### GetEnableAutoComplete

`func (o *CurrentUser) GetEnableAutoComplete() bool`

GetEnableAutoComplete returns the EnableAutoComplete field if non-nil, zero value otherwise.

### GetEnableAutoCompleteOk

`func (o *CurrentUser) GetEnableAutoCompleteOk() (*bool, bool)`

GetEnableAutoCompleteOk returns a tuple with the EnableAutoComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoComplete

`func (o *CurrentUser) SetEnableAutoComplete(v bool)`

SetEnableAutoComplete sets EnableAutoComplete field to given value.


### GetDisableCroppedThumbnails

`func (o *CurrentUser) GetDisableCroppedThumbnails() bool`

GetDisableCroppedThumbnails returns the DisableCroppedThumbnails field if non-nil, zero value otherwise.

### GetDisableCroppedThumbnailsOk

`func (o *CurrentUser) GetDisableCroppedThumbnailsOk() (*bool, bool)`

GetDisableCroppedThumbnailsOk returns a tuple with the DisableCroppedThumbnails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableCroppedThumbnails

`func (o *CurrentUser) SetDisableCroppedThumbnails(v bool)`

SetDisableCroppedThumbnails sets DisableCroppedThumbnails field to given value.


### GetEnableSafeMode

`func (o *CurrentUser) GetEnableSafeMode() bool`

GetEnableSafeMode returns the EnableSafeMode field if non-nil, zero value otherwise.

### GetEnableSafeModeOk

`func (o *CurrentUser) GetEnableSafeModeOk() (*bool, bool)`

GetEnableSafeModeOk returns a tuple with the EnableSafeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableSafeMode

`func (o *CurrentUser) SetEnableSafeMode(v bool)`

SetEnableSafeMode sets EnableSafeMode field to given value.


### GetDisableResponsiveMode

`func (o *CurrentUser) GetDisableResponsiveMode() bool`

GetDisableResponsiveMode returns the DisableResponsiveMode field if non-nil, zero value otherwise.

### GetDisableResponsiveModeOk

`func (o *CurrentUser) GetDisableResponsiveModeOk() (*bool, bool)`

GetDisableResponsiveModeOk returns a tuple with the DisableResponsiveMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableResponsiveMode

`func (o *CurrentUser) SetDisableResponsiveMode(v bool)`

SetDisableResponsiveMode sets DisableResponsiveMode field to given value.


### GetNoFlagging

`func (o *CurrentUser) GetNoFlagging() bool`

GetNoFlagging returns the NoFlagging field if non-nil, zero value otherwise.

### GetNoFlaggingOk

`func (o *CurrentUser) GetNoFlaggingOk() (*bool, bool)`

GetNoFlaggingOk returns a tuple with the NoFlagging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoFlagging

`func (o *CurrentUser) SetNoFlagging(v bool)`

SetNoFlagging sets NoFlagging field to given value.


### GetDisableUserDmails

`func (o *CurrentUser) GetDisableUserDmails() bool`

GetDisableUserDmails returns the DisableUserDmails field if non-nil, zero value otherwise.

### GetDisableUserDmailsOk

`func (o *CurrentUser) GetDisableUserDmailsOk() (*bool, bool)`

GetDisableUserDmailsOk returns a tuple with the DisableUserDmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableUserDmails

`func (o *CurrentUser) SetDisableUserDmails(v bool)`

SetDisableUserDmails sets DisableUserDmails field to given value.


### GetEnableCompactUploader

`func (o *CurrentUser) GetEnableCompactUploader() bool`

GetEnableCompactUploader returns the EnableCompactUploader field if non-nil, zero value otherwise.

### GetEnableCompactUploaderOk

`func (o *CurrentUser) GetEnableCompactUploaderOk() (*bool, bool)`

GetEnableCompactUploaderOk returns a tuple with the EnableCompactUploader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableCompactUploader

`func (o *CurrentUser) SetEnableCompactUploader(v bool)`

SetEnableCompactUploader sets EnableCompactUploader field to given value.


### GetReplacementsBeta

`func (o *CurrentUser) GetReplacementsBeta() bool`

GetReplacementsBeta returns the ReplacementsBeta field if non-nil, zero value otherwise.

### GetReplacementsBetaOk

`func (o *CurrentUser) GetReplacementsBetaOk() (*bool, bool)`

GetReplacementsBetaOk returns a tuple with the ReplacementsBeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacementsBeta

`func (o *CurrentUser) SetReplacementsBeta(v bool)`

SetReplacementsBeta sets ReplacementsBeta field to given value.


### GetUpdatedAt

`func (o *CurrentUser) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CurrentUser) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CurrentUser) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetEmail

`func (o *CurrentUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CurrentUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CurrentUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetLastLoggedInAt

`func (o *CurrentUser) GetLastLoggedInAt() string`

GetLastLoggedInAt returns the LastLoggedInAt field if non-nil, zero value otherwise.

### GetLastLoggedInAtOk

`func (o *CurrentUser) GetLastLoggedInAtOk() (*string, bool)`

GetLastLoggedInAtOk returns a tuple with the LastLoggedInAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoggedInAt

`func (o *CurrentUser) SetLastLoggedInAt(v string)`

SetLastLoggedInAt sets LastLoggedInAt field to given value.


### GetLastForumReadAt

`func (o *CurrentUser) GetLastForumReadAt() string`

GetLastForumReadAt returns the LastForumReadAt field if non-nil, zero value otherwise.

### GetLastForumReadAtOk

`func (o *CurrentUser) GetLastForumReadAtOk() (*string, bool)`

GetLastForumReadAtOk returns a tuple with the LastForumReadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastForumReadAt

`func (o *CurrentUser) SetLastForumReadAt(v string)`

SetLastForumReadAt sets LastForumReadAt field to given value.


### GetRecentTags

`func (o *CurrentUser) GetRecentTags() string`

GetRecentTags returns the RecentTags field if non-nil, zero value otherwise.

### GetRecentTagsOk

`func (o *CurrentUser) GetRecentTagsOk() (*string, bool)`

GetRecentTagsOk returns a tuple with the RecentTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentTags

`func (o *CurrentUser) SetRecentTags(v string)`

SetRecentTags sets RecentTags field to given value.


### GetCommentThreshold

`func (o *CurrentUser) GetCommentThreshold() int32`

GetCommentThreshold returns the CommentThreshold field if non-nil, zero value otherwise.

### GetCommentThresholdOk

`func (o *CurrentUser) GetCommentThresholdOk() (*int32, bool)`

GetCommentThresholdOk returns a tuple with the CommentThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentThreshold

`func (o *CurrentUser) SetCommentThreshold(v int32)`

SetCommentThreshold sets CommentThreshold field to given value.


### GetDefaultImageSize

`func (o *CurrentUser) GetDefaultImageSize() string`

GetDefaultImageSize returns the DefaultImageSize field if non-nil, zero value otherwise.

### GetDefaultImageSizeOk

`func (o *CurrentUser) GetDefaultImageSizeOk() (*string, bool)`

GetDefaultImageSizeOk returns a tuple with the DefaultImageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImageSize

`func (o *CurrentUser) SetDefaultImageSize(v string)`

SetDefaultImageSize sets DefaultImageSize field to given value.


### GetFavoriteTags

`func (o *CurrentUser) GetFavoriteTags() string`

GetFavoriteTags returns the FavoriteTags field if non-nil, zero value otherwise.

### GetFavoriteTagsOk

`func (o *CurrentUser) GetFavoriteTagsOk() (*string, bool)`

GetFavoriteTagsOk returns a tuple with the FavoriteTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoriteTags

`func (o *CurrentUser) SetFavoriteTags(v string)`

SetFavoriteTags sets FavoriteTags field to given value.


### GetBlacklistedTags

`func (o *CurrentUser) GetBlacklistedTags() string`

GetBlacklistedTags returns the BlacklistedTags field if non-nil, zero value otherwise.

### GetBlacklistedTagsOk

`func (o *CurrentUser) GetBlacklistedTagsOk() (*string, bool)`

GetBlacklistedTagsOk returns a tuple with the BlacklistedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlacklistedTags

`func (o *CurrentUser) SetBlacklistedTags(v string)`

SetBlacklistedTags sets BlacklistedTags field to given value.


### GetTimeZone

`func (o *CurrentUser) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *CurrentUser) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *CurrentUser) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.


### GetPerPage

`func (o *CurrentUser) GetPerPage() int32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *CurrentUser) GetPerPageOk() (*int32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *CurrentUser) SetPerPage(v int32)`

SetPerPage sets PerPage field to given value.


### GetCustomStyle

`func (o *CurrentUser) GetCustomStyle() string`

GetCustomStyle returns the CustomStyle field if non-nil, zero value otherwise.

### GetCustomStyleOk

`func (o *CurrentUser) GetCustomStyleOk() (*string, bool)`

GetCustomStyleOk returns a tuple with the CustomStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomStyle

`func (o *CurrentUser) SetCustomStyle(v string)`

SetCustomStyle sets CustomStyle field to given value.


### GetFavoriteCount

`func (o *CurrentUser) GetFavoriteCount() int32`

GetFavoriteCount returns the FavoriteCount field if non-nil, zero value otherwise.

### GetFavoriteCountOk

`func (o *CurrentUser) GetFavoriteCountOk() (*int32, bool)`

GetFavoriteCountOk returns a tuple with the FavoriteCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoriteCount

`func (o *CurrentUser) SetFavoriteCount(v int32)`

SetFavoriteCount sets FavoriteCount field to given value.


### GetApiRegenMultiplier

`func (o *CurrentUser) GetApiRegenMultiplier() int32`

GetApiRegenMultiplier returns the ApiRegenMultiplier field if non-nil, zero value otherwise.

### GetApiRegenMultiplierOk

`func (o *CurrentUser) GetApiRegenMultiplierOk() (*int32, bool)`

GetApiRegenMultiplierOk returns a tuple with the ApiRegenMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiRegenMultiplier

`func (o *CurrentUser) SetApiRegenMultiplier(v int32)`

SetApiRegenMultiplier sets ApiRegenMultiplier field to given value.


### GetApiBurstLimit

`func (o *CurrentUser) GetApiBurstLimit() int32`

GetApiBurstLimit returns the ApiBurstLimit field if non-nil, zero value otherwise.

### GetApiBurstLimitOk

`func (o *CurrentUser) GetApiBurstLimitOk() (*int32, bool)`

GetApiBurstLimitOk returns a tuple with the ApiBurstLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBurstLimit

`func (o *CurrentUser) SetApiBurstLimit(v int32)`

SetApiBurstLimit sets ApiBurstLimit field to given value.


### GetRemainingApiLimit

`func (o *CurrentUser) GetRemainingApiLimit() int32`

GetRemainingApiLimit returns the RemainingApiLimit field if non-nil, zero value otherwise.

### GetRemainingApiLimitOk

`func (o *CurrentUser) GetRemainingApiLimitOk() (*int32, bool)`

GetRemainingApiLimitOk returns a tuple with the RemainingApiLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingApiLimit

`func (o *CurrentUser) SetRemainingApiLimit(v int32)`

SetRemainingApiLimit sets RemainingApiLimit field to given value.


### GetStatementTimeout

`func (o *CurrentUser) GetStatementTimeout() int32`

GetStatementTimeout returns the StatementTimeout field if non-nil, zero value otherwise.

### GetStatementTimeoutOk

`func (o *CurrentUser) GetStatementTimeoutOk() (*int32, bool)`

GetStatementTimeoutOk returns a tuple with the StatementTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatementTimeout

`func (o *CurrentUser) SetStatementTimeout(v int32)`

SetStatementTimeout sets StatementTimeout field to given value.


### GetFavoriteLimit

`func (o *CurrentUser) GetFavoriteLimit() int32`

GetFavoriteLimit returns the FavoriteLimit field if non-nil, zero value otherwise.

### GetFavoriteLimitOk

`func (o *CurrentUser) GetFavoriteLimitOk() (*int32, bool)`

GetFavoriteLimitOk returns a tuple with the FavoriteLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavoriteLimit

`func (o *CurrentUser) SetFavoriteLimit(v int32)`

SetFavoriteLimit sets FavoriteLimit field to given value.


### GetTagQueryLimit

`func (o *CurrentUser) GetTagQueryLimit() int32`

GetTagQueryLimit returns the TagQueryLimit field if non-nil, zero value otherwise.

### GetTagQueryLimitOk

`func (o *CurrentUser) GetTagQueryLimitOk() (*int32, bool)`

GetTagQueryLimitOk returns a tuple with the TagQueryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagQueryLimit

`func (o *CurrentUser) SetTagQueryLimit(v int32)`

SetTagQueryLimit sets TagQueryLimit field to given value.


### GetHasMail

`func (o *CurrentUser) GetHasMail() bool`

GetHasMail returns the HasMail field if non-nil, zero value otherwise.

### GetHasMailOk

`func (o *CurrentUser) GetHasMailOk() (*bool, bool)`

GetHasMailOk returns a tuple with the HasMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMail

`func (o *CurrentUser) SetHasMail(v bool)`

SetHasMail sets HasMail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


