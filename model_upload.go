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

// checks if the Upload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Upload{}

// Upload struct for Upload
type Upload struct {
	Id         int32   `json:"id"`
	Source     string  `json:"source"`
	Rating     Ratings `json:"rating"`
	UploaderId int32   `json:"uploader_id"`
	TagString  string  `json:"tag_string"`
	// Note: The \"error\" status will be proceeded by an error, ex: \"error: RuntimeError - No file or source URL provided\"
	Status    string          `json:"status"`
	Backtrace NullableString  `json:"backtrace"`
	PostId    NullableFloat32 `json:"post_id"`
	// Deprecated
	Md5Confirmation interface{}                `json:"md5_confirmation"`
	CreatedAt       time.Time                  `json:"created_at"`
	UpdatedAt       time.Time                  `json:"updated_at"`
	ParentId        NullableFloat32            `json:"parent_id"`
	Md5             NullableString             `json:"md5"`
	FileExt         GetArtistIdOrNameParameter `json:"file_ext"`
	FileSize        NullableFloat32            `json:"file_size"`
	ImageWidth      NullableFloat32            `json:"image_width"`
	ImageHeight     NullableFloat32            `json:"image_height"`
	Description     string                     `json:"description"`
	UploaderName    string                     `json:"uploader_name"`
}

type _Upload Upload

// NewUpload instantiates a new Upload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpload(id int32, source string, rating Ratings, uploaderId int32, tagString string, status string, backtrace NullableString, postId NullableFloat32, md5Confirmation interface{}, createdAt time.Time, updatedAt time.Time, parentId NullableFloat32, md5 NullableString, fileExt GetArtistIdOrNameParameter, fileSize NullableFloat32, imageWidth NullableFloat32, imageHeight NullableFloat32, description string, uploaderName string) *Upload {
	this := Upload{}
	this.Id = id
	this.Source = source
	this.Rating = rating
	this.UploaderId = uploaderId
	this.TagString = tagString
	this.Status = status
	this.Backtrace = backtrace
	this.PostId = postId
	this.Md5Confirmation = md5Confirmation
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.ParentId = parentId
	this.Md5 = md5
	this.FileExt = fileExt
	this.FileSize = fileSize
	this.ImageWidth = imageWidth
	this.ImageHeight = imageHeight
	this.Description = description
	this.UploaderName = uploaderName
	return &this
}

// NewUploadWithDefaults instantiates a new Upload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUploadWithDefaults() *Upload {
	this := Upload{}
	return &this
}

// GetId returns the Id field value
func (o *Upload) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Upload) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Upload) SetId(v int32) {
	o.Id = v
}

// GetSource returns the Source field value
func (o *Upload) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *Upload) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *Upload) SetSource(v string) {
	o.Source = v
}

// GetRating returns the Rating field value
func (o *Upload) GetRating() Ratings {
	if o == nil {
		var ret Ratings
		return ret
	}

	return o.Rating
}

// GetRatingOk returns a tuple with the Rating field value
// and a boolean to check if the value has been set.
func (o *Upload) GetRatingOk() (*Ratings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rating, true
}

// SetRating sets field value
func (o *Upload) SetRating(v Ratings) {
	o.Rating = v
}

// GetUploaderId returns the UploaderId field value
func (o *Upload) GetUploaderId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UploaderId
}

// GetUploaderIdOk returns a tuple with the UploaderId field value
// and a boolean to check if the value has been set.
func (o *Upload) GetUploaderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploaderId, true
}

// SetUploaderId sets field value
func (o *Upload) SetUploaderId(v int32) {
	o.UploaderId = v
}

// GetTagString returns the TagString field value
func (o *Upload) GetTagString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TagString
}

// GetTagStringOk returns a tuple with the TagString field value
// and a boolean to check if the value has been set.
func (o *Upload) GetTagStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagString, true
}

// SetTagString sets field value
func (o *Upload) SetTagString(v string) {
	o.TagString = v
}

// GetStatus returns the Status field value
func (o *Upload) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Upload) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Upload) SetStatus(v string) {
	o.Status = v
}

// GetBacktrace returns the Backtrace field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Upload) GetBacktrace() string {
	if o == nil || o.Backtrace.Get() == nil {
		var ret string
		return ret
	}

	return *o.Backtrace.Get()
}

// GetBacktraceOk returns a tuple with the Backtrace field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Upload) GetBacktraceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Backtrace.Get(), o.Backtrace.IsSet()
}

// SetBacktrace sets field value
func (o *Upload) SetBacktrace(v string) {
	o.Backtrace.Set(&v)
}

// GetPostId returns the PostId field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *Upload) GetPostId() float32 {
	if o == nil || o.PostId.Get() == nil {
		var ret float32
		return ret
	}

	return *o.PostId.Get()
}

// GetPostIdOk returns a tuple with the PostId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Upload) GetPostIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PostId.Get(), o.PostId.IsSet()
}

// SetPostId sets field value
func (o *Upload) SetPostId(v float32) {
	o.PostId.Set(&v)
}

// GetMd5Confirmation returns the Md5Confirmation field value
// If the value is explicit nil, the zero value for interface{} will be returned
// Deprecated
func (o *Upload) GetMd5Confirmation() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Md5Confirmation
}

// GetMd5ConfirmationOk returns a tuple with the Md5Confirmation field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *Upload) GetMd5ConfirmationOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Md5Confirmation) {
		return nil, false
	}
	return &o.Md5Confirmation, true
}

// SetMd5Confirmation sets field value
// Deprecated
func (o *Upload) SetMd5Confirmation(v interface{}) {
	o.Md5Confirmation = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Upload) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Upload) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Upload) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Upload) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Upload) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Upload) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetParentId returns the ParentId field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *Upload) GetParentId() float32 {
	if o == nil || o.ParentId.Get() == nil {
		var ret float32
		return ret
	}

	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Upload) GetParentIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// SetParentId sets field value
func (o *Upload) SetParentId(v float32) {
	o.ParentId.Set(&v)
}

// GetMd5 returns the Md5 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Upload) GetMd5() string {
	if o == nil || o.Md5.Get() == nil {
		var ret string
		return ret
	}

	return *o.Md5.Get()
}

// GetMd5Ok returns a tuple with the Md5 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Upload) GetMd5Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Md5.Get(), o.Md5.IsSet()
}

// SetMd5 sets field value
func (o *Upload) SetMd5(v string) {
	o.Md5.Set(&v)
}

// GetFileExt returns the FileExt field value
func (o *Upload) GetFileExt() GetArtistIdOrNameParameter {
	if o == nil {
		var ret GetArtistIdOrNameParameter
		return ret
	}

	return o.FileExt
}

// GetFileExtOk returns a tuple with the FileExt field value
// and a boolean to check if the value has been set.
func (o *Upload) GetFileExtOk() (*GetArtistIdOrNameParameter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileExt, true
}

// SetFileExt sets field value
func (o *Upload) SetFileExt(v GetArtistIdOrNameParameter) {
	o.FileExt = v
}

// GetFileSize returns the FileSize field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *Upload) GetFileSize() float32 {
	if o == nil || o.FileSize.Get() == nil {
		var ret float32
		return ret
	}

	return *o.FileSize.Get()
}

// GetFileSizeOk returns a tuple with the FileSize field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Upload) GetFileSizeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileSize.Get(), o.FileSize.IsSet()
}

// SetFileSize sets field value
func (o *Upload) SetFileSize(v float32) {
	o.FileSize.Set(&v)
}

// GetImageWidth returns the ImageWidth field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *Upload) GetImageWidth() float32 {
	if o == nil || o.ImageWidth.Get() == nil {
		var ret float32
		return ret
	}

	return *o.ImageWidth.Get()
}

// GetImageWidthOk returns a tuple with the ImageWidth field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Upload) GetImageWidthOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageWidth.Get(), o.ImageWidth.IsSet()
}

// SetImageWidth sets field value
func (o *Upload) SetImageWidth(v float32) {
	o.ImageWidth.Set(&v)
}

// GetImageHeight returns the ImageHeight field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *Upload) GetImageHeight() float32 {
	if o == nil || o.ImageHeight.Get() == nil {
		var ret float32
		return ret
	}

	return *o.ImageHeight.Get()
}

// GetImageHeightOk returns a tuple with the ImageHeight field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Upload) GetImageHeightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageHeight.Get(), o.ImageHeight.IsSet()
}

// SetImageHeight sets field value
func (o *Upload) SetImageHeight(v float32) {
	o.ImageHeight.Set(&v)
}

// GetDescription returns the Description field value
func (o *Upload) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Upload) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Upload) SetDescription(v string) {
	o.Description = v
}

// GetUploaderName returns the UploaderName field value
func (o *Upload) GetUploaderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UploaderName
}

// GetUploaderNameOk returns a tuple with the UploaderName field value
// and a boolean to check if the value has been set.
func (o *Upload) GetUploaderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploaderName, true
}

// SetUploaderName sets field value
func (o *Upload) SetUploaderName(v string) {
	o.UploaderName = v
}

func (o Upload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Upload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["source"] = o.Source
	toSerialize["rating"] = o.Rating
	toSerialize["uploader_id"] = o.UploaderId
	toSerialize["tag_string"] = o.TagString
	toSerialize["status"] = o.Status
	toSerialize["backtrace"] = o.Backtrace.Get()
	toSerialize["post_id"] = o.PostId.Get()
	if o.Md5Confirmation != nil {
		toSerialize["md5_confirmation"] = o.Md5Confirmation
	}
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["parent_id"] = o.ParentId.Get()
	toSerialize["md5"] = o.Md5.Get()
	toSerialize["file_ext"] = o.FileExt
	toSerialize["file_size"] = o.FileSize.Get()
	toSerialize["image_width"] = o.ImageWidth.Get()
	toSerialize["image_height"] = o.ImageHeight.Get()
	toSerialize["description"] = o.Description
	toSerialize["uploader_name"] = o.UploaderName
	return toSerialize, nil
}

func (o *Upload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"source",
		"rating",
		"uploader_id",
		"tag_string",
		"status",
		"backtrace",
		"post_id",
		"md5_confirmation",
		"created_at",
		"updated_at",
		"parent_id",
		"md5",
		"file_ext",
		"file_size",
		"image_width",
		"image_height",
		"description",
		"uploader_name",
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

	varUpload := _Upload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpload)

	if err != nil {
		return err
	}

	*o = Upload(varUpload)

	return err
}

type NullableUpload struct {
	value *Upload
	isSet bool
}

func (v NullableUpload) Get() *Upload {
	return v.value
}

func (v *NullableUpload) Set(val *Upload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpload(val *Upload) *NullableUpload {
	return &NullableUpload{value: val, isSet: true}
}

func (v NullableUpload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
