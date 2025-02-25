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

// checks if the CreatePostVote200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePostVote200Response{}

// CreatePostVote200Response struct for CreatePostVote200Response
type CreatePostVote200Response struct {
	Score    int32 `json:"score"`
	Up       int32 `json:"up"`
	Down     int32 `json:"down"`
	OurScore int32 `json:"our_score"`
}

type _CreatePostVote200Response CreatePostVote200Response

// NewCreatePostVote200Response instantiates a new CreatePostVote200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePostVote200Response(score int32, up int32, down int32, ourScore int32) *CreatePostVote200Response {
	this := CreatePostVote200Response{}
	this.Score = score
	this.Up = up
	this.Down = down
	this.OurScore = ourScore
	return &this
}

// NewCreatePostVote200ResponseWithDefaults instantiates a new CreatePostVote200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePostVote200ResponseWithDefaults() *CreatePostVote200Response {
	this := CreatePostVote200Response{}
	return &this
}

// GetScore returns the Score field value
func (o *CreatePostVote200Response) GetScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *CreatePostVote200Response) GetScoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *CreatePostVote200Response) SetScore(v int32) {
	o.Score = v
}

// GetUp returns the Up field value
func (o *CreatePostVote200Response) GetUp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Up
}

// GetUpOk returns a tuple with the Up field value
// and a boolean to check if the value has been set.
func (o *CreatePostVote200Response) GetUpOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Up, true
}

// SetUp sets field value
func (o *CreatePostVote200Response) SetUp(v int32) {
	o.Up = v
}

// GetDown returns the Down field value
func (o *CreatePostVote200Response) GetDown() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Down
}

// GetDownOk returns a tuple with the Down field value
// and a boolean to check if the value has been set.
func (o *CreatePostVote200Response) GetDownOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Down, true
}

// SetDown sets field value
func (o *CreatePostVote200Response) SetDown(v int32) {
	o.Down = v
}

// GetOurScore returns the OurScore field value
func (o *CreatePostVote200Response) GetOurScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OurScore
}

// GetOurScoreOk returns a tuple with the OurScore field value
// and a boolean to check if the value has been set.
func (o *CreatePostVote200Response) GetOurScoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OurScore, true
}

// SetOurScore sets field value
func (o *CreatePostVote200Response) SetOurScore(v int32) {
	o.OurScore = v
}

func (o CreatePostVote200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePostVote200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["score"] = o.Score
	toSerialize["up"] = o.Up
	toSerialize["down"] = o.Down
	toSerialize["our_score"] = o.OurScore
	return toSerialize, nil
}

func (o *CreatePostVote200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"score",
		"up",
		"down",
		"our_score",
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

	varCreatePostVote200Response := _CreatePostVote200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreatePostVote200Response)

	if err != nil {
		return err
	}

	*o = CreatePostVote200Response(varCreatePostVote200Response)

	return err
}

type NullableCreatePostVote200Response struct {
	value *CreatePostVote200Response
	isSet bool
}

func (v NullableCreatePostVote200Response) Get() *CreatePostVote200Response {
	return v.value
}

func (v *NullableCreatePostVote200Response) Set(val *CreatePostVote200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePostVote200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePostVote200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePostVote200Response(val *CreatePostVote200Response) *NullableCreatePostVote200Response {
	return &NullableCreatePostVote200Response{value: val, isSet: true}
}

func (v NullableCreatePostVote200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePostVote200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
