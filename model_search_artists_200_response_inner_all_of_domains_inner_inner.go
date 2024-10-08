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

// SearchArtists200ResponseInnerAllOfDomainsInnerInner struct for SearchArtists200ResponseInnerAllOfDomainsInnerInner
type SearchArtists200ResponseInnerAllOfDomainsInnerInner struct {
	Float32 *float32
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SearchArtists200ResponseInnerAllOfDomainsInnerInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into Float32
	err = json.Unmarshal(data, &dst.Float32);
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			return nil // data stored in dst.Float32, return on the first match
		}
	} else {
		dst.Float32 = nil
	}

	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SearchArtists200ResponseInnerAllOfDomainsInnerInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SearchArtists200ResponseInnerAllOfDomainsInnerInner) MarshalJSON() ([]byte, error) {
	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSearchArtists200ResponseInnerAllOfDomainsInnerInner struct {
	value *SearchArtists200ResponseInnerAllOfDomainsInnerInner
	isSet bool
}

func (v NullableSearchArtists200ResponseInnerAllOfDomainsInnerInner) Get() *SearchArtists200ResponseInnerAllOfDomainsInnerInner {
	return v.value
}

func (v *NullableSearchArtists200ResponseInnerAllOfDomainsInnerInner) Set(val *SearchArtists200ResponseInnerAllOfDomainsInnerInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchArtists200ResponseInnerAllOfDomainsInnerInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchArtists200ResponseInnerAllOfDomainsInnerInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchArtists200ResponseInnerAllOfDomainsInnerInner(val *SearchArtists200ResponseInnerAllOfDomainsInnerInner) *NullableSearchArtists200ResponseInnerAllOfDomainsInnerInner {
	return &NullableSearchArtists200ResponseInnerAllOfDomainsInnerInner{value: val, isSet: true}
}

func (v NullableSearchArtists200ResponseInnerAllOfDomainsInnerInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchArtists200ResponseInnerAllOfDomainsInnerInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


