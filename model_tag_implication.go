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

// checks if the TagImplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TagImplication{}

// TagImplication struct for TagImplication
type TagImplication struct {
	Id             int32         `json:"id"`
	Reason         string        `json:"reason"`
	CreatorId      int32         `json:"creator_id"`
	CreatedAt      time.Time     `json:"created_at"`
	ForumPostId    NullableInt32 `json:"forum_post_id"`
	AntecedentName string        `json:"antecedent_name"`
	ConsequentName string        `json:"consequent_name"`
	// Note: The \"error\" status will be proceeded by an error, ex: \"error: Validation failed: A tag alias for tag_name already exists\"
	Status          TagRequestStatuses `json:"status"`
	ForumTopicId    NullableInt32      `json:"forum_topic_id"`
	UpdatedAt       time.Time          `json:"updated_at"`
	DescendantNames []string           `json:"descendant_names"`
	ApproverId      NullableInt32      `json:"approver_id"`
}

type _TagImplication TagImplication

// NewTagImplication instantiates a new TagImplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagImplication(id int32, reason string, creatorId int32, createdAt time.Time, forumPostId NullableInt32, antecedentName string, consequentName string, status TagRequestStatuses, forumTopicId NullableInt32, updatedAt time.Time, descendantNames []string, approverId NullableInt32) *TagImplication {
	this := TagImplication{}
	this.Id = id
	this.Reason = reason
	this.CreatorId = creatorId
	this.CreatedAt = createdAt
	this.ForumPostId = forumPostId
	this.AntecedentName = antecedentName
	this.ConsequentName = consequentName
	this.Status = status
	this.ForumTopicId = forumTopicId
	this.UpdatedAt = updatedAt
	this.DescendantNames = descendantNames
	this.ApproverId = approverId
	return &this
}

// NewTagImplicationWithDefaults instantiates a new TagImplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagImplicationWithDefaults() *TagImplication {
	this := TagImplication{}
	return &this
}

// GetId returns the Id field value
func (o *TagImplication) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TagImplication) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TagImplication) SetId(v int32) {
	o.Id = v
}

// GetReason returns the Reason field value
func (o *TagImplication) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *TagImplication) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *TagImplication) SetReason(v string) {
	o.Reason = v
}

// GetCreatorId returns the CreatorId field value
func (o *TagImplication) GetCreatorId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatorId
}

// GetCreatorIdOk returns a tuple with the CreatorId field value
// and a boolean to check if the value has been set.
func (o *TagImplication) GetCreatorIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatorId, true
}

// SetCreatorId sets field value
func (o *TagImplication) SetCreatorId(v int32) {
	o.CreatorId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *TagImplication) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *TagImplication) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *TagImplication) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetForumPostId returns the ForumPostId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *TagImplication) GetForumPostId() int32 {
	if o == nil || o.ForumPostId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ForumPostId.Get()
}

// GetForumPostIdOk returns a tuple with the ForumPostId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagImplication) GetForumPostIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForumPostId.Get(), o.ForumPostId.IsSet()
}

// SetForumPostId sets field value
func (o *TagImplication) SetForumPostId(v int32) {
	o.ForumPostId.Set(&v)
}

// GetAntecedentName returns the AntecedentName field value
func (o *TagImplication) GetAntecedentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AntecedentName
}

// GetAntecedentNameOk returns a tuple with the AntecedentName field value
// and a boolean to check if the value has been set.
func (o *TagImplication) GetAntecedentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AntecedentName, true
}

// SetAntecedentName sets field value
func (o *TagImplication) SetAntecedentName(v string) {
	o.AntecedentName = v
}

// GetConsequentName returns the ConsequentName field value
func (o *TagImplication) GetConsequentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsequentName
}

// GetConsequentNameOk returns a tuple with the ConsequentName field value
// and a boolean to check if the value has been set.
func (o *TagImplication) GetConsequentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsequentName, true
}

// SetConsequentName sets field value
func (o *TagImplication) SetConsequentName(v string) {
	o.ConsequentName = v
}

// GetStatus returns the Status field value
func (o *TagImplication) GetStatus() TagRequestStatuses {
	if o == nil {
		var ret TagRequestStatuses
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TagImplication) GetStatusOk() (*TagRequestStatuses, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TagImplication) SetStatus(v TagRequestStatuses) {
	o.Status = v
}

// GetForumTopicId returns the ForumTopicId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *TagImplication) GetForumTopicId() int32 {
	if o == nil || o.ForumTopicId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ForumTopicId.Get()
}

// GetForumTopicIdOk returns a tuple with the ForumTopicId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagImplication) GetForumTopicIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForumTopicId.Get(), o.ForumTopicId.IsSet()
}

// SetForumTopicId sets field value
func (o *TagImplication) SetForumTopicId(v int32) {
	o.ForumTopicId.Set(&v)
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *TagImplication) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *TagImplication) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *TagImplication) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetDescendantNames returns the DescendantNames field value
func (o *TagImplication) GetDescendantNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DescendantNames
}

// GetDescendantNamesOk returns a tuple with the DescendantNames field value
// and a boolean to check if the value has been set.
func (o *TagImplication) GetDescendantNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DescendantNames, true
}

// SetDescendantNames sets field value
func (o *TagImplication) SetDescendantNames(v []string) {
	o.DescendantNames = v
}

// GetApproverId returns the ApproverId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *TagImplication) GetApproverId() int32 {
	if o == nil || o.ApproverId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ApproverId.Get()
}

// GetApproverIdOk returns a tuple with the ApproverId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TagImplication) GetApproverIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApproverId.Get(), o.ApproverId.IsSet()
}

// SetApproverId sets field value
func (o *TagImplication) SetApproverId(v int32) {
	o.ApproverId.Set(&v)
}

func (o TagImplication) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TagImplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["reason"] = o.Reason
	toSerialize["creator_id"] = o.CreatorId
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["forum_post_id"] = o.ForumPostId.Get()
	toSerialize["antecedent_name"] = o.AntecedentName
	toSerialize["consequent_name"] = o.ConsequentName
	toSerialize["status"] = o.Status
	toSerialize["forum_topic_id"] = o.ForumTopicId.Get()
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["descendant_names"] = o.DescendantNames
	toSerialize["approver_id"] = o.ApproverId.Get()
	return toSerialize, nil
}

func (o *TagImplication) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"reason",
		"creator_id",
		"created_at",
		"forum_post_id",
		"antecedent_name",
		"consequent_name",
		"status",
		"forum_topic_id",
		"updated_at",
		"descendant_names",
		"approver_id",
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

	varTagImplication := _TagImplication{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTagImplication)

	if err != nil {
		return err
	}

	*o = TagImplication(varTagImplication)

	return err
}

type NullableTagImplication struct {
	value *TagImplication
	isSet bool
}

func (v NullableTagImplication) Get() *TagImplication {
	return v.value
}

func (v *NullableTagImplication) Set(val *TagImplication) {
	v.value = val
	v.isSet = true
}

func (v NullableTagImplication) IsSet() bool {
	return v.isSet
}

func (v *NullableTagImplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagImplication(val *TagImplication) *NullableTagImplication {
	return &NullableTagImplication{value: val, isSet: true}
}

func (v NullableTagImplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagImplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
