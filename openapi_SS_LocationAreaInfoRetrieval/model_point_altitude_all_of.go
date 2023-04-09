/*
SS_LocationAreaInfoRetrieval

API for SEAL Location Area Info Retrieval.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_LocationAreaInfoRetrieval

import (
	"encoding/json"
)

// checks if the PointAltitudeAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PointAltitudeAllOf{}

// PointAltitudeAllOf struct for PointAltitudeAllOf
type PointAltitudeAllOf struct {
	Point GeographicalCoordinates `json:"point"`
	// Indicates value of altitude.
	Altitude float64 `json:"altitude"`
}

// NewPointAltitudeAllOf instantiates a new PointAltitudeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPointAltitudeAllOf(point GeographicalCoordinates, altitude float64) *PointAltitudeAllOf {
	this := PointAltitudeAllOf{}
	this.Point = point
	this.Altitude = altitude
	return &this
}

// NewPointAltitudeAllOfWithDefaults instantiates a new PointAltitudeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPointAltitudeAllOfWithDefaults() *PointAltitudeAllOf {
	this := PointAltitudeAllOf{}
	return &this
}

// GetPoint returns the Point field value
func (o *PointAltitudeAllOf) GetPoint() GeographicalCoordinates {
	if o == nil {
		var ret GeographicalCoordinates
		return ret
	}

	return o.Point
}

// GetPointOk returns a tuple with the Point field value
// and a boolean to check if the value has been set.
func (o *PointAltitudeAllOf) GetPointOk() (*GeographicalCoordinates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Point, true
}

// SetPoint sets field value
func (o *PointAltitudeAllOf) SetPoint(v GeographicalCoordinates) {
	o.Point = v
}

// GetAltitude returns the Altitude field value
func (o *PointAltitudeAllOf) GetAltitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Altitude
}

// GetAltitudeOk returns a tuple with the Altitude field value
// and a boolean to check if the value has been set.
func (o *PointAltitudeAllOf) GetAltitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Altitude, true
}

// SetAltitude sets field value
func (o *PointAltitudeAllOf) SetAltitude(v float64) {
	o.Altitude = v
}

func (o PointAltitudeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PointAltitudeAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["point"] = o.Point
	toSerialize["altitude"] = o.Altitude
	return toSerialize, nil
}

type NullablePointAltitudeAllOf struct {
	value *PointAltitudeAllOf
	isSet bool
}

func (v NullablePointAltitudeAllOf) Get() *PointAltitudeAllOf {
	return v.value
}

func (v *NullablePointAltitudeAllOf) Set(val *PointAltitudeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePointAltitudeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePointAltitudeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePointAltitudeAllOf(val *PointAltitudeAllOf) *NullablePointAltitudeAllOf {
	return &NullablePointAltitudeAllOf{value: val, isSet: true}
}

func (v NullablePointAltitudeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePointAltitudeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
