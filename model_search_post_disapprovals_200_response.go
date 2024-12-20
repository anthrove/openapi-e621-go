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

// SearchPostDisapprovals200Response struct for SearchPostDisapprovals200Response
type SearchPostDisapprovals200Response struct {
	SearchPostDisapprovals200ResponseAnyOf *SearchPostDisapprovals200ResponseAnyOf
	ArrayOfPostDisapproval                 *[]PostDisapproval
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SearchPostDisapprovals200Response) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SearchPostDisapprovals200ResponseAnyOf
	err = json.Unmarshal(data, &dst.SearchPostDisapprovals200ResponseAnyOf)
	if err == nil {
		jsonSearchPostDisapprovals200ResponseAnyOf, _ := json.Marshal(dst.SearchPostDisapprovals200ResponseAnyOf)
		if string(jsonSearchPostDisapprovals200ResponseAnyOf) == "{}" { // empty struct
			dst.SearchPostDisapprovals200ResponseAnyOf = nil
		} else {
			return nil // data stored in dst.SearchPostDisapprovals200ResponseAnyOf, return on the first match
		}
	} else {
		dst.SearchPostDisapprovals200ResponseAnyOf = nil
	}

	// try to unmarshal JSON data into ArrayOfPostDisapproval
	err = json.Unmarshal(data, &dst.ArrayOfPostDisapproval)
	if err == nil {
		jsonArrayOfPostDisapproval, _ := json.Marshal(dst.ArrayOfPostDisapproval)
		if string(jsonArrayOfPostDisapproval) == "{}" { // empty struct
			dst.ArrayOfPostDisapproval = nil
		} else {
			return nil // data stored in dst.ArrayOfPostDisapproval, return on the first match
		}
	} else {
		dst.ArrayOfPostDisapproval = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SearchPostDisapprovals200Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SearchPostDisapprovals200Response) MarshalJSON() ([]byte, error) {
	if src.SearchPostDisapprovals200ResponseAnyOf != nil {
		return json.Marshal(&src.SearchPostDisapprovals200ResponseAnyOf)
	}

	if src.ArrayOfPostDisapproval != nil {
		return json.Marshal(&src.ArrayOfPostDisapproval)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSearchPostDisapprovals200Response struct {
	value *SearchPostDisapprovals200Response
	isSet bool
}

func (v NullableSearchPostDisapprovals200Response) Get() *SearchPostDisapprovals200Response {
	return v.value
}

func (v *NullableSearchPostDisapprovals200Response) Set(val *SearchPostDisapprovals200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchPostDisapprovals200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchPostDisapprovals200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchPostDisapprovals200Response(val *SearchPostDisapprovals200Response) *NullableSearchPostDisapprovals200Response {
	return &NullableSearchPostDisapprovals200Response{value: val, isSet: true}
}

func (v NullableSearchPostDisapprovals200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchPostDisapprovals200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
