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
)

// checks if the PostRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostRelationships{}

// PostRelationships struct for PostRelationships
type PostRelationships struct {
	ParentId          NullableInt32 `json:"parent_id"`
	HasChildren       bool          `json:"has_children"`
	HasActiveChildren bool          `json:"has_active_children"`
	Children          []int32       `json:"children"`
}

type _PostRelationships PostRelationships

// NewPostRelationships instantiates a new PostRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostRelationships(parentId NullableInt32, hasChildren bool, hasActiveChildren bool, children []int32) *PostRelationships {
	this := PostRelationships{}
	this.ParentId = parentId
	this.HasChildren = hasChildren
	this.HasActiveChildren = hasActiveChildren
	this.Children = children
	return &this
}

// NewPostRelationshipsWithDefaults instantiates a new PostRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostRelationshipsWithDefaults() *PostRelationships {
	this := PostRelationships{}
	return &this
}

// GetParentId returns the ParentId field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *PostRelationships) GetParentId() int32 {
	if o == nil || o.ParentId.Get() == nil {
		var ret int32
		return ret
	}

	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostRelationships) GetParentIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// SetParentId sets field value
func (o *PostRelationships) SetParentId(v int32) {
	o.ParentId.Set(&v)
}

// GetHasChildren returns the HasChildren field value
func (o *PostRelationships) GetHasChildren() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasChildren
}

// GetHasChildrenOk returns a tuple with the HasChildren field value
// and a boolean to check if the value has been set.
func (o *PostRelationships) GetHasChildrenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasChildren, true
}

// SetHasChildren sets field value
func (o *PostRelationships) SetHasChildren(v bool) {
	o.HasChildren = v
}

// GetHasActiveChildren returns the HasActiveChildren field value
func (o *PostRelationships) GetHasActiveChildren() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasActiveChildren
}

// GetHasActiveChildrenOk returns a tuple with the HasActiveChildren field value
// and a boolean to check if the value has been set.
func (o *PostRelationships) GetHasActiveChildrenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasActiveChildren, true
}

// SetHasActiveChildren sets field value
func (o *PostRelationships) SetHasActiveChildren(v bool) {
	o.HasActiveChildren = v
}

// GetChildren returns the Children field value
func (o *PostRelationships) GetChildren() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
func (o *PostRelationships) GetChildrenOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Children, true
}

// SetChildren sets field value
func (o *PostRelationships) SetChildren(v []int32) {
	o.Children = v
}

func (o PostRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["parent_id"] = o.ParentId.Get()
	toSerialize["has_children"] = o.HasChildren
	toSerialize["has_active_children"] = o.HasActiveChildren
	toSerialize["children"] = o.Children
	return toSerialize, nil
}

func (o *PostRelationships) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"parent_id",
		"has_children",
		"has_active_children",
		"children",
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

	varPostRelationships := _PostRelationships{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostRelationships)

	if err != nil {
		return err
	}

	*o = PostRelationships(varPostRelationships)

	return err
}

type NullablePostRelationships struct {
	value *PostRelationships
	isSet bool
}

func (v NullablePostRelationships) Get() *PostRelationships {
	return v.value
}

func (v *NullablePostRelationships) Set(val *PostRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePostRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePostRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostRelationships(val *PostRelationships) *NullablePostRelationships {
	return &NullablePostRelationships{value: val, isSet: true}
}

func (v NullablePostRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
