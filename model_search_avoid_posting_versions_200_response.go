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

// SearchAvoidPostingVersions200Response struct for SearchAvoidPostingVersions200Response
type SearchAvoidPostingVersions200Response struct {
	SearchAvoidPostingVersions200ResponseAnyOf *SearchAvoidPostingVersions200ResponseAnyOf
	ArrayOfAvoidPostingVersion *[]AvoidPostingVersion
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SearchAvoidPostingVersions200Response) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SearchAvoidPostingVersions200ResponseAnyOf
	err = json.Unmarshal(data, &dst.SearchAvoidPostingVersions200ResponseAnyOf);
	if err == nil {
		jsonSearchAvoidPostingVersions200ResponseAnyOf, _ := json.Marshal(dst.SearchAvoidPostingVersions200ResponseAnyOf)
		if string(jsonSearchAvoidPostingVersions200ResponseAnyOf) == "{}" { // empty struct
			dst.SearchAvoidPostingVersions200ResponseAnyOf = nil
		} else {
			return nil // data stored in dst.SearchAvoidPostingVersions200ResponseAnyOf, return on the first match
		}
	} else {
		dst.SearchAvoidPostingVersions200ResponseAnyOf = nil
	}

	// try to unmarshal JSON data into ArrayOfAvoidPostingVersion
	err = json.Unmarshal(data, &dst.ArrayOfAvoidPostingVersion);
	if err == nil {
		jsonArrayOfAvoidPostingVersion, _ := json.Marshal(dst.ArrayOfAvoidPostingVersion)
		if string(jsonArrayOfAvoidPostingVersion) == "{}" { // empty struct
			dst.ArrayOfAvoidPostingVersion = nil
		} else {
			return nil // data stored in dst.ArrayOfAvoidPostingVersion, return on the first match
		}
	} else {
		dst.ArrayOfAvoidPostingVersion = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SearchAvoidPostingVersions200Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SearchAvoidPostingVersions200Response) MarshalJSON() ([]byte, error) {
	if src.SearchAvoidPostingVersions200ResponseAnyOf != nil {
		return json.Marshal(&src.SearchAvoidPostingVersions200ResponseAnyOf)
	}

	if src.ArrayOfAvoidPostingVersion != nil {
		return json.Marshal(&src.ArrayOfAvoidPostingVersion)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSearchAvoidPostingVersions200Response struct {
	value *SearchAvoidPostingVersions200Response
	isSet bool
}

func (v NullableSearchAvoidPostingVersions200Response) Get() *SearchAvoidPostingVersions200Response {
	return v.value
}

func (v *NullableSearchAvoidPostingVersions200Response) Set(val *SearchAvoidPostingVersions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchAvoidPostingVersions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchAvoidPostingVersions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchAvoidPostingVersions200Response(val *SearchAvoidPostingVersions200Response) *NullableSearchAvoidPostingVersions200Response {
	return &NullableSearchAvoidPostingVersions200Response{value: val, isSet: true}
}

func (v NullableSearchAvoidPostingVersions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchAvoidPostingVersions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


