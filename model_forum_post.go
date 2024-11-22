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

// checks if the ForumPost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForumPost{}

// ForumPost struct for ForumPost
type ForumPost struct {
	Id            int32           `json:"id"`
	TopicId       int32           `json:"topic_id"`
	CreatorId     int32           `json:"creator_id"`
	UpdaterId     int32           `json:"updater_id"`
	Body          string          `json:"body"`
	IsHidden      bool            `json:"is_hidden"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
	WarningType   WarningTypes    `json:"warning_type"`
	WarningUserId NullableFloat32 `json:"warning_user_id"`
}

type _ForumPost ForumPost

// NewForumPost instantiates a new ForumPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForumPost(id int32, topicId int32, creatorId int32, updaterId int32, body string, isHidden bool, createdAt time.Time, updatedAt time.Time, warningType WarningTypes, warningUserId NullableFloat32) *ForumPost {
	this := ForumPost{}
	this.Id = id
	this.TopicId = topicId
	this.CreatorId = creatorId
	this.UpdaterId = updaterId
	this.Body = body
	this.IsHidden = isHidden
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.WarningType = warningType
	this.WarningUserId = warningUserId
	return &this
}

// NewForumPostWithDefaults instantiates a new ForumPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForumPostWithDefaults() *ForumPost {
	this := ForumPost{}
	return &this
}

// GetId returns the Id field value
func (o *ForumPost) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ForumPost) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ForumPost) SetId(v int32) {
	o.Id = v
}

// GetTopicId returns the TopicId field value
func (o *ForumPost) GetTopicId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TopicId
}

// GetTopicIdOk returns a tuple with the TopicId field value
// and a boolean to check if the value has been set.
func (o *ForumPost) GetTopicIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TopicId, true
}

// SetTopicId sets field value
func (o *ForumPost) SetTopicId(v int32) {
	o.TopicId = v
}

// GetCreatorId returns the CreatorId field value
func (o *ForumPost) GetCreatorId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value
// and a boolean to check if the value has been set.
func (o *ForumPost) GetCreatorIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorId, true
}

// SetCreatorId sets field value
func (o *ForumPost) SetCreatorId(v int32) {
	o.CreatorId = v
}

// GetUpdaterId returns the UpdaterId field value
func (o *ForumPost) GetUpdaterId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdaterId
}

// GetUpdaterIdOk returns a tuple with the UpdaterId field value
// and a boolean to check if the value has been set.
func (o *ForumPost) GetUpdaterIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdaterId, true
}

// SetUpdaterId sets field value
func (o *ForumPost) SetUpdaterId(v int32) {
	o.UpdaterId = v
}

// GetBody returns the Body field value
func (o *ForumPost) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *ForumPost) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *ForumPost) SetBody(v string) {
	o.Body = v
}

// GetIsHidden returns the IsHidden field value
func (o *ForumPost) GetIsHidden() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsHidden
}

// GetIsHiddenOk returns a tuple with the IsHidden field value
// and a boolean to check if the value has been set.
func (o *ForumPost) GetIsHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsHidden, true
}

// SetIsHidden sets field value
func (o *ForumPost) SetIsHidden(v bool) {
	o.IsHidden = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ForumPost) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ForumPost) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ForumPost) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ForumPost) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ForumPost) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ForumPost) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetWarningType returns the WarningType field value
func (o *ForumPost) GetWarningType() WarningTypes {
	if o == nil {
		var ret WarningTypes
		return ret
	}

	return o.WarningType
}

// GetWarningTypeOk returns a tuple with the WarningType field value
// and a boolean to check if the value has been set.
func (o *ForumPost) GetWarningTypeOk() (*WarningTypes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WarningType, true
}

// SetWarningType sets field value
func (o *ForumPost) SetWarningType(v WarningTypes) {
	o.WarningType = v
}

// GetWarningUserId returns the WarningUserId field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *ForumPost) GetWarningUserId() float32 {
	if o == nil || o.WarningUserId.Get() == nil {
		var ret float32
		return ret
	}

	return *o.WarningUserId.Get()
}

// GetWarningUserIdOk returns a tuple with the WarningUserId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ForumPost) GetWarningUserIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.WarningUserId.Get(), o.WarningUserId.IsSet()
}

// SetWarningUserId sets field value
func (o *ForumPost) SetWarningUserId(v float32) {
	o.WarningUserId.Set(&v)
}

func (o ForumPost) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForumPost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["topic_id"] = o.TopicId
	toSerialize["creator_id"] = o.CreatorId
	toSerialize["updater_id"] = o.UpdaterId
	toSerialize["body"] = o.Body
	toSerialize["is_hidden"] = o.IsHidden
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["warning_type"] = o.WarningType
	toSerialize["warning_user_id"] = o.WarningUserId.Get()
	return toSerialize, nil
}

func (o *ForumPost) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"topic_id",
		"creator_id",
		"updater_id",
		"body",
		"is_hidden",
		"created_at",
		"updated_at",
		"warning_type",
		"warning_user_id",
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

	varForumPost := _ForumPost{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varForumPost)

	if err != nil {
		return err
	}

	*o = ForumPost(varForumPost)

	return err
}

type NullableForumPost struct {
	value *ForumPost
	isSet bool
}

func (v NullableForumPost) Get() *ForumPost {
	return v.value
}

func (v *NullableForumPost) Set(val *ForumPost) {
	v.value = val
	v.isSet = true
}

func (v NullableForumPost) IsSet() bool {
	return v.isSet
}

func (v *NullableForumPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForumPost(val *ForumPost) *NullableForumPost {
	return &NullableForumPost{value: val, isSet: true}
}

func (v NullableForumPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForumPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
