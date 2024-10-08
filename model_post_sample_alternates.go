/*
E621

OpenAPI definition for E621's API. You can find the source [here](https://github.com/DonovanDMC/E621OpenAPI)<br> This document is intended to compliment E621's existing [API Documentation](https://e621.net/help/api).<br> Note if E621's api is under attack and/or cloudflare protections are enabled, the \"Try it out\" buttons here will not work.<br> If they are not working, you can check this [Unofficial Status Page](https://status.e621.ws). 

API version: d69c34e
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi-e621-go

import (
	"encoding/json"
)

// checks if the PostSampleAlternates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostSampleAlternates{}

// PostSampleAlternates struct for PostSampleAlternates
type PostSampleAlternates struct {
	Var480p *PostSampleAlternate `json:"480p,omitempty"`
	Var720p *PostSampleAlternate `json:"720p,omitempty"`
	Original *PostSampleAlternate `json:"original,omitempty"`
}

// NewPostSampleAlternates instantiates a new PostSampleAlternates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostSampleAlternates() *PostSampleAlternates {
	this := PostSampleAlternates{}
	return &this
}

// NewPostSampleAlternatesWithDefaults instantiates a new PostSampleAlternates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostSampleAlternatesWithDefaults() *PostSampleAlternates {
	this := PostSampleAlternates{}
	return &this
}

// GetVar480p returns the Var480p field value if set, zero value otherwise.
func (o *PostSampleAlternates) GetVar480p() PostSampleAlternate {
	if o == nil || IsNil(o.Var480p) {
		var ret PostSampleAlternate
		return ret
	}
	return *o.Var480p
}

// GetVar480pOk returns a tuple with the Var480p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSampleAlternates) GetVar480pOk() (*PostSampleAlternate, bool) {
	if o == nil || IsNil(o.Var480p) {
		return nil, false
	}
	return o.Var480p, true
}

// HasVar480p returns a boolean if a field has been set.
func (o *PostSampleAlternates) HasVar480p() bool {
	if o != nil && !IsNil(o.Var480p) {
		return true
	}

	return false
}

// SetVar480p gets a reference to the given PostSampleAlternate and assigns it to the Var480p field.
func (o *PostSampleAlternates) SetVar480p(v PostSampleAlternate) {
	o.Var480p = &v
}

// GetVar720p returns the Var720p field value if set, zero value otherwise.
func (o *PostSampleAlternates) GetVar720p() PostSampleAlternate {
	if o == nil || IsNil(o.Var720p) {
		var ret PostSampleAlternate
		return ret
	}
	return *o.Var720p
}

// GetVar720pOk returns a tuple with the Var720p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSampleAlternates) GetVar720pOk() (*PostSampleAlternate, bool) {
	if o == nil || IsNil(o.Var720p) {
		return nil, false
	}
	return o.Var720p, true
}

// HasVar720p returns a boolean if a field has been set.
func (o *PostSampleAlternates) HasVar720p() bool {
	if o != nil && !IsNil(o.Var720p) {
		return true
	}

	return false
}

// SetVar720p gets a reference to the given PostSampleAlternate and assigns it to the Var720p field.
func (o *PostSampleAlternates) SetVar720p(v PostSampleAlternate) {
	o.Var720p = &v
}

// GetOriginal returns the Original field value if set, zero value otherwise.
func (o *PostSampleAlternates) GetOriginal() PostSampleAlternate {
	if o == nil || IsNil(o.Original) {
		var ret PostSampleAlternate
		return ret
	}
	return *o.Original
}

// GetOriginalOk returns a tuple with the Original field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostSampleAlternates) GetOriginalOk() (*PostSampleAlternate, bool) {
	if o == nil || IsNil(o.Original) {
		return nil, false
	}
	return o.Original, true
}

// HasOriginal returns a boolean if a field has been set.
func (o *PostSampleAlternates) HasOriginal() bool {
	if o != nil && !IsNil(o.Original) {
		return true
	}

	return false
}

// SetOriginal gets a reference to the given PostSampleAlternate and assigns it to the Original field.
func (o *PostSampleAlternates) SetOriginal(v PostSampleAlternate) {
	o.Original = &v
}

func (o PostSampleAlternates) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostSampleAlternates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Var480p) {
		toSerialize["480p"] = o.Var480p
	}
	if !IsNil(o.Var720p) {
		toSerialize["720p"] = o.Var720p
	}
	if !IsNil(o.Original) {
		toSerialize["original"] = o.Original
	}
	return toSerialize, nil
}

type NullablePostSampleAlternates struct {
	value *PostSampleAlternates
	isSet bool
}

func (v NullablePostSampleAlternates) Get() *PostSampleAlternates {
	return v.value
}

func (v *NullablePostSampleAlternates) Set(val *PostSampleAlternates) {
	v.value = val
	v.isSet = true
}

func (v NullablePostSampleAlternates) IsSet() bool {
	return v.isSet
}

func (v *NullablePostSampleAlternates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostSampleAlternates(val *PostSampleAlternates) *NullablePostSampleAlternates {
	return &NullablePostSampleAlternates{value: val, isSet: true}
}

func (v NullablePostSampleAlternates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostSampleAlternates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

