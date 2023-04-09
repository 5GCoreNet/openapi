/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
)

// checks if the NotifSummaryReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotifSummaryReport{}

// NotifSummaryReport Represents summarized notifications based on processing instructions.
type NotifSummaryReport struct {
	EventId DccfEvent `json:"eventId"`
	// indicating a time in seconds.
	ProcInterval int32 `json:"procInterval"`
	// List of event parameter reports.
	EventReports []EventParamReport `json:"eventReports"`
}

// NewNotifSummaryReport instantiates a new NotifSummaryReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifSummaryReport(eventId DccfEvent, procInterval int32, eventReports []EventParamReport) *NotifSummaryReport {
	this := NotifSummaryReport{}
	this.EventId = eventId
	this.ProcInterval = procInterval
	this.EventReports = eventReports
	return &this
}

// NewNotifSummaryReportWithDefaults instantiates a new NotifSummaryReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifSummaryReportWithDefaults() *NotifSummaryReport {
	this := NotifSummaryReport{}
	return &this
}

// GetEventId returns the EventId field value
func (o *NotifSummaryReport) GetEventId() DccfEvent {
	if o == nil {
		var ret DccfEvent
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *NotifSummaryReport) GetEventIdOk() (*DccfEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *NotifSummaryReport) SetEventId(v DccfEvent) {
	o.EventId = v
}

// GetProcInterval returns the ProcInterval field value
func (o *NotifSummaryReport) GetProcInterval() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProcInterval
}

// GetProcIntervalOk returns a tuple with the ProcInterval field value
// and a boolean to check if the value has been set.
func (o *NotifSummaryReport) GetProcIntervalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcInterval, true
}

// SetProcInterval sets field value
func (o *NotifSummaryReport) SetProcInterval(v int32) {
	o.ProcInterval = v
}

// GetEventReports returns the EventReports field value
func (o *NotifSummaryReport) GetEventReports() []EventParamReport {
	if o == nil {
		var ret []EventParamReport
		return ret
	}

	return o.EventReports
}

// GetEventReportsOk returns a tuple with the EventReports field value
// and a boolean to check if the value has been set.
func (o *NotifSummaryReport) GetEventReportsOk() ([]EventParamReport, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventReports, true
}

// SetEventReports sets field value
func (o *NotifSummaryReport) SetEventReports(v []EventParamReport) {
	o.EventReports = v
}

func (o NotifSummaryReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotifSummaryReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventId"] = o.EventId
	toSerialize["procInterval"] = o.ProcInterval
	toSerialize["eventReports"] = o.EventReports
	return toSerialize, nil
}

type NullableNotifSummaryReport struct {
	value *NotifSummaryReport
	isSet bool
}

func (v NullableNotifSummaryReport) Get() *NotifSummaryReport {
	return v.value
}

func (v *NullableNotifSummaryReport) Set(val *NotifSummaryReport) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifSummaryReport) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifSummaryReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifSummaryReport(val *NotifSummaryReport) *NullableNotifSummaryReport {
	return &NullableNotifSummaryReport{value: val, isSet: true}
}

func (v NullableNotifSummaryReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifSummaryReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
