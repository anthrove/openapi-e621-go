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

// SearchNoteVersions200Response struct for SearchNoteVersions200Response
type SearchNoteVersions200Response struct {
	SearchNoteVersions200ResponseAnyOf *SearchNoteVersions200ResponseAnyOf
	ArrayOfNoteVersion                 *[]NoteVersion
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SearchNoteVersions200Response) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SearchNoteVersions200ResponseAnyOf
	err = json.Unmarshal(data, &dst.SearchNoteVersions200ResponseAnyOf)
	if err == nil {
		jsonSearchNoteVersions200ResponseAnyOf, _ := json.Marshal(dst.SearchNoteVersions200ResponseAnyOf)
		if string(jsonSearchNoteVersions200ResponseAnyOf) == "{}" { // empty struct
			dst.SearchNoteVersions200ResponseAnyOf = nil
		} else {
			return nil // data stored in dst.SearchNoteVersions200ResponseAnyOf, return on the first match
		}
	} else {
		dst.SearchNoteVersions200ResponseAnyOf = nil
	}

	// try to unmarshal JSON data into ArrayOfNoteVersion
	err = json.Unmarshal(data, &dst.ArrayOfNoteVersion)
	if err == nil {
		jsonArrayOfNoteVersion, _ := json.Marshal(dst.ArrayOfNoteVersion)
		if string(jsonArrayOfNoteVersion) == "{}" { // empty struct
			dst.ArrayOfNoteVersion = nil
		} else {
			return nil // data stored in dst.ArrayOfNoteVersion, return on the first match
		}
	} else {
		dst.ArrayOfNoteVersion = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SearchNoteVersions200Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SearchNoteVersions200Response) MarshalJSON() ([]byte, error) {
	if src.SearchNoteVersions200ResponseAnyOf != nil {
		return json.Marshal(&src.SearchNoteVersions200ResponseAnyOf)
	}

	if src.ArrayOfNoteVersion != nil {
		return json.Marshal(&src.ArrayOfNoteVersion)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSearchNoteVersions200Response struct {
	value *SearchNoteVersions200Response
	isSet bool
}

func (v NullableSearchNoteVersions200Response) Get() *SearchNoteVersions200Response {
	return v.value
}

func (v *NullableSearchNoteVersions200Response) Set(val *SearchNoteVersions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchNoteVersions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchNoteVersions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchNoteVersions200Response(val *SearchNoteVersions200Response) *NullableSearchNoteVersions200Response {
	return &NullableSearchNoteVersions200Response{value: val, isSet: true}
}

func (v NullableSearchNoteVersions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchNoteVersions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
