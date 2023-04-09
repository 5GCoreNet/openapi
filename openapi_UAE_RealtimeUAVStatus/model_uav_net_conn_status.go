/*
UAE Server Real-time UAV Status Service

UAE Server Real-time UAV Status Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_UAE_RealtimeUAVStatus

import (
	"encoding/json"
	"time"
)

// checks if the UavNetConnStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UavNetConnStatus{}

// UavNetConnStatus Represents UAV network connection status information.
type UavNetConnStatus struct {
	StatusInfo MonitoringType `json:"statusInfo"`
	// string with format \"date-time\" as defined in OpenAPI.
	Timestamp time.Time `json:"timestamp"`
}

// NewUavNetConnStatus instantiates a new UavNetConnStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUavNetConnStatus(statusInfo MonitoringType, timestamp time.Time) *UavNetConnStatus {
	this := UavNetConnStatus{}
	this.StatusInfo = statusInfo
	this.Timestamp = timestamp
	return &this
}

// NewUavNetConnStatusWithDefaults instantiates a new UavNetConnStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUavNetConnStatusWithDefaults() *UavNetConnStatus {
	this := UavNetConnStatus{}
	return &this
}

// GetStatusInfo returns the StatusInfo field value
func (o *UavNetConnStatus) GetStatusInfo() MonitoringType {
	if o == nil {
		var ret MonitoringType
		return ret
	}

	return o.StatusInfo
}

// GetStatusInfoOk returns a tuple with the StatusInfo field value
// and a boolean to check if the value has been set.
func (o *UavNetConnStatus) GetStatusInfoOk() (*MonitoringType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusInfo, true
}

// SetStatusInfo sets field value
func (o *UavNetConnStatus) SetStatusInfo(v MonitoringType) {
	o.StatusInfo = v
}

// GetTimestamp returns the Timestamp field value
func (o *UavNetConnStatus) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *UavNetConnStatus) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *UavNetConnStatus) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

func (o UavNetConnStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UavNetConnStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["statusInfo"] = o.StatusInfo
	toSerialize["timestamp"] = o.Timestamp
	return toSerialize, nil
}

type NullableUavNetConnStatus struct {
	value *UavNetConnStatus
	isSet bool
}

func (v NullableUavNetConnStatus) Get() *UavNetConnStatus {
	return v.value
}

func (v *NullableUavNetConnStatus) Set(val *UavNetConnStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableUavNetConnStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableUavNetConnStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUavNetConnStatus(val *UavNetConnStatus) *NullableUavNetConnStatus {
	return &NullableUavNetConnStatus{value: val, isSet: true}
}

func (v NullableUavNetConnStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUavNetConnStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
