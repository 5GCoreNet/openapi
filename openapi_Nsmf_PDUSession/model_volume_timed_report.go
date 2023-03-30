/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
	"time"
)

// checks if the VolumeTimedReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeTimedReport{}

// VolumeTimedReport Contains Usage data information.
type VolumeTimedReport struct {
	// string with format 'date-time' as defined in OpenAPI.
	StartTimeStamp time.Time `json:"startTimeStamp"`
	// string with format 'date-time' as defined in OpenAPI.
	EndTimeStamp time.Time `json:"endTimeStamp"`
	// string with format 'int64' as defined in OpenAPI.
	DownlinkVolume int64 `json:"downlinkVolume"`
	// string with format 'int64' as defined in OpenAPI.
	UplinkVolume int64 `json:"uplinkVolume"`
}

// NewVolumeTimedReport instantiates a new VolumeTimedReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeTimedReport(startTimeStamp time.Time, endTimeStamp time.Time, downlinkVolume int64, uplinkVolume int64) *VolumeTimedReport {
	this := VolumeTimedReport{}
	this.StartTimeStamp = startTimeStamp
	this.EndTimeStamp = endTimeStamp
	this.DownlinkVolume = downlinkVolume
	this.UplinkVolume = uplinkVolume
	return &this
}

// NewVolumeTimedReportWithDefaults instantiates a new VolumeTimedReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeTimedReportWithDefaults() *VolumeTimedReport {
	this := VolumeTimedReport{}
	return &this
}

// GetStartTimeStamp returns the StartTimeStamp field value
func (o *VolumeTimedReport) GetStartTimeStamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTimeStamp
}

// GetStartTimeStampOk returns a tuple with the StartTimeStamp field value
// and a boolean to check if the value has been set.
func (o *VolumeTimedReport) GetStartTimeStampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTimeStamp, true
}

// SetStartTimeStamp sets field value
func (o *VolumeTimedReport) SetStartTimeStamp(v time.Time) {
	o.StartTimeStamp = v
}

// GetEndTimeStamp returns the EndTimeStamp field value
func (o *VolumeTimedReport) GetEndTimeStamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTimeStamp
}

// GetEndTimeStampOk returns a tuple with the EndTimeStamp field value
// and a boolean to check if the value has been set.
func (o *VolumeTimedReport) GetEndTimeStampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTimeStamp, true
}

// SetEndTimeStamp sets field value
func (o *VolumeTimedReport) SetEndTimeStamp(v time.Time) {
	o.EndTimeStamp = v
}

// GetDownlinkVolume returns the DownlinkVolume field value
func (o *VolumeTimedReport) GetDownlinkVolume() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DownlinkVolume
}

// GetDownlinkVolumeOk returns a tuple with the DownlinkVolume field value
// and a boolean to check if the value has been set.
func (o *VolumeTimedReport) GetDownlinkVolumeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DownlinkVolume, true
}

// SetDownlinkVolume sets field value
func (o *VolumeTimedReport) SetDownlinkVolume(v int64) {
	o.DownlinkVolume = v
}

// GetUplinkVolume returns the UplinkVolume field value
func (o *VolumeTimedReport) GetUplinkVolume() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UplinkVolume
}

// GetUplinkVolumeOk returns a tuple with the UplinkVolume field value
// and a boolean to check if the value has been set.
func (o *VolumeTimedReport) GetUplinkVolumeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UplinkVolume, true
}

// SetUplinkVolume sets field value
func (o *VolumeTimedReport) SetUplinkVolume(v int64) {
	o.UplinkVolume = v
}

func (o VolumeTimedReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeTimedReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["startTimeStamp"] = o.StartTimeStamp
	toSerialize["endTimeStamp"] = o.EndTimeStamp
	toSerialize["downlinkVolume"] = o.DownlinkVolume
	toSerialize["uplinkVolume"] = o.UplinkVolume
	return toSerialize, nil
}

type NullableVolumeTimedReport struct {
	value *VolumeTimedReport
	isSet bool
}

func (v NullableVolumeTimedReport) Get() *VolumeTimedReport {
	return v.value
}

func (v *NullableVolumeTimedReport) Set(val *VolumeTimedReport) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeTimedReport) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeTimedReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeTimedReport(val *VolumeTimedReport) *NullableVolumeTimedReport {
	return &NullableVolumeTimedReport{value: val, isSet: true}
}

func (v NullableVolumeTimedReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeTimedReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


