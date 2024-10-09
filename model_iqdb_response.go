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

// checks if the IQDBResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IQDBResponse{}

// IQDBResponse struct for IQDBResponse
type IQDBResponse struct {
	Hash   string           `json:"hash"`
	PostId int32            `json:"post_id"`
	Score  float32          `json:"score"`
	Post   IQDBResponsePost `json:"post"`
}

type _IQDBResponse IQDBResponse

// NewIQDBResponse instantiates a new IQDBResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIQDBResponse(hash string, postId int32, score float32, post IQDBResponsePost) *IQDBResponse {
	this := IQDBResponse{}
	this.Hash = hash
	this.PostId = postId
	this.Score = score
	this.Post = post
	return &this
}

// NewIQDBResponseWithDefaults instantiates a new IQDBResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIQDBResponseWithDefaults() *IQDBResponse {
	this := IQDBResponse{}
	return &this
}

// GetHash returns the Hash field value
func (o *IQDBResponse) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *IQDBResponse) GetHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *IQDBResponse) SetHash(v string) {
	o.Hash = v
}

// GetPostId returns the PostId field value
func (o *IQDBResponse) GetPostId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PostId
}

// GetPostIdOk returns a tuple with the PostId field value
// and a boolean to check if the value has been set.
func (o *IQDBResponse) GetPostIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostId, true
}

// SetPostId sets field value
func (o *IQDBResponse) SetPostId(v int32) {
	o.PostId = v
}

// GetScore returns the Score field value
func (o *IQDBResponse) GetScore() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *IQDBResponse) GetScoreOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *IQDBResponse) SetScore(v float32) {
	o.Score = v
}

// GetPost returns the Post field value
func (o *IQDBResponse) GetPost() IQDBResponsePost {
	if o == nil {
		var ret IQDBResponsePost
		return ret
	}

	return o.Post
}

// GetPostOk returns a tuple with the Post field value
// and a boolean to check if the value has been set.
func (o *IQDBResponse) GetPostOk() (*IQDBResponsePost, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Post, true
}

// SetPost sets field value
func (o *IQDBResponse) SetPost(v IQDBResponsePost) {
	o.Post = v
}

func (o IQDBResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IQDBResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hash"] = o.Hash
	toSerialize["post_id"] = o.PostId
	toSerialize["score"] = o.Score
	toSerialize["post"] = o.Post
	return toSerialize, nil
}

func (o *IQDBResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hash",
		"post_id",
		"score",
		"post",
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

	varIQDBResponse := _IQDBResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIQDBResponse)

	if err != nil {
		return err
	}

	*o = IQDBResponse(varIQDBResponse)

	return err
}

type NullableIQDBResponse struct {
	value *IQDBResponse
	isSet bool
}

func (v NullableIQDBResponse) Get() *IQDBResponse {
	return v.value
}

func (v *NullableIQDBResponse) Set(val *IQDBResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIQDBResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIQDBResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIQDBResponse(val *IQDBResponse) *NullableIQDBResponse {
	return &NullableIQDBResponse{value: val, isSet: true}
}

func (v NullableIQDBResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIQDBResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
