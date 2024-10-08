/*
E621

OpenAPI definition for E621's API. You can find the source [here](https://github.com/DonovanDMC/E621OpenAPI)<br> This document is intended to compliment E621's existing [API Documentation](https://e621.net/help/api).<br> Note if E621's api is under attack and/or cloudflare protections are enabled, the \"Try it out\" buttons here will not work.<br> If they are not working, you can check this [Unofficial Status Page](https://status.e621.ws). 

API version: d69c34e
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi-e621-go

import (
	"encoding/json"
	"fmt"
)

// SearchForumPosts200Response struct for SearchForumPosts200Response
type SearchForumPosts200Response struct {
	SearchForumPosts200ResponseAnyOf *SearchForumPosts200ResponseAnyOf
	ArrayOfForumPost *[]ForumPost
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SearchForumPosts200Response) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SearchForumPosts200ResponseAnyOf
	err = json.Unmarshal(data, &dst.SearchForumPosts200ResponseAnyOf);
	if err == nil {
		jsonSearchForumPosts200ResponseAnyOf, _ := json.Marshal(dst.SearchForumPosts200ResponseAnyOf)
		if string(jsonSearchForumPosts200ResponseAnyOf) == "{}" { // empty struct
			dst.SearchForumPosts200ResponseAnyOf = nil
		} else {
			return nil // data stored in dst.SearchForumPosts200ResponseAnyOf, return on the first match
		}
	} else {
		dst.SearchForumPosts200ResponseAnyOf = nil
	}

	// try to unmarshal JSON data into ArrayOfForumPost
	err = json.Unmarshal(data, &dst.ArrayOfForumPost);
	if err == nil {
		jsonArrayOfForumPost, _ := json.Marshal(dst.ArrayOfForumPost)
		if string(jsonArrayOfForumPost) == "{}" { // empty struct
			dst.ArrayOfForumPost = nil
		} else {
			return nil // data stored in dst.ArrayOfForumPost, return on the first match
		}
	} else {
		dst.ArrayOfForumPost = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SearchForumPosts200Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SearchForumPosts200Response) MarshalJSON() ([]byte, error) {
	if src.SearchForumPosts200ResponseAnyOf != nil {
		return json.Marshal(&src.SearchForumPosts200ResponseAnyOf)
	}

	if src.ArrayOfForumPost != nil {
		return json.Marshal(&src.ArrayOfForumPost)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSearchForumPosts200Response struct {
	value *SearchForumPosts200Response
	isSet bool
}

func (v NullableSearchForumPosts200Response) Get() *SearchForumPosts200Response {
	return v.value
}

func (v *NullableSearchForumPosts200Response) Set(val *SearchForumPosts200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchForumPosts200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchForumPosts200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchForumPosts200Response(val *SearchForumPosts200Response) *NullableSearchForumPosts200Response {
	return &NullableSearchForumPosts200Response{value: val, isSet: true}
}

func (v NullableSearchForumPosts200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchForumPosts200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


