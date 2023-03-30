/*
SS_Events

API for SEAL Events management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_Events

import (
	"encoding/json"
)

// checks if the MonitorLocationInterestFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitorLocationInterestFilter{}

// MonitorLocationInterestFilter Represents the location monitoring filter information.
type MonitorLocationInterestFilter struct {
	// List of VAL Users or UE IDs for which location monitoring is requested.
	TgtUes []ValTargetUe `json:"tgtUes"`
	LocInt LocationInfo `json:"locInt"`
	// indicating a time in seconds.
	NotInt int32 `json:"notInt"`
}

// NewMonitorLocationInterestFilter instantiates a new MonitorLocationInterestFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorLocationInterestFilter(tgtUes []ValTargetUe, locInt LocationInfo, notInt int32) *MonitorLocationInterestFilter {
	this := MonitorLocationInterestFilter{}
	this.TgtUes = tgtUes
	this.LocInt = locInt
	this.NotInt = notInt
	return &this
}

// NewMonitorLocationInterestFilterWithDefaults instantiates a new MonitorLocationInterestFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorLocationInterestFilterWithDefaults() *MonitorLocationInterestFilter {
	this := MonitorLocationInterestFilter{}
	return &this
}

// GetTgtUes returns the TgtUes field value
func (o *MonitorLocationInterestFilter) GetTgtUes() []ValTargetUe {
	if o == nil {
		var ret []ValTargetUe
		return ret
	}

	return o.TgtUes
}

// GetTgtUesOk returns a tuple with the TgtUes field value
// and a boolean to check if the value has been set.
func (o *MonitorLocationInterestFilter) GetTgtUesOk() ([]ValTargetUe, bool) {
	if o == nil {
		return nil, false
	}
	return o.TgtUes, true
}

// SetTgtUes sets field value
func (o *MonitorLocationInterestFilter) SetTgtUes(v []ValTargetUe) {
	o.TgtUes = v
}

// GetLocInt returns the LocInt field value
func (o *MonitorLocationInterestFilter) GetLocInt() LocationInfo {
	if o == nil {
		var ret LocationInfo
		return ret
	}

	return o.LocInt
}

// GetLocIntOk returns a tuple with the LocInt field value
// and a boolean to check if the value has been set.
func (o *MonitorLocationInterestFilter) GetLocIntOk() (*LocationInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocInt, true
}

// SetLocInt sets field value
func (o *MonitorLocationInterestFilter) SetLocInt(v LocationInfo) {
	o.LocInt = v
}

// GetNotInt returns the NotInt field value
func (o *MonitorLocationInterestFilter) GetNotInt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NotInt
}

// GetNotIntOk returns a tuple with the NotInt field value
// and a boolean to check if the value has been set.
func (o *MonitorLocationInterestFilter) GetNotIntOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotInt, true
}

// SetNotInt sets field value
func (o *MonitorLocationInterestFilter) SetNotInt(v int32) {
	o.NotInt = v
}

func (o MonitorLocationInterestFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitorLocationInterestFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tgtUes"] = o.TgtUes
	toSerialize["locInt"] = o.LocInt
	toSerialize["notInt"] = o.NotInt
	return toSerialize, nil
}

type NullableMonitorLocationInterestFilter struct {
	value *MonitorLocationInterestFilter
	isSet bool
}

func (v NullableMonitorLocationInterestFilter) Get() *MonitorLocationInterestFilter {
	return v.value
}

func (v *NullableMonitorLocationInterestFilter) Set(val *MonitorLocationInterestFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorLocationInterestFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorLocationInterestFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorLocationInterestFilter(val *MonitorLocationInterestFilter) *NullableMonitorLocationInterestFilter {
	return &NullableMonitorLocationInterestFilter{value: val, isSet: true}
}

func (v NullableMonitorLocationInterestFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorLocationInterestFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


