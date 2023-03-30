/*
UAE Server Real-time UAV Status Service

UAE Server Real-time UAV Status Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_UAE_RealtimeUAVStatus

import (
	"encoding/json"
)

// checks if the MinorLocationQoS type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MinorLocationQoS{}

// MinorLocationQoS Contain Minor Location QoS.
type MinorLocationQoS struct {
	// Indicates value of accuracy.
	HAccuracy *float32 `json:"hAccuracy,omitempty"`
	// Indicates value of accuracy.
	VAccuracy *float32 `json:"vAccuracy,omitempty"`
}

// NewMinorLocationQoS instantiates a new MinorLocationQoS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinorLocationQoS() *MinorLocationQoS {
	this := MinorLocationQoS{}
	return &this
}

// NewMinorLocationQoSWithDefaults instantiates a new MinorLocationQoS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinorLocationQoSWithDefaults() *MinorLocationQoS {
	this := MinorLocationQoS{}
	return &this
}

// GetHAccuracy returns the HAccuracy field value if set, zero value otherwise.
func (o *MinorLocationQoS) GetHAccuracy() float32 {
	if o == nil || IsNil(o.HAccuracy) {
		var ret float32
		return ret
	}
	return *o.HAccuracy
}

// GetHAccuracyOk returns a tuple with the HAccuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinorLocationQoS) GetHAccuracyOk() (*float32, bool) {
	if o == nil || IsNil(o.HAccuracy) {
		return nil, false
	}
	return o.HAccuracy, true
}

// HasHAccuracy returns a boolean if a field has been set.
func (o *MinorLocationQoS) HasHAccuracy() bool {
	if o != nil && !IsNil(o.HAccuracy) {
		return true
	}

	return false
}

// SetHAccuracy gets a reference to the given float32 and assigns it to the HAccuracy field.
func (o *MinorLocationQoS) SetHAccuracy(v float32) {
	o.HAccuracy = &v
}

// GetVAccuracy returns the VAccuracy field value if set, zero value otherwise.
func (o *MinorLocationQoS) GetVAccuracy() float32 {
	if o == nil || IsNil(o.VAccuracy) {
		var ret float32
		return ret
	}
	return *o.VAccuracy
}

// GetVAccuracyOk returns a tuple with the VAccuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinorLocationQoS) GetVAccuracyOk() (*float32, bool) {
	if o == nil || IsNil(o.VAccuracy) {
		return nil, false
	}
	return o.VAccuracy, true
}

// HasVAccuracy returns a boolean if a field has been set.
func (o *MinorLocationQoS) HasVAccuracy() bool {
	if o != nil && !IsNil(o.VAccuracy) {
		return true
	}

	return false
}

// SetVAccuracy gets a reference to the given float32 and assigns it to the VAccuracy field.
func (o *MinorLocationQoS) SetVAccuracy(v float32) {
	o.VAccuracy = &v
}

func (o MinorLocationQoS) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MinorLocationQoS) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HAccuracy) {
		toSerialize["hAccuracy"] = o.HAccuracy
	}
	if !IsNil(o.VAccuracy) {
		toSerialize["vAccuracy"] = o.VAccuracy
	}
	return toSerialize, nil
}

type NullableMinorLocationQoS struct {
	value *MinorLocationQoS
	isSet bool
}

func (v NullableMinorLocationQoS) Get() *MinorLocationQoS {
	return v.value
}

func (v *NullableMinorLocationQoS) Set(val *MinorLocationQoS) {
	v.value = val
	v.isSet = true
}

func (v NullableMinorLocationQoS) IsSet() bool {
	return v.isSet
}

func (v *NullableMinorLocationQoS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinorLocationQoS(val *MinorLocationQoS) *NullableMinorLocationQoS {
	return &NullableMinorLocationQoS{value: val, isSet: true}
}

func (v NullableMinorLocationQoS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinorLocationQoS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


