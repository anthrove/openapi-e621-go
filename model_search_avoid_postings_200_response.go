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

// SearchAvoidPostings200Response struct for SearchAvoidPostings200Response
type SearchAvoidPostings200Response struct {
	SearchAvoidPostings200ResponseAnyOf *SearchAvoidPostings200ResponseAnyOf
	ArrayOfAvoidPosting                 *[]AvoidPosting
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SearchAvoidPostings200Response) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SearchAvoidPostings200ResponseAnyOf
	err = json.Unmarshal(data, &dst.SearchAvoidPostings200ResponseAnyOf)
	if err == nil {
		jsonSearchAvoidPostings200ResponseAnyOf, _ := json.Marshal(dst.SearchAvoidPostings200ResponseAnyOf)
		if string(jsonSearchAvoidPostings200ResponseAnyOf) == "{}" { // empty struct
			dst.SearchAvoidPostings200ResponseAnyOf = nil
		} else {
			return nil // data stored in dst.SearchAvoidPostings200ResponseAnyOf, return on the first match
		}
	} else {
		dst.SearchAvoidPostings200ResponseAnyOf = nil
	}

	// try to unmarshal JSON data into ArrayOfAvoidPosting
	err = json.Unmarshal(data, &dst.ArrayOfAvoidPosting)
	if err == nil {
		jsonArrayOfAvoidPosting, _ := json.Marshal(dst.ArrayOfAvoidPosting)
		if string(jsonArrayOfAvoidPosting) == "{}" { // empty struct
			dst.ArrayOfAvoidPosting = nil
		} else {
			return nil // data stored in dst.ArrayOfAvoidPosting, return on the first match
		}
	} else {
		dst.ArrayOfAvoidPosting = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SearchAvoidPostings200Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SearchAvoidPostings200Response) MarshalJSON() ([]byte, error) {
	if src.SearchAvoidPostings200ResponseAnyOf != nil {
		return json.Marshal(&src.SearchAvoidPostings200ResponseAnyOf)
	}

	if src.ArrayOfAvoidPosting != nil {
		return json.Marshal(&src.ArrayOfAvoidPosting)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSearchAvoidPostings200Response struct {
	value *SearchAvoidPostings200Response
	isSet bool
}

func (v NullableSearchAvoidPostings200Response) Get() *SearchAvoidPostings200Response {
	return v.value
}

func (v *NullableSearchAvoidPostings200Response) Set(val *SearchAvoidPostings200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchAvoidPostings200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchAvoidPostings200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchAvoidPostings200Response(val *SearchAvoidPostings200Response) *NullableSearchAvoidPostings200Response {
	return &NullableSearchAvoidPostings200Response{value: val, isSet: true}
}

func (v NullableSearchAvoidPostings200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchAvoidPostings200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
