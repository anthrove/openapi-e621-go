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

// checks if the MarkBlipRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarkBlipRequest{}

// MarkBlipRequest struct for MarkBlipRequest
type MarkBlipRequest struct {
	RecordType string `json:"record_type"`
}

type _MarkBlipRequest MarkBlipRequest

// NewMarkBlipRequest instantiates a new MarkBlipRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarkBlipRequest(recordType string) *MarkBlipRequest {
	this := MarkBlipRequest{}
	this.RecordType = recordType
	return &this
}

// NewMarkBlipRequestWithDefaults instantiates a new MarkBlipRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarkBlipRequestWithDefaults() *MarkBlipRequest {
	this := MarkBlipRequest{}
	return &this
}

// GetRecordType returns the RecordType field value
func (o *MarkBlipRequest) GetRecordType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value
// and a boolean to check if the value has been set.
func (o *MarkBlipRequest) GetRecordTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordType, true
}

// SetRecordType sets field value
func (o *MarkBlipRequest) SetRecordType(v string) {
	o.RecordType = v
}

func (o MarkBlipRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarkBlipRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["record_type"] = o.RecordType
	return toSerialize, nil
}

func (o *MarkBlipRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"record_type",
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

	varMarkBlipRequest := _MarkBlipRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMarkBlipRequest)

	if err != nil {
		return err
	}

	*o = MarkBlipRequest(varMarkBlipRequest)

	return err
}

type NullableMarkBlipRequest struct {
	value *MarkBlipRequest
	isSet bool
}

func (v NullableMarkBlipRequest) Get() *MarkBlipRequest {
	return v.value
}

func (v *NullableMarkBlipRequest) Set(val *MarkBlipRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMarkBlipRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMarkBlipRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarkBlipRequest(val *MarkBlipRequest) *NullableMarkBlipRequest {
	return &NullableMarkBlipRequest{value: val, isSet: true}
}

func (v NullableMarkBlipRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarkBlipRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
