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

// checks if the UserFeedback type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserFeedback{}

// UserFeedback struct for UserFeedback
type UserFeedback struct {
	Id        int32              `json:"id"`
	UserId    int32              `json:"user_id"`
	CreatorId int32              `json:"creator_id"`
	Category  FeedbackCategories `json:"category"`
	Body      string             `json:"body"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	UpdaterId int32              `json:"updater_id"`
	IsDeleted bool               `json:"is_deleted"`
}

type _UserFeedback UserFeedback

// NewUserFeedback instantiates a new UserFeedback object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFeedback(id int32, userId int32, creatorId int32, category FeedbackCategories, body string, createdAt time.Time, updatedAt time.Time, updaterId int32, isDeleted bool) *UserFeedback {
	this := UserFeedback{}
	this.Id = id
	this.UserId = userId
	this.CreatorId = creatorId
	this.Category = category
	this.Body = body
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.UpdaterId = updaterId
	this.IsDeleted = isDeleted
	return &this
}

// NewUserFeedbackWithDefaults instantiates a new UserFeedback object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFeedbackWithDefaults() *UserFeedback {
	this := UserFeedback{}
	return &this
}

// GetId returns the Id field value
func (o *UserFeedback) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserFeedback) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserFeedback) SetId(v int32) {
	o.Id = v
}

// GetUserId returns the UserId field value
func (o *UserFeedback) GetUserId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UserFeedback) GetUserIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UserFeedback) SetUserId(v int32) {
	o.UserId = v
}

// GetCreatorId returns the CreatorId field value
func (o *UserFeedback) GetCreatorId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value
// and a boolean to check if the value has been set.
func (o *UserFeedback) GetCreatorIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorId, true
}

// SetCreatorId sets field value
func (o *UserFeedback) SetCreatorId(v int32) {
	o.CreatorId = v
}

// GetCategory returns the Category field value
func (o *UserFeedback) GetCategory() FeedbackCategories {
	if o == nil {
		var ret FeedbackCategories
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *UserFeedback) GetCategoryOk() (*FeedbackCategories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *UserFeedback) SetCategory(v FeedbackCategories) {
	o.Category = v
}

// GetBody returns the Body field value
func (o *UserFeedback) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *UserFeedback) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *UserFeedback) SetBody(v string) {
	o.Body = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *UserFeedback) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *UserFeedback) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *UserFeedback) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *UserFeedback) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *UserFeedback) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *UserFeedback) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetUpdaterId returns the UpdaterId field value
func (o *UserFeedback) GetUpdaterId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdaterId
}

// GetUpdaterIdOk returns a tuple with the UpdaterId field value
// and a boolean to check if the value has been set.
func (o *UserFeedback) GetUpdaterIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdaterId, true
}

// SetUpdaterId sets field value
func (o *UserFeedback) SetUpdaterId(v int32) {
	o.UpdaterId = v
}

// GetIsDeleted returns the IsDeleted field value
func (o *UserFeedback) GetIsDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value
// and a boolean to check if the value has been set.
func (o *UserFeedback) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeleted, true
}

// SetIsDeleted sets field value
func (o *UserFeedback) SetIsDeleted(v bool) {
	o.IsDeleted = v
}

func (o UserFeedback) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserFeedback) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["user_id"] = o.UserId
	toSerialize["creator_id"] = o.CreatorId
	toSerialize["category"] = o.Category
	toSerialize["body"] = o.Body
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["updater_id"] = o.UpdaterId
	toSerialize["is_deleted"] = o.IsDeleted
	return toSerialize, nil
}

func (o *UserFeedback) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"user_id",
		"creator_id",
		"category",
		"body",
		"created_at",
		"updated_at",
		"updater_id",
		"is_deleted",
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

	varUserFeedback := _UserFeedback{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserFeedback)

	if err != nil {
		return err
	}

	*o = UserFeedback(varUserFeedback)

	return err
}

type NullableUserFeedback struct {
	value *UserFeedback
	isSet bool
}

func (v NullableUserFeedback) Get() *UserFeedback {
	return v.value
}

func (v *NullableUserFeedback) Set(val *UserFeedback) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFeedback) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFeedback) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFeedback(val *UserFeedback) *NullableUserFeedback {
	return &NullableUserFeedback{value: val, isSet: true}
}

func (v NullableUserFeedback) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFeedback) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
