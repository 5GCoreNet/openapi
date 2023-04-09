/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
	"time"
)

// checks if the SACEventReportItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SACEventReportItem{}

// SACEventReportItem Represents a report triggered by a subscribed event type
type SACEventReportItem struct {
	EventType  SACEventType  `json:"eventType"`
	EventState SACEventState `json:"eventState"`
	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp       time.Time       `json:"timeStamp"`
	EventFilter     Snssai          `json:"eventFilter"`
	SliceStautsInfo *SACEventStatus `json:"sliceStautsInfo,omitempty"`
}

// NewSACEventReportItem instantiates a new SACEventReportItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSACEventReportItem(eventType SACEventType, eventState SACEventState, timeStamp time.Time, eventFilter Snssai) *SACEventReportItem {
	this := SACEventReportItem{}
	this.EventType = eventType
	this.EventState = eventState
	this.TimeStamp = timeStamp
	this.EventFilter = eventFilter
	return &this
}

// NewSACEventReportItemWithDefaults instantiates a new SACEventReportItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSACEventReportItemWithDefaults() *SACEventReportItem {
	this := SACEventReportItem{}
	return &this
}

// GetEventType returns the EventType field value
func (o *SACEventReportItem) GetEventType() SACEventType {
	if o == nil {
		var ret SACEventType
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *SACEventReportItem) GetEventTypeOk() (*SACEventType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *SACEventReportItem) SetEventType(v SACEventType) {
	o.EventType = v
}

// GetEventState returns the EventState field value
func (o *SACEventReportItem) GetEventState() SACEventState {
	if o == nil {
		var ret SACEventState
		return ret
	}

	return o.EventState
}

// GetEventStateOk returns a tuple with the EventState field value
// and a boolean to check if the value has been set.
func (o *SACEventReportItem) GetEventStateOk() (*SACEventState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventState, true
}

// SetEventState sets field value
func (o *SACEventReportItem) SetEventState(v SACEventState) {
	o.EventState = v
}

// GetTimeStamp returns the TimeStamp field value
func (o *SACEventReportItem) GetTimeStamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value
// and a boolean to check if the value has been set.
func (o *SACEventReportItem) GetTimeStampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeStamp, true
}

// SetTimeStamp sets field value
func (o *SACEventReportItem) SetTimeStamp(v time.Time) {
	o.TimeStamp = v
}

// GetEventFilter returns the EventFilter field value
func (o *SACEventReportItem) GetEventFilter() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.EventFilter
}

// GetEventFilterOk returns a tuple with the EventFilter field value
// and a boolean to check if the value has been set.
func (o *SACEventReportItem) GetEventFilterOk() (*Snssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventFilter, true
}

// SetEventFilter sets field value
func (o *SACEventReportItem) SetEventFilter(v Snssai) {
	o.EventFilter = v
}

// GetSliceStautsInfo returns the SliceStautsInfo field value if set, zero value otherwise.
func (o *SACEventReportItem) GetSliceStautsInfo() SACEventStatus {
	if o == nil || IsNil(o.SliceStautsInfo) {
		var ret SACEventStatus
		return ret
	}
	return *o.SliceStautsInfo
}

// GetSliceStautsInfoOk returns a tuple with the SliceStautsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SACEventReportItem) GetSliceStautsInfoOk() (*SACEventStatus, bool) {
	if o == nil || IsNil(o.SliceStautsInfo) {
		return nil, false
	}
	return o.SliceStautsInfo, true
}

// HasSliceStautsInfo returns a boolean if a field has been set.
func (o *SACEventReportItem) HasSliceStautsInfo() bool {
	if o != nil && !IsNil(o.SliceStautsInfo) {
		return true
	}

	return false
}

// SetSliceStautsInfo gets a reference to the given SACEventStatus and assigns it to the SliceStautsInfo field.
func (o *SACEventReportItem) SetSliceStautsInfo(v SACEventStatus) {
	o.SliceStautsInfo = &v
}

func (o SACEventReportItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SACEventReportItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventType"] = o.EventType
	toSerialize["eventState"] = o.EventState
	toSerialize["timeStamp"] = o.TimeStamp
	toSerialize["eventFilter"] = o.EventFilter
	if !IsNil(o.SliceStautsInfo) {
		toSerialize["sliceStautsInfo"] = o.SliceStautsInfo
	}
	return toSerialize, nil
}

type NullableSACEventReportItem struct {
	value *SACEventReportItem
	isSet bool
}

func (v NullableSACEventReportItem) Get() *SACEventReportItem {
	return v.value
}

func (v *NullableSACEventReportItem) Set(val *SACEventReportItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSACEventReportItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSACEventReportItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSACEventReportItem(val *SACEventReportItem) *NullableSACEventReportItem {
	return &NullableSACEventReportItem{value: val, isSet: true}
}

func (v NullableSACEventReportItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSACEventReportItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
