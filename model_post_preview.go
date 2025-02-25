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

// checks if the PostPreview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostPreview{}

// PostPreview struct for PostPreview
type PostPreview struct {
	Width  int32          `json:"width"`
	Height int32          `json:"height"`
	Url    NullableString `json:"url"`
}

type _PostPreview PostPreview

// NewPostPreview instantiates a new PostPreview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostPreview(width int32, height int32, url NullableString) *PostPreview {
	this := PostPreview{}
	this.Width = width
	this.Height = height
	this.Url = url
	return &this
}

// NewPostPreviewWithDefaults instantiates a new PostPreview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostPreviewWithDefaults() *PostPreview {
	this := PostPreview{}
	return &this
}

// GetWidth returns the Width field value
func (o *PostPreview) GetWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *PostPreview) GetWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *PostPreview) SetWidth(v int32) {
	o.Width = v
}

// GetHeight returns the Height field value
func (o *PostPreview) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *PostPreview) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *PostPreview) SetHeight(v int32) {
	o.Height = v
}

// GetUrl returns the Url field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PostPreview) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}

	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostPreview) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// SetUrl sets field value
func (o *PostPreview) SetUrl(v string) {
	o.Url.Set(&v)
}

func (o PostPreview) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostPreview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["width"] = o.Width
	toSerialize["height"] = o.Height
	toSerialize["url"] = o.Url.Get()
	return toSerialize, nil
}

func (o *PostPreview) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"width",
		"height",
		"url",
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

	varPostPreview := _PostPreview{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostPreview)

	if err != nil {
		return err
	}

	*o = PostPreview(varPostPreview)

	return err
}

type NullablePostPreview struct {
	value *PostPreview
	isSet bool
}

func (v NullablePostPreview) Get() *PostPreview {
	return v.value
}

func (v *NullablePostPreview) Set(val *PostPreview) {
	v.value = val
	v.isSet = true
}

func (v NullablePostPreview) IsSet() bool {
	return v.isSet
}

func (v *NullablePostPreview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostPreview(val *PostPreview) *NullablePostPreview {
	return &NullablePostPreview{value: val, isSet: true}
}

func (v NullablePostPreview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostPreview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
