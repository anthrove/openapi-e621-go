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

// checks if the PoolVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PoolVersion{}

// PoolVersion struct for PoolVersion
type PoolVersion struct {
	Id                 int32          `json:"id"`
	PoolId             int32          `json:"pool_id"`
	PostIds            []int32        `json:"post_ids"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	UpdaterId          int32          `json:"updater_id"`
	Name               string         `json:"name"`
	NameChanged        bool           `json:"name_changed"`
	Description        string         `json:"description"`
	DescriptionChanged bool           `json:"description_changed"`
	IsActive           bool           `json:"is_active"`
	IsLocked           bool           `json:"is_locked"`
	Category           PoolCategories `json:"category"`
	Version            float32        `json:"version"`
	AddedPostIds       []int32        `json:"added_post_ids"`
	RemovedPostIds     []int32        `json:"removed_post_ids"`
}

type _PoolVersion PoolVersion

// NewPoolVersion instantiates a new PoolVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoolVersion(id int32, poolId int32, postIds []int32, createdAt time.Time, updatedAt time.Time, updaterId int32, name string, nameChanged bool, description string, descriptionChanged bool, isActive bool, isLocked bool, category PoolCategories, version float32, addedPostIds []int32, removedPostIds []int32) *PoolVersion {
	this := PoolVersion{}
	this.Id = id
	this.PoolId = poolId
	this.PostIds = postIds
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.UpdaterId = updaterId
	this.Name = name
	this.NameChanged = nameChanged
	this.Description = description
	this.DescriptionChanged = descriptionChanged
	this.IsActive = isActive
	this.IsLocked = isLocked
	this.Category = category
	this.Version = version
	this.AddedPostIds = addedPostIds
	this.RemovedPostIds = removedPostIds
	return &this
}

// NewPoolVersionWithDefaults instantiates a new PoolVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPoolVersionWithDefaults() *PoolVersion {
	this := PoolVersion{}
	return &this
}

// GetId returns the Id field value
func (o *PoolVersion) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PoolVersion) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PoolVersion) SetId(v int32) {
	o.Id = v
}

// GetPoolId returns the PoolId field value
func (o *PoolVersion) GetPoolId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value
// and a boolean to check if the value has been set.
func (o *PoolVersion) GetPoolIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolId, true
}

// SetPoolId sets field value
func (o *PoolVersion) SetPoolId(v int32) {
	o.PoolId = v
}

// GetPostIds returns the PostIds field value
func (o *PoolVersion) GetPostIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.PostIds
}

// GetPostIdsOk returns a tuple with the PostIds field value
// and a boolean to check if the value has been set.
func (o *PoolVersion) GetPostIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostIds, true
}

// SetPostIds sets field value
func (o *PoolVersion) SetPostIds(v []int32) {
	o.PostIds = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PoolVersion) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PoolVersion) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PoolVersion) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PoolVersion) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PoolVersion) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PoolVersion) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetUpdaterId returns the UpdaterId field value
func (o *PoolVersion) GetUpdaterId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdaterId
}

// GetUpdaterIdOk returns a tuple with the UpdaterId field value
// and a boolean to check if the value has been set.
func (o *PoolVersion) GetUpdaterIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdaterId, true
}

// SetUpdaterId sets field value
func (o *PoolVersion) SetUpdaterId(v int32) {
	o.UpdaterId = v
}

// GetName returns the Name field value
func (o *PoolVersion) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PoolVersion) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PoolVersion) SetName(v string) {
	o.Name = v
}

// GetNameChanged returns the NameChanged field value
func (o *PoolVersion) GetNameChanged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NameChanged
}

// GetNameChangedOk returns a tuple with the NameChanged field value
// and a boolean to check if the value has been set.
func (o *PoolVersion) GetNameChangedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameChanged, true
}

// SetNameChanged sets field value
func (o *PoolVersion) SetNameChanged(v bool) {
	o.NameChanged = v
}

// GetDescription returns the Description field value
func (o *PoolVersion) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PoolVersion) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PoolVersion) SetDescription(v string) {
	o.Description = v
}

// GetDescriptionChanged returns the DescriptionChanged field value
func (o *PoolVersion) GetDescriptionChanged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DescriptionChanged
}

// GetDescriptionChangedOk returns a tuple with the DescriptionChanged field value
// and a boolean to check if the value has been set.
func (o *PoolVersion) GetDescriptionChangedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DescriptionChanged, true
}

// SetDescriptionChanged sets field value
func (o *PoolVersion) SetDescriptionChanged(v bool) {
	o.DescriptionChanged = v
}

// GetIsActive returns the IsActive field value
func (o *PoolVersion) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *PoolVersion) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *PoolVersion) SetIsActive(v bool) {
	o.IsActive = v
}

// GetIsLocked returns the IsLocked field value
func (o *PoolVersion) GetIsLocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value
// and a boolean to check if the value has been set.
func (o *PoolVersion) GetIsLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLocked, true
}

// SetIsLocked sets field value
func (o *PoolVersion) SetIsLocked(v bool) {
	o.IsLocked = v
}

// GetCategory returns the Category field value
func (o *PoolVersion) GetCategory() PoolCategories {
	if o == nil {
		var ret PoolCategories
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *PoolVersion) GetCategoryOk() (*PoolCategories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *PoolVersion) SetCategory(v PoolCategories) {
	o.Category = v
}

// GetVersion returns the Version field value
func (o *PoolVersion) GetVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *PoolVersion) GetVersionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *PoolVersion) SetVersion(v float32) {
	o.Version = v
}

// GetAddedPostIds returns the AddedPostIds field value
func (o *PoolVersion) GetAddedPostIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.AddedPostIds
}

// GetAddedPostIdsOk returns a tuple with the AddedPostIds field value
// and a boolean to check if the value has been set.
func (o *PoolVersion) GetAddedPostIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddedPostIds, true
}

// SetAddedPostIds sets field value
func (o *PoolVersion) SetAddedPostIds(v []int32) {
	o.AddedPostIds = v
}

// GetRemovedPostIds returns the RemovedPostIds field value
func (o *PoolVersion) GetRemovedPostIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.RemovedPostIds
}

// GetRemovedPostIdsOk returns a tuple with the RemovedPostIds field value
// and a boolean to check if the value has been set.
func (o *PoolVersion) GetRemovedPostIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemovedPostIds, true
}

// SetRemovedPostIds sets field value
func (o *PoolVersion) SetRemovedPostIds(v []int32) {
	o.RemovedPostIds = v
}

func (o PoolVersion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PoolVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["pool_id"] = o.PoolId
	toSerialize["post_ids"] = o.PostIds
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["updater_id"] = o.UpdaterId
	toSerialize["name"] = o.Name
	toSerialize["name_changed"] = o.NameChanged
	toSerialize["description"] = o.Description
	toSerialize["description_changed"] = o.DescriptionChanged
	toSerialize["is_active"] = o.IsActive
	toSerialize["is_locked"] = o.IsLocked
	toSerialize["category"] = o.Category
	toSerialize["version"] = o.Version
	toSerialize["added_post_ids"] = o.AddedPostIds
	toSerialize["removed_post_ids"] = o.RemovedPostIds
	return toSerialize, nil
}

func (o *PoolVersion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"pool_id",
		"post_ids",
		"created_at",
		"updated_at",
		"updater_id",
		"name",
		"name_changed",
		"description",
		"description_changed",
		"is_active",
		"is_locked",
		"category",
		"version",
		"added_post_ids",
		"removed_post_ids",
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

	varPoolVersion := _PoolVersion{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPoolVersion)

	if err != nil {
		return err
	}

	*o = PoolVersion(varPoolVersion)

	return err
}

type NullablePoolVersion struct {
	value *PoolVersion
	isSet bool
}

func (v NullablePoolVersion) Get() *PoolVersion {
	return v.value
}

func (v *NullablePoolVersion) Set(val *PoolVersion) {
	v.value = val
	v.isSet = true
}

func (v NullablePoolVersion) IsSet() bool {
	return v.isSet
}

func (v *NullablePoolVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoolVersion(val *PoolVersion) *NullablePoolVersion {
	return &NullablePoolVersion{value: val, isSet: true}
}

func (v NullablePoolVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoolVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
