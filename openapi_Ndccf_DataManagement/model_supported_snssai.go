/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
)

// checks if the SupportedSnssai type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SupportedSnssai{}

// SupportedSnssai Supported S-NSSAIs
type SupportedSnssai struct {
	SNssai         ExtSnssai `json:"sNssai"`
	RestrictionInd *bool     `json:"restrictionInd,omitempty"`
}

// NewSupportedSnssai instantiates a new SupportedSnssai object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSupportedSnssai(sNssai ExtSnssai) *SupportedSnssai {
	this := SupportedSnssai{}
	this.SNssai = sNssai
	var restrictionInd bool = false
	this.RestrictionInd = &restrictionInd
	return &this
}

// NewSupportedSnssaiWithDefaults instantiates a new SupportedSnssai object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSupportedSnssaiWithDefaults() *SupportedSnssai {
	this := SupportedSnssai{}
	var restrictionInd bool = false
	this.RestrictionInd = &restrictionInd
	return &this
}

// GetSNssai returns the SNssai field value
func (o *SupportedSnssai) GetSNssai() ExtSnssai {
	if o == nil {
		var ret ExtSnssai
		return ret
	}

	return o.SNssai
}

// GetSNssaiOk returns a tuple with the SNssai field value
// and a boolean to check if the value has been set.
func (o *SupportedSnssai) GetSNssaiOk() (*ExtSnssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SNssai, true
}

// SetSNssai sets field value
func (o *SupportedSnssai) SetSNssai(v ExtSnssai) {
	o.SNssai = v
}

// GetRestrictionInd returns the RestrictionInd field value if set, zero value otherwise.
func (o *SupportedSnssai) GetRestrictionInd() bool {
	if o == nil || IsNil(o.RestrictionInd) {
		var ret bool
		return ret
	}
	return *o.RestrictionInd
}

// GetRestrictionIndOk returns a tuple with the RestrictionInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SupportedSnssai) GetRestrictionIndOk() (*bool, bool) {
	if o == nil || IsNil(o.RestrictionInd) {
		return nil, false
	}
	return o.RestrictionInd, true
}

// HasRestrictionInd returns a boolean if a field has been set.
func (o *SupportedSnssai) HasRestrictionInd() bool {
	if o != nil && !IsNil(o.RestrictionInd) {
		return true
	}

	return false
}

// SetRestrictionInd gets a reference to the given bool and assigns it to the RestrictionInd field.
func (o *SupportedSnssai) SetRestrictionInd(v bool) {
	o.RestrictionInd = &v
}

func (o SupportedSnssai) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SupportedSnssai) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sNssai"] = o.SNssai
	if !IsNil(o.RestrictionInd) {
		toSerialize["restrictionInd"] = o.RestrictionInd
	}
	return toSerialize, nil
}

type NullableSupportedSnssai struct {
	value *SupportedSnssai
	isSet bool
}

func (v NullableSupportedSnssai) Get() *SupportedSnssai {
	return v.value
}

func (v *NullableSupportedSnssai) Set(val *SupportedSnssai) {
	v.value = val
	v.isSet = true
}

func (v NullableSupportedSnssai) IsSet() bool {
	return v.isSet
}

func (v *NullableSupportedSnssai) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSupportedSnssai(val *SupportedSnssai) *NullableSupportedSnssai {
	return &NullableSupportedSnssai{value: val, isSet: true}
}

func (v NullableSupportedSnssai) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSupportedSnssai) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
