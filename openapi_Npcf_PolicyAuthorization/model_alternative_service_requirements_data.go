/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
)

// checks if the AlternativeServiceRequirementsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlternativeServiceRequirementsData{}

// AlternativeServiceRequirementsData Contains an alternative QoS related parameter set.
type AlternativeServiceRequirementsData struct {
	// Reference to this alternative QoS related parameter set.
	AltQosParamSetRef string `json:"altQosParamSetRef"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	GbrUl *string `json:"gbrUl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	GbrDl *string `json:"gbrDl,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	Pdb *int32 `json:"pdb,omitempty"`
}

// NewAlternativeServiceRequirementsData instantiates a new AlternativeServiceRequirementsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlternativeServiceRequirementsData(altQosParamSetRef string) *AlternativeServiceRequirementsData {
	this := AlternativeServiceRequirementsData{}
	this.AltQosParamSetRef = altQosParamSetRef
	return &this
}

// NewAlternativeServiceRequirementsDataWithDefaults instantiates a new AlternativeServiceRequirementsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlternativeServiceRequirementsDataWithDefaults() *AlternativeServiceRequirementsData {
	this := AlternativeServiceRequirementsData{}
	return &this
}

// GetAltQosParamSetRef returns the AltQosParamSetRef field value
func (o *AlternativeServiceRequirementsData) GetAltQosParamSetRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AltQosParamSetRef
}

// GetAltQosParamSetRefOk returns a tuple with the AltQosParamSetRef field value
// and a boolean to check if the value has been set.
func (o *AlternativeServiceRequirementsData) GetAltQosParamSetRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AltQosParamSetRef, true
}

// SetAltQosParamSetRef sets field value
func (o *AlternativeServiceRequirementsData) SetAltQosParamSetRef(v string) {
	o.AltQosParamSetRef = v
}

// GetGbrUl returns the GbrUl field value if set, zero value otherwise.
func (o *AlternativeServiceRequirementsData) GetGbrUl() string {
	if o == nil || IsNil(o.GbrUl) {
		var ret string
		return ret
	}
	return *o.GbrUl
}

// GetGbrUlOk returns a tuple with the GbrUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeServiceRequirementsData) GetGbrUlOk() (*string, bool) {
	if o == nil || IsNil(o.GbrUl) {
		return nil, false
	}
	return o.GbrUl, true
}

// HasGbrUl returns a boolean if a field has been set.
func (o *AlternativeServiceRequirementsData) HasGbrUl() bool {
	if o != nil && !IsNil(o.GbrUl) {
		return true
	}

	return false
}

// SetGbrUl gets a reference to the given string and assigns it to the GbrUl field.
func (o *AlternativeServiceRequirementsData) SetGbrUl(v string) {
	o.GbrUl = &v
}

// GetGbrDl returns the GbrDl field value if set, zero value otherwise.
func (o *AlternativeServiceRequirementsData) GetGbrDl() string {
	if o == nil || IsNil(o.GbrDl) {
		var ret string
		return ret
	}
	return *o.GbrDl
}

// GetGbrDlOk returns a tuple with the GbrDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeServiceRequirementsData) GetGbrDlOk() (*string, bool) {
	if o == nil || IsNil(o.GbrDl) {
		return nil, false
	}
	return o.GbrDl, true
}

// HasGbrDl returns a boolean if a field has been set.
func (o *AlternativeServiceRequirementsData) HasGbrDl() bool {
	if o != nil && !IsNil(o.GbrDl) {
		return true
	}

	return false
}

// SetGbrDl gets a reference to the given string and assigns it to the GbrDl field.
func (o *AlternativeServiceRequirementsData) SetGbrDl(v string) {
	o.GbrDl = &v
}

// GetPdb returns the Pdb field value if set, zero value otherwise.
func (o *AlternativeServiceRequirementsData) GetPdb() int32 {
	if o == nil || IsNil(o.Pdb) {
		var ret int32
		return ret
	}
	return *o.Pdb
}

// GetPdbOk returns a tuple with the Pdb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlternativeServiceRequirementsData) GetPdbOk() (*int32, bool) {
	if o == nil || IsNil(o.Pdb) {
		return nil, false
	}
	return o.Pdb, true
}

// HasPdb returns a boolean if a field has been set.
func (o *AlternativeServiceRequirementsData) HasPdb() bool {
	if o != nil && !IsNil(o.Pdb) {
		return true
	}

	return false
}

// SetPdb gets a reference to the given int32 and assigns it to the Pdb field.
func (o *AlternativeServiceRequirementsData) SetPdb(v int32) {
	o.Pdb = &v
}

func (o AlternativeServiceRequirementsData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlternativeServiceRequirementsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["altQosParamSetRef"] = o.AltQosParamSetRef
	if !IsNil(o.GbrUl) {
		toSerialize["gbrUl"] = o.GbrUl
	}
	if !IsNil(o.GbrDl) {
		toSerialize["gbrDl"] = o.GbrDl
	}
	if !IsNil(o.Pdb) {
		toSerialize["pdb"] = o.Pdb
	}
	return toSerialize, nil
}

type NullableAlternativeServiceRequirementsData struct {
	value *AlternativeServiceRequirementsData
	isSet bool
}

func (v NullableAlternativeServiceRequirementsData) Get() *AlternativeServiceRequirementsData {
	return v.value
}

func (v *NullableAlternativeServiceRequirementsData) Set(val *AlternativeServiceRequirementsData) {
	v.value = val
	v.isSet = true
}

func (v NullableAlternativeServiceRequirementsData) IsSet() bool {
	return v.isSet
}

func (v *NullableAlternativeServiceRequirementsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlternativeServiceRequirementsData(val *AlternativeServiceRequirementsData) *NullableAlternativeServiceRequirementsData {
	return &NullableAlternativeServiceRequirementsData{value: val, isSet: true}
}

func (v NullableAlternativeServiceRequirementsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlternativeServiceRequirementsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
