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

// checks if the PostFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostFile{}

// PostFile struct for PostFile
type PostFile struct {
	Width  int32          `json:"width"`
	Height int32          `json:"height"`
	Ext    string         `json:"ext"`
	Size   int64          `json:"size"`
	Md5    string         `json:"md5"`
	Url    NullableString `json:"url"`
}

type _PostFile PostFile

// NewPostFile instantiates a new PostFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostFile(width int32, height int32, ext string, size int64, md5 string, url NullableString) *PostFile {
	this := PostFile{}
	this.Width = width
	this.Height = height
	this.Ext = ext
	this.Size = size
	this.Md5 = md5
	this.Url = url
	return &this
}

// NewPostFileWithDefaults instantiates a new PostFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostFileWithDefaults() *PostFile {
	this := PostFile{}
	return &this
}

// GetWidth returns the Width field value
func (o *PostFile) GetWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *PostFile) GetWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *PostFile) SetWidth(v int32) {
	o.Width = v
}

// GetHeight returns the Height field value
func (o *PostFile) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *PostFile) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *PostFile) SetHeight(v int32) {
	o.Height = v
}

// GetExt returns the Ext field value
func (o *PostFile) GetExt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ext
}

// GetExtOk returns a tuple with the Ext field value
// and a boolean to check if the value has been set.
func (o *PostFile) GetExtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ext, true
}

// SetExt sets field value
func (o *PostFile) SetExt(v string) {
	o.Ext = v
}

// GetSize returns the Size field value
func (o *PostFile) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *PostFile) GetSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *PostFile) SetSize(v int64) {
	o.Size = v
}

// GetMd5 returns the Md5 field value
func (o *PostFile) GetMd5() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value
// and a boolean to check if the value has been set.
func (o *PostFile) GetMd5Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Md5, true
}

// SetMd5 sets field value
func (o *PostFile) SetMd5(v string) {
	o.Md5 = v
}

// GetUrl returns the Url field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PostFile) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}

	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostFile) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// SetUrl sets field value
func (o *PostFile) SetUrl(v string) {
	o.Url.Set(&v)
}

func (o PostFile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["width"] = o.Width
	toSerialize["height"] = o.Height
	toSerialize["ext"] = o.Ext
	toSerialize["size"] = o.Size
	toSerialize["md5"] = o.Md5
	toSerialize["url"] = o.Url.Get()
	return toSerialize, nil
}

func (o *PostFile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"width",
		"height",
		"ext",
		"size",
		"md5",
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

	varPostFile := _PostFile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostFile)

	if err != nil {
		return err
	}

	*o = PostFile(varPostFile)

	return err
}

type NullablePostFile struct {
	value *PostFile
	isSet bool
}

func (v NullablePostFile) Get() *PostFile {
	return v.value
}

func (v *NullablePostFile) Set(val *PostFile) {
	v.value = val
	v.isSet = true
}

func (v NullablePostFile) IsSet() bool {
	return v.isSet
}

func (v *NullablePostFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostFile(val *PostFile) *NullablePostFile {
	return &NullablePostFile{value: val, isSet: true}
}

func (v NullablePostFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
