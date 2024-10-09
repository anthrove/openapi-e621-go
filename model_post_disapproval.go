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
	"time"
)

// checks if the PostDisapproval type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostDisapproval{}

// PostDisapproval struct for PostDisapproval
type PostDisapproval struct {
	Id        int32          `json:"id"`
	UserId    int32          `json:"user_id"`
	PostId    int32          `json:"post_id"`
	Reason    string         `json:"reason"`
	Message   NullableString `json:"message"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type _PostDisapproval PostDisapproval

// NewPostDisapproval instantiates a new PostDisapproval object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostDisapproval(id int32, userId int32, postId int32, reason string, message NullableString, createdAt time.Time, updatedAt time.Time) *PostDisapproval {
	this := PostDisapproval{}
	this.Id = id
	this.UserId = userId
	this.PostId = postId
	this.Reason = reason
	this.Message = message
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewPostDisapprovalWithDefaults instantiates a new PostDisapproval object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostDisapprovalWithDefaults() *PostDisapproval {
	this := PostDisapproval{}
	return &this
}

// GetId returns the Id field value
func (o *PostDisapproval) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PostDisapproval) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PostDisapproval) SetId(v int32) {
	o.Id = v
}

// GetUserId returns the UserId field value
func (o *PostDisapproval) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *PostDisapproval) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *PostDisapproval) SetUserId(v int32) {
	o.UserId = v
}

// GetPostId returns the PostId field value
func (o *PostDisapproval) GetPostId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PostId
}

// GetPostIdOk returns a tuple with the PostId field value
// and a boolean to check if the value has been set.
func (o *PostDisapproval) GetPostIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostId, true
}

// SetPostId sets field value
func (o *PostDisapproval) SetPostId(v int32) {
	o.PostId = v
}

// GetReason returns the Reason field value
func (o *PostDisapproval) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *PostDisapproval) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *PostDisapproval) SetReason(v string) {
	o.Reason = v
}

// GetMessage returns the Message field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PostDisapproval) GetMessage() string {
	if o == nil || o.Message.Get() == nil {
		var ret string
		return ret
	}

	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostDisapproval) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// SetMessage sets field value
func (o *PostDisapproval) SetMessage(v string) {
	o.Message.Set(&v)
}

// GetCreatedAt returns the CreatedAt field value
func (o *PostDisapproval) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PostDisapproval) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PostDisapproval) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PostDisapproval) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PostDisapproval) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PostDisapproval) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o PostDisapproval) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostDisapproval) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["user_id"] = o.UserId
	toSerialize["post_id"] = o.PostId
	toSerialize["reason"] = o.Reason
	toSerialize["message"] = o.Message.Get()
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *PostDisapproval) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"user_id",
		"post_id",
		"reason",
		"message",
		"created_at",
		"updated_at",
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

	varPostDisapproval := _PostDisapproval{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostDisapproval)

	if err != nil {
		return err
	}

	*o = PostDisapproval(varPostDisapproval)

	return err
}

type NullablePostDisapproval struct {
	value *PostDisapproval
	isSet bool
}

func (v NullablePostDisapproval) Get() *PostDisapproval {
	return v.value
}

func (v *NullablePostDisapproval) Set(val *PostDisapproval) {
	v.value = val
	v.isSet = true
}

func (v NullablePostDisapproval) IsSet() bool {
	return v.isSet
}

func (v *NullablePostDisapproval) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostDisapproval(val *PostDisapproval) *NullablePostDisapproval {
	return &NullablePostDisapproval{value: val, isSet: true}
}

func (v NullablePostDisapproval) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostDisapproval) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
