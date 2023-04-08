/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.5.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CommonData

import (
	"encoding/json"
)

// checks if the OdbData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OdbData{}

// OdbData Contains information regarding operater  determined  barring.
type OdbData struct {
	RoamingOdb *RoamingOdb `json:"roamingOdb,omitempty"`
}

// NewOdbData instantiates a new OdbData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOdbData() *OdbData {
	this := OdbData{}
	return &this
}

// NewOdbDataWithDefaults instantiates a new OdbData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOdbDataWithDefaults() *OdbData {
	this := OdbData{}
	return &this
}

// GetRoamingOdb returns the RoamingOdb field value if set, zero value otherwise.
func (o *OdbData) GetRoamingOdb() RoamingOdb {
	if o == nil || isNil(o.RoamingOdb) {
		var ret RoamingOdb
		return ret
	}
	return *o.RoamingOdb
}

// GetRoamingOdbOk returns a tuple with the RoamingOdb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OdbData) GetRoamingOdbOk() (*RoamingOdb, bool) {
	if o == nil || isNil(o.RoamingOdb) {
		return nil, false
	}
	return o.RoamingOdb, true
}

// HasRoamingOdb returns a boolean if a field has been set.
func (o *OdbData) HasRoamingOdb() bool {
	if o != nil && !isNil(o.RoamingOdb) {
		return true
	}

	return false
}

// SetRoamingOdb gets a reference to the given RoamingOdb and assigns it to the RoamingOdb field.
func (o *OdbData) SetRoamingOdb(v RoamingOdb) {
	o.RoamingOdb = &v
}

func (o OdbData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OdbData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RoamingOdb) {
		toSerialize["roamingOdb"] = o.RoamingOdb
	}
	return toSerialize, nil
}

type NullableOdbData struct {
	value *OdbData
	isSet bool
}

func (v NullableOdbData) Get() *OdbData {
	return v.value
}

func (v *NullableOdbData) Set(val *OdbData) {
	v.value = val
	v.isSet = true
}

func (v NullableOdbData) IsSet() bool {
	return v.isSet
}

func (v *NullableOdbData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOdbData(val *OdbData) *NullableOdbData {
	return &NullableOdbData{value: val, isSet: true}
}

func (v NullableOdbData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOdbData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


