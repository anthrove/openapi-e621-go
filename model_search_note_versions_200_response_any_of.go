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

// checks if the SearchNoteVersions200ResponseAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchNoteVersions200ResponseAnyOf{}

// SearchNoteVersions200ResponseAnyOf No Results
type SearchNoteVersions200ResponseAnyOf struct {
	NoteVersions []string `json:"note_versions"`
}

type _SearchNoteVersions200ResponseAnyOf SearchNoteVersions200ResponseAnyOf

// NewSearchNoteVersions200ResponseAnyOf instantiates a new SearchNoteVersions200ResponseAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchNoteVersions200ResponseAnyOf(noteVersions []string) *SearchNoteVersions200ResponseAnyOf {
	this := SearchNoteVersions200ResponseAnyOf{}
	this.NoteVersions = noteVersions
	return &this
}

// NewSearchNoteVersions200ResponseAnyOfWithDefaults instantiates a new SearchNoteVersions200ResponseAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchNoteVersions200ResponseAnyOfWithDefaults() *SearchNoteVersions200ResponseAnyOf {
	this := SearchNoteVersions200ResponseAnyOf{}
	return &this
}

// GetNoteVersions returns the NoteVersions field value
func (o *SearchNoteVersions200ResponseAnyOf) GetNoteVersions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.NoteVersions
}

// GetNoteVersionsOk returns a tuple with the NoteVersions field value
// and a boolean to check if the value has been set.
func (o *SearchNoteVersions200ResponseAnyOf) GetNoteVersionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NoteVersions, true
}

// SetNoteVersions sets field value
func (o *SearchNoteVersions200ResponseAnyOf) SetNoteVersions(v []string) {
	o.NoteVersions = v
}

func (o SearchNoteVersions200ResponseAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchNoteVersions200ResponseAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["note_versions"] = o.NoteVersions
	return toSerialize, nil
}

func (o *SearchNoteVersions200ResponseAnyOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"note_versions",
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

	varSearchNoteVersions200ResponseAnyOf := _SearchNoteVersions200ResponseAnyOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchNoteVersions200ResponseAnyOf)

	if err != nil {
		return err
	}

	*o = SearchNoteVersions200ResponseAnyOf(varSearchNoteVersions200ResponseAnyOf)

	return err
}

type NullableSearchNoteVersions200ResponseAnyOf struct {
	value *SearchNoteVersions200ResponseAnyOf
	isSet bool
}

func (v NullableSearchNoteVersions200ResponseAnyOf) Get() *SearchNoteVersions200ResponseAnyOf {
	return v.value
}

func (v *NullableSearchNoteVersions200ResponseAnyOf) Set(val *SearchNoteVersions200ResponseAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchNoteVersions200ResponseAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchNoteVersions200ResponseAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchNoteVersions200ResponseAnyOf(val *SearchNoteVersions200ResponseAnyOf) *NullableSearchNoteVersions200ResponseAnyOf {
	return &NullableSearchNoteVersions200ResponseAnyOf{value: val, isSet: true}
}

func (v NullableSearchNoteVersions200ResponseAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchNoteVersions200ResponseAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}