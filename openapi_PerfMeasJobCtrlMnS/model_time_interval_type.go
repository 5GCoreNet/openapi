/*
TS 28.550 Performance Measurement Job Control Service

OAS 3.0.1 specification of the Performance Measurement Job Control Service @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 16.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_PerfMeasJobCtrlMnS

import (
	"encoding/json"
)

// checks if the TimeIntervalType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeIntervalType{}

// TimeIntervalType struct for TimeIntervalType
type TimeIntervalType struct {
	IntervalStart *string `json:"intervalStart,omitempty"`
	IntervalEnd   *string `json:"intervalEnd,omitempty"`
}

// NewTimeIntervalType instantiates a new TimeIntervalType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeIntervalType() *TimeIntervalType {
	this := TimeIntervalType{}
	return &this
}

// NewTimeIntervalTypeWithDefaults instantiates a new TimeIntervalType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeIntervalTypeWithDefaults() *TimeIntervalType {
	this := TimeIntervalType{}
	return &this
}

// GetIntervalStart returns the IntervalStart field value if set, zero value otherwise.
func (o *TimeIntervalType) GetIntervalStart() string {
	if o == nil || IsNil(o.IntervalStart) {
		var ret string
		return ret
	}
	return *o.IntervalStart
}

// GetIntervalStartOk returns a tuple with the IntervalStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeIntervalType) GetIntervalStartOk() (*string, bool) {
	if o == nil || IsNil(o.IntervalStart) {
		return nil, false
	}
	return o.IntervalStart, true
}

// HasIntervalStart returns a boolean if a field has been set.
func (o *TimeIntervalType) HasIntervalStart() bool {
	if o != nil && !IsNil(o.IntervalStart) {
		return true
	}

	return false
}

// SetIntervalStart gets a reference to the given string and assigns it to the IntervalStart field.
func (o *TimeIntervalType) SetIntervalStart(v string) {
	o.IntervalStart = &v
}

// GetIntervalEnd returns the IntervalEnd field value if set, zero value otherwise.
func (o *TimeIntervalType) GetIntervalEnd() string {
	if o == nil || IsNil(o.IntervalEnd) {
		var ret string
		return ret
	}
	return *o.IntervalEnd
}

// GetIntervalEndOk returns a tuple with the IntervalEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeIntervalType) GetIntervalEndOk() (*string, bool) {
	if o == nil || IsNil(o.IntervalEnd) {
		return nil, false
	}
	return o.IntervalEnd, true
}

// HasIntervalEnd returns a boolean if a field has been set.
func (o *TimeIntervalType) HasIntervalEnd() bool {
	if o != nil && !IsNil(o.IntervalEnd) {
		return true
	}

	return false
}

// SetIntervalEnd gets a reference to the given string and assigns it to the IntervalEnd field.
func (o *TimeIntervalType) SetIntervalEnd(v string) {
	o.IntervalEnd = &v
}

func (o TimeIntervalType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeIntervalType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IntervalStart) {
		toSerialize["intervalStart"] = o.IntervalStart
	}
	if !IsNil(o.IntervalEnd) {
		toSerialize["intervalEnd"] = o.IntervalEnd
	}
	return toSerialize, nil
}

type NullableTimeIntervalType struct {
	value *TimeIntervalType
	isSet bool
}

func (v NullableTimeIntervalType) Get() *TimeIntervalType {
	return v.value
}

func (v *NullableTimeIntervalType) Set(val *TimeIntervalType) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeIntervalType) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeIntervalType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeIntervalType(val *TimeIntervalType) *NullableTimeIntervalType {
	return &NullableTimeIntervalType{value: val, isSet: true}
}

func (v NullableTimeIntervalType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeIntervalType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
