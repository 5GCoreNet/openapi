/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
)

// checks if the SACEventReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SACEventReport{}

// SACEventReport Event notification
type SACEventReport struct {
	Report SACEventReportItem `json:"report"`
	NotifyCorrelationId *string `json:"notifyCorrelationId,omitempty"`
}

// NewSACEventReport instantiates a new SACEventReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSACEventReport(report SACEventReportItem) *SACEventReport {
	this := SACEventReport{}
	this.Report = report
	return &this
}

// NewSACEventReportWithDefaults instantiates a new SACEventReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSACEventReportWithDefaults() *SACEventReport {
	this := SACEventReport{}
	return &this
}

// GetReport returns the Report field value
func (o *SACEventReport) GetReport() SACEventReportItem {
	if o == nil {
		var ret SACEventReportItem
		return ret
	}

	return o.Report
}

// GetReportOk returns a tuple with the Report field value
// and a boolean to check if the value has been set.
func (o *SACEventReport) GetReportOk() (*SACEventReportItem, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Report, true
}

// SetReport sets field value
func (o *SACEventReport) SetReport(v SACEventReportItem) {
	o.Report = v
}

// GetNotifyCorrelationId returns the NotifyCorrelationId field value if set, zero value otherwise.
func (o *SACEventReport) GetNotifyCorrelationId() string {
	if o == nil || IsNil(o.NotifyCorrelationId) {
		var ret string
		return ret
	}
	return *o.NotifyCorrelationId
}

// GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SACEventReport) GetNotifyCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyCorrelationId) {
		return nil, false
	}
	return o.NotifyCorrelationId, true
}

// HasNotifyCorrelationId returns a boolean if a field has been set.
func (o *SACEventReport) HasNotifyCorrelationId() bool {
	if o != nil && !IsNil(o.NotifyCorrelationId) {
		return true
	}

	return false
}

// SetNotifyCorrelationId gets a reference to the given string and assigns it to the NotifyCorrelationId field.
func (o *SACEventReport) SetNotifyCorrelationId(v string) {
	o.NotifyCorrelationId = &v
}

func (o SACEventReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SACEventReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["report"] = o.Report
	if !IsNil(o.NotifyCorrelationId) {
		toSerialize["notifyCorrelationId"] = o.NotifyCorrelationId
	}
	return toSerialize, nil
}

type NullableSACEventReport struct {
	value *SACEventReport
	isSet bool
}

func (v NullableSACEventReport) Get() *SACEventReport {
	return v.value
}

func (v *NullableSACEventReport) Set(val *SACEventReport) {
	v.value = val
	v.isSet = true
}

func (v NullableSACEventReport) IsSet() bool {
	return v.isSet
}

func (v *NullableSACEventReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSACEventReport(val *SACEventReport) *NullableSACEventReport {
	return &NullableSACEventReport{value: val, isSet: true}
}

func (v NullableSACEventReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSACEventReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


