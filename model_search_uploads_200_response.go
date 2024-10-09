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

// SearchUploads200Response struct for SearchUploads200Response
type SearchUploads200Response struct {
	SearchUploads200ResponseAnyOf *SearchUploads200ResponseAnyOf
	ArrayOfUpload                 *[]Upload
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SearchUploads200Response) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SearchUploads200ResponseAnyOf
	err = json.Unmarshal(data, &dst.SearchUploads200ResponseAnyOf)
	if err == nil {
		jsonSearchUploads200ResponseAnyOf, _ := json.Marshal(dst.SearchUploads200ResponseAnyOf)
		if string(jsonSearchUploads200ResponseAnyOf) == "{}" { // empty struct
			dst.SearchUploads200ResponseAnyOf = nil
		} else {
			return nil // data stored in dst.SearchUploads200ResponseAnyOf, return on the first match
		}
	} else {
		dst.SearchUploads200ResponseAnyOf = nil
	}

	// try to unmarshal JSON data into ArrayOfUpload
	err = json.Unmarshal(data, &dst.ArrayOfUpload)
	if err == nil {
		jsonArrayOfUpload, _ := json.Marshal(dst.ArrayOfUpload)
		if string(jsonArrayOfUpload) == "{}" { // empty struct
			dst.ArrayOfUpload = nil
		} else {
			return nil // data stored in dst.ArrayOfUpload, return on the first match
		}
	} else {
		dst.ArrayOfUpload = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SearchUploads200Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SearchUploads200Response) MarshalJSON() ([]byte, error) {
	if src.SearchUploads200ResponseAnyOf != nil {
		return json.Marshal(&src.SearchUploads200ResponseAnyOf)
	}

	if src.ArrayOfUpload != nil {
		return json.Marshal(&src.ArrayOfUpload)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSearchUploads200Response struct {
	value *SearchUploads200Response
	isSet bool
}

func (v NullableSearchUploads200Response) Get() *SearchUploads200Response {
	return v.value
}

func (v *NullableSearchUploads200Response) Set(val *SearchUploads200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchUploads200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchUploads200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchUploads200Response(val *SearchUploads200Response) *NullableSearchUploads200Response {
	return &NullableSearchUploads200Response{value: val, isSet: true}
}

func (v NullableSearchUploads200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchUploads200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}