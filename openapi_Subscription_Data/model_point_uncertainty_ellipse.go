/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
)

// checks if the PointUncertaintyEllipse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PointUncertaintyEllipse{}

// PointUncertaintyEllipse Ellipsoid point with uncertainty ellipse.
type PointUncertaintyEllipse struct {
	GADShape
	Point GeographicalCoordinates `json:"point"`
	UncertaintyEllipse UncertaintyEllipse `json:"uncertaintyEllipse"`
	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}

// NewPointUncertaintyEllipse instantiates a new PointUncertaintyEllipse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPointUncertaintyEllipse(point GeographicalCoordinates, uncertaintyEllipse UncertaintyEllipse, confidence int32, shape SupportedGADShapes) *PointUncertaintyEllipse {
	this := PointUncertaintyEllipse{}
	this.Shape = shape
	this.Point = point
	this.UncertaintyEllipse = uncertaintyEllipse
	this.Confidence = confidence
	return &this
}

// NewPointUncertaintyEllipseWithDefaults instantiates a new PointUncertaintyEllipse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPointUncertaintyEllipseWithDefaults() *PointUncertaintyEllipse {
	this := PointUncertaintyEllipse{}
	return &this
}

// GetPoint returns the Point field value
func (o *PointUncertaintyEllipse) GetPoint() GeographicalCoordinates {
	if o == nil {
		var ret GeographicalCoordinates
		return ret
	}

	return o.Point
}

// GetPointOk returns a tuple with the Point field value
// and a boolean to check if the value has been set.
func (o *PointUncertaintyEllipse) GetPointOk() (*GeographicalCoordinates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Point, true
}

// SetPoint sets field value
func (o *PointUncertaintyEllipse) SetPoint(v GeographicalCoordinates) {
	o.Point = v
}

// GetUncertaintyEllipse returns the UncertaintyEllipse field value
func (o *PointUncertaintyEllipse) GetUncertaintyEllipse() UncertaintyEllipse {
	if o == nil {
		var ret UncertaintyEllipse
		return ret
	}

	return o.UncertaintyEllipse
}

// GetUncertaintyEllipseOk returns a tuple with the UncertaintyEllipse field value
// and a boolean to check if the value has been set.
func (o *PointUncertaintyEllipse) GetUncertaintyEllipseOk() (*UncertaintyEllipse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UncertaintyEllipse, true
}

// SetUncertaintyEllipse sets field value
func (o *PointUncertaintyEllipse) SetUncertaintyEllipse(v UncertaintyEllipse) {
	o.UncertaintyEllipse = v
}

// GetConfidence returns the Confidence field value
func (o *PointUncertaintyEllipse) GetConfidence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value
// and a boolean to check if the value has been set.
func (o *PointUncertaintyEllipse) GetConfidenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Confidence, true
}

// SetConfidence sets field value
func (o *PointUncertaintyEllipse) SetConfidence(v int32) {
	o.Confidence = v
}

func (o PointUncertaintyEllipse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PointUncertaintyEllipse) ToMap() (map[string]interface{}, error) {
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
	toSerialize["uncertaintyEllipse"] = o.UncertaintyEllipse
	toSerialize["confidence"] = o.Confidence
	return toSerialize, nil
}

type NullablePointUncertaintyEllipse struct {
	value *PointUncertaintyEllipse
	isSet bool
}

func (v NullablePointUncertaintyEllipse) Get() *PointUncertaintyEllipse {
	return v.value
}

func (v *NullablePointUncertaintyEllipse) Set(val *PointUncertaintyEllipse) {
	v.value = val
	v.isSet = true
}

func (v NullablePointUncertaintyEllipse) IsSet() bool {
	return v.isSet
}

func (v *NullablePointUncertaintyEllipse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePointUncertaintyEllipse(val *PointUncertaintyEllipse) *NullablePointUncertaintyEllipse {
	return &NullablePointUncertaintyEllipse{value: val, isSet: true}
}

func (v NullablePointUncertaintyEllipse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePointUncertaintyEllipse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


