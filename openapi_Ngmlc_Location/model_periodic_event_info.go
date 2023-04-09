/*
Ngmlc_Location

GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ngmlc_Location

import (
	"encoding/json"
)

// checks if the PeriodicEventInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PeriodicEventInfo{}

// PeriodicEventInfo Indicates the information of periodic event reporting.
type PeriodicEventInfo struct {
	// Number of required periodic event reports.
	ReportingAmount int32 `json:"reportingAmount"`
	// Event reporting periodic interval.
	ReportingInterval int32 `json:"reportingInterval"`
}

// NewPeriodicEventInfo instantiates a new PeriodicEventInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeriodicEventInfo(reportingAmount int32, reportingInterval int32) *PeriodicEventInfo {
	this := PeriodicEventInfo{}
	this.ReportingAmount = reportingAmount
	this.ReportingInterval = reportingInterval
	return &this
}

// NewPeriodicEventInfoWithDefaults instantiates a new PeriodicEventInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeriodicEventInfoWithDefaults() *PeriodicEventInfo {
	this := PeriodicEventInfo{}
	return &this
}

// GetReportingAmount returns the ReportingAmount field value
func (o *PeriodicEventInfo) GetReportingAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReportingAmount
}

// GetReportingAmountOk returns a tuple with the ReportingAmount field value
// and a boolean to check if the value has been set.
func (o *PeriodicEventInfo) GetReportingAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportingAmount, true
}

// SetReportingAmount sets field value
func (o *PeriodicEventInfo) SetReportingAmount(v int32) {
	o.ReportingAmount = v
}

// GetReportingInterval returns the ReportingInterval field value
func (o *PeriodicEventInfo) GetReportingInterval() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReportingInterval
}

// GetReportingIntervalOk returns a tuple with the ReportingInterval field value
// and a boolean to check if the value has been set.
func (o *PeriodicEventInfo) GetReportingIntervalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportingInterval, true
}

// SetReportingInterval sets field value
func (o *PeriodicEventInfo) SetReportingInterval(v int32) {
	o.ReportingInterval = v
}

func (o PeriodicEventInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PeriodicEventInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reportingAmount"] = o.ReportingAmount
	toSerialize["reportingInterval"] = o.ReportingInterval
	return toSerialize, nil
}

type NullablePeriodicEventInfo struct {
	value *PeriodicEventInfo
	isSet bool
}

func (v NullablePeriodicEventInfo) Get() *PeriodicEventInfo {
	return v.value
}

func (v *NullablePeriodicEventInfo) Set(val *PeriodicEventInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePeriodicEventInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePeriodicEventInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeriodicEventInfo(val *PeriodicEventInfo) *NullablePeriodicEventInfo {
	return &NullablePeriodicEventInfo{value: val, isSet: true}
}

func (v NullablePeriodicEventInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeriodicEventInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
