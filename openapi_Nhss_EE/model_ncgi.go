/*
Nhss_EE

HSS Event Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_EE

import (
	"encoding/json"
)

// checks if the Ncgi type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ncgi{}

// Ncgi Contains the NCGI (NR Cell Global Identity), as described in 3GPP 23.003
type Ncgi struct {
	PlmnId PlmnId `json:"plmnId"`
	// 36-bit string identifying an NR Cell Id as specified in clause 9.3.1.7 of 3GPP TS 38.413,  in hexadecimal representation. Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent 4 bits. The most significant character  representing the 4 most significant bits of the Cell Id shall appear first in the string, and  the character representing the 4 least significant bit of the Cell Id shall appear last in the  string.  
	NrCellId string `json:"nrCellId"`
	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).  
	Nid *string `json:"nid,omitempty"`
}

// NewNcgi instantiates a new Ncgi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNcgi(plmnId PlmnId, nrCellId string) *Ncgi {
	this := Ncgi{}
	this.PlmnId = plmnId
	this.NrCellId = nrCellId
	return &this
}

// NewNcgiWithDefaults instantiates a new Ncgi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNcgiWithDefaults() *Ncgi {
	this := Ncgi{}
	return &this
}

// GetPlmnId returns the PlmnId field value
func (o *Ncgi) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *Ncgi) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *Ncgi) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetNrCellId returns the NrCellId field value
func (o *Ncgi) GetNrCellId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NrCellId
}

// GetNrCellIdOk returns a tuple with the NrCellId field value
// and a boolean to check if the value has been set.
func (o *Ncgi) GetNrCellIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NrCellId, true
}

// SetNrCellId sets field value
func (o *Ncgi) SetNrCellId(v string) {
	o.NrCellId = v
}

// GetNid returns the Nid field value if set, zero value otherwise.
func (o *Ncgi) GetNid() string {
	if o == nil || IsNil(o.Nid) {
		var ret string
		return ret
	}
	return *o.Nid
}

// GetNidOk returns a tuple with the Nid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ncgi) GetNidOk() (*string, bool) {
	if o == nil || IsNil(o.Nid) {
		return nil, false
	}
	return o.Nid, true
}

// HasNid returns a boolean if a field has been set.
func (o *Ncgi) HasNid() bool {
	if o != nil && !IsNil(o.Nid) {
		return true
	}

	return false
}

// SetNid gets a reference to the given string and assigns it to the Nid field.
func (o *Ncgi) SetNid(v string) {
	o.Nid = &v
}

func (o Ncgi) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ncgi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plmnId"] = o.PlmnId
	toSerialize["nrCellId"] = o.NrCellId
	if !IsNil(o.Nid) {
		toSerialize["nid"] = o.Nid
	}
	return toSerialize, nil
}

type NullableNcgi struct {
	value *Ncgi
	isSet bool
}

func (v NullableNcgi) Get() *Ncgi {
	return v.value
}

func (v *NullableNcgi) Set(val *Ncgi) {
	v.value = val
	v.isSet = true
}

func (v NullableNcgi) IsSet() bool {
	return v.isSet
}

func (v *NullableNcgi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNcgi(val *Ncgi) *NullableNcgi {
	return &NullableNcgi{value: val, isSet: true}
}

func (v NullableNcgi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNcgi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


