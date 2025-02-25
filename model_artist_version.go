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

// checks if the ArtistVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArtistVersion{}

// ArtistVersion struct for ArtistVersion
type ArtistVersion struct {
	Id           int32     `json:"id"`
	ArtistId     int32     `json:"artist_id"`
	Name         string    `json:"name"`
	UpdaterId    int32     `json:"updater_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	IsActive     bool      `json:"is_active"`
	OtherNames   []string  `json:"other_names"`
	NotesChanged bool      `json:"notes_changed"`
	Urls         []string  `json:"urls"`
}

type _ArtistVersion ArtistVersion

// NewArtistVersion instantiates a new ArtistVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtistVersion(id int32, artistId int32, name string, updaterId int32, createdAt time.Time, updatedAt time.Time, isActive bool, otherNames []string, notesChanged bool, urls []string) *ArtistVersion {
	this := ArtistVersion{}
	this.Id = id
	this.ArtistId = artistId
	this.Name = name
	this.UpdaterId = updaterId
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.IsActive = isActive
	this.OtherNames = otherNames
	this.NotesChanged = notesChanged
	this.Urls = urls
	return &this
}

// NewArtistVersionWithDefaults instantiates a new ArtistVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtistVersionWithDefaults() *ArtistVersion {
	this := ArtistVersion{}
	return &this
}

// GetId returns the Id field value
func (o *ArtistVersion) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArtistVersion) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArtistVersion) SetId(v int32) {
	o.Id = v
}

// GetArtistId returns the ArtistId field value
func (o *ArtistVersion) GetArtistId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ArtistId
}

// GetArtistIdOk returns a tuple with the ArtistId field value
// and a boolean to check if the value has been set.
func (o *ArtistVersion) GetArtistIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ArtistId, true
}

// SetArtistId sets field value
func (o *ArtistVersion) SetArtistId(v int32) {
	o.ArtistId = v
}

// GetName returns the Name field value
func (o *ArtistVersion) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArtistVersion) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArtistVersion) SetName(v string) {
	o.Name = v
}

// GetUpdaterId returns the UpdaterId field value
func (o *ArtistVersion) GetUpdaterId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdaterId
}

// GetUpdaterIdOk returns a tuple with the UpdaterId field value
// and a boolean to check if the value has been set.
func (o *ArtistVersion) GetUpdaterIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdaterId, true
}

// SetUpdaterId sets field value
func (o *ArtistVersion) SetUpdaterId(v int32) {
	o.UpdaterId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ArtistVersion) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ArtistVersion) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ArtistVersion) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ArtistVersion) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ArtistVersion) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ArtistVersion) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetIsActive returns the IsActive field value
func (o *ArtistVersion) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *ArtistVersion) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *ArtistVersion) SetIsActive(v bool) {
	o.IsActive = v
}

// GetOtherNames returns the OtherNames field value
func (o *ArtistVersion) GetOtherNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OtherNames
}

// GetOtherNamesOk returns a tuple with the OtherNames field value
// and a boolean to check if the value has been set.
func (o *ArtistVersion) GetOtherNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OtherNames, true
}

// SetOtherNames sets field value
func (o *ArtistVersion) SetOtherNames(v []string) {
	o.OtherNames = v
}

// GetNotesChanged returns the NotesChanged field value
func (o *ArtistVersion) GetNotesChanged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NotesChanged
}

// GetNotesChangedOk returns a tuple with the NotesChanged field value
// and a boolean to check if the value has been set.
func (o *ArtistVersion) GetNotesChangedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotesChanged, true
}

// SetNotesChanged sets field value
func (o *ArtistVersion) SetNotesChanged(v bool) {
	o.NotesChanged = v
}

// GetUrls returns the Urls field value
func (o *ArtistVersion) GetUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value
// and a boolean to check if the value has been set.
func (o *ArtistVersion) GetUrlsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Urls, true
}

// SetUrls sets field value
func (o *ArtistVersion) SetUrls(v []string) {
	o.Urls = v
}

func (o ArtistVersion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArtistVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["artist_id"] = o.ArtistId
	toSerialize["name"] = o.Name
	toSerialize["updater_id"] = o.UpdaterId
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["is_active"] = o.IsActive
	toSerialize["other_names"] = o.OtherNames
	toSerialize["notes_changed"] = o.NotesChanged
	toSerialize["urls"] = o.Urls
	return toSerialize, nil
}

func (o *ArtistVersion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"artist_id",
		"name",
		"updater_id",
		"created_at",
		"updated_at",
		"is_active",
		"other_names",
		"notes_changed",
		"urls",
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

	varArtistVersion := _ArtistVersion{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArtistVersion)

	if err != nil {
		return err
	}

	*o = ArtistVersion(varArtistVersion)

	return err
}

type NullableArtistVersion struct {
	value *ArtistVersion
	isSet bool
}

func (v NullableArtistVersion) Get() *ArtistVersion {
	return v.value
}

func (v *NullableArtistVersion) Set(val *ArtistVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableArtistVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableArtistVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtistVersion(val *ArtistVersion) *NullableArtistVersion {
	return &NullableArtistVersion{value: val, isSet: true}
}

func (v NullableArtistVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtistVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
