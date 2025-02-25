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

// SearchUserFeedbacks200Response struct for SearchUserFeedbacks200Response
type SearchUserFeedbacks200Response struct {
	SearchUserFeedbacks200ResponseAnyOf *SearchUserFeedbacks200ResponseAnyOf
	ArrayOfUserFeedback                 *[]UserFeedback
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SearchUserFeedbacks200Response) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SearchUserFeedbacks200ResponseAnyOf
	err = json.Unmarshal(data, &dst.SearchUserFeedbacks200ResponseAnyOf)
	if err == nil {
		jsonSearchUserFeedbacks200ResponseAnyOf, _ := json.Marshal(dst.SearchUserFeedbacks200ResponseAnyOf)
		if string(jsonSearchUserFeedbacks200ResponseAnyOf) == "{}" { // empty struct
			dst.SearchUserFeedbacks200ResponseAnyOf = nil
		} else {
			return nil // data stored in dst.SearchUserFeedbacks200ResponseAnyOf, return on the first match
		}
	} else {
		dst.SearchUserFeedbacks200ResponseAnyOf = nil
	}

	// try to unmarshal JSON data into ArrayOfUserFeedback
	err = json.Unmarshal(data, &dst.ArrayOfUserFeedback)
	if err == nil {
		jsonArrayOfUserFeedback, _ := json.Marshal(dst.ArrayOfUserFeedback)
		if string(jsonArrayOfUserFeedback) == "{}" { // empty struct
			dst.ArrayOfUserFeedback = nil
		} else {
			return nil // data stored in dst.ArrayOfUserFeedback, return on the first match
		}
	} else {
		dst.ArrayOfUserFeedback = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SearchUserFeedbacks200Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SearchUserFeedbacks200Response) MarshalJSON() ([]byte, error) {
	if src.SearchUserFeedbacks200ResponseAnyOf != nil {
		return json.Marshal(&src.SearchUserFeedbacks200ResponseAnyOf)
	}

	if src.ArrayOfUserFeedback != nil {
		return json.Marshal(&src.ArrayOfUserFeedback)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSearchUserFeedbacks200Response struct {
	value *SearchUserFeedbacks200Response
	isSet bool
}

func (v NullableSearchUserFeedbacks200Response) Get() *SearchUserFeedbacks200Response {
	return v.value
}

func (v *NullableSearchUserFeedbacks200Response) Set(val *SearchUserFeedbacks200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchUserFeedbacks200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchUserFeedbacks200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchUserFeedbacks200Response(val *SearchUserFeedbacks200Response) *NullableSearchUserFeedbacks200Response {
	return &NullableSearchUserFeedbacks200Response{value: val, isSet: true}
}

func (v NullableSearchUserFeedbacks200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchUserFeedbacks200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
