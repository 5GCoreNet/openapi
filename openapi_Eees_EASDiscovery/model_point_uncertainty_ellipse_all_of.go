/*
Eees_EASDiscovery

API for EAS Discovery. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_EASDiscovery

import (
	"encoding/json"
)

// checks if the PointUncertaintyEllipseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PointUncertaintyEllipseAllOf{}

// PointUncertaintyEllipseAllOf struct for PointUncertaintyEllipseAllOf
type PointUncertaintyEllipseAllOf struct {
	Point              GeographicalCoordinates `json:"point"`
	UncertaintyEllipse UncertaintyEllipse      `json:"uncertaintyEllipse"`
	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}

// NewPointUncertaintyEllipseAllOf instantiates a new PointUncertaintyEllipseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPointUncertaintyEllipseAllOf(point GeographicalCoordinates, uncertaintyEllipse UncertaintyEllipse, confidence int32) *PointUncertaintyEllipseAllOf {
	this := PointUncertaintyEllipseAllOf{}
	this.Point = point
	this.UncertaintyEllipse = uncertaintyEllipse
	this.Confidence = confidence
	return &this
}

// NewPointUncertaintyEllipseAllOfWithDefaults instantiates a new PointUncertaintyEllipseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPointUncertaintyEllipseAllOfWithDefaults() *PointUncertaintyEllipseAllOf {
	this := PointUncertaintyEllipseAllOf{}
	return &this
}

// GetPoint returns the Point field value
func (o *PointUncertaintyEllipseAllOf) GetPoint() GeographicalCoordinates {
	if o == nil {
		var ret GeographicalCoordinates
		return ret
	}

	return o.Point
}

// GetPointOk returns a tuple with the Point field value
// and a boolean to check if the value has been set.
func (o *PointUncertaintyEllipseAllOf) GetPointOk() (*GeographicalCoordinates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Point, true
}

// SetPoint sets field value
func (o *PointUncertaintyEllipseAllOf) SetPoint(v GeographicalCoordinates) {
	o.Point = v
}

// GetUncertaintyEllipse returns the UncertaintyEllipse field value
func (o *PointUncertaintyEllipseAllOf) GetUncertaintyEllipse() UncertaintyEllipse {
	if o == nil {
		var ret UncertaintyEllipse
		return ret
	}

	return o.UncertaintyEllipse
}

// GetUncertaintyEllipseOk returns a tuple with the UncertaintyEllipse field value
// and a boolean to check if the value has been set.
func (o *PointUncertaintyEllipseAllOf) GetUncertaintyEllipseOk() (*UncertaintyEllipse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UncertaintyEllipse, true
}

// SetUncertaintyEllipse sets field value
func (o *PointUncertaintyEllipseAllOf) SetUncertaintyEllipse(v UncertaintyEllipse) {
	o.UncertaintyEllipse = v
}

// GetConfidence returns the Confidence field value
func (o *PointUncertaintyEllipseAllOf) GetConfidence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value
// and a boolean to check if the value has been set.
func (o *PointUncertaintyEllipseAllOf) GetConfidenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Confidence, true
}

// SetConfidence sets field value
func (o *PointUncertaintyEllipseAllOf) SetConfidence(v int32) {
	o.Confidence = v
}

func (o PointUncertaintyEllipseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PointUncertaintyEllipseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["point"] = o.Point
	toSerialize["uncertaintyEllipse"] = o.UncertaintyEllipse
	toSerialize["confidence"] = o.Confidence
	return toSerialize, nil
}

type NullablePointUncertaintyEllipseAllOf struct {
	value *PointUncertaintyEllipseAllOf
	isSet bool
}

func (v NullablePointUncertaintyEllipseAllOf) Get() *PointUncertaintyEllipseAllOf {
	return v.value
}

func (v *NullablePointUncertaintyEllipseAllOf) Set(val *PointUncertaintyEllipseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePointUncertaintyEllipseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePointUncertaintyEllipseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePointUncertaintyEllipseAllOf(val *PointUncertaintyEllipseAllOf) *NullablePointUncertaintyEllipseAllOf {
	return &NullablePointUncertaintyEllipseAllOf{value: val, isSet: true}
}

func (v NullablePointUncertaintyEllipseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePointUncertaintyEllipseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
