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

// checks if the User type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &User{}

// User struct for User
type User struct {
	Id              float32         `json:"id"`
	CreatedAt       time.Time       `json:"created_at"`
	Name            string          `json:"name"`
	Level           float32         `json:"level"`
	BaseUploadLimit float32         `json:"base_upload_limit"`
	PostUploadCount float32         `json:"post_upload_count"`
	PostUpdateCount float32         `json:"post_update_count"`
	NoteUpdateCount float32         `json:"note_update_count"`
	IsBanned        bool            `json:"is_banned"`
	CanApprovePosts bool            `json:"can_approve_posts"`
	CanUploadFree   bool            `json:"can_upload_free"`
	LevelString     string          `json:"level_string"`
	AvatarId        NullableFloat32 `json:"avatar_id"`
}

type _User User

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(id float32, createdAt time.Time, name string, level float32, baseUploadLimit float32, postUploadCount float32, postUpdateCount float32, noteUpdateCount float32, isBanned bool, canApprovePosts bool, canUploadFree bool, levelString string, avatarId NullableFloat32) *User {
	this := User{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Name = name
	this.Level = level
	this.BaseUploadLimit = baseUploadLimit
	this.PostUploadCount = postUploadCount
	this.PostUpdateCount = postUpdateCount
	this.NoteUpdateCount = noteUpdateCount
	this.IsBanned = isBanned
	this.CanApprovePosts = canApprovePosts
	this.CanUploadFree = canUploadFree
	this.LevelString = levelString
	this.AvatarId = avatarId
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetId returns the Id field value
func (o *User) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *User) SetId(v float32) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *User) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *User) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *User) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetName returns the Name field value
func (o *User) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *User) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *User) SetName(v string) {
	o.Name = v
}

// GetLevel returns the Level field value
func (o *User) GetLevel() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *User) GetLevelOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *User) SetLevel(v float32) {
	o.Level = v
}

// GetBaseUploadLimit returns the BaseUploadLimit field value
func (o *User) GetBaseUploadLimit() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.BaseUploadLimit
}

// GetBaseUploadLimitOk returns a tuple with the BaseUploadLimit field value
// and a boolean to check if the value has been set.
func (o *User) GetBaseUploadLimitOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseUploadLimit, true
}

// SetBaseUploadLimit sets field value
func (o *User) SetBaseUploadLimit(v float32) {
	o.BaseUploadLimit = v
}

// GetPostUploadCount returns the PostUploadCount field value
func (o *User) GetPostUploadCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PostUploadCount
}

// GetPostUploadCountOk returns a tuple with the PostUploadCount field value
// and a boolean to check if the value has been set.
func (o *User) GetPostUploadCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostUploadCount, true
}

// SetPostUploadCount sets field value
func (o *User) SetPostUploadCount(v float32) {
	o.PostUploadCount = v
}

// GetPostUpdateCount returns the PostUpdateCount field value
func (o *User) GetPostUpdateCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PostUpdateCount
}

// GetPostUpdateCountOk returns a tuple with the PostUpdateCount field value
// and a boolean to check if the value has been set.
func (o *User) GetPostUpdateCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostUpdateCount, true
}

// SetPostUpdateCount sets field value
func (o *User) SetPostUpdateCount(v float32) {
	o.PostUpdateCount = v
}

// GetNoteUpdateCount returns the NoteUpdateCount field value
func (o *User) GetNoteUpdateCount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NoteUpdateCount
}

// GetNoteUpdateCountOk returns a tuple with the NoteUpdateCount field value
// and a boolean to check if the value has been set.
func (o *User) GetNoteUpdateCountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NoteUpdateCount, true
}

// SetNoteUpdateCount sets field value
func (o *User) SetNoteUpdateCount(v float32) {
	o.NoteUpdateCount = v
}

// GetIsBanned returns the IsBanned field value
func (o *User) GetIsBanned() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsBanned
}

// GetIsBannedOk returns a tuple with the IsBanned field value
// and a boolean to check if the value has been set.
func (o *User) GetIsBannedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsBanned, true
}

// SetIsBanned sets field value
func (o *User) SetIsBanned(v bool) {
	o.IsBanned = v
}

// GetCanApprovePosts returns the CanApprovePosts field value
func (o *User) GetCanApprovePosts() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanApprovePosts
}

// GetCanApprovePostsOk returns a tuple with the CanApprovePosts field value
// and a boolean to check if the value has been set.
func (o *User) GetCanApprovePostsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanApprovePosts, true
}

// SetCanApprovePosts sets field value
func (o *User) SetCanApprovePosts(v bool) {
	o.CanApprovePosts = v
}

// GetCanUploadFree returns the CanUploadFree field value
func (o *User) GetCanUploadFree() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CanUploadFree
}

// GetCanUploadFreeOk returns a tuple with the CanUploadFree field value
// and a boolean to check if the value has been set.
func (o *User) GetCanUploadFreeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CanUploadFree, true
}

// SetCanUploadFree sets field value
func (o *User) SetCanUploadFree(v bool) {
	o.CanUploadFree = v
}

// GetLevelString returns the LevelString field value
func (o *User) GetLevelString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LevelString
}

// GetLevelStringOk returns a tuple with the LevelString field value
// and a boolean to check if the value has been set.
func (o *User) GetLevelStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LevelString, true
}

// SetLevelString sets field value
func (o *User) SetLevelString(v string) {
	o.LevelString = v
}

// GetAvatarId returns the AvatarId field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *User) GetAvatarId() float32 {
	if o == nil || o.AvatarId.Get() == nil {
		var ret float32
		return ret
	}

	return *o.AvatarId.Get()
}

// GetAvatarIdOk returns a tuple with the AvatarId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetAvatarIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AvatarId.Get(), o.AvatarId.IsSet()
}

// SetAvatarId sets field value
func (o *User) SetAvatarId(v float32) {
	o.AvatarId.Set(&v)
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o User) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["name"] = o.Name
	toSerialize["level"] = o.Level
	toSerialize["base_upload_limit"] = o.BaseUploadLimit
	toSerialize["post_upload_count"] = o.PostUploadCount
	toSerialize["post_update_count"] = o.PostUpdateCount
	toSerialize["note_update_count"] = o.NoteUpdateCount
	toSerialize["is_banned"] = o.IsBanned
	toSerialize["can_approve_posts"] = o.CanApprovePosts
	toSerialize["can_upload_free"] = o.CanUploadFree
	toSerialize["level_string"] = o.LevelString
	toSerialize["avatar_id"] = o.AvatarId.Get()
	return toSerialize, nil
}

func (o *User) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"name",
		"level",
		"base_upload_limit",
		"post_upload_count",
		"post_update_count",
		"note_update_count",
		"is_banned",
		"can_approve_posts",
		"can_upload_free",
		"level_string",
		"avatar_id",
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

	varUser := _User{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUser)

	if err != nil {
		return err
	}

	*o = User(varUser)

	return err
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
