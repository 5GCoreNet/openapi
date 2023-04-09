/*
Nmbsmf_TMGI

MB-SMF TMGI Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsmf_TMGI

import (
	"encoding/json"
	"time"
)

// checks if the TmgiAllocated type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TmgiAllocated{}

// TmgiAllocated Data within TMGI Allocate Response
type TmgiAllocated struct {
	// One or more TMGI values
	TmgiList []Tmgi `json:"tmgiList"`
	// string with format 'date-time' as defined in OpenAPI.
	ExpirationTime time.Time `json:"expirationTime"`
	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).
	Nid *string `json:"nid,omitempty"`
}

// NewTmgiAllocated instantiates a new TmgiAllocated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTmgiAllocated(tmgiList []Tmgi, expirationTime time.Time) *TmgiAllocated {
	this := TmgiAllocated{}
	this.TmgiList = tmgiList
	this.ExpirationTime = expirationTime
	return &this
}

// NewTmgiAllocatedWithDefaults instantiates a new TmgiAllocated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTmgiAllocatedWithDefaults() *TmgiAllocated {
	this := TmgiAllocated{}
	return &this
}

// GetTmgiList returns the TmgiList field value
func (o *TmgiAllocated) GetTmgiList() []Tmgi {
	if o == nil {
		var ret []Tmgi
		return ret
	}

	return o.TmgiList
}

// GetTmgiListOk returns a tuple with the TmgiList field value
// and a boolean to check if the value has been set.
func (o *TmgiAllocated) GetTmgiListOk() ([]Tmgi, bool) {
	if o == nil {
		return nil, false
	}
	return o.TmgiList, true
}

// SetTmgiList sets field value
func (o *TmgiAllocated) SetTmgiList(v []Tmgi) {
	o.TmgiList = v
}

// GetExpirationTime returns the ExpirationTime field value
func (o *TmgiAllocated) GetExpirationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ExpirationTime
}

// GetExpirationTimeOk returns a tuple with the ExpirationTime field value
// and a boolean to check if the value has been set.
func (o *TmgiAllocated) GetExpirationTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationTime, true
}

// SetExpirationTime sets field value
func (o *TmgiAllocated) SetExpirationTime(v time.Time) {
	o.ExpirationTime = v
}

// GetNid returns the Nid field value if set, zero value otherwise.
func (o *TmgiAllocated) GetNid() string {
	if o == nil || IsNil(o.Nid) {
		var ret string
		return ret
	}
	return *o.Nid
}

// GetNidOk returns a tuple with the Nid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TmgiAllocated) GetNidOk() (*string, bool) {
	if o == nil || IsNil(o.Nid) {
		return nil, false
	}
	return o.Nid, true
}

// HasNid returns a boolean if a field has been set.
func (o *TmgiAllocated) HasNid() bool {
	if o != nil && !IsNil(o.Nid) {
		return true
	}

	return false
}

// SetNid gets a reference to the given string and assigns it to the Nid field.
func (o *TmgiAllocated) SetNid(v string) {
	o.Nid = &v
}

func (o TmgiAllocated) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TmgiAllocated) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tmgiList"] = o.TmgiList
	toSerialize["expirationTime"] = o.ExpirationTime
	if !IsNil(o.Nid) {
		toSerialize["nid"] = o.Nid
	}
	return toSerialize, nil
}

type NullableTmgiAllocated struct {
	value *TmgiAllocated
	isSet bool
}

func (v NullableTmgiAllocated) Get() *TmgiAllocated {
	return v.value
}

func (v *NullableTmgiAllocated) Set(val *TmgiAllocated) {
	v.value = val
	v.isSet = true
}

func (v NullableTmgiAllocated) IsSet() bool {
	return v.isSet
}

func (v *NullableTmgiAllocated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTmgiAllocated(val *TmgiAllocated) *NullableTmgiAllocated {
	return &NullableTmgiAllocated{value: val, isSet: true}
}

func (v NullableTmgiAllocated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTmgiAllocated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
