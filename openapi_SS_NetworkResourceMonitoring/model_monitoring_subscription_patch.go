/*
SS_NetworkResourceMonitoring

API for SEAL Network Resource Monitoring.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_NetworkResourceMonitoring

import (
	"encoding/json"
)

// checks if the MonitoringSubscriptionPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitoringSubscriptionPatch{}

// MonitoringSubscriptionPatch Represents the monitoring subscription modification request.
type MonitoringSubscriptionPatch struct {
	MeasReqs   *MeasurementRequirements `json:"measReqs,omitempty"`
	ReportReqs *ReportingRequirements   `json:"reportReqs,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	NotifUri *string `json:"notifUri,omitempty"`
}

// NewMonitoringSubscriptionPatch instantiates a new MonitoringSubscriptionPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitoringSubscriptionPatch() *MonitoringSubscriptionPatch {
	this := MonitoringSubscriptionPatch{}
	return &this
}

// NewMonitoringSubscriptionPatchWithDefaults instantiates a new MonitoringSubscriptionPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitoringSubscriptionPatchWithDefaults() *MonitoringSubscriptionPatch {
	this := MonitoringSubscriptionPatch{}
	return &this
}

// GetMeasReqs returns the MeasReqs field value if set, zero value otherwise.
func (o *MonitoringSubscriptionPatch) GetMeasReqs() MeasurementRequirements {
	if o == nil || IsNil(o.MeasReqs) {
		var ret MeasurementRequirements
		return ret
	}
	return *o.MeasReqs
}

// GetMeasReqsOk returns a tuple with the MeasReqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSubscriptionPatch) GetMeasReqsOk() (*MeasurementRequirements, bool) {
	if o == nil || IsNil(o.MeasReqs) {
		return nil, false
	}
	return o.MeasReqs, true
}

// HasMeasReqs returns a boolean if a field has been set.
func (o *MonitoringSubscriptionPatch) HasMeasReqs() bool {
	if o != nil && !IsNil(o.MeasReqs) {
		return true
	}

	return false
}

// SetMeasReqs gets a reference to the given MeasurementRequirements and assigns it to the MeasReqs field.
func (o *MonitoringSubscriptionPatch) SetMeasReqs(v MeasurementRequirements) {
	o.MeasReqs = &v
}

// GetReportReqs returns the ReportReqs field value if set, zero value otherwise.
func (o *MonitoringSubscriptionPatch) GetReportReqs() ReportingRequirements {
	if o == nil || IsNil(o.ReportReqs) {
		var ret ReportingRequirements
		return ret
	}
	return *o.ReportReqs
}

// GetReportReqsOk returns a tuple with the ReportReqs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSubscriptionPatch) GetReportReqsOk() (*ReportingRequirements, bool) {
	if o == nil || IsNil(o.ReportReqs) {
		return nil, false
	}
	return o.ReportReqs, true
}

// HasReportReqs returns a boolean if a field has been set.
func (o *MonitoringSubscriptionPatch) HasReportReqs() bool {
	if o != nil && !IsNil(o.ReportReqs) {
		return true
	}

	return false
}

// SetReportReqs gets a reference to the given ReportingRequirements and assigns it to the ReportReqs field.
func (o *MonitoringSubscriptionPatch) SetReportReqs(v ReportingRequirements) {
	o.ReportReqs = &v
}

// GetNotifUri returns the NotifUri field value if set, zero value otherwise.
func (o *MonitoringSubscriptionPatch) GetNotifUri() string {
	if o == nil || IsNil(o.NotifUri) {
		var ret string
		return ret
	}
	return *o.NotifUri
}

// GetNotifUriOk returns a tuple with the NotifUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitoringSubscriptionPatch) GetNotifUriOk() (*string, bool) {
	if o == nil || IsNil(o.NotifUri) {
		return nil, false
	}
	return o.NotifUri, true
}

// HasNotifUri returns a boolean if a field has been set.
func (o *MonitoringSubscriptionPatch) HasNotifUri() bool {
	if o != nil && !IsNil(o.NotifUri) {
		return true
	}

	return false
}

// SetNotifUri gets a reference to the given string and assigns it to the NotifUri field.
func (o *MonitoringSubscriptionPatch) SetNotifUri(v string) {
	o.NotifUri = &v
}

func (o MonitoringSubscriptionPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitoringSubscriptionPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MeasReqs) {
		toSerialize["measReqs"] = o.MeasReqs
	}
	if !IsNil(o.ReportReqs) {
		toSerialize["reportReqs"] = o.ReportReqs
	}
	if !IsNil(o.NotifUri) {
		toSerialize["notifUri"] = o.NotifUri
	}
	return toSerialize, nil
}

type NullableMonitoringSubscriptionPatch struct {
	value *MonitoringSubscriptionPatch
	isSet bool
}

func (v NullableMonitoringSubscriptionPatch) Get() *MonitoringSubscriptionPatch {
	return v.value
}

func (v *NullableMonitoringSubscriptionPatch) Set(val *MonitoringSubscriptionPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitoringSubscriptionPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitoringSubscriptionPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitoringSubscriptionPatch(val *MonitoringSubscriptionPatch) *NullableMonitoringSubscriptionPatch {
	return &NullableMonitoringSubscriptionPatch{value: val, isSet: true}
}

func (v NullableMonitoringSubscriptionPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitoringSubscriptionPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
