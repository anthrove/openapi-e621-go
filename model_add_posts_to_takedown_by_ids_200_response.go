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

// checks if the AddPostsToTakedownByIds200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddPostsToTakedownByIds200Response{}

// AddPostsToTakedownByIds200Response struct for AddPostsToTakedownByIds200Response
type AddPostsToTakedownByIds200Response struct {
	AddedCount float32 `json:"added_count"`
	AddedPostIds []float32 `json:"added_post_ids"`
}

type _AddPostsToTakedownByIds200Response AddPostsToTakedownByIds200Response

// NewAddPostsToTakedownByIds200Response instantiates a new AddPostsToTakedownByIds200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddPostsToTakedownByIds200Response(addedCount float32, addedPostIds []float32) *AddPostsToTakedownByIds200Response {
	this := AddPostsToTakedownByIds200Response{}
	this.AddedCount = addedCount
	this.AddedPostIds = addedPostIds
	return &this
}

// NewAddPostsToTakedownByIds200ResponseWithDefaults instantiates a new AddPostsToTakedownByIds200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddPostsToTakedownByIds200ResponseWithDefaults() *AddPostsToTakedownByIds200Response {
	this := AddPostsToTakedownByIds200Response{}
	return &this
}

// GetAddedCount returns the AddedCount field value
func (o *AddPostsToTakedownByIds200Response) GetAddedCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AddedCount
}

// GetAddedCountOk returns a tuple with the AddedCount field value
// and a boolean to check if the value has been set.
func (o *AddPostsToTakedownByIds200Response) GetAddedCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AddedCount, true
}

// SetAddedCount sets field value
func (o *AddPostsToTakedownByIds200Response) SetAddedCount(v float32) {
	o.AddedCount = v
}

// GetAddedPostIds returns the AddedPostIds field value
func (o *AddPostsToTakedownByIds200Response) GetAddedPostIds() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.AddedPostIds
}

// GetAddedPostIdsOk returns a tuple with the AddedPostIds field value
// and a boolean to check if the value has been set.
func (o *AddPostsToTakedownByIds200Response) GetAddedPostIdsOk() ([]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddedPostIds, true
}

// SetAddedPostIds sets field value
func (o *AddPostsToTakedownByIds200Response) SetAddedPostIds(v []float32) {
	o.AddedPostIds = v
}

func (o AddPostsToTakedownByIds200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddPostsToTakedownByIds200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["added_count"] = o.AddedCount
	toSerialize["added_post_ids"] = o.AddedPostIds
	return toSerialize, nil
}

func (o *AddPostsToTakedownByIds200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"added_count",
		"added_post_ids",
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

	varAddPostsToTakedownByIds200Response := _AddPostsToTakedownByIds200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAddPostsToTakedownByIds200Response)

	if err != nil {
		return err
	}

	*o = AddPostsToTakedownByIds200Response(varAddPostsToTakedownByIds200Response)

	return err
}

type NullableAddPostsToTakedownByIds200Response struct {
	value *AddPostsToTakedownByIds200Response
	isSet bool
}

func (v NullableAddPostsToTakedownByIds200Response) Get() *AddPostsToTakedownByIds200Response {
	return v.value
}

func (v *NullableAddPostsToTakedownByIds200Response) Set(val *AddPostsToTakedownByIds200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableAddPostsToTakedownByIds200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableAddPostsToTakedownByIds200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddPostsToTakedownByIds200Response(val *AddPostsToTakedownByIds200Response) *NullableAddPostsToTakedownByIds200Response {
	return &NullableAddPostsToTakedownByIds200Response{value: val, isSet: true}
}

func (v NullableAddPostsToTakedownByIds200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddPostsToTakedownByIds200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


