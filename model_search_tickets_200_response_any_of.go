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

// checks if the SearchTickets200ResponseAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchTickets200ResponseAnyOf{}

// SearchTickets200ResponseAnyOf No Results
type SearchTickets200ResponseAnyOf struct {
	Tickets []string `json:"tickets"`
}

type _SearchTickets200ResponseAnyOf SearchTickets200ResponseAnyOf

// NewSearchTickets200ResponseAnyOf instantiates a new SearchTickets200ResponseAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchTickets200ResponseAnyOf(tickets []string) *SearchTickets200ResponseAnyOf {
	this := SearchTickets200ResponseAnyOf{}
	this.Tickets = tickets
	return &this
}

// NewSearchTickets200ResponseAnyOfWithDefaults instantiates a new SearchTickets200ResponseAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchTickets200ResponseAnyOfWithDefaults() *SearchTickets200ResponseAnyOf {
	this := SearchTickets200ResponseAnyOf{}
	return &this
}

// GetTickets returns the Tickets field value
func (o *SearchTickets200ResponseAnyOf) GetTickets() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tickets
}

// GetTicketsOk returns a tuple with the Tickets field value
// and a boolean to check if the value has been set.
func (o *SearchTickets200ResponseAnyOf) GetTicketsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tickets, true
}

// SetTickets sets field value
func (o *SearchTickets200ResponseAnyOf) SetTickets(v []string) {
	o.Tickets = v
}

func (o SearchTickets200ResponseAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchTickets200ResponseAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tickets"] = o.Tickets
	return toSerialize, nil
}

func (o *SearchTickets200ResponseAnyOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tickets",
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

	varSearchTickets200ResponseAnyOf := _SearchTickets200ResponseAnyOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchTickets200ResponseAnyOf)

	if err != nil {
		return err
	}

	*o = SearchTickets200ResponseAnyOf(varSearchTickets200ResponseAnyOf)

	return err
}

type NullableSearchTickets200ResponseAnyOf struct {
	value *SearchTickets200ResponseAnyOf
	isSet bool
}

func (v NullableSearchTickets200ResponseAnyOf) Get() *SearchTickets200ResponseAnyOf {
	return v.value
}

func (v *NullableSearchTickets200ResponseAnyOf) Set(val *SearchTickets200ResponseAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchTickets200ResponseAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchTickets200ResponseAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchTickets200ResponseAnyOf(val *SearchTickets200ResponseAnyOf) *NullableSearchTickets200ResponseAnyOf {
	return &NullableSearchTickets200ResponseAnyOf{value: val, isSet: true}
}

func (v NullableSearchTickets200ResponseAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchTickets200ResponseAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
