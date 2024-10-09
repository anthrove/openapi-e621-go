/*
E621

OpenAPI definition for E621's API. You can find the source [here](https://github.com/DonovanDMC/E621OpenAPI)<br> This document is intended to compliment E621's existing [API Documentation](https://e621.net/help/api).<br> Note if E621's api is under attack and/or cloudflare protections are enabled, the \"Try it out\" buttons here will not work.<br> If they are not working, you can check this [Unofficial Status Page](https://status.e621.ws).

API version: d69c34e
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SearchEmailBlacklists200ResponseAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchEmailBlacklists200ResponseAnyOf{}

// SearchEmailBlacklists200ResponseAnyOf No Results
type SearchEmailBlacklists200ResponseAnyOf struct {
	EmailBlacklists []string `json:"email_blacklists"`
}

type _SearchEmailBlacklists200ResponseAnyOf SearchEmailBlacklists200ResponseAnyOf

// NewSearchEmailBlacklists200ResponseAnyOf instantiates a new SearchEmailBlacklists200ResponseAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchEmailBlacklists200ResponseAnyOf(emailBlacklists []string) *SearchEmailBlacklists200ResponseAnyOf {
	this := SearchEmailBlacklists200ResponseAnyOf{}
	this.EmailBlacklists = emailBlacklists
	return &this
}

// NewSearchEmailBlacklists200ResponseAnyOfWithDefaults instantiates a new SearchEmailBlacklists200ResponseAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchEmailBlacklists200ResponseAnyOfWithDefaults() *SearchEmailBlacklists200ResponseAnyOf {
	this := SearchEmailBlacklists200ResponseAnyOf{}
	return &this
}

// GetEmailBlacklists returns the EmailBlacklists field value
func (o *SearchEmailBlacklists200ResponseAnyOf) GetEmailBlacklists() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EmailBlacklists
}

// GetEmailBlacklistsOk returns a tuple with the EmailBlacklists field value
// and a boolean to check if the value has been set.
func (o *SearchEmailBlacklists200ResponseAnyOf) GetEmailBlacklistsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EmailBlacklists, true
}

// SetEmailBlacklists sets field value
func (o *SearchEmailBlacklists200ResponseAnyOf) SetEmailBlacklists(v []string) {
	o.EmailBlacklists = v
}

func (o SearchEmailBlacklists200ResponseAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchEmailBlacklists200ResponseAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email_blacklists"] = o.EmailBlacklists
	return toSerialize, nil
}

func (o *SearchEmailBlacklists200ResponseAnyOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email_blacklists",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSearchEmailBlacklists200ResponseAnyOf := _SearchEmailBlacklists200ResponseAnyOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchEmailBlacklists200ResponseAnyOf)

	if err != nil {
		return err
	}

	*o = SearchEmailBlacklists200ResponseAnyOf(varSearchEmailBlacklists200ResponseAnyOf)

	return err
}

type NullableSearchEmailBlacklists200ResponseAnyOf struct {
	value *SearchEmailBlacklists200ResponseAnyOf
	isSet bool
}

func (v NullableSearchEmailBlacklists200ResponseAnyOf) Get() *SearchEmailBlacklists200ResponseAnyOf {
	return v.value
}

func (v *NullableSearchEmailBlacklists200ResponseAnyOf) Set(val *SearchEmailBlacklists200ResponseAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchEmailBlacklists200ResponseAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchEmailBlacklists200ResponseAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchEmailBlacklists200ResponseAnyOf(val *SearchEmailBlacklists200ResponseAnyOf) *NullableSearchEmailBlacklists200ResponseAnyOf {
	return &NullableSearchEmailBlacklists200ResponseAnyOf{value: val, isSet: true}
}

func (v NullableSearchEmailBlacklists200ResponseAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchEmailBlacklists200ResponseAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
