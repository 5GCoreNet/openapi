/*
Nnwdaf_EventsSubscription

Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_EventsSubscription

import (
	"encoding/json"
	"time"
)

// checks if the UeMobility type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UeMobility{}

// UeMobility Represents UE mobility information.
type UeMobility struct {
	// string with format 'date-time' as defined in OpenAPI.
	Ts            *time.Time                  `json:"ts,omitempty"`
	RecurringTime *ScheduledCommunicationTime `json:"recurringTime,omitempty"`
	// indicating a time in seconds.
	Duration *int32 `json:"duration,omitempty"`
	// string with format 'float' as defined in OpenAPI.
	DurationVariance *float32       `json:"durationVariance,omitempty"`
	LocInfos         []LocationInfo `json:"locInfos,omitempty"`
}

// NewUeMobility instantiates a new UeMobility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeMobility() *UeMobility {
	this := UeMobility{}
	return &this
}

// NewUeMobilityWithDefaults instantiates a new UeMobility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeMobilityWithDefaults() *UeMobility {
	this := UeMobility{}
	return &this
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *UeMobility) GetTs() time.Time {
	if o == nil || IsNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeMobility) GetTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// HasTs returns a boolean if a field has been set.
func (o *UeMobility) HasTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *UeMobility) SetTs(v time.Time) {
	o.Ts = &v
}

// GetRecurringTime returns the RecurringTime field value if set, zero value otherwise.
func (o *UeMobility) GetRecurringTime() ScheduledCommunicationTime {
	if o == nil || IsNil(o.RecurringTime) {
		var ret ScheduledCommunicationTime
		return ret
	}
	return *o.RecurringTime
}

// GetRecurringTimeOk returns a tuple with the RecurringTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeMobility) GetRecurringTimeOk() (*ScheduledCommunicationTime, bool) {
	if o == nil || IsNil(o.RecurringTime) {
		return nil, false
	}
	return o.RecurringTime, true
}

// HasRecurringTime returns a boolean if a field has been set.
func (o *UeMobility) HasRecurringTime() bool {
	if o != nil && !IsNil(o.RecurringTime) {
		return true
	}

	return false
}

// SetRecurringTime gets a reference to the given ScheduledCommunicationTime and assigns it to the RecurringTime field.
func (o *UeMobility) SetRecurringTime(v ScheduledCommunicationTime) {
	o.RecurringTime = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *UeMobility) GetDuration() int32 {
	if o == nil || IsNil(o.Duration) {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeMobility) GetDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *UeMobility) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *UeMobility) SetDuration(v int32) {
	o.Duration = &v
}

// GetDurationVariance returns the DurationVariance field value if set, zero value otherwise.
func (o *UeMobility) GetDurationVariance() float32 {
	if o == nil || IsNil(o.DurationVariance) {
		var ret float32
		return ret
	}
	return *o.DurationVariance
}

// GetDurationVarianceOk returns a tuple with the DurationVariance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeMobility) GetDurationVarianceOk() (*float32, bool) {
	if o == nil || IsNil(o.DurationVariance) {
		return nil, false
	}
	return o.DurationVariance, true
}

// HasDurationVariance returns a boolean if a field has been set.
func (o *UeMobility) HasDurationVariance() bool {
	if o != nil && !IsNil(o.DurationVariance) {
		return true
	}

	return false
}

// SetDurationVariance gets a reference to the given float32 and assigns it to the DurationVariance field.
func (o *UeMobility) SetDurationVariance(v float32) {
	o.DurationVariance = &v
}

// GetLocInfos returns the LocInfos field value if set, zero value otherwise.
func (o *UeMobility) GetLocInfos() []LocationInfo {
	if o == nil || IsNil(o.LocInfos) {
		var ret []LocationInfo
		return ret
	}
	return o.LocInfos
}

// GetLocInfosOk returns a tuple with the LocInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeMobility) GetLocInfosOk() ([]LocationInfo, bool) {
	if o == nil || IsNil(o.LocInfos) {
		return nil, false
	}
	return o.LocInfos, true
}

// HasLocInfos returns a boolean if a field has been set.
func (o *UeMobility) HasLocInfos() bool {
	if o != nil && !IsNil(o.LocInfos) {
		return true
	}

	return false
}

// SetLocInfos gets a reference to the given []LocationInfo and assigns it to the LocInfos field.
func (o *UeMobility) SetLocInfos(v []LocationInfo) {
	o.LocInfos = v
}

func (o UeMobility) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UeMobility) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}
	if !IsNil(o.RecurringTime) {
		toSerialize["recurringTime"] = o.RecurringTime
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.DurationVariance) {
		toSerialize["durationVariance"] = o.DurationVariance
	}
	if !IsNil(o.LocInfos) {
		toSerialize["locInfos"] = o.LocInfos
	}
	return toSerialize, nil
}

type NullableUeMobility struct {
	value *UeMobility
	isSet bool
}

func (v NullableUeMobility) Get() *UeMobility {
	return v.value
}

func (v *NullableUeMobility) Set(val *UeMobility) {
	v.value = val
	v.isSet = true
}

func (v NullableUeMobility) IsSet() bool {
	return v.isSet
}

func (v *NullableUeMobility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeMobility(val *UeMobility) *NullableUeMobility {
	return &NullableUeMobility{value: val, isSet: true}
}

func (v NullableUeMobility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeMobility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
