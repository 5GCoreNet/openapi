/*
EES ACR Management Event_API

API for EES ACR Management Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACRManagementEvent

import (
	"encoding/json"
)

// checks if the FailureAcrMgntEventInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FailureAcrMgntEventInfo{}

// FailureAcrMgntEventInfo Represents a failure ACR management event.
type FailureAcrMgntEventInfo struct {
	Event AcrMgntEvent `json:"event"`
	FailureCode AcrMgntEventFailureCode `json:"failureCode"`
}

// NewFailureAcrMgntEventInfo instantiates a new FailureAcrMgntEventInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFailureAcrMgntEventInfo(event AcrMgntEvent, failureCode AcrMgntEventFailureCode) *FailureAcrMgntEventInfo {
	this := FailureAcrMgntEventInfo{}
	this.Event = event
	this.FailureCode = failureCode
	return &this
}

// NewFailureAcrMgntEventInfoWithDefaults instantiates a new FailureAcrMgntEventInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFailureAcrMgntEventInfoWithDefaults() *FailureAcrMgntEventInfo {
	this := FailureAcrMgntEventInfo{}
	return &this
}

// GetEvent returns the Event field value
func (o *FailureAcrMgntEventInfo) GetEvent() AcrMgntEvent {
	if o == nil {
		var ret AcrMgntEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *FailureAcrMgntEventInfo) GetEventOk() (*AcrMgntEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *FailureAcrMgntEventInfo) SetEvent(v AcrMgntEvent) {
	o.Event = v
}

// GetFailureCode returns the FailureCode field value
func (o *FailureAcrMgntEventInfo) GetFailureCode() AcrMgntEventFailureCode {
	if o == nil {
		var ret AcrMgntEventFailureCode
		return ret
	}

	return o.FailureCode
}

// GetFailureCodeOk returns a tuple with the FailureCode field value
// and a boolean to check if the value has been set.
func (o *FailureAcrMgntEventInfo) GetFailureCodeOk() (*AcrMgntEventFailureCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailureCode, true
}

// SetFailureCode sets field value
func (o *FailureAcrMgntEventInfo) SetFailureCode(v AcrMgntEventFailureCode) {
	o.FailureCode = v
}

func (o FailureAcrMgntEventInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FailureAcrMgntEventInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	toSerialize["failureCode"] = o.FailureCode
	return toSerialize, nil
}

type NullableFailureAcrMgntEventInfo struct {
	value *FailureAcrMgntEventInfo
	isSet bool
}

func (v NullableFailureAcrMgntEventInfo) Get() *FailureAcrMgntEventInfo {
	return v.value
}

func (v *NullableFailureAcrMgntEventInfo) Set(val *FailureAcrMgntEventInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFailureAcrMgntEventInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFailureAcrMgntEventInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailureAcrMgntEventInfo(val *FailureAcrMgntEventInfo) *NullableFailureAcrMgntEventInfo {
	return &NullableFailureAcrMgntEventInfo{value: val, isSet: true}
}

func (v NullableFailureAcrMgntEventInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailureAcrMgntEventInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


