/*
E621

OpenAPI definition for E621's API. You can find the source [here](https://github.com/DonovanDMC/E621OpenAPI)<br> This document is intended to compliment E621's existing [API Documentation](https://e621.net/help/api).<br> Note if E621's api is under attack and/or cloudflare protections are enabled, the \"Try it out\" buttons here will not work.<br> If they are not working, you can check this [Unofficial Status Page](https://status.e621.ws).

API version: d69c34e
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// UsersAPIService UsersAPI service
type UsersAPIService service

type UsersAPIEditCurrentUserRequest struct {
	ctx                               context.Context
	ApiService                        *UsersAPIService
	id                                int32
	userCommentThreshold              *float32
	userDefaultImageSize              *string
	userFavoriteTags                  *string
	userBlacklistedTags               *string
	userTimeZone                      *string
	userPerPage                       *float32
	userCustomStyle                   *string
	userDescriptionCollapsedInitially *bool
	userHideComments                  *bool
	userReceiveEmailNotifications     *bool
	userEnableKeyboardNavigation      *bool
	userEnablePrivacyMode             *bool
	userDisableUserDmails             *bool
	userBlacklistUsers                *bool
	userShowPostStatistics            *bool
	userStyleUsernames                *bool
	userShowHiddenComments            *bool
	userEnableAutocomplete            *bool
	userDisableCroppedThumbnails      *bool
	userEnableSafeMode                *bool
	userDisableResponsiveMode         *bool
	userDmailFilterAttributesId       *float32
	userDmailFilterAttributesWords    *string
	userProfileAbout                  *string
	userProfileArtinfo                *string
	userAvatarId                      *int32
	userEnableCompactUploader         *bool
}

func (r UsersAPIEditCurrentUserRequest) UserCommentThreshold(userCommentThreshold float32) UsersAPIEditCurrentUserRequest {
	r.userCommentThreshold = &userCommentThreshold
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserDefaultImageSize(userDefaultImageSize string) UsersAPIEditCurrentUserRequest {
	r.userDefaultImageSize = &userDefaultImageSize
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserFavoriteTags(userFavoriteTags string) UsersAPIEditCurrentUserRequest {
	r.userFavoriteTags = &userFavoriteTags
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserBlacklistedTags(userBlacklistedTags string) UsersAPIEditCurrentUserRequest {
	r.userBlacklistedTags = &userBlacklistedTags
	return r
}

// https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
func (r UsersAPIEditCurrentUserRequest) UserTimeZone(userTimeZone string) UsersAPIEditCurrentUserRequest {
	r.userTimeZone = &userTimeZone
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserPerPage(userPerPage float32) UsersAPIEditCurrentUserRequest {
	r.userPerPage = &userPerPage
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserCustomStyle(userCustomStyle string) UsersAPIEditCurrentUserRequest {
	r.userCustomStyle = &userCustomStyle
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserDescriptionCollapsedInitially(userDescriptionCollapsedInitially bool) UsersAPIEditCurrentUserRequest {
	r.userDescriptionCollapsedInitially = &userDescriptionCollapsedInitially
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserHideComments(userHideComments bool) UsersAPIEditCurrentUserRequest {
	r.userHideComments = &userHideComments
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserReceiveEmailNotifications(userReceiveEmailNotifications bool) UsersAPIEditCurrentUserRequest {
	r.userReceiveEmailNotifications = &userReceiveEmailNotifications
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserEnableKeyboardNavigation(userEnableKeyboardNavigation bool) UsersAPIEditCurrentUserRequest {
	r.userEnableKeyboardNavigation = &userEnableKeyboardNavigation
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserEnablePrivacyMode(userEnablePrivacyMode bool) UsersAPIEditCurrentUserRequest {
	r.userEnablePrivacyMode = &userEnablePrivacyMode
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserDisableUserDmails(userDisableUserDmails bool) UsersAPIEditCurrentUserRequest {
	r.userDisableUserDmails = &userDisableUserDmails
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserBlacklistUsers(userBlacklistUsers bool) UsersAPIEditCurrentUserRequest {
	r.userBlacklistUsers = &userBlacklistUsers
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserShowPostStatistics(userShowPostStatistics bool) UsersAPIEditCurrentUserRequest {
	r.userShowPostStatistics = &userShowPostStatistics
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserStyleUsernames(userStyleUsernames bool) UsersAPIEditCurrentUserRequest {
	r.userStyleUsernames = &userStyleUsernames
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserShowHiddenComments(userShowHiddenComments bool) UsersAPIEditCurrentUserRequest {
	r.userShowHiddenComments = &userShowHiddenComments
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserEnableAutocomplete(userEnableAutocomplete bool) UsersAPIEditCurrentUserRequest {
	r.userEnableAutocomplete = &userEnableAutocomplete
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserDisableCroppedThumbnails(userDisableCroppedThumbnails bool) UsersAPIEditCurrentUserRequest {
	r.userDisableCroppedThumbnails = &userDisableCroppedThumbnails
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserEnableSafeMode(userEnableSafeMode bool) UsersAPIEditCurrentUserRequest {
	r.userEnableSafeMode = &userEnableSafeMode
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserDisableResponsiveMode(userDisableResponsiveMode bool) UsersAPIEditCurrentUserRequest {
	r.userDisableResponsiveMode = &userDisableResponsiveMode
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserDmailFilterAttributesId(userDmailFilterAttributesId float32) UsersAPIEditCurrentUserRequest {
	r.userDmailFilterAttributesId = &userDmailFilterAttributesId
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserDmailFilterAttributesWords(userDmailFilterAttributesWords string) UsersAPIEditCurrentUserRequest {
	r.userDmailFilterAttributesWords = &userDmailFilterAttributesWords
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserProfileAbout(userProfileAbout string) UsersAPIEditCurrentUserRequest {
	r.userProfileAbout = &userProfileAbout
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserProfileArtinfo(userProfileArtinfo string) UsersAPIEditCurrentUserRequest {
	r.userProfileArtinfo = &userProfileArtinfo
	return r
}

func (r UsersAPIEditCurrentUserRequest) UserAvatarId(userAvatarId int32) UsersAPIEditCurrentUserRequest {
	r.userAvatarId = &userAvatarId
	return r
}

// You must have uploaded at least 10 posts.
func (r UsersAPIEditCurrentUserRequest) UserEnableCompactUploader(userEnableCompactUploader bool) UsersAPIEditCurrentUserRequest {
	r.userEnableCompactUploader = &userEnableCompactUploader
	return r
}

func (r UsersAPIEditCurrentUserRequest) Execute() (*http.Response, error) {
	return r.ApiService.EditCurrentUserExecute(r)
}

/*
EditCurrentUser Edit Current User

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID of the user. The actual value is ignored, but something must be supplied.
	@return UsersAPIEditCurrentUserRequest
*/
func (a *UsersAPIService) EditCurrentUser(ctx context.Context, id int32) UsersAPIEditCurrentUserRequest {
	return UsersAPIEditCurrentUserRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *UsersAPIService) EditCurrentUserExecute(r UsersAPIEditCurrentUserRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPatch
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.EditCurrentUser")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{id}.json"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.userCommentThreshold != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[comment_threshold]", r.userCommentThreshold, "")
	}
	if r.userDefaultImageSize != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[default_image_size]", r.userDefaultImageSize, "")
	}
	if r.userFavoriteTags != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[favorite_tags]", r.userFavoriteTags, "")
	}
	if r.userBlacklistedTags != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[blacklisted_tags]", r.userBlacklistedTags, "")
	}
	if r.userTimeZone != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[time_zone]", r.userTimeZone, "")
	}
	if r.userPerPage != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[per_page]", r.userPerPage, "")
	}
	if r.userCustomStyle != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[custom_style]", r.userCustomStyle, "")
	}
	if r.userDescriptionCollapsedInitially != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[description_collapsed_initially]", r.userDescriptionCollapsedInitially, "")
	}
	if r.userHideComments != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[hide_comments]", r.userHideComments, "")
	}
	if r.userReceiveEmailNotifications != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[receive_email_notifications]", r.userReceiveEmailNotifications, "")
	}
	if r.userEnableKeyboardNavigation != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[enable_keyboard_navigation]", r.userEnableKeyboardNavigation, "")
	}
	if r.userEnablePrivacyMode != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[enable_privacy_mode]", r.userEnablePrivacyMode, "")
	}
	if r.userDisableUserDmails != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[disable_user_dmails]", r.userDisableUserDmails, "")
	}
	if r.userBlacklistUsers != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[blacklist_users]", r.userBlacklistUsers, "")
	}
	if r.userShowPostStatistics != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[show_post_statistics]", r.userShowPostStatistics, "")
	}
	if r.userStyleUsernames != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[style_usernames]", r.userStyleUsernames, "")
	}
	if r.userShowHiddenComments != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[show_hidden_comments]", r.userShowHiddenComments, "")
	}
	if r.userEnableAutocomplete != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[enable_autocomplete]", r.userEnableAutocomplete, "")
	}
	if r.userDisableCroppedThumbnails != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[disable_cropped_thumbnails]", r.userDisableCroppedThumbnails, "")
	}
	if r.userEnableSafeMode != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[enable_safe_mode]", r.userEnableSafeMode, "")
	}
	if r.userDisableResponsiveMode != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[disable_responsive_mode]", r.userDisableResponsiveMode, "")
	}
	if r.userDmailFilterAttributesId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[dmail_filter_attributes][id]", r.userDmailFilterAttributesId, "")
	}
	if r.userDmailFilterAttributesWords != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[dmail_filter_attributes][words]", r.userDmailFilterAttributesWords, "")
	}
	if r.userProfileAbout != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[profile_about]", r.userProfileAbout, "")
	}
	if r.userProfileArtinfo != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[profile_artinfo]", r.userProfileArtinfo, "")
	}
	if r.userAvatarId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[avatar_id]", r.userAvatarId, "")
	}
	if r.userEnableCompactUploader != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "user[enable_compact_uploader]", r.userEnableCompactUploader, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type UsersAPIFixUserCountsRequest struct {
	ctx        context.Context
	ApiService *UsersAPIService
}

func (r UsersAPIFixUserCountsRequest) Execute() (*http.Response, error) {
	return r.ApiService.FixUserCountsExecute(r)
}

/*
FixUserCounts Fix User Counts

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return UsersAPIFixUserCountsRequest
*/
func (a *UsersAPIService) FixUserCounts(ctx context.Context) UsersAPIFixUserCountsRequest {
	return UsersAPIFixUserCountsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *UsersAPIService) FixUserCountsExecute(r UsersAPIFixUserCountsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.FixUserCounts")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/maintenance/user/count_fixes.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type UsersAPIGetCurrentUserRequest struct {
	ctx        context.Context
	ApiService *UsersAPIService
}

func (r UsersAPIGetCurrentUserRequest) Execute() (*CurrentUser, *http.Response, error) {
	return r.ApiService.GetCurrentUserExecute(r)
}

/*
GetCurrentUser Get Current User

This is a crude but effective way to get the currently authenticated user without scraping HTML.
Note that this route does not include some properties included in the show action.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return UsersAPIGetCurrentUserRequest
*/
func (a *UsersAPIService) GetCurrentUser(ctx context.Context) UsersAPIGetCurrentUserRequest {
	return UsersAPIGetCurrentUserRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CurrentUser
func (a *UsersAPIService) GetCurrentUserExecute(r UsersAPIGetCurrentUserRequest) (*CurrentUser, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CurrentUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetCurrentUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/upload_limit.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UsersAPIGetUserRequest struct {
	ctx        context.Context
	ApiService *UsersAPIService
	id         string
}

func (r UsersAPIGetUserRequest) Execute() (*GetUser200Response, *http.Response, error) {
	return r.ApiService.GetUserExecute(r)
}

/*
GetUser Get User

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The ID or Username of the user.
	@return UsersAPIGetUserRequest
*/
func (a *UsersAPIService) GetUser(ctx context.Context, id string) UsersAPIGetUserRequest {
	return UsersAPIGetUserRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return GetUser200Response
func (a *UsersAPIService) GetUserExecute(r UsersAPIGetUserRequest) (*GetUser200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetUser200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.GetUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users/{id}.json"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UsersAPISearchUsersRequest struct {
	ctx                   context.Context
	ApiService            *UsersAPIService
	limit                 *int32
	page                  *int32
	searchId              *int32
	searchIpAddr          *string
	searchOrder           *string
	searchNameMatches     *string
	searchAboutMe         *string
	searchAvatarId        *float32
	searchLevel           *float32
	searchMinLevel        *float32
	searchMaxLevel        *float32
	searchCanUploadFree   *bool
	searchCanApprovePosts *bool
	searchEmailMatches    *string
}

// The maximum number of results to return. Between 0 and 320.
func (r UsersAPISearchUsersRequest) Limit(limit int32) UsersAPISearchUsersRequest {
	r.limit = &limit
	return r
}

// The page number of results to get. Between 1 and 750.
func (r UsersAPISearchUsersRequest) Page(page int32) UsersAPISearchUsersRequest {
	r.page = &page
	return r
}

// Search for a specific id.
func (r UsersAPISearchUsersRequest) SearchId(searchId int32) UsersAPISearchUsersRequest {
	r.searchId = &searchId
	return r
}

// Must be Admin+ to use. See [postgres&#39; documentation](https://www.postgresql.org/docs/9.3/functions-net.html) for information on how this is parsed. Specifically, \&quot;is contained within or equals\&quot; (&#x60;&lt;&lt;&#x3D;&#x60;).
func (r UsersAPISearchUsersRequest) SearchIpAddr(searchIpAddr string) UsersAPISearchUsersRequest {
	r.searchIpAddr = &searchIpAddr
	return r
}

func (r UsersAPISearchUsersRequest) SearchOrder(searchOrder string) UsersAPISearchUsersRequest {
	r.searchOrder = &searchOrder
	return r
}

func (r UsersAPISearchUsersRequest) SearchNameMatches(searchNameMatches string) UsersAPISearchUsersRequest {
	r.searchNameMatches = &searchNameMatches
	return r
}

func (r UsersAPISearchUsersRequest) SearchAboutMe(searchAboutMe string) UsersAPISearchUsersRequest {
	r.searchAboutMe = &searchAboutMe
	return r
}

func (r UsersAPISearchUsersRequest) SearchAvatarId(searchAvatarId float32) UsersAPISearchUsersRequest {
	r.searchAvatarId = &searchAvatarId
	return r
}

func (r UsersAPISearchUsersRequest) SearchLevel(searchLevel float32) UsersAPISearchUsersRequest {
	r.searchLevel = &searchLevel
	return r
}

func (r UsersAPISearchUsersRequest) SearchMinLevel(searchMinLevel float32) UsersAPISearchUsersRequest {
	r.searchMinLevel = &searchMinLevel
	return r
}

func (r UsersAPISearchUsersRequest) SearchMaxLevel(searchMaxLevel float32) UsersAPISearchUsersRequest {
	r.searchMaxLevel = &searchMaxLevel
	return r
}

func (r UsersAPISearchUsersRequest) SearchCanUploadFree(searchCanUploadFree bool) UsersAPISearchUsersRequest {
	r.searchCanUploadFree = &searchCanUploadFree
	return r
}

func (r UsersAPISearchUsersRequest) SearchCanApprovePosts(searchCanApprovePosts bool) UsersAPISearchUsersRequest {
	r.searchCanApprovePosts = &searchCanApprovePosts
	return r
}

// You must be Admin+.
func (r UsersAPISearchUsersRequest) SearchEmailMatches(searchEmailMatches string) UsersAPISearchUsersRequest {
	r.searchEmailMatches = &searchEmailMatches
	return r
}

func (r UsersAPISearchUsersRequest) Execute() ([]SearchUsers200ResponseInner, *http.Response, error) {
	return r.ApiService.SearchUsersExecute(r)
}

/*
SearchUsers Search Users

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return UsersAPISearchUsersRequest
*/
func (a *UsersAPIService) SearchUsers(ctx context.Context) UsersAPISearchUsersRequest {
	return UsersAPISearchUsersRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []SearchUsers200ResponseInner
func (a *UsersAPIService) SearchUsersExecute(r UsersAPISearchUsersRequest) ([]SearchUsers200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []SearchUsers200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.SearchUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/users.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "")
	}
	if r.searchId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[id]", r.searchId, "")
	}
	if r.searchIpAddr != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[ip_addr]", r.searchIpAddr, "")
	}
	if r.searchOrder != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[order]", r.searchOrder, "")
	}
	if r.searchNameMatches != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[name_matches]", r.searchNameMatches, "")
	}
	if r.searchAboutMe != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[about_me]", r.searchAboutMe, "")
	}
	if r.searchAvatarId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[avatar_id]", r.searchAvatarId, "")
	}
	if r.searchLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[level]", r.searchLevel, "")
	}
	if r.searchMinLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[min_level]", r.searchMinLevel, "")
	}
	if r.searchMaxLevel != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[max_level]", r.searchMaxLevel, "")
	}
	if r.searchCanUploadFree != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[can_upload_free]", r.searchCanUploadFree, "")
	}
	if r.searchCanApprovePosts != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[can_approve_posts]", r.searchCanApprovePosts, "")
	}
	if r.searchEmailMatches != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search[email_matches]", r.searchEmailMatches, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
