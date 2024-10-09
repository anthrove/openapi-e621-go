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

// ListIPBans200Response struct for ListIPBans200Response
type ListIPBans200Response struct {
	ListIPBans200ResponseAnyOf *ListIPBans200ResponseAnyOf
	ArrayOfIPBan               *[]IPBan
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ListIPBans200Response) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ListIPBans200ResponseAnyOf
	err = json.Unmarshal(data, &dst.ListIPBans200ResponseAnyOf)
	if err == nil {
		jsonListIPBans200ResponseAnyOf, _ := json.Marshal(dst.ListIPBans200ResponseAnyOf)
		if string(jsonListIPBans200ResponseAnyOf) == "{}" { // empty struct
			dst.ListIPBans200ResponseAnyOf = nil
		} else {
			return nil // data stored in dst.ListIPBans200ResponseAnyOf, return on the first match
		}
	} else {
		dst.ListIPBans200ResponseAnyOf = nil
	}

	// try to unmarshal JSON data into ArrayOfIPBan
	err = json.Unmarshal(data, &dst.ArrayOfIPBan)
	if err == nil {
		jsonArrayOfIPBan, _ := json.Marshal(dst.ArrayOfIPBan)
		if string(jsonArrayOfIPBan) == "{}" { // empty struct
			dst.ArrayOfIPBan = nil
		} else {
			return nil // data stored in dst.ArrayOfIPBan, return on the first match
		}
	} else {
		dst.ArrayOfIPBan = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ListIPBans200Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ListIPBans200Response) MarshalJSON() ([]byte, error) {
	if src.ListIPBans200ResponseAnyOf != nil {
		return json.Marshal(&src.ListIPBans200ResponseAnyOf)
	}

	if src.ArrayOfIPBan != nil {
		return json.Marshal(&src.ArrayOfIPBan)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableListIPBans200Response struct {
	value *ListIPBans200Response
	isSet bool
}

func (v NullableListIPBans200Response) Get() *ListIPBans200Response {
	return v.value
}

func (v *NullableListIPBans200Response) Set(val *ListIPBans200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListIPBans200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListIPBans200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListIPBans200Response(val *ListIPBans200Response) *NullableListIPBans200Response {
	return &NullableListIPBans200Response{value: val, isSet: true}
}

func (v NullableListIPBans200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListIPBans200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
