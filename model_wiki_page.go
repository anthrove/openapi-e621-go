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

// checks if the WikiPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WikiPage{}

// WikiPage struct for WikiPage
type WikiPage struct {
	Id float32 `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Title string `json:"title"`
	Body string `json:"body"`
	CreatorId float32 `json:"creator_id"`
	IsLocked bool `json:"is_locked"`
	UpdaterId float32 `json:"updater_id"`
	IsDeleted bool `json:"is_deleted"`
	OtherNames []string `json:"other_names"`
	Parent NullableString `json:"parent"`
	CreatorName string `json:"creator_name"`
	CategoryId TagCategories `json:"category_id"`
}

type _WikiPage WikiPage

// NewWikiPage instantiates a new WikiPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWikiPage(id float32, createdAt time.Time, updatedAt time.Time, title string, body string, creatorId float32, isLocked bool, updaterId float32, isDeleted bool, otherNames []string, parent NullableString, creatorName string, categoryId TagCategories) *WikiPage {
	this := WikiPage{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Title = title
	this.Body = body
	this.CreatorId = creatorId
	this.IsLocked = isLocked
	this.UpdaterId = updaterId
	this.IsDeleted = isDeleted
	this.OtherNames = otherNames
	this.Parent = parent
	this.CreatorName = creatorName
	this.CategoryId = categoryId
	return &this
}

// NewWikiPageWithDefaults instantiates a new WikiPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWikiPageWithDefaults() *WikiPage {
	this := WikiPage{}
	return &this
}

// GetId returns the Id field value
func (o *WikiPage) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WikiPage) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WikiPage) SetId(v float32) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *WikiPage) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WikiPage) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WikiPage) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *WikiPage) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *WikiPage) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *WikiPage) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetTitle returns the Title field value
func (o *WikiPage) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *WikiPage) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *WikiPage) SetTitle(v string) {
	o.Title = v
}

// GetBody returns the Body field value
func (o *WikiPage) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *WikiPage) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *WikiPage) SetBody(v string) {
	o.Body = v
}

// GetCreatorId returns the CreatorId field value
func (o *WikiPage) GetCreatorId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value
// and a boolean to check if the value has been set.
func (o *WikiPage) GetCreatorIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorId, true
}

// SetCreatorId sets field value
func (o *WikiPage) SetCreatorId(v float32) {
	o.CreatorId = v
}

// GetIsLocked returns the IsLocked field value
func (o *WikiPage) GetIsLocked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value
// and a boolean to check if the value has been set.
func (o *WikiPage) GetIsLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsLocked, true
}

// SetIsLocked sets field value
func (o *WikiPage) SetIsLocked(v bool) {
	o.IsLocked = v
}

// GetUpdaterId returns the UpdaterId field value
func (o *WikiPage) GetUpdaterId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UpdaterId
}

// GetUpdaterIdOk returns a tuple with the UpdaterId field value
// and a boolean to check if the value has been set.
func (o *WikiPage) GetUpdaterIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdaterId, true
}

// SetUpdaterId sets field value
func (o *WikiPage) SetUpdaterId(v float32) {
	o.UpdaterId = v
}

// GetIsDeleted returns the IsDeleted field value
func (o *WikiPage) GetIsDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDeleted
}

// GetIsDeletedOk returns a tuple with the IsDeleted field value
// and a boolean to check if the value has been set.
func (o *WikiPage) GetIsDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDeleted, true
}

// SetIsDeleted sets field value
func (o *WikiPage) SetIsDeleted(v bool) {
	o.IsDeleted = v
}

// GetOtherNames returns the OtherNames field value
func (o *WikiPage) GetOtherNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OtherNames
}

// GetOtherNamesOk returns a tuple with the OtherNames field value
// and a boolean to check if the value has been set.
func (o *WikiPage) GetOtherNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OtherNames, true
}

// SetOtherNames sets field value
func (o *WikiPage) SetOtherNames(v []string) {
	o.OtherNames = v
}

// GetParent returns the Parent field value
// If the value is explicit nil, the zero value for string will be returned
func (o *WikiPage) GetParent() string {
	if o == nil || o.Parent.Get() == nil {
		var ret string
		return ret
	}

	return *o.Parent.Get()
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WikiPage) GetParentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parent.Get(), o.Parent.IsSet()
}

// SetParent sets field value
func (o *WikiPage) SetParent(v string) {
	o.Parent.Set(&v)
}

// GetCreatorName returns the CreatorName field value
func (o *WikiPage) GetCreatorName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatorName
}

// GetCreatorNameOk returns a tuple with the CreatorName field value
// and a boolean to check if the value has been set.
func (o *WikiPage) GetCreatorNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorName, true
}

// SetCreatorName sets field value
func (o *WikiPage) SetCreatorName(v string) {
	o.CreatorName = v
}

// GetCategoryId returns the CategoryId field value
func (o *WikiPage) GetCategoryId() TagCategories {
	if o == nil {
		var ret TagCategories
		return ret
	}

	return o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value
// and a boolean to check if the value has been set.
func (o *WikiPage) GetCategoryIdOk() (*TagCategories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CategoryId, true
}

// SetCategoryId sets field value
func (o *WikiPage) SetCategoryId(v TagCategories) {
	o.CategoryId = v
}

func (o WikiPage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WikiPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["title"] = o.Title
	toSerialize["body"] = o.Body
	toSerialize["creator_id"] = o.CreatorId
	toSerialize["is_locked"] = o.IsLocked
	toSerialize["updater_id"] = o.UpdaterId
	toSerialize["is_deleted"] = o.IsDeleted
	toSerialize["other_names"] = o.OtherNames
	toSerialize["parent"] = o.Parent.Get()
	toSerialize["creator_name"] = o.CreatorName
	toSerialize["category_id"] = o.CategoryId
	return toSerialize, nil
}

func (o *WikiPage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"title",
		"body",
		"creator_id",
		"is_locked",
		"updater_id",
		"is_deleted",
		"other_names",
		"parent",
		"creator_name",
		"category_id",
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

	varWikiPage := _WikiPage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWikiPage)

	if err != nil {
		return err
	}

	*o = WikiPage(varWikiPage)

	return err
}

type NullableWikiPage struct {
	value *WikiPage
	isSet bool
}

func (v NullableWikiPage) Get() *WikiPage {
	return v.value
}

func (v *NullableWikiPage) Set(val *WikiPage) {
	v.value = val
	v.isSet = true
}

func (v NullableWikiPage) IsSet() bool {
	return v.isSet
}

func (v *NullableWikiPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWikiPage(val *WikiPage) *NullableWikiPage {
	return &NullableWikiPage{value: val, isSet: true}
}

func (v NullableWikiPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWikiPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


