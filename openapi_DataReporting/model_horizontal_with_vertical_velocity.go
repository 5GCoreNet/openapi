/*
3gpp-data-reporting

API for 3GPP Data Reporting.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_DataReporting

import (
	"encoding/json"
)

// checks if the HorizontalWithVerticalVelocity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HorizontalWithVerticalVelocity{}

// HorizontalWithVerticalVelocity Horizontal and vertical velocity.
type HorizontalWithVerticalVelocity struct {
	// Indicates value of horizontal speed.
	HSpeed float32 `json:"hSpeed"`
	// Indicates value of angle.
	Bearing int32 `json:"bearing"`
	// Indicates value of vertical speed.
	VSpeed     float32           `json:"vSpeed"`
	VDirection VerticalDirection `json:"vDirection"`
}

// NewHorizontalWithVerticalVelocity instantiates a new HorizontalWithVerticalVelocity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHorizontalWithVerticalVelocity(hSpeed float32, bearing int32, vSpeed float32, vDirection VerticalDirection) *HorizontalWithVerticalVelocity {
	this := HorizontalWithVerticalVelocity{}
	this.HSpeed = hSpeed
	this.Bearing = bearing
	this.VSpeed = vSpeed
	this.VDirection = vDirection
	return &this
}

// NewHorizontalWithVerticalVelocityWithDefaults instantiates a new HorizontalWithVerticalVelocity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHorizontalWithVerticalVelocityWithDefaults() *HorizontalWithVerticalVelocity {
	this := HorizontalWithVerticalVelocity{}
	return &this
}

// GetHSpeed returns the HSpeed field value
func (o *HorizontalWithVerticalVelocity) GetHSpeed() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.HSpeed
}

// GetHSpeedOk returns a tuple with the HSpeed field value
// and a boolean to check if the value has been set.
func (o *HorizontalWithVerticalVelocity) GetHSpeedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HSpeed, true
}

// SetHSpeed sets field value
func (o *HorizontalWithVerticalVelocity) SetHSpeed(v float32) {
	o.HSpeed = v
}

// GetBearing returns the Bearing field value
func (o *HorizontalWithVerticalVelocity) GetBearing() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Bearing
}

// GetBearingOk returns a tuple with the Bearing field value
// and a boolean to check if the value has been set.
func (o *HorizontalWithVerticalVelocity) GetBearingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bearing, true
}

// SetBearing sets field value
func (o *HorizontalWithVerticalVelocity) SetBearing(v int32) {
	o.Bearing = v
}

// GetVSpeed returns the VSpeed field value
func (o *HorizontalWithVerticalVelocity) GetVSpeed() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.VSpeed
}

// GetVSpeedOk returns a tuple with the VSpeed field value
// and a boolean to check if the value has been set.
func (o *HorizontalWithVerticalVelocity) GetVSpeedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VSpeed, true
}

// SetVSpeed sets field value
func (o *HorizontalWithVerticalVelocity) SetVSpeed(v float32) {
	o.VSpeed = v
}

// GetVDirection returns the VDirection field value
func (o *HorizontalWithVerticalVelocity) GetVDirection() VerticalDirection {
	if o == nil {
		var ret VerticalDirection
		return ret
	}

	return o.VDirection
}

// GetVDirectionOk returns a tuple with the VDirection field value
// and a boolean to check if the value has been set.
func (o *HorizontalWithVerticalVelocity) GetVDirectionOk() (*VerticalDirection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VDirection, true
}

// SetVDirection sets field value
func (o *HorizontalWithVerticalVelocity) SetVDirection(v VerticalDirection) {
	o.VDirection = v
}

func (o HorizontalWithVerticalVelocity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HorizontalWithVerticalVelocity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hSpeed"] = o.HSpeed
	toSerialize["bearing"] = o.Bearing
	toSerialize["vSpeed"] = o.VSpeed
	toSerialize["vDirection"] = o.VDirection
	return toSerialize, nil
}

type NullableHorizontalWithVerticalVelocity struct {
	value *HorizontalWithVerticalVelocity
	isSet bool
}

func (v NullableHorizontalWithVerticalVelocity) Get() *HorizontalWithVerticalVelocity {
	return v.value
}

func (v *NullableHorizontalWithVerticalVelocity) Set(val *HorizontalWithVerticalVelocity) {
	v.value = val
	v.isSet = true
}

func (v NullableHorizontalWithVerticalVelocity) IsSet() bool {
	return v.isSet
}

func (v *NullableHorizontalWithVerticalVelocity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHorizontalWithVerticalVelocity(val *HorizontalWithVerticalVelocity) *NullableHorizontalWithVerticalVelocity {
	return &NullableHorizontalWithVerticalVelocity{value: val, isSet: true}
}

func (v NullableHorizontalWithVerticalVelocity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHorizontalWithVerticalVelocity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
