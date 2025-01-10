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

// checks if the BulkUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkUpdateRequest{}

// BulkUpdateRequest struct for BulkUpdateRequest
type BulkUpdateRequest struct {
	Id           int32         `json:"id"`
	CreatorId    int32         `json:"creator_id"`
	ForumTopicId NullableInt32 `json:"forum_topic_id"`
	Script       string        `json:"script"`
	Status       string        `json:"status"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	ApproverId   NullableInt32 `json:"approver_id"`
	ForumPostId  NullableInt32 `json:"forum_post_id"`
	Title        string        `json:"title"`
}

type _BulkUpdateRequest BulkUpdateRequest

// NewBulkUpdateRequest instantiates a new BulkUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkUpdateRequest(id int32, creatorId int32, forumTopicId NullableInt32, script string, status string, createdAt time.Time, updatedAt time.Time, approverId NullableInt32, forumPostId NullableInt32, title string) *BulkUpdateRequest {
	this := BulkUpdateRequest{}
	this.Id = id
	this.CreatorId = creatorId
	this.ForumTopicId = forumTopicId
	this.Script = script
	this.Status = status
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.ApproverId = approverId
	this.ForumPostId = forumPostId
	this.Title = title
	return &this
}

// NewBulkUpdateRequestWithDefaults instantiates a new BulkUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkUpdateRequestWithDefaults() *BulkUpdateRequest {
	this := BulkUpdateRequest{}
	return &this
}

// GetId returns the Id field value
func (o *BulkUpdateRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkUpdateRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkUpdateRequest) SetId(v int32) {
	o.Id = v
}

// GetCreatorId returns the CreatorId field value
func (o *BulkUpdateRequest) GetCreatorId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value
// and a boolean to check if the value has been set.
func (o *BulkUpdateRequest) GetCreatorIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorId, true
}

// SetCreatorId sets field value
func (o *BulkUpdateRequest) SetCreatorId(v int32) {
	o.CreatorId = v
}

// GetForumTopicId returns the ForumTopicId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *BulkUpdateRequest) GetForumTopicId() int32 {
	if o == nil || o.ForumTopicId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ForumTopicId.Get()
}

// GetForumTopicIdOk returns a tuple with the ForumTopicId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkUpdateRequest) GetForumTopicIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForumTopicId.Get(), o.ForumTopicId.IsSet()
}

// SetForumTopicId sets field value
func (o *BulkUpdateRequest) SetForumTopicId(v int32) {
	o.ForumTopicId.Set(&v)
}

// GetScript returns the Script field value
func (o *BulkUpdateRequest) GetScript() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Script
}

// GetScriptOk returns a tuple with the Script field value
// and a boolean to check if the value has been set.
func (o *BulkUpdateRequest) GetScriptOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Script, true
}

// SetScript sets field value
func (o *BulkUpdateRequest) SetScript(v string) {
	o.Script = v
}

// GetStatus returns the Status field value
func (o *BulkUpdateRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BulkUpdateRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BulkUpdateRequest) SetStatus(v string) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *BulkUpdateRequest) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *BulkUpdateRequest) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *BulkUpdateRequest) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *BulkUpdateRequest) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *BulkUpdateRequest) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *BulkUpdateRequest) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetApproverId returns the ApproverId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *BulkUpdateRequest) GetApproverId() int32 {
	if o == nil || o.ApproverId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ApproverId.Get()
}

// GetApproverIdOk returns a tuple with the ApproverId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkUpdateRequest) GetApproverIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApproverId.Get(), o.ApproverId.IsSet()
}

// SetApproverId sets field value
func (o *BulkUpdateRequest) SetApproverId(v int32) {
	o.ApproverId.Set(&v)
}

// GetForumPostId returns the ForumPostId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *BulkUpdateRequest) GetForumPostId() int32 {
	if o == nil || o.ForumPostId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ForumPostId.Get()
}

// GetForumPostIdOk returns a tuple with the ForumPostId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkUpdateRequest) GetForumPostIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForumPostId.Get(), o.ForumPostId.IsSet()
}

// SetForumPostId sets field value
func (o *BulkUpdateRequest) SetForumPostId(v int32) {
	o.ForumPostId.Set(&v)
}

// GetTitle returns the Title field value
func (o *BulkUpdateRequest) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *BulkUpdateRequest) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *BulkUpdateRequest) SetTitle(v string) {
	o.Title = v
}

func (o BulkUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["creator_id"] = o.CreatorId
	toSerialize["forum_topic_id"] = o.ForumTopicId.Get()
	toSerialize["script"] = o.Script
	toSerialize["status"] = o.Status
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["approver_id"] = o.ApproverId.Get()
	toSerialize["forum_post_id"] = o.ForumPostId.Get()
	toSerialize["title"] = o.Title
	return toSerialize, nil
}

func (o *BulkUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"creator_id",
		"forum_topic_id",
		"script",
		"status",
		"created_at",
		"updated_at",
		"approver_id",
		"forum_post_id",
		"title",
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

	varBulkUpdateRequest := _BulkUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBulkUpdateRequest)

	if err != nil {
		return err
	}

	*o = BulkUpdateRequest(varBulkUpdateRequest)

	return err
}

type NullableBulkUpdateRequest struct {
	value *BulkUpdateRequest
	isSet bool
}

func (v NullableBulkUpdateRequest) Get() *BulkUpdateRequest {
	return v.value
}

func (v *NullableBulkUpdateRequest) Set(val *BulkUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkUpdateRequest(val *BulkUpdateRequest) *NullableBulkUpdateRequest {
	return &NullableBulkUpdateRequest{value: val, isSet: true}
}

func (v NullableBulkUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
