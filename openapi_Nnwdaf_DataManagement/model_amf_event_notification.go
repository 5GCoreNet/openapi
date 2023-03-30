/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
)

// checks if the AmfEventNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmfEventNotification{}

// AmfEventNotification Data within a AMF Event Notification request
type AmfEventNotification struct {
	NotifyCorrelationId *string `json:"notifyCorrelationId,omitempty"`
	SubsChangeNotifyCorrelationId *string `json:"subsChangeNotifyCorrelationId,omitempty"`
	ReportList []AmfEventReport `json:"reportList,omitempty"`
	EventSubsSyncInfo *AmfEventSubsSyncInfo `json:"eventSubsSyncInfo,omitempty"`
}

// NewAmfEventNotification instantiates a new AmfEventNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfEventNotification() *AmfEventNotification {
	this := AmfEventNotification{}
	return &this
}

// NewAmfEventNotificationWithDefaults instantiates a new AmfEventNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfEventNotificationWithDefaults() *AmfEventNotification {
	this := AmfEventNotification{}
	return &this
}

// GetNotifyCorrelationId returns the NotifyCorrelationId field value if set, zero value otherwise.
func (o *AmfEventNotification) GetNotifyCorrelationId() string {
	if o == nil || IsNil(o.NotifyCorrelationId) {
		var ret string
		return ret
	}
	return *o.NotifyCorrelationId
}

// GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventNotification) GetNotifyCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.NotifyCorrelationId) {
		return nil, false
	}
	return o.NotifyCorrelationId, true
}

// HasNotifyCorrelationId returns a boolean if a field has been set.
func (o *AmfEventNotification) HasNotifyCorrelationId() bool {
	if o != nil && !IsNil(o.NotifyCorrelationId) {
		return true
	}

	return false
}

// SetNotifyCorrelationId gets a reference to the given string and assigns it to the NotifyCorrelationId field.
func (o *AmfEventNotification) SetNotifyCorrelationId(v string) {
	o.NotifyCorrelationId = &v
}

// GetSubsChangeNotifyCorrelationId returns the SubsChangeNotifyCorrelationId field value if set, zero value otherwise.
func (o *AmfEventNotification) GetSubsChangeNotifyCorrelationId() string {
	if o == nil || IsNil(o.SubsChangeNotifyCorrelationId) {
		var ret string
		return ret
	}
	return *o.SubsChangeNotifyCorrelationId
}

// GetSubsChangeNotifyCorrelationIdOk returns a tuple with the SubsChangeNotifyCorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventNotification) GetSubsChangeNotifyCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubsChangeNotifyCorrelationId) {
		return nil, false
	}
	return o.SubsChangeNotifyCorrelationId, true
}

// HasSubsChangeNotifyCorrelationId returns a boolean if a field has been set.
func (o *AmfEventNotification) HasSubsChangeNotifyCorrelationId() bool {
	if o != nil && !IsNil(o.SubsChangeNotifyCorrelationId) {
		return true
	}

	return false
}

// SetSubsChangeNotifyCorrelationId gets a reference to the given string and assigns it to the SubsChangeNotifyCorrelationId field.
func (o *AmfEventNotification) SetSubsChangeNotifyCorrelationId(v string) {
	o.SubsChangeNotifyCorrelationId = &v
}

// GetReportList returns the ReportList field value if set, zero value otherwise.
func (o *AmfEventNotification) GetReportList() []AmfEventReport {
	if o == nil || IsNil(o.ReportList) {
		var ret []AmfEventReport
		return ret
	}
	return o.ReportList
}

// GetReportListOk returns a tuple with the ReportList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventNotification) GetReportListOk() ([]AmfEventReport, bool) {
	if o == nil || IsNil(o.ReportList) {
		return nil, false
	}
	return o.ReportList, true
}

// HasReportList returns a boolean if a field has been set.
func (o *AmfEventNotification) HasReportList() bool {
	if o != nil && !IsNil(o.ReportList) {
		return true
	}

	return false
}

// SetReportList gets a reference to the given []AmfEventReport and assigns it to the ReportList field.
func (o *AmfEventNotification) SetReportList(v []AmfEventReport) {
	o.ReportList = v
}

// GetEventSubsSyncInfo returns the EventSubsSyncInfo field value if set, zero value otherwise.
func (o *AmfEventNotification) GetEventSubsSyncInfo() AmfEventSubsSyncInfo {
	if o == nil || IsNil(o.EventSubsSyncInfo) {
		var ret AmfEventSubsSyncInfo
		return ret
	}
	return *o.EventSubsSyncInfo
}

// GetEventSubsSyncInfoOk returns a tuple with the EventSubsSyncInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventNotification) GetEventSubsSyncInfoOk() (*AmfEventSubsSyncInfo, bool) {
	if o == nil || IsNil(o.EventSubsSyncInfo) {
		return nil, false
	}
	return o.EventSubsSyncInfo, true
}

// HasEventSubsSyncInfo returns a boolean if a field has been set.
func (o *AmfEventNotification) HasEventSubsSyncInfo() bool {
	if o != nil && !IsNil(o.EventSubsSyncInfo) {
		return true
	}

	return false
}

// SetEventSubsSyncInfo gets a reference to the given AmfEventSubsSyncInfo and assigns it to the EventSubsSyncInfo field.
func (o *AmfEventNotification) SetEventSubsSyncInfo(v AmfEventSubsSyncInfo) {
	o.EventSubsSyncInfo = &v
}

func (o AmfEventNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmfEventNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NotifyCorrelationId) {
		toSerialize["notifyCorrelationId"] = o.NotifyCorrelationId
	}
	if !IsNil(o.SubsChangeNotifyCorrelationId) {
		toSerialize["subsChangeNotifyCorrelationId"] = o.SubsChangeNotifyCorrelationId
	}
	if !IsNil(o.ReportList) {
		toSerialize["reportList"] = o.ReportList
	}
	if !IsNil(o.EventSubsSyncInfo) {
		toSerialize["eventSubsSyncInfo"] = o.EventSubsSyncInfo
	}
	return toSerialize, nil
}

type NullableAmfEventNotification struct {
	value *AmfEventNotification
	isSet bool
}

func (v NullableAmfEventNotification) Get() *AmfEventNotification {
	return v.value
}

func (v *NullableAmfEventNotification) Set(val *AmfEventNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfEventNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfEventNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfEventNotification(val *AmfEventNotification) *NullableAmfEventNotification {
	return &NullableAmfEventNotification{value: val, isSet: true}
}

func (v NullableAmfEventNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfEventNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


