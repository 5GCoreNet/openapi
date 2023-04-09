/*
3gpp-am-policyauthorization

API for AM policy authorization.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AMPolicyAuthorization

import (
	"encoding/json"
)

// checks if the PointAltitude type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PointAltitude{}

// PointAltitude Ellipsoid point with altitude.
type PointAltitude struct {
	GADShape
	Point GeographicalCoordinates `json:"point"`
	// Indicates value of altitude.
	Altitude float64 `json:"altitude"`
}

// NewPointAltitude instantiates a new PointAltitude object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPointAltitude(point GeographicalCoordinates, altitude float64, shape SupportedGADShapes) *PointAltitude {
	this := PointAltitude{}
	this.Shape = shape
	this.Point = point
	this.Altitude = altitude
	return &this
}

// NewPointAltitudeWithDefaults instantiates a new PointAltitude object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPointAltitudeWithDefaults() *PointAltitude {
	this := PointAltitude{}
	return &this
}

// GetPoint returns the Point field value
func (o *PointAltitude) GetPoint() GeographicalCoordinates {
	if o == nil {
		var ret GeographicalCoordinates
		return ret
	}

	return o.Point
}

// GetPointOk returns a tuple with the Point field value
// and a boolean to check if the value has been set.
func (o *PointAltitude) GetPointOk() (*GeographicalCoordinates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Point, true
}

// SetPoint sets field value
func (o *PointAltitude) SetPoint(v GeographicalCoordinates) {
	o.Point = v
}

// GetAltitude returns the Altitude field value
func (o *PointAltitude) GetAltitude() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Altitude
}

// GetAltitudeOk returns a tuple with the Altitude field value
// and a boolean to check if the value has been set.
func (o *PointAltitude) GetAltitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Altitude, true
}

// SetAltitude sets field value
func (o *PointAltitude) SetAltitude(v float64) {
	o.Altitude = v
}

func (o PointAltitude) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PointAltitude) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedGADShape, errGADShape := json.Marshal(o.GADShape)
	if errGADShape != nil {
		return map[string]interface{}{}, errGADShape
	}
	errGADShape = json.Unmarshal([]byte(serializedGADShape), &toSerialize)
	if errGADShape != nil {
		return map[string]interface{}{}, errGADShape
	}
	toSerialize["point"] = o.Point
	toSerialize["altitude"] = o.Altitude
	return toSerialize, nil
}

type NullablePointAltitude struct {
	value *PointAltitude
	isSet bool
}

func (v NullablePointAltitude) Get() *PointAltitude {
	return v.value
}

func (v *NullablePointAltitude) Set(val *PointAltitude) {
	v.value = val
	v.isSet = true
}

func (v NullablePointAltitude) IsSet() bool {
	return v.isSet
}

func (v *NullablePointAltitude) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePointAltitude(val *PointAltitude) *NullablePointAltitude {
	return &NullablePointAltitude{value: val, isSet: true}
}

func (v NullablePointAltitude) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePointAltitude) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
