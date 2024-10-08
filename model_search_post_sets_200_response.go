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

// SearchPostSets200Response struct for SearchPostSets200Response
type SearchPostSets200Response struct {
	SearchPostSets200ResponseAnyOf *SearchPostSets200ResponseAnyOf
	ArrayOfPostSet *[]PostSet
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SearchPostSets200Response) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SearchPostSets200ResponseAnyOf
	err = json.Unmarshal(data, &dst.SearchPostSets200ResponseAnyOf);
	if err == nil {
		jsonSearchPostSets200ResponseAnyOf, _ := json.Marshal(dst.SearchPostSets200ResponseAnyOf)
		if string(jsonSearchPostSets200ResponseAnyOf) == "{}" { // empty struct
			dst.SearchPostSets200ResponseAnyOf = nil
		} else {
			return nil // data stored in dst.SearchPostSets200ResponseAnyOf, return on the first match
		}
	} else {
		dst.SearchPostSets200ResponseAnyOf = nil
	}

	// try to unmarshal JSON data into ArrayOfPostSet
	err = json.Unmarshal(data, &dst.ArrayOfPostSet);
	if err == nil {
		jsonArrayOfPostSet, _ := json.Marshal(dst.ArrayOfPostSet)
		if string(jsonArrayOfPostSet) == "{}" { // empty struct
			dst.ArrayOfPostSet = nil
		} else {
			return nil // data stored in dst.ArrayOfPostSet, return on the first match
		}
	} else {
		dst.ArrayOfPostSet = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SearchPostSets200Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SearchPostSets200Response) MarshalJSON() ([]byte, error) {
	if src.SearchPostSets200ResponseAnyOf != nil {
		return json.Marshal(&src.SearchPostSets200ResponseAnyOf)
	}

	if src.ArrayOfPostSet != nil {
		return json.Marshal(&src.ArrayOfPostSet)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSearchPostSets200Response struct {
	value *SearchPostSets200Response
	isSet bool
}

func (v NullableSearchPostSets200Response) Get() *SearchPostSets200Response {
	return v.value
}

func (v *NullableSearchPostSets200Response) Set(val *SearchPostSets200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchPostSets200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchPostSets200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchPostSets200Response(val *SearchPostSets200Response) *NullableSearchPostSets200Response {
	return &NullableSearchPostSets200Response{value: val, isSet: true}
}

func (v NullableSearchPostSets200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchPostSets200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


