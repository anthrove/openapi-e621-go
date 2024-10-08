/*
E621

OpenAPI definition for E621's API. You can find the source [here](https://github.com/DonovanDMC/E621OpenAPI)<br> This document is intended to compliment E621's existing [API Documentation](https://e621.net/help/api).<br> Note if E621's api is under attack and/or cloudflare protections are enabled, the \"Try it out\" buttons here will not work.<br> If they are not working, you can check this [Unofficial Status Page](https://status.e621.ws). 

API version: d69c34e
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi-e621-go

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ListIPBans200ResponseAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListIPBans200ResponseAnyOf{}

// ListIPBans200ResponseAnyOf No Results
type ListIPBans200ResponseAnyOf struct {
	IpBans []string `json:"ip_bans"`
}

type _ListIPBans200ResponseAnyOf ListIPBans200ResponseAnyOf

// NewListIPBans200ResponseAnyOf instantiates a new ListIPBans200ResponseAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListIPBans200ResponseAnyOf(ipBans []string) *ListIPBans200ResponseAnyOf {
	this := ListIPBans200ResponseAnyOf{}
	this.IpBans = ipBans
	return &this
}

// NewListIPBans200ResponseAnyOfWithDefaults instantiates a new ListIPBans200ResponseAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListIPBans200ResponseAnyOfWithDefaults() *ListIPBans200ResponseAnyOf {
	this := ListIPBans200ResponseAnyOf{}
	return &this
}

// GetIpBans returns the IpBans field value
func (o *ListIPBans200ResponseAnyOf) GetIpBans() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.IpBans
}

// GetIpBansOk returns a tuple with the IpBans field value
// and a boolean to check if the value has been set.
func (o *ListIPBans200ResponseAnyOf) GetIpBansOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IpBans, true
}

// SetIpBans sets field value
func (o *ListIPBans200ResponseAnyOf) SetIpBans(v []string) {
	o.IpBans = v
}

func (o ListIPBans200ResponseAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListIPBans200ResponseAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ip_bans"] = o.IpBans
	return toSerialize, nil
}

func (o *ListIPBans200ResponseAnyOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ip_bans",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varListIPBans200ResponseAnyOf := _ListIPBans200ResponseAnyOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListIPBans200ResponseAnyOf)

	if err != nil {
		return err
	}

	*o = ListIPBans200ResponseAnyOf(varListIPBans200ResponseAnyOf)

	return err
}

type NullableListIPBans200ResponseAnyOf struct {
	value *ListIPBans200ResponseAnyOf
	isSet bool
}

func (v NullableListIPBans200ResponseAnyOf) Get() *ListIPBans200ResponseAnyOf {
	return v.value
}

func (v *NullableListIPBans200ResponseAnyOf) Set(val *ListIPBans200ResponseAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableListIPBans200ResponseAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableListIPBans200ResponseAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListIPBans200ResponseAnyOf(val *ListIPBans200ResponseAnyOf) *NullableListIPBans200ResponseAnyOf {
	return &NullableListIPBans200ResponseAnyOf{value: val, isSet: true}
}

func (v NullableListIPBans200ResponseAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListIPBans200ResponseAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


