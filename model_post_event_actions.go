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

// PostEventActions the model 'PostEventActions'
type PostEventActions string

// List of PostEventActions
const (
	POSTEVENTACTIONS_DELETED              PostEventActions = "deleted"
	POSTEVENTACTIONS_UNDELETED            PostEventActions = "undeleted"
	POSTEVENTACTIONS_APPROVED             PostEventActions = "approved"
	POSTEVENTACTIONS_UNAPPROVED           PostEventActions = "unapproved"
	POSTEVENTACTIONS_FLAG_CREATED         PostEventActions = "flag_created"
	POSTEVENTACTIONS_FLAG_REMOVED         PostEventActions = "flag_removed"
	POSTEVENTACTIONS_FAVORITES_MOVED      PostEventActions = "favorites_moved"
	POSTEVENTACTIONS_FAVORITES_RECEIVED   PostEventActions = "favorites_received"
	POSTEVENTACTIONS_RATING_LOCKED        PostEventActions = "rating_locked"
	POSTEVENTACTIONS_RATING_UNLOCKED      PostEventActions = "rating_unlocked"
	POSTEVENTACTIONS_STATUS_LOCKED        PostEventActions = "status_locked"
	POSTEVENTACTIONS_STATUS_UNLOCKED      PostEventActions = "status_unlocked"
	POSTEVENTACTIONS_NOTE_LOCKED          PostEventActions = "note_locked"
	POSTEVENTACTIONS_NOTE_UNLOCKED        PostEventActions = "note_unlocked"
	POSTEVENTACTIONS_COMMENT_LOCKED       PostEventActions = "comment_locked"
	POSTEVENTACTIONS_COMMENT_UNLOCKED     PostEventActions = "comment_unlocked"
	POSTEVENTACTIONS_REPLACEMENT_ACCEPTED PostEventActions = "replacement_accepted"
	POSTEVENTACTIONS_REPLACEMENT_REJECTED PostEventActions = "replacement_rejected"
	POSTEVENTACTIONS_REPLACEMENT_PROMOTED PostEventActions = "replacement_promoted"
	POSTEVENTACTIONS_REPLACEMENT_DELETED  PostEventActions = "replacement_deleted"
	POSTEVENTACTIONS_EXPUNGED             PostEventActions = "expunged"
	POSTEVENTACTIONS_CHANGED_BG_COLOR     PostEventActions = "changed_bg_color"
)

// All allowed values of PostEventActions enum
var AllowedPostEventActionsEnumValues = []PostEventActions{
	"deleted",
	"undeleted",
	"approved",
	"unapproved",
	"flag_created",
	"flag_removed",
	"favorites_moved",
	"favorites_received",
	"rating_locked",
	"rating_unlocked",
	"status_locked",
	"status_unlocked",
	"note_locked",
	"note_unlocked",
	"comment_locked",
	"comment_unlocked",
	"replacement_accepted",
	"replacement_rejected",
	"replacement_promoted",
	"replacement_deleted",
	"expunged",
	"changed_bg_color",
}

func (v *PostEventActions) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PostEventActions(value)
	for _, existing := range AllowedPostEventActionsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PostEventActions", value)
}

// NewPostEventActionsFromValue returns a pointer to a valid PostEventActions
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPostEventActionsFromValue(v string) (*PostEventActions, error) {
	ev := PostEventActions(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PostEventActions: valid values are %v", v, AllowedPostEventActionsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PostEventActions) IsValid() bool {
	for _, existing := range AllowedPostEventActionsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PostEventActions value
func (v PostEventActions) Ptr() *PostEventActions {
	return &v
}

type NullablePostEventActions struct {
	value *PostEventActions
	isSet bool
}

func (v NullablePostEventActions) Get() *PostEventActions {
	return v.value
}

func (v *NullablePostEventActions) Set(val *PostEventActions) {
	v.value = val
	v.isSet = true
}

func (v NullablePostEventActions) IsSet() bool {
	return v.isSet
}

func (v *NullablePostEventActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostEventActions(val *PostEventActions) *NullablePostEventActions {
	return &NullablePostEventActions{value: val, isSet: true}
}

func (v NullablePostEventActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostEventActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
