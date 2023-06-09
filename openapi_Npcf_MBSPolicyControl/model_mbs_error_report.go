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

// checks if the MbsErrorReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsErrorReport{}

// MbsErrorReport Represents the reporting of MBS Policy decision level failure(s) and/or MBS PCC rule level failure(s).
type MbsErrorReport struct {
	MbsReports []MbsReport `json:"mbsReports,omitempty"`
}

// NewMbsErrorReport instantiates a new MbsErrorReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsErrorReport() *MbsErrorReport {
	this := MbsErrorReport{}
	return &this
}

// NewMbsErrorReportWithDefaults instantiates a new MbsErrorReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsErrorReportWithDefaults() *MbsErrorReport {
	this := MbsErrorReport{}
	return &this
}

// GetMbsReports returns the MbsReports field value if set, zero value otherwise.
func (o *MbsErrorReport) GetMbsReports() []MbsReport {
	if o == nil || IsNil(o.MbsReports) {
		var ret []MbsReport
		return ret
	}
	return o.MbsReports
}

// GetMbsReportsOk returns a tuple with the MbsReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsErrorReport) GetMbsReportsOk() ([]MbsReport, bool) {
	if o == nil || IsNil(o.MbsReports) {
		return nil, false
	}
	return o.MbsReports, true
}

// HasMbsReports returns a boolean if a field has been set.
func (o *MbsErrorReport) HasMbsReports() bool {
	if o != nil && !IsNil(o.MbsReports) {
		return true
	}

	return false
}

// SetMbsReports gets a reference to the given []MbsReport and assigns it to the MbsReports field.
func (o *MbsErrorReport) SetMbsReports(v []MbsReport) {
	o.MbsReports = v
}

func (o MbsErrorReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsErrorReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MbsReports) {
		toSerialize["mbsReports"] = o.MbsReports
	}
	return toSerialize, nil
}

type NullableMbsErrorReport struct {
	value *MbsErrorReport
	isSet bool
}

func (v NullableMbsErrorReport) Get() *MbsErrorReport {
	return v.value
}

func (v *NullableMbsErrorReport) Set(val *MbsErrorReport) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsErrorReport) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsErrorReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsErrorReport(val *MbsErrorReport) *NullableMbsErrorReport {
	return &NullableMbsErrorReport{value: val, isSet: true}
}

func (v NullableMbsErrorReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsErrorReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
