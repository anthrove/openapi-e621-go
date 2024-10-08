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

// checks if the SearchUserNameChangeRequests200ResponseAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchUserNameChangeRequests200ResponseAnyOf{}

// SearchUserNameChangeRequests200ResponseAnyOf No Results
type SearchUserNameChangeRequests200ResponseAnyOf struct {
	UserNameChangeRequests []string `json:"user_name_change_requests"`
}

type _SearchUserNameChangeRequests200ResponseAnyOf SearchUserNameChangeRequests200ResponseAnyOf

// NewSearchUserNameChangeRequests200ResponseAnyOf instantiates a new SearchUserNameChangeRequests200ResponseAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchUserNameChangeRequests200ResponseAnyOf(userNameChangeRequests []string) *SearchUserNameChangeRequests200ResponseAnyOf {
	this := SearchUserNameChangeRequests200ResponseAnyOf{}
	this.UserNameChangeRequests = userNameChangeRequests
	return &this
}

// NewSearchUserNameChangeRequests200ResponseAnyOfWithDefaults instantiates a new SearchUserNameChangeRequests200ResponseAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchUserNameChangeRequests200ResponseAnyOfWithDefaults() *SearchUserNameChangeRequests200ResponseAnyOf {
	this := SearchUserNameChangeRequests200ResponseAnyOf{}
	return &this
}

// GetUserNameChangeRequests returns the UserNameChangeRequests field value
func (o *SearchUserNameChangeRequests200ResponseAnyOf) GetUserNameChangeRequests() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.UserNameChangeRequests
}

// GetUserNameChangeRequestsOk returns a tuple with the UserNameChangeRequests field value
// and a boolean to check if the value has been set.
func (o *SearchUserNameChangeRequests200ResponseAnyOf) GetUserNameChangeRequestsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserNameChangeRequests, true
}

// SetUserNameChangeRequests sets field value
func (o *SearchUserNameChangeRequests200ResponseAnyOf) SetUserNameChangeRequests(v []string) {
	o.UserNameChangeRequests = v
}

func (o SearchUserNameChangeRequests200ResponseAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchUserNameChangeRequests200ResponseAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user_name_change_requests"] = o.UserNameChangeRequests
	return toSerialize, nil
}

func (o *SearchUserNameChangeRequests200ResponseAnyOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_name_change_requests",
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

	varSearchUserNameChangeRequests200ResponseAnyOf := _SearchUserNameChangeRequests200ResponseAnyOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchUserNameChangeRequests200ResponseAnyOf)

	if err != nil {
		return err
	}

	*o = SearchUserNameChangeRequests200ResponseAnyOf(varSearchUserNameChangeRequests200ResponseAnyOf)

	return err
}

type NullableSearchUserNameChangeRequests200ResponseAnyOf struct {
	value *SearchUserNameChangeRequests200ResponseAnyOf
	isSet bool
}

func (v NullableSearchUserNameChangeRequests200ResponseAnyOf) Get() *SearchUserNameChangeRequests200ResponseAnyOf {
	return v.value
}

func (v *NullableSearchUserNameChangeRequests200ResponseAnyOf) Set(val *SearchUserNameChangeRequests200ResponseAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchUserNameChangeRequests200ResponseAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchUserNameChangeRequests200ResponseAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchUserNameChangeRequests200ResponseAnyOf(val *SearchUserNameChangeRequests200ResponseAnyOf) *NullableSearchUserNameChangeRequests200ResponseAnyOf {
	return &NullableSearchUserNameChangeRequests200ResponseAnyOf{value: val, isSet: true}
}

func (v NullableSearchUserNameChangeRequests200ResponseAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchUserNameChangeRequests200ResponseAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


