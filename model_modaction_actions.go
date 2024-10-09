/*
E621

OpenAPI definition for E621's API. You can find the source [here](https://github.com/DonovanDMC/E621OpenAPI)<br> This document is intended to compliment E621's existing [API Documentation](https://e621.net/help/api).<br> Note if E621's api is under attack and/or cloudflare protections are enabled, the \"Try it out\" buttons here will not work.<br> If they are not working, you can check this [Unofficial Status Page](https://status.e621.ws).

API version: d69c34e
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// ModactionActions the model 'ModactionActions'
type ModactionActions string

// List of ModactionActions
const (
	MODACTIONACTIONS_ARTIST_PAGE_RENAME       ModactionActions = "artist_page_rename"
	MODACTIONACTIONS_ARTIST_PAGE_LOCK         ModactionActions = "artist_page_lock"
	MODACTIONACTIONS_ARTIST_PAGE_UNLOCK       ModactionActions = "artist_page_unlock"
	MODACTIONACTIONS_ARTIST_USER_LINKED       ModactionActions = "artist_user_linked"
	MODACTIONACTIONS_ARTIST_USER_UNLINKED     ModactionActions = "artist_user_unlinked"
	MODACTIONACTIONS_AVOID_POSTING_CREATE     ModactionActions = "avoid_posting_create"
	MODACTIONACTIONS_AVOID_POSTING_UPDATE     ModactionActions = "avoid_posting_update"
	MODACTIONACTIONS_AVOID_POSTING_DELETE     ModactionActions = "avoid_posting_delete"
	MODACTIONACTIONS_AVOID_POSTING_UNDELETE   ModactionActions = "avoid_posting_undelete"
	MODACTIONACTIONS_AVOID_POSTING_DESTROY    ModactionActions = "avoid_posting_destroy"
	MODACTIONACTIONS_BLIP_DELETE              ModactionActions = "blip_delete"
	MODACTIONACTIONS_BLIP_HIDE                ModactionActions = "blip_hide"
	MODACTIONACTIONS_BLIP_UNHIDE              ModactionActions = "blip_unhide"
	MODACTIONACTIONS_BLIP_UPDATE              ModactionActions = "blip_update"
	MODACTIONACTIONS_COMMENT_DELETE           ModactionActions = "comment_delete"
	MODACTIONACTIONS_COMMENT_HIDE             ModactionActions = "comment_hide"
	MODACTIONACTIONS_COMMENT_UNHIDE           ModactionActions = "comment_unhide"
	MODACTIONACTIONS_COMMENT_UPDATE           ModactionActions = "comment_update"
	MODACTIONACTIONS_FORUM_CATEGORY_CREATE    ModactionActions = "forum_category_create"
	MODACTIONACTIONS_FORUM_CATEGORY_DELETE    ModactionActions = "forum_category_delete"
	MODACTIONACTIONS_FORUM_CATEGORY_UPDATE    ModactionActions = "forum_category_update"
	MODACTIONACTIONS_FORUM_POST_DELETE        ModactionActions = "forum_post_delete"
	MODACTIONACTIONS_FORUM_POST_HIDE          ModactionActions = "forum_post_hide"
	MODACTIONACTIONS_FORUM_POST_UNHIDE        ModactionActions = "forum_post_unhide"
	MODACTIONACTIONS_FORUM_POST_UPDATE        ModactionActions = "forum_post_update"
	MODACTIONACTIONS_FORUM_TOPIC_DELETE       ModactionActions = "forum_topic_delete"
	MODACTIONACTIONS_FORUM_TOPIC_HIDE         ModactionActions = "forum_topic_hide"
	MODACTIONACTIONS_FORUM_TOPIC_UNHIDE       ModactionActions = "forum_topic_unhide"
	MODACTIONACTIONS_FORUM_TOPIC_LOCK         ModactionActions = "forum_topic_lock"
	MODACTIONACTIONS_FORUM_TOPIC_UNLOCK       ModactionActions = "forum_topic_unlock"
	MODACTIONACTIONS_FORUM_TOPIC_STICK        ModactionActions = "forum_topic_stick"
	MODACTIONACTIONS_FORUM_TOPIC_UNSTICK      ModactionActions = "forum_topic_unstick"
	MODACTIONACTIONS_FORUM_TOPIC_UPDATE       ModactionActions = "forum_topic_update"
	MODACTIONACTIONS_HELP_CREATE              ModactionActions = "help_create"
	MODACTIONACTIONS_HELP_DELETE              ModactionActions = "help_delete"
	MODACTIONACTIONS_HELP_UPDATE              ModactionActions = "help_update"
	MODACTIONACTIONS_IP_BAN_CREATE            ModactionActions = "ip_ban_create"
	MODACTIONACTIONS_IP_BAN_DELETE            ModactionActions = "ip_ban_delete"
	MODACTIONACTIONS_MASCOT_CREATE            ModactionActions = "mascot_create"
	MODACTIONACTIONS_MASCOT_UPDATE            ModactionActions = "mascot_update"
	MODACTIONACTIONS_MASCOT_DELETE            ModactionActions = "mascot_delete"
	MODACTIONACTIONS_POOL_DELETE              ModactionActions = "pool_delete"
	MODACTIONACTIONS_REPORT_REASON_CREATE     ModactionActions = "report_reason_create"
	MODACTIONACTIONS_REPORT_REASON_DELETE     ModactionActions = "report_reason_delete"
	MODACTIONACTIONS_REPORT_REASON_UPDATE     ModactionActions = "report_reason_update"
	MODACTIONACTIONS_SET_UPDATE               ModactionActions = "set_update"
	MODACTIONACTIONS_SET_DELETE               ModactionActions = "set_delete"
	MODACTIONACTIONS_SET_CHANGE_VISIBILITY    ModactionActions = "set_change_visibility"
	MODACTIONACTIONS_TAG_ALIAS_CREATE         ModactionActions = "tag_alias_create"
	MODACTIONACTIONS_TAG_ALIAS_UPDATE         ModactionActions = "tag_alias_update"
	MODACTIONACTIONS_TAG_IMPLICATION_CREATE   ModactionActions = "tag_implication_create"
	MODACTIONACTIONS_TAG_IMPLICATION_UPDATE   ModactionActions = "tag_implication_update"
	MODACTIONACTIONS_TICKET_CLAIM             ModactionActions = "ticket_claim"
	MODACTIONACTIONS_TICKET_UNCLAIM           ModactionActions = "ticket_unclaim"
	MODACTIONACTIONS_TICKET_UPDATE            ModactionActions = "ticket_update"
	MODACTIONACTIONS_UPLOAD_WHITELIST_CREATE  ModactionActions = "upload_whitelist_create"
	MODACTIONACTIONS_UPLOAD_WHITELIST_UPDATE  ModactionActions = "upload_whitelist_update"
	MODACTIONACTIONS_UPLOAD_WHITELIST_DELETE  ModactionActions = "upload_whitelist_delete"
	MODACTIONACTIONS_USER_BLACKLIST_CHANGED   ModactionActions = "user_blacklist_changed"
	MODACTIONACTIONS_USER_TEXT_CHANGE         ModactionActions = "user_text_change"
	MODACTIONACTIONS_USER_UPLOAD_LIMIT_CHANGE ModactionActions = "user_upload_limit_change"
	MODACTIONACTIONS_USER_FLAGS_CHANGE        ModactionActions = "user_flags_change"
	MODACTIONACTIONS_USER_LEVEL_CHANGE        ModactionActions = "user_level_change"
	MODACTIONACTIONS_USER_NAME_CHANGE         ModactionActions = "user_name_change"
	MODACTIONACTIONS_USER_DELETE              ModactionActions = "user_delete"
	MODACTIONACTIONS_USER_BAN                 ModactionActions = "user_ban"
	MODACTIONACTIONS_USER_BAN_UPDATE          ModactionActions = "user_ban_update"
	MODACTIONACTIONS_USER_UNBAN               ModactionActions = "user_unban"
	MODACTIONACTIONS_USER_FEEDBACK_CREATE     ModactionActions = "user_feedback_create"
	MODACTIONACTIONS_USER_FEEDBACK_UPDATE     ModactionActions = "user_feedback_update"
	MODACTIONACTIONS_USER_FEEDBACK_DELETE     ModactionActions = "user_feedback_delete"
	MODACTIONACTIONS_USER_FEEDBACK_UNDELETE   ModactionActions = "user_feedback_undelete"
	MODACTIONACTIONS_USER_FEEDBACK_DESTROY    ModactionActions = "user_feedback_destroy"
	MODACTIONACTIONS_WIKI_PAGE_RENAME         ModactionActions = "wiki_page_rename"
	MODACTIONACTIONS_WIKI_PAGE_DELETE         ModactionActions = "wiki_page_delete"
	MODACTIONACTIONS_WIKI_PAGE_LOCK           ModactionActions = "wiki_page_lock"
	MODACTIONACTIONS_WIKI_PAGE_UNLOCK         ModactionActions = "wiki_page_unlock"
	MODACTIONACTIONS_MASS_UPDATE              ModactionActions = "mass_update"
	MODACTIONACTIONS_NUKE_TAG                 ModactionActions = "nuke_tag"
	MODACTIONACTIONS_TAKEDOWN_DELETE          ModactionActions = "takedown_delete"
	MODACTIONACTIONS_TAKEDOWN_PROCESS         ModactionActions = "takedown_process"
	MODACTIONACTIONS_CREATED_POSITIVE_RECORD  ModactionActions = "created_positive_record"
	MODACTIONACTIONS_CREATED_NEUTRAL_RECORD   ModactionActions = "created_neutral_record"
	MODACTIONACTIONS_CREATED_NEGATIVE_RECORD  ModactionActions = "created_negative_record"
	MODACTIONACTIONS_CREATED_FLAG_REASON      ModactionActions = "created_flag_reason"
	MODACTIONACTIONS_EDITED_FLAG_REASON       ModactionActions = "edited_flag_reason"
	MODACTIONACTIONS_DELETED_FLAG_REASON      ModactionActions = "deleted_flag_reason"
	MODACTIONACTIONS_POST_MOVE_FAVORITES      ModactionActions = "post_move_favorites"
	MODACTIONACTIONS_POST_DELETE              ModactionActions = "post_delete"
	MODACTIONACTIONS_POST_UNDELETE            ModactionActions = "post_undelete"
	MODACTIONACTIONS_POST_DESTROY             ModactionActions = "post_destroy"
	MODACTIONACTIONS_POST_RATING_LOCK         ModactionActions = "post_rating_lock"
	MODACTIONACTIONS_POST_UNAPPROVE           ModactionActions = "post_unapprove"
	MODACTIONACTIONS_POST_REPLACEMENT_ACCEPT  ModactionActions = "post_replacement_accept"
	MODACTIONACTIONS_POST_REPLACEMENT_REJECT  ModactionActions = "post_replacement_reject"
	MODACTIONACTIONS_POST_REPLACEMENT_DELETE  ModactionActions = "post_replacement_delete"
)

// All allowed values of ModactionActions enum
var AllowedModactionActionsEnumValues = []ModactionActions{
	"artist_page_rename",
	"artist_page_lock",
	"artist_page_unlock",
	"artist_user_linked",
	"artist_user_unlinked",
	"avoid_posting_create",
	"avoid_posting_update",
	"avoid_posting_delete",
	"avoid_posting_undelete",
	"avoid_posting_destroy",
	"blip_delete",
	"blip_hide",
	"blip_unhide",
	"blip_update",
	"comment_delete",
	"comment_hide",
	"comment_unhide",
	"comment_update",
	"forum_category_create",
	"forum_category_delete",
	"forum_category_update",
	"forum_post_delete",
	"forum_post_hide",
	"forum_post_unhide",
	"forum_post_update",
	"forum_topic_delete",
	"forum_topic_hide",
	"forum_topic_unhide",
	"forum_topic_lock",
	"forum_topic_unlock",
	"forum_topic_stick",
	"forum_topic_unstick",
	"forum_topic_update",
	"help_create",
	"help_delete",
	"help_update",
	"ip_ban_create",
	"ip_ban_delete",
	"mascot_create",
	"mascot_update",
	"mascot_delete",
	"pool_delete",
	"report_reason_create",
	"report_reason_delete",
	"report_reason_update",
	"set_update",
	"set_delete",
	"set_change_visibility",
	"tag_alias_create",
	"tag_alias_update",
	"tag_implication_create",
	"tag_implication_update",
	"ticket_claim",
	"ticket_unclaim",
	"ticket_update",
	"upload_whitelist_create",
	"upload_whitelist_update",
	"upload_whitelist_delete",
	"user_blacklist_changed",
	"user_text_change",
	"user_upload_limit_change",
	"user_flags_change",
	"user_level_change",
	"user_name_change",
	"user_delete",
	"user_ban",
	"user_ban_update",
	"user_unban",
	"user_feedback_create",
	"user_feedback_update",
	"user_feedback_delete",
	"user_feedback_undelete",
	"user_feedback_destroy",
	"wiki_page_rename",
	"wiki_page_delete",
	"wiki_page_lock",
	"wiki_page_unlock",
	"mass_update",
	"nuke_tag",
	"takedown_delete",
	"takedown_process",
	"created_positive_record",
	"created_neutral_record",
	"created_negative_record",
	"created_flag_reason",
	"edited_flag_reason",
	"deleted_flag_reason",
	"post_move_favorites",
	"post_delete",
	"post_undelete",
	"post_destroy",
	"post_rating_lock",
	"post_unapprove",
	"post_replacement_accept",
	"post_replacement_reject",
	"post_replacement_delete",
}

func (v *ModactionActions) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ModactionActions(value)
	for _, existing := range AllowedModactionActionsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ModactionActions", value)
}

// NewModactionActionsFromValue returns a pointer to a valid ModactionActions
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewModactionActionsFromValue(v string) (*ModactionActions, error) {
	ev := ModactionActions(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ModactionActions: valid values are %v", v, AllowedModactionActionsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ModactionActions) IsValid() bool {
	for _, existing := range AllowedModactionActionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ModactionActions value
func (v ModactionActions) Ptr() *ModactionActions {
	return &v
}

type NullableModactionActions struct {
	value *ModactionActions
	isSet bool
}

func (v NullableModactionActions) Get() *ModactionActions {
	return v.value
}

func (v *NullableModactionActions) Set(val *ModactionActions) {
	v.value = val
	v.isSet = true
}

func (v NullableModactionActions) IsSet() bool {
	return v.isSet
}

func (v *NullableModactionActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModactionActions(val *ModactionActions) *NullableModactionActions {
	return &NullableModactionActions{value: val, isSet: true}
}

func (v NullableModactionActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModactionActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
