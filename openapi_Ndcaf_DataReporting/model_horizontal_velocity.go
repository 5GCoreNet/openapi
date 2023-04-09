/*
Ndcaf_DataReporting

Data Collection AF: Data Collection and Reporting Configuration API and Data Reporting API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndcaf_DataReporting

import (
	"encoding/json"
)

// checks if the HorizontalVelocity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HorizontalVelocity{}

// HorizontalVelocity Horizontal velocity.
type HorizontalVelocity struct {
	// Indicates value of horizontal speed.
	HSpeed float32 `json:"hSpeed"`
	// Indicates value of angle.
	Bearing int32 `json:"bearing"`
}

// NewHorizontalVelocity instantiates a new HorizontalVelocity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHorizontalVelocity(hSpeed float32, bearing int32) *HorizontalVelocity {
	this := HorizontalVelocity{}
	this.HSpeed = hSpeed
	this.Bearing = bearing
	return &this
}

// NewHorizontalVelocityWithDefaults instantiates a new HorizontalVelocity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHorizontalVelocityWithDefaults() *HorizontalVelocity {
	this := HorizontalVelocity{}
	return &this
}

// GetHSpeed returns the HSpeed field value
func (o *HorizontalVelocity) GetHSpeed() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.HSpeed
}

// GetHSpeedOk returns a tuple with the HSpeed field value
// and a boolean to check if the value has been set.
func (o *HorizontalVelocity) GetHSpeedOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HSpeed, true
}

// SetHSpeed sets field value
func (o *HorizontalVelocity) SetHSpeed(v float32) {
	o.HSpeed = v
}

// GetBearing returns the Bearing field value
func (o *HorizontalVelocity) GetBearing() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Bearing
}

// GetBearingOk returns a tuple with the Bearing field value
// and a boolean to check if the value has been set.
func (o *HorizontalVelocity) GetBearingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bearing, true
}

// SetBearing sets field value
func (o *HorizontalVelocity) SetBearing(v int32) {
	o.Bearing = v
}

func (o HorizontalVelocity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HorizontalVelocity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hSpeed"] = o.HSpeed
	toSerialize["bearing"] = o.Bearing
	return toSerialize, nil
}

type NullableHorizontalVelocity struct {
	value *HorizontalVelocity
	isSet bool
}

func (v NullableHorizontalVelocity) Get() *HorizontalVelocity {
	return v.value
}

func (v *NullableHorizontalVelocity) Set(val *HorizontalVelocity) {
	v.value = val
	v.isSet = true
}

func (v NullableHorizontalVelocity) IsSet() bool {
	return v.isSet
}

func (v *NullableHorizontalVelocity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHorizontalVelocity(val *HorizontalVelocity) *NullableHorizontalVelocity {
	return &NullableHorizontalVelocity{value: val, isSet: true}
}

func (v NullableHorizontalVelocity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHorizontalVelocity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
