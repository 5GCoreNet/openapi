/*
NSSF NS Selection

NSSF Network Slice Selection Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnssf_NSSelection

import (
	"encoding/json"
)

// checks if the TaiRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaiRange{}

// TaiRange Range of TAIs (Tracking Area Identities)
type TaiRange struct {
	PlmnId       PlmnId     `json:"plmnId"`
	TacRangeList []TacRange `json:"tacRangeList"`
	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).
	Nid *string `json:"nid,omitempty"`
}

// NewTaiRange instantiates a new TaiRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaiRange(plmnId PlmnId, tacRangeList []TacRange) *TaiRange {
	this := TaiRange{}
	this.PlmnId = plmnId
	this.TacRangeList = tacRangeList
	return &this
}

// NewTaiRangeWithDefaults instantiates a new TaiRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaiRangeWithDefaults() *TaiRange {
	this := TaiRange{}
	return &this
}

// GetPlmnId returns the PlmnId field value
func (o *TaiRange) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *TaiRange) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *TaiRange) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

// GetTacRangeList returns the TacRangeList field value
func (o *TaiRange) GetTacRangeList() []TacRange {
	if o == nil {
		var ret []TacRange
		return ret
	}

	return o.TacRangeList
}

// GetTacRangeListOk returns a tuple with the TacRangeList field value
// and a boolean to check if the value has been set.
func (o *TaiRange) GetTacRangeListOk() ([]TacRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.TacRangeList, true
}

// SetTacRangeList sets field value
func (o *TaiRange) SetTacRangeList(v []TacRange) {
	o.TacRangeList = v
}

// GetNid returns the Nid field value if set, zero value otherwise.
func (o *TaiRange) GetNid() string {
	if o == nil || IsNil(o.Nid) {
		var ret string
		return ret
	}
	return *o.Nid
}

// GetNidOk returns a tuple with the Nid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaiRange) GetNidOk() (*string, bool) {
	if o == nil || IsNil(o.Nid) {
		return nil, false
	}
	return o.Nid, true
}

// HasNid returns a boolean if a field has been set.
func (o *TaiRange) HasNid() bool {
	if o != nil && !IsNil(o.Nid) {
		return true
	}

	return false
}

// SetNid gets a reference to the given string and assigns it to the Nid field.
func (o *TaiRange) SetNid(v string) {
	o.Nid = &v
}

func (o TaiRange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaiRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plmnId"] = o.PlmnId
	toSerialize["tacRangeList"] = o.TacRangeList
	if !IsNil(o.Nid) {
		toSerialize["nid"] = o.Nid
	}
	return toSerialize, nil
}

type NullableTaiRange struct {
	value *TaiRange
	isSet bool
}

func (v NullableTaiRange) Get() *TaiRange {
	return v.value
}

func (v *NullableTaiRange) Set(val *TaiRange) {
	v.value = val
	v.isSet = true
}

func (v NullableTaiRange) IsSet() bool {
	return v.isSet
}

func (v *NullableTaiRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaiRange(val *TaiRange) *NullableTaiRange {
	return &NullableTaiRange{value: val, isSet: true}
}

func (v NullableTaiRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaiRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
