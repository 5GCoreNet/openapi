/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3caDataManagement

import (
	"encoding/json"
)

// checks if the AmfEventState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmfEventState{}

// AmfEventState Represents the state of a subscribed event
type AmfEventState struct {
	Active bool `json:"active"`
	RemainReports *int32 `json:"remainReports,omitempty"`
	// indicating a time in seconds.
	RemainDuration *int32 `json:"remainDuration,omitempty"`
}

// NewAmfEventState instantiates a new AmfEventState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfEventState(active bool) *AmfEventState {
	this := AmfEventState{}
	this.Active = active
	return &this
}

// NewAmfEventStateWithDefaults instantiates a new AmfEventState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfEventStateWithDefaults() *AmfEventState {
	this := AmfEventState{}
	return &this
}

// GetActive returns the Active field value
func (o *AmfEventState) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *AmfEventState) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *AmfEventState) SetActive(v bool) {
	o.Active = v
}

// GetRemainReports returns the RemainReports field value if set, zero value otherwise.
func (o *AmfEventState) GetRemainReports() int32 {
	if o == nil || isNil(o.RemainReports) {
		var ret int32
		return ret
	}
	return *o.RemainReports
}

// GetRemainReportsOk returns a tuple with the RemainReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventState) GetRemainReportsOk() (*int32, bool) {
	if o == nil || isNil(o.RemainReports) {
		return nil, false
	}
	return o.RemainReports, true
}

// HasRemainReports returns a boolean if a field has been set.
func (o *AmfEventState) HasRemainReports() bool {
	if o != nil && !isNil(o.RemainReports) {
		return true
	}

	return false
}

// SetRemainReports gets a reference to the given int32 and assigns it to the RemainReports field.
func (o *AmfEventState) SetRemainReports(v int32) {
	o.RemainReports = &v
}

// GetRemainDuration returns the RemainDuration field value if set, zero value otherwise.
func (o *AmfEventState) GetRemainDuration() int32 {
	if o == nil || isNil(o.RemainDuration) {
		var ret int32
		return ret
	}
	return *o.RemainDuration
}

// GetRemainDurationOk returns a tuple with the RemainDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventState) GetRemainDurationOk() (*int32, bool) {
	if o == nil || isNil(o.RemainDuration) {
		return nil, false
	}
	return o.RemainDuration, true
}

// HasRemainDuration returns a boolean if a field has been set.
func (o *AmfEventState) HasRemainDuration() bool {
	if o != nil && !isNil(o.RemainDuration) {
		return true
	}

	return false
}

// SetRemainDuration gets a reference to the given int32 and assigns it to the RemainDuration field.
func (o *AmfEventState) SetRemainDuration(v int32) {
	o.RemainDuration = &v
}

func (o AmfEventState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmfEventState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	if !isNil(o.RemainReports) {
		toSerialize["remainReports"] = o.RemainReports
	}
	if !isNil(o.RemainDuration) {
		toSerialize["remainDuration"] = o.RemainDuration
	}
	return toSerialize, nil
}

type NullableAmfEventState struct {
	value *AmfEventState
	isSet bool
}

func (v NullableAmfEventState) Get() *AmfEventState {
	return v.value
}

func (v *NullableAmfEventState) Set(val *AmfEventState) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfEventState) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfEventState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfEventState(val *AmfEventState) *NullableAmfEventState {
	return &NullableAmfEventState{value: val, isSet: true}
}

func (v NullableAmfEventState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfEventState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


