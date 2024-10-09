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

// checks if the PostVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostVersion{}

// PostVersion struct for PostVersion
type PostVersion struct {
	Id                  float32         `json:"id"`
	PostId              float32         `json:"post_id"`
	Tags                string          `json:"tags"`
	UpdaterId           float32         `json:"updater_id"`
	UpdatedAt           time.Time       `json:"updated_at"`
	Rating              Ratings         `json:"rating"`
	ParentId            NullableFloat32 `json:"parent_id"`
	Source              string          `json:"source"`
	Description         string          `json:"description"`
	Reason              NullableString  `json:"reason"`
	LockedTags          NullableString  `json:"locked_tags"`
	AddedTags           []string        `json:"added_tags"`
	RemovedTags         []string        `json:"removed_tags"`
	AddedLockedTags     []string        `json:"added_locked_tags"`
	RemovedLockedTags   []string        `json:"removed_locked_tags"`
	RatingChanged       bool            `json:"rating_changed"`
	ParentChanged       bool            `json:"parent_changed"`
	SourceChanged       bool            `json:"source_changed"`
	DescriptionChanged  bool            `json:"description_changed"`
	Version             float32         `json:"version"`
	ObsoleteAddedTags   string          `json:"obsolete_added_tags"`
	ObsoleteRemovedTags string          `json:"obsolete_removed_tags"`
	UnchangedTags       string          `json:"unchanged_tags"`
	UpdaterName         string          `json:"updater_name"`
}

type _PostVersion PostVersion

// NewPostVersion instantiates a new PostVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostVersion(id float32, postId float32, tags string, updaterId float32, updatedAt time.Time, rating Ratings, parentId NullableFloat32, source string, description string, reason NullableString, lockedTags NullableString, addedTags []string, removedTags []string, addedLockedTags []string, removedLockedTags []string, ratingChanged bool, parentChanged bool, sourceChanged bool, descriptionChanged bool, version float32, obsoleteAddedTags string, obsoleteRemovedTags string, unchangedTags string, updaterName string) *PostVersion {
	this := PostVersion{}
	this.Id = id
	this.PostId = postId
	this.Tags = tags
	this.UpdaterId = updaterId
	this.UpdatedAt = updatedAt
	this.Rating = rating
	this.ParentId = parentId
	this.Source = source
	this.Description = description
	this.Reason = reason
	this.LockedTags = lockedTags
	this.AddedTags = addedTags
	this.RemovedTags = removedTags
	this.AddedLockedTags = addedLockedTags
	this.RemovedLockedTags = removedLockedTags
	this.RatingChanged = ratingChanged
	this.ParentChanged = parentChanged
	this.SourceChanged = sourceChanged
	this.DescriptionChanged = descriptionChanged
	this.Version = version
	this.ObsoleteAddedTags = obsoleteAddedTags
	this.ObsoleteRemovedTags = obsoleteRemovedTags
	this.UnchangedTags = unchangedTags
	this.UpdaterName = updaterName
	return &this
}

// NewPostVersionWithDefaults instantiates a new PostVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostVersionWithDefaults() *PostVersion {
	this := PostVersion{}
	return &this
}

// GetId returns the Id field value
func (o *PostVersion) GetId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PostVersion) GetIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PostVersion) SetId(v float32) {
	o.Id = v
}

// GetPostId returns the PostId field value
func (o *PostVersion) GetPostId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PostId
}

// GetPostIdOk returns a tuple with the PostId field value
// and a boolean to check if the value has been set.
func (o *PostVersion) GetPostIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostId, true
}

// SetPostId sets field value
func (o *PostVersion) SetPostId(v float32) {
	o.PostId = v
}

// GetTags returns the Tags field value
func (o *PostVersion) GetTags() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *PostVersion) GetTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value
func (o *PostVersion) SetTags(v string) {
	o.Tags = v
}

// GetUpdaterId returns the UpdaterId field value
func (o *PostVersion) GetUpdaterId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UpdaterId
}

// GetUpdaterIdOk returns a tuple with the UpdaterId field value
// and a boolean to check if the value has been set.
func (o *PostVersion) GetUpdaterIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdaterId, true
}

// SetUpdaterId sets field value
func (o *PostVersion) SetUpdaterId(v float32) {
	o.UpdaterId = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *PostVersion) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PostVersion) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *PostVersion) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetRating returns the Rating field value
func (o *PostVersion) GetRating() Ratings {
	if o == nil {
		var ret Ratings
		return ret
	}

	return o.Rating
}

// GetRatingOk returns a tuple with the Rating field value
// and a boolean to check if the value has been set.
func (o *PostVersion) GetRatingOk() (*Ratings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rating, true
}

// SetRating sets field value
func (o *PostVersion) SetRating(v Ratings) {
	o.Rating = v
}

// GetParentId returns the ParentId field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *PostVersion) GetParentId() float32 {
	if o == nil || o.ParentId.Get() == nil {
		var ret float32
		return ret
	}

	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostVersion) GetParentIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// SetParentId sets field value
func (o *PostVersion) SetParentId(v float32) {
	o.ParentId.Set(&v)
}

// GetSource returns the Source field value
func (o *PostVersion) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *PostVersion) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *PostVersion) SetSource(v string) {
	o.Source = v
}

// GetDescription returns the Description field value
func (o *PostVersion) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *PostVersion) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *PostVersion) SetDescription(v string) {
	o.Description = v
}

// GetReason returns the Reason field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PostVersion) GetReason() string {
	if o == nil || o.Reason.Get() == nil {
		var ret string
		return ret
	}

	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostVersion) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// SetReason sets field value
func (o *PostVersion) SetReason(v string) {
	o.Reason.Set(&v)
}

// GetLockedTags returns the LockedTags field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PostVersion) GetLockedTags() string {
	if o == nil || o.LockedTags.Get() == nil {
		var ret string
		return ret
	}

	return *o.LockedTags.Get()
}

// GetLockedTagsOk returns a tuple with the LockedTags field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostVersion) GetLockedTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LockedTags.Get(), o.LockedTags.IsSet()
}

// SetLockedTags sets field value
func (o *PostVersion) SetLockedTags(v string) {
	o.LockedTags.Set(&v)
}

// GetAddedTags returns the AddedTags field value
func (o *PostVersion) GetAddedTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AddedTags
}

// GetAddedTagsOk returns a tuple with the AddedTags field value
// and a boolean to check if the value has been set.
func (o *PostVersion) GetAddedTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddedTags, true
}

// SetAddedTags sets field value
func (o *PostVersion) SetAddedTags(v []string) {
	o.AddedTags = v
}

// GetRemovedTags returns the RemovedTags field value
func (o *PostVersion) GetRemovedTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RemovedTags
}

// GetRemovedTagsOk returns a tuple with the RemovedTags field value
// and a boolean to check if the value has been set.
func (o *PostVersion) GetRemovedTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemovedTags, true
}

// SetRemovedTags sets field value
func (o *PostVersion) SetRemovedTags(v []string) {
	o.RemovedTags = v
}

// GetAddedLockedTags returns the AddedLockedTags field value
func (o *PostVersion) GetAddedLockedTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AddedLockedTags
}

// GetAddedLockedTagsOk returns a tuple with the AddedLockedTags field value
// and a boolean to check if the value has been set.
func (o *PostVersion) GetAddedLockedTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AddedLockedTags, true
}

// SetAddedLockedTags sets field value
func (o *PostVersion) SetAddedLockedTags(v []string) {
	o.AddedLockedTags = v
}

// GetRemovedLockedTags returns the RemovedLockedTags field value
func (o *PostVersion) GetRemovedLockedTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RemovedLockedTags
}

// GetRemovedLockedTagsOk returns a tuple with the RemovedLockedTags field value
// and a boolean to check if the value has been set.
func (o *PostVersion) GetRemovedLockedTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RemovedLockedTags, true
}

// SetRemovedLockedTags sets field value
func (o *PostVersion) SetRemovedLockedTags(v []string) {
	o.RemovedLockedTags = v
}

// GetRatingChanged returns the RatingChanged field value
func (o *PostVersion) GetRatingChanged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RatingChanged
}

// GetRatingChangedOk returns a tuple with the RatingChanged field value
// and a boolean to check if the value has been set.
func (o *PostVersion) GetRatingChangedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RatingChanged, true
}

// SetRatingChanged sets field value
func (o *PostVersion) SetRatingChanged(v bool) {
	o.RatingChanged = v
}

// GetParentChanged returns the ParentChanged field value
func (o *PostVersion) GetParentChanged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ParentChanged
}

// GetParentChangedOk returns a tuple with the ParentChanged field value
// and a boolean to check if the value has been set.
func (o *PostVersion) GetParentChangedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentChanged, true
}

// SetParentChanged sets field value
func (o *PostVersion) SetParentChanged(v bool) {
	o.ParentChanged = v
}

// GetSourceChanged returns the SourceChanged field value
func (o *PostVersion) GetSourceChanged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SourceChanged
}

// GetSourceChangedOk returns a tuple with the SourceChanged field value
// and a boolean to check if the value has been set.
func (o *PostVersion) GetSourceChangedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceChanged, true
}

// SetSourceChanged sets field value
func (o *PostVersion) SetSourceChanged(v bool) {
	o.SourceChanged = v
}

// GetDescriptionChanged returns the DescriptionChanged field value
func (o *PostVersion) GetDescriptionChanged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DescriptionChanged
}

// GetDescriptionChangedOk returns a tuple with the DescriptionChanged field value
// and a boolean to check if the value has been set.
func (o *PostVersion) GetDescriptionChangedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DescriptionChanged, true
}

// SetDescriptionChanged sets field value
func (o *PostVersion) SetDescriptionChanged(v bool) {
	o.DescriptionChanged = v
}

// GetVersion returns the Version field value
func (o *PostVersion) GetVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *PostVersion) GetVersionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *PostVersion) SetVersion(v float32) {
	o.Version = v
}

// GetObsoleteAddedTags returns the ObsoleteAddedTags field value
func (o *PostVersion) GetObsoleteAddedTags() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObsoleteAddedTags
}

// GetObsoleteAddedTagsOk returns a tuple with the ObsoleteAddedTags field value
// and a boolean to check if the value has been set.
func (o *PostVersion) GetObsoleteAddedTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObsoleteAddedTags, true
}

// SetObsoleteAddedTags sets field value
func (o *PostVersion) SetObsoleteAddedTags(v string) {
	o.ObsoleteAddedTags = v
}

// GetObsoleteRemovedTags returns the ObsoleteRemovedTags field value
func (o *PostVersion) GetObsoleteRemovedTags() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObsoleteRemovedTags
}

// GetObsoleteRemovedTagsOk returns a tuple with the ObsoleteRemovedTags field value
// and a boolean to check if the value has been set.
func (o *PostVersion) GetObsoleteRemovedTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObsoleteRemovedTags, true
}

// SetObsoleteRemovedTags sets field value
func (o *PostVersion) SetObsoleteRemovedTags(v string) {
	o.ObsoleteRemovedTags = v
}

// GetUnchangedTags returns the UnchangedTags field value
func (o *PostVersion) GetUnchangedTags() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UnchangedTags
}

// GetUnchangedTagsOk returns a tuple with the UnchangedTags field value
// and a boolean to check if the value has been set.
func (o *PostVersion) GetUnchangedTagsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UnchangedTags, true
}

// SetUnchangedTags sets field value
func (o *PostVersion) SetUnchangedTags(v string) {
	o.UnchangedTags = v
}

// GetUpdaterName returns the UpdaterName field value
func (o *PostVersion) GetUpdaterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdaterName
}

// GetUpdaterNameOk returns a tuple with the UpdaterName field value
// and a boolean to check if the value has been set.
func (o *PostVersion) GetUpdaterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdaterName, true
}

// SetUpdaterName sets field value
func (o *PostVersion) SetUpdaterName(v string) {
	o.UpdaterName = v
}

func (o PostVersion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["post_id"] = o.PostId
	toSerialize["tags"] = o.Tags
	toSerialize["updater_id"] = o.UpdaterId
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["rating"] = o.Rating
	toSerialize["parent_id"] = o.ParentId.Get()
	toSerialize["source"] = o.Source
	toSerialize["description"] = o.Description
	toSerialize["reason"] = o.Reason.Get()
	toSerialize["locked_tags"] = o.LockedTags.Get()
	toSerialize["added_tags"] = o.AddedTags
	toSerialize["removed_tags"] = o.RemovedTags
	toSerialize["added_locked_tags"] = o.AddedLockedTags
	toSerialize["removed_locked_tags"] = o.RemovedLockedTags
	toSerialize["rating_changed"] = o.RatingChanged
	toSerialize["parent_changed"] = o.ParentChanged
	toSerialize["source_changed"] = o.SourceChanged
	toSerialize["description_changed"] = o.DescriptionChanged
	toSerialize["version"] = o.Version
	toSerialize["obsolete_added_tags"] = o.ObsoleteAddedTags
	toSerialize["obsolete_removed_tags"] = o.ObsoleteRemovedTags
	toSerialize["unchanged_tags"] = o.UnchangedTags
	toSerialize["updater_name"] = o.UpdaterName
	return toSerialize, nil
}

func (o *PostVersion) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"post_id",
		"tags",
		"updater_id",
		"updated_at",
		"rating",
		"parent_id",
		"source",
		"description",
		"reason",
		"locked_tags",
		"added_tags",
		"removed_tags",
		"added_locked_tags",
		"removed_locked_tags",
		"rating_changed",
		"parent_changed",
		"source_changed",
		"description_changed",
		"version",
		"obsolete_added_tags",
		"obsolete_removed_tags",
		"unchanged_tags",
		"updater_name",
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

	varPostVersion := _PostVersion{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostVersion)

	if err != nil {
		return err
	}

	*o = PostVersion(varPostVersion)

	return err
}

type NullablePostVersion struct {
	value *PostVersion
	isSet bool
}

func (v NullablePostVersion) Get() *PostVersion {
	return v.value
}

func (v *NullablePostVersion) Set(val *PostVersion) {
	v.value = val
	v.isSet = true
}

func (v NullablePostVersion) IsSet() bool {
	return v.isSet
}

func (v *NullablePostVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostVersion(val *PostVersion) *NullablePostVersion {
	return &NullablePostVersion{value: val, isSet: true}
}

func (v NullablePostVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
