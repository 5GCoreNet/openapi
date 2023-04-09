/*
MDA Report

OAS 3.0.1 specification of the MDA Report © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MdaReport

import (
	"encoding/json"
	"time"
)

// checks if the TimeWindow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeWindow{}

// TimeWindow struct for TimeWindow
type TimeWindow struct {
	MDAOutputStartTime *time.Time `json:"mDAOutputStartTime,omitempty"`
	MDAOutputEndTime   *time.Time `json:"mDAOutputEndTime,omitempty"`
}

// NewTimeWindow instantiates a new TimeWindow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeWindow() *TimeWindow {
	this := TimeWindow{}
	return &this
}

// NewTimeWindowWithDefaults instantiates a new TimeWindow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeWindowWithDefaults() *TimeWindow {
	this := TimeWindow{}
	return &this
}

// GetMDAOutputStartTime returns the MDAOutputStartTime field value if set, zero value otherwise.
func (o *TimeWindow) GetMDAOutputStartTime() time.Time {
	if o == nil || IsNil(o.MDAOutputStartTime) {
		var ret time.Time
		return ret
	}
	return *o.MDAOutputStartTime
}

// GetMDAOutputStartTimeOk returns a tuple with the MDAOutputStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeWindow) GetMDAOutputStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.MDAOutputStartTime) {
		return nil, false
	}
	return o.MDAOutputStartTime, true
}

// HasMDAOutputStartTime returns a boolean if a field has been set.
func (o *TimeWindow) HasMDAOutputStartTime() bool {
	if o != nil && !IsNil(o.MDAOutputStartTime) {
		return true
	}

	return false
}

// SetMDAOutputStartTime gets a reference to the given time.Time and assigns it to the MDAOutputStartTime field.
func (o *TimeWindow) SetMDAOutputStartTime(v time.Time) {
	o.MDAOutputStartTime = &v
}

// GetMDAOutputEndTime returns the MDAOutputEndTime field value if set, zero value otherwise.
func (o *TimeWindow) GetMDAOutputEndTime() time.Time {
	if o == nil || IsNil(o.MDAOutputEndTime) {
		var ret time.Time
		return ret
	}
	return *o.MDAOutputEndTime
}

// GetMDAOutputEndTimeOk returns a tuple with the MDAOutputEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeWindow) GetMDAOutputEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.MDAOutputEndTime) {
		return nil, false
	}
	return o.MDAOutputEndTime, true
}

// HasMDAOutputEndTime returns a boolean if a field has been set.
func (o *TimeWindow) HasMDAOutputEndTime() bool {
	if o != nil && !IsNil(o.MDAOutputEndTime) {
		return true
	}

	return false
}

// SetMDAOutputEndTime gets a reference to the given time.Time and assigns it to the MDAOutputEndTime field.
func (o *TimeWindow) SetMDAOutputEndTime(v time.Time) {
	o.MDAOutputEndTime = &v
}

func (o TimeWindow) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeWindow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MDAOutputStartTime) {
		toSerialize["mDAOutputStartTime"] = o.MDAOutputStartTime
	}
	if !IsNil(o.MDAOutputEndTime) {
		toSerialize["mDAOutputEndTime"] = o.MDAOutputEndTime
	}
	return toSerialize, nil
}

type NullableTimeWindow struct {
	value *TimeWindow
	isSet bool
}

func (v NullableTimeWindow) Get() *TimeWindow {
	return v.value
}

func (v *NullableTimeWindow) Set(val *TimeWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeWindow(val *TimeWindow) *NullableTimeWindow {
	return &NullableTimeWindow{value: val, isSet: true}
}

func (v NullableTimeWindow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
