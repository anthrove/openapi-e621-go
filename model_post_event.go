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

// checks if the PostEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostEvent{}

// PostEvent struct for PostEvent
type PostEvent struct {
	Id        float32          `json:"id"`
	CreatorId NullableFloat32  `json:"creator_id"`
	PostId    float32          `json:"post_id"`
	Action    PostEventActions `json:"action"`
	CreatedAt time.Time        `json:"created_at"`
}

type _PostEvent PostEvent

// NewPostEvent instantiates a new PostEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostEvent(id float32, creatorId NullableFloat32, postId float32, action PostEventActions, createdAt time.Time) *PostEvent {
	this := PostEvent{}
	this.Id = id
	this.CreatorId = creatorId
	this.PostId = postId
	this.Action = action
	this.CreatedAt = createdAt
	return &this
}

// NewPostEventWithDefaults instantiates a new PostEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostEventWithDefaults() *PostEvent {
	this := PostEvent{}
	return &this
}

// GetId returns the Id field value
func (o *PostEvent) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PostEvent) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PostEvent) SetId(v float32) {
	o.Id = v
}

// GetCreatorId returns the CreatorId field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *PostEvent) GetCreatorId() float32 {
	if o == nil || o.CreatorId.Get() == nil {
		var ret float32
		return ret
	}

	return *o.CreatorId.Get()
}

// GetCreatorIdOk returns a tuple with the CreatorId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostEvent) GetCreatorIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatorId.Get(), o.CreatorId.IsSet()
}

// SetCreatorId sets field value
func (o *PostEvent) SetCreatorId(v float32) {
	o.CreatorId.Set(&v)
}

// GetPostId returns the PostId field value
func (o *PostEvent) GetPostId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PostId
}

// GetPostIdOk returns a tuple with the PostId field value
// and a boolean to check if the value has been set.
func (o *PostEvent) GetPostIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostId, true
}

// SetPostId sets field value
func (o *PostEvent) SetPostId(v float32) {
	o.PostId = v
}

// GetAction returns the Action field value
func (o *PostEvent) GetAction() PostEventActions {
	if o == nil {
		var ret PostEventActions
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *PostEvent) GetActionOk() (*PostEventActions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *PostEvent) SetAction(v PostEventActions) {
	o.Action = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PostEvent) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PostEvent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PostEvent) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o PostEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["creator_id"] = o.CreatorId.Get()
	toSerialize["post_id"] = o.PostId
	toSerialize["action"] = o.Action
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *PostEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"creator_id",
		"post_id",
		"action",
		"created_at",
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

	varPostEvent := _PostEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostEvent)

	if err != nil {
		return err
	}

	*o = PostEvent(varPostEvent)

	return err
}

type NullablePostEvent struct {
	value *PostEvent
	isSet bool
}

func (v NullablePostEvent) Get() *PostEvent {
	return v.value
}

func (v *NullablePostEvent) Set(val *PostEvent) {
	v.value = val
	v.isSet = true
}

func (v NullablePostEvent) IsSet() bool {
	return v.isSet
}

func (v *NullablePostEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostEvent(val *PostEvent) *NullablePostEvent {
	return &NullablePostEvent{value: val, isSet: true}
}

func (v NullablePostEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}