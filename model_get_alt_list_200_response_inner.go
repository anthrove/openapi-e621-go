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

// GetAltList200ResponseInner - struct for GetAltList200ResponseInner
type GetAltList200ResponseInner struct {
	ArrayOfFloat32 *[]float32
	Float32        *float32
}

// []float32AsGetAltList200ResponseInner is a convenience function that returns []float32 wrapped in GetAltList200ResponseInner
func ArrayOfFloat32AsGetAltList200ResponseInner(v *[]float32) GetAltList200ResponseInner {
	return GetAltList200ResponseInner{
		ArrayOfFloat32: v,
	}
}

// float32AsGetAltList200ResponseInner is a convenience function that returns float32 wrapped in GetAltList200ResponseInner
func Float32AsGetAltList200ResponseInner(v *float32) GetAltList200ResponseInner {
	return GetAltList200ResponseInner{
		Float32: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetAltList200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ArrayOfFloat32
	err = newStrictDecoder(data).Decode(&dst.ArrayOfFloat32)
	if err == nil {
		jsonArrayOfFloat32, _ := json.Marshal(dst.ArrayOfFloat32)
		if string(jsonArrayOfFloat32) == "{}" { // empty struct
			dst.ArrayOfFloat32 = nil
		} else {
			if err = validator.Validate(dst.ArrayOfFloat32); err != nil {
				dst.ArrayOfFloat32 = nil
			} else {
				match++
			}
		}
	} else {
		dst.ArrayOfFloat32 = nil
	}

	// try to unmarshal data into Float32
	err = newStrictDecoder(data).Decode(&dst.Float32)
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			if err = validator.Validate(dst.Float32); err != nil {
				dst.Float32 = nil
			} else {
				match++
			}
		}
	} else {
		dst.Float32 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ArrayOfFloat32 = nil
		dst.Float32 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetAltList200ResponseInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetAltList200ResponseInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetAltList200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.ArrayOfFloat32 != nil {
		return json.Marshal(&src.ArrayOfFloat32)
	}

	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetAltList200ResponseInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ArrayOfFloat32 != nil {
		return obj.ArrayOfFloat32
	}

	if obj.Float32 != nil {
		return obj.Float32
	}

	// all schemas are nil
	return nil
}

type NullableGetAltList200ResponseInner struct {
	value *GetAltList200ResponseInner
	isSet bool
}

func (v NullableGetAltList200ResponseInner) Get() *GetAltList200ResponseInner {
	return v.value
}

func (v *NullableGetAltList200ResponseInner) Set(val *GetAltList200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAltList200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAltList200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAltList200ResponseInner(val *GetAltList200ResponseInner) *NullableGetAltList200ResponseInner {
	return &NullableGetAltList200ResponseInner{value: val, isSet: true}
}

func (v NullableGetAltList200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAltList200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
