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

// TagRequestStatuses the model 'TagRequestStatuses'
type TagRequestStatuses string

// List of TagRequestStatuses
const (
	TAGREQUESTSTATUSES_ACTIVE     TagRequestStatuses = "active"
	TAGREQUESTSTATUSES_DELETED    TagRequestStatuses = "deleted"
	TAGREQUESTSTATUSES_PROCESSING TagRequestStatuses = "processing"
	TAGREQUESTSTATUSES_QUEUED     TagRequestStatuses = "queued"
	TAGREQUESTSTATUSES_RETIRED    TagRequestStatuses = "retired"
	TAGREQUESTSTATUSES_ERROR      TagRequestStatuses = "error"
	TAGREQUESTSTATUSES_PENDING    TagRequestStatuses = "pending"
)

// All allowed values of TagRequestStatuses enum
var AllowedTagRequestStatusesEnumValues = []TagRequestStatuses{
	"active",
	"deleted",
	"processing",
	"queued",
	"retired",
	"error",
	"pending",
}

func (v *TagRequestStatuses) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TagRequestStatuses(value)
	for _, existing := range AllowedTagRequestStatusesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TagRequestStatuses", value)
}

// NewTagRequestStatusesFromValue returns a pointer to a valid TagRequestStatuses
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTagRequestStatusesFromValue(v string) (*TagRequestStatuses, error) {
	ev := TagRequestStatuses(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TagRequestStatuses: valid values are %v", v, AllowedTagRequestStatusesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TagRequestStatuses) IsValid() bool {
	for _, existing := range AllowedTagRequestStatusesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TagRequestStatuses value
func (v TagRequestStatuses) Ptr() *TagRequestStatuses {
	return &v
}

type NullableTagRequestStatuses struct {
	value *TagRequestStatuses
	isSet bool
}

func (v NullableTagRequestStatuses) Get() *TagRequestStatuses {
	return v.value
}

func (v *NullableTagRequestStatuses) Set(val *TagRequestStatuses) {
	v.value = val
	v.isSet = true
}

func (v NullableTagRequestStatuses) IsSet() bool {
	return v.isSet
}

func (v *NullableTagRequestStatuses) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagRequestStatuses(val *TagRequestStatuses) *NullableTagRequestStatuses {
	return &NullableTagRequestStatuses{value: val, isSet: true}
}

func (v NullableTagRequestStatuses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagRequestStatuses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
