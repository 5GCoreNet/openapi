/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
	"time"
)

// checks if the DispersionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DispersionInfo{}

// DispersionInfo Represents the Dispersion information. When subscribed event is \"DISPERSION\", the  \"disperInfos\" attribute shall be included. 
type DispersionInfo struct {
	// string with format 'date-time' as defined in OpenAPI.
	TsStart time.Time `json:"tsStart"`
	// indicating a time in seconds.
	TsDuration int32 `json:"tsDuration"`
	DisperCollects []DispersionCollection `json:"disperCollects"`
	DisperType DispersionType `json:"disperType"`
}

// NewDispersionInfo instantiates a new DispersionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDispersionInfo(tsStart time.Time, tsDuration int32, disperCollects []DispersionCollection, disperType DispersionType) *DispersionInfo {
	this := DispersionInfo{}
	this.TsStart = tsStart
	this.TsDuration = tsDuration
	this.DisperCollects = disperCollects
	this.DisperType = disperType
	return &this
}

// NewDispersionInfoWithDefaults instantiates a new DispersionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDispersionInfoWithDefaults() *DispersionInfo {
	this := DispersionInfo{}
	return &this
}

// GetTsStart returns the TsStart field value
func (o *DispersionInfo) GetTsStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TsStart
}

// GetTsStartOk returns a tuple with the TsStart field value
// and a boolean to check if the value has been set.
func (o *DispersionInfo) GetTsStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TsStart, true
}

// SetTsStart sets field value
func (o *DispersionInfo) SetTsStart(v time.Time) {
	o.TsStart = v
}

// GetTsDuration returns the TsDuration field value
func (o *DispersionInfo) GetTsDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TsDuration
}

// GetTsDurationOk returns a tuple with the TsDuration field value
// and a boolean to check if the value has been set.
func (o *DispersionInfo) GetTsDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TsDuration, true
}

// SetTsDuration sets field value
func (o *DispersionInfo) SetTsDuration(v int32) {
	o.TsDuration = v
}

// GetDisperCollects returns the DisperCollects field value
func (o *DispersionInfo) GetDisperCollects() []DispersionCollection {
	if o == nil {
		var ret []DispersionCollection
		return ret
	}

	return o.DisperCollects
}

// GetDisperCollectsOk returns a tuple with the DisperCollects field value
// and a boolean to check if the value has been set.
func (o *DispersionInfo) GetDisperCollectsOk() ([]DispersionCollection, bool) {
	if o == nil {
		return nil, false
	}
	return o.DisperCollects, true
}

// SetDisperCollects sets field value
func (o *DispersionInfo) SetDisperCollects(v []DispersionCollection) {
	o.DisperCollects = v
}

// GetDisperType returns the DisperType field value
func (o *DispersionInfo) GetDisperType() DispersionType {
	if o == nil {
		var ret DispersionType
		return ret
	}

	return o.DisperType
}

// GetDisperTypeOk returns a tuple with the DisperType field value
// and a boolean to check if the value has been set.
func (o *DispersionInfo) GetDisperTypeOk() (*DispersionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisperType, true
}

// SetDisperType sets field value
func (o *DispersionInfo) SetDisperType(v DispersionType) {
	o.DisperType = v
}

func (o DispersionInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DispersionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tsStart"] = o.TsStart
	toSerialize["tsDuration"] = o.TsDuration
	toSerialize["disperCollects"] = o.DisperCollects
	toSerialize["disperType"] = o.DisperType
	return toSerialize, nil
}

type NullableDispersionInfo struct {
	value *DispersionInfo
	isSet bool
}

func (v NullableDispersionInfo) Get() *DispersionInfo {
	return v.value
}

func (v *NullableDispersionInfo) Set(val *DispersionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDispersionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDispersionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDispersionInfo(val *DispersionInfo) *NullableDispersionInfo {
	return &NullableDispersionInfo{value: val, isSet: true}
}

func (v NullableDispersionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDispersionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


