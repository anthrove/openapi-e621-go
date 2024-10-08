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

// checks if the SearchAvoidPostings200ResponseAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchAvoidPostings200ResponseAnyOf{}

// SearchAvoidPostings200ResponseAnyOf No Results
type SearchAvoidPostings200ResponseAnyOf struct {
	AvoidPostings []string `json:"avoid_postings"`
}

type _SearchAvoidPostings200ResponseAnyOf SearchAvoidPostings200ResponseAnyOf

// NewSearchAvoidPostings200ResponseAnyOf instantiates a new SearchAvoidPostings200ResponseAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchAvoidPostings200ResponseAnyOf(avoidPostings []string) *SearchAvoidPostings200ResponseAnyOf {
	this := SearchAvoidPostings200ResponseAnyOf{}
	this.AvoidPostings = avoidPostings
	return &this
}

// NewSearchAvoidPostings200ResponseAnyOfWithDefaults instantiates a new SearchAvoidPostings200ResponseAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchAvoidPostings200ResponseAnyOfWithDefaults() *SearchAvoidPostings200ResponseAnyOf {
	this := SearchAvoidPostings200ResponseAnyOf{}
	return &this
}

// GetAvoidPostings returns the AvoidPostings field value
func (o *SearchAvoidPostings200ResponseAnyOf) GetAvoidPostings() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AvoidPostings
}

// GetAvoidPostingsOk returns a tuple with the AvoidPostings field value
// and a boolean to check if the value has been set.
func (o *SearchAvoidPostings200ResponseAnyOf) GetAvoidPostingsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvoidPostings, true
}

// SetAvoidPostings sets field value
func (o *SearchAvoidPostings200ResponseAnyOf) SetAvoidPostings(v []string) {
	o.AvoidPostings = v
}

func (o SearchAvoidPostings200ResponseAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchAvoidPostings200ResponseAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["avoid_postings"] = o.AvoidPostings
	return toSerialize, nil
}

func (o *SearchAvoidPostings200ResponseAnyOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"avoid_postings",
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

	varSearchAvoidPostings200ResponseAnyOf := _SearchAvoidPostings200ResponseAnyOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchAvoidPostings200ResponseAnyOf)

	if err != nil {
		return err
	}

	*o = SearchAvoidPostings200ResponseAnyOf(varSearchAvoidPostings200ResponseAnyOf)

	return err
}

type NullableSearchAvoidPostings200ResponseAnyOf struct {
	value *SearchAvoidPostings200ResponseAnyOf
	isSet bool
}

func (v NullableSearchAvoidPostings200ResponseAnyOf) Get() *SearchAvoidPostings200ResponseAnyOf {
	return v.value
}

func (v *NullableSearchAvoidPostings200ResponseAnyOf) Set(val *SearchAvoidPostings200ResponseAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchAvoidPostings200ResponseAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchAvoidPostings200ResponseAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchAvoidPostings200ResponseAnyOf(val *SearchAvoidPostings200ResponseAnyOf) *NullableSearchAvoidPostings200ResponseAnyOf {
	return &NullableSearchAvoidPostings200ResponseAnyOf{value: val, isSet: true}
}

func (v NullableSearchAvoidPostings200ResponseAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchAvoidPostings200ResponseAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


