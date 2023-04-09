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

// checks if the PlmnSnssai type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlmnSnssai{}

// PlmnSnssai List of network slices (S-NSSAIs) for a given PLMN ID
type PlmnSnssai struct {
	PlmnId     PlmnId      `json:"plmnId"`
	SNssaiList []ExtSnssai `json:"sNssaiList"`
	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).
	Nid *string `json:"nid,omitempty"`
}

// NewPlmnSnssai instantiates a new PlmnSnssai object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlmnSnssai(plmnId PlmnId, sNssaiList []ExtSnssai) *PlmnSnssai {
	this := PlmnSnssai{}
	this.PlmnId = plmnId
	this.SNssaiList = sNssaiList
	return &this
}

// NewPlmnSnssaiWithDefaults instantiates a new PlmnSnssai object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlmnSnssaiWithDefaults() *PlmnSnssai {
	this := PlmnSnssai{}
	return &this
}

// GetPlmnId returns the PlmnId field value
func (o *PlmnSnssai) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *PlmnSnssai) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *PlmnSnssai) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetSNssaiList returns the SNssaiList field value
func (o *PlmnSnssai) GetSNssaiList() []ExtSnssai {
	if o == nil {
		var ret []ExtSnssai
		return ret
	}

	return o.SNssaiList
}

// GetSNssaiListOk returns a tuple with the SNssaiList field value
// and a boolean to check if the value has been set.
func (o *PlmnSnssai) GetSNssaiListOk() ([]ExtSnssai, bool) {
	if o == nil {
		return nil, false
	}
	return o.SNssaiList, true
}

// SetSNssaiList sets field value
func (o *PlmnSnssai) SetSNssaiList(v []ExtSnssai) {
	o.SNssaiList = v
}

// GetNid returns the Nid field value if set, zero value otherwise.
func (o *PlmnSnssai) GetNid() string {
	if o == nil || IsNil(o.Nid) {
		var ret string
		return ret
	}
	return *o.Nid
}

// GetNidOk returns a tuple with the Nid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlmnSnssai) GetNidOk() (*string, bool) {
	if o == nil || IsNil(o.Nid) {
		return nil, false
	}
	return o.Nid, true
}

// HasNid returns a boolean if a field has been set.
func (o *PlmnSnssai) HasNid() bool {
	if o != nil && !IsNil(o.Nid) {
		return true
	}

	return false
}

// SetNid gets a reference to the given string and assigns it to the Nid field.
func (o *PlmnSnssai) SetNid(v string) {
	o.Nid = &v
}

func (o PlmnSnssai) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlmnSnssai) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plmnId"] = o.PlmnId
	toSerialize["sNssaiList"] = o.SNssaiList
	if !IsNil(o.Nid) {
		toSerialize["nid"] = o.Nid
	}
	return toSerialize, nil
}

type NullablePlmnSnssai struct {
	value *PlmnSnssai
	isSet bool
}

func (v NullablePlmnSnssai) Get() *PlmnSnssai {
	return v.value
}

func (v *NullablePlmnSnssai) Set(val *PlmnSnssai) {
	v.value = val
	v.isSet = true
}

func (v NullablePlmnSnssai) IsSet() bool {
	return v.isSet
}

func (v *NullablePlmnSnssai) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlmnSnssai(val *PlmnSnssai) *NullablePlmnSnssai {
	return &NullablePlmnSnssai{value: val, isSet: true}
}

func (v NullablePlmnSnssai) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlmnSnssai) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
