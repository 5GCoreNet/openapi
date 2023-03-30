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

// checks if the EllipsoidArcAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EllipsoidArcAllOf{}

// EllipsoidArcAllOf struct for EllipsoidArcAllOf
type EllipsoidArcAllOf struct {
	Point GeographicalCoordinates `json:"point"`
	// Indicates value of the inner radius.
	InnerRadius int32 `json:"innerRadius"`
	// Indicates value of uncertainty.
	UncertaintyRadius float32 `json:"uncertaintyRadius"`
	// Indicates value of angle.
	OffsetAngle int32 `json:"offsetAngle"`
	// Indicates value of angle.
	IncludedAngle int32 `json:"includedAngle"`
	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}

// NewEllipsoidArcAllOf instantiates a new EllipsoidArcAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEllipsoidArcAllOf(point GeographicalCoordinates, innerRadius int32, uncertaintyRadius float32, offsetAngle int32, includedAngle int32, confidence int32) *EllipsoidArcAllOf {
	this := EllipsoidArcAllOf{}
	this.Point = point
	this.InnerRadius = innerRadius
	this.UncertaintyRadius = uncertaintyRadius
	this.OffsetAngle = offsetAngle
	this.IncludedAngle = includedAngle
	this.Confidence = confidence
	return &this
}

// NewEllipsoidArcAllOfWithDefaults instantiates a new EllipsoidArcAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEllipsoidArcAllOfWithDefaults() *EllipsoidArcAllOf {
	this := EllipsoidArcAllOf{}
	return &this
}

// GetPoint returns the Point field value
func (o *EllipsoidArcAllOf) GetPoint() GeographicalCoordinates {
	if o == nil {
		var ret GeographicalCoordinates
		return ret
	}

	return o.Point
}

// GetPointOk returns a tuple with the Point field value
// and a boolean to check if the value has been set.
func (o *EllipsoidArcAllOf) GetPointOk() (*GeographicalCoordinates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Point, true
}

// SetPoint sets field value
func (o *EllipsoidArcAllOf) SetPoint(v GeographicalCoordinates) {
	o.Point = v
}

// GetInnerRadius returns the InnerRadius field value
func (o *EllipsoidArcAllOf) GetInnerRadius() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InnerRadius
}

// GetInnerRadiusOk returns a tuple with the InnerRadius field value
// and a boolean to check if the value has been set.
func (o *EllipsoidArcAllOf) GetInnerRadiusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InnerRadius, true
}

// SetInnerRadius sets field value
func (o *EllipsoidArcAllOf) SetInnerRadius(v int32) {
	o.InnerRadius = v
}

// GetUncertaintyRadius returns the UncertaintyRadius field value
func (o *EllipsoidArcAllOf) GetUncertaintyRadius() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.UncertaintyRadius
}

// GetUncertaintyRadiusOk returns a tuple with the UncertaintyRadius field value
// and a boolean to check if the value has been set.
func (o *EllipsoidArcAllOf) GetUncertaintyRadiusOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UncertaintyRadius, true
}

// SetUncertaintyRadius sets field value
func (o *EllipsoidArcAllOf) SetUncertaintyRadius(v float32) {
	o.UncertaintyRadius = v
}

// GetOffsetAngle returns the OffsetAngle field value
func (o *EllipsoidArcAllOf) GetOffsetAngle() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OffsetAngle
}

// GetOffsetAngleOk returns a tuple with the OffsetAngle field value
// and a boolean to check if the value has been set.
func (o *EllipsoidArcAllOf) GetOffsetAngleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OffsetAngle, true
}

// SetOffsetAngle sets field value
func (o *EllipsoidArcAllOf) SetOffsetAngle(v int32) {
	o.OffsetAngle = v
}

// GetIncludedAngle returns the IncludedAngle field value
func (o *EllipsoidArcAllOf) GetIncludedAngle() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.IncludedAngle
}

// GetIncludedAngleOk returns a tuple with the IncludedAngle field value
// and a boolean to check if the value has been set.
func (o *EllipsoidArcAllOf) GetIncludedAngleOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IncludedAngle, true
}

// SetIncludedAngle sets field value
func (o *EllipsoidArcAllOf) SetIncludedAngle(v int32) {
	o.IncludedAngle = v
}

// GetConfidence returns the Confidence field value
func (o *EllipsoidArcAllOf) GetConfidence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value
// and a boolean to check if the value has been set.
func (o *EllipsoidArcAllOf) GetConfidenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Confidence, true
}

// SetConfidence sets field value
func (o *EllipsoidArcAllOf) SetConfidence(v int32) {
	o.Confidence = v
}

func (o EllipsoidArcAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EllipsoidArcAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["point"] = o.Point
	toSerialize["innerRadius"] = o.InnerRadius
	toSerialize["uncertaintyRadius"] = o.UncertaintyRadius
	toSerialize["offsetAngle"] = o.OffsetAngle
	toSerialize["includedAngle"] = o.IncludedAngle
	toSerialize["confidence"] = o.Confidence
	return toSerialize, nil
}

type NullableEllipsoidArcAllOf struct {
	value *EllipsoidArcAllOf
	isSet bool
}

func (v NullableEllipsoidArcAllOf) Get() *EllipsoidArcAllOf {
	return v.value
}

func (v *NullableEllipsoidArcAllOf) Set(val *EllipsoidArcAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEllipsoidArcAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEllipsoidArcAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEllipsoidArcAllOf(val *EllipsoidArcAllOf) *NullableEllipsoidArcAllOf {
	return &NullableEllipsoidArcAllOf{value: val, isSet: true}
}

func (v NullableEllipsoidArcAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEllipsoidArcAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


