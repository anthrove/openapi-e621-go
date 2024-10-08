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

// checks if the DTextResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DTextResponse{}

// DTextResponse struct for DTextResponse
type DTextResponse struct {
	Html string `json:"html"`
	Posts map[string]interface{} `json:"posts"`
}

type _DTextResponse DTextResponse

// NewDTextResponse instantiates a new DTextResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDTextResponse(html string, posts map[string]interface{}) *DTextResponse {
	this := DTextResponse{}
	this.Html = html
	this.Posts = posts
	return &this
}

// NewDTextResponseWithDefaults instantiates a new DTextResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDTextResponseWithDefaults() *DTextResponse {
	this := DTextResponse{}
	return &this
}

// GetHtml returns the Html field value
func (o *DTextResponse) GetHtml() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Html
}

// GetHtmlOk returns a tuple with the Html field value
// and a boolean to check if the value has been set.
func (o *DTextResponse) GetHtmlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Html, true
}

// SetHtml sets field value
func (o *DTextResponse) SetHtml(v string) {
	o.Html = v
}

// GetPosts returns the Posts field value
func (o *DTextResponse) GetPosts() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Posts
}

// GetPostsOk returns a tuple with the Posts field value
// and a boolean to check if the value has been set.
func (o *DTextResponse) GetPostsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Posts, true
}

// SetPosts sets field value
func (o *DTextResponse) SetPosts(v map[string]interface{}) {
	o.Posts = v
}

func (o DTextResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DTextResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["html"] = o.Html
	toSerialize["posts"] = o.Posts
	return toSerialize, nil
}

func (o *DTextResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"html",
		"posts",
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

	varDTextResponse := _DTextResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDTextResponse)

	if err != nil {
		return err
	}

	*o = DTextResponse(varDTextResponse)

	return err
}

type NullableDTextResponse struct {
	value *DTextResponse
	isSet bool
}

func (v NullableDTextResponse) Get() *DTextResponse {
	return v.value
}

func (v *NullableDTextResponse) Set(val *DTextResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDTextResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDTextResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDTextResponse(val *DTextResponse) *NullableDTextResponse {
	return &NullableDTextResponse{value: val, isSet: true}
}

func (v NullableDTextResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDTextResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


