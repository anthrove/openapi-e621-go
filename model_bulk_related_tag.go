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

// checks if the BulkRelatedTag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkRelatedTag{}

// BulkRelatedTag struct for BulkRelatedTag
type BulkRelatedTag struct {
	Name       string        `json:"name"`
	CategoryId TagCategories `json:"category_id"`
	Count      int32         `json:"count"`
}

type _BulkRelatedTag BulkRelatedTag

// NewBulkRelatedTag instantiates a new BulkRelatedTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkRelatedTag(name string, categoryId TagCategories, count int32) *BulkRelatedTag {
	this := BulkRelatedTag{}
	this.Name = name
	this.CategoryId = categoryId
	this.Count = count
	return &this
}

// NewBulkRelatedTagWithDefaults instantiates a new BulkRelatedTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkRelatedTagWithDefaults() *BulkRelatedTag {
	this := BulkRelatedTag{}
	return &this
}

// GetName returns the Name field value
func (o *BulkRelatedTag) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BulkRelatedTag) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BulkRelatedTag) SetName(v string) {
	o.Name = v
}

// GetCategoryId returns the CategoryId field value
func (o *BulkRelatedTag) GetCategoryId() TagCategories {
	if o == nil {
		var ret TagCategories
		return ret
	}

	return o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value
// and a boolean to check if the value has been set.
func (o *BulkRelatedTag) GetCategoryIdOk() (*TagCategories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CategoryId, true
}

// SetCategoryId sets field value
func (o *BulkRelatedTag) SetCategoryId(v TagCategories) {
	o.CategoryId = v
}

// GetCount returns the Count field value
func (o *BulkRelatedTag) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *BulkRelatedTag) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *BulkRelatedTag) SetCount(v int32) {
	o.Count = v
}

func (o BulkRelatedTag) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkRelatedTag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["category_id"] = o.CategoryId
	toSerialize["count"] = o.Count
	return toSerialize, nil
}

func (o *BulkRelatedTag) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"category_id",
		"count",
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

	varBulkRelatedTag := _BulkRelatedTag{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBulkRelatedTag)

	if err != nil {
		return err
	}

	*o = BulkRelatedTag(varBulkRelatedTag)

	return err
}

type NullableBulkRelatedTag struct {
	value *BulkRelatedTag
	isSet bool
}

func (v NullableBulkRelatedTag) Get() *BulkRelatedTag {
	return v.value
}

func (v *NullableBulkRelatedTag) Set(val *BulkRelatedTag) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkRelatedTag) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkRelatedTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkRelatedTag(val *BulkRelatedTag) *NullableBulkRelatedTag {
	return &NullableBulkRelatedTag{value: val, isSet: true}
}

func (v NullableBulkRelatedTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkRelatedTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
