/*
CAPIF_Discover_Service_API

API for discovering service APIs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_Discover_Service_API

import (
	"encoding/json"
)

// checks if the PointUncertaintyCircleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PointUncertaintyCircleAllOf{}

// PointUncertaintyCircleAllOf struct for PointUncertaintyCircleAllOf
type PointUncertaintyCircleAllOf struct {
	Point GeographicalCoordinates `json:"point"`
	// Indicates value of uncertainty.
	Uncertainty float32 `json:"uncertainty"`
}

// NewPointUncertaintyCircleAllOf instantiates a new PointUncertaintyCircleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPointUncertaintyCircleAllOf(point GeographicalCoordinates, uncertainty float32) *PointUncertaintyCircleAllOf {
	this := PointUncertaintyCircleAllOf{}
	this.Point = point
	this.Uncertainty = uncertainty
	return &this
}

// NewPointUncertaintyCircleAllOfWithDefaults instantiates a new PointUncertaintyCircleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPointUncertaintyCircleAllOfWithDefaults() *PointUncertaintyCircleAllOf {
	this := PointUncertaintyCircleAllOf{}
	return &this
}

// GetPoint returns the Point field value
func (o *PointUncertaintyCircleAllOf) GetPoint() GeographicalCoordinates {
	if o == nil {
		var ret GeographicalCoordinates
		return ret
	}

	return o.Point
}

// GetPointOk returns a tuple with the Point field value
// and a boolean to check if the value has been set.
func (o *PointUncertaintyCircleAllOf) GetPointOk() (*GeographicalCoordinates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Point, true
}

// SetPoint sets field value
func (o *PointUncertaintyCircleAllOf) SetPoint(v GeographicalCoordinates) {
	o.Point = v
}

// GetUncertainty returns the Uncertainty field value
func (o *PointUncertaintyCircleAllOf) GetUncertainty() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Uncertainty
}

// GetUncertaintyOk returns a tuple with the Uncertainty field value
// and a boolean to check if the value has been set.
func (o *PointUncertaintyCircleAllOf) GetUncertaintyOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uncertainty, true
}

// SetUncertainty sets field value
func (o *PointUncertaintyCircleAllOf) SetUncertainty(v float32) {
	o.Uncertainty = v
}

func (o PointUncertaintyCircleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PointUncertaintyCircleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["point"] = o.Point
	toSerialize["uncertainty"] = o.Uncertainty
	return toSerialize, nil
}

type NullablePointUncertaintyCircleAllOf struct {
	value *PointUncertaintyCircleAllOf
	isSet bool
}

func (v NullablePointUncertaintyCircleAllOf) Get() *PointUncertaintyCircleAllOf {
	return v.value
}

func (v *NullablePointUncertaintyCircleAllOf) Set(val *PointUncertaintyCircleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePointUncertaintyCircleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePointUncertaintyCircleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePointUncertaintyCircleAllOf(val *PointUncertaintyCircleAllOf) *NullablePointUncertaintyCircleAllOf {
	return &NullablePointUncertaintyCircleAllOf{value: val, isSet: true}
}

func (v NullablePointUncertaintyCircleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePointUncertaintyCircleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


