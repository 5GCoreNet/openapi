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

// checks if the PeiUpdateInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PeiUpdateInfo{}

// PeiUpdateInfo struct for PeiUpdateInfo
type PeiUpdateInfo struct {
	// String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.  
	Pei string `json:"pei"`
}

// NewPeiUpdateInfo instantiates a new PeiUpdateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeiUpdateInfo(pei string) *PeiUpdateInfo {
	this := PeiUpdateInfo{}
	this.Pei = pei
	return &this
}

// NewPeiUpdateInfoWithDefaults instantiates a new PeiUpdateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeiUpdateInfoWithDefaults() *PeiUpdateInfo {
	this := PeiUpdateInfo{}
	return &this
}

// GetPei returns the Pei field value
func (o *PeiUpdateInfo) GetPei() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pei
}

// GetPeiOk returns a tuple with the Pei field value
// and a boolean to check if the value has been set.
func (o *PeiUpdateInfo) GetPeiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pei, true
}

// SetPei sets field value
func (o *PeiUpdateInfo) SetPei(v string) {
	o.Pei = v
}

func (o PeiUpdateInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PeiUpdateInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pei"] = o.Pei
	return toSerialize, nil
}

type NullablePeiUpdateInfo struct {
	value *PeiUpdateInfo
	isSet bool
}

func (v NullablePeiUpdateInfo) Get() *PeiUpdateInfo {
	return v.value
}

func (v *NullablePeiUpdateInfo) Set(val *PeiUpdateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePeiUpdateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePeiUpdateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeiUpdateInfo(val *PeiUpdateInfo) *NullablePeiUpdateInfo {
	return &NullablePeiUpdateInfo{value: val, isSet: true}
}

func (v NullablePeiUpdateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeiUpdateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

