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
	"gopkg.in/validator.v2"
)

// GetUser200Response - struct for GetUser200Response
type GetUser200Response struct {
	FullCurrentUser *FullCurrentUser
	FullUser        *FullUser
}

// FullCurrentUserAsGetUser200Response is a convenience function that returns FullCurrentUser wrapped in GetUser200Response
func FullCurrentUserAsGetUser200Response(v *FullCurrentUser) GetUser200Response {
	return GetUser200Response{
		FullCurrentUser: v,
	}
}

// FullUserAsGetUser200Response is a convenience function that returns FullUser wrapped in GetUser200Response
func FullUserAsGetUser200Response(v *FullUser) GetUser200Response {
	return GetUser200Response{
		FullUser: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetUser200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FullCurrentUser
	err = newStrictDecoder(data).Decode(&dst.FullCurrentUser)
	if err == nil {
		jsonFullCurrentUser, _ := json.Marshal(dst.FullCurrentUser)
		if string(jsonFullCurrentUser) == "{}" { // empty struct
			dst.FullCurrentUser = nil
		} else {
			if err = validator.Validate(dst.FullCurrentUser); err != nil {
				dst.FullCurrentUser = nil
			} else {
				match++
			}
		}
	} else {
		dst.FullCurrentUser = nil
	}

	// try to unmarshal data into FullUser
	err = newStrictDecoder(data).Decode(&dst.FullUser)
	if err == nil {
		jsonFullUser, _ := json.Marshal(dst.FullUser)
		if string(jsonFullUser) == "{}" { // empty struct
			dst.FullUser = nil
		} else {
			if err = validator.Validate(dst.FullUser); err != nil {
				dst.FullUser = nil
			} else {
				match++
			}
		}
	} else {
		dst.FullUser = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.FullCurrentUser = nil
		dst.FullUser = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetUser200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetUser200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetUser200Response) MarshalJSON() ([]byte, error) {
	if src.FullCurrentUser != nil {
		return json.Marshal(&src.FullCurrentUser)
	}

	if src.FullUser != nil {
		return json.Marshal(&src.FullUser)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetUser200Response) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.FullCurrentUser != nil {
		return obj.FullCurrentUser
	}

	if obj.FullUser != nil {
		return obj.FullUser
	}

	// all schemas are nil
	return nil
}

type NullableGetUser200Response struct {
	value *GetUser200Response
	isSet bool
}

func (v NullableGetUser200Response) Get() *GetUser200Response {
	return v.value
}

func (v *NullableGetUser200Response) Set(val *GetUser200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUser200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUser200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUser200Response(val *GetUser200Response) *NullableGetUser200Response {
	return &NullableGetUser200Response{value: val, isSet: true}
}

func (v NullableGetUser200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUser200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
