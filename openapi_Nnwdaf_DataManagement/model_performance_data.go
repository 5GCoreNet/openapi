/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
)

// checks if the PerformanceData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerformanceData{}

// PerformanceData Contains Performance Data.
type PerformanceData struct {
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds. 
	Pdb *int32 `json:"pdb,omitempty"`
	// Unsigned integer indicating Packet Loss Rate (see clauses 5.7.2.8 and 5.7.4 of 3GPP TS 23.501), expressed in tenth of percent. 
	Plr *int32 `json:"plr,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	ThrputUl *string `json:"thrputUl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	ThrputDl *string `json:"thrputDl,omitempty"`
}

// NewPerformanceData instantiates a new PerformanceData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerformanceData() *PerformanceData {
	this := PerformanceData{}
	return &this
}

// NewPerformanceDataWithDefaults instantiates a new PerformanceData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerformanceDataWithDefaults() *PerformanceData {
	this := PerformanceData{}
	return &this
}

// GetPdb returns the Pdb field value if set, zero value otherwise.
func (o *PerformanceData) GetPdb() int32 {
	if o == nil || isNil(o.Pdb) {
		var ret int32
		return ret
	}
	return *o.Pdb
}

// GetPdbOk returns a tuple with the Pdb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceData) GetPdbOk() (*int32, bool) {
	if o == nil || isNil(o.Pdb) {
		return nil, false
	}
	return o.Pdb, true
}

// HasPdb returns a boolean if a field has been set.
func (o *PerformanceData) HasPdb() bool {
	if o != nil && !isNil(o.Pdb) {
		return true
	}

	return false
}

// SetPdb gets a reference to the given int32 and assigns it to the Pdb field.
func (o *PerformanceData) SetPdb(v int32) {
	o.Pdb = &v
}

// GetPlr returns the Plr field value if set, zero value otherwise.
func (o *PerformanceData) GetPlr() int32 {
	if o == nil || isNil(o.Plr) {
		var ret int32
		return ret
	}
	return *o.Plr
}

// GetPlrOk returns a tuple with the Plr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceData) GetPlrOk() (*int32, bool) {
	if o == nil || isNil(o.Plr) {
		return nil, false
	}
	return o.Plr, true
}

// HasPlr returns a boolean if a field has been set.
func (o *PerformanceData) HasPlr() bool {
	if o != nil && !isNil(o.Plr) {
		return true
	}

	return false
}

// SetPlr gets a reference to the given int32 and assigns it to the Plr field.
func (o *PerformanceData) SetPlr(v int32) {
	o.Plr = &v
}

// GetThrputUl returns the ThrputUl field value if set, zero value otherwise.
func (o *PerformanceData) GetThrputUl() string {
	if o == nil || isNil(o.ThrputUl) {
		var ret string
		return ret
	}
	return *o.ThrputUl
}

// GetThrputUlOk returns a tuple with the ThrputUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceData) GetThrputUlOk() (*string, bool) {
	if o == nil || isNil(o.ThrputUl) {
		return nil, false
	}
	return o.ThrputUl, true
}

// HasThrputUl returns a boolean if a field has been set.
func (o *PerformanceData) HasThrputUl() bool {
	if o != nil && !isNil(o.ThrputUl) {
		return true
	}

	return false
}

// SetThrputUl gets a reference to the given string and assigns it to the ThrputUl field.
func (o *PerformanceData) SetThrputUl(v string) {
	o.ThrputUl = &v
}

// GetThrputDl returns the ThrputDl field value if set, zero value otherwise.
func (o *PerformanceData) GetThrputDl() string {
	if o == nil || isNil(o.ThrputDl) {
		var ret string
		return ret
	}
	return *o.ThrputDl
}

// GetThrputDlOk returns a tuple with the ThrputDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerformanceData) GetThrputDlOk() (*string, bool) {
	if o == nil || isNil(o.ThrputDl) {
		return nil, false
	}
	return o.ThrputDl, true
}

// HasThrputDl returns a boolean if a field has been set.
func (o *PerformanceData) HasThrputDl() bool {
	if o != nil && !isNil(o.ThrputDl) {
		return true
	}

	return false
}

// SetThrputDl gets a reference to the given string and assigns it to the ThrputDl field.
func (o *PerformanceData) SetThrputDl(v string) {
	o.ThrputDl = &v
}

func (o PerformanceData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PerformanceData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Pdb) {
		toSerialize["pdb"] = o.Pdb
	}
	if !isNil(o.Plr) {
		toSerialize["plr"] = o.Plr
	}
	if !isNil(o.ThrputUl) {
		toSerialize["thrputUl"] = o.ThrputUl
	}
	if !isNil(o.ThrputDl) {
		toSerialize["thrputDl"] = o.ThrputDl
	}
	return toSerialize, nil
}

type NullablePerformanceData struct {
	value *PerformanceData
	isSet bool
}

func (v NullablePerformanceData) Get() *PerformanceData {
	return v.value
}

func (v *NullablePerformanceData) Set(val *PerformanceData) {
	v.value = val
	v.isSet = true
}

func (v NullablePerformanceData) IsSet() bool {
	return v.isSet
}

func (v *NullablePerformanceData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerformanceData(val *PerformanceData) *NullablePerformanceData {
	return &NullablePerformanceData{value: val, isSet: true}
}

func (v NullablePerformanceData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerformanceData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


