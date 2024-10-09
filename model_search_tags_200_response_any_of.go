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

// checks if the SearchTags200ResponseAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchTags200ResponseAnyOf{}

// SearchTags200ResponseAnyOf No Results
type SearchTags200ResponseAnyOf struct {
	Tags []string `json:"tags"`
}

type _SearchTags200ResponseAnyOf SearchTags200ResponseAnyOf

// NewSearchTags200ResponseAnyOf instantiates a new SearchTags200ResponseAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchTags200ResponseAnyOf(tags []string) *SearchTags200ResponseAnyOf {
	this := SearchTags200ResponseAnyOf{}
	this.Tags = tags
	return &this
}

// NewSearchTags200ResponseAnyOfWithDefaults instantiates a new SearchTags200ResponseAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchTags200ResponseAnyOfWithDefaults() *SearchTags200ResponseAnyOf {
	this := SearchTags200ResponseAnyOf{}
	return &this
}

// GetTags returns the Tags field value
func (o *SearchTags200ResponseAnyOf) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *SearchTags200ResponseAnyOf) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *SearchTags200ResponseAnyOf) SetTags(v []string) {
	o.Tags = v
}

func (o SearchTags200ResponseAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchTags200ResponseAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tags"] = o.Tags
	return toSerialize, nil
}

func (o *SearchTags200ResponseAnyOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tags",
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

	varSearchTags200ResponseAnyOf := _SearchTags200ResponseAnyOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchTags200ResponseAnyOf)

	if err != nil {
		return err
	}

	*o = SearchTags200ResponseAnyOf(varSearchTags200ResponseAnyOf)

	return err
}

type NullableSearchTags200ResponseAnyOf struct {
	value *SearchTags200ResponseAnyOf
	isSet bool
}

func (v NullableSearchTags200ResponseAnyOf) Get() *SearchTags200ResponseAnyOf {
	return v.value
}

func (v *NullableSearchTags200ResponseAnyOf) Set(val *SearchTags200ResponseAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchTags200ResponseAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchTags200ResponseAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchTags200ResponseAnyOf(val *SearchTags200ResponseAnyOf) *NullableSearchTags200ResponseAnyOf {
	return &NullableSearchTags200ResponseAnyOf{value: val, isSet: true}
}

func (v NullableSearchTags200ResponseAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchTags200ResponseAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
