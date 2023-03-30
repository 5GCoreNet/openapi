/*
Ngmlc_Location

GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ngmlc_Location

import (
	"encoding/json"
)

// checks if the HorizontalWithVerticalVelocityAndUncertainty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HorizontalWithVerticalVelocityAndUncertainty{}

// HorizontalWithVerticalVelocityAndUncertainty Horizontal and vertical velocity with speed uncertainty.
type HorizontalWithVerticalVelocityAndUncertainty struct {
	// Indicates value of horizontal speed.
	HSpeed float32 `json:"hSpeed"`
	// Indicates value of angle.
	Bearing int32 `json:"bearing"`
	// Indicates value of vertical speed.
	VSpeed float32 `json:"vSpeed"`
	VDirection VerticalDirection `json:"vDirection"`
	// Indicates value of speed uncertainty.
	HUncertainty float32 `json:"hUncertainty"`
	// Indicates value of speed uncertainty.
	VUncertainty float32 `json:"vUncertainty"`
}

// NewHorizontalWithVerticalVelocityAndUncertainty instantiates a new HorizontalWithVerticalVelocityAndUncertainty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHorizontalWithVerticalVelocityAndUncertainty(hSpeed float32, bearing int32, vSpeed float32, vDirection VerticalDirection, hUncertainty float32, vUncertainty float32) *HorizontalWithVerticalVelocityAndUncertainty {
	this := HorizontalWithVerticalVelocityAndUncertainty{}
	this.HSpeed = hSpeed
	this.Bearing = bearing
	this.VSpeed = vSpeed
	this.VDirection = vDirection
	this.HUncertainty = hUncertainty
	this.VUncertainty = vUncertainty
	return &this
}

// NewHorizontalWithVerticalVelocityAndUncertaintyWithDefaults instantiates a new HorizontalWithVerticalVelocityAndUncertainty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHorizontalWithVerticalVelocityAndUncertaintyWithDefaults() *HorizontalWithVerticalVelocityAndUncertainty {
	this := HorizontalWithVerticalVelocityAndUncertainty{}
	return &this
}

// GetHSpeed returns the HSpeed field value
func (o *HorizontalWithVerticalVelocityAndUncertainty) GetHSpeed() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.HSpeed
}

// GetHSpeedOk returns a tuple with the HSpeed field value
// and a boolean to check if the value has been set.
func (o *HorizontalWithVerticalVelocityAndUncertainty) GetHSpeedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HSpeed, true
}

// SetHSpeed sets field value
func (o *HorizontalWithVerticalVelocityAndUncertainty) SetHSpeed(v float32) {
	o.HSpeed = v
}

// GetBearing returns the Bearing field value
func (o *HorizontalWithVerticalVelocityAndUncertainty) GetBearing() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Bearing
}

// GetBearingOk returns a tuple with the Bearing field value
// and a boolean to check if the value has been set.
func (o *HorizontalWithVerticalVelocityAndUncertainty) GetBearingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bearing, true
}

// SetBearing sets field value
func (o *HorizontalWithVerticalVelocityAndUncertainty) SetBearing(v int32) {
	o.Bearing = v
}

// GetVSpeed returns the VSpeed field value
func (o *HorizontalWithVerticalVelocityAndUncertainty) GetVSpeed() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.VSpeed
}

// GetVSpeedOk returns a tuple with the VSpeed field value
// and a boolean to check if the value has been set.
func (o *HorizontalWithVerticalVelocityAndUncertainty) GetVSpeedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VSpeed, true
}

// SetVSpeed sets field value
func (o *HorizontalWithVerticalVelocityAndUncertainty) SetVSpeed(v float32) {
	o.VSpeed = v
}

// GetVDirection returns the VDirection field value
func (o *HorizontalWithVerticalVelocityAndUncertainty) GetVDirection() VerticalDirection {
	if o == nil {
		var ret VerticalDirection
		return ret
	}

	return o.VDirection
}

// GetVDirectionOk returns a tuple with the VDirection field value
// and a boolean to check if the value has been set.
func (o *HorizontalWithVerticalVelocityAndUncertainty) GetVDirectionOk() (*VerticalDirection, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VDirection, true
}

// SetVDirection sets field value
func (o *HorizontalWithVerticalVelocityAndUncertainty) SetVDirection(v VerticalDirection) {
	o.VDirection = v
}

// GetHUncertainty returns the HUncertainty field value
func (o *HorizontalWithVerticalVelocityAndUncertainty) GetHUncertainty() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.HUncertainty
}

// GetHUncertaintyOk returns a tuple with the HUncertainty field value
// and a boolean to check if the value has been set.
func (o *HorizontalWithVerticalVelocityAndUncertainty) GetHUncertaintyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HUncertainty, true
}

// SetHUncertainty sets field value
func (o *HorizontalWithVerticalVelocityAndUncertainty) SetHUncertainty(v float32) {
	o.HUncertainty = v
}

// GetVUncertainty returns the VUncertainty field value
func (o *HorizontalWithVerticalVelocityAndUncertainty) GetVUncertainty() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.VUncertainty
}

// GetVUncertaintyOk returns a tuple with the VUncertainty field value
// and a boolean to check if the value has been set.
func (o *HorizontalWithVerticalVelocityAndUncertainty) GetVUncertaintyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VUncertainty, true
}

// SetVUncertainty sets field value
func (o *HorizontalWithVerticalVelocityAndUncertainty) SetVUncertainty(v float32) {
	o.VUncertainty = v
}

func (o HorizontalWithVerticalVelocityAndUncertainty) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HorizontalWithVerticalVelocityAndUncertainty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hSpeed"] = o.HSpeed
	toSerialize["bearing"] = o.Bearing
	toSerialize["vSpeed"] = o.VSpeed
	toSerialize["vDirection"] = o.VDirection
	toSerialize["hUncertainty"] = o.HUncertainty
	toSerialize["vUncertainty"] = o.VUncertainty
	return toSerialize, nil
}

type NullableHorizontalWithVerticalVelocityAndUncertainty struct {
	value *HorizontalWithVerticalVelocityAndUncertainty
	isSet bool
}

func (v NullableHorizontalWithVerticalVelocityAndUncertainty) Get() *HorizontalWithVerticalVelocityAndUncertainty {
	return v.value
}

func (v *NullableHorizontalWithVerticalVelocityAndUncertainty) Set(val *HorizontalWithVerticalVelocityAndUncertainty) {
	v.value = val
	v.isSet = true
}

func (v NullableHorizontalWithVerticalVelocityAndUncertainty) IsSet() bool {
	return v.isSet
}

func (v *NullableHorizontalWithVerticalVelocityAndUncertainty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHorizontalWithVerticalVelocityAndUncertainty(val *HorizontalWithVerticalVelocityAndUncertainty) *NullableHorizontalWithVerticalVelocityAndUncertainty {
	return &NullableHorizontalWithVerticalVelocityAndUncertainty{value: val, isSet: true}
}

func (v NullableHorizontalWithVerticalVelocityAndUncertainty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHorizontalWithVerticalVelocityAndUncertainty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


