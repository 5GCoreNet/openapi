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

// checks if the MonitorUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitorUpdateData{}

// MonitorUpdateData Represents Monitor Authorize Data to update.
type MonitorUpdateData struct {
	DiscType             DiscoveryType                   `json:"discType"`
	OpenUpdateData       *MonitorUpdateDataForOpen       `json:"openUpdateData,omitempty"`
	RestrictedUpdateData *MonitorUpdateDataForRestricted `json:"restrictedUpdateData,omitempty"`
}

// NewMonitorUpdateData instantiates a new MonitorUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorUpdateData(discType DiscoveryType) *MonitorUpdateData {
	this := MonitorUpdateData{}
	this.DiscType = discType
	return &this
}

// NewMonitorUpdateDataWithDefaults instantiates a new MonitorUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorUpdateDataWithDefaults() *MonitorUpdateData {
	this := MonitorUpdateData{}
	return &this
}

// GetDiscType returns the DiscType field value
func (o *MonitorUpdateData) GetDiscType() DiscoveryType {
	if o == nil {
		var ret DiscoveryType
		return ret
	}

	return o.DiscType
}

// GetDiscTypeOk returns a tuple with the DiscType field value
// and a boolean to check if the value has been set.
func (o *MonitorUpdateData) GetDiscTypeOk() (*DiscoveryType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DiscType, true
}

// SetDiscType sets field value
func (o *MonitorUpdateData) SetDiscType(v DiscoveryType) {
	o.DiscType = v
}

// GetOpenUpdateData returns the OpenUpdateData field value if set, zero value otherwise.
func (o *MonitorUpdateData) GetOpenUpdateData() MonitorUpdateDataForOpen {
	if o == nil || IsNil(o.OpenUpdateData) {
		var ret MonitorUpdateDataForOpen
		return ret
	}
	return *o.OpenUpdateData
}

// GetOpenUpdateDataOk returns a tuple with the OpenUpdateData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorUpdateData) GetOpenUpdateDataOk() (*MonitorUpdateDataForOpen, bool) {
	if o == nil || IsNil(o.OpenUpdateData) {
		return nil, false
	}
	return o.OpenUpdateData, true
}

// HasOpenUpdateData returns a boolean if a field has been set.
func (o *MonitorUpdateData) HasOpenUpdateData() bool {
	if o != nil && !IsNil(o.OpenUpdateData) {
		return true
	}

	return false
}

// SetOpenUpdateData gets a reference to the given MonitorUpdateDataForOpen and assigns it to the OpenUpdateData field.
func (o *MonitorUpdateData) SetOpenUpdateData(v MonitorUpdateDataForOpen) {
	o.OpenUpdateData = &v
}

// GetRestrictedUpdateData returns the RestrictedUpdateData field value if set, zero value otherwise.
func (o *MonitorUpdateData) GetRestrictedUpdateData() MonitorUpdateDataForRestricted {
	if o == nil || IsNil(o.RestrictedUpdateData) {
		var ret MonitorUpdateDataForRestricted
		return ret
	}
	return *o.RestrictedUpdateData
}

// GetRestrictedUpdateDataOk returns a tuple with the RestrictedUpdateData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorUpdateData) GetRestrictedUpdateDataOk() (*MonitorUpdateDataForRestricted, bool) {
	if o == nil || IsNil(o.RestrictedUpdateData) {
		return nil, false
	}
	return o.RestrictedUpdateData, true
}

// HasRestrictedUpdateData returns a boolean if a field has been set.
func (o *MonitorUpdateData) HasRestrictedUpdateData() bool {
	if o != nil && !IsNil(o.RestrictedUpdateData) {
		return true
	}

	return false
}

// SetRestrictedUpdateData gets a reference to the given MonitorUpdateDataForRestricted and assigns it to the RestrictedUpdateData field.
func (o *MonitorUpdateData) SetRestrictedUpdateData(v MonitorUpdateDataForRestricted) {
	o.RestrictedUpdateData = &v
}

func (o MonitorUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitorUpdateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["discType"] = o.DiscType
	if !IsNil(o.OpenUpdateData) {
		toSerialize["openUpdateData"] = o.OpenUpdateData
	}
	if !IsNil(o.RestrictedUpdateData) {
		toSerialize["restrictedUpdateData"] = o.RestrictedUpdateData
	}
	return toSerialize, nil
}

type NullableMonitorUpdateData struct {
	value *MonitorUpdateData
	isSet bool
}

func (v NullableMonitorUpdateData) Get() *MonitorUpdateData {
	return v.value
}

func (v *NullableMonitorUpdateData) Set(val *MonitorUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorUpdateData(val *MonitorUpdateData) *NullableMonitorUpdateData {
	return &NullableMonitorUpdateData{value: val, isSet: true}
}

func (v NullableMonitorUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
