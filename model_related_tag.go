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

// checks if the RelatedTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RelatedTag{}

// RelatedTag struct for RelatedTag
type RelatedTag struct {
	Name       string        `json:"name"`
	CategoryId TagCategories `json:"category_id"`
}

type _RelatedTag RelatedTag

// NewRelatedTag instantiates a new RelatedTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelatedTag(name string, categoryId TagCategories) *RelatedTag {
	this := RelatedTag{}
	this.Name = name
	this.CategoryId = categoryId
	return &this
}

// NewRelatedTagWithDefaults instantiates a new RelatedTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelatedTagWithDefaults() *RelatedTag {
	this := RelatedTag{}
	return &this
}

// GetName returns the Name field value
func (o *RelatedTag) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RelatedTag) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RelatedTag) SetName(v string) {
	o.Name = v
}

// GetCategoryId returns the CategoryId field value
func (o *RelatedTag) GetCategoryId() TagCategories {
	if o == nil {
		var ret TagCategories
		return ret
	}

	return o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value
// and a boolean to check if the value has been set.
func (o *RelatedTag) GetCategoryIdOk() (*TagCategories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CategoryId, true
}

// SetCategoryId sets field value
func (o *RelatedTag) SetCategoryId(v TagCategories) {
	o.CategoryId = v
}

func (o RelatedTag) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RelatedTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["category_id"] = o.CategoryId
	return toSerialize, nil
}

func (o *RelatedTag) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"category_id",
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

	varRelatedTag := _RelatedTag{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRelatedTag)

	if err != nil {
		return err
	}

	*o = RelatedTag(varRelatedTag)

	return err
}

type NullableRelatedTag struct {
	value *RelatedTag
	isSet bool
}

func (v NullableRelatedTag) Get() *RelatedTag {
	return v.value
}

func (v *NullableRelatedTag) Set(val *RelatedTag) {
	v.value = val
	v.isSet = true
}

func (v NullableRelatedTag) IsSet() bool {
	return v.isSet
}

func (v *NullableRelatedTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelatedTag(val *RelatedTag) *NullableRelatedTag {
	return &NullableRelatedTag{value: val, isSet: true}
}

func (v NullableRelatedTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelatedTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
