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

// checks if the Pool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Pool{}

// Pool struct for Pool
type Pool struct {
	Id          int32          `json:"id"`
	Name        string         `json:"name"`
	UpdatedAt   time.Time      `json:"updated_at"`
	CreatorId   int32          `json:"creator_id"`
	Description string         `json:"description"`
	IsActive    bool           `json:"is_active"`
	Category    PoolCategories `json:"category"`
	PostIds     []int32        `json:"post_ids"`
	CreatedAt   time.Time      `json:"created_at"`
	CreatorName string         `json:"creator_name"`
	PostCount   int32          `json:"post_count"`
}

type _Pool Pool

// NewPool instantiates a new Pool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPool(id int32, name string, updatedAt time.Time, creatorId int32, description string, isActive bool, category PoolCategories, postIds []int32, createdAt time.Time, creatorName string, postCount int32) *Pool {
	this := Pool{}
	this.Id = id
	this.Name = name
	this.UpdatedAt = updatedAt
	this.CreatorId = creatorId
	this.Description = description
	this.IsActive = isActive
	this.Category = category
	this.PostIds = postIds
	this.CreatedAt = createdAt
	this.CreatorName = creatorName
	this.PostCount = postCount
	return &this
}

// NewPoolWithDefaults instantiates a new Pool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolWithDefaults() *Pool {
	this := Pool{}
	return &this
}

// GetId returns the Id field value
func (o *Pool) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Pool) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Pool) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Pool) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Pool) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Pool) SetName(v string) {
	o.Name = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Pool) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Pool) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Pool) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetCreatorId returns the CreatorId field value
func (o *Pool) GetCreatorId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value
// and a boolean to check if the value has been set.
func (o *Pool) GetCreatorIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorId, true
}

// SetCreatorId sets field value
func (o *Pool) SetCreatorId(v int32) {
	o.CreatorId = v
}

// GetDescription returns the Description field value
func (o *Pool) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Pool) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Pool) SetDescription(v string) {
	o.Description = v
}

// GetIsActive returns the IsActive field value
func (o *Pool) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *Pool) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *Pool) SetIsActive(v bool) {
	o.IsActive = v
}

// GetCategory returns the Category field value
func (o *Pool) GetCategory() PoolCategories {
	if o == nil {
		var ret PoolCategories
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *Pool) GetCategoryOk() (*PoolCategories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *Pool) SetCategory(v PoolCategories) {
	o.Category = v
}

// GetPostIds returns the PostIds field value
func (o *Pool) GetPostIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.PostIds
}

// GetPostIdsOk returns a tuple with the PostIds field value
// and a boolean to check if the value has been set.
func (o *Pool) GetPostIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostIds, true
}

// SetPostIds sets field value
func (o *Pool) SetPostIds(v []int32) {
	o.PostIds = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Pool) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Pool) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Pool) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatorName returns the CreatorName field value
func (o *Pool) GetCreatorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatorName
}

// GetCreatorNameOk returns a tuple with the CreatorName field value
// and a boolean to check if the value has been set.
func (o *Pool) GetCreatorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorName, true
}

// SetCreatorName sets field value
func (o *Pool) SetCreatorName(v string) {
	o.CreatorName = v
}

// GetPostCount returns the PostCount field value
func (o *Pool) GetPostCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PostCount
}

// GetPostCountOk returns a tuple with the PostCount field value
// and a boolean to check if the value has been set.
func (o *Pool) GetPostCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostCount, true
}

// SetPostCount sets field value
func (o *Pool) SetPostCount(v int32) {
	o.PostCount = v
}

func (o Pool) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Pool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["creator_id"] = o.CreatorId
	toSerialize["description"] = o.Description
	toSerialize["is_active"] = o.IsActive
	toSerialize["category"] = o.Category
	toSerialize["post_ids"] = o.PostIds
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["creator_name"] = o.CreatorName
	toSerialize["post_count"] = o.PostCount
	return toSerialize, nil
}

func (o *Pool) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"updated_at",
		"creator_id",
		"description",
		"is_active",
		"category",
		"post_ids",
		"created_at",
		"creator_name",
		"post_count",
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

	varPool := _Pool{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPool)

	if err != nil {
		return err
	}

	*o = Pool(varPool)

	return err
}

type NullablePool struct {
	value *Pool
	isSet bool
}

func (v NullablePool) Get() *Pool {
	return v.value
}

func (v *NullablePool) Set(val *Pool) {
	v.value = val
	v.isSet = true
}

func (v NullablePool) IsSet() bool {
	return v.isSet
}

func (v *NullablePool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePool(val *Pool) *NullablePool {
	return &NullablePool{value: val, isSet: true}
}

func (v NullablePool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
