/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
)

// checks if the SscModes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SscModes{}

// SscModes struct for SscModes
type SscModes struct {
	DefaultSscMode SscMode `json:"defaultSscMode"`
	AllowedSscModes []SscMode `json:"allowedSscModes,omitempty"`
}

// NewSscModes instantiates a new SscModes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSscModes(defaultSscMode SscMode) *SscModes {
	this := SscModes{}
	this.DefaultSscMode = defaultSscMode
	return &this
}

// NewSscModesWithDefaults instantiates a new SscModes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSscModesWithDefaults() *SscModes {
	this := SscModes{}
	return &this
}

// GetDefaultSscMode returns the DefaultSscMode field value
func (o *SscModes) GetDefaultSscMode() SscMode {
	if o == nil {
		var ret SscMode
		return ret
	}

	return o.DefaultSscMode
}

// GetDefaultSscModeOk returns a tuple with the DefaultSscMode field value
// and a boolean to check if the value has been set.
func (o *SscModes) GetDefaultSscModeOk() (*SscMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultSscMode, true
}

// SetDefaultSscMode sets field value
func (o *SscModes) SetDefaultSscMode(v SscMode) {
	o.DefaultSscMode = v
}

// GetAllowedSscModes returns the AllowedSscModes field value if set, zero value otherwise.
func (o *SscModes) GetAllowedSscModes() []SscMode {
	if o == nil || IsNil(o.AllowedSscModes) {
		var ret []SscMode
		return ret
	}
	return o.AllowedSscModes
}

// GetAllowedSscModesOk returns a tuple with the AllowedSscModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SscModes) GetAllowedSscModesOk() ([]SscMode, bool) {
	if o == nil || IsNil(o.AllowedSscModes) {
		return nil, false
	}
	return o.AllowedSscModes, true
}

// HasAllowedSscModes returns a boolean if a field has been set.
func (o *SscModes) HasAllowedSscModes() bool {
	if o != nil && !IsNil(o.AllowedSscModes) {
		return true
	}

	return false
}

// SetAllowedSscModes gets a reference to the given []SscMode and assigns it to the AllowedSscModes field.
func (o *SscModes) SetAllowedSscModes(v []SscMode) {
	o.AllowedSscModes = v
}

func (o SscModes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SscModes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["defaultSscMode"] = o.DefaultSscMode
	if !IsNil(o.AllowedSscModes) {
		toSerialize["allowedSscModes"] = o.AllowedSscModes
	}
	return toSerialize, nil
}

type NullableSscModes struct {
	value *SscModes
	isSet bool
}

func (v NullableSscModes) Get() *SscModes {
	return v.value
}

func (v *NullableSscModes) Set(val *SscModes) {
	v.value = val
	v.isSet = true
}

func (v NullableSscModes) IsSet() bool {
	return v.isSet
}

func (v *NullableSscModes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSscModes(val *SscModes) *NullableSscModes {
	return &NullableSscModes{value: val, isSet: true}
}

func (v NullableSscModes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSscModes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

