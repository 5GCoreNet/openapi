/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// checks if the SscModes1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SscModes1{}

// SscModes1 struct for SscModes1
type SscModes1 struct {
	DefaultSscMode  SscMode   `json:"defaultSscMode"`
	AllowedSscModes []SscMode `json:"allowedSscModes,omitempty"`
}

// NewSscModes1 instantiates a new SscModes1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSscModes1(defaultSscMode SscMode) *SscModes1 {
	this := SscModes1{}
	this.DefaultSscMode = defaultSscMode
	return &this
}

// NewSscModes1WithDefaults instantiates a new SscModes1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSscModes1WithDefaults() *SscModes1 {
	this := SscModes1{}
	return &this
}

// GetDefaultSscMode returns the DefaultSscMode field value
func (o *SscModes1) GetDefaultSscMode() SscMode {
	if o == nil {
		var ret SscMode
		return ret
	}

	return o.DefaultSscMode
}

// GetDefaultSscModeOk returns a tuple with the DefaultSscMode field value
// and a boolean to check if the value has been set.
func (o *SscModes1) GetDefaultSscModeOk() (*SscMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultSscMode, true
}

// SetDefaultSscMode sets field value
func (o *SscModes1) SetDefaultSscMode(v SscMode) {
	o.DefaultSscMode = v
}

// GetAllowedSscModes returns the AllowedSscModes field value if set, zero value otherwise.
func (o *SscModes1) GetAllowedSscModes() []SscMode {
	if o == nil || IsNil(o.AllowedSscModes) {
		var ret []SscMode
		return ret
	}
	return o.AllowedSscModes
}

// GetAllowedSscModesOk returns a tuple with the AllowedSscModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SscModes1) GetAllowedSscModesOk() ([]SscMode, bool) {
	if o == nil || IsNil(o.AllowedSscModes) {
		return nil, false
	}
	return o.AllowedSscModes, true
}

// HasAllowedSscModes returns a boolean if a field has been set.
func (o *SscModes1) HasAllowedSscModes() bool {
	if o != nil && !IsNil(o.AllowedSscModes) {
		return true
	}

	return false
}

// SetAllowedSscModes gets a reference to the given []SscMode and assigns it to the AllowedSscModes field.
func (o *SscModes1) SetAllowedSscModes(v []SscMode) {
	o.AllowedSscModes = v
}

func (o SscModes1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SscModes1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["defaultSscMode"] = o.DefaultSscMode
	if !IsNil(o.AllowedSscModes) {
		toSerialize["allowedSscModes"] = o.AllowedSscModes
	}
	return toSerialize, nil
}

type NullableSscModes1 struct {
	value *SscModes1
	isSet bool
}

func (v NullableSscModes1) Get() *SscModes1 {
	return v.value
}

func (v *NullableSscModes1) Set(val *SscModes1) {
	v.value = val
	v.isSet = true
}

func (v NullableSscModes1) IsSet() bool {
	return v.isSet
}

func (v *NullableSscModes1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSscModes1(val *SscModes1) *NullableSscModes1 {
	return &NullableSscModes1{value: val, isSet: true}
}

func (v NullableSscModes1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSscModes1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
