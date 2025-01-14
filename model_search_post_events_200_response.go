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

// checks if the SearchPostEvents200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchPostEvents200Response{}

// SearchPostEvents200Response struct for SearchPostEvents200Response
type SearchPostEvents200Response struct {
	PostEvents []PostEvent `json:"post_events"`
}

type _SearchPostEvents200Response SearchPostEvents200Response

// NewSearchPostEvents200Response instantiates a new SearchPostEvents200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchPostEvents200Response(postEvents []PostEvent) *SearchPostEvents200Response {
	this := SearchPostEvents200Response{}
	this.PostEvents = postEvents
	return &this
}

// NewSearchPostEvents200ResponseWithDefaults instantiates a new SearchPostEvents200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchPostEvents200ResponseWithDefaults() *SearchPostEvents200Response {
	this := SearchPostEvents200Response{}
	return &this
}

// GetPostEvents returns the PostEvents field value
func (o *SearchPostEvents200Response) GetPostEvents() []PostEvent {
	if o == nil {
		var ret []PostEvent
		return ret
	}

	return o.PostEvents
}

// GetPostEventsOk returns a tuple with the PostEvents field value
// and a boolean to check if the value has been set.
func (o *SearchPostEvents200Response) GetPostEventsOk() ([]PostEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostEvents, true
}

// SetPostEvents sets field value
func (o *SearchPostEvents200Response) SetPostEvents(v []PostEvent) {
	o.PostEvents = v
}

func (o SearchPostEvents200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchPostEvents200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["post_events"] = o.PostEvents
	return toSerialize, nil
}

func (o *SearchPostEvents200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"post_events",
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

	varSearchPostEvents200Response := _SearchPostEvents200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchPostEvents200Response)

	if err != nil {
		return err
	}

	*o = SearchPostEvents200Response(varSearchPostEvents200Response)

	return err
}

type NullableSearchPostEvents200Response struct {
	value *SearchPostEvents200Response
	isSet bool
}

func (v NullableSearchPostEvents200Response) Get() *SearchPostEvents200Response {
	return v.value
}

func (v *NullableSearchPostEvents200Response) Set(val *SearchPostEvents200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchPostEvents200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchPostEvents200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchPostEvents200Response(val *SearchPostEvents200Response) *NullableSearchPostEvents200Response {
	return &NullableSearchPostEvents200Response{value: val, isSet: true}
}

func (v NullableSearchPostEvents200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchPostEvents200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
