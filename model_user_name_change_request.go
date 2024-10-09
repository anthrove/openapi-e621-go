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

// checks if the UserNameChangeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserNameChangeRequest{}

// UserNameChangeRequest struct for UserNameChangeRequest
type UserNameChangeRequest struct {
	Id           float32   `json:"id"`
	ApproverId   float32   `json:"approver_id"`
	UserId       float32   `json:"user_id"`
	OriginalName string    `json:"original_name"`
	DesiredName  string    `json:"desired_name"`
	ChangeReason *string   `json:"change_reason,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Status       string    `json:"status"`
}

type _UserNameChangeRequest UserNameChangeRequest

// NewUserNameChangeRequest instantiates a new UserNameChangeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserNameChangeRequest(id float32, approverId float32, userId float32, originalName string, desiredName string, createdAt time.Time, updatedAt time.Time, status string) *UserNameChangeRequest {
	this := UserNameChangeRequest{}
	this.Id = id
	this.ApproverId = approverId
	this.UserId = userId
	this.OriginalName = originalName
	this.DesiredName = desiredName
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.Status = status
	return &this
}

// NewUserNameChangeRequestWithDefaults instantiates a new UserNameChangeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserNameChangeRequestWithDefaults() *UserNameChangeRequest {
	this := UserNameChangeRequest{}
	return &this
}

// GetId returns the Id field value
func (o *UserNameChangeRequest) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserNameChangeRequest) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserNameChangeRequest) SetId(v float32) {
	o.Id = v
}

// GetApproverId returns the ApproverId field value
func (o *UserNameChangeRequest) GetApproverId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.ApproverId
}

// GetApproverIdOk returns a tuple with the ApproverId field value
// and a boolean to check if the value has been set.
func (o *UserNameChangeRequest) GetApproverIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApproverId, true
}

// SetApproverId sets field value
func (o *UserNameChangeRequest) SetApproverId(v float32) {
	o.ApproverId = v
}

// GetUserId returns the UserId field value
func (o *UserNameChangeRequest) GetUserId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UserNameChangeRequest) GetUserIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UserNameChangeRequest) SetUserId(v float32) {
	o.UserId = v
}

// GetOriginalName returns the OriginalName field value
func (o *UserNameChangeRequest) GetOriginalName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalName
}

// GetOriginalNameOk returns a tuple with the OriginalName field value
// and a boolean to check if the value has been set.
func (o *UserNameChangeRequest) GetOriginalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalName, true
}

// SetOriginalName sets field value
func (o *UserNameChangeRequest) SetOriginalName(v string) {
	o.OriginalName = v
}

// GetDesiredName returns the DesiredName field value
func (o *UserNameChangeRequest) GetDesiredName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DesiredName
}

// GetDesiredNameOk returns a tuple with the DesiredName field value
// and a boolean to check if the value has been set.
func (o *UserNameChangeRequest) GetDesiredNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DesiredName, true
}

// SetDesiredName sets field value
func (o *UserNameChangeRequest) SetDesiredName(v string) {
	o.DesiredName = v
}

// GetChangeReason returns the ChangeReason field value if set, zero value otherwise.
func (o *UserNameChangeRequest) GetChangeReason() string {
	if o == nil || IsNil(o.ChangeReason) {
		var ret string
		return ret
	}
	return *o.ChangeReason
}

// GetChangeReasonOk returns a tuple with the ChangeReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserNameChangeRequest) GetChangeReasonOk() (*string, bool) {
	if o == nil || IsNil(o.ChangeReason) {
		return nil, false
	}
	return o.ChangeReason, true
}

// HasChangeReason returns a boolean if a field has been set.
func (o *UserNameChangeRequest) HasChangeReason() bool {
	if o != nil && !IsNil(o.ChangeReason) {
		return true
	}

	return false
}

// SetChangeReason gets a reference to the given string and assigns it to the ChangeReason field.
func (o *UserNameChangeRequest) SetChangeReason(v string) {
	o.ChangeReason = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *UserNameChangeRequest) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *UserNameChangeRequest) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *UserNameChangeRequest) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *UserNameChangeRequest) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *UserNameChangeRequest) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *UserNameChangeRequest) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetStatus returns the Status field value
func (o *UserNameChangeRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *UserNameChangeRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *UserNameChangeRequest) SetStatus(v string) {
	o.Status = v
}

func (o UserNameChangeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserNameChangeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["approver_id"] = o.ApproverId
	toSerialize["user_id"] = o.UserId
	toSerialize["original_name"] = o.OriginalName
	toSerialize["desired_name"] = o.DesiredName
	if !IsNil(o.ChangeReason) {
		toSerialize["change_reason"] = o.ChangeReason
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *UserNameChangeRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"approver_id",
		"user_id",
		"original_name",
		"desired_name",
		"created_at",
		"updated_at",
		"status",
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

	varUserNameChangeRequest := _UserNameChangeRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserNameChangeRequest)

	if err != nil {
		return err
	}

	*o = UserNameChangeRequest(varUserNameChangeRequest)

	return err
}

type NullableUserNameChangeRequest struct {
	value *UserNameChangeRequest
	isSet bool
}

func (v NullableUserNameChangeRequest) Get() *UserNameChangeRequest {
	return v.value
}

func (v *NullableUserNameChangeRequest) Set(val *UserNameChangeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserNameChangeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserNameChangeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserNameChangeRequest(val *UserNameChangeRequest) *NullableUserNameChangeRequest {
	return &NullableUserNameChangeRequest{value: val, isSet: true}
}

func (v NullableUserNameChangeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserNameChangeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
