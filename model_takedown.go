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

// checks if the Takedown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Takedown{}

// Takedown struct for Takedown
type Takedown struct {
	Id           int32           `json:"id"`
	Status       string          `json:"status"`
	ApproverId   NullableFloat32 `json:"approver_id"`
	ReasonHidden bool            `json:"reason_hidden"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
	PostCount    int32           `json:"post_count"`
}

type _Takedown Takedown

// NewTakedown instantiates a new Takedown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTakedown(id int32, status string, approverId NullableFloat32, reasonHidden bool, createdAt time.Time, updatedAt time.Time, postCount int32) *Takedown {
	this := Takedown{}
	this.Id = id
	this.Status = status
	this.ApproverId = approverId
	this.ReasonHidden = reasonHidden
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.PostCount = postCount
	return &this
}

// NewTakedownWithDefaults instantiates a new Takedown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTakedownWithDefaults() *Takedown {
	this := Takedown{}
	return &this
}

// GetId returns the Id field value
func (o *Takedown) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Takedown) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Takedown) SetId(v int32) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *Takedown) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Takedown) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Takedown) SetStatus(v string) {
	o.Status = v
}

// GetApproverId returns the ApproverId field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *Takedown) GetApproverId() float32 {
	if o == nil || o.ApproverId.Get() == nil {
		var ret float32
		return ret
	}

	return *o.ApproverId.Get()
}

// GetApproverIdOk returns a tuple with the ApproverId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Takedown) GetApproverIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApproverId.Get(), o.ApproverId.IsSet()
}

// SetApproverId sets field value
func (o *Takedown) SetApproverId(v float32) {
	o.ApproverId.Set(&v)
}

// GetReasonHidden returns the ReasonHidden field value
func (o *Takedown) GetReasonHidden() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ReasonHidden
}

// GetReasonHiddenOk returns a tuple with the ReasonHidden field value
// and a boolean to check if the value has been set.
func (o *Takedown) GetReasonHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReasonHidden, true
}

// SetReasonHidden sets field value
func (o *Takedown) SetReasonHidden(v bool) {
	o.ReasonHidden = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Takedown) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Takedown) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Takedown) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Takedown) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Takedown) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Takedown) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetPostCount returns the PostCount field value
func (o *Takedown) GetPostCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PostCount
}

// GetPostCountOk returns a tuple with the PostCount field value
// and a boolean to check if the value has been set.
func (o *Takedown) GetPostCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostCount, true
}

// SetPostCount sets field value
func (o *Takedown) SetPostCount(v int32) {
	o.PostCount = v
}

func (o Takedown) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Takedown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["approver_id"] = o.ApproverId.Get()
	toSerialize["reason_hidden"] = o.ReasonHidden
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["post_count"] = o.PostCount
	return toSerialize, nil
}

func (o *Takedown) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"status",
		"approver_id",
		"reason_hidden",
		"created_at",
		"updated_at",
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

	varTakedown := _Takedown{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTakedown)

	if err != nil {
		return err
	}

	*o = Takedown(varTakedown)

	return err
}

type NullableTakedown struct {
	value *Takedown
	isSet bool
}

func (v NullableTakedown) Get() *Takedown {
	return v.value
}

func (v *NullableTakedown) Set(val *Takedown) {
	v.value = val
	v.isSet = true
}

func (v NullableTakedown) IsSet() bool {
	return v.isSet
}

func (v *NullableTakedown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTakedown(val *Takedown) *NullableTakedown {
	return &NullableTakedown{value: val, isSet: true}
}

func (v NullableTakedown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTakedown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
