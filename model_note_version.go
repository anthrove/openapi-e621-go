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

// checks if the NoteVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NoteVersion{}

// NoteVersion struct for NoteVersion
type NoteVersion struct {
	Id        int32     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	X         int32     `json:"x"`
	Y         int32     `json:"y"`
	Width     int32     `json:"width"`
	Height    int32     `json:"height"`
	Body      string    `json:"body"`
	Version   float32   `json:"version"`
	IsActive  bool      `json:"is_active"`
	NoteId    int32     `json:"note_id"`
	PostId    int32     `json:"post_id"`
	UpdaterId int32     `json:"updater_id"`
}

type _NoteVersion NoteVersion

// NewNoteVersion instantiates a new NoteVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNoteVersion(id int32, createdAt time.Time, updatedAt time.Time, x int32, y int32, width int32, height int32, body string, version float32, isActive bool, noteId int32, postId int32, updaterId int32) *NoteVersion {
	this := NoteVersion{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.X = x
	this.Y = y
	this.Width = width
	this.Height = height
	this.Body = body
	this.Version = version
	this.IsActive = isActive
	this.NoteId = noteId
	this.PostId = postId
	this.UpdaterId = updaterId
	return &this
}

// NewNoteVersionWithDefaults instantiates a new NoteVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNoteVersionWithDefaults() *NoteVersion {
	this := NoteVersion{}
	return &this
}

// GetId returns the Id field value
func (o *NoteVersion) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NoteVersion) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NoteVersion) SetId(v int32) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *NoteVersion) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *NoteVersion) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *NoteVersion) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *NoteVersion) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *NoteVersion) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *NoteVersion) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetX returns the X field value
func (o *NoteVersion) GetX() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *NoteVersion) GetXOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *NoteVersion) SetX(v int32) {
	o.X = v
}

// GetY returns the Y field value
func (o *NoteVersion) GetY() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *NoteVersion) GetYOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *NoteVersion) SetY(v int32) {
	o.Y = v
}

// GetWidth returns the Width field value
func (o *NoteVersion) GetWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *NoteVersion) GetWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *NoteVersion) SetWidth(v int32) {
	o.Width = v
}

// GetHeight returns the Height field value
func (o *NoteVersion) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *NoteVersion) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *NoteVersion) SetHeight(v int32) {
	o.Height = v
}

// GetBody returns the Body field value
func (o *NoteVersion) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *NoteVersion) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *NoteVersion) SetBody(v string) {
	o.Body = v
}

// GetVersion returns the Version field value
func (o *NoteVersion) GetVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *NoteVersion) GetVersionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *NoteVersion) SetVersion(v float32) {
	o.Version = v
}

// GetIsActive returns the IsActive field value
func (o *NoteVersion) GetIsActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value
// and a boolean to check if the value has been set.
func (o *NoteVersion) GetIsActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsActive, true
}

// SetIsActive sets field value
func (o *NoteVersion) SetIsActive(v bool) {
	o.IsActive = v
}

// GetNoteId returns the NoteId field value
func (o *NoteVersion) GetNoteId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NoteId
}

// GetNoteIdOk returns a tuple with the NoteId field value
// and a boolean to check if the value has been set.
func (o *NoteVersion) GetNoteIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NoteId, true
}

// SetNoteId sets field value
func (o *NoteVersion) SetNoteId(v int32) {
	o.NoteId = v
}

// GetPostId returns the PostId field value
func (o *NoteVersion) GetPostId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PostId
}

// GetPostIdOk returns a tuple with the PostId field value
// and a boolean to check if the value has been set.
func (o *NoteVersion) GetPostIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostId, true
}

// SetPostId sets field value
func (o *NoteVersion) SetPostId(v int32) {
	o.PostId = v
}

// GetUpdaterId returns the UpdaterId field value
func (o *NoteVersion) GetUpdaterId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdaterId
}

// GetUpdaterIdOk returns a tuple with the UpdaterId field value
// and a boolean to check if the value has been set.
func (o *NoteVersion) GetUpdaterIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdaterId, true
}

// SetUpdaterId sets field value
func (o *NoteVersion) SetUpdaterId(v int32) {
	o.UpdaterId = v
}

func (o NoteVersion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NoteVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y
	toSerialize["width"] = o.Width
	toSerialize["height"] = o.Height
	toSerialize["body"] = o.Body
	toSerialize["version"] = o.Version
	toSerialize["is_active"] = o.IsActive
	toSerialize["note_id"] = o.NoteId
	toSerialize["post_id"] = o.PostId
	toSerialize["updater_id"] = o.UpdaterId
	return toSerialize, nil
}

func (o *NoteVersion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"x",
		"y",
		"width",
		"height",
		"body",
		"version",
		"is_active",
		"note_id",
		"post_id",
		"updater_id",
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

	varNoteVersion := _NoteVersion{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varNoteVersion)

	if err != nil {
		return err
	}

	*o = NoteVersion(varNoteVersion)

	return err
}

type NullableNoteVersion struct {
	value *NoteVersion
	isSet bool
}

func (v NullableNoteVersion) Get() *NoteVersion {
	return v.value
}

func (v *NullableNoteVersion) Set(val *NoteVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableNoteVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableNoteVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoteVersion(val *NoteVersion) *NullableNoteVersion {
	return &NullableNoteVersion{value: val, isSet: true}
}

func (v NullableNoteVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoteVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
