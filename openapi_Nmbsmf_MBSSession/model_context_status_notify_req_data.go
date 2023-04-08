/*
Nmbsmf-MBSSession

MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsmf_MBSSession

import (
	"encoding/json"
)

// checks if the ContextStatusNotifyReqData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContextStatusNotifyReqData{}

// ContextStatusNotifyReqData Context Status Notification
type ContextStatusNotifyReqData struct {
	ReportList []ContextStatusEventReport `json:"reportList"`
	NotifyCorrelationId *string `json:"notifyCorrelationId,omitempty"`
}

// NewContextStatusNotifyReqData instantiates a new ContextStatusNotifyReqData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextStatusNotifyReqData(reportList []ContextStatusEventReport) *ContextStatusNotifyReqData {
	this := ContextStatusNotifyReqData{}
	this.ReportList = reportList
	return &this
}

// NewContextStatusNotifyReqDataWithDefaults instantiates a new ContextStatusNotifyReqData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextStatusNotifyReqDataWithDefaults() *ContextStatusNotifyReqData {
	this := ContextStatusNotifyReqData{}
	return &this
}

// GetReportList returns the ReportList field value
func (o *ContextStatusNotifyReqData) GetReportList() []ContextStatusEventReport {
	if o == nil {
		var ret []ContextStatusEventReport
		return ret
	}

	return o.ReportList
}

// GetReportListOk returns a tuple with the ReportList field value
// and a boolean to check if the value has been set.
func (o *ContextStatusNotifyReqData) GetReportListOk() ([]ContextStatusEventReport, bool) {
	if o == nil {
		return nil, false
	}
	return o.ReportList, true
}

// SetReportList sets field value
func (o *ContextStatusNotifyReqData) SetReportList(v []ContextStatusEventReport) {
	o.ReportList = v
}

// GetNotifyCorrelationId returns the NotifyCorrelationId field value if set, zero value otherwise.
func (o *ContextStatusNotifyReqData) GetNotifyCorrelationId() string {
	if o == nil || isNil(o.NotifyCorrelationId) {
		var ret string
		return ret
	}
	return *o.NotifyCorrelationId
}

// GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextStatusNotifyReqData) GetNotifyCorrelationIdOk() (*string, bool) {
	if o == nil || isNil(o.NotifyCorrelationId) {
		return nil, false
	}
	return o.NotifyCorrelationId, true
}

// HasNotifyCorrelationId returns a boolean if a field has been set.
func (o *ContextStatusNotifyReqData) HasNotifyCorrelationId() bool {
	if o != nil && !isNil(o.NotifyCorrelationId) {
		return true
	}

	return false
}

// SetNotifyCorrelationId gets a reference to the given string and assigns it to the NotifyCorrelationId field.
func (o *ContextStatusNotifyReqData) SetNotifyCorrelationId(v string) {
	o.NotifyCorrelationId = &v
}

func (o ContextStatusNotifyReqData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContextStatusNotifyReqData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reportList"] = o.ReportList
	if !isNil(o.NotifyCorrelationId) {
		toSerialize["notifyCorrelationId"] = o.NotifyCorrelationId
	}
	return toSerialize, nil
}

type NullableContextStatusNotifyReqData struct {
	value *ContextStatusNotifyReqData
	isSet bool
}

func (v NullableContextStatusNotifyReqData) Get() *ContextStatusNotifyReqData {
	return v.value
}

func (v *NullableContextStatusNotifyReqData) Set(val *ContextStatusNotifyReqData) {
	v.value = val
	v.isSet = true
}

func (v NullableContextStatusNotifyReqData) IsSet() bool {
	return v.isSet
}

func (v *NullableContextStatusNotifyReqData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextStatusNotifyReqData(val *ContextStatusNotifyReqData) *NullableContextStatusNotifyReqData {
	return &NullableContextStatusNotifyReqData{value: val, isSet: true}
}

func (v NullableContextStatusNotifyReqData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextStatusNotifyReqData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


