/*
LMF Location

LMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nlmf_Location

import (
	"encoding/json"
)

// checks if the EventReportingStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventReportingStatus{}

// EventReportingStatus Indicates the status of event reporting.
type EventReportingStatus struct {
	// Number of event reports received from the target UE.
	EventReportCounter *int32 `json:"eventReportCounter,omitempty"`
	// Duration of event reporting.
	EventReportDuration *int32 `json:"eventReportDuration,omitempty"`
}

// NewEventReportingStatus instantiates a new EventReportingStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventReportingStatus() *EventReportingStatus {
	this := EventReportingStatus{}
	return &this
}

// NewEventReportingStatusWithDefaults instantiates a new EventReportingStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventReportingStatusWithDefaults() *EventReportingStatus {
	this := EventReportingStatus{}
	return &this
}

// GetEventReportCounter returns the EventReportCounter field value if set, zero value otherwise.
func (o *EventReportingStatus) GetEventReportCounter() int32 {
	if o == nil || isNil(o.EventReportCounter) {
		var ret int32
		return ret
	}
	return *o.EventReportCounter
}

// GetEventReportCounterOk returns a tuple with the EventReportCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventReportingStatus) GetEventReportCounterOk() (*int32, bool) {
	if o == nil || isNil(o.EventReportCounter) {
		return nil, false
	}
	return o.EventReportCounter, true
}

// HasEventReportCounter returns a boolean if a field has been set.
func (o *EventReportingStatus) HasEventReportCounter() bool {
	if o != nil && !isNil(o.EventReportCounter) {
		return true
	}

	return false
}

// SetEventReportCounter gets a reference to the given int32 and assigns it to the EventReportCounter field.
func (o *EventReportingStatus) SetEventReportCounter(v int32) {
	o.EventReportCounter = &v
}

// GetEventReportDuration returns the EventReportDuration field value if set, zero value otherwise.
func (o *EventReportingStatus) GetEventReportDuration() int32 {
	if o == nil || isNil(o.EventReportDuration) {
		var ret int32
		return ret
	}
	return *o.EventReportDuration
}

// GetEventReportDurationOk returns a tuple with the EventReportDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventReportingStatus) GetEventReportDurationOk() (*int32, bool) {
	if o == nil || isNil(o.EventReportDuration) {
		return nil, false
	}
	return o.EventReportDuration, true
}

// HasEventReportDuration returns a boolean if a field has been set.
func (o *EventReportingStatus) HasEventReportDuration() bool {
	if o != nil && !isNil(o.EventReportDuration) {
		return true
	}

	return false
}

// SetEventReportDuration gets a reference to the given int32 and assigns it to the EventReportDuration field.
func (o *EventReportingStatus) SetEventReportDuration(v int32) {
	o.EventReportDuration = &v
}

func (o EventReportingStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventReportingStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EventReportCounter) {
		toSerialize["eventReportCounter"] = o.EventReportCounter
	}
	if !isNil(o.EventReportDuration) {
		toSerialize["eventReportDuration"] = o.EventReportDuration
	}
	return toSerialize, nil
}

type NullableEventReportingStatus struct {
	value *EventReportingStatus
	isSet bool
}

func (v NullableEventReportingStatus) Get() *EventReportingStatus {
	return v.value
}

func (v *NullableEventReportingStatus) Set(val *EventReportingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEventReportingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEventReportingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventReportingStatus(val *EventReportingStatus) *NullableEventReportingStatus {
	return &NullableEventReportingStatus{value: val, isSet: true}
}

func (v NullableEventReportingStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventReportingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


