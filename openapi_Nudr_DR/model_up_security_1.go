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

// checks if the UpSecurity1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpSecurity1{}

// UpSecurity1 Contains Userplain security information.
type UpSecurity1 struct {
	UpIntegr UpIntegrity       `json:"upIntegr"`
	UpConfid UpConfidentiality `json:"upConfid"`
}

// NewUpSecurity1 instantiates a new UpSecurity1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpSecurity1(upIntegr UpIntegrity, upConfid UpConfidentiality) *UpSecurity1 {
	this := UpSecurity1{}
	this.UpIntegr = upIntegr
	this.UpConfid = upConfid
	return &this
}

// NewUpSecurity1WithDefaults instantiates a new UpSecurity1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpSecurity1WithDefaults() *UpSecurity1 {
	this := UpSecurity1{}
	return &this
}

// GetUpIntegr returns the UpIntegr field value
func (o *UpSecurity1) GetUpIntegr() UpIntegrity {
	if o == nil {
		var ret UpIntegrity
		return ret
	}

	return o.UpIntegr
}

// GetUpIntegrOk returns a tuple with the UpIntegr field value
// and a boolean to check if the value has been set.
func (o *UpSecurity1) GetUpIntegrOk() (*UpIntegrity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpIntegr, true
}

// SetUpIntegr sets field value
func (o *UpSecurity1) SetUpIntegr(v UpIntegrity) {
	o.UpIntegr = v
}

// GetUpConfid returns the UpConfid field value
func (o *UpSecurity1) GetUpConfid() UpConfidentiality {
	if o == nil {
		var ret UpConfidentiality
		return ret
	}

	return o.UpConfid
}

// GetUpConfidOk returns a tuple with the UpConfid field value
// and a boolean to check if the value has been set.
func (o *UpSecurity1) GetUpConfidOk() (*UpConfidentiality, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpConfid, true
}

// SetUpConfid sets field value
func (o *UpSecurity1) SetUpConfid(v UpConfidentiality) {
	o.UpConfid = v
}

func (o UpSecurity1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpSecurity1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["upIntegr"] = o.UpIntegr
	toSerialize["upConfid"] = o.UpConfid
	return toSerialize, nil
}

type NullableUpSecurity1 struct {
	value *UpSecurity1
	isSet bool
}

func (v NullableUpSecurity1) Get() *UpSecurity1 {
	return v.value
}

func (v *NullableUpSecurity1) Set(val *UpSecurity1) {
	v.value = val
	v.isSet = true
}

func (v NullableUpSecurity1) IsSet() bool {
	return v.isSet
}

func (v *NullableUpSecurity1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpSecurity1(val *UpSecurity1) *NullableUpSecurity1 {
	return &NullableUpSecurity1{value: val, isSet: true}
}

func (v NullableUpSecurity1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpSecurity1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
