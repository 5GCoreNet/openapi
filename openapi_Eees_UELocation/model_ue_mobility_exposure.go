/*
EES UE Location Information_API

API for EES UE Location Information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_UELocation

import (
	"encoding/json"
	"time"
)

// checks if the UeMobilityExposure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UeMobilityExposure{}

// UeMobilityExposure Represents a UE mobility information.
type UeMobilityExposure struct {
	// string with format \"date-time\" as defined in OpenAPI.
	Ts *time.Time `json:"ts,omitempty"`
	RecurringTime *ScheduledCommunicationTime `json:"recurringTime,omitempty"`
	// Unsigned integer identifying a period of time in units of seconds.
	Duration int32 `json:"duration"`
	// string with format 'float' as defined in OpenAPI.
	DurationVariance *float32 `json:"durationVariance,omitempty"`
	LocInfo []UeLocationInfo `json:"locInfo"`
}

// NewUeMobilityExposure instantiates a new UeMobilityExposure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeMobilityExposure(duration int32, locInfo []UeLocationInfo) *UeMobilityExposure {
	this := UeMobilityExposure{}
	this.Duration = duration
	this.LocInfo = locInfo
	return &this
}

// NewUeMobilityExposureWithDefaults instantiates a new UeMobilityExposure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeMobilityExposureWithDefaults() *UeMobilityExposure {
	this := UeMobilityExposure{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *UeMobilityExposure) GetTs() time.Time {
	if o == nil || IsNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeMobilityExposure) GetTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *UeMobilityExposure) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *UeMobilityExposure) SetTs(v time.Time) {
	o.Ts = &v
}

// GetRecurringTime returns the RecurringTime field value if set, zero value otherwise.
func (o *UeMobilityExposure) GetRecurringTime() ScheduledCommunicationTime {
	if o == nil || IsNil(o.RecurringTime) {
		var ret ScheduledCommunicationTime
		return ret
	}
	return *o.RecurringTime
}

// GetRecurringTimeOk returns a tuple with the RecurringTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeMobilityExposure) GetRecurringTimeOk() (*ScheduledCommunicationTime, bool) {
	if o == nil || IsNil(o.RecurringTime) {
		return nil, false
	}
	return o.RecurringTime, true
}

// HasRecurringTime returns a boolean if a field has been set.
func (o *UeMobilityExposure) HasRecurringTime() bool {
	if o != nil && !IsNil(o.RecurringTime) {
		return true
	}

	return false
}

// SetRecurringTime gets a reference to the given ScheduledCommunicationTime and assigns it to the RecurringTime field.
func (o *UeMobilityExposure) SetRecurringTime(v ScheduledCommunicationTime) {
	o.RecurringTime = &v
}

// GetDuration returns the Duration field value
func (o *UeMobilityExposure) GetDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *UeMobilityExposure) GetDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *UeMobilityExposure) SetDuration(v int32) {
	o.Duration = v
}

// GetDurationVariance returns the DurationVariance field value if set, zero value otherwise.
func (o *UeMobilityExposure) GetDurationVariance() float32 {
	if o == nil || IsNil(o.DurationVariance) {
		var ret float32
		return ret
	}
	return *o.DurationVariance
}

// GetDurationVarianceOk returns a tuple with the DurationVariance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeMobilityExposure) GetDurationVarianceOk() (*float32, bool) {
	if o == nil || IsNil(o.DurationVariance) {
		return nil, false
	}
	return o.DurationVariance, true
}

// HasDurationVariance returns a boolean if a field has been set.
func (o *UeMobilityExposure) HasDurationVariance() bool {
	if o != nil && !IsNil(o.DurationVariance) {
		return true
	}

	return false
}

// SetDurationVariance gets a reference to the given float32 and assigns it to the DurationVariance field.
func (o *UeMobilityExposure) SetDurationVariance(v float32) {
	o.DurationVariance = &v
}

// GetLocInfo returns the LocInfo field value
func (o *UeMobilityExposure) GetLocInfo() []UeLocationInfo {
	if o == nil {
		var ret []UeLocationInfo
		return ret
	}

	return o.LocInfo
}

// GetLocInfoOk returns a tuple with the LocInfo field value
// and a boolean to check if the value has been set.
func (o *UeMobilityExposure) GetLocInfoOk() ([]UeLocationInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.LocInfo, true
}

// SetLocInfo sets field value
func (o *UeMobilityExposure) SetLocInfo(v []UeLocationInfo) {
	o.LocInfo = v
}

func (o UeMobilityExposure) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UeMobilityExposure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.RecurringTime) {
		toSerialize["recurringTime"] = o.RecurringTime
	}
	toSerialize["duration"] = o.Duration
	if !IsNil(o.DurationVariance) {
		toSerialize["durationVariance"] = o.DurationVariance
	}
	toSerialize["locInfo"] = o.LocInfo
	return toSerialize, nil
}

type NullableUeMobilityExposure struct {
	value *UeMobilityExposure
	isSet bool
}

func (v NullableUeMobilityExposure) Get() *UeMobilityExposure {
	return v.value
}

func (v *NullableUeMobilityExposure) Set(val *UeMobilityExposure) {
	v.value = val
	v.isSet = true
}

func (v NullableUeMobilityExposure) IsSet() bool {
	return v.isSet
}

func (v *NullableUeMobilityExposure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeMobilityExposure(val *UeMobilityExposure) *NullableUeMobilityExposure {
	return &NullableUeMobilityExposure{value: val, isSet: true}
}

func (v NullableUeMobilityExposure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeMobilityExposure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


