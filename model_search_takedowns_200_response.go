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

// SearchTakedowns200Response struct for SearchTakedowns200Response
type SearchTakedowns200Response struct {
	SearchTakedowns200ResponseAnyOf *SearchTakedowns200ResponseAnyOf
	ArrayOfTakedown                 *[]Takedown
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SearchTakedowns200Response) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SearchTakedowns200ResponseAnyOf
	err = json.Unmarshal(data, &dst.SearchTakedowns200ResponseAnyOf)
	if err == nil {
		jsonSearchTakedowns200ResponseAnyOf, _ := json.Marshal(dst.SearchTakedowns200ResponseAnyOf)
		if string(jsonSearchTakedowns200ResponseAnyOf) == "{}" { // empty struct
			dst.SearchTakedowns200ResponseAnyOf = nil
		} else {
			return nil // data stored in dst.SearchTakedowns200ResponseAnyOf, return on the first match
		}
	} else {
		dst.SearchTakedowns200ResponseAnyOf = nil
	}

	// try to unmarshal JSON data into ArrayOfTakedown
	err = json.Unmarshal(data, &dst.ArrayOfTakedown)
	if err == nil {
		jsonArrayOfTakedown, _ := json.Marshal(dst.ArrayOfTakedown)
		if string(jsonArrayOfTakedown) == "{}" { // empty struct
			dst.ArrayOfTakedown = nil
		} else {
			return nil // data stored in dst.ArrayOfTakedown, return on the first match
		}
	} else {
		dst.ArrayOfTakedown = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SearchTakedowns200Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SearchTakedowns200Response) MarshalJSON() ([]byte, error) {
	if src.SearchTakedowns200ResponseAnyOf != nil {
		return json.Marshal(&src.SearchTakedowns200ResponseAnyOf)
	}

	if src.ArrayOfTakedown != nil {
		return json.Marshal(&src.ArrayOfTakedown)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSearchTakedowns200Response struct {
	value *SearchTakedowns200Response
	isSet bool
}

func (v NullableSearchTakedowns200Response) Get() *SearchTakedowns200Response {
	return v.value
}

func (v *NullableSearchTakedowns200Response) Set(val *SearchTakedowns200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchTakedowns200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchTakedowns200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchTakedowns200Response(val *SearchTakedowns200Response) *NullableSearchTakedowns200Response {
	return &NullableSearchTakedowns200Response{value: val, isSet: true}
}

func (v NullableSearchTakedowns200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchTakedowns200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}