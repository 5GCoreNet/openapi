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

// checks if the MonitorEvents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MonitorEvents{}

// MonitorEvents List of event types to be monitored in the context of events monitoring service.
type MonitorEvents struct {
	// List of monitoring events related to VAL UE.
	CnEvnts []MonitoringType `json:"cnEvnts,omitempty"`
	// List of analytics events related to VAL UE.
	AnlEvnts []AnalyticsEvent `json:"anlEvnts,omitempty"`
}

// NewMonitorEvents instantiates a new MonitorEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMonitorEvents() *MonitorEvents {
	this := MonitorEvents{}
	return &this
}

// NewMonitorEventsWithDefaults instantiates a new MonitorEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMonitorEventsWithDefaults() *MonitorEvents {
	this := MonitorEvents{}
	return &this
}

// GetCnEvnts returns the CnEvnts field value if set, zero value otherwise.
func (o *MonitorEvents) GetCnEvnts() []MonitoringType {
	if o == nil || IsNil(o.CnEvnts) {
		var ret []MonitoringType
		return ret
	}
	return o.CnEvnts
}

// GetCnEvntsOk returns a tuple with the CnEvnts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorEvents) GetCnEvntsOk() ([]MonitoringType, bool) {
	if o == nil || IsNil(o.CnEvnts) {
		return nil, false
	}
	return o.CnEvnts, true
}

// HasCnEvnts returns a boolean if a field has been set.
func (o *MonitorEvents) HasCnEvnts() bool {
	if o != nil && !IsNil(o.CnEvnts) {
		return true
	}

	return false
}

// SetCnEvnts gets a reference to the given []MonitoringType and assigns it to the CnEvnts field.
func (o *MonitorEvents) SetCnEvnts(v []MonitoringType) {
	o.CnEvnts = v
}

// GetAnlEvnts returns the AnlEvnts field value if set, zero value otherwise.
func (o *MonitorEvents) GetAnlEvnts() []AnalyticsEvent {
	if o == nil || IsNil(o.AnlEvnts) {
		var ret []AnalyticsEvent
		return ret
	}
	return o.AnlEvnts
}

// GetAnlEvntsOk returns a tuple with the AnlEvnts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorEvents) GetAnlEvntsOk() ([]AnalyticsEvent, bool) {
	if o == nil || IsNil(o.AnlEvnts) {
		return nil, false
	}
	return o.AnlEvnts, true
}

// HasAnlEvnts returns a boolean if a field has been set.
func (o *MonitorEvents) HasAnlEvnts() bool {
	if o != nil && !IsNil(o.AnlEvnts) {
		return true
	}

	return false
}

// SetAnlEvnts gets a reference to the given []AnalyticsEvent and assigns it to the AnlEvnts field.
func (o *MonitorEvents) SetAnlEvnts(v []AnalyticsEvent) {
	o.AnlEvnts = v
}

func (o MonitorEvents) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MonitorEvents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CnEvnts) {
		toSerialize["cnEvnts"] = o.CnEvnts
	}
	if !IsNil(o.AnlEvnts) {
		toSerialize["anlEvnts"] = o.AnlEvnts
	}
	return toSerialize, nil
}

type NullableMonitorEvents struct {
	value *MonitorEvents
	isSet bool
}

func (v NullableMonitorEvents) Get() *MonitorEvents {
	return v.value
}

func (v *NullableMonitorEvents) Set(val *MonitorEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableMonitorEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableMonitorEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMonitorEvents(val *MonitorEvents) *NullableMonitorEvents {
	return &NullableMonitorEvents{value: val, isSet: true}
}

func (v NullableMonitorEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMonitorEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
