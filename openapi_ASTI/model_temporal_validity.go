/*
3gpp-asti

API for ASTI.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ASTI

import (
	"encoding/json"
	"time"
)

// checks if the TemporalValidity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemporalValidity{}

// TemporalValidity Indicates the time interval(s) during which the AF request is to be applied.
type TemporalValidity struct {
	// string with format 'date-time' as defined in OpenAPI.
	StartTime *time.Time `json:"startTime,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	StopTime *time.Time `json:"stopTime,omitempty"`
}

// NewTemporalValidity instantiates a new TemporalValidity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemporalValidity() *TemporalValidity {
	this := TemporalValidity{}
	return &this
}

// NewTemporalValidityWithDefaults instantiates a new TemporalValidity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemporalValidityWithDefaults() *TemporalValidity {
	this := TemporalValidity{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *TemporalValidity) GetStartTime() time.Time {
	if o == nil || isNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemporalValidity) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *TemporalValidity) HasStartTime() bool {
	if o != nil && !isNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *TemporalValidity) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStopTime returns the StopTime field value if set, zero value otherwise.
func (o *TemporalValidity) GetStopTime() time.Time {
	if o == nil || isNil(o.StopTime) {
		var ret time.Time
		return ret
	}
	return *o.StopTime
}

// GetStopTimeOk returns a tuple with the StopTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemporalValidity) GetStopTimeOk() (*time.Time, bool) {
	if o == nil || isNil(o.StopTime) {
		return nil, false
	}
	return o.StopTime, true
}

// HasStopTime returns a boolean if a field has been set.
func (o *TemporalValidity) HasStopTime() bool {
	if o != nil && !isNil(o.StopTime) {
		return true
	}

	return false
}

// SetStopTime gets a reference to the given time.Time and assigns it to the StopTime field.
func (o *TemporalValidity) SetStopTime(v time.Time) {
	o.StopTime = &v
}

func (o TemporalValidity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemporalValidity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.StartTime) {
		toSerialize["startTime"] = o.StartTime
	}
	if !isNil(o.StopTime) {
		toSerialize["stopTime"] = o.StopTime
	}
	return toSerialize, nil
}

type NullableTemporalValidity struct {
	value *TemporalValidity
	isSet bool
}

func (v NullableTemporalValidity) Get() *TemporalValidity {
	return v.value
}

func (v *NullableTemporalValidity) Set(val *TemporalValidity) {
	v.value = val
	v.isSet = true
}

func (v NullableTemporalValidity) IsSet() bool {
	return v.isSet
}

func (v *NullableTemporalValidity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemporalValidity(val *TemporalValidity) *NullableTemporalValidity {
	return &NullableTemporalValidity{value: val, isSet: true}
}

func (v NullableTemporalValidity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemporalValidity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


