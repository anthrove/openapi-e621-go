/*
E621

OpenAPI definition for E621's API. You can find the source [here](https://github.com/DonovanDMC/E621OpenAPI)<br> This document is intended to compliment E621's existing [API Documentation](https://e621.net/help/api).<br> Note if E621's api is under attack and/or cloudflare protections are enabled, the \"Try it out\" buttons here will not work.<br> If they are not working, you can check this [Unofficial Status Page](https://status.e621.ws). 

API version: d69c34e
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi-e621-go

import (
	"encoding/json"
)

// checks if the QueryIQDPostRequest1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryIQDPostRequest1{}

// QueryIQDPostRequest1 struct for QueryIQDPostRequest1
type QueryIQDPostRequest1 struct {
	ScoreCutoff *float32 `json:"score_cutoff,omitempty"`
	Url *string `json:"url,omitempty"`
	PostId *string `json:"post_id,omitempty"`
	Hash *string `json:"hash,omitempty"`
}

// NewQueryIQDPostRequest1 instantiates a new QueryIQDPostRequest1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryIQDPostRequest1() *QueryIQDPostRequest1 {
	this := QueryIQDPostRequest1{}
	return &this
}

// NewQueryIQDPostRequest1WithDefaults instantiates a new QueryIQDPostRequest1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryIQDPostRequest1WithDefaults() *QueryIQDPostRequest1 {
	this := QueryIQDPostRequest1{}
	return &this
}

// GetScoreCutoff returns the ScoreCutoff field value if set, zero value otherwise.
func (o *QueryIQDPostRequest1) GetScoreCutoff() float32 {
	if o == nil || IsNil(o.ScoreCutoff) {
		var ret float32
		return ret
	}
	return *o.ScoreCutoff
}

// GetScoreCutoffOk returns a tuple with the ScoreCutoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIQDPostRequest1) GetScoreCutoffOk() (*float32, bool) {
	if o == nil || IsNil(o.ScoreCutoff) {
		return nil, false
	}
	return o.ScoreCutoff, true
}

// HasScoreCutoff returns a boolean if a field has been set.
func (o *QueryIQDPostRequest1) HasScoreCutoff() bool {
	if o != nil && !IsNil(o.ScoreCutoff) {
		return true
	}

	return false
}

// SetScoreCutoff gets a reference to the given float32 and assigns it to the ScoreCutoff field.
func (o *QueryIQDPostRequest1) SetScoreCutoff(v float32) {
	o.ScoreCutoff = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *QueryIQDPostRequest1) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIQDPostRequest1) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *QueryIQDPostRequest1) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *QueryIQDPostRequest1) SetUrl(v string) {
	o.Url = &v
}

// GetPostId returns the PostId field value if set, zero value otherwise.
func (o *QueryIQDPostRequest1) GetPostId() string {
	if o == nil || IsNil(o.PostId) {
		var ret string
		return ret
	}
	return *o.PostId
}

// GetPostIdOk returns a tuple with the PostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIQDPostRequest1) GetPostIdOk() (*string, bool) {
	if o == nil || IsNil(o.PostId) {
		return nil, false
	}
	return o.PostId, true
}

// HasPostId returns a boolean if a field has been set.
func (o *QueryIQDPostRequest1) HasPostId() bool {
	if o != nil && !IsNil(o.PostId) {
		return true
	}

	return false
}

// SetPostId gets a reference to the given string and assigns it to the PostId field.
func (o *QueryIQDPostRequest1) SetPostId(v string) {
	o.PostId = &v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *QueryIQDPostRequest1) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryIQDPostRequest1) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *QueryIQDPostRequest1) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *QueryIQDPostRequest1) SetHash(v string) {
	o.Hash = &v
}

func (o QueryIQDPostRequest1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryIQDPostRequest1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScoreCutoff) {
		toSerialize["score_cutoff"] = o.ScoreCutoff
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.PostId) {
		toSerialize["post_id"] = o.PostId
	}
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	return toSerialize, nil
}

type NullableQueryIQDPostRequest1 struct {
	value *QueryIQDPostRequest1
	isSet bool
}

func (v NullableQueryIQDPostRequest1) Get() *QueryIQDPostRequest1 {
	return v.value
}

func (v *NullableQueryIQDPostRequest1) Set(val *QueryIQDPostRequest1) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryIQDPostRequest1) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryIQDPostRequest1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryIQDPostRequest1(val *QueryIQDPostRequest1) *NullableQueryIQDPostRequest1 {
	return &NullableQueryIQDPostRequest1{value: val, isSet: true}
}

func (v NullableQueryIQDPostRequest1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryIQDPostRequest1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


