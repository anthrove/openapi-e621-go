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

// checks if the SearchMascots200ResponseAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchMascots200ResponseAnyOf{}

// SearchMascots200ResponseAnyOf No Results
type SearchMascots200ResponseAnyOf struct {
	Mascots []string `json:"mascots"`
}

type _SearchMascots200ResponseAnyOf SearchMascots200ResponseAnyOf

// NewSearchMascots200ResponseAnyOf instantiates a new SearchMascots200ResponseAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchMascots200ResponseAnyOf(mascots []string) *SearchMascots200ResponseAnyOf {
	this := SearchMascots200ResponseAnyOf{}
	this.Mascots = mascots
	return &this
}

// NewSearchMascots200ResponseAnyOfWithDefaults instantiates a new SearchMascots200ResponseAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchMascots200ResponseAnyOfWithDefaults() *SearchMascots200ResponseAnyOf {
	this := SearchMascots200ResponseAnyOf{}
	return &this
}

// GetMascots returns the Mascots field value
func (o *SearchMascots200ResponseAnyOf) GetMascots() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Mascots
}

// GetMascotsOk returns a tuple with the Mascots field value
// and a boolean to check if the value has been set.
func (o *SearchMascots200ResponseAnyOf) GetMascotsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mascots, true
}

// SetMascots sets field value
func (o *SearchMascots200ResponseAnyOf) SetMascots(v []string) {
	o.Mascots = v
}

func (o SearchMascots200ResponseAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchMascots200ResponseAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mascots"] = o.Mascots
	return toSerialize, nil
}

func (o *SearchMascots200ResponseAnyOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mascots",
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

	varSearchMascots200ResponseAnyOf := _SearchMascots200ResponseAnyOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchMascots200ResponseAnyOf)

	if err != nil {
		return err
	}

	*o = SearchMascots200ResponseAnyOf(varSearchMascots200ResponseAnyOf)

	return err
}

type NullableSearchMascots200ResponseAnyOf struct {
	value *SearchMascots200ResponseAnyOf
	isSet bool
}

func (v NullableSearchMascots200ResponseAnyOf) Get() *SearchMascots200ResponseAnyOf {
	return v.value
}

func (v *NullableSearchMascots200ResponseAnyOf) Set(val *SearchMascots200ResponseAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchMascots200ResponseAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchMascots200ResponseAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchMascots200ResponseAnyOf(val *SearchMascots200ResponseAnyOf) *NullableSearchMascots200ResponseAnyOf {
	return &NullableSearchMascots200ResponseAnyOf{value: val, isSet: true}
}

func (v NullableSearchMascots200ResponseAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchMascots200ResponseAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
