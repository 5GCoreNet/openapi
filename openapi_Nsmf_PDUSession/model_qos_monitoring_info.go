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

// checks if the QosMonitoringInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QosMonitoringInfo{}

// QosMonitoringInfo QoS Monitoring Information
type QosMonitoringInfo struct {
	QosMonitoringInd *bool `json:"qosMonitoringInd,omitempty"`
}

// NewQosMonitoringInfo instantiates a new QosMonitoringInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosMonitoringInfo() *QosMonitoringInfo {
	this := QosMonitoringInfo{}
	var qosMonitoringInd bool = false
	this.QosMonitoringInd = &qosMonitoringInd
	return &this
}

// NewQosMonitoringInfoWithDefaults instantiates a new QosMonitoringInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosMonitoringInfoWithDefaults() *QosMonitoringInfo {
	this := QosMonitoringInfo{}
	var qosMonitoringInd bool = false
	this.QosMonitoringInd = &qosMonitoringInd
	return &this
}

// GetQosMonitoringInd returns the QosMonitoringInd field value if set, zero value otherwise.
func (o *QosMonitoringInfo) GetQosMonitoringInd() bool {
	if o == nil || IsNil(o.QosMonitoringInd) {
		var ret bool
		return ret
	}
	return *o.QosMonitoringInd
}

// GetQosMonitoringIndOk returns a tuple with the QosMonitoringInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosMonitoringInfo) GetQosMonitoringIndOk() (*bool, bool) {
	if o == nil || IsNil(o.QosMonitoringInd) {
		return nil, false
	}
	return o.QosMonitoringInd, true
}

// HasQosMonitoringInd returns a boolean if a field has been set.
func (o *QosMonitoringInfo) HasQosMonitoringInd() bool {
	if o != nil && !IsNil(o.QosMonitoringInd) {
		return true
	}

	return false
}

// SetQosMonitoringInd gets a reference to the given bool and assigns it to the QosMonitoringInd field.
func (o *QosMonitoringInfo) SetQosMonitoringInd(v bool) {
	o.QosMonitoringInd = &v
}

func (o QosMonitoringInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QosMonitoringInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QosMonitoringInd) {
		toSerialize["qosMonitoringInd"] = o.QosMonitoringInd
	}
	return toSerialize, nil
}

type NullableQosMonitoringInfo struct {
	value *QosMonitoringInfo
	isSet bool
}

func (v NullableQosMonitoringInfo) Get() *QosMonitoringInfo {
	return v.value
}

func (v *NullableQosMonitoringInfo) Set(val *QosMonitoringInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableQosMonitoringInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableQosMonitoringInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosMonitoringInfo(val *QosMonitoringInfo) *NullableQosMonitoringInfo {
	return &NullableQosMonitoringInfo{value: val, isSet: true}
}

func (v NullableQosMonitoringInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosMonitoringInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
