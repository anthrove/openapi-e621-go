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

// checks if the Tag type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tag{}

// Tag struct for Tag
type Tag struct {
	Id                   int32         `json:"id"`
	Name                 string        `json:"name"`
	PostCount            int32         `json:"post_count"`
	RelatedTags          []string      `json:"related_tags"`
	RelatedTagsUpdatedAt NullableTime  `json:"related_tags_updated_at"`
	Category             TagCategories `json:"category"`
	IsLocked             bool          `json:"is_locked"`
	CreatedAt            time.Time     `json:"created_at"`
	UpdatedAt            time.Time     `json:"updated_at"`
}

type _Tag Tag

// NewTag instantiates a new Tag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTag(id int32, name string, postCount int32, relatedTags []string, relatedTagsUpdatedAt NullableTime, category TagCategories, isLocked bool, createdAt time.Time, updatedAt time.Time) *Tag {
	this := Tag{}
	this.Id = id
	this.Name = name
	this.PostCount = postCount
	this.RelatedTags = relatedTags
	this.RelatedTagsUpdatedAt = relatedTagsUpdatedAt
	this.Category = category
	this.IsLocked = isLocked
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewTagWithDefaults instantiates a new Tag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagWithDefaults() *Tag {
	this := Tag{}
	return &this
}

// GetId returns the Id field value
func (o *Tag) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Tag) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Tag) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Tag) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Tag) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Tag) SetName(v string) {
	o.Name = v
}

// GetPostCount returns the PostCount field value
func (o *Tag) GetPostCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PostCount
}

// GetPostCountOk returns a tuple with the PostCount field value
// and a boolean to check if the value has been set.
func (o *Tag) GetPostCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostCount, true
}

// SetPostCount sets field value
func (o *Tag) SetPostCount(v int32) {
	o.PostCount = v
}

// GetRelatedTags returns the RelatedTags field value
func (o *Tag) GetRelatedTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RelatedTags
}

// GetRelatedTagsOk returns a tuple with the RelatedTags field value
// and a boolean to check if the value has been set.
func (o *Tag) GetRelatedTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RelatedTags, true
}

// SetRelatedTags sets field value
func (o *Tag) SetRelatedTags(v []string) {
	o.RelatedTags = v
}

// GetRelatedTagsUpdatedAt returns the RelatedTagsUpdatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Tag) GetRelatedTagsUpdatedAt() time.Time {
	if o == nil || o.RelatedTagsUpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.RelatedTagsUpdatedAt.Get()
}

// GetRelatedTagsUpdatedAtOk returns a tuple with the RelatedTagsUpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Tag) GetRelatedTagsUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RelatedTagsUpdatedAt.Get(), o.RelatedTagsUpdatedAt.IsSet()
}

// SetRelatedTagsUpdatedAt sets field value
func (o *Tag) SetRelatedTagsUpdatedAt(v time.Time) {
	o.RelatedTagsUpdatedAt.Set(&v)
}

// GetCategory returns the Category field value
func (o *Tag) GetCategory() TagCategories {
	if o == nil {
		var ret TagCategories
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *Tag) GetCategoryOk() (*TagCategories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *Tag) SetCategory(v TagCategories) {
	o.Category = v
}

// GetIsLocked returns the IsLocked field value
func (o *Tag) GetIsLocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value
// and a boolean to check if the value has been set.
func (o *Tag) GetIsLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLocked, true
}

// SetIsLocked sets field value
func (o *Tag) SetIsLocked(v bool) {
	o.IsLocked = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Tag) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Tag) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Tag) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Tag) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Tag) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Tag) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o Tag) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tag) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["post_count"] = o.PostCount
	toSerialize["related_tags"] = o.RelatedTags
	toSerialize["related_tags_updated_at"] = o.RelatedTagsUpdatedAt.Get()
	toSerialize["category"] = o.Category
	toSerialize["is_locked"] = o.IsLocked
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *Tag) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"post_count",
		"related_tags",
		"related_tags_updated_at",
		"category",
		"is_locked",
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

	varTag := _Tag{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTag)

	if err != nil {
		return err
	}

	*o = Tag(varTag)

	return err
}

type NullableTag struct {
	value *Tag
	isSet bool
}

func (v NullableTag) Get() *Tag {
	return v.value
}

func (v *NullableTag) Set(val *Tag) {
	v.value = val
	v.isSet = true
}

func (v NullableTag) IsSet() bool {
	return v.isSet
}

func (v *NullableTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTag(val *Tag) *NullableTag {
	return &NullableTag{value: val, isSet: true}
}

func (v NullableTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
