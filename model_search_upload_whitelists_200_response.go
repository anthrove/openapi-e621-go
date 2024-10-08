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

// SearchUploadWhitelists200Response struct for SearchUploadWhitelists200Response
type SearchUploadWhitelists200Response struct {
	SearchUploadWhitelists200ResponseAnyOf *SearchUploadWhitelists200ResponseAnyOf
	ArrayOfUploadWhitelist *[]UploadWhitelist
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SearchUploadWhitelists200Response) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SearchUploadWhitelists200ResponseAnyOf
	err = json.Unmarshal(data, &dst.SearchUploadWhitelists200ResponseAnyOf);
	if err == nil {
		jsonSearchUploadWhitelists200ResponseAnyOf, _ := json.Marshal(dst.SearchUploadWhitelists200ResponseAnyOf)
		if string(jsonSearchUploadWhitelists200ResponseAnyOf) == "{}" { // empty struct
			dst.SearchUploadWhitelists200ResponseAnyOf = nil
		} else {
			return nil // data stored in dst.SearchUploadWhitelists200ResponseAnyOf, return on the first match
		}
	} else {
		dst.SearchUploadWhitelists200ResponseAnyOf = nil
	}

	// try to unmarshal JSON data into ArrayOfUploadWhitelist
	err = json.Unmarshal(data, &dst.ArrayOfUploadWhitelist);
	if err == nil {
		jsonArrayOfUploadWhitelist, _ := json.Marshal(dst.ArrayOfUploadWhitelist)
		if string(jsonArrayOfUploadWhitelist) == "{}" { // empty struct
			dst.ArrayOfUploadWhitelist = nil
		} else {
			return nil // data stored in dst.ArrayOfUploadWhitelist, return on the first match
		}
	} else {
		dst.ArrayOfUploadWhitelist = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SearchUploadWhitelists200Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SearchUploadWhitelists200Response) MarshalJSON() ([]byte, error) {
	if src.SearchUploadWhitelists200ResponseAnyOf != nil {
		return json.Marshal(&src.SearchUploadWhitelists200ResponseAnyOf)
	}

	if src.ArrayOfUploadWhitelist != nil {
		return json.Marshal(&src.ArrayOfUploadWhitelist)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSearchUploadWhitelists200Response struct {
	value *SearchUploadWhitelists200Response
	isSet bool
}

func (v NullableSearchUploadWhitelists200Response) Get() *SearchUploadWhitelists200Response {
	return v.value
}

func (v *NullableSearchUploadWhitelists200Response) Set(val *SearchUploadWhitelists200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchUploadWhitelists200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchUploadWhitelists200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchUploadWhitelists200Response(val *SearchUploadWhitelists200Response) *NullableSearchUploadWhitelists200Response {
	return &NullableSearchUploadWhitelists200Response{value: val, isSet: true}
}

func (v NullableSearchUploadWhitelists200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchUploadWhitelists200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

