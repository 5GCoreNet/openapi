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

// checks if the ContextStatusEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContextStatusEvent{}

// ContextStatusEvent Context Status Event
type ContextStatusEvent struct {
	EventType          ContextStatusEventType `json:"eventType"`
	ImmediateReportInd *bool                  `json:"immediateReportInd,omitempty"`
	ReportingMode      *ReportingMode         `json:"reportingMode,omitempty"`
}

// NewContextStatusEvent instantiates a new ContextStatusEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextStatusEvent(eventType ContextStatusEventType) *ContextStatusEvent {
	this := ContextStatusEvent{}
	this.EventType = eventType
	var immediateReportInd bool = false
	this.ImmediateReportInd = &immediateReportInd
	return &this
}

// NewContextStatusEventWithDefaults instantiates a new ContextStatusEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextStatusEventWithDefaults() *ContextStatusEvent {
	this := ContextStatusEvent{}
	var immediateReportInd bool = false
	this.ImmediateReportInd = &immediateReportInd
	return &this
}

// GetEventType returns the EventType field value
func (o *ContextStatusEvent) GetEventType() ContextStatusEventType {
	if o == nil {
		var ret ContextStatusEventType
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *ContextStatusEvent) GetEventTypeOk() (*ContextStatusEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *ContextStatusEvent) SetEventType(v ContextStatusEventType) {
	o.EventType = v
}

// GetImmediateReportInd returns the ImmediateReportInd field value if set, zero value otherwise.
func (o *ContextStatusEvent) GetImmediateReportInd() bool {
	if o == nil || IsNil(o.ImmediateReportInd) {
		var ret bool
		return ret
	}
	return *o.ImmediateReportInd
}

// GetImmediateReportIndOk returns a tuple with the ImmediateReportInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextStatusEvent) GetImmediateReportIndOk() (*bool, bool) {
	if o == nil || IsNil(o.ImmediateReportInd) {
		return nil, false
	}
	return o.ImmediateReportInd, true
}

// HasImmediateReportInd returns a boolean if a field has been set.
func (o *ContextStatusEvent) HasImmediateReportInd() bool {
	if o != nil && !IsNil(o.ImmediateReportInd) {
		return true
	}

	return false
}

// SetImmediateReportInd gets a reference to the given bool and assigns it to the ImmediateReportInd field.
func (o *ContextStatusEvent) SetImmediateReportInd(v bool) {
	o.ImmediateReportInd = &v
}

// GetReportingMode returns the ReportingMode field value if set, zero value otherwise.
func (o *ContextStatusEvent) GetReportingMode() ReportingMode {
	if o == nil || IsNil(o.ReportingMode) {
		var ret ReportingMode
		return ret
	}
	return *o.ReportingMode
}

// GetReportingModeOk returns a tuple with the ReportingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextStatusEvent) GetReportingModeOk() (*ReportingMode, bool) {
	if o == nil || IsNil(o.ReportingMode) {
		return nil, false
	}
	return o.ReportingMode, true
}

// HasReportingMode returns a boolean if a field has been set.
func (o *ContextStatusEvent) HasReportingMode() bool {
	if o != nil && !IsNil(o.ReportingMode) {
		return true
	}

	return false
}

// SetReportingMode gets a reference to the given ReportingMode and assigns it to the ReportingMode field.
func (o *ContextStatusEvent) SetReportingMode(v ReportingMode) {
	o.ReportingMode = &v
}

func (o ContextStatusEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContextStatusEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventType"] = o.EventType
	if !IsNil(o.ImmediateReportInd) {
		toSerialize["immediateReportInd"] = o.ImmediateReportInd
	}
	if !IsNil(o.ReportingMode) {
		toSerialize["reportingMode"] = o.ReportingMode
	}
	return toSerialize, nil
}

type NullableContextStatusEvent struct {
	value *ContextStatusEvent
	isSet bool
}

func (v NullableContextStatusEvent) Get() *ContextStatusEvent {
	return v.value
}

func (v *NullableContextStatusEvent) Set(val *ContextStatusEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableContextStatusEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableContextStatusEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextStatusEvent(val *ContextStatusEvent) *NullableContextStatusEvent {
	return &NullableContextStatusEvent{value: val, isSet: true}
}

func (v NullableContextStatusEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextStatusEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
