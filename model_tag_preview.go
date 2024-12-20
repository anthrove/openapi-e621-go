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

// checks if the TagPreview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagPreview{}

// TagPreview struct for TagPreview
type TagPreview struct {
	// The name if type=tag, else the antecedent.
	A string `json:"a"`
	// The consequent, only present if type=alias or type=implication.
	B        *string        `json:"b,omitempty"`
	Type     string         `json:"type"`
	TagTypeA TagCategories  `json:"tagTypeA"`
	TagTypeB *TagCategories `json:"tagTypeB,omitempty"`
}

type _TagPreview TagPreview

// NewTagPreview instantiates a new TagPreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagPreview(a string, type_ string, tagTypeA TagCategories) *TagPreview {
	this := TagPreview{}
	this.A = a
	this.Type = type_
	this.TagTypeA = tagTypeA
	return &this
}

// NewTagPreviewWithDefaults instantiates a new TagPreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagPreviewWithDefaults() *TagPreview {
	this := TagPreview{}
	return &this
}

// GetA returns the A field value
func (o *TagPreview) GetA() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.A
}

// GetAOk returns a tuple with the A field value
// and a boolean to check if the value has been set.
func (o *TagPreview) GetAOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.A, true
}

// SetA sets field value
func (o *TagPreview) SetA(v string) {
	o.A = v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *TagPreview) GetB() string {
	if o == nil || IsNil(o.B) {
		var ret string
		return ret
	}
	return *o.B
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPreview) GetBOk() (*string, bool) {
	if o == nil || IsNil(o.B) {
		return nil, false
	}
	return o.B, true
}

// HasB returns a boolean if a field has been set.
func (o *TagPreview) HasB() bool {
	if o != nil && !IsNil(o.B) {
		return true
	}

	return false
}

// SetB gets a reference to the given string and assigns it to the B field.
func (o *TagPreview) SetB(v string) {
	o.B = &v
}

// GetType returns the Type field value
func (o *TagPreview) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TagPreview) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TagPreview) SetType(v string) {
	o.Type = v
}

// GetTagTypeA returns the TagTypeA field value
func (o *TagPreview) GetTagTypeA() TagCategories {
	if o == nil {
		var ret TagCategories
		return ret
	}

	return o.TagTypeA
}

// GetTagTypeAOk returns a tuple with the TagTypeA field value
// and a boolean to check if the value has been set.
func (o *TagPreview) GetTagTypeAOk() (*TagCategories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagTypeA, true
}

// SetTagTypeA sets field value
func (o *TagPreview) SetTagTypeA(v TagCategories) {
	o.TagTypeA = v
}

// GetTagTypeB returns the TagTypeB field value if set, zero value otherwise.
func (o *TagPreview) GetTagTypeB() TagCategories {
	if o == nil || IsNil(o.TagTypeB) {
		var ret TagCategories
		return ret
	}
	return *o.TagTypeB
}

// GetTagTypeBOk returns a tuple with the TagTypeB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagPreview) GetTagTypeBOk() (*TagCategories, bool) {
	if o == nil || IsNil(o.TagTypeB) {
		return nil, false
	}
	return o.TagTypeB, true
}

// HasTagTypeB returns a boolean if a field has been set.
func (o *TagPreview) HasTagTypeB() bool {
	if o != nil && !IsNil(o.TagTypeB) {
		return true
	}

	return false
}

// SetTagTypeB gets a reference to the given TagCategories and assigns it to the TagTypeB field.
func (o *TagPreview) SetTagTypeB(v TagCategories) {
	o.TagTypeB = &v
}

func (o TagPreview) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagPreview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["a"] = o.A
	if !IsNil(o.B) {
		toSerialize["b"] = o.B
	}
	toSerialize["type"] = o.Type
	toSerialize["tagTypeA"] = o.TagTypeA
	if !IsNil(o.TagTypeB) {
		toSerialize["tagTypeB"] = o.TagTypeB
	}
	return toSerialize, nil
}

func (o *TagPreview) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"a",
		"type",
		"tagTypeA",
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

	varTagPreview := _TagPreview{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTagPreview)

	if err != nil {
		return err
	}

	*o = TagPreview(varTagPreview)

	return err
}

type NullableTagPreview struct {
	value *TagPreview
	isSet bool
}

func (v NullableTagPreview) Get() *TagPreview {
	return v.value
}

func (v *NullableTagPreview) Set(val *TagPreview) {
	v.value = val
	v.isSet = true
}

func (v NullableTagPreview) IsSet() bool {
	return v.isSet
}

func (v *NullableTagPreview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagPreview(val *TagPreview) *NullableTagPreview {
	return &NullableTagPreview{value: val, isSet: true}
}

func (v NullableTagPreview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagPreview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
