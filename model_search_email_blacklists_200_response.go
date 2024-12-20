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

// SearchEmailBlacklists200Response struct for SearchEmailBlacklists200Response
type SearchEmailBlacklists200Response struct {
	SearchEmailBlacklists200ResponseAnyOf *SearchEmailBlacklists200ResponseAnyOf
	ArrayOfEmailBlacklist                 *[]EmailBlacklist
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SearchEmailBlacklists200Response) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SearchEmailBlacklists200ResponseAnyOf
	err = json.Unmarshal(data, &dst.SearchEmailBlacklists200ResponseAnyOf)
	if err == nil {
		jsonSearchEmailBlacklists200ResponseAnyOf, _ := json.Marshal(dst.SearchEmailBlacklists200ResponseAnyOf)
		if string(jsonSearchEmailBlacklists200ResponseAnyOf) == "{}" { // empty struct
			dst.SearchEmailBlacklists200ResponseAnyOf = nil
		} else {
			return nil // data stored in dst.SearchEmailBlacklists200ResponseAnyOf, return on the first match
		}
	} else {
		dst.SearchEmailBlacklists200ResponseAnyOf = nil
	}

	// try to unmarshal JSON data into ArrayOfEmailBlacklist
	err = json.Unmarshal(data, &dst.ArrayOfEmailBlacklist)
	if err == nil {
		jsonArrayOfEmailBlacklist, _ := json.Marshal(dst.ArrayOfEmailBlacklist)
		if string(jsonArrayOfEmailBlacklist) == "{}" { // empty struct
			dst.ArrayOfEmailBlacklist = nil
		} else {
			return nil // data stored in dst.ArrayOfEmailBlacklist, return on the first match
		}
	} else {
		dst.ArrayOfEmailBlacklist = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SearchEmailBlacklists200Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SearchEmailBlacklists200Response) MarshalJSON() ([]byte, error) {
	if src.SearchEmailBlacklists200ResponseAnyOf != nil {
		return json.Marshal(&src.SearchEmailBlacklists200ResponseAnyOf)
	}

	if src.ArrayOfEmailBlacklist != nil {
		return json.Marshal(&src.ArrayOfEmailBlacklist)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSearchEmailBlacklists200Response struct {
	value *SearchEmailBlacklists200Response
	isSet bool
}

func (v NullableSearchEmailBlacklists200Response) Get() *SearchEmailBlacklists200Response {
	return v.value
}

func (v *NullableSearchEmailBlacklists200Response) Set(val *SearchEmailBlacklists200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchEmailBlacklists200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchEmailBlacklists200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchEmailBlacklists200Response(val *SearchEmailBlacklists200Response) *NullableSearchEmailBlacklists200Response {
	return &NullableSearchEmailBlacklists200Response{value: val, isSet: true}
}

func (v NullableSearchEmailBlacklists200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchEmailBlacklists200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
