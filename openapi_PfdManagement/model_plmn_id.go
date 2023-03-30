/*
3gpp-pfd-management

API for PFD management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_PfdManagement

import (
	"encoding/json"
)

// checks if the PlmnId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlmnId{}

// PlmnId When PlmnId needs to be converted to string (e.g. when used in maps as key), the string  shall be composed of three digits \"mcc\" followed by \"-\" and two or three digits \"mnc\". 
type PlmnId struct {
	// Mobile Country Code part of the PLMN, comprising 3 digits, as defined in clause 9.3.3.5 of 3GPP TS 38.413.  
	Mcc string `json:"mcc"`
	// Mobile Network Code part of the PLMN, comprising 2 or 3 digits, as defined in clause 9.3.3.5 of 3GPP TS 38.413.
	Mnc string `json:"mnc"`
}

// NewPlmnId instantiates a new PlmnId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlmnId(mcc string, mnc string) *PlmnId {
	this := PlmnId{}
	this.Mcc = mcc
	this.Mnc = mnc
	return &this
}

// NewPlmnIdWithDefaults instantiates a new PlmnId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlmnIdWithDefaults() *PlmnId {
	this := PlmnId{}
	return &this
}

// GetMcc returns the Mcc field value
func (o *PlmnId) GetMcc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value
// and a boolean to check if the value has been set.
func (o *PlmnId) GetMccOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mcc, true
}

// SetMcc sets field value
func (o *PlmnId) SetMcc(v string) {
	o.Mcc = v
}

// GetMnc returns the Mnc field value
func (o *PlmnId) GetMnc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mnc
}

// GetMncOk returns a tuple with the Mnc field value
// and a boolean to check if the value has been set.
func (o *PlmnId) GetMncOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mnc, true
}

// SetMnc sets field value
func (o *PlmnId) SetMnc(v string) {
	o.Mnc = v
}

func (o PlmnId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlmnId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mcc"] = o.Mcc
	toSerialize["mnc"] = o.Mnc
	return toSerialize, nil
}

type NullablePlmnId struct {
	value *PlmnId
	isSet bool
}

func (v NullablePlmnId) Get() *PlmnId {
	return v.value
}

func (v *NullablePlmnId) Set(val *PlmnId) {
	v.value = val
	v.isSet = true
}

func (v NullablePlmnId) IsSet() bool {
	return v.isSet
}

func (v *NullablePlmnId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlmnId(val *PlmnId) *NullablePlmnId {
	return &NullablePlmnId{value: val, isSet: true}
}

func (v NullablePlmnId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlmnId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


