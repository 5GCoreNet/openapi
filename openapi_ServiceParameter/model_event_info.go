/*
3gpp-service-parameter

API for AF service paramter   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ServiceParameter

import (
	"encoding/json"
)

// checks if the EventInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventInfo{}

// EventInfo Indicates the event information.
type EventInfo struct {
	FailureCause *Failure `json:"failureCause,omitempty"`
}

// NewEventInfo instantiates a new EventInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventInfo() *EventInfo {
	this := EventInfo{}
	return &this
}

// NewEventInfoWithDefaults instantiates a new EventInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventInfoWithDefaults() *EventInfo {
	this := EventInfo{}
	return &this
}

// GetFailureCause returns the FailureCause field value if set, zero value otherwise.
func (o *EventInfo) GetFailureCause() Failure {
	if o == nil || IsNil(o.FailureCause) {
		var ret Failure
		return ret
	}
	return *o.FailureCause
}

// GetFailureCauseOk returns a tuple with the FailureCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventInfo) GetFailureCauseOk() (*Failure, bool) {
	if o == nil || IsNil(o.FailureCause) {
		return nil, false
	}
	return o.FailureCause, true
}

// HasFailureCause returns a boolean if a field has been set.
func (o *EventInfo) HasFailureCause() bool {
	if o != nil && !IsNil(o.FailureCause) {
		return true
	}

	return false
}

// SetFailureCause gets a reference to the given Failure and assigns it to the FailureCause field.
func (o *EventInfo) SetFailureCause(v Failure) {
	o.FailureCause = &v
}

func (o EventInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FailureCause) {
		toSerialize["failureCause"] = o.FailureCause
	}
	return toSerialize, nil
}

type NullableEventInfo struct {
	value *EventInfo
	isSet bool
}

func (v NullableEventInfo) Get() *EventInfo {
	return v.value
}

func (v *NullableEventInfo) Set(val *EventInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEventInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEventInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventInfo(val *EventInfo) *NullableEventInfo {
	return &NullableEventInfo{value: val, isSet: true}
}

func (v NullableEventInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


