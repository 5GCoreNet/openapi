/*
3gpp-ecs-address-provision

API for ECS Address Provisioning.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EcsAddressProvision

import (
	"encoding/json"
)

// checks if the Local3dPointUncertaintyEllipsoidAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Local3dPointUncertaintyEllipsoidAllOf{}

// Local3dPointUncertaintyEllipsoidAllOf struct for Local3dPointUncertaintyEllipsoidAllOf
type Local3dPointUncertaintyEllipsoidAllOf struct {
	LocalOrigin LocalOrigin `json:"localOrigin"`
	Point RelativeCartesianLocation `json:"point"`
	UncertaintyEllipsoid UncertaintyEllipsoid `json:"uncertaintyEllipsoid"`
	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}

// NewLocal3dPointUncertaintyEllipsoidAllOf instantiates a new Local3dPointUncertaintyEllipsoidAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocal3dPointUncertaintyEllipsoidAllOf(localOrigin LocalOrigin, point RelativeCartesianLocation, uncertaintyEllipsoid UncertaintyEllipsoid, confidence int32) *Local3dPointUncertaintyEllipsoidAllOf {
	this := Local3dPointUncertaintyEllipsoidAllOf{}
	this.LocalOrigin = localOrigin
	this.Point = point
	this.UncertaintyEllipsoid = uncertaintyEllipsoid
	this.Confidence = confidence
	return &this
}

// NewLocal3dPointUncertaintyEllipsoidAllOfWithDefaults instantiates a new Local3dPointUncertaintyEllipsoidAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocal3dPointUncertaintyEllipsoidAllOfWithDefaults() *Local3dPointUncertaintyEllipsoidAllOf {
	this := Local3dPointUncertaintyEllipsoidAllOf{}
	return &this
}

// GetLocalOrigin returns the LocalOrigin field value
func (o *Local3dPointUncertaintyEllipsoidAllOf) GetLocalOrigin() LocalOrigin {
	if o == nil {
		var ret LocalOrigin
		return ret
	}

	return o.LocalOrigin
}

// GetLocalOriginOk returns a tuple with the LocalOrigin field value
// and a boolean to check if the value has been set.
func (o *Local3dPointUncertaintyEllipsoidAllOf) GetLocalOriginOk() (*LocalOrigin, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocalOrigin, true
}

// SetLocalOrigin sets field value
func (o *Local3dPointUncertaintyEllipsoidAllOf) SetLocalOrigin(v LocalOrigin) {
	o.LocalOrigin = v
}

// GetPoint returns the Point field value
func (o *Local3dPointUncertaintyEllipsoidAllOf) GetPoint() RelativeCartesianLocation {
	if o == nil {
		var ret RelativeCartesianLocation
		return ret
	}

	return o.Point
}

// GetPointOk returns a tuple with the Point field value
// and a boolean to check if the value has been set.
func (o *Local3dPointUncertaintyEllipsoidAllOf) GetPointOk() (*RelativeCartesianLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Point, true
}

// SetPoint sets field value
func (o *Local3dPointUncertaintyEllipsoidAllOf) SetPoint(v RelativeCartesianLocation) {
	o.Point = v
}

// GetUncertaintyEllipsoid returns the UncertaintyEllipsoid field value
func (o *Local3dPointUncertaintyEllipsoidAllOf) GetUncertaintyEllipsoid() UncertaintyEllipsoid {
	if o == nil {
		var ret UncertaintyEllipsoid
		return ret
	}

	return o.UncertaintyEllipsoid
}

// GetUncertaintyEllipsoidOk returns a tuple with the UncertaintyEllipsoid field value
// and a boolean to check if the value has been set.
func (o *Local3dPointUncertaintyEllipsoidAllOf) GetUncertaintyEllipsoidOk() (*UncertaintyEllipsoid, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UncertaintyEllipsoid, true
}

// SetUncertaintyEllipsoid sets field value
func (o *Local3dPointUncertaintyEllipsoidAllOf) SetUncertaintyEllipsoid(v UncertaintyEllipsoid) {
	o.UncertaintyEllipsoid = v
}

// GetConfidence returns the Confidence field value
func (o *Local3dPointUncertaintyEllipsoidAllOf) GetConfidence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value
// and a boolean to check if the value has been set.
func (o *Local3dPointUncertaintyEllipsoidAllOf) GetConfidenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Confidence, true
}

// SetConfidence sets field value
func (o *Local3dPointUncertaintyEllipsoidAllOf) SetConfidence(v int32) {
	o.Confidence = v
}

func (o Local3dPointUncertaintyEllipsoidAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Local3dPointUncertaintyEllipsoidAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["localOrigin"] = o.LocalOrigin
	toSerialize["point"] = o.Point
	toSerialize["uncertaintyEllipsoid"] = o.UncertaintyEllipsoid
	toSerialize["confidence"] = o.Confidence
	return toSerialize, nil
}

type NullableLocal3dPointUncertaintyEllipsoidAllOf struct {
	value *Local3dPointUncertaintyEllipsoidAllOf
	isSet bool
}

func (v NullableLocal3dPointUncertaintyEllipsoidAllOf) Get() *Local3dPointUncertaintyEllipsoidAllOf {
	return v.value
}

func (v *NullableLocal3dPointUncertaintyEllipsoidAllOf) Set(val *Local3dPointUncertaintyEllipsoidAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLocal3dPointUncertaintyEllipsoidAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLocal3dPointUncertaintyEllipsoidAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocal3dPointUncertaintyEllipsoidAllOf(val *Local3dPointUncertaintyEllipsoidAllOf) *NullableLocal3dPointUncertaintyEllipsoidAllOf {
	return &NullableLocal3dPointUncertaintyEllipsoidAllOf{value: val, isSet: true}
}

func (v NullableLocal3dPointUncertaintyEllipsoidAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocal3dPointUncertaintyEllipsoidAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


