/*
Nudm_EE

Nudm Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_EE

import (
	"encoding/json"
	"time"
)

// checks if the MonitoringReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitoringReport{}

// MonitoringReport struct for MonitoringReport
type MonitoringReport struct {
	ReferenceId int32 `json:"referenceId"`
	EventType EventType `json:"eventType"`
	Report *Report `json:"report,omitempty"`
	ReachabilityForSmsReport *ReachabilityForSmsReport `json:"reachabilityForSmsReport,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi *string `json:"gpsi,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp"`
	ReachabilityReport *ReachabilityReport `json:"reachabilityReport,omitempty"`
}

// NewMonitoringReport instantiates a new MonitoringReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringReport(referenceId int32, eventType EventType, timeStamp time.Time) *MonitoringReport {
	this := MonitoringReport{}
	this.ReferenceId = referenceId
	this.EventType = eventType
	this.TimeStamp = timeStamp
	return &this
}

// NewMonitoringReportWithDefaults instantiates a new MonitoringReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringReportWithDefaults() *MonitoringReport {
	this := MonitoringReport{}
	return &this
}

// GetReferenceId returns the ReferenceId field value
func (o *MonitoringReport) GetReferenceId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value
// and a boolean to check if the value has been set.
func (o *MonitoringReport) GetReferenceIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceId, true
}

// SetReferenceId sets field value
func (o *MonitoringReport) SetReferenceId(v int32) {
	o.ReferenceId = v
}

// GetEventType returns the EventType field value
func (o *MonitoringReport) GetEventType() EventType {
	if o == nil {
		var ret EventType
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *MonitoringReport) GetEventTypeOk() (*EventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *MonitoringReport) SetEventType(v EventType) {
	o.EventType = v
}

// GetReport returns the Report field value if set, zero value otherwise.
func (o *MonitoringReport) GetReport() Report {
	if o == nil || isNil(o.Report) {
		var ret Report
		return ret
	}
	return *o.Report
}

// GetReportOk returns a tuple with the Report field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringReport) GetReportOk() (*Report, bool) {
	if o == nil || isNil(o.Report) {
		return nil, false
	}
	return o.Report, true
}

// HasReport returns a boolean if a field has been set.
func (o *MonitoringReport) HasReport() bool {
	if o != nil && !isNil(o.Report) {
		return true
	}

	return false
}

// SetReport gets a reference to the given Report and assigns it to the Report field.
func (o *MonitoringReport) SetReport(v Report) {
	o.Report = &v
}

// GetReachabilityForSmsReport returns the ReachabilityForSmsReport field value if set, zero value otherwise.
func (o *MonitoringReport) GetReachabilityForSmsReport() ReachabilityForSmsReport {
	if o == nil || isNil(o.ReachabilityForSmsReport) {
		var ret ReachabilityForSmsReport
		return ret
	}
	return *o.ReachabilityForSmsReport
}

// GetReachabilityForSmsReportOk returns a tuple with the ReachabilityForSmsReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringReport) GetReachabilityForSmsReportOk() (*ReachabilityForSmsReport, bool) {
	if o == nil || isNil(o.ReachabilityForSmsReport) {
		return nil, false
	}
	return o.ReachabilityForSmsReport, true
}

// HasReachabilityForSmsReport returns a boolean if a field has been set.
func (o *MonitoringReport) HasReachabilityForSmsReport() bool {
	if o != nil && !isNil(o.ReachabilityForSmsReport) {
		return true
	}

	return false
}

// SetReachabilityForSmsReport gets a reference to the given ReachabilityForSmsReport and assigns it to the ReachabilityForSmsReport field.
func (o *MonitoringReport) SetReachabilityForSmsReport(v ReachabilityForSmsReport) {
	o.ReachabilityForSmsReport = &v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *MonitoringReport) GetGpsi() string {
	if o == nil || isNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringReport) GetGpsiOk() (*string, bool) {
	if o == nil || isNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *MonitoringReport) HasGpsi() bool {
	if o != nil && !isNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *MonitoringReport) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetTimeStamp returns the TimeStamp field value
func (o *MonitoringReport) GetTimeStamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value
// and a boolean to check if the value has been set.
func (o *MonitoringReport) GetTimeStampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeStamp, true
}

// SetTimeStamp sets field value
func (o *MonitoringReport) SetTimeStamp(v time.Time) {
	o.TimeStamp = v
}

// GetReachabilityReport returns the ReachabilityReport field value if set, zero value otherwise.
func (o *MonitoringReport) GetReachabilityReport() ReachabilityReport {
	if o == nil || isNil(o.ReachabilityReport) {
		var ret ReachabilityReport
		return ret
	}
	return *o.ReachabilityReport
}

// GetReachabilityReportOk returns a tuple with the ReachabilityReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringReport) GetReachabilityReportOk() (*ReachabilityReport, bool) {
	if o == nil || isNil(o.ReachabilityReport) {
		return nil, false
	}
	return o.ReachabilityReport, true
}

// HasReachabilityReport returns a boolean if a field has been set.
func (o *MonitoringReport) HasReachabilityReport() bool {
	if o != nil && !isNil(o.ReachabilityReport) {
		return true
	}

	return false
}

// SetReachabilityReport gets a reference to the given ReachabilityReport and assigns it to the ReachabilityReport field.
func (o *MonitoringReport) SetReachabilityReport(v ReachabilityReport) {
	o.ReachabilityReport = &v
}

func (o MonitoringReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitoringReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["referenceId"] = o.ReferenceId
	toSerialize["eventType"] = o.EventType
	if !isNil(o.Report) {
		toSerialize["report"] = o.Report
	}
	if !isNil(o.ReachabilityForSmsReport) {
		toSerialize["reachabilityForSmsReport"] = o.ReachabilityForSmsReport
	}
	if !isNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	toSerialize["timeStamp"] = o.TimeStamp
	if !isNil(o.ReachabilityReport) {
		toSerialize["reachabilityReport"] = o.ReachabilityReport
	}
	return toSerialize, nil
}

type NullableMonitoringReport struct {
	value *MonitoringReport
	isSet bool
}

func (v NullableMonitoringReport) Get() *MonitoringReport {
	return v.value
}

func (v *NullableMonitoringReport) Set(val *MonitoringReport) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringReport) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringReport(val *MonitoringReport) *NullableMonitoringReport {
	return &NullableMonitoringReport{value: val, isSet: true}
}

func (v NullableMonitoringReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


