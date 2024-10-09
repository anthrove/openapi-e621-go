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

// checks if the UploadWhitelist type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UploadWhitelist{}

// UploadWhitelist struct for UploadWhitelist
type UploadWhitelist struct {
	Id        float32   `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Pattern   string    `json:"pattern"`
	Note      string    `json:"note"`
	Hidden    bool      `json:"hidden"`
	Allowed   bool      `json:"allowed"`
	Reason    string    `json:"reason"`
}

type _UploadWhitelist UploadWhitelist

// NewUploadWhitelist instantiates a new UploadWhitelist object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUploadWhitelist(id float32, createdAt time.Time, updatedAt time.Time, pattern string, note string, hidden bool, allowed bool, reason string) *UploadWhitelist {
	this := UploadWhitelist{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Pattern = pattern
	this.Note = note
	this.Hidden = hidden
	this.Allowed = allowed
	this.Reason = reason
	return &this
}

// NewUploadWhitelistWithDefaults instantiates a new UploadWhitelist object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadWhitelistWithDefaults() *UploadWhitelist {
	this := UploadWhitelist{}
	return &this
}

// GetId returns the Id field value
func (o *UploadWhitelist) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UploadWhitelist) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UploadWhitelist) SetId(v float32) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *UploadWhitelist) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *UploadWhitelist) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *UploadWhitelist) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *UploadWhitelist) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *UploadWhitelist) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *UploadWhitelist) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetPattern returns the Pattern field value
func (o *UploadWhitelist) GetPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
func (o *UploadWhitelist) GetPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pattern, true
}

// SetPattern sets field value
func (o *UploadWhitelist) SetPattern(v string) {
	o.Pattern = v
}

// GetNote returns the Note field value
func (o *UploadWhitelist) GetNote() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Note
}

// GetNoteOk returns a tuple with the Note field value
// and a boolean to check if the value has been set.
func (o *UploadWhitelist) GetNoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Note, true
}

// SetNote sets field value
func (o *UploadWhitelist) SetNote(v string) {
	o.Note = v
}

// GetHidden returns the Hidden field value
func (o *UploadWhitelist) GetHidden() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value
// and a boolean to check if the value has been set.
func (o *UploadWhitelist) GetHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hidden, true
}

// SetHidden sets field value
func (o *UploadWhitelist) SetHidden(v bool) {
	o.Hidden = v
}

// GetAllowed returns the Allowed field value
func (o *UploadWhitelist) GetAllowed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value
// and a boolean to check if the value has been set.
func (o *UploadWhitelist) GetAllowedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Allowed, true
}

// SetAllowed sets field value
func (o *UploadWhitelist) SetAllowed(v bool) {
	o.Allowed = v
}

// GetReason returns the Reason field value
func (o *UploadWhitelist) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *UploadWhitelist) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *UploadWhitelist) SetReason(v string) {
	o.Reason = v
}

func (o UploadWhitelist) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UploadWhitelist) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["pattern"] = o.Pattern
	toSerialize["note"] = o.Note
	toSerialize["hidden"] = o.Hidden
	toSerialize["allowed"] = o.Allowed
	toSerialize["reason"] = o.Reason
	return toSerialize, nil
}

func (o *UploadWhitelist) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"updated_at",
		"pattern",
		"note",
		"hidden",
		"allowed",
		"reason",
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

	varUploadWhitelist := _UploadWhitelist{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUploadWhitelist)

	if err != nil {
		return err
	}

	*o = UploadWhitelist(varUploadWhitelist)

	return err
}

type NullableUploadWhitelist struct {
	value *UploadWhitelist
	isSet bool
}

func (v NullableUploadWhitelist) Get() *UploadWhitelist {
	return v.value
}

func (v *NullableUploadWhitelist) Set(val *UploadWhitelist) {
	v.value = val
	v.isSet = true
}

func (v NullableUploadWhitelist) IsSet() bool {
	return v.isSet
}

func (v *NullableUploadWhitelist) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUploadWhitelist(val *UploadWhitelist) *NullableUploadWhitelist {
	return &NullableUploadWhitelist{value: val, isSet: true}
}

func (v NullableUploadWhitelist) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUploadWhitelist) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
