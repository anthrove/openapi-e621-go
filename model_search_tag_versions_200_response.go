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

// SearchTagVersions200Response struct for SearchTagVersions200Response
type SearchTagVersions200Response struct {
	SearchTagVersions200ResponseAnyOf *SearchTagVersions200ResponseAnyOf
	ArrayOfTagTypeVersion             *[]TagTypeVersion
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SearchTagVersions200Response) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SearchTagVersions200ResponseAnyOf
	err = json.Unmarshal(data, &dst.SearchTagVersions200ResponseAnyOf)
	if err == nil {
		jsonSearchTagVersions200ResponseAnyOf, _ := json.Marshal(dst.SearchTagVersions200ResponseAnyOf)
		if string(jsonSearchTagVersions200ResponseAnyOf) == "{}" { // empty struct
			dst.SearchTagVersions200ResponseAnyOf = nil
		} else {
			return nil // data stored in dst.SearchTagVersions200ResponseAnyOf, return on the first match
		}
	} else {
		dst.SearchTagVersions200ResponseAnyOf = nil
	}

	// try to unmarshal JSON data into ArrayOfTagTypeVersion
	err = json.Unmarshal(data, &dst.ArrayOfTagTypeVersion)
	if err == nil {
		jsonArrayOfTagTypeVersion, _ := json.Marshal(dst.ArrayOfTagTypeVersion)
		if string(jsonArrayOfTagTypeVersion) == "{}" { // empty struct
			dst.ArrayOfTagTypeVersion = nil
		} else {
			return nil // data stored in dst.ArrayOfTagTypeVersion, return on the first match
		}
	} else {
		dst.ArrayOfTagTypeVersion = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SearchTagVersions200Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SearchTagVersions200Response) MarshalJSON() ([]byte, error) {
	if src.SearchTagVersions200ResponseAnyOf != nil {
		return json.Marshal(&src.SearchTagVersions200ResponseAnyOf)
	}

	if src.ArrayOfTagTypeVersion != nil {
		return json.Marshal(&src.ArrayOfTagTypeVersion)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSearchTagVersions200Response struct {
	value *SearchTagVersions200Response
	isSet bool
}

func (v NullableSearchTagVersions200Response) Get() *SearchTagVersions200Response {
	return v.value
}

func (v *NullableSearchTagVersions200Response) Set(val *SearchTagVersions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchTagVersions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchTagVersions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchTagVersions200Response(val *SearchTagVersions200Response) *NullableSearchTagVersions200Response {
	return &NullableSearchTagVersions200Response{value: val, isSet: true}
}

func (v NullableSearchTagVersions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchTagVersions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
