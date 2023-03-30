/*
SS_Events

API for SEAL Events management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_Events

import (
	"encoding/json"
)

// checks if the ValidityConditions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidityConditions{}

// ValidityConditions List of monitoring and/or analytics events related to VAL UE.
type ValidityConditions struct {
	LocArea *LocationArea5G `json:"locArea,omitempty"`
	// Time window validity conditions.
	TmWdws []TimeWindow `json:"tmWdws,omitempty"`
}

// NewValidityConditions instantiates a new ValidityConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidityConditions() *ValidityConditions {
	this := ValidityConditions{}
	return &this
}

// NewValidityConditionsWithDefaults instantiates a new ValidityConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidityConditionsWithDefaults() *ValidityConditions {
	this := ValidityConditions{}
	return &this
}

// GetLocArea returns the LocArea field value if set, zero value otherwise.
func (o *ValidityConditions) GetLocArea() LocationArea5G {
	if o == nil || IsNil(o.LocArea) {
		var ret LocationArea5G
		return ret
	}
	return *o.LocArea
}

// GetLocAreaOk returns a tuple with the LocArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidityConditions) GetLocAreaOk() (*LocationArea5G, bool) {
	if o == nil || IsNil(o.LocArea) {
		return nil, false
	}
	return o.LocArea, true
}

// HasLocArea returns a boolean if a field has been set.
func (o *ValidityConditions) HasLocArea() bool {
	if o != nil && !IsNil(o.LocArea) {
		return true
	}

	return false
}

// SetLocArea gets a reference to the given LocationArea5G and assigns it to the LocArea field.
func (o *ValidityConditions) SetLocArea(v LocationArea5G) {
	o.LocArea = &v
}

// GetTmWdws returns the TmWdws field value if set, zero value otherwise.
func (o *ValidityConditions) GetTmWdws() []TimeWindow {
	if o == nil || IsNil(o.TmWdws) {
		var ret []TimeWindow
		return ret
	}
	return o.TmWdws
}

// GetTmWdwsOk returns a tuple with the TmWdws field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidityConditions) GetTmWdwsOk() ([]TimeWindow, bool) {
	if o == nil || IsNil(o.TmWdws) {
		return nil, false
	}
	return o.TmWdws, true
}

// HasTmWdws returns a boolean if a field has been set.
func (o *ValidityConditions) HasTmWdws() bool {
	if o != nil && !IsNil(o.TmWdws) {
		return true
	}

	return false
}

// SetTmWdws gets a reference to the given []TimeWindow and assigns it to the TmWdws field.
func (o *ValidityConditions) SetTmWdws(v []TimeWindow) {
	o.TmWdws = v
}

func (o ValidityConditions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidityConditions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LocArea) {
		toSerialize["locArea"] = o.LocArea
	}
	if !IsNil(o.TmWdws) {
		toSerialize["tmWdws"] = o.TmWdws
	}
	return toSerialize, nil
}

type NullableValidityConditions struct {
	value *ValidityConditions
	isSet bool
}

func (v NullableValidityConditions) Get() *ValidityConditions {
	return v.value
}

func (v *NullableValidityConditions) Set(val *ValidityConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableValidityConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableValidityConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidityConditions(val *ValidityConditions) *NullableValidityConditions {
	return &NullableValidityConditions{value: val, isSet: true}
}

func (v NullableValidityConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidityConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


