/*
E621

OpenAPI definition for E621's API. You can find the source [here](https://github.com/DonovanDMC/E621OpenAPI)<br> This document is intended to compliment E621's existing [API Documentation](https://e621.net/help/api).<br> Note if E621's api is under attack and/or cloudflare protections are enabled, the \"Try it out\" buttons here will not work.<br> If they are not working, you can check this [Unofficial Status Page](https://status.e621.ws). 

API version: d69c34e
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi-e621-go

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the PostApproval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostApproval{}

// PostApproval struct for PostApproval
type PostApproval struct {
	Id float32 `json:"id"`
	UserId float32 `json:"user_id"`
	PostId float32 `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type _PostApproval PostApproval

// NewPostApproval instantiates a new PostApproval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostApproval(id float32, userId float32, postId float32, createdAt time.Time, updatedAt time.Time) *PostApproval {
	this := PostApproval{}
	this.Id = id
	this.UserId = userId
	this.PostId = postId
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewPostApprovalWithDefaults instantiates a new PostApproval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostApprovalWithDefaults() *PostApproval {
	this := PostApproval{}
	return &this
}

// GetId returns the Id field value
func (o *PostApproval) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PostApproval) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PostApproval) SetId(v float32) {
	o.Id = v
}

// GetUserId returns the UserId field value
func (o *PostApproval) GetUserId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *PostApproval) GetUserIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *PostApproval) SetUserId(v float32) {
	o.UserId = v
}

// GetPostId returns the PostId field value
func (o *PostApproval) GetPostId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PostId
}

// GetPostIdOk returns a tuple with the PostId field value
// and a boolean to check if the value has been set.
func (o *PostApproval) GetPostIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostId, true
}

// SetPostId sets field value
func (o *PostApproval) SetPostId(v float32) {
	o.PostId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PostApproval) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PostApproval) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PostApproval) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PostApproval) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PostApproval) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PostApproval) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o PostApproval) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostApproval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["user_id"] = o.UserId
	toSerialize["post_id"] = o.PostId
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *PostApproval) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"user_id",
		"post_id",
		"created_at",
		"updated_at",
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

	varPostApproval := _PostApproval{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostApproval)

	if err != nil {
		return err
	}

	*o = PostApproval(varPostApproval)

	return err
}

type NullablePostApproval struct {
	value *PostApproval
	isSet bool
}

func (v NullablePostApproval) Get() *PostApproval {
	return v.value
}

func (v *NullablePostApproval) Set(val *PostApproval) {
	v.value = val
	v.isSet = true
}

func (v NullablePostApproval) IsSet() bool {
	return v.isSet
}

func (v *NullablePostApproval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostApproval(val *PostApproval) *NullablePostApproval {
	return &NullablePostApproval{value: val, isSet: true}
}

func (v NullablePostApproval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostApproval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

