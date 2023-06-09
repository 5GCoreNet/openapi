/*
SS_Events

API for SEAL Events management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_Events

import (
	"encoding/json"
)

// checks if the LocationDevMonReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationDevMonReport{}

// LocationDevMonReport Location deviation monitoring report.
type LocationDevMonReport struct {
	// List of VAL Users or UE IDs for which report is related to.
	TgtUes    []ValTargetUe      `json:"tgtUes"`
	LocInfo   LocationInfo       `json:"locInfo"`
	NotifType LocDevNotification `json:"notifType"`
}

// NewLocationDevMonReport instantiates a new LocationDevMonReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationDevMonReport(tgtUes []ValTargetUe, locInfo LocationInfo, notifType LocDevNotification) *LocationDevMonReport {
	this := LocationDevMonReport{}
	this.TgtUes = tgtUes
	this.LocInfo = locInfo
	this.NotifType = notifType
	return &this
}

// NewLocationDevMonReportWithDefaults instantiates a new LocationDevMonReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationDevMonReportWithDefaults() *LocationDevMonReport {
	this := LocationDevMonReport{}
	return &this
}

// GetTgtUes returns the TgtUes field value
func (o *LocationDevMonReport) GetTgtUes() []ValTargetUe {
	if o == nil {
		var ret []ValTargetUe
		return ret
	}

	return o.TgtUes
}

// GetTgtUesOk returns a tuple with the TgtUes field value
// and a boolean to check if the value has been set.
func (o *LocationDevMonReport) GetTgtUesOk() ([]ValTargetUe, bool) {
	if o == nil {
		return nil, false
	}
	return o.TgtUes, true
}

// SetTgtUes sets field value
func (o *LocationDevMonReport) SetTgtUes(v []ValTargetUe) {
	o.TgtUes = v
}

// GetLocInfo returns the LocInfo field value
func (o *LocationDevMonReport) GetLocInfo() LocationInfo {
	if o == nil {
		var ret LocationInfo
		return ret
	}

	return o.LocInfo
}

// GetLocInfoOk returns a tuple with the LocInfo field value
// and a boolean to check if the value has been set.
func (o *LocationDevMonReport) GetLocInfoOk() (*LocationInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocInfo, true
}

// SetLocInfo sets field value
func (o *LocationDevMonReport) SetLocInfo(v LocationInfo) {
	o.LocInfo = v
}

// GetNotifType returns the NotifType field value
func (o *LocationDevMonReport) GetNotifType() LocDevNotification {
	if o == nil {
		var ret LocDevNotification
		return ret
	}

	return o.NotifType
}

// GetNotifTypeOk returns a tuple with the NotifType field value
// and a boolean to check if the value has been set.
func (o *LocationDevMonReport) GetNotifTypeOk() (*LocDevNotification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifType, true
}

// SetNotifType sets field value
func (o *LocationDevMonReport) SetNotifType(v LocDevNotification) {
	o.NotifType = v
}

func (o LocationDevMonReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationDevMonReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tgtUes"] = o.TgtUes
	toSerialize["locInfo"] = o.LocInfo
	toSerialize["notifType"] = o.NotifType
	return toSerialize, nil
}

type NullableLocationDevMonReport struct {
	value *LocationDevMonReport
	isSet bool
}

func (v NullableLocationDevMonReport) Get() *LocationDevMonReport {
	return v.value
}

func (v *NullableLocationDevMonReport) Set(val *LocationDevMonReport) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationDevMonReport) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationDevMonReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationDevMonReport(val *LocationDevMonReport) *NullableLocationDevMonReport {
	return &NullableLocationDevMonReport{value: val, isSet: true}
}

func (v NullableLocationDevMonReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationDevMonReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
