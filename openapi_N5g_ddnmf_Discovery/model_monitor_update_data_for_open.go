/*
N5g-ddnmf_Discovery API

N5g-ddnmf_Discovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_N5g_ddnmf_Discovery

import (
	"encoding/json"
)

// checks if the MonitorUpdateDataForOpen type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitorUpdateDataForOpen{}

// MonitorUpdateDataForOpen Represents Monitor Update Data for the Discovery Type \"OPEN\".
type MonitorUpdateDataForOpen struct {
	// Contains the ProSe Application ID name.
	ProseAppIdName string `json:"proseAppIdName"`
	Ttl            int32  `json:"ttl"`
}

// NewMonitorUpdateDataForOpen instantiates a new MonitorUpdateDataForOpen object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorUpdateDataForOpen(proseAppIdName string, ttl int32) *MonitorUpdateDataForOpen {
	this := MonitorUpdateDataForOpen{}
	this.ProseAppIdName = proseAppIdName
	this.Ttl = ttl
	return &this
}

// NewMonitorUpdateDataForOpenWithDefaults instantiates a new MonitorUpdateDataForOpen object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorUpdateDataForOpenWithDefaults() *MonitorUpdateDataForOpen {
	this := MonitorUpdateDataForOpen{}
	return &this
}

// GetProseAppIdName returns the ProseAppIdName field value
func (o *MonitorUpdateDataForOpen) GetProseAppIdName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProseAppIdName
}

// GetProseAppIdNameOk returns a tuple with the ProseAppIdName field value
// and a boolean to check if the value has been set.
func (o *MonitorUpdateDataForOpen) GetProseAppIdNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProseAppIdName, true
}

// SetProseAppIdName sets field value
func (o *MonitorUpdateDataForOpen) SetProseAppIdName(v string) {
	o.ProseAppIdName = v
}

// GetTtl returns the Ttl field value
func (o *MonitorUpdateDataForOpen) GetTtl() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value
// and a boolean to check if the value has been set.
func (o *MonitorUpdateDataForOpen) GetTtlOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ttl, true
}

// SetTtl sets field value
func (o *MonitorUpdateDataForOpen) SetTtl(v int32) {
	o.Ttl = v
}

func (o MonitorUpdateDataForOpen) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitorUpdateDataForOpen) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["proseAppIdName"] = o.ProseAppIdName
	toSerialize["ttl"] = o.Ttl
	return toSerialize, nil
}

type NullableMonitorUpdateDataForOpen struct {
	value *MonitorUpdateDataForOpen
	isSet bool
}

func (v NullableMonitorUpdateDataForOpen) Get() *MonitorUpdateDataForOpen {
	return v.value
}

func (v *NullableMonitorUpdateDataForOpen) Set(val *MonitorUpdateDataForOpen) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorUpdateDataForOpen) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorUpdateDataForOpen) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorUpdateDataForOpen(val *MonitorUpdateDataForOpen) *NullableMonitorUpdateDataForOpen {
	return &NullableMonitorUpdateDataForOpen{value: val, isSet: true}
}

func (v NullableMonitorUpdateDataForOpen) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorUpdateDataForOpen) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
