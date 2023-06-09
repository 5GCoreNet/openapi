/*
Npcf_MBSPolicyControl API

MBS Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_MBSPolicyControl

import (
	"encoding/json"
)

// checks if the MbsPolicyCtxtDataUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsPolicyCtxtDataUpdate{}

// MbsPolicyCtxtDataUpdate Contains the parameters to request the modification of an existing MBS Policy Association.
type MbsPolicyCtxtDataUpdate struct {
	MbsServInfo    *MbsServiceInfo `json:"mbsServInfo,omitempty"`
	MbsPcrts       []MbsPcrt       `json:"mbsPcrts,omitempty"`
	MbsErrorReport *MbsErrorReport `json:"mbsErrorReport,omitempty"`
}

// NewMbsPolicyCtxtDataUpdate instantiates a new MbsPolicyCtxtDataUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsPolicyCtxtDataUpdate() *MbsPolicyCtxtDataUpdate {
	this := MbsPolicyCtxtDataUpdate{}
	return &this
}

// NewMbsPolicyCtxtDataUpdateWithDefaults instantiates a new MbsPolicyCtxtDataUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsPolicyCtxtDataUpdateWithDefaults() *MbsPolicyCtxtDataUpdate {
	this := MbsPolicyCtxtDataUpdate{}
	return &this
}

// GetMbsServInfo returns the MbsServInfo field value if set, zero value otherwise.
func (o *MbsPolicyCtxtDataUpdate) GetMbsServInfo() MbsServiceInfo {
	if o == nil || IsNil(o.MbsServInfo) {
		var ret MbsServiceInfo
		return ret
	}
	return *o.MbsServInfo
}

// GetMbsServInfoOk returns a tuple with the MbsServInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsPolicyCtxtDataUpdate) GetMbsServInfoOk() (*MbsServiceInfo, bool) {
	if o == nil || IsNil(o.MbsServInfo) {
		return nil, false
	}
	return o.MbsServInfo, true
}

// HasMbsServInfo returns a boolean if a field has been set.
func (o *MbsPolicyCtxtDataUpdate) HasMbsServInfo() bool {
	if o != nil && !IsNil(o.MbsServInfo) {
		return true
	}

	return false
}

// SetMbsServInfo gets a reference to the given MbsServiceInfo and assigns it to the MbsServInfo field.
func (o *MbsPolicyCtxtDataUpdate) SetMbsServInfo(v MbsServiceInfo) {
	o.MbsServInfo = &v
}

// GetMbsPcrts returns the MbsPcrts field value if set, zero value otherwise.
func (o *MbsPolicyCtxtDataUpdate) GetMbsPcrts() []MbsPcrt {
	if o == nil || IsNil(o.MbsPcrts) {
		var ret []MbsPcrt
		return ret
	}
	return o.MbsPcrts
}

// GetMbsPcrtsOk returns a tuple with the MbsPcrts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsPolicyCtxtDataUpdate) GetMbsPcrtsOk() ([]MbsPcrt, bool) {
	if o == nil || IsNil(o.MbsPcrts) {
		return nil, false
	}
	return o.MbsPcrts, true
}

// HasMbsPcrts returns a boolean if a field has been set.
func (o *MbsPolicyCtxtDataUpdate) HasMbsPcrts() bool {
	if o != nil && !IsNil(o.MbsPcrts) {
		return true
	}

	return false
}

// SetMbsPcrts gets a reference to the given []MbsPcrt and assigns it to the MbsPcrts field.
func (o *MbsPolicyCtxtDataUpdate) SetMbsPcrts(v []MbsPcrt) {
	o.MbsPcrts = v
}

// GetMbsErrorReport returns the MbsErrorReport field value if set, zero value otherwise.
func (o *MbsPolicyCtxtDataUpdate) GetMbsErrorReport() MbsErrorReport {
	if o == nil || IsNil(o.MbsErrorReport) {
		var ret MbsErrorReport
		return ret
	}
	return *o.MbsErrorReport
}

// GetMbsErrorReportOk returns a tuple with the MbsErrorReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsPolicyCtxtDataUpdate) GetMbsErrorReportOk() (*MbsErrorReport, bool) {
	if o == nil || IsNil(o.MbsErrorReport) {
		return nil, false
	}
	return o.MbsErrorReport, true
}

// HasMbsErrorReport returns a boolean if a field has been set.
func (o *MbsPolicyCtxtDataUpdate) HasMbsErrorReport() bool {
	if o != nil && !IsNil(o.MbsErrorReport) {
		return true
	}

	return false
}

// SetMbsErrorReport gets a reference to the given MbsErrorReport and assigns it to the MbsErrorReport field.
func (o *MbsPolicyCtxtDataUpdate) SetMbsErrorReport(v MbsErrorReport) {
	o.MbsErrorReport = &v
}

func (o MbsPolicyCtxtDataUpdate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsPolicyCtxtDataUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MbsServInfo) {
		toSerialize["mbsServInfo"] = o.MbsServInfo
	}
	if !IsNil(o.MbsPcrts) {
		toSerialize["mbsPcrts"] = o.MbsPcrts
	}
	if !IsNil(o.MbsErrorReport) {
		toSerialize["mbsErrorReport"] = o.MbsErrorReport
	}
	return toSerialize, nil
}

type NullableMbsPolicyCtxtDataUpdate struct {
	value *MbsPolicyCtxtDataUpdate
	isSet bool
}

func (v NullableMbsPolicyCtxtDataUpdate) Get() *MbsPolicyCtxtDataUpdate {
	return v.value
}

func (v *NullableMbsPolicyCtxtDataUpdate) Set(val *MbsPolicyCtxtDataUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsPolicyCtxtDataUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsPolicyCtxtDataUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsPolicyCtxtDataUpdate(val *MbsPolicyCtxtDataUpdate) *NullableMbsPolicyCtxtDataUpdate {
	return &NullableMbsPolicyCtxtDataUpdate{value: val, isSet: true}
}

func (v NullableMbsPolicyCtxtDataUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsPolicyCtxtDataUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
