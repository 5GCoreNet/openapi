/*
AMInfluence

AMInfluence API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AMInfluence

import (
	"encoding/json"
)

// checks if the Point type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Point{}

// Point Ellipsoid Point.
type Point struct {
	GADShape
	Point GeographicalCoordinates `json:"point"`
}

// NewPoint instantiates a new Point object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPoint(point GeographicalCoordinates, shape SupportedGADShapes) *Point {
	this := Point{}
	this.Shape = shape
	this.Point = point
	return &this
}

// NewPointWithDefaults instantiates a new Point object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPointWithDefaults() *Point {
	this := Point{}
	return &this
}

// GetPoint returns the Point field value
func (o *Point) GetPoint() GeographicalCoordinates {
	if o == nil {
		var ret GeographicalCoordinates
		return ret
	}

	return o.Point
}

// GetPointOk returns a tuple with the Point field value
// and a boolean to check if the value has been set.
func (o *Point) GetPointOk() (*GeographicalCoordinates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Point, true
}

// SetPoint sets field value
func (o *Point) SetPoint(v GeographicalCoordinates) {
	o.Point = v
}

func (o Point) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Point) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullablePoint struct {
	value *Point
	isSet bool
}

func (v NullablePoint) Get() *Point {
	return v.value
}

func (v *NullablePoint) Set(val *Point) {
	v.value = val
	v.isSet = true
}

func (v NullablePoint) IsSet() bool {
	return v.isSet
}

func (v *NullablePoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePoint(val *Point) *NullablePoint {
	return &NullablePoint{value: val, isSet: true}
}

func (v NullablePoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
