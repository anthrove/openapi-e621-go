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

// GetUser200Response struct for GetUser200Response
type GetUser200Response struct {
	FullCurrentUser *FullCurrentUser
	User *User
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *GetUser200Response) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into FullCurrentUser
	err = json.Unmarshal(data, &dst.FullCurrentUser);
	if err == nil {
		jsonFullCurrentUser, _ := json.Marshal(dst.FullCurrentUser)
		if string(jsonFullCurrentUser) == "{}" { // empty struct
			dst.FullCurrentUser = nil
		} else {
			return nil // data stored in dst.FullCurrentUser, return on the first match
		}
	} else {
		dst.FullCurrentUser = nil
	}

	// try to unmarshal JSON data into User
	err = json.Unmarshal(data, &dst.User);
	if err == nil {
		jsonUser, _ := json.Marshal(dst.User)
		if string(jsonUser) == "{}" { // empty struct
			dst.User = nil
		} else {
			return nil // data stored in dst.User, return on the first match
		}
	} else {
		dst.User = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(GetUser200Response)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *GetUser200Response) MarshalJSON() ([]byte, error) {
	if src.FullCurrentUser != nil {
		return json.Marshal(&src.FullCurrentUser)
	}

	if src.User != nil {
		return json.Marshal(&src.User)
	}

	return nil, nil // no data in anyOf schemas
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


