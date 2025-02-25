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

// TagCategories the model 'TagCategories'
type TagCategories int32

// List of TagCategories
const (
	TAGCATEGORIES__0 TagCategories = 0
	TAGCATEGORIES__1 TagCategories = 1
	TAGCATEGORIES__3 TagCategories = 3
	TAGCATEGORIES__4 TagCategories = 4
	TAGCATEGORIES__5 TagCategories = 5
	TAGCATEGORIES__6 TagCategories = 6
	TAGCATEGORIES__7 TagCategories = 7
	TAGCATEGORIES__8 TagCategories = 8
)

// All allowed values of TagCategories enum
var AllowedTagCategoriesEnumValues = []TagCategories{
	0,
	1,
	3,
	4,
	5,
	6,
	7,
	8,
}

func (v *TagCategories) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TagCategories(value)
	for _, existing := range AllowedTagCategoriesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TagCategories", value)
}

// NewTagCategoriesFromValue returns a pointer to a valid TagCategories
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTagCategoriesFromValue(v int32) (*TagCategories, error) {
	ev := TagCategories(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TagCategories: valid values are %v", v, AllowedTagCategoriesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TagCategories) IsValid() bool {
	for _, existing := range AllowedTagCategoriesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TagCategories value
func (v TagCategories) Ptr() *TagCategories {
	return &v
}

type NullableTagCategories struct {
	value *TagCategories
	isSet bool
}

func (v NullableTagCategories) Get() *TagCategories {
	return v.value
}

func (v *NullableTagCategories) Set(val *TagCategories) {
	v.value = val
	v.isSet = true
}

func (v NullableTagCategories) IsSet() bool {
	return v.isSet
}

func (v *NullableTagCategories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagCategories(val *TagCategories) *NullableTagCategories {
	return &NullableTagCategories{value: val, isSet: true}
}

func (v NullableTagCategories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagCategories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
