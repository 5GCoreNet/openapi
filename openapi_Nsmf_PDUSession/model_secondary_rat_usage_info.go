/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
)

// checks if the SecondaryRatUsageInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SecondaryRatUsageInfo{}

// SecondaryRatUsageInfo Secondary RAT Usage Information to report usage data for a secondary RAT for QoS flows and/or the whole PDU session. 
type SecondaryRatUsageInfo struct {
	SecondaryRatType RatType `json:"secondaryRatType"`
	QosFlowsUsageData []QosFlowUsageReport `json:"qosFlowsUsageData,omitempty"`
	PduSessionUsageData []VolumeTimedReport `json:"pduSessionUsageData,omitempty"`
}

// NewSecondaryRatUsageInfo instantiates a new SecondaryRatUsageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecondaryRatUsageInfo(secondaryRatType RatType) *SecondaryRatUsageInfo {
	this := SecondaryRatUsageInfo{}
	this.SecondaryRatType = secondaryRatType
	return &this
}

// NewSecondaryRatUsageInfoWithDefaults instantiates a new SecondaryRatUsageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecondaryRatUsageInfoWithDefaults() *SecondaryRatUsageInfo {
	this := SecondaryRatUsageInfo{}
	return &this
}

// GetSecondaryRatType returns the SecondaryRatType field value
func (o *SecondaryRatUsageInfo) GetSecondaryRatType() RatType {
	if o == nil {
		var ret RatType
		return ret
	}

	return o.SecondaryRatType
}

// GetSecondaryRatTypeOk returns a tuple with the SecondaryRatType field value
// and a boolean to check if the value has been set.
func (o *SecondaryRatUsageInfo) GetSecondaryRatTypeOk() (*RatType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecondaryRatType, true
}

// SetSecondaryRatType sets field value
func (o *SecondaryRatUsageInfo) SetSecondaryRatType(v RatType) {
	o.SecondaryRatType = v
}

// GetQosFlowsUsageData returns the QosFlowsUsageData field value if set, zero value otherwise.
func (o *SecondaryRatUsageInfo) GetQosFlowsUsageData() []QosFlowUsageReport {
	if o == nil || IsNil(o.QosFlowsUsageData) {
		var ret []QosFlowUsageReport
		return ret
	}
	return o.QosFlowsUsageData
}

// GetQosFlowsUsageDataOk returns a tuple with the QosFlowsUsageData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecondaryRatUsageInfo) GetQosFlowsUsageDataOk() ([]QosFlowUsageReport, bool) {
	if o == nil || IsNil(o.QosFlowsUsageData) {
		return nil, false
	}
	return o.QosFlowsUsageData, true
}

// HasQosFlowsUsageData returns a boolean if a field has been set.
func (o *SecondaryRatUsageInfo) HasQosFlowsUsageData() bool {
	if o != nil && !IsNil(o.QosFlowsUsageData) {
		return true
	}

	return false
}

// SetQosFlowsUsageData gets a reference to the given []QosFlowUsageReport and assigns it to the QosFlowsUsageData field.
func (o *SecondaryRatUsageInfo) SetQosFlowsUsageData(v []QosFlowUsageReport) {
	o.QosFlowsUsageData = v
}

// GetPduSessionUsageData returns the PduSessionUsageData field value if set, zero value otherwise.
func (o *SecondaryRatUsageInfo) GetPduSessionUsageData() []VolumeTimedReport {
	if o == nil || IsNil(o.PduSessionUsageData) {
		var ret []VolumeTimedReport
		return ret
	}
	return o.PduSessionUsageData
}

// GetPduSessionUsageDataOk returns a tuple with the PduSessionUsageData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecondaryRatUsageInfo) GetPduSessionUsageDataOk() ([]VolumeTimedReport, bool) {
	if o == nil || IsNil(o.PduSessionUsageData) {
		return nil, false
	}
	return o.PduSessionUsageData, true
}

// HasPduSessionUsageData returns a boolean if a field has been set.
func (o *SecondaryRatUsageInfo) HasPduSessionUsageData() bool {
	if o != nil && !IsNil(o.PduSessionUsageData) {
		return true
	}

	return false
}

// SetPduSessionUsageData gets a reference to the given []VolumeTimedReport and assigns it to the PduSessionUsageData field.
func (o *SecondaryRatUsageInfo) SetPduSessionUsageData(v []VolumeTimedReport) {
	o.PduSessionUsageData = v
}

func (o SecondaryRatUsageInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SecondaryRatUsageInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["secondaryRatType"] = o.SecondaryRatType
	if !IsNil(o.QosFlowsUsageData) {
		toSerialize["qosFlowsUsageData"] = o.QosFlowsUsageData
	}
	if !IsNil(o.PduSessionUsageData) {
		toSerialize["pduSessionUsageData"] = o.PduSessionUsageData
	}
	return toSerialize, nil
}

type NullableSecondaryRatUsageInfo struct {
	value *SecondaryRatUsageInfo
	isSet bool
}

func (v NullableSecondaryRatUsageInfo) Get() *SecondaryRatUsageInfo {
	return v.value
}

func (v *NullableSecondaryRatUsageInfo) Set(val *SecondaryRatUsageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSecondaryRatUsageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSecondaryRatUsageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecondaryRatUsageInfo(val *SecondaryRatUsageInfo) *NullableSecondaryRatUsageInfo {
	return &NullableSecondaryRatUsageInfo{value: val, isSet: true}
}

func (v NullableSecondaryRatUsageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecondaryRatUsageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


