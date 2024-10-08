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

// checks if the AvoidPosting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AvoidPosting{}

// AvoidPosting struct for AvoidPosting
type AvoidPosting struct {
	Id float32 `json:"id"`
	CreatorId float32 `json:"creator_id"`
	UpdaterId float32 `json:"updater_id"`
	ArtistId float32 `json:"artist_id"`
	// Only visible to Janitor+
	StaffNotes *string `json:"staff_notes,omitempty"`
	Details string `json:"details"`
	IsActive bool `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type _AvoidPosting AvoidPosting

// NewAvoidPosting instantiates a new AvoidPosting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvoidPosting(id float32, creatorId float32, updaterId float32, artistId float32, details string, isActive bool, createdAt time.Time, updatedAt time.Time) *AvoidPosting {
	this := AvoidPosting{}
	this.Id = id
	this.CreatorId = creatorId
	this.UpdaterId = updaterId
	this.ArtistId = artistId
	this.Details = details
	this.IsActive = isActive
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewAvoidPostingWithDefaults instantiates a new AvoidPosting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvoidPostingWithDefaults() *AvoidPosting {
	this := AvoidPosting{}
	return &this
}

// GetId returns the Id field value
func (o *AvoidPosting) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AvoidPosting) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AvoidPosting) SetId(v float32) {
	o.Id = v
}

// GetCreatorId returns the CreatorId field value
func (o *AvoidPosting) GetCreatorId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value
// and a boolean to check if the value has been set.
func (o *AvoidPosting) GetCreatorIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorId, true
}

// SetCreatorId sets field value
func (o *AvoidPosting) SetCreatorId(v float32) {
	o.CreatorId = v
}

// GetUpdaterId returns the UpdaterId field value
func (o *AvoidPosting) GetUpdaterId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UpdaterId
}

// GetUpdaterIdOk returns a tuple with the UpdaterId field value
// and a boolean to check if the value has been set.
func (o *AvoidPosting) GetUpdaterIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdaterId, true
}

// SetUpdaterId sets field value
func (o *AvoidPosting) SetUpdaterId(v float32) {
	o.UpdaterId = v
}

// GetArtistId returns the ArtistId field value
func (o *AvoidPosting) GetArtistId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ArtistId
}

// GetArtistIdOk returns a tuple with the ArtistId field value
// and a boolean to check if the value has been set.
func (o *AvoidPosting) GetArtistIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArtistId, true
}

// SetArtistId sets field value
func (o *AvoidPosting) SetArtistId(v float32) {
	o.ArtistId = v
}

// GetStaffNotes returns the StaffNotes field value if set, zero value otherwise.
func (o *AvoidPosting) GetStaffNotes() string {
	if o == nil || IsNil(o.StaffNotes) {
		var ret string
		return ret
	}
	return *o.StaffNotes
}

// GetStaffNotesOk returns a tuple with the StaffNotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AvoidPosting) GetStaffNotesOk() (*string, bool) {
	if o == nil || IsNil(o.StaffNotes) {
		return nil, false
	}
	return o.StaffNotes, true
}

// HasStaffNotes returns a boolean if a field has been set.
func (o *AvoidPosting) HasStaffNotes() bool {
	if o != nil && !IsNil(o.StaffNotes) {
		return true
	}

	return false
}

// SetStaffNotes gets a reference to the given string and assigns it to the StaffNotes field.
func (o *AvoidPosting) SetStaffNotes(v string) {
	o.StaffNotes = &v
}

// GetDetails returns the Details field value
func (o *AvoidPosting) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *AvoidPosting) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *AvoidPosting) SetDetails(v string) {
	o.Details = v
}

// GetIsActive returns the IsActive field value
func (o *AvoidPosting) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *AvoidPosting) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *AvoidPosting) SetIsActive(v bool) {
	o.IsActive = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AvoidPosting) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AvoidPosting) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AvoidPosting) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *AvoidPosting) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *AvoidPosting) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *AvoidPosting) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o AvoidPosting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AvoidPosting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["creator_id"] = o.CreatorId
	toSerialize["updater_id"] = o.UpdaterId
	toSerialize["artist_id"] = o.ArtistId
	if !IsNil(o.StaffNotes) {
		toSerialize["staff_notes"] = o.StaffNotes
	}
	toSerialize["details"] = o.Details
	toSerialize["is_active"] = o.IsActive
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *AvoidPosting) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"creator_id",
		"updater_id",
		"artist_id",
		"details",
		"is_active",
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

	varAvoidPosting := _AvoidPosting{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAvoidPosting)

	if err != nil {
		return err
	}

	*o = AvoidPosting(varAvoidPosting)

	return err
}

type NullableAvoidPosting struct {
	value *AvoidPosting
	isSet bool
}

func (v NullableAvoidPosting) Get() *AvoidPosting {
	return v.value
}

func (v *NullableAvoidPosting) Set(val *AvoidPosting) {
	v.value = val
	v.isSet = true
}

func (v NullableAvoidPosting) IsSet() bool {
	return v.isSet
}

func (v *NullableAvoidPosting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvoidPosting(val *AvoidPosting) *NullableAvoidPosting {
	return &NullableAvoidPosting{value: val, isSet: true}
}

func (v NullableAvoidPosting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvoidPosting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


